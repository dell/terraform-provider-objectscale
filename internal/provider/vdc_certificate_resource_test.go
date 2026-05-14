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
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"regexp"
	"testing"
	"time"

	"github.com/bytedance/mockey"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func generateTestKey() string {
	key, _ := rsa.GenerateKey(rand.Reader, 2048)
	return string(pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}))
}

func generateTestCert() string {
	key, _ := rsa.GenerateKey(rand.Reader, 2048)
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "test"},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
	}
	certBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	return string(pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	}))
}

func TestAccVDCCertificateResource_Create(t *testing.T) {
	testKey := generateTestKey()
	testCert := generateTestCert()

	loginM := loginMocker()
	defer loginM.UnPatch()
	getM := mockey.Mock(GetVDCKeystore).To(func(ctx context.Context, c interface{}) (string, error) {
		return testCert, nil
	}).Build()
	putM := mockey.Mock(PutVDCKeystore).To(func(ctx context.Context, c interface{}, pk, cc string) error {
		return nil
	}).Build()
	defer getM.UnPatch()
	defer putM.UnPatch()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + fmt.Sprintf(`
					resource "objectscale_vdc_certificate" "test" {
						private_key       = %q
						certificate_chain = %q
					}
				`, testKey, testCert),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_vdc_certificate.test", "id", "vdc_certificate"),
					resource.TestCheckResourceAttrSet("objectscale_vdc_certificate.test", "current_certificate_chain"),
				),
			},
		},
	})
}

func TestAccVDCCertificateResource_CreateGetError(t *testing.T) {
	testKey := generateTestKey()
	testCert := generateTestCert()

	loginM := loginMocker()
	defer loginM.UnPatch()
	getM := mockey.Mock(GetVDCKeystore).To(func(ctx context.Context, c interface{}) (string, error) {
		return "", fmt.Errorf("connection refused")
	}).Build()
	defer getM.UnPatch()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + fmt.Sprintf(`
					resource "objectscale_vdc_certificate" "test" {
						private_key       = %q
						certificate_chain = %q
					}
				`, testKey, testCert),
				ExpectError: regexp.MustCompile(`Error reading current VDC certificate`),
			},
		},
	})
}

func TestAccVDCCertificateResource_PutError(t *testing.T) {
	testKey := generateTestKey()
	testCert := generateTestCert()
	differentCert := generateTestCert()

	loginM := loginMocker()
	defer loginM.UnPatch()
	getM := mockey.Mock(GetVDCKeystore).To(func(ctx context.Context, c interface{}) (string, error) {
		return differentCert, nil
	}).Build()
	putM := mockey.Mock(PutVDCKeystore).To(func(ctx context.Context, c interface{}, pk, cc string) error {
		return fmt.Errorf("server error")
	}).Build()
	defer getM.UnPatch()
	defer putM.UnPatch()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + fmt.Sprintf(`
					resource "objectscale_vdc_certificate" "test" {
						private_key       = %q
						certificate_chain = %q
					}
				`, testKey, testCert),
				ExpectError: regexp.MustCompile(`Error updating VDC certificate`),
			},
		},
	})
}

func TestAccVDCCertificateResource_ReadError(t *testing.T) {
	testKey := generateTestKey()
	testCert := generateTestCert()

	loginM := loginMocker()
	defer loginM.UnPatch()
	getM := mockey.Mock(GetVDCKeystore).To(func(ctx context.Context, c interface{}) (string, error) {
		return "", fmt.Errorf("read error")
	}).Build()
	defer getM.UnPatch()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + fmt.Sprintf(`
					resource "objectscale_vdc_certificate" "test" {
						private_key       = %q
						certificate_chain = %q
					}
				`, testKey, testCert),
				ExpectError: regexp.MustCompile(`Error reading.*VDC certificate`),
			},
		},
	})
}
