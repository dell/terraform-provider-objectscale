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

# Changelog

## [1.1.0] - 2026-06-23

### Added
- VDC (Virtual Data Center) resource and data source support
- Object certificate keystore management capabilities
- VDC certificate keystore management
- SAML provider resource and data source
- Service provider resource and data source

### Changed
- Enhanced create request handling to properly manage missing fields

### Fixed
- IAM user permissions boundary handling improvements
- Create request modified to handle missing fields (#56)

### Security
- Upgraded Go dependencies to resolve BlackDuck security issues (#62)

## [1.0.0] - 2025-12-26

### Added

#### Resources
- `objectscale_bucket` - S3 bucket management
- `objectscale_namespace` - Namespace management
- `objectscale_management_user` - Management user provisioning
- `objectscale_object_user` - Object user management
- `objectscale_object_user_secret_key` - Object user secret key management
- `objectscale_object_certificate` - Object certificate keystore management
- `objectscale_replication_group` - Replication group configuration
- `objectscale_vdc_certificate` - VDC certificate keystore management
- `objectscale_iam_user` - IAM user management
- `objectscale_iam_user_access_key` - IAM user access key management
- `objectscale_iam_group` - IAM group management
- `objectscale_iam_group_membership` - IAM group membership management
- `objectscale_iam_role` - IAM role management
- `objectscale_iam_policy` - IAM policy management
- `objectscale_iam_inline_policy` - IAM inline policy management
- `objectscale_iam_policy_attachment` - IAM policy attachment management
- `objectscale_iam_saml_provider` - SAML provider configuration
- `objectscale_iam_service_provider` - Service provider configuration

#### Data Sources
- `objectscale_bucket` - Bucket information lookup
- `objectscale_namespace` - Namespace information lookup
- `objectscale_storage_pool` - Storage pool information
- `objectscale_vdc` - Virtual Data Center information
- `objectscale_management_user` - Management user lookup
- `objectscale_object_user` - Object user lookup
- `objectscale_object_certificate` - Object certificate lookup
- `objectscale_vdc_certificate` - VDC certificate lookup
- `objectscale_replication_group` - Replication group lookup
- `objectscale_iam_user` - IAM user lookup
- `objectscale_iam_groups` - IAM groups listing
- `objectscale_iam_role` - IAM role lookup
- `objectscale_iam_policy` - IAM policy lookup
- `objectscale_iam_inline_policy` - IAM inline policy lookup
- `objectscale_iam_saml_provider` - SAML provider lookup
- `objectscale_iam_service_provider` - Service provider lookup
- `objectscale_iam_service_provider_metadata` - Service provider metadata

### Changed
- Improved test case coverage and reliability
- Enhanced parameter documentation across all resources

### Fixed
- IAM user permissions boundary handling
- IAM group resource operations
- IAM role resource functionality
- IAM inline policy resource operations
- Bucket resource update operations
- Object user resource handling
- Replication group length validators
- Tags handling for IAM role resources
- Create request handling for missing fields

### Security
- Upgraded Go dependencies to resolve BlackDuck security issues

---

## Release Notes Format

Each release follows this structure:

- **Added**: New features, resources, or data sources
- **Changed**: Changes in existing functionality
- **Deprecated**: Soon-to-be removed features
- **Removed**: Removed features
- **Fixed**: Bug fixes
- **Security**: Security vulnerability fixes

[1.1.0]: https://github.com/dell/terraform-provider-objectscale/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/dell/terraform-provider-objectscale/releases/tag/v1.0.0