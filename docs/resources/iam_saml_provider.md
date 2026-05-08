# Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.
#
# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://mozilla.org/MPL/2.0/

title: "objectscale_iam_saml_provider resource"
linkTitle: "objectscale_iam_saml_provider"
page_title: "objectscale_iam_saml_provider Resource - terraform-provider-objectscale"
subcategory: "Identity & Access Management (IAM)"
description: |-
  Manages an ObjectScale IAM SAML Identity Provider (external IdP) registration.
---

# objectscale_iam_saml_provider (Resource)

Manages an ObjectScale IAM SAML Identity Provider registration. Used to register
external IdPs (Okta, Azure AD, ADFS, …) with an ObjectScale namespace so users
can federate via SAML 2.0 and assume IAM roles for S3 access.

> **Note**: ObjectScale also requires a Service Provider (SP) configuration. See
> [`objectscale_iam_service_provider`](./iam_service_provider.md). Both
> resources are required for end-to-end federation.

## Example Usage

```hcl
resource "objectscale_iam_saml_provider" "corp" {
  name                   = "corp-saml"
  namespace              = "ns1"
  saml_metadata_document = file("${path.module}/idp-metadata.xml")
}
```

## Schema

### Required

- `name` (String) — IdP name. **ForceNew** — cannot be renamed; changing this
  attribute destroys + recreates the resource.
- `saml_metadata_document` (String) — Raw SAML metadata XML for the IdP.
  Triggers an in-place Update on change.

### Optional

- `namespace` (String) — Namespace binding. Sent as `x-emc-namespace` header.
  **ForceNew**.

### Read-Only

- `id` (String) — The provider ARN, also the resource ID.
- `arn` (String) — Provider ARN: `urn:ecs:iam::<namespace>:saml-provider/<name>`.
- `create_date` (String) — ISO 8601 creation timestamp.
- `valid_until` (String) — ISO 8601 expiry of the IdP signing certificate.

## Import

Import by ARN:

```bash
terraform import objectscale_iam_saml_provider.corp \
  "urn:ecs:iam::ns1:saml-provider/corp-saml"
```
