# Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.

# Read a single SAML IdP by ARN.
data "objectscale_iam_saml_provider" "corp" {
  saml_provider_arn = "urn:ecs:iam::ns1:saml-provider/corp-saml"
  namespace         = "ns1"
}

output "valid_until" {
  value = data.objectscale_iam_saml_provider.corp.valid_until
}
