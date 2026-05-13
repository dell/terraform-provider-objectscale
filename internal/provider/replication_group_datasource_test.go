// Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.
//
// Licensed under the Mozilla Public License Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
