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
# After `terraform apply` of this example file it will create a bucket with the name set in `name` attribute on the ObjectScale

resource "objectscale_bucket" "example_bucket" {
  # Required: Owner of the bucket
  owner = "admin1"

  # Required: Name of the bucket
  name = "example-bucket-2"

  # Required: Virtual pool URL associated with the bucket (Get it using Replication Datasource)
  replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"

  # Required: Namespace for bucket isolation
  namespace = "ns1"

  # Optional: Size of each block in bytes
  block_size = 7000

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
  default_group_file_read_permission = "true"
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

  # Optional: Days after which audited delete expires
  audit_delete_expiration = 90

  # Optional: Enable object lock
  # is_object_lock_enabled = true

  # Optional: Enable advanced metadata search
  # enable_advanced_metadata_search = true

  # Optional: Advanced metadata search target name
  # advanced_metadata_search_target_name = "search-target"

  # Optional: Advanced metadata search target stream
  # advanced_metadata_search_target_stream = "stream-1"

  # Optional: Enable local metadata reads
  # local_object_metadata_reads = true

  # Optional: Versioning status (Enabled/Suspended)
  versioning_status = "Enabled"

  # bucket_policy = ""

  # Optional: Bucket policy in JSON format
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

  # Optional: User ACLs
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

  # Optional: User ACLs
  group_acl = [
    {
      name       = "all_users"
      permission = ["full_control"]
    }
  ]
}
# After the execution of above resource block, bucket would have been created on the ObjectScale array. For more information, Please check the terraform state file. 