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
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Test to Fetch Namespaces.
func TestAccIAMGroupResource(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	defer testUserTokenCleanup(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// invalid config testing
			{
				Config: ProviderConfigForTesting + `
				resource "objectscale_iam_group" "test" {
					// name is missing
					namespace = "ns1"
				}
				`,
				ExpectError: regexp.MustCompile(".*The argument \"name\" is required, but no definition was found.*"),
			},
			// create and read testing
			{
				Config: ProviderConfigForTesting + `
				resource"objectscale_iam_group" "test" {
					name = "testacc_group"
					namespace = "ns1"
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_group.test", "name", "testacc_group"),
					resource.TestCheckResourceAttr("objectscale_iam_group.test", "namespace", "ns1"),
				),
			},
			// update and testing (update not supported by API)
			{
				Config: ProviderConfigForTesting + `
				resource"objectscale_iam_group" "test" {
					name = "testacc_group_updated"
					namespace = "ns1"
				}
				`,
				ExpectError: regexp.MustCompile(".*Update operation is not supported.*"),
			},
			// import testing
			{
				// test for valid import
				ResourceName:      "objectscale_iam_group.test",
				ImportStateId:     "testacc_group:ns1",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// test for invalid import
				ResourceName:  "objectscale_iam_group.test",
				ImportStateId: "invalid_import_format",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(".*invalid format: expected 'group_name:namespace'.*"),
			},
		},
	})
}
