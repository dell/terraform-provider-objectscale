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

	"terraform-provider-objectscale/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
}

func (r *IAMGroupMembershipResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_group_membership"
}

func (r *IAMGroupMembershipResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema.Description = "Resource for managing IAM Group Memberships in ObjectScale."

	resp.Schema = schema.Schema{
		Description:         "Manages Group Membership for a User.",
		MarkdownDescription: "Manages Group Membership for a User.",
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

func (r *IAMGroupMembershipResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
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

	// Check that user was added
	found := false
	for _, user := range members.GetGroupResult.Users {
		if user.UserName != nil && *user.UserName == plan.User.ValueString() {
			found = true
			break
		}
	}
	if !found {
		resp.Diagnostics.AddError("Error adding user to group", "User was not found in group after addition.")
		return
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

	// Check that user is still a member
	found := false
	for _, user := range members.GetGroupResult.Users {
		if user.UserName != nil && *user.UserName == state.User.ValueString() {
			found = true
			break
		}
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)

}

func (r *IAMGroupMembershipResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Update operation is not supported
	resp.Diagnostics.AddError("Update Group membership operation is not supported.", "Update operation is not supported.")
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
	//check that user was removed
	members, _, err := r.client.GenClient.IamApi.IamServiceGetGroup(ctx).GroupName(state.GroupName.ValueString()).XEmcNamespace(state.Namespace.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading group members", err.Error())
		return
	}

	// Check that user was removed
	found := false
	for _, user := range members.GetGroupResult.Users {
		if user.UserName != nil && *user.UserName == state.User.ValueString() {
			found = true
			break
		}
	}
	if found {
		resp.Diagnostics.AddError("Error removing user from group", "User is still found in group after removal.")
		return
	}

	// Remove resource from state

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
