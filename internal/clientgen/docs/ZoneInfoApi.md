# \ZoneInfoApi

All URIs are relative to *https://objectscale.local:4443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ZoneInfoServiceGetLocalVdc**](ZoneInfoApi.md#ZoneInfoServiceGetLocalVdc) | **Get** /object/vdcs/vdc/local | Gets the details for the local VDC
[**ZoneInfoServiceGetVdcById**](ZoneInfoApi.md#ZoneInfoServiceGetVdcById) | **Get** /object/vdcs/vdcid/{vdcId} | Gets the details for a VDC specified by VDC Id
[**ZoneInfoServiceGetVdcByName**](ZoneInfoApi.md#ZoneInfoServiceGetVdcByName) | **Get** /object/vdcs/vdc/{vdcName} | Gets the details for a VDC specified by name
[**ZoneInfoServiceInsertVdcInfo**](ZoneInfoApi.md#ZoneInfoServiceInsertVdcInfo) | **Put** /object/vdcs/vdc/{vdcName} | Inserts attributes for the current VDC or a VDC to connect to
[**ZoneInfoServiceListAllVdc**](ZoneInfoApi.md#ZoneInfoServiceListAllVdc) | **Get** /object/vdcs/vdc/list | Gets the details of all configured VDCs



## ZoneInfoServiceGetLocalVdc

> Vdc ZoneInfoServiceGetLocalVdc(ctx).Execute()

Gets the details for the local VDC



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
    resp, r, err := apiClient.ZoneInfoApi.ZoneInfoServiceGetLocalVdc(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZoneInfoApi.ZoneInfoServiceGetLocalVdc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZoneInfoServiceGetLocalVdc`: Vdc
    fmt.Fprintf(os.Stdout, "Response from `ZoneInfoApi.ZoneInfoServiceGetLocalVdc`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiZoneInfoServiceGetLocalVdcRequest struct via the builder pattern


### Return type

[**Vdc**](Vdc.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneInfoServiceGetVdcById

> Vdc ZoneInfoServiceGetVdcById(ctx, vdcId).Execute()

Gets the details for a VDC specified by VDC Id



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
    vdcId := "vdcId_example" // string | VDC identifier for which VDC Information is to be retrieved

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZoneInfoApi.ZoneInfoServiceGetVdcById(context.Background(), vdcId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZoneInfoApi.ZoneInfoServiceGetVdcById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZoneInfoServiceGetVdcById`: Vdc
    fmt.Fprintf(os.Stdout, "Response from `ZoneInfoApi.ZoneInfoServiceGetVdcById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdcId** | **string** | VDC identifier for which VDC Information is to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiZoneInfoServiceGetVdcByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Vdc**](Vdc.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneInfoServiceGetVdcByName

> Vdc ZoneInfoServiceGetVdcByName(ctx, vdcName).Execute()

Gets the details for a VDC specified by name



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
    vdcName := "vdcName_example" // string | VDC name for which VDC Information is to be retrieved

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZoneInfoApi.ZoneInfoServiceGetVdcByName(context.Background(), vdcName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZoneInfoApi.ZoneInfoServiceGetVdcByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZoneInfoServiceGetVdcByName`: Vdc
    fmt.Fprintf(os.Stdout, "Response from `ZoneInfoApi.ZoneInfoServiceGetVdcByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdcName** | **string** | VDC name for which VDC Information is to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiZoneInfoServiceGetVdcByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Vdc**](Vdc.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneInfoServiceInsertVdcInfo

> map[string]interface{} ZoneInfoServiceInsertVdcInfo(ctx, vdcName).ZoneInfoServiceInsertVdcInfoRequest(zoneInfoServiceInsertVdcInfoRequest).SkipMemoryProfileChecks(skipMemoryProfileChecks).Execute()

Inserts attributes for the current VDC or a VDC to connect to



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
    vdcName := "vdcName_example" // string | VDC name for which mapping needs to be inserted
    zoneInfoServiceInsertVdcInfoRequest := *openapiclient.NewZoneInfoServiceInsertVdcInfoRequest() // ZoneInfoServiceInsertVdcInfoRequest | 
    skipMemoryProfileChecks := "skipMemoryProfileChecks_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZoneInfoApi.ZoneInfoServiceInsertVdcInfo(context.Background(), vdcName).ZoneInfoServiceInsertVdcInfoRequest(zoneInfoServiceInsertVdcInfoRequest).SkipMemoryProfileChecks(skipMemoryProfileChecks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZoneInfoApi.ZoneInfoServiceInsertVdcInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZoneInfoServiceInsertVdcInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ZoneInfoApi.ZoneInfoServiceInsertVdcInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdcName** | **string** | VDC name for which mapping needs to be inserted | 

### Other Parameters

Other parameters are passed through a pointer to a apiZoneInfoServiceInsertVdcInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zoneInfoServiceInsertVdcInfoRequest** | [**ZoneInfoServiceInsertVdcInfoRequest**](ZoneInfoServiceInsertVdcInfoRequest.md) |  | 
 **skipMemoryProfileChecks** | **string** |  | 

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


## ZoneInfoServiceListAllVdc

> ZoneInfoServiceListAllVdcResponse ZoneInfoServiceListAllVdc(ctx).Execute()

Gets the details of all configured VDCs



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
    resp, r, err := apiClient.ZoneInfoApi.ZoneInfoServiceListAllVdc(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZoneInfoApi.ZoneInfoServiceListAllVdc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZoneInfoServiceListAllVdc`: ZoneInfoServiceListAllVdcResponse
    fmt.Fprintf(os.Stdout, "Response from `ZoneInfoApi.ZoneInfoServiceListAllVdc`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiZoneInfoServiceListAllVdcRequest struct via the builder pattern


### Return type

[**ZoneInfoServiceListAllVdcResponse**](ZoneInfoServiceListAllVdcResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

