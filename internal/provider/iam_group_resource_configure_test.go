// iam_group_resource_configure_test.go
package provider

import (
	"context"
	"testing"

	"terraform-provider-objectscale/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func TestAccIAMGroupResource_Configure_InvalidType(t *testing.T) {
	ctx := context.Background()

	r := &IAMGroupResource{}

	req := resource.ConfigureRequest{
		ProviderData: 12345, // wrong type to trigger the !ok branch
	}
	resp := &resource.ConfigureResponse{
		Diagnostics: diag.Diagnostics{},
	}

	r.Configure(ctx, req, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatalf("expected diagnostics error when ProviderData type is invalid, got none")
	}

	got := resp.Diagnostics[0].Summary() // ✅ Summary is a method in your version
	want := "Unexpected Resource Configure Type"
	if got != want {
		t.Errorf("unexpected diagnostic summary: got %q, want %q", got, want)
	}
}

func TestAccIAMGroupResource_Configure_NilProviderData_NoPanic(t *testing.T) {
	ctx := context.Background()

	r := &IAMGroupResource{}
	req := resource.ConfigureRequest{ProviderData: nil}
	resp := &resource.ConfigureResponse{Diagnostics: diag.Diagnostics{}}

	r.Configure(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("did not expect error when ProviderData is nil")
	}
}

func TestAccIAMGroupResource_Configure_ValidType_SetsClient(t *testing.T) {
	ctx := context.Background()

	r := &IAMGroupResource{}

	fake := &client.Client{} // minimal acceptable object
	req := resource.ConfigureRequest{ProviderData: fake}
	resp := &resource.ConfigureResponse{Diagnostics: diag.Diagnostics{}}

	r.Configure(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected diagnostics: %+v", resp.Diagnostics)
	}
	if r.client != fake {
		t.Fatalf("expected resource client to be set to provider client")
	}
}
