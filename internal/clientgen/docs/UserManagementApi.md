# \UserManagementApi

All URIs are relative to *https://objectscale.local:4443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserManagementServiceAddUser**](UserManagementApi.md#UserManagementServiceAddUser) | **Post** /object/users | Creates a user for the specified namespace
[**UserManagementServiceAddUserTag**](UserManagementApi.md#UserManagementServiceAddUserTag) | **Post** /object/users/{uid}/tags | Updates user tags for the specified user - this is append operation
[**UserManagementServiceGetAllUsers**](UserManagementApi.md#UserManagementServiceGetAllUsers) | **Get** /object/users | Gets identifiers for all configured users
[**UserManagementServiceGetUserInfo**](UserManagementApi.md#UserManagementServiceGetUserInfo) | **Get** /object/users/{uid}/info | Gets user details for the specified user belonging to specified namespace
[**UserManagementServiceGetUserLockWithNamespace**](UserManagementApi.md#UserManagementServiceGetUserLockWithNamespace) | **Get** /object/users/lock/{uid}/{namespace} | Gets the user lock details for the specified user belonging to specified namespace
[**UserManagementServiceGetUserLockWithoutNamespace**](UserManagementApi.md#UserManagementServiceGetUserLockWithoutNamespace) | **Get** /object/users/lock/{uid} | Gets the user lock details for the specified user
[**UserManagementServiceGetUserTagsWithNamespace**](UserManagementApi.md#UserManagementServiceGetUserTagsWithNamespace) | **Get** /object/users/{uid}/tags | Gets the user tags details for the specified user belonging to specified namespace
[**UserManagementServiceGetUsersForNamespace**](UserManagementApi.md#UserManagementServiceGetUsersForNamespace) | **Get** /object/users/{namespace} | Gets all user identifiers for the specified namespace
[**UserManagementServiceQueryUsers**](UserManagementApi.md#UserManagementServiceQueryUsers) | **Get** /object/users/query | Gets user details for the specified user belonging to specified namespace
[**UserManagementServiceRemoveUser**](UserManagementApi.md#UserManagementServiceRemoveUser) | **Post** /object/users/deactivate | Deletes the specified user and its associated secret keys
[**UserManagementServiceRemoveUserTags**](UserManagementApi.md#UserManagementServiceRemoveUserTags) | **Delete** /object/users/{uid}/tags | Deletes user tags for specified user
[**UserManagementServiceSetUserLock**](UserManagementApi.md#UserManagementServiceSetUserLock) | **Put** /object/users/lock | Locks the specified user
[**UserManagementServiceUpdateUserTag**](UserManagementApi.md#UserManagementServiceUpdateUserTag) | **Put** /object/users/{uid}/tags | Updates user tags for the specified user



## UserManagementServiceAddUser

> UserManagementServiceAddUserResponse UserManagementServiceAddUser(ctx).UserManagementServiceAddUserRequest(userManagementServiceAddUserRequest).Execute()

Creates a user for the specified namespace



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
    userManagementServiceAddUserRequest := *openapiclient.NewUserManagementServiceAddUserRequest("User_example", "Namespace_example") // UserManagementServiceAddUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserManagementApi.UserManagementServiceAddUser(context.Background()).UserManagementServiceAddUserRequest(userManagementServiceAddUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.UserManagementServiceAddUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserManagementServiceAddUser`: UserManagementServiceAddUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.UserManagementServiceAddUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserManagementServiceAddUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userManagementServiceAddUserRequest** | [**UserManagementServiceAddUserRequest**](UserManagementServiceAddUserRequest.md) |  | 

### Return type

[**UserManagementServiceAddUserResponse**](UserManagementServiceAddUserResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserManagementServiceAddUserTag

> map[string]interface{} UserManagementServiceAddUserTag(ctx, uid).UserManagementServiceAddUserTagRequest(userManagementServiceAddUserTagRequest).Namespace(namespace).Execute()

Updates user tags for the specified user - this is append operation



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
    uid := "uid_example" // string | User Name for the User Tags which are being added
    userManagementServiceAddUserTagRequest := *openapiclient.NewUserManagementServiceAddUserTagRequest() // UserManagementServiceAddUserTagRequest | 
    namespace := "namespace_example" // string | Namespace for the User Tags which are being added (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserManagementApi.UserManagementServiceAddUserTag(context.Background(), uid).UserManagementServiceAddUserTagRequest(userManagementServiceAddUserTagRequest).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.UserManagementServiceAddUserTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserManagementServiceAddUserTag`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.UserManagementServiceAddUserTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | User Name for the User Tags which are being added | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserManagementServiceAddUserTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userManagementServiceAddUserTagRequest** | [**UserManagementServiceAddUserTagRequest**](UserManagementServiceAddUserTagRequest.md) |  | 
 **namespace** | **string** | Namespace for the User Tags which are being added | 

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


## UserManagementServiceGetAllUsers

> UserManagementServiceGetAllUsersResponse UserManagementServiceGetAllUsers(ctx).Limit(limit).Marker(marker).Userid(userid).Execute()

Gets identifiers for all configured users



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
    userid := "userid_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserManagementApi.UserManagementServiceGetAllUsers(context.Background()).Limit(limit).Marker(marker).Userid(userid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.UserManagementServiceGetAllUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserManagementServiceGetAllUsers`: UserManagementServiceGetAllUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.UserManagementServiceGetAllUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserManagementServiceGetAllUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | Number of objects requested in current fetch. | 
 **marker** | **string** | Reference to last object returned. | 
 **userid** | **string** |  | 

### Return type

[**UserManagementServiceGetAllUsersResponse**](UserManagementServiceGetAllUsersResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserManagementServiceGetUserInfo

> UserManagementServiceGetUserInfoResponse UserManagementServiceGetUserInfo(ctx, uid).Namespace(namespace).Execute()

Gets user details for the specified user belonging to specified namespace



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
    uid := "uid_example" // string | Valid user identifier
    namespace := "namespace_example" // string | Optional when userscope is GLOBAL. Required when userscope is NAMESPACE. The namespace to which user belong (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserManagementApi.UserManagementServiceGetUserInfo(context.Background(), uid).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.UserManagementServiceGetUserInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserManagementServiceGetUserInfo`: UserManagementServiceGetUserInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.UserManagementServiceGetUserInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Valid user identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserManagementServiceGetUserInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | Optional when userscope is GLOBAL. Required when userscope is NAMESPACE. The namespace to which user belong | 

### Return type

[**UserManagementServiceGetUserInfoResponse**](UserManagementServiceGetUserInfoResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserManagementServiceGetUserLockWithNamespace

> UserManagementServiceGetUserLockWithNamespaceResponse UserManagementServiceGetUserLockWithNamespace(ctx, uid, namespace).Execute()

Gets the user lock details for the specified user belonging to specified namespace



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
    uid := "uid_example" // string | User name for which user lock status should be returned
    namespace := "namespace_example" // string | Namespace to which user belongs

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserManagementApi.UserManagementServiceGetUserLockWithNamespace(context.Background(), uid, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.UserManagementServiceGetUserLockWithNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserManagementServiceGetUserLockWithNamespace`: UserManagementServiceGetUserLockWithNamespaceResponse
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.UserManagementServiceGetUserLockWithNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | User name for which user lock status should be returned | 
**namespace** | **string** | Namespace to which user belongs | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserManagementServiceGetUserLockWithNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserManagementServiceGetUserLockWithNamespaceResponse**](UserManagementServiceGetUserLockWithNamespaceResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserManagementServiceGetUserLockWithoutNamespace

> UserManagementServiceGetUserLockWithoutNamespaceResponse UserManagementServiceGetUserLockWithoutNamespace(ctx, uid).Execute()

Gets the user lock details for the specified user



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
    uid := "uid_example" // string | User name for which user lock details should be returned

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserManagementApi.UserManagementServiceGetUserLockWithoutNamespace(context.Background(), uid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.UserManagementServiceGetUserLockWithoutNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserManagementServiceGetUserLockWithoutNamespace`: UserManagementServiceGetUserLockWithoutNamespaceResponse
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.UserManagementServiceGetUserLockWithoutNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | User name for which user lock details should be returned | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserManagementServiceGetUserLockWithoutNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserManagementServiceGetUserLockWithoutNamespaceResponse**](UserManagementServiceGetUserLockWithoutNamespaceResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserManagementServiceGetUserTagsWithNamespace

> UserManagementServiceGetUserTagsWithNamespaceResponse UserManagementServiceGetUserTagsWithNamespace(ctx, uid).Namespace(namespace).Execute()

Gets the user tags details for the specified user belonging to specified namespace



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
    uid := "uid_example" // string | User name for which user tags should be returned
    namespace := "namespace_example" // string | Namespace to which user belongs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserManagementApi.UserManagementServiceGetUserTagsWithNamespace(context.Background(), uid).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.UserManagementServiceGetUserTagsWithNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserManagementServiceGetUserTagsWithNamespace`: UserManagementServiceGetUserTagsWithNamespaceResponse
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.UserManagementServiceGetUserTagsWithNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | User name for which user tags should be returned | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserManagementServiceGetUserTagsWithNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | Namespace to which user belongs | 

### Return type

[**UserManagementServiceGetUserTagsWithNamespaceResponse**](UserManagementServiceGetUserTagsWithNamespaceResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserManagementServiceGetUsersForNamespace

> UserManagementServiceGetUsersForNamespaceResponse UserManagementServiceGetUsersForNamespace(ctx, namespace).Limit(limit).Execute()

Gets all user identifiers for the specified namespace



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
    namespace := "namespace_example" // string | Namespace for which users should be returned
    limit := "limit_example" // string | Number of objects requested in current fetch. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserManagementApi.UserManagementServiceGetUsersForNamespace(context.Background(), namespace).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.UserManagementServiceGetUsersForNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserManagementServiceGetUsersForNamespace`: UserManagementServiceGetUsersForNamespaceResponse
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.UserManagementServiceGetUsersForNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Namespace for which users should be returned | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserManagementServiceGetUsersForNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **string** | Number of objects requested in current fetch. | 

### Return type

[**UserManagementServiceGetUsersForNamespaceResponse**](UserManagementServiceGetUsersForNamespaceResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserManagementServiceQueryUsers

> UserManagementServiceQueryUsersResponse UserManagementServiceQueryUsers(ctx).Namespace(namespace).Limit(limit).Marker(marker).Tag(tag).Value(value).Execute()

Gets user details for the specified user belonging to specified namespace



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
    namespace := "namespace_example" // string | Namespace for which users should be returned (optional)
    limit := "limit_example" // string | Number of objects requested in current fetch. (optional)
    marker := "marker_example" // string | Reference to last object returned. (optional)
    tag := "tag_example" // string | User Tag Name (optional)
    value := "value_example" // string | User Tag Value (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserManagementApi.UserManagementServiceQueryUsers(context.Background()).Namespace(namespace).Limit(limit).Marker(marker).Tag(tag).Value(value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.UserManagementServiceQueryUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserManagementServiceQueryUsers`: UserManagementServiceQueryUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.UserManagementServiceQueryUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserManagementServiceQueryUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **string** | Namespace for which users should be returned | 
 **limit** | **string** | Number of objects requested in current fetch. | 
 **marker** | **string** | Reference to last object returned. | 
 **tag** | **string** | User Tag Name | 
 **value** | **string** | User Tag Value | 

### Return type

[**UserManagementServiceQueryUsersResponse**](UserManagementServiceQueryUsersResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserManagementServiceRemoveUser

> map[string]interface{} UserManagementServiceRemoveUser(ctx).UserManagementServiceRemoveUserRequest(userManagementServiceRemoveUserRequest).Execute()

Deletes the specified user and its associated secret keys



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
    userManagementServiceRemoveUserRequest := *openapiclient.NewUserManagementServiceRemoveUserRequest("User_example") // UserManagementServiceRemoveUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserManagementApi.UserManagementServiceRemoveUser(context.Background()).UserManagementServiceRemoveUserRequest(userManagementServiceRemoveUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.UserManagementServiceRemoveUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserManagementServiceRemoveUser`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.UserManagementServiceRemoveUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserManagementServiceRemoveUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userManagementServiceRemoveUserRequest** | [**UserManagementServiceRemoveUserRequest**](UserManagementServiceRemoveUserRequest.md) |  | 

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


## UserManagementServiceRemoveUserTags

> map[string]interface{} UserManagementServiceRemoveUserTags(ctx, uid).UserManagementServiceRemoveUserTagsRequest(userManagementServiceRemoveUserTagsRequest).Namespace(namespace).Execute()

Deletes user tags for specified user



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
    uid := "uid_example" // string | UserName for User Tags which is being deleted
    userManagementServiceRemoveUserTagsRequest := *openapiclient.NewUserManagementServiceRemoveUserTagsRequest() // UserManagementServiceRemoveUserTagsRequest | 
    namespace := "namespace_example" // string | Namespace for the User Tags which are being deleted (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserManagementApi.UserManagementServiceRemoveUserTags(context.Background(), uid).UserManagementServiceRemoveUserTagsRequest(userManagementServiceRemoveUserTagsRequest).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.UserManagementServiceRemoveUserTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserManagementServiceRemoveUserTags`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.UserManagementServiceRemoveUserTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UserName for User Tags which is being deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserManagementServiceRemoveUserTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userManagementServiceRemoveUserTagsRequest** | [**UserManagementServiceRemoveUserTagsRequest**](UserManagementServiceRemoveUserTagsRequest.md) |  | 
 **namespace** | **string** | Namespace for the User Tags which are being deleted | 

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


## UserManagementServiceSetUserLock

> map[string]interface{} UserManagementServiceSetUserLock(ctx).UserManagementServiceSetUserLockRequest(userManagementServiceSetUserLockRequest).Execute()

Locks the specified user



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
    userManagementServiceSetUserLockRequest := *openapiclient.NewUserManagementServiceSetUserLockRequest("User_example", false) // UserManagementServiceSetUserLockRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserManagementApi.UserManagementServiceSetUserLock(context.Background()).UserManagementServiceSetUserLockRequest(userManagementServiceSetUserLockRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.UserManagementServiceSetUserLock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserManagementServiceSetUserLock`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.UserManagementServiceSetUserLock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserManagementServiceSetUserLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userManagementServiceSetUserLockRequest** | [**UserManagementServiceSetUserLockRequest**](UserManagementServiceSetUserLockRequest.md) |  | 

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


## UserManagementServiceUpdateUserTag

> map[string]interface{} UserManagementServiceUpdateUserTag(ctx, uid).UserManagementServiceUpdateUserTagRequest(userManagementServiceUpdateUserTagRequest).Namespace(namespace).Execute()

Updates user tags for the specified user



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
    uid := "uid_example" // string | User Name for the User Tags which are being modified
    userManagementServiceUpdateUserTagRequest := *openapiclient.NewUserManagementServiceUpdateUserTagRequest() // UserManagementServiceUpdateUserTagRequest | 
    namespace := "namespace_example" // string | Namespace for the User Tags which are being modified (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserManagementApi.UserManagementServiceUpdateUserTag(context.Background(), uid).UserManagementServiceUpdateUserTagRequest(userManagementServiceUpdateUserTagRequest).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.UserManagementServiceUpdateUserTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserManagementServiceUpdateUserTag`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.UserManagementServiceUpdateUserTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | User Name for the User Tags which are being modified | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserManagementServiceUpdateUserTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userManagementServiceUpdateUserTagRequest** | [**UserManagementServiceUpdateUserTagRequest**](UserManagementServiceUpdateUserTagRequest.md) |  | 
 **namespace** | **string** | Namespace for the User Tags which are being modified | 

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

