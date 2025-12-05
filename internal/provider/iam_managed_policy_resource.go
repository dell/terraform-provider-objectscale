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
	"terraform-provider-objectscale/internal/helper"
	"terraform-provider-objectscale/internal/models"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var _ resource.Resource = &IAMManagedPolicyResource{}
var _ resource.ResourceWithImportState = &IAMManagedPolicyResource{}

// IAMManagedPolicyResource is the resource implementation.
type IAMManagedPolicyResource struct {
	client *client.Client
}

// NewIAMManagedPolicyResource is a helper function to simplify the provider implementation.
func NewIAMManagedPolicyResource() resource.Resource {
	return &IAMManagedPolicyResource{}
}

// Metadata returns the resource type name.
func (r *IAMManagedPolicyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_managed_policy"
}

// Configure adds the provider configured client to the resource.
func (r *IAMManagedPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

// Schema defines the schema for the resource.
func (r *IAMManagedPolicyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "Manages IAM managed policies for ObjectScale entities (user, group, or role).",
		MarkdownDescription: "Manages IAM managed policies for ObjectScale entities (user, group, or role).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Unique identifier for the IAM managed policy resource.",
				MarkdownDescription: "Unique identifier for the IAM managed policy resource.",
				Computed:            true,
			},
			"namespace": schema.StringAttribute{
				Description:         "Namespace to which the IAM entity belongs.",
				MarkdownDescription: "Namespace to which the IAM entity belongs.",
				Required:            true,
			},
			"username": schema.StringAttribute{
				Description:         "Name of the user. Exactly one of username, groupname, or rolename must be set.",
				MarkdownDescription: "Name of the user. Exactly one of username, groupname, or rolename must be set.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.ExactlyOneOf(
						path.MatchRoot("groupname"),
						path.MatchRoot("rolename"),
					),
				},
			},
			"groupname": schema.StringAttribute{
				Description:         "Name of the group. Exactly one of username, groupname, or rolename must be set.",
				MarkdownDescription: "Name of the group. Exactly one of username, groupname, or rolename must be set.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.ExactlyOneOf(
						path.MatchRoot("username"),
						path.MatchRoot("rolename"),
					),
				},
			},
			"rolename": schema.StringAttribute{
				Description:         "Name of the role. Exactly one of username, groupname, or rolename must be set.",
				MarkdownDescription: "Name of the role. Exactly one of username, groupname, or rolename must be set.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.ExactlyOneOf(
						path.MatchRoot("username"),
						path.MatchRoot("groupname"),
					),
				},
			},
			"policy_arns": schema.ListAttribute{
				Description:         "List of IAM managed policy arns to associate with the entity.",
				MarkdownDescription: "List of IAM managed policy arns to associate with the entity.",
				Required:            true,
				ElementType:         types.StringType,
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *IAMManagedPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.IAMManagedPolicyResourceModel

	// Get current state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Determine namespace
	namespace := state.Namespace.ValueString()

	// Determine entity type and name
	var entityType, entityName string
	if !state.Username.IsNull() && !state.Username.IsUnknown() {
		entityType = "User"
		entityName = state.Username.ValueString()
	} else if !state.Groupname.IsNull() && !state.Groupname.IsUnknown() {
		entityType = "Group"
		entityName = state.Groupname.ValueString()
	} else if !state.Rolename.IsNull() && !state.Rolename.IsUnknown() {
		entityType = "Role"
		entityName = state.Rolename.ValueString()
	}

	// Call ListAttached<entity>Policies API
	var policyARNs []string
	var marker string

	switch entityType {
	case "User":
		for {
			listReq := r.client.GenClient.IamApi.IamServiceListAttachedUserPolicies(ctx).
				XEmcNamespace(namespace).
				UserName(entityName)

			if marker != "" {
				listReq = listReq.Marker(marker)
			}

			listResp, _, err := listReq.Execute()
			if err != nil {
				resp.Diagnostics.AddError("Read Error", fmt.Sprintf("Failed to list policy arns: %s", err.Error()))
				return
			}

			for _, p := range listResp.ListAttachedUserPoliciesResult.AttachedPolicies {
				policyARNs = append(policyARNs, *p.PolicyArn)
			}

			markerPtr := listResp.ListAttachedUserPoliciesResult.Marker
			if markerPtr == nil || *markerPtr == "" {
				break
			}
			marker = *markerPtr
		}
	case "Group":
		for {
			listReq := r.client.GenClient.IamApi.IamServiceListAttachedGroupPolicies(ctx).
				XEmcNamespace(namespace).
				GroupName(entityName)

			if marker != "" {
				listReq = listReq.Marker(marker)
			}

			listResp, _, err := listReq.Execute()
			if err != nil {
				resp.Diagnostics.AddError("Read Error", fmt.Sprintf("Failed to list policy arns: %s", err.Error()))
				return
			}

			for _, p := range listResp.ListAttachedGroupPoliciesResult.AttachedPolicies {
				policyARNs = append(policyARNs, *p.PolicyArn)
			}

			markerPtr := listResp.ListAttachedGroupPoliciesResult.Marker
			if markerPtr == nil || *markerPtr == "" {
				break
			}
			marker = *markerPtr
		}
	case "Role":
		for {
			listReq := r.client.GenClient.IamApi.IamServiceListAttachedRolePolicies(ctx).
				XEmcNamespace(namespace).
				RoleName(entityName)

			if marker != "" {
				listReq = listReq.Marker(marker)
			}

			listResp, _, err := listReq.Execute()
			if err != nil {
				resp.Diagnostics.AddError("Read Error", fmt.Sprintf("Failed to list policy arns: %s", err.Error()))
				return
			}

			for _, p := range listResp.ListAttachedRolePoliciesResult.AttachedPolicies {
				policyARNs = append(policyARNs, *p.PolicyArn)
			}

			markerPtr := listResp.ListAttachedRolePoliciesResult.Marker
			if markerPtr == nil || *markerPtr == "" {
				break
			}
			marker = *markerPtr
		}
	}

	if policyARNs == nil {
		policyARNs = []string{}
	}

	// Update state
	state.PolicyARNs = make([]types.String, len(policyARNs))
	for i, arn := range policyARNs {
		state.PolicyARNs[i] = types.StringValue(arn)
	}
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Create creates the resource and sets the updated Terraform state on success.
func (r *IAMManagedPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.IAMManagedPolicyResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	updatedModel, err := helper.ApplyPolicyARNs(r.client, ctx, plan, nil)
	if err != nil {
		resp.Diagnostics.AddError("Create Error", err.Error())
		return
	}

	diags = resp.State.Set(ctx, updatedModel)
	resp.Diagnostics.Append(diags...)
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *IAMManagedPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan models.IAMManagedPolicyResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state models.IAMManagedPolicyResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	updatedModel, err := helper.ApplyPolicyARNs(r.client, ctx, plan, &state)
	if err != nil {
		resp.Diagnostics.AddError("Update Error", err.Error())
		return
	}

	diags = resp.State.Set(ctx, updatedModel)
	resp.Diagnostics.Append(diags...)
}

// Delete deletes the resource and removes the Terraform state.
func (r *IAMManagedPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Info(ctx, "Deleting IAM Managed Policy resource")

	var state models.IAMManagedPolicyResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.PolicyARNs = []types.String{}
	_, err := helper.ApplyPolicyARNs(r.client, ctx, state, nil)
	if err != nil {
		resp.Diagnostics.AddError("Delete Error", err.Error())
		return
	}

	// Remove resource from Terraform state
	resp.State.RemoveResource(ctx)

	tflog.Info(ctx, "Done with deleting IAM Managed Policy resource")
}

// ImportState imports the existing resource into the Terraform state.
func (r *IAMManagedPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Expected format: <namespace>:<entity_type>:<entity_name>
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	parts := strings.Split(req.ID, ":")
	if len(parts) != 3 {
		resp.Diagnostics.AddError(
			"Invalid import ID format",
			"Expected format: <namespace>:<entity_type>:<entity_name>. Example: ns1:role:Role001",
		)
		return
	}

	namespace := parts[0]
	entityType := parts[1]
	entityName := parts[2]

	// Set namespace
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("namespace"), namespace)...)

	// Set entity type-specific attribute
	switch strings.ToLower(entityType) {
	case "user":
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("username"), entityName)...)
	case "group":
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("groupname"), entityName)...)
	case "role":
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("rolename"), entityName)...)
	default:
		resp.Diagnostics.AddError(
			"Invalid entity type",
			fmt.Sprintf("Entity type must be one of: user, group, role. Got: %s", entityType),
		)
		return
	}

	// ID will be set automatically by Terraform after import
}
