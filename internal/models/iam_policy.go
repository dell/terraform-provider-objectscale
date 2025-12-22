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

// IamPolicyDataSourceModel represents the schema for the IamPolicy data source.
type IamPolicyDataSourceModel struct {
	ID          types.String                        `tfsdk:"id"`
	ARN         types.String                        `tfsdk:"arn"`
	Namespace   types.String                        `tfsdk:"namespace"`
	User        types.String                        `tfsdk:"user"`
	Group       types.String                        `tfsdk:"group"`
	Role        types.String                        `tfsdk:"role"`
	IamPolicies []IamPolicyDataSourceIamPolicyModel `tfsdk:"policies"`
}

// IamPolicyDataSourceIamPolicyModel represents the schema for the iam_policies attribute.
type IamPolicyDataSourceIamPolicyModel struct {
	ARN                           types.String                               `tfsdk:"arn"`
	AttachmentCount               types.Int32                                `tfsdk:"attachment_count"`
	CreateDate                    types.String                               `tfsdk:"create_date"`
	DefaultVersionID              types.String                               `tfsdk:"default_version_id"`
	Description                   types.String                               `tfsdk:"description"`
	IsAttachable                  types.Bool                                 `tfsdk:"is_attachable"`
	Path                          types.String                               `tfsdk:"path"`
	PermissionsBoundaryUsageCount types.Int32                                `tfsdk:"permissions_boundary_usage_count"`
	PolicyID                      types.String                               `tfsdk:"policy_id"`
	PolicyName                    types.String                               `tfsdk:"policy_name"`
	UpdateDate                    types.String                               `tfsdk:"update_date"`
	Versions                      []IamPolicyDataSourceIamPolicyVersionModel `tfsdk:"versions"`
}

// IamPolicyDataSourceIamPolicyVersionModel represents the schema for the versions attribute.
type IamPolicyDataSourceIamPolicyVersionModel struct {
	IsDefaultVersion types.Bool   `tfsdk:"is_default_version"`
	VersionID        types.String `tfsdk:"version_id"`
	CreateDate       types.String `tfsdk:"create_date"`
	Document         types.String `tfsdk:"document"`
}
