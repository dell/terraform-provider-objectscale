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
var _ datasource.DataSource = &StoragePoolDataSource{}

func NewStoragePoolDataSource() datasource.DataSource {
	return &StoragePoolDataSource{}
}

type StoragePoolDataSource struct {
	datasourceProviderConfig
}

func (d *StoragePoolDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_storage_pool"
}

// datasource item schema.
func (d *StoragePoolDataSource) itemSchema() schema.ListNestedAttribute {
	return schema.ListNestedAttribute{
		Description:         "List of Storage Pools fetched using this datasource.",
		MarkdownDescription: "List of Storage Pools fetched using this datasource.",
		Computed:            true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				// Id types.String `tfsdk:"id"`
				"id": schema.StringAttribute{
					Description:         "Storage pool id",
					MarkdownDescription: "Storage pool id",
					Computed:            true,
				},
				// Name types.String `tfsdk:"name"`
				"name": schema.StringAttribute{
					Description:         "Storage pool name",
					MarkdownDescription: "Storage pool name",
					Computed:            true,
				},
				// Description types.String `tfsdk:"description"`
				"description": schema.StringAttribute{
					Description:         "Description",
					MarkdownDescription: "Description",
					Computed:            true,
				},
				// IsColdStorageEnabled types.Bool `tfsdk:"is_cold_storage_enabled"`
				"is_cold_storage_enabled": schema.BoolAttribute{
					Description:         "Flag indicating that cold storage encoding is enabled",
					MarkdownDescription: "Flag indicating that cold storage encoding is enabled",
					Computed:            true,
				},
				// NumberOfDataBlocks types.Int64 `tfsdk:"number_of_data_blocks"`
				"number_of_data_blocks": schema.Int32Attribute{
					Description:         "Number of Data Blocks in EC Scheme",
					MarkdownDescription: "Number of Data Blocks in EC Scheme",
					Computed:            true,
				},
				// NumberOfCodeBlocks types.Int64 `tfsdk:"number_of_code_blocks"`
				"number_of_code_blocks": schema.Int32Attribute{
					Description:         "Number of Code Blocks in EC Scheme",
					MarkdownDescription: "Number of Code Blocks in EC Scheme",
					Computed:            true,
				},
				// WarningAlertAt types.Int64 `tfsdk:"warning_alert_at"`
				"warning_alert_at": schema.Int32Attribute{
					Description:         "Threshold percent at which warning alert is raised.  Valid values are from -1 to 100. Value of -1 means do not alert",
					MarkdownDescription: "Threshold percent at which warning alert is raised.  Valid values are from -1 to 100. Value of -1 means do not alert",
					Computed:            true,
				},
				// ErrorAlertAt types.Int64 `tfsdk:"error_alert_at"`
				"error_alert_at": schema.Int32Attribute{
					Description:         "Threshold percent at which error alert is raised.  Valid values are from -1 to 100. Value of -1 means do not alert",
					MarkdownDescription: "Threshold percent at which error alert is raised.  Valid values are from -1 to 100. Value of -1 means do not alert",
					Computed:            true,
				},
				// CriticalAlertAt types.Int64 `tfsdk:"critical_alert_at"`
				"critical_alert_at": schema.Int32Attribute{
					Description:         "Threshold percent at which critical alert is raised.  Valid values are from -1 to 100. Value of -1 means do not alert",
					MarkdownDescription: "Threshold percent at which critical alert is raised.  Valid values are from -1 to 100. Value of -1 means do not alert",
					Computed:            true,
				},
				// Status types.Int32 `tfsdk:"status"`
				"status": schema.Int32Attribute{
					Description:         "flag for status, -1 for null, 0 ~ 6 for value",
					MarkdownDescription: "flag for status, -1 for null, 0 ~ 6 for value",
					Computed:            true,
				},
				// Label types.String `tfsdk:"label"`
				"label": schema.StringAttribute{
					Description:         "Lbel of VArray",
					MarkdownDescription: "Lbel of VArray",
					Computed:            true,
				},
				// DriveTechnology types.String `tfsdk:"drive_technology"`
				"drive_technology": schema.StringAttribute{
					Description:         "Drive technology of VArray",
					MarkdownDescription: "Drive technology of VArray",
					Computed:            true,
				},
			},
		},
	}
}

// Schema describes the data source arguments.
func (d *StoragePoolDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This datasource can be used to fetch details of Storage Pools from Dell ObjectScale.",
		Description:         "This datasource can be used to fetch details of Storage Pools from Dell ObjectScale.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier of the Storage Pool to be fetched." +
					" Conflicts with `vdc_id` and `name`.",
				MarkdownDescription: "Identifier of the Storage Pool to be fetched." +
					" Conflicts with `vdc_id` and `name`.",
				Optional: true,
			},
			"vdc_id": schema.StringAttribute{
				Description: "ID of Virtual Datacenters from which Storage Pool(s) are to be fetched." +
					" If none given, the local VDC is used." +
					" Conflicts with `id`.",
				MarkdownDescription: "ID of Virtual Datacenters from which Storage Pool(s) are to be fetched." +
					" If none given, the local VDC is used." +
					" Conflicts with `id`.",
				Optional: true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the Storage Pool to be fetched." +
					" This is an offline filter." +
					" Conflicts with `id`.",
				MarkdownDescription: "Name of the Storage Pool to be fetched." +
					" This is an offline filter." +
					" Conflicts with `id`.",
				Optional: true,
			},
			"storage_pools": d.itemSchema(),
		},
	}
}

func (d *StoragePoolDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		// validate that id and vdc_id cannot be set at the same time
		datasourcevalidator.Conflicting(
			path.MatchRoot("id"),
			path.MatchRoot("vdc_id"),
		),
		// validate that only one of name and id cannot be set at the same time
		datasourcevalidator.Conflicting(
			path.MatchRoot("name"),
			path.MatchRoot("id"),
		),
	}
}

func (d *StoragePoolDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.StoragePoolDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var allSpResp []clientgen.ObjectVarrayServiceGetVirtualArrayResponse

	if id := helper.ValueToPointer[string](data.ID); id != nil {
		// get by id
		dsResp, _, err := d.client.GenClient.ObjectVarrayApi.ObjectVarrayServiceGetVirtualArray(ctx, *id).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error fetching Storage Pool with ID: "+*id, err.Error())
			return
		}
		allSpResp = append(allSpResp, *dsResp)
	} else {
		// list Storage Pool request
		dsreq := d.client.GenClient.ObjectVarrayApi.ObjectVarrayServiceGetVirtualArrays(ctx)
		if vdcID := helper.ValueToPointer[string](data.VdcID); vdcID != nil {
			// add vdc id filter if present
			dsreq = dsreq.VdcId(*vdcID)
		}
		dsresp, _, err := dsreq.Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error fetching Storage Pools", err.Error())
			return
		}
		if name := helper.ValueToPointer[string](data.Name); name != nil {
			// iterate through all Storage Pools and filter by name
			for _, sp := range dsresp.Varray {
				if *sp.Name == *name {
					allSpResp = append(allSpResp, sp)
				}
			}
			if (len(allSpResp)) == 0 {
				resp.Diagnostics.AddError("Error fetching Storage Pool by name", "Storage Pool with name: "+*name+" not found.")
				return
			}
		} else {
			allSpResp = dsresp.Varray
		}
	}

	// hardcoding a response value to save into the Terraform state.
	data.ID = types.StringValue("storage_pool_datasource")
	data.StoragePools = d.updateState(allSpResp)

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read storage pool data source done")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d StoragePoolDataSource) updateState(vdcs []clientgen.ObjectVarrayServiceGetVirtualArrayResponse) []models.StoragePoolDataSourceItem {
	return helper.SliceTransform(vdcs, func(v clientgen.ObjectVarrayServiceGetVirtualArrayResponse) models.StoragePoolDataSourceItem {
		return models.StoragePoolDataSourceItem{
			Id:                   helper.TfStringNN(v.Id),
			Name:                 helper.TfStringNN(v.Name),
			Description:          helper.TfStringNN(v.Description),
			IsColdStorageEnabled: helper.TfBoolNN(v.IsColdStorageEnabled),
			NumberOfDataBlocks:   helper.TfInt32NN(v.NumberOfDataBlocks),
			NumberOfCodeBlocks:   helper.TfInt32NN(v.NumberOfCodeBlocks),
			WarningAlertAt:       helper.TfInt32NN(v.WarningAlertAt),
			ErrorAlertAt:         helper.TfInt32NN(v.ErrorAlertAt),
			CriticalAlertAt:      helper.TfInt32NN(v.CriticalAlertAt),
			Status:               helper.TfInt32NN(v.Status),
			Label:                helper.TfStringNN(v.Label),
			DriveTechnology:      helper.TfStringNN(v.DriveTechnology),
		}
	})
}
