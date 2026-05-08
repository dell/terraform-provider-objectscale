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
	"encoding/xml"
	"fmt"
	"net/url"
	"strings"
)

// URLEncodeMetadata percent-encodes a SAML metadata XML document for safe
// inclusion in a query string (`SAMLMetadataDocument=...`).
func URLEncodeMetadata(xml string) string {
	return url.QueryEscape(xml)
}

// SPMetadata is the parsed view of an ObjectScale Service Provider's
// EntityDescriptor SAML metadata document, exposing the fields needed for
// IdP-side onboarding.
type SPMetadata struct {
	EntityID             string
	ACSURL               string
	AuthnRequestsSigned  bool
	WantAssertionsSigned bool
	SigningCertificate   string
	NameIDFormats        []string
}

// internal raw EntityDescriptor structure used only for parsing.
type entityDescriptor struct {
	XMLName        xml.Name        `xml:"EntityDescriptor"`
	EntityID       string          `xml:"entityID,attr"`
	SPSSODescriptors []spssoDescriptor `xml:"SPSSODescriptor"`
}

type spssoDescriptor struct {
	AuthnRequestsSigned  string                  `xml:"AuthnRequestsSigned,attr"`
	WantAssertionsSigned string                  `xml:"WantAssertionsSigned,attr"`
	KeyDescriptor        []keyDescriptor         `xml:"KeyDescriptor"`
	NameIDFormats        []string                `xml:"NameIDFormat"`
	ACS                  []assertionConsumerSvc  `xml:"AssertionConsumerService"`
}

type keyDescriptor struct {
	Use      string  `xml:"use,attr"`
	KeyInfo  keyInfo `xml:"KeyInfo"`
}

type keyInfo struct {
	X509Data x509Data `xml:"X509Data"`
}

type x509Data struct {
	X509Certificate string `xml:"X509Certificate"`
}

type assertionConsumerSvc struct {
	Location string `xml:"Location,attr"`
	Binding  string `xml:"Binding,attr"`
}

// ParseSPMetadata parses an EntityDescriptor and returns the relevant fields.
//
// Missing optional fields are returned as zero values. Partial / malformed
// XML returns a typed error but never panics.
func ParseSPMetadata(rawXML string) (SPMetadata, error) {
	if strings.TrimSpace(rawXML) == "" {
		return SPMetadata{}, fmt.Errorf("empty SP metadata document")
	}
	var ent entityDescriptor
	if err := xml.Unmarshal([]byte(rawXML), &ent); err != nil {
		return SPMetadata{}, fmt.Errorf("parse SP metadata: %w", err)
	}
	out := SPMetadata{EntityID: ent.EntityID}

	if len(ent.SPSSODescriptors) == 0 {
		return out, nil
	}
	d := ent.SPSSODescriptors[0]
	out.AuthnRequestsSigned = parseBoolAttr(d.AuthnRequestsSigned)
	out.WantAssertionsSigned = parseBoolAttr(d.WantAssertionsSigned)
	out.NameIDFormats = append([]string(nil), d.NameIDFormats...)

	// Pick first signing certificate (use="signing"), else first KeyDescriptor.
	for _, kd := range d.KeyDescriptor {
		if kd.Use == "" || kd.Use == "signing" {
			out.SigningCertificate = strings.TrimSpace(kd.KeyInfo.X509Data.X509Certificate)
			break
		}
	}
	if out.SigningCertificate == "" && len(d.KeyDescriptor) > 0 {
		out.SigningCertificate = strings.TrimSpace(d.KeyDescriptor[0].KeyInfo.X509Data.X509Certificate)
	}

	// Pick first ACS entry.
	if len(d.ACS) > 0 {
		out.ACSURL = d.ACS[0].Location
	}
	return out, nil
}

func parseBoolAttr(v string) bool {
	switch strings.ToLower(strings.TrimSpace(v)) {
	case "true", "1":
		return true
	default:
		return false
	}
}
