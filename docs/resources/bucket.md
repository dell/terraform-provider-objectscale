---
# Copyright (c) 2025-2026 Dell Inc., or its subsidiaries. All Rights Reserved.
#
# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://mozilla.org/MPL/2.0/
#
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

title: "objectscale_bucket resource"
linkTitle: "objectscale_bucket"
page_title: "objectscale_bucket Resource - terraform-provider-objectscale"
subcategory: "Object Storage Containers"
description: |-
  This resource provisions and manages S3 buckets on Dell ObjectScale.
---

# objectscale_bucket (Resource)

This resource provisions and manages S3 buckets on Dell ObjectScale.


## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the bucket.
- `namespace` (String) Namespace for bucket isolation.
- `owner` (String) Owner of the bucket.
- `replication_group` (String) Replication group associated with the bucket.

### Optional

- `advanced_metadata_search_target_name` (String) Advanced metadata search target name.
- `advanced_metadata_search_target_stream` (String) Advanced metadata search target stream.
- `audit_delete_expiration` (Number) Days after which audited delete expires.
									- If not set, or set to -1 or -2, reflections are retained infinitely.
									- If set to 0, reflections are deleted immediately and not retained.
									- Any other positive value specifies the number of days to retain reflections before deletion.
- `auto_commit_period` (Number) Auto-commit period in seconds.
- `block_size` (Number) Size of each block in bytes.
- `bucket_policy` (String) Bucket policy in JSON format.
- `custom_group_acl` (Attributes Set) List of custom group ACLs for the bucket. (see [below for nested schema](#nestedatt--custom_group_acl))
- `default_group` (String) Default group name.
- `default_group_dir_execute_permission` (Boolean) Default group directory execute permission.
- `default_group_dir_read_permission` (Boolean) Default group directory read permission.
- `default_group_dir_write_permission` (Boolean) Default group directory write permission.
- `default_group_file_execute_permission` (Boolean) Default group file execute permission.
- `default_group_file_read_permission` (Boolean) Default group file read permission.
- `default_group_file_write_permission` (Boolean) Default group file write permission.
- `default_object_lock_retention_days` (Number) Default object lock retention days.
- `default_object_lock_retention_mode` (String) Default object lock retention mode.
- `default_object_lock_retention_years` (Number) Default object lock retention years.
- `default_retention` (Number) Default retention period in seconds.
- `enable_advanced_metadata_search` (Boolean) Enable advanced metadata search.
- `filesystem_enabled` (Boolean) Enable filesystem access.
- `group_acl` (Attributes Set) List of group ACLs for the bucket. (see [below for nested schema](#nestedatt--group_acl))
- `is_encryption_enabled` (Boolean) Enable server-side encryption.
- `is_metadata_enabled` (Boolean) Is search metadata enabled.
- `is_object_lock_enabled` (Boolean) Enable object lock.
- `is_object_lock_with_ado_allowed` (Boolean) Allow object lock with ADO.
- `is_stale_allowed` (Boolean) Allow stale reads.
- `is_tso_read_only` (Boolean) Enable TSO read-only mode.
- `local_object_metadata_reads` (Boolean) Enable or disable local object metadata reads for OBS CAS ADO RW buckets.

				- When enabled, the bucket will attempt to read object metadata from locally replicated data, improving availability and reducing latency if the remote VDC is far away or unavailable.
				- This may result in stale object metadata being returned if the metadata is not fully replicated to the local VDC. For example, deletion status, Litigation Hold, or Event Based Retention information may be outdated.
				- If the object metadata is not available locally, it will be requested from the remote VDC.
- `min_max_governor` (Attributes) Retention governance settings. (see [below for nested schema](#nestedatt--min_max_governor))
- `notification_size` (Number) Size threshold for notifications.
- `retention` (Number) Retention period in days.
- `search_metadata` (Attributes Set) List of metadata definitions. (see [below for nested schema](#nestedatt--search_metadata))
- `tag` (Attributes Set) Key-value tags for the bucket. (see [below for nested schema](#nestedatt--tag))
- `user_acl` (Attributes Set) List of user ACLs for the bucket. (see [below for nested schema](#nestedatt--user_acl))
- `versioning_status` (String) Versioning status (Enabled/Suspended).

### Read-Only

- `api_type` (String) API type for the bucket.
- `block_size_in_count` (Number) Block size in count.
- `created` (String) Creation date of the bucket resource.
- `id` (String) Unique identifier for the bucket resource.
- `is_empty_bucket_in_progress` (Boolean) Indicates if empty bucket operation is in progress.
- `locked` (Boolean) Indicates if the bucket is locked.
- `max_keys` (Number) Maximum number of keys for search.
- `md_tokens` (Boolean) Metadata tokens for advanced search.
- `notification_size_in_count` (Number) Notification size in count.
- `soft_quota` (String) Soft quota for the bucket.

<a id="nestedatt--custom_group_acl"></a>
### Nested Schema for `custom_group_acl`

Required:

- `name` (String) Custom group for the ACL entry.
- `permission` (Set of String) List of permissions for the custom group. Valid values: `full_control`, `read`, `delete`, `write`, `write_acl`, `read_acl`, `execute`, `privileged_write`, `none`.


<a id="nestedatt--group_acl"></a>
### Nested Schema for `group_acl`

Required:

- `name` (String) Group for the ACL entry.
- `permission` (Set of String) List of permissions for the custom group. Valid values: `full_control`, `read`, `delete`, `write`, `write_acl`, `read_acl`, `execute`, `privileged_write`, `none`.


<a id="nestedatt--min_max_governor"></a>
### Nested Schema for `min_max_governor`

Optional:

- `enforce_retention` (Boolean) Enforce retention.
- `maximum_fixed_retention` (Number) Maximum fixed retention.
- `maximum_variable_retention` (Number) Maximum variable retention.
- `minimum_fixed_retention` (Number) Minimum fixed retention.
- `minimum_variable_retention` (Number) Minimum variable retention.


<a id="nestedatt--search_metadata"></a>
### Nested Schema for `search_metadata`

Optional:

- `datatype` (String) Metadata datatype.
- `name` (String) Metadata name.
- `type` (String) Metadata type.


<a id="nestedatt--tag"></a>
### Nested Schema for `tag`

Required:

- `key` (String) Tag key.
- `value` (String) Tag value.


<a id="nestedatt--user_acl"></a>
### Nested Schema for `user_acl`

Required:

- `name` (String) User for the ACL entry.
- `permission` (Set of String) List of permissions for the custom group. Valid values: `full_control`, `read`, `delete`, `write`, `write_acl`, `read_acl`, `execute`, `privileged_write`, `none`.

Unless specified otherwise, all fields of this resource can be updated.

## Import

Import is supported using the following syntax:

```shell
# Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://mozilla.org/MPL/2.0/


# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# The command is
# terraform import objectscale_bucket.bucket <bucket_name:namespace>
# Example:
terraform import objectscale_bucket.bucket example_bucket:name_space_1
# after running this command, populate the name field and other required parameters in the config file to start managing this resource.
# Note: running "terraform show" after importing shows the current config/state of the resource. You can copy/paste that config to make it easier to manage the resource.
```
