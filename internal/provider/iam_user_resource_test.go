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
	"net/http"
	"regexp"
	"terraform-provider-objectscale/internal/clientgen"
	"testing"

	"github.com/bytedance/mockey"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// TestAccIamUserResource_NoBoundary verifies that creating a user without permissions_boundary_arn succeeds.
func TestAccIamUserResource_NoBoundary(t *testing.T) {
	userId := "AIDTEST00NOBOUNDARY"
	userName := "test_user_no_boundary"
	arn := "urn:ecs:iam::ns1:user/test_user_no_boundary"
	createDate := "2026-06-03T00:00:00Z"
	path := "/"

	getUserResp := &clientgen.IamServiceGetUserResponse{
		GetUserResult: &clientgen.IamServiceGetUserResponseGetUserResult{
			User: &clientgen.IamServiceGetUserResponseGetUserResultUser{
				UserId:              &userId,
				UserName:            &userName,
				Arn:                 &arn,
				CreateDate:          &createDate,
				Path:                &path,
				PermissionsBoundary: nil,
			},
		},
	}

	createM := mockey.Mock((*clientgen.IamApiService).IamServiceCreateUserExecute).
		Return(&clientgen.IamServiceCreateUserResponse{}, &http.Response{StatusCode: 200}, nil).Build()
	defer createM.UnPatch()

	getM := mockey.Mock((*clientgen.IamApiService).IamServiceGetUserExecute).
		Return(getUserResp, &http.Response{StatusCode: 200}, nil).Build()
	defer getM.UnPatch()

	deleteM := mockey.Mock((*clientgen.IamApiService).IamServiceDeleteUserExecute).
		Return(&clientgen.BasicResponse{}, &http.Response{StatusCode: 200}, nil).Build()
	defer deleteM.UnPatch()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
				resource "objectscale_iam_user" "no_boundary" {
					name      = "test_user_no_boundary"
					namespace = "ns1"
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_user.no_boundary", "name", "test_user_no_boundary"),
					resource.TestCheckResourceAttr("objectscale_iam_user.no_boundary", "namespace", "ns1"),
					resource.TestCheckResourceAttr("objectscale_iam_user.no_boundary", "id", userId),
					resource.TestCheckNoResourceAttr("objectscale_iam_user.no_boundary", "permissions_boundary_arn"),
				),
			},
		},
	})
}

// Test to Create and Update User Resource.
func TestAccIamUserResource(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create user with invalid permission boundary
			{
				Config: ProviderConfigForTesting + `
                resource "objectscale_iam_user" "test_user" {
                    name = "test_user"
                    namespace    = "ns1"
                    permissions_boundary_arn    = "urn:ecs:iam:::policy/ECSS3Access"
                    tags = [{"key":"example_key", "value":"example_value"}]
				}
				`,
				ExpectError: regexp.MustCompile(".*not found in the namespace*"),
			},
			// Step 2: Create user
			{
				Config: ProviderConfigForTesting + `
                resource "objectscale_iam_user" "test_user" {
                    name = "test_user"
                    namespace    = "ns1"
                    permissions_boundary_arn    = "urn:ecs:iam:::policy/ECSS3FullAccess"
                    tags = [{"key":"example_key", "value":"example_value"}]
                }
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_user.test_user", "name", "test_user"),
					resource.TestCheckResourceAttr("objectscale_iam_user.test_user", "namespace", "ns1"),
					resource.TestCheckResourceAttr("objectscale_iam_user.test_user", "permissions_boundary_arn", "urn:ecs:iam:::policy/ECSS3FullAccess"),
				),
			},
			// Step 3: Update user tags and permissions_boundary_arn
			{
				Config: ProviderConfigForTesting + `
                resource "objectscale_iam_user" "test_user" {
                    name = "test_user"
                    namespace    = "ns1"
                    permissions_boundary_arn    = "urn:ecs:iam:::policy/ECSDenyAll"
					tags = [{key = "example_key1", value = "example_value1"}]
                }
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_user.test_user", "permissions_boundary_arn", "urn:ecs:iam:::policy/ECSDenyAll"),
				),
			},
			// Step 4: Update user tags with invalid key (should fail)
			{
				Config: ProviderConfigForTesting + `
                resource "objectscale_iam_user" "test_user" {
                    name = "test_user"
                    namespace    = "ns1"
                    permissions_boundary_arn    = "urn:ecs:iam:::policy/ECSDenyAll"
					tags = [{key = "", value = "example_value1"}]
                }
                `,
				ExpectError: regexp.MustCompile(".*has a length 0*"),
			},
			// Step 5: Remove permissions_boundary_arn
			{
				Config: ProviderConfigForTesting + `
                resource "objectscale_iam_user" "test_user" {
                    name = "test_user"
                    namespace    = "ns1"
                    permissions_boundary_arn    = ""
					tags = [{key = "example_key1", value = "example_value1"}]
                }
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_user.test_user", "permissions_boundary_arn", ""),
				),
			},
			// Step 6: Attempt to update user name (should fail)
			{
				Config: ProviderConfigForTesting + `
				resource "objectscale_iam_user" "test_user" {
					name = "test_user_updated"
					namespace = "ns1"
				}
				`,
				ExpectError: regexp.MustCompile(".*Name is not updatable*"),
			},
			// Step 7: Attempt to update permission boundary with invalid value (should fail)
			{
				Config: ProviderConfigForTesting + `
				resource "objectscale_iam_user" "test_user" {
                    name = "test_user"
                    namespace    = "ns1"
                    permissions_boundary_arn    = "urn:ecs:iam:::policy/ECSDeny"
					tags = [{key = "example_key1", value = "example_value1"}]
				}
				`,
				ExpectError: regexp.MustCompile(".*not found in the namespace*"),
			},
			// Step 8: Attempt to import with invalid format (should fail)
			{

				ResourceName:  "objectscale_iam_user.test_user",
				ImportState:   true,
				ImportStateId: "invalid-format", // missing namespace
				ExpectError:   regexp.MustCompile("invalid format"),
			},
			// Step 9:import testing
			{
				ResourceName:      "objectscale_iam_user.test_user",
				ImportStateId:     "test_user:ns1",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
