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
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &ObjectCertificateDataSource{}

func NewObjectCertificateDataSource() datasource.DataSource {
	return &ObjectCertificateDataSource{}
}

// ObjectCertificateDataSource reads the current Object data-plane certificate chain.
type ObjectCertificateDataSource struct {
	datasourceProviderConfig
}

func (d *ObjectCertificateDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_object_certificate"
}

func (d *ObjectCertificateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This datasource reads the current Object data-plane (S3) TLS certificate chain from Dell ObjectScale.",
		MarkdownDescription: "This datasource reads the current Object data-plane (S3) TLS certificate chain from Dell ObjectScale.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Identifier for this data source.",
				MarkdownDescription: "Identifier for this data source.",
				Computed:            true,
			},
			"certificate_chain": schema.StringAttribute{
				Description:         "Current Object certificate chain in PEM format.",
				MarkdownDescription: "Current Object certificate chain in PEM format.",
				Computed:            true,
			},
		},
	}
}

func (d *ObjectCertificateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.ObjectCertificateDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	chain, err := GetObjectCertKeystore(ctx, d.client)
	if err != nil {
		resp.Diagnostics.AddError("Error reading Object certificate", err.Error())
		return
	}

	normalizedChain := helper.NormalizeLineEndings(chain)

	data.ID = types.StringValue("object_certificate_datasource")
	data.CertificateChain = types.StringValue(normalizedChain)

	tflog.Trace(ctx, "read Object certificate data source")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
