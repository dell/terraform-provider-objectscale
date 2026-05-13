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

// role_Test1 is assumed to exist in the test ObjectScale cluster.

func TestAccIAMRoleDataSource_basic(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
				data "objectscale_iam_role" "all" {
					namespace = "ns1"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Make sure at least one role is returned
					resource.TestCheckResourceAttrSet("data.objectscale_iam_role.all", "roles.0.role_id"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_role.all", "roles.0.role_name"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_role.all", "roles.0.arn"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_role.all", "roles.0.path"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_role.all", "roles.0.create_date"),
				),
			},
		},
	})
}

func TestAccIAMRoleDataSource_withRoleNameFilter(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
				data "objectscale_iam_role" "by_name" {
					namespace = "ns1"
					role_name  = "roleTest1"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.objectscale_iam_role.by_name", "roles.0.role_name", "roleTest1"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_role.by_name", "roles.0.role_id"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_role.by_name", "roles.0.arn"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_role.by_name", "roles.0.path"),
				),
			},
		},
	})
}

func TestAccIAMRoleDataSource_getRole_404(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_role" "test" {
					  namespace = "ns1"
					  role_name = "missing_role_abc"
					}
				`,
				ExpectError: regexp.MustCompile(`(?i)(Error retrieving IAM role|Failed retrieving role)`),
			},
		},
	})
}

func TestAccIAMRoleDataSource_ListRole_404(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_role" "test1" {
					  namespace = "nonexistent_namespace"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.objectscale_iam_role.test1", "roles.#", "0"),
				),
			},
		},
	})
}
