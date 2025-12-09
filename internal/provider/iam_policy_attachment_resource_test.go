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

var testingManagedPolicyResourceInputParams testingInputsForIAMManagedPolicyResource

type testingInputsForIAMManagedPolicyResource struct {
	Namespace  string
	Username   string
	Groupname  string
	Rolename   string
	PolicyARN1 string
	PolicyARN2 string
	PolicyARN3 string
}

func init() {
	testingManagedPolicyResourceInputParams = testingInputsForIAMManagedPolicyResource{
		Namespace:  "ns1",
		Username:   "userTest1",
		Groupname:  "groupTest1",
		Rolename:   "roleTest1",
		PolicyARN1: "urn:ecs:iam:::policy/ECSS3ReadOnlyAccess",
		PolicyARN2: "urn:ecs:iam:::policy/IAMReadOnlyAccess",
		PolicyARN3: "urn:ecs:iam:::policy/ECSS3FullAccess",
	}
}

func TestAccIAMManagedPolicyResourceForUserCRUD(t *testing.T) {
	resourceName := "objectscale_iam_managed_policy.example"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + testAccIAMManagedPolicyResourceForInvalidUserConfig(testingManagedPolicyResourceInputParams),
				ExpectError: regexp.MustCompile("Create Error"),
			},
			{
				Config: ProviderConfigForTesting + testAccIAMManagedPolicyResourceForUserConfig1(testingManagedPolicyResourceInputParams),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "policy_arns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.0", testingManagedPolicyResourceInputParams.PolicyARN1),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.1", testingManagedPolicyResourceInputParams.PolicyARN2),
				),
			},
			{
				Config: ProviderConfigForTesting + testAccIAMManagedPolicyResourceForUserConfig2(testingManagedPolicyResourceInputParams),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "policy_arns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.0", testingManagedPolicyResourceInputParams.PolicyARN3),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.1", testingManagedPolicyResourceInputParams.PolicyARN2),
				),
			},
			{
				Config:      ProviderConfigForTesting + testAccIAMManagedPolicyResourceForInvalidUserConfig(testingManagedPolicyResourceInputParams),
				ExpectError: regexp.MustCompile("Update Error"),
			},
		},
	})
}

func TestAccIAMManagedPolicyResourceForGroupCRUD(t *testing.T) {
	resourceName := "objectscale_iam_managed_policy.example"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + testAccIAMManagedPolicyResourceForInvalidGroupConfig(testingManagedPolicyResourceInputParams),
				ExpectError: regexp.MustCompile("Create Error"),
			},
			{
				Config: ProviderConfigForTesting + testAccIAMManagedPolicyResourceForGroupConfig1(testingManagedPolicyResourceInputParams),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "policy_arns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.0", testingManagedPolicyResourceInputParams.PolicyARN1),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.1", testingManagedPolicyResourceInputParams.PolicyARN2),
				),
			},
			{
				Config: ProviderConfigForTesting + testAccIAMManagedPolicyResourceForGroupConfig2(testingManagedPolicyResourceInputParams),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "policy_arns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.0", testingManagedPolicyResourceInputParams.PolicyARN3),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.1", testingManagedPolicyResourceInputParams.PolicyARN2),
				),
			},
			{
				Config:      ProviderConfigForTesting + testAccIAMManagedPolicyResourceForInvalidGroupConfig(testingManagedPolicyResourceInputParams),
				ExpectError: regexp.MustCompile("Update Error"),
			},
		},
	})
}

func TestAccIAMManagedPolicyResourceForRoleCRUD(t *testing.T) {
	resourceName := "objectscale_iam_managed_policy.example"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + testAccIAMManagedPolicyResourceForInvalidRoleConfig(testingManagedPolicyResourceInputParams),
				ExpectError: regexp.MustCompile("Create Error"),
			},
			{
				Config: ProviderConfigForTesting + testAccIAMManagedPolicyResourceForRoleConfig1(testingManagedPolicyResourceInputParams),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "policy_arns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.0", testingManagedPolicyResourceInputParams.PolicyARN1),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.1", testingManagedPolicyResourceInputParams.PolicyARN2),
				),
			},
			{
				Config: ProviderConfigForTesting + testAccIAMManagedPolicyResourceForRoleConfig2(testingManagedPolicyResourceInputParams),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "policy_arns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.0", testingManagedPolicyResourceInputParams.PolicyARN3),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.1", testingManagedPolicyResourceInputParams.PolicyARN2),
				),
			},
			{
				Config:      ProviderConfigForTesting + testAccIAMManagedPolicyResourceForInvalidRoleConfig(testingManagedPolicyResourceInputParams),
				ExpectError: regexp.MustCompile("Update Error"),
			},
		},
	})
}

func TestAccIAMManagedPolicyResourceForErrorScenarios(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + testAccIAMManagedPolicyResourceForErrorConfig1(testingManagedPolicyResourceInputParams),
				ExpectError: regexp.MustCompile("Invalid Attribute Combination"),
			},
			{
				Config:      ProviderConfigForTesting + testAccIAMManagedPolicyResourceForErrorConfig2(testingManagedPolicyResourceInputParams),
				ExpectError: regexp.MustCompile("Invalid Attribute Combination"),
			},
		},
	})
}

func TestAccIAMManagedPolicyResourceForImport(t *testing.T) {
	resourceName := "objectscale_iam_managed_policy.example"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:        ProviderConfigForTesting + `resource "objectscale_iam_managed_policy" "example" {}`,
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: fmt.Sprintf("%s:%s", testingManagedPolicyResourceInputParams.Namespace, testingManagedPolicyResourceInputParams.Username),
				ExpectError:   regexp.MustCompile("Invalid import ID format"),
			},
			{
				Config:        ProviderConfigForTesting + `resource "objectscale_iam_managed_policy" "example" {}`,
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: fmt.Sprintf("%s:%s:%s", testingManagedPolicyResourceInputParams.Namespace, "invalid_type", testingManagedPolicyResourceInputParams.Username),
				ExpectError:   regexp.MustCompile("Invalid entity type"),
			},
			{
				Config:        ProviderConfigForTesting + `resource "objectscale_iam_managed_policy" "example" {}`,
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: fmt.Sprintf("%s:%s:%s", testingManagedPolicyResourceInputParams.Namespace, "user", testingManagedPolicyResourceInputParams.Username),
				ExpectError:   nil,
			},
			{
				Config:        ProviderConfigForTesting + `resource "objectscale_iam_managed_policy" "example" {}`,
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: fmt.Sprintf("%s:%s:%s", testingManagedPolicyResourceInputParams.Namespace, "group", testingManagedPolicyResourceInputParams.Groupname),
				ExpectError:   nil,
			},
			{
				Config:        ProviderConfigForTesting + `resource "objectscale_iam_managed_policy" "example" {}`,
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: fmt.Sprintf("%s:%s:%s", testingManagedPolicyResourceInputParams.Namespace, "role", testingManagedPolicyResourceInputParams.Rolename),
				ExpectError:   nil,
			},
		},
	})
}

func testAccIAMManagedPolicyResourceForUserConfig1(testingManagedPolicyResourceInputParams testingInputsForIAMManagedPolicyResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_managed_policy" "example" {
    namespace = "%s"
    username  = "%s"

    policy_arns = [
      "%s",
      "%s"
    ]
  }
		`,
		testingManagedPolicyResourceInputParams.Namespace,
		testingManagedPolicyResourceInputParams.Username,
		testingManagedPolicyResourceInputParams.PolicyARN1,
		testingManagedPolicyResourceInputParams.PolicyARN2,
	)
}

func testAccIAMManagedPolicyResourceForUserConfig2(testingManagedPolicyResourceInputParams testingInputsForIAMManagedPolicyResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_managed_policy" "example" {
    namespace = "%s"
    username  = "%s"

    policy_arns = [
      "%s",
      "%s"
    ]
  }
		`,
		testingManagedPolicyResourceInputParams.Namespace,
		testingManagedPolicyResourceInputParams.Username,
		testingManagedPolicyResourceInputParams.PolicyARN2,
		testingManagedPolicyResourceInputParams.PolicyARN3,
	)
}

func testAccIAMManagedPolicyResourceForGroupConfig1(testingManagedPolicyResourceInputParams testingInputsForIAMManagedPolicyResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_managed_policy" "example" {
    namespace = "%s"
    groupname  = "%s"

    policy_arns = [
      "%s",
      "%s"
    ]
  }
		`,
		testingManagedPolicyResourceInputParams.Namespace,
		testingManagedPolicyResourceInputParams.Groupname,
		testingManagedPolicyResourceInputParams.PolicyARN1,
		testingManagedPolicyResourceInputParams.PolicyARN2,
	)
}

func testAccIAMManagedPolicyResourceForGroupConfig2(testingManagedPolicyResourceInputParams testingInputsForIAMManagedPolicyResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_managed_policy" "example" {
    namespace = "%s"
    groupname  = "%s"

    policy_arns = [
      "%s",
      "%s"
    ]
  }
		`,
		testingManagedPolicyResourceInputParams.Namespace,
		testingManagedPolicyResourceInputParams.Groupname,
		testingManagedPolicyResourceInputParams.PolicyARN2,
		testingManagedPolicyResourceInputParams.PolicyARN3,
	)
}

func testAccIAMManagedPolicyResourceForRoleConfig1(testingManagedPolicyResourceInputParams testingInputsForIAMManagedPolicyResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_managed_policy" "example" {
    namespace = "%s"
    rolename  = "%s"

    policy_arns = [
      "%s",
      "%s"
    ]
  }
		`,
		testingManagedPolicyResourceInputParams.Namespace,
		testingManagedPolicyResourceInputParams.Rolename,
		testingManagedPolicyResourceInputParams.PolicyARN1,
		testingManagedPolicyResourceInputParams.PolicyARN2,
	)
}

func testAccIAMManagedPolicyResourceForRoleConfig2(testingManagedPolicyResourceInputParams testingInputsForIAMManagedPolicyResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_managed_policy" "example" {
    namespace = "%s"
    rolename  = "%s"

    policy_arns = [
      "%s",
      "%s"
    ]
  }
		`,
		testingManagedPolicyResourceInputParams.Namespace,
		testingManagedPolicyResourceInputParams.Rolename,
		testingManagedPolicyResourceInputParams.PolicyARN2,
		testingManagedPolicyResourceInputParams.PolicyARN3,
	)
}

func testAccIAMManagedPolicyResourceForErrorConfig1(testingManagedPolicyResourceInputParams testingInputsForIAMManagedPolicyResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_managed_policy" "example" {
    namespace = "%s"

    policy_arns = []
  }
		`,
		testingManagedPolicyResourceInputParams.Namespace,
	)
}

func testAccIAMManagedPolicyResourceForErrorConfig2(testingManagedPolicyResourceInputParams testingInputsForIAMManagedPolicyResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_managed_policy" "example" {
    namespace = "%s"
    username  = "%s"
    groupname  = "%s"
    rolename  = "%s"

    policy_arns = []
  }
		`,
		testingManagedPolicyResourceInputParams.Namespace,
		testingManagedPolicyResourceInputParams.Username,
		testingManagedPolicyResourceInputParams.Groupname,
		testingManagedPolicyResourceInputParams.Rolename,
	)
}

func testAccIAMManagedPolicyResourceForInvalidUserConfig(testingManagedPolicyResourceInputParams testingInputsForIAMManagedPolicyResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_managed_policy" "example" {
    namespace = "%s"
    username  = "%s"

    policy_arns = [
      "%s",
      "%s"
    ]
  }
		`,
		testingManagedPolicyResourceInputParams.Namespace,
		"INVALID_USERNAME",
		testingManagedPolicyResourceInputParams.PolicyARN1,
		testingManagedPolicyResourceInputParams.PolicyARN2,
	)
}

func testAccIAMManagedPolicyResourceForInvalidGroupConfig(testingManagedPolicyResourceInputParams testingInputsForIAMManagedPolicyResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_managed_policy" "example" {
    namespace = "%s"
    groupname  = "%s"

    policy_arns = [
      "%s",
      "%s"
    ]
  }
		`,
		testingManagedPolicyResourceInputParams.Namespace,
		"INVALID_GROUPNAME",
		testingManagedPolicyResourceInputParams.PolicyARN1,
		testingManagedPolicyResourceInputParams.PolicyARN2,
	)
}

func testAccIAMManagedPolicyResourceForInvalidRoleConfig(testingManagedPolicyResourceInputParams testingInputsForIAMManagedPolicyResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_managed_policy" "example" {
    namespace = "%s"
    rolename  = "%s"

    policy_arns = [
      "%s",
      "%s"
    ]
  }
		`,
		testingManagedPolicyResourceInputParams.Namespace,
		"INVALID_ROLENAME",
		testingManagedPolicyResourceInputParams.PolicyARN1,
		testingManagedPolicyResourceInputParams.PolicyARN2,
	)
}
