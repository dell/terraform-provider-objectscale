# \UserSecretKeyApi

All URIs are relative to *https://objectscale.local:4443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserSecretKeyServiceCreateNewKeyForUser**](UserSecretKeyApi.md#UserSecretKeyServiceCreateNewKeyForUser) | **Post** /object/user-secret-keys/{uid} | Creates a secret key with the given details for the specified user
[**UserSecretKeyServiceDeleteKeyForUser**](UserSecretKeyApi.md#UserSecretKeyServiceDeleteKeyForUser) | **Post** /object/user-secret-keys/{uid}/deactivate | Deletes a specified secret key for a user
[**UserSecretKeyServiceGetKeysExistForUser**](UserSecretKeyApi.md#UserSecretKeyServiceGetKeysExistForUser) | **Get** /object/user-secret-keys/exist/{uid}/{namespace} | Returns indication if secret keys for the specified user and namespace exist
[**UserSecretKeyServiceGetKeysForUser**](UserSecretKeyApi.md#UserSecretKeyServiceGetKeysForUser) | **Get** /object/user-secret-keys/{uid} | Gets all secret keys for the specified user
[**UserSecretKeyServiceGetKeysForUser1**](UserSecretKeyApi.md#UserSecretKeyServiceGetKeysForUser1) | **Get** /object/user-secret-keys/{uid}/{namespace} | Gets all secret keys for the specified user and namespace



## UserSecretKeyServiceCreateNewKeyForUser

> UserSecretKeyServiceCreateNewKeyForUserResponse UserSecretKeyServiceCreateNewKeyForUser(ctx, uid).UserSecretKeyServiceCreateNewKeyForUserRequest(userSecretKeyServiceCreateNewKeyForUserRequest).Execute()

Creates a secret key with the given details for the specified user



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
    uid := "uid_example" // string | Valid user identifier to create a key for
    userSecretKeyServiceCreateNewKeyForUserRequest := *openapiclient.NewUserSecretKeyServiceCreateNewKeyForUserRequest() // UserSecretKeyServiceCreateNewKeyForUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserSecretKeyApi.UserSecretKeyServiceCreateNewKeyForUser(context.Background(), uid).UserSecretKeyServiceCreateNewKeyForUserRequest(userSecretKeyServiceCreateNewKeyForUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSecretKeyApi.UserSecretKeyServiceCreateNewKeyForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserSecretKeyServiceCreateNewKeyForUser`: UserSecretKeyServiceCreateNewKeyForUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UserSecretKeyApi.UserSecretKeyServiceCreateNewKeyForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Valid user identifier to create a key for | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserSecretKeyServiceCreateNewKeyForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userSecretKeyServiceCreateNewKeyForUserRequest** | [**UserSecretKeyServiceCreateNewKeyForUserRequest**](UserSecretKeyServiceCreateNewKeyForUserRequest.md) |  | 

### Return type

[**UserSecretKeyServiceCreateNewKeyForUserResponse**](UserSecretKeyServiceCreateNewKeyForUserResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSecretKeyServiceDeleteKeyForUser

> map[string]interface{} UserSecretKeyServiceDeleteKeyForUser(ctx, uid).UserSecretKeyServiceDeleteKeyForUserRequest(userSecretKeyServiceDeleteKeyForUserRequest).Execute()

Deletes a specified secret key for a user



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
    uid := "uid_example" // string | Valid user identifier to delete the key from
    userSecretKeyServiceDeleteKeyForUserRequest := *openapiclient.NewUserSecretKeyServiceDeleteKeyForUserRequest() // UserSecretKeyServiceDeleteKeyForUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserSecretKeyApi.UserSecretKeyServiceDeleteKeyForUser(context.Background(), uid).UserSecretKeyServiceDeleteKeyForUserRequest(userSecretKeyServiceDeleteKeyForUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSecretKeyApi.UserSecretKeyServiceDeleteKeyForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserSecretKeyServiceDeleteKeyForUser`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UserSecretKeyApi.UserSecretKeyServiceDeleteKeyForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Valid user identifier to delete the key from | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserSecretKeyServiceDeleteKeyForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userSecretKeyServiceDeleteKeyForUserRequest** | [**UserSecretKeyServiceDeleteKeyForUserRequest**](UserSecretKeyServiceDeleteKeyForUserRequest.md) |  | 

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


## UserSecretKeyServiceGetKeysExistForUser

> UserSecretKeyServiceGetKeysExistForUserResponse UserSecretKeyServiceGetKeysExistForUser(ctx, uid, namespace).Execute()

Returns indication if secret keys for the specified user and namespace exist



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
    uid := "uid_example" // string | Valid user identifier to get the keys from
    namespace := "namespace_example" // string | the namespace to get all secret keys

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserSecretKeyApi.UserSecretKeyServiceGetKeysExistForUser(context.Background(), uid, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSecretKeyApi.UserSecretKeyServiceGetKeysExistForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserSecretKeyServiceGetKeysExistForUser`: UserSecretKeyServiceGetKeysExistForUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UserSecretKeyApi.UserSecretKeyServiceGetKeysExistForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Valid user identifier to get the keys from | 
**namespace** | **string** | the namespace to get all secret keys | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserSecretKeyServiceGetKeysExistForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserSecretKeyServiceGetKeysExistForUserResponse**](UserSecretKeyServiceGetKeysExistForUserResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSecretKeyServiceGetKeysForUser

> UserSecretKeyServiceGetKeysForUserResponse UserSecretKeyServiceGetKeysForUser(ctx, uid).Execute()

Gets all secret keys for the specified user



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
    uid := "uid_example" // string | Valid user identifier to get the keys from

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserSecretKeyApi.UserSecretKeyServiceGetKeysForUser(context.Background(), uid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSecretKeyApi.UserSecretKeyServiceGetKeysForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserSecretKeyServiceGetKeysForUser`: UserSecretKeyServiceGetKeysForUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UserSecretKeyApi.UserSecretKeyServiceGetKeysForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Valid user identifier to get the keys from | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserSecretKeyServiceGetKeysForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserSecretKeyServiceGetKeysForUserResponse**](UserSecretKeyServiceGetKeysForUserResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSecretKeyServiceGetKeysForUser1

> UserSecretKeyServiceGetKeysForUser1Response UserSecretKeyServiceGetKeysForUser1(ctx, uid, namespace).Execute()

Gets all secret keys for the specified user and namespace



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
    uid := "uid_example" // string | Valid user identifier to get the keys from
    namespace := "namespace_example" // string | the namespace to get all secret keys

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserSecretKeyApi.UserSecretKeyServiceGetKeysForUser1(context.Background(), uid, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSecretKeyApi.UserSecretKeyServiceGetKeysForUser1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserSecretKeyServiceGetKeysForUser1`: UserSecretKeyServiceGetKeysForUser1Response
    fmt.Fprintf(os.Stdout, "Response from `UserSecretKeyApi.UserSecretKeyServiceGetKeysForUser1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Valid user identifier to get the keys from | 
**namespace** | **string** | the namespace to get all secret keys | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserSecretKeyServiceGetKeysForUser1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserSecretKeyServiceGetKeysForUser1Response**](UserSecretKeyServiceGetKeysForUser1Response.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

