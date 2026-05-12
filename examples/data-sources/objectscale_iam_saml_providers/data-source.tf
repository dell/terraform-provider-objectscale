# Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.

# List all SAML IdPs in a namespace, with optional pagination.
data "objectscale_iam_saml_providers" "all" {
  namespace = "ns1"
  max_items = 100
}

output "provider_count" {
  value = length(data.objectscale_iam_saml_providers.all.providers)
}

output "provider_arns" {
  value = [for p in data.objectscale_iam_saml_providers.all.providers : p.arn]
}
