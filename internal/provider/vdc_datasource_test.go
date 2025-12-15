package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Test to Fetch Replication Group.
func TestAccVDCDs(t *testing.T) {
	defer testUserTokenCleanup(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
				data "objectscale_vdc" "all" {
				}
				`,
			},
			{
				Config: ProviderConfigForTesting + `
				data "objectscale_vdc" "name" {
					name = "vdc1"
				}
				`,
			},
			{
				Config: ProviderConfigForTesting + `
				data "objectscale_vdc" "local" {
					local = true
				}
				`,
			},
			{
				Config: ProviderConfigForTesting + `
				data "objectscale_vdc" "local" {
					local = true
				}
				data "objectscale_vdc" "id" {
					id = data.objectscale_vdc.local.vdcs[0].id
				}
				`,
			},
		},
	})
}
