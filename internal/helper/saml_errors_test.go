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

package helper

import (
	"errors"
	"fmt"
	"testing"
)

// fakeHTTPErr implements HTTPStatusError.
type fakeHTTPErr struct {
	status int
	msg    string
}

func (e *fakeHTTPErr) Error() string   { return e.msg }
func (e *fakeHTTPErr) HTTPStatus() int { return e.status }

// U-17 — Parse error response — 400.
func TestU17_ClassifyError_400(t *testing.T) {
	err := &fakeHTTPErr{status: 400, msg: "bad request"}
	if got := ClassifyError(err); got != SAMLErrBadRequest {
		t.Fatalf("U-17 (typed): got %v want SAMLErrBadRequest", got)
	}
	// textual fallback
	if got := ClassifyError(fmt.Errorf("HTTP 400: bad request")); got != SAMLErrBadRequest {
		t.Fatalf("U-17 (text): got %v want SAMLErrBadRequest", got)
	}
}

// U-18 — Parse error response — 404.
func TestU18_ClassifyError_404(t *testing.T) {
	err := &fakeHTTPErr{status: 404, msg: "not found"}
	if got := ClassifyError(err); got != SAMLErrNotFound {
		t.Fatalf("U-18 (typed): got %v want SAMLErrNotFound", got)
	}
	if !IsSAMLNotFound(err) {
		t.Fatalf("U-18: IsSAMLNotFound = false, want true")
	}
	if got := ClassifyError(fmt.Errorf("HTTP 404: provider not found")); got != SAMLErrNotFound {
		t.Fatalf("U-18 (text): got %v want SAMLErrNotFound", got)
	}
}

// U-18b — Other status codes classify correctly.
func TestU18b_ClassifyError_OtherStatuses(t *testing.T) {
	cases := []struct {
		status int
		want   SAMLErrorClass
	}{
		{401, SAMLErrUnauthorized},
		{403, SAMLErrForbidden},
		{409, SAMLErrConflict},
		{429, SAMLErrRateLimited},
		{500, SAMLErrServer},
		{502, SAMLErrServer},
		{503, SAMLErrServer},
	}
	for _, c := range cases {
		err := &fakeHTTPErr{status: c.status, msg: fmt.Sprintf("HTTP %d", c.status)}
		if got := ClassifyError(err); got != c.want {
			t.Errorf("status %d: got %v want %v", c.status, got, c.want)
		}
	}
}

// U-18c — Nil error → SAMLErrUnknown.
func TestU18c_ClassifyError_Nil(t *testing.T) {
	if got := ClassifyError(nil); got != SAMLErrUnknown {
		t.Fatalf("nil err: got %v want SAMLErrUnknown", got)
	}
	if IsSAMLNotFound(nil) {
		t.Fatalf("IsSAMLNotFound(nil) = true, want false")
	}
}

// U-18d — Wrapped error preserves classification.
func TestU18d_ClassifyError_Wrapped(t *testing.T) {
	inner := &fakeHTTPErr{status: 404, msg: "missing"}
	wrapped := fmt.Errorf("during get: %w", inner)
	if got := ClassifyError(wrapped); got != SAMLErrNotFound {
		t.Fatalf("wrapped 404: got %v want SAMLErrNotFound", got)
	}
	// errors.As must surface the inner.
	var hse HTTPStatusError
	if !errors.As(wrapped, &hse) || hse.HTTPStatus() != 404 {
		t.Fatalf("errors.As failed for wrapped HTTPStatusError")
	}
}
