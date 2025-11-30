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
# Running `terraform apply` will set the specified inline policies for that user/group/role in the ObjectScale
resource "objectscale_iam_inline_policy" "example" {
  # Namespace to which the IAM entity belongs must be provided
  namespace = "ns1"

  # Exactly one of username, groupname, or rolename must be provided
  username = "userTest1"
  # groupname = "groupTest1"
  # rolename  = "roleTest1"

  # List of inline policies to be set on the specified IAM entity must be provided
  # Ensure that you provide a valid JSON for the policy documents
  policies = [
    {
      name = "inlinePolicyTest1"
      document = jsonencode({
        Version = "2012-10-17",
        Statement = [
          {
            Sid    = "VisualEditor0",
            Effect = "Allow",
            Action = [
              "iam:GetPolicyVersion",
              "iam:GetUser",
              "iam:GetPolicy",
              "iam:GetGroupPolicy",
              "iam:GetRole",
              "iam:GetAccessKeyLastUsed",
              "iam:GetGroup",
              "iam:GetUserPolicy",
              "iam:GetSAMLProvider",
              "iam:GetRolePolicy",
              "iam:GetContextKeysForCustomPolicy",
              "iam:GetContextKeysForPrincipalPolicy",
              "iam:SimulateCustomPolicy",
              "iam:SimulatePrincipalPolicy"
            ],
            Resource = "*"
          }
        ]
      })
    },
    {
      name = "inlinePolicyTest2"
      document = jsonencode({
        Version = "2012-10-17",
        Statement = [
          {
            Sid    = "VisualEditor0",
            Effect = "Allow",
            Action = [
              "iam:DeleteAccessKey",
              "iam:UpdateSAMLProvider",
              "iam:CreateRole",
              "iam:RemoveUserFromGroup",
              "iam:AddUserToGroup",
              "iam:UpdateUser",
              "iam:CreateAccessKey",
              "iam:UpdateAccessKey",
              "iam:CreateSAMLProvider",
              "iam:DeleteRole",
              "iam:UpdateRole",
              "iam:DeleteGroup",
              "iam:UpdateGroup",
              "iam:CreateUser",
              "iam:CreateGroup",
              "iam:DeleteSAMLProvider",
              "iam:DeleteUser"
            ],
            Resource = "*"
          }
        ]
      })
    }
  ]
}
