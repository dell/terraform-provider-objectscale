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

package helper

import (
	"strings"
	"testing"
)

// U-01 — Parse valid ARN.
func TestU01_ParseSAMLProviderARN_Valid(t *testing.T) {
	got, err := ParseSAMLProviderARN("urn:ecs:iam::ns:saml-provider/TestProvider")
	if err != nil {
		t.Fatalf("U-01: unexpected error: %v", err)
	}
	if got.Namespace != "ns" || got.Name != "TestProvider" {
		t.Fatalf("U-01: got %+v, want namespace=ns name=TestProvider", got)
	}
}

// U-01b — Round-trip via BuildSAMLProviderARN.
func TestU01b_BuildSAMLProviderARN_RoundTrip(t *testing.T) {
	in := SAMLProviderARN{Namespace: "n1", Name: "Foo"}
	got := in.String()
	want := "urn:ecs:iam::n1:saml-provider/Foo"
	if got != want {
		t.Fatalf("U-01b: got %q want %q", got, want)
	}
	parsed, err := ParseSAMLProviderARN(got)
	if err != nil {
		t.Fatalf("U-01b: round-trip parse error: %v", err)
	}
	if parsed != in {
		t.Fatalf("U-01b: got %+v want %+v", parsed, in)
	}
}

// U-02 — Parse invalid ARN.
func TestU02_ParseSAMLProviderARN_Invalid(t *testing.T) {
	cases := []string{
		"invalid-arn",
		"arn:aws:iam::ns:saml-provider/Foo",
		"urn:ecs:iam::ns:user/Foo",
		"urn:ecs:iam::ns:saml-provider/",
		"urn:ecs:iam::ns:saml-provider/foo/bar",
	}
	for _, c := range cases {
		if _, err := ParseSAMLProviderARN(c); err == nil {
			t.Errorf("U-02: expected error for %q", c)
		}
	}
}

// U-14 — Validate name (valid).
func TestU14_ValidateSAMLProviderName_Valid(t *testing.T) {
	for _, name := range []string{"my-provider", "Corp_Idp", "okta.example", "a"} {
		if err := ValidateSAMLProviderName(name); err != nil {
			t.Errorf("U-14: unexpected error for %q: %v", name, err)
		}
	}
}

// U-15 — Validate name (empty).
func TestU15_ValidateSAMLProviderName_Empty(t *testing.T) {
	if err := ValidateSAMLProviderName(""); err == nil {
		t.Fatalf("U-15: expected error for empty name")
	}
}

// U-16 — Validate name (special characters).
func TestU16_ValidateSAMLProviderName_Special(t *testing.T) {
	cases := []string{"my provider!", "foo/bar", "foo:bar", "foo bar"}
	for _, c := range cases {
		if err := ValidateSAMLProviderName(c); err == nil {
			t.Errorf("U-16: expected error for %q", c)
		}
	}
}

// U-16b — name too long.
func TestU16b_ValidateSAMLProviderName_TooLong(t *testing.T) {
	long := strings.Repeat("a", 129)
	if err := ValidateSAMLProviderName(long); err == nil {
		t.Fatalf("U-16b: expected error for 129-char name")
	}
}

// U-24 — Validate SP DNS (valid).
func TestU24_ValidateSPDNS_Valid(t *testing.T) {
	if err := ValidateSPDNS("objectscale.example.com"); err != nil {
		t.Fatalf("U-24: unexpected error: %v", err)
	}
}

// U-25 — Validate SP DNS (empty).
func TestU25_ValidateSPDNS_Empty(t *testing.T) {
	if err := ValidateSPDNS(""); err == nil {
		t.Fatalf("U-25: expected error for empty DNS")
	}
	if err := ValidateSPDNS("   "); err == nil {
		t.Fatalf("U-25: expected error for whitespace DNS")
	}
}
