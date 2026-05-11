# \IamProviderApi

All URIs are relative to *https://objectscale.local:4443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceProviderCreate**](IamProviderApi.md#ServiceProviderCreate) | **Post** /ecs-service-provider | Creates a service provider using the specified attributes
[**ServiceProviderDelete**](IamProviderApi.md#ServiceProviderDelete) | **Delete** /ecs-service-provider | Deletes a service provider
[**ServiceProviderGet**](IamProviderApi.md#ServiceProviderGet) | **Get** /ecs-service-provider | Returns a service provider if it exists
[**ServiceProviderGetMetadata**](IamProviderApi.md#ServiceProviderGetMetadata) | **Get** /ecs-service-provider/metadata | Returns metadata for a service provider
[**ServiceProviderUpdate**](IamProviderApi.md#ServiceProviderUpdate) | **Put** /ecs-service-provider | Creates a service provider using the specified attributes



## ServiceProviderCreate

> ServiceProviderCreateResponse ServiceProviderCreate(ctx).IamServiceProviderControllerProcessCreateServiceProviderRequest(iamServiceProviderControllerProcessCreateServiceProviderRequest).Execute()

Creates a service provider using the specified attributes



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
    iamServiceProviderControllerProcessCreateServiceProviderRequest := *openapiclient.NewIamServiceProviderControllerProcessCreateServiceProviderRequest() // IamServiceProviderControllerProcessCreateServiceProviderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProviderApi.ServiceProviderCreate(context.Background()).IamServiceProviderControllerProcessCreateServiceProviderRequest(iamServiceProviderControllerProcessCreateServiceProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProviderApi.ServiceProviderCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceProviderCreate`: ServiceProviderCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `IamProviderApi.ServiceProviderCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceProviderCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamServiceProviderControllerProcessCreateServiceProviderRequest** | [**IamServiceProviderControllerProcessCreateServiceProviderRequest**](IamServiceProviderControllerProcessCreateServiceProviderRequest.md) |  | 

### Return type

[**ServiceProviderCreateResponse**](ServiceProviderCreateResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceProviderDelete

> ServiceProviderDeleteResponse ServiceProviderDelete(ctx).Execute()

Deletes a service provider



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
    resp, r, err := apiClient.IamProviderApi.ServiceProviderDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProviderApi.ServiceProviderDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceProviderDelete`: ServiceProviderDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `IamProviderApi.ServiceProviderDelete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServiceProviderDeleteRequest struct via the builder pattern


### Return type

[**ServiceProviderDeleteResponse**](ServiceProviderDeleteResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceProviderGet

> ServiceProviderGetResponse ServiceProviderGet(ctx).Execute()

Returns a service provider if it exists



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
    resp, r, err := apiClient.IamProviderApi.ServiceProviderGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProviderApi.ServiceProviderGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceProviderGet`: ServiceProviderGetResponse
    fmt.Fprintf(os.Stdout, "Response from `IamProviderApi.ServiceProviderGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServiceProviderGetRequest struct via the builder pattern


### Return type

[**ServiceProviderGetResponse**](ServiceProviderGetResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceProviderGetMetadata

> string ServiceProviderGetMetadata(ctx).Execute()

Returns metadata for a service provider



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
    resp, r, err := apiClient.IamProviderApi.ServiceProviderGetMetadata(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProviderApi.ServiceProviderGetMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceProviderGetMetadata`: string
    fmt.Fprintf(os.Stdout, "Response from `IamProviderApi.ServiceProviderGetMetadata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServiceProviderGetMetadataRequest struct via the builder pattern


### Return type

**string**

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceProviderUpdate

> ServiceProviderUpdateResponse ServiceProviderUpdate(ctx).IamServiceProviderControllerProcessUpdateServiceProviderRequest(iamServiceProviderControllerProcessUpdateServiceProviderRequest).Execute()

Creates a service provider using the specified attributes



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
    iamServiceProviderControllerProcessUpdateServiceProviderRequest := *openapiclient.NewIamServiceProviderControllerProcessUpdateServiceProviderRequest() // IamServiceProviderControllerProcessUpdateServiceProviderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IamProviderApi.ServiceProviderUpdate(context.Background()).IamServiceProviderControllerProcessUpdateServiceProviderRequest(iamServiceProviderControllerProcessUpdateServiceProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IamProviderApi.ServiceProviderUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceProviderUpdate`: ServiceProviderUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `IamProviderApi.ServiceProviderUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceProviderUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamServiceProviderControllerProcessUpdateServiceProviderRequest** | [**IamServiceProviderControllerProcessUpdateServiceProviderRequest**](IamServiceProviderControllerProcessUpdateServiceProviderRequest.md) |  | 

### Return type

[**ServiceProviderUpdateResponse**](ServiceProviderUpdateResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

