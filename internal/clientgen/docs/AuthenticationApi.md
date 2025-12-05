# \AuthenticationApi

All URIs are relative to *https://objectscale.local:4443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticationResourceGetLoginToken**](AuthenticationApi.md#AuthenticationResourceGetLoginToken) | **Get** /login | Authenticates a user and obtains an authentication token
[**AuthenticationResourceLogout**](AuthenticationApi.md#AuthenticationResourceLogout) | **Get** /logout | User logout



## AuthenticationResourceGetLoginToken

> map[string]interface{} AuthenticationResourceGetLoginToken(ctx).Service(service).Namespace(namespace).Execute()

Authenticates a user and obtains an authentication token



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
    service := "service_example" // string | Optional query parameter, to specify a URL to redirect to on successful          authentication (optional)
    namespace := "namespace_example" // string | namespaces available for the user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.AuthenticationResourceGetLoginToken(context.Background()).Service(service).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.AuthenticationResourceGetLoginToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticationResourceGetLoginToken`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.AuthenticationResourceGetLoginToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationResourceGetLoginTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | **string** | Optional query parameter, to specify a URL to redirect to on successful          authentication | 
 **namespace** | **string** | namespaces available for the user. | 

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


## AuthenticationResourceLogout

> map[string]interface{} AuthenticationResourceLogout(ctx).Force(force).Username(username).Execute()

User logout



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
    force := "force_example" // string | The user with exceeded tokens limit can still call '/logout?force=true' . This will delete all the current logged in session for the user and invalidates all the tokens, (All the tokens are deleted for the user)               Default value: false (optional)
    username := "username_example" // string | A system administrator ( root ) user can log-out the logged in management-user calling '/logout?username=[USER_ID]'. (In an event of a user exceeded it's token limits, a sysadmin user can logout the user) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.AuthenticationResourceLogout(context.Background()).Force(force).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.AuthenticationResourceLogout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticationResourceLogout`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.AuthenticationResourceLogout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationResourceLogoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **force** | **string** | The user with exceeded tokens limit can still call &#39;/logout?force&#x3D;true&#39; . This will delete all the current logged in session for the user and invalidates all the tokens, (All the tokens are deleted for the user)               Default value: false | 
 **username** | **string** | A system administrator ( root ) user can log-out the logged in management-user calling &#39;/logout?username&#x3D;[USER_ID]&#39;. (In an event of a user exceeded it&#39;s token limits, a sysadmin user can logout the user) | 

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

