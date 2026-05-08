# Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.
#
# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://mozilla.org/MPL/2.0/

title: "objectscale_iam_service_provider resource"
linkTitle: "objectscale_iam_service_provider"
page_title: "objectscale_iam_service_provider Resource - terraform-provider-objectscale"
subcategory: "Identity & Access Management (IAM)"
description: |-
  Manages the ObjectScale SAML Service Provider configuration (singleton).
---

# objectscale_iam_service_provider (Resource)

Manages the ObjectScale SAML Service Provider configuration. **Singleton** —
exactly one Service Provider per cluster.

## Example Usage

```hcl
resource "objectscale_iam_service_provider" "sp" {
  dns           = "objectscale.example.com"
  java_keystore = filebase64("${path.module}/keystore.jks")
  key_alias     = "saml"
  key_password  = var.keystore_password
}
```

## Schema

### Required

- `dns` (String) — SP base URL for SAML ACS.
- `java_keystore` (String, **Sensitive**) — Base64-encoded JKS.
- `key_alias` (String) — Keystore entry alias.
- `key_password` (String, **Sensitive**) — Keystore password.

### Read-Only

- `id` (String) — Always `objectscale-sp` (singleton).
- `uuid` (String) — Entity ID component.
- `unique_id` (String) — Keystore unique ID.
- `etag` (String) — Optimistic concurrency tag.
- `create_time` (String) — ISO 8601 creation timestamp.
- `last_modified` (String) — ISO 8601 last-modified timestamp.

## Import

```bash
terraform import objectscale_iam_service_provider.sp objectscale-sp
```
