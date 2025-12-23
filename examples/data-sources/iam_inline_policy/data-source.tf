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

data "objectscale_iam_inline_policy" "usernamefilter" {
  namespace = "ns1"
  username  = "sample_user_1"
}

output "objectscale_iam_inline_policy_username" {
  value = data.objectscale_iam_inline_policy.usernamefilter
}

data "objectscale_iam_inline_policy" "groupnamefilter" {
  namespace = "ns1"
  groupname = "group_008"
}

output "objectscale_iam_inline_policy_groupname" {
  value = data.objectscale_iam_inline_policy.groupnamefilter
}

data "objectscale_iam_inline_policy" "rolenamefilter" {
  namespace = "ns1"
  rolename  = "roleTest1"
}

output "objectscale_iam_inline_policy_rolename" {
  value = data.objectscale_iam_inline_policy.rolenamefilter
}
