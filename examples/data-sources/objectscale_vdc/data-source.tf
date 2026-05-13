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

# Example: Get all VDCs
data "objectscale_vdc" "all" {
}

# Example: Get VDC by name
data "objectscale_vdc" "name" {
  name = "vdc1"
}

# Example: Get the local VDC
data "objectscale_vdc" "local" {
  local = true
}

# Example: Get VDC by ID
data "objectscale_vdc" "id" {
  id = "urn:storageos:VirtualDataCenterData:1fcdfca4-4d0f-4a0a-b16a-d7aee4aa99a7"
}
