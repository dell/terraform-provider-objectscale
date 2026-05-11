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

var _ datasource.DataSource = &IAMServiceProviderDataSource{}

// NewIAMServiceProviderDataSource returns the singleton SP datasource.
func NewIAMServiceProviderDataSource() datasource.DataSource {
	return &IAMServiceProviderDataSource{}
}

// IAMServiceProviderDataSource reads the singleton SP config.
type IAMServiceProviderDataSource struct {
	datasourceProviderConfig
}

func (d *IAMServiceProviderDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_service_provider"
}

func (d *IAMServiceProviderDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "Reads the ObjectScale SAML Service Provider configuration (singleton).",
		MarkdownDescription: "Reads the ObjectScale SAML Service Provider configuration (singleton).",
		Attributes: map[string]schema.Attribute{
			"id":            schema.StringAttribute{Computed: true},
			"dns":           schema.StringAttribute{Computed: true},
			"uuid":          schema.StringAttribute{Computed: true},
			"unique_id":     schema.StringAttribute{Computed: true},
			"etag":          schema.StringAttribute{Computed: true},
			"key_alias":     schema.StringAttribute{Computed: true},
			"create_time":   schema.StringAttribute{Computed: true},
			"last_modified": schema.StringAttribute{Computed: true},
			"java_keystore": schema.StringAttribute{Computed: true, Sensitive: true},
			"key_password":  schema.StringAttribute{Computed: true, Sensitive: true},
		},
	}
}

func (d *IAMServiceProviderDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	getRes, _, err := d.client.GenClient.IamProviderApi.ServiceProviderGet(ctx).Execute()
	if err != nil {
		resp.Diagnostics.AddError("GetServiceProvider failed", classifyDiag(err).Error())
		return
	}
	sp := getRes.GetServiceProviderResult.ServiceProvider
	state := models.IAMServiceProviderDataSourceModel{
		ID:           types.StringValue("objectscale-sp"),
		DNS:          helper.TfStringNN(sp.Dns),
		UUID:         helper.TfStringNN(sp.Uuid),
		UniqueID:     helper.TfStringNN(sp.UniqueId),
		Etag:         helper.TfStringNN(sp.Etag),
		KeyAlias:     helper.TfStringNN(sp.KeyAlias),
		CreateTime:   helper.TfStringNN(sp.CreateTime),
		LastModified: helper.TfStringNN(sp.LastModified),
		JavaKeystore: helper.TfStringNN(sp.JavaKeystore),
		KeyPassword:  helper.TfStringNN(sp.KeyPassword),
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
