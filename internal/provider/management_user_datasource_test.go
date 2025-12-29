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

func TestAccManagementUserDataSourcePositiveScenarios(t *testing.T) {
	defer testUserTokenCleanup(t)

	datasourceName := "data.objectscale_management_user.example"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// List all management users
			{
				Config: ProviderConfigForTesting + testAccManagementUserDataSourceListConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "id", "management_user_datasource"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_users.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_users.0.user_id"),
				),
			},
			// Get management user by valid name
			{
				Config: ProviderConfigForTesting + testAccManagementUserDataSourceGetValidConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "id", "management_user_datasource"),
					resource.TestCheckResourceAttr(datasourceName, "management_users.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "management_users.0.user_id", "testlocaluser1"),
				),
			},
		},
	})
}

func TestAccManagementUserDataSourceErrorScenarios(t *testing.T) {
	defer testUserTokenCleanup(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Invalid management user name
			{
				Config:      ProviderConfigForTesting + testAccManagementUserDataSourceGetInvalidConfig(),
				ExpectError: regexp.MustCompile(`Get Management User failed`),
			},
		},
	})
}

func testAccManagementUserDataSourceListConfig() string {
	return `
    data "objectscale_management_user" "example" {
    }
    `
}

func testAccManagementUserDataSourceGetValidConfig() string {
	return `
    data "objectscale_management_user" "example" {
        name = "testlocaluser1"
    }
    `
}

func testAccManagementUserDataSourceGetInvalidConfig() string {
	return `
    data "objectscale_management_user" "example" {
         name = "invalid@name"
    }
    `
}
