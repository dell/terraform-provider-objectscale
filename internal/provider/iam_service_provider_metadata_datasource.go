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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &IAMServiceProviderMetadataDataSource{}

// NewIAMServiceProviderMetadataDataSource returns the SP metadata datasource.
func NewIAMServiceProviderMetadataDataSource() datasource.DataSource {
	return &IAMServiceProviderMetadataDataSource{}
}

// IAMServiceProviderMetadataDataSource reads + parses the SP metadata XML.
type IAMServiceProviderMetadataDataSource struct {
	datasourceProviderConfig
}

func (d *IAMServiceProviderMetadataDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_service_provider_metadata"
}

func (d *IAMServiceProviderMetadataDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "Reads ObjectScale SAML SP metadata XML and parses key fields for IdP onboarding.",
		MarkdownDescription: "Reads ObjectScale SAML SP metadata XML and parses key fields for IdP onboarding.",
		Attributes: map[string]schema.Attribute{
			"id":                     schema.StringAttribute{Computed: true},
			"metadata_xml":           schema.StringAttribute{Computed: true, Description: "Raw EntityDescriptor XML."},
			"entity_id":              schema.StringAttribute{Computed: true},
			"acs_url":                schema.StringAttribute{Computed: true},
			"authn_requests_signed":  schema.BoolAttribute{Computed: true},
			"want_assertions_signed": schema.BoolAttribute{Computed: true},
			"signing_certificate":    schema.StringAttribute{Computed: true},
			"name_id_formats": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Supported NameIDFormat values.",
			},
		},
	}
}

func (d *IAMServiceProviderMetadataDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	rawXML, _, err := d.client.GenClient.IamProviderApi.ServiceProviderGetMetadata(ctx).Execute()
	if err != nil {
		resp.Diagnostics.AddError("GetServiceProviderMetadata failed", classifyDiag(err).Error())
		return
	}
	parsed, err := helper.ParseSPMetadata(rawXML)
	if err != nil {
		resp.Diagnostics.AddError("Parse SP metadata failed", err.Error())
		return
	}
	values := make([]attr.Value, 0, len(parsed.NameIDFormats))
	for _, n := range parsed.NameIDFormats {
		values = append(values, types.StringValue(n))
	}
	listVal, listDiags := types.ListValue(types.StringType, values)
	resp.Diagnostics.Append(listDiags...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := models.IAMServiceProviderMetadataDataSourceModel{
		ID:                   types.StringValue("objectscale-sp-metadata"),
		MetadataXML:          types.StringValue(rawXML),
		EntityID:             types.StringValue(parsed.EntityID),
		ACSURL:               types.StringValue(parsed.ACSURL),
		AuthnRequestsSigned:  types.BoolValue(parsed.AuthnRequestsSigned),
		WantAssertionsSigned: types.BoolValue(parsed.WantAssertionsSigned),
		SigningCertificate:   types.StringValue(parsed.SigningCertificate),
		NameIDFormats:        listVal,
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
