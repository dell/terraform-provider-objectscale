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

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ObjectUserResourceModel struct {
	Tags      types.List   `tfsdk:"tags"`
	Name      types.String `tfsdk:"name"`
	Namespace types.String `tfsdk:"namespace"`
	Locked    types.Bool   `tfsdk:"locked"`
	Created   types.String `tfsdk:"created"`
	Id        types.String `tfsdk:"id"`
}

type ObjectUserTags struct {
	// A single-valued attribute indicating the user's IDP domain
	Name types.String `tfsdk:"name"`
	// Attributes
	Value types.String `tfsdk:"value"`
}

type ObjectUserDatasourceModel struct {
	Name      types.String `tfsdk:"name"`
	Namespace types.String `tfsdk:"namespace"`
	Id        types.String `tfsdk:"id"`
	Tag       types.String `tfsdk:"tag"`
	Value     types.String `tfsdk:"value"`
	Users     []ObjectUser `tfsdk:"users"`
}
type ObjectUser struct {
	Tags      types.List          `tfsdk:"tags"`
	Name      types.String        `tfsdk:"name"`
	Namespace types.String        `tfsdk:"namespace"`
	Locked    types.Bool          `tfsdk:"locked"`
	Created   types.String        `tfsdk:"created"`
	Id        types.String        `tfsdk:"id"`
	SecretKey ObjectUserAccessKey `tfsdk:"secret_keys"`
}

type ObjectUserAccessKey struct {
	SecretKey1Id        types.String `tfsdk:"secret_key_1_id"`
	SecretKey1          types.String `tfsdk:"secret_key_1"`
	SecretKey1Exist     types.Bool   `tfsdk:"secret_key_1_exist"`
	KeyTimestamp1       types.String `tfsdk:"key_timestamp_1"`
	KeyExpiryTimestamp1 types.String `tfsdk:"key_expiry_timestamp_1"`
	SecretKey2Id        types.String `tfsdk:"secret_key_2_id"`
	SecretKey2          types.String `tfsdk:"secret_key_2"`
	SecretKey2Exist     types.Bool   `tfsdk:"secret_key_2_exist"`
	KeyTimestamp2       types.String `tfsdk:"key_timestamp_2"`
	KeyExpiryTimestamp2 types.String `tfsdk:"key_expiry_timestamp_2"`
}

type ObjectUserSecretKeyResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	SecretKey          types.String `tfsdk:"secret_key"`
	KeyTimestamp       types.String `tfsdk:"key_timestamp"`
	KeyExpiryTimestamp types.String `tfsdk:"key_expiry_timestamp"`
	UserName           types.String `tfsdk:"username"`
	Namespace          types.String `tfsdk:"namespace"`
	ExpiryInMins       types.String `tfsdk:"expiry_in_mins"`
}
