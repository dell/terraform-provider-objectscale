# Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.

# Configure ObjectScale itself as a SAML Service Provider. Singleton: at most
# one resource of this type per cluster.
resource "objectscale_iam_service_provider" "sp" {
  dns           = "objectscale.example.com"
  java_keystore = filebase64("${path.module}/keystore.jks")
  key_alias     = "saml"
  key_password  = var.keystore_password
}

variable "keystore_password" {
  type      = string
  sensitive = true
}
