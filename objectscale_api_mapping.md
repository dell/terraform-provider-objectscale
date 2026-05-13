<!-- Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License. -->

# Dell ObjectScale Terraform Provider - REST API Endpoint Mapping

## Overview
This document provides an expert analysis mapping all Terraform datasources and resources to their corresponding ObjectScale REST API endpoints, including the HTTP Method and Path.

## API Endpoint Mapping

### 1. Bucket API Service

| Type | Name | Primary REST Endpoint(s) | Key Functionality |
|------|------|--------------------------|-------------------|
| Resource | `objectscale_bucket` | `POST /object/bucket` <br> `GET /object/bucket/{bucketName}/info` <br> `POST /object/bucket/{bucketName}/deactivate` <br> `PUT /object/bucket/{bucketName}/acl` <br> `PUT /object/bucket/{bucketName}/policy` <br> `PUT /object/bucket/{bucketName}/quota` <br> `PUT /object/bucket/{bucketName}/versioning` <br> `PUT /object/bucket/{bucketName}/tags` | Manages the full lifecycle of S3 buckets. |
| Datasource | `objectscale_bucket` | `GET /object/bucket` | Retrieves S3 bucket information. |

### 2. Namespace API Service

| Type | Name | Primary REST Endpoint(s) | Key Functionality |
|------|------|--------------------------|-------------------|
| Resource | `objectscale_namespace` | `POST /object/namespaces/namespace` <br> `GET /object/namespaces/namespace/{id}` <br> `PUT /object/namespaces/namespace/{namespace}` <br> `POST /object/namespaces/namespace/{namespace}/deactivate` | Manages ObjectScale namespaces. |
| Datasource | `objectscale_namespace` | `GET /object/namespaces` | Retrieves namespace information. |

### 3. IAM API Service

| Type | Name | Primary REST Endpoint(s) | Key Functionality |
|------|------|--------------------------|-------------------|
| Resource | `objectscale_iam_user` | `POST /iam?Action=CreateUser` <br> `POST /iam?Action=DeleteUser` <br> `POST /iam?Action=TagUser` | Manages IAM users. |
| Resource | `objectscale_iam_group` | `POST /iam?Action=CreateGroup` <br> `POST /iam?Action=DeleteGroup` | Manages IAM groups. |
| Resource | `objectscale_iam_role` | `POST /iam?Action=CreateRole` <br> `POST /iam?Action=DeleteRole` <br> `POST /iam?Action=UpdateRole` | Manages IAM roles. |
| Resource | `objectscale_iam_policy` | `POST /iam?Action=CreatePolicy` <br> `POST /iam?Action=DeletePolicy` | Manages IAM policies. |
| Resource | `objectscale_iam_inline_policy` | `POST /iam?Action=PutUserPolicy` <br> `POST /iam?Action=PutGroupPolicy` <br> `POST /iam?Action=PutRolePolicy` | Manages inline policies. |
| Resource | `objectscale_iam_policy_attachment` | `POST /iam?Action=AttachUserPolicy` <br> `POST /iam?Action=DetachUserPolicy` | Manages policy attachments. |
| Resource | `objectscale_iam_group_membership` | `POST /iam?Action=AddUserToGroup` <br> `POST /iam?Action=RemoveUserFromGroup` | Manages group memberships. |
| Resource | `objectscale_iam_user_access_key` | `POST /iam?Action=CreateAccessKey` <br> `POST /iam?Action=DeleteAccessKey` | Manages user access keys. |
| Datasource | `objectscale_iam_user` | `POST /iam?Action=GetUser` <br> `POST /iam?Action=ListUsers` | Retrieves IAM user information. |
| Datasource | `objectscale_iam_groups` | `POST /iam?Action=ListGroups` | Retrieves IAM group information. |
| Datasource | `objectscale_iam_role` | `POST /iam?Action=GetRole` <br> `POST /iam?Action=ListRoles` | Retrieves IAM role information. |
| Datasource | `objectscale_iam_policy` | `POST /iam?Action=GetPolicy` <br> `POST /iam?Action=ListPolicies` | Retrieves IAM policy information. |
| Datasource | `objectscale_iam_inline_policy` | `POST /iam?Action=GetUserPolicy` <br> `POST /iam?Action=ListUserPolicies` | Retrieves inline policy information. |

### 4. User & Key Management API Services

| Type | Name | Primary REST Endpoint(s) | Key Functionality |
|------|------|--------------------------|-------------------|
| Resource | `objectscale_object_user` | `POST /object/users` <br> `POST /object/users/deactivate` | Manages ObjectScale object users. |
| Resource | `objectscale_object_user_secret_key` | `POST /object/user-secret-keys/{uid}` <br> `POST /object/user-secret-keys/{uid}/deactivate` | Manages secret keys for object users. |
| Datasource | `objectscale_object_user` | `GET /object/users/query` <br> `GET /object/users/{uid}/info` | Retrieves object user information. |

### 5. Management User API Service

| Type | Name | Primary REST Endpoint(s) | Key Functionality |
|------|------|--------------------------|-------------------|
| Resource | `objectscale_management_user` | `POST /vdc/users` <br> `PUT /vdc/users/{userid}` <br> `POST /vdc/users/{userid}/deactivate` | Manages local management users. |
| Datasource | `objectscale_management_user` | `GET /vdc/users` <br> `GET /vdc/users/{userid}` | Retrieves management user information. |

### 6. VDC & Storage API Services

| Type | Name | Primary REST Endpoint(s) | Key Functionality |
|------|------|--------------------------|-------------------|
| Resource | `objectscale_replication_group` | `POST /vdc/data-service/vpools` <br> `PUT /vdc/data-service/vpools/{id}` | Manages replication groups (vpools). |
| Datasource | `objectscale_replication_group` | `GET /vdc/data-service/vpools` | Retrieves replication group information. |
| Datasource | `objectscale_storage_pool` | `GET /vdc/data-services/varrays` | Retrieves storage pools (varrays). |
| Datasource | `objectscale_vdc` | `GET /vdc/zones/zone/local` <br> `GET /vdc/zones/zone/all` | Retrieves VDC information. |

## Expert Summary

This analysis confirms the provider follows standard RESTful practices and AWS IAM-compatible API design patterns. The mapping reveals a clear and logical connection between the Terraform resources and the underlying ObjectScale API, which is essential for effective infrastructure automation and management.
