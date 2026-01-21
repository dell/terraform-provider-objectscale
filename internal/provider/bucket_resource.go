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
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"terraform-provider-objectscale/internal/clientgen"
	"terraform-provider-objectscale/internal/helper"
	"terraform-provider-objectscale/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &BucketResource{}
var _ resource.ResourceWithImportState = &BucketResource{}

func NewBucketResource() resource.Resource {
	return &BucketResource{}
}

// BucketResource defines the resource implementation.
type BucketResource struct {
	resourceProviderConfig
}

// models.BucketResourceModel describes the resource data model.

func (r *BucketResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bucket"
}

func (r *BucketResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Unique identifier for the bucket resource.",
				MarkdownDescription: "Unique identifier for the bucket resource.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				Description:         "Name of the bucket.",
				MarkdownDescription: "Name of the bucket.",
				Required:            true,
			},
			"owner": schema.StringAttribute{
				Description:         "Owner of the bucket.",
				MarkdownDescription: "Owner of the bucket.",
				Required:            true,
			},
			"namespace": schema.StringAttribute{
				Description:         "Namespace for bucket isolation.",
				MarkdownDescription: "Namespace for bucket isolation.",
				Required:            true,
			},
			"replication_group": schema.StringAttribute{
				Description:         "Replication group associated with the bucket.",
				MarkdownDescription: "Replication group associated with the bucket.",
				Required:            true,
			},
			"created": schema.StringAttribute{
				Description:         "Creation date of the bucket resource.",
				MarkdownDescription: "Creation date of the bucket resource.",
				Computed:            true,
			},
			"soft_quota": schema.StringAttribute{
				Description:         "Soft quota for the bucket.",
				MarkdownDescription: "Soft quota for the bucket.",
				Computed:            true,
			},
			"filesystem_enabled": schema.BoolAttribute{
				Description:         "Enable filesystem access.",
				MarkdownDescription: "Enable filesystem access.",
				Optional:            true,
				Computed:            true,
			},
			"locked": schema.BoolAttribute{
				Description:         "Indicates if the bucket is locked.",
				MarkdownDescription: "Indicates if the bucket is locked.",
				Computed:            true,
			},
			"block_size": schema.Int64Attribute{
				Description:         "Size of each block in bytes.",
				MarkdownDescription: "Size of each block in bytes.",
				Optional:            true,
				Computed:            true,
			},
			"notification_size": schema.Int64Attribute{
				Description:         "Size threshold for notifications.",
				MarkdownDescription: "Size threshold for notifications.",
				Optional:            true,
				Computed:            true,
			},
			"auto_commit_period": schema.Int64Attribute{
				Description:         "Auto-commit period in seconds.",
				MarkdownDescription: "Auto-commit period in seconds.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 30),
				},
			},
			"api_type": schema.StringAttribute{
				Description:         "API type for the bucket.",
				MarkdownDescription: "API type for the bucket.",
				Computed:            true,
			},
			"tag": schema.SetNestedAttribute{
				Description:         "Key-value tags for the bucket.",
				MarkdownDescription: "Key-value tags for the bucket.",
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Description:         "Tag key.",
							MarkdownDescription: "Tag key.",
							Required:            true,
						},
						"value": schema.StringAttribute{
							Description:         "Tag value.",
							MarkdownDescription: "Tag value.",
							Required:            true,
						},
					},
				},
			},
			"retention": schema.Int64Attribute{
				Description:         "Retention period in days.",
				MarkdownDescription: "Retention period in days.",
				Optional:            true,
				Computed:            true,
			},
			"default_group_file_read_permission": schema.BoolAttribute{
				Description:         "Default group file read permission.",
				MarkdownDescription: "Default group file read permission.",
				Optional:            true,
				Computed:            true,
			},
			"default_group_file_write_permission": schema.BoolAttribute{
				Description:         "Default group file write permission.",
				MarkdownDescription: "Default group file write permission.",
				Optional:            true,
				Computed:            true,
			},
			"default_group_file_execute_permission": schema.BoolAttribute{
				Description:         "Default group file execute permission.",
				MarkdownDescription: "Default group file execute permission.",
				Optional:            true,
				Computed:            true,
			},
			"default_group_dir_read_permission": schema.BoolAttribute{
				Description:         "Default group directory read permission.",
				MarkdownDescription: "Default group directory read permission.",
				Optional:            true,
				Computed:            true,
			},
			"default_group_dir_write_permission": schema.BoolAttribute{
				Description:         "Default group directory write permission.",
				MarkdownDescription: "Default group directory write permission.",
				Optional:            true,
				Computed:            true,
			},
			"default_group_dir_execute_permission": schema.BoolAttribute{
				Description:         "Default group directory execute permission.",
				MarkdownDescription: "Default group directory execute permission.",
				Optional:            true,
				Computed:            true,
			},
			"default_group": schema.StringAttribute{
				Description:         "Default group name.",
				MarkdownDescription: "Default group name.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					// stringvalidator.AtLeastOneOf(
					// 	path.MatchRoot("default_group_dir_execute_permission"),
					// 	path.MatchRoot("default_group_dir_write_permission"),
					// 	path.MatchRoot("default_group_dir_read_permission"),
					// 	path.MatchRoot("default_group_file_execute_permission"),
					// 	path.MatchRoot("default_group_file_write_permission"),
					// 	path.MatchRoot("default_group_file_read_permission"),
					// ),
					// stringvalidator.AlsoRequires(
					// 	path.MatchRoot("filesystem_enabled"),
					// ),
					stringvalidator.ConflictsWith(
						path.MatchRoot("group_acl"),
						path.MatchRoot("user_acl"),
						path.MatchRoot("custom_group_acl"),
					),
				},
			},
			"is_metadata_enabled": schema.BoolAttribute{
				Description:         "Is search metadata enabled.",
				MarkdownDescription: "Is search metadata enabled.",
				Optional:            true,
				Computed:            true,
			},
			"md_tokens": schema.BoolAttribute{
				Description:         "Metadata tokens for advanced search.",
				MarkdownDescription: "Metadata tokens for advanced search.",
				Computed:            true,
			},
			"max_keys": schema.Int64Attribute{
				Description:         "Maximum number of keys for search.",
				MarkdownDescription: "Maximum number of keys for search.",
				// Optional:            true,
				Computed: true,
			},
			"search_metadata": schema.SetNestedAttribute{
				Description:         "List of metadata definitions.",
				MarkdownDescription: "List of metadata definitions.",
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Description:         "Metadata type.",
							MarkdownDescription: "Metadata type.",
							Optional:            true,
							Computed:            true,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"User",
									"System",
									"Head",
								),
							},
						},
						"name": schema.StringAttribute{
							Description:         "Metadata name.",
							MarkdownDescription: "Metadata name.",
							Optional:            true,
							Computed:            true,
						},
						"datatype": schema.StringAttribute{
							Description:         "Metadata datatype.",
							MarkdownDescription: "Metadata datatype.",
							Optional:            true,
							Computed:            true,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"datetime",
									"decimal",
									"integer",
									"string",
								),
							},
						},
					},
				},
			},
			"min_max_governor": schema.SingleNestedAttribute{
				Description:         "Retention governance settings.",
				MarkdownDescription: "Retention governance settings.",
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"enforce_retention": schema.BoolAttribute{
						Description:         "Enforce retention.",
						MarkdownDescription: "Enforce retention.",
						Optional:            true,
						Computed:            true,
					},
					"minimum_fixed_retention": schema.Int64Attribute{
						Description:         "Minimum fixed retention.",
						MarkdownDescription: "Minimum fixed retention.",
						Optional:            true,
						Computed:            true,
					},
					"maximum_fixed_retention": schema.Int64Attribute{
						Description:         "Maximum fixed retention.",
						MarkdownDescription: "Maximum fixed retention.",
						Optional:            true,
						Computed:            true,
					},
					"minimum_variable_retention": schema.Int64Attribute{
						Description:         "Minimum variable retention.",
						MarkdownDescription: "Minimum variable retention.",
						Optional:            true,
						Computed:            true,
					},
					"maximum_variable_retention": schema.Int64Attribute{
						Description:         "Maximum variable retention.",
						MarkdownDescription: "Maximum variable retention.",
						Optional:            true,
						Computed:            true,
					},
				},
			},
			"audit_delete_expiration": schema.Int64Attribute{
				Description: `Days after which audited delete expires.
							If not set, or set to -1 or -2, reflections are retained infinitely.
							If set to 0, reflections are deleted immediately and not retained.
							Any other positive value specifies the number of days to retain reflections before deletion.`,
				MarkdownDescription: `Days after which audited delete expires.
									- If not set, or set to -1 or -2, reflections are retained infinitely.
									- If set to 0, reflections are deleted immediately and not retained.
									- Any other positive value specifies the number of days to retain reflections before deletion.
				`,
				Optional: true,
				Computed: true,
			},
			"is_stale_allowed": schema.BoolAttribute{
				Description:         "Allow stale reads.",
				MarkdownDescription: "Allow stale reads.",
				Optional:            true,
				Computed:            true,
			},
			"is_object_lock_with_ado_allowed": schema.BoolAttribute{
				Description:         "Allow object lock with ADO.",
				MarkdownDescription: "Allow object lock with ADO.",
				Optional:            true,
				Computed:            true,
			},
			"is_tso_read_only": schema.BoolAttribute{
				Description:         "Enable TSO read-only mode.",
				MarkdownDescription: "Enable TSO read-only mode.",
				Optional:            true,
				Computed:            true,
			},
			"is_object_lock_enabled": schema.BoolAttribute{
				Description:         "Enable object lock.",
				MarkdownDescription: "Enable object lock.",
				Optional:            true,
				Computed:            true,
			},
			"default_object_lock_retention_mode": schema.StringAttribute{
				Description:         "Default object lock retention mode.",
				MarkdownDescription: "Default object lock retention mode.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Compliance",
						"Governance",
					),
				},
			},
			"default_object_lock_retention_years": schema.Int64Attribute{
				Description:         "Default object lock retention years.",
				MarkdownDescription: "Default object lock retention years.",
				Optional:            true,
				Computed:            true,
			},
			"default_object_lock_retention_days": schema.Int64Attribute{
				Description:         "Default object lock retention days.",
				MarkdownDescription: "Default object lock retention days.",
				Optional:            true,
				Computed:            true,
			},
			"is_encryption_enabled": schema.BoolAttribute{
				Description:         "Enable server-side encryption.",
				MarkdownDescription: "Enable server-side encryption.",
				Optional:            true,
				Computed:            true,
			},
			"default_retention": schema.Int64Attribute{
				Description:         "Default retention period in seconds.",
				MarkdownDescription: "Default retention period in seconds.",
				Optional:            true,
				Computed:            true,
			},
			"is_empty_bucket_in_progress": schema.BoolAttribute{
				Description:         "Indicates if empty bucket operation is in progress.",
				MarkdownDescription: "Indicates if empty bucket operation is in progress.",
				Computed:            true,
			},
			"block_size_in_count": schema.Int64Attribute{
				Description:         "Block size in count.",
				MarkdownDescription: "Block size in count.",
				Computed:            true,
			},
			"notification_size_in_count": schema.Int64Attribute{
				Description:         "Notification size in count.",
				MarkdownDescription: "Notification size in count.",
				Computed:            true,
			},
			"enable_advanced_metadata_search": schema.BoolAttribute{
				Description:         "Enable advanced metadata search.",
				MarkdownDescription: "Enable advanced metadata search.",
				Optional:            true,
				Computed:            true,
			},
			"advanced_metadata_search_target_name": schema.StringAttribute{
				Description:         "Advanced metadata search target name.",
				MarkdownDescription: "Advanced metadata search target name.",
				Optional:            true,
				Computed:            true,
			},
			"advanced_metadata_search_target_stream": schema.StringAttribute{
				Description:         "Advanced metadata search target stream.",
				MarkdownDescription: "Advanced metadata search target stream.",
				Optional:            true,
				Computed:            true,
			},
			"local_object_metadata_reads": schema.BoolAttribute{
				Description: `Enable or disable local object metadata reads for OBS CAS ADO RW buckets.
				When enabled, the bucket will attempt to read object metadata from locally replicated data, improving availability and reducing latency when the remote VDC is far away or unavailable. However, this may result in stale object metadata being returned if replication is incomplete, such as outdated deletion status, Litigation Hold, or Event Based Retention information. If the metadata is not available locally, it will be fetched from the remote VDC.`,
				MarkdownDescription: `Enable or disable local object metadata reads for OBS CAS ADO RW buckets.

				- When enabled, the bucket will attempt to read object metadata from locally replicated data, improving availability and reducing latency if the remote VDC is far away or unavailable.
				- This may result in stale object metadata being returned if the metadata is not fully replicated to the local VDC. For example, deletion status, Litigation Hold, or Event Based Retention information may be outdated.
				- If the object metadata is not available locally, it will be requested from the remote VDC.`,
				Optional: true,
				Computed: true,
			},
			"versioning_status": schema.StringAttribute{
				Description:         "Versioning status (Enabled/Suspended).",
				MarkdownDescription: "Versioning status (Enabled/Suspended).",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"ENABLED",
						"SUSPENDED",
					),
				},
			},
			"bucket_policy": schema.StringAttribute{
				Description:         "Bucket policy in JSON format.",
				MarkdownDescription: "Bucket policy in JSON format.",
				Optional:            true,
				Computed:            true,
			},
			"user_acl": schema.SetNestedAttribute{
				Description:         "List of user ACLs for the bucket.",
				MarkdownDescription: "List of user ACLs for the bucket.",
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description:         "User for the ACL entry.",
							MarkdownDescription: "User for the ACL entry.",
							Required:            true,
						},
						"permission": schema.SetAttribute{
							Description:         "List of permissions for the custom group. Valid values: full_control, read, delete, write, write_acl, read_acl, execute, privileged_write, none.",
							MarkdownDescription: "List of permissions for the custom group. Valid values: `full_control`, `read`, `delete`, `write`, `write_acl`, `read_acl`, `execute`, `privileged_write`, `none`.",
							ElementType:         types.StringType,
							Required:            true,
						},
					},
				},
			},
			"group_acl": schema.SetNestedAttribute{
				Description:         "List of group ACLs for the bucket.",
				MarkdownDescription: "List of group ACLs for the bucket.",
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description:         "Group for the ACL entry.",
							MarkdownDescription: "Group for the ACL entry.",
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"all_users",
									"log_delivery",
									"other",
									"public",
								),
							},
						},
						"permission": schema.SetAttribute{
							Description:         "List of permissions for the custom group. Valid values: full_control, read, delete, write, write_acl, read_acl, execute, privileged_write, none.",
							MarkdownDescription: "List of permissions for the custom group. Valid values: `full_control`, `read`, `delete`, `write`, `write_acl`, `read_acl`, `execute`, `privileged_write`, `none`.",
							ElementType:         types.StringType,
							Required:            true,
						},
					},
				},
			},
			"custom_group_acl": schema.SetNestedAttribute{
				Description:         "List of custom group ACLs for the bucket.",
				MarkdownDescription: "List of custom group ACLs for the bucket.",
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description:         "Custom group for the ACL entry.",
							MarkdownDescription: "Custom group for the ACL entry.",
							Required:            true,
						},
						"permission": schema.SetAttribute{
							Description:         "List of permissions for the custom group. Valid values: full_control, read, delete, write, write_acl, read_acl, execute, privileged_write, none.",
							MarkdownDescription: "List of permissions for the custom group. Valid values: `full_control`, `read`, `delete`, `write`, `write_acl`, `read_acl`, `execute`, `privileged_write`, `none`.",
							ElementType:         types.StringType,
							Required:            true,
						},
					},
				},
			},
		},
	}
}

func (r *BucketResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var config models.BucketResourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validation: if enforce_retention is true, then default_retention and retention must be the same (if both are set)
	if !config.MinMaxGovernor.IsNull() && !config.MinMaxGovernor.IsUnknown() {
		var minMax models.MinMaxGovernorModel
		diags := config.MinMaxGovernor.As(ctx, &minMax, basetypes.ObjectAsOptions{})
		resp.Diagnostics.Append(diags...)
		if !minMax.EnforceRetention.IsNull() && minMax.EnforceRetention.ValueBool() {
			if !config.DefaultRetention.IsNull() && !config.DefaultRetention.IsUnknown() &&
				!config.Retention.IsNull() && !config.Retention.IsUnknown() &&
				config.DefaultRetention.ValueInt64() != config.Retention.ValueInt64() {
				resp.Diagnostics.AddAttributeError(
					path.Root("default_retention"),
					"Default Retention and Retention Mismatch",
					"When 'enforce_retention' is true, 'default_retention' and 'retention' must be the same value.",
				)
			}
		}
	}

	if !config.DefaultGroupFileReadPermission.IsNull() && config.DefaultGroupFileReadPermission.ValueBool() ||
		!config.DefaultGroupFileWritePermission.IsNull() && config.DefaultGroupFileWritePermission.ValueBool() ||
		!config.DefaultGroupFileExecutePermission.IsNull() && config.DefaultGroupFileExecutePermission.ValueBool() ||
		!config.DefaultGroupDirReadPermission.IsNull() && config.DefaultGroupDirReadPermission.ValueBool() ||
		!config.DefaultGroupDirWritePermission.IsNull() && config.DefaultGroupDirWritePermission.ValueBool() ||
		!config.DefaultGroupDirExecutePermission.IsNull() && config.DefaultGroupDirExecutePermission.ValueBool() {
		if config.DefaultGroup.IsNull() || config.DefaultGroup.ValueString() == "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("default_group"),
				"Missing Default Group",
				"At least one default_group_*_permission is set to true, but default_group is not specified. Please set default_group.",
			)
		}
	}

	// Validate search_metadata.name for type "User"
	searchMetadataModels := []models.MetadataModel{}
	diags := config.SearchMetadata.ElementsAs(ctx, &searchMetadataModels, true)
	resp.Diagnostics.Append(diags...)

	for idx, md := range searchMetadataModels {
		if md.Type.ValueString() == "User" {
			name := md.Name.ValueString()
			if name == "" {
				resp.Diagnostics.AddAttributeError(
					path.Root("search_metadata").AtListIndex(idx).AtName("name"),
					"Empty Metadata Name",
					"Metadata name for type 'User' must not be empty.",
				)
				continue
			}
			if !strings.HasPrefix(name, "x-amz-meta-") {
				resp.Diagnostics.AddAttributeError(
					path.Root("search_metadata").AtListIndex(idx).AtName("name"),
					"Invalid Metadata Name Prefix",
					"Metadata name for type 'User' must start with 'x-amz-meta-'.",
				)
			} else {
				// Check the rest of the name after the prefix
				suffix := name[len("x-amz-meta-"):]
				if suffix == "" {
					resp.Diagnostics.AddAttributeError(
						path.Root("search_metadata").AtListIndex(idx).AtName("name"),
						"Empty Metadata Name Suffix",
						"Metadata name for type 'User' must have characters after 'x-amz-meta-'.",
					)
				} else if !regexp.MustCompile(`^[a-z0-9\-]+$`).MatchString(suffix) {
					resp.Diagnostics.AddAttributeError(
						path.Root("search_metadata").AtListIndex(idx).AtName("name"),
						"Invalid Metadata Name",
						"Metadata name for type 'User' must only contain lowercase letters, numbers, or hyphens after 'x-amz-meta-'.",
					)
				}
			}
		} else if md.Type.ValueString() == "System" {
			name := md.Name.ValueString()
			if name != "CreateTime" && name != "Owner" && name != "Size" && name != "LastModified" && name != "ObjectName" {
				resp.Diagnostics.AddAttributeError(
					path.Root("search_metadata").AtListIndex(idx).AtName("name"),
					"Invalid System Metadata Name",
					"Metadata name for type 'System' must be one of: CreateTime, Owner, Size, LastModified, ObjectName.",
				)
			}
		}
	}
}

func (r *BucketResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "creating bucket")
	var plan models.BucketResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Build the request from the plan
	reqBody := r.modelToJson(plan)

	_, _, err := r.client.GenClient.BucketApi.BucketServiceCreateBucket(ctx).BucketServiceCreateBucketRequest(reqBody).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error creating bucket", err.Error())
		return
	}

	// Handle user_acl, group_acl, and custom_group_acl
	if len(plan.UserAcl.Elements()) > 0 || len(plan.GroupAcl.Elements()) > 0 || len(plan.CustomGroupAcl.Elements()) > 0 {
		var userAclList []clientgen.BucketServiceSetBucketACLRequestAclUserAclInner
		var groupAclList []clientgen.BucketServiceSetBucketACLRequestAclGroupAclInner
		var customAclList []clientgen.BucketServiceSetBucketACLRequestAclCustomgroupAclInner
		for _, aclVal := range plan.UserAcl.Elements() {
			acl, ok := aclVal.(types.Object)
			if !ok {
				resp.Diagnostics.AddError("Type Assertion Error", "Failed to assert aclVal to types.Object for user_acl")
				return
			}
			userVal, ok := acl.Attributes()["name"].(types.String)
			if !ok {
				resp.Diagnostics.AddError("Type Assertion Error", "Failed to assert acl.Attributes()[\"name\"] to types.String for user_acl")
				return
			}
			user := userVal.ValueString()
			permAttr, ok := acl.Attributes()["permission"]
			if !ok {
				continue
			}
			permList, ok := permAttr.(types.Set)
			if !ok {
				continue
			}
			var permissions []string
			for _, p := range permList.Elements() {
				permStr, ok := p.(types.String)
				if !ok {
					continue
				}
				permissions = append(permissions, permStr.ValueString())
			}
			userAclList = append(userAclList, clientgen.BucketServiceSetBucketACLRequestAclUserAclInner{
				User:       &user,
				Permission: permissions,
			})
		}
		for _, aclVal := range plan.GroupAcl.Elements() {
			acl, ok := aclVal.(types.Object)
			if !ok {
				resp.Diagnostics.AddError("Type Assertion Error", "Failed to assert aclVal to types.Object for group_acl")
				return
			}
			groupVal, ok := acl.Attributes()["name"].(types.String)
			if !ok {
				resp.Diagnostics.AddError("Type Assertion Error", "Failed to assert acl.Attributes()[\"name\"] to types.String for group_acl")
				return
			}
			group := groupVal.ValueString()
			permAttr, ok := acl.Attributes()["permission"]
			if !ok {
				continue
			}
			permList, ok := permAttr.(types.Set)
			if !ok {
				continue
			}
			var permissions []string
			for _, p := range permList.Elements() {
				permStr, ok := p.(types.String)
				if !ok {
					continue
				}
				permissions = append(permissions, permStr.ValueString())
			}
			groupAclList = append(groupAclList, clientgen.BucketServiceSetBucketACLRequestAclGroupAclInner{
				Group:      &group,
				Permission: permissions,
			})
		}
		for _, aclVal := range plan.CustomGroupAcl.Elements() {
			acl, ok := aclVal.(types.Object)
			if !ok {
				resp.Diagnostics.AddError("Type Assertion Error", "Failed to assert aclVal to types.Object for custom_group_acl")
				return
			}
			nameAttr, ok := acl.Attributes()["name"].(types.String)
			if !ok {
				resp.Diagnostics.AddError("Type Assertion Error", "Failed to assert acl.Attributes()[\"name\"] to types.String for custom_group_acl")
				return
			}
			customGroup := nameAttr.ValueString()
			permAttr, ok := acl.Attributes()["permission"]
			if !ok {
				continue
			}
			permList, ok := permAttr.(types.Set)
			if !ok {
				continue
			}
			var permissions []string
			for _, p := range permList.Elements() {
				permStr, ok := p.(types.String)
				if !ok {
					continue
				}
				permissions = append(permissions, permStr.ValueString())
			}
			customAclList = append(customAclList, clientgen.BucketServiceSetBucketACLRequestAclCustomgroupAclInner{
				Customgroup: &customGroup,
				Permission:  permissions,
			})
		}
		namespace := plan.Namespace.ValueString()
		bucketName := plan.Name.ValueString()
		_, _, err := r.client.GenClient.BucketApi.
			BucketServiceSetBucketACL(ctx, plan.Name.ValueString()).
			BucketServiceSetBucketACLRequest(
				clientgen.BucketServiceSetBucketACLRequest{
					Bucket: &bucketName,
					Acl: &clientgen.BucketServiceSetBucketACLRequestAcl{
						UserAcl:        userAclList,
						GroupAcl:       groupAclList,
						CustomgroupAcl: customAclList,
					},
					Namespace: &namespace,
				},
			).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error setting bucket ACL", err.Error())
			_, _, err := r.client.GenClient.BucketApi.BucketServiceDeactivateBucket(ctx, plan.Name.ValueString()).Namespace(plan.Namespace.ValueString()).EmptyBucket("false").Execute()
			if err != nil {
				resp.Diagnostics.AddError(
					"Error deleting Bucket",
					err.Error(),
				)
			}
			return
		}
	}

	// Handle bucket policy
	if plan.BucketPolicy.ValueString() != "" {
		var policyMap map[string]interface{}
		err := json.Unmarshal([]byte(plan.BucketPolicy.ValueString()), &policyMap)
		if err != nil {
			resp.Diagnostics.AddError("Error parsing bucket policy JSON", err.Error())
			_, _, err := r.client.GenClient.BucketApi.BucketServiceDeactivateBucket(ctx, plan.Name.ValueString()).Namespace(plan.Namespace.ValueString()).EmptyBucket("false").Execute()
			if err != nil {
				resp.Diagnostics.AddError(
					"Error deleting Bucket",
					err.Error(),
				)
			}
			return
		}
		_, _, err = r.client.GenClient.BucketApi.
			BucketServiceSetBucketPolicy(ctx, plan.Name.ValueString()).
			Namespace(plan.Namespace.ValueString()).
			Body(policyMap).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error setting bucket policy", err.Error())
			_, _, err := r.client.GenClient.BucketApi.BucketServiceDeactivateBucket(ctx, plan.Name.ValueString()).Namespace(plan.Namespace.ValueString()).EmptyBucket("false").Execute()
			if err != nil {
				resp.Diagnostics.AddError(
					"Error deleting Bucket",
					err.Error(),
				)
			}
			return
		}
	}

	// Use setStateFromAPI to populate state from API after creation
	aclFromPlan := len(plan.UserAcl.Elements()) > 0 || len(plan.GroupAcl.Elements()) > 0 || len(plan.CustomGroupAcl.Elements()) > 0
	data, diags := r.setStateFromAPI(
		ctx,
		plan.Name.ValueString(),
		plan.Namespace.ValueString(),
		plan.BucketPolicy.ValueString(),
		aclFromPlan,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BucketResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.BucketResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Use setStateFromAPI to populate state from API after creation
	aclFromPlan := len(state.UserAcl.Elements()) > 0 || len(state.GroupAcl.Elements()) > 0 || len(state.CustomGroupAcl.Elements()) > 0
	data, diags := r.setStateFromAPI(
		ctx,
		state.Name.ValueString(),
		state.Namespace.ValueString(),
		state.BucketPolicy.ValueString(),
		aclFromPlan,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated plan into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BucketResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Update operation is not supported
	// resp.Diagnostics.AddError("Update Bucket operation is not supported.", "Update operation is not supported.")

	// Check if owner has changed
	var state models.BucketResourceModel
	var plan models.BucketResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	bucketName := state.Name.ValueString()
	namespace := state.Namespace.ValueString()

	// Only error if the plan values for default_group and fs_access_enabled are known (not unknown)
	if state.Name.ValueString() != plan.Name.ValueString() ||
		state.Namespace.ValueString() != plan.Namespace.ValueString() ||
		state.ReplicationGroup.ValueString() != plan.ReplicationGroup.ValueString() ||
		(!plan.FsAccessEnabled.IsUnknown() && state.FsAccessEnabled.ValueBool() != plan.FsAccessEnabled.ValueBool()) ||
		(!plan.DefaultGroup.IsUnknown() && state.DefaultGroup.ValueString() != plan.DefaultGroup.ValueString()) ||
		(!plan.SearchMetadata.IsUnknown() && !state.SearchMetadata.Equal(plan.SearchMetadata)) {
		resp.Diagnostics.AddError(
			"Immutable Field Change Detected",
			"Changing 'name', 'namespace', 'fs_access_enabled', 'replication_group', 'default_group', or 'search_metadata' is not supported. Please create a new resource instead.",
		)
		return
	}

	// Once IsObjectLockWithAdoAllowed is true, it cannot be changed back to false
	if state.IsObjectLockWithAdoAllowed.ValueBool() && !plan.IsObjectLockWithAdoAllowed.ValueBool() {
		resp.Diagnostics.AddError(
			"Immutable Field Change Detected",
			"Once 'is_object_lock_with_ado_allowed' is enabled (true), it cannot be changed back to false.",
		)
		return
	}

	// Handle LocalObjectMetadataReads update
	if state.LocalObjectMetadataReads.ValueBool() != plan.LocalObjectMetadataReads.ValueBool() {
		enabledStr := "Disable"
		if plan.LocalObjectMetadataReads.ValueBool() {
			enabledStr = "Enable"
		}
		_, _, err := r.client.GenClient.BucketApi.BucketServiceSetEventualReadsForBucket(ctx, bucketName).Namespace(namespace).Enabled(enabledStr).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error updating LocalObjectMetadataReads", err.Error())
			return
		}
	}

	//Handle Owner update
	if !plan.Owner.IsNull() && state.Owner.ValueString() != plan.Owner.ValueString() {
		resetPreviousOwners := true
		_, _, err := r.client.GenClient.BucketApi.
			BucketServiceUpdateBucketOwner(ctx, bucketName).
			BucketServiceUpdateBucketOwnerRequest(
				clientgen.BucketServiceUpdateBucketOwnerRequest{
					NewOwner:            plan.Owner.ValueString(),
					Namespace:           namespace,
					ResetPreviousOwners: &resetPreviousOwners,
				},
			).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error updating bucket owner", err.Error())
			return
		}
	}

	// Handle IsStaleAllowed update
	if (!plan.IsStaleAllowed.IsNull() && state.IsStaleAllowed.ValueBool() != plan.IsStaleAllowed.ValueBool()) ||
		(!plan.IsTsoReadOnly.IsNull() && state.IsTsoReadOnly.ValueBool() != plan.IsTsoReadOnly.ValueBool()) {
		isStaleAllowed := plan.IsStaleAllowed.ValueBool()
		isTsoReadonly := plan.IsTsoReadOnly.ValueBool()
		_, _, err := r.client.GenClient.BucketApi.
			BucketServiceUpdateBucketIsStaleAllowed(ctx, bucketName).
			BucketServiceUpdateBucketIsStaleAllowedRequest(
				clientgen.BucketServiceUpdateBucketIsStaleAllowedRequest{
					IsStaleAllowed: fmt.Sprintf("%v", isStaleAllowed),
					Namespace:      &namespace,
					IsTsoReadOnly:  &isTsoReadonly,
				},
			).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error updating IsStaleAllowed", err.Error())
			return
		}
	}

	// Handle AutoCommitPeriod update
	if !plan.AutoCommitPeriod.IsNull() && state.AutoCommitPeriod.ValueInt64() != plan.AutoCommitPeriod.ValueInt64() {
		autoCommitPeriod := plan.AutoCommitPeriod.ValueInt64()
		_, _, err := r.client.GenClient.BucketApi.
			BucketServiceSetBucketAutoCommitPeriod(ctx, bucketName).
			BucketServiceSetBucketAutoCommitPeriodRequest(
				clientgen.BucketServiceSetBucketAutoCommitPeriodRequest{
					Autocommit: fmt.Sprintf("%d", autoCommitPeriod),
					Namespace:  namespace,
				},
			).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error updating AutoCommitPeriod", err.Error())
			return
		}
	}

	// Handle Retention update
	if !plan.Retention.IsNull() && !plan.Retention.IsUnknown() && state.Retention.ValueInt64() != plan.Retention.ValueInt64() {
		retention := plan.Retention.ValueInt64()
		_, _, err := r.client.GenClient.BucketApi.
			BucketServiceSetBucketRetention(ctx, bucketName).
			BucketServiceSetBucketRetentionRequest(
				clientgen.BucketServiceSetBucketRetentionRequest{
					Period:    &retention,
					Namespace: &namespace,
				},
			).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error updating Retention", err.Error())
			return
		}
	}

	// Handle VersioningStatus update
	if !plan.VersioningStatus.IsNull() && !plan.VersioningStatus.IsUnknown() && state.VersioningStatus.ValueString() != plan.VersioningStatus.ValueString() {
		versioningStatus := plan.VersioningStatus.ValueString()
		_, _, err := r.client.GenClient.BucketApi.BucketServiceSetBucketVersioning(ctx, bucketName).
			BucketServiceSetBucketVersioningRequest(
				clientgen.BucketServiceSetBucketVersioningRequest{
					Status: &versioningStatus,
				},
			).Namespace(namespace).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error updating VersioningStatus", err.Error())
			return
		}
	}

	//Handle DefaultGroup and related permissions update
	// Only update DefaultGroup and related permissions if any of them are set in the plan (not null/unknown)
	// This avoids issues where computed values are not yet known during apply.
	if !plan.DefaultGroup.IsNull() && !plan.DefaultGroup.IsUnknown() &&
		(state.DefaultGroup.ValueString() != plan.DefaultGroup.ValueString() ||
			state.DefaultGroupFileReadPermission.ValueBool() != plan.DefaultGroupFileReadPermission.ValueBool() ||
			state.DefaultGroupFileWritePermission.ValueBool() != plan.DefaultGroupFileWritePermission.ValueBool() ||
			state.DefaultGroupFileExecutePermission.ValueBool() != plan.DefaultGroupFileExecutePermission.ValueBool() ||
			state.DefaultGroupDirReadPermission.ValueBool() != plan.DefaultGroupDirReadPermission.ValueBool() ||
			state.DefaultGroupDirWritePermission.ValueBool() != plan.DefaultGroupDirWritePermission.ValueBool() ||
			state.DefaultGroupDirExecutePermission.ValueBool() != plan.DefaultGroupDirExecutePermission.ValueBool() ||
			// Also trigger if any of the plan permissions are explicitly set (not null/unknown)
			(!plan.DefaultGroupFileReadPermission.IsNull() && !plan.DefaultGroupFileReadPermission.IsUnknown()) ||
			(!plan.DefaultGroupFileWritePermission.IsNull() && !plan.DefaultGroupFileWritePermission.IsUnknown()) ||
			(!plan.DefaultGroupFileExecutePermission.IsNull() && !plan.DefaultGroupFileExecutePermission.IsUnknown()) ||
			(!plan.DefaultGroupDirReadPermission.IsNull() && !plan.DefaultGroupDirReadPermission.IsUnknown()) ||
			(!plan.DefaultGroupDirWritePermission.IsNull() && !plan.DefaultGroupDirWritePermission.IsUnknown()) ||
			(!plan.DefaultGroupDirExecutePermission.IsNull() && !plan.DefaultGroupDirExecutePermission.IsUnknown())) {

		defaultGroupFileReadPermission := helper.ValueToPointer[bool](plan.DefaultGroupFileReadPermission)
		defaultGroupFileWritePermission := helper.ValueToPointer[bool](plan.DefaultGroupFileWritePermission)
		defaultGroupFileExecutePermission := helper.ValueToPointer[bool](plan.DefaultGroupFileExecutePermission)
		defaultGroupDirReadPermission := helper.ValueToPointer[bool](plan.DefaultGroupDirReadPermission)
		defaultGroupDirWritePermission := helper.ValueToPointer[bool](plan.DefaultGroupDirWritePermission)
		defaultGroupDirExecutePermission := helper.ValueToPointer[bool](plan.DefaultGroupDirExecutePermission)
		defaultGroup := plan.DefaultGroup.ValueString()

		_, _, err := r.client.GenClient.BucketApi.
			BucketServiceSetBucketDefaultGroup(ctx, bucketName).
			BucketServiceSetBucketDefaultGroupRequest(
				clientgen.BucketServiceSetBucketDefaultGroupRequest{
					DefaultGroup:                      &defaultGroup,
					DefaultGroupFileReadPermission:    defaultGroupFileReadPermission,
					DefaultGroupFileWritePermission:   defaultGroupFileWritePermission,
					DefaultGroupFileExecutePermission: defaultGroupFileExecutePermission,
					DefaultGroupDirReadPermission:     defaultGroupDirReadPermission,
					DefaultGroupDirWritePermission:    defaultGroupDirWritePermission,
					DefaultGroupDirExecutePermission:  defaultGroupDirExecutePermission,
					Namespace:                         &namespace,
				},
			).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error updating DefaultGroup or related permissions", err.Error())
			return
		}
	}

	// Handle tag updates
	// Convert state and plan tags to maps for easy comparison

	stateTags := make(map[string]string)
	for _, tagObj := range state.Tag.Elements() {
		tagObjTyped, ok := tagObj.(types.Object)
		if !ok {
			continue
		}
		tagMap := tagObjTyped.Attributes()
		keyVal, _ := tagMap["key"].(types.String)
		valueVal, _ := tagMap["value"].(types.String)
		stateTags[keyVal.ValueString()] = valueVal.ValueString()
	}

	planTags := make(map[string]string)
	for _, tagObj := range plan.Tag.Elements() {
		tagObjTyped, ok := tagObj.(types.Object)
		if !ok {
			continue
		}
		tagMap := tagObjTyped.Attributes()
		keyVal, _ := tagMap["key"].(types.String)
		valueVal, _ := tagMap["value"].(types.String)
		planTags[keyVal.ValueString()] = valueVal.ValueString()
	}

	// Delete tags that exist in state but not in plan
	for key := range stateTags {
		if _, exists := planTags[key]; !exists {
			val := stateTags[key]
			tagList := []clientgen.BucketServiceCreateBucketRequestTagSetInner{
				{
					Key:   &key,
					Value: &val,
				},
			}
			_, _, err := r.client.GenClient.BucketApi.
				BucketServiceDeleteBucketTags(ctx, bucketName).
				BucketServiceDeleteBucketTagsRequest(
					clientgen.BucketServiceDeleteBucketTagsRequest{
						TagSet:    tagList,
						Namespace: &namespace,
					},
				).
				Execute()
			if err != nil {
				resp.Diagnostics.AddError("Error deleting tag", fmt.Sprintf("Tag: %s, Error: %s", key, err.Error()))
				return
			}
		}
	}

	// Update existing tags if value changed, create new tags if not present in state
	for key, planVal := range planTags {
		stateVal, exists := stateTags[key]
		if exists {
			if stateVal != planVal {
				tagList := []clientgen.BucketServiceCreateBucketRequestTagSetInner{
					{
						Key:   &key,
						Value: &planVal,
					},
				}
				_, _, err := r.client.GenClient.BucketApi.
					BucketServiceUpdateBucketTags(ctx, bucketName).
					BucketServiceUpdateBucketTagsRequest(
						clientgen.BucketServiceUpdateBucketTagsRequest{
							TagSet:    tagList,
							Namespace: &namespace,
						},
					).
					Execute()
				if err != nil {
					resp.Diagnostics.AddError("Error updating tag", fmt.Sprintf("Tag: %s, Error: %s", key, err.Error()))
					return
				}
			}
		} else {
			// Create new tag
			tagList := []clientgen.BucketServiceCreateBucketRequestTagSetInner{
				{
					Key:   &key,
					Value: &planVal,
				},
			}
			_, _, err := r.client.GenClient.BucketApi.
				BucketServiceAddBucketTags(ctx, bucketName).
				BucketServiceAddBucketTagsRequest(
					clientgen.BucketServiceAddBucketTagsRequest{
						TagSet:    tagList,
						Namespace: &namespace,
					},
				).
				Execute()
			if err != nil {
				resp.Diagnostics.AddError("Error creating tag", fmt.Sprintf("Tag: %s, Error: %s", key, err.Error()))
				return
			}
		}
	}

	// Handle BlockSize and NotificationSize updates
	if (!plan.BlockSize.IsUnknown() && state.BlockSize.ValueInt64() != plan.BlockSize.ValueInt64()) ||
		(!plan.NotificationSize.IsUnknown() && state.NotificationSize.ValueInt64() != plan.NotificationSize.ValueInt64()) {
		blockSize := plan.BlockSize.ValueInt64()
		notificationSize := plan.NotificationSize.ValueInt64()
		_, _, err := r.client.GenClient.BucketApi.
			BucketServiceUpdateBucketQuota(ctx, bucketName).
			BucketServiceUpdateBucketQuotaRequest(
				clientgen.BucketServiceUpdateBucketQuotaRequest{
					BlockSize:        &blockSize,
					NotificationSize: &notificationSize,
					Namespace:        &namespace,
				},
			).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error updating BlockSize or NotificationSize", fmt.Sprintf("BlockSize: %d, NotificationSize: %d, Error: %s", blockSize, notificationSize, err.Error()))
			return
		}
	}

	// Handle IsObjectLockEnabled update
	if !plan.IsObjectLockEnabled.IsUnknown() && state.IsObjectLockEnabled.ValueBool() != plan.IsObjectLockEnabled.ValueBool() {
		isObjectLockEnabled := plan.IsObjectLockEnabled.ValueBool()
		// Prepare DefaultRetention Rule if any of the related fields are set
		var rule *clientgen.BucketServicePutBucketDefaultLockConfigurationRequestRule
		mode := plan.DefaultObjectLockRetentionMode.ValueString()
		years := plan.DefaultObjectLockRetentionYears.ValueInt64()
		days := plan.DefaultObjectLockRetentionDays.ValueInt64()

		// Only set Rule if at least one of the fields is set
		if mode != "" || years > 0 || days > 0 {
			var defaultRetention clientgen.BucketServicePutBucketDefaultLockConfigurationRequestRuleDefaultRetention

			if mode != "" {
				defaultRetention.Mode = &mode
			}
			// Only one of Years or Days should be set, prefer Years if both are set
			if years > 0 {
				y := int32(years)
				defaultRetention.Years = &y
			} else if days > 0 {
				d := int32(days)
				defaultRetention.Days = &d
			}

			rule = &clientgen.BucketServicePutBucketDefaultLockConfigurationRequestRule{
				DefaultRetention: &defaultRetention,
			}
		}

		var enabled string
		if isObjectLockEnabled {
			enabled = "Enabled"
		} else {
			enabled = "Disabled"
		}
		_, _, err := r.client.GenClient.BucketApi.
			BucketServicePutBucketDefaultLockConfiguration(ctx, bucketName).
			BucketServicePutBucketDefaultLockConfigurationRequest(
				clientgen.BucketServicePutBucketDefaultLockConfigurationRequest{
					ObjectLockEnabled: &enabled,
					Rule:              rule,
				},
			).
			Namespace(namespace).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error updating IsObjectLockEnabled", err.Error())
			return
		}
	}

	// Handle IsObjectLockWithAdoAllowed update
	if !plan.IsObjectLockWithAdoAllowed.IsUnknown() && state.IsObjectLockWithAdoAllowed.ValueBool() != plan.IsObjectLockWithAdoAllowed.ValueBool() {
		if plan.IsObjectLockWithAdoAllowed.ValueBool() {
			_, _, err := r.client.GenClient.BucketApi.
				BucketServiceEnableObjectLockWithAdoAllowedForExistingBucket(ctx, bucketName).
				Namespace(namespace).
				Execute()
			if err != nil {
				resp.Diagnostics.AddError("Error updating IsObjectLockWithAdoAllowed", err.Error())
				return
			}
		}
	}

	// Handle EnableAdvancedMetadataSearch and related fields update
	// Handle EnableAdvancedMetadataSearch toggle
	if !plan.EnableAdvancedMetadataSearch.IsUnknown() && state.EnableAdvancedMetadataSearch.ValueBool() != plan.EnableAdvancedMetadataSearch.ValueBool() {
		var err error
		if plan.EnableAdvancedMetadataSearch.ValueBool() {
			_, _, err = r.client.GenClient.BucketApi.
				BucketServiceActivateAdvancedMetadataSearch(ctx, bucketName).Namespace(namespace).
				Execute()
		} else {
			_, _, err = r.client.GenClient.BucketApi.
				BucketServiceDeactivateAdvancedMetadataSearch(ctx, bucketName).Namespace(namespace).
				Execute()
		}
		if err != nil {
			resp.Diagnostics.AddError("Error Updating AdvancedMetadataSearch Status", err.Error())
			return
		}
	}

	// Handle AdvancedMetadataSearchTargetName or AdvancedMetadataSearchTargetStream update
	if (!plan.AdvancedMetadataSearchTargetName.IsUnknown() && state.AdvancedMetadataSearchTargetName.ValueString() != plan.AdvancedMetadataSearchTargetName.ValueString()) ||
		(!plan.AdvancedMetadataSearchTargetStream.IsUnknown() && state.AdvancedMetadataSearchTargetStream.ValueString() != plan.AdvancedMetadataSearchTargetStream.ValueString()) {

		targetName := plan.AdvancedMetadataSearchTargetName.ValueString()
		targetStream := plan.AdvancedMetadataSearchTargetStream.ValueString()

		_, _, err := r.client.GenClient.BucketApi.
			BucketServiceSetAdvancedMetadataSearchTarget(ctx, bucketName).
			BucketServiceSetAdvancedMetadataSearchTargetRequest(
				clientgen.BucketServiceSetAdvancedMetadataSearchTargetRequest{
					TargetName: &targetName,
					StreamName: &targetStream,
				},
			).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error updating advanced metadata search target settings", err.Error())
			return
		}
	}

	// Handle AuditDeleteExpiration update
	if !plan.AuditDeleteExpiration.IsUnknown() && state.AuditDeleteExpiration.ValueInt64() != plan.AuditDeleteExpiration.ValueInt64() {
		auditDeleteExpiration := plan.AuditDeleteExpiration.ValueInt64()
		_, _, err := r.client.GenClient.BucketApi.
			BucketServiceSetBucketAuditDeleteExpiration(ctx, bucketName).Namespace(namespace).Expiration(fmt.Sprintf("%d", auditDeleteExpiration)).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error updating AuditDeleteExpiration", err.Error())
			return
		}
	}

	// Handle BucketPolicy update
	if plan.BucketPolicy.ValueString() != "" {
		var policyMap map[string]interface{}
		err := json.Unmarshal([]byte(plan.BucketPolicy.ValueString()), &policyMap)
		if err != nil {
			resp.Diagnostics.AddError("Error parsing bucket policy JSON", err.Error())
			return
		}
		_, _, err = r.client.GenClient.BucketApi.
			BucketServiceSetBucketPolicy(ctx, plan.Name.ValueString()).
			Namespace(plan.Namespace.ValueString()).
			Body(policyMap).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error setting bucket policy", err.Error())
			return
		}
	} else {
		_, _, err := r.client.GenClient.BucketApi.
			BucketServiceDeleteBucketPolicy(ctx, plan.Name.ValueString()).
			Namespace(plan.Namespace.ValueString()).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error deleting bucket policy", err.Error())
			return
		}
	}

	// Handle DefaultRetention update
	if (!plan.DefaultRetention.IsNull() && !plan.DefaultRetention.IsUnknown() && state.DefaultRetention.ValueInt64() != plan.DefaultRetention.ValueInt64()) ||
		(!plan.MinMaxGovernor.IsNull() && !plan.MinMaxGovernor.IsUnknown() && !state.MinMaxGovernor.Equal(plan.MinMaxGovernor)) {
		defaultRetention := plan.DefaultRetention.ValueInt64()
		minMaxGovernor := helper.ValueObjectTransform(plan.MinMaxGovernor, r.minMaxGovernorJson)
		_, _, err := r.client.GenClient.BucketApi.
			BucketServiceSetBucketRetention(ctx, bucketName).
			BucketServiceSetBucketRetentionRequest(
				clientgen.BucketServiceSetBucketRetentionRequest{
					Period:         &defaultRetention,
					Namespace:      &namespace,
					MinMaxGovernor: &minMaxGovernor,
				},
			).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error updating DefaultRetention", err.Error())
			return
		}
	}

	aclMap := func(acls types.Set) map[string][]string {
		result := make(map[string][]string)
		for _, aclVal := range acls.Elements() {
			acl, ok := aclVal.(types.Object)
			if !ok {
				continue
			}
			nameAttr, ok := acl.Attributes()["name"].(types.String)
			if !ok {
				continue
			}
			user := nameAttr.ValueString()
			permAttr, ok := acl.Attributes()["permission"]
			if !ok {
				continue
			}
			permList, ok := permAttr.(types.Set)
			if !ok {
				continue
			}
			var permissions []string
			for _, p := range permList.Elements() {
				permStr, ok := p.(types.String)
				if !ok {
					continue
				}
				permissions = append(permissions, permStr.ValueString())
			}
			result[user] = permissions
		}
		return result
	}

	stateUserAcl := aclMap(state.UserAcl)
	planUserAcl := aclMap(plan.UserAcl)
	stateGroupAcl := aclMap(state.GroupAcl)
	planGroupAcl := aclMap(plan.GroupAcl)
	stateCustomGroupAcl := aclMap(state.CustomGroupAcl)
	planCustomGroupAcl := aclMap(plan.CustomGroupAcl)

	aclChanged := false
	if len(stateUserAcl) != len(planUserAcl) || len(stateGroupAcl) != len(planGroupAcl) || len(stateCustomGroupAcl) != len(planCustomGroupAcl) {
		aclChanged = true
	} else {
		// Check for differences in users/groups and permissions
		for user, perms := range planUserAcl {
			if statePerms, ok := stateUserAcl[user]; !ok || len(perms) != len(statePerms) {
				aclChanged = true
				break
			} else {
				for i, p := range perms {
					if statePerms[i] != p {
						aclChanged = true
						break
					}
				}
			}
		}
		for group, perms := range planGroupAcl {
			if statePerms, ok := stateGroupAcl[group]; !ok || len(perms) != len(statePerms) {
				aclChanged = true
				break
			} else {
				for i, p := range perms {
					if statePerms[i] != p {
						aclChanged = true
						break
					}
				}
			}
		}
		for custom, perms := range planCustomGroupAcl {
			if statePerms, ok := stateCustomGroupAcl[custom]; !ok || len(perms) != len(statePerms) {
				aclChanged = true
				break
			} else {
				for i, p := range perms {
					if statePerms[i] != p {
						aclChanged = true
						break
					}
				}
			}
		}
	}

	if aclChanged {
		var userAclList []clientgen.BucketServiceSetBucketACLRequestAclUserAclInner
		var groupAclList []clientgen.BucketServiceSetBucketACLRequestAclGroupAclInner
		var customAclList []clientgen.BucketServiceSetBucketACLRequestAclCustomgroupAclInner

		for user, perms := range planUserAcl {
			userAclList = append(userAclList, clientgen.BucketServiceSetBucketACLRequestAclUserAclInner{
				User:       &user,
				Permission: perms,
			})
		}
		for group, perms := range planGroupAcl {
			groupAclList = append(groupAclList, clientgen.BucketServiceSetBucketACLRequestAclGroupAclInner{
				Group:      &group,
				Permission: perms,
			})
		}
		for custom, perms := range planCustomGroupAcl {
			customAclList = append(customAclList, clientgen.BucketServiceSetBucketACLRequestAclCustomgroupAclInner{
				Customgroup: &custom,
				Permission:  perms,
			})
		}
		namespace := plan.Namespace.ValueString()
		_, _, err := r.client.GenClient.BucketApi.
			BucketServiceSetBucketACL(ctx, plan.Name.ValueString()).
			BucketServiceSetBucketACLRequest(
				clientgen.BucketServiceSetBucketACLRequest{
					Bucket: &bucketName,
					Acl: &clientgen.BucketServiceSetBucketACLRequestAcl{
						UserAcl:        userAclList,
						GroupAcl:       groupAclList,
						CustomgroupAcl: customAclList,
					},
					Namespace: &namespace,
				},
			).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error updating bucket ACL", err.Error())
			return
		}
	}

	// Refresh state after owner update
	// Use setStateFromAPI to populate state from API after creation
	aclFromPlan := len(plan.UserAcl.Elements()) > 0 || len(plan.GroupAcl.Elements()) > 0 || len(plan.CustomGroupAcl.Elements()) > 0

	data, diags := r.setStateFromAPI(
		ctx,
		plan.Name.ValueString(),
		plan.Namespace.ValueString(),
		plan.BucketPolicy.ValueString(),
		aclFromPlan,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *BucketResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Info(ctx, "deleting Bucket")
	var state models.BucketResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, _, err := r.client.GenClient.BucketApi.BucketServiceDeactivateBucket(ctx, state.Name.ValueString()).Namespace(state.Namespace.ValueString()).EmptyBucket("true").Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting Bucket",
			err.Error(),
		)
	}
}

func (r *BucketResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Info(ctx, "importing Bucket")
	parts := strings.SplitN(req.ID, ":", 2)
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Error importing Bucket", "invalid format: expected 'bucket_name:namespace'")
		return
	}
	bucket_name := parts[0]
	namespace := parts[1]

	data, diags := r.setStateFromAPI(
		ctx,
		bucket_name,
		namespace,
		"", // No bucket policy from state
		true,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated plan into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BucketResource) tagListJson(in models.TagModel) clientgen.BucketServiceCreateBucketRequestTagSetInner {
	return clientgen.BucketServiceCreateBucketRequestTagSetInner{
		Key:   helper.ValueToPointer[string](in.Key),
		Value: helper.ValueToPointer[string](in.Value),
	}
}

func (r *BucketResource) metadataJson(in models.MetadataModel) clientgen.BucketServiceCreateBucketRequestSearchMetadataInner {
	return clientgen.BucketServiceCreateBucketRequestSearchMetadataInner{
		Type:     helper.ValueToPointer[string](in.Type),
		Name:     helper.ValueToPointer[string](in.Name),
		Datatype: helper.ValueToPointer[string](in.Datatype),
	}
}

func (r *BucketResource) minMaxGovernorJson(in models.MinMaxGovernorModel) clientgen.BucketServiceCreateBucketRequestMinMaxGovernor {
	return clientgen.BucketServiceCreateBucketRequestMinMaxGovernor{
		EnforceRetention:         helper.ValueToPointer[bool](in.EnforceRetention),
		MinimumFixedRetention:    helper.ValueToPointer[int64](in.MinimumFixedRetention),
		MaximumFixedRetention:    helper.ValueToPointer[int64](in.MaximumFixedRetention),
		MinimumVariableRetention: helper.ValueToPointer[int64](in.MinimumVariableRetention),
		MaximumVariableRetention: helper.ValueToPointer[int64](in.MaximumVariableRetention),
	}
}

func (r *BucketResource) modelToJson(plan models.BucketResourceModel) clientgen.BucketServiceCreateBucketRequest {

	minMaxGovernor := helper.ValueObjectTransform(plan.MinMaxGovernor, r.minMaxGovernorJson)

	// If VersioningStatus is set (not null/unknown and not empty) and FilesystemEnabled is not mentioned (null/unknown), set FilesystemEnabled to false
	filesystemEnabled := helper.ValueToPointer[bool](plan.FsAccessEnabled)
	if !plan.VersioningStatus.IsNull() && !plan.VersioningStatus.IsUnknown() && plan.VersioningStatus.ValueString() != "" &&
		(plan.FsAccessEnabled.IsNull() || plan.FsAccessEnabled.IsUnknown()) {
		val := false
		filesystemEnabled = &val
	}

	return clientgen.BucketServiceCreateBucketRequest{
		Name:                               plan.Name.ValueString(),
		Owner:                              helper.ValueToPointer[string](plan.Owner),
		Namespace:                          helper.ValueToPointer[string](plan.Namespace),
		Vpool:                              helper.ValueToPointer[string](plan.ReplicationGroup),
		FilesystemEnabled:                  filesystemEnabled,
		BlockSize:                          helper.ValueToPointer[int64](plan.BlockSize),
		NotificationSize:                   helper.ValueToPointer[int64](plan.NotificationSize),
		AutocommitPeriod:                   helper.ValueToPointer[int64](plan.AutoCommitPeriod),
		TagSet:                             helper.ValueListTransform(plan.Tag, r.tagListJson),
		Retention:                          helper.ValueToPointer[int64](plan.Retention),
		DefaultGroupFileReadPermission:     helper.ValueToPointer[bool](plan.DefaultGroupFileReadPermission),
		DefaultGroupFileWritePermission:    helper.ValueToPointer[bool](plan.DefaultGroupFileWritePermission),
		DefaultGroupFileExecutePermission:  helper.ValueToPointer[bool](plan.DefaultGroupFileExecutePermission),
		DefaultGroupDirReadPermission:      helper.ValueToPointer[bool](plan.DefaultGroupDirReadPermission),
		DefaultGroupDirWritePermission:     helper.ValueToPointer[bool](plan.DefaultGroupDirWritePermission),
		DefaultGroupDirExecutePermission:   helper.ValueToPointer[bool](plan.DefaultGroupDirExecutePermission),
		DefaultGroup:                       helper.ValueToPointer[string](plan.DefaultGroup),
		SearchMetadata:                     helper.ValueListTransform(plan.SearchMetadata, r.metadataJson),
		MetadataTokens:                     helper.ValueToPointer[bool](plan.MdTokens),
		MinMaxGovernor:                     &minMaxGovernor,
		AuditedDeleteExpiration:            helper.ValueToPointer[int64](plan.AuditDeleteExpiration),
		IsStaleAllowed:                     helper.ValueToPointer[bool](plan.IsStaleAllowed),
		IsObjectLockWithAdoAllowed:         helper.ValueToPointer[bool](plan.IsObjectLockWithAdoAllowed),
		IsTsoReadOnly:                      helper.ValueToPointer[bool](plan.IsTsoReadOnly),
		IsObjectLockEnabled:                helper.ValueToPointer[bool](plan.IsObjectLockEnabled),
		IsEncryptionEnabled:                helper.ValueToPointer[bool](plan.IsEncryptionEnabled),
		EnableAdvancedMetadataSearch:       helper.ValueToPointer[bool](plan.EnableAdvancedMetadataSearch),
		AdvancedMetadataSearchTargetName:   helper.ValueToPointer[string](plan.AdvancedMetadataSearchTargetName),
		AdvancedMetadataSearchTargetStream: helper.ValueToPointer[string](plan.AdvancedMetadataSearchTargetStream),
		LocalObjectMetadataReads:           helper.ValueToPointer[bool](plan.LocalObjectMetadataReads),
		VersioningStatus:                   helper.ValueToPointer[string](plan.VersioningStatus),
	}
}

// mapBucketToModel maps a BucketServiceGetBucketInfoResponse to models.BucketModel.
func getBucketToModel(b clientgen.BucketServiceGetBucketInfoResponse) models.BucketResourceModel {
	m := models.BucketResourceModel{
		Id:               helper.TfString(b.Id),
		Owner:            helper.TfString(b.Owner),
		Name:             helper.TfString(b.Name),
		ReplicationGroup: helper.TfString(b.Vpool),
		Namespace:        helper.TfString(b.Namespace),
		BlockSize:        helper.TfInt64(b.BlockSize),
		NotificationSize: helper.TfInt64(b.NotificationSize),
		FsAccessEnabled:  helper.TfBool(b.FsAccessEnabled),
		Tag: helper.SetNotNull(b.TagSet,
			func(v clientgen.BucketServiceCreateBucketRequestTagSetInner) types.Object {
				return helper.Object(models.Tags{
					Key:   helper.TfStringNN(v.Key),
					Value: helper.TfStringNN(v.Value),
				})
			}),
		IsEncryptionEnabled: func() types.Bool {
			if b.IsEncryptionEnabled != nil {
				// Accepts "true" or "false" as string
				return types.BoolValue(strings.ToLower(*b.IsEncryptionEnabled) == "true")
			}
			return types.BoolNull()
		}(),
		DefaultGroupFileReadPermission:    helper.TfBool(b.DefaultGroupFileReadPermission),
		DefaultGroupFileWritePermission:   helper.TfBool(b.DefaultGroupFileWritePermission),
		DefaultGroupFileExecutePermission: helper.TfBool(b.DefaultGroupFileExecutePermission),
		DefaultGroupDirReadPermission:     helper.TfBool(b.DefaultGroupDirReadPermission),
		DefaultGroupDirWritePermission:    helper.TfBool(b.DefaultGroupDirWritePermission),
		DefaultGroupDirExecutePermission:  helper.TfBool(b.DefaultGroupDirExecutePermission),
		DefaultGroup:                      helper.TfString(b.DefaultGroup),
		AutoCommitPeriod:                  helper.TfInt64(b.AutoCommitPeriod),
		Retention:                         helper.TfInt64(b.Retention),
		IsStaleAllowed:                    helper.TfBool(b.IsStaleAllowed),
		IsObjectLockWithAdoAllowed:        helper.TfBool(b.IsObjectLockWithAdoAllowed),
		IsTsoReadOnly:                     helper.TfBool(b.IsTsoReadOnly),
		DefaultRetention:                  helper.TfInt64(b.DefaultRetention),
		SearchMetadata: helper.SetNotNull(b.SearchMetadata.Metadata,
			func(v clientgen.BucketServiceCreateBucketRequestSearchMetadataInner) types.Object {
				return helper.Object(models.MetadataModel{
					Type:     helper.TfString(v.Type),
					Name:     helper.TfString(v.Name),
					Datatype: helper.TfString(v.Datatype),
				})
			}),

		IsEnabled: helper.TfBool(b.SearchMetadata.IsEnabled),
		MdTokens:  helper.TfBool(b.SearchMetadata.MdTokens),
		MaxKeys:   helper.TfInt64From32(b.SearchMetadata.MaxKeys),
		MinMaxGovernor: func() types.Object {
			return helper.Object(models.MinMaxGovernorModel{
				EnforceRetention:         helper.TfBool(b.MinMaxGovernor.EnforceRetention),
				MinimumFixedRetention:    helper.TfInt64(b.MinMaxGovernor.MinimumFixedRetention),
				MaximumFixedRetention:    helper.TfInt64(b.MinMaxGovernor.MaximumFixedRetention),
				MinimumVariableRetention: helper.TfInt64(b.MinMaxGovernor.MinimumVariableRetention),
				MaximumVariableRetention: helper.TfInt64(b.MinMaxGovernor.MaximumVariableRetention),
			})
		}(),
		AuditDeleteExpiration:              helper.TfInt64(b.AuditDeleteExpiration),
		IsObjectLockEnabled:                helper.TfBool(b.IsObjectLockEnabled),
		EnableAdvancedMetadataSearch:       helper.TfBool(b.EnableAdvancedMetadataSearch),
		AdvancedMetadataSearchTargetName:   helper.TfString(b.AdvancedMetadataSearchTargetName),
		AdvancedMetadataSearchTargetStream: helper.TfString(b.AdvancedMetadataSearchTargetStream),
		LocalObjectMetadataReads:           helper.TfBool(b.LocalObjectMetadataReads),
		VersioningStatus:                   helper.TfString(b.VersioningStatus),
		Created:                            helper.TfString(b.Created),
	}
	return m
}

func (r *BucketResource) setStateFromAPI(ctx context.Context, name, namespace, bucketPolicy string, aclFromPlan bool) (*models.BucketResourceModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	bucketData, _, err := r.client.GenClient.BucketApi.BucketServiceGetBucketInfo(ctx, name).Namespace(namespace).Execute()
	if err != nil {
		diags.AddError(
			"Error Reading Buckets",
			fmt.Sprintf("An error was encountered reading buckets from ObjectScale IAM: %s", err.Error()),
		)
		return nil, diags
	}

	// Fetch bucket policy if not provided
	if bucketPolicy == "" {
		_, _, err := r.client.GenClient.BucketApi.
			BucketServiceGetBucketPolicy(ctx, name).
			Namespace(namespace).
			Execute()
		if err != nil {
			diags.AddError(
				"Error Reading Bucket Policy",
				fmt.Sprintf("An error was encountered reading bucket policy from ObjectScale IAM: %s", err.Error()),
			)
			return nil, diags
		}
	}

	// Fetch ACLs only if any ACL is set in the plan/state
	userAclList := []models.AclModel{}
	groupAclList := []models.AclModel{}
	customGroupAclList := []models.AclModel{}
	if aclFromPlan {
		aclResp, _, err := r.client.GenClient.BucketApi.
			BucketServiceGetBucketACL(ctx, name).
			Namespace(namespace).
			Execute()
		if err != nil {
			diags.AddError(
				"Error Reading Bucket ACL",
				fmt.Sprintf("An error was encountered reading bucket ACL from ObjectScale IAM: %s", err.Error()),
			)
			return nil, diags
		}
		if aclResp != nil && aclResp.Acl != nil {
			for _, u := range aclResp.Acl.UserAcl {
				if u.User != nil {
					userAclList = append(userAclList, models.AclModel{
						Name:       types.StringValue(*u.User),
						Permission: helper.SetNotNull(u.Permission, types.StringValue),
					})
				}
			}
			for _, g := range aclResp.Acl.GroupAcl {
				if g.Group != nil {
					groupAclList = append(groupAclList, models.AclModel{
						Name:       types.StringValue(*g.Group),
						Permission: helper.SetNotNull(g.Permission, types.StringValue),
					})
				}
			}
			for _, c := range aclResp.Acl.CustomgroupAcl {
				if c.Customgroup != nil {
					customGroupAclList = append(customGroupAclList, models.AclModel{
						Name:       types.StringValue(*c.Customgroup),
						Permission: helper.SetNotNull(c.Permission, types.StringValue),
					})
				}
			}
		}
	}

	data := getBucketToModel(*bucketData)
	// Set BucketPolicy
	if bucketPolicy != "" {
		data.BucketPolicy = types.StringValue(bucketPolicy)
	} else {
		data.BucketPolicy = types.StringNull()
	}

	// Set ACLs: if not set in plan, keep as in plan (to avoid diff drift)
	// Always set ACLs to a known, non-null value (empty list if none from API)
	aclAttrTypes := map[string]attr.Type{
		"name":       types.StringType,
		"permission": types.SetType{ElemType: types.StringType},
	}
	aclObjectType := types.ObjectType{AttrTypes: aclAttrTypes}
	//aclListType := types.ListType{ElemType: aclObjectType}

	// Helper to convert []models.AclModel to []attr.Value
	aclModelsToList := func(acls []models.AclModel) []attr.Value {
		var objs []attr.Value
		for _, v := range acls {
			obj, _ := types.ObjectValue(aclAttrTypes, map[string]attr.Value{
				"name":       v.Name,
				"permission": v.Permission,
			})
			objs = append(objs, obj)
		}
		return objs
	}

	userAclListVal, _ := types.SetValue(aclObjectType, aclModelsToList(userAclList))
	groupAclListVal, _ := types.SetValue(aclObjectType, aclModelsToList(groupAclList))
	customGroupAclListVal, _ := types.SetValue(aclObjectType, aclModelsToList(customGroupAclList))

	data.UserAcl = userAclListVal
	data.GroupAcl = groupAclListVal
	data.CustomGroupAcl = customGroupAclListVal
	return &data, diags
}
