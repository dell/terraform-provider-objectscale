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
    "/iam?Action=ListGroups",
    "/iam?Action=ListGroupsForUser",
    "/iam?Action=GetUser",
    "/iam?Action=GetGroup",
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
    "/iam?Action=CreateGroup",
    "/iam?Action=GetGroup",
    "/iam?Action=GetGroupPolicy",
    "/iam?Action=ListAttachedGroupPolicies",
    "/iam?Action=ListGroupPolicies",
    "/iam?Action=ListGroups",
    "/iam?Action=ListGroupsForUser",
    "/iam?Action=AddUserToGroup",
    "/iam?Action=AttachGroupPolicy",
    "/iam?Action=PutGroupPolicy",
    "/iam?Action=DeleteGroup",
    "/iam?Action=DeleteGroupPolicy",
    "/iam?Action=DetachGroupPolicy",
    "/iam?Action=DeleteGroup",
    "/iam?Action=RemoveUserFromGroup",
    "/object/bucket",
]
