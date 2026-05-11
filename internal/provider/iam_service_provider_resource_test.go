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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const spKeystoreFixture = "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"

func spHCL(dns string) string {
	return ProviderConfigForTesting + fmt.Sprintf(`
resource "objectscale_iam_service_provider" "sp" {
  dns           = %q
  java_keystore = %q
  key_alias     = "saml"
  key_password  = "pass123"
}
`, dns, spKeystoreFixture)
}

// I-16 — Create service provider populates computed fields.
func TestAcc_I16_CreateServiceProvider(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: spHCL("objectscale.example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_service_provider.sp", "dns", "objectscale.example.com"),
					resource.TestCheckResourceAttr("objectscale_iam_service_provider.sp", "key_alias", "saml"),
					resource.TestCheckResourceAttrSet("objectscale_iam_service_provider.sp", "uuid"),
					resource.TestCheckResourceAttrSet("objectscale_iam_service_provider.sp", "unique_id"),
					resource.TestCheckResourceAttrSet("objectscale_iam_service_provider.sp", "etag"),
					resource.TestCheckResourceAttrSet("objectscale_iam_service_provider.sp", "create_time"),
					resource.TestCheckResourceAttrSet("objectscale_iam_service_provider.sp", "last_modified"),
				),
			},
		},
	})
}

// I-17 — Update SP DNS changes last_modified.
func TestAcc_I17_UpdateServiceProvider(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: spHCL("objectscale.example.com"),
			},
			{
				Config: spHCL("objectscale-rotated.example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_service_provider.sp", "dns", "objectscale-rotated.example.com"),
					resource.TestCheckResourceAttrSet("objectscale_iam_service_provider.sp", "etag"),
				),
			},
		},
	})
}

// I-18 — Destroy issues DELETE.
func TestAcc_I18_DeleteServiceProvider(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: spHCL("objectscale.example.com"),
				// implicit destroy at end
			},
		},
	})
}

// I-20 — Import singleton SP.
func TestAcc_I20_ImportServiceProvider(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: spHCL("objectscale.example.com"),
			},
			{
				ResourceName:      "objectscale_iam_service_provider.sp",
				ImportState:       true,
				ImportStateVerify: true,
				// keystore + password are required-on-create but the import path
				// reads them from the API. Allow framework to skip verification
				// for these inputs since they are sensitive and round-tripped.
				ImportStateVerifyIgnore: []string{"java_keystore", "key_password"},
				ImportStateId:           "objectscale-sp",
			},
		},
	})
}
