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
package provider

import (
	"fmt"
	"regexp"
	"terraform-provider-objectscale/internal/clientgen"
	"testing"

	"github.com/bytedance/mockey"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Test to Create and Update User Resource
func TestAccIamPolicyResource(t *testing.T) {
	defer testUserTokenCleanup(t)
	var apiMocker *mockey.Mocker
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create policy (CREATE FAIL)
			{
				PreConfig: func() {
					apiMocker = mockey.Mock((*clientgen.IamApiService).IamServiceCreatePolicyExecute).
						Return(nil, nil, fmt.Errorf("error")).Build()
				},
				Config: ProviderConfigForTesting + `
				resource "objectscale_iam_policy" "testacc_policy" {
					name = "testacc_policy"
					namespace = "ns1"
					description = "An example policy"
					policy_document = jsonencode({
  
						"Version": "2012-10-17",
						
						"Statement": [
							
							{
							
							"Action": [
								
								"s3:ListBucket",
								
								"s3:ListAllMyBuckets",
							
							],
							
							"Resource": "*",
							
							"Effect": "Allow",
							
							"Sid": "VisualEditor0"
							
							}
						
						]

						})
				}
				`,
				ExpectError: regexp.MustCompile(".*Error creating IAM Policy.*"),
			},
			// Step 2: Create policy (READ FAIL)
			{
				PreConfig: func() {
					apiMocker.UnPatch()
					apiMocker = mockey.Mock((*clientgen.IamApiService).IamServiceGetPolicyExecute).
						Return(nil, nil, fmt.Errorf("error")).Build()
				},
				Config: ProviderConfigForTesting + `
				resource "objectscale_iam_policy" "testacc_policy" {
					name = "testacc_policy"
					namespace = "ns1"
					description = "An example policy"
					policy_document = jsonencode({
  
						"Version": "2012-10-17",
						
						"Statement": [
							
							{
							
							"Action": [
								
								"s3:ListBucket",
								
								"s3:ListAllMyBuckets",
							
							],
							
							"Resource": "*",
							
							"Effect": "Allow",
							
							"Sid": "VisualEditor0"
							
							}
						
						]

						})
				}
				`,
				ExpectError: regexp.MustCompile(".*Error reading IAM Policy.*"),
			},
			// Step 3: Create policy (OK)
			{
				PreConfig: func() {
					apiMocker.UnPatch()
				},
				Config: ProviderConfigForTesting + `
				resource "objectscale_iam_policy" "testacc_policy" {
					name = "testacc_policy"
					namespace = "ns1"
					description = "An example policy"
					policy_document = jsonencode({
  
						"Version": "2012-10-17",
						
						"Statement": [
							
							{
							
							"Action": [
								
								"s3:ListBucket",
								
								"s3:ListAllMyBuckets",
							
							],
							
							"Resource": "*",
							
							"Effect": "Allow",
							
							"Sid": "VisualEditor0"
							
							}
						
						]

						})
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_policy.testacc_policy", "name", "testacc_policy"),
					resource.TestCheckResourceAttr("objectscale_iam_policy.testacc_policy", "namespace", "ns1"),
					resource.TestCheckResourceAttr("objectscale_iam_policy.testacc_policy", "description", "An example policy"),
				),
			},
			// Step 4: Update policy (FAIL)
			{
				PreConfig: func() {
					apiMocker = mockey.Mock((*clientgen.IamApiService).IamServiceCreatePolicyVersionExecute).
						Return(nil, nil, fmt.Errorf("error")).Build()
				},
				Config: ProviderConfigForTesting + `
				resource "objectscale_iam_policy" "testacc_policy" {
					name = "testacc_policy"
					namespace = "ns1"
					description = "An example policy"
					policy_document = jsonencode({
  
						"Version": "2012-10-17",
						
						"Statement": [
							
							{
							
							"Action": [
								
								"s3:ListBucket",
								
								"s3:ListAllMyBuckets",

        						"iam:GetUserPolicy"
							
							],
							
							"Resource": "*",
							
							"Effect": "Allow",
							
							"Sid": "VisualEditor0"
							
							}
						
						]

						})
				}
				`,
				ExpectError: regexp.MustCompile(".*Error creating new IAM Policy Version.*"),
			},
			// Step 5: Update policy (INVALID PARAMETER)
			{
				PreConfig: func() {
					apiMocker.UnPatch()
				},
				Config: ProviderConfigForTesting + `
				resource "objectscale_iam_policy" "testacc_policy" {
					name = "testacc_policy"
					namespace = "ns1"
					description = "An example policy UPDATED"
					policy_document = jsonencode({
  
						"Version": "2012-10-17",
						
						"Statement": [
							
							{
							
							"Action": [
								
								"s3:ListBucket",
								
								"s3:ListAllMyBuckets",

        						"iam:GetUserPolicy"
							
							],
							
							"Resource": "*",
							
							"Effect": "Allow",
							
							"Sid": "VisualEditor0"
							
							}
						
						]

						})
				}
				`,
				ExpectError: regexp.MustCompile(".*Invalid Update.*"),
			},
			// Step 6: Update policy (OK)
			{
				PreConfig: func() {
					apiMocker.UnPatch()
				},
				Config: ProviderConfigForTesting + `
				resource "objectscale_iam_policy" "testacc_policy" {
					name = "testacc_policy"
					namespace = "ns1"
					description = "An example policy"
					policy_document = jsonencode({
  
						"Version": "2012-10-17",
						
						"Statement": [
							
							{
							
							"Action": [
								
								"s3:ListBucket",
								
								"s3:ListAllMyBuckets",

        						"iam:GetUserPolicy"
							
							],
							
							"Resource": "*",
							
							"Effect": "Allow",
							
							"Sid": "VisualEditor0"
							
							}
						
						]

						})
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_policy.testacc_policy", "name", "testacc_policy"),
					resource.TestCheckResourceAttr("objectscale_iam_policy.testacc_policy", "namespace", "ns1"),
					resource.TestCheckResourceAttr("objectscale_iam_policy.testacc_policy", "description", "An example policy"),
				),
			},
			// Step 7: Update policy with deletion of redundant policy statements (OK)
			{
				Config: ProviderConfigForTesting + `
				resource "objectscale_iam_policy" "testacc_policy" {
					name = "testacc_policy"
					namespace = "ns1"
					description = "An example policy"
					policy_document = jsonencode({
  
						"Version": "2012-10-17",
						
						"Statement": [
							
							{
							
							"Action": [
								
								"s3:ListBucket",
			
        						"iam:GetUserPolicy"
							
							],
							
							"Resource": "*",
							
							"Effect": "Allow",
							
							"Sid": "VisualEditor0"
							
							}
						
						]

						})
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("objectscale_iam_policy.testacc_policy", "name", "testacc_policy"),
					resource.TestCheckResourceAttr("objectscale_iam_policy.testacc_policy", "namespace", "ns1"),
					resource.TestCheckResourceAttr("objectscale_iam_policy.testacc_policy", "description", "An example policy"),
				),
			},
			// Step 8: Import state
			{
				ResourceName: "objectscale_iam_policy.testacc_policy",
				// get resource arn for import : "policy_arn:namespace"
				ImportStateId: "urn:ecs:iam::ns1:policy/testacc_policy#ns1",
				ImportState:   true,
			},
		},
	})
	apiMocker.UnPatch()
}
