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

// IAMSAMLProviderResourceModel is the state model for
// `objectscale_iam_saml_provider` resource.
type IAMSAMLProviderResourceModel struct {
	ID                   types.String `tfsdk:"id"`
	Name                 types.String `tfsdk:"name"`
	SAMLMetadataDocument types.String `tfsdk:"saml_metadata_document"`
	Namespace            types.String `tfsdk:"namespace"`
	Arn                  types.String `tfsdk:"arn"`
	CreateDate           types.String `tfsdk:"create_date"`
	ValidUntil           types.String `tfsdk:"valid_until"`
}

// IAMSAMLProvider is one entry in the providers list.
type IAMSAMLProvider struct {
	Arn                  types.String `tfsdk:"arn"`
	Name                 types.String `tfsdk:"name"`
	SAMLMetadataDocument types.String `tfsdk:"saml_metadata_document"`
	CreateDate           types.String `tfsdk:"create_date"`
	ValidUntil           types.String `tfsdk:"valid_until"`
}

// IAMSAMLProviderDataSourceModel is the state model for
// `objectscale_iam_saml_provider` datasource.
// When saml_provider_arn is provided, returns a single provider.
// When saml_provider_arn is not provided, returns a list of providers.
type IAMSAMLProviderDataSourceModel struct {
	ID              types.String      `tfsdk:"id"`
	SAMLProviderArn types.String      `tfsdk:"saml_provider_arn"`
	Namespace       types.String      `tfsdk:"namespace"`
	Providers       []IAMSAMLProvider `tfsdk:"providers"`
}
