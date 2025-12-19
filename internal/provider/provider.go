/*
Copyright (c) 2023-2024 Dell Inc., or its subsidiaries. All Rights Reserved.

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
	"terraform-provider-objectscale/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure ObjectScaleProvider satisfies various provider interfaces.
var _ provider.Provider = &ObjectScaleProvider{}

// ObjectScaleProvider defines the provider implementation.
type ObjectScaleProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// ObjectScaleProviderModel describes the provider data model.
type ObjectScaleProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
	Insecure types.Bool   `tfsdk:"insecure"`
	Timeout  types.Int64  `tfsdk:"timeout"`
}

// Metadata describes the provider arguments.
func (p *ObjectScaleProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "objectscale"
	resp.Version = p.version
}

// Schema describes the provider arguments.
func (p *ObjectScaleProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The Terraform provider for Dell Objectscale can be used to interact with a Dell Objectscale array in order to manage the array resources.",
		Description:         "The Terraform provider for Dell Objectscale can be used to interact with a Dell Objectscale array in order to manage the array resources.",
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "The API endpoint, ex. https://10.225.100.1:4443",
				Description:         "The API endpoint, ex. https://10.225.100.1:4443",
				Required:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "The username",
				Description:         "The username",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "The password",
				Description:         "The password",
				Required:            true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"insecure": schema.BoolAttribute{
				MarkdownDescription: "whether to skip SSL validation",
				Description:         "whether to skip SSL validation",
				Optional:            true,
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: "The timeout in seconds",
				Description:         "The timeout in seconds",
				Optional:            true,
			},
		},
	}
}

// Configure configures the provider.
func (p *ObjectScaleProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data ObjectScaleProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	client, err := client.NewClient(
		data.Endpoint.ValueString(),
		data.Username.ValueString(),
		data.Password.ValueString(),
		data.Insecure.ValueBool(),
		data.Timeout.ValueInt64(),
	)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create objectscale client",
			err.Error(),
		)
		return
	}

	// client configuration for data sources and resources
	resp.DataSourceData = client
	resp.ResourceData = client
}

// Resources describes the provider resources.
func (p *ObjectScaleProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewNamespaceResource,
		NewIAMUserResource,
		NewIAMInlinePolicyResource,
		NewIAMPolicyAttachmentResource,
		NewIAMGroupResource,
		NewBucketResource,
		NewIAMGroupMembershipResource,
		NewIAMUserAccessKeyResource,
		NewIAMRoleResource,
		NewObjectUserResource,
	}
}

// DataSources describes the provider data sources.
func (p *ObjectScaleProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewNamespaceDataSource,
		NewIAMGroupsDataSource,
		NewIAMUserDataSource,
		NewReplicationGroupDataSource,
		NewIAMPolicyDataSource,
		NewIAMRoleDataSource,
		NewBucketDataSource,
	}
}

// New returns a new provider instance.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ObjectScaleProvider{
			version: version,
		}
	}
}

// datasourceProviderConfig defines the provider config struct
type datasourceProviderConfig struct {
	client *client.Client
}

func (d *datasourceProviderConfig) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}
	d.client = client
}

// resourceProviderConfig defines the provider config struct
type resourceProviderConfig struct {
	client *client.Client
}

func (r *resourceProviderConfig) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}
	r.client = client
}
