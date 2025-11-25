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

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

var testingInputParams testingInputsForIAMInlinePolicyResource

type testingInputsForIAMInlinePolicyResource struct {
	Namespace              string
	Username               string
	Groupname              string
	Rolename               string
	PolicyName1            string
	PolicyName2            string
	PolicyName3            string
	PolicyDocument1        string
	PolicyDocument2        string
	PolicyDocument2Updated string
	PolicyDocument3        string
}

func init() {
	testingInputParams = testingInputsForIAMInlinePolicyResource{
		Namespace:   "ns1",
		Username:    "userTest1",
		Groupname:   "groupTest1",
		Rolename:    "roleTest1",
		PolicyName1: "inlinePolicyTest1",
		PolicyName2: "inlinePolicyTest2",
		PolicyName3: "inlinePolicyTest3",
		PolicyDocument1: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "VisualEditor0",
      "Effect": "Allow",
      "Action": [
        "iam:GetPolicyVersion",
        "iam:GetUser",
        "iam:GetPolicy",
        "iam:GetGroupPolicy",
        "iam:GetRole",
        "iam:GetAccessKeyLastUsed",
        "iam:GetGroup",
        "iam:GetUserPolicy",
        "iam:GetSAMLProvider",
        "iam:GetRolePolicy",
        "iam:GetContextKeysForCustomPolicy",
        "iam:GetContextKeysForPrincipalPolicy",
        "iam:SimulateCustomPolicy",
        "iam:SimulatePrincipalPolicy"
      ],
      "Resource": "*"
    }
  ]
}`,
		PolicyDocument2: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "VisualEditor0",
      "Effect": "Allow",
      "Action": [
        "iam:DeleteAccessKey",
        "iam:UpdateSAMLProvider",
        "iam:CreateRole",
        "iam:RemoveUserFromGroup",
        "iam:AddUserToGroup",
        "iam:UpdateUser",
        "iam:CreateAccessKey",
        "iam:UpdateAccessKey",
        "iam:CreateSAMLProvider",
        "iam:DeleteRole",
        "iam:UpdateRole",
        "iam:DeleteGroup",
        "iam:UpdateGroup",
        "iam:CreateUser",
        "iam:CreateGroup",
        "iam:DeleteSAMLProvider",
        "iam:DeleteUser"
      ],
      "Resource": "*"
    }
  ]
}`,
		PolicyDocument2Updated: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "VisualEditor0",
      "Effect": "Allow",
      "Action": [
        "iam:GetPolicyVersion",
        "iam:GetUser",
        "iam:GetPolicy",
        "iam:GetGroupPolicy",
        "iam:GetRole",
        "iam:GetAccessKeyLastUsed",
        "iam:GetGroup",
        "iam:GetUserPolicy",
        "iam:GetSAMLProvider",
        "iam:GetRolePolicy",
        "iam:GetContextKeysForCustomPolicy",
        "iam:GetContextKeysForPrincipalPolicy",
        "iam:SimulateCustomPolicy",
        "iam:SimulatePrincipalPolicy",
        "iam:DeleteAccessKey",
        "iam:UpdateSAMLProvider",
        "iam:CreateRole",
        "iam:RemoveUserFromGroup",
        "iam:AddUserToGroup",
        "iam:UpdateUser",
        "iam:CreateAccessKey",
        "iam:UpdateAccessKey",
        "iam:CreateSAMLProvider",
        "iam:DeleteRole",
        "iam:UpdateRole",
        "iam:DeleteGroup",
        "iam:UpdateGroup",
        "iam:CreateUser",
        "iam:CreateGroup",
        "iam:DeleteSAMLProvider",
        "iam:DeleteUser"
      ],
      "Resource": "*"
    }
  ]
}`,
		PolicyDocument3: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "VisualEditor0",
      "Effect": "Allow",
      "Action": [
        "iam:TagUser",
        "iam:TagRole",
        "iam:UntagUser",
        "iam:UntagRole"
      ],
      "Resource": "*"
    }
  ]
}`,
	}
}

func TestAccIAMInlinePolicyForUserCRUD(t *testing.T) {
	resourceName := "objectscale_iam_inline_policy.example"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + testAccIAMInlinePolicyForUserConfig1(testingInputParams),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "policies.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.name", testingInputParams.PolicyName1),
					resource.TestCheckResourceAttr(resourceName, "policies.1.name", testingInputParams.PolicyName2),
				),
			},
			{
				Config: ProviderConfigForTesting + testAccIAMInlinePolicyForUserConfig2(testingInputParams),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "policies.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.name", testingInputParams.PolicyName2),
					resource.TestCheckResourceAttr(resourceName, "policies.1.name", testingInputParams.PolicyName3),
				),
			},
		},
	})
}

func TestAccIAMInlinePolicyForGroupCRUD(t *testing.T) {
	resourceName := "objectscale_iam_inline_policy.example"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + testAccIAMInlinePolicyForGroupConfig1(testingInputParams),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "policies.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.name", testingInputParams.PolicyName1),
					resource.TestCheckResourceAttr(resourceName, "policies.1.name", testingInputParams.PolicyName2),
				),
			},
			{
				Config: ProviderConfigForTesting + testAccIAMInlinePolicyForGroupConfig2(testingInputParams),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "policies.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.name", testingInputParams.PolicyName2),
					resource.TestCheckResourceAttr(resourceName, "policies.1.name", testingInputParams.PolicyName3),
				),
			},
		},
	})
}

func TestAccIAMInlinePolicyForRoleCRUD(t *testing.T) {
	resourceName := "objectscale_iam_inline_policy.example"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + testAccIAMInlinePolicyForRoleConfig1(testingInputParams),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "policies.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.name", testingInputParams.PolicyName1),
					resource.TestCheckResourceAttr(resourceName, "policies.1.name", testingInputParams.PolicyName2),
				),
			},
			{
				Config: ProviderConfigForTesting + testAccIAMInlinePolicyForRoleConfig2(testingInputParams),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "policies.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.name", testingInputParams.PolicyName2),
					resource.TestCheckResourceAttr(resourceName, "policies.1.name", testingInputParams.PolicyName3),
				),
			},
		},
	})
}

func TestAccIAMInlinePolicyErrorScenarios(t *testing.T) {
	resourceName := "objectscale_iam_inline_policy.example"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + testAccIAMInlinePolicyErrorConfig1(testingInputParams),
				ExpectError: regexp.MustCompile("Validation Error"),
			},
			{
				Config:      ProviderConfigForTesting + testAccIAMInlinePolicyErrorConfig2(testingInputParams),
				ExpectError: regexp.MustCompile("Validation Error"),
			},
			{
				Config:        ProviderConfigForTesting + `resource "objectscale_iam_inline_policy" "example" {}`,
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: fmt.Sprintf("%s:%s", testingInputParams.Namespace, testingInputParams.Username),
				ExpectError:   regexp.MustCompile("Invalid import ID format"),
			},
		},
	})
}

func TestAccIAMInlinePolicyImport(t *testing.T) {
	resourceName := "objectscale_iam_inline_policy.example"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:        ProviderConfigForTesting + `resource "objectscale_iam_inline_policy" "example" {}`,
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: fmt.Sprintf("%s:%s:%s", testingInputParams.Namespace, "user", testingInputParams.Username),
				ExpectError:   nil,
			},
			{
				Config:        ProviderConfigForTesting + `resource "objectscale_iam_inline_policy" "example" {}`,
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: fmt.Sprintf("%s:%s:%s", testingInputParams.Namespace, "group", testingInputParams.Groupname),
				ExpectError:   nil,
			},
			{
				Config:        ProviderConfigForTesting + `resource "objectscale_iam_inline_policy" "example" {}`,
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: fmt.Sprintf("%s:%s:%s", testingInputParams.Namespace, "role", testingInputParams.Rolename),
				ExpectError:   nil,
			},
		},
	})
}

func testAccIAMInlinePolicyForUserConfig1(testingInputParams testingInputsForIAMInlinePolicyResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_inline_policy" "example" {
    namespace = "%s"
    username  = "%s"

    policies = [
      {
        name     = "%s"
        document = <<EOT
%s
EOT
      },
      {
        name     = "%s"
        document = <<EOT
%s
EOT
      }
    ]
  }
		`,
		testingInputParams.Namespace,
		testingInputParams.Username,
		testingInputParams.PolicyName1,
		testingInputParams.PolicyDocument1,
		testingInputParams.PolicyName2,
		testingInputParams.PolicyDocument2,
	)
}

func testAccIAMInlinePolicyForUserConfig2(testingInputParams testingInputsForIAMInlinePolicyResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_inline_policy" "example" {
    namespace = "%s"
    username  = "%s"

    policies = [
      {
        name     = "%s"
        document = <<EOT
%s
EOT
      },
      {
        name     = "%s"
        document = <<EOT
%s
EOT
      }
    ]
  }
		`,
		testingInputParams.Namespace,
		testingInputParams.Username,
		testingInputParams.PolicyName2,
		testingInputParams.PolicyDocument2Updated,
		testingInputParams.PolicyName3,
		testingInputParams.PolicyDocument3,
	)
}

func testAccIAMInlinePolicyForGroupConfig1(testingInputParams testingInputsForIAMInlinePolicyResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_inline_policy" "example" {
    namespace = "%s"
    groupname  = "%s"

    policies = [
      {
        name     = "%s"
        document = <<EOT
%s
EOT
      },
      {
        name     = "%s"
        document = <<EOT
%s
EOT
      }
    ]
  }
		`,
		testingInputParams.Namespace,
		testingInputParams.Groupname,
		testingInputParams.PolicyName1,
		testingInputParams.PolicyDocument1,
		testingInputParams.PolicyName2,
		testingInputParams.PolicyDocument2,
	)
}

func testAccIAMInlinePolicyForGroupConfig2(testingInputParams testingInputsForIAMInlinePolicyResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_inline_policy" "example" {
    namespace = "%s"
    groupname  = "%s"

    policies = [
      {
        name     = "%s"
        document = <<EOT
%s
EOT
      },
      {
        name     = "%s"
        document = <<EOT
%s
EOT
      }
    ]
  }
		`,
		testingInputParams.Namespace,
		testingInputParams.Groupname,
		testingInputParams.PolicyName2,
		testingInputParams.PolicyDocument2Updated,
		testingInputParams.PolicyName3,
		testingInputParams.PolicyDocument3,
	)
}

func testAccIAMInlinePolicyForRoleConfig1(testingInputParams testingInputsForIAMInlinePolicyResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_inline_policy" "example" {
    namespace = "%s"
    rolename  = "%s"

    policies = [
      {
        name     = "%s"
        document = <<EOT
%s
EOT
      },
      {
        name     = "%s"
        document = <<EOT
%s
EOT
      }
    ]
  }
		`,
		testingInputParams.Namespace,
		testingInputParams.Rolename,
		testingInputParams.PolicyName1,
		testingInputParams.PolicyDocument1,
		testingInputParams.PolicyName2,
		testingInputParams.PolicyDocument2,
	)
}

func testAccIAMInlinePolicyForRoleConfig2(testingInputParams testingInputsForIAMInlinePolicyResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_inline_policy" "example" {
    namespace = "%s"
    rolename  = "%s"

    policies = [
      {
        name     = "%s"
        document = <<EOT
%s
EOT
      },
      {
        name     = "%s"
        document = <<EOT
%s
EOT
      }
    ]
  }
		`,
		testingInputParams.Namespace,
		testingInputParams.Rolename,
		testingInputParams.PolicyName2,
		testingInputParams.PolicyDocument2Updated,
		testingInputParams.PolicyName3,
		testingInputParams.PolicyDocument3,
	)
}

func testAccIAMInlinePolicyErrorConfig1(testingInputParams testingInputsForIAMInlinePolicyResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_inline_policy" "example" {
    namespace = "%s"

    policies = []
  }
		`,
		testingInputParams.Namespace,
	)
}

func testAccIAMInlinePolicyErrorConfig2(testingInputParams testingInputsForIAMInlinePolicyResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_inline_policy" "example" {
    namespace = "%s"
    username  = "%s"
    groupname  = "%s"
    rolename  = "%s"

    policies = []
  }
		`,
		testingInputParams.Namespace,
		testingInputParams.Username,
		testingInputParams.Groupname,
		testingInputParams.Rolename,
	)
}
