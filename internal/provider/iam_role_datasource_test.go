package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// role_Test1 is assumed to exist in the test ObjectScale cluster.

func TestAccIAMRoleDataSource_basic(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
				data "objectscale_iam_role" "all" {
					namespace = "ns1"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Make sure at least one role is returned
					resource.TestCheckResourceAttrSet("data.objectscale_iam_role.all", "roles.0.role_id"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_role.all", "roles.0.role_name"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_role.all", "roles.0.arn"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_role.all", "roles.0.path"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_role.all", "roles.0.create_date"),
				),
			},
		},
	})
}

func TestAccIAMRoleDataSource_withRoleNameFilter(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
				data "objectscale_iam_role" "by_name" {
					namespace = "ns1"
					role_name  = "roleTest1"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.objectscale_iam_role.by_name", "roles.0.role_name", "roleTest1"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_role.by_name", "roles.0.role_id"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_role.by_name", "roles.0.arn"),
					resource.TestCheckResourceAttrSet("data.objectscale_iam_role.by_name", "roles.0.path"),
				),
			},
		},
	})
}

func TestAccIAMRoleDataSource_getRole_404(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_role" "test" {
					  namespace = "ns1"
					  role_name = "missing_role_abc"
					}
				`,
				ExpectError: regexp.MustCompile(`(?i)(Error retrieving IAM role|Failed retrieving role)`),
			},
		},
	})
}

func TestAccIAMRoleDataSource_ListRole_404(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_iam_role" "test1" {
					  namespace = "nonexistent_namespace"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.objectscale_iam_role.test1", "roles.#", "0"),
				),
			},
		},
	})
}
