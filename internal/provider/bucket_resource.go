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
				Optional:            true,
				Computed:            true,
			},
			"fs_access_enabled": schema.BoolAttribute{
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
			},
			"api_type": schema.StringAttribute{
				Description:         "API type for the bucket.",
				MarkdownDescription: "API type for the bucket.",
				Computed:            true,
			},
			"tag": schema.ListNestedAttribute{
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
			},
			"is_enabled": schema.BoolAttribute{
				Description:         "Is search metadata enabled.",
				MarkdownDescription: "Is search metadata enabled.",
				Optional:            true,
				Computed:            true,
			},
			"md_tokens": schema.BoolAttribute{
				Description:         "Metadata tokens for advanced search.",
				MarkdownDescription: "Metadata tokens for advanced search.",
				Optional:            true,
				Computed:            true,
			},
			"max_keys": schema.Int64Attribute{
				Description:         "Maximum number of keys for search.",
				MarkdownDescription: "Maximum number of keys for search.",
				Optional:            true,
				Computed:            true,
			},
			"metadata": schema.ListNestedAttribute{
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
				Description:         "Days after which audited delete expires.",
				MarkdownDescription: "Days after which audited delete expires.",
				Optional:            true,
				Computed:            true,
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
			"is_encryption_enabled": schema.StringAttribute{
				Description:         "Enable server-side encryption.",
				MarkdownDescription: "Enable server-side encryption.",
				Optional:            true,
				Computed:            true,
			},
			"default_retention": schema.Int64Attribute{
				Description:         "Default retention period.",
				MarkdownDescription: "Default retention period.",
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
				Description:         "Enable local metadata reads.",
				MarkdownDescription: "Enable local metadata reads.",
				Optional:            true,
				Computed:            true,
			},
			"versioning_status": schema.StringAttribute{
				Description:         "Versioning status (Enabled/Suspended).",
				MarkdownDescription: "Versioning status (Enabled/Suspended).",
				Optional:            true,
				Computed:            true,
			},
		},
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
	reqBody := r.modelToJson(ctx, plan)

	bucket, _, err := r.client.GenClient.BucketApi.BucketServiceCreateBucket(ctx).BucketServiceCreateBucketRequest(reqBody).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error creating bucket", err.Error())
		return
	}

	// Call the API to get buckets with pagination
	apiReq := r.client.GenClient.BucketApi.BucketServiceGetBuckets(ctx).Namespace(plan.Namespace.ValueString())
	if *bucket.Name != "" {
		apiReq = apiReq.Name(*bucket.Name + "*")
	}

	allBuckets, err := helper.GetAllInstances(apiReq)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Buckets",
			fmt.Sprintf("An error was encountered reading buckets from ObjectScale IAM: %s", err.Error()),
		)
		return
	}

	// Filter allBuckets to only include the bucket with the matching ID
	filteredBuckets := make([]clientgen.BucketServiceGetBucketsResponseObjectBucketInner, 0, len(allBuckets))
	for _, b := range allBuckets {
		if b.Id != nil && bucket.Id != nil && *b.Id == *bucket.Id {
			filteredBuckets = append(filteredBuckets, b)
		}
	}

	// If a name prefix was provided and no buckets were found, return an error
	if plan.Name.ValueString() != "" && len(filteredBuckets) == 0 {
		resp.Diagnostics.AddError(
			"No buckets found with the specified prefix",
			fmt.Sprintf("No buckets found in namespace '%s' with prefix '%s'. Please check the prefix.", plan.Namespace, plan.Name),
		)
		return
	}

	data := getBucketToModel(filteredBuckets[0])

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BucketResource) tagListJson(in models.TagModel) clientgen.BucketServiceCreateBucketRequestTagSetInner {
	return clientgen.BucketServiceCreateBucketRequestTagSetInner{
		Key:   helper.ValueToPointer[string](in.Key),
		Value: helper.ValueToPointer[string](in.Value),
	}
}

func (r *BucketResource) metadataJson(in models.MetadataModel) clientgen.BucketServiceCreateBucketRequestMetadataInner {
	return clientgen.BucketServiceCreateBucketRequestMetadataInner{
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

func (r *BucketResource) modelToJson(ctx context.Context, plan models.BucketResourceModel) clientgen.BucketServiceCreateBucketRequest {

	minMaxGovernor := helper.ValueObjectTransform(plan.MinMaxGovernor, r.minMaxGovernorJson)

	return clientgen.BucketServiceCreateBucketRequest{
		Name:                               plan.Name.ValueString(),
		Owner:                              helper.ValueToPointer[string](plan.Owner),
		Namespace:                          helper.ValueToPointer[string](plan.Namespace),
		Vpool:                              helper.ValueToPointer[string](plan.ReplicationGroup),
		FilesystemEnabled:                  helper.ValueToPointer[bool](plan.FsAccessEnabled),
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
		Metadata:                           helper.ValueListTransform(plan.Metadata, r.metadataJson),
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

func (r *BucketResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.BucketResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call the API to get buckets with pagination
	apiReq := r.client.GenClient.BucketApi.BucketServiceGetBuckets(ctx).Namespace(state.Namespace.ValueString())
	if state.Name.ValueString() != "" {
		apiReq = apiReq.Name(state.Name.ValueString() + "*")
	}

	allBuckets, err := helper.GetAllInstances(apiReq)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Buckets",
			fmt.Sprintf("An error was encountered reading buckets from ObjectScale IAM: %s", err.Error()),
		)
		return
	}

	// Filter allBuckets to only include the bucket with the matching ID
	filteredBuckets := make([]clientgen.BucketServiceGetBucketsResponseObjectBucketInner, 0, len(allBuckets))
	for _, b := range allBuckets {
		if b.Id != nil && state.Id.ValueString() != "" && *b.Id == state.Id.ValueString() {
			filteredBuckets = append(filteredBuckets, b)
		}
	}

	// If a name prefix was provided and no buckets were found, return an error
	// if state.Name.ValueString() != "" && len(filteredBuckets) == 0 {
	// 	resp.Diagnostics.AddError(
	// 		"No buckets found with the specified prefix",
	// 		fmt.Sprintf("No buckets found in namespace '%s' with prefix '%s'. Please check the prefix.", state.Namespace.ValueString(), state.Name.ValueString()),
	// 	)
	// 	return
	// }

	data := getBucketToModel(filteredBuckets[0])

	// Save updated plan into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// mapBucketToModel maps a BucketServiceGetBucketsResponseObjectBucketInner to models.BucketModel.
func getBucketToModel(b clientgen.BucketServiceGetBucketsResponseObjectBucketInner) models.BucketResourceModel {
	m := models.BucketResourceModel{
		Id:               helper.TfString(b.Id),
		Owner:            helper.TfString(b.Owner),
		Name:             helper.TfString(b.Name),
		ReplicationGroup: helper.TfString(b.Vpool),
		Namespace:        helper.TfString(b.Namespace),
		BlockSize:        helper.TfInt64(b.BlockSize),
		NotificationSize: helper.TfInt64(b.NotificationSize),
		FsAccessEnabled:  helper.TfBool(b.FsAccessEnabled),
		//HeadType:         helper.TfString(b.HeadType),
		Tag: helper.ListNotNull(b.TagSet,
			func(v clientgen.BucketServiceCreateBucketRequestTagSetInner) types.Object {
				return helper.Object(models.Tags{
					Key:   helper.TfStringNN(v.Key),
					Value: helper.TfStringNN(v.Value),
				})
			}),
		IsEncryptionEnabled:               helper.TfString(b.IsEncryptionEnabled),
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
		Metadata: helper.ListNotNull(b.SearchMetadata.Metadata,
			func(v clientgen.BucketServiceCreateBucketRequestMetadataInner) types.Object {
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
		//AuditedDeleteExpiration:            helper.TfInt64(b.AuditedDeleteExpiration),
		IsObjectLockEnabled: helper.TfBool(b.IsObjectLockEnabled),
		//StoragePolicy:                      helper.TfString(b.StoragePolicy),
		EnableAdvancedMetadataSearch:       helper.TfBool(b.EnableAdvancedMetadataSearch),
		AdvancedMetadataSearchTargetName:   helper.TfString(b.AdvancedMetadataSearchTargetName),
		AdvancedMetadataSearchTargetStream: helper.TfString(b.AdvancedMetadataSearchTargetStream),
		LocalObjectMetadataReads:           helper.TfBool(b.LocalObjectMetadataReads),
		VersioningStatus:                   helper.TfString(b.VersioningStatus),
		Created:                            helper.TfString(b.Created),
	}
	return m
}

func (r *BucketResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Update operation is not supported
	resp.Diagnostics.AddError("Update Bucket operation is not supported.", "Update operation is not supported.")
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
	// Call the API to get buckets with pagination
	apiReq := r.client.GenClient.BucketApi.BucketServiceGetBuckets(ctx).Namespace(namespace)
	if bucket_name != "" {
		apiReq = apiReq.Name(bucket_name + "*")
	}

	allBuckets, err := helper.GetAllInstances(apiReq)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Buckets",
			fmt.Sprintf("An error was encountered reading buckets from ObjectScale IAM: %s", err.Error()),
		)
		return
	}

	// Filter allBuckets to only include the bucket with the matching ID
	filteredBuckets := make([]clientgen.BucketServiceGetBucketsResponseObjectBucketInner, 0, len(allBuckets))
	for _, b := range allBuckets {
		if b.Id != nil && bucket_name != "" && *b.Name == bucket_name {
			filteredBuckets = append(filteredBuckets, b)
		}
	}

	// If a name prefix was provided and no buckets were found, return an error
	if bucket_name != "" && len(filteredBuckets) == 0 {
		resp.Diagnostics.AddError(
			"No buckets found with the specified prefix",
			fmt.Sprintf("No buckets found in namespace '%s' with prefix '%s'. Please check the prefix.", namespace, bucket_name),
		)
		return
	}

	data := getBucketToModel(filteredBuckets[0])
	// Save updated plan into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
