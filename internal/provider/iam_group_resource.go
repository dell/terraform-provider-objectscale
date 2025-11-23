// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"

	"terraform-provider-objectscale/internal/client"
	"terraform-provider-objectscale/internal/clientgen"
	"terraform-provider-objectscale/internal/helper"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &IAMGroupResource{}
var _ resource.ResourceWithImportState = &IAMGroupResource{}

func NewIAMGroupResource() resource.Resource {
	return &IAMGroupResource{}
}

// IAMGroupResource defines the resource implementation.
type IAMGroupResource struct {
	client *client.Client
}

// IAMGroupResourceModel describes the resource data model.
type IAMGroupResourceModel struct {
	GroupName  types.String `tfsdk:"name"`
	GroupId    types.String `tfsdk:"id"`
	Arn        types.String `tfsdk:"arn"`
	Path       types.String `tfsdk:"path"`
	CreateDate types.String `tfsdk:"create_date"`
	Namespace  types.String `tfsdk:"namespace"`
}

func (r *IAMGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_group"
}

func (r *IAMGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Description:         "Simple name identifying the group. Required",
				MarkdownDescription: "Simple name identifying the group. Required",
				Required:            true,
			},
			"namespace": schema.StringAttribute{
				Description:         "Namespace under which group exists. Required",
				MarkdownDescription: "Namespace under which group exists. Required",
				Required:            true,
			},

			"path": schema.StringAttribute{
				Description:         "The path to the IAM Group. Defaults to / and only / is allowed",
				MarkdownDescription: "The path to the IAM Group. Defaults to / and only / is allowed",
				Computed:            true,
			},

			"arn": schema.StringAttribute{
				Description:         "Arn that identifies the Group. Computed",
				MarkdownDescription: "Arn that identifies the Group. Computed",
				Computed:            true,
			},

			"id": schema.StringAttribute{
				Description:         "Unique Id associated with the Group.",
				MarkdownDescription: "Unique Id associated with the Group.",
				Computed:            true,
			},

			"create_date": schema.StringAttribute{
				Description:         "ISO 8601 format DateTime when group was created.",
				MarkdownDescription: "ISO 8601 format DateTime when group was created.",
				Computed:            true,
			},
		},
	}
}

func (r *IAMGroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *IAMGroupResource) modelToJson(plan IAMGroupResourceModel) clientgen.IamServiceCreateGroupResponseCreateGroupResultGroup {
	return clientgen.IamServiceCreateGroupResponseCreateGroupResultGroup{
		GroupName:  helper.ValueToPointer[string](plan.GroupName),
		GroupId:    helper.ValueToPointer[string](plan.GroupId),
		Path:       helper.ValueToPointer[string](plan.Path),
		CreateDate: helper.ValueToPointer[string](plan.CreateDate),
		Arn:        helper.ValueToPointer[string](plan.Arn),
	}
}

func (r *IAMGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "creating group")
	var plan IAMGroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	iam_group, _, err := r.client.GenClient.IamApi.IamServiceCreateGroup(ctx).GroupName(plan.GroupName.ValueString()).XEmcNamespace(plan.Namespace.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error creating Group", err.Error())
		return
	}

	data := r.getModel(&clientgen.IamServiceCreateGroupResponseCreateGroupResultGroup{
		GroupId:    iam_group.CreateGroupResult.Group.GroupId,
		GroupName:  iam_group.CreateGroupResult.Group.GroupName,
		Arn:        iam_group.CreateGroupResult.Group.Arn,
		CreateDate: iam_group.CreateGroupResult.Group.CreateDate,
		Path:       iam_group.CreateGroupResult.Group.Path,
	}, plan.Namespace)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IAMGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IAMGroupResourceModel

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

func (r *IAMGroupResource) getModel(
	iam_group *clientgen.IamServiceCreateGroupResponseCreateGroupResultGroup,
	namespace types.String) IAMGroupResourceModel {

	return IAMGroupResourceModel{

		GroupId:    helper.TfStringNN(iam_group.GroupId),
		GroupName:  helper.TfStringNN(iam_group.GroupName),
		Arn:        helper.TfStringNN(iam_group.Arn),
		Namespace:  namespace,
		CreateDate: helper.TfStringNN(iam_group.CreateDate),
		Path:       helper.TfStringNN(iam_group.Path),
	}
}

func (r *IAMGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// resp.Diagnostics.AddError("[Update] Update operation is not available.", "Update operation is not available.")
	tflog.Info(ctx, "updating group")
	// TODO: Add update logic
	var plan, state IAMGroupResourceModel

	// Read Terraform plan and state data into the models
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// To prevent the non-updatable fields from being changed
	if !plan.GroupName.Equal(state.GroupName) {
		resp.Diagnostics.AddError("Error updating group", "Fields of `name`, `arn` and `create_date` are not updatable")
		return
	}

	iam_group, _, err := r.client.GenClient.IamApi.IamServiceGetGroup(ctx).GroupName(state.GroupName.ValueString()).XEmcNamespace(state.Namespace.ValueString()).Execute()

	if err != nil {
		resp.Diagnostics.AddError("Error reading group", err.Error())
		return
	}

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

func (r *IAMGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Info(ctx, "deleting IAM Group")
	var state IAMGroupResourceModel

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

func (r *IAMGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// tflog.Info(ctx, "importing IAM user")

	// iam_user, _, err := r.client.GenClient.IamApi.IamServiceGetUser(ctx).UserName(req.ID.ValueString()).XEmcNamespace(state.Namespace.ValueString()).Execute()

	// if err != nil {
	// 	resp.Diagnostics.AddError("Error reading user", err.Error())
	// 	return
	// }

	// data := r.getModel(namespace, types.StringNull())
	// // Save updated plan into Terraform state
	// resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	resp.Diagnostics.AddError("[Import] Import operation is not available.", "Import operation is not available.")

}
