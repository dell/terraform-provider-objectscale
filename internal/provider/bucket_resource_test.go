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
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Test to Fetch Namespaces.
func TestAccBucketResourceNegative(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	defer testUserTokenCleanup(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Missing required attribute: name
			{
				Config: ProviderConfigForTesting + `
				resource "objectscale_bucket" "test" {
					owner = "admin1"
					replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"
					namespace = "ns1"
				}
				`,
				ExpectError: regexp.MustCompile(`(?i).*The argument "name" is required.*`),
			},
			// Missing required attribute: owner
			{
				Config: ProviderConfigForTesting + `
				resource "objectscale_bucket" "test" {
					name = "example-bucket-2"
					replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"
					namespace = "ns1"
				}
				`,
				ExpectError: regexp.MustCompile(`(?i).*The argument "owner" is required.*`),
			},
			// Missing required attribute: namespace
			{
				Config: ProviderConfigForTesting + `
				resource "objectscale_bucket" "test" {
					name = "example-bucket-2"
					owner = "admin1"
					replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"
				}
				`,
				ExpectError: regexp.MustCompile(`(?i).*The argument "namespace" is required.*`),
			},
			// Missing required attribute: replication_group
			{
				Config: ProviderConfigForTesting + `
				resource "objectscale_bucket" "test" {
					name = "example-bucket-2"
					owner = "admin1"
					namespace = "ns1"
				}
				`,
				ExpectError: regexp.MustCompile(`(?i).*The argument "replication_group" is required.*`),
			},
			// Invalid type for block_size (should be int, given string)
			{
				Config: ProviderConfigForTesting + `
				resource "objectscale_bucket" "test" {
					name = "example-bucket-2"
					owner = "admin1"
					namespace = "ns1"
					replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"
					block_size = "not-an-int"
				}
				`,
				ExpectError: regexp.MustCompile(`Incorrect attribute value type`),
			},
		},
	})
}

func TestAccBucketResourcePositive(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with unit tests because it will try to create the context")
	}
	defer testUserTokenCleanup(t)

	resourceName := "objectscale_bucket.test"

	// Only the following metadata fields are supported: "created_by" and "purpose"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `resource "objectscale_bucket" "example_bucket" {
						# Required: Owner of the bucket
						owner = "admin1"

						# Required: Name of the bucket
						name = "example-bucket-test"

						# Required: Virtual pool URL associated with the bucket (Get it using Replication Datasource)
						replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"

						# Required: Namespace for bucket isolation
						namespace = "ns1"

						# # Optional: Size of each block in bytes
						block_size = 4096

						# # Optional: Size threshold for notifications
						notification_size = 1024

						# # Optional: Enable filesystem interface
						# filesystem_enabled = true

						# # Optional: Type of bucket head (e.g., object)
						# head_type = "object"

						# # Optional: Key-value tags for bucket
						tag = [
							{
							key   = "Env"
							value = "Production"
							}
						]

						# # Optional: Enable server-side encryption
						is_encryption_enabled = true

						# # Optional: Default group permissions
						default_group_file_read_permission    = "true"
						# default_group_file_write_permission   = "false"
						# default_group_file_execute_permission = "false"
						# default_group_dir_read_permission     = "true"
						# default_group_dir_write_permission    = "false"
						# default_group_dir_execute_permission  = "true"

						# # Optional: Default group name
						default_group = "default-group"

						# # Optional: Auto-commit period in seconds
						auto_commit_period = 600

						# # Optional: Retention period in days
						retention = 30

						# # Optional: Allow stale reads
						# is_stale_allowed = false

						# # Optional: Allow object lock with ADO
						# is_object_lock_with_ado_allowed = false

						# # Optional: Enable TSO read-only mode
						# is_tso_read_only = false

						is_enabled = true

						# # Optional: Metadata search configuration
						metadata = [
							{
							type     = "System"
							name     = "CreateTime"
							datatype = "Datetime"
							},
							{
							type     = "User"
							name     = "Krunal"
							datatype = "String"
							}
						]

						# # Optional: Metadata tokens for advanced search
						# metadata_tokens = false

						# # Optional: Retention governance settings
						min_max_governor =  {
							enforce_retention           = true
							minimum_fixed_retention     = 10
							maximum_fixed_retention     = 365
							minimum_variable_retention  = 5
							maximum_variable_retention  = 180
						}

						# # Optional: Days after which audited delete expires
						audit_delete_expiration = 90

						# # Optional: Enable object lock
						is_object_lock_enabled = true

						# # Optional: Enable advanced metadata search
						enable_advanced_metadata_search = true

						# # Optional: Advanced metadata search target name
						advanced_metadata_search_target_name = "search-target"

						# # Optional: Advanced metadata search target stream
						advanced_metadata_search_target_stream = "stream-1"

						# # Optional: Enable local metadata reads
						# local_object_metadata_reads = true

						# # Optional: Versioning status (Enabled/Suspended)
						# versioning_status = "Enabled"
						}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "example-bucket-positive"),
					resource.TestCheckResourceAttr(resourceName, "owner", "admin1"),
					resource.TestCheckResourceAttr(resourceName, "namespace", "ns1"),
					resource.TestCheckResourceAttr(resourceName, "replication_group", "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"),
					resource.TestCheckResourceAttr(resourceName, "block_size", "4096"),
				),
			},
			// Update block_size to test update, expect error
			{
				Config: ProviderConfigForTesting + `
				resource "objectscale_bucket" "test" {
					name = "example-bucket-positive"
					owner = "root"
					namespace = "ns1"
					replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"
					block_size = 4096
				}
				`,
				ExpectError: regexp.MustCompile(`Update Bucket operation is not supported`),
			},
		},
	})
}

func TestAccBucketResourceImport(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with unit tests because it will try to create the context")
	}
	defer testUserTokenCleanup(t)

	resourceName := "objectscale_bucket.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
				resource "objectscale_bucket" "test" {
					name = "example-bucket-positive-1"
					owner = "admin1"
					namespace = "ns1"
					replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"
					block_size = 4096
				}
				`,
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateId:     "example-bucket:ns1",
				ImportStateVerify: true,
				ExpectError:       regexp.MustCompile(`Error importing Bucket`),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateId:     "example-bucket-positive-invalid:ns1",
				ImportStateVerify: true,
				ExpectError:       regexp.MustCompile(`Error Reading Buckets`),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateId:     "example-bucket-positive-1:ns1",
				ImportStateVerify: true,
			},
		},
	})
}
