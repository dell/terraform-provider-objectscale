/*
Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

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

# Example: Get storage pools on the local VDC
data "objectscale_vdc" "all" {
}

# Example: Get storage pools on a remote VDC
data "objectscale_vdc" "remote_vdc1" {
  name = "remote_vdc1"
}
data "objectscale_storage_pool" "by_name" {
  vdc_id = data.objectscale_vdc.remote_vdc1.vdcs[0].id
}

# Example: Get storage pool by name on the local VDC
data "objectscale_storage_pool" "by_name_local" {
  name = "sp1"
}

# Example: Get storage pool by name on a remote VDC
data "objectscale_vdc" "remote_vdc1" {
  name = "remote_vdc1"
}
data "objectscale_storage_pool" "by_name_remote" {
  name   = "remote_sp1"
  vdc_id = data.objectscale_vdc.remote_vdc1.vdcs[0].id
}

# Example: Get storage pool by ID
data "objectscale_storage_pool" "id" {
  id = "urn:storageos:VirtualArray:e7b72e23-4ee5-45c3-9e13-33b3c4c3c373"
}
