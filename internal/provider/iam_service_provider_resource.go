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
	_ resource.Resource                = &IAMServiceProviderResource{}
	_ resource.ResourceWithImportState = &IAMServiceProviderResource{}
)

// NewIAMServiceProviderResource returns the singleton SP resource.
func NewIAMServiceProviderResource() resource.Resource { return &IAMServiceProviderResource{} }

// IAMServiceProviderResource manages the (singleton) ObjectScale SAML SP config.
type IAMServiceProviderResource struct {
	resourceProviderConfig
}

func (r *IAMServiceProviderResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_service_provider"
}

func (r *IAMServiceProviderResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "Manages the ObjectScale SAML Service Provider configuration (singleton per cluster).",
		MarkdownDescription: "Manages the ObjectScale SAML Service Provider configuration (singleton per cluster).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Singleton ID; always equal to `objectscale-sp`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"dns": schema.StringAttribute{
				Required:    true,
				Description: "Service Provider base URL for the SAML ACS.",
			},
			"java_keystore": schema.StringAttribute{
				Required:    true,
				Sensitive:   true,
				Description: "Base64-encoded Java KeyStore.",
			},
			"key_alias": schema.StringAttribute{
				Required:    true,
				Description: "KeyStore entry alias.",
			},
			"key_password": schema.StringAttribute{
				Required:    true,
				Sensitive:   true,
				Description: "KeyStore password.",
			},
			"uuid":          schema.StringAttribute{Computed: true, Description: "Entity Id component."},
			"unique_id":     schema.StringAttribute{Computed: true, Description: "KeyStore unique ID."},
			"etag":          schema.StringAttribute{Computed: true, Description: "Optimistic concurrency tag."},
			"create_time":   schema.StringAttribute{Computed: true, Description: "ISO 8601 creation timestamp."},
			"last_modified": schema.StringAttribute{Computed: true, Description: "ISO 8601 last-modified timestamp."},
		},
	}
}

func (r *IAMServiceProviderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.IAMServiceProviderResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := helper.ValidateSPDNS(plan.DNS.ValueString()); err != nil {
		resp.Diagnostics.AddAttributeError(path.Root("dns"), "Invalid SP DNS", err.Error())
		return
	}
	tflog.Debug(ctx, "creating SP config", map[string]interface{}{"dns": plan.DNS.ValueString()})
	body := r.buildRequestBody(&plan)
	_, _, err := r.client.GenClient.IamProviderApi.ServiceProviderCreate(ctx).IamServiceProviderControllerProcessCreateServiceProviderRequest(body).Execute()
	if err != nil {
		// Singleton: if it already exists, update in place.
		if helper.ClassifyError(err) == helper.SAMLErrConflict {
			tflog.Info(ctx, "SP already exists, updating in place")
			updateBody := r.buildUpdateBody(&plan)
			if _, _, uErr := r.client.GenClient.IamProviderApi.ServiceProviderUpdate(ctx).IamServiceProviderControllerProcessUpdateServiceProviderRequest(updateBody).Execute(); uErr != nil {
				resp.Diagnostics.AddError("UpdateServiceProvider (upsert) failed", classifyDiag(uErr).Error())
				return
			}
		} else {
			resp.Diagnostics.AddError("CreateServiceProvider failed", classifyDiag(err).Error())
			return
		}
	}
	getRes, _, err := r.client.GenClient.IamProviderApi.ServiceProviderGet(ctx).Execute()
	if err != nil {
		resp.Diagnostics.AddError("GetServiceProvider after create failed", classifyDiag(err).Error())
		return
	}
	data := r.getModel(getRes, plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IAMServiceProviderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.IAMServiceProviderResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	getRes, _, err := r.client.GenClient.IamProviderApi.ServiceProviderGet(ctx).Execute()
	if err != nil {
		if helper.IsSAMLNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("GetServiceProvider failed", classifyDiag(err).Error())
		return
	}
	data := r.getModel(getRes, state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IAMServiceProviderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan models.IAMServiceProviderResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	body := r.buildUpdateBody(&plan)
	if _, _, err := r.client.GenClient.IamProviderApi.ServiceProviderUpdate(ctx).IamServiceProviderControllerProcessUpdateServiceProviderRequest(body).Execute(); err != nil {
		resp.Diagnostics.AddError("UpdateServiceProvider failed", classifyDiag(err).Error())
		return
	}
	getRes, _, err := r.client.GenClient.IamProviderApi.ServiceProviderGet(ctx).Execute()
	if err != nil {
		resp.Diagnostics.AddError("GetServiceProvider after update failed", classifyDiag(err).Error())
		return
	}
	data := r.getModel(getRes, plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IAMServiceProviderResource) Delete(ctx context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	if _, _, err := r.client.GenClient.IamProviderApi.ServiceProviderDelete(ctx).Execute(); err != nil && !helper.IsSAMLNotFound(err) {
		resp.Diagnostics.AddError("DeleteServiceProvider failed", classifyDiag(err).Error())
		return
	}
}

// ImportState — singleton, ID ignored.
func (r *IAMServiceProviderResource) ImportState(ctx context.Context, _ resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), "objectscale-sp")...)
}

func (r *IAMServiceProviderResource) getModel(
	getRes *clientgen.ServiceProviderGetResponse,
	prior models.IAMServiceProviderResourceModel,
) models.IAMServiceProviderResourceModel {
	sp := getRes.GetServiceProviderResult.ServiceProvider
	m := models.IAMServiceProviderResourceModel{
		ID:           types.StringValue("objectscale-sp"),
		UUID:         helper.TfStringNN(sp.Uuid),
		UniqueID:     helper.TfStringNN(sp.UniqueId),
		Etag:         helper.TfStringNN(sp.Etag),
		CreateTime:   helper.TfStringNN(sp.CreateTime),
		LastModified: helper.TfStringNN(sp.LastModified),
		DNS:          helper.TfStringNN(sp.Dns),
		KeyAlias:     helper.TfStringNN(sp.KeyAlias),
		JavaKeystore: prior.JavaKeystore,
		KeyPassword:  prior.KeyPassword,
	}
	if sp.JavaKeystore != nil && *sp.JavaKeystore != "" {
		m.JavaKeystore = helper.TfStringNN(sp.JavaKeystore)
	}
	if sp.KeyPassword != nil && *sp.KeyPassword != "" {
		m.KeyPassword = helper.TfStringNN(sp.KeyPassword)
	}
	return m
}

func (r *IAMServiceProviderResource) buildRequestBody(plan *models.IAMServiceProviderResourceModel) clientgen.IamServiceProviderControllerProcessCreateServiceProviderRequest {
	dns := plan.DNS.ValueString()
	jks := plan.JavaKeystore.ValueString()
	alias := plan.KeyAlias.ValueString()
	pwd := plan.KeyPassword.ValueString()
	return clientgen.IamServiceProviderControllerProcessCreateServiceProviderRequest{
		ServiceProvider: &clientgen.ServiceProvider{
			Dns:          &dns,
			JavaKeystore: &jks,
			KeyAlias:     &alias,
			KeyPassword:  &pwd,
		},
	}
}

func (r *IAMServiceProviderResource) buildUpdateBody(plan *models.IAMServiceProviderResourceModel) clientgen.IamServiceProviderControllerProcessUpdateServiceProviderRequest {
	dns := plan.DNS.ValueString()
	jks := plan.JavaKeystore.ValueString()
	alias := plan.KeyAlias.ValueString()
	pwd := plan.KeyPassword.ValueString()
	return clientgen.IamServiceProviderControllerProcessUpdateServiceProviderRequest{
		ServiceProvider: &clientgen.ServiceProvider{
			Dns:          &dns,
			JavaKeystore: &jks,
			KeyAlias:     &alias,
			KeyPassword:  &pwd,
		},
	}
}
