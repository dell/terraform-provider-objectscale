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

### Example: Get all policies in a namespace
data "objectscale_iam_policy" "all_policies" {
  namespace = "ns1"
}

output "all_policies" {
  value = data.objectscale_iam_policy.all_policies.policies
}

### Example: Get a policy by ARN

data "objectscale_iam_policy" "policy_by_arn" {
  namespace = "ns1"
  arn       = "urn:ecs:iam:::policy/ECSS3FullAccess"
}

output "policy_by_arn" {
  value = data.objectscale_iam_policy.policy_by_arn.policies[0]
}

### Example: Get policies attached to a user

data "objectscale_iam_policy" "user_policies" {
  namespace = "ns1"
  user      = "user1"
}

### Example: Get policies attached to a group

data "objectscale_iam_policy" "group_policies" {
  namespace = "ns1"
  group     = "group1"
}

### Example: Get policies attached to a role

data "objectscale_iam_policy" "role_policies" {
  namespace = "ns1"
  role      = "role1"
}

