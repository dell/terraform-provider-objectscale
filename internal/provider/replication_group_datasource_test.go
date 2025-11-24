package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var rgDSConfig = `
	data "objectscale_replication_group" "all" {
	}
`

// Test to Fetch Replication Group.
func TestAccRGDs(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + rgDSConfig,
			},
		},
	})
}
