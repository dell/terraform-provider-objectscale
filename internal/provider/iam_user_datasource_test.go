// Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.
//
// Licensed under the Mozilla Public License Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// user_001,sample_user_1 and group_008 are assumed to exist in the test ObjectScale cluster.
func TestAccIAMUserDataSource_basic(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
				data "objectscale_iam_user" "all" {
					namespace = "ns1"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Make sure at least one user is returned
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.all", "users.0.id"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.all", "users.0.arn"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.all", "users.0.create_date"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.all", "users.0.path"),
				),
			},
		},
	})
}

func TestAccIAMUserDataSource_withUsernameFilter(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
				data "objectscale_iam_user" "by_username" {
					namespace = "ns1"
					username  = "sample_user_1"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.objectscale_iam_user.by_username", "users.0.username", "sample_user_1"),
				),
			},
		},
	})
}

func TestAccIAMUserDataSource_withGroupnameFilter(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
				data "objectscale_iam_user" "by_group" {
					namespace  = "ns1"
					groupname  = "group_008"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify that the returned users belong to the group
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.by_group", "users.0.id"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.by_group", "users.0.arn"),
				),
			},
		},
	})
}

func TestAccIAMUserDataSource_missingNamespace(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_user" "invalid" {}
				`,
				ExpectError: regexp.MustCompile(`The argument "namespace" is required`),
			},
		},
	})
}

func TestAccIAMUserDataSource_groupNoMatch(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_user" "none" {
						namespace = "ns1"
						groupname = "group_does_not_exist"
					}
				`,
				ExpectError: regexp.MustCompile(
					`Unable to retrieve IAM group "group_does_not_exist": 404 Not Found`,
				),
			},
		},
	})
}

func TestAccIAMUserDataSource_UserTagsAndAccessKeys(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_user" "with_keys_tags" {
						namespace = "ns1"
						username  = "user_001"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check that at least one user is returned
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.with_keys_tags", "users.#"),

					// Check tags only if any exist
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.with_keys_tags", "users.0.tags.#"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.with_keys_tags", "users.0.tags.0.key"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.with_keys_tags", "users.0.tags.0.value"),

					// Check access keys only if any exist
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.with_keys_tags", "users.0.access_keys.#"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.with_keys_tags", "users.0.access_keys.0.access_key_id"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.with_keys_tags", "users.0.access_keys.0.create_date"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.with_keys_tags", "users.0.access_keys.0.status"),
				),
			},
		},
	})
}
