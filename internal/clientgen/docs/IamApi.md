# \IamApi

All URIs are relative to *https://objectscale.local:4443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamServiceDeleteGroupPolicy**](IamApi.md#IamServiceDeleteGroupPolicy) | **Post** /iam?Action&#x3D;DeleteGroupPolicy | Delete specific inlinePolicy for IAM Group.
[**IamServiceDeleteRolePolicy**](IamApi.md#IamServiceDeleteRolePolicy) | **Post** /iam?Action&#x3D;DeleteRolePolicy | Deletes the specified inline policy that is embedded in the specified IAM role.
[**IamServiceDeleteUserPolicy**](IamApi.md#IamServiceDeleteUserPolicy) | **Post** /iam?Action&#x3D;DeleteUserPolicy | Delete specific inlinePolicy for IAM User.
[**IamServiceGetGroup**](IamApi.md#IamServiceGetGroup) | **Post** /iam?Action&#x3D;GetGroup | Retrieve list of users in IAM group.
[**IamServiceGetGroupPolicy**](IamApi.md#IamServiceGetGroupPolicy) | **Post** /iam?Action&#x3D;GetGroupPolicy | Get specific inlinePolicy for IAM Group.
[**IamServiceGetRolePolicy**](IamApi.md#IamServiceGetRolePolicy) | **Post** /iam?Action&#x3D;GetRolePolicy | Gets tthe specified inline policy document that is embedded with the specified IAM role.
[**IamServiceGetUser**](IamApi.md#IamServiceGetUser) | **Post** /iam?Action&#x3D;GetUser | Retrieve IAM user.
[**IamServiceGetUserPolicy**](IamApi.md#IamServiceGetUserPolicy) | **Post** /iam?Action&#x3D;GetUserPolicy | Get specific inlinePolicy for IAM User.
[**IamServiceListAccessKeys**](IamApi.md#IamServiceListAccessKeys) | **Post** /iam?Action&#x3D;ListAccessKeys | List AccessKeys for a user.
[**IamServiceListGroupPolicies**](IamApi.md#IamServiceListGroupPolicies) | **Post** /iam?Action&#x3D;ListGroupPolicies | List Inline Policies for IAM Group.
[**IamServiceListGroups**](IamApi.md#IamServiceListGroups) | **Post** /iam?Action&#x3D;ListGroups | Lists the IAM groups.
[**IamServiceListGroupsForUser**](IamApi.md#IamServiceListGroupsForUser) | **Post** /iam?Action&#x3D;ListGroupsForUser | List Groups for IAM User
[**IamServiceListRolePolicies**](IamApi.md#IamServiceListRolePolicies) | **Post** /iam?Action&#x3D;ListRolePolicies | Lists the names of the inline policies that are embedded in the specified IAM role.
[**IamServiceListUserPolicies**](IamApi.md#IamServiceListUserPolicies) | **Post** /iam?Action&#x3D;ListUserPolicies | List Inline Policies for IAM User.
[**IamServiceListUserTags**](IamApi.md#IamServiceListUserTags) | **Post** /iam?Action&#x3D;ListUserTags | Lists the tags that are attached to the specified IAM User.
[**IamServiceListUsers**](IamApi.md#IamServiceListUsers) | **Post** /iam?Action&#x3D;ListUsers | Lists the IAM users.
[**IamServicePutGroupPolicy**](IamApi.md#IamServicePutGroupPolicy) | **Post** /iam?Action&#x3D;PutGroupPolicy | Add or Update Inline Policy for IAM Group.
[**IamServicePutRolePolicy**](IamApi.md#IamServicePutRolePolicy) | **Post** /iam?Action&#x3D;PutRolePolicy | Adds or updates an inline policy document that is embedded in the specified IAM role.
[**IamServicePutUserPolicy**](IamApi.md#IamServicePutUserPolicy) | **Post** /iam?Action&#x3D;PutUserPolicy | Add or Update Inline Policy for IAM User.



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

