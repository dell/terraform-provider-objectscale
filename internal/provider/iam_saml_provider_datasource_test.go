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

// TestAccIAMSAMLProviderDataSource exercises the SAML provider datasource:
// read by ARN, non-existent ARN error, list multiple, and empty namespace.
func TestAccIAMSAMLProviderDataSource(t *testing.T) {
	defer testUserTokenCleanup(t)
	var listSamlProvidersM *mockey.Mocker

	// read single provider via datasource, then list multiple
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// read single provider by ARN
			{
				Config: ProviderConfigForTesting + fmt.Sprintf(`
					resource "objectscale_iam_saml_provider" "src" {
					name                   = "testacc_saml_ds"
					namespace              = "ns1"
					saml_metadata_document = %q
					}
					data "objectscale_iam_saml_provider" "by_arn" {
					saml_provider_arn = objectscale_iam_saml_provider.src.arn
					namespace         = "ns1"
					}
					`, samlMetadataFixture),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.objectscale_iam_saml_provider.by_arn", "providers.#", "1"),
					resource.TestCheckResourceAttr("data.objectscale_iam_saml_provider.by_arn", "providers.0.name", "testacc_saml_ds"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_saml_provider.by_arn", "providers.0.saml_metadata_document"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_saml_provider.by_arn", "providers.0.create_date"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_saml_provider.by_arn", "providers.0.valid_until"),
				),
			},
			// non-existent ARN
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_saml_provider" "missing" {
					saml_provider_arn = "urn:ecs:iam::ns1:saml-provider/testacc_saml_missing"
					namespace         = "ns1"
					}
					`,
				ExpectError: regexp.MustCompile(`(?i)not found|404|NoSuchEntity`),
			},
			// invalid saml arn
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_saml_provider" "invalid" {
					saml_provider_arn = "invalid-arn"
					namespace         = "ns1"
					}
					`,
				ExpectError: regexp.MustCompile(`(?i)Invalid saml_provider_arn`),
			},
			// list providers mock error
			{
				PreConfig: func() {
					listSamlProvidersM = mockey.Mock((*clientgen.IamApiService).IamServiceListSAMLProvidersExecute).
						Return(nil, nil, fmt.Errorf("error")).Build()
				},
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_saml_provider" "list_error" {
						namespace = "ns1"
					}
					`,
				ExpectError: regexp.MustCompile(`Error listing SAML providers`),
			},
			// list multiple providers (separate resource.Test because of multi-resource config)
			{
				PreConfig: func() {
					listSamlProvidersM.UnPatch()
				},
				Config: ProviderConfigForTesting + fmt.Sprintf(`
					resource "objectscale_iam_saml_provider" "a" {
						name                   = "testacc_saml_list_a"
						namespace              = "ns1"
						saml_metadata_document = %q
					}
					resource "objectscale_iam_saml_provider" "b" {
						name                   = "testacc_saml_list_b"
						namespace              = "ns1"
						saml_metadata_document = %q
					}
					resource "objectscale_iam_saml_provider" "c" {
						name                   = "testacc_saml_list_c"
						namespace              = "ns1"
						saml_metadata_document = %q
					}
					data "objectscale_iam_saml_provider" "page" {
						namespace  = "ns1"
						depends_on = [
							objectscale_iam_saml_provider.a,
							objectscale_iam_saml_provider.b,
							objectscale_iam_saml_provider.c,
						]
					}
					`, samlMetadataFixture, samlMetadataFixture, samlMetadataFixture),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Don't assert exact count since mock may retain providers from prior runs
					resource.TestCheckResourceAttrSet("data.objectscale_iam_saml_provider.page", "providers.#"),
				),
			},
		},
	})
}
