package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Main Acceptance Test: all scenarios.
func TestAccIAMGroupsDataSource_PositiveScenarios(t *testing.T) {
	defer testUserTokenCleanup(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{

			// 1. Fetch a single group using group_name filter
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_groups" "one" {
						namespace  = "ns1"
						group_name = "group_008"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.objectscale_iam_groups.one", "id",
					),
					resource.TestCheckResourceAttrSet(
						"data.objectscale_iam_groups.one", "groups.0.group_name",
					),
				),
			},

			// 2. Fetch all groups (no group_name specified)
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_groups" "all" {
						namespace = "ns1"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.objectscale_iam_groups.all", "groups.0.group_name",
					),
				),
			},

			// 3. Fetch groups for a specific user
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_groups" "user_groups" {
						namespace = "ns1"
						user_name = "user_001"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.objectscale_iam_groups.user_groups", "id",
					),
					resource.TestCheckResourceAttrSet(
						"data.objectscale_iam_groups.user_groups", "groups.0.group_name",
					),
				),
			},
		},
	})
}

// Error Scenarios for IAM Groups Data Source.
func TestAccIAMGroupsDataSource_ErrorScenarios(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,

		// IMPORTANT → prevent destroy-phase login failures
		PreventPostDestroyRefresh: true,

		Steps: []resource.TestStep{

			// Missing namespace
			{
				Config: ProviderConfigForTesting + `
                    data "objectscale_iam_groups" "missing_ns" {
                        group_name = "testgroup"
                    }
                `,
				ExpectError: regexp.MustCompile(`namespace`),
			},

			// Invalid group_name → error
			{
				Config: ProviderConfigForTesting + `
                    data "objectscale_iam_groups" "bad_group_name" {
                        namespace  = "ns1"
                        group_name = "INVALID"
                    }
                `,
				ExpectError: regexp.MustCompile(`Failed to retrieve group INVALID`),
			},

			// Nonexistent user → triggers "Error Listing IAM Groups for User"
			{
				Config: ProviderConfigForTesting + `
                    data "objectscale_iam_groups" "invalid_user" {
                        namespace  = "ns1"
                        user_name  = "nonexistent_user_12345"
                    }
                `,
				ExpectError: regexp.MustCompile(`(?i)Error listing groups for user`),
			},
		},
	})
}
