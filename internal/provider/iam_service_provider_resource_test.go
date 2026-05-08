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

const spKeystoreFixture = "TU9DS19LRVlTVE9SRV9CQVNFNjQ="

func spHCL(dns string) string {
	return ProviderConfigForTesting + fmt.Sprintf(`
resource "objectscale_iam_service_provider" "sp" {
  dns           = %q
  java_keystore = %q
  key_alias     = "saml"
  key_password  = "mockpassword"
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
					resource.TestCheckResourceAttr("objectscale_iam_service_provider.sp", "etag", "etag-update"),
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
