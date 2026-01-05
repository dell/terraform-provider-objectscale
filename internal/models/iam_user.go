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

type IAMUserResourceModel struct {
	Arn                     types.String `tfsdk:"arn"`
	CreateDate              types.String `tfsdk:"create_date"`
	Path                    types.String `tfsdk:"path"`
	PermissionsBoundaryArn  types.String `tfsdk:"permissions_boundary_arn"`
	PermissionsBoundaryType types.String `tfsdk:"permissions_boundary_type"`
	Tags                    types.Set   `tfsdk:"tags"`
	Id                      types.String `tfsdk:"id"`
	Name                    types.String `tfsdk:"name"`
	Namespace               types.String `tfsdk:"namespace"`
}

type Tags struct {
	// A single-valued attribute indicating the user's IDP domain
	Key types.String `tfsdk:"key"`
	// Attributes
	Value types.String `tfsdk:"value"`
}
type IAMUserDatasourceModel struct {
	ID        types.String `tfsdk:"id"`
	Namespace types.String `tfsdk:"namespace"`
	Username  types.String `tfsdk:"username"`
	Groupname types.String `tfsdk:"groupname"`
	Users     []IAMUser    `tfsdk:"users"`
}

type IAMUser struct {
	ID                  types.String       `tfsdk:"id"`
	UserName            types.String       `tfsdk:"username"`
	Arn                 types.String       `tfsdk:"arn"`
	Path                types.String       `tfsdk:"path"`
	CreateDate          types.String       `tfsdk:"create_date"`
	PermissionsBoundary types.String       `tfsdk:"permissions_boundary"` // can be empty
	Tags                []IAMUserTag       `tfsdk:"tags"`
	AccessKeys          []IAMUserAccessKey `tfsdk:"access_keys"`
}

type IAMUserTag struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

type IAMUserAccessKey struct {
	AccessKeyId types.String `tfsdk:"access_key_id"`
	CreateDate  types.String `tfsdk:"create_date"`
	Status      types.String `tfsdk:"status"`
}
