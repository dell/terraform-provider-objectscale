/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

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
var _ resource.Resource = &IAMGroupResource{}
var _ resource.ResourceWithImportState = &IAMGroupResource{}

func NewIAMGroupResource() resource.Resource {
	return &IAMGroupResource{}
}

// IAMGroupResource defines the resource implementation.
type IAMGroupResource struct {
	client *client.Client
}

// models.IAMGroupResourceModel describes the resource data model.

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

func (r *IAMGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "creating group")
	var plan models.IAMGroupResourceModel

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
	var state models.IAMGroupResourceModel

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
	namespace types.String) models.IAMGroupResourceModel {

	return models.IAMGroupResourceModel{

		GroupId:    helper.TfStringNN(iam_group.GroupId),
		GroupName:  helper.TfStringNN(iam_group.GroupName),
		Arn:        helper.TfStringNN(iam_group.Arn),
		Namespace:  namespace,
		CreateDate: helper.TfStringNN(iam_group.CreateDate),
		Path:       helper.TfStringNN(iam_group.Path),
	}
}

func (r *IAMGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Update operation is not supported
	resp.Diagnostics.AddError("[POST /iam?Action=UpdateGroup] UpdateGroup operation is not supported.", "Update operation is not supported.")
}

func (r *IAMGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Info(ctx, "deleting IAM Group")
	var state models.IAMGroupResourceModel

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
