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
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Test to Create and Update User Resource
func TestAccIamUserResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create user
			{
				Config: ProviderConfigForTesting + `
                resource "objectscale_iam_user" "test_user" {
                    name = "test_user"
                    namespace    = "ns1"
                    permissions_boundary_arn    = "urn:ecs:iam:::policy/ECSS3FullAccess"
                    tags = [{"key":"example_key", "value":"example_value"}]
                }
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_user.test_user", "name", "test_user"),
					resource.TestCheckResourceAttr("objectscale_iam_user.test_user", "namespace", "ns1"),
					resource.TestCheckResourceAttr("objectscale_iam_user.test_user", "permissions_boundary_arn", "urn:ecs:iam:::policy/ECSS3FullAccess"),
				),
			},
			// Step 2: Update user tags and permissions_boundary_arn
			{
				Config: ProviderConfigForTesting + `
                resource "objectscale_iam_user" "test_user" {
                    name = "test_user"
                    namespace    = "ns1"
                    permissions_boundary_arn    = "urn:ecs:iam:::policy/ECSDenyAll"
					tags = [{key = "example_key1", value = "example_value1"}]
                }
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_user.test_user", "permissions_boundary_arn", "urn:ecs:iam:::policy/ECSDenyAll"),
				),
			},
			// Step 3: Remove permissions_boundary_arn
			{
				Config: ProviderConfigForTesting + `
                resource "objectscale_iam_user" "test_user" {
                    name = "test_user"
                    namespace    = "ns1"
                    permissions_boundary_arn    = ""
					tags = [{key = "example_key1", value = "example_value1"}]
                }
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_user.test_user", "permissions_boundary_arn", ""),
				),
			},
			// Step 4: Attempt to update user name (should fail)
			{
				Config: ProviderConfigForTesting + `
				resource "objectscale_iam_user" "test_user" {
					name = "test_user_updated"
					namespace = "ns1"
				}
				`,
				ExpectError: regexp.MustCompile(".*Name is not updatable*"),
			},
			// Step 5: Attempt to untag without key (should fail)
			{

				ResourceName:  "objectscale_iam_user.test_user",
				ImportState:   true,
				ImportStateId: "invalid-format", // missing namespace
				ExpectError:   regexp.MustCompile("invalid format"),
			},
			// Step 6:import testing
			{
				ResourceName:      "objectscale_iam_user.test_user",
				ImportStateId:     "test_user:ns1",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
