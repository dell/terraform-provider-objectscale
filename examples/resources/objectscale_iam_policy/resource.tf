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
resource "objectscale_iam_policy" "testacc_policy" {
  name        = "testacc_policy"
  namespace   = "ns1"
  description = "An example policy"
  policy_document = jsonencode({

    "Version" : "2012-10-17",

    "Statement" : [

      {

        "Action" : [

          "s3:ListBucket",

          "iam:GetUserPolicy"

        ],

        "Resource" : "*",

        "Effect" : "Allow",

        "Sid" : "VisualEditor0"

      }

    ]

  })
}