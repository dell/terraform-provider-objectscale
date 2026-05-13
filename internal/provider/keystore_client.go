/*
Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"terraform-provider-objectscale/internal/client"
	"terraform-provider-objectscale/internal/models"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const (
	vdcKeystorePath        = "/vdc/keystore"
	objectCertKeystorePath = "/object-cert/keystore"
	maxRetries             = 3
)

// doKeystoreRequest executes an HTTP request against the ObjectScale keystore API with auth headers.
func doKeystoreRequest(ctx context.Context, c *client.Client, method, path string, body []byte) ([]byte, int, error) {
	cfg := c.GenClient.GetConfig()
	url := cfg.Servers[0].URL + path
	var reqBody io.Reader
	if body != nil {
		reqBody = bytes.NewReader(body)
	}
	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, 0, fmt.Errorf("error creating request: %w", err)
	}
	for k, v := range cfg.DefaultHeader {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := cfg.HTTPClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("error executing request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("error reading response: %w", err)
	}

	return respBody, resp.StatusCode, nil
}

// parseKeystoreError checks if a response body contains an ObjectScale error code.
func parseKeystoreError(body []byte) *models.KeystoreErrorResponse {
	var errResp models.KeystoreErrorResponse
	if err := json.Unmarshal(body, &errResp); err != nil {
		return nil
	}
	if errResp.Code != 0 {
		return &errResp
	}
	return nil
}

// mapKeystoreError maps an ObjectScale API error code to a human-readable message.
func mapKeystoreError(errResp *models.KeystoreErrorResponse) (string, string) {
	switch errResp.Code {
	case 999:
		return "Insufficient Permissions", fmt.Sprintf("SECURITY_ADMIN role required. API code: %d. Details: %s", errResp.Code, errResp.Details)
	case 1005:
		return "Missing Required Field", fmt.Sprintf("Both private_key and certificate_chain are required. API code: %d. Details: %s", errResp.Code, errResp.Details)
	case 1008:
		return "Invalid Certificate or Key Format", fmt.Sprintf("Ensure valid PEM with PKCS#1 or PKCS#8 RSA key. API code: %d. Details: %s", errResp.Code, errResp.Details)
	case 1013:
		return "Certificate Already Deployed", fmt.Sprintf("The certificate chain is already active. No changes made. API code: %d. Details: %s", errResp.Code, errResp.Details)
	default:
		return "ObjectScale API Error", fmt.Sprintf("Unexpected error code: %d. Description: %s. Details: %s", errResp.Code, errResp.Description, errResp.Details)
	}
}

// GetVDCKeystore reads the current VDC certificate chain via GET /vdc/keystore.
var GetVDCKeystore = func(ctx context.Context, c *client.Client) (string, error) {
	tflog.Debug(ctx, "reading VDC keystore certificate")
	body, statusCode, err := doKeystoreRequest(ctx, c, http.MethodGet, vdcKeystorePath, nil)
	if err != nil {
		return "", fmt.Errorf("error reading VDC keystore: %w", err)
	}
	return parseGetResponse(body, statusCode, "VDC keystore")
}

// GetObjectCertKeystore reads the current Object certificate chain via GET /object-cert/keystore.
var GetObjectCertKeystore = func(ctx context.Context, c *client.Client) (string, error) {
	tflog.Debug(ctx, "reading Object certificate keystore")
	body, statusCode, err := doKeystoreRequest(ctx, c, http.MethodGet, objectCertKeystorePath, nil)
	if err != nil {
		return "", fmt.Errorf("error reading Object certificate keystore: %w", err)
	}
	return parseGetResponse(body, statusCode, "Object certificate keystore")
}

// parseGetResponse parses a GET keystore response, handling error codes.
func parseGetResponse(body []byte, statusCode int, context string) (string, error) {
	if statusCode == http.StatusUnauthorized {
		return "", fmt.Errorf("%s: authentication failed (HTTP 401). Check credentials", context)
	}

	// Check for error body (even on HTTP 200)
	if errResp := parseKeystoreError(body); errResp != nil {
		summary, detail := mapKeystoreError(errResp)
		return "", fmt.Errorf("%s: %s - %s", context, summary, detail)
	}

	var getResp models.KeystoreGetResponse
	if err := json.Unmarshal(body, &getResp); err != nil {
		return "", fmt.Errorf("%s: error parsing response: %w", context, err)
	}
	if getResp.Chain == "" {
		return "", fmt.Errorf("%s: empty certificate chain in response", context)
	}

	return getResp.Chain, nil
}

// PutVDCKeystore replaces the VDC certificate via PUT /vdc/keystore with retry support.
var PutVDCKeystore = func(ctx context.Context, c *client.Client, privateKey, certChain string) error {
	tflog.Debug(ctx, "updating VDC keystore certificate", map[string]interface{}{"has_private_key": true})
	payload := models.KeystorePutRequest{
		KeyAndCertificate: &models.KeyAndCertificate{
			PrivateKey:       privateKey,
			CertificateChain: certChain,
		},
	}
	return executePutWithRetry(ctx, c, vdcKeystorePath, payload, "VDC keystore")
}

// PutObjectCertKeystore replaces the Object certificate via PUT /object-cert/keystore with retry support.
var PutObjectCertKeystore = func(ctx context.Context, c *client.Client, privateKey, certChain string) error {
	tflog.Debug(ctx, "updating Object certificate keystore", map[string]interface{}{"has_private_key": true})
	payload := models.KeystorePutRequest{
		KeyAndCertificate: &models.KeyAndCertificate{
			PrivateKey:       privateKey,
			CertificateChain: certChain,
		},
	}
	return executePutWithRetry(ctx, c, objectCertKeystorePath, payload, "Object certificate keystore")
}

// PutObjectCertSelfSigned generates a self-signed Object certificate via PUT /object-cert/keystore.
var PutObjectCertSelfSigned = func(ctx context.Context, c *client.Client, ipAddresses []string) (string, error) {
	tflog.Debug(ctx, "generating self-signed Object certificate", map[string]interface{}{"ip_count": len(ipAddresses)})
	selfsigned := true
	payload := models.KeystorePutRequest{
		SystemSelfsigned: &selfsigned,
		IPAddresses:      ipAddresses,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error marshaling self-signed request: %w", err)
	}

	respBody, statusCode, err := doKeystoreRequest(ctx, c, http.MethodPut, objectCertKeystorePath, body)
	if err != nil {
		return "", fmt.Errorf("error sending self-signed request: %w", err)
	}

	if statusCode == http.StatusUnauthorized {
		return "", fmt.Errorf("object certificate keystore: authentication failed (HTTP 401). Check credentials")
	}

	if errResp := parseKeystoreError(respBody); errResp != nil {
		summary, detail := mapKeystoreError(errResp)
		return "", fmt.Errorf("object certificate keystore: %s - %s", summary, detail)
	}

	if statusCode >= 500 {
		return "", fmt.Errorf("object certificate keystore: server error (HTTP %d)", statusCode)
	}

	var getResp models.KeystoreGetResponse
	if err := json.Unmarshal(respBody, &getResp); err != nil {
		// Self-signed PUT may not return chain in some versions; read it back
		chain, getErr := GetObjectCertKeystore(ctx, c)
		if getErr != nil {
			return "", fmt.Errorf("self-signed certificate generated but failed to read back chain: %w", getErr)
		}
		return chain, nil
	}
	if getResp.Chain != "" {
		return getResp.Chain, nil
	}
	// Read back if no chain in response
	chain, getErr := GetObjectCertKeystore(ctx, c)
	if getErr != nil {
		return "", fmt.Errorf("self-signed certificate generated but failed to read back chain: %w", getErr)
	}
	return chain, nil
}

// executePutWithRetry executes a PUT request with exponential backoff retry for HTTP 5xx errors.
func executePutWithRetry(ctx context.Context, c *client.Client, path string, payload models.KeystorePutRequest, context string) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling request: %w", err)
	}

	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		respBody, statusCode, err := doKeystoreRequest(ctx, c, http.MethodPut, path, body)
		if err != nil {
			lastErr = err
			if attempt < maxRetries-1 {
				sleepDuration := time.Duration(1<<uint(attempt)) * time.Second
				tflog.Warn(ctx, fmt.Sprintf("%s: request failed, retrying in %v (attempt %d/%d)", context, sleepDuration, attempt+1, maxRetries))
				time.Sleep(sleepDuration)
				continue
			}
			return fmt.Errorf("%s: %w", context, lastErr)
		}

		if statusCode == http.StatusUnauthorized {
			return fmt.Errorf("%s: authentication failed (HTTP 401). Check credentials", context)
		}

		if statusCode >= 500 {
			lastErr = fmt.Errorf("%s: server error (HTTP %d)", context, statusCode)
			if attempt < maxRetries-1 {
				sleepDuration := time.Duration(1<<uint(attempt)) * time.Second
				tflog.Warn(ctx, fmt.Sprintf("%s: server error, retrying in %v (attempt %d/%d)", context, sleepDuration, attempt+1, maxRetries))
				time.Sleep(sleepDuration)
				continue
			}
			return lastErr
		}

		// Check for error body (even on HTTP 200)
		if errResp := parseKeystoreError(respBody); errResp != nil {
			summary, detail := mapKeystoreError(errResp)
			return fmt.Errorf("%s: %s - %s", context, summary, detail)
		}

		return nil
	}
	return fmt.Errorf("%s: max retries exceeded: %w", context, lastErr)
}
