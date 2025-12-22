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

import "github.com/hashicorp/terraform-plugin-framework/types"

type IAMRoleDatasourceModel struct {
	ID        types.String `tfsdk:"id"`
	Namespace types.String `tfsdk:"namespace"`
	RoleName  types.String `tfsdk:"role_name"`
	Roles     []IAMRole    `tfsdk:"roles"`
}

type IAMRole struct {
	RoleId              types.String                `tfsdk:"role_id"`
	RoleName            types.String                `tfsdk:"role_name"`
	Arn                 types.String                `tfsdk:"arn"`
	AssumeRolePolicy    types.String                `tfsdk:"assume_role_policy"`
	Path                types.String                `tfsdk:"path"`
	Description         types.String                `tfsdk:"description"`
	CreateDate          types.String                `tfsdk:"create_date"`
	MaxSessionDuration  types.Int64                 `tfsdk:"max_session_duration"`
	PermissionsBoundary *IAMRolePermissionsBoundary `tfsdk:"permissions_boundary"`
	Tags                []IAMRoleTag                `tfsdk:"tags"`
}

type IAMRoleTag struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

type IAMRolePermissionsBoundary struct {
	PermissionsBoundaryArn  types.String `tfsdk:"permissions_boundary_arn"`
	PermissionsBoundaryType types.String `tfsdk:"permissions_boundary_type"`
}

type IAMRoleResourceModel struct {
	Name                     types.String `tfsdk:"name"`
	Namespace                types.String `tfsdk:"namespace"`
	AssumeRolePolicyDocument types.String `tfsdk:"assume_role_policy_document"`
	Description              types.String `tfsdk:"description"`
	MaxSessionDuration       types.Int32  `tfsdk:"max_session_duration"`
	Path                     types.String `tfsdk:"path"`
	PermissionsBoundaryArn   types.String `tfsdk:"permissions_boundary_arn"`
	PermissionsBoundaryType  types.String `tfsdk:"permissions_boundary_type"`
	Tags                     types.List   `tfsdk:"tags"`
}
