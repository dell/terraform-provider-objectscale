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
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"strings"
	"testing"
	"time"
)

// generateTestRSAKeyPKCS1 generates a test RSA private key in PKCS#1 PEM format.
func generateTestRSAKeyPKCS1() string {
	key, _ := rsa.GenerateKey(rand.Reader, 2048)
	pkcs1Bytes := x509.MarshalPKCS1PrivateKey(key)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: pkcs1Bytes,
	}
	return string(pem.EncodeToMemory(block))
}

// generateTestRSAKeyPKCS8 generates a test RSA private key in PKCS#8 PEM format.
func generateTestRSAKeyPKCS8() string {
	key, _ := rsa.GenerateKey(rand.Reader, 2048)
	pkcs8Bytes, _ := x509.MarshalPKCS8PrivateKey(key)
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: pkcs8Bytes,
	}
	return string(pem.EncodeToMemory(block))
}

// generateTestCertificate generates a self-signed test certificate PEM.
func generateTestCertificate() string {
	key, _ := rsa.GenerateKey(rand.Reader, 2048)
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "test"},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
	}
	certBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	block := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	}
	return string(pem.EncodeToMemory(block))
}

func TestValidateAndNormalizePrivateKey_PKCS1Passthrough(t *testing.T) {
	pkcs1Key := generateTestRSAKeyPKCS1()
	result, err := ValidateAndNormalizePrivateKey(pkcs1Key)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(result, "RSA PRIVATE KEY") {
		t.Error("expected PKCS#1 key to be preserved")
	}
}

func TestValidateAndNormalizePrivateKey_PKCS8Rejected(t *testing.T) {
	pkcs8Key := generateTestRSAKeyPKCS8()
	_, err := ValidateAndNormalizePrivateKey(pkcs8Key)
	if err == nil {
		t.Fatal("expected error for PKCS#8 key")
	}
	if !strings.Contains(err.Error(), "PKCS#8") {
		t.Errorf("expected PKCS#8 error message, got: %v", err)
	}
	if !strings.Contains(err.Error(), "openssl rsa") {
		t.Error("expected conversion guidance in error message")
	}
}

func TestValidateAndNormalizePrivateKey_InvalidPEM(t *testing.T) {
	_, err := ValidateAndNormalizePrivateKey("not-a-pem-key")
	if err == nil {
		t.Error("expected error for invalid PEM")
	}
	if !strings.Contains(err.Error(), "invalid PEM") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestValidateAndNormalizePrivateKey_EncryptedKey(t *testing.T) {
	encryptedPEM := "-----BEGIN ENCRYPTED PRIVATE KEY-----\nfakedata\n-----END ENCRYPTED PRIVATE KEY-----\n"
	_, err := ValidateAndNormalizePrivateKey(encryptedPEM)
	if err == nil {
		t.Error("expected error for encrypted key")
	}
	if !strings.Contains(err.Error(), "encrypted") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestValidateAndNormalizePrivateKey_UnsupportedBlockType(t *testing.T) {
	dsaPEM := "-----BEGIN DSA PRIVATE KEY-----\nbm90YXZhbGlka2V5\n-----END DSA PRIVATE KEY-----\n"
	_, err := ValidateAndNormalizePrivateKey(dsaPEM)
	if err == nil {
		t.Error("expected error for unsupported block type")
	}
	if !strings.Contains(err.Error(), "unsupported private key type") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestNormalizeLineEndings(t *testing.T) {
	input := "line1\r\nline2\r\nline3\n"
	result := NormalizeLineEndings(input)
	if strings.Contains(result, "\r\n") {
		t.Error("CRLF should be normalized to LF")
	}
	if result != "line1\nline2\nline3\n" {
		t.Errorf("unexpected result: %q", result)
	}
}

func TestCompareCertificateChains_Identical(t *testing.T) {
	cert := generateTestCertificate()
	certCRLF := strings.ReplaceAll(cert, "\n", "\r\n")
	if !CompareCertificateChains(cert, certCRLF) {
		t.Error("chains should be identical after normalization")
	}
}

func TestCompareCertificateChains_Different(t *testing.T) {
	cert1 := generateTestCertificate()
	cert2 := generateTestCertificate()
	if CompareCertificateChains(cert1, cert2) {
		t.Error("different chains should not compare as equal")
	}
}

func TestValidatePEMCertificate_Valid(t *testing.T) {
	cert := generateTestCertificate()
	if err := ValidatePEMCertificate(cert); err != nil {
		t.Fatalf("unexpected error for valid cert: %v", err)
	}
}

func TestValidatePEMCertificate_Invalid(t *testing.T) {
	if err := ValidatePEMCertificate("not-a-cert"); err == nil {
		t.Error("expected error for invalid cert")
	}
}

func TestValidatePEMCertificate_WrongType(t *testing.T) {
	key := generateTestRSAKeyPKCS1()
	err := ValidatePEMCertificate(key)
	if err == nil {
		t.Error("expected error for non-CERTIFICATE PEM block")
	}
	if !strings.Contains(err.Error(), "expected CERTIFICATE block") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestValidatePEMPrivateKey_Valid(t *testing.T) {
	key := generateTestRSAKeyPKCS1()
	if err := ValidatePEMPrivateKey(key); err != nil {
		t.Fatalf("unexpected error for valid key: %v", err)
	}
}

func TestValidatePEMPrivateKey_PKCS8Rejected(t *testing.T) {
	key := generateTestRSAKeyPKCS8()
	err := ValidatePEMPrivateKey(key)
	if err == nil {
		t.Fatal("expected error for PKCS#8 key")
	}
	if !strings.Contains(err.Error(), "PKCS#8") {
		t.Errorf("expected PKCS#8 error, got: %v", err)
	}
}

func TestValidatePEMPrivateKey_Invalid(t *testing.T) {
	if err := ValidatePEMPrivateKey("not-a-key"); err == nil {
		t.Error("expected error for invalid key")
	}
}

func TestValidatePEMPrivateKey_Encrypted(t *testing.T) {
	encryptedPEM := "-----BEGIN ENCRYPTED PRIVATE KEY-----\nfakedata\n-----END ENCRYPTED PRIVATE KEY-----\n"
	err := ValidatePEMPrivateKey(encryptedPEM)
	if err == nil {
		t.Error("expected error for encrypted key")
	}
	if !strings.Contains(err.Error(), "encrypted") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestValidatePEMPrivateKey_UnsupportedType(t *testing.T) {
	cert := generateTestCertificate()
	err := ValidatePEMPrivateKey(cert)
	if err == nil {
		t.Error("expected error for certificate passed as key")
	}
	if !strings.Contains(err.Error(), "unsupported") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestValidateAndNormalizePrivateKey_CRLFNormalized(t *testing.T) {
	pkcs1Key := generateTestRSAKeyPKCS1()
	crlfKey := strings.ReplaceAll(pkcs1Key, "\n", "\r\n")
	result, err := ValidateAndNormalizePrivateKey(crlfKey)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.Contains(result, "\r\n") {
		t.Error("CRLF should be normalized to LF")
	}
}

func TestCompareCertificateChains_WithWhitespace(t *testing.T) {
	cert := generateTestCertificate()
	certWithSpaces := "\n\n" + cert + "\n\n"
	if !CompareCertificateChains(cert, certWithSpaces) {
		t.Error("chains should be identical after trimming")
	}
}
