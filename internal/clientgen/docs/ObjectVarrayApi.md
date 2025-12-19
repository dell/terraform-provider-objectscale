# \ObjectVarrayApi

All URIs are relative to *https://objectscale.local:4443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ObjectVarrayServiceCreateVirtualArray**](ObjectVarrayApi.md#ObjectVarrayServiceCreateVirtualArray) | **Post** /vdc/data-services/varrays | Create a storage pool with the specified details
[**ObjectVarrayServiceDeleteVirtualArray**](ObjectVarrayApi.md#ObjectVarrayServiceDeleteVirtualArray) | **Delete** /vdc/data-services/varrays/{id} | Deletes the storage pool for the specified identifier
[**ObjectVarrayServiceGetVirtualArray**](ObjectVarrayApi.md#ObjectVarrayServiceGetVirtualArray) | **Get** /vdc/data-services/varrays/{id} | Gets the details for the specified storage pool
[**ObjectVarrayServiceGetVirtualArrays**](ObjectVarrayApi.md#ObjectVarrayServiceGetVirtualArrays) | **Get** /vdc/data-services/varrays | Gets a list of storage pools from the local VDC
[**ObjectVarrayServiceUpdateVirtualArray**](ObjectVarrayApi.md#ObjectVarrayServiceUpdateVirtualArray) | **Put** /vdc/data-services/varrays/{id} | Updates storage pool for the specified identifier



## ObjectVarrayServiceCreateVirtualArray

> ObjectVarrayServiceCreateVirtualArrayResponse ObjectVarrayServiceCreateVirtualArray(ctx).ObjectVarrayServiceCreateVirtualArrayRequest(objectVarrayServiceCreateVirtualArrayRequest).Execute()

Create a storage pool with the specified details



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
    objectVarrayServiceCreateVirtualArrayRequest := *openapiclient.NewObjectVarrayServiceCreateVirtualArrayRequest("Name_example", false) // ObjectVarrayServiceCreateVirtualArrayRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectVarrayApi.ObjectVarrayServiceCreateVirtualArray(context.Background()).ObjectVarrayServiceCreateVirtualArrayRequest(objectVarrayServiceCreateVirtualArrayRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectVarrayApi.ObjectVarrayServiceCreateVirtualArray``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ObjectVarrayServiceCreateVirtualArray`: ObjectVarrayServiceCreateVirtualArrayResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectVarrayApi.ObjectVarrayServiceCreateVirtualArray`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiObjectVarrayServiceCreateVirtualArrayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectVarrayServiceCreateVirtualArrayRequest** | [**ObjectVarrayServiceCreateVirtualArrayRequest**](ObjectVarrayServiceCreateVirtualArrayRequest.md) |  | 

### Return type

[**ObjectVarrayServiceCreateVirtualArrayResponse**](ObjectVarrayServiceCreateVirtualArrayResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObjectVarrayServiceDeleteVirtualArray

> map[string]interface{} ObjectVarrayServiceDeleteVirtualArray(ctx, id).Execute()

Deletes the storage pool for the specified identifier



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
    id := "id_example" // string | storage pool identifier to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectVarrayApi.ObjectVarrayServiceDeleteVirtualArray(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectVarrayApi.ObjectVarrayServiceDeleteVirtualArray``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ObjectVarrayServiceDeleteVirtualArray`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ObjectVarrayApi.ObjectVarrayServiceDeleteVirtualArray`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | storage pool identifier to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiObjectVarrayServiceDeleteVirtualArrayRequest struct via the builder pattern


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


## ObjectVarrayServiceGetVirtualArray

> ObjectVarrayServiceGetVirtualArrayResponse ObjectVarrayServiceGetVirtualArray(ctx, id).Execute()

Gets the details for the specified storage pool



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
    id := "id_example" // string | Storage pool identifier to be retrieved

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectVarrayApi.ObjectVarrayServiceGetVirtualArray(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectVarrayApi.ObjectVarrayServiceGetVirtualArray``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ObjectVarrayServiceGetVirtualArray`: ObjectVarrayServiceGetVirtualArrayResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectVarrayApi.ObjectVarrayServiceGetVirtualArray`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Storage pool identifier to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiObjectVarrayServiceGetVirtualArrayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectVarrayServiceGetVirtualArrayResponse**](ObjectVarrayServiceGetVirtualArrayResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObjectVarrayServiceGetVirtualArrays

> ObjectVarrayServiceGetVirtualArraysResponse ObjectVarrayServiceGetVirtualArrays(ctx).VdcId(vdcId).Execute()

Gets a list of storage pools from the local VDC



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
    vdcId := "vdcId_example" // string | virtual data center identifier for which list of storage poold is to be retrieved (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectVarrayApi.ObjectVarrayServiceGetVirtualArrays(context.Background()).VdcId(vdcId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectVarrayApi.ObjectVarrayServiceGetVirtualArrays``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ObjectVarrayServiceGetVirtualArrays`: ObjectVarrayServiceGetVirtualArraysResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectVarrayApi.ObjectVarrayServiceGetVirtualArrays`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiObjectVarrayServiceGetVirtualArraysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vdcId** | **string** | virtual data center identifier for which list of storage poold is to be retrieved | 

### Return type

[**ObjectVarrayServiceGetVirtualArraysResponse**](ObjectVarrayServiceGetVirtualArraysResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObjectVarrayServiceUpdateVirtualArray

> ObjectVarrayServiceUpdateVirtualArrayResponse ObjectVarrayServiceUpdateVirtualArray(ctx, id).ObjectVarrayServiceUpdateVirtualArrayRequest(objectVarrayServiceUpdateVirtualArrayRequest).Execute()

Updates storage pool for the specified identifier



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
    id := "id_example" // string | Storage pool identifier to be updated
    objectVarrayServiceUpdateVirtualArrayRequest := *openapiclient.NewObjectVarrayServiceUpdateVirtualArrayRequest("Name_example", false) // ObjectVarrayServiceUpdateVirtualArrayRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectVarrayApi.ObjectVarrayServiceUpdateVirtualArray(context.Background(), id).ObjectVarrayServiceUpdateVirtualArrayRequest(objectVarrayServiceUpdateVirtualArrayRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectVarrayApi.ObjectVarrayServiceUpdateVirtualArray``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ObjectVarrayServiceUpdateVirtualArray`: ObjectVarrayServiceUpdateVirtualArrayResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectVarrayApi.ObjectVarrayServiceUpdateVirtualArray`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Storage pool identifier to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiObjectVarrayServiceUpdateVirtualArrayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectVarrayServiceUpdateVirtualArrayRequest** | [**ObjectVarrayServiceUpdateVirtualArrayRequest**](ObjectVarrayServiceUpdateVirtualArrayRequest.md) |  | 

### Return type

[**ObjectVarrayServiceUpdateVirtualArrayResponse**](ObjectVarrayServiceUpdateVirtualArrayResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

