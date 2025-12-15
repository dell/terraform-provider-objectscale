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
var _ datasource.DataSource = &VDCDataSource{}

func NewVDCDataSource() datasource.DataSource {
	return &VDCDataSource{}
}

type VDCDataSource struct {
	datasourceProviderConfig
}

func (d *VDCDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vdc"
}

// datasource item schema
func (d *VDCDataSource) itemSchema() schema.ListNestedAttribute {
	return schema.ListNestedAttribute{
		Description:         "List of Virtual Data Centers fetched using this datasource.",
		MarkdownDescription: "List of Virtual Data Centers fetched using this datasource.",
		Computed:            true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				// VdcId types.String `tfsdk:"vdc_id"`
				"vdc_id": schema.StringAttribute{
					Description:         "VDC id",
					MarkdownDescription: "VDC id",
					Computed:            true,
				},
				// VdcName types.String `tfsdk:"vdc_name"`
				"vdc_name": schema.StringAttribute{
					Description:         "VDC name",
					MarkdownDescription: "VDC name",
					Computed:            true,
				},
				// InterVdcEndPoints types.String `tfsdk:"inter_vdc_end_points"`
				"inter_vdc_end_points": schema.StringAttribute{
					Description:         "VDC end points",
					MarkdownDescription: "VDC end points",
					Computed:            true,
				},
				// InterVdcCmdEndPoints types.String `tfsdk:"inter_vdc_cmd_end_points"`
				"inter_vdc_cmd_end_points": schema.StringAttribute{
					Description:         "VDC cmd end points",
					MarkdownDescription: "VDC cmd end points",
					Computed:            true,
				},
				// SecretKeys types.String `tfsdk:"secret_keys"`
				"secret_keys": schema.StringAttribute{
					Description:         "Secret key for this VDC",
					MarkdownDescription: "Secret key for this VDC",
					Computed:            true,
				},
				// PermanentlyFailed types.Bool `tfsdk:"permanently_failed"`
				"permanently_failed": schema.BoolAttribute{
					Description:         "True if VDC is permanently failed, false otherwise",
					MarkdownDescription: "True if VDC is permanently failed, false otherwise",
					Computed:            true,
				},
				// Local types.Bool `tfsdk:"local"`
				"local": schema.BoolAttribute{
					Description:         "True if this VDC is local, false otherwise",
					MarkdownDescription: "True if this VDC is local, false otherwise",
					Computed:            true,
				},
				// IsEncryptionEnabled types.Bool `tfsdk:"is_encryption_enabled"`
				"is_encryption_enabled": schema.BoolAttribute{
					Description:         "True if this VDC is enabled for encryption, false otherwise",
					MarkdownDescription: "True if this VDC is enabled for encryption, false otherwise",
					Computed:            true,
				},
				// ManagementEndPoints types.String `tfsdk:"management_end_points"`
				"management_end_points": schema.StringAttribute{
					Description:         "The management end points for the VDC",
					MarkdownDescription: "The management end points for the VDC",
					Computed:            true,
				},
				// Hosted types.Bool `tfsdk:"hosted"`
				"hosted": schema.BoolAttribute{
					Description:         "Indicates whether the VDC is hosted",
					MarkdownDescription: "Indicates whether the VDC is hosted",
					Computed:            true,
				},
				// Name types.String `tfsdk:"name"`
				"name": schema.StringAttribute{
					Description:         "Name assigned to this resource in ECS (name-defined, not unique)",
					MarkdownDescription: "Name assigned to this resource in ECS (name-defined, not unique)",
					Computed:            true,
				},
				// Id types.String `tfsdk:"id"`
				"id": schema.StringAttribute{
					Description:         "Unique, immutable identifier generated by ECS for this resource",
					MarkdownDescription: "Unique, immutable identifier generated by ECS for this resource",
					Computed:            true,
				},
				// CreationTime types.Int64 `tfsdk:"creation_time"`
				"creation_time": schema.Int64Attribute{
					Description:         "Timestamp when this resource was created in ECS",
					MarkdownDescription: "Timestamp when this resource was created in ECS",
					Computed:            true,
				},
				// Inactive types.Bool `tfsdk:"inactive"`
				"inactive": schema.BoolAttribute{
					Description:         "Indicates whether the resource is inactive (pre-removal state)",
					MarkdownDescription: "Indicates whether the resource is inactive (pre-removal state)",
					Computed:            true,
				},
				// Global types.Bool `tfsdk:"global"`
				"global": schema.BoolAttribute{
					Description:         "Indicates whether the resource is global",
					MarkdownDescription: "Indicates whether the resource is global",
					Computed:            true,
				},
				// Remote types.Bool `tfsdk:"remote"`
				"remote": schema.BoolAttribute{
					Description:         "Indicates whether the resource is remote",
					MarkdownDescription: "Indicates whether the resource is remote",
					Computed:            true,
				},
				// Internal types.Bool `tfsdk:"internal"`
				"internal": schema.BoolAttribute{
					Description:         "Indicates whether the resource is an internal resource",
					MarkdownDescription: "Indicates whether the resource is an internal resource",
					Computed:            true,
				},
			},
		},
	}
}

// Schema describes the data source arguments.
func (d *VDCDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This datasource can be used to fetch details of Virtual Data Centers from Dell ObjectScale.",
		Description:         "This datasource can be used to fetch details of Virtual Data Centers from Dell ObjectScale.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Identifier of the VDC to be fetched.",
				MarkdownDescription: "Identifier of the VDC to be fetched.",
				Optional:            true,
			},
			"name": schema.StringAttribute{
				Description:         "Name of the VDC to be fetched.",
				MarkdownDescription: "Name of the VDC to be fetched.",
				Optional:            true,
			},
			"local": schema.BoolAttribute{
				Description:         "Whether to fetch the local VDC.",
				MarkdownDescription: "Whether to fetch the local VDC.",
				Optional:            true,
			},
			"vdcs": d.itemSchema(),
		},
	}
}

func (d *VDCDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		// validate that only one of arn, name, group or role can be set
		datasourcevalidator.Conflicting(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
			path.MatchRoot("local"),
		),
	}
}

func (d *VDCDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.VDCDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// namespace := data.Namespace.ValueString()
	var allPolicyResp []clientgen.Vdc

	if id := helper.ValueToPointer[string](data.ID); id != nil {
		// get by id
		dsResp, _, err := d.client.GenClient.ZoneInfoApi.ZoneInfoServiceGetVdcById(ctx, *id).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error fetching VDC with ID: "+*id, err.Error())
			return
		}
		allPolicyResp = append(allPolicyResp, *dsResp)
	} else if name := helper.ValueToPointer[string](data.Name); name != nil {
		// get by name
		dsresp, _, err := d.client.GenClient.ZoneInfoApi.ZoneInfoServiceGetVdcByName(ctx, *name).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error fetching VDC with name: "+*name, err.Error())
			return
		}
		allPolicyResp = append(allPolicyResp, *dsresp)
	} else if local := helper.ValueToPointer[bool](data.Local); local != nil && *local {
		// get local
		dsresp, _, err := d.client.GenClient.ZoneInfoApi.ZoneInfoServiceGetLocalVdc(ctx).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error fetching local VDC", err.Error())
			return
		}
		allPolicyResp = append(allPolicyResp, *dsresp)
	} else {
		// get all VDCs
		dsresp, _, err := d.client.GenClient.ZoneInfoApi.ZoneInfoServiceListAllVdc(ctx).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error fetching VDCs", err.Error())
			return
		}
		allPolicyResp = dsresp.Vdc
	}

	IamPolicyList := d.updateState(allPolicyResp)

	// hardcoding a response value to save into the Terraform state.
	data.ID = types.StringValue("vdc_datasource")
	data.Vdcs = IamPolicyList

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read vdc data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d VDCDataSource) updateState(vdcs []clientgen.Vdc) []models.VdcDsItem {
	return helper.SliceTransform(vdcs, func(v clientgen.Vdc) models.VdcDsItem {
		return models.VdcDsItem{
			VdcId:   helper.TfStringNN(v.VdcId),
			VdcName: helper.TfStringNN(v.VdcName),

			// VDC end points
			InterVdcEndPoints: helper.TfStringNN(v.InterVdcEndPoints),
			// VDC cmd end points
			InterVdcCmdEndPoints: helper.TfStringNN(v.InterVdcCmdEndPoints),
			// Secret key for this VDC
			SecretKeys: helper.TfStringNN(v.SecretKeys),
			// True if VDC is permanently failed, false otherwise.
			PermanentlyFailed: helper.TfBoolNN(v.PermanentlyFailed),
			// True if this VDC is local, false otherwise
			Local: helper.TfBoolNN(v.Local),
			// True if this VDC is enabled for encryption, false otherwise
			IsEncryptionEnabled: helper.TfBoolNN(v.IsEncryptionEnabled),
			// The management end points for the VDC
			ManagementEndPoints: helper.TfStringNN(v.ManagementEndPoints),
			// Hosted
			Hosted: helper.TfBoolNN(v.Hosted),
			// Name assigned to this resource in ECS
			Name: helper.TfStringNN(v.Name),
			// Unique identifier generated by ECS
			Id: helper.TfStringNN(v.Id),
			// Timestamp when this resource was created in ECS
			CreationTime: helper.TfInt64NN(v.CreationTime),
			// Indicates whether the resource is inactive
			Inactive: helper.TfBoolNN(v.Inactive),
			// Indicates whether the resource is global
			Global: helper.TfBoolNN(v.Global),
			// Indicates whether the resource is remote
			Remote: helper.TfBoolNN(v.Remote),
			// Indicates whether the resource is an internal resource
			Internal: helper.TfBoolNN(v.Internal),
		}
	})
}
