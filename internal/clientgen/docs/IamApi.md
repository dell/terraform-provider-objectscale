# \IamApi

All URIs are relative to *https://objectscale.local:4443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamServiceGetGroup**](IamApi.md#IamServiceGetGroup) | **Post** /iam?Action&#x3D;GetGroup | Retrieve list of users in IAM group.
[**IamServiceGetUser**](IamApi.md#IamServiceGetUser) | **Post** /iam?Action&#x3D;GetUser | Retrieve IAM user.
[**IamServiceListAccessKeys**](IamApi.md#IamServiceListAccessKeys) | **Post** /iam?Action&#x3D;ListAccessKeys | List AccessKeys for a user.
[**IamServiceListUserTags**](IamApi.md#IamServiceListUserTags) | **Post** /iam?Action&#x3D;ListUserTags | Lists the tags that are attached to the specified IAM User.
[**IamServiceListUsers**](IamApi.md#IamServiceListUsers) | **Post** /iam?Action&#x3D;ListUsers | Lists the IAM users.



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

