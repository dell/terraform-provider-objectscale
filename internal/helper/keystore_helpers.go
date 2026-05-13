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
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
)

// ValidateAndNormalizePrivateKey validates that the private key is in PKCS#1 (RSA PRIVATE KEY) PEM format
// and normalizes line endings. PKCS#8 keys are automatically converted to PKCS#1 format for compatibility
// with ObjectScale 4.1 which requires PKCS#1 format.
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
		// PKCS#8 — convert to PKCS#1 automatically
		pkcs1Key, err := convertPKCS8ToPKCS1(block.Bytes)
		if err != nil {
			return "", fmt.Errorf("failed to convert PKCS#8 to PKCS#1: %w", err)
		}
		pkcs1Block := &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: pkcs1Key,
		}
		pkcs1PEM := pem.EncodeToMemory(pkcs1Block)
		return string(pkcs1PEM), nil

	case "ENCRYPTED PRIVATE KEY":
		return "", errors.New("encrypted (passphrase-protected) private keys are not supported")

	default:
		return "", fmt.Errorf("unsupported private key type: %s", block.Type)
	}
}

// convertPKCS8ToPKCS1 converts a PKCS#8 encoded RSA private key to PKCS#1 format.
func convertPKCS8ToPKCS1(pkcs8Bytes []byte) ([]byte, error) {
	key, err := x509.ParsePKCS8PrivateKey(pkcs8Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PKCS#8 private key: %w", err)
	}

	rsaKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("PKCS#8 key is not an RSA key")
	}

	return x509.MarshalPKCS1PrivateKey(rsaKey), nil
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
// PKCS#8 keys are automatically converted to PKCS#1 format for compatibility.
func ValidatePEMPrivateKey(pemStr string) error {
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil {
		return errors.New("invalid PEM: no PEM block found")
	}
	switch block.Type {
	case "RSA PRIVATE KEY":
		return nil
	case "PRIVATE KEY":
		// PKCS#8 is acceptable - will be converted to PKCS#1 by ValidateAndNormalizePrivateKey
		return nil
	case "ENCRYPTED PRIVATE KEY":
		return errors.New("encrypted (passphrase-protected) private keys are not supported")
	default:
		return fmt.Errorf("unsupported private key type: %s", block.Type)
	}
}
