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

	"terraform-provider-objectscale/internal/clientgen"
	"github.com/bytedance/mockey"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Test to Create and Update User Resource
func TestAccObjectUserResource(t *testing.T) {
	defer testUserTokenCleanup(t)
	var upM *mockey.Mocker
	upMConfig := func() {
		upM = mockey.Mock((*clientgen.UserManagementApiService).UserManagementServiceAddUserExecute).
			Return(nil, nil, fmt.Errorf("error")).Build()
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Fail create user operation
			{
				// mocked list policies error
				PreConfig: upMConfig,
				Config: ProviderConfigForTesting + `
				resource "objectscale_object_user" "test_user" {
			        name = "test_user"
			        namespace    = "ns1"
				}
				`,
				ExpectError: regexp.MustCompile(`Error creating user`),
			},
			// Step 2: Create user
			{
				PreConfig: func() {
					upM.UnPatch()
				},
				Config: ProviderConfigForTesting + `
                resource "objectscale_object_user" "test_user" {
                    name = "test_user"
                    namespace    = "ns1"
                    tags = [{"name":"example_name", "value":"example_value"}]
                }
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_object_user.test_user", "name", "test_user"),
					resource.TestCheckResourceAttr("objectscale_object_user.test_user", "namespace", "ns1"),
				),
			},
			// Step 3: Fail removal of user tags
			{
				PreConfig: func() {
					upM = mockey.Mock((*clientgen.UserManagementApiService).UserManagementServiceRemoveUserTagsExecute).
						Return(nil, nil, fmt.Errorf("{}")).Build()
				},
				Config: ProviderConfigForTesting + `
				resource "objectscale_object_user" "test_user" {
					name = "test_user"
			        namespace    = "ns1"
					tags = [{name = "example_name1", value = "example_value1"}]
				}
				`,
				ExpectError: regexp.MustCompile(`Error removing user tags`),
			},
			// Step 4: Fail update of user tags
			{
				PreConfig: func() {
					upM.UnPatch()
					upM = mockey.Mock((*clientgen.UserManagementApiService).UserManagementServiceUpdateUserTagExecute).
						Return(nil, nil, fmt.Errorf("{}")).Build()
				},
				Config: ProviderConfigForTesting + `
				resource "objectscale_object_user" "test_user" {
					name = "test_user"
			        namespace    = "ns1"
					tags = [{name = "example_name", value = "example_value1"}]
				}
				`,
				ExpectError: regexp.MustCompile(`Error updating user tags`),
			},
			// Step 5: Fail addition of user tags
			{
				PreConfig: func() {
					upM.UnPatch()
					upM = mockey.Mock((*clientgen.UserManagementApiService).UserManagementServiceAddUserTagExecute).
						Return(nil, nil, fmt.Errorf("{}")).Build()
				},
				Config: ProviderConfigForTesting + `
				resource "objectscale_object_user" "test_user" {
					name = "test_user"
			        namespace    = "ns1"
					tags = [{name = "example_name2", value = "example_value2"}]
				}
				`,
				ExpectError: regexp.MustCompile(`Error adding user tags`),
			},
			// Step 6: Update user tags and lock status
			{
				PreConfig: func() {
					upM.UnPatch()
				},
				Config: ProviderConfigForTesting + `
			    resource "objectscale_object_user" "test_user" {
			        name = "test_user"
			        namespace    = "ns1"
			        locked = true
					tags = [{name = "example_name1", value = "example_value1"}]
			    }
			    `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_object_user.test_user", "locked", "true"),
				),
			},
			// Step 7: Add user tags and lock users
			{
				Config: ProviderConfigForTesting + `
			    resource "objectscale_object_user" "test_user" {
			        name = "test_user"
			        namespace    = "ns1"
			        locked = false
					tags = [{name = "example_name1", value = "example_value2"}, {name = "example_name3", value = "example_value3"}]
			    }
			    `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_object_user.test_user", "locked", "false"),
				),
			},
			// Step 8: Import testing
			{
				ResourceName:      "objectscale_object_user.test_user",
				ImportStateId:     "test_user",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
