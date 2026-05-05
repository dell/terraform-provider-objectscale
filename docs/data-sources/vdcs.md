---
# Copyright (c) 2025-2026 Dell Inc., or its subsidiaries. All Rights Reserved.
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

title: "objectscale_vdcs usage guide"
linkTitle: "objectscale_vdcs"
page_title: "objectscale_vdcs - Listing All VDCs - terraform-provider-objectscale"
subcategory: "Storage Topology & Capacity Domains"
description: |-
  Guide for listing all Virtual Data Centers (VDCs) across single-site or federated Dell ObjectScale deployments using the objectscale_vdc data source.
---

# Listing All VDCs (Federation View)

The `objectscale_vdc` data source supports listing **all** Virtual Data Centers in a federated ObjectScale deployment when invoked without any filter arguments. This page documents the "list all" (federation) usage pattern.

> **Note:** There is no separate `objectscale_vdcs` data source. The `objectscale_vdc` data source serves both single-VDC lookups and federation-wide listing. When no filter (`id`, `name`, or `local`) is specified, all VDCs in the federation are returned.

## API Endpoint

This usage pattern calls the ObjectScale REST API:

```
GET /object/vdcs/vdc/list
```

## Example Usage

### List All VDCs in a Federation

```terraform
# Retrieve every VDC visible to the authenticated user
data "objectscale_vdc" "federation" {
}

# Output the full list
output "all_vdc_names" {
  value = [for v in data.objectscale_vdc.federation.vdcs : v.vdc_name]
}

output "all_vdc_ids" {
  value = [for v in data.objectscale_vdc.federation.vdcs : v.vdc_id]
}

output "vdc_count" {
  value = length(data.objectscale_vdc.federation.vdcs)
}
```

### Use Federation List to Feed Downstream Resources

```terraform
# List all VDCs, then use the first one's ID for a storage pool lookup
data "objectscale_vdc" "all" {
}

data "objectscale_storage_pool" "by_vdc" {
  vdc_id = data.objectscale_vdc.all.vdcs[0].id
}
```

### Iterate Over All VDCs

```terraform
data "objectscale_vdc" "all" {
}

# Create a local map of VDC name → VDC ID
locals {
  vdc_map = { for v in data.objectscale_vdc.all.vdcs : v.vdc_name => v.id }
}
```

## Schema

The schema is identical to `objectscale_vdc`. See the [objectscale_vdc data source](vdc.md) for the full attribute reference.

### Key Attributes Returned per VDC

| Attribute | Type | Description |
|-----------|------|-------------|
| `vdc_id` | String | VDC identifier |
| `vdc_name` | String | VDC name |
| `local` | Boolean | `true` if the VDC is the local site |
| `remote` | Boolean | `true` if the VDC is a remote site |
| `permanently_failed` | Boolean | `true` if the VDC has permanently failed |
| `is_encryption_enabled` | Boolean | Encryption status |
| `inter_vdc_end_points` | String | Inter-VDC data endpoints |
| `inter_vdc_cmd_end_points` | String | Inter-VDC command endpoints |
| `management_end_points` | String | Management API endpoints |
| `creation_time` | Number | Unix timestamp of VDC creation |
| `id` | String | Immutable ECS-generated resource identifier |
| `name` | String | User-assigned resource name |

## Filter Behaviour Summary

| Filter | API Called | Result |
|--------|-----------|--------|
| _(none)_ | `GET /object/vdcs/vdc/list` | All VDCs in federation |
| `name = "vdc1"` | `GET /object/vdcs/vdc/vdc1` | Single VDC by name |
| `id = "urn:..."` | `GET /object/vdcs/vdcid/{id}` | Single VDC by ID |
| `local = true` | `GET /object/vdcs/vdc/local` | Local VDC only |

> **Constraint:** Only one of `id`, `name`, or `local` may be set at a time. Setting multiple produces an `Invalid Attribute Combination` error.

## Related Documentation

- [objectscale_vdc data source](vdc.md) — Single VDC lookups (by name, ID, or local flag)
- [ObjectScale VDC REST API](https://www.dell.com/support/kbdoc/en-us/000021062/dell-emc-ecs-api-reference) — Dell ECS/ObjectScale API Reference
