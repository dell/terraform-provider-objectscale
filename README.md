<!--
Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
-->

# Terraform Provider for Dell Technologies ObjectScale
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v2.0%20adopted-ff69b4.svg)](https://github.com/dell/terraform-provider-objectscale/blob/main/about/CODE_OF_CONDUCT.md)
[![License](https://img.shields.io/badge/License-MPL_2.0-blue.svg)](https://github.com/dell/terraform-provider-objectscale/blob/main/LICENSE)

The Terraform Provider for Dell Technologies (Dell) ObjectScale allows Data Center and IT administrators to use Hashicorp Terraform to automate and orchestrate the provisioning and management of Dell ObjectScale storage systems.

The logged-in user configured in the Terraform provider must possess adequate permissions against the target Dell ObjectScale System.

## Table of Contents

* [Code of Conduct](https://github.com/dell/dell-terraform-providers/blob/main/docs/CODE_OF_CONDUCT.md)
* [Maintainer Guide](https://github.com/dell/dell-terraform-providers/blob/main/docs/MAINTAINER_GUIDE.md)
* [Committer Guide](https://github.com/dell/dell-terraform-providers/blob/main/docs/COMMITTER_GUIDE.md)
* [Contributing Guide](https://github.com/dell/dell-terraform-providers/blob/main/docs/CONTRIBUTING.md)
* [List of Adopters](https://github.com/dell/dell-terraform-providers/blob/main/docs/ADOPTERS.md)
* [Security](https://github.com/dell/dell-terraform-providers/blob/main/docs/SECURITY.md)
* [Support](#support)
* [License](#license)
* [Prerequisites](#prerequisites)
* [List of DataSources in Terraform Provider for Dell ObjectScale](#list-of-datasources-in-terraform-provider-for-dell-objectscale)
* [List of Resources in Terraform Provider for Dell ObjectScale](#list-of-resources-in-terraform-provider-for-dell-objectscale)
* [Releasing, Maintenance and Deprecation](#releasing-maintenance-and-deprecation)

## Support
For any Terraform Provider for Dell ObjectScale issues, questions or feedback, please follow our [support process](https://github.com/dell/dell-terraform-providers/blob/main/docs/SUPPORT.md)

## License
The Terraform Provider for Dell ObjectScale is released and licensed under the MPL-2.0 license. See [LICENSE](https://github.com/dell/terraform-provider-objectscale/blob/main/LICENSE) for the full terms.

## Prerequisites

| **Terraform Provider** | **ObjectScale Version** |         **OS**         |   **Terraform**   | **Golang** |
|------------------------|:------------------------|:-----------------------|-------------------|------------|
| v1.0.0                 | 4.1.x                   | RHEL 9.6, UBUNTU 22.04 | 1.14.3 and 1.13.5 | 1.25.6     |

## List of Data Sources in Terraform Provider for Dell ObjectScale

### Identity & Access Management (IAM)
* [IAM Group](docs/data-sources/iam_groups.md)
* [IAM Policy](docs/data-sources/iam_policy.md)
* [IAM Role](docs/data-sources/iam_role.md)
* [IAM User](docs/data-sources/iam_user.md)
* [IAM Inline Policy](docs/data-sources/iam_inline_policy.md)

### Namespacing & Tenancy
* [Namespace](docs/data-sources/namespace.md)

### User Management
* [Object User](docs/data-sources/object_user.md)
* [Management User](docs/data-sources/management_user.md)

### Object Storage Containers
* [Bucket](docs/data-sources/bucket.md)

### Data Protection
* [Replication Group](docs/data-sources/replication_group.md)

### Storage Topology & Capacity Domains
* [Storage Pool](docs/data-sources/storage_pool.md)
* [VDC](docs/data-sources/vdc.md)

## List of Resources in Terraform Provider for Dell ObjectScale

### Identity & Access Management (IAM)
* [IAM Group](docs/resources/iam_group.md)
* [IAM Policy](docs/resources/iam_policy.md)
* [IAM Inline Policy](docs/resources/iam_inline_policy.md)
* [IAM Policy Attachment](docs/resources/iam_policy_attachment.md)
* [IAM Role](docs/resources/iam_role.md)
* [IAM User](docs/resources/iam_user.md)
* [IAM User Access Key](docs/resources/iam_user_access_key.md)
* [IAM Group Membership](docs/resources/iam_group_membership.md)

### Object Storage Containers
* [Bucket](docs/resources/bucket.md)

### Namespace and Tenancy
* [Namespace](docs/resources/namespace.md)

### User Management
* [Object User ](docs/resources/object_user.md)
* [Object User Secret Key](docs/resources/object_user_secret_key.md)
* [Management User](docs/data-sources/management_user.md)

### Data Protection
* [Replication Group](docs/resources/replication_group.md)

## Installation and execution of Terraform Provider for Dell ObjectScale

## Installation from public repository

The provider will be fetched from the public repository and installed by Terraform automatically.
Create a file called `main.tf` in your workspace with the following contents

```tf
terraform {
  required_providers {
    objectscale = { 
      version = "1.0.0"
      source = "registry.terraform.io/dell/objectscale"
    }
  }
}
```
Then, in that workspace, run
```
terraform init
``` 

## Installation from source code

1. Clone this repo
2. In the root of this repo run
```
make install
```
Then follow [installation from public repo](#installation-from-public-repository)

## SSL Certificate Verification

For SSL verifcation on RHEL, these steps can be performed:
* Copy the CA certificate to the `/etc/pki/ca-trust/source/anchors` path of the host by any external means.
* Import the SSL certificate to host by running
```
update-ca-trust extract
```

For SSL verification on Ubuntu, these steps can be performed:
* Copy the CA certificate to the `/etc/ssl/certs` path of the host by any external means.
* Import the SSL certificate to host by running:
 ```
  update-ca-certificates
```

## Releasing, Maintenance and Deprecation

Terraform Provider for Dell Technologies ObjectScale follows [Semantic Versioning](https://semver.org/).

New versions will be released on a regular basis whenever significant updates—such as bug fixes or new features—are introduced to the provider.

Released code versions are located on tags in the form of "vx.y.z" where x.y.z corresponds to the version number.

## Documentation

For more detailed information, please refer to 
  * [Dell Terraform Providers Documentation](https://dell.github.io/terraform-docs/)
  * [Dell Terraform Registry](https://registry.terraform.io/providers/dell/objectscale/latest/docs)
