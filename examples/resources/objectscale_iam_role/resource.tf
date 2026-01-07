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
resource "objectscale_iam_role" "example" {
  name                     = "example-role"
  namespace                = "ns1"
  description              = "An example role updated"
  permissions_boundary_arn = "urn:ecs:iam:::policy/ECSS3Access"
  max_session_duration     = 4000
  assume_role_policy_document = jsonencode({
    Version = "2012-11-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          AWS = [
            "urn:ecs:iam::ns1:user/sample_user_1"
          ]
        }
        Action = "sts:AssumeRole"
      }
    ]
  })
  tags = [
    {
      "key" : "key1",
      "value" : "value1"
    },
    {
      "key" : "key2",
      "value" : "value2"
    }
  ]
}