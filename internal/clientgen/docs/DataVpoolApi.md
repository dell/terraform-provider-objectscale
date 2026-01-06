# \DataVpoolApi

All URIs are relative to *https://objectscale.local:4443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataServiceVpoolServiceAddToVpool**](DataVpoolApi.md#DataServiceVpoolServiceAddToVpool) | **Put** /vdc/data-service/vpools/{id}/addvarrays | Adds one or more storage pools (as  VDC:storage pool tuples) to the specified replication group
[**DataServiceVpoolServiceCreateDataServiceVpool**](DataVpoolApi.md#DataServiceVpoolServiceCreateDataServiceVpool) | **Post** /vdc/data-service/vpools | Creates a replication group that includes the specified storage pools (VDC:storage pool tuple)
[**DataServiceVpoolServiceGetDataServiceStore**](DataVpoolApi.md#DataServiceVpoolServiceGetDataServiceStore) | **Get** /vdc/data-service/vpools/{id} | Gets the details for the specified replication group
[**DataServiceVpoolServiceGetDataServiceVpools**](DataVpoolApi.md#DataServiceVpoolServiceGetDataServiceVpools) | **Get** /vdc/data-service/vpools | Lists all configured replication groups
[**DataServiceVpoolServicePutDataServiceVpool**](DataVpoolApi.md#DataServiceVpoolServicePutDataServiceVpool) | **Put** /vdc/data-service/vpools/{id} | Updates the name and description for a replication group
[**DataServiceVpoolServiceRemoveFromVpool**](DataVpoolApi.md#DataServiceVpoolServiceRemoveFromVpool) | **Put** /vdc/data-service/vpools/{id}/removevarrays | Deletes a storage pool (VDC:storage pool tuple) from a specified replication group



## DataServiceVpoolServiceAddToVpool

> map[string]interface{} DataServiceVpoolServiceAddToVpool(ctx, id).DataServiceVpoolServiceAddToVpoolRequest(dataServiceVpoolServiceAddToVpoolRequest).Execute()

Adds one or more storage pools (as  VDC:storage pool tuples) to the specified replication group



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
    id := "id_example" // string | Replication group identifier for which storage pool needs to be added
    dataServiceVpoolServiceAddToVpoolRequest := *openapiclient.NewDataServiceVpoolServiceAddToVpoolRequest([]openapiclient.DataServiceVpoolServiceGetDataServiceVpoolsResponseDataServiceVpoolInnerVarrayMappingsInner{*openapiclient.NewDataServiceVpoolServiceGetDataServiceVpoolsResponseDataServiceVpoolInnerVarrayMappingsInner()}) // DataServiceVpoolServiceAddToVpoolRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataVpoolApi.DataServiceVpoolServiceAddToVpool(context.Background(), id).DataServiceVpoolServiceAddToVpoolRequest(dataServiceVpoolServiceAddToVpoolRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataVpoolApi.DataServiceVpoolServiceAddToVpool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataServiceVpoolServiceAddToVpool`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataVpoolApi.DataServiceVpoolServiceAddToVpool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Replication group identifier for which storage pool needs to be added | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataServiceVpoolServiceAddToVpoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataServiceVpoolServiceAddToVpoolRequest** | [**DataServiceVpoolServiceAddToVpoolRequest**](DataServiceVpoolServiceAddToVpoolRequest.md) |  | 

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


## DataServiceVpoolServiceCreateDataServiceVpool

> DataServiceVpoolServiceCreateDataServiceVpoolResponse DataServiceVpoolServiceCreateDataServiceVpool(ctx).DataServiceVpoolServiceCreateDataServiceVpoolRequest(dataServiceVpoolServiceCreateDataServiceVpoolRequest).Execute()

Creates a replication group that includes the specified storage pools (VDC:storage pool tuple)



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
    dataServiceVpoolServiceCreateDataServiceVpoolRequest := *openapiclient.NewDataServiceVpoolServiceCreateDataServiceVpoolRequest("Name_example") // DataServiceVpoolServiceCreateDataServiceVpoolRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataVpoolApi.DataServiceVpoolServiceCreateDataServiceVpool(context.Background()).DataServiceVpoolServiceCreateDataServiceVpoolRequest(dataServiceVpoolServiceCreateDataServiceVpoolRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataVpoolApi.DataServiceVpoolServiceCreateDataServiceVpool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataServiceVpoolServiceCreateDataServiceVpool`: DataServiceVpoolServiceCreateDataServiceVpoolResponse
    fmt.Fprintf(os.Stdout, "Response from `DataVpoolApi.DataServiceVpoolServiceCreateDataServiceVpool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataServiceVpoolServiceCreateDataServiceVpoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataServiceVpoolServiceCreateDataServiceVpoolRequest** | [**DataServiceVpoolServiceCreateDataServiceVpoolRequest**](DataServiceVpoolServiceCreateDataServiceVpoolRequest.md) |  | 

### Return type

[**DataServiceVpoolServiceCreateDataServiceVpoolResponse**](DataServiceVpoolServiceCreateDataServiceVpoolResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataServiceVpoolServiceGetDataServiceStore

> DataServiceVpoolServiceGetDataServiceStoreResponse DataServiceVpoolServiceGetDataServiceStore(ctx, id).Execute()

Gets the details for the specified replication group



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
    id := "id_example" // string | Replication group identifier for which details needs to be retrieved

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataVpoolApi.DataServiceVpoolServiceGetDataServiceStore(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataVpoolApi.DataServiceVpoolServiceGetDataServiceStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataServiceVpoolServiceGetDataServiceStore`: DataServiceVpoolServiceGetDataServiceStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `DataVpoolApi.DataServiceVpoolServiceGetDataServiceStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Replication group identifier for which details needs to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataServiceVpoolServiceGetDataServiceStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataServiceVpoolServiceGetDataServiceStoreResponse**](DataServiceVpoolServiceGetDataServiceStoreResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataServiceVpoolServiceGetDataServiceVpools

> DataServiceVpoolServiceGetDataServiceVpoolsResponse DataServiceVpoolServiceGetDataServiceVpools(ctx).Execute()

Lists all configured replication groups



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
    resp, r, err := apiClient.DataVpoolApi.DataServiceVpoolServiceGetDataServiceVpools(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataVpoolApi.DataServiceVpoolServiceGetDataServiceVpools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataServiceVpoolServiceGetDataServiceVpools`: DataServiceVpoolServiceGetDataServiceVpoolsResponse
    fmt.Fprintf(os.Stdout, "Response from `DataVpoolApi.DataServiceVpoolServiceGetDataServiceVpools`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDataServiceVpoolServiceGetDataServiceVpoolsRequest struct via the builder pattern


### Return type

[**DataServiceVpoolServiceGetDataServiceVpoolsResponse**](DataServiceVpoolServiceGetDataServiceVpoolsResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataServiceVpoolServicePutDataServiceVpool

> map[string]interface{} DataServiceVpoolServicePutDataServiceVpool(ctx, id).DataServiceVpoolServicePutDataServiceVpoolRequest(dataServiceVpoolServicePutDataServiceVpoolRequest).Execute()

Updates the name and description for a replication group



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
    id := "id_example" // string | Replication group identifier for which details needs to be updated
    dataServiceVpoolServicePutDataServiceVpoolRequest := *openapiclient.NewDataServiceVpoolServicePutDataServiceVpoolRequest() // DataServiceVpoolServicePutDataServiceVpoolRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataVpoolApi.DataServiceVpoolServicePutDataServiceVpool(context.Background(), id).DataServiceVpoolServicePutDataServiceVpoolRequest(dataServiceVpoolServicePutDataServiceVpoolRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataVpoolApi.DataServiceVpoolServicePutDataServiceVpool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataServiceVpoolServicePutDataServiceVpool`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataVpoolApi.DataServiceVpoolServicePutDataServiceVpool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Replication group identifier for which details needs to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataServiceVpoolServicePutDataServiceVpoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataServiceVpoolServicePutDataServiceVpoolRequest** | [**DataServiceVpoolServicePutDataServiceVpoolRequest**](DataServiceVpoolServicePutDataServiceVpoolRequest.md) |  | 

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


## DataServiceVpoolServiceRemoveFromVpool

> map[string]interface{} DataServiceVpoolServiceRemoveFromVpool(ctx, id).DataServiceVpoolServiceRemoveFromVpoolRequest(dataServiceVpoolServiceRemoveFromVpoolRequest).SkipBootstrapCheck(skipBootstrapCheck).ForcePSOzones(forcePSOzones).Execute()

Deletes a storage pool (VDC:storage pool tuple) from a specified replication group



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
    id := "id_example" // string | Replication group identifier for which storage pool needs to be removed
    dataServiceVpoolServiceRemoveFromVpoolRequest := *openapiclient.NewDataServiceVpoolServiceRemoveFromVpoolRequest([]openapiclient.DataServiceVpoolServiceGetDataServiceVpoolsResponseDataServiceVpoolInnerVarrayMappingsInner{*openapiclient.NewDataServiceVpoolServiceGetDataServiceVpoolsResponseDataServiceVpoolInnerVarrayMappingsInner()}) // DataServiceVpoolServiceRemoveFromVpoolRequest | 
    skipBootstrapCheck := "skipBootstrapCheck_example" // string |  (optional)
    forcePSOzones := "forcePSOzones_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataVpoolApi.DataServiceVpoolServiceRemoveFromVpool(context.Background(), id).DataServiceVpoolServiceRemoveFromVpoolRequest(dataServiceVpoolServiceRemoveFromVpoolRequest).SkipBootstrapCheck(skipBootstrapCheck).ForcePSOzones(forcePSOzones).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataVpoolApi.DataServiceVpoolServiceRemoveFromVpool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataServiceVpoolServiceRemoveFromVpool`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataVpoolApi.DataServiceVpoolServiceRemoveFromVpool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Replication group identifier for which storage pool needs to be removed | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataServiceVpoolServiceRemoveFromVpoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataServiceVpoolServiceRemoveFromVpoolRequest** | [**DataServiceVpoolServiceRemoveFromVpoolRequest**](DataServiceVpoolServiceRemoveFromVpoolRequest.md) |  | 
 **skipBootstrapCheck** | **string** |  | 
 **forcePSOzones** | **string** |  | 

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

