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

package models

import "github.com/hashicorp/terraform-plugin-framework/types"

// IAMServiceProviderResourceModel is the state model for
// `objectscale_iam_service_provider` resource (singleton).
type IAMServiceProviderResourceModel struct {
	ID           types.String `tfsdk:"id"`
	DNS          types.String `tfsdk:"dns"`
	JavaKeystore types.String `tfsdk:"java_keystore"`
	KeyAlias     types.String `tfsdk:"key_alias"`
	KeyPassword  types.String `tfsdk:"key_password"`
	UUID         types.String `tfsdk:"uuid"`
	UniqueID     types.String `tfsdk:"unique_id"`
	Etag         types.String `tfsdk:"etag"`
	CreateTime   types.String `tfsdk:"create_time"`
	LastModified types.String `tfsdk:"last_modified"`
}

// IAMServiceProviderDataSourceModel mirrors the resource model with all
// attributes computed.
type IAMServiceProviderDataSourceModel = IAMServiceProviderResourceModel

// IAMServiceProviderMetadataDataSourceModel is the state model for
// `objectscale_iam_service_provider_metadata` datasource.
type IAMServiceProviderMetadataDataSourceModel struct {
	ID                   types.String `tfsdk:"id"`
	MetadataXML          types.String `tfsdk:"metadata_xml"`
	EntityID             types.String `tfsdk:"entity_id"`
	ACSURL               types.String `tfsdk:"acs_url"`
	AuthnRequestsSigned  types.Bool   `tfsdk:"authn_requests_signed"`
	WantAssertionsSigned types.Bool   `tfsdk:"want_assertions_signed"`
	SigningCertificate   types.String `tfsdk:"signing_certificate"`
	NameIDFormats        types.List   `tfsdk:"name_id_formats"`
}
