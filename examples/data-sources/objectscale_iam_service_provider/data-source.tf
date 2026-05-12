# Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.

# Read the singleton SAML Service Provider configuration.
data "objectscale_iam_service_provider" "sp" {}

output "sp_dns" {
  value = data.objectscale_iam_service_provider.sp.dns
}

output "sp_uuid" {
  value = data.objectscale_iam_service_provider.sp.uuid
}
