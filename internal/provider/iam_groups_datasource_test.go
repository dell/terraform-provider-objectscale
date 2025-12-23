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

// Main Acceptance Test: all scenarios.
func TestAccIAMGroupsDataSource_PositiveScenarios(t *testing.T) {
	defer testUserTokenCleanup(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{

			// 1. Fetch a single group using group_name filter
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_groups" "one" {
						namespace  = "ns1"
						group_name = "group_008"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.objectscale_iam_groups.one", "id",
					),
					resource.TestCheckResourceAttrSet(
						"data.objectscale_iam_groups.one", "groups.0.group_name",
					),
				),
			},

			// 2. Fetch all groups (no group_name specified)
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_groups" "all" {
						namespace = "ns1"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.objectscale_iam_groups.all", "groups.0.group_name",
					),
				),
			},

			// 3. Fetch groups for a specific user
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_groups" "user_groups" {
						namespace = "ns1"
						user_name = "user_001"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.objectscale_iam_groups.user_groups", "id",
					),
					resource.TestCheckResourceAttrSet(
						"data.objectscale_iam_groups.user_groups", "groups.0.group_name",
					),
				),
			},
		},
	})
}

// Error Scenarios for IAM Groups Data Source.
func TestAccIAMGroupsDataSource_ErrorScenarios(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,

		// IMPORTANT → prevent destroy-phase login failures
		PreventPostDestroyRefresh: true,

		Steps: []resource.TestStep{

			// Missing namespace
			{
				Config: ProviderConfigForTesting + `
                    data "objectscale_iam_groups" "missing_ns" {
                        group_name = "testgroup"
                    }
                `,
				ExpectError: regexp.MustCompile(`namespace`),
			},

			// Invalid namespace → error
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_groups" "bad_ns" {
						namespace  = "INVALID_NS"
					}
				`,
				ExpectError: regexp.MustCompile(`The namespace "INVALID_NS" does not exist.`),
			},

			// Invalid group_name → error
			{
				Config: ProviderConfigForTesting + `
                    data "objectscale_iam_groups" "bad_group_name" {
                        namespace  = "ns1"
                        group_name = "INVALID"
                    }
                `,
				ExpectError: regexp.MustCompile(`Failed to retrieve group INVALID`),
			},

			//Nonexistent user → triggers "Error Retrieving IAM Groups for User"
			{
				Config: ProviderConfigForTesting + `
                    data "objectscale_iam_groups" "invalid_user" {
                        namespace  = "ns1"
                        user_name  = "nonexistent_user_12345"
                    }
                `,
				ExpectError: regexp.MustCompile(`(?i)Error retrieving groups for user`),
			},
		},
	})
}
