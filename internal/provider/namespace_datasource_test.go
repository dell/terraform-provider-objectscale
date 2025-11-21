package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Test to Fetch Namespaces.
func TestAccNSDs(t *testing.T) {
	// if os.Getenv("TF_ACC") == "" {
	// 	t.Skip("Dont run with units tests because it will try to create the context")
	// }

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
				data "objectscale_replication_group" "all" {
				}
				`,
			},
		},
	})
}
