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

# Available actions: Create, Update, Delete and Import
# Create, Update, and Delete operations require SECURITY_ADMIN role.
# Read operation requires any one of SECURITY_ADMIN, SYSTEM_ADMIN, SYSTEM_MONITOR roles.
# Running `terraform apply` will create/update the management user with the specified roles in the ObjectScale
resource "objectscale_management_user" "example" {
  # Required. Possible values are LOCAL_USER/AD_LDAP_USER/AD_LDAP_GROUP
  type = "LOCAL_USER"
  # Required. Format is "user1" for LOCAL_USER and "user1@domain" for AD_LDAP_USER/AD_LDAP_GROUP 
  name = "localuser1"

  # Required for LOCAL_USER and not applicable for AD_LDAP_USER/AD_LDAP_GROUP
  password = "pass123"

  # Optional parameters. If set to true, assigns the management user to that role.
  system_administrator   = true
  system_monitor         = true
  security_administrator = true
}
