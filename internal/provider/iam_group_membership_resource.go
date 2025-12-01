// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"

	"terraform-provider-objectscale/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &IAMGroupMembershipResource{}
var _ resource.ResourceWithImportState = &IAMGroupMembershipResource{}

func NewIAMGroupMembershipResource() resource.Resource {
	return &IAMGroupMembershipResource{}
}

// IAMGroupMembershipResource defines the resource implementation.
type IAMGroupMembershipResource struct {
	client *client.Client
}

// IAMGroupMembershipResourceModel describes the resource data model.
type IAMGroupMembershipResourceModel struct {
	GroupName types.String `tfsdk:"name"`
	Namespace types.String `tfsdk:"namespace"`
	User      types.String `tfsdk:"user"`

	Users types.Set `tfsdk:"users"`
}

func (r *IAMGroupMembershipResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_group_membership"
}

func (r *IAMGroupMembershipResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"user": schema.StringAttribute{
				Description:         "User to be added to the group. Required",
				MarkdownDescription: "User to be added to the group. Required",
				Required:            true,
			},

			//Optional expose the current group membership list as computed for read only use
			"users": schema.ListAttribute{
				Description:         "List of users who are members of the group.",
				MarkdownDescription: "List of users who are members of the group.",
				ElementType:         types.StringType,
				Computed:            true,
			},
		},
	}
}

func (r *IAMGroupMembershipResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// func (r *IAMGroupMembershipResource) modelToJson(plan IAMGroupMembershipResourceModel) clientgen.IamServiceCreateGroupResponseCreateGroupResultGroup {
// 	return clientgen.IamServiceCreateGroupResponseCreateGroupResultGroup{
// 		GroupName:  helper.ValueToPointer[string](plan.GroupName),
// 		GroupId:    helper.ValueToPointer[string](plan.GroupId),
// 		Path:       helper.ValueToPointer[string](plan.Path),
// 		CreateDate: helper.ValueToPointer[string](plan.CreateDate),
// 		Arn:        helper.ValueToPointer[string](plan.Arn),
// 	}
// }

func (r *IAMGroupMembershipResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "creating user")
	var plan IAMGroupMembershipResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, _, err := r.client.GenClient.IamApi.IamServiceAddUserToGroup(ctx).GroupName(plan.GroupName.ValueString()).XEmcNamespace(plan.Namespace.ValueString()).UserName(plan.User.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error creating user", err.Error())
		return
	}

	// Read full membership list
	members, _, err := r.client.GenClient.IamApi.IamServiceGetGroup(ctx).GroupName(plan.GroupName.ValueString()).XEmcNamespace(plan.Namespace.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading group members", err.Error())
		return
	}

	for _, user := range members.GetGroupResult.Users {
		plan.Users = append(plan.Users, types.StringValue(user))
	}
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IAMGroupMembershipResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state IAMGroupMembershipResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	members, _, err := r.client.GenClient.IamApi.IamServiceGetGroup(ctx).GroupName(state.GroupName.ValueString()).XEmcNamespace(state.Namespace.ValueString()).Execute()

	if err != nil {
		// If group is gone, remove:
		// resp.State.RemoveResource(ctx); return
		resp.Diagnostics.AddError("Read membership failed", err.Error())
		return
	}

	state.Users = make([]types.String, 0, len(members.GetGroupResult.Users))
	for _, m := range members.GetGroupResult.Users {
		state.Users = append(state.Users, types.StringValue(m))
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)

}

// func (r *IAMGroupMembershipResource) getModel(
// 	// iam_group *clientgen.IamServiceCreateUserResponseCreateUserResultUser,
// 	iam_group *clientgen.IamServiceCreateGroupResponseCreateGroupResultGroup,
// 	namespace types.String) IAMGroupMembershipResourceModel {

// 	return IAMGroupMembershipResourceModel{

// 		GroupId:    helper.TfStringNN(iam_group.GroupId),
// 		GroupName:  helper.TfStringNN(iam_group.GroupName),
// 		Arn:        helper.TfStringNN(iam_group.Arn),
// 		Namespace:  namespace,
// 		CreateDate: helper.TfStringNN(iam_group.CreateDate),
// 		Path:       helper.TfStringNN(iam_group.Path),
// 	}
// }

func (r *IAMGroupMembershipResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data IAMGroupMembershipResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IAMGroupMembershipResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IAMGroupMembershipResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// API call: remove USER from GROUP
	_, _, err := r.client.GenClient.IamApi.IamServiceRemoveUserFromGroup(ctx).GroupName(state.GroupName.ValueString()).XEmcNamespace(state.Namespace.ValueString()).UserName(state.User.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Remove user failed", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *IAMGroupMembershipResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
