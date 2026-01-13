/*
Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"fmt"
	"os"
	"regexp"
	"terraform-provider-objectscale/internal/clientgen"
	"testing"
	"time"

	"github.com/bytedance/mockey"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Test to Fetch Replication Group.
func TestAccRgRs(t *testing.T) {
	if os.Getenv("TF_ACC_MANUAL") == "" {
		t.Skip("This test is skipped unless TF_ACC_MANUAL is set")
	}
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	defer testUserTokenCleanup(t)

	mockTime := mockey.Mock(time.Sleep).Return().Build()
	defer mockTime.UnPatch()

	var mockAPI *mockey.Mocker
	unPatchFunc := func() {
		mockAPI.UnPatch()
	}
	ProviderConfigForRgTesting := ProviderConfigForTesting + `
	data "objectscale_vdc" "vdc1" {
		name = "vdc1"
	}

	data "objectscale_vdc" "vdc2" {
		name = "vdc2"
	}

	data "objectscale_vdc" "vdc3" {
		name = "vdc3"
	}

	data "objectscale_storage_pool" "sp1" {
		name   = "sp1"
		vdc_id = data.objectscale_vdc.vdc1.vdcs[0].id
	}

	data "objectscale_storage_pool" "sp2" {
		name   = "sp1"
		vdc_id = data.objectscale_vdc.vdc2.vdcs[0].id
	}

	data "objectscale_storage_pool" "sp3" {
		name   = "sp1"
		vdc_id = data.objectscale_vdc.vdc3.vdcs[0].id
	}
	`
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Create Passive site with full replication as true - Negative
				Config: ProviderConfigForRgTesting + `
				resource objectscale_replication_group test_replication_group {
					name = "test"
					zone_mappings = [
						{
							vdc = data.objectscale_vdc.vdc1.vdcs[0].id
							storage_pool = data.objectscale_storage_pool.sp1.storage_pools[0].id
						},
						{
						    vdc = data.objectscale_vdc.vdc2.vdcs[0].id
						    storage_pool = data.objectscale_storage_pool.sp2.storage_pools[0].id
						},
						{
						    vdc = data.objectscale_vdc.vdc3.vdcs[0].id
						    storage_pool = data.objectscale_storage_pool.sp3.storage_pools[0].id
							is_replication_target = true							
						}
					]
					replicate_to_all_sites = true
				}
				`,
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("replicate_to_all_sites can be set to true only for Active replication"),
			},
			{
				// Create Passive with 2 zones only - Negative
				Config: ProviderConfigForRgTesting + `
				resource objectscale_replication_group test_replication_group {
					name = "test"
					zone_mappings = [
						{
							vdc = data.objectscale_vdc.vdc1.vdcs[0].id
							storage_pool = data.objectscale_storage_pool.sp1.storage_pools[0].id
						},
						{
						    vdc = data.objectscale_vdc.vdc2.vdcs[0].id
						    storage_pool = data.objectscale_storage_pool.sp2.storage_pools[0].id
							is_replication_target = true
						},
					]
				}
				`,
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("Invalid number of zones for Passive Replication Group"),
			},
			{
				// Create Passive with no 2 targets and 1 source - Negative
				Config: ProviderConfigForRgTesting + `
				resource objectscale_replication_group test_replication_group {
					name = "test"
					zone_mappings = [
						{
							vdc = data.objectscale_vdc.vdc1.vdcs[0].id
							storage_pool = data.objectscale_storage_pool.sp1.storage_pools[0].id
							is_replication_target = true
						},
						{
						    vdc = data.objectscale_vdc.vdc2.vdcs[0].id
						    storage_pool = data.objectscale_storage_pool.sp2.storage_pools[0].id
							is_replication_target = true
						},
						{
						    vdc = data.objectscale_vdc.vdc3.vdcs[0].id
						    storage_pool = data.objectscale_storage_pool.sp3.storage_pools[0].id
						}
					]
				}
				`,
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("Invalid number of source and target zones for Passive Replication Group"),
			},
			{
				// Create mock error
				PreConfig: func() {
					mockAPI = mockey.Mock((*clientgen.DataVpoolApiService).DataServiceVpoolServiceCreateDataServiceVpoolExecute).Return(
						nil, nil, fmt.Errorf("mock error"),
					).Build()
				},
				Config: ProviderConfigForRgTesting + `
				resource objectscale_replication_group test_replication_group {
					name = "test"
					zone_mappings = [
						{
							vdc = data.objectscale_vdc.vdc1.vdcs[0].id
							storage_pool = data.objectscale_storage_pool.sp1.storage_pools[0].id
						}
					]
				}
				`,
				ExpectError: regexp.MustCompile("Error creating Replication Group"),
			},
			{
				// Create with 2 zones
				PreConfig: unPatchFunc,
				Config: ProviderConfigForRgTesting + `
				resource objectscale_replication_group test_replication_group {
					name = "test"
					zone_mappings = [
						{
							vdc = data.objectscale_vdc.vdc1.vdcs[0].id
							storage_pool = data.objectscale_storage_pool.sp1.storage_pools[0].id
						},
						{
						    vdc = data.objectscale_vdc.vdc2.vdcs[0].id
						    storage_pool = data.objectscale_storage_pool.sp2.storage_pools[0].id
						}
					]
				}
				`,
			},
			{
				ResourceName:      "objectscale_replication_group.test_replication_group",
				ImportState:       true,
				ImportStateId:     "test",
				ImportStateVerify: true,
			},
			{
				// import invalid
				ResourceName:  "objectscale_replication_group.test_replication_group",
				ImportState:   true,
				ImportStateId: "invalid-id",
				ExpectError:   regexp.MustCompile("Could not find replication group with name"),
			},
			{
				// mock import error when getting list of replication groups
				PreConfig: func() {
					mockAPI = mockey.Mock((*clientgen.DataVpoolApiService).DataServiceVpoolServiceGetDataServiceVpoolsExecute).Return(
						nil, nil, fmt.Errorf("mock error"),
					).Build()
				},
				ResourceName:  "objectscale_replication_group.test_replication_group",
				ImportState:   true,
				ImportStateId: "test",
				ExpectError:   regexp.MustCompile("Error getting the list of replication groups"),
			},
			{
				// mock refresh error
				PreConfig: func() {
					unPatchFunc()
					mockAPI = mockey.Mock((*clientgen.DataVpoolApiService).DataServiceVpoolServiceGetDataServiceStoreExecute).Return(
						nil, nil, fmt.Errorf("mock error"),
					).Build()
				},
				RefreshState: true,
				ExpectError:  regexp.MustCompile("Error reading Replication Group state"),
			},
			{
				// mock refresh bad data
				PreConfig: func() {
					unPatchFunc()
					mockAPI = mockey.Mock((*clientgen.DataVpoolApiService).DataServiceVpoolServiceGetDataServiceStoreExecute).Return(
						&clientgen.DataServiceVpoolServiceGetDataServiceStoreResponse{}, nil, nil,
					).Build()
				},
				RefreshState: true,
				ExpectError:  regexp.MustCompile("Replication Group not found with ID"),
			},
			{
				// Update replicate_to_all_sites to true - Negative
				PreConfig: unPatchFunc,
				Config: ProviderConfigForRgTesting + `
				resource objectscale_replication_group test_replication_group {
					name = "test"
					zone_mappings = [
						{
							vdc = data.objectscale_vdc.vdc1.vdcs[0].id
							storage_pool = data.objectscale_storage_pool.sp1.storage_pools[0].id
						},
						{
						    vdc = data.objectscale_vdc.vdc2.vdcs[0].id
						    storage_pool = data.objectscale_storage_pool.sp2.storage_pools[0].id
						}
					]
					replicate_to_all_sites = true
				}
				`,
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("replicate_to_all_sites cannot be modified"),
			},
			{
				// Add a duplicate zone - Negative
				Config: ProviderConfigForRgTesting + `
				resource objectscale_replication_group test_replication_group {
					name = "test"
					zone_mappings = [
						{
							vdc = data.objectscale_vdc.vdc1.vdcs[0].id
							storage_pool = data.objectscale_storage_pool.sp1.storage_pools[0].id
						},
						{
						    vdc = data.objectscale_vdc.vdc2.vdcs[0].id
						    storage_pool = data.objectscale_storage_pool.sp2.storage_pools[0].id
						},
						{
						    vdc = data.objectscale_vdc.vdc2.vdcs[0].id
						    storage_pool = data.objectscale_storage_pool.sp2.storage_pools[0].id
							is_replication_target = true
						}
					]

				}
				`,
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("Duplicate zone mapping"),
			},
			{
				// Add a target zone in active group - Negative
				Config: ProviderConfigForRgTesting + `
				resource objectscale_replication_group test_replication_group {
					name = "test"
					zone_mappings = [
						{
							vdc = data.objectscale_vdc.vdc1.vdcs[0].id
							storage_pool = data.objectscale_storage_pool.sp1.storage_pools[0].id
						},
						{
						    vdc = data.objectscale_vdc.vdc2.vdcs[0].id
						    storage_pool = data.objectscale_storage_pool.sp2.storage_pools[0].id
						},
						{
						    vdc = data.objectscale_vdc.vdc3.vdcs[0].id
						    storage_pool = data.objectscale_storage_pool.sp3.storage_pools[0].id
							is_replication_target = true
						}
					]

				}
				`,
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("Cannot update replication group type from Active to Passive"),
			},
			{
				// Add and remove a zone at the same time - Negative
				Config: ProviderConfigForRgTesting + `
				resource objectscale_replication_group test_replication_group {
					name = "test"
					zone_mappings = [
						{
							vdc = data.objectscale_vdc.vdc1.vdcs[0].id
							storage_pool = data.objectscale_storage_pool.sp1.storage_pools[0].id
						},
						{
						    vdc = data.objectscale_vdc.vdc3.vdcs[0].id
						    storage_pool = data.objectscale_storage_pool.sp3.storage_pools[0].id
						}
					]

				}
				`,
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("Cannot add and remove zones in one operation"),
			},
			{
				// Update mock error basic
				PreConfig: func() {
					mockAPI = mockey.Mock((*clientgen.DataVpoolApiService).DataServiceVpoolServicePutDataServiceVpoolExecute).Return(
						nil, nil, fmt.Errorf("mock error")).Build()
				},
				Config: ProviderConfigForRgTesting + `
				resource objectscale_replication_group test_replication_group {
					name = "test_updated"
					zone_mappings = [
						{
							vdc = data.objectscale_vdc.vdc1.vdcs[0].id
							storage_pool = data.objectscale_storage_pool.sp1.storage_pools[0].id
						},
						{
						    vdc = data.objectscale_vdc.vdc2.vdcs[0].id
						    storage_pool = data.objectscale_storage_pool.sp2.storage_pools[0].id
						}
					]
				}
				`,
				ExpectError: regexp.MustCompile("Error updating Replication Group attributes"),
			},
			{
				// Add zones mock error
				PreConfig: func() {
					unPatchFunc()
					mockAPI = mockey.Mock((*clientgen.DataVpoolApiService).DataServiceVpoolServiceAddToVpoolExecute).Return(
						nil, nil, fmt.Errorf("mock error")).Build()
				},
				Config: ProviderConfigForRgTesting + `
				resource objectscale_replication_group test_replication_group {
					name = "test"
					zone_mappings = [
						{
							vdc = data.objectscale_vdc.vdc1.vdcs[0].id
							storage_pool = data.objectscale_storage_pool.sp1.storage_pools[0].id
						},
						{
						    vdc = data.objectscale_vdc.vdc2.vdcs[0].id
						    storage_pool = data.objectscale_storage_pool.sp2.storage_pools[0].id
						},
						{
						    vdc = data.objectscale_vdc.vdc3.vdcs[0].id
						    storage_pool = data.objectscale_storage_pool.sp3.storage_pools[0].id
						}
					]
				}
				`,
				ExpectError: regexp.MustCompile("Error adding new zones to Replication Group"),
			},
			{
				// Remove zones mock error
				PreConfig: func() {
					unPatchFunc()
					mockAPI = mockey.Mock((*clientgen.DataVpoolApiService).DataServiceVpoolServiceRemoveFromVpoolExecute).Return(
						nil, nil, fmt.Errorf("mock error")).Build()
				},
				Config: ProviderConfigForRgTesting + `
				resource objectscale_replication_group test_replication_group {
					name = "test"
					zone_mappings = [
						{
							vdc = data.objectscale_vdc.vdc1.vdcs[0].id
							storage_pool = data.objectscale_storage_pool.sp1.storage_pools[0].id
						}
					]
				}
				`,
				ExpectError: regexp.MustCompile("Error removing zones from Replication Group"),
			},
			{
				// Add 1 more zone and change name
				PreConfig: unPatchFunc,
				Config: ProviderConfigForRgTesting + `
				resource objectscale_replication_group test_replication_group {
					name = "test_1"
					zone_mappings = [
						{
							vdc = data.objectscale_vdc.vdc1.vdcs[0].id
							storage_pool = data.objectscale_storage_pool.sp1.storage_pools[0].id
						},
						{
						    vdc = data.objectscale_vdc.vdc2.vdcs[0].id
						    storage_pool = data.objectscale_storage_pool.sp2.storage_pools[0].id
						},
						{
						    vdc = data.objectscale_vdc.vdc3.vdcs[0].id
						    storage_pool = data.objectscale_storage_pool.sp3.storage_pools[0].id
						}
					]
				}
				`,
			},
			{
				ResourceName:      "objectscale_replication_group.test_replication_group",
				ImportState:       true,
				ImportStateId:     "test_1",
				ImportStateVerify: true,
			},
			{
				// revert the change
				Config: ProviderConfigForRgTesting + `
				resource objectscale_replication_group test_replication_group {
					name = "test"
					zone_mappings = [
						{
							vdc = data.objectscale_vdc.vdc1.vdcs[0].id
							storage_pool = data.objectscale_storage_pool.sp1.storage_pools[0].id
						},
						{
						    vdc = data.objectscale_vdc.vdc2.vdcs[0].id
							storage_pool = data.objectscale_storage_pool.sp2.storage_pools[0].id
						}
					]
				}
				`,
			},
		},
	})
}
