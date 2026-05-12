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

package helper

import (
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
)

// ValidateAndNormalizePrivateKey validates that the private key is in PKCS#1 (RSA PRIVATE KEY) PEM format
// and normalizes line endings. PKCS#8 keys are rejected with an explicit error because ObjectScale 4.1
// does not support PKCS#8 format. Users targeting ObjectScale 4.3+ who have PKCS#8 keys should convert
// them externally (e.g., openssl rsa -in key.pem -out key-pkcs1.pem) before supplying them to this provider.
func ValidateAndNormalizePrivateKey(pemKey string) (string, error) {
	normalized := NormalizeLineEndings(pemKey)
	block, _ := pem.Decode([]byte(normalized))
	if block == nil {
		return "", errors.New("invalid PEM: failed to decode private key")
	}

	switch block.Type {
	case "RSA PRIVATE KEY":
		// PKCS#1 — accepted
		return normalized, nil

	case "PRIVATE KEY":
		// PKCS#8 — not supported; reject with guidance
		return "", errors.New(
			"PKCS#8 private keys (PEM type 'PRIVATE KEY') are not supported. " +
				"ObjectScale 4.1 requires PKCS#1 format (PEM type 'RSA PRIVATE KEY'). " +
				"Convert your key with: openssl rsa -in key.pem -out key-pkcs1.pem")

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

// ValidatePEMPrivateKey validates that the given string contains a valid PEM private key in PKCS#1 format.
// PKCS#8 keys are rejected because ObjectScale 4.1 does not support them.
func ValidatePEMPrivateKey(pemStr string) error {
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil {
		return errors.New("invalid PEM: no PEM block found")
	}
	switch block.Type {
	case "RSA PRIVATE KEY":
		return nil
	case "PRIVATE KEY":
		return errors.New(
			"PKCS#8 private keys (PEM type 'PRIVATE KEY') are not supported. " +
				"ObjectScale 4.1 requires PKCS#1 format (PEM type 'RSA PRIVATE KEY'). " +
				"Convert your key with: openssl rsa -in key.pem -out key-pkcs1.pem")
	case "ENCRYPTED PRIVATE KEY":
		return errors.New("encrypted (passphrase-protected) private keys are not supported")
	default:
		return fmt.Errorf("unsupported private key type: %s", block.Type)
	}
}
