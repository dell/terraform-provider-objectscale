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
	"strings"

	"terraform-provider-objectscale/internal/clientgen"
	"terraform-provider-objectscale/internal/helper"
	"terraform-provider-objectscale/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &IAMRoleResource{}
var _ resource.ResourceWithImportState = &IAMRoleResource{}

func NewIAMRoleResource() resource.Resource {
	return &IAMRoleResource{}
}

// IAMRoleResource defines the resource implementation.
type IAMRoleResource struct {
	resourceProviderConfig
}

// models.IAMRoleResourceModel describes the resource data model.

func (r *IAMRoleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_role"
}

func (r *IAMRoleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "Manages an ObjectScale IAM Role.",
		MarkdownDescription: "Manages an ObjectScale IAM Role.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Description:         "Simple name identifying the Role. Required",
				MarkdownDescription: "Simple name identifying the Role. Required",
				Required:            true,
			},
			"namespace": schema.StringAttribute{
				Description:         "Namespace under which Role exists. Required",
				MarkdownDescription: "Namespace under which Role exists. Required",
				Required:            true,
			},

			"assume_role_policy_document": schema.StringAttribute{
				Description:         "The trust relationship policy document that grants an entity permission to assume the role.",
				MarkdownDescription: "The trust relationship policy document that grants an entity permission to assume the role.",
				Required:            true,
			},

			"max_session_duration": schema.Int32Attribute{
				Description:         "The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.",
				MarkdownDescription: "The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.",
				Optional:            true,
				Computed:            true,
			},

			"description": schema.StringAttribute{
				Description:         "The description of the role.",
				MarkdownDescription: "The description of the role.",
				Optional:            true,
				Computed:            true,
			},

			"path": schema.StringAttribute{
				Description:         "The path to the IAM Role. Defaults to / and only / is allowed",
				MarkdownDescription: "The path to the IAM Role. Defaults to / and only / is allowed",
				Computed:            true,
			},

			"permissions_boundary_arn": schema.StringAttribute{
				Description:         "Arn of the permissions boundary.",
				MarkdownDescription: "Arn of the permissions boundary.",
				Optional:            true,
				Computed:            true,
			},
			"permissions_boundary_type": schema.StringAttribute{
				Description:         "Type of the permissions boundary.",
				MarkdownDescription: "Type of the permissions boundary.",
				Optional:            true,
				Computed:            true,
			},

			"tags": schema.ListNestedAttribute{
				Description:         "The list of Tags associated with the role.. Default: []. Updatable.",
				MarkdownDescription: "The list of Tags associated with the role.. Default: []. Updatable.",
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Description:         "Key of the tag associated to the role.",
							MarkdownDescription: "Key of the tag associated to the role.",
							Optional:            true,
							Computed:            true,
						},
						"value": schema.StringAttribute{
							Description:         "Value of the tag associated to the role.",
							MarkdownDescription: "Value of the tag associated to the role.",
							Optional:            true,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (r *IAMRoleResource) tagJson(a models.Tags) clientgen.IamTagKeyValue {
	return clientgen.IamTagKeyValue{
		Key:   helper.ValueToPointer[string](a.Key),
		Value: helper.ValueToPointer[string](a.Value),
	}
}

func (r *IAMRoleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.IAMRoleResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	creq := r.client.GenClient.IamApi.IamServiceCreateRole(ctx).
		RoleName(plan.Name.ValueString()).
		XEmcNamespace(plan.Namespace.ValueString()).
		AssumeRolePolicyDocument(plan.AssumeRolePolicyDocument.ValueString())
	if max_session_duration := helper.ValueToPointer[int32](plan.MaxSessionDuration); max_session_duration != nil {
		creq = creq.MaxSessionDuration(*max_session_duration)
	}
	if description := helper.ValueToPointer[string](plan.Description); description != nil {
		creq = creq.Description(*description)
	}
	if permissions_boundary := helper.ValueToPointer[string](plan.PermissionsBoundaryArn); permissions_boundary != nil {
		creq = creq.PermissionsBoundary(*permissions_boundary)
	}
	if ptags := helper.ValueListTransform(plan.Tags, r.tagJson); len(ptags) > 0 {
		creq = creq.TagsMemberN(ptags)
	}

	_, _, err := creq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error creating Role", err.Error())
		return
	}

	iam_role, _, err := r.client.GenClient.IamApi.IamServiceGetRole(ctx).
		RoleName(plan.Name.ValueString()).
		XEmcNamespace(plan.Namespace.ValueString()).
		Execute()
	data := r.getModel(&clientgen.IamRole{
		RoleName:                 iam_role.GetRoleResult.Role.RoleName,
		AssumeRolePolicyDocument: iam_role.GetRoleResult.Role.AssumeRolePolicyDocument,
		MaxSessionDuration:       iam_role.GetRoleResult.Role.MaxSessionDuration,
		Description:              iam_role.GetRoleResult.Role.Description,
		PermissionsBoundary:      iam_role.GetRoleResult.Role.PermissionsBoundary,
		Tags:                     iam_role.GetRoleResult.Role.Tags,
	}, plan.Namespace)

	if err != nil {
		resp.Diagnostics.AddError("Error reading role", err.Error())
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IAMRoleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.IAMRoleResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	iam_role, _, err := r.client.GenClient.IamApi.IamServiceGetRole(ctx).
		RoleName(state.Name.ValueString()).
		XEmcNamespace(state.Namespace.ValueString()).
		Execute()
	data := r.getModel(&clientgen.IamRole{
		RoleName:                 iam_role.GetRoleResult.Role.RoleName,
		AssumeRolePolicyDocument: iam_role.GetRoleResult.Role.AssumeRolePolicyDocument,
		MaxSessionDuration:       iam_role.GetRoleResult.Role.MaxSessionDuration,
		Description:              iam_role.GetRoleResult.Role.Description,
		PermissionsBoundary:      iam_role.GetRoleResult.Role.PermissionsBoundary,
		Tags:                     iam_role.GetRoleResult.Role.Tags,
	}, state.Namespace)

	if err != nil {
		resp.Diagnostics.AddError("Error reading role", err.Error())
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IAMRoleResource) getModel(
	iam_role *clientgen.IamRole,
	namespace types.String) models.IAMRoleResourceModel {
	var permissionsBoundaryArn, permissionsBoundaryType basetypes.StringValue

	// Check if PermissionsBoundary exists
	if iam_role.PermissionsBoundary != nil {
		permissionsBoundaryArn = helper.TfStringNN(iam_role.PermissionsBoundary.PermissionsBoundaryArn)
		permissionsBoundaryType = helper.TfStringNN(iam_role.PermissionsBoundary.PermissionsBoundaryType)
	} else {
		// Set empty values if missing
		permissionsBoundaryArn = types.StringValue("")
		permissionsBoundaryType = types.StringValue("")
	}
	return models.IAMRoleResourceModel{

		Name:                     helper.TfStringNN(iam_role.RoleName),
		Namespace:                namespace,
		AssumeRolePolicyDocument: helper.TfStringNN(iam_role.AssumeRolePolicyDocument),
		Description:              helper.TfStringNN(iam_role.Description),
		MaxSessionDuration:       helper.TfInt32NN(iam_role.MaxSessionDuration),
		Path:                     helper.TfStringNN(iam_role.Path),
		PermissionsBoundaryType:  permissionsBoundaryType,
		PermissionsBoundaryArn:   permissionsBoundaryArn,
		Tags: helper.ListNotNull(iam_role.Tags, func(tag clientgen.IamTagKeyValue) types.Object {
			return helper.Object(models.IAMRoleTag{
				Key:   helper.TfStringNN(tag.Key),
				Value: helper.TfStringNN(tag.Value),
			})
		}),
	}
}

func (r *IAMRoleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state models.IAMRoleResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// check for changes in plan and state
	if !(helper.IsChangedNN(plan.MaxSessionDuration, state.MaxSessionDuration) || helper.IsChangedNN(plan.Description, state.Description)) {
		resp.Diagnostics.AddError("Only 'max_session_duration' or 'description' can be changed", "invalid attribute change detected")
		return
	}

	updReq := r.client.GenClient.IamApi.IamServiceUpdateRole(ctx).
		RoleName(plan.Name.ValueString()).
		XEmcNamespace(plan.Namespace.ValueString())

	if max_session_duration := helper.ValueToPointer[int32](plan.MaxSessionDuration); max_session_duration != nil {
		updReq = updReq.MaxSessionDuration(*max_session_duration)
	}
	if description := helper.ValueToPointer[string](plan.Description); description != nil {
		updReq = updReq.Description(*description)
	}

	_, _, err := updReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error updating Role", err.Error())
		return
	}

	iam_role, _, err := r.client.GenClient.IamApi.IamServiceGetRole(ctx).
		RoleName(plan.Name.ValueString()).
		XEmcNamespace(plan.Namespace.ValueString()).
		Execute()
	data := r.getModel(&clientgen.IamRole{
		RoleName:                 iam_role.GetRoleResult.Role.RoleName,
		AssumeRolePolicyDocument: iam_role.GetRoleResult.Role.AssumeRolePolicyDocument,
		MaxSessionDuration:       iam_role.GetRoleResult.Role.MaxSessionDuration,
		Description:              iam_role.GetRoleResult.Role.Description,
		PermissionsBoundary:      iam_role.GetRoleResult.Role.PermissionsBoundary,
		Tags:                     iam_role.GetRoleResult.Role.Tags,
	}, plan.Namespace)

	if err != nil {
		resp.Diagnostics.AddError("Error reading role", err.Error())
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IAMRoleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Info(ctx, "deleting IAM Role")
	var state models.IAMRoleResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, _, err := r.client.GenClient.IamApi.IamServiceDeleteRole(ctx).RoleName(state.Name.ValueString()).XEmcNamespace(state.Namespace.ValueString()).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting IAM Role",
			err.Error(),
		)
	}
}

func (r *IAMRoleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Info(ctx, "importing IAM Role")
	parts := strings.SplitN(req.ID, ":", 2)
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Error importing IAM Role", "invalid format: expected 'role_name:namespace'")
		return
	}
	role_name := parts[0]
	namespace := parts[1]
	iam_role, _, err := r.client.GenClient.IamApi.IamServiceGetRole(ctx).
		RoleName(role_name).
		XEmcNamespace(namespace).
		Execute()
	data := r.getModel(&clientgen.IamRole{
		RoleName:                 iam_role.GetRoleResult.Role.RoleName,
		AssumeRolePolicyDocument: iam_role.GetRoleResult.Role.AssumeRolePolicyDocument,
		MaxSessionDuration:       iam_role.GetRoleResult.Role.MaxSessionDuration,
		Description:              iam_role.GetRoleResult.Role.Description,
		PermissionsBoundary:      iam_role.GetRoleResult.Role.PermissionsBoundary,
		Tags:                     iam_role.GetRoleResult.Role.Tags,
	}, types.StringValue(namespace))

	if err != nil {
		resp.Diagnostics.AddError("Error reading role", err.Error())
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
