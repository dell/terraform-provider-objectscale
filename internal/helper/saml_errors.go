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
	"strings"
)

// SAMLErrorClass classifies an HTTP error returned by the ObjectScale IAM /
// Service Provider APIs into a stable category that resource code can switch on.
type SAMLErrorClass int

const (
	SAMLErrUnknown      SAMLErrorClass = iota
	SAMLErrBadRequest                  // 400
	SAMLErrUnauthorized                // 401
	SAMLErrForbidden                   // 403
	SAMLErrNotFound                    // 404
	SAMLErrConflict                    // 409
	SAMLErrRateLimited                 // 429
	SAMLErrServer                      // 5xx
)

// HTTPStatusError is the common interface implemented by manual clientgen
// wrappers' error types. It exposes the HTTP status the server returned.
type HTTPStatusError interface {
	error
	HTTPStatus() int
}

// ClassifyError inspects an error returned from the ObjectScale client and
// returns a SAMLErrorClass.
//
// It first looks for an HTTPStatusError; if not found, it heuristically
// inspects the textual error for status hints (used by parsed XML/JSON error
// bodies emitted by the OpenAPI generator).
func ClassifyError(err error) SAMLErrorClass {
	if err == nil {
		return SAMLErrUnknown
	}
	var hse HTTPStatusError
	if errors.As(err, &hse) {
		return classifyByStatus(hse.HTTPStatus())
	}
	msg := strings.ToLower(err.Error())
	switch {
	case strings.Contains(msg, "400"):
		return SAMLErrBadRequest
	case strings.Contains(msg, "401"):
		return SAMLErrUnauthorized
	case strings.Contains(msg, "403"):
		return SAMLErrForbidden
	case strings.Contains(msg, "404") || strings.Contains(msg, "not found"):
		return SAMLErrNotFound
	case strings.Contains(msg, "409") || strings.Contains(msg, "already exists"):
		return SAMLErrConflict
	case strings.Contains(msg, "429"):
		return SAMLErrRateLimited
	case strings.Contains(msg, "500") || strings.Contains(msg, "502") ||
		strings.Contains(msg, "503") || strings.Contains(msg, "504"):
		return SAMLErrServer
	default:
		return SAMLErrUnknown
	}
}

func classifyByStatus(status int) SAMLErrorClass {
	switch {
	case status == 400:
		return SAMLErrBadRequest
	case status == 401:
		return SAMLErrUnauthorized
	case status == 403:
		return SAMLErrForbidden
	case status == 404:
		return SAMLErrNotFound
	case status == 409:
		return SAMLErrConflict
	case status == 429:
		return SAMLErrRateLimited
	case status >= 500 && status < 600:
		return SAMLErrServer
	default:
		return SAMLErrUnknown
	}
}

// IsSAMLNotFound is a convenience that returns true if the error classifies
// as 404 / not-found. Used by Read/Delete to gracefully accept missing
// resources.
func IsSAMLNotFound(err error) bool {
	return ClassifyError(err) == SAMLErrNotFound
}
