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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
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
			},
			"password": schema.StringAttribute{
				Description:         "Password for the management user. Required **only** when creating LOCAL_USER; ignored for AD/LDAP users and groups.",
				MarkdownDescription: "Password for the management user. Required **only** when creating LOCAL_USER; ignored for AD/LDAP users and groups.",
				Optional:            true,
				Sensitive:           true,
				PlanModifiers: []planmodifier.String{
					// Keep password from being diffed/unknown in plans unless explicitly changed.
					stringplanmodifier.UseStateForUnknown(),
				},
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

}

// Create creates the resource and sets the updated Terraform state on success.
func (r *ManagementUserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

}

// Update updates the resource and sets the updated Terraform state on success.
func (r *ManagementUserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

}

// Delete deletes the resource and removes the Terraform state.
func (r *ManagementUserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}

// ImportState imports the existing resource into the Terraform state.
func (r *ManagementUserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

}
