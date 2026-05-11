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

	"terraform-provider-objectscale/internal/helper"
	"terraform-provider-objectscale/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &IAMSAMLProviderDataSource{}

// NewIAMSAMLProviderDataSource returns the SAML provider datasource.
func NewIAMSAMLProviderDataSource() datasource.DataSource { return &IAMSAMLProviderDataSource{} }

// IAMSAMLProviderDataSource reads SAML Identity Providers.
// When saml_provider_arn is provided, returns a single provider.
// When saml_provider_arn is not provided, returns a list of providers.
type IAMSAMLProviderDataSource struct {
	datasourceProviderConfig
}

func (d *IAMSAMLProviderDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_saml_provider"
}

func (d *IAMSAMLProviderDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	providerObj := schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"arn":                    schema.StringAttribute{Computed: true, Description: "Provider ARN."},
			"name":                   schema.StringAttribute{Computed: true, Description: "Provider name."},
			"saml_metadata_document": schema.StringAttribute{Computed: true, Description: "SAML metadata XML document."},
			"create_date":            schema.StringAttribute{Computed: true, Description: "Creation timestamp."},
			"valid_until":            schema.StringAttribute{Computed: true, Description: "Certificate expiration timestamp."},
		},
	}
	resp.Schema = schema.Schema{
		Description:         "Reads ObjectScale IAM SAML Identity Providers. When saml_provider_arn is provided, returns a single provider. When not provided, returns a list of providers.",
		MarkdownDescription: "Reads ObjectScale IAM SAML Identity Providers. When `saml_provider_arn` is provided, returns a single provider. When not provided, returns a list of providers.",
		Attributes: map[string]schema.Attribute{
			"id":                schema.StringAttribute{Computed: true, Description: "Identifier."},
			"saml_provider_arn": schema.StringAttribute{Optional: true, Description: "Provider ARN to look up. If not provided, lists all providers."},
			"namespace":         schema.StringAttribute{Optional: true, Description: "Optional `x-emc-namespace` header (management user only)."},
			"providers": schema.ListNestedAttribute{
				Computed:            true,
				NestedObject:        providerObj,
				Description:         "List of SAML providers.",
				MarkdownDescription: "List of SAML providers.",
			},
		},
	}
}

func (d *IAMSAMLProviderDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state models.IAMSAMLProviderDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If saml_provider_arn is provided, get a single provider
	if !state.SAMLProviderArn.IsNull() && !state.SAMLProviderArn.IsUnknown() {
		arn := state.SAMLProviderArn.ValueString()
		parsed, err := helper.ParseSAMLProviderARN(arn)
		if err != nil {
			resp.Diagnostics.AddError("Invalid saml_provider_arn", err.Error())
			return
		}
		getRes, _, err := d.client.GenClient.IamApi.IamServiceGetSAMLProvider(ctx).SAMLProviderArn(arn).XEmcNamespace(state.Namespace.ValueString()).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error getting SAML provider by ARN", classifyDiag(err).Error())
			return
		}
		state.ID = state.SAMLProviderArn
		state.Providers = []models.IAMSAMLProvider{
			{
				Arn:                  types.StringValue(arn),
				Name:                 types.StringValue(parsed.Name),
				SAMLMetadataDocument: helper.TfStringNN(getRes.GetSAMLProviderResult.SAMLMetadataDocument),
				CreateDate:           helper.TfStringNN(getRes.GetSAMLProviderResult.CreateDate),
				ValidUntil:           helper.TfStringNN(getRes.GetSAMLProviderResult.ValidUntil),
			},
		}
	} else {
		// List all providers using helper pagination
		req := d.client.GenClient.IamApi.IamServiceListSAMLProviders(ctx).XEmcNamespace(state.Namespace.ValueString())
		allProviders, err := helper.GetAllInstances(req)
		if err != nil {
			resp.Diagnostics.AddError("Error listing SAML providers", classifyDiag(err).Error())
			return
		}

		var all []models.IAMSAMLProvider
		for _, p := range allProviders {
			parsed, _ := helper.ParseSAMLProviderARN(*p.Arn)
			all = append(all, models.IAMSAMLProvider{
				Arn:        helper.TfStringNN(p.Arn),
				Name:       types.StringValue(parsed.Name),
				CreateDate: helper.TfStringNN(p.CreateDate),
				ValidUntil: helper.TfStringNN(p.ValidUntil),
			})
		}

		state.ID = types.StringValue("saml-providers")
		state.Providers = all
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
