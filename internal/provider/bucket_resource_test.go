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
			// Validation
			{
				Config: ProviderConfigForTesting + `
				resource "objectscale_bucket" "test" {
					name = "example-bucket-2"
					owner = "admin1"
					namespace = "ns1"
					replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"
					default_group_file_read_permission = "true"
					filesystem_enabled = true
				}
				`,
				ExpectError: regexp.MustCompile(`Missing Default Group`),
			},
			// Validation
			{
				Config: ProviderConfigForTesting + `
				resource "objectscale_bucket" "test" {
					name = "example-bucket-2"
					owner = "admin1"
					namespace = "ns1"
					replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"
					search_metadata = [
						{
						type     = "User"
						name     = ""
						datatype = "string"
						}
					]
				}
				`,
				ExpectError: regexp.MustCompile(`Empty Metadata Name`),
			},
			// Validation
			{
				Config: ProviderConfigForTesting + `
				resource "objectscale_bucket" "test" {
					name = "example-bucket-2"
					owner = "admin1"
					namespace = "ns1"
					replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"
					search_metadata = [
						{
						type     = "User"
						name     = "adad"
						datatype = "string"
						}
					]
				}
				`,
				ExpectError: regexp.MustCompile(`Invalid Metadata Name Prefix`),
			},
			// Validation
			{
				Config: ProviderConfigForTesting + `
				resource "objectscale_bucket" "test" {
					name = "example-bucket-2"
					owner = "admin1"
					namespace = "ns1"
					replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"
					search_metadata = [
						{
						type     = "User"
						name     = "x-amz-meta-"
						datatype = "string"
						}
					]
				}
				`,
				ExpectError: regexp.MustCompile(`Empty Metadata Name Suffix`),
			},
			// Validation
			{
				Config: ProviderConfigForTesting + `
				resource "objectscale_bucket" "test" {
					name = "example-bucket-2"
					owner = "admin1"
					namespace = "ns1"
					replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"
					search_metadata = [
						{
						type     = "User"
						name     = "x-amz-meta-ABCD"
						datatype = "string"
						}
					]
				}
				`,
				ExpectError: regexp.MustCompile(`Invalid Metadata Name`),
			},
			// Validation
			{
				Config: ProviderConfigForTesting + `
				resource "objectscale_bucket" "test" {
					name = "example-bucket-2"
					owner = "admin1"
					namespace = "ns1"
					replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"
					search_metadata = [
						{
						type     = "System"
						name     = "x-amz-meta-ABCD"
						datatype = "string"
						}
					]
				}
				`,
				ExpectError: regexp.MustCompile(`Invalid System Metadata Name`),
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

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creation Test Step
			{
				Config: ProviderConfigForTesting + `resource "objectscale_bucket" "test" {
						# Required: Owner of the bucket
						owner = "admin1"

						# Required: Name of the bucket
						name = "example-bucket-positive"

						# Required: Virtual pool URL associated with the bucket (Get it using Replication Datasource)
						replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"

						# Required: Namespace for bucket isolation
						namespace = "ns1"

						# Optional: Size of each block in bytes
						block_size = 4096

						# Optional: Size threshold for notifications
						notification_size = 2024

						# Optional: Enable filesystem interface
						# filesystem_enabled = true

						# Optional: Key-value tags for bucket
						tag = [
							{
							key   = "Env"
							value = "Production"
							},
							{
							key   = "Tag"
							value = "Production"
							}
						]

						# Optional: Enable server-side encryption
						# is_encryption_enabled = true

						# Optional: Default group permissions
						# default_group_file_read_permission = "true"
						# default_group_file_write_permission   = "false"
						# default_group_file_execute_permission = "false"
						# default_group_dir_read_permission     = "true"
						# default_group_dir_write_permission    = "false"
						# default_group_dir_execute_permission  = "true"

						# Optional: Default group name
						# default_group = "default"

						# Optional: Auto-commit period in seconds
						auto_commit_period = 28

						# Optional: Retention period in days
						retention = 30

						default_retention = 30

						# Optional: Allow stale reads
						# is_stale_allowed = true

						# Optional: Allow object lock with ADO
						is_object_lock_with_ado_allowed = true

						# Optional: Enable TSO read-only mode
						# is_tso_read_only = true

						# Optional: Enable metadata search
						is_metadata_enabled = true

						# Optional: Metadata search configuration
						search_metadata = [
							{
							type     = "System"
							name     = "CreateTime"
							datatype = "datetime"
							},
							{
							type     = "User"
							name     = "x-amz-meta-abc"
							datatype = "string"
							}
						]

						# Optional: Retention governance settings
						min_max_governor = {
							enforce_retention          = true
							minimum_fixed_retention    = 10
							maximum_fixed_retention    = 365
							minimum_variable_retention = 5
							maximum_variable_retention = 180
						}

						bucket_policy = jsonencode({
							Version = "2012-10-17",
							Id      = "null",
							Statement = [
							{
								Sid       = "AllowPublicRead",
								Effect    = "Allow",
								Principal = "*",
								Action    = "s3:GetObject",
								Resource  = "arn:aws:s3:::example-bucket-positive/*"
							}
							]
						})

					    user_acl = [
							{
							name       = "admin1"
							permission = ["full_control"]
							},
							{
							name       = "root"
							permission = ["read", "write"]
							}
						]	

						group_acl = [
							{
							name       = "all_users"
							permission = ["full_control"]
							}
						]

						custom_group_acl = [
							{
							name       = "testing"
							permission = ["full_control"]
							}
						]
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
			// Update Failure Test Step
			{
				// Update block_size to test update, expect error
				Config: ProviderConfigForTesting + `resource "objectscale_bucket" "test" {
						# Required: Owner of the bucket
						owner = "admin1"

						# Required: Name of the bucket
						name = "example-bucket-update"

						# Required: Virtual pool URL associated with the bucket (Get it using Replication Datasource)
						replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"

						# Required: Namespace for bucket isolation
						namespace = "ns1"

						# Optional: Size of each block in bytes
						block_size = 4096

						# Optional: Size threshold for notifications
						notification_size = 2024

						# Optional: Enable filesystem interface
						# filesystem_enabled = true

						# Optional: Key-value tags for bucket
						tag = [
							{
							key   = "Env"
							value = "Production"
							}
						]

						# Optional: Enable server-side encryption
						# is_encryption_enabled = true

						# Optional: Default group permissions
						# default_group_file_read_permission = "true"
						# default_group_file_write_permission   = "false"
						# default_group_file_execute_permission = "false"
						# default_group_dir_read_permission     = "true"
						# default_group_dir_write_permission    = "false"
						# default_group_dir_execute_permission  = "true"

						# Optional: Default group name
						# default_group = "default"

						# Optional: Auto-commit period in seconds
						auto_commit_period = 28

						# Optional: Retention period in days
						retention = 30

						# Optional: Allow stale reads
						# is_stale_allowed = true

						# Optional: Allow object lock with ADO
						is_object_lock_with_ado_allowed = true

						# Optional: Enable TSO read-only mode
						# is_tso_read_only = true

						# Optional: Enable metadata search
						is_metadata_enabled = true

						# Optional: Metadata search configuration
						search_metadata = [
							{
							type     = "System"
							name     = "CreateTime"
							datatype = "datetime"
							},
							{
							type     = "User"
							name     = "x-amz-meta-abc"
							datatype = "string"
							}
						]

						# Optional: Retention governance settings
						min_max_governor = {
							enforce_retention          = true
							minimum_fixed_retention    = 10
							maximum_fixed_retention    = 365
							minimum_variable_retention = 5
							maximum_variable_retention = 180
						}


						user_acl = [
								{
								name       = "admin1"
								permission = ["full_control"]
								},
								{
								name       = "root"
								permission = ["read", "write"]
								}
						]
			    }
				`,
				ExpectError: regexp.MustCompile(`Immutable Field Change Detected`),
			},
			// Update Successful Test Step
			{
				// Update block_size to test update, expect error
				Config: ProviderConfigForTesting + `resource "objectscale_bucket" "test" {
						# Required: Owner of the bucket
						owner = "root"

						# Required: Name of the bucket
						name = "example-bucket-positive"

						# Required: Virtual pool URL associated with the bucket (Get it using Replication Datasource)
						replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"

						# Required: Namespace for bucket isolation
						namespace = "ns1"

						# Optional: Size of each block in bytes
						block_size = 5096

						# Optional: Size threshold for notifications
						notification_size = 2024

						# Optional: Enable filesystem interface
						# filesystem_enabled = true

						# Optional: Key-value tags for bucket
						tag = [
							{
							key   = "Env"
							value = "Prod"
							},
							{
							key   = "test"
							value = "Production"
							},
							{
							key   = "Dev"
							value = "Production"
							}
						]

						# Optional: Enable server-side encryption
						# is_encryption_enabled = true

						# Optional: Default group permissions
						# default_group_file_read_permission = "true"
						# default_group_file_write_permission   = "false"
						# default_group_file_execute_permission = "false"
						# default_group_dir_read_permission     = "true"
						# default_group_dir_write_permission    = "false"
						# default_group_dir_execute_permission  = "true"

						# Optional: Default group name
						# default_group = "default"

						# Optional: Auto-commit period in seconds
						auto_commit_period = 28

						# Optional: Retention period in days
						retention = 30

						default_retention = 30

						# Optional: Allow stale reads
						# is_stale_allowed = true

						# Optional: Allow object lock with ADO
						is_object_lock_with_ado_allowed = true

						# Optional: Enable TSO read-only mode
						# is_tso_read_only = true

						# Optional: Enable metadata search
						is_metadata_enabled = true

						# Optional: Metadata search configuration
						search_metadata = [
							{
							type     = "System"
							name     = "CreateTime"
							datatype = "datetime"
							},
							{
							type     = "User"
							name     = "x-amz-meta-abc"
							datatype = "string"
							}
						]

						versioning_status = "Enabled"

						# Optional: Retention governance settings
						min_max_governor = {
							enforce_retention          = true
							minimum_fixed_retention    = 10
							maximum_fixed_retention    = 365
							minimum_variable_retention = 5
							maximum_variable_retention = 180
						}

						bucket_policy = jsonencode({
							Version = "2012-10-17",
							Id      = "null",
							Statement = [
							{
								Sid       = "AllowPublicWrite",
								Effect    = "Allow",
								Principal = "*",
								Action    = "s3:GetObject",
								Resource  = "arn:aws:s3:::example-bucket-positive/*"
							}
							]
						})
						
						user_acl = [
								{
								name       = "admin1"
								permission = ["full_control"]
								},
								{
								name       = "root"
								permission = ["read", "write"]
								}
						]

						group_acl = [
							{
							name       = "log_delivery"
							permission = ["full_control"]
							}
						]

						custom_group_acl = [
							{
							name       = "development"
							permission = ["full_control"]
							}
						]
			    }
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "example-bucket-positive"),
					resource.TestCheckResourceAttr(resourceName, "owner", "root"),
					resource.TestCheckResourceAttr(resourceName, "namespace", "ns1"),
					resource.TestCheckResourceAttr(resourceName, "replication_group", "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"),
					resource.TestCheckResourceAttr(resourceName, "block_size", "5096"),
				),
			},
		},
	})
}

func TestAccBucketResourcePositive2(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with unit tests because it will try to create the context")
	}
	defer testUserTokenCleanup(t)

	resourceName := "objectscale_bucket.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creation Test Step Negative for Bucket Policy
			{
				Config: ProviderConfigForTesting + `resource "objectscale_bucket" "test" {
						# Required: Owner of the bucket
						owner = "admin1"

						# Required: Name of the bucket
						name = "example-bucket-positive-2"

						# Required: Virtual pool URL associated with the bucket (Get it using Replication Datasource)
						replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"

						# Required: Namespace for bucket isolation
						namespace = "ns1"

						# Optional: Size of each block in bytes
						block_size = 4096

						# Optional: Size threshold for notifications
						notification_size = 2024

						# Optional: Enable filesystem interface
						filesystem_enabled = true

						# Optional: Key-value tags for bucket
						tag = [
							{
							key   = "Env"
							value = "Production"
							},
							{
							key   = "Tag"
							value = "Production"
							}
						]

						# Optional: Enable server-side encryption
						# is_encryption_enabled = true

						# Optional: Default group permissions
						default_group_file_read_permission = "true"
						# default_group_file_write_permission   = "false"
						# default_group_file_execute_permission = "false"
						# default_group_dir_read_permission     = "true"
						# default_group_dir_write_permission    = "false"
						# default_group_dir_execute_permission  = "true"

						# Optional: Default group name
						default_group = "default"

						# Optional: Auto-commit period in seconds
						auto_commit_period = 28

						# Optional: Retention period in days
						retention = 30

						# Optional: Allow stale reads
						is_stale_allowed = true

						# Optional: Allow object lock with ADO
						is_object_lock_with_ado_allowed = false

						# Optional: Enable TSO read-only mode
						is_tso_read_only = true

						# Optional: Enable metadata search
						is_metadata_enabled = true

						# Optional: Metadata search configuration
						search_metadata = [
							{
							type     = "System"
							name     = "CreateTime"
							datatype = "datetime"
							},
							{
							type     = "User"
							name     = "x-amz-meta-abc"
							datatype = "string"
							}
						]

						bucket_policy = jsonencode({
							Version = "2012-10-17",
							Id      = "null",
							Statement = [
							{
								Sid       = "AllowPublicRead",
								Effect    = "Allow",
								Principal = "*",
								Action    = "s3:GetObject",
								Resource  = "arn:aws:s3:::example-bucket-2/*"
							}
							]
						})
				}
				`,
				ExpectError: regexp.MustCompile("Error setting bucket policy"),
			},
			// Creation Test Step Negative for Bucket Policy
			{
				Config: ProviderConfigForTesting + `resource "objectscale_bucket" "test" {
						# Required: Owner of the bucket
						owner = "admin1"

						# Required: Name of the bucket
						name = "example-bucket-positive-2"

						# Required: Virtual pool URL associated with the bucket (Get it using Replication Datasource)
						replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"

						# Required: Namespace for bucket isolation
						namespace = "ns1"

						# Optional: Size of each block in bytes
						block_size = 4096

						# Optional: Size threshold for notifications
						notification_size = 2024

						# Optional: Enable filesystem interface
						# filesystem_enabled = true

						# Optional: Key-value tags for bucket
						tag = [
							{
							key   = "Env"
							value = "Production"
							},
							{
							key   = "Tag"
							value = "Production"
							}
						]

						# Optional: Enable server-side encryption
						# is_encryption_enabled = true

						# Optional: Default group permissions
						# default_group_file_read_permission = "true"
						# default_group_file_write_permission   = "false"
						# default_group_file_execute_permission = "false"
						# default_group_dir_read_permission     = "true"
						# default_group_dir_write_permission    = "false"
						# default_group_dir_execute_permission  = "true"

						# Optional: Default group name
						# default_group = "default"

						# Optional: Auto-commit period in seconds
						auto_commit_period = 28

						# Optional: Retention period in days
						retention = 30

						# Optional: Allow stale reads
						is_stale_allowed = true

						# Optional: Allow object lock with ADO
						is_object_lock_with_ado_allowed = false

						# Optional: Enable TSO read-only mode
						is_tso_read_only = true

						# Optional: Enable metadata search
						is_metadata_enabled = true

						# Optional: Metadata search configuration
						search_metadata = [
							{
							type     = "System"
							name     = "CreateTime"
							datatype = "datetime"
							},
							{
							type     = "User"
							name     = "x-amz-meta-abc"
							datatype = "string"
							}
						]

						group_acl = [
							{
							name       = "all_users"
							permission = ["invalid"]
							}
						]
				}
				`,
				ExpectError: regexp.MustCompile("Error setting bucket ACL"),
			},
			// Creation Test Step
			{
				Config: ProviderConfigForTesting + `resource "objectscale_bucket" "test" {
						# Required: Owner of the bucket
						owner = "admin1"

						# Required: Name of the bucket
						name = "example-bucket-positive-2"

						# Required: Virtual pool URL associated with the bucket (Get it using Replication Datasource)
						replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"

						# Required: Namespace for bucket isolation
						namespace = "ns1"

						# Optional: Size of each block in bytes
						block_size = 4096

						# Optional: Size threshold for notifications
						notification_size = 2024

						# Optional: Enable filesystem interface
						filesystem_enabled = true

						# Optional: Key-value tags for bucket
						tag = [
							{
							key   = "Env"
							value = "Production"
							},
							{
							key   = "Tag"
							value = "Production"
							}
						]

						# Optional: Enable server-side encryption
						# is_encryption_enabled = true

						# Optional: Default group permissions
						default_group_file_read_permission = "true"
						# default_group_file_write_permission   = "false"
						# default_group_file_execute_permission = "false"
						# default_group_dir_read_permission     = "true"
						# default_group_dir_write_permission    = "false"
						# default_group_dir_execute_permission  = "true"

						# Optional: Default group name
						default_group = "default"

						# Optional: Auto-commit period in seconds
						auto_commit_period = 28

						# Optional: Retention period in days
						retention = 30

						# Optional: Allow stale reads
						is_stale_allowed = true

						# Optional: Allow object lock with ADO
						is_object_lock_with_ado_allowed = true

						# Optional: Enable TSO read-only mode
						is_tso_read_only = true

						# Optional: Enable metadata search
						is_metadata_enabled = true

						# Optional: Metadata search configuration
						search_metadata = [
							{
							type     = "System"
							name     = "CreateTime"
							datatype = "datetime"
							},
							{
							type     = "User"
							name     = "x-amz-meta-abc"
							datatype = "string"
							}
						]
						
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "example-bucket-positive-2"),
					resource.TestCheckResourceAttr(resourceName, "owner", "admin1"),
					resource.TestCheckResourceAttr(resourceName, "namespace", "ns1"),
					resource.TestCheckResourceAttr(resourceName, "replication_group", "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"),
					resource.TestCheckResourceAttr(resourceName, "block_size", "4096"),
				),
			},
			// Update Failure Test Step
			{
				Config: ProviderConfigForTesting + `resource "objectscale_bucket" "test" {
						# Required: Owner of the bucket
						owner = "root"

						# Required: Name of the bucket
						name = "example-bucket-positive-2"

						# Required: Virtual pool URL associated with the bucket (Get it using Replication Datasource)
						replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"

						# Required: Namespace for bucket isolation
						namespace = "ns1"

						# Optional: Size of each block in bytes
						block_size = 5096

						# Optional: Size threshold for notifications
						notification_size = 2024

						# Optional: Enable filesystem interface
						filesystem_enabled = true

						# Optional: Key-value tags for bucket
						tag = [
							{
							key   = "Env"
							value = "Production"
							},
							{
							key   = "Tag"
							value = "Production"
							}
						]

						# Optional: Enable server-side encryption
						# is_encryption_enabled = true

						# Optional: Default group permissions
						default_group_file_read_permission = "true"
						# default_group_file_write_permission   = "false"
						default_group_file_execute_permission = "true"
						# default_group_dir_read_permission     = "true"
						# default_group_dir_write_permission    = "false"
						# default_group_dir_execute_permission  = "true"

						# Optional: Default group name
						default_group = "default"

						# Optional: Auto-commit period in seconds
						auto_commit_period = 25

						# Optional: Retention period in days
						retention = 28

						# Optional: Allow stale reads
						is_stale_allowed = true

						# Optional: Allow object lock with ADO
						is_object_lock_with_ado_allowed = false

						# Optional: Enable TSO read-only mode
						is_tso_read_only = true

						# Optional: Enable metadata search
						is_metadata_enabled = true

						# Optional: Metadata search configuration
						search_metadata = [
							{
							type     = "System"
							name     = "CreateTime"
							datatype = "datetime"
							},
							{
							type     = "User"
							name     = "x-amz-meta-abc"
							datatype = "string"
							}
						]

						
			    }
				`,
				ExpectError: regexp.MustCompile("Immutable Field Change Detected"),
			},
			// Update Failure Test Step
			{
				Config: ProviderConfigForTesting + `resource "objectscale_bucket" "test" {
						# Required: Owner of the bucket
						owner = "root"

						# Required: Name of the bucket
						name = "example-bucket-positive-2"

						# Required: Virtual pool URL associated with the bucket (Get it using Replication Datasource)
						replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"

						# Required: Namespace for bucket isolation
						namespace = "ns1"

						# Optional: Size of each block in bytes
						block_size = 5096

						# Optional: Size threshold for notifications
						notification_size = 2024

						is_object_lock_enabled = true

						# Optional: Enable filesystem interface
						filesystem_enabled = true

						# Optional: Key-value tags for bucket
						tag = [
							{
							key   = "Env"
							value = "Production"
							},
							{
							key   = "Tag"
							value = "Production"
							}
						]

						# Optional: Enable server-side encryption
						# is_encryption_enabled = true

						# Optional: Default group permissions
						default_group_file_read_permission = "true"
						# default_group_file_write_permission   = "false"
						default_group_file_execute_permission = "true"
						# default_group_dir_read_permission     = "true"
						# default_group_dir_write_permission    = "false"
						# default_group_dir_execute_permission  = "true"

						# Optional: Default group name
						default_group = "default"

						# Optional: Auto-commit period in seconds
						auto_commit_period = 25

						# Optional: Retention period in days
						retention = 28

						# Optional: Allow stale reads
						is_stale_allowed = true

						# Optional: Allow object lock with ADO
						is_object_lock_with_ado_allowed = true

						# Optional: Enable TSO read-only mode
						is_tso_read_only = true

						# Optional: Enable metadata search
						is_metadata_enabled = true

						# Optional: Metadata search configuration
						search_metadata = [
							{
							type     = "System"
							name     = "CreateTime"
							datatype = "datetime"
							},
							{
							type     = "User"
							name     = "x-amz-meta-abc"
							datatype = "string"
							}
						]

						
			    }
				`,
				ExpectError: regexp.MustCompile("Error updating IsObjectLockEnabled"),
			},
			// Update Failure Test Step
			{
				Config: ProviderConfigForTesting + `resource "objectscale_bucket" "test" {
						# Required: Owner of the bucket
						owner = "root"

						# Required: Name of the bucket
						name = "example-bucket-positive-2"

						# Required: Virtual pool URL associated with the bucket (Get it using Replication Datasource)
						replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"

						# Required: Namespace for bucket isolation
						namespace = "ns1"

						# Optional: Size of each block in bytes
						block_size = 5096

						# Optional: Size threshold for notifications
						notification_size = 2024

						# Optional: Enable filesystem interface
						filesystem_enabled = true

						# Optional: Key-value tags for bucket
						tag = [
							{
							key   = "Env"
							value = "Production"
							},
							{
							key   = "Tag"
							value = "Production"
							}
						]

						# Optional: Enable server-side encryption
						# is_encryption_enabled = true

						# Optional: Default group permissions
						default_group_file_read_permission = "true"
						# default_group_file_write_permission   = "false"
						default_group_file_execute_permission = "true"
						# default_group_dir_read_permission     = "true"
						# default_group_dir_write_permission    = "false"
						# default_group_dir_execute_permission  = "true"

						# Optional: Default group name
						default_group = "default"

						# Optional: Auto-commit period in seconds
						auto_commit_period = 25

						# Optional: Retention period in days
						retention = 28

						# Optional: Allow stale reads
						is_stale_allowed = true

						# Optional: Allow object lock with ADO
						is_object_lock_with_ado_allowed = true

						# Optional: Enable TSO read-only mode
						is_tso_read_only = true

						# Optional: Enable metadata search
						is_metadata_enabled = true

						# Optional: Metadata search configuration
						search_metadata = [
							{
							type     = "System"
							name     = "CreateTime"
							datatype = "datetime"
							},
							{
							type     = "User"
							name     = "x-amz-meta-abc"
							datatype = "string"
							}
						]

						# Optional: Enable advanced metadata search
						enable_advanced_metadata_search = true

						# Optional: Advanced metadata search target name
						advanced_metadata_search_target_name = "search-target"

						# Optional: Advanced metadata search target stream
						advanced_metadata_search_target_stream = "stream-1"

						
			    }
				`,
				ExpectError: regexp.MustCompile("Error Updating AdvancedMetadataSearch Status"),
			},
			// Update Failure Test Step
			{
				Config: ProviderConfigForTesting + `resource "objectscale_bucket" "test" {
						# Required: Owner of the bucket
						owner = "root"

						# Required: Name of the bucket
						name = "example-bucket-positive-2"

						# Required: Virtual pool URL associated with the bucket (Get it using Replication Datasource)
						replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"

						# Required: Namespace for bucket isolation
						namespace = "ns1"

						# Optional: Size of each block in bytes
						block_size = 5096

						# Optional: Size threshold for notifications
						notification_size = 2024

						# Optional: Enable filesystem interface
						filesystem_enabled = true

						# Optional: Key-value tags for bucket
						tag = [
							{
							key   = "Env"
							value = "Production"
							},
							{
							key   = "Tag"
							value = "Production"
							}
						]

						# Optional: Enable server-side encryption
						# is_encryption_enabled = true

						# Optional: Default group permissions
						default_group_file_read_permission = "true"
						# default_group_file_write_permission   = "false"
						default_group_file_execute_permission = "true"
						# default_group_dir_read_permission     = "true"
						# default_group_dir_write_permission    = "false"
						# default_group_dir_execute_permission  = "true"

						# Optional: Default group name
						default_group = "default"

						# Optional: Auto-commit period in seconds
						auto_commit_period = 25

						# Optional: Retention period in days
						retention = 28

						# Optional: Allow stale reads
						is_stale_allowed = true

						# Optional: Allow object lock with ADO
						is_object_lock_with_ado_allowed = true

						# Optional: Enable TSO read-only mode
						is_tso_read_only = false

						# Optional: Enable metadata search
						is_metadata_enabled = true

						local_object_metadata_reads = true

						# Optional: Metadata search configuration
						search_metadata = [
							{
							type     = "System"
							name     = "CreateTime"
							datatype = "datetime"
							},
							{
							type     = "User"
							name     = "x-amz-meta-abc"
							datatype = "string"
							}
						]

						
			    }
				`,
				ExpectError: regexp.MustCompile("Error updating LocalObjectMetadataReads"),
			},
			// Update Failure Test Step
			{
				Config: ProviderConfigForTesting + `resource "objectscale_bucket" "test" {
						# Required: Owner of the bucket
						owner = "root"

						# Required: Name of the bucket
						name = "example-bucket-positive-2"

						# Required: Virtual pool URL associated with the bucket (Get it using Replication Datasource)
						replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"

						# Required: Namespace for bucket isolation
						namespace = "ns1"

						# Optional: Size of each block in bytes
						block_size = 5096

						# Optional: Size threshold for notifications
						notification_size = 2024

						# Optional: Enable filesystem interface
						filesystem_enabled = true

						# Optional: Key-value tags for bucket
						tag = [
							{
							key   = "Env"
							value = "Production"
							},
							{
							key   = "Tag"
							value = "Production"
							}
						]

						# Optional: Enable server-side encryption
						# is_encryption_enabled = true

						# Optional: Default group permissions
						default_group_file_read_permission = "true"
						# default_group_file_write_permission   = "false"
						default_group_file_execute_permission = "true"
						# default_group_dir_read_permission     = "true"
						# default_group_dir_write_permission    = "false"
						# default_group_dir_execute_permission  = "true"

						# Optional: Default group name
						default_group = "default"

						# Optional: Auto-commit period in seconds
						auto_commit_period = 25

						# Optional: Retention period in days
						retention = 28

						# Optional: Allow stale reads
						is_stale_allowed = true

						# Optional: Allow object lock with ADO
						is_object_lock_with_ado_allowed = true

						# Optional: Enable TSO read-only mode
						is_tso_read_only = false

						# Optional: Enable metadata search
						is_metadata_enabled = true

						# Optional: Metadata search configuration
						search_metadata = [
							{
							type     = "System"
							name     = "CreateTime"
							datatype = "datetime"
							},
							{
							type     = "User"
							name     = "x-amz-meta-abc"
							datatype = "string"
							}
						]

						
			    }
				`,
				ExpectError: regexp.MustCompile("Error updating IsStaleAllowed"),
			},
			// Update Failure Test Step
			{
				Config: ProviderConfigForTesting + `resource "objectscale_bucket" "test" {
						# Required: Owner of the bucket
						owner = "root"

						# Required: Name of the bucket
						name = "example-bucket-positive-2"

						# Required: Virtual pool URL associated with the bucket (Get it using Replication Datasource)
						replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"

						# Required: Namespace for bucket isolation
						namespace = "ns1"

						# Optional: Size of each block in bytes
						block_size = 5096

						# Optional: Size threshold for notifications
						notification_size = 2024

						# Optional: Enable filesystem interface
						filesystem_enabled = true

						# Optional: Key-value tags for bucket
						tag = [
							{
							key   = "Env"
							value = "Production"
							},
							{
							key   = "Tag"
							value = "Production"
							}
						]

						# Optional: Enable server-side encryption
						# is_encryption_enabled = true

						# Optional: Default group permissions
						default_group_file_read_permission = "true"
						# default_group_file_write_permission   = "false"
						default_group_file_execute_permission = "true"
						# default_group_dir_read_permission     = "true"
						# default_group_dir_write_permission    = "false"
						# default_group_dir_execute_permission  = "true"

						# Optional: Default group name
						default_group = "default"

						# Optional: Auto-commit period in seconds
						auto_commit_period = 25

						# Optional: Retention period in days
						retention = 28

						# Optional: Allow stale reads
						is_stale_allowed = true

						# Optional: Allow object lock with ADO
						is_object_lock_with_ado_allowed = true

						default_retention = 30

						# Optional: Enable TSO read-only mode
						is_tso_read_only = true

						# Optional: Enable metadata search
						is_metadata_enabled = true

						# Optional: Metadata search configuration
						search_metadata = [
							{
							type     = "System"
							name     = "CreateTime"
							datatype = "datetime"
							},
							{
							type     = "User"
							name     = "x-amz-meta-abc"
							datatype = "string"
							}
						]

						

						min_max_governor = {
							enforce_retention          = true
							minimum_fixed_retention    = 10
							maximum_fixed_retention    = 365
							minimum_variable_retention = 5
							maximum_variable_retention = 180
						}
						
			    }
				`,
				ExpectError: regexp.MustCompile("Default Retention and Retention Mismatch"),
			},
			// Update Successful Test Step
			{
				Config: ProviderConfigForTesting + `resource "objectscale_bucket" "test" {
						# Required: Owner of the bucket
						owner = "root"

						# Required: Name of the bucket
						name = "example-bucket-positive-2"

						# Required: Virtual pool URL associated with the bucket (Get it using Replication Datasource)
						replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"

						# Required: Namespace for bucket isolation
						namespace = "ns1"

						# Optional: Size of each block in bytes
						block_size = 5096

						# Optional: Size threshold for notifications
						notification_size = 2024

						# Optional: Enable filesystem interface
						filesystem_enabled = true

						# Optional: Key-value tags for bucket
						tag = [
							{
							key   = "Env"
							value = "Production"
							},
							{
							key   = "Tag"
							value = "Production"
							}
						]

						# Optional: Enable server-side encryption
						# is_encryption_enabled = true

						# Optional: Default group permissions
						default_group_file_read_permission = "true"
						# default_group_file_write_permission   = "false"
						default_group_file_execute_permission = "true"
						# default_group_dir_read_permission     = "true"
						# default_group_dir_write_permission    = "false"
						# default_group_dir_execute_permission  = "true"

						# Optional: Default group name
						default_group = "default"

						# Optional: Auto-commit period in seconds
						auto_commit_period = 25

						# Optional: Retention period in days
						retention = 30

						# Optional: Allow stale reads
						is_stale_allowed = true

						# Optional: Allow object lock with ADO
						is_object_lock_with_ado_allowed = true

						default_retention = 30

						# Optional: Enable TSO read-only mode
						is_tso_read_only = true

						# Optional: Enable metadata search
						is_metadata_enabled = true

						# Optional: Metadata search configuration
						search_metadata = [
							{
							type     = "System"
							name     = "CreateTime"
							datatype = "datetime"
							},
							{
							type     = "User"
							name     = "x-amz-meta-abc"
							datatype = "string"
							}
						]

						

						min_max_governor = {
							enforce_retention          = true
							minimum_fixed_retention    = 10
							maximum_fixed_retention    = 365
							minimum_variable_retention = 5
							maximum_variable_retention = 180
						}
						
			    }
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "example-bucket-positive-2"),
					resource.TestCheckResourceAttr(resourceName, "owner", "root"),
				),
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
				ExpectError:       regexp.MustCompile(`Error importing Buckets`),
			},
			{
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: "example-bucket-positive-1:ns1",
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "example-bucket-positive-1"),
				),
			},
		},
	})
}
