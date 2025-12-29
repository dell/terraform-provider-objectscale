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
	"terraform-provider-objectscale/internal/helper"
	"terraform-provider-objectscale/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var _ datasource.DataSource = &ManagementUserDataSource{}

// NewManagementUserDataSource is a helper function to simplify the provider implementation.
func NewManagementUserDataSource() datasource.DataSource {
	return &ManagementUserDataSource{}
}

// ManagementUserDataSource is the data source implementation.
type ManagementUserDataSource struct {
	datasourceProviderConfig
}

// Metadata returns the data source type name.
func (d *ManagementUserDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_management_user"
}

// Schema defines the schema for the data source.
func (d *ManagementUserDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This datasource can be used to fetch details of Management Users from ObjectScale.",
		MarkdownDescription: "This datasource can be used to fetch details of Management Users from ObjectScale.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Identifier",
				MarkdownDescription: "Identifier",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				Description:         "Management user name.",
				MarkdownDescription: "Management user name.",
				Optional:            true,
			},
			"management_users": schema.ListNestedAttribute{
				Description:         "List of management user information.",
				MarkdownDescription: "List of management user information.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"user_id": schema.StringAttribute{
							Description:         "User Id.",
							MarkdownDescription: "User Id.",
							Computed:            true,
						},
						"is_system_admin": schema.BoolAttribute{
							Description:         "Flag indicating whether management user is System Admin.",
							MarkdownDescription: "Flag indicating whether management user is System Admin.",
							Computed:            true,
						},
						"is_system_monitor": schema.BoolAttribute{
							Description:         "Flag indicating whether management user is System Monitor.",
							MarkdownDescription: "Flag indicating whether management user is System Monitor.",
							Computed:            true,
						},
						"is_security_admin": schema.BoolAttribute{
							Description:         "Flag indicating whether management user is Security Admin.",
							MarkdownDescription: "Flag indicating whether management user is Security Admin.",
							Computed:            true,
						},
						"is_external_group": schema.BoolAttribute{
							Description:         "If set to true, its a domain.",
							MarkdownDescription: "If set to true, its a domain.",
							Computed:            true,
						},
						"is_locked": schema.BoolAttribute{
							Description:         "If set to true, the user is locked.",
							MarkdownDescription: "If set to true, the user is locked.",
							Computed:            true,
						},
						"last_time_password_changed": schema.StringAttribute{
							Description:         "Value of last time password changed.",
							MarkdownDescription: "Value of last time password changed.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *ManagementUserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state models.ManagementUserDataSourceModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	userID := ""
	if !state.Name.IsNull() {
		userID = state.Name.ValueString()
	}

	var managementUsers []models.ManagementUserInfo
	if userID != "" {
		getResp, _, err := d.client.GenClient.MgmtUserInfoApi.MgmtUserInfoServiceGetLocalUserInfo(ctx, userID).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Get Management User failed", err.Error())
			return
		}
		managementUsers = append(managementUsers, models.ManagementUserInfo{
			UserId:                  helper.TfString(getResp.UserId),
			IsSystemAdmin:           helper.TfBool(getResp.IsSystemAdmin),
			IsSystemMonitor:         helper.TfBool(getResp.IsSystemMonitor),
			IsSecurityAdmin:         helper.TfBool(getResp.IsSecurityAdmin),
			IsExternalGroup:         helper.TfBool(getResp.IsExternalGroup),
			IsLocked:                helper.TfBool(getResp.IsLocked),
			LastTimePasswordChanged: helper.TfString(getResp.LastTimePasswordChanged),
		})
	} else {
		listResp, _, err := d.client.GenClient.MgmtUserInfoApi.MgmtUserInfoServiceGetLocalUserInfos(ctx).Execute()
		if err != nil {
			resp.Diagnostics.AddError("List Management Users failed", err.Error())
			return
		}
		allManagementUsers := listResp.MgmtUserInfo
		for _, mgmtUser := range allManagementUsers {
			managementUsers = append(managementUsers, models.ManagementUserInfo{
				UserId:                  helper.TfString(mgmtUser.UserId),
				IsSystemAdmin:           helper.TfBool(mgmtUser.IsSystemAdmin),
				IsSystemMonitor:         helper.TfBool(mgmtUser.IsSystemMonitor),
				IsSecurityAdmin:         helper.TfBool(mgmtUser.IsSecurityAdmin),
				IsExternalGroup:         helper.TfBool(mgmtUser.IsExternalGroup),
				IsLocked:                helper.TfBool(mgmtUser.IsLocked),
				LastTimePasswordChanged: helper.TfString(mgmtUser.LastTimePasswordChanged),
			})
		}
	}

	// Set state
	state.ID = types.StringValue("management_user_datasource")
	state.ManagementUsers = managementUsers
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
