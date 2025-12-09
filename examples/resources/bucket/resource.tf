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
  name = "example-bucket"

  # Required: Virtual pool URL associated with the bucket (Get it using Replication Datasource)
  replication_group = "urn:storageos:ReplicationGroupInfo:1cb09936-67a2-4692-abd2-eb1277ef7364:global"

  # Required: Namespace for bucket isolation
  namespace = "ns1"

  # # Optional: Size of each block in bytes
  # block_size = 4096

  # # Optional: Size threshold for notifications
  # notification_size = 1024

  # # Optional: Enable filesystem interface
  # filesystem_enabled = true

  # # Optional: Type of bucket head (e.g., object)
  # head_type = "object"

  # # Optional: Key-value tags for bucket
  tag = [
    {
      key   = "Environment"
      value = "Production"
    }
  ]

  # # Optional: Enable server-side encryption
  # is_encryption_enabled = true

  # # Optional: Default group permissions
  # default_group_file_read_permission    = "true"
  # default_group_file_write_permission   = "false"
  # default_group_file_execute_permission = "false"
  # default_group_dir_read_permission     = "true"
  # default_group_dir_write_permission    = "false"
  # default_group_dir_execute_permission  = "true"

  # # Optional: Default group name
  # default_group = "default-group"

  # # Optional: Auto-commit period in seconds
  # autocommit_period = 60

  # # Optional: Retention period in days
  # retention = 30

  # # Optional: Allow stale reads
  # is_stale_allowed = false

  # # Optional: Allow object lock with ADO
  # is_object_lock_with_ado_allowed = false

  # # Optional: Enable TSO read-only mode
  # is_tso_read_only = false

  # # Optional: Metadata search configuration
  # search_metadata = [
  #   {
  #     type     = "custom"
  #     name     = "project"
  #     datatype = "string"
  #   }
  # ]

  # # Optional: Metadata tokens for advanced search
  # metadata_tokens = false

  # # Optional: Retention governance settings
  # min_max_governor {
  #   enforce_retention           = true
  #   minimum_fixed_retention     = 10
  #   maximum_fixed_retention     = 365
  #   minimum_variable_retention  = 5
  #   maximum_variable_retention  = 180
  # }

  # # Optional: Days after which audited delete expires
  # audited_delete_expiration = 90

  # # Optional: Enable object lock
  # is_object_lock_enabled = true

  # # Optional: Storage policy type
  # storage_policy = "standard"

  # # Optional: Enable advanced metadata search
  # enable_advanced_metadata_search = true

  # # Optional: Advanced metadata search target name
  # advanced_metadata_search_target_name = "search-target"

  # # Optional: Advanced metadata search target stream
  # advanced_metadata_search_target_stream = "stream-1"

  # # Optional: Enable local metadata reads
  # local_object_metadata_reads = true

  # # Optional: Versioning status (Enabled/Suspended)
  # versioning_status = "Enabled"
}

# After the execution of above resource block, namespace would have been created on the ObjectScale array. For more information, Please check the terraform state file. 
