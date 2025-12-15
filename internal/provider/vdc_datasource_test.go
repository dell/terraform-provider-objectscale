package provider

import (
	"fmt"
	"os"
	"regexp"
	"terraform-provider-objectscale/internal/clientgen"
	"testing"

	"github.com/bytedance/mockey"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Test to Fetch Replication Group.
func TestAccVDCDs(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	defer testUserTokenCleanup(t)
	var upM *mockey.Mocker
	upMConfig := func() {
		upM = mockey.Mock((*clientgen.ZoneInfoApiService).ZoneInfoServiceListAllVdcExecute).
			Return(nil, nil, fmt.Errorf("error")).Build()
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// get all
				Config: ProviderConfigForTesting + `
				data "objectscale_vdc" "all" {
				}
				`,
			},
			{
				// get all error mocked
				PreConfig: upMConfig,
				Config: ProviderConfigForTesting + `
				data "objectscale_vdc" "all" {
				}
				`,
				ExpectError: regexp.MustCompile(`Error fetching VDCs`),
			},
			{
				// get by name
				PreConfig: func() {
					upM.UnPatch()
				},
				Config: ProviderConfigForTesting + `
				data "objectscale_vdc" "name" {
					name = "vdc1"
				}
				`,
			},
			{
				// get by invalid name
				Config: ProviderConfigForTesting + `
				data "objectscale_vdc" "name" {
					name = "invalid-id"
				}
				`,
				ExpectError: regexp.MustCompile(`Error fetching VDC with name`),
			},
			{
				// get local vdc
				Config: ProviderConfigForTesting + `
				data "objectscale_vdc" "local" {
					local = true
				}
				`,
			},
			{
				// get local vdc error mocked
				PreConfig: func() {
					upM = mockey.Mock((*clientgen.ZoneInfoApiService).ZoneInfoServiceGetLocalVdcExecute).
						Return(nil, nil, fmt.Errorf("error")).Build()
				},
				Config: ProviderConfigForTesting + `
				data "objectscale_vdc" "local" {
					local = true
				}
				`,
				ExpectError: regexp.MustCompile(`Error fetching local VDC`),
			},
			{
				// get by ID
				PreConfig: func() {
					upM.UnPatch()
				},
				Config: ProviderConfigForTesting + `
				data "objectscale_vdc" "local" {
					local = true
				}
				data "objectscale_vdc" "id" {
					id = data.objectscale_vdc.local.vdcs[0].id
				}
				`,
			},
			{
				// get by invalid id
				Config: ProviderConfigForTesting + `
				data "objectscale_vdc" "id" {
					id = "invalid-id"
				}
				`,
				ExpectError: regexp.MustCompile(`Error fetching VDC with ID`),
			},
		},
	})
}
