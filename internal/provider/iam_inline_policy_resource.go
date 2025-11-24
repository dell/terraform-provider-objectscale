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
var _ resource.Resource = &IAMInlinePolicyResource{}
var _ resource.ResourceWithImportState = &IAMInlinePolicyResource{}

// IAMInlinePolicyResource is the resource implementation.
type IAMInlinePolicyResource struct {
	client *client.Client
}

// NewIAMInlinePolicyResource is a helper function to simplify the provider implementation.
func NewIAMInlinePolicyResource() resource.Resource {
	return &IAMInlinePolicyResource{}
}

// Metadata returns the resource type name.
func (r *IAMInlinePolicyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_inline_policy"
}

// Configure adds the provider configured client to the resource.
func (r *IAMInlinePolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *IAMInlinePolicyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "Manages IAM inline policies for ObjectScale entities (user, group, or role).",
		MarkdownDescription: "Manages IAM inline policies for ObjectScale entities (user, group, or role).",
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
						},
					},
				},
			},
		},
	}
}

// ValidateConfig validates the resource configuration.
func (r *IAMInlinePolicyResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var config models.IAMInlinePolicyResourceModel

	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	count := 0
	if !config.Username.IsNull() && !config.Username.IsUnknown() {
		count++
	}
	if !config.Groupname.IsNull() && !config.Groupname.IsUnknown() {
		count++
	}
	if !config.Rolename.IsNull() && !config.Rolename.IsUnknown() {
		count++
	}

	if count == 0 {
		resp.Diagnostics.AddError(
			"Validation Error",
			"Exactly one of username, groupname, or rolename must be provided.",
		)
	} else if count > 1 {
		resp.Diagnostics.AddError(
			"Validation Error",
			"Only one of username, groupname, or rolename can be provided.",
		)
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
	switch entityType {
	case "User":
		listResp, _, err := r.client.GenClient.IamApi.IamServiceListUserPolicies(ctx).
			XEmcNamespace(namespace).
			UserName(entityName).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Read Error", fmt.Sprintf("Failed to list policies: %s", err))
			return
		}
		policyNames = listResp.ListUserPoliciesResult.PolicyNames

	case "Group":
		listResp, _, err := r.client.GenClient.IamApi.IamServiceListGroupPolicies(ctx).
			XEmcNamespace(namespace).
			GroupName(entityName).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Read Error", fmt.Sprintf("Failed to list policies: %s", err))
			return
		}
		policyNames = listResp.ListGroupPoliciesResult.PolicyNames

	case "Role":
		listResp, _, err := r.client.GenClient.IamApi.IamServiceListRolePolicies(ctx).
			XEmcNamespace(namespace).
			RoleName(entityName).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Read Error", fmt.Sprintf("Failed to list policies: %s", err))
			return
		}
		policyNames = listResp.ListRolePoliciesResult.PolicyNames
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
				resp.Diagnostics.AddError("Read Error", fmt.Sprintf("Failed to get policy %s: %s", policyName, err))
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
				resp.Diagnostics.AddError("Read Error", fmt.Sprintf("Failed to get policy %s: %s", policyName, err))
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
				resp.Diagnostics.AddError("Read Error", fmt.Sprintf("Failed to get policy %s: %s", policyName, err))
				return
			}
			policyDoc = *getResp.GetRolePolicyResult.PolicyDocument
		}

		policies = append(policies, models.IAMInlinePolicyModel{
			Name:     types.StringValue(policyName),
			Document: types.StringValue(policyDoc),
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

	updatedModel, err := r.applyPolicies(ctx, plan)
	if err != nil {
		resp.Diagnostics.AddError("Apply Error", err.Error())
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

	updatedModel, err := r.applyPolicies(ctx, plan)
	if err != nil {
		resp.Diagnostics.AddError("Apply Error", err.Error())
		return
	}

	diags = resp.State.Set(ctx, updatedModel)
	resp.Diagnostics.Append(diags...)
}

func (r *IAMInlinePolicyResource) applyPolicies(ctx context.Context, plan models.IAMInlinePolicyResourceModel) (models.IAMInlinePolicyResourceModel, error) {
	// Determine namespace
	namespace := plan.Namespace.ValueString()

	// Determine entity type and name
	var entityType, entityName string
	if !plan.Username.IsNull() && !plan.Username.IsUnknown() {
		entityType = "User"
		entityName = plan.Username.ValueString()
	} else if !plan.Groupname.IsNull() && !plan.Groupname.IsUnknown() {
		entityType = "Group"
		entityName = plan.Groupname.ValueString()
	} else if !plan.Rolename.IsNull() && !plan.Rolename.IsUnknown() {
		entityType = "Role"
		entityName = plan.Rolename.ValueString()
	}

	// Step 1: Get current policies from ObjectScale
	var currentPolicies []string
	switch entityType {
	case "User":
		listResp, _, err := r.client.GenClient.IamApi.IamServiceListUserPolicies(ctx).
			XEmcNamespace(namespace).
			UserName(entityName).
			Execute()
		if err != nil {
			return plan, fmt.Errorf("failed to list policies: %w", err)
		}
		currentPolicies = listResp.ListUserPoliciesResult.PolicyNames

	case "Group":
		listResp, _, err := r.client.GenClient.IamApi.IamServiceListGroupPolicies(ctx).
			XEmcNamespace(namespace).
			GroupName(entityName).
			Execute()
		if err != nil {
			return plan, fmt.Errorf("failed to list policies: %w", err)
		}
		currentPolicies = listResp.ListGroupPoliciesResult.PolicyNames

	case "Role":
		listResp, _, err := r.client.GenClient.IamApi.IamServiceListRolePolicies(ctx).
			XEmcNamespace(namespace).
			RoleName(entityName).
			Execute()
		if err != nil {
			return plan, fmt.Errorf("failed to list policies: %w", err)
		}
		currentPolicies = listResp.ListRolePoliciesResult.PolicyNames
	}

	// Convert desired policies to map for quick lookup
	desiredMap := make(map[string]string)
	for _, p := range plan.Policies {
		desiredMap[p.Name.ValueString()] = p.Document.ValueString()
	}

	// Step 2: Delete policies not in desired config
	for _, existing := range currentPolicies {
		if _, found := desiredMap[existing]; !found {
			switch entityType {
			case "User":
				_, _, err := r.client.GenClient.IamApi.IamServiceDeleteUserPolicy(ctx).
					XEmcNamespace(namespace).
					UserName(entityName).
					PolicyName(existing).
					Execute()
				if err != nil {
					return plan, fmt.Errorf("failed to delete policy %s: %w", existing, err)
				}

			case "Group":
				_, _, err := r.client.GenClient.IamApi.IamServiceDeleteGroupPolicy(ctx).
					XEmcNamespace(namespace).
					GroupName(entityName).
					PolicyName(existing).
					Execute()
				if err != nil {
					return plan, fmt.Errorf("failed to delete policy %s: %w", existing, err)
				}

			case "Role":
				_, _, err := r.client.GenClient.IamApi.IamServiceDeleteRolePolicy(ctx).
					XEmcNamespace(namespace).
					RoleName(entityName).
					PolicyName(existing).
					Execute()
				if err != nil {
					return plan, fmt.Errorf("failed to delete policy %s: %w", existing, err)
				}
			}
		}
	}

	// Step 3: Create or Update desired policies
	for name, doc := range desiredMap {
		switch entityType {
		case "User":
			_, _, err := r.client.GenClient.IamApi.IamServicePutUserPolicy(ctx).
				XEmcNamespace(namespace).
				UserName(entityName).
				PolicyName(name).
				PolicyDocument(doc).
				Execute()
			if err != nil {
				return plan, fmt.Errorf("failed to apply policy %s: %w", name, err)
			}

		case "Group":
			_, _, err := r.client.GenClient.IamApi.IamServicePutGroupPolicy(ctx).
				XEmcNamespace(namespace).
				GroupName(entityName).
				PolicyName(name).
				PolicyDocument(doc).
				Execute()
			if err != nil {
				return plan, fmt.Errorf("failed to apply policy %s: %w", name, err)
			}

		case "Role":
			_, _, err := r.client.GenClient.IamApi.IamServicePutRolePolicy(ctx).
				XEmcNamespace(namespace).
				RoleName(entityName).
				PolicyName(name).
				PolicyDocument(doc).
				Execute()
			if err != nil {
				return plan, fmt.Errorf("failed to apply policy %s: %w", name, err)
			}
		}
	}

	// Set ID - format: <namespace>:<entity_type>:<entity_name>
	plan.ID = types.StringValue(fmt.Sprintf("%s:%s:%s", namespace, strings.ToLower(entityType), entityName))

	return plan, nil
}

// Delete deletes the resource and removes the Terraform state.
func (r *IAMInlinePolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Info(ctx, "Deleting IAM Inline Policy resource (no API call, just removing state)")

	var state models.IAMInlinePolicyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Remove resource from Terraform state
	resp.State.RemoveResource(ctx)

	tflog.Info(ctx, "Done with Deleting IAM Inline Policy resource")
}

// ImportState imports the existing resource into the Terraform state.
func (r *IAMInlinePolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Expected format: <namespace>:<entity_type>:<entity_name>
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
