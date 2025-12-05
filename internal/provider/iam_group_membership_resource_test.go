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
	"context"
	"fmt"
	"os"
	"regexp"
	"terraform-provider-objectscale/internal/clientgen"
	"testing"

	"github.com/bytedance/mockey"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	resourceconfig "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Test to Fetch Namespaces
func TestAccIAMGroupMembershipResource(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// invalid config testing
			{
				Config: ProviderConfigForTesting + `
				resource "objectscale_iam_group" "example" {
					name      = "example-group"
					namespace = "ns1"
					}
					
					resource "objectscale_iam_group_membership" "example_membership" {
					name      = objectscale_iam_group.example.name
					namespace = objectscale_iam_group.example.namespace
					// user is missing
				}
				`,
				ExpectError: regexp.MustCompile(".*The argument \"user\" is required, but no definition was found..*"),
			},
			{
				Config: ProviderConfigForTesting + `
				resource "objectscale_iam_group" "example" {
					name      = "example-group"
					namespace = "ns1"
					}
					
					resource "objectscale_iam_group_membership" "example_membership" {
						name      = objectscale_iam_group.example.name
						namespace = objectscale_iam_group.example.namespace
						user      = "testacc-user"
						}
						`,
				ExpectError: regexp.MustCompile(".*NoSuchEntity.*"),
			},
			{
				Config: ProviderConfigForTesting + `
				resource "objectscale_iam_group" "example" {
					name      = "example-group"
					namespace = "ns1"
				}
		
				resource "objectscale_iam_group_membership" "example_membership" {
					name      = objectscale_iam_group.example.name
					namespace = objectscale_iam_group.example.namespace
					user      = "test-user"
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_group_membership.example_membership", "name", "example-group"),
					resource.TestCheckResourceAttr("objectscale_iam_group_membership.example_membership", "namespace", "ns1"),
					resource.TestCheckResourceAttr("objectscale_iam_group_membership.example_membership", "user", "test-user"),
				),
			},
			{
				Config: ProviderConfigForTesting + `
				resource "objectscale_iam_group" "example" {
					name      = "example-group"
					namespace = "ns1"
				}
		
				resource "objectscale_iam_group_membership" "example_membership" {
					name      = objectscale_iam_group.example.name
					namespace = objectscale_iam_group.example.namespace
					user      = "test-user1"
				}
				`,
				ExpectError: regexp.MustCompile(".*Update operation is not supported.*"),
			},
			{
				// test for import
				ResourceName: "objectscale_iam_group_membership.example_membership",
				ImportState:  true,
				ExpectError:  regexp.MustCompile(".*Import operation is not available.*"),
			},
		},
	})
}

func TestAccIAMGroupMembershipResource_InvalidAPIClient(t *testing.T) {

	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	var apiMocker, api1mocker *mockey.Mocker

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					apiMocker = mockey.Mock((*clientgen.IamApiService).IamServiceAddUserToGroupExecute).
						Return(nil, nil, fmt.Errorf("error")).
						Build()
				},
				Config: ProviderConfigForTesting + `
				resource "objectscale_iam_group_membership" "example_membership" {
					name      = "mocked_group"
					namespace = "ns1"
					user      = "mocked_user"
				}
				`,
				ExpectError: regexp.MustCompile(".*Error adding user to group.*"),
			},
			{
				PreConfig: func() {
					apiMocker.UnPatch() // cleanup after the previous step
					apiMocker = mockey.Mock((*clientgen.IamApiService).IamServiceAddUserToGroupExecute).
						Return(nil, nil, nil).
						Build()
					api1mocker = mockey.Mock((*clientgen.IamApiService).IamServiceGetGroupExecute).
						Return(nil, nil, fmt.Errorf("error")).
						Build()
				},
				Config: ProviderConfigForTesting + `
				resource "objectscale_iam_group_membership" "example_membership" {
					name      = "mocked_group"
					namespace = "ns1"
					user      = "mocked_user"
				}
				`,
				ExpectError: regexp.MustCompile(".*Error reading Group.*"),
			},
			{
				PreConfig: func() {
					apiMocker.UnPatch() // cleanup after the previous step
					apiMocker = mockey.Mock((*clientgen.IamApiService).IamServiceRemoveUserFromGroupExecute).
						Return(nil, nil, fmt.Errorf("error")).
						Build()
				},
				Config: ProviderConfigForTesting + `
				resource "objectscale_iam_group_membership" "example_membership" {
					name      = "mocked_group"
					namespace = "ns1"
					user      = "mocked_user"
				}
				`,
				Destroy:     true,
				ExpectError: regexp.MustCompile(".*Remove user failed.*"),
			},
		},
	})

	api1mocker.UnPatch()
	apiMocker.UnPatch()
}

func TestAccIAMGroupMembershipResource_Configure_InvalidType(t *testing.T) {
	ctx := context.Background()

	r := &IAMGroupMembershipResource{}

	req := resourceconfig.ConfigureRequest{
		ProviderData: 12345, // wrong type to trigger the !ok branch
	}
	resp := &resourceconfig.ConfigureResponse{
		Diagnostics: diag.Diagnostics{},
	}

	r.Configure(ctx, req, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatalf("expected diagnostics error when ProviderData type is invalid, got none")
	}

	got := resp.Diagnostics[0].Summary() // ✅ Summary is a method in your version
	want := "Unexpected Resource Configure Type"
	if got != want {
		t.Errorf("unexpected diagnostic summary: got %q, want %q", got, want)
	}
}
