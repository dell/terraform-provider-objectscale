# Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.
#
# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://mozilla.org/MPL/2.0/

title: "objectscale_iam_saml_provider data source"
linkTitle: "objectscale_iam_saml_provider"
page_title: "objectscale_iam_saml_provider Data Source - terraform-provider-objectscale"
subcategory: "Identity & Access Management (IAM)"
description: |-
  Reads a single ObjectScale IAM SAML Identity Provider by ARN.
---

# objectscale_iam_saml_provider (Data Source)

Reads a single ObjectScale IAM SAML Identity Provider by ARN.

## Example Usage

```hcl
data "objectscale_iam_saml_provider" "corp" {
  saml_provider_arn = "urn:ecs:iam::ns1:saml-provider/corp-saml"
  namespace         = "ns1"
}
```

## Schema

### Required

- `saml_provider_arn` (String) — Lookup key.

### Optional

- `namespace` (String) — `x-emc-namespace` header (management user only).

### Read-Only

- `id` (String)
- `name` (String) — Parsed from the ARN.
- `saml_metadata_document` (String)
- `create_date` (String) — ISO 8601 timestamp.
- `valid_until` (String) — ISO 8601 expiry.
- `request_id` (String) — Tracking ID for the API request.
