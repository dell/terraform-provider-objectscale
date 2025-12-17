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
	"regexp"
	"strings"
	"terraform-provider-objectscale/internal/clientgen"
	"terraform-provider-objectscale/internal/helper"
	"terraform-provider-objectscale/internal/models"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Constants for valid type values (used in Schema validators)
const (
	ManagementUserTypeLocal       = "LOCAL_USER"
	ManagementUserTypeADLDAPUser  = "AD_LDAP_USER"
	ManagementUserTypeADLDAPGroup = "AD_LDAP_GROUP"
)

// Ensure the implementation satisfies the expected interfaces.
var _ resource.Resource = &ManagementUserResource{}
var _ resource.ResourceWithImportState = &ManagementUserResource{}

// ManagementUserResource is the resource implementation.
type ManagementUserResource struct {
	resourceProviderConfig
}

// NewManagementUserResource is a helper function to simplify the provider implementation.
func NewManagementUserResource() resource.Resource {
	return &ManagementUserResource{}
}

// Metadata returns the resource type name.
func (r *ManagementUserResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_management_user"
}

// Schema defines the schema for the resource.
func (r *ManagementUserResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "Manages ObjectScale management users. Supported types: LOCAL_USER, AD_LDAP_USER, AD_LDAP_GROUP.",
		MarkdownDescription: "Manages ObjectScale management users. Supported types: LOCAL_USER, AD_LDAP_USER, AD_LDAP_GROUP.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Unique identifier for the management user.",
				MarkdownDescription: "Unique identifier for the management user.",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				Description:         "Type of management user. Allowed values: LOCAL_USER, AD_LDAP_USER, AD_LDAP_GROUP.",
				MarkdownDescription: "Type of management user. Allowed values: LOCAL_USER, AD_LDAP_USER, AD_LDAP_GROUP.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						ManagementUserTypeLocal,
						ManagementUserTypeADLDAPUser,
						ManagementUserTypeADLDAPGroup,
					),
				},
			},
			"name": schema.StringAttribute{
				Description:         `Management user id. Format is as follows: For LOCAL_USER use "user1". For AD/LDAP User/Group use "user1@domain".`,
				MarkdownDescription: `Management user id. Format is as follows: For LOCAL_USER use "user1". For AD/LDAP User/Group use "user1@domain".`,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[^A-Z]+$`),
						"Uppercase letters are not allowed in 'name'.",
					),
				},
			},
			"password": schema.StringAttribute{
				Description:         "Password for the management user. Required **only** when creating LOCAL_USER; ignored for AD/LDAP users and groups.",
				MarkdownDescription: "Password for the management user. Required **only** when creating LOCAL_USER; ignored for AD/LDAP users and groups.",
				Optional:            true,
				Sensitive:           true,
			},
			"system_administrator": schema.BoolAttribute{
				Description:         "If set to true, assigns the management user to the System Admin role. System Administrators perform system level administration (VDC administration) and namespace administration.",
				MarkdownDescription: "If set to true, assigns the management user to the System Admin role. System Administrators perform system level administration (VDC administration) and namespace administration.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"system_monitor": schema.BoolAttribute{
				Description:         "If set to true, assigns the management user to the System Monitor role. System Monitors have read-only access to the ObjectScale Portal.",
				MarkdownDescription: "If set to true, assigns the management user to the System Monitor role. System Monitors have read-only access to the ObjectScale Portal.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"security_administrator": schema.BoolAttribute{
				Description:         "If set to true, assigns the management user to the Security Admin role. Security Administrators perform user management and security related administration.",
				MarkdownDescription: "If set to true, assigns the management user to the Security Admin role. Security Administrators perform user management and security related administration.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *ManagementUserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.ManagementUserResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	userID := state.Name.ValueString()
	prevPassword := state.Password

	// get management user
	getResp, _, err := r.client.GenClient.MgmtUserInfoApi.MgmtUserInfoServiceGetLocalUserInfo(ctx, userID).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Read Management User failed", err.Error())
		return
	}

	newState := mapToModel(getResp, prevPassword)
	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
}

// Create creates the resource and sets the Terraform state on success.
func (r *ManagementUserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.ManagementUserResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	userID := plan.Name.ValueString()
	mgmtUserType := plan.Type.ValueString()

	// Validate name format based on type
	if mgmtUserType == ManagementUserTypeLocal && strings.Contains(userID, "@") {
		resp.Diagnostics.AddError(
			"Invalid Name Format for LOCAL_USER",
			"For type LOCAL_USER, 'name' must not contain '@'. Please provide a valid name format.",
		)
		return
	}
	if (mgmtUserType == ManagementUserTypeADLDAPUser || mgmtUserType == ManagementUserTypeADLDAPGroup) && !strings.Contains(userID, "@") {
		resp.Diagnostics.AddError(
			"Invalid Name Format for AD_LDAP_USER/AD_LDAP_GROUP",
			"For type AD_LDAP_USER or AD_LDAP_GROUP, 'name' must contain '@'. Please provide a valid name format.",
		)
		return
	}

	// Validate conditional password for LOCAL_USER
	if mgmtUserType == ManagementUserTypeLocal && !isNonEmptyString(plan.Password) {
		resp.Diagnostics.AddError(
			"Password is required for LOCAL_USER",
			"For type LOCAL_USER, 'password' must be provided during creation.",
		)
		return
	}

	// build create request payload
	createRequest := clientgen.MgmtUserInfoServiceCreateLocalUserInfoRequest{
		UserId:          userID,
		IsSystemAdmin:   helper.ValueToPointer[bool](plan.SystemAdministrator),
		IsSystemMonitor: helper.ValueToPointer[bool](plan.SystemMonitor),
		IsSecurityAdmin: helper.ValueToPointer[bool](plan.SecurityAdministrator),
	}

	switch mgmtUserType {
	case ManagementUserTypeLocal:
		createRequest.IsExternalGroup = helper.ValueToPointer[bool](types.BoolValue(false))
		createRequest.Password = helper.ValueToPointer[string](plan.Password)
	case ManagementUserTypeADLDAPUser:
		createRequest.IsExternalGroup = helper.ValueToPointer[bool](types.BoolValue(false))
	case ManagementUserTypeADLDAPGroup:
		createRequest.IsExternalGroup = helper.ValueToPointer[bool](types.BoolValue(true))
	}

	// create management user
	_, _, err := r.client.GenClient.MgmtUserInfoApi.MgmtUserInfoServiceCreateLocalUserInfo(ctx).MgmtUserInfoServiceCreateLocalUserInfoRequest(createRequest).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Create Management User failed", err.Error())
		return
	}

	// get management user
	getResp, _, err := r.client.GenClient.MgmtUserInfoApi.MgmtUserInfoServiceGetLocalUserInfo(ctx, userID).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Read Management User failed", err.Error())
		return
	}

	newState := mapToModel(getResp, plan.Password)
	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *ManagementUserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan models.ManagementUserResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state models.ManagementUserResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// To prevent the non-updatable fields from being changed
	if !plan.Type.Equal(state.Type) || !plan.Name.Equal(state.Name) {
		resp.Diagnostics.AddError("Error updating management user", "The attributes `type` and `name` are not updatable")
		return
	}

	userID := plan.Name.ValueString()
	mgmtUserType := plan.Type.ValueString()

	// build update request payload
	updateRequest := clientgen.MgmtUserInfoServiceModifyLocalUserInfoRequest{
		IsSystemAdmin:   helper.ValueToPointer[bool](plan.SystemAdministrator),
		IsSystemMonitor: helper.ValueToPointer[bool](plan.SystemMonitor),
		IsSecurityAdmin: helper.ValueToPointer[bool](plan.SecurityAdministrator),
	}

	// Only LOCAL_USER can change password; include only if explicitly provided
	if mgmtUserType == ManagementUserTypeLocal && isNonEmptyString(plan.Password) {
		// Optional: only send if changed vs state
		if state.Password.IsNull() || state.Password.IsUnknown() || plan.Password.ValueString() != state.Password.ValueString() {
			updateRequest.Password = helper.ValueToPointer[string](plan.Password)
		}
	}

	// update management user
	_, _, err := r.client.GenClient.MgmtUserInfoApi.MgmtUserInfoServiceModifyLocalUserInfo(ctx, userID).MgmtUserInfoServiceModifyLocalUserInfoRequest(updateRequest).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Update Management User failed", err.Error())
		return
	}

	// get management user
	getResp, _, err := r.client.GenClient.MgmtUserInfoApi.MgmtUserInfoServiceGetLocalUserInfo(ctx, userID).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Read Management User failed", err.Error())
		return
	}

	newState := mapToModel(getResp, plan.Password)
	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
}

// Delete deletes the resource and removes the Terraform state.
func (r *ManagementUserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Info(ctx, "Deleting Management User resource")

	var state models.ManagementUserResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	userID := state.Name.ValueString()

	// delete management user
	_, _, err := r.client.GenClient.MgmtUserInfoApi.MgmtUserInfoServiceDeleteLocalUserInfo(ctx, userID).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Delete Management User failed", err.Error())
		return
	}

	// Remove resource from Terraform state
	resp.State.RemoveResource(ctx)
	tflog.Info(ctx, "Done with deleting Management User resource")
}

// ImportState imports the existing resource into the Terraform state.
func (r *ManagementUserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	userID := req.ID
	prevPassword := types.StringNull()

	// get management user
	getResp, _, err := r.client.GenClient.MgmtUserInfoApi.MgmtUserInfoServiceGetLocalUserInfo(ctx, userID).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Import Management User failed", err.Error())
		return
	}

	newState := mapToModel(getResp, prevPassword)
	diags := resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
}

func mapToModel(resp *clientgen.MgmtUserInfoServiceGetLocalUserInfoResponse, password types.String) models.ManagementUserResourceModel {
	mgmtUserType := deriveTypeFromAPI(resp)
	return models.ManagementUserResourceModel{
		ID:                    helper.TfString(resp.UserId),
		Type:                  types.StringValue(mgmtUserType),
		Name:                  helper.TfString(resp.UserId),
		Password:              password,
		SystemAdministrator:   helper.TfBool(resp.IsSystemAdmin),
		SystemMonitor:         helper.TfBool(resp.IsSystemMonitor),
		SecurityAdministrator: helper.TfBool(resp.IsSecurityAdmin),
	}
}

func deriveTypeFromAPI(resp *clientgen.MgmtUserInfoServiceGetLocalUserInfoResponse) string {
	if *resp.IsExternalGroup {
		return ManagementUserTypeADLDAPGroup
	}
	// If not a group, infer by name format.
	if strings.Contains(*resp.UserId, "@") {
		return ManagementUserTypeADLDAPUser
	}
	return ManagementUserTypeLocal
}

func isNonEmptyString(v types.String) bool {
	return !v.IsNull() && !v.IsUnknown() && strings.TrimSpace(v.ValueString()) != ""
}
