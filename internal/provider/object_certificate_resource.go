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

	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ObjectCertificateResource{}
var _ resource.ResourceWithImportState = &ObjectCertificateResource{}
var _ resource.ResourceWithConfigValidators = &ObjectCertificateResource{}

func NewObjectCertificateResource() resource.Resource {
	return &ObjectCertificateResource{}
}

// ObjectCertificateResource manages the Object data-plane (S3) TLS certificate.
type ObjectCertificateResource struct {
	resourceProviderConfig
}

func (r *ObjectCertificateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_object_certificate"
}

func (r *ObjectCertificateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This resource manages the Object data-plane (S3) TLS certificate on Dell ObjectScale. Supports custom certificate upload or self-signed certificate generation.",
		MarkdownDescription: "This resource manages the Object data-plane (S3) TLS certificate on Dell ObjectScale. Supports custom certificate upload or self-signed certificate generation.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Identifier for the Object certificate resource.",
				MarkdownDescription: "Identifier for the Object certificate resource.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"private_key": schema.StringAttribute{
				Description:         "Private key in PEM PKCS#1 format (RSA PRIVATE KEY). PKCS#8 keys are not supported; convert with: openssl rsa -in key.pem -out key-pkcs1.pem. Required when system_selfsigned is not set. Mutually exclusive with system_selfsigned.",
				MarkdownDescription: "Private key in PEM PKCS#1 format (`RSA PRIVATE KEY`). PKCS#8 keys are not supported; convert with: `openssl rsa -in key.pem -out key-pkcs1.pem`. Required when `system_selfsigned` is not set. Mutually exclusive with `system_selfsigned`.",
				Optional:            true,
				Sensitive:           true,
			},
			"certificate_chain": schema.StringAttribute{
				Description:         "Certificate chain in PEM format. Required when system_selfsigned is not set. Mutually exclusive with system_selfsigned.",
				MarkdownDescription: "Certificate chain in PEM format. Required when `system_selfsigned` is not set. Mutually exclusive with `system_selfsigned`.",
				Optional:            true,
			},
			"system_selfsigned": schema.BoolAttribute{
				Description:         "Generate a self-signed certificate. Mutually exclusive with private_key and certificate_chain. Forces resource replacement.",
				MarkdownDescription: "Generate a self-signed certificate. Mutually exclusive with `private_key` and `certificate_chain`. Forces resource replacement.",
				Optional:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"ip_addresses": schema.ListAttribute{
				Description:         "List of IP addresses for self-signed certificate SANs. Only used when system_selfsigned is true.",
				MarkdownDescription: "List of IP addresses for self-signed certificate SANs. Only used when `system_selfsigned` is `true`.",
				Optional:            true,
				ElementType:         types.StringType,
			},
			"current_certificate_chain": schema.StringAttribute{
				Description:         "The currently active certificate chain as read from the ObjectScale API.",
				MarkdownDescription: "The currently active certificate chain as read from the ObjectScale API.",
				Computed:            true,
			},
		},
	}
}

func (r *ObjectCertificateResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.Conflicting(
			path.MatchRoot("system_selfsigned"),
			path.MatchRoot("private_key"),
		),
		resourcevalidator.Conflicting(
			path.MatchRoot("system_selfsigned"),
			path.MatchRoot("certificate_chain"),
		),
		resourcevalidator.RequiredTogether(
			path.MatchRoot("private_key"),
			path.MatchRoot("certificate_chain"),
		),
	}
}

func (r *ObjectCertificateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.ObjectCertificateResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.applyObjectCertificate(ctx, &plan, &resp.Diagnostics, &resp.State)
}

func (r *ObjectCertificateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.ObjectCertificateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	chain, err := GetObjectCertKeystore(ctx, r.client)
	if err != nil {
		resp.Diagnostics.AddError("Error reading Object certificate", err.Error())
		return
	}

	state.CurrentCertificateChain = types.StringValue(helper.NormalizeLineEndings(chain))

	tflog.Trace(ctx, "read Object certificate resource")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ObjectCertificateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan models.ObjectCertificateResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.applyObjectCertificate(ctx, &plan, &resp.Diagnostics, &resp.State)
}

func (r *ObjectCertificateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Warn(ctx, "Object certificate cannot be deleted. Removing from Terraform state only. The certificate remains active on ObjectScale.")
	// No API call — Object cert always exists. Just remove from state.
}

func (r *ObjectCertificateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	chain, err := GetObjectCertKeystore(ctx, r.client)
	if err != nil {
		resp.Diagnostics.AddError("Error importing Object certificate", err.Error())
		return
	}

	normalizedChain := helper.NormalizeLineEndings(chain)

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), "object_certificate")...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("current_certificate_chain"), normalizedChain)...)
}

// applyObjectCertificate is shared between Create and Update.
func (r *ObjectCertificateResource) applyObjectCertificate(ctx context.Context, plan *models.ObjectCertificateResourceModel, diagnostics *diag.Diagnostics, state *tfsdk.State) {
	isSelfSigned := !plan.SystemSelfsigned.IsNull() && plan.SystemSelfsigned.ValueBool()

	if isSelfSigned {
		r.applySelfSignedCert(ctx, plan, diagnostics, state)
	} else {
		r.applyCustomCert(ctx, plan, diagnostics, state)
	}
}

// applySelfSignedCert generates a self-signed certificate.
func (r *ObjectCertificateResource) applySelfSignedCert(ctx context.Context, plan *models.ObjectCertificateResourceModel, diagnostics *diag.Diagnostics, state *tfsdk.State) {
	var ipAddresses []string
	if !plan.IPAddresses.IsNull() && !plan.IPAddresses.IsUnknown() {
		diagnostics.Append(plan.IPAddresses.ElementsAs(ctx, &ipAddresses, false)...)
		if diagnostics.HasError() {
			return
		}
	}

	chain, err := PutObjectCertSelfSigned(ctx, r.client, ipAddresses)
	if err != nil {
		diagnostics.AddError("Error generating self-signed Object certificate", err.Error())
		return
	}

	plan.ID = types.StringValue("object_certificate")
	plan.CurrentCertificateChain = types.StringValue(helper.NormalizeLineEndings(chain))

	diagnostics.Append(state.Set(ctx, plan)...)
}

// applyCustomCert uploads a custom certificate.
func (r *ObjectCertificateResource) applyCustomCert(ctx context.Context, plan *models.ObjectCertificateResourceModel, diagnostics *diag.Diagnostics, state *tfsdk.State) {
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
	currentChain, err := GetObjectCertKeystore(ctx, r.client)
	if err != nil {
		diagnostics.AddError("Error reading current Object certificate for idempotency check", err.Error())
		return
	}

	if helper.CompareCertificateChains(currentChain, normalizedCertChain) {
		tflog.Info(ctx, "Object certificate chain unchanged, skipping PUT")
	} else {
		if err := PutObjectCertKeystore(ctx, r.client, normalizedKey, normalizedCertChain); err != nil {
			diagnostics.AddError("Error updating Object certificate", err.Error())
			return
		}
	}

	// Read back current chain (Object cert propagation is immediate)
	updatedChain, err := GetObjectCertKeystore(ctx, r.client)
	if err != nil {
		// Non-fatal: use the plan chain as fallback
		updatedChain = currentChain
		tflog.Warn(ctx, "Failed to read back Object certificate after update, using previous chain")
	}

	plan.ID = types.StringValue("object_certificate")
	plan.CurrentCertificateChain = types.StringValue(helper.NormalizeLineEndings(updatedChain))

	diagnostics.Append(state.Set(ctx, plan)...)
}
