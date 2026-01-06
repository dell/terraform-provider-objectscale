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

# Example of getting the pre-requisite VDCs and Storage Pools
data "objectscale_vdc" "vdcs" {
  for_each = toset(["vdc1", "vdc2", "vdc3"])
  name     = each.value
}

data "objectscale_storage_pool" "sps" {
  for_each = toset(["vdc1", "vdc2", "vdc3"])
  # this example assumes that each VDC has a storage pool named "sp1"
  name   = "sp1"
  vdc_id = data.objectscale_vdc.vdcs[each.value].vdcs[0].id
}

# Example 1: Active Replication group with 3 VDCs
resource "objectscale_replication_group" "active" {
  name        = "ActiveRG"
  description = "Active replication group"
  zone_mappings = [
    {
      vdc          = data.objectscale_vdc.vdcs["vdc1"].vdcs[0].id
      storage_pool = data.objectscale_storage_pool.sps["vdc1"].storage_pools[0].id
    },
    {
      vdc          = data.objectscale_vdc.vdcs["vdc2"].vdcs[0].id
      storage_pool = data.objectscale_storage_pool.sps["vdc2"].storage_pools[0].id
    },
    {
      vdc          = data.objectscale_vdc.vdcs["vdc3"].vdcs[0].id
      storage_pool = data.objectscale_storage_pool.sps["vdc3"].storage_pools[0].id
    }
  ]
  enable_rebalancing   = true
  allow_all_namespaces = true
}

# Example 2: Active Replication group with data being replicated fully to all sites
resource "objectscale_replication_group" "fully_replicated" {
  name        = "FullyReplicatedRG"
  description = "Fully replicated Active replication group"
  zone_mappings = [
    {
      vdc          = data.objectscale_vdc.vdcs["vdc1"].vdcs[0].id
      storage_pool = data.objectscale_storage_pool.sps["vdc1"].storage_pools[0].id
    },
    {
      vdc          = data.objectscale_vdc.vdcs["vdc2"].vdcs[0].id
      storage_pool = data.objectscale_storage_pool.sps["vdc2"].storage_pools[0].id
    }
  ]
  replicate_to_all_sites = true
}

# Example 3: Passive Replication group
# Passive replication always has 3 zones, with 1 zone as a target
resource "objectscale_replication_group" "passive" {
  name        = "PassiveRG"
  description = "Passive replication group"
  zone_mappings = [
    {
      vdc          = data.objectscale_vdc.vdcs["vdc1"].vdcs[0].id
      storage_pool = data.objectscale_storage_pool.sps["vdc1"].storage_pools[0].id
    },
    {
      vdc          = data.objectscale_vdc.vdcs["vdc2"].vdcs[0].id
      storage_pool = data.objectscale_storage_pool.sps["vdc2"].storage_pools[0].id
    },
    {
      vdc          = data.objectscale_vdc.vdcs["vdc3"].vdcs[0].id
      storage_pool = data.objectscale_storage_pool.sps["vdc3"].storage_pools[0].id
      # this is the target VDC
      is_replication_target = true
    }
  ]
}

# Example 4: A replication group that should not be removed from the state by accident
# A replication group cannot be deleted in ObjectScale.
# If a destroy plan is applied, this resource simply removes itself from state, but the replication group remains in ObjectScale.
# If your usecase requires that the resource throw an error if it is being destroyed, set the prevent_destroy lifecycle attribute
resource "objectscale_replication_group" "non_destroyable" {
  name = "NotDestroyableRG"
  zone_mappings = [
    {
      vdc          = data.objectscale_vdc.vdcs["vdc1"].vdcs[0].id
      storage_pool = data.objectscale_storage_pool.sps["vdc1"].storage_pools[0].id
    }
  ]

  lifecycle {
    # this lifecycle attribute prevents accidental removal of this resource from state
    prevent_destroy = true
  }
}
