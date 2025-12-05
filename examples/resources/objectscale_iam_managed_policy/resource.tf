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
# Before running `terraform apply`, ensure that the specified user/group/role exists in the ObjectScale
# Running `terraform apply` will set the specified managed policies for that user/group/role in the ObjectScale
resource "objectscale_iam_managed_policy" "example" {
  # Namespace to which the IAM entity belongs must be provided
  namespace = "ns1"

  # Exactly one of username, groupname, or rolename must be provided
  username = "userTest1"
  # groupname = "groupTest1"
  # rolename  = "roleTest1"

  # List of managed policies to be set on the specified IAM entity must be provided
  policy_arns = [
    "urn:ecs:iam:::policy/ECSS3ReadOnlyAccess",
    "urn:ecs:iam:::policy/IAMReadOnlyAccess"
  ]
}
