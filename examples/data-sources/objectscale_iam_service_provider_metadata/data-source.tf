# Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.

# Read and parse the SAML SP metadata XML for IdP-side onboarding.
data "objectscale_iam_service_provider_metadata" "md" {}

output "entity_id" {
  value = data.objectscale_iam_service_provider_metadata.md.entity_id
}

output "acs_url" {
  value = data.objectscale_iam_service_provider_metadata.md.acs_url
}

output "signing_certificate" {
  value = data.objectscale_iam_service_provider_metadata.md.signing_certificate
}

output "name_id_formats" {
  value = data.objectscale_iam_service_provider_metadata.md.name_id_formats
}

# Save the raw EntityDescriptor XML to disk so it can be uploaded to the
# external IdP admin console.
resource "local_file" "sp_metadata" {
  content  = data.objectscale_iam_service_provider_metadata.md.metadata_xml
  filename = "${path.module}/objectscale-sp-metadata.xml"
}
