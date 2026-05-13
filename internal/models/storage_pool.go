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

package models

import "github.com/hashicorp/terraform-plugin-framework/types"

type StoragePoolDataSourceModel struct {
	ID           types.String                `tfsdk:"id"`
	VdcID        types.String                `tfsdk:"vdc_id"`
	Name         types.String                `tfsdk:"name"`
	StoragePools []StoragePoolDataSourceItem `tfsdk:"storage_pools"`
}

type StoragePoolDataSourceItem struct {
	Id                   types.String `tfsdk:"id"`
	Name                 types.String `tfsdk:"name"`
	Description          types.String `tfsdk:"description"`
	IsColdStorageEnabled types.Bool   `tfsdk:"is_cold_storage_enabled"`
	NumberOfDataBlocks   types.Int32  `tfsdk:"number_of_data_blocks"`
	NumberOfCodeBlocks   types.Int32  `tfsdk:"number_of_code_blocks"`
	WarningAlertAt       types.Int32  `tfsdk:"warning_alert_at"`
	ErrorAlertAt         types.Int32  `tfsdk:"error_alert_at"`
	CriticalAlertAt      types.Int32  `tfsdk:"critical_alert_at"`
	Status               types.Int32  `tfsdk:"status"`
	Label                types.String `tfsdk:"label"`
	DriveTechnology      types.String `tfsdk:"drive_technology"`
}
