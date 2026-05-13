/*
Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

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
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"terraform-provider-objectscale/internal/client"
	"terraform-provider-objectscale/internal/models"
	"testing"
)

// newTestClient creates a client.Client pointing to a test HTTP server.
func newTestClient(server *httptest.Server) *client.Client {
	return &client.Client{
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
		AuthHeaders: map[string]string{
			"X-SDS-AUTH-TOKEN": "test-token",
		},
	}
}

func TestGetVDCKeystore_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/vdc/keystore" || r.Method != http.MethodGet {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		resp := models.KeystoreGetResponse{Chain: "-----BEGIN CERTIFICATE-----\ntest\n-----END CERTIFICATE-----"}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	defer server.Close()

	c := newTestClient(server)
	chain, err := GetVDCKeystore(context.Background(), c)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain == "" {
		t.Error("expected non-empty chain")
	}
}

func TestGetVDCKeystore_Error999(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := models.KeystoreErrorResponse{Code: 999, Description: "Insufficient permissions", Details: "No SECURITY_ADMIN role"}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	defer server.Close()

	c := newTestClient(server)
	_, err := GetVDCKeystore(context.Background(), c)
	if err == nil {
		t.Fatal("expected error for code 999")
	}
	if !contains(err.Error(), "Insufficient Permissions") {
		t.Errorf("expected permission error, got: %v", err)
	}
}

func TestPutVDCKeystore_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/vdc/keystore" || r.Method != http.MethodPut {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("{}")); err != nil {
			t.Errorf("failed to write response: %v", err)
		}
	}))
	defer server.Close()

	c := newTestClient(server)
	err := PutVDCKeystore(context.Background(), c, "key", "chain")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPutVDCKeystore_Error1013(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		resp := models.KeystoreErrorResponse{Code: 1013, Description: "Duplicate cert", Details: "Already deployed"}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	defer server.Close()

	c := newTestClient(server)
	err := PutVDCKeystore(context.Background(), c, "key", "chain")
	if err == nil {
		t.Fatal("expected error for code 1013")
	}
	if !contains(err.Error(), "Certificate Already Deployed") {
		t.Errorf("expected duplicate cert error, got: %v", err)
	}
}

func TestGetObjectCertKeystore_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/object-cert/keystore" || r.Method != http.MethodGet {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		resp := models.KeystoreGetResponse{Chain: "-----BEGIN CERTIFICATE-----\nobjtest\n-----END CERTIFICATE-----"}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	defer server.Close()

	c := newTestClient(server)
	chain, err := GetObjectCertKeystore(context.Background(), c)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain == "" {
		t.Error("expected non-empty chain")
	}
}

func TestPutObjectCertKeystore_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/object-cert/keystore" || r.Method != http.MethodPut {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("{}")); err != nil {
			t.Errorf("failed to write response: %v", err)
		}
	}))
	defer server.Close()

	c := newTestClient(server)
	err := PutObjectCertKeystore(context.Background(), c, "key", "chain")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPutObjectCertSelfSigned_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/object-cert/keystore" || r.Method != http.MethodPut {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		resp := models.KeystoreGetResponse{Chain: "-----BEGIN CERTIFICATE-----\nselfsigned\n-----END CERTIFICATE-----"}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	defer server.Close()

	c := newTestClient(server)
	chain, err := PutObjectCertSelfSigned(context.Background(), c, []string{"10.0.0.1"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain == "" {
		t.Error("expected non-empty chain")
	}
}

func TestPutVDCKeystore_Error1008(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := models.KeystoreErrorResponse{Code: 1008, Description: "Invalid format", Details: "Bad PEM"}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	defer server.Close()

	c := newTestClient(server)
	err := PutVDCKeystore(context.Background(), c, "key", "chain")
	if err == nil {
		t.Fatal("expected error for code 1008")
	}
	if !contains(err.Error(), "Invalid Certificate or Key Format") {
		t.Errorf("expected format error, got: %v", err)
	}
}

func TestGetVDCKeystore_Auth401(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	c := newTestClient(server)
	_, err := GetVDCKeystore(context.Background(), c)
	if err == nil {
		t.Fatal("expected error for 401")
	}
	if !contains(err.Error(), "authentication failed") {
		t.Errorf("expected auth error, got: %v", err)
	}
}

func TestPutVDCKeystore_Auth401(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	c := newTestClient(server)
	err := PutVDCKeystore(context.Background(), c, "key", "chain")
	if err == nil {
		t.Fatal("expected error for 401")
	}
	if !contains(err.Error(), "authentication failed") {
		t.Errorf("expected auth error, got: %v", err)
	}
}

func TestGetVDCKeystore_EmptyChain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := models.KeystoreGetResponse{Chain: ""}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	defer server.Close()

	c := newTestClient(server)
	_, err := GetVDCKeystore(context.Background(), c)
	if err == nil {
		t.Fatal("expected error for empty chain")
	}
	if !contains(err.Error(), "empty certificate chain") {
		t.Errorf("expected empty chain error, got: %v", err)
	}
}

func TestPutObjectCertSelfSigned_Auth401(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	c := newTestClient(server)
	_, err := PutObjectCertSelfSigned(context.Background(), c, []string{"10.0.0.1"})
	if err == nil {
		t.Fatal("expected error for 401")
	}
	if !contains(err.Error(), "authentication failed") {
		t.Errorf("expected auth error, got: %v", err)
	}
}

func TestParseKeystoreError_NoError(t *testing.T) {
	body := []byte(`{"chain":"test"}`)
	result := parseKeystoreError(body)
	if result != nil {
		t.Error("expected nil for non-error response")
	}
}

func TestParseKeystoreError_WithError(t *testing.T) {
	body := []byte(`{"code":999,"description":"Permission denied","details":"No admin role"}`)
	result := parseKeystoreError(body)
	if result == nil {
		t.Fatal("expected error response")
	}
	if result.Code != 999 {
		t.Errorf("expected code 999, got %d", result.Code)
	}
}

func TestParseKeystoreError_InvalidJSON(t *testing.T) {
	body := []byte(`not json`)
	result := parseKeystoreError(body)
	if result != nil {
		t.Error("expected nil for invalid JSON")
	}
}

func TestMapKeystoreError_AllCodes(t *testing.T) {
	tests := []struct {
		code            int
		expectedSummary string
	}{
		{999, "Insufficient Permissions"},
		{1005, "Missing Required Field"},
		{1008, "Invalid Certificate or Key Format"},
		{1013, "Certificate Already Deployed"},
		{9999, "ObjectScale API Error"},
	}

	for _, tt := range tests {
		errResp := &models.KeystoreErrorResponse{Code: tt.code, Description: "desc", Details: "detail"}
		summary, _ := mapKeystoreError(errResp)
		if summary != tt.expectedSummary {
			t.Errorf("code %d: expected summary %q, got %q", tt.code, tt.expectedSummary, summary)
		}
	}
}

func TestPutObjectCertSelfSigned_Error999(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := models.KeystoreErrorResponse{Code: 999, Description: "Insufficient permissions", Details: "No role"}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	defer server.Close()

	c := newTestClient(server)
	_, err := PutObjectCertSelfSigned(context.Background(), c, []string{"10.0.0.1"})
	if err == nil {
		t.Fatal("expected error for code 999")
	}
	if !contains(err.Error(), "Insufficient Permissions") {
		t.Errorf("expected permission error, got: %v", err)
	}
}

func TestPutObjectCertSelfSigned_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("internal error")); err != nil {
			t.Errorf("failed to write response: %v", err)
		}
	}))
	defer server.Close()

	c := newTestClient(server)
	_, err := PutObjectCertSelfSigned(context.Background(), c, nil)
	if err == nil {
		t.Fatal("expected error for server error")
	}
	if !contains(err.Error(), "server error") {
		t.Errorf("expected server error, got: %v", err)
	}
}

func TestPutVDCKeystore_RetryOn500(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		if callCount < 3 {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte("server error")); err != nil {
				t.Errorf("failed to write response: %v", err)
			}
			return
		}
		// Third attempt succeeds
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("{}")); err != nil {
			t.Errorf("failed to write response: %v", err)
		}
	}))
	defer server.Close()

	c := newTestClient(server)
	err := PutVDCKeystore(context.Background(), c, "key", "chain")
	if err != nil {
		t.Fatalf("expected success after retries, got: %v", err)
	}
	if callCount != 3 {
		t.Errorf("expected 3 calls (2 retries), got %d", callCount)
	}
}

func TestPutVDCKeystore_RetryExhausted(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("error")); err != nil {
			t.Errorf("failed to write response: %v", err)
		}
	}))
	defer server.Close()

	c := newTestClient(server)
	err := PutVDCKeystore(context.Background(), c, "key", "chain")
	if err == nil {
		t.Fatal("expected error after max retries")
	}
	if !contains(err.Error(), "server error") {
		t.Errorf("expected server error, got: %v", err)
	}
}

func TestGetVDCKeystore_BadJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("not json at all")); err != nil {
			t.Errorf("failed to write response: %v", err)
		}
	}))
	defer server.Close()

	c := newTestClient(server)
	_, err := GetVDCKeystore(context.Background(), c)
	if err == nil {
		t.Fatal("expected error for bad JSON")
	}
	if !contains(err.Error(), "error parsing response") {
		t.Errorf("expected parsing error, got: %v", err)
	}
}

func TestDoKeystoreRequest_ConnectionError(t *testing.T) {
	// Use a client pointing to a non-existent server
	c := &client.Client{
		BaseURL:     "http://127.0.0.1:1",
		HTTPClient:  &http.Client{Timeout: 1},
		AuthHeaders: map[string]string{},
	}
	_, err := GetVDCKeystore(context.Background(), c)
	if err == nil {
		t.Fatal("expected connection error")
	}
}

func TestGetObjectCertKeystore_Auth401(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	c := newTestClient(server)
	_, err := GetObjectCertKeystore(context.Background(), c)
	if err == nil {
		t.Fatal("expected error for 401")
	}
	if !contains(err.Error(), "authentication failed") {
		t.Errorf("expected auth error, got: %v", err)
	}
}

func TestGetObjectCertKeystore_Error999(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := models.KeystoreErrorResponse{Code: 999, Description: "Permission denied", Details: "No role"}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	defer server.Close()

	c := newTestClient(server)
	_, err := GetObjectCertKeystore(context.Background(), c)
	if err == nil {
		t.Fatal("expected error for code 999")
	}
}

func TestGetObjectCertKeystore_EmptyChain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := models.KeystoreGetResponse{Chain: ""}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	defer server.Close()

	c := newTestClient(server)
	_, err := GetObjectCertKeystore(context.Background(), c)
	if err == nil {
		t.Fatal("expected error for empty chain")
	}
}

func TestPutObjectCertKeystore_Error999(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := models.KeystoreErrorResponse{Code: 999, Description: "Permission denied", Details: "No role"}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	defer server.Close()

	c := newTestClient(server)
	err := PutObjectCertKeystore(context.Background(), c, "key", "chain")
	if err == nil {
		t.Fatal("expected error for code 999")
	}
}

func TestPutObjectCertKeystore_Auth401(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	c := newTestClient(server)
	err := PutObjectCertKeystore(context.Background(), c, "key", "chain")
	if err == nil {
		t.Fatal("expected error for 401")
	}
}

func TestPutVDCKeystore_ConnectionError(t *testing.T) {
	c := &client.Client{
		BaseURL:     "http://127.0.0.1:1",
		HTTPClient:  &http.Client{Timeout: 1},
		AuthHeaders: map[string]string{},
	}
	err := PutVDCKeystore(context.Background(), c, "key", "chain")
	if err == nil {
		t.Fatal("expected connection error")
	}
}

func TestPutObjectCertSelfSigned_FallbackReadBack(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		if callCount == 1 {
			// First call is the PUT - return success but no chain
			w.WriteHeader(http.StatusOK)
			if _, err := w.Write([]byte(`{"status":"ok"}`)); err != nil {
				t.Errorf("failed to write response: %v", err)
			}
			return
		}
		// Second call is the GET read-back
		resp := models.KeystoreGetResponse{Chain: "-----BEGIN CERTIFICATE-----\nreadback\n-----END CERTIFICATE-----"}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	defer server.Close()

	c := newTestClient(server)
	chain, err := PutObjectCertSelfSigned(context.Background(), c, []string{"10.0.0.1"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chain == "" {
		t.Error("expected non-empty chain from read-back")
	}
}

func TestPutObjectCertSelfSigned_FallbackReadBackFail(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		if callCount == 1 {
			// PUT succeeds with empty chain in response
			w.WriteHeader(http.StatusOK)
			if _, err := w.Write([]byte(`{"chain":""}`)); err != nil {
				t.Errorf("failed to write response: %v", err)
			}
			return
		}
		// GET read-back also fails
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	c := newTestClient(server)
	_, err := PutObjectCertSelfSigned(context.Background(), c, nil)
	if err == nil {
		t.Fatal("expected error when read-back fails")
	}
}

// contains is a helper for string containment checks.
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsSubstr(s, substr))
}

func containsSubstr(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
