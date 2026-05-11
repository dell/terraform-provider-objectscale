/*
Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.

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

	"terraform-provider-objectscale/internal/clientgen"
	"terraform-provider-objectscale/internal/helper"
	"terraform-provider-objectscale/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &IAMSAMLProviderResource{}
	_ resource.ResourceWithImportState = &IAMSAMLProviderResource{}
)

// NewIAMSAMLProviderResource returns the SAML Identity Provider resource.
func NewIAMSAMLProviderResource() resource.Resource {
	return &IAMSAMLProviderResource{}
}

// IAMSAMLProviderResource is the resource implementation for SAML IdPs.
type IAMSAMLProviderResource struct {
	resourceProviderConfig
}

// Metadata sets the resource type name.
func (r *IAMSAMLProviderResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_saml_provider"
}

// Schema returns the resource schema.
func (r *IAMSAMLProviderResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "Manages an ObjectScale IAM SAML Identity Provider (external IdP) registration.",
		MarkdownDescription: "Manages an ObjectScale IAM SAML Identity Provider (external IdP) registration.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "The provider ARN, also used as resource ID.",
				MarkdownDescription: "The provider ARN, also used as resource ID.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Description:         "The SAML Provider name. Cannot be changed after creation.",
				MarkdownDescription: "The SAML Provider name. Cannot be changed after creation.",
				Required:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"saml_metadata_document": schema.StringAttribute{
				Description:         "Raw SAML metadata XML for the IdP.",
				MarkdownDescription: "Raw SAML metadata XML for the IdP.",
				Required:            true,
			},
			"namespace": schema.StringAttribute{
				Description:         "Namespace of the SAML Provider. Cannot be changed after creation.",
				MarkdownDescription: "Namespace of the SAML Provider. Cannot be changed after creation.",
				Optional:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"arn": schema.StringAttribute{
				Description:         "ARN of the SAML Provider.",
				MarkdownDescription: "ARN of the SAML Provider.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"create_date": schema.StringAttribute{
				Description:         "ISO 8601 creation timestamp of the SAML Provider.",
				MarkdownDescription: "ISO 8601 creation timestamp of the SAML Provider.",
				Computed:            true,
			},
			"valid_until": schema.StringAttribute{
				Description:         "ISO 8601 timestamp at which the SAML Provider metadata signing certificate expires.",
				MarkdownDescription: "ISO 8601 timestamp at which the SAML Provider metadata signing certificate expires.",
				Computed:            true,
			},
		},
	}
}

// Create handles plan → create + read-after-write.
func (r *IAMSAMLProviderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.IAMSAMLProviderResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	name := plan.Name.ValueString()
	metadata := plan.SAMLMetadataDocument.ValueString()
	namespace := plan.Namespace.ValueString()

	tflog.Debug(ctx, "creating SAML IdP", map[string]interface{}{"name": name, "namespace": namespace})
	createRes, _, err := r.client.GenClient.IamApi.IamServiceCreateSAMLProvider(ctx).Name(name).SAMLMetadataDocument(metadata).XEmcNamespace(namespace).Execute()
	if err != nil {
		resp.Diagnostics.AddError("CreateSAMLProvider failed", classifyDiag(err).Error())
		return
	}

	getRes, _, err := r.client.GenClient.IamApi.IamServiceGetSAMLProvider(ctx).SAMLProviderArn(*createRes.CreateSAMLProviderResult.SAMLProviderArn).XEmcNamespace(namespace).Execute()
	if err != nil {
		resp.Diagnostics.AddError("GetSAMLProvider after create failed", classifyDiag(err).Error())
		return
	}

	arn := helper.TfStringNN(createRes.CreateSAMLProviderResult.SAMLProviderArn)
	data := r.getModel(getRes, arn)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read refreshes from the API; 404 removes the resource from state.
func (r *IAMSAMLProviderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.IAMSAMLProviderResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	getRes, _, err := r.client.GenClient.IamApi.IamServiceGetSAMLProvider(ctx).SAMLProviderArn(state.Arn.ValueString()).XEmcNamespace(state.Namespace.ValueString()).Execute()
	if err != nil {
		if helper.IsSAMLNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("GetSAMLProvider failed", classifyDiag(err).Error())
		return
	}

	data := r.getModel(getRes, state.Arn)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update handles in-place metadata updates.
func (r *IAMSAMLProviderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state models.IAMSAMLProviderResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	arn := state.Arn.ValueString()
	metadata := plan.SAMLMetadataDocument.ValueString()
	namespace := state.Namespace.ValueString()

	_, _, err := r.client.GenClient.IamApi.IamServiceUpdateSAMLProvider(ctx).SAMLProviderArn(arn).SAMLMetadataDocument(metadata).XEmcNamespace(namespace).Execute()
	if err != nil {
		resp.Diagnostics.AddError("UpdateSAMLProvider failed", classifyDiag(err).Error())
		return
	}

	getRes, _, err := r.client.GenClient.IamApi.IamServiceGetSAMLProvider(ctx).SAMLProviderArn(arn).XEmcNamespace(namespace).Execute()
	if err != nil {
		resp.Diagnostics.AddError("GetSAMLProvider after update failed", classifyDiag(err).Error())
		return
	}

	data := r.getModel(getRes, state.Arn)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete handles destroy. 404 is treated as success (I-10).
func (r *IAMSAMLProviderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state models.IAMSAMLProviderResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, _, err := r.client.GenClient.IamApi.IamServiceDeleteSAMLProvider(ctx).SAMLProviderArn(state.Arn.ValueString()).XEmcNamespace(state.Namespace.ValueString()).Execute()
	if err != nil && !helper.IsSAMLNotFound(err) {
		resp.Diagnostics.AddError("DeleteSAMLProvider failed", classifyDiag(err).Error())
		return
	}
}

// ImportState by ARN. Namespace is recovered from the ARN.
func (r *IAMSAMLProviderResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parsed, err := helper.ParseSAMLProviderARN(req.ID)
	if err != nil {
		resp.Diagnostics.AddError("Invalid import ID", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("arn"), req.ID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), parsed.Name)...)
	if parsed.Namespace != "" {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("namespace"), parsed.Namespace)...)
	}
}

// classifyDiag wraps an error in a more helpful diagnostic message.
func classifyDiag(err error) error {
	switch helper.ClassifyError(err) {
	case helper.SAMLErrBadRequest:
		return fmt.Errorf("validation error from ObjectScale: %w", err)
	case helper.SAMLErrUnauthorized:
		return fmt.Errorf("invalid or expired auth token: %w", err)
	case helper.SAMLErrForbidden:
		return fmt.Errorf("insufficient permissions: %w", err)
	case helper.SAMLErrNotFound:
		return fmt.Errorf("resource not found: %w", err)
	case helper.SAMLErrConflict:
		return fmt.Errorf("provider already exists: %w", err)
	default:
		return err
	}
}

func (r *IAMSAMLProviderResource) getModel(
	getRes *clientgen.IamServiceGetSAMLProviderResponse,
	arn types.String) models.IAMSAMLProviderResourceModel {
	parsed, _ := helper.ParseSAMLProviderARN(arn.ValueString())
	return models.IAMSAMLProviderResourceModel{
		ID:                   arn,
		Arn:                  arn,
		Name:                 types.StringValue(parsed.Name),
		Namespace:            types.StringValue(parsed.Namespace),
		SAMLMetadataDocument: helper.TfStringNN(getRes.GetSAMLProviderResult.SAMLMetadataDocument),
		CreateDate:           helper.TfStringNN(getRes.GetSAMLProviderResult.CreateDate),
		ValidUntil:           helper.TfStringNN(getRes.GetSAMLProviderResult.ValidUntil),
	}
}
