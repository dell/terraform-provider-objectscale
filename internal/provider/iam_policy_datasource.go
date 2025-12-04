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
	"terraform-provider-objectscale/internal/clientgen"
	"terraform-provider-objectscale/internal/helper"
	"terraform-provider-objectscale/internal/models"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &IAMPolicyDataSource{}

func NewIAMPolicyDataSource() datasource.DataSource {
	return &IAMPolicyDataSource{}
}

type IAMPolicyDataSource struct {
	datasourceProviderConfig
}

func (d *IAMPolicyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_policy"
}

// datasource item schema
func (d *IAMPolicyDataSource) itemSchema() schema.ListNestedAttribute {
	return schema.ListNestedAttribute{
		Description:         "List of IAM Policies.",
		MarkdownDescription: "List of IAM Policies.",
		Computed:            true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"arn": schema.StringAttribute{
					Description:         "The resource name of the policy.",
					MarkdownDescription: "The resource name of the policy.",
					Computed:            true,
				},
				"attachment_count": schema.Int32Attribute{
					Description:         "The number of entities (users, groups, and roles) that the policy is attached to.",
					MarkdownDescription: "The number of entities (users, groups, and roles) that the policy is attached to.",
					Computed:            true,
				},
				"create_date": schema.StringAttribute{
					Description:         "The date and time, in ISO 8601 date-time format, when the policy was created.",
					MarkdownDescription: "The date and time, in ISO 8601 date-time format, when the policy was created.",
					Computed:            true,
				},
				"default_version_id": schema.StringAttribute{
					Description:         "The identifier for the version of the policy that is set as the default version.",
					MarkdownDescription: "The identifier for the version of the policy that is set as the default version.",
					Computed:            true,
				},
				"description": schema.StringAttribute{
					Description:         "A friendly description of the policy.",
					MarkdownDescription: "A friendly description of the policy.",
					Computed:            true,
				},
				"is_attachable": schema.BoolAttribute{
					Description:         "Specifies whether the policy can be attached to user, group, or role.",
					MarkdownDescription: "Specifies whether the policy can be attached to user, group, or role.",
					Computed:            true,
				},
				"path": schema.StringAttribute{
					Description:         "The path to the policy",
					MarkdownDescription: "The path to the policy",
					Computed:            true,
				},
				"permissions_boundary_usage_count": schema.Int32Attribute{
					Description:         "Resource name of the policy that is used to set permissions boundary for the policy.",
					MarkdownDescription: "Resource name of the policy that is used to set permissions boundary for the policy.",
					Computed:            true,
				},
				"policy_id": schema.StringAttribute{
					Description:         "The stable and unique string identifying the policy.",
					MarkdownDescription: "The stable and unique string identifying the policy.",
					Computed:            true,
				},
				"policy_name": schema.StringAttribute{
					Description:         "The friendly name of the policy.",
					MarkdownDescription: "The friendly name of the policy.",
					Computed:            true,
				},
				"update_date": schema.StringAttribute{
					Description:         "The date and time, in ISO 8601 date-time format, when the policy was created.",
					MarkdownDescription: "The date and time, in ISO 8601 date-time format, when the policy was created.",
					Computed:            true,
				},
			},
		},
	}
}

// Schema describes the data source arguments.
func (d *IAMPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "ObjectScale supports access by multiple tenants, where each tenant is defined by a iam_policy.",
		Description:         "ObjectScale supports access by multiple tenants, where each tenant is defined by a iam_policy.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Identifier",
				MarkdownDescription: "Identifier",
				Computed:            true,
			},
			"namespace": schema.StringAttribute{
				Description:         "Name of the namespace from where the IAM.",
				MarkdownDescription: "Name of the namespace from where the IAM.",
				Required:            true,
			},
			"arn": schema.StringAttribute{
				Description:         "ARN of the IAM Policy to be fetched.",
				MarkdownDescription: "ARN of the IAM Policy to be fetched.",
				Optional:            true,
			},
			"user": schema.StringAttribute{
				Description:         "Name of the user whose attached policies are to be fetched.",
				MarkdownDescription: "Name of the user whose attached policies are to be fetched.",
				Optional:            true,
			},
			"group": schema.StringAttribute{
				Description:         "Name of the group whose attached policies are to be fetched.",
				MarkdownDescription: "Name of the group whose attached policies are to be fetched.",
				Optional:            true,
			},
			"role": schema.StringAttribute{
				Description:         "Name of the role whose attached policies are to be fetched.",
				MarkdownDescription: "Name of the role whose attached policies are to be fetched.",
				Optional:            true,
			},
			"policies": d.itemSchema(),
		},
	}
}

func (d *IAMPolicyDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		// validate that only one of arn, user, group or role can be set
		datasourcevalidator.Conflicting(
			path.MatchRoot("arn"),
			path.MatchRoot("user"),
			path.MatchRoot("group"),
			path.MatchRoot("role"),
		),
	}
}

func (d *IAMPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.IamPolicyDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	namespace := data.Namespace.ValueString()
	var allPolicyResp []clientgen.IamServiceGetPolicyResponseGetPolicyResultPolicy

	if arn := helper.ValueToPointer[string](data.ARN); arn != nil {
		// get by arn
		NsResp, _, err := d.client.GenClient.IamApi.IamServiceGetPolicy(ctx).XEmcNamespace(namespace).
			PolicyArn(*arn).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error fetching IAM policies with ARN: "+*arn, err.Error())
			return
		}
		allPolicyResp = append(allPolicyResp, *NsResp.GetPolicyResult.Policy)
	} else if user := helper.ValueToPointer[string](data.User); user != nil {
		// get by username
		dsreq := d.client.GenClient.IamApi.IamServiceListAttachedUserPolicies(ctx).XEmcNamespace(namespace).
			UserName(*user)
		dsresp, err := helper.GetAllInstances(dsreq)
		if err != nil {
			resp.Diagnostics.AddError("Error listing IAM policies attached to user: "+*user, err.Error())
			return
		}
		allPolicyResp = dsresp
	} else if group := helper.ValueToPointer[string](data.Group); group != nil {
		// get by group
		dsreq := d.client.GenClient.IamApi.IamServiceListAttachedGroupPolicies(ctx).XEmcNamespace(namespace).
			GroupName(*group)
		dsresp, err := helper.GetAllInstances(dsreq)
		if err != nil {
			resp.Diagnostics.AddError("Error listing IAM policies attached to group: "+*group, err.Error())
			return
		}
		allPolicyResp = dsresp
	} else if role := helper.ValueToPointer[string](data.Role); role != nil {
		// get by role
		dsreq := d.client.GenClient.IamApi.IamServiceListAttachedRolePolicies(ctx).XEmcNamespace(namespace).
			RoleName(*role)
		dsresp, err := helper.GetAllInstances(dsreq)
		if err != nil {
			resp.Diagnostics.AddError("Error listing IAM policies attached to role: "+*role, err.Error())
			return
		}
		allPolicyResp = dsresp
	} else {
		// get all policies
		dsreq := d.client.GenClient.IamApi.IamServiceListPolicies(ctx).XEmcNamespace(namespace)
		dsresp, err := helper.GetAllInstances(dsreq)
		if err != nil {
			resp.Diagnostics.AddError("Error listing IAM policies", err.Error())
			return
		}
		allPolicyResp = dsresp
	}

	IamPolicyList := d.updateState(allPolicyResp)

	// hardcoding a response value to save into the Terraform state.
	data.ID = types.StringValue("iam_policy_datasource")
	data.IamPolicies = IamPolicyList

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read iam_policy data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d IAMPolicyDataSource) updateState(iam_policys []clientgen.IamServiceGetPolicyResponseGetPolicyResultPolicy) []models.IamPolicyDataSourceIamPolicyModel {
	return helper.SliceTransform(iam_policys, func(v clientgen.IamServiceGetPolicyResponseGetPolicyResultPolicy) models.IamPolicyDataSourceIamPolicyModel {
		return models.IamPolicyDataSourceIamPolicyModel{
			ARN:                           helper.TfStringNN(v.Arn),
			AttachmentCount:               helper.TfInt32NN(v.AttachmentCount),
			CreateDate:                    helper.TfStringNN(v.CreateDate),
			DefaultVersionID:              helper.TfStringNN(v.DefaultVersionId),
			Description:                   helper.TfStringNN(v.Description),
			IsAttachable:                  helper.TfBoolNN(v.IsAttachable),
			Path:                          helper.TfStringNN(v.Path),
			PermissionsBoundaryUsageCount: helper.TfInt32NN(v.PermissionsBoundaryUsageCount),
			PolicyID:                      helper.TfStringNN(v.PolicyId),
			PolicyName:                    helper.TfStringNN(v.PolicyName),
			UpdateDate:                    helper.TfStringNN(v.UpdateDate),
		}
	})
}
