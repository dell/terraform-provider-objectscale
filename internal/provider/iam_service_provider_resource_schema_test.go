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

package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func serviceProviderSchema(t *testing.T) rschema.Schema {
	t.Helper()
	r := NewIAMServiceProviderResource()
	resp := &resource.SchemaResponse{}
	r.(*IAMServiceProviderResource).Schema(context.Background(), resource.SchemaRequest{}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("schema errors: %v", resp.Diagnostics)
	}
	return resp.Schema
}

// U-26 — Sensitive: java_keystore and key_password are sensitive.
func TestU26_ServiceProviderSchema_SensitiveFields(t *testing.T) {
	s := serviceProviderSchema(t)

	for _, name := range []string{"java_keystore", "key_password"} {
		attr, ok := s.Attributes[name].(rschema.StringAttribute)
		if !ok {
			t.Fatalf("U-26: attribute %q missing or not StringAttribute", name)
		}
		if !attr.IsSensitive() {
			t.Errorf("U-26: attribute %q is not Sensitive", name)
		}
		if !attr.IsRequired() {
			t.Errorf("U-26: attribute %q is not Required", name)
		}
	}
}

// Sanity — `dns` and `key_alias` are required but not sensitive.
func TestU26b_ServiceProviderSchema_NonSensitiveRequired(t *testing.T) {
	s := serviceProviderSchema(t)
	for _, name := range []string{"dns", "key_alias"} {
		attr, ok := s.Attributes[name].(rschema.StringAttribute)
		if !ok {
			t.Fatalf("attribute %q missing", name)
		}
		if !attr.IsRequired() {
			t.Errorf("attribute %q not Required", name)
		}
		if attr.IsSensitive() {
			t.Errorf("attribute %q unexpectedly Sensitive", name)
		}
	}
}

// Sanity — computed fields exist.
func TestU26c_ServiceProviderSchema_ComputedFields(t *testing.T) {
	s := serviceProviderSchema(t)
	for _, name := range []string{"id", "uuid", "unique_id", "etag", "create_time", "last_modified"} {
		attr, ok := s.Attributes[name]
		if !ok {
			t.Fatalf("attribute %q missing", name)
		}
		if !attr.IsComputed() {
			t.Errorf("attribute %q not Computed", name)
		}
	}
}
