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
	"context"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
)

// OBSVersion represents the detected ObjectScale version
type OBSVersion string

const (
	OBSVersion41   OBSVersion = "4.1"
	OBSVersion43Plus OBSVersion = "4.3+"
	OBSVersionUnknown OBSVersion = "unknown"
)

// DetectOBSDetectedVersion attempts to detect the OBS version based on API response.
// Currently uses error-based detection: if PKCS#8 is rejected with error 1008, assume OBS 4.1.
// This can be enhanced with a dedicated version endpoint if available.
func DetectOBSDetectedVersion(ctx context.Context, pkcs8Rejected bool) OBSVersion {
	if pkcs8Rejected {
		// PKCS#8 rejection indicates OBS 4.1
		return OBSVersion41
	}
	// Default to 4.3+ if PKCS#8 is accepted
	return OBSVersion43Plus
}

// SupportsPKCS8 returns true if the detected OBS version supports PKCS#8
func SupportsPKCS8(version OBSVersion) bool {
	return version == OBSVersion43Plus || version == OBSVersionUnknown
}

// ValidateAndNormalizePrivateKey validates that the private key is in PKCS#1 (RSA PRIVATE KEY) or PKCS#8 (PRIVATE KEY) PEM format
// and normalizes line endings. PKCS#8 is supported on OBS 4.3+, but OBS 4.1 only supports PKCS#1.
// The provider accepts both formats and lets the API handle version-specific validation.
func ValidateAndNormalizePrivateKey(pemKey string) (string, error) {
	normalized := NormalizeLineEndings(pemKey)
	block, _ := pem.Decode([]byte(normalized))
	if block == nil {
		return "", errors.New("invalid PEM: failed to decode private key")
	}

	switch block.Type {
	case "RSA PRIVATE KEY":
		// PKCS#1 — accepted on all OBS versions
		return normalized, nil

	case "PRIVATE KEY":
		// PKCS#8 — accepted by provider, API will validate based on OBS version
		// OBS 4.1 rejects PKCS#8 with error 1008, OBS 4.3+ accepts it
		return normalized, nil

	case "ENCRYPTED PRIVATE KEY":
		return "", errors.New("encrypted (passphrase-protected) private keys are not supported")

	default:
		return "", fmt.Errorf("unsupported private key type: %s", block.Type)
	}
}

// NormalizeLineEndings replaces all \r\n with \n.
func NormalizeLineEndings(s string) string {
	return strings.ReplaceAll(s, "\r\n", "\n")
}

// CompareCertificateChains compares two PEM certificate chains after normalization.
// Returns true if they are identical after line ending normalization and trimming.
func CompareCertificateChains(a, b string) bool {
	normA := strings.TrimSpace(NormalizeLineEndings(a))
	normB := strings.TrimSpace(NormalizeLineEndings(b))
	return normA == normB
}

// ValidatePEMCertificate validates that the given string contains at least one PEM CERTIFICATE block.
func ValidatePEMCertificate(pemStr string) error {
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil {
		return errors.New("invalid PEM: no PEM block found")
	}
	if block.Type != "CERTIFICATE" {
		return fmt.Errorf("expected CERTIFICATE block, got %s", block.Type)
	}
	return nil
}

// ValidatePEMPrivateKey validates that the given string contains a valid PEM private key in PKCS#1 or PKCS#8 format.
// PKCS#8 is supported on OBS 4.3+, but OBS 4.1 only supports PKCS#1.
// The provider accepts both formats and lets the API handle version-specific validation.
func ValidatePEMPrivateKey(pemStr string) error {
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil {
		return errors.New("invalid PEM: no PEM block found")
	}
	switch block.Type {
	case "RSA PRIVATE KEY":
		// PKCS#1 — accepted on all OBS versions
		return nil
	case "PRIVATE KEY":
		// PKCS#8 — accepted by provider, API will validate based on OBS version
		// OBS 4.1 rejects PKCS#8 with error 1008, OBS 4.3+ accepts it
		return nil
	case "ENCRYPTED PRIVATE KEY":
		return errors.New("encrypted (passphrase-protected) private keys are not supported")
	default:
		return fmt.Errorf("unsupported private key type: %s", block.Type)
	}
}
