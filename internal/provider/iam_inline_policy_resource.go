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
	"terraform-provider-objectscale/internal/helper"
	"terraform-provider-objectscale/internal/models"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var _ resource.Resource = &IAMInlinePolicyResource{}
var _ resource.ResourceWithImportState = &IAMInlinePolicyResource{}

// IAMInlinePolicyResource is the resource implementation.
type IAMInlinePolicyResource struct {
	resourceProviderConfig
}

// NewIAMInlinePolicyResource is a helper function to simplify the provider implementation.
func NewIAMInlinePolicyResource() resource.Resource {
	return &IAMInlinePolicyResource{}
}

// Metadata returns the resource type name.
func (r *IAMInlinePolicyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_inline_policy"
}

// Schema defines the schema for the resource.
func (r *IAMInlinePolicyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This resource manages IAM inline policies for Dell ObjectScale entities (user, group, or role).",
		MarkdownDescription: "This resource manages IAM inline policies for Dell ObjectScale entities (user, group, or role).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Unique identifier for the IAM inline policy resource.",
				MarkdownDescription: "Unique identifier for the IAM inline policy resource.",
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
			"policies": schema.ListNestedAttribute{
				Description:         "List of IAM inline policies to associate with the entity.",
				MarkdownDescription: "List of IAM inline policies to associate with the entity.",
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description:         "Name of the IAM inline policy.",
							MarkdownDescription: "Name of the IAM inline policy.",
							Required:            true,
						},
						"document": schema.StringAttribute{
							Description:         "Policy document in JSON format.",
							MarkdownDescription: "Policy document in JSON format.",
							Required:            true,
							CustomType:          jsontypes.NormalizedType{},
						},
					},
				},
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *IAMInlinePolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.IAMInlinePolicyResourceModel

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

	// Step 1: Call List<entity>Policies API
	var policyNames []string
	var marker string

	switch entityType {
	case "User":
		for {
			listReq := r.client.GenClient.IamApi.IamServiceListUserPolicies(ctx).
				XEmcNamespace(namespace).
				UserName(entityName)

			if marker != "" {
				listReq = listReq.Marker(marker)
			}

			listResp, _, err := listReq.Execute()
			if err != nil {
				resp.Diagnostics.AddError("Read Error", fmt.Sprintf("Failed to list policies: %s", err.Error()))
				return
			}

			policyNames = append(policyNames, listResp.ListUserPoliciesResult.PolicyNames...)

			markerPtr := listResp.ListUserPoliciesResult.Marker
			if markerPtr == nil || *markerPtr == "" {
				break
			}
			marker = *markerPtr
		}
	case "Group":
		for {
			listReq := r.client.GenClient.IamApi.IamServiceListGroupPolicies(ctx).
				XEmcNamespace(namespace).
				GroupName(entityName)

			if marker != "" {
				listReq = listReq.Marker(marker)
			}

			listResp, _, err := listReq.Execute()
			if err != nil {
				resp.Diagnostics.AddError("Read Error", fmt.Sprintf("Failed to list policies: %s", err.Error()))
				return
			}

			policyNames = append(policyNames, listResp.ListGroupPoliciesResult.PolicyNames...)

			markerPtr := listResp.ListGroupPoliciesResult.Marker
			if markerPtr == nil || *markerPtr == "" {
				break
			}
			marker = *markerPtr
		}
	case "Role":
		for {
			listReq := r.client.GenClient.IamApi.IamServiceListRolePolicies(ctx).
				XEmcNamespace(namespace).
				RoleName(entityName)

			if marker != "" {
				listReq = listReq.Marker(marker)
			}

			listResp, _, err := listReq.Execute()
			if err != nil {
				resp.Diagnostics.AddError("Read Error", fmt.Sprintf("Failed to list policies: %s", err.Error()))
				return
			}

			policyNames = append(policyNames, listResp.ListRolePoliciesResult.PolicyNames...)

			markerPtr := listResp.ListRolePoliciesResult.Marker
			if markerPtr == nil || *markerPtr == "" {
				break
			}
			marker = *markerPtr
		}
	}

	// Step 2: For each policy name, call Get<entity>Policy API
	var policies []models.IAMInlinePolicyModel
	for _, policyName := range policyNames {
		var policyDoc string
		switch entityType {
		case "User":
			getResp, _, err := r.client.GenClient.IamApi.IamServiceGetUserPolicy(ctx).
				XEmcNamespace(namespace).
				UserName(entityName).
				PolicyName(policyName).
				Execute()
			if err != nil {
				resp.Diagnostics.AddError("Read Error", fmt.Sprintf("Failed to get policy %s: %s", policyName, err.Error()))
				return
			}
			policyDoc = *getResp.GetUserPolicyResult.PolicyDocument

		case "Group":
			getResp, _, err := r.client.GenClient.IamApi.IamServiceGetGroupPolicy(ctx).
				XEmcNamespace(namespace).
				GroupName(entityName).
				PolicyName(policyName).
				Execute()
			if err != nil {
				resp.Diagnostics.AddError("Read Error", fmt.Sprintf("Failed to get policy %s: %s", policyName, err.Error()))
				return
			}
			policyDoc = *getResp.GetGroupPolicyResult.PolicyDocument

		case "Role":
			getResp, _, err := r.client.GenClient.IamApi.IamServiceGetRolePolicy(ctx).
				XEmcNamespace(namespace).
				RoleName(entityName).
				PolicyName(policyName).
				Execute()
			if err != nil {
				resp.Diagnostics.AddError("Read Error", fmt.Sprintf("Failed to get policy %s: %s", policyName, err.Error()))
				return
			}
			policyDoc = *getResp.GetRolePolicyResult.PolicyDocument
		}

		policies = append(policies, models.IAMInlinePolicyModel{
			Name:     types.StringValue(policyName),
			Document: jsontypes.NewNormalizedValue(policyDoc),
		})
	}
	if policies == nil {
		policies = []models.IAMInlinePolicyModel{}
	}

	// Update state
	state.Policies = policies
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Create creates the resource and sets the updated Terraform state on success.
func (r *IAMInlinePolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.IAMInlinePolicyResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	updatedModel, err := helper.ApplyPolicies(r.client, ctx, plan, nil)
	if err != nil {
		resp.Diagnostics.AddError("Create Error", err.Error())
		return
	}

	diags = resp.State.Set(ctx, updatedModel)
	resp.Diagnostics.Append(diags...)
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *IAMInlinePolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan models.IAMInlinePolicyResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state models.IAMInlinePolicyResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Determine entity type and name from plan
	var entityTypeFromPlan, entityNameFromPlan string
	if !plan.Username.IsNull() && !plan.Username.IsUnknown() {
		entityTypeFromPlan = "User"
		entityNameFromPlan = plan.Username.ValueString()
	} else if !plan.Groupname.IsNull() && !plan.Groupname.IsUnknown() {
		entityTypeFromPlan = "Group"
		entityNameFromPlan = plan.Groupname.ValueString()
	} else if !plan.Rolename.IsNull() && !plan.Rolename.IsUnknown() {
		entityTypeFromPlan = "Role"
		entityNameFromPlan = plan.Rolename.ValueString()
	}

	// Determine entity type and name from state
	var entityTypeFromState, entityNameFromState string
	if !state.Username.IsNull() && !state.Username.IsUnknown() {
		entityTypeFromState = "User"
		entityNameFromState = state.Username.ValueString()
	} else if !state.Groupname.IsNull() && !state.Groupname.IsUnknown() {
		entityTypeFromState = "Group"
		entityNameFromState = state.Groupname.ValueString()
	} else if !state.Rolename.IsNull() && !state.Rolename.IsUnknown() {
		entityTypeFromState = "Role"
		entityNameFromState = state.Rolename.ValueString()
	}

	if plan.Namespace.ValueString() != state.Namespace.ValueString() ||
		entityTypeFromPlan != entityTypeFromState ||
		entityNameFromPlan != entityNameFromState {
		resp.Diagnostics.AddError("Update Error", "The attributes `namespace`,`username`,`groupname`,`rolename` are not updatable")
		return
	}

	updatedModel, err := helper.ApplyPolicies(r.client, ctx, plan, &state)
	if err != nil {
		resp.Diagnostics.AddError("Update Error", err.Error())
		return
	}

	diags = resp.State.Set(ctx, updatedModel)
	resp.Diagnostics.Append(diags...)
}

// Delete deletes the resource and removes the Terraform state.
func (r *IAMInlinePolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Info(ctx, "Deleting IAM Inline Policy resource")

	var state models.IAMInlinePolicyResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.Policies = []models.IAMInlinePolicyModel{}
	_, err := helper.ApplyPolicies(r.client, ctx, state, nil)
	if err != nil {
		resp.Diagnostics.AddError("Delete Error", err.Error())
		return
	}

	// Remove resource from Terraform state
	resp.State.RemoveResource(ctx)

	tflog.Info(ctx, "Done with deleting IAM Inline Policy resource")
}

// ImportState imports the existing resource into the Terraform state.
func (r *IAMInlinePolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
