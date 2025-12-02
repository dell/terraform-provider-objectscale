package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Main Acceptance Test: all scenarios.
func TestAccBucketDataSource_PositiveScenarios(t *testing.T) {

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
				ExpectError: regexp.MustCompile(`Namespace INVALID does not exist`),
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
