# Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.

# Register an external SAML IdP (e.g. Okta / Azure AD / ADFS) with ObjectScale.
resource "objectscale_iam_saml_provider" "corp" {
  name      = "corp-saml"
  namespace = "ns1"

  # Provide the IdP's SAML metadata XML (typically downloaded from the IdP admin
  # console). Updates trigger an in-place rotation; changing `name` or
  # `namespace` triggers destroy + recreate.
  saml_metadata_document = file("${path.module}/idp-metadata.xml")
}

output "saml_provider_arn" {
  value = objectscale_iam_saml_provider.corp.arn
}
