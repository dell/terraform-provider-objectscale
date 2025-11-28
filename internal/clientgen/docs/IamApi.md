# \IamApi

All URIs are relative to *https://objectscale.local:4443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamServiceAddUserToGroup**](IamApi.md#IamServiceAddUserToGroup) | **Post** /iam?Action&#x3D;AddUserToGroup | Add user to a group.
[**IamServiceAttachGroupPolicy**](IamApi.md#IamServiceAttachGroupPolicy) | **Post** /iam?Action&#x3D;AttachGroupPolicy | Attach a Managed Policy to Group.
[**IamServiceCreateGroup**](IamApi.md#IamServiceCreateGroup) | **Post** /iam?Action&#x3D;CreateGroup | Creates a new IAM Group.
[**IamServiceDeleteGroup**](IamApi.md#IamServiceDeleteGroup) | **Post** /iam?Action&#x3D;DeleteGroup | Delete an IAM Group.
[**IamServiceDeleteGroupPolicy**](IamApi.md#IamServiceDeleteGroupPolicy) | **Post** /iam?Action&#x3D;DeleteGroupPolicy | Delete specific inlinePolicy for IAM Group.
[**IamServiceDetachGroupPolicy**](IamApi.md#IamServiceDetachGroupPolicy) | **Post** /iam?Action&#x3D;DetachGroupPolicy | Remove a Managed Policy attached to Group.
[**IamServiceGetGroup**](IamApi.md#IamServiceGetGroup) | **Post** /iam?Action&#x3D;GetGroup | Retrieve list of users in IAM group.
[**IamServiceGetGroupPolicy**](IamApi.md#IamServiceGetGroupPolicy) | **Post** /iam?Action&#x3D;GetGroupPolicy | Get specific inlinePolicy for IAM Group.
[**IamServiceGetUser**](IamApi.md#IamServiceGetUser) | **Post** /iam?Action&#x3D;GetUser | Retrieve IAM user.
[**IamServiceListAccessKeys**](IamApi.md#IamServiceListAccessKeys) | **Post** /iam?Action&#x3D;ListAccessKeys | List AccessKeys for a user.
[**IamServiceListAttachedGroupPolicies**](IamApi.md#IamServiceListAttachedGroupPolicies) | **Post** /iam?Action&#x3D;ListAttachedGroupPolicies | List Managed Policies for IAM Group.
[**IamServiceListGroupPolicies**](IamApi.md#IamServiceListGroupPolicies) | **Post** /iam?Action&#x3D;ListGroupPolicies | List Inline Policies for IAM Group.
[**IamServiceListGroups**](IamApi.md#IamServiceListGroups) | **Post** /iam?Action&#x3D;ListGroups | Lists the IAM groups.
[**IamServiceListGroupsForUser**](IamApi.md#IamServiceListGroupsForUser) | **Post** /iam?Action&#x3D;ListGroupsForUser | List Groups for IAM User
[**IamServiceListUserTags**](IamApi.md#IamServiceListUserTags) | **Post** /iam?Action&#x3D;ListUserTags | Lists the tags that are attached to the specified IAM User.
[**IamServiceListUsers**](IamApi.md#IamServiceListUsers) | **Post** /iam?Action&#x3D;ListUsers | Lists the IAM users.
[**IamServicePutGroupPolicy**](IamApi.md#IamServicePutGroupPolicy) | **Post** /iam?Action&#x3D;PutGroupPolicy | Add or Update Inline Policy for IAM Group.
[**IamServiceRemoveUserFromGroup**](IamApi.md#IamServiceRemoveUserFromGroup) | **Post** /iam?Action&#x3D;RemoveUserFromGroup | Remove User from a Group.



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

