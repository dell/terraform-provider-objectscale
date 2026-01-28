---
# Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.
#
# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://mozilla.org/MPL/2.0/
#
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

title: "objectscale_replication_group resource"
linkTitle: "objectscale_replication_group"
page_title: "objectscale_replication_group Resource - terraform-provider-objectscale"
subcategory: "Data Protection"
description: |-
  This resource allows end user to Provision and manage Dell ObjectScale Replication Groups.
---

# objectscale_replication_group (Resource)

This resource allows end user to Provision and manage Dell ObjectScale Replication Groups.

~> **Note:** Deletion of Replication Group is not supported. If this resource gets planned for deletion, it will simply be removed from the state. But the Replication Group will not be destroyed on the ObjectScale array.

!> **Caution:** This resource does support removal of zones from Replication Group. But be cautious that removing zones from replication group may result in data loss.
We recommend contacting customer support before performing this operation.
Data loss may occur if prerequisite procedures are not properly followed.
Verify the following conditions:<br/>- Ensure that Geo replication is up-to-date.<br/>- Replication to/from VDC for the Replication Group will be disabled.<br/>- Recovery will be initiated. Data may not be available until recovery is complete.<br/>- Removal is permanent; the site cannot be added back to this replication group.<br/>- Data associated with this replication group will be permanently deleted from this VDC.<br/>- In cases where XOR encoding is utilized and the RG falls below 3 VDCs, the XOR encoded data will have to be replaced with fully replicated copies, which could significantly increase storage required to fully protect the data.


## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the Replication Group.
- `zone_mappings` (Attributes Set) List of zones (VDC + Storage Pool) which will be used for replication. (see [below for nested schema](#nestedatt--zone_mappings))

### Optional

- `allow_all_namespaces` (Boolean) Whether to allow all namespaces.
- `description` (String) Description of the Replication Group.
- `enable_rebalancing` (Boolean) Enable Rebalancing.
- `replicate_to_all_sites` (Boolean) Whether to replicate to all sites (for Active configuration). Cannot be updated.

### Read-Only

- `id` (String) Identifier of the Replication Group.
- `type` (String) Type of the Replication Group (Active/Passive). Cannot be updated.

<a id="nestedatt--zone_mappings"></a>
### Nested Schema for `zone_mappings`

Required:

- `storage_pool` (String) Storage Pool ID.
- `vdc` (String) Virtual Data Center ID.

Optional:

- `is_replication_target` (Boolean) In passive replication groups, one zone acts as the target. This attribute must be set to `true` for the zone which will act as the replication target.

Unless specified otherwise, all fields of this resource can be updated.

## Import

Import is supported using the following syntax:

```shell
# Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.

# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://mozilla.org/MPL/2.0/


# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# The command is
# terraform import objectscale_replication_group.<resource_name> <name of the replication group>
# Examples:
terraform import objectscale_replication_group.rg1 "ReplicationGroup1"

# after running this command, populate the other parameters in the config file to start managing this resource.
# Note: running "terraform show" after importing shows the current config/state of the resource. You can copy/paste that config to make it easier to manage the resource.
```
