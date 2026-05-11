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
	"fmt"
	"regexp"
	"strings"
)

// SAMLProviderARN represents a parsed ObjectScale SAML provider ARN of the form
//
//	urn:ecs:iam::<namespace>:saml-provider/<name>
//
// Namespace may be empty for cluster-level providers.
type SAMLProviderARN struct {
	Namespace string
	Name      string
}

// String renders the canonical ARN.
func (a SAMLProviderARN) String() string {
	return BuildSAMLProviderARN(a.Namespace, a.Name)
}

// BuildSAMLProviderARN renders a canonical ARN for the given namespace + name.
// namespace may be empty.
func BuildSAMLProviderARN(namespace, name string) string {
	return fmt.Sprintf("urn:ecs:iam::%s:saml-provider/%s", namespace, name)
}

// ParseSAMLProviderARN parses a string of the form
// "urn:ecs:iam::<ns>:saml-provider/<name>" and returns the components.
//
// The function is strict: any deviation from the prefix, the empty resource
// type segment, or a missing name yields an error.
func ParseSAMLProviderARN(arn string) (SAMLProviderARN, error) {
	if !strings.HasPrefix(arn, "urn:ecs:iam::") {
		return SAMLProviderARN{}, fmt.Errorf("invalid SAML provider ARN %q: must start with urn:ecs:iam:: prefix", arn)
	}
	rest := strings.TrimPrefix(arn, "urn:ecs:iam::")
	// rest is "<namespace>:saml-provider/<name>"
	colonIdx := strings.Index(rest, ":saml-provider/")
	if colonIdx < 0 {
		return SAMLProviderARN{}, fmt.Errorf("invalid SAML provider ARN %q: missing :saml-provider/ segment", arn)
	}
	namespace := rest[:colonIdx]
	name := rest[colonIdx+len(":saml-provider/"):]
	if name == "" {
		return SAMLProviderARN{}, fmt.Errorf("invalid SAML provider ARN %q: empty name", arn)
	}
	if strings.ContainsAny(name, "/:") {
		return SAMLProviderARN{}, fmt.Errorf("invalid SAML provider ARN %q: name contains forbidden characters", arn)
	}
	return SAMLProviderARN{Namespace: namespace, Name: name}, nil
}

// samlNameRE matches the AWS-compatible IAM identifier charset used by
// ObjectScale: letters, digits and `+=,.@_-`.
var samlNameRE = regexp.MustCompile(`^[A-Za-z0-9+=,.@_-]+$`)

// ValidateSAMLProviderName enforces the documented charset and length.
//
//   - 1..128 characters
//   - charset: A-Z, a-z, 0-9, `+`, `=`, `,`, `.`, `@`, `_`, `-`
func ValidateSAMLProviderName(name string) error {
	if name == "" {
		return fmt.Errorf("SAML provider name must not be empty")
	}
	if len(name) > 128 {
		return fmt.Errorf("SAML provider name %q exceeds 128 characters", name)
	}
	if !samlNameRE.MatchString(name) {
		return fmt.Errorf("SAML provider name %q contains invalid characters; allowed: A-Z a-z 0-9 + = , . @ _ -", name)
	}
	return nil
}

// ValidateSPDNS performs minimal validation on the Service Provider DNS host
// (a non-empty string with at least one dot or a single label that is not just
// whitespace).
func ValidateSPDNS(dns string) error {
	if strings.TrimSpace(dns) == "" {
		return fmt.Errorf("service provider dns must not be empty")
	}
	return nil
}
