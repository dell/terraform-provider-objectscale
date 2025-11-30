# \NamespaceApi

All URIs are relative to *https://objectscale.local:4443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NamespaceServiceCreateNamespace**](NamespaceApi.md#NamespaceServiceCreateNamespace) | **Post** /object/namespaces/namespace | Creates a namespace with the given details
[**NamespaceServiceCreateRetentionClass**](NamespaceApi.md#NamespaceServiceCreateRetentionClass) | **Post** /object/namespaces/namespace/{namespace}/retention | Creates a retention class for the specified namespace
[**NamespaceServiceDeactivateNamespace**](NamespaceApi.md#NamespaceServiceDeactivateNamespace) | **Post** /object/namespaces/namespace/{namespace}/deactivate | Deactivates and deletes the given namespace and all associated user mappings
[**NamespaceServiceGetNamespace**](NamespaceApi.md#NamespaceServiceGetNamespace) | **Get** /object/namespaces/namespace/{id} | Gets the details for the specified namespace
[**NamespaceServiceGetNamespaceQuota**](NamespaceApi.md#NamespaceServiceGetNamespaceQuota) | **Get** /object/namespaces/namespace/{namespace}/quota | Gets the namespace quota for a specified namespace
[**NamespaceServiceGetNamespaces**](NamespaceApi.md#NamespaceServiceGetNamespaces) | **Get** /object/namespaces | Gets the list of all configured namespaces
[**NamespaceServiceGetRetentionClass**](NamespaceApi.md#NamespaceServiceGetRetentionClass) | **Get** /object/namespaces/namespace/{namespace}/retention/{class} | Gets the retention period for the given namespace and retention class
[**NamespaceServiceGetRetentionClasses**](NamespaceApi.md#NamespaceServiceGetRetentionClasses) | **Get** /object/namespaces/namespace/{namespace}/retention | Gets the list of retention classes for the specified namespace
[**NamespaceServiceRemoveNamespaceQuota**](NamespaceApi.md#NamespaceServiceRemoveNamespaceQuota) | **Delete** /object/namespaces/namespace/{namespace}/quota | Deletes the namespace quota for the specified namespace
[**NamespaceServiceUpdateNamespace**](NamespaceApi.md#NamespaceServiceUpdateNamespace) | **Put** /object/namespaces/namespace/{namespace} | Updates namespace details like replication group list, namespace admins and user mappings
[**NamespaceServiceUpdateNamespaceQuota**](NamespaceApi.md#NamespaceServiceUpdateNamespaceQuota) | **Put** /object/namespaces/namespace/{namespace}/quota | Updates the namespace quota for a specified namespace
[**NamespaceServiceUpdateRetentionClass**](NamespaceApi.md#NamespaceServiceUpdateRetentionClass) | **Put** /object/namespaces/namespace/{namespace}/retention/{class} | Updates the retention class details for a specified retention class for a namespace



## NamespaceServiceCreateNamespace

> NamespaceServiceCreateNamespaceResponse NamespaceServiceCreateNamespace(ctx).NamespaceServiceCreateNamespaceRequest(namespaceServiceCreateNamespaceRequest).Execute()

Creates a namespace with the given details



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
    namespaceServiceCreateNamespaceRequest := *openapiclient.NewNamespaceServiceCreateNamespaceRequest("Namespace_example") // NamespaceServiceCreateNamespaceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespaceApi.NamespaceServiceCreateNamespace(context.Background()).NamespaceServiceCreateNamespaceRequest(namespaceServiceCreateNamespaceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceApi.NamespaceServiceCreateNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NamespaceServiceCreateNamespace`: NamespaceServiceCreateNamespaceResponse
    fmt.Fprintf(os.Stdout, "Response from `NamespaceApi.NamespaceServiceCreateNamespace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNamespaceServiceCreateNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespaceServiceCreateNamespaceRequest** | [**NamespaceServiceCreateNamespaceRequest**](NamespaceServiceCreateNamespaceRequest.md) |  | 

### Return type

[**NamespaceServiceCreateNamespaceResponse**](NamespaceServiceCreateNamespaceResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NamespaceServiceCreateRetentionClass

> map[string]interface{} NamespaceServiceCreateRetentionClass(ctx, namespace).NamespaceServiceCreateRetentionClassRequest(namespaceServiceCreateRetentionClassRequest).Execute()

Creates a retention class for the specified namespace



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
    namespace := "namespace_example" // string | Namespace identifier for which retention class needs to created.
    namespaceServiceCreateRetentionClassRequest := *openapiclient.NewNamespaceServiceCreateRetentionClassRequest() // NamespaceServiceCreateRetentionClassRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespaceApi.NamespaceServiceCreateRetentionClass(context.Background(), namespace).NamespaceServiceCreateRetentionClassRequest(namespaceServiceCreateRetentionClassRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceApi.NamespaceServiceCreateRetentionClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NamespaceServiceCreateRetentionClass`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NamespaceApi.NamespaceServiceCreateRetentionClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Namespace identifier for which retention class needs to created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNamespaceServiceCreateRetentionClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespaceServiceCreateRetentionClassRequest** | [**NamespaceServiceCreateRetentionClassRequest**](NamespaceServiceCreateRetentionClassRequest.md) |  | 

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


## NamespaceServiceDeactivateNamespace

> map[string]interface{} NamespaceServiceDeactivateNamespace(ctx, namespace).Execute()

Deactivates and deletes the given namespace and all associated user mappings



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
    namespace := "namespace_example" // string | An active namespace identifier which needs to be deactivated/deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespaceApi.NamespaceServiceDeactivateNamespace(context.Background(), namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceApi.NamespaceServiceDeactivateNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NamespaceServiceDeactivateNamespace`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NamespaceApi.NamespaceServiceDeactivateNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | An active namespace identifier which needs to be deactivated/deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiNamespaceServiceDeactivateNamespaceRequest struct via the builder pattern


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


## NamespaceServiceGetNamespace

> NamespaceServiceGetNamespaceResponse NamespaceServiceGetNamespace(ctx, id).Execute()

Gets the details for the specified namespace



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
    id := "id_example" // string | Namespace identifier for which details needs to be retrieved.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespaceApi.NamespaceServiceGetNamespace(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceApi.NamespaceServiceGetNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NamespaceServiceGetNamespace`: NamespaceServiceGetNamespaceResponse
    fmt.Fprintf(os.Stdout, "Response from `NamespaceApi.NamespaceServiceGetNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Namespace identifier for which details needs to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNamespaceServiceGetNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NamespaceServiceGetNamespaceResponse**](NamespaceServiceGetNamespaceResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NamespaceServiceGetNamespaceQuota

> NamespaceServiceGetNamespaceQuotaResponse NamespaceServiceGetNamespaceQuota(ctx, namespace).Execute()

Gets the namespace quota for a specified namespace



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
    namespace := "namespace_example" // string | Namespace identifier for which namespace quota details needs to retrieved.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespaceApi.NamespaceServiceGetNamespaceQuota(context.Background(), namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceApi.NamespaceServiceGetNamespaceQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NamespaceServiceGetNamespaceQuota`: NamespaceServiceGetNamespaceQuotaResponse
    fmt.Fprintf(os.Stdout, "Response from `NamespaceApi.NamespaceServiceGetNamespaceQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Namespace identifier for which namespace quota details needs to retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNamespaceServiceGetNamespaceQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NamespaceServiceGetNamespaceQuotaResponse**](NamespaceServiceGetNamespaceQuotaResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NamespaceServiceGetNamespaces

> NamespaceServiceGetNamespacesResponse NamespaceServiceGetNamespaces(ctx).Limit(limit).Marker(marker).Name(name).Execute()

Gets the list of all configured namespaces



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
    limit := "limit_example" // string | Number of objects requested in current fetch. (optional)
    marker := "marker_example" // string | Reference to last object returned. (optional)
    name := "name_example" // string | Case sensitive prefix of the Namespace name with a wild card(*) Ex : any_prefix_string* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespaceApi.NamespaceServiceGetNamespaces(context.Background()).Limit(limit).Marker(marker).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceApi.NamespaceServiceGetNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NamespaceServiceGetNamespaces`: NamespaceServiceGetNamespacesResponse
    fmt.Fprintf(os.Stdout, "Response from `NamespaceApi.NamespaceServiceGetNamespaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNamespaceServiceGetNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | Number of objects requested in current fetch. | 
 **marker** | **string** | Reference to last object returned. | 
 **name** | **string** | Case sensitive prefix of the Namespace name with a wild card(*) Ex : any_prefix_string* | 

### Return type

[**NamespaceServiceGetNamespacesResponse**](NamespaceServiceGetNamespacesResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NamespaceServiceGetRetentionClass

> NamespaceServiceGetRetentionClassResponse NamespaceServiceGetRetentionClass(ctx, namespace, class).Execute()

Gets the retention period for the given namespace and retention class



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
    namespace := "namespace_example" // string | Namespace for which retention period needs to retrieved
    class := "class_example" // string | Class name for which retention period needs to retrieved

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespaceApi.NamespaceServiceGetRetentionClass(context.Background(), namespace, class).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceApi.NamespaceServiceGetRetentionClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NamespaceServiceGetRetentionClass`: NamespaceServiceGetRetentionClassResponse
    fmt.Fprintf(os.Stdout, "Response from `NamespaceApi.NamespaceServiceGetRetentionClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Namespace for which retention period needs to retrieved | 
**class** | **string** | Class name for which retention period needs to retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiNamespaceServiceGetRetentionClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NamespaceServiceGetRetentionClassResponse**](NamespaceServiceGetRetentionClassResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NamespaceServiceGetRetentionClasses

> NamespaceServiceGetRetentionClassesResponse NamespaceServiceGetRetentionClasses(ctx, namespace).Execute()

Gets the list of retention classes for the specified namespace



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
    namespace := "namespace_example" // string | Namespace identifier for which retention classes needs to retrieved

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespaceApi.NamespaceServiceGetRetentionClasses(context.Background(), namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceApi.NamespaceServiceGetRetentionClasses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NamespaceServiceGetRetentionClasses`: NamespaceServiceGetRetentionClassesResponse
    fmt.Fprintf(os.Stdout, "Response from `NamespaceApi.NamespaceServiceGetRetentionClasses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Namespace identifier for which retention classes needs to retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiNamespaceServiceGetRetentionClassesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NamespaceServiceGetRetentionClassesResponse**](NamespaceServiceGetRetentionClassesResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NamespaceServiceRemoveNamespaceQuota

> map[string]interface{} NamespaceServiceRemoveNamespaceQuota(ctx, namespace).Execute()

Deletes the namespace quota for the specified namespace



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
    namespace := "namespace_example" // string | Namespace identifier for which namespace quota details needs to deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespaceApi.NamespaceServiceRemoveNamespaceQuota(context.Background(), namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceApi.NamespaceServiceRemoveNamespaceQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NamespaceServiceRemoveNamespaceQuota`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NamespaceApi.NamespaceServiceRemoveNamespaceQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Namespace identifier for which namespace quota details needs to deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNamespaceServiceRemoveNamespaceQuotaRequest struct via the builder pattern


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


## NamespaceServiceUpdateNamespace

> map[string]interface{} NamespaceServiceUpdateNamespace(ctx, namespace).NamespaceServiceUpdateNamespaceRequest(namespaceServiceUpdateNamespaceRequest).Execute()

Updates namespace details like replication group list, namespace admins and user mappings



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
    namespace := "namespace_example" // string | Namespace identifier whose details needs to be updated
    namespaceServiceUpdateNamespaceRequest := *openapiclient.NewNamespaceServiceUpdateNamespaceRequest([]openapiclient.NamespaceServiceGetNamespacesResponseNamespaceInnerUserMappingInner{*openapiclient.NewNamespaceServiceGetNamespacesResponseNamespaceInnerUserMappingInner("Domain_example")}) // NamespaceServiceUpdateNamespaceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespaceApi.NamespaceServiceUpdateNamespace(context.Background(), namespace).NamespaceServiceUpdateNamespaceRequest(namespaceServiceUpdateNamespaceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceApi.NamespaceServiceUpdateNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NamespaceServiceUpdateNamespace`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NamespaceApi.NamespaceServiceUpdateNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Namespace identifier whose details needs to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiNamespaceServiceUpdateNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespaceServiceUpdateNamespaceRequest** | [**NamespaceServiceUpdateNamespaceRequest**](NamespaceServiceUpdateNamespaceRequest.md) |  | 

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


## NamespaceServiceUpdateNamespaceQuota

> map[string]interface{} NamespaceServiceUpdateNamespaceQuota(ctx, namespace).NamespaceServiceUpdateNamespaceQuotaRequest(namespaceServiceUpdateNamespaceQuotaRequest).Execute()

Updates the namespace quota for a specified namespace



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
    namespace := "namespace_example" // string | Namespace identifier for which namespace quota details need to be updated.
    namespaceServiceUpdateNamespaceQuotaRequest := *openapiclient.NewNamespaceServiceUpdateNamespaceQuotaRequest() // NamespaceServiceUpdateNamespaceQuotaRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespaceApi.NamespaceServiceUpdateNamespaceQuota(context.Background(), namespace).NamespaceServiceUpdateNamespaceQuotaRequest(namespaceServiceUpdateNamespaceQuotaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceApi.NamespaceServiceUpdateNamespaceQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NamespaceServiceUpdateNamespaceQuota`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NamespaceApi.NamespaceServiceUpdateNamespaceQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Namespace identifier for which namespace quota details need to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNamespaceServiceUpdateNamespaceQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespaceServiceUpdateNamespaceQuotaRequest** | [**NamespaceServiceUpdateNamespaceQuotaRequest**](NamespaceServiceUpdateNamespaceQuotaRequest.md) |  | 

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


## NamespaceServiceUpdateRetentionClass

> map[string]interface{} NamespaceServiceUpdateRetentionClass(ctx, namespace, class).NamespaceServiceUpdateRetentionClassRequest(namespaceServiceUpdateRetentionClassRequest).Execute()

Updates the retention class details for a specified retention class for a namespace



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
    namespace := "namespace_example" // string | Namespace identifier for which retention class needs to retrieved.
    class := "class_example" // string | Retention class for which details needs to updated.
    namespaceServiceUpdateRetentionClassRequest := *openapiclient.NewNamespaceServiceUpdateRetentionClassRequest() // NamespaceServiceUpdateRetentionClassRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespaceApi.NamespaceServiceUpdateRetentionClass(context.Background(), namespace, class).NamespaceServiceUpdateRetentionClassRequest(namespaceServiceUpdateRetentionClassRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceApi.NamespaceServiceUpdateRetentionClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NamespaceServiceUpdateRetentionClass`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NamespaceApi.NamespaceServiceUpdateRetentionClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Namespace identifier for which retention class needs to retrieved. | 
**class** | **string** | Retention class for which details needs to updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNamespaceServiceUpdateRetentionClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **namespaceServiceUpdateRetentionClassRequest** | [**NamespaceServiceUpdateRetentionClassRequest**](NamespaceServiceUpdateRetentionClassRequest.md) |  | 

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

