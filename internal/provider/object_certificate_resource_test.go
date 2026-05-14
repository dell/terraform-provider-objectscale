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
	"fmt"
	"regexp"
	"testing"

	"github.com/bytedance/mockey"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccObjectCertificateResource_CustomCert(t *testing.T) {
	testKey := generateTestKey()
	testCert := generateTestCert()

	loginM := loginMocker()
	defer loginM.UnPatch()
	getM := mockey.Mock(GetObjectCertKeystore).To(func(ctx context.Context, c interface{}) (string, error) {
		return testCert, nil
	}).Build()
	putM := mockey.Mock(PutObjectCertKeystore).To(func(ctx context.Context, c interface{}, pk, cc string) error {
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
					resource "objectscale_object_certificate" "test" {
						private_key       = %q
						certificate_chain = %q
					}
				`, testKey, testCert),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_object_certificate.test", "id", "object_certificate"),
					resource.TestCheckResourceAttrSet("objectscale_object_certificate.test", "current_certificate_chain"),
				),
			},
		},
	})
}

func TestAccObjectCertificateResource_SelfSigned(t *testing.T) {
	selfSignedChain := "-----BEGIN CERTIFICATE-----\nselfsigned\n-----END CERTIFICATE-----"
	loginM := loginMocker()
	defer loginM.UnPatch()
	getM := mockey.Mock(GetObjectCertKeystore).To(func(ctx context.Context, c interface{}) (string, error) {
		return selfSignedChain, nil
	}).Build()
	putM := mockey.Mock(PutObjectCertSelfSigned).To(func(ctx context.Context, c interface{}, ips []string) (string, error) {
		return selfSignedChain, nil
	}).Build()
	defer getM.UnPatch()
	defer putM.UnPatch()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
					resource "objectscale_object_certificate" "test" {
						system_selfsigned = true
						ip_addresses      = ["10.0.0.1", "10.0.0.2"]
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_object_certificate.test", "id", "object_certificate"),
					resource.TestCheckResourceAttrSet("objectscale_object_certificate.test", "current_certificate_chain"),
				),
			},
		},
	})
}

func TestAccObjectCertificateResource_SelfSignedError(t *testing.T) {
	loginM := loginMocker()
	defer loginM.UnPatch()
	putM := mockey.Mock(PutObjectCertSelfSigned).To(func(ctx context.Context, c interface{}, ips []string) (string, error) {
		return "", fmt.Errorf("server error")
	}).Build()
	defer putM.UnPatch()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
					resource "objectscale_object_certificate" "test" {
						system_selfsigned = true
					}
				`,
				ExpectError: regexp.MustCompile(`Error generating self-signed Object certificate`),
			},
		},
	})
}

func TestAccObjectCertificateResource_CustomCertGetError(t *testing.T) {
	testKey := generateTestKey()
	testCert := generateTestCert()

	loginM := loginMocker()
	defer loginM.UnPatch()
	getM := mockey.Mock(GetObjectCertKeystore).To(func(ctx context.Context, c interface{}) (string, error) {
		return "", fmt.Errorf("connection refused")
	}).Build()
	defer getM.UnPatch()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + fmt.Sprintf(`
					resource "objectscale_object_certificate" "test" {
						private_key       = %q
						certificate_chain = %q
					}
				`, testKey, testCert),
				ExpectError: regexp.MustCompile(`Error reading current Object certificate`),
			},
		},
	})
}

func TestAccObjectCertificateResource_CustomCertPutError(t *testing.T) {
	testKey := generateTestKey()
	testCert := generateTestCert()
	differentCert := generateTestCert()

	loginM := loginMocker()
	defer loginM.UnPatch()
	getM := mockey.Mock(GetObjectCertKeystore).To(func(ctx context.Context, c interface{}) (string, error) {
		return differentCert, nil
	}).Build()
	putM := mockey.Mock(PutObjectCertKeystore).To(func(ctx context.Context, c interface{}, pk, cc string) error {
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
					resource "objectscale_object_certificate" "test" {
						private_key       = %q
						certificate_chain = %q
					}
				`, testKey, testCert),
				ExpectError: regexp.MustCompile(`Error updating Object certificate`),
			},
		},
	})
}

func TestAccObjectCertificateResource_ConflictValidation(t *testing.T) {
	testKey := generateTestKey()
	testCert := generateTestCert()

	loginM := loginMocker()
	defer loginM.UnPatch()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + fmt.Sprintf(`
					resource "objectscale_object_certificate" "test" {
						system_selfsigned = true
						private_key       = %q
						certificate_chain = %q
					}
				`, testKey, testCert),
				ExpectError: regexp.MustCompile(`Invalid Attribute Combination|Conflicting`),
			},
		},
	})
}

func TestAccObjectCertificateResource_ReadError(t *testing.T) {
	selfSignedChain := "-----BEGIN CERTIFICATE-----\nselfsigned\n-----END CERTIFICATE-----"
	loginM := loginMocker()
	defer loginM.UnPatch()
	putM := mockey.Mock(PutObjectCertSelfSigned).To(func(ctx context.Context, c interface{}, ips []string) (string, error) {
		return selfSignedChain, nil
	}).Build()
	defer putM.UnPatch()
	getM := mockey.Mock(GetObjectCertKeystore).To(func(ctx context.Context, c interface{}) (string, error) {
		return "", fmt.Errorf("read error")
	}).Build()
	defer getM.UnPatch()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
					resource "objectscale_object_certificate" "test" {
						system_selfsigned = true
					}
				`,
				ExpectError: regexp.MustCompile(`Error reading Object certificate`),
			},
		},
	})
}
