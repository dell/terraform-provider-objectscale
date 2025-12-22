# \MgmtUserInfoApi

All URIs are relative to *https://objectscale.local:4443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MgmtUserInfoServiceCreateLocalUserInfo**](MgmtUserInfoApi.md#MgmtUserInfoServiceCreateLocalUserInfo) | **Post** /vdc/users | Creates a local VDC user with the specified details
[**MgmtUserInfoServiceDeleteLocalUserInfo**](MgmtUserInfoApi.md#MgmtUserInfoServiceDeleteLocalUserInfo) | **Post** /vdc/users/{userid}/deactivate | Deletes local user information for the specified user identifier
[**MgmtUserInfoServiceGetLocalUserInfo**](MgmtUserInfoApi.md#MgmtUserInfoServiceGetLocalUserInfo) | **Get** /vdc/users/{userid} | Gets local user details for the specified user identifier
[**MgmtUserInfoServiceGetLocalUserInfos**](MgmtUserInfoApi.md#MgmtUserInfoServiceGetLocalUserInfos) | **Get** /vdc/users | Lists all local management users
[**MgmtUserInfoServiceModifyLocalUserInfo**](MgmtUserInfoApi.md#MgmtUserInfoServiceModifyLocalUserInfo) | **Put** /vdc/users/{userid} | Updates local user details for the specified user identifier



## MgmtUserInfoServiceCreateLocalUserInfo

> MgmtUserInfoServiceCreateLocalUserInfoResponse MgmtUserInfoServiceCreateLocalUserInfo(ctx).MgmtUserInfoServiceCreateLocalUserInfoRequest(mgmtUserInfoServiceCreateLocalUserInfoRequest).Execute()

Creates a local VDC user with the specified details



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
    mgmtUserInfoServiceCreateLocalUserInfoRequest := *openapiclient.NewMgmtUserInfoServiceCreateLocalUserInfoRequest("UserId_example") // MgmtUserInfoServiceCreateLocalUserInfoRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MgmtUserInfoApi.MgmtUserInfoServiceCreateLocalUserInfo(context.Background()).MgmtUserInfoServiceCreateLocalUserInfoRequest(mgmtUserInfoServiceCreateLocalUserInfoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MgmtUserInfoApi.MgmtUserInfoServiceCreateLocalUserInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MgmtUserInfoServiceCreateLocalUserInfo`: MgmtUserInfoServiceCreateLocalUserInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `MgmtUserInfoApi.MgmtUserInfoServiceCreateLocalUserInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMgmtUserInfoServiceCreateLocalUserInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mgmtUserInfoServiceCreateLocalUserInfoRequest** | [**MgmtUserInfoServiceCreateLocalUserInfoRequest**](MgmtUserInfoServiceCreateLocalUserInfoRequest.md) |  | 

### Return type

[**MgmtUserInfoServiceCreateLocalUserInfoResponse**](MgmtUserInfoServiceCreateLocalUserInfoResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MgmtUserInfoServiceDeleteLocalUserInfo

> map[string]interface{} MgmtUserInfoServiceDeleteLocalUserInfo(ctx, userid).Execute()

Deletes local user information for the specified user identifier



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
    userid := "userid_example" // string | User identifier for which local user information needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MgmtUserInfoApi.MgmtUserInfoServiceDeleteLocalUserInfo(context.Background(), userid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MgmtUserInfoApi.MgmtUserInfoServiceDeleteLocalUserInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MgmtUserInfoServiceDeleteLocalUserInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MgmtUserInfoApi.MgmtUserInfoServiceDeleteLocalUserInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** | User identifier for which local user information needs to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMgmtUserInfoServiceDeleteLocalUserInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## MgmtUserInfoServiceGetLocalUserInfo

> MgmtUserInfoServiceGetLocalUserInfoResponse MgmtUserInfoServiceGetLocalUserInfo(ctx, userid).Execute()

Gets local user details for the specified user identifier



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
    userid := "userid_example" // string | User identifier for which local user information needs to be retrieved

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MgmtUserInfoApi.MgmtUserInfoServiceGetLocalUserInfo(context.Background(), userid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MgmtUserInfoApi.MgmtUserInfoServiceGetLocalUserInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MgmtUserInfoServiceGetLocalUserInfo`: MgmtUserInfoServiceGetLocalUserInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `MgmtUserInfoApi.MgmtUserInfoServiceGetLocalUserInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** | User identifier for which local user information needs to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiMgmtUserInfoServiceGetLocalUserInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MgmtUserInfoServiceGetLocalUserInfoResponse**](MgmtUserInfoServiceGetLocalUserInfoResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MgmtUserInfoServiceGetLocalUserInfos

> MgmtUserInfoServiceGetLocalUserInfosResponse MgmtUserInfoServiceGetLocalUserInfos(ctx).Execute()

Lists all local management users



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MgmtUserInfoApi.MgmtUserInfoServiceGetLocalUserInfos(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MgmtUserInfoApi.MgmtUserInfoServiceGetLocalUserInfos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MgmtUserInfoServiceGetLocalUserInfos`: MgmtUserInfoServiceGetLocalUserInfosResponse
    fmt.Fprintf(os.Stdout, "Response from `MgmtUserInfoApi.MgmtUserInfoServiceGetLocalUserInfos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMgmtUserInfoServiceGetLocalUserInfosRequest struct via the builder pattern


### Return type

[**MgmtUserInfoServiceGetLocalUserInfosResponse**](MgmtUserInfoServiceGetLocalUserInfosResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MgmtUserInfoServiceModifyLocalUserInfo

> map[string]interface{} MgmtUserInfoServiceModifyLocalUserInfo(ctx, userid).MgmtUserInfoServiceModifyLocalUserInfoRequest(mgmtUserInfoServiceModifyLocalUserInfoRequest).Execute()

Updates local user details for the specified user identifier



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
    userid := "userid_example" // string | User identifier for which local user information needs to be updated.
    mgmtUserInfoServiceModifyLocalUserInfoRequest := *openapiclient.NewMgmtUserInfoServiceModifyLocalUserInfoRequest() // MgmtUserInfoServiceModifyLocalUserInfoRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MgmtUserInfoApi.MgmtUserInfoServiceModifyLocalUserInfo(context.Background(), userid).MgmtUserInfoServiceModifyLocalUserInfoRequest(mgmtUserInfoServiceModifyLocalUserInfoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MgmtUserInfoApi.MgmtUserInfoServiceModifyLocalUserInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MgmtUserInfoServiceModifyLocalUserInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MgmtUserInfoApi.MgmtUserInfoServiceModifyLocalUserInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** | User identifier for which local user information needs to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMgmtUserInfoServiceModifyLocalUserInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mgmtUserInfoServiceModifyLocalUserInfoRequest** | [**MgmtUserInfoServiceModifyLocalUserInfoRequest**](MgmtUserInfoServiceModifyLocalUserInfoRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

