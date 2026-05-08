# Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.

title: "objectscale_iam_service_provider_metadata data source"
page_title: "objectscale_iam_service_provider_metadata Data Source - terraform-provider-objectscale"
subcategory: "Identity & Access Management (IAM)"
description: |-
  Reads ObjectScale SAML SP metadata XML and parses key fields for IdP onboarding.
---

# objectscale_iam_service_provider_metadata (Data Source)

Reads the ObjectScale SAML SP metadata document (`/ecs-service-provider/metadata`)
and parses the EntityDescriptor XML into structured fields. Use this datasource
to extract the SP `entity_id`, ACS URL, signing certificate, and supported
`NameIDFormat` values when onboarding the ObjectScale SP with an external IdP.

## Example Usage

```hcl
data "objectscale_iam_service_provider_metadata" "md" {}

output "acs_url" {
  value = data.objectscale_iam_service_provider_metadata.md.acs_url
}

output "signing_certificate" {
  value = data.objectscale_iam_service_provider_metadata.md.signing_certificate
}
```

## Schema

### Read-Only

- `id` (String)
- `metadata_xml` (String) — Raw EntityDescriptor XML.
- `entity_id` (String)
- `acs_url` (String)
- `authn_requests_signed` (Boolean)
- `want_assertions_signed` (Boolean)
- `signing_certificate` (String) — Base64 X.509 certificate.
- `name_id_formats` (List of String) — Supported `NameIDFormat` values.
