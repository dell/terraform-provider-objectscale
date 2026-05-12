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
	"terraform-provider-objectscale/internal/helper"
	"terraform-provider-objectscale/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VDCCertificateResource{}
var _ resource.ResourceWithImportState = &VDCCertificateResource{}

func NewVDCCertificateResource() resource.Resource {
	return &VDCCertificateResource{}
}

// VDCCertificateResource manages the VDC management-plane TLS certificate.
type VDCCertificateResource struct {
	resourceProviderConfig
}

func (r *VDCCertificateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vdc_certificate"
}

func (r *VDCCertificateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This resource manages the VDC management-plane TLS certificate on Dell ObjectScale. Certificates always exist on a VDC; this resource replaces (not creates/deletes) the active certificate.",
		MarkdownDescription: "This resource manages the VDC management-plane TLS certificate on Dell ObjectScale. Certificates always exist on a VDC; this resource replaces (not creates/deletes) the active certificate.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Identifier for the VDC certificate resource.",
				MarkdownDescription: "Identifier for the VDC certificate resource.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"private_key": schema.StringAttribute{
				Description:         "Private key in PEM PKCS#1 format (RSA PRIVATE KEY). PKCS#8 keys are not supported; convert with: openssl rsa -in key.pem -out key-pkcs1.pem",
				MarkdownDescription: "Private key in PEM PKCS#1 format (`RSA PRIVATE KEY`). PKCS#8 keys are not supported; convert with: `openssl rsa -in key.pem -out key-pkcs1.pem`",
				Required:            true,
				Sensitive:           true,
			},
			"certificate_chain": schema.StringAttribute{
				Description:         "Certificate chain in PEM format. Must contain at least one CERTIFICATE block.",
				MarkdownDescription: "Certificate chain in PEM format. Must contain at least one CERTIFICATE block.",
				Required:            true,
			},
			"current_certificate_chain": schema.StringAttribute{
				Description:         "The currently active certificate chain as read from the ObjectScale API. May be stale for up to 1 hour after a VDC certificate update.",
				MarkdownDescription: "The currently active certificate chain as read from the ObjectScale API. May be stale for up to 1 hour after a VDC certificate update.",
				Computed:            true,
			},
		},
	}
}

func (r *VDCCertificateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.VDCCertificateResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.applyVDCCertificate(ctx, &plan, &resp.Diagnostics, &resp.State)
}

func (r *VDCCertificateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.VDCCertificateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	chain, err := GetVDCKeystore(ctx, r.client)
	if err != nil {
		resp.Diagnostics.AddError("Error reading VDC certificate", err.Error())
		return
	}

	state.CurrentCertificateChain = types.StringValue(helper.NormalizeLineEndings(chain))

	tflog.Trace(ctx, "read VDC certificate resource")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *VDCCertificateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan models.VDCCertificateResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.applyVDCCertificate(ctx, &plan, &resp.Diagnostics, &resp.State)
}

func (r *VDCCertificateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Warn(ctx, "VDC certificate cannot be deleted. Removing from Terraform state only. The certificate remains active on ObjectScale.")
	// No API call — VDC always has a certificate. Just remove from state.
}

func (r *VDCCertificateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	chain, err := GetVDCKeystore(ctx, r.client)
	if err != nil {
		resp.Diagnostics.AddError("Error importing VDC certificate", err.Error())
		return
	}

	normalizedChain := helper.NormalizeLineEndings(chain)

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), "vdc_certificate")...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("private_key"), "")...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("certificate_chain"), "")...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("current_certificate_chain"), normalizedChain)...)
}

// applyVDCCertificate is shared between Create and Update.
func (r *VDCCertificateResource) applyVDCCertificate(ctx context.Context, plan *models.VDCCertificateResourceModel, diagnostics *diag.Diagnostics, state *tfsdk.State) {
	// Validate PEM private key
	privateKeyRaw := plan.PrivateKey.ValueString()
	if err := helper.ValidatePEMPrivateKey(privateKeyRaw); err != nil {
		diagnostics.AddError("Invalid Private Key", err.Error())
		return
	}

	// Validate PEM certificate chain
	certChainRaw := plan.CertificateChain.ValueString()
	if err := helper.ValidatePEMCertificate(certChainRaw); err != nil {
		diagnostics.AddError("Invalid Certificate Chain", err.Error())
		return
	}

	// Validate private key format (must be PKCS#1)
	normalizedKey, err := helper.ValidateAndNormalizePrivateKey(privateKeyRaw)
	if err != nil {
		diagnostics.AddError("Invalid Private Key Format", err.Error())
		return
	}

	normalizedCertChain := helper.NormalizeLineEndings(certChainRaw)

	// Idempotency check: GET current chain and compare
	currentChain, err := GetVDCKeystore(ctx, r.client)
	if err != nil {
		diagnostics.AddError("Error reading current VDC certificate for idempotency check", err.Error())
		return
	}

	if helper.CompareCertificateChains(currentChain, normalizedCertChain) {
		tflog.Info(ctx, "VDC certificate chain unchanged, skipping PUT")
	} else {
		// Chains differ — execute PUT
		if err := PutVDCKeystore(ctx, r.client, normalizedKey, normalizedCertChain); err != nil {
			diagnostics.AddError("Error updating VDC certificate", err.Error())
			return
		}
		diagnostics.AddWarning("VDC Certificate Propagation Delay",
			"VDC certificate propagation may take up to 1 hour. The current_certificate_chain attribute may show the old certificate during this period.")
	}

	// Set state
	plan.ID = types.StringValue("vdc_certificate")
	plan.CurrentCertificateChain = types.StringValue(helper.NormalizeLineEndings(currentChain))

	diagnostics.Append(state.Set(ctx, plan)...)
}
