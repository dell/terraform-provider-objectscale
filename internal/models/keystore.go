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

// VDCCertificateDataSourceModel is the tfsdk model for the VDC certificate data source.
type VDCCertificateDataSourceModel struct {
	ID               types.String `tfsdk:"id"`
	CertificateChain types.String `tfsdk:"certificate_chain"`
}

// ObjectCertificateDataSourceModel is the tfsdk model for the Object certificate data source.
type ObjectCertificateDataSourceModel struct {
	ID               types.String `tfsdk:"id"`
	CertificateChain types.String `tfsdk:"certificate_chain"`
}

// VDCCertificateResourceModel is the tfsdk model for the VDC certificate resource.
type VDCCertificateResourceModel struct {
	ID                      types.String `tfsdk:"id"`
	PrivateKey              types.String `tfsdk:"private_key"`
	CertificateChain        types.String `tfsdk:"certificate_chain"`
	CurrentCertificateChain types.String `tfsdk:"current_certificate_chain"`
}

// ObjectCertificateResourceModel is the tfsdk model for the Object certificate resource.
type ObjectCertificateResourceModel struct {
	ID                      types.String `tfsdk:"id"`
	PrivateKey              types.String `tfsdk:"private_key"`
	CertificateChain        types.String `tfsdk:"certificate_chain"`
	SystemSelfsigned        types.Bool   `tfsdk:"system_selfsigned"`
	IPAddresses             types.List   `tfsdk:"ip_addresses"`
	CurrentCertificateChain types.String `tfsdk:"current_certificate_chain"`
}

// KeystoreGetResponse represents the JSON response from GET /vdc/keystore or GET /object-cert/keystore.
type KeystoreGetResponse struct {
	Chain string `json:"chain"`
}

// KeystorePutRequest represents the JSON request body for PUT /vdc/keystore or PUT /object-cert/keystore.
type KeystorePutRequest struct {
	KeyAndCertificate *KeyAndCertificate `json:"key_and_certificate,omitempty"`
	SystemSelfsigned  *bool              `json:"system_selfsigned,omitempty"`
	IPAddresses       []string           `json:"ip_addresses,omitempty"`
}

// KeyAndCertificate is the nested structure for custom certificate upload.
type KeyAndCertificate struct {
	PrivateKey       string `json:"private_key"`
	CertificateChain string `json:"certificate_chain"`
}

// KeystoreErrorResponse represents an error response from the ObjectScale keystore API.
type KeystoreErrorResponse struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	Details     string `json:"details"`
	Retryable   bool   `json:"retryable"`
}
