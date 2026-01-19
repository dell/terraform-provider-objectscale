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
package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// Test to Create and Update Object User Secret Key Resource.
func testAccSecretkeyImportStateIDFunc(resourceName string, username string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}
		return rs.Primary.Attributes["id"] + ":" + username + ":ns1", nil
	}
}
func TestAccObjectUserSecretKeyResource(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + ObjectUserParams,
			},
			// Step 1: Create secret key for a user without specifying username (should fail)
			{
				Config: ProviderConfigForTesting + `
                resource "objectscale_object_user_secret_key" "test_user_secret_key" {
					depends_on = [
						objectscale_object_user.object_user_create_test
					]
                    namespace      = "ns1"
                }
				`,
				ExpectError: regexp.MustCompile(".*The argument \"username\" is required, but no definition was found.*"),
			},
			// Step 2: Create secret key
			{
				Config: ProviderConfigForTesting + ObjectUserParams + CreateObjectUserSecretKeyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_object_user_secret_key.test_user_secret_key", "username", "sample_user_ousk"),
					resource.TestCheckResourceAttr("objectscale_object_user_secret_key.test_user_secret_key", "namespace", "ns1"),
				),
			},
			// Step 3: Update secret key of the object user secret key
			{
				Config: ProviderConfigForTesting + ObjectUserParams + `
				resource "objectscale_object_user_secret_key" "test_user_secret_key" {
					depends_on = [
						objectscale_object_user.object_user_create_test
					]

					username       = "sample_user_ousk"
					namespace      = "ns1"
					secret_key	   = "abcd"
				}
				`,
				ExpectError: regexp.MustCompile(".*Update operation is not supported.*"),
			},
			// Step 4: Attempt to import with invalid format (should fail)
			{
				Config: ProviderConfigForTesting + ObjectUserParams + `
				resource "objectscale_object_user_secret_key" "test_user_secret_key" {
					depends_on = [
						objectscale_object_user.object_user_create_test
					]
				}
				`,
				ResourceName:  "objectscale_object_user_secret_key.test_user_secret_key",
				ImportState:   true,
				ImportStateId: "invalid-format",
				ExpectError:   regexp.MustCompile("invalid format"),
			},
			// Step 5: Import testing
			{
				Config: ProviderConfigForTesting + ObjectUserParams + `
				resource "objectscale_object_user_secret_key" "test_user_secret_key" {
					depends_on = [
						objectscale_object_user.object_user_create_test
					]
				}
				`,
				ResourceName:      "objectscale_object_user_secret_key.test_user_secret_key",
				ImportState:       true,
				ImportStateIdFunc: testAccSecretkeyImportStateIDFunc("objectscale_object_user_secret_key.test_user_secret_key", "sample_user_ousk"),
			},
		},
	})
}
func TestAccObjectUserSecretKey2Resource(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + ObjectUserParams2 + CreateObjectUserSecretKeyConfig2,
			},
			// Step 1: Create secret key with expiry for existing
			{
				Config: ProviderConfigForTesting + ObjectUserParams2 + CreateObjectUserSecretKeyConfig2 + CreateObjectUserSecondSecretKeyConfig2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_object_user_secret_key.test_user_secret_key2", "username", "sample_user_ousk_2"),
					resource.TestCheckResourceAttr("objectscale_object_user_secret_key.test_user_secret_key2", "namespace", "ns1"),
				),
			},
			// Step 2: Import testing
			{
				Config: ProviderConfigForTesting + ObjectUserParams2 + CreateObjectUserSecretKeyConfig2 + `
				resource "objectscale_object_user_secret_key" "test_user_second_secret_key2" {
					depends_on = [
						objectscale_object_user_secret_key.test_user_secret_key2
					]
				}
				`,
				ResourceName:      "objectscale_object_user_secret_key.test_user_second_secret_key2",
				ImportState:       true,
				ImportStateIdFunc: testAccSecretkeyImportStateIDFunc("objectscale_object_user_secret_key.test_user_second_secret_key2", "sample_user_ousk_2"),
			},
		},
	})
}

var ObjectUserParams = `
resource "objectscale_object_user" "object_user_create_test" {
	name = "sample_user_ousk"
	namespace    = "ns1"
}
`
var CreateObjectUserSecretKeyConfig = `
resource "objectscale_object_user_secret_key" "test_user_secret_key" {
	depends_on = [
		objectscale_object_user.object_user_create_test
	]

	username       = "sample_user_ousk"
	namespace      = "ns1"
}
`
var UpdateObjectUserSecretKeyConfig = `
resource "objectscale_object_user_secret_key" "test_user_secret_key" {
	depends_on = [
		objectscale_object_user.object_user_create_test
	]

	username       = "sample_user_ousk"
	namespace      = "ns1"
	secret_key	   = "abcd"
}
`

var ObjectUserParams2 = `
resource "objectscale_object_user" "object_user_create_test2" {
	name = "sample_user_ousk_2"
	namespace    = "ns1"
}
`
var CreateObjectUserSecretKeyConfig2 = `
resource "objectscale_object_user_secret_key" "test_user_secret_key2" {
	depends_on = [
		objectscale_object_user.object_user_create_test2
	]

	username       = "sample_user_ousk_2"
	namespace      = "ns1"
}
`
var CreateObjectUserSecondSecretKeyConfig2 = `
resource "objectscale_object_user_secret_key" "test_user_second_secret_key2" {
	depends_on = [
		objectscale_object_user_secret_key.test_user_secret_key2
	]

	username       = "sample_user_ousk_2"
	namespace      = "ns1"
	expiry_in_mins = "2"
}
`
