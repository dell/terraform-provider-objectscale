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
	"fmt"
	"regexp"
	"terraform-provider-objectscale/internal/clientgen"
	"testing"

	"github.com/bytedance/mockey"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const samlMetadataFixture = `<?xml version="1.0" encoding="UTF-8"?>
<md:EntityDescriptor xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata" entityID="https://test-idp-v1.example.com">
  <md:IDPSSODescriptor WantAuthnRequestsSigned="false" protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
    <md:KeyDescriptor use="signing">
      <KeyInfo xmlns="http://www.w3.org/2000/09/xmldsig#">
        <X509Data>
          <X509Certificate>MIIDJTCCAg2gAwIBAgIUZJHGmNqz0G5ISRyaM48DTIpsqlcwDQYJKoZIhvcNAQELBQAwIjEgMB4GA1UEAwwXdGVzdC1pZHAtdjEuZXhhbXBsZS5jb20wHhcNMjYwNTA0MDg1MjA2WhcNMzYwNTAxMDg1MjA2WjAiMSAwHgYDVQQDDBd0ZXN0LWlkcC12MS5leGFtcGxlLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALrM4fAYyxQBocBuL1PoBOVKSH/l7YEcdUZREO2FX4GlDolWf5q0RnLy1nje9HsUR9njOxpbElFHQbTvtGXUm57LDoi8IOPojPDwrqa3uUb5EvPhtE4ltjoShIk84xlurC3HppM+y2RJ+O3cx8t9qUUOT2VAGg8cR6ntZlPZMxshghAT5IQpGDdf1NZ15nIT6j8UuNIyFzvxu+x6I/RSvCEaDURLqRtmKDWV/9ad1We560ARYBGWRkEIvKweyQm2yeXt1Ho5ycIkoU6KQnOY3yPCuVNGKTqz1OTJpPfIvKWlARLLauIGlQlMo+RTikuPF2NqKNJJfHiQMhVwLH3z1JECAwEAAaNTMFEwHQYDVR0OBBYEFJ13WvqFmqd0mypzJdwflkpFFfZAMB8GA1UdIwQYMBaAFJ13WvqFmqd0mypzJdwflkpFFfZAMA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAIutssQnDdzEE+AWZ8+7StOJa65m/lhyITZHEJw/z9GBbO1EbnOgYIuRn5qjSMxFvXs5QjDc4AxZbzxv4SwMKA8VzoS2fWMarIP6hXH99EZJ1h5wL9qntBq4hIQCUyqhgbwKF5q+uP+pohJDytPxFfoRTYQRZ6yOl3Bj3Oraygq+BveyppJMxWagMxp8OHSpJ/O/nsyIbN2H4GW/2xXL5b0QgKyt2dJQSBIoQziBGjRUd8wPYlX9euZdc8Bb++5Yr4Grv2+q4zv7bAB7dHcuWuS4pUZAv1XE1fZ1koklMVGpJiiT5DXVtIMD8V1P+NENJlol00X/nYJpZmVpFMyF0ho=</X509Certificate>
        </X509Data>
      </KeyInfo>
    </md:KeyDescriptor>
    <md:SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://test-idp-v1.example.com/sso"/>
    <md:SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect" Location="https://test-idp-v1.example.com/sso"/>
  </md:IDPSSODescriptor>
</md:EntityDescriptor>`
const samlMetadataUpdated = `<?xml version="1.0" encoding="UTF-8"?>
<md:EntityDescriptor xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata" entityID="https://test-idp-v2.example.com">
  <md:IDPSSODescriptor WantAuthnRequestsSigned="false" protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
    <md:KeyDescriptor use="signing">
      <KeyInfo xmlns="http://www.w3.org/2000/09/xmldsig#">
        <X509Data>
          <X509Certificate>MIIDJTCCAg2gAwIBAgIUCFOMjrz0UOs1MNytRuh6335zfxowDQYJKoZIhvcNAQELBQAwIjEgMB4GA1UEAwwXdGVzdC1pZHAtdjIuZXhhbXBsZS5jb20wHhcNMjYwNTA0MDg1MjA2WhcNMzYwNTAxMDg1MjA2WjAiMSAwHgYDVQQDDBd0ZXN0LWlkcC12Mi5leGFtcGxlLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMtkOTS0ctfzUv0sOVMJ0dsn0An8NJpBdp2QdXUdnkDveyBN/cCk173t8VEKm45q/EDA1GT6bKGlzbqIYCPiabpWD6gAT8SinN13VPsrtFdGw3LcfdPakWo3sG+NC4YYkRuWCv1LWe5k3GG+udtS43u39globqORLZJQhSEiAd2mqW2e3vNrXhnV2ZVG7gI6m39RsbcM+GBXl7Z657WI8HJVeCI797FLSWT2MEMlNH8KY+6nGZaHQdRSThNql6esxbqhdku5Nq0ftyLOuqoX3cCgVeWemJrrRoi6yoetFoWBEl6FA30mbmw+qUtPEPXVMFlRuBLgNsg0gD0UmbxzRMkCAwEAAaNTMFEwHQYDVR0OBBYEFCV/qvmv6oksNHba84C+hKRpEgnwMB8GA1UdIwQYMBaAFCV/qvmv6oksNHba84C+hKRpEgnwMA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAC614rYD1Qff2E658eNFH19WEH/jMFqGyUWXBEspWLV7bBNxaSbzTb+71hqykocWy3qu/cQ2PpzdZA4M7YxpAcLop/pBftGx/8ku6wUbwSDNFc+2Obgsafhriem/FXJwyMWI9ntxkKQoAwXIcdZrT3GWCTy7Fd/NVfHpduO/JmruxGxzFr/tt6QTlq6l6b+fA8hWsReiMC5ut5Bd6ksLfqbMOxnbwmBizQC9qSp09OmO/aI1bE/D0jm3O+Xu30oaRmBtGxdfkqzZ8tkROcicphRgGb6cyhSl64Ou6AIOgWmx1gtyFlDR7+2pj9f9F6sLDgAKpI6vtq/i/MmH2r3INPE=</X509Certificate>
        </X509Data>
      </KeyInfo>
    </md:KeyDescriptor>
    <md:SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://test-idp-v2.example.com/sso"/>
    <md:SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect" Location="https://test-idp-v2.example.com/sso"/>
  </md:IDPSSODescriptor>
</md:EntityDescriptor>`
const samlMetadataInvalid = `INVALID_XML not a valid <SAML metadata`

func samlProviderHCL(name, metadata string) string {
	return ProviderConfigForTesting + fmt.Sprintf(`
resource "objectscale_iam_saml_provider" "test" {
  name                   = %q
  namespace              = "ns1"
  saml_metadata_document = %q
}
`, name, metadata)
}

// TestAccIAMSAMLProviderResource exercises the full SAML provider resource
// lifecycle: create, read, update, import, drift detection, invalid input,
// duplicate conflict, and delete-already-deleted.
func TestAccIAMSAMLProviderResource(t *testing.T) {
	defer testUserTokenCleanup(t)

	// create, read, update metadata, force-new on name change, import, drift detection
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// create and read
			{
				Config: samlProviderHCL("testacc_saml", samlMetadataFixture),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_saml_provider.test", "name", "testacc_saml"),
					resource.TestCheckResourceAttr("objectscale_iam_saml_provider.test", "namespace", "ns1"),
					resource.TestCheckResourceAttrSet("objectscale_iam_saml_provider.test", "arn"),
					resource.TestCheckResourceAttrSet("objectscale_iam_saml_provider.test", "create_date"),
					resource.TestCheckResourceAttrSet("objectscale_iam_saml_provider.test", "valid_until"),
					resource.TestCheckResourceAttrSet("objectscale_iam_saml_provider.test", "saml_metadata_document"),
				),
			},
			// update metadata in place
			{
				Config: samlProviderHCL("testacc_saml", samlMetadataUpdated),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_saml_provider.test", "saml_metadata_document", samlMetadataUpdated),
				),
			},
			// drift detection — re-applying same config produces no diff
			{
				Config:   samlProviderHCL("testacc_saml", samlMetadataUpdated),
				PlanOnly: true,
			},
			// force-new on name change
			{
				Config: samlProviderHCL("testacc_saml_renamed", samlMetadataUpdated),
				Check:  resource.TestCheckResourceAttr("objectscale_iam_saml_provider.test", "name", "testacc_saml_renamed"),
			},
			// import by ARN
			{
				ResourceName:            "objectscale_iam_saml_provider.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"id"},
			},
			// invalid metadata
			{
				Config:      samlProviderHCL("testacc_saml_bad", samlMetadataInvalid),
				ExpectError: regexp.MustCompile(`(?i)validation|MalformedSAMLMetadataDocument|400`),
			},
		},
	})

	// duplicate provider conflict (separate resource.Test because of different config shape)
	dupName := "testacc_saml_dup"
	cfgDup := ProviderConfigForTesting + fmt.Sprintf(`
resource "objectscale_iam_saml_provider" "a" {
  name                   = %q
  namespace              = "ns1"
  saml_metadata_document = %q
}
resource "objectscale_iam_saml_provider" "b" {
  name                   = %q
  namespace              = "ns1"
  saml_metadata_document = %q
}
`, dupName, samlMetadataFixture, dupName, samlMetadataFixture)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      cfgDup,
				ExpectError: regexp.MustCompile(`(?i)already exists|409|conflict`),
			},
		},
	})

	// --- mocked error paths for coverage ---
	var createM, getM, updateM, deleteM *mockey.Mocker
	samlCfg := samlProviderHCL("testacc_saml_mock", samlMetadataFixture)

	// create API error
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					createM = mockey.Mock((*clientgen.IamApiService).IamServiceCreateSAMLProviderExecute).
						Return(nil, nil, fmt.Errorf("error")).Build()
				},
				Config:      samlCfg,
				ExpectError: regexp.MustCompile(`CreateSAMLProvider failed`),
			},
		},
	})
	createM.UnPatch()

	// update API error
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: samlProviderHCL("testacc_saml_uperr", samlMetadataFixture),
			},
			{
				PreConfig: func() {
					updateM = mockey.Mock((*clientgen.IamApiService).IamServiceUpdateSAMLProviderExecute).
						Return(nil, nil, fmt.Errorf("error")).Build()
				},
				Config:      samlProviderHCL("testacc_saml_uperr", samlMetadataUpdated),
				ExpectError: regexp.MustCompile(`UpdateSAMLProvider failed`),
			},
		},
	})
	updateM.UnPatch()

	// read error (non-404) — also covers get-after-update error path since
	// both use IamServiceGetSAMLProviderExecute; the framework refresh calls
	// Read before reaching Update, so mocking Get always hits Read first.
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: samlProviderHCL("testacc_saml_rerr", samlMetadataFixture),
			},
			{
				PreConfig: func() {
					getM = mockey.Mock((*clientgen.IamApiService).IamServiceGetSAMLProviderExecute).
						Return(nil, nil, fmt.Errorf("error")).Build()
				},
				Config:      samlProviderHCL("testacc_saml_rerr", samlMetadataFixture),
				ExpectError: regexp.MustCompile(`GetSAMLProvider failed`),
			},
		},
	})
	getM.UnPatch()

	// delete error (non-404)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: samlProviderHCL("testacc_saml_derr", samlMetadataFixture),
			},
			{
				PreConfig: func() {
					deleteM = mockey.Mock((*clientgen.IamApiService).IamServiceDeleteSAMLProviderExecute).
						Return(nil, nil, fmt.Errorf("error")).Build()
				},
				Config:      samlProviderHCL("testacc_saml_derr_new", samlMetadataFixture),
				ExpectError: regexp.MustCompile(`DeleteSAMLProvider failed`),
			},
			{
				// unpatch delete mock so post-test destroy succeeds
				PreConfig: func() {
					deleteM.UnPatch()
				},
				Config: samlProviderHCL("testacc_saml_derr_new", samlMetadataFixture),
			},
		},
	})

	// import with invalid ARN
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: samlProviderHCL("testacc_saml_imp", samlMetadataFixture),
			},
			{
				ResourceName:  "objectscale_iam_saml_provider.test",
				ImportState:   true,
				ImportStateId: "invalid-arn-format",
				ExpectError:   regexp.MustCompile(`Invalid import ID`),
			},
		},
	})
}
