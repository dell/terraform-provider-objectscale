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
	"fmt"
	"strings"

	"terraform-provider-objectscale/internal/client"
	"terraform-provider-objectscale/internal/clientgen"
	"terraform-provider-objectscale/internal/helper"
	"terraform-provider-objectscale/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	client *client.Client
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
			"owner": schema.StringAttribute{
				Description:         "Owner of the bucket.",
				MarkdownDescription: "Owner of the bucket.",
				Required:            true,
			},
			"name": schema.StringAttribute{
				Description:         "Name of the bucket.",
				MarkdownDescription: "Name of the bucket.",
				Required:            true,
			},
			"vpool": schema.StringAttribute{
				Description:         "Virtual pool URL associated with the bucket.",
				MarkdownDescription: "Virtual pool URL associated with the bucket.",
				Required:            true,
			},
			"namespace": schema.StringAttribute{
				Description:         "Namespace for bucket isolation.",
				MarkdownDescription: "Namespace for bucket isolation.",
				Required:            true,
			},
			"block_size": schema.Int64Attribute{
				Description:         "Size of each block in bytes.",
				MarkdownDescription: "Size of each block in bytes.",
				Optional:            true,
			},
			"notification_size": schema.Int64Attribute{
				Description:         "Size threshold for notifications.",
				MarkdownDescription: "Size threshold for notifications.",
				Optional:            true,
			},
			"filesystem_enabled": schema.BoolAttribute{
				Description:         "Enable filesystem interface.",
				MarkdownDescription: "Enable filesystem interface.",
				Optional:            true,
			},
			"head_type": schema.StringAttribute{
				Description:         "Type of bucket head (e.g., object).",
				MarkdownDescription: "Type of bucket head (e.g., object).",
				Optional:            true,
			},
			"tag_set": schema.ListNestedAttribute{
				Description:         "Key-value tags for bucket.",
				MarkdownDescription: "Key-value tags for bucket.",
				Optional:            true,
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
			"is_encryption_enabled": schema.BoolAttribute{
				Description:         "Enable server-side encryption.",
				MarkdownDescription: "Enable server-side encryption.",
				Optional:            true,
			},
			"default_group_file_read_permission": schema.StringAttribute{
				Description:         "Default group file read permission.",
				MarkdownDescription: "Default group file read permission.",
				Optional:            true,
			},
			"default_group_file_write_permission": schema.StringAttribute{
				Description:         "Default group file write permission.",
				MarkdownDescription: "Default group file write permission.",
				Optional:            true,
			},
			"default_group_file_execute_permission": schema.StringAttribute{
				Description:         "Default group file execute permission.",
				MarkdownDescription: "Default group file execute permission.",
				Optional:            true,
			},
			"default_group_dir_read_permission": schema.StringAttribute{
				Description:         "Default group directory read permission.",
				MarkdownDescription: "Default group directory read permission.",
				Optional:            true,
			},
			"default_group_dir_write_permission": schema.StringAttribute{
				Description:         "Default group directory write permission.",
				MarkdownDescription: "Default group directory write permission.",
				Optional:            true,
			},
			"default_group_dir_execute_permission": schema.StringAttribute{
				Description:         "Default group directory execute permission.",
				MarkdownDescription: "Default group directory execute permission.",
				Optional:            true,
			},
			"default_group": schema.StringAttribute{
				Description:         "Default group name.",
				MarkdownDescription: "Default group name.",
				Optional:            true,
			},
			"autocommit_period": schema.Int64Attribute{
				Description:         "Auto-commit period in seconds.",
				MarkdownDescription: "Auto-commit period in seconds.",
				Optional:            true,
			},
			"retention": schema.Int64Attribute{
				Description:         "Retention period in days.",
				MarkdownDescription: "Retention period in days.",
				Optional:            true,
			},
			"is_stale_allowed": schema.BoolAttribute{
				Description:         "Allow stale reads.",
				MarkdownDescription: "Allow stale reads.",
				Optional:            true,
			},
			"is_object_lock_with_ado_allowed": schema.BoolAttribute{
				Description:         "Allow object lock with ADO.",
				MarkdownDescription: "Allow object lock with ADO.",
				Optional:            true,
			},
			"is_tso_read_only": schema.BoolAttribute{
				Description:         "Enable TSO read-only mode.",
				MarkdownDescription: "Enable TSO read-only mode.",
				Optional:            true,
			},
			"search_metadata": schema.ListNestedAttribute{
				Description:         "Metadata search configuration.",
				MarkdownDescription: "Metadata search configuration.",
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Description:         "Metadata type.",
							MarkdownDescription: "Metadata type.",
							Required:            true,
						},
						"name": schema.StringAttribute{
							Description:         "Metadata name.",
							MarkdownDescription: "Metadata name.",
							Required:            true,
						},
						"datatype": schema.StringAttribute{
							Description:         "Metadata datatype.",
							MarkdownDescription: "Metadata datatype.",
							Required:            true,
						},
					},
				},
			},
			"metadata_tokens": schema.BoolAttribute{
				Description:         "Metadata tokens for advanced search.",
				MarkdownDescription: "Metadata tokens for advanced search.",
				Optional:            true,
			},
			"min_max_governor": schema.SingleNestedAttribute{
				Description:         "Retention governance settings.",
				MarkdownDescription: "Retention governance settings.",
				Optional:            true,
				Attributes: map[string]schema.Attribute{
					"enforce_retention": schema.BoolAttribute{
						Description:         "Enforce retention.",
						MarkdownDescription: "Enforce retention.",
						Optional:            true,
					},
					"minimum_fixed_retention": schema.Int64Attribute{
						Description:         "Minimum fixed retention.",
						MarkdownDescription: "Minimum fixed retention.",
						Optional:            true,
					},
					"maximum_fixed_retention": schema.Int64Attribute{
						Description:         "Maximum fixed retention.",
						MarkdownDescription: "Maximum fixed retention.",
						Optional:            true,
					},
					"minimum_variable_retention": schema.Int64Attribute{
						Description:         "Minimum variable retention.",
						MarkdownDescription: "Minimum variable retention.",
						Optional:            true,
					},
					"maximum_variable_retention": schema.Int64Attribute{
						Description:         "Maximum variable retention.",
						MarkdownDescription: "Maximum variable retention.",
						Optional:            true,
					},
				},
			},
			"audited_delete_expiration": schema.Int64Attribute{
				Description:         "Days after which audited delete expires.",
				MarkdownDescription: "Days after which audited delete expires.",
				Optional:            true,
			},
			"is_object_lock_enabled": schema.BoolAttribute{
				Description:         "Enable object lock.",
				MarkdownDescription: "Enable object lock.",
				Optional:            true,
			},
			"storage_policy": schema.StringAttribute{
				Description:         "Storage policy type.",
				MarkdownDescription: "Storage policy type.",
				Optional:            true,
			},
			"enable_advanced_metadata_search": schema.BoolAttribute{
				Description:         "Enable advanced metadata search.",
				MarkdownDescription: "Enable advanced metadata search.",
				Optional:            true,
			},
			"advanced_metadata_search_target_name": schema.StringAttribute{
				Description:         "Advanced metadata search target name.",
				MarkdownDescription: "Advanced metadata search target name.",
				Optional:            true,
			},
			"advanced_metadata_search_target_stream": schema.StringAttribute{
				Description:         "Advanced metadata search target stream.",
				MarkdownDescription: "Advanced metadata search target stream.",
				Optional:            true,
			},
			"local_object_metadata_reads": schema.BoolAttribute{
				Description:         "Enable local metadata reads.",
				MarkdownDescription: "Enable local metadata reads.",
				Optional:            true,
			},
			"versioning_status": schema.StringAttribute{
				Description:         "Versioning status (Enabled/Suspended).",
				MarkdownDescription: "Versioning status (Enabled/Suspended).",
				Optional:            true,
			},
		},
	}
}

func (r *BucketResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *BucketResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "creating group")
	var plan models.BucketResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	bucket, _, err := r.client.GenClient.BucketApi.BucketServiceCreateBucket(ctx).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error creating bucket", err.Error())
		return
	}

	// apiReq := r.client.GenClient.BucketApi.BucketServiceGetBuckets(ctx).Namespace(ns)
	// if *bucket.Name != "" {
	// 	apiReq = apiReq.Name(*bucket.Name + "*")
	// }

	// apiResp, httpResp, err := apiReq.Execute()
	// if err != nil {
	// 	resp.Diagnostics.AddError(
	// 		"Unable to read buckets",
	// 		fmt.Sprintf("Error: %s", err),
	// 	)
	// 	return
	// }

	data := models.BucketResourceModel{
		Id:                                types.StringValue(*bucket.Id),
		Owner:                             plan.Owner,
		Name:                              plan.Name,
		ReplicationGroup:                  plan.ReplicationGroup,
		Namespace:                         plan.Namespace,
		BlockSize:                         plan.BlockSize,
		NotificationSize:                  plan.NotificationSize,
		FsAccessEnabled:                   plan.FsAccessEnabled,
		Tag:                               plan.Tag,
		IsEncryptionEnabled:               plan.IsEncryptionEnabled,
		DefaultGroupFileReadPermission:    plan.DefaultGroupFileReadPermission,
		DefaultGroupFileWritePermission:   plan.DefaultGroupFileWritePermission,
		DefaultGroupFileExecutePermission: plan.DefaultGroupFileExecutePermission,
		DefaultGroupDirReadPermission:     plan.DefaultGroupDirReadPermission,
		DefaultGroupDirWritePermission:    plan.DefaultGroupDirWritePermission,
		DefaultGroupDirExecutePermission:  plan.DefaultGroupDirExecutePermission,
		DefaultGroup:                      plan.DefaultGroup,
		AutoCommitPeriod:                  plan.AutoCommitPeriod,
		Retention:                         plan.Retention,
		IsStaleAllowed:                    plan.IsStaleAllowed,
		IsObjectLockWithAdoAllowed:        plan.IsObjectLockWithAdoAllowed,
		IsTsoReadOnly:                     plan.IsTsoReadOnly,
		SearchMetadata:                    plan.SearchMetadata,
		// MetadataTokens:                     plan.MetadataTokens,
		MinMaxGovernor: plan.MinMaxGovernor,
		// AuditedDeleteExpiration:            plan.AuditedDeleteExpiration,
		IsObjectLockEnabled: plan.IsObjectLockEnabled,
		// StoragePolicy:                      plan.StoragePolicy,
		EnableAdvancedMetadataSearch:       plan.EnableAdvancedMetadataSearch,
		AdvancedMetadataSearchTargetName:   plan.AdvancedMetadataSearchTargetName,
		AdvancedMetadataSearchTargetStream: plan.AdvancedMetadataSearchTargetStream,
		LocalObjectMetadataReads:           plan.LocalObjectMetadataReads,
		VersioningStatus:                   plan.VersioningStatus,
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

	iam_group, _, err := r.client.GenClient.IamApi.IamServiceGetGroup(ctx).GroupName(state.GroupName.ValueString()).XEmcNamespace(state.Namespace.ValueString()).Execute()

	if err != nil {
		resp.Diagnostics.AddError("Error reading Group", err.Error())
		return
	}

	// data := r.getModel(iam_group)
	data := r.getModel(&clientgen.IamServiceCreateGroupResponseCreateGroupResultGroup{
		GroupId:    iam_group.GetGroupResult.Group.GroupId,
		GroupName:  iam_group.GetGroupResult.Group.GroupName,
		Arn:        iam_group.GetGroupResult.Group.Arn,
		CreateDate: iam_group.GetGroupResult.Group.CreateDate,
		Path:       iam_group.GetGroupResult.Group.Path,
	}, state.Namespace)
	// Save updated plan into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BucketResource) getModel(
	iam_group *clientgen.IamServiceCreateGroupResponseCreateGroupResultGroup,
	namespace types.String) models.BucketResourceModel {

	return models.BucketResourceModel{

	}
}

func (r *BucketResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Update operation is not supported
	resp.Diagnostics.AddError("[POST /iam?Action=UpdateGroup] UpdateGroup operation is not supported.", "Update operation is not supported.")
}

func (r *BucketResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Info(ctx, "deleting IAM Group")
	var state models.BucketResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, _, err := r.client.GenClient.IamApi.IamServiceDeleteGroup(ctx).GroupName(state.GroupName.ValueString()).XEmcNamespace(state.Namespace.ValueString()).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting IAM Group",
			err.Error(),
		)
	}
}

func (r *BucketResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Info(ctx, "importing IAM Group")
	parts := strings.SplitN(req.ID, ":", 2)
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Error importing IAM Group", "invalid format: expected 'group_name:namespace'")
		return
	}
	group_name := parts[0]
	namespace := parts[1]
	iam_group, _, err := r.client.GenClient.IamApi.IamServiceGetGroup(ctx).GroupName(group_name).XEmcNamespace(namespace).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading Group", err.Error())
		return
	}
	data := r.getModel(&clientgen.IamServiceCreateGroupResponseCreateGroupResultGroup{
		GroupId:    iam_group.GetGroupResult.Group.GroupId,
		GroupName:  iam_group.GetGroupResult.Group.GroupName,
		Arn:        iam_group.GetGroupResult.Group.Arn,
		CreateDate: iam_group.GetGroupResult.Group.CreateDate,
		Path:       iam_group.GetGroupResult.Group.Path,
	}, types.StringValue(namespace))
	// Save updated plan into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
