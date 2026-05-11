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
	"net/url"
	"strings"
	"testing"
)

// U-13 — URL-encode SAML metadata.
func TestU13_URLEncodeMetadata(t *testing.T) {
	xml := `<EntityDescriptor xmlns="urn:oasis:names:tc:SAML:2.0:metadata" entityID="https://idp.example.com/idp"/>`
	enc := URLEncodeMetadata(xml)
	dec, err := url.QueryUnescape(enc)
	if err != nil {
		t.Fatalf("U-13: re-decode failed: %v", err)
	}
	if dec != xml {
		t.Fatalf("U-13: round-trip mismatch.\n got: %q\nwant: %q", dec, xml)
	}
	// must percent-escape angle brackets etc.
	if strings.Contains(enc, "<") || strings.Contains(enc, ">") {
		t.Fatalf("U-13: encoded form still contains raw < or >: %q", enc)
	}
}

const sampleSPMetadata = `<?xml version="1.0" encoding="UTF-8"?>
<EntityDescriptor xmlns="urn:oasis:names:tc:SAML:2.0:metadata"
                  entityID="urn:objectscale:sp:test">
  <SPSSODescriptor protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol"
                   AuthnRequestsSigned="true"
                   WantAssertionsSigned="true">
    <KeyDescriptor use="signing">
      <KeyInfo xmlns="http://www.w3.org/2000/09/xmldsig#">
        <X509Data>
          <X509Certificate>MIIDdDCCAlygAwIBAgIEdummycert</X509Certificate>
        </X509Data>
      </KeyInfo>
    </KeyDescriptor>
    <NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress</NameIDFormat>
    <NameIDFormat>urn:oasis:names:tc:SAML:2.0:nameid-format:persistent</NameIDFormat>
    <AssertionConsumerService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST"
                              Location="https://objectscale.example.com/saml/acs"
                              index="0"/>
  </SPSSODescriptor>
</EntityDescriptor>`

// U-23 — Parse SP metadata XML (full).
func TestU23_ParseSPMetadata_Full(t *testing.T) {
	got, err := ParseSPMetadata(sampleSPMetadata)
	if err != nil {
		t.Fatalf("U-23: unexpected error: %v", err)
	}
	if got.EntityID != "urn:objectscale:sp:test" {
		t.Errorf("U-23: entity_id = %q, want urn:objectscale:sp:test", got.EntityID)
	}
	if got.ACSURL != "https://objectscale.example.com/saml/acs" {
		t.Errorf("U-23: acs_url = %q, want .../saml/acs", got.ACSURL)
	}
	if !got.AuthnRequestsSigned {
		t.Errorf("U-23: AuthnRequestsSigned = false, want true")
	}
	if !got.WantAssertionsSigned {
		t.Errorf("U-23: WantAssertionsSigned = false, want true")
	}
	if got.SigningCertificate != "MIIDdDCCAlygAwIBAgIEdummycert" {
		t.Errorf("U-23: signing_certificate = %q", got.SigningCertificate)
	}
	if len(got.NameIDFormats) != 2 {
		t.Fatalf("U-23: name_id_formats len = %d, want 2", len(got.NameIDFormats))
	}
}

// U-27 — Parse SP metadata: partial XML, no panic, no error, zero fields.
func TestU27_ParseSPMetadata_Partial(t *testing.T) {
	cases := []string{
		`<EntityDescriptor entityID="urn:test"/>`,
		`<EntityDescriptor entityID="urn:test"><SPSSODescriptor></SPSSODescriptor></EntityDescriptor>`,
	}
	for _, c := range cases {
		got, err := ParseSPMetadata(c)
		if err != nil {
			t.Errorf("U-27: unexpected error for %q: %v", c, err)
			continue
		}
		if got.EntityID != "urn:test" {
			t.Errorf("U-27: entity_id = %q for %q", got.EntityID, c)
		}
	}
}

// U-27b — Empty XML returns error (not panic).
func TestU27b_ParseSPMetadata_Empty(t *testing.T) {
	if _, err := ParseSPMetadata(""); err == nil {
		t.Fatalf("U-27b: expected error for empty XML")
	}
	if _, err := ParseSPMetadata("not xml at all"); err == nil {
		t.Fatalf("U-27b: expected error for non-XML")
	}
}
