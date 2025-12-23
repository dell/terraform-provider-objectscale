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

func TestAccIAMInlinePolicyDataSource_PositiveScenarios(t *testing.T) {
	defer testUserTokenCleanup(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{

			// 1. Fetch inline policies for a user
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_inline_policy" "user" {
						namespace = "ns1"
						username  = "sample_user_1"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.objectscale_iam_inline_policy.user", "id",
					),
					resource.TestCheckResourceAttrSet(
						"data.objectscale_iam_inline_policy.user", "policies.0.name",
					),
					resource.TestCheckResourceAttrSet(
						"data.objectscale_iam_inline_policy.user", "policies.0.document",
					),
				),
			},

			// 2. Fetch inline policies for a group
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_inline_policy" "group" {
						namespace = "ns1"
						groupname = "group_008"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.objectscale_iam_inline_policy.group", "id",
					),
					resource.TestCheckResourceAttrSet(
						"data.objectscale_iam_inline_policy.group", "policies.0.name",
					),
					resource.TestCheckResourceAttrSet(
						"data.objectscale_iam_inline_policy.group", "policies.0.document",
					),
				),
			},

			// 3. Fetch inline policies for a role
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_inline_policy" "role" {
						namespace = "ns1"
						rolename  = "roleTest1"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.objectscale_iam_inline_policy.role", "id",
					),
					resource.TestCheckResourceAttrSet(
						"data.objectscale_iam_inline_policy.role", "policies.0.name",
					),
					resource.TestCheckResourceAttrSet(
						"data.objectscale_iam_inline_policy.role", "policies.0.document",
					),
				),
			},
		},
	})
}

func TestAccIAMInlinePolicyDataSource_NegativeScenarios(t *testing.T) {
	defer testUserTokenCleanup(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{

			// 1. More than one entity specified (username + groupname)
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_inline_policy" "invalid_multiple_entities" {
						namespace = "ns1"
						username  = "sample_user_1"
						groupname = "group_008"
					}
				`,
				ExpectError: regexp.MustCompile(
					"Exactly one of username, groupname, or rolename must be specified",
				),
			},

			// 2. No entity specified
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_inline_policy" "invalid_no_entity" {
						namespace = "ns1"
					}
				`,
				ExpectError: regexp.MustCompile(
					"Exactly one of username, groupname, or rolename must be specified",
				),
			},

			// 3. Invalid namespace
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_inline_policy" "invalid_namespace" {
						namespace = "non_existing_namespace"
						username  = "sample_user_1"
					}
				`,
				ExpectError: regexp.MustCompile(
					"Namespace does not exist or error reading IAM inline policies",
				),
			},
		},
	})
}
