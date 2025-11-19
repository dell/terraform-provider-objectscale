# \IamApi

All URIs are relative to *https://objectscale.local:4443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamServiceGetGroup**](IamApi.md#IamServiceGetGroup) | **Post** /iam?Action&#x3D;GetGroup | Retrieve list of users in IAM group.
[**IamServiceListGroups**](IamApi.md#IamServiceListGroups) | **Post** /iam?Action&#x3D;ListGroups | Lists the IAM groups.



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

