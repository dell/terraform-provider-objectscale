package provider

import (
	"fmt"
	"regexp"
	"terraform-provider-objectscale/internal/clientgen"
	"testing"

	"github.com/bytedance/mockey"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

var rgDSConfig = `
	data "objectscale_replication_group" "all" {
	}
`

var rgDSConfig2 = `
	data "objectscale_replication_group" "all" {
		name = "rg1"
	}
`

// Test to Fetch Replication Group.
func TestAccRGDs(t *testing.T) {
	defer testUserTokenCleanup(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + rgDSConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.objectscale_replication_group.all", "replication_groups.0.id"),
					resource.TestCheckResourceAttrSet("data.objectscale_replication_group.all", "replication_groups.1.id"),
				),
			},
			{
				Config: ProviderConfigForTesting + rgDSConfig2,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.objectscale_replication_group.all", "replication_groups.0.id"),
				),
			},
		},
	})
}

func TestAccRGDsErrorGetAll(t *testing.T) {
	defer testUserTokenCleanup(t)
	var FunctionMocker *mockey.Mocker
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					FunctionMocker = mockey.Mock((*clientgen.DataVpoolApiService).DataServiceVpoolServiceGetDataServiceVpoolsExecute).
						Return(nil, nil, fmt.Errorf("error")).Build()
				},
				Config:      ProviderConfigForTesting + rgDSConfig,
				ExpectError: regexp.MustCompile("Error getting the list of replication group"),
			},
		},
	})
	FunctionMocker.UnPatch()
}
