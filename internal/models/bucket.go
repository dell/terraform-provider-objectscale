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

package models

import "github.com/hashicorp/terraform-plugin-framework/types"

// BucketDatasourceModel represents the Terraform data source model for listing buckets.
// It contains the namespace, optional prefix, and a list of BucketModel entries.
type BucketDatasourceModel struct {
	ID               types.String  `tfsdk:"id"`
	Namespace        types.String  `tfsdk:"namespace"`
	BucketNamePrefix types.String  `tfsdk:"bucket_name_prefix"`
	Buckets          []BucketModel `tfsdk:"buckets"`
}

// BucketModel maps to the BucketServiceGetBucketsResponseObjectBucketInner API struct.
// It represents a single bucket and all its properties as exposed by the API.
type BucketModel struct {
	Name                               types.String         `tfsdk:"name"`
	Id                                 types.String         `tfsdk:"id"`
	Created                            types.String         `tfsdk:"created"`
	Softquota                          types.String         `tfsdk:"softquota"`
	FsAccessEnabled                    types.Bool           `tfsdk:"fs_access_enabled"`
	Locked                             types.Bool           `tfsdk:"locked"`
	Vpool                              types.String         `tfsdk:"vpool"`
	Namespace                          types.String         `tfsdk:"namespace"`
	Owner                              types.String         `tfsdk:"owner"`
	IsStaleAllowed                     types.Bool           `tfsdk:"is_stale_allowed"`
	IsObjectLockWithAdoAllowed         types.Bool           `tfsdk:"is_object_lock_with_ado_allowed"`
	IsTsoReadOnly                      types.Bool           `tfsdk:"is_tso_read_only"`
	IsObjectLockEnabled                types.Bool           `tfsdk:"is_object_lock_enabled"`
	DefaultObjectLockRetentionMode     types.String         `tfsdk:"default_object_lock_retention_mode"`
	DefaultObjectLockRetentionYears    types.Int64          `tfsdk:"default_object_lock_retention_years"`
	DefaultObjectLockRetentionDays     types.Int64          `tfsdk:"default_object_lock_retention_days"`
	IsEncryptionEnabled                types.String         `tfsdk:"is_encryption_enabled"`
	DefaultRetention                   types.Int64          `tfsdk:"default_retention"`
	BlockSize                          types.Int64          `tfsdk:"block_size"`
	AutoCommitPeriod                   types.Int64          `tfsdk:"auto_commit_period"`
	NotificationSize                   types.Int64          `tfsdk:"notification_size"`
	ApiType                            types.String         `tfsdk:"api_type"`
	Tag                                []TagModel           `tfsdk:"tag"`
	Retention                          types.Int64          `tfsdk:"retention"`
	DefaultGroupFileReadPermission     types.Bool           `tfsdk:"default_group_file_read_permission"`
	DefaultGroupFileWritePermission    types.Bool           `tfsdk:"default_group_file_write_permission"`
	DefaultGroupFileExecutePermission  types.Bool           `tfsdk:"default_group_file_execute_permission"`
	DefaultGroupDirReadPermission      types.Bool           `tfsdk:"default_group_dir_read_permission"`
	DefaultGroupDirWritePermission     types.Bool           `tfsdk:"default_group_dir_write_permission"`
	DefaultGroupDirExecutePermission   types.Bool           `tfsdk:"default_group_dir_execute_permission"`
	DefaultGroup                       types.String         `tfsdk:"default_group"`
	SearchMetadata                     *SearchMetadataModel `tfsdk:"search_metadata"`
	MinMaxGovernor                     *MinMaxGovernorModel `tfsdk:"min_max_governor"`
	AuditDeleteExpiration              types.Int64          `tfsdk:"audit_delete_expiration"`
	IsEmptyBucketInProgress            types.Bool           `tfsdk:"is_empty_bucket_in_progress"`
	BlockSizeInCount                   types.Int64          `tfsdk:"block_size_in_count"`
	NotificationSizeInCount            types.Int64          `tfsdk:"notification_size_in_count"`
	EnableAdvancedMetadataSearch       types.Bool           `tfsdk:"enable_advanced_metadata_search"`
	AdvancedMetadataSearchTargetName   types.String         `tfsdk:"advanced_metadata_search_target_name"`
	AdvancedMetadataSearchTargetStream types.String         `tfsdk:"advanced_metadata_search_target_stream"`
	LocalObjectMetadataReads           types.Bool           `tfsdk:"local_object_metadata_reads"`
	VersioningStatus                   types.String         `tfsdk:"versioning_status"`
}

// TagModel represents a key-value tag for a bucket.
type TagModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

// SearchMetadataModel represents the search metadata block for a bucket, including metadata keys and flags.
type SearchMetadataModel struct {
	Metadata  []MetadataModel `tfsdk:"metadata"`
	IsEnabled types.Bool      `tfsdk:"is_enabled"`
	MdTokens  types.Bool      `tfsdk:"md_tokens"`
	MaxKeys   types.Int64     `tfsdk:"max_keys"`
}

// MetadataModel represents a single metadata key definition for a bucket.
type MetadataModel struct {
	Type     types.String `tfsdk:"type"`
	Name     types.String `tfsdk:"name"`
	Datatype types.String `tfsdk:"datatype"`
}

// MinMaxGovernorModel represents the min/max retention governor settings for a bucket.
type MinMaxGovernorModel struct {
	EnforceRetention         types.Bool  `tfsdk:"enforce_retention"`
	MinimumFixedRetention    types.Int64 `tfsdk:"minimum_fixed_retention"`
	MaximumFixedRetention    types.Int64 `tfsdk:"maximum_fixed_retention"`
	MinimumVariableRetention types.Int64 `tfsdk:"minimum_variable_retention"`
	MaximumVariableRetention types.Int64 `tfsdk:"maximum_variable_retention"`
}

// BucketResourceModel represents the Terraform resource model for a bucket.
// It includes all required and optional properties for bucket creation and management.
type BucketResourceModel struct {
	Id                                 types.String `tfsdk:"id"`
	Name                               types.String `tfsdk:"name"`
	Owner                              types.String `tfsdk:"owner"`
	Namespace                          types.String `tfsdk:"namespace"`
	ReplicationGroup                   types.String `tfsdk:"replication_group"`
	Created                            types.String `tfsdk:"created"`
	SoftQuota                          types.String `tfsdk:"soft_quota"`
	FsAccessEnabled                    types.Bool   `tfsdk:"filesystem_enabled"`
	Locked                             types.Bool   `tfsdk:"locked"`
	BlockSize                          types.Int64  `tfsdk:"block_size"`
	NotificationSize                   types.Int64  `tfsdk:"notification_size"`
	AutoCommitPeriod                   types.Int64  `tfsdk:"auto_commit_period"`
	ApiType                            types.String `tfsdk:"api_type"`
	Tag                                types.List   `tfsdk:"tag"`
	Retention                          types.Int64  `tfsdk:"retention"`
	DefaultGroupFileReadPermission     types.Bool   `tfsdk:"default_group_file_read_permission"`
	DefaultGroupFileWritePermission    types.Bool   `tfsdk:"default_group_file_write_permission"`
	DefaultGroupFileExecutePermission  types.Bool   `tfsdk:"default_group_file_execute_permission"`
	DefaultGroupDirReadPermission      types.Bool   `tfsdk:"default_group_dir_read_permission"`
	DefaultGroupDirWritePermission     types.Bool   `tfsdk:"default_group_dir_write_permission"`
	DefaultGroupDirExecutePermission   types.Bool   `tfsdk:"default_group_dir_execute_permission"`
	DefaultGroup                       types.String `tfsdk:"default_group"`
	SearchMetadata                     types.List   `tfsdk:"search_metadata"`
	IsEnabled                          types.Bool   `tfsdk:"is_metadata_enabled"`
	MdTokens                           types.Bool   `tfsdk:"md_tokens"`
	MaxKeys                            types.Int64  `tfsdk:"max_keys"`
	MinMaxGovernor                     types.Object `tfsdk:"min_max_governor"`
	AuditDeleteExpiration              types.Int64  `tfsdk:"audit_delete_expiration"`
	IsStaleAllowed                     types.Bool   `tfsdk:"is_stale_allowed"`
	IsObjectLockWithAdoAllowed         types.Bool   `tfsdk:"is_object_lock_with_ado_allowed"`
	IsTsoReadOnly                      types.Bool   `tfsdk:"is_tso_read_only"`
	IsObjectLockEnabled                types.Bool   `tfsdk:"is_object_lock_enabled"`
	DefaultObjectLockRetentionMode     types.String `tfsdk:"default_object_lock_retention_mode"`
	DefaultObjectLockRetentionYears    types.Int64  `tfsdk:"default_object_lock_retention_years"`
	DefaultObjectLockRetentionDays     types.Int64  `tfsdk:"default_object_lock_retention_days"`
	IsEncryptionEnabled                types.Bool   `tfsdk:"is_encryption_enabled"`
	DefaultRetention                   types.Int64  `tfsdk:"default_retention"`
	IsEmptyBucketInProgress            types.Bool   `tfsdk:"is_empty_bucket_in_progress"`
	BlockSizeInCount                   types.Int64  `tfsdk:"block_size_in_count"`
	NotificationSizeInCount            types.Int64  `tfsdk:"notification_size_in_count"`
	EnableAdvancedMetadataSearch       types.Bool   `tfsdk:"enable_advanced_metadata_search"`
	AdvancedMetadataSearchTargetName   types.String `tfsdk:"advanced_metadata_search_target_name"`
	AdvancedMetadataSearchTargetStream types.String `tfsdk:"advanced_metadata_search_target_stream"`
	LocalObjectMetadataReads           types.Bool   `tfsdk:"local_object_metadata_reads"`
	VersioningStatus                   types.String `tfsdk:"versioning_status"`

	//Policy related fields
	BucketPolicy types.String `tfsdk:"bucket_policy"`

	//ACL related fields
	UserAcl        types.List `tfsdk:"user_acl"`
	GroupAcl       types.List `tfsdk:"group_acl"`
	CustomGroupAcl types.List `tfsdk:"custom_group_acl"`
}

type AclModel struct {
	Name       types.String `tfsdk:"name"`
	Permission types.List   `tfsdk:"permission"`
}
