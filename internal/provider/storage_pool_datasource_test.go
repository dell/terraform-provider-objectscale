package provider

import (
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Test to Fetch Replication Group.
func TestAccStoragePoolDs(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	defer testUserTokenCleanup(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// get all on local
				Config: ProviderConfigForTesting + `
				data "objectscale_storage_pool" "all_local" {
				}
				`,
			},
			{
				// get by vdc id
				Config: ProviderConfigForTesting + `
				data "objectscale_vdc" "local" {
					local = true
				}
				data "objectscale_storage_pool" "by_vdc_id" {
					vdc_id = data.objectscale_vdc.local.vdcs[0].id
				}
				`,
			},
			{
				// get by invalid vdc id
				Config: ProviderConfigForTesting + `
				data "objectscale_storage_pool" "by_invalid_vdc_id" {
					vdc_id = "invalid-id"
				}
				`,
				ExpectError: regexp.MustCompile(`Error fetching Storage Pools`),
			},
			{
				// get by name
				Config: ProviderConfigForTesting + `
				data "objectscale_storage_pool" "by_name" {
					name = "sp1"
				}
				`,
			},
			{
				// get by invalid name
				Config: ProviderConfigForTesting + `
				data "objectscale_storage_pool" "by_invalid_name" {
					name = "invalid-name"
				}
				`,
				ExpectError: regexp.MustCompile(`Storage Pool with name: invalid-name not found`),
			},
			{
				// get by invalid ID
				Config: ProviderConfigForTesting + `
				data "objectscale_storage_pool" "by_invalid_id" {
					id = "invalid-id"
				}
				`,
				ExpectError: regexp.MustCompile(`Error fetching Storage Pool with ID: invalid-id`),
			},
			{
				// get by ID
				Config: ProviderConfigForTesting + `
				data "objectscale_storage_pool" "by_name" {
					name = "sp1"
				}
				data "objectscale_storage_pool" "by_id" {
					id = data.objectscale_storage_pool.by_name.storage_pools[0].id
				}
				`,
			},
		},
	})
}
