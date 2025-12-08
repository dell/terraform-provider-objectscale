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
	"os"
	"regexp"
	"terraform-provider-objectscale/internal/clientgen"
	"testing"

	"github.com/bytedance/mockey"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAcc_IamPolicyDataSource(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	defer testUserTokenCleanup(t)
	var upM *mockey.Mocker
	upMConfig := func() {
		upM = mockey.Mock((*clientgen.IamApiService).IamServiceListPoliciesExecute).
			Return(nil, nil, fmt.Errorf("error")).Build()
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// get all in namespace
				// namespace ns1 must exist
				Config: ProviderConfigForTesting + `
				data "objectscale_iam_policy" "all" {
					namespace = "ns1"
					lifecycle {
						# objectscale has some default policies so atleast one must exist
						postcondition {
							condition = length(self.policies) > 0
							error_message = "atleast one policy in ns1 must exist"
						}
					}
				}
				`,
			},
			{
				// mocked list policies error
				PreConfig: upMConfig,
				Config: ProviderConfigForTesting + `
				data "objectscale_iam_policy" "all" {
					namespace = "ns1"
				}
				`,
				ExpectError: regexp.MustCompile(`Error listing IAM policies`),
			},
			{
				// get by arn
				// atleast one policy in ns1 must exist
				PreConfig: func() {
					upM.UnPatch()
				},
				Config: ProviderConfigForTesting + `
				data "objectscale_iam_policy" "all" {
					namespace = "ns1"
				}
				data "objectscale_iam_policy" "iam_policy" {
					namespace = "ns1"
					arn = data.objectscale_iam_policy.all.policies[0].arn
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"data.objectscale_iam_policy.all",
						"policies.0.arn",
						"data.objectscale_iam_policy.iam_policy",
						"policies.0.arn"),
				),
			},
			{
				// get by invalid arn
				Config: ProviderConfigForTesting + `
				data "objectscale_iam_policy" "all" {
					namespace = "ns1"
					arn = "invalid-id"
				}
				`,
				ExpectError: regexp.MustCompile(`Error fetching IAM policies with ARN`),
			},
			{
				// get by user name
				// This works stably only if testaccpreq user is created in ns1
				// with only one policy attached to it that is IAMReadOnlyAccess
				Config: ProviderConfigForTesting + `
				data "objectscale_iam_policy" "iam_policy" {
					namespace = "ns1"
					user = "testaccpreq"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.objectscale_iam_policy.iam_policy",
						"policies.0.policy_name",
						"IAMReadOnlyAccess",
					),
				),
			},
			{
				// get by invalid user name
				Config: ProviderConfigForTesting + `
				data "objectscale_iam_policy" "all" {
					namespace = "ns1"
					user = "invalid-id"
				}
				`,
				ExpectError: regexp.MustCompile(`Error listing IAM policies attached to user`),
			},
			{
				// get by group name
				// This works stably only if testaccpreq group is created in ns1
				// with only one policy attached to it that is IAMReadOnlyAccess
				Config: ProviderConfigForTesting + `
				data "objectscale_iam_policy" "iam_policy" {
					namespace = "ns1"
					group = "testaccpreq"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.objectscale_iam_policy.iam_policy",
						"policies.0.policy_name",
						"IAMReadOnlyAccess",
					),
				),
			},
			{
				// get by invalid group name
				Config: ProviderConfigForTesting + `
				data "objectscale_iam_policy" "all" {
					namespace = "ns1"
					group = "invalid-id"
				}
				`,
				ExpectError: regexp.MustCompile(`Error listing IAM policies attached to group`),
			},
			{
				// get by role name
				// This works stably only if testaccpreq role is created in ns1
				// with only one policy attached to it that is IAMReadOnlyAccess
				Config: ProviderConfigForTesting + `
				data "objectscale_iam_policy" "iam_policy" {
					namespace = "ns1"
					role = "testaccpreq"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.objectscale_iam_policy.iam_policy",
						"policies.0.policy_name",
						"IAMReadOnlyAccess",
					),
				),
			},
			{
				// get by invalid role
				Config: ProviderConfigForTesting + `
				data "objectscale_iam_policy" "all" {
					namespace = "ns1"
					role = "invalid-id"
				}
				`,
				ExpectError: regexp.MustCompile(`Error listing IAM policies attached to role`),
			},
		},
	})
}
