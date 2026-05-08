# Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.

title: "objectscale_iam_service_provider data source"
page_title: "objectscale_iam_service_provider Data Source - terraform-provider-objectscale"
subcategory: "Identity & Access Management (IAM)"
description: |-
  Reads the ObjectScale SAML Service Provider configuration (singleton).
---

# objectscale_iam_service_provider (Data Source)

Reads the singleton ObjectScale SAML Service Provider configuration.

## Example Usage

```hcl
data "objectscale_iam_service_provider" "sp" {}
```

## Schema

### Read-Only

- `id` (String) — Always `objectscale-sp`.
- `dns` (String)
- `uuid` (String)
- `unique_id` (String)
- `etag` (String)
- `key_alias` (String)
- `create_time` (String) — ISO 8601.
- `last_modified` (String) — ISO 8601.
- `java_keystore` (String, **Sensitive**) — Base64 JKS.
- `key_password` (String, **Sensitive**) — Keystore password.
