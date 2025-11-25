package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Test to Create and Update User Resource
func TestAccIamUserResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create user
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
			// Step 2: Update user role
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
		},
	})
}
