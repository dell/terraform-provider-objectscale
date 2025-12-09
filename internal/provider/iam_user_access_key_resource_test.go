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

// Test to Create and Update IAM User Access Key Resource
func testAccImportStateIDFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}
		return rs.Primary.Attributes["id"] + ":sample_user_1:ns1", nil
	}
}
func TestAccIamUserAccessKeyResource(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create accesss key for a user without specifying username (should fail)
			{
				Config: ProviderConfigForTesting + `
                resource "objectscale_iam_user_access_key" "test_user_access_key" {
                    namespace    = "ns1"
                }
				`,
				ExpectError: regexp.MustCompile(".*The argument \"username\" is required, but no definition was found.*"),
			},
			// Step 2: Create accesss key
			{
				Config: ProviderConfigForTesting + `
                resource "objectscale_iam_user_access_key" "test_user_access_key" {
                    username     = "sample_user_1"
                    namespace    = "ns1"
                }
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_user_access_key.test_user_access_key", "username", "sample_user_1"),
					resource.TestCheckResourceAttr("objectscale_iam_user_access_key.test_user_access_key", "namespace", "ns1"),
				),
			},
			// Step 3: Update state of the user access key to Inactive
			{
				Config: ProviderConfigForTesting + `
			    resource "objectscale_iam_user_access_key" "test_user_access_key" {
			        username     = "sample_user_1"
			        namespace    = "ns1"
					status	     = "Inactive"
			    }
			    `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_user_access_key.test_user_access_key", "username", "sample_user_1"),
					resource.TestCheckResourceAttr("objectscale_iam_user_access_key.test_user_access_key", "namespace", "ns1"),
					resource.TestCheckResourceAttr("objectscale_iam_user_access_key.test_user_access_key", "status", "Inactive"),
				),
			},
			// Step 4: Attempt to import with invalid format (should fail)
			{

				ResourceName:  "objectscale_iam_user_access_key.test_user_access_key",
				ImportState:   true,
				ImportStateId: "invalid-format", // missing namespace
				ExpectError:   regexp.MustCompile("invalid format"),
			},
			// Step 5:import testing
			{
				ResourceName:      "objectscale_iam_user_access_key.test_user_access_key",
				ImportState:       true,
				ImportStateIdFunc: testAccImportStateIDFunc("objectscale_iam_user_access_key.test_user_access_key"),
			},
		},
	})
}
