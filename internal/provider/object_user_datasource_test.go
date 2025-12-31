package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// sample_user with tag "Department" and value "Finance" in namespace "ns1" is assumed to exist in the test ObjectScale cluster.
func TestAccObjectUserDataSource_basic(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
				data "objectscale_object_user" "all" {
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Make sure at least one user is returned
					resource.TestCheckResourceAttrSet("data.objectscale_object_user.all", "users.0.id"),
					resource.TestCheckResourceAttrSet("data.objectscale_object_user.all", "users.0.name"),
					resource.TestCheckResourceAttrSet("data.objectscale_object_user.all", "users.0.created"),
					resource.TestCheckResourceAttrSet("data.objectscale_object_user.all", "users.0.locked"),
					resource.TestCheckResourceAttrSet("data.objectscale_object_user.all", "users.0.tags.#"),
					resource.TestCheckResourceAttrSet("data.objectscale_object_user.all", "users.0.namespace"),
					resource.TestCheckResourceAttrSet("data.objectscale_object_user.all", "users.0.secret_keys.%"),
				),
			},
		},
	})
}

func TestAccObjectUserDataSource_withUsernameFilter(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
				data "objectscale_object_user" "by_username" {
					name  = "sample_user"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.objectscale_object_user.by_username", "users.0.name", "sample_user"),
					resource.TestCheckResourceAttr("data.objectscale_object_user.by_username", "users.0.id", "sample_user"),
					resource.TestCheckResourceAttr("data.objectscale_object_user.by_username", "users.0.locked", "false"),
					resource.TestCheckResourceAttr("data.objectscale_object_user.by_username", "users.0.tags.#", "1"),
					resource.TestCheckResourceAttr("data.objectscale_object_user.by_username", "users.0.namespace", "ns1"),
					resource.TestCheckResourceAttr("data.objectscale_object_user.by_username", "users.0.secret_keys.%", "10"),
				),
			},
		},
	})
}

func TestAccObjectUserDataSource_withNamespaceFilter(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
				data "objectscale_object_user" "by_namespace" {
					namespace  = "ns1"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify that the returned users belong to the namespace
					resource.TestCheckResourceAttrSet("data.objectscale_object_user.by_namespace", "users.0.id"),
					resource.TestCheckResourceAttrSet("data.objectscale_object_user.by_namespace", "users.0.name"),
					resource.TestCheckResourceAttrSet("data.objectscale_object_user.by_namespace", "users.0.created"),
					resource.TestCheckResourceAttrSet("data.objectscale_object_user.by_namespace", "users.0.locked"),
					resource.TestCheckResourceAttrSet("data.objectscale_object_user.by_namespace", "users.0.tags.#"),
					resource.TestCheckResourceAttr("data.objectscale_object_user.by_namespace", "users.0.namespace", "ns1"),
					resource.TestCheckResourceAttrSet("data.objectscale_object_user.by_namespace", "users.0.secret_keys.%"),
				),
			},
		},
	})
}

func TestAccObjectUserDataSource_UserTagsAndAccessKeys(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_object_user" "with_keys_tags" {
						tag = "Department"
						value = "Finance"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check that at least one user is returned
					resource.TestCheckResourceAttrSet("data.objectscale_object_user.with_keys_tags", "users.0.id"),

					// Check tags only if any exist
					resource.TestCheckResourceAttrSet("data.objectscale_object_user.with_keys_tags", "users.0.tags.#"),
					resource.TestCheckResourceAttr("data.objectscale_object_user.with_keys_tags", "users.0.tags.0.name", "Department"),
					resource.TestCheckResourceAttr("data.objectscale_object_user.with_keys_tags", "users.0.tags.0.value", "Finance"),
				),
			},
		},
	})
}
