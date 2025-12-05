# Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://mozilla.org/MPL/2.0/


# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

RequiredAPIs = [
    "/login",
    "/object/namespaces",
    "/object/namespaces/namespace/{id}",
    "/object/namespaces/namespace",
    "/object/namespaces/namespace/{namespace}",
    "/object/namespaces/namespace/{namespace}/deactivate",
    "/iam?Action=ListGroupsForUser",
    "/iam?Action=GetUser",
    "/iam?Action=ListUsers",
    "/iam?Action=ListUserTags",
    "/iam?Action=ListAccessKeys",
    "/object/namespaces/namespace/{namespace}/retention*",
    "/object/namespaces/namespace/{namespace}/quota",
    "/vdc/data-service/vpools",
    "/vdc/data-service/vpools/{id}",
    "/vdc/data-service/vpools/{id}/addvarrays",
    "/vdc/data-service/vpools/{id}/removevarrays",
    "/vdc/data-service/vpools",
    "/iam?Action=ListGroups",
    "/iam?Action=GetGroup",
    "/iam?Action=CreateGroup",
    "/iam?Action=DeleteGroup",
    "/iam?Action=AddUserToGroup",
    "/iam?Action=RemoveUserFromGroup",
    "/iam?Action=ListUserPolicies",
    "/iam?Action=ListGroupPolicies",
    "/iam?Action=ListRolePolicies",
    "/iam?Action=GetUserPolicy",
    "/iam?Action=GetGroupPolicy",
    "/iam?Action=GetRolePolicy",
    "/iam?Action=PutUserPolicy",
    "/iam?Action=PutGroupPolicy",
    "/iam?Action=PutRolePolicy",
    "/iam?Action=DeleteUserPolicy",
    "/iam?Action=DeleteGroupPolicy",
    "/iam?Action=DeleteRolePolicy",

    # user resource APIs
    "/iam?Action=CreateUser",
    "/iam?Action=DeleteUser",
    "/iam?Action=UntagUser",
    "/iam?Action=TagUser",
    "/iam?Action=PutUserPermissionsBoundary",
    "/iam?Action=DeleteUserPermissionsBoundary",
    
    "/iam?Action=ListRoles",
    "/iam?Action=GetRole",
    "/iam?Action=CreateRole",
    "/iam?Action=ListRoles",
    "/iam?Action=GetRole",
    "/iam?Action=ListRolePolicies",
    "/iam?Action=GetRolePolicy",
    "/iam?Action=ListRoleTags",  
    "/iam?Action=UpdateRole",
    "/iam?Action=UpdateAssumeRolePolicy",
    "/iam?Action=PutRolePolicy",
    "/iam?Action=DeleteRolePolicy",
    "/iam?Action=PutRolePermissionsBoundary",
    "/iam?Action=DeleteRolePermissionsBoundary",
    "/iam?Action=TagRole",
    "/iam?Action=UntagRole",
    "/iam?Action=DeleteRole",
    
    # Bucket API endpoints
    "/object/bucket",
    "/object/bucket/{bucketName}/deactivate",
    "/object/bucket/{bucketName}/empty-bucket-status",
    "/object/bucket/{bucketName}/tags",
    "/object/bucket/{bucketName}/autocommit",
    "/object/bucket/{bucketName}/retention",
    "/object/bucket/{bucketName}/object-lock-config",
    "/object/bucket/{bucketName}/allow-object-lock-with-ado",
    "/object/bucket/{bucketName}/info",
    "/object/bucket/{bucketName}/owner",
    "/object/bucket/{bucketName}/isstaleallowed",
    "/object/bucket/{bucketName}/lock",
    "/object/bucket/{bucketName}/lock/{IsLocked}",
    "/object/bucket/{bucketName}/quota",
    "/object/bucket/{bucketName}/acl",
    "/object/bucket/{bucketName}/policy",
    "/object/bucket/{bucketName}/acl",
    "/object/bucket/{bucketName}/defaultGroup",
    "/object/bucket/{bucketName}/metadata",
    "/object/bucket/acl/permissions",
    "/object/bucket/acl/groups",
    "/object/bucket/searchmetadata",
    "/object/bucket/{bucketName}/searchmetadata",
    "/object/bucket/{bucketName}/advancedMetadataSearch",
    "/object/bucket/{bucketName}/advancedMetadataSearchTarget",
    "/object/bucket/{bucketName}/auditDeleteExpiration",
    "/object/bucket/test-policy",
    "/object/bucket/test-policy-edit",
    "/object/bucket/{bucketName}/set-local-object-metadata-reads",
    "/object/bucket/{bucketName}/versioning",
    "/object/bucket/{bucketName}/notification",
    
    # Access Key API endpoints
    "/iam?Action=CreateAccessKey",
    "/iam?Action=DeleteAccessKey",
    "/iam?Action=UpdateAccessKey",
    
    # Policy API endpoints
    "/iam?Action=GetPolicy",
    "/iam?Action=ListPolicies",
    "/iam?Action=ListAttachedGroupPolicies",
    "/iam?Action=ListAttachedRolePolicies",
    "/iam?Action=ListAttachedUserPolicies",
    "/iam?Action=AttachGroupPolicy",
    "/iam?Action=DetachGroupPolicy",
    "/iam?Action=AttachRolePolicy",
    "/iam?Action=DetachRolePolicy",
    "/iam?Action=AttachUserPolicy",
    "/iam?Action=DetachUserPolicy"
]
