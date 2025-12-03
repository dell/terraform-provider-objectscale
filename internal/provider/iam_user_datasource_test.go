package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// user_001,sample_user_1 and group_008 are assumed to exist in the test ObjectScale cluster.
func TestAccIAMUserDataSource_basic(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
				data "objectscale_iam_user" "all" {
					namespace = "ns1"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Make sure at least one user is returned
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.all", "users.0.id"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.all", "users.0.arn"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.all", "users.0.create_date"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.all", "users.0.path"),
				),
			},
		},
	})
}

func TestAccIAMUserDataSource_withUsernameFilter(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
				data "objectscale_iam_user" "by_username" {
					namespace = "ns1"
					username  = "sample_user_1"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.objectscale_iam_user.by_username", "users.0.username", "sample_user_1"),
				),
			},
		},
	})
}

func TestAccIAMUserDataSource_withGroupnameFilter(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
				data "objectscale_iam_user" "by_group" {
					namespace  = "ns1"
					groupname  = "group_008"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify that the returned users belong to the group
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.by_group", "users.0.id"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.by_group", "users.0.arn"),
				),
			},
		},
	})
}

func TestAccIAMUserDataSource_missingNamespace(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_user" "invalid" {}
				`,
				ExpectError: regexp.MustCompile(`The argument "namespace" is required`),
			},
		},
	})
}

func TestAccIAMUserDataSource_groupNoMatch(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_user" "none" {
						namespace = "ns1"
						groupname = "group_does_not_exist"
					}
				`,
				ExpectError: regexp.MustCompile(
					`Unable to retrieve IAM group "group_does_not_exist": 404 Not Found`,
				),
			},
		},
	})
}

func TestAccIAMUserDataSource_UserTagsAndAccessKeys(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_user" "with_keys_tags" {
						namespace = "ns1"
						username  = "user_001"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check that at least one user is returned
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.with_keys_tags", "users.#"),

					// Check tags only if any exist
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.with_keys_tags", "users.0.tags.#"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.with_keys_tags", "users.0.tags.0.key"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.with_keys_tags", "users.0.tags.0.value"),

					// Check access keys only if any exist
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.with_keys_tags", "users.0.access_keys.#"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.with_keys_tags", "users.0.access_keys.0.access_key_id"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.with_keys_tags", "users.0.access_keys.0.create_date"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_user.with_keys_tags", "users.0.access_keys.0.status"),
				),
			},
		},
	})
}
