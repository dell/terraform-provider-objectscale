# Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.

title: "objectscale_iam_saml_providers data source"
page_title: "objectscale_iam_saml_providers Data Source - terraform-provider-objectscale"
subcategory: "Identity & Access Management (IAM)"
description: |-
  Lists ObjectScale IAM SAML Identity Providers, with optional pagination.
---

# objectscale_iam_saml_providers (Data Source)

Lists all ObjectScale IAM SAML Identity Providers, with optional pagination.
The data source auto-paginates until either the API reports `IsTruncated=false`
or the caller-specified `max_items` cap is reached.

## Example Usage

```hcl
data "objectscale_iam_saml_providers" "all" {
  namespace = "ns1"
  max_items = 100
}
```

## Schema

### Optional

- `max_items` (Number) — Maximum providers to return.
- `namespace` (String) — `x-emc-namespace` header.

### Read-Only

- `id` (String)
- `providers` (List of Object) — Each entry has `arn`, `name`, `create_date`, `valid_until`.
- `is_truncated` (Boolean) — True if the last API page was truncated.
- `marker` (String) — API continuation token of the last page.
