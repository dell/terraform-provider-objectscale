// Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.
//
// Licensed under the Mozilla Public License Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Main Acceptance Test: all scenarios.
func TestAccBucketDataSource_PositiveScenarios(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{

			// 1. Fetch a single group using bucket_name_prefix filter
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_bucket" "one" {
						namespace  = "ns1"
						bucket_name_prefix = "bucket"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.objectscale_bucket.one", "id",
					),
					resource.TestCheckResourceAttrSet(
						"data.objectscale_bucket.one", "buckets.0.name",
					),
				),
			},

			// 2. Fetch all groups (no group_name specified)
			{
				Config: ProviderConfigForTesting + `
					data "objectscale_bucket" "all" {
						namespace = "ns1"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.objectscale_bucket.all", "buckets.0.name",
					),
				),
			},
		},
	})
}

// Error Scenarios for IAM Groups Data Source.
func TestAccBucketDataSource_ErrorScenarios(t *testing.T) {
	defer testUserTokenCleanup(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,

		// IMPORTANT → prevent destroy-phase login failures
		PreventPostDestroyRefresh: true,

		Steps: []resource.TestStep{

			// Missing namespace
			{
				Config: ProviderConfigForTesting + `
			        data "objectscale_bucket" "missing_ns" {
			        }
			    `,
				ExpectError: regexp.MustCompile(`namespace`),
			},

			//Invalid namespace → error
			{
				Config: ProviderConfigForTesting + `
			        data "objectscale_bucket" "bad_bucket_name_prefix" {
			            namespace  = "INVALID"
			        }
			    `,
				ExpectError: regexp.MustCompile(`Error Reading Buckets`),
			},

			// Nonexistent Bucket Name → triggers "Error Listing Bucket"
			{
				Config: ProviderConfigForTesting + `
			        data "objectscale_bucket" "invalid_bucket_name_prefix" {
			            namespace  = "ns1"
			            bucket_name_prefix  = "INVALID"
			        }
			    `,
				ExpectError: regexp.MustCompile(`No buckets found`),
			},
		},
	})
}
