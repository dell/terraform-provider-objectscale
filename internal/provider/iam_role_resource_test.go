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
			// Step 4: Update role with tags attribute
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
					"key" : "key11",
					"value" : "value11"
					},
					{
					"key" : "key22",
					"value" : "value22"
					}
				]
				}
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_role.example", "tags.0.key", "key11"),
					resource.TestCheckResourceAttr("objectscale_iam_role.example", "tags.0.value", "value11"),
					resource.TestCheckResourceAttr("objectscale_iam_role.example", "tags.1.key", "key22"),
					resource.TestCheckResourceAttr("objectscale_iam_role.example", "tags.1.value", "value22"),
				),
			},
			// Step 5: Update role with Permissions boundary attribute (fail)
			{
				Config: ProviderConfigForTesting + `
                resource "objectscale_iam_role" "example" {
				name      = "example-role"
				namespace = "ns1"
				description = "An example role updated"
				permissions_boundary_arn = "urn:ecs:iam:::policy/ECSS3DenyAll"
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
					"key" : "key11",
					"value" : "value11"
					},
					{
					"key" : "key22",
					"value" : "value22"
					}
				]
				}
                `,
				ExpectError: regexp.MustCompile(".*not found in the namespace*"),
			},
			//Step 6: Update permissions_boundary_arn
			{Config: ProviderConfigForTesting + `
                resource "objectscale_iam_role" "example" {
				name      = "example-role"
				namespace = "ns1"
				description = "An example role updated"
				permissions_boundary_arn = "urn:ecs:iam:::policy/ECSDenyAll"
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
					"key" : "key11",
					"value" : "value11"
					},
					{
					"key" : "key22",
					"value" : "value22"
					}
				]
				}
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_role.example", "permissions_boundary_arn", "urn:ecs:iam:::policy/ECSDenyAll"),
				),
			},
			//Step 6: Update permissions_boundary_arn to null
			{Config: ProviderConfigForTesting + `
                resource "objectscale_iam_role" "example" {
				name      = "example-role"
				namespace = "ns1"
				description = "An example role updated"
				permissions_boundary_arn = ""
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
					"key" : "key11",
					"value" : "value11"
					},
					{
					"key" : "key22",
					"value" : "value22"
					}
				]
				}
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_role.example", "permissions_boundary_arn", ""),
				),
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
