package provider

import (
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Test to Fetch Namespaces
func TestAccIamGroupResource(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// create and read testing
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
				resource"objectscale_iam_group" "test" {
					name = "testacc_group"
					namespace = "ns1"
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_group.test", "name", "testacc_group"),
					resource.TestCheckResourceAttr("objectscale_iam_group.test", "namespace", "ns1"),
				),
			},
			// update and testing (update not supported by API)
			{
				Config: ProviderConfigForTesting + `
				resource"objectscale_iam_group" "test" {
					name = "testacc_group_updated"
					namespace = "ns1"
				}
				`,
				ExpectError: regexp.MustCompile(".*Update operation is not supported.*"),
			},
			// import testing
			{
				ResourceName:      "objectscale_iam_group.test",
				ImportStateId:     "testacc_group:ns1",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
