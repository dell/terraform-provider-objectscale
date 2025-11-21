package provider

import (
	"fmt"
	"regexp"
	"terraform-provider-objectscale/internal/helper"
	"testing"

	. "github.com/bytedance/mockey"
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
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + rgDSConfig,
			},
		},
	})
}

func TestAccClusterEmailDatasourceErrorGetAll(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.UpdateReplicationGroupState).Return(nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfigForTesting + rgDSConfig,
				ExpectError: regexp.MustCompile("mock error"),
			},
		},
	})
}
