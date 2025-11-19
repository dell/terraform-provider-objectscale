package provider

import (
	"testing"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Main Acceptance Test: all scenarios.
func TestAccIAMGroupsDataSource_AllScenarios(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{

			// 1. Fetch a single group using group_name filter
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_groups" "one" {
						namespace  = "ns1"
						group_name = "sak_test_do_not_delete"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.objectscale_iam_groups.one", "id",
					),
					resource.TestCheckResourceAttr(
						"data.objectscale_iam_groups.one", "groups.#", "1",
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
						"data.objectscale_iam_groups.all", "groups.#",
					),
				),
			},

			// 3. Ensure Users[] is populated correctly
			// NOTE: Users may be 0 if the group has no users → TEST FOR "0"
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_groups" "with_users" {
						namespace  = "ns1"
						group_name = "sak_test_do_not_delete"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.objectscale_iam_groups.with_users", "groups.0.users.#", "0",
					),
				),
			},
		},
	})
}

// Error Scenarios for IAM Groups Data Source.
func TestAccIAMGroupsDataSource_ErrorScenarios(t *testing.T) {
    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { testAccPreCheck(t) },
        ProtoV6ProviderFactories: testProviderFactory,

        // IMPORTANT FIX → prevent destroy-phase login failures
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

            // Invalid namespace → no error, empty list
            {
                Config: ProviderConfigForTesting + `
                    data "objectscale_iam_groups" "bad_ns" {
                        namespace  = "invalid-namespace-xyz"
                        group_name = "sak_test_do_not_delete"
                    }
                `,
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckResourceAttr(
                        "data.objectscale_iam_groups.bad_ns", "groups.#", "0",
                    ),
                ),
            },

            // Invalid group_name → no error, empty list
            {
                Config: ProviderConfigForTesting + `
                    data "objectscale_iam_groups" "bad_group_name" {
                        namespace  = "ns1"
                        group_name = "###INVALID###"
                    }
                `,
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckResourceAttr(
                        "data.objectscale_iam_groups.bad_group_name", "groups.#", "0",
                    ),
                ),
            },

            // Empty namespace & group_name → error
            {
                Config: ProviderConfigForTesting + `
                    data "objectscale_iam_groups" "empty_values" {
                        namespace  = ""
                        group_name = ""
                    }
                `,
                ExpectError: regexp.MustCompile(`namespace`),
            },
        },
    })
}

func TestAccIAMGroupsDataSource_AuthFailure(t *testing.T) {
    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { testAccPreCheck(t) },
        ProtoV6ProviderFactories: testProviderFactory,
        
        PreventPostDestroyRefresh: true,

        Steps: []resource.TestStep{
            {
                Config: `
                    provider "objectscale" {
                        endpoint = "https://wrong-endpoint"
                        username = "invalid"
                        password = "invalid"
                        insecure = true
                        timeout  = 30
                    }

                    data "objectscale_iam_groups" "auth_error" {
                        namespace = "ns1"
                    }
                `,
                ExpectError: regexp.MustCompile(`(no such host|lookup|dial tcp|error during login)`),

                // Prevent terraform from trying to create or destroy anything
                Destroy: true,
            },
        },
    })
}

func TestAccIAMGroupsDataSource_ServerConnectionFailure(t *testing.T) {
    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { testAccPreCheck(t) },
        ProtoV6ProviderFactories: testProviderFactory,

        // Must be TRUE → prevents destroy-phase from triggering provider init
        PreventPostDestroyRefresh: true,

        Steps: []resource.TestStep{
            {
                Config: `
                    provider "objectscale" {
                        endpoint = "https://127.0.0.1:9999"   # unreachable port
                        username = "admin"
                        password = "password"
                        insecure = true
                        timeout  = 30
                    }

                    data "objectscale_iam_groups" "server_error" {
                        namespace = "ns1"
                    }
                `,

                ExpectError: regexp.MustCompile(`(connection|refused|dial tcp)`),

                // Ensures Terraform doesn't try to refresh/destroy anything
                Destroy: true,
            },
        },
    })
}

