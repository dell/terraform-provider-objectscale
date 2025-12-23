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
	"strings"
	"terraform-provider-objectscale/internal/clientgen"
	"terraform-provider-objectscale/internal/helper"
	"terraform-provider-objectscale/internal/models"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &IAMPolicyResource{}
	_ resource.ResourceWithConfigure = &IAMPolicyResource{}
)

func NewIAMPolicyResource() resource.Resource {
	return &IAMPolicyResource{}
}

type IAMPolicyResource struct {
	resourceProviderConfig
}

func (r *IAMPolicyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_policy"
}

func (r *IAMPolicyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "Manages an ObjectScale IAM Policy.",
		MarkdownDescription: "Manages an ObjectScale IAM Policy.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Description:         "The name of the IAM Policy.",
				MarkdownDescription: "The name of the IAM Policy.",
				Required:            true,
			},
			"policy_document": schema.StringAttribute{
				Description:         "A map of policy document versions to their JSON content.",
				MarkdownDescription: "A map of policy document versions to their JSON content.",
				Required:            true,
				CustomType:          jsontypes.NormalizedType{},
			},
			"namespace": schema.StringAttribute{
				Description:         "The namespace in which to create the IAM Policy.",
				MarkdownDescription: "The namespace in which to create the IAM Policy.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				Description:         "The description of the IAM Policy.",
				MarkdownDescription: "The description of the IAM Policy.",
				Optional:            true,
			},

			"version_id": schema.StringAttribute{
				Description:         "The ID of the default policy document version.",
				MarkdownDescription: "The ID of the default policy document version.",
				Computed:            true,
			},

			"arn": schema.StringAttribute{
				Description:         "The Amazon Resource Name (ARN) of the IAM Policy.",
				MarkdownDescription: "The Amazon Resource Name (ARN) of the IAM Policy.",
				Computed:            true,
			},

			"create_date": schema.StringAttribute{
				Description:         "The creation date of the IAM Policy.",
				MarkdownDescription: "The creation date of the IAM Policy.",
				Computed:            true,
			},
		},
	}
}

func (r *IAMPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.IamPolicyResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	creq := r.client.GenClient.IamApi.IamServiceCreatePolicy(ctx).
		PolicyName(plan.PolicyName.ValueString()).
		PolicyDocument(plan.PolicyDocument.ValueString()).
		XEmcNamespace(plan.Namespace.ValueString()).
		Description(plan.Description.ValueString())

	iam_policy, _, err := creq.Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating IAM Policy",
			"Could not create IAM Policy: "+err.Error(),
		)
		return
	}

	data := r.getModel(&clientgen.IamServiceCreatePolicyResponseCreatePolicyResultPolicy{
		PolicyName:       iam_policy.CreatePolicyResult.Policy.PolicyName,
		Arn:              iam_policy.CreatePolicyResult.Policy.Arn,
		CreateDate:       iam_policy.CreatePolicyResult.Policy.CreateDate,
		DefaultVersionId: iam_policy.CreatePolicyResult.Policy.DefaultVersionId,
		Description:      iam_policy.CreatePolicyResult.Policy.Description,
	}, plan.PolicyDocument, plan.Namespace)

	// save into state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IAMPolicyResource) getModel(
	iam_policy *clientgen.IamServiceCreatePolicyResponseCreatePolicyResultPolicy,
	policyDocument jsontypes.Normalized,
	namespace types.String) models.IamPolicyResourceModel {

	return models.IamPolicyResourceModel{
		Arn:            helper.TfStringNN(iam_policy.Arn),
		CreateDate:     helper.TfStringNN(iam_policy.CreateDate),
		VersionId:      helper.TfStringNN(iam_policy.DefaultVersionId),
		PolicyName:     helper.TfStringNN(iam_policy.PolicyName),
		Description:    helper.TfString(iam_policy.Description),
		Namespace:      namespace,
		PolicyDocument: policyDocument,
	}
}

func (r *IAMPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.IamPolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	iam_policy, _, err := r.client.GenClient.IamApi.IamServiceGetPolicy(ctx).
		PolicyArn(state.Arn.ValueString()).
		XEmcNamespace(state.Namespace.ValueString()).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError("Error reading IAM Policy", err.Error())
		return
	}

	iam_policy_document, _, err := r.client.GenClient.IamApi.IamServiceGetPolicyVersion(ctx).
		PolicyArn(state.Arn.ValueString()).
		VersionId(*iam_policy.GetPolicyResult.Policy.DefaultVersionId).
		XEmcNamespace(state.Namespace.ValueString()).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError("Error reading IAM Policy", err.Error())
		return
	}

	var policyDocument jsontypes.Normalized
	policyDocument = jsontypes.NewNormalizedValue(IAMPolicyDataSource{}.decodeDocument(iam_policy_document.GetPolicyVersionResult.PolicyVersion.Document).ValueString())

	data := r.getModel(&clientgen.IamServiceCreatePolicyResponseCreatePolicyResultPolicy{
		PolicyName:       iam_policy.GetPolicyResult.Policy.PolicyName,
		Arn:              iam_policy.GetPolicyResult.Policy.Arn,
		CreateDate:       iam_policy.GetPolicyResult.Policy.CreateDate,
		DefaultVersionId: iam_policy.GetPolicyResult.Policy.DefaultVersionId,
		Description:      iam_policy.GetPolicyResult.Policy.Description,
	}, policyDocument, state.Namespace)

	// Save updated plan into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IAMPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan models.IamPolicyResourceModel
	var state models.IamPolicyResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if helper.IsChangedNN(plan.PolicyName, state.PolicyName) || helper.IsChangedNN(plan.Description, state.Description) {
		resp.Diagnostics.AddError("Unexpected Update Parameter : Only Policy Document is updateable", "Invalid Update")
		return
	}

	listreq := r.client.GenClient.IamApi.IamServiceListPolicyVersions(ctx).
		PolicyArn(state.Arn.ValueString()).
		XEmcNamespace(plan.Namespace.ValueString())

	versionsResp, _, err := listreq.Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating IAM Policy",
			"Could not update IAM Policy: "+err.Error(),
		)
		return
	}

	for _, v := range versionsResp.ListPolicyVersionsResult.Versions {
		if !*v.IsDefaultVersion {
			dreq := r.client.GenClient.IamApi.IamServiceDeletePolicyVersion(ctx).
				PolicyArn(state.Arn.ValueString()).
				VersionId(*v.VersionId).
				XEmcNamespace(plan.Namespace.ValueString())

			_, _, err := dreq.Execute()
			if err != nil {
				resp.Diagnostics.AddError(
					"Error deleting non-default IAM Policy Version : "+*v.VersionId,
					err.Error(),
				)
				return
			}
		}
	}

	updReq := r.client.GenClient.IamApi.IamServiceCreatePolicyVersion(ctx).
		PolicyArn(state.Arn.ValueString()).
		PolicyDocument(plan.PolicyDocument.ValueString()).
		SetAsDefault(true).
		XEmcNamespace(plan.Namespace.ValueString())

	iam_policy_version, _, err := updReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating new IAM Policy Version",
			err.Error(),
		)
		return
	}

	policyNameStr := state.PolicyName.ValueString()
	arnStr := state.Arn.ValueString()
	createDateStr := state.CreateDate.ValueString()
	descriptionStr := state.Description.ValueString()

	data := r.getModel(&clientgen.IamServiceCreatePolicyResponseCreatePolicyResultPolicy{
		PolicyName:       &policyNameStr,
		Arn:              &arnStr,
		CreateDate:       &createDateStr,
		DefaultVersionId: iam_policy_version.CreatePolicyVersionResult.PolicyVersion.VersionId,
		Description:      &descriptionStr,
	}, plan.PolicyDocument, plan.Namespace)

	// save into state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IAMPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state models.IamPolicyResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dreq := r.client.GenClient.IamApi.IamServiceDeletePolicy(ctx).
		PolicyArn(state.Arn.ValueString()).
		XEmcNamespace(state.Namespace.ValueString())

	_, _, err := dreq.Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting IAM Policy",
			"Could not delete IAM Policy: "+err.Error(),
		)
		return
	}

}

// Import state function
func (r *IAMPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.SplitN(req.ID, "#", 2)
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Error importing IAM Group", "invalid format: expected 'group_name#namespace'")
		return
	}
	policyArn := parts[0]
	namespace := parts[1]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("arn"), policyArn)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("namespace"), namespace)...)
}
