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

// MANUAL — hand-written wrapper for the ObjectScale SAML Service Provider
// (`/ecs-service-provider`) operations. Excluded from `make build_client`
// regeneration via the `_manual.go` suffix.

package clientgen

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// ServiceProvider is the JSON shape of `service_provider` returned by
// `GET /ecs-service-provider`.
type ServiceProvider struct {
	DNS          string `json:"dns,omitempty"`
	UUID         string `json:"uuid,omitempty"`
	UniqueId     string `json:"unique_id,omitempty"`
	Etag         string `json:"etag,omitempty"`
	JavaKeystore string `json:"java_keystore,omitempty"`
	KeyAlias     string `json:"key_alias,omitempty"`
	KeyPassword  string `json:"key_password,omitempty"`
	CreateTime   string `json:"create_time,omitempty"`
	LastModified string `json:"last_modified,omitempty"`
}

// ServiceProviderEnvelope is the wire envelope for create/get/update.
type ServiceProviderEnvelope struct {
	ServiceProviderCreate ServiceProvider `json:"service_provider_create,omitempty"`
	ServiceProvider       ServiceProvider `json:"service_provider,omitempty"`
}

// ServiceProviderHTTPError is returned for any non-2xx response.
type ServiceProviderHTTPError struct {
	Status int
	Body   []byte
	Action string
}

func (e *ServiceProviderHTTPError) Error() string {
	return fmt.Sprintf("ServiceProvider %s: HTTP %d: %s", e.Action, e.Status, string(e.Body))
}

// HTTPStatus implements helper.HTTPStatusError.
func (e *ServiceProviderHTTPError) HTTPStatus() int { return e.Status }

// CreateServiceProvider issues `POST /ecs-service-provider`.
func (c *APIClient) CreateServiceProvider(ctx context.Context, sp ServiceProvider) (*ServiceProvider, *http.Response, error) {
	body, err := json.Marshal(ServiceProviderEnvelope{ServiceProvider: sp})
	if err != nil {
		return nil, nil, fmt.Errorf("marshal CreateServiceProvider: %w", err)
	}
	req, err := c.buildSPRequest(ctx, http.MethodPost, "/ecs-service-provider", body)
	if err != nil {
		return nil, nil, err
	}
	var env ServiceProviderEnvelope
	resp, err := c.executeSPRequest(req, "Create", &env)
	if err != nil {
		return nil, resp, err
	}
	return c.pickSP(env), resp, nil
}

// GetServiceProvider issues `GET /ecs-service-provider`.
func (c *APIClient) GetServiceProvider(ctx context.Context) (*ServiceProvider, *http.Response, error) {
	req, err := c.buildSPRequest(ctx, http.MethodGet, "/ecs-service-provider", nil)
	if err != nil {
		return nil, nil, err
	}
	var env ServiceProviderEnvelope
	resp, err := c.executeSPRequest(req, "Get", &env)
	if err != nil {
		return nil, resp, err
	}
	return c.pickSP(env), resp, nil
}

// UpdateServiceProvider issues `PUT /ecs-service-provider`.
func (c *APIClient) UpdateServiceProvider(ctx context.Context, sp ServiceProvider) (*ServiceProvider, *http.Response, error) {
	body, err := json.Marshal(ServiceProviderEnvelope{ServiceProvider: sp})
	if err != nil {
		return nil, nil, fmt.Errorf("marshal UpdateServiceProvider: %w", err)
	}
	req, err := c.buildSPRequest(ctx, http.MethodPut, "/ecs-service-provider", body)
	if err != nil {
		return nil, nil, err
	}
	var env ServiceProviderEnvelope
	resp, err := c.executeSPRequest(req, "Update", &env)
	if err != nil {
		return nil, resp, err
	}
	return c.pickSP(env), resp, nil
}

// DeleteServiceProvider issues `DELETE /ecs-service-provider`.
func (c *APIClient) DeleteServiceProvider(ctx context.Context) (*http.Response, error) {
	req, err := c.buildSPRequest(ctx, http.MethodDelete, "/ecs-service-provider", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.executeSPRequest(req, "Delete", nil)
	return resp, err
}

// GetServiceProviderMetadata issues `GET /ecs-service-provider/metadata`
// and returns the raw XML body.
func (c *APIClient) GetServiceProviderMetadata(ctx context.Context) (string, *http.Response, error) {
	req, err := c.buildSPRequest(ctx, http.MethodGet, "/ecs-service-provider/metadata", nil)
	if err != nil {
		return "", nil, err
	}
	// metadata is XML — request explicit XML accept
	req.Header.Set("Accept", "application/xml")
	for k, v := range c.cfg.DefaultHeader {
		if req.Header.Get(k) == "" {
			req.Header.Set(k, v)
		}
	}
	if c.cfg.UserAgent != "" && req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", c.cfg.UserAgent)
	}
	resp, err := c.cfg.HTTPClient.Do(req)
	if err != nil {
		return "", nil, fmt.Errorf("execute GetServiceProviderMetadata: %w", err)
	}
	body, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return "", resp, fmt.Errorf("read metadata: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", resp, &ServiceProviderHTTPError{Status: resp.StatusCode, Body: body, Action: "GetMetadata"}
	}
	return string(body), resp, nil
}

func (c *APIClient) buildSPRequest(ctx context.Context, method, path string, body []byte) (*http.Request, error) {
	server, err := c.cfg.ServerURLWithContext(ctx, "")
	if err != nil {
		return nil, fmt.Errorf("resolve server url: %w", err)
	}
	var rdr io.Reader
	if body != nil {
		rdr = bytes.NewReader(body)
	}
	req, err := http.NewRequestWithContext(ctx, method, strings.TrimRight(server, "/")+path, rdr)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func (c *APIClient) executeSPRequest(req *http.Request, action string, out interface{}) (*http.Response, error) {
	for k, v := range c.cfg.DefaultHeader {
		if req.Header.Get(k) == "" {
			req.Header.Set(k, v)
		}
	}
	if c.cfg.UserAgent != "" && req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", c.cfg.UserAgent)
	}
	resp, err := c.cfg.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("execute ServiceProvider %s: %w", action, err)
	}
	body, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return resp, fmt.Errorf("read ServiceProvider %s response: %w", action, err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return resp, &ServiceProviderHTTPError{Status: resp.StatusCode, Body: body, Action: action}
	}
	if len(body) == 0 || out == nil {
		return resp, nil
	}
	ct := resp.Header.Get("Content-Type")
	if xmlCheck.MatchString(ct) {
		if err := xml.Unmarshal(body, out); err != nil {
			return resp, fmt.Errorf("decode XML ServiceProvider %s: %w", action, err)
		}
	} else {
		if err := json.Unmarshal(body, out); err != nil {
			return resp, fmt.Errorf("decode JSON ServiceProvider %s: %w", action, err)
		}
	}
	return resp, nil
}

// pickSP returns whichever of the envelope's two views is populated.
func (c *APIClient) pickSP(env ServiceProviderEnvelope) *ServiceProvider {
	if env.ServiceProvider.DNS != "" || env.ServiceProvider.UUID != "" || env.ServiceProvider.UniqueId != "" {
		sp := env.ServiceProvider
		return &sp
	}
	if env.ServiceProviderCreate.DNS != "" || env.ServiceProviderCreate.UUID != "" {
		sp := env.ServiceProviderCreate
		return &sp
	}
	// neither populated — return zero value to allow caller to detect via 200+empty
	sp := env.ServiceProvider
	return &sp
}
