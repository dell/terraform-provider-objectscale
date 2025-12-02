# \IamApi

All URIs are relative to *https://objectscale.local:4443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamServiceAddUserToGroup**](IamApi.md#IamServiceAddUserToGroup) | **Post** /iam?Action&#x3D;AddUserToGroup | Add user to a group.
[**IamServiceAttachGroupPolicy**](IamApi.md#IamServiceAttachGroupPolicy) | **Post** /iam?Action&#x3D;AttachGroupPolicy | Attach a Managed Policy to Group.
[**IamServiceAttachRolePolicy**](IamApi.md#IamServiceAttachRolePolicy) | **Post** /iam?Action&#x3D;AttachRolePolicy | Attaches the specified managed policy to the specified IAM role.
[**IamServiceCreateAccessKey**](IamApi.md#IamServiceCreateAccessKey) | **Post** /iam?Action&#x3D;CreateAccessKey | Create AccessKey for User.
[**IamServiceCreateGroup**](IamApi.md#IamServiceCreateGroup) | **Post** /iam?Action&#x3D;CreateGroup | Creates a new IAM Group.
[**IamServiceCreateRole**](IamApi.md#IamServiceCreateRole) | **Post** /iam?Action&#x3D;CreateRole | Creates a new IAM role.
[**IamServiceDeleteAccessKey**](IamApi.md#IamServiceDeleteAccessKey) | **Post** /iam?Action&#x3D;DeleteAccessKey | Delete access key.
[**IamServiceDeleteGroup**](IamApi.md#IamServiceDeleteGroup) | **Post** /iam?Action&#x3D;DeleteGroup | Delete an IAM Group.
[**IamServiceDeleteGroupPolicy**](IamApi.md#IamServiceDeleteGroupPolicy) | **Post** /iam?Action&#x3D;DeleteGroupPolicy | Delete specific inlinePolicy for IAM Group.
[**IamServiceDeleteRole**](IamApi.md#IamServiceDeleteRole) | **Post** /iam?Action&#x3D;DeleteRole | Deletes the specified IAM role.
[**IamServiceDeleteRolePermissionsBoundary**](IamApi.md#IamServiceDeleteRolePermissionsBoundary) | **Post** /iam?Action&#x3D;DeleteRolePermissionsBoundary | Deletes the permissions boundary for the specified IAM role.
[**IamServiceDeleteRolePolicy**](IamApi.md#IamServiceDeleteRolePolicy) | **Post** /iam?Action&#x3D;DeleteRolePolicy | Deletes the specified inline policy that is embedded in the specified IAM role.
[**IamServiceDeleteUserPolicy**](IamApi.md#IamServiceDeleteUserPolicy) | **Post** /iam?Action&#x3D;DeleteUserPolicy | Delete specific inlinePolicy for IAM User.
[**IamServiceDetachGroupPolicy**](IamApi.md#IamServiceDetachGroupPolicy) | **Post** /iam?Action&#x3D;DetachGroupPolicy | Remove a Managed Policy attached to Group.
[**IamServiceDetachRolePolicy**](IamApi.md#IamServiceDetachRolePolicy) | **Post** /iam?Action&#x3D;DetachRolePolicy | Removes the specified managed policy from the specified IAM role.
[**IamServiceGetGroup**](IamApi.md#IamServiceGetGroup) | **Post** /iam?Action&#x3D;GetGroup | Retrieve list of users in IAM group.
[**IamServiceGetGroupPolicy**](IamApi.md#IamServiceGetGroupPolicy) | **Post** /iam?Action&#x3D;GetGroupPolicy | Get specific inlinePolicy for IAM Group.
[**IamServiceGetRole**](IamApi.md#IamServiceGetRole) | **Post** /iam?Action&#x3D;GetRole | Gets information about the specified IAM role.
[**IamServiceGetRolePolicy**](IamApi.md#IamServiceGetRolePolicy) | **Post** /iam?Action&#x3D;GetRolePolicy | Gets tthe specified inline policy document that is embedded with the specified IAM role.
[**IamServiceGetUser**](IamApi.md#IamServiceGetUser) | **Post** /iam?Action&#x3D;GetUser | Retrieve IAM user.
[**IamServiceGetUserPolicy**](IamApi.md#IamServiceGetUserPolicy) | **Post** /iam?Action&#x3D;GetUserPolicy | Get specific inlinePolicy for IAM User.
[**IamServiceListAccessKeys**](IamApi.md#IamServiceListAccessKeys) | **Post** /iam?Action&#x3D;ListAccessKeys | List AccessKeys for a user.
[**IamServiceListAttachedGroupPolicies**](IamApi.md#IamServiceListAttachedGroupPolicies) | **Post** /iam?Action&#x3D;ListAttachedGroupPolicies | List Managed Policies for IAM Group.
[**IamServiceListAttachedRolePolicies**](IamApi.md#IamServiceListAttachedRolePolicies) | **Post** /iam?Action&#x3D;ListAttachedRolePolicies | Lists all managed policies that are attached to the specified IAM Role.
[**IamServiceListAttachedUserPolicies**](IamApi.md#IamServiceListAttachedUserPolicies) | **Post** /iam?Action&#x3D;ListAttachedUserPolicies | List Managed Policies for IAM User.
[**IamServiceListGroupPolicies**](IamApi.md#IamServiceListGroupPolicies) | **Post** /iam?Action&#x3D;ListGroupPolicies | List Inline Policies for IAM Group.
[**IamServiceListGroups**](IamApi.md#IamServiceListGroups) | **Post** /iam?Action&#x3D;ListGroups | Lists the IAM groups.
[**IamServiceListGroupsForUser**](IamApi.md#IamServiceListGroupsForUser) | **Post** /iam?Action&#x3D;ListGroupsForUser | List Groups for IAM User
[**IamServiceListPolicies**](IamApi.md#IamServiceListPolicies) | **Post** /iam?Action&#x3D;ListPolicies | Lists the IAM users.
[**IamServiceListRolePolicies**](IamApi.md#IamServiceListRolePolicies) | **Post** /iam?Action&#x3D;ListRolePolicies | Lists the names of the inline policies that are embedded in the specified IAM role.
[**IamServiceListRoleTags**](IamApi.md#IamServiceListRoleTags) | **Post** /iam?Action&#x3D;ListRoleTags | Lists the tags that are attached to the specified IAM role.
[**IamServiceListRoles**](IamApi.md#IamServiceListRoles) | **Post** /iam?Action&#x3D;ListRoles | Lists the IAM roles.
[**IamServiceListUserPolicies**](IamApi.md#IamServiceListUserPolicies) | **Post** /iam?Action&#x3D;ListUserPolicies | List Inline Policies for IAM User.
[**IamServiceListUserTags**](IamApi.md#IamServiceListUserTags) | **Post** /iam?Action&#x3D;ListUserTags | Lists the tags that are attached to the specified IAM User.
[**IamServiceListUsers**](IamApi.md#IamServiceListUsers) | **Post** /iam?Action&#x3D;ListUsers | Lists the IAM users.
[**IamServicePutGroupPolicy**](IamApi.md#IamServicePutGroupPolicy) | **Post** /iam?Action&#x3D;PutGroupPolicy | Add or Update Inline Policy for IAM Group.
[**IamServicePutRolePermissionsBoundary**](IamApi.md#IamServicePutRolePermissionsBoundary) | **Post** /iam?Action&#x3D;PutRolePermissionsBoundary | Adds or updates the policy that is specified as the IAM role&#39;s permissions boundary.
[**IamServicePutRolePolicy**](IamApi.md#IamServicePutRolePolicy) | **Post** /iam?Action&#x3D;PutRolePolicy | Adds or updates an inline policy document that is embedded in the specified IAM role.
[**IamServicePutUserPolicy**](IamApi.md#IamServicePutUserPolicy) | **Post** /iam?Action&#x3D;PutUserPolicy | Add or Update Inline Policy for IAM User.
[**IamServiceRemoveUserFromGroup**](IamApi.md#IamServiceRemoveUserFromGroup) | **Post** /iam?Action&#x3D;RemoveUserFromGroup | Remove User from a Group.
[**IamServiceTagRole**](IamApi.md#IamServiceTagRole) | **Post** /iam?Action&#x3D;TagRole | Adds one or more tags to a specified IAM Role.
[**IamServiceUntagRole**](IamApi.md#IamServiceUntagRole) | **Post** /iam?Action&#x3D;UntagRole | Removes the specified tags from a specified IAM Role.
[**IamServiceUpdateAssumeRolePolicy**](IamApi.md#IamServiceUpdateAssumeRolePolicy) | **Post** /iam?Action&#x3D;UpdateAssumeRolePolicy | Updates the policy that grants an IAM entity permission to assume a role.
[**IamServiceUpdateRole**](IamApi.md#IamServiceUpdateRole) | **Post** /iam?Action&#x3D;UpdateRole | Updates the description or maximum session duration setting of the specified IAM role.



## IamServiceAddUserToGroup

> IamServiceAddUserToGroupResponse IamServiceAddUserToGroup(ctx).GroupName(groupName).UserName(userName).XEmcNamespace(xEmcNamespace).Execute()

Add user to a group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    groupName := "groupName_example" // string | The name of the group to update. (optional)
    userName := "userName_example" // string | The name of the user to add. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceAddUserToGroup(context.Background()).GroupName(groupName).UserName(userName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceAddUserToGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceAddUserToGroup`: IamServiceAddUserToGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceAddUserToGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceAddUserToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupName** | **string** | The name of the group to update. | 
 **userName** | **string** | The name of the user to add. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceAddUserToGroupResponse**](IamServiceAddUserToGroupResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceAttachGroupPolicy

> IamServiceAttachGroupPolicyResponse IamServiceAttachGroupPolicy(ctx).PolicyArn(policyArn).GroupName(groupName).XEmcNamespace(xEmcNamespace).Execute()

Attach a Managed Policy to Group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    policyArn := "policyArn_example" // string | Arn of the policy to attach. (optional)
    groupName := "groupName_example" // string | Name of the group to attach the policy. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceAttachGroupPolicy(context.Background()).PolicyArn(policyArn).GroupName(groupName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceAttachGroupPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceAttachGroupPolicy`: IamServiceAttachGroupPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceAttachGroupPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceAttachGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyArn** | **string** | Arn of the policy to attach. | 
 **groupName** | **string** | Name of the group to attach the policy. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceAttachGroupPolicyResponse**](IamServiceAttachGroupPolicyResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceAttachRolePolicy

> IamServiceAttachRolePolicyResponse IamServiceAttachRolePolicy(ctx).PolicyArn(policyArn).RoleName(roleName).XEmcNamespace(xEmcNamespace).Execute()

Attaches the specified managed policy to the specified IAM role.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    policyArn := "policyArn_example" // string | Arn that identifies the policy. (optional)
    roleName := "roleName_example" // string | Simple name identifying the role. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceAttachRolePolicy(context.Background()).PolicyArn(policyArn).RoleName(roleName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceAttachRolePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceAttachRolePolicy`: IamServiceAttachRolePolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceAttachRolePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceAttachRolePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyArn** | **string** | Arn that identifies the policy. | 
 **roleName** | **string** | Simple name identifying the role. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceAttachRolePolicyResponse**](IamServiceAttachRolePolicyResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceCreateAccessKey

> IamServiceCreateAccessKeyResponse IamServiceCreateAccessKey(ctx).UserName(userName).XEmcNamespace(xEmcNamespace).Execute()

Create AccessKey for User.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    userName := "userName_example" // string | The name of the user that new AccessKey belongs to. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceCreateAccessKey(context.Background()).UserName(userName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceCreateAccessKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceCreateAccessKey`: IamServiceCreateAccessKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceCreateAccessKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceCreateAccessKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userName** | **string** | The name of the user that new AccessKey belongs to. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceCreateAccessKeyResponse**](IamServiceCreateAccessKeyResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceCreateGroup

> IamServiceCreateGroupResponse IamServiceCreateGroup(ctx).GroupName(groupName).Path(path).XEmcNamespace(xEmcNamespace).Execute()

Creates a new IAM Group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    groupName := "groupName_example" // string | The name of the group to create. (optional)
    path := "path_example" // string | The path for the group. Optional, defaults to \"/\" and only \"/\" is allowed. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceCreateGroup(context.Background()).GroupName(groupName).Path(path).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceCreateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceCreateGroup`: IamServiceCreateGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceCreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupName** | **string** | The name of the group to create. | 
 **path** | **string** | The path for the group. Optional, defaults to \&quot;/\&quot; and only \&quot;/\&quot; is allowed. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceCreateGroupResponse**](IamServiceCreateGroupResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceCreateRole

> IamServiceCreateRoleResponse IamServiceCreateRole(ctx).RoleName(roleName).AssumeRolePolicyDocument(assumeRolePolicyDocument).MaxSessionDuration(maxSessionDuration).Description(description).Path(path).PermissionsBoundary(permissionsBoundary).TagsMemberN(tagsMemberN).XEmcNamespace(xEmcNamespace).Execute()

Creates a new IAM role.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    roleName := "roleName_example" // string | The name of the role to create. (optional)
    assumeRolePolicyDocument := "assumeRolePolicyDocument_example" // string | The trust relationship policy document that grants an entity permission to assume the role (optional)
    maxSessionDuration := int32(56) // int32 | The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied.  This setting can have a value from 1 hour to 12 hours (optional)
    description := "description_example" // string | A description of the role. (optional)
    path := "path_example" // string | The path to the role. Optional, defaults to \"/\" and only \"/\" is allowed. (optional)
    permissionsBoundary := "permissionsBoundary_example" // string | The ARN of the policy that is used to set the permissions boundary for the role. (optional)
    tagsMemberN := map[string]interface{}{ ... } // map[string]interface{} | A list of tags that you want to attach to the role being created. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceCreateRole(context.Background()).RoleName(roleName).AssumeRolePolicyDocument(assumeRolePolicyDocument).MaxSessionDuration(maxSessionDuration).Description(description).Path(path).PermissionsBoundary(permissionsBoundary).TagsMemberN(tagsMemberN).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceCreateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceCreateRole`: IamServiceCreateRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceCreateRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleName** | **string** | The name of the role to create. | 
 **assumeRolePolicyDocument** | **string** | The trust relationship policy document that grants an entity permission to assume the role | 
 **maxSessionDuration** | **int32** | The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied.  This setting can have a value from 1 hour to 12 hours | 
 **description** | **string** | A description of the role. | 
 **path** | **string** | The path to the role. Optional, defaults to \&quot;/\&quot; and only \&quot;/\&quot; is allowed. | 
 **permissionsBoundary** | **string** | The ARN of the policy that is used to set the permissions boundary for the role. | 
 **tagsMemberN** | [**map[string]interface{}**](map[string]interface{}.md) | A list of tags that you want to attach to the role being created. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceCreateRoleResponse**](IamServiceCreateRoleResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceDeleteAccessKey

> IamServiceDeleteAccessKeyResponse IamServiceDeleteAccessKey(ctx).AccessKeyId(accessKeyId).UserName(userName).XEmcNamespace(xEmcNamespace).Execute()

Delete access key.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    accessKeyId := "accessKeyId_example" // string | The access key ID for the access key ID and secret access key you want to delete. (optional)
    userName := "userName_example" // string | Name of the user to delete accesskeys. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceDeleteAccessKey(context.Background()).AccessKeyId(accessKeyId).UserName(userName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceDeleteAccessKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceDeleteAccessKey`: IamServiceDeleteAccessKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceDeleteAccessKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceDeleteAccessKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessKeyId** | **string** | The access key ID for the access key ID and secret access key you want to delete. | 
 **userName** | **string** | Name of the user to delete accesskeys. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceDeleteAccessKeyResponse**](IamServiceDeleteAccessKeyResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceDeleteGroup

> IamServiceDeleteGroupResponse IamServiceDeleteGroup(ctx).GroupName(groupName).XEmcNamespace(xEmcNamespace).Execute()

Delete an IAM Group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    groupName := "groupName_example" // string | The name of the group to delete. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceDeleteGroup(context.Background()).GroupName(groupName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceDeleteGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceDeleteGroup`: IamServiceDeleteGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceDeleteGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceDeleteGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupName** | **string** | The name of the group to delete. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceDeleteGroupResponse**](IamServiceDeleteGroupResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceDeleteGroupPolicy

> IamServiceDeleteGroupPolicyResponse IamServiceDeleteGroupPolicy(ctx).GroupName(groupName).PolicyName(policyName).XEmcNamespace(xEmcNamespace).Execute()

Delete specific inlinePolicy for IAM Group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    groupName := "groupName_example" // string | Name of the group to delete the inline policy. (optional)
    policyName := "policyName_example" // string | Name of the policy whose Policy Document needs to be deleted. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceDeleteGroupPolicy(context.Background()).GroupName(groupName).PolicyName(policyName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceDeleteGroupPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceDeleteGroupPolicy`: IamServiceDeleteGroupPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceDeleteGroupPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceDeleteGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupName** | **string** | Name of the group to delete the inline policy. | 
 **policyName** | **string** | Name of the policy whose Policy Document needs to be deleted. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceDeleteGroupPolicyResponse**](IamServiceDeleteGroupPolicyResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceDeleteRole

> IamServiceDeleteRoleResponse IamServiceDeleteRole(ctx).RoleName(roleName).XEmcNamespace(xEmcNamespace).Execute()

Deletes the specified IAM role.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    roleName := "roleName_example" // string | Simple name identifying the role. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceDeleteRole(context.Background()).RoleName(roleName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceDeleteRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceDeleteRole`: IamServiceDeleteRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceDeleteRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleName** | **string** | Simple name identifying the role. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceDeleteRoleResponse**](IamServiceDeleteRoleResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceDeleteRolePermissionsBoundary

> IamServiceDeleteRolePermissionsBoundaryResponse IamServiceDeleteRolePermissionsBoundary(ctx).RoleName(roleName).XEmcNamespace(xEmcNamespace).Execute()

Deletes the permissions boundary for the specified IAM role.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    roleName := "roleName_example" // string | Simple name identifying the role. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceDeleteRolePermissionsBoundary(context.Background()).RoleName(roleName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceDeleteRolePermissionsBoundary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceDeleteRolePermissionsBoundary`: IamServiceDeleteRolePermissionsBoundaryResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceDeleteRolePermissionsBoundary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceDeleteRolePermissionsBoundaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleName** | **string** | Simple name identifying the role. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceDeleteRolePermissionsBoundaryResponse**](IamServiceDeleteRolePermissionsBoundaryResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceDeleteRolePolicy

> IamServiceDeleteRolePolicyResponse IamServiceDeleteRolePolicy(ctx).RoleName(roleName).PolicyName(policyName).XEmcNamespace(xEmcNamespace).Execute()

Deletes the specified inline policy that is embedded in the specified IAM role.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    roleName := "roleName_example" // string | Simple name identifying the role. (optional)
    policyName := "policyName_example" // string | Simple name identifying the policy. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceDeleteRolePolicy(context.Background()).RoleName(roleName).PolicyName(policyName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceDeleteRolePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceDeleteRolePolicy`: IamServiceDeleteRolePolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceDeleteRolePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceDeleteRolePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleName** | **string** | Simple name identifying the role. | 
 **policyName** | **string** | Simple name identifying the policy. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceDeleteRolePolicyResponse**](IamServiceDeleteRolePolicyResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceDeleteUserPolicy

> IamServiceDeleteUserPolicyResponse IamServiceDeleteUserPolicy(ctx).UserName(userName).PolicyName(policyName).XEmcNamespace(xEmcNamespace).Execute()

Delete specific inlinePolicy for IAM User.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    userName := "userName_example" // string | Name of the user to delete the inline policy. (optional)
    policyName := "policyName_example" // string | Name of the policy whose Policy Document needs to be deleted. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceDeleteUserPolicy(context.Background()).UserName(userName).PolicyName(policyName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceDeleteUserPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceDeleteUserPolicy`: IamServiceDeleteUserPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceDeleteUserPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceDeleteUserPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userName** | **string** | Name of the user to delete the inline policy. | 
 **policyName** | **string** | Name of the policy whose Policy Document needs to be deleted. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceDeleteUserPolicyResponse**](IamServiceDeleteUserPolicyResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceDetachGroupPolicy

> IamServiceDetachGroupPolicyResponse IamServiceDetachGroupPolicy(ctx).PolicyArn(policyArn).GroupName(groupName).XEmcNamespace(xEmcNamespace).Execute()

Remove a Managed Policy attached to Group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    policyArn := "policyArn_example" // string | Arn of the policy to remove. (optional)
    groupName := "groupName_example" // string | Name of the group to remove the policy. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceDetachGroupPolicy(context.Background()).PolicyArn(policyArn).GroupName(groupName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceDetachGroupPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceDetachGroupPolicy`: IamServiceDetachGroupPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceDetachGroupPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceDetachGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyArn** | **string** | Arn of the policy to remove. | 
 **groupName** | **string** | Name of the group to remove the policy. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceDetachGroupPolicyResponse**](IamServiceDetachGroupPolicyResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceDetachRolePolicy

> IamServiceDetachRolePolicyResponse IamServiceDetachRolePolicy(ctx).PolicyArn(policyArn).RoleName(roleName).XEmcNamespace(xEmcNamespace).Execute()

Removes the specified managed policy from the specified IAM role.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    policyArn := "policyArn_example" // string | Arn that identifies the policy. (optional)
    roleName := "roleName_example" // string | Simple name identifying the role. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceDetachRolePolicy(context.Background()).PolicyArn(policyArn).RoleName(roleName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceDetachRolePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceDetachRolePolicy`: IamServiceDetachRolePolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceDetachRolePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceDetachRolePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyArn** | **string** | Arn that identifies the policy. | 
 **roleName** | **string** | Simple name identifying the role. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceDetachRolePolicyResponse**](IamServiceDetachRolePolicyResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceGetGroup

> IamServiceGetGroupResponse IamServiceGetGroup(ctx).GroupName(groupName).Marker(marker).MaxItems(maxItems).XEmcNamespace(xEmcNamespace).Execute()

Retrieve list of users in IAM group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    groupName := "groupName_example" // string | The name of the group. (optional)
    marker := "marker_example" // string | Marker is obtained from paginated response from the previous query. Use this only if the response indicates it is truncated. (optional)
    maxItems := int32(56) // int32 | Indicates the maximum number of elements to be returned in the response. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceGetGroup(context.Background()).GroupName(groupName).Marker(marker).MaxItems(maxItems).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceGetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceGetGroup`: IamServiceGetGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceGetGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupName** | **string** | The name of the group. | 
 **marker** | **string** | Marker is obtained from paginated response from the previous query. Use this only if the response indicates it is truncated. | 
 **maxItems** | **int32** | Indicates the maximum number of elements to be returned in the response. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceGetGroupResponse**](IamServiceGetGroupResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceGetGroupPolicy

> IamServiceGetGroupPolicyResponse IamServiceGetGroupPolicy(ctx).GroupName(groupName).PolicyName(policyName).XEmcNamespace(xEmcNamespace).Execute()

Get specific inlinePolicy for IAM Group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    groupName := "groupName_example" // string | Name of the group to retrieve the inline policy. (optional)
    policyName := "policyName_example" // string | Name of the policy whose Policy Document needs to be retrieved. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceGetGroupPolicy(context.Background()).GroupName(groupName).PolicyName(policyName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceGetGroupPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceGetGroupPolicy`: IamServiceGetGroupPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceGetGroupPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceGetGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupName** | **string** | Name of the group to retrieve the inline policy. | 
 **policyName** | **string** | Name of the policy whose Policy Document needs to be retrieved. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceGetGroupPolicyResponse**](IamServiceGetGroupPolicyResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceGetRole

> IamServiceGetRoleResponse IamServiceGetRole(ctx).RoleName(roleName).XEmcNamespace(xEmcNamespace).Execute()

Gets information about the specified IAM role.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    roleName := "roleName_example" // string | Simple name identifying the role. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceGetRole(context.Background()).RoleName(roleName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceGetRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceGetRole`: IamServiceGetRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceGetRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleName** | **string** | Simple name identifying the role. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceGetRoleResponse**](IamServiceGetRoleResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceGetRolePolicy

> IamServiceGetRolePolicyResponse IamServiceGetRolePolicy(ctx).RoleName(roleName).PolicyName(policyName).XEmcNamespace(xEmcNamespace).Execute()

Gets tthe specified inline policy document that is embedded with the specified IAM role.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    roleName := "roleName_example" // string | Simple name identifying the role. (optional)
    policyName := "policyName_example" // string | Simple name identifying the policy. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceGetRolePolicy(context.Background()).RoleName(roleName).PolicyName(policyName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceGetRolePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceGetRolePolicy`: IamServiceGetRolePolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceGetRolePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceGetRolePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleName** | **string** | Simple name identifying the role. | 
 **policyName** | **string** | Simple name identifying the policy. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceGetRolePolicyResponse**](IamServiceGetRolePolicyResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceGetUser

> IamServiceGetUserResponse IamServiceGetUser(ctx).UserName(userName).XEmcNamespace(xEmcNamespace).Execute()

Retrieve IAM user.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    userName := "userName_example" // string | The name of the user to retrieve. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceGetUser(context.Background()).UserName(userName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceGetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceGetUser`: IamServiceGetUserResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceGetUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userName** | **string** | The name of the user to retrieve. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceGetUserResponse**](IamServiceGetUserResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceGetUserPolicy

> IamServiceGetUserPolicyResponse IamServiceGetUserPolicy(ctx).UserName(userName).PolicyName(policyName).XEmcNamespace(xEmcNamespace).Execute()

Get specific inlinePolicy for IAM User.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    userName := "userName_example" // string | Name of the user to retrieve the inline policy. (optional)
    policyName := "policyName_example" // string | Name of the policy whose Policy Document needs to be retrieved. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceGetUserPolicy(context.Background()).UserName(userName).PolicyName(policyName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceGetUserPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceGetUserPolicy`: IamServiceGetUserPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceGetUserPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceGetUserPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userName** | **string** | Name of the user to retrieve the inline policy. | 
 **policyName** | **string** | Name of the policy whose Policy Document needs to be retrieved. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceGetUserPolicyResponse**](IamServiceGetUserPolicyResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceListAccessKeys

> IamServiceListAccessKeysResponse IamServiceListAccessKeys(ctx).Marker(marker).MaxItems(maxItems).PathPrefix(pathPrefix).UserName(userName).XEmcNamespace(xEmcNamespace).Execute()

List AccessKeys for a user.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    marker := "marker_example" // string | Marker is obtained from paginated response from the previous query. Use this only if the response indicates it is truncated. (optional)
    maxItems := int32(56) // int32 | Indicates the maximum number of elements to be returned in the response. (optional)
    pathPrefix := "pathPrefix_example" // string | Path prefix for filtering the results. Optional, default to \"/\". Only \"/\" is allowed. (optional)
    userName := "userName_example" // string | Name of the user to list accesskeys. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceListAccessKeys(context.Background()).Marker(marker).MaxItems(maxItems).PathPrefix(pathPrefix).UserName(userName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceListAccessKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceListAccessKeys`: IamServiceListAccessKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceListAccessKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceListAccessKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marker** | **string** | Marker is obtained from paginated response from the previous query. Use this only if the response indicates it is truncated. | 
 **maxItems** | **int32** | Indicates the maximum number of elements to be returned in the response. | 
 **pathPrefix** | **string** | Path prefix for filtering the results. Optional, default to \&quot;/\&quot;. Only \&quot;/\&quot; is allowed. | 
 **userName** | **string** | Name of the user to list accesskeys. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceListAccessKeysResponse**](IamServiceListAccessKeysResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceListAttachedGroupPolicies

> IamServiceListAttachedGroupPoliciesResponse IamServiceListAttachedGroupPolicies(ctx).GroupName(groupName).Marker(marker).MaxItems(maxItems).PathPrefix(pathPrefix).XEmcNamespace(xEmcNamespace).Execute()

List Managed Policies for IAM Group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    groupName := "groupName_example" // string | The name of the group to list attached policies for. (optional)
    marker := "marker_example" // string | Marker is obtained from paginated response from the previous query. Use this only if the response indicates it is truncated. (optional)
    maxItems := int32(56) // int32 | Indicates the maximum number of elements to be returned in the response. (optional)
    pathPrefix := "pathPrefix_example" // string | Path prefix for filtering the results. Optional, default to \"/\". Only \"/\" is allowed. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceListAttachedGroupPolicies(context.Background()).GroupName(groupName).Marker(marker).MaxItems(maxItems).PathPrefix(pathPrefix).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceListAttachedGroupPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceListAttachedGroupPolicies`: IamServiceListAttachedGroupPoliciesResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceListAttachedGroupPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceListAttachedGroupPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupName** | **string** | The name of the group to list attached policies for. | 
 **marker** | **string** | Marker is obtained from paginated response from the previous query. Use this only if the response indicates it is truncated. | 
 **maxItems** | **int32** | Indicates the maximum number of elements to be returned in the response. | 
 **pathPrefix** | **string** | Path prefix for filtering the results. Optional, default to \&quot;/\&quot;. Only \&quot;/\&quot; is allowed. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceListAttachedGroupPoliciesResponse**](IamServiceListAttachedGroupPoliciesResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceListAttachedRolePolicies

> IamServiceListAttachedRolePoliciesResponse IamServiceListAttachedRolePolicies(ctx).Marker(marker).MaxItems(maxItems).PathPrefix(pathPrefix).RoleName(roleName).XEmcNamespace(xEmcNamespace).Execute()

Lists all managed policies that are attached to the specified IAM Role.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    marker := "marker_example" // string | For pagination, the value of the Marker element in the response that you received to indicate where the next call should start. (optional)
    maxItems := int32(56) // int32 | Use this only when paginating results to indicate the maximum number of items you want in the response.  If additional items exist beyond the maximum you specify, the IsTruncated response element is true and  Marker contains a value to include in the subsequent call that tells the service where to continue from. (optional)
    pathPrefix := "pathPrefix_example" // string | The path to the IAM role. (optional)
    roleName := "roleName_example" // string | Simple name identifying the role. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceListAttachedRolePolicies(context.Background()).Marker(marker).MaxItems(maxItems).PathPrefix(pathPrefix).RoleName(roleName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceListAttachedRolePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceListAttachedRolePolicies`: IamServiceListAttachedRolePoliciesResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceListAttachedRolePolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceListAttachedRolePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marker** | **string** | For pagination, the value of the Marker element in the response that you received to indicate where the next call should start. | 
 **maxItems** | **int32** | Use this only when paginating results to indicate the maximum number of items you want in the response.  If additional items exist beyond the maximum you specify, the IsTruncated response element is true and  Marker contains a value to include in the subsequent call that tells the service where to continue from. | 
 **pathPrefix** | **string** | The path to the IAM role. | 
 **roleName** | **string** | Simple name identifying the role. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceListAttachedRolePoliciesResponse**](IamServiceListAttachedRolePoliciesResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceListAttachedUserPolicies

> IamServiceListAttachedUserPoliciesResponse IamServiceListAttachedUserPolicies(ctx).UserName(userName).Marker(marker).MaxItems(maxItems).PathPrefix(pathPrefix).XEmcNamespace(xEmcNamespace).Execute()

List Managed Policies for IAM User.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    userName := "userName_example" // string | The name of the user to list attached policies for. (optional)
    marker := "marker_example" // string | Marker is obtained from paginated response from the previous query. Use this only if the response indicates it is truncated. (optional)
    maxItems := int32(56) // int32 | Indicates the maximum number of elements to be returned in the response. (optional)
    pathPrefix := "pathPrefix_example" // string | Path prefix for filtering the results. Optional, default to \"/\". Only \"/\" is allowed. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceListAttachedUserPolicies(context.Background()).UserName(userName).Marker(marker).MaxItems(maxItems).PathPrefix(pathPrefix).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceListAttachedUserPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceListAttachedUserPolicies`: IamServiceListAttachedUserPoliciesResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceListAttachedUserPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceListAttachedUserPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userName** | **string** | The name of the user to list attached policies for. | 
 **marker** | **string** | Marker is obtained from paginated response from the previous query. Use this only if the response indicates it is truncated. | 
 **maxItems** | **int32** | Indicates the maximum number of elements to be returned in the response. | 
 **pathPrefix** | **string** | Path prefix for filtering the results. Optional, default to \&quot;/\&quot;. Only \&quot;/\&quot; is allowed. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceListAttachedUserPoliciesResponse**](IamServiceListAttachedUserPoliciesResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceListGroupPolicies

> IamServiceListGroupPoliciesResponse IamServiceListGroupPolicies(ctx).GroupName(groupName).Marker(marker).MaxItems(maxItems).XEmcNamespace(xEmcNamespace).Execute()

List Inline Policies for IAM Group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    groupName := "groupName_example" // string | The name of the group to list attached policies for. (optional)
    marker := "marker_example" // string | Marker is obtained from paginated response from the previous query. Use this only if the response indicates it is truncated. (optional)
    maxItems := int32(56) // int32 | Indicates the maximum number of elements to be returned in the response. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceListGroupPolicies(context.Background()).GroupName(groupName).Marker(marker).MaxItems(maxItems).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceListGroupPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceListGroupPolicies`: IamServiceListGroupPoliciesResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceListGroupPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceListGroupPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupName** | **string** | The name of the group to list attached policies for. | 
 **marker** | **string** | Marker is obtained from paginated response from the previous query. Use this only if the response indicates it is truncated. | 
 **maxItems** | **int32** | Indicates the maximum number of elements to be returned in the response. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceListGroupPoliciesResponse**](IamServiceListGroupPoliciesResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceListGroups

> IamServiceListGroupsResponse IamServiceListGroups(ctx).Marker(marker).MaxItems(maxItems).PathPrefix(pathPrefix).XEmcNamespace(xEmcNamespace).Execute()

Lists the IAM groups.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    marker := "marker_example" // string | Marker is obtained from paginated response from the previous query. Use this only if the response indicates it is truncated. (optional)
    maxItems := int32(56) // int32 | Indicates the maximum number of elements to be returned in the response. (optional)
    pathPrefix := "pathPrefix_example" // string | Path prefix for filtering the results. Optional, default to \"/\". Only \"/\" is allowed. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceListGroups(context.Background()).Marker(marker).MaxItems(maxItems).PathPrefix(pathPrefix).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceListGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceListGroups`: IamServiceListGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceListGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceListGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marker** | **string** | Marker is obtained from paginated response from the previous query. Use this only if the response indicates it is truncated. | 
 **maxItems** | **int32** | Indicates the maximum number of elements to be returned in the response. | 
 **pathPrefix** | **string** | Path prefix for filtering the results. Optional, default to \&quot;/\&quot;. Only \&quot;/\&quot; is allowed. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceListGroupsResponse**](IamServiceListGroupsResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceListGroupsForUser

> IamServiceListGroupsForUserResponse IamServiceListGroupsForUser(ctx).UserName(userName).Marker(marker).MaxItems(maxItems).XEmcNamespace(xEmcNamespace).Execute()

List Groups for IAM User



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    userName := "userName_example" // string | Simple name identifying the user. (optional)
    marker := "marker_example" // string | For pagination, the value of the Marker element in the response that you received to indicate where the next call should start. (optional)
    maxItems := int32(56) // int32 | Use this only when paginating results to indicate the maximum number of items you want in the response.  If additional items exist beyond the maximum you specify, the IsTruncated response element is true and  Marker contains a value to include in the subsequent call that tells the service where to continue from. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceListGroupsForUser(context.Background()).UserName(userName).Marker(marker).MaxItems(maxItems).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceListGroupsForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceListGroupsForUser`: IamServiceListGroupsForUserResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceListGroupsForUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceListGroupsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userName** | **string** | Simple name identifying the user. | 
 **marker** | **string** | For pagination, the value of the Marker element in the response that you received to indicate where the next call should start. | 
 **maxItems** | **int32** | Use this only when paginating results to indicate the maximum number of items you want in the response.  If additional items exist beyond the maximum you specify, the IsTruncated response element is true and  Marker contains a value to include in the subsequent call that tells the service where to continue from. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceListGroupsForUserResponse**](IamServiceListGroupsForUserResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceListPolicies

> IamServiceListPoliciesResponse IamServiceListPolicies(ctx).Marker(marker).MaxItems(maxItems).OnlyAttached(onlyAttached).PathPrefix(pathPrefix).PolicyUsageFilter(policyUsageFilter).PolicyScope(policyScope).XEmcNamespace(xEmcNamespace).Execute()

Lists the IAM users.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    marker := "marker_example" // string | Marker is obtained from paginated response from the previous query. Use this only if the response indicates it is truncated. (optional)
    maxItems := int32(56) // int32 | Indicates the maximum number of elements to be returned in the response. (optional)
    onlyAttached := true // bool | A flag to filter the results to only the attached policies. (optional)
    pathPrefix := "pathPrefix_example" // string | Path prefix for filtering the results. Optional, default to \"/\". Only \"/\" is allowed. (optional)
    policyUsageFilter := "policyUsageFilter_example" // string | The policy usage method to use for filtering the results. Values {PermissionsPolicy, PermissionsBoundary} (optional)
    policyScope := "policyScope_example" // string | The scope to use for filtering the results. One of {All, ECS, AWS, Local} (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceListPolicies(context.Background()).Marker(marker).MaxItems(maxItems).OnlyAttached(onlyAttached).PathPrefix(pathPrefix).PolicyUsageFilter(policyUsageFilter).PolicyScope(policyScope).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceListPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceListPolicies`: IamServiceListPoliciesResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceListPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceListPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marker** | **string** | Marker is obtained from paginated response from the previous query. Use this only if the response indicates it is truncated. | 
 **maxItems** | **int32** | Indicates the maximum number of elements to be returned in the response. | 
 **onlyAttached** | **bool** | A flag to filter the results to only the attached policies. | 
 **pathPrefix** | **string** | Path prefix for filtering the results. Optional, default to \&quot;/\&quot;. Only \&quot;/\&quot; is allowed. | 
 **policyUsageFilter** | **string** | The policy usage method to use for filtering the results. Values {PermissionsPolicy, PermissionsBoundary} | 
 **policyScope** | **string** | The scope to use for filtering the results. One of {All, ECS, AWS, Local} | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceListPoliciesResponse**](IamServiceListPoliciesResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceListRolePolicies

> IamServiceListRolePoliciesResponse IamServiceListRolePolicies(ctx).Marker(marker).MaxItems(maxItems).RoleName(roleName).XEmcNamespace(xEmcNamespace).Execute()

Lists the names of the inline policies that are embedded in the specified IAM role.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    marker := "marker_example" // string | For pagination, the value of the Marker element in the response that you received to indicate where the next call should start. (optional)
    maxItems := int32(56) // int32 | Use this only when paginating results to indicate the maximum number of items you want in the response.  If additional items exist beyond the maximum you specify, the IsTruncated response element is true and  Marker contains a value to include in the subsequent call that tells the service where to continue from. (optional)
    roleName := "roleName_example" // string | Simple name identifying the role. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceListRolePolicies(context.Background()).Marker(marker).MaxItems(maxItems).RoleName(roleName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceListRolePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceListRolePolicies`: IamServiceListRolePoliciesResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceListRolePolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceListRolePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marker** | **string** | For pagination, the value of the Marker element in the response that you received to indicate where the next call should start. | 
 **maxItems** | **int32** | Use this only when paginating results to indicate the maximum number of items you want in the response.  If additional items exist beyond the maximum you specify, the IsTruncated response element is true and  Marker contains a value to include in the subsequent call that tells the service where to continue from. | 
 **roleName** | **string** | Simple name identifying the role. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceListRolePoliciesResponse**](IamServiceListRolePoliciesResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceListRoleTags

> IamServiceListRoleTagsResponse IamServiceListRoleTags(ctx).Marker(marker).MaxItems(maxItems).RoleName(roleName).XEmcNamespace(xEmcNamespace).Execute()

Lists the tags that are attached to the specified IAM role.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    marker := "marker_example" // string | For pagination, the value of the Marker element in the response that you received to indicate where the next call should start. (optional)
    maxItems := int32(56) // int32 | Use this only when paginating results to indicate the maximum number of items you want in the response.  If additional items exist beyond the maximum you specify, the IsTruncated response element is true and  Marker contains a value to include in the subsequent call that tells the service where to continue from. (optional)
    roleName := "roleName_example" // string | Simple name identifying the role. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceListRoleTags(context.Background()).Marker(marker).MaxItems(maxItems).RoleName(roleName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceListRoleTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceListRoleTags`: IamServiceListRoleTagsResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceListRoleTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceListRoleTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marker** | **string** | For pagination, the value of the Marker element in the response that you received to indicate where the next call should start. | 
 **maxItems** | **int32** | Use this only when paginating results to indicate the maximum number of items you want in the response.  If additional items exist beyond the maximum you specify, the IsTruncated response element is true and  Marker contains a value to include in the subsequent call that tells the service where to continue from. | 
 **roleName** | **string** | Simple name identifying the role. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceListRoleTagsResponse**](IamServiceListRoleTagsResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceListRoles

> IamServiceListRolesResponse IamServiceListRoles(ctx).Marker(marker).MaxItems(maxItems).PathPrefix(pathPrefix).XEmcNamespace(xEmcNamespace).Execute()

Lists the IAM roles.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    marker := "marker_example" // string | For pagination, the value of the Marker element in the response that you received to indicate where the next call should start. (optional)
    maxItems := int32(56) // int32 | Use this only when paginating results to indicate the maximum number of items you want in the response.  If additional items exist beyond the maximum you specify, the IsTruncated response element is true and  Marker contains a value to include in the subsequent call that tells the service where to continue from. (optional)
    pathPrefix := "pathPrefix_example" // string | The path to the roles. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceListRoles(context.Background()).Marker(marker).MaxItems(maxItems).PathPrefix(pathPrefix).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceListRoles`: IamServiceListRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceListRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marker** | **string** | For pagination, the value of the Marker element in the response that you received to indicate where the next call should start. | 
 **maxItems** | **int32** | Use this only when paginating results to indicate the maximum number of items you want in the response.  If additional items exist beyond the maximum you specify, the IsTruncated response element is true and  Marker contains a value to include in the subsequent call that tells the service where to continue from. | 
 **pathPrefix** | **string** | The path to the roles. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceListRolesResponse**](IamServiceListRolesResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceListUserPolicies

> IamServiceListUserPoliciesResponse IamServiceListUserPolicies(ctx).UserName(userName).Marker(marker).MaxItems(maxItems).XEmcNamespace(xEmcNamespace).Execute()

List Inline Policies for IAM User.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    userName := "userName_example" // string | The name of the user to list attached policies for. (optional)
    marker := "marker_example" // string | Marker is obtained from paginated response from the previous query. Use this only if the response indicates it is truncated. (optional)
    maxItems := int32(56) // int32 | Indicates the maximum number of elements to be returned in the response. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceListUserPolicies(context.Background()).UserName(userName).Marker(marker).MaxItems(maxItems).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceListUserPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceListUserPolicies`: IamServiceListUserPoliciesResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceListUserPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceListUserPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userName** | **string** | The name of the user to list attached policies for. | 
 **marker** | **string** | Marker is obtained from paginated response from the previous query. Use this only if the response indicates it is truncated. | 
 **maxItems** | **int32** | Indicates the maximum number of elements to be returned in the response. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceListUserPoliciesResponse**](IamServiceListUserPoliciesResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceListUserTags

> IamServiceListUserTagsResponse IamServiceListUserTags(ctx).UserName(userName).Marker(marker).MaxItems(maxItems).XEmcNamespace(xEmcNamespace).Execute()

Lists the tags that are attached to the specified IAM User.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    userName := "userName_example" // string | Simple name identifying the user. (optional)
    marker := "marker_example" // string | For pagination, the value of the Marker element in the response that you received to indicate where the next call should start. (optional)
    maxItems := int32(56) // int32 | Use this only when paginating results to indicate the maximum number of items you want in the response.  If additional items exist beyond the maximum you specify, the IsTruncated response element is true and  Marker contains a value to include in the subsequent call that tells the service where to continue from. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceListUserTags(context.Background()).UserName(userName).Marker(marker).MaxItems(maxItems).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceListUserTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceListUserTags`: IamServiceListUserTagsResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceListUserTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceListUserTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userName** | **string** | Simple name identifying the user. | 
 **marker** | **string** | For pagination, the value of the Marker element in the response that you received to indicate where the next call should start. | 
 **maxItems** | **int32** | Use this only when paginating results to indicate the maximum number of items you want in the response.  If additional items exist beyond the maximum you specify, the IsTruncated response element is true and  Marker contains a value to include in the subsequent call that tells the service where to continue from. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceListUserTagsResponse**](IamServiceListUserTagsResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceListUsers

> IamServiceListUsersResponse IamServiceListUsers(ctx).Marker(marker).MaxItems(maxItems).PathPrefix(pathPrefix).XEmcNamespace(xEmcNamespace).Execute()

Lists the IAM users.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    marker := "marker_example" // string | Marker is obtained from paginated response from the previous query. Use this only if the response indicates it is truncated. (optional)
    maxItems := int32(56) // int32 | Indicates the maximum number of elements to be returned in the response. (optional)
    pathPrefix := "pathPrefix_example" // string | Path prefix for filtering the results. Optional, default to \"/\". Only \"/\" is allowed. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceListUsers(context.Background()).Marker(marker).MaxItems(maxItems).PathPrefix(pathPrefix).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceListUsers`: IamServiceListUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marker** | **string** | Marker is obtained from paginated response from the previous query. Use this only if the response indicates it is truncated. | 
 **maxItems** | **int32** | Indicates the maximum number of elements to be returned in the response. | 
 **pathPrefix** | **string** | Path prefix for filtering the results. Optional, default to \&quot;/\&quot;. Only \&quot;/\&quot; is allowed. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceListUsersResponse**](IamServiceListUsersResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServicePutGroupPolicy

> IamServicePutGroupPolicyResponse IamServicePutGroupPolicy(ctx).PolicyDocument(policyDocument).GroupName(groupName).PolicyName(policyName).XEmcNamespace(xEmcNamespace).Execute()

Add or Update Inline Policy for IAM Group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    policyDocument := "policyDocument_example" // string | The policy document in JSON format. (optional)
    groupName := "groupName_example" // string | Simple name identifying the group. (optional)
    policyName := "policyName_example" // string | Simple name identifying the policy. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServicePutGroupPolicy(context.Background()).PolicyDocument(policyDocument).GroupName(groupName).PolicyName(policyName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServicePutGroupPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServicePutGroupPolicy`: IamServicePutGroupPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServicePutGroupPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServicePutGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyDocument** | **string** | The policy document in JSON format. | 
 **groupName** | **string** | Simple name identifying the group. | 
 **policyName** | **string** | Simple name identifying the policy. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServicePutGroupPolicyResponse**](IamServicePutGroupPolicyResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServicePutRolePermissionsBoundary

> IamServicePutRolePermissionsBoundaryResponse IamServicePutRolePermissionsBoundary(ctx).PolicyArn(policyArn).RoleName(roleName).XEmcNamespace(xEmcNamespace).Execute()

Adds or updates the policy that is specified as the IAM role's permissions boundary.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    policyArn := "policyArn_example" // string | Arn that identifies the policy. (optional)
    roleName := "roleName_example" // string | Simple name identifying the role. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServicePutRolePermissionsBoundary(context.Background()).PolicyArn(policyArn).RoleName(roleName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServicePutRolePermissionsBoundary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServicePutRolePermissionsBoundary`: IamServicePutRolePermissionsBoundaryResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServicePutRolePermissionsBoundary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServicePutRolePermissionsBoundaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyArn** | **string** | Arn that identifies the policy. | 
 **roleName** | **string** | Simple name identifying the role. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServicePutRolePermissionsBoundaryResponse**](IamServicePutRolePermissionsBoundaryResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServicePutRolePolicy

> IamServicePutRolePolicyResponse IamServicePutRolePolicy(ctx).PolicyDocument(policyDocument).RoleName(roleName).PolicyName(policyName).XEmcNamespace(xEmcNamespace).Execute()

Adds or updates an inline policy document that is embedded in the specified IAM role.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    policyDocument := "policyDocument_example" // string | The policy document in JSON format. (optional)
    roleName := "roleName_example" // string | Simple name identifying the role. (optional)
    policyName := "policyName_example" // string | Simple name identifying the policy. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServicePutRolePolicy(context.Background()).PolicyDocument(policyDocument).RoleName(roleName).PolicyName(policyName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServicePutRolePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServicePutRolePolicy`: IamServicePutRolePolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServicePutRolePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServicePutRolePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyDocument** | **string** | The policy document in JSON format. | 
 **roleName** | **string** | Simple name identifying the role. | 
 **policyName** | **string** | Simple name identifying the policy. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServicePutRolePolicyResponse**](IamServicePutRolePolicyResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServicePutUserPolicy

> IamServicePutUserPolicyResponse IamServicePutUserPolicy(ctx).PolicyDocument(policyDocument).PolicyName(policyName).UserName(userName).XEmcNamespace(xEmcNamespace).Execute()

Add or Update Inline Policy for IAM User.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    policyDocument := "policyDocument_example" // string | The policy document in JSON format. (optional)
    policyName := "policyName_example" // string | Simple name identifying the policy. (optional)
    userName := "userName_example" // string | Simple name identifying the user. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServicePutUserPolicy(context.Background()).PolicyDocument(policyDocument).PolicyName(policyName).UserName(userName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServicePutUserPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServicePutUserPolicy`: IamServicePutUserPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServicePutUserPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServicePutUserPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyDocument** | **string** | The policy document in JSON format. | 
 **policyName** | **string** | Simple name identifying the policy. | 
 **userName** | **string** | Simple name identifying the user. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServicePutUserPolicyResponse**](IamServicePutUserPolicyResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceRemoveUserFromGroup

> IamServiceRemoveUserFromGroupResponse IamServiceRemoveUserFromGroup(ctx).GroupName(groupName).UserName(userName).XEmcNamespace(xEmcNamespace).Execute()

Remove User from a Group.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    groupName := "groupName_example" // string | The name of the group to update. (optional)
    userName := "userName_example" // string | The name of the user to be removed. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceRemoveUserFromGroup(context.Background()).GroupName(groupName).UserName(userName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceRemoveUserFromGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceRemoveUserFromGroup`: IamServiceRemoveUserFromGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceRemoveUserFromGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceRemoveUserFromGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupName** | **string** | The name of the group to update. | 
 **userName** | **string** | The name of the user to be removed. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceRemoveUserFromGroupResponse**](IamServiceRemoveUserFromGroupResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceTagRole

> map[string]interface{} IamServiceTagRole(ctx).RoleName(roleName).TagsMemberN(tagsMemberN).XEmcNamespace(xEmcNamespace).Execute()

Adds one or more tags to a specified IAM Role.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    roleName := "roleName_example" // string | Simple name identifying the role. (optional)
    tagsMemberN := map[string]interface{}{ ... } // map[string]interface{} | A list of tags that you want to attach to the role. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceTagRole(context.Background()).RoleName(roleName).TagsMemberN(tagsMemberN).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceTagRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceTagRole`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceTagRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceTagRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleName** | **string** | Simple name identifying the role. | 
 **tagsMemberN** | [**map[string]interface{}**](map[string]interface{}.md) | A list of tags that you want to attach to the role. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

**map[string]interface{}**

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceUntagRole

> map[string]interface{} IamServiceUntagRole(ctx).RoleName(roleName).TagKeys(tagKeys).XEmcNamespace(xEmcNamespace).Execute()

Removes the specified tags from a specified IAM Role.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    roleName := "roleName_example" // string | Simple name identifying the role. (optional)
    tagKeys := map[string][]openapiclient.IamServiceUntagRoleTagKeysParameter{ ... } // IamServiceUntagRoleTagKeysParameter | A list of tags that you want to remove from the role. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceUntagRole(context.Background()).RoleName(roleName).TagKeys(tagKeys).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceUntagRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceUntagRole`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceUntagRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceUntagRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleName** | **string** | Simple name identifying the role. | 
 **tagKeys** | [**IamServiceUntagRoleTagKeysParameter**](IamServiceUntagRoleTagKeysParameter.md) | A list of tags that you want to remove from the role. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

**map[string]interface{}**

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceUpdateAssumeRolePolicy

> IamServiceUpdateAssumeRolePolicyResponse IamServiceUpdateAssumeRolePolicy(ctx).PolicyDocument(policyDocument).RoleName(roleName).XEmcNamespace(xEmcNamespace).Execute()

Updates the policy that grants an IAM entity permission to assume a role.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    policyDocument := "policyDocument_example" // string | The policy that grants an entity permission to assume the role. (optional)
    roleName := "roleName_example" // string | Simple name identifying the role. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceUpdateAssumeRolePolicy(context.Background()).PolicyDocument(policyDocument).RoleName(roleName).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceUpdateAssumeRolePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceUpdateAssumeRolePolicy`: IamServiceUpdateAssumeRolePolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceUpdateAssumeRolePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceUpdateAssumeRolePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyDocument** | **string** | The policy that grants an entity permission to assume the role. | 
 **roleName** | **string** | Simple name identifying the role. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceUpdateAssumeRolePolicyResponse**](IamServiceUpdateAssumeRolePolicyResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamServiceUpdateRole

> IamServiceUpdateRoleResponse IamServiceUpdateRole(ctx).RoleName(roleName).MaxSessionDuration(maxSessionDuration).Description(description).XEmcNamespace(xEmcNamespace).Execute()

Updates the description or maximum session duration setting of the specified IAM role.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    roleName := "roleName_example" // string | Simple name identifying the role. (optional)
    maxSessionDuration := int32(56) // int32 | The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied.  This setting can have a value from 1 hour to 12 hours (optional)
    description := "description_example" // string | The new description that you want to apply to the specified role. (optional)
    xEmcNamespace := "xEmcNamespace_example" // string | ECS namespace IAM entity belongs to, only required when request performed by management user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamApi.IamServiceUpdateRole(context.Background()).RoleName(roleName).MaxSessionDuration(maxSessionDuration).Description(description).XEmcNamespace(xEmcNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamApi.IamServiceUpdateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IamServiceUpdateRole`: IamServiceUpdateRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `IamApi.IamServiceUpdateRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamServiceUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleName** | **string** | Simple name identifying the role. | 
 **maxSessionDuration** | **int32** | The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied.  This setting can have a value from 1 hour to 12 hours | 
 **description** | **string** | The new description that you want to apply to the specified role. | 
 **xEmcNamespace** | **string** | ECS namespace IAM entity belongs to, only required when request performed by management user | 

### Return type

[**IamServiceUpdateRoleResponse**](IamServiceUpdateRoleResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

