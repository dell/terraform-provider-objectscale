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

var testingPolicyAttachmentResourceInputParams testingInputsForIAMPolicyAttachmentResource

type testingInputsForIAMPolicyAttachmentResource struct {
	Namespace  string
	Username   string
	Groupname  string
	Rolename   string
	PolicyARN1 string
	PolicyARN2 string
	PolicyARN3 string
}

func init() {
	testingPolicyAttachmentResourceInputParams = testingInputsForIAMPolicyAttachmentResource{
		Namespace:  "ns1",
		Username:   "userTest1",
		Groupname:  "groupTest1",
		Rolename:   "roleTest1",
		PolicyARN1: "urn:ecs:iam:::policy/ECSS3ReadOnlyAccess",
		PolicyARN2: "urn:ecs:iam:::policy/IAMReadOnlyAccess",
		PolicyARN3: "urn:ecs:iam:::policy/ECSS3FullAccess",
	}
}

func TestAccIAMPolicyAttachmentResourceForUserCRUD(t *testing.T) {
	defer testUserTokenCleanup(t)

	resourceName := "objectscale_iam_policy_attachment.example"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + testAccIAMPolicyAttachmentResourceForInvalidUserConfig(testingPolicyAttachmentResourceInputParams),
				ExpectError: regexp.MustCompile("Create Error"),
			},
			{
				Config: ProviderConfigForTesting + testAccIAMPolicyAttachmentResourceForUserConfig1(testingPolicyAttachmentResourceInputParams),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "policy_arns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.0", testingPolicyAttachmentResourceInputParams.PolicyARN1),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.1", testingPolicyAttachmentResourceInputParams.PolicyARN2),
				),
			},
			{
				Config: ProviderConfigForTesting + testAccIAMPolicyAttachmentResourceForUserConfig2(testingPolicyAttachmentResourceInputParams),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "policy_arns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.0", testingPolicyAttachmentResourceInputParams.PolicyARN3),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.1", testingPolicyAttachmentResourceInputParams.PolicyARN2),
				),
			},
			{
				Config:      ProviderConfigForTesting + testAccIAMPolicyAttachmentResourceForInvalidUserConfig(testingPolicyAttachmentResourceInputParams),
				ExpectError: regexp.MustCompile("Update Error"),
			},
		},
	})
}

func TestAccIAMPolicyAttachmentResourceForGroupCRUD(t *testing.T) {
	defer testUserTokenCleanup(t)

	resourceName := "objectscale_iam_policy_attachment.example"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + testAccIAMPolicyAttachmentResourceForInvalidGroupConfig(testingPolicyAttachmentResourceInputParams),
				ExpectError: regexp.MustCompile("Create Error"),
			},
			{
				Config: ProviderConfigForTesting + testAccIAMPolicyAttachmentResourceForGroupConfig1(testingPolicyAttachmentResourceInputParams),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "policy_arns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.0", testingPolicyAttachmentResourceInputParams.PolicyARN1),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.1", testingPolicyAttachmentResourceInputParams.PolicyARN2),
				),
			},
			{
				Config: ProviderConfigForTesting + testAccIAMPolicyAttachmentResourceForGroupConfig2(testingPolicyAttachmentResourceInputParams),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "policy_arns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.0", testingPolicyAttachmentResourceInputParams.PolicyARN3),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.1", testingPolicyAttachmentResourceInputParams.PolicyARN2),
				),
			},
			{
				Config:      ProviderConfigForTesting + testAccIAMPolicyAttachmentResourceForInvalidGroupConfig(testingPolicyAttachmentResourceInputParams),
				ExpectError: regexp.MustCompile("Update Error"),
			},
		},
	})
}

func TestAccIAMPolicyAttachmentResourceForRoleCRUD(t *testing.T) {
	defer testUserTokenCleanup(t)

	resourceName := "objectscale_iam_policy_attachment.example"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + testAccIAMPolicyAttachmentResourceForInvalidRoleConfig(testingPolicyAttachmentResourceInputParams),
				ExpectError: regexp.MustCompile("Create Error"),
			},
			{
				Config: ProviderConfigForTesting + testAccIAMPolicyAttachmentResourceForRoleConfig1(testingPolicyAttachmentResourceInputParams),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "policy_arns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.0", testingPolicyAttachmentResourceInputParams.PolicyARN1),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.1", testingPolicyAttachmentResourceInputParams.PolicyARN2),
				),
			},
			{
				Config: ProviderConfigForTesting + testAccIAMPolicyAttachmentResourceForRoleConfig2(testingPolicyAttachmentResourceInputParams),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "policy_arns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.0", testingPolicyAttachmentResourceInputParams.PolicyARN3),
					resource.TestCheckResourceAttr(resourceName, "policy_arns.1", testingPolicyAttachmentResourceInputParams.PolicyARN2),
				),
			},
			{
				Config:      ProviderConfigForTesting + testAccIAMPolicyAttachmentResourceForInvalidRoleConfig(testingPolicyAttachmentResourceInputParams),
				ExpectError: regexp.MustCompile("Update Error"),
			},
		},
	})
}

func TestAccIAMPolicyAttachmentResourceForErrorScenarios(t *testing.T) {
	defer testUserTokenCleanup(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + testAccIAMPolicyAttachmentResourceForErrorConfig1(testingPolicyAttachmentResourceInputParams),
				ExpectError: regexp.MustCompile("Invalid Attribute Combination"),
			},
			{
				Config:      ProviderConfigForTesting + testAccIAMPolicyAttachmentResourceForErrorConfig2(testingPolicyAttachmentResourceInputParams),
				ExpectError: regexp.MustCompile("Invalid Attribute Combination"),
			},
		},
	})
}

func TestAccIAMPolicyAttachmentResourceForImport(t *testing.T) {
	defer testUserTokenCleanup(t)

	resourceName := "objectscale_iam_policy_attachment.example"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:        ProviderConfigForTesting + `resource "objectscale_iam_policy_attachment" "example" {}`,
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: fmt.Sprintf("%s:%s", testingPolicyAttachmentResourceInputParams.Namespace, testingPolicyAttachmentResourceInputParams.Username),
				ExpectError:   regexp.MustCompile("Invalid import ID format"),
			},
			{
				Config:        ProviderConfigForTesting + `resource "objectscale_iam_policy_attachment" "example" {}`,
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: fmt.Sprintf("%s:%s:%s", testingPolicyAttachmentResourceInputParams.Namespace, "invalid_type", testingPolicyAttachmentResourceInputParams.Username),
				ExpectError:   regexp.MustCompile("Invalid entity type"),
			},
			{
				Config:        ProviderConfigForTesting + `resource "objectscale_iam_policy_attachment" "example" {}`,
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: fmt.Sprintf("%s:%s:%s", testingPolicyAttachmentResourceInputParams.Namespace, "user", testingPolicyAttachmentResourceInputParams.Username),
				ExpectError:   nil,
			},
			{
				Config:        ProviderConfigForTesting + `resource "objectscale_iam_policy_attachment" "example" {}`,
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: fmt.Sprintf("%s:%s:%s", testingPolicyAttachmentResourceInputParams.Namespace, "group", testingPolicyAttachmentResourceInputParams.Groupname),
				ExpectError:   nil,
			},
			{
				Config:        ProviderConfigForTesting + `resource "objectscale_iam_policy_attachment" "example" {}`,
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: fmt.Sprintf("%s:%s:%s", testingPolicyAttachmentResourceInputParams.Namespace, "role", testingPolicyAttachmentResourceInputParams.Rolename),
				ExpectError:   nil,
			},
		},
	})
}

func testAccIAMPolicyAttachmentResourceForUserConfig1(testingPolicyAttachmentResourceInputParams testingInputsForIAMPolicyAttachmentResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_policy_attachment" "example" {
    namespace = "%s"
    username  = "%s"

    policy_arns = [
      "%s",
      "%s"
    ]
  }
		`,
		testingPolicyAttachmentResourceInputParams.Namespace,
		testingPolicyAttachmentResourceInputParams.Username,
		testingPolicyAttachmentResourceInputParams.PolicyARN1,
		testingPolicyAttachmentResourceInputParams.PolicyARN2,
	)
}

func testAccIAMPolicyAttachmentResourceForUserConfig2(testingPolicyAttachmentResourceInputParams testingInputsForIAMPolicyAttachmentResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_policy_attachment" "example" {
    namespace = "%s"
    username  = "%s"

    policy_arns = [
      "%s",
      "%s"
    ]
  }
		`,
		testingPolicyAttachmentResourceInputParams.Namespace,
		testingPolicyAttachmentResourceInputParams.Username,
		testingPolicyAttachmentResourceInputParams.PolicyARN2,
		testingPolicyAttachmentResourceInputParams.PolicyARN3,
	)
}

func testAccIAMPolicyAttachmentResourceForGroupConfig1(testingPolicyAttachmentResourceInputParams testingInputsForIAMPolicyAttachmentResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_policy_attachment" "example" {
    namespace = "%s"
    groupname  = "%s"

    policy_arns = [
      "%s",
      "%s"
    ]
  }
		`,
		testingPolicyAttachmentResourceInputParams.Namespace,
		testingPolicyAttachmentResourceInputParams.Groupname,
		testingPolicyAttachmentResourceInputParams.PolicyARN1,
		testingPolicyAttachmentResourceInputParams.PolicyARN2,
	)
}

func testAccIAMPolicyAttachmentResourceForGroupConfig2(testingPolicyAttachmentResourceInputParams testingInputsForIAMPolicyAttachmentResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_policy_attachment" "example" {
    namespace = "%s"
    groupname  = "%s"

    policy_arns = [
      "%s",
      "%s"
    ]
  }
		`,
		testingPolicyAttachmentResourceInputParams.Namespace,
		testingPolicyAttachmentResourceInputParams.Groupname,
		testingPolicyAttachmentResourceInputParams.PolicyARN2,
		testingPolicyAttachmentResourceInputParams.PolicyARN3,
	)
}

func testAccIAMPolicyAttachmentResourceForRoleConfig1(testingPolicyAttachmentResourceInputParams testingInputsForIAMPolicyAttachmentResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_policy_attachment" "example" {
    namespace = "%s"
    rolename  = "%s"

    policy_arns = [
      "%s",
      "%s"
    ]
  }
		`,
		testingPolicyAttachmentResourceInputParams.Namespace,
		testingPolicyAttachmentResourceInputParams.Rolename,
		testingPolicyAttachmentResourceInputParams.PolicyARN1,
		testingPolicyAttachmentResourceInputParams.PolicyARN2,
	)
}

func testAccIAMPolicyAttachmentResourceForRoleConfig2(testingPolicyAttachmentResourceInputParams testingInputsForIAMPolicyAttachmentResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_policy_attachment" "example" {
    namespace = "%s"
    rolename  = "%s"

    policy_arns = [
      "%s",
      "%s"
    ]
  }
		`,
		testingPolicyAttachmentResourceInputParams.Namespace,
		testingPolicyAttachmentResourceInputParams.Rolename,
		testingPolicyAttachmentResourceInputParams.PolicyARN2,
		testingPolicyAttachmentResourceInputParams.PolicyARN3,
	)
}

func testAccIAMPolicyAttachmentResourceForErrorConfig1(testingPolicyAttachmentResourceInputParams testingInputsForIAMPolicyAttachmentResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_policy_attachment" "example" {
    namespace = "%s"

    policy_arns = []
  }
		`,
		testingPolicyAttachmentResourceInputParams.Namespace,
	)
}

func testAccIAMPolicyAttachmentResourceForErrorConfig2(testingPolicyAttachmentResourceInputParams testingInputsForIAMPolicyAttachmentResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_policy_attachment" "example" {
    namespace = "%s"
    username  = "%s"
    groupname  = "%s"
    rolename  = "%s"

    policy_arns = []
  }
		`,
		testingPolicyAttachmentResourceInputParams.Namespace,
		testingPolicyAttachmentResourceInputParams.Username,
		testingPolicyAttachmentResourceInputParams.Groupname,
		testingPolicyAttachmentResourceInputParams.Rolename,
	)
}

func testAccIAMPolicyAttachmentResourceForInvalidUserConfig(testingPolicyAttachmentResourceInputParams testingInputsForIAMPolicyAttachmentResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_policy_attachment" "example" {
    namespace = "%s"
    username  = "%s"

    policy_arns = [
      "%s",
      "%s"
    ]
  }
		`,
		testingPolicyAttachmentResourceInputParams.Namespace,
		"INVALID_USERNAME",
		testingPolicyAttachmentResourceInputParams.PolicyARN1,
		testingPolicyAttachmentResourceInputParams.PolicyARN2,
	)
}

func testAccIAMPolicyAttachmentResourceForInvalidGroupConfig(testingPolicyAttachmentResourceInputParams testingInputsForIAMPolicyAttachmentResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_policy_attachment" "example" {
    namespace = "%s"
    groupname  = "%s"

    policy_arns = [
      "%s",
      "%s"
    ]
  }
		`,
		testingPolicyAttachmentResourceInputParams.Namespace,
		"INVALID_GROUPNAME",
		testingPolicyAttachmentResourceInputParams.PolicyARN1,
		testingPolicyAttachmentResourceInputParams.PolicyARN2,
	)
}

func testAccIAMPolicyAttachmentResourceForInvalidRoleConfig(testingPolicyAttachmentResourceInputParams testingInputsForIAMPolicyAttachmentResource) string {
	return fmt.Sprintf(`
	resource "objectscale_iam_policy_attachment" "example" {
    namespace = "%s"
    rolename  = "%s"

    policy_arns = [
      "%s",
      "%s"
    ]
  }
		`,
		testingPolicyAttachmentResourceInputParams.Namespace,
		"INVALID_ROLENAME",
		testingPolicyAttachmentResourceInputParams.PolicyARN1,
		testingPolicyAttachmentResourceInputParams.PolicyARN2,
	)
}
