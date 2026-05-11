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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// I-21 — Read singleton SP via datasource.
func TestAcc_I21_DataSource_ReadServiceProvider(t *testing.T) {
	defer testUserTokenCleanup(t)
	cfg := ProviderConfigForTesting + `
resource "objectscale_iam_service_provider" "sp" {
  dns           = "objectscale-ds-test.example.com"
  java_keystore = "` + spKeystoreFixture + `"
  key_alias     = "saml"
  key_password  = "pass123"
}
data "objectscale_iam_service_provider" "read" {
  depends_on = [objectscale_iam_service_provider.sp]
}
`
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: cfg,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.objectscale_iam_service_provider.read", "id", "objectscale-sp"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_service_provider.read", "dns"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_service_provider.read", "uuid"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_service_provider.read", "unique_id"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_service_provider.read", "etag"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_service_provider.read", "key_alias"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_service_provider.read", "create_time"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_service_provider.read", "last_modified"),
				),
			},
		},
	})
}

// I-22 — Read SP metadata via datasource.
func TestAcc_I22_DataSource_ReadServiceProviderMetadata(t *testing.T) {
	defer testUserTokenCleanup(t)
	cfg := ProviderConfigForTesting + `
resource "objectscale_iam_service_provider" "sp" {
  dns           = "objectscale-meta-test.example.com"
  java_keystore = "` + spKeystoreFixture + `"
  key_alias     = "saml"
  key_password  = "pass123"
}
data "objectscale_iam_service_provider_metadata" "meta" {
  depends_on = [objectscale_iam_service_provider.sp]
}
`
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: cfg,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.objectscale_iam_service_provider_metadata.meta", "id", "objectscale-sp-metadata"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_service_provider_metadata.meta", "metadata_xml"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_service_provider_metadata.meta", "entity_id"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_service_provider_metadata.meta", "acs_url"),
				),
			},
		},
	})
}
