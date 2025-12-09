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

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Test to Create and Update User Resource
func TestAccIamRoleResource(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create role with invalid permission boundary
			{
				Config: ProviderConfigForTesting + `
                resource "objectscale_iam_role" "example" {
				name      = "example-role"
				namespace = "ns1"
				description = "An example role"
				permissions_boundary_arn = "urn:ecs:iam:::policy/ECSS3Access"
				max_session_duration = 4000
				assume_role_policy_document = jsonencode({
					Version = "2012-11-17"
					Statement = [
					{
						Effect = "Allow"
						Principal = {
						AWS = [
							"urn:ecs:iam::ns1:user/sample_user_1"
						]
						}
						Action = "sts:AssumeRole"
					}
					]
				})
				tags = [
					{
					"key" : "key1",
					"value" : "value1"
					},
					{
					"key" : "key2",
					"value" : "value2"
					}
				]
				}
				`,
				ExpectError: regexp.MustCompile(".*not found in the namespace*"),
			},
			// Step 2: Create role
			{
				Config: ProviderConfigForTesting + `
                resource "objectscale_iam_role" "example" {
				name      = "example-role"
				namespace = "ns1"
				description = "An example role"
				permissions_boundary_arn = "urn:ecs:iam:::policy/ECSS3FullAccess"
				max_session_duration = 4000
				assume_role_policy_document = jsonencode({
					Version = "2012-11-17"
					Statement = [
					{
						Effect = "Allow"
						Principal = {
						AWS = [
							"urn:ecs:iam::ns1:user/sample_user_1"
						]
						}
						Action = "sts:AssumeRole"
					}
					]
				})
				tags = [
					{
					"key" : "key1",
					"value" : "value1"
					},
					{
					"key" : "key2",
					"value" : "value2"
					}
				]
				}
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_role.example", "name", "example-role"),
					resource.TestCheckResourceAttr("objectscale_iam_role.example", "namespace", "ns1"),
					resource.TestCheckResourceAttr("objectscale_iam_role.example", "permissions_boundary_arn", "urn:ecs:iam:::policy/ECSS3FullAccess"),
				),
			},

			// Step 3: Update description and max_session_duration
			{
				Config: ProviderConfigForTesting + `
                resource "objectscale_iam_role" "example" {
				name      = "example-role"
				namespace = "ns1"
				description = "An example role updated"
				permissions_boundary_arn = "urn:ecs:iam:::policy/ECSS3FullAccess"
				max_session_duration = 3600
				assume_role_policy_document = jsonencode({
					Version = "2012-11-17"
					Statement = [
					{
						Effect = "Allow"
						Principal = {
						AWS = [
							"urn:ecs:iam::ns1:user/sample_user_1"
						]
						}
						Action = "sts:AssumeRole"
					}
					]
				})
				tags = [
					{
					"key" : "key1",
					"value" : "value1"
					},
					{
					"key" : "key2",
					"value" : "value2"
					}
				]
				}
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_role.example", "description", "An example role updated"),
					resource.TestCheckResourceAttr("objectscale_iam_role.example", "max_session_duration", "3600"),
				),
			},
			// Step 4: Update role with invalid attribute (should fail)
			{
				Config: ProviderConfigForTesting + `
                resource "objectscale_iam_role" "example" {
				name      = "example-role"
				namespace = "ns1"
				description = "An example role updated"
				permissions_boundary_arn = "urn:ecs:iam:::policy/ECSS3DenyAccess"
				max_session_duration = 3600
				assume_role_policy_document = jsonencode({
					Version = "2012-11-17"
					Statement = [
					{
						Effect = "Allow"
						Principal = {
						AWS = [
							"urn:ecs:iam::ns1:user/sample_user_1"
						]
						}
						Action = "sts:AssumeRole"
					}
					]
				})
				tags = [
					{
					"key" : "key1",
					"value" : "value1"
					},
					{
					"key" : "key2",
					"value" : "value2"
					}
				]
				}
                `,
				ExpectError: regexp.MustCompile(".*invalid attribute change detected*"),
			},
			// Step 5: Attempt to import with invalid format (should fail)
			{

				ResourceName:  "objectscale_iam_role.example",
				ImportState:   true,
				ImportStateId: "invalid-format", // missing namespace
				ExpectError:   regexp.MustCompile("invalid format"),
			},
			// Step 6:import testing
			{
				ResourceName:                         "objectscale_iam_role.example",
				ImportStateId:                        "example-role:ns1",
				ImportState:                          true,
				ImportStateVerifyIdentifierAttribute: "name",
			},
		},
	})
}
