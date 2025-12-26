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

package models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ManagementUserResourceModel maps the Management User resource data.
type ManagementUserResourceModel struct {
	ID                    types.String `tfsdk:"id"`
	Type                  types.String `tfsdk:"type"`
	Name                  types.String `tfsdk:"name"`
	Password              types.String `tfsdk:"password"`
	SystemAdministrator   types.Bool   `tfsdk:"system_administrator"`
	SystemMonitor         types.Bool   `tfsdk:"system_monitor"`
	SecurityAdministrator types.Bool   `tfsdk:"security_administrator"`
}

// ManagementUserDataSourceModel maps the Management User data source data.
type ManagementUserDataSourceModel struct {
	ID              types.String         `tfsdk:"id"`
	Name            types.String         `tfsdk:"name"`
	ManagementUsers []ManagementUserInfo `tfsdk:"management_users"`
}

// ManagementUserInfo represents a single Management User in the data source results.
type ManagementUserInfo struct {
	UserId                  types.String `tfsdk:"user_id"`
	IsSystemAdmin           types.Bool   `tfsdk:"is_system_admin"`
	IsSystemMonitor         types.Bool   `tfsdk:"is_system_monitor"`
	IsSecurityAdmin         types.Bool   `tfsdk:"is_security_admin"`
	IsExternalGroup         types.Bool   `tfsdk:"is_external_group"`
	IsLocked                types.Bool   `tfsdk:"is_locked"`
	LastTimePasswordChanged types.String `tfsdk:"last_time_password_changed"`
}
