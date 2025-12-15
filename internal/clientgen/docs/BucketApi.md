# \BucketApi

All URIs are relative to *https://objectscale.local:4443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BucketServiceActivateAdvancedMetadataSearch**](BucketApi.md#BucketServiceActivateAdvancedMetadataSearch) | **Put** /object/bucket/{bucketName}/advancedMetadataSearch | Enables advanced metadata search functionality for a bucket.
[**BucketServiceAddBucketTags**](BucketApi.md#BucketServiceAddBucketTags) | **Post** /object/bucket/{bucketName}/tags | Adds the provided tags for the specified bucket.
[**BucketServiceCreateBucket**](BucketApi.md#BucketServiceCreateBucket) | **Post** /object/bucket | Creates a bucket in which users can create objects
[**BucketServiceDeactivateAdvancedMetadataSearch**](BucketApi.md#BucketServiceDeactivateAdvancedMetadataSearch) | **Delete** /object/bucket/{bucketName}/advancedMetadataSearch | Disables advanced metadata search functionality for a bucket.
[**BucketServiceDeactivateBucket**](BucketApi.md#BucketServiceDeactivateBucket) | **Post** /object/bucket/{bucketName}/deactivate | Deletes the specified bucket
[**BucketServiceDeactivateMetaSearch**](BucketApi.md#BucketServiceDeactivateMetaSearch) | **Delete** /object/bucket/{bucketName}/searchmetadata | Disables the metadata search functionality for a bucket.
[**BucketServiceDeleteBucketHeadMetadata**](BucketApi.md#BucketServiceDeleteBucketHeadMetadata) | **Delete** /object/bucket/{bucketName}/metadata | Deletes additional metadata associated with the bucket for a given head-type
[**BucketServiceDeleteBucketPolicy**](BucketApi.md#BucketServiceDeleteBucketPolicy) | **Delete** /object/bucket/{bucketName}/policy | Deletes the bucket policy for the specified bucket.
[**BucketServiceDeleteBucketTags**](BucketApi.md#BucketServiceDeleteBucketTags) | **Delete** /object/bucket/{bucketName}/tags | Deletes the provided tags for the specified bucket.
[**BucketServiceEnableObjectLockWithAdoAllowedForExistingBucket**](BucketApi.md#BucketServiceEnableObjectLockWithAdoAllowedForExistingBucket) | **Put** /object/bucket/{bucketName}/allow-object-lock-with-ado | Sets flag on the bucket to allow Object Lock and ADO to be enabled together.
[**BucketServiceGetBucketACL**](BucketApi.md#BucketServiceGetBucketACL) | **Get** /object/bucket/{bucketName}/acl | Gets the ACL for the given bucket
[**BucketServiceGetBucketDefaultLockConfiguration**](BucketApi.md#BucketServiceGetBucketDefaultLockConfiguration) | **Get** /object/bucket/{bucketName}/object-lock-config | Gets bucket default lock configuration.
[**BucketServiceGetBucketHeadMetadata**](BucketApi.md#BucketServiceGetBucketHeadMetadata) | **Get** /object/bucket/{bucketName}/metadata | Retrieves additional metadata associated with the bucket for a given head-type
[**BucketServiceGetBucketInfo**](BucketApi.md#BucketServiceGetBucketInfo) | **Get** /object/bucket/{bucketName}/info | Gets bucket information for the specified bucket
[**BucketServiceGetBucketLock**](BucketApi.md#BucketServiceGetBucketLock) | **Get** /object/bucket/{bucketName}/lock | Gets lock information for the specified bucket
[**BucketServiceGetBucketNotificationConfig**](BucketApi.md#BucketServiceGetBucketNotificationConfig) | **Get** /object/bucket/{bucketName}/notification | Gets the notification configuration for the specified bucket.
[**BucketServiceGetBucketPolicy**](BucketApi.md#BucketServiceGetBucketPolicy) | **Get** /object/bucket/{bucketName}/policy | Gets policy on the specified bucket
[**BucketServiceGetBucketQuota**](BucketApi.md#BucketServiceGetBucketQuota) | **Get** /object/bucket/{bucketName}/quota | Gets the quota for the given bucket and namespace
[**BucketServiceGetBucketRetention**](BucketApi.md#BucketServiceGetBucketRetention) | **Get** /object/bucket/{bucketName}/retention | Gets the retention period setting for the specified bucket
[**BucketServiceGetBucketVersioning**](BucketApi.md#BucketServiceGetBucketVersioning) | **Get** /object/bucket/{bucketName}/versioning | Gets the versioning status for the specified bucket.
[**BucketServiceGetBuckets**](BucketApi.md#BucketServiceGetBuckets) | **Get** /object/bucket | Gets the list of buckets for the specified namespace
[**BucketServiceGetEmptyBucketStatus**](BucketApi.md#BucketServiceGetEmptyBucketStatus) | **Get** /object/bucket/{bucketName}/empty-bucket-status | Get empty bucket status
[**BucketServiceGetGroups**](BucketApi.md#BucketServiceGetGroups) | **Get** /object/bucket/acl/groups | Gets all ACL groups
[**BucketServiceGetPermissions**](BucketApi.md#BucketServiceGetPermissions) | **Get** /object/bucket/acl/permissions | Gets all ACL permissions
[**BucketServiceGetSearchMetaData**](BucketApi.md#BucketServiceGetSearchMetaData) | **Get** /object/bucket/searchmetadata | Lists the system metadata keys available.
[**BucketServicePutBucketDefaultLockConfiguration**](BucketApi.md#BucketServicePutBucketDefaultLockConfiguration) | **Put** /object/bucket/{bucketName}/object-lock-config | Puts bucket default lock configuration.
[**BucketServicePutBucketNotificationConfig**](BucketApi.md#BucketServicePutBucketNotificationConfig) | **Put** /object/bucket/{bucketName}/notification | Creates or replaces the notification configuration for the bucket.
[**BucketServiceRemoveBucketQuota**](BucketApi.md#BucketServiceRemoveBucketQuota) | **Delete** /object/bucket/{bucketName}/quota | Deletes the quota setting for the given bucket and namespace
[**BucketServiceSetAdvancedMetadataSearchTarget**](BucketApi.md#BucketServiceSetAdvancedMetadataSearchTarget) | **Put** /object/bucket/{bucketName}/advancedMetadataSearchTarget | Sets advanced metadata search target for a bucket.
[**BucketServiceSetBucketACL**](BucketApi.md#BucketServiceSetBucketACL) | **Put** /object/bucket/{bucketName}/acl | Updates the ACL for the given bucket and namespace.
[**BucketServiceSetBucketAuditDeleteExpiration**](BucketApi.md#BucketServiceSetBucketAuditDeleteExpiration) | **Put** /object/bucket/{bucketName}/auditDeleteExpiration | Updates the audit delete expiration for the specified bucket.
[**BucketServiceSetBucketAutoCommitPeriod**](BucketApi.md#BucketServiceSetBucketAutoCommitPeriod) | **Put** /object/bucket/{bucketName}/autocommit | Updates the auto-commit period setting for the specified bucket.
[**BucketServiceSetBucketDefaultGroup**](BucketApi.md#BucketServiceSetBucketDefaultGroup) | **Put** /object/bucket/{bucketName}/defaultGroup | Updates the defaultGroup &amp; defaultGroupPermissions for the given bucket and namespace.
[**BucketServiceSetBucketHeadMetadata**](BucketApi.md#BucketServiceSetBucketHeadMetadata) | **Put** /object/bucket/{bucketName}/metadata | Attaches additional metadata associated with the bucket for a given head-type
[**BucketServiceSetBucketLock**](BucketApi.md#BucketServiceSetBucketLock) | **Put** /object/bucket/{bucketName}/lock/{IsLocked} | Locks or unlocks the specified bucket
[**BucketServiceSetBucketPolicy**](BucketApi.md#BucketServiceSetBucketPolicy) | **Put** /object/bucket/{bucketName}/policy | Add/Replace the policy for the specified bucket in namespace
[**BucketServiceSetBucketRetention**](BucketApi.md#BucketServiceSetBucketRetention) | **Put** /object/bucket/{bucketName}/retention | Updates the default retention period setting for the specified bucket
[**BucketServiceSetBucketVersioning**](BucketApi.md#BucketServiceSetBucketVersioning) | **Put** /object/bucket/{bucketName}/versioning | Updates the versioning status for the specified bucket
[**BucketServiceSetEventualReadsForBucket**](BucketApi.md#BucketServiceSetEventualReadsForBucket) | **Put** /object/bucket/{bucketName}/set-local-object-metadata-reads | Updates local object metadata read flag for a bucket.
[**BucketServiceTestPolicy**](BucketApi.md#BucketServiceTestPolicy) | **Post** /object/bucket/test-policy | Validates a DM policy
[**BucketServiceTestPolicyEdit**](BucketApi.md#BucketServiceTestPolicyEdit) | **Post** /object/bucket/test-policy-edit | Validates a DM policy edit operation
[**BucketServiceUpdateBucketIsStaleAllowed**](BucketApi.md#BucketServiceUpdateBucketIsStaleAllowed) | **Post** /object/bucket/{bucketName}/isstaleallowed | Updates isStaleAllowed details for the specified bucket
[**BucketServiceUpdateBucketOwner**](BucketApi.md#BucketServiceUpdateBucketOwner) | **Post** /object/bucket/{bucketName}/owner | Updates the owner for the specified bucket
[**BucketServiceUpdateBucketQuota**](BucketApi.md#BucketServiceUpdateBucketQuota) | **Put** /object/bucket/{bucketName}/quota | Updates the quota for the given bucket
[**BucketServiceUpdateBucketTags**](BucketApi.md#BucketServiceUpdateBucketTags) | **Put** /object/bucket/{bucketName}/tags | Updates the provided tags for the specified bucket. Note that the operation will over write the existing tags with the new values.



## BucketServiceActivateAdvancedMetadataSearch

> map[string]interface{} BucketServiceActivateAdvancedMetadataSearch(ctx, bucketName).Namespace(namespace).Execute()

Enables advanced metadata search functionality for a bucket.



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
    bucketName := "bucketName_example" // string | Bucket name for which advanced metadata search will be enabled.
    namespace := "namespace_example" // string | Namespace associated. If it is null, then current user's namespace is                    used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceActivateAdvancedMetadataSearch(context.Background(), bucketName).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceActivateAdvancedMetadataSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceActivateAdvancedMetadataSearch`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceActivateAdvancedMetadataSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Bucket name for which advanced metadata search will be enabled. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceActivateAdvancedMetadataSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | Namespace associated. If it is null, then current user&#39;s namespace is                    used. | 

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


## BucketServiceAddBucketTags

> map[string]interface{} BucketServiceAddBucketTags(ctx, bucketName).BucketServiceAddBucketTagsRequest(bucketServiceAddBucketTagsRequest).Execute()

Adds the provided tags for the specified bucket.



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
    bucketName := "bucketName_example" // string | Bucket name for which specified tags will be updated.
    bucketServiceAddBucketTagsRequest := *openapiclient.NewBucketServiceAddBucketTagsRequest() // BucketServiceAddBucketTagsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceAddBucketTags(context.Background(), bucketName).BucketServiceAddBucketTagsRequest(bucketServiceAddBucketTagsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceAddBucketTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceAddBucketTags`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceAddBucketTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Bucket name for which specified tags will be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceAddBucketTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucketServiceAddBucketTagsRequest** | [**BucketServiceAddBucketTagsRequest**](BucketServiceAddBucketTagsRequest.md) |  | 

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


## BucketServiceCreateBucket

> BucketServiceCreateBucketResponse BucketServiceCreateBucket(ctx).BucketServiceCreateBucketRequest(bucketServiceCreateBucketRequest).Execute()

Creates a bucket in which users can create objects



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
    bucketServiceCreateBucketRequest := *openapiclient.NewBucketServiceCreateBucketRequest("Name_example") // BucketServiceCreateBucketRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceCreateBucket(context.Background()).BucketServiceCreateBucketRequest(bucketServiceCreateBucketRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceCreateBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceCreateBucket`: BucketServiceCreateBucketResponse
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceCreateBucket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceCreateBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucketServiceCreateBucketRequest** | [**BucketServiceCreateBucketRequest**](BucketServiceCreateBucketRequest.md) |  | 

### Return type

[**BucketServiceCreateBucketResponse**](BucketServiceCreateBucketResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BucketServiceDeactivateAdvancedMetadataSearch

> map[string]interface{} BucketServiceDeactivateAdvancedMetadataSearch(ctx, bucketName).Namespace(namespace).Execute()

Disables advanced metadata search functionality for a bucket.



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
    bucketName := "bucketName_example" // string | Bucket name for which advanced metadata search will be disabled.
    namespace := "namespace_example" // string | Namespace associated. If it is null, then current user's namespace is                    used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceDeactivateAdvancedMetadataSearch(context.Background(), bucketName).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceDeactivateAdvancedMetadataSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceDeactivateAdvancedMetadataSearch`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceDeactivateAdvancedMetadataSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Bucket name for which advanced metadata search will be disabled. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceDeactivateAdvancedMetadataSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | Namespace associated. If it is null, then current user&#39;s namespace is                    used. | 

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


## BucketServiceDeactivateBucket

> map[string]interface{} BucketServiceDeactivateBucket(ctx, bucketName).Namespace(namespace).EmptyBucket(emptyBucket).Execute()

Deletes the specified bucket



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
    bucketName := "bucketName_example" // string | Bucket name to be deleted
    namespace := "namespace_example" // string | Namespace associated. If it is null, then current user's namespace is used. (optional)
    emptyBucket := "emptyBucket_example" // string | Optional: <b>true</b> | <b>false</b> (default).      If emptyBucket=true the contents of the bucket will be emptied as part of the delete.     The request will return a 202 Accepted if the bucket is not already empty and cleanup was initiated to run in the background.     <br>     The bucket will be read only during the operation.  If the task successfully removes all related items the buket will be deleted.     If the task is unable to remove all items or is aborted the bucket will be put back into a writable state.     <br>     Progress can be monitored through call to:     <br>     /object/bucket/{bucketName}/emtpy-bucket-status     <br>     <br>     If emptyBucket=false or not present the delete bucket operation will fail if the bucket is not empty. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceDeactivateBucket(context.Background(), bucketName).Namespace(namespace).EmptyBucket(emptyBucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceDeactivateBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceDeactivateBucket`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceDeactivateBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Bucket name to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceDeactivateBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | Namespace associated. If it is null, then current user&#39;s namespace is used. | 
 **emptyBucket** | **string** | Optional: &lt;b&gt;true&lt;/b&gt; | &lt;b&gt;false&lt;/b&gt; (default).      If emptyBucket&#x3D;true the contents of the bucket will be emptied as part of the delete.     The request will return a 202 Accepted if the bucket is not already empty and cleanup was initiated to run in the background.     &lt;br&gt;     The bucket will be read only during the operation.  If the task successfully removes all related items the buket will be deleted.     If the task is unable to remove all items or is aborted the bucket will be put back into a writable state.     &lt;br&gt;     Progress can be monitored through call to:     &lt;br&gt;     /object/bucket/{bucketName}/emtpy-bucket-status     &lt;br&gt;     &lt;br&gt;     If emptyBucket&#x3D;false or not present the delete bucket operation will fail if the bucket is not empty. | 

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


## BucketServiceDeactivateMetaSearch

> map[string]interface{} BucketServiceDeactivateMetaSearch(ctx, bucketName).Namespace(namespace).Execute()

Disables the metadata search functionality for a bucket.



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
    bucketName := "bucketName_example" // string | Bucket name for which metadata search mode will be disabled.
    namespace := "namespace_example" // string | Namespace associated. If it is null, then current user's namespace is                    used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceDeactivateMetaSearch(context.Background(), bucketName).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceDeactivateMetaSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceDeactivateMetaSearch`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceDeactivateMetaSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Bucket name for which metadata search mode will be disabled. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceDeactivateMetaSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | Namespace associated. If it is null, then current user&#39;s namespace is                    used. | 

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


## BucketServiceDeleteBucketHeadMetadata

> map[string]interface{} BucketServiceDeleteBucketHeadMetadata(ctx, bucketName).HeadType(headType).Namespace(namespace).Execute()

Deletes additional metadata associated with the bucket for a given head-type



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
    bucketName := "bucketName_example" // string | name of the bucket for which the metadata is to be removed
    headType := "headType_example" // string | the head-type of the metadata to be removed (HDFS, S3, etc) (optional)
    namespace := "namespace_example" // string | namespace of the bucket (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceDeleteBucketHeadMetadata(context.Background(), bucketName).HeadType(headType).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceDeleteBucketHeadMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceDeleteBucketHeadMetadata`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceDeleteBucketHeadMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | name of the bucket for which the metadata is to be removed | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceDeleteBucketHeadMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **headType** | **string** | the head-type of the metadata to be removed (HDFS, S3, etc) | 
 **namespace** | **string** | namespace of the bucket | 

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


## BucketServiceDeleteBucketPolicy

> map[string]interface{} BucketServiceDeleteBucketPolicy(ctx, bucketName).Namespace(namespace).Execute()

Deletes the bucket policy for the specified bucket.



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
    bucketName := "bucketName_example" // string | Bucket name for which the policy will be deleted.
    namespace := "namespace_example" // string | Namespace of the bucket. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceDeleteBucketPolicy(context.Background(), bucketName).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceDeleteBucketPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceDeleteBucketPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceDeleteBucketPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Bucket name for which the policy will be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceDeleteBucketPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | Namespace of the bucket. | 

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


## BucketServiceDeleteBucketTags

> map[string]interface{} BucketServiceDeleteBucketTags(ctx, bucketName).BucketServiceDeleteBucketTagsRequest(bucketServiceDeleteBucketTagsRequest).Execute()

Deletes the provided tags for the specified bucket.



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
    bucketName := "bucketName_example" // string | Bucket name for which specified tags will be updated.
    bucketServiceDeleteBucketTagsRequest := *openapiclient.NewBucketServiceDeleteBucketTagsRequest() // BucketServiceDeleteBucketTagsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceDeleteBucketTags(context.Background(), bucketName).BucketServiceDeleteBucketTagsRequest(bucketServiceDeleteBucketTagsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceDeleteBucketTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceDeleteBucketTags`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceDeleteBucketTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Bucket name for which specified tags will be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceDeleteBucketTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucketServiceDeleteBucketTagsRequest** | [**BucketServiceDeleteBucketTagsRequest**](BucketServiceDeleteBucketTagsRequest.md) |  | 

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


## BucketServiceEnableObjectLockWithAdoAllowedForExistingBucket

> map[string]interface{} BucketServiceEnableObjectLockWithAdoAllowedForExistingBucket(ctx, bucketName).Namespace(namespace).Execute()

Sets flag on the bucket to allow Object Lock and ADO to be enabled together.



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
    bucketName := "bucketName_example" // string | Bucket name for which object lock and ADO will be enabled.
    namespace := "namespace_example" // string | Namespace for the bucket for which object lock and ADO will be enabled. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceEnableObjectLockWithAdoAllowedForExistingBucket(context.Background(), bucketName).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceEnableObjectLockWithAdoAllowedForExistingBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceEnableObjectLockWithAdoAllowedForExistingBucket`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceEnableObjectLockWithAdoAllowedForExistingBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Bucket name for which object lock and ADO will be enabled. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceEnableObjectLockWithAdoAllowedForExistingBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | Namespace for the bucket for which object lock and ADO will be enabled. | 

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


## BucketServiceGetBucketACL

> BucketServiceGetBucketACLResponse BucketServiceGetBucketACL(ctx, bucketName).Namespace(namespace).Execute()

Gets the ACL for the given bucket



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
    bucketName := "bucketName_example" // string | Name of the bucket for which ACL is to be fetched.
    namespace := "namespace_example" // string | Namespace with which bucket is associated. If it is null, the current user's namespace is used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceGetBucketACL(context.Background(), bucketName).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceGetBucketACL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceGetBucketACL`: BucketServiceGetBucketACLResponse
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceGetBucketACL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Name of the bucket for which ACL is to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceGetBucketACLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | Namespace with which bucket is associated. If it is null, the current user&#39;s namespace is used. | 

### Return type

[**BucketServiceGetBucketACLResponse**](BucketServiceGetBucketACLResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BucketServiceGetBucketDefaultLockConfiguration

> BucketServiceGetBucketDefaultLockConfigurationResponse BucketServiceGetBucketDefaultLockConfiguration(ctx, bucketName).Namespace(namespace).Execute()

Gets bucket default lock configuration.



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
    bucketName := "bucketName_example" // string | Bucket name for which default lock configuration will be retrieved.
    namespace := "namespace_example" // string | Namespace of the bucket for which default lock configuration will be retrieved. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceGetBucketDefaultLockConfiguration(context.Background(), bucketName).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceGetBucketDefaultLockConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceGetBucketDefaultLockConfiguration`: BucketServiceGetBucketDefaultLockConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceGetBucketDefaultLockConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Bucket name for which default lock configuration will be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceGetBucketDefaultLockConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | Namespace of the bucket for which default lock configuration will be retrieved. | 

### Return type

[**BucketServiceGetBucketDefaultLockConfigurationResponse**](BucketServiceGetBucketDefaultLockConfigurationResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BucketServiceGetBucketHeadMetadata

> BucketServiceGetBucketHeadMetadataResponse BucketServiceGetBucketHeadMetadata(ctx, bucketName).HeadType(headType).Namespace(namespace).Execute()

Retrieves additional metadata associated with the bucket for a given head-type



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
    bucketName := "bucketName_example" // string | name of the bucket for which head metadata is to be fetched.
    headType := "headType_example" // string | the head-type of the metadata to be queried (HDFS, S3, etc) (optional)
    namespace := "namespace_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceGetBucketHeadMetadata(context.Background(), bucketName).HeadType(headType).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceGetBucketHeadMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceGetBucketHeadMetadata`: BucketServiceGetBucketHeadMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceGetBucketHeadMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | name of the bucket for which head metadata is to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceGetBucketHeadMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **headType** | **string** | the head-type of the metadata to be queried (HDFS, S3, etc) | 
 **namespace** | **string** |  | 

### Return type

[**BucketServiceGetBucketHeadMetadataResponse**](BucketServiceGetBucketHeadMetadataResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BucketServiceGetBucketInfo

> BucketServiceGetBucketInfoResponse BucketServiceGetBucketInfo(ctx, bucketName).Namespace(namespace).Execute()

Gets bucket information for the specified bucket



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
    bucketName := "bucketName_example" // string | Bucket name for which information will be retrieved
    namespace := "namespace_example" // string | Namespace associated. If it is null, then current user's namespace is used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceGetBucketInfo(context.Background(), bucketName).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceGetBucketInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceGetBucketInfo`: BucketServiceGetBucketInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceGetBucketInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Bucket name for which information will be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceGetBucketInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | Namespace associated. If it is null, then current user&#39;s namespace is used. | 

### Return type

[**BucketServiceGetBucketInfoResponse**](BucketServiceGetBucketInfoResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BucketServiceGetBucketLock

> BucketServiceGetBucketLockResponse BucketServiceGetBucketLock(ctx, bucketName).Namespace(namespace).Execute()

Gets lock information for the specified bucket



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
    bucketName := "bucketName_example" // string | Name of the bucket for which lock information is to be retrieved
    namespace := "namespace_example" // string | Namespace associated (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceGetBucketLock(context.Background(), bucketName).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceGetBucketLock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceGetBucketLock`: BucketServiceGetBucketLockResponse
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceGetBucketLock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Name of the bucket for which lock information is to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceGetBucketLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | Namespace associated | 

### Return type

[**BucketServiceGetBucketLockResponse**](BucketServiceGetBucketLockResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BucketServiceGetBucketNotificationConfig

> BucketServiceGetBucketNotificationConfigResponse BucketServiceGetBucketNotificationConfig(ctx, bucketName).Namespace(namespace).Execute()

Gets the notification configuration for the specified bucket.



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
    bucketName := "bucketName_example" // string | Bucket name for which notification config will be retrieved.
    namespace := "namespace_example" // string | Namespace associated with the bucket. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceGetBucketNotificationConfig(context.Background(), bucketName).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceGetBucketNotificationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceGetBucketNotificationConfig`: BucketServiceGetBucketNotificationConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceGetBucketNotificationConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Bucket name for which notification config will be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceGetBucketNotificationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | Namespace associated with the bucket. | 

### Return type

[**BucketServiceGetBucketNotificationConfigResponse**](BucketServiceGetBucketNotificationConfigResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BucketServiceGetBucketPolicy

> map[string]interface{} BucketServiceGetBucketPolicy(ctx, bucketName).Namespace(namespace).Execute()

Gets policy on the specified bucket



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
    bucketName := "bucketName_example" // string | Name of the bucket for which the policy is to be displayed.
    namespace := "namespace_example" // string | Namespace of the bucket. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceGetBucketPolicy(context.Background(), bucketName).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceGetBucketPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceGetBucketPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceGetBucketPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Name of the bucket for which the policy is to be displayed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceGetBucketPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | Namespace of the bucket. | 

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


## BucketServiceGetBucketQuota

> BucketServiceGetBucketQuotaResponse BucketServiceGetBucketQuota(ctx, bucketName).Namespace(namespace).Execute()

Gets the quota for the given bucket and namespace



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
    bucketName := "bucketName_example" // string | Name of the bucket which for which quota is to be retrieved
    namespace := "namespace_example" // string | Namespace with which bucket is associated. If it is null, the current user's namespace is used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceGetBucketQuota(context.Background(), bucketName).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceGetBucketQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceGetBucketQuota`: BucketServiceGetBucketQuotaResponse
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceGetBucketQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Name of the bucket which for which quota is to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceGetBucketQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | Namespace with which bucket is associated. If it is null, the current user&#39;s namespace is used. | 

### Return type

[**BucketServiceGetBucketQuotaResponse**](BucketServiceGetBucketQuotaResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BucketServiceGetBucketRetention

> BucketServiceGetBucketRetentionResponse BucketServiceGetBucketRetention(ctx, bucketName).Namespace(namespace).Execute()

Gets the retention period setting for the specified bucket



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
    bucketName := "bucketName_example" // string | Bucket name for which the retention period setting will be retrieved
    namespace := "namespace_example" // string | Namespace associated. If it is null, then current user's namespace is used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceGetBucketRetention(context.Background(), bucketName).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceGetBucketRetention``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceGetBucketRetention`: BucketServiceGetBucketRetentionResponse
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceGetBucketRetention`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Bucket name for which the retention period setting will be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceGetBucketRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | Namespace associated. If it is null, then current user&#39;s namespace is used. | 

### Return type

[**BucketServiceGetBucketRetentionResponse**](BucketServiceGetBucketRetentionResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BucketServiceGetBucketVersioning

> BucketServiceGetBucketVersioningResponse BucketServiceGetBucketVersioning(ctx, bucketName).Namespace(namespace).Execute()

Gets the versioning status for the specified bucket.



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
    bucketName := "bucketName_example" // string | Bucket name for which versioning will be retrieved.
    namespace := "namespace_example" // string | Namespace associated with the bucket. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceGetBucketVersioning(context.Background(), bucketName).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceGetBucketVersioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceGetBucketVersioning`: BucketServiceGetBucketVersioningResponse
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceGetBucketVersioning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Bucket name for which versioning will be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceGetBucketVersioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | Namespace associated with the bucket. | 

### Return type

[**BucketServiceGetBucketVersioningResponse**](BucketServiceGetBucketVersioningResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BucketServiceGetBuckets

> BucketServiceGetBucketsResponse BucketServiceGetBuckets(ctx).Namespace(namespace).Marker(marker).Limit(limit).Name(name).Execute()

Gets the list of buckets for the specified namespace



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
    namespace := "namespace_example" // string | Namespace for which buckets should be listed. (optional)
    marker := "marker_example" // string | reference to last object returned. (optional)
    limit := "limit_example" // string | number of objects requested in current fetch. (optional)
    name := "name_example" // string | Case sensitive prefix of the Bucket name with a wild card(*) Ex : any_prefix_string* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceGetBuckets(context.Background()).Namespace(namespace).Marker(marker).Limit(limit).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceGetBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceGetBuckets`: BucketServiceGetBucketsResponse
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceGetBuckets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceGetBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **string** | Namespace for which buckets should be listed. | 
 **marker** | **string** | reference to last object returned. | 
 **limit** | **string** | number of objects requested in current fetch. | 
 **name** | **string** | Case sensitive prefix of the Bucket name with a wild card(*) Ex : any_prefix_string* | 

### Return type

[**BucketServiceGetBucketsResponse**](BucketServiceGetBucketsResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BucketServiceGetEmptyBucketStatus

> BucketServiceGetEmptyBucketStatusResponse BucketServiceGetEmptyBucketStatus(ctx, bucketName).Namespace(namespace).Execute()

Get empty bucket status



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
    bucketName := "bucketName_example" // string | Name of the bucket for which lock information is to be retrieved
    namespace := "namespace_example" // string | Namespace associated with the bucket. If not present the user's namespace is used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceGetEmptyBucketStatus(context.Background(), bucketName).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceGetEmptyBucketStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceGetEmptyBucketStatus`: BucketServiceGetEmptyBucketStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceGetEmptyBucketStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Name of the bucket for which lock information is to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceGetEmptyBucketStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | Namespace associated with the bucket. If not present the user&#39;s namespace is used. | 

### Return type

[**BucketServiceGetEmptyBucketStatusResponse**](BucketServiceGetEmptyBucketStatusResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BucketServiceGetGroups

> BucketServiceGetGroupsResponse BucketServiceGetGroups(ctx).Execute()

Gets all ACL groups



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
    resp, r, err := apiClient.BucketApi.BucketServiceGetGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceGetGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceGetGroups`: BucketServiceGetGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceGetGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceGetGroupsRequest struct via the builder pattern


### Return type

[**BucketServiceGetGroupsResponse**](BucketServiceGetGroupsResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BucketServiceGetPermissions

> BucketServiceGetPermissionsResponse BucketServiceGetPermissions(ctx).Execute()

Gets all ACL permissions



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
    resp, r, err := apiClient.BucketApi.BucketServiceGetPermissions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceGetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceGetPermissions`: BucketServiceGetPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceGetPermissions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceGetPermissionsRequest struct via the builder pattern


### Return type

[**BucketServiceGetPermissionsResponse**](BucketServiceGetPermissionsResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BucketServiceGetSearchMetaData

> BucketServiceGetSearchMetaDataResponse BucketServiceGetSearchMetaData(ctx).Execute()

Lists the system metadata keys available.



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
    resp, r, err := apiClient.BucketApi.BucketServiceGetSearchMetaData(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceGetSearchMetaData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceGetSearchMetaData`: BucketServiceGetSearchMetaDataResponse
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceGetSearchMetaData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceGetSearchMetaDataRequest struct via the builder pattern


### Return type

[**BucketServiceGetSearchMetaDataResponse**](BucketServiceGetSearchMetaDataResponse.md)

### Authorization

[AuthToken](../README.md#AuthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BucketServicePutBucketDefaultLockConfiguration

> map[string]interface{} BucketServicePutBucketDefaultLockConfiguration(ctx, bucketName).BucketServicePutBucketDefaultLockConfigurationRequest(bucketServicePutBucketDefaultLockConfigurationRequest).Namespace(namespace).Execute()

Puts bucket default lock configuration.



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
    bucketName := "bucketName_example" // string | Bucket name for which default lock configuration will be updated.
    bucketServicePutBucketDefaultLockConfigurationRequest := *openapiclient.NewBucketServicePutBucketDefaultLockConfigurationRequest() // BucketServicePutBucketDefaultLockConfigurationRequest | 
    namespace := "namespace_example" // string | Namespace of the bucket for which default lock configuration will be updated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServicePutBucketDefaultLockConfiguration(context.Background(), bucketName).BucketServicePutBucketDefaultLockConfigurationRequest(bucketServicePutBucketDefaultLockConfigurationRequest).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServicePutBucketDefaultLockConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServicePutBucketDefaultLockConfiguration`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServicePutBucketDefaultLockConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Bucket name for which default lock configuration will be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServicePutBucketDefaultLockConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucketServicePutBucketDefaultLockConfigurationRequest** | [**BucketServicePutBucketDefaultLockConfigurationRequest**](BucketServicePutBucketDefaultLockConfigurationRequest.md) |  | 
 **namespace** | **string** | Namespace of the bucket for which default lock configuration will be updated. | 

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


## BucketServicePutBucketNotificationConfig

> map[string]interface{} BucketServicePutBucketNotificationConfig(ctx, bucketName).BucketServicePutBucketNotificationConfigRequest(bucketServicePutBucketNotificationConfigRequest).Namespace(namespace).XAmzSkipDestinationValidation(xAmzSkipDestinationValidation).Execute()

Creates or replaces the notification configuration for the bucket.



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
    bucketName := "bucketName_example" // string | Bucket name for which notification configuration will be updated.
    bucketServicePutBucketNotificationConfigRequest := *openapiclient.NewBucketServicePutBucketNotificationConfigRequest() // BucketServicePutBucketNotificationConfigRequest | 
    namespace := "namespace_example" // string | Namespace associated with the bucket for which notification config will be updated. (optional)
    xAmzSkipDestinationValidation := "xAmzSkipDestinationValidation_example" // string | Optional header to skip destination validation.                                                                             If set to true, destination validation will be skipped. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServicePutBucketNotificationConfig(context.Background(), bucketName).BucketServicePutBucketNotificationConfigRequest(bucketServicePutBucketNotificationConfigRequest).Namespace(namespace).XAmzSkipDestinationValidation(xAmzSkipDestinationValidation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServicePutBucketNotificationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServicePutBucketNotificationConfig`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServicePutBucketNotificationConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Bucket name for which notification configuration will be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServicePutBucketNotificationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucketServicePutBucketNotificationConfigRequest** | [**BucketServicePutBucketNotificationConfigRequest**](BucketServicePutBucketNotificationConfigRequest.md) |  | 
 **namespace** | **string** | Namespace associated with the bucket for which notification config will be updated. | 
 **xAmzSkipDestinationValidation** | **string** | Optional header to skip destination validation.                                                                             If set to true, destination validation will be skipped. | 

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


## BucketServiceRemoveBucketQuota

> map[string]interface{} BucketServiceRemoveBucketQuota(ctx, bucketName).Namespace(namespace).Execute()

Deletes the quota setting for the given bucket and namespace



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
    bucketName := "bucketName_example" // string | Name of the bucket for which the quota is to be deleted
    namespace := "namespace_example" // string | Namespace with which bucket is associated. If it is null, the current user's namespace is used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceRemoveBucketQuota(context.Background(), bucketName).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceRemoveBucketQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceRemoveBucketQuota`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceRemoveBucketQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Name of the bucket for which the quota is to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceRemoveBucketQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | Namespace with which bucket is associated. If it is null, the current user&#39;s namespace is used. | 

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


## BucketServiceSetAdvancedMetadataSearchTarget

> map[string]interface{} BucketServiceSetAdvancedMetadataSearchTarget(ctx, bucketName).BucketServiceSetAdvancedMetadataSearchTargetRequest(bucketServiceSetAdvancedMetadataSearchTargetRequest).Namespace(namespace).Execute()

Sets advanced metadata search target for a bucket.



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
    bucketName := "bucketName_example" // string | Bucket name on which advanced metadata search target will be set.
    bucketServiceSetAdvancedMetadataSearchTargetRequest := *openapiclient.NewBucketServiceSetAdvancedMetadataSearchTargetRequest() // BucketServiceSetAdvancedMetadataSearchTargetRequest | 
    namespace := "namespace_example" // string | Namespace associated. If it is null, then current user's namespace is                    used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceSetAdvancedMetadataSearchTarget(context.Background(), bucketName).BucketServiceSetAdvancedMetadataSearchTargetRequest(bucketServiceSetAdvancedMetadataSearchTargetRequest).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceSetAdvancedMetadataSearchTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceSetAdvancedMetadataSearchTarget`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceSetAdvancedMetadataSearchTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Bucket name on which advanced metadata search target will be set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceSetAdvancedMetadataSearchTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucketServiceSetAdvancedMetadataSearchTargetRequest** | [**BucketServiceSetAdvancedMetadataSearchTargetRequest**](BucketServiceSetAdvancedMetadataSearchTargetRequest.md) |  | 
 **namespace** | **string** | Namespace associated. If it is null, then current user&#39;s namespace is                    used. | 

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


## BucketServiceSetBucketACL

> map[string]interface{} BucketServiceSetBucketACL(ctx, bucketName).BucketServiceSetBucketACLRequest(bucketServiceSetBucketACLRequest).Execute()

Updates the ACL for the given bucket and namespace.



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
    bucketName := "bucketName_example" // string | Name of the bucket for which the ACL is to be updated.
    bucketServiceSetBucketACLRequest := *openapiclient.NewBucketServiceSetBucketACLRequest() // BucketServiceSetBucketACLRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceSetBucketACL(context.Background(), bucketName).BucketServiceSetBucketACLRequest(bucketServiceSetBucketACLRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceSetBucketACL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceSetBucketACL`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceSetBucketACL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Name of the bucket for which the ACL is to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceSetBucketACLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucketServiceSetBucketACLRequest** | [**BucketServiceSetBucketACLRequest**](BucketServiceSetBucketACLRequest.md) |  | 

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


## BucketServiceSetBucketAuditDeleteExpiration

> map[string]interface{} BucketServiceSetBucketAuditDeleteExpiration(ctx, bucketName).Namespace(namespace).Expiration(expiration).Execute()

Updates the audit delete expiration for the specified bucket.



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
    bucketName := "bucketName_example" // string | Name of the bucket for which audit delete expiration will be updated
    namespace := "namespace_example" // string | Namespace associated. If it is null, then current user's namespace is used. (optional)
    expiration := "expiration_example" // string | Bucket's audit delete expiration in seconds (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceSetBucketAuditDeleteExpiration(context.Background(), bucketName).Namespace(namespace).Expiration(expiration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceSetBucketAuditDeleteExpiration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceSetBucketAuditDeleteExpiration`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceSetBucketAuditDeleteExpiration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Name of the bucket for which audit delete expiration will be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceSetBucketAuditDeleteExpirationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | Namespace associated. If it is null, then current user&#39;s namespace is used. | 
 **expiration** | **string** | Bucket&#39;s audit delete expiration in seconds | 

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


## BucketServiceSetBucketAutoCommitPeriod

> map[string]interface{} BucketServiceSetBucketAutoCommitPeriod(ctx, bucketName).BucketServiceSetBucketAutoCommitPeriodRequest(bucketServiceSetBucketAutoCommitPeriodRequest).Execute()

Updates the auto-commit period setting for the specified bucket.



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
    bucketName := "bucketName_example" // string | Bucket name for which retention period will be updated.
    bucketServiceSetBucketAutoCommitPeriodRequest := *openapiclient.NewBucketServiceSetBucketAutoCommitPeriodRequest("Autocommit_example", "Namespace_example") // BucketServiceSetBucketAutoCommitPeriodRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceSetBucketAutoCommitPeriod(context.Background(), bucketName).BucketServiceSetBucketAutoCommitPeriodRequest(bucketServiceSetBucketAutoCommitPeriodRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceSetBucketAutoCommitPeriod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceSetBucketAutoCommitPeriod`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceSetBucketAutoCommitPeriod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Bucket name for which retention period will be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceSetBucketAutoCommitPeriodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucketServiceSetBucketAutoCommitPeriodRequest** | [**BucketServiceSetBucketAutoCommitPeriodRequest**](BucketServiceSetBucketAutoCommitPeriodRequest.md) |  | 

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


## BucketServiceSetBucketDefaultGroup

> map[string]interface{} BucketServiceSetBucketDefaultGroup(ctx, bucketName).BucketServiceSetBucketDefaultGroupRequest(bucketServiceSetBucketDefaultGroupRequest).Execute()

Updates the defaultGroup & defaultGroupPermissions for the given bucket and namespace.



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
    bucketName := "bucketName_example" // string | Name of the bucket for which the default group is to be updated.
    bucketServiceSetBucketDefaultGroupRequest := *openapiclient.NewBucketServiceSetBucketDefaultGroupRequest() // BucketServiceSetBucketDefaultGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceSetBucketDefaultGroup(context.Background(), bucketName).BucketServiceSetBucketDefaultGroupRequest(bucketServiceSetBucketDefaultGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceSetBucketDefaultGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceSetBucketDefaultGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceSetBucketDefaultGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Name of the bucket for which the default group is to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceSetBucketDefaultGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucketServiceSetBucketDefaultGroupRequest** | [**BucketServiceSetBucketDefaultGroupRequest**](BucketServiceSetBucketDefaultGroupRequest.md) |  | 

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


## BucketServiceSetBucketHeadMetadata

> map[string]interface{} BucketServiceSetBucketHeadMetadata(ctx, bucketName).BucketServiceSetBucketHeadMetadataRequest(bucketServiceSetBucketHeadMetadataRequest).Namespace(namespace).Execute()

Attaches additional metadata associated with the bucket for a given head-type



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
    bucketName := "bucketName_example" // string | name of the bucket for which the metadata is to be added
    bucketServiceSetBucketHeadMetadataRequest := *openapiclient.NewBucketServiceSetBucketHeadMetadataRequest() // BucketServiceSetBucketHeadMetadataRequest | 
    namespace := "namespace_example" // string | namespace of the bucket (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceSetBucketHeadMetadata(context.Background(), bucketName).BucketServiceSetBucketHeadMetadataRequest(bucketServiceSetBucketHeadMetadataRequest).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceSetBucketHeadMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceSetBucketHeadMetadata`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceSetBucketHeadMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | name of the bucket for which the metadata is to be added | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceSetBucketHeadMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucketServiceSetBucketHeadMetadataRequest** | [**BucketServiceSetBucketHeadMetadataRequest**](BucketServiceSetBucketHeadMetadataRequest.md) |  | 
 **namespace** | **string** | namespace of the bucket | 

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


## BucketServiceSetBucketLock

> map[string]interface{} BucketServiceSetBucketLock(ctx, bucketName, isLocked).BucketServiceSetBucketLockRequest(bucketServiceSetBucketLockRequest).Execute()

Locks or unlocks the specified bucket



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
    bucketName := "bucketName_example" // string | Name of the bucket which is to be locked/unlocked.
    isLocked := "isLocked_example" // string | Set to \"true\" for lock bucket and \"false\" for unlock bucket.
    bucketServiceSetBucketLockRequest := *openapiclient.NewBucketServiceSetBucketLockRequest() // BucketServiceSetBucketLockRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceSetBucketLock(context.Background(), bucketName, isLocked).BucketServiceSetBucketLockRequest(bucketServiceSetBucketLockRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceSetBucketLock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceSetBucketLock`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceSetBucketLock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Name of the bucket which is to be locked/unlocked. | 
**isLocked** | **string** | Set to \&quot;true\&quot; for lock bucket and \&quot;false\&quot; for unlock bucket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceSetBucketLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bucketServiceSetBucketLockRequest** | [**BucketServiceSetBucketLockRequest**](BucketServiceSetBucketLockRequest.md) |  | 

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


## BucketServiceSetBucketPolicy

> map[string]interface{} BucketServiceSetBucketPolicy(ctx, bucketName).Body(body).Namespace(namespace).Execute()

Add/Replace the policy for the specified bucket in namespace



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
    bucketName := "bucketName_example" // string | Name of the bucket for which the policy is to be updated.
    body := map[string]interface{}{ ... } // map[string]interface{} | 
    namespace := "namespace_example" // string | namespace of the bucket (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceSetBucketPolicy(context.Background(), bucketName).Body(body).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceSetBucketPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceSetBucketPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceSetBucketPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Name of the bucket for which the policy is to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceSetBucketPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 
 **namespace** | **string** | namespace of the bucket | 

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


## BucketServiceSetBucketRetention

> map[string]interface{} BucketServiceSetBucketRetention(ctx, bucketName).BucketServiceSetBucketRetentionRequest(bucketServiceSetBucketRetentionRequest).Execute()

Updates the default retention period setting for the specified bucket



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
    bucketName := "bucketName_example" // string | Bucket name for which retention period will be updated.
    bucketServiceSetBucketRetentionRequest := *openapiclient.NewBucketServiceSetBucketRetentionRequest() // BucketServiceSetBucketRetentionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceSetBucketRetention(context.Background(), bucketName).BucketServiceSetBucketRetentionRequest(bucketServiceSetBucketRetentionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceSetBucketRetention``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceSetBucketRetention`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceSetBucketRetention`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Bucket name for which retention period will be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceSetBucketRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucketServiceSetBucketRetentionRequest** | [**BucketServiceSetBucketRetentionRequest**](BucketServiceSetBucketRetentionRequest.md) |  | 

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


## BucketServiceSetBucketVersioning

> map[string]interface{} BucketServiceSetBucketVersioning(ctx, bucketName).BucketServiceSetBucketVersioningRequest(bucketServiceSetBucketVersioningRequest).Namespace(namespace).Execute()

Updates the versioning status for the specified bucket



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
    bucketName := "bucketName_example" // string | Bucket name for which versioning will be updated.
    bucketServiceSetBucketVersioningRequest := *openapiclient.NewBucketServiceSetBucketVersioningRequest() // BucketServiceSetBucketVersioningRequest | 
    namespace := "namespace_example" // string | Namespace associated with the bucket for which versioning will be updated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceSetBucketVersioning(context.Background(), bucketName).BucketServiceSetBucketVersioningRequest(bucketServiceSetBucketVersioningRequest).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceSetBucketVersioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceSetBucketVersioning`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceSetBucketVersioning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Bucket name for which versioning will be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceSetBucketVersioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucketServiceSetBucketVersioningRequest** | [**BucketServiceSetBucketVersioningRequest**](BucketServiceSetBucketVersioningRequest.md) |  | 
 **namespace** | **string** | Namespace associated with the bucket for which versioning will be updated. | 

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


## BucketServiceSetEventualReadsForBucket

> map[string]interface{} BucketServiceSetEventualReadsForBucket(ctx, bucketName).Namespace(namespace).Enabled(enabled).Execute()

Updates local object metadata read flag for a bucket.



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
    bucketName := "bucketName_example" // string | Bucket name for which the OBS CAS local object metadata reads should be updated.
    namespace := "namespace_example" // string | Namespace of the bucket for which the OBS CAS local object metadata reads should be updated. (optional)
    enabled := "enabled_example" // string | Enable or disable OBS CAS local object metadata reads on the bucket buckets. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceSetEventualReadsForBucket(context.Background(), bucketName).Namespace(namespace).Enabled(enabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceSetEventualReadsForBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceSetEventualReadsForBucket`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceSetEventualReadsForBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Bucket name for which the OBS CAS local object metadata reads should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceSetEventualReadsForBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | Namespace of the bucket for which the OBS CAS local object metadata reads should be updated. | 
 **enabled** | **string** | Enable or disable OBS CAS local object metadata reads on the bucket buckets. | 

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


## BucketServiceTestPolicy

> map[string]interface{} BucketServiceTestPolicy(ctx).BucketServiceTestPolicyRequest(bucketServiceTestPolicyRequest).BucketName(bucketName).Account(account).Execute()

Validates a DM policy



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
    bucketServiceTestPolicyRequest := *openapiclient.NewBucketServiceTestPolicyRequest() // BucketServiceTestPolicyRequest | 
    bucketName := "bucketName_example" // string | Bucket name for which DM policy should be validated. (optional)
    account := "account_example" // string | Namespace for which DM policy should be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceTestPolicy(context.Background()).BucketServiceTestPolicyRequest(bucketServiceTestPolicyRequest).BucketName(bucketName).Account(account).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceTestPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceTestPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceTestPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceTestPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucketServiceTestPolicyRequest** | [**BucketServiceTestPolicyRequest**](BucketServiceTestPolicyRequest.md) |  | 
 **bucketName** | **string** | Bucket name for which DM policy should be validated. | 
 **account** | **string** | Namespace for which DM policy should be validated. | 

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


## BucketServiceTestPolicyEdit

> map[string]interface{} BucketServiceTestPolicyEdit(ctx).BucketServiceTestPolicyEditRequest(bucketServiceTestPolicyEditRequest).BucketName(bucketName).Account(account).Execute()

Validates a DM policy edit operation



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
    bucketServiceTestPolicyEditRequest := *openapiclient.NewBucketServiceTestPolicyEditRequest() // BucketServiceTestPolicyEditRequest | 
    bucketName := "bucketName_example" // string | Bucket name for which DM policy should be validated. (optional)
    account := "account_example" // string | Namespace for which DM policy should be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceTestPolicyEdit(context.Background()).BucketServiceTestPolicyEditRequest(bucketServiceTestPolicyEditRequest).BucketName(bucketName).Account(account).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceTestPolicyEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceTestPolicyEdit`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceTestPolicyEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceTestPolicyEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucketServiceTestPolicyEditRequest** | [**BucketServiceTestPolicyEditRequest**](BucketServiceTestPolicyEditRequest.md) |  | 
 **bucketName** | **string** | Bucket name for which DM policy should be validated. | 
 **account** | **string** | Namespace for which DM policy should be validated. | 

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


## BucketServiceUpdateBucketIsStaleAllowed

> map[string]interface{} BucketServiceUpdateBucketIsStaleAllowed(ctx, bucketName).BucketServiceUpdateBucketIsStaleAllowedRequest(bucketServiceUpdateBucketIsStaleAllowedRequest).Execute()

Updates isStaleAllowed details for the specified bucket



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
    bucketName := "bucketName_example" // string | Name of the bucket for which isStaleAllowed is to be updated
    bucketServiceUpdateBucketIsStaleAllowedRequest := *openapiclient.NewBucketServiceUpdateBucketIsStaleAllowedRequest("IsStaleAllowed_example") // BucketServiceUpdateBucketIsStaleAllowedRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceUpdateBucketIsStaleAllowed(context.Background(), bucketName).BucketServiceUpdateBucketIsStaleAllowedRequest(bucketServiceUpdateBucketIsStaleAllowedRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceUpdateBucketIsStaleAllowed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceUpdateBucketIsStaleAllowed`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceUpdateBucketIsStaleAllowed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Name of the bucket for which isStaleAllowed is to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceUpdateBucketIsStaleAllowedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucketServiceUpdateBucketIsStaleAllowedRequest** | [**BucketServiceUpdateBucketIsStaleAllowedRequest**](BucketServiceUpdateBucketIsStaleAllowedRequest.md) |  | 

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


## BucketServiceUpdateBucketOwner

> map[string]interface{} BucketServiceUpdateBucketOwner(ctx, bucketName).BucketServiceUpdateBucketOwnerRequest(bucketServiceUpdateBucketOwnerRequest).Execute()

Updates the owner for the specified bucket



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
    bucketName := "bucketName_example" // string | Name of the bucket for which owner will be updated
    bucketServiceUpdateBucketOwnerRequest := *openapiclient.NewBucketServiceUpdateBucketOwnerRequest("Namespace_example", "NewOwner_example") // BucketServiceUpdateBucketOwnerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceUpdateBucketOwner(context.Background(), bucketName).BucketServiceUpdateBucketOwnerRequest(bucketServiceUpdateBucketOwnerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceUpdateBucketOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceUpdateBucketOwner`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceUpdateBucketOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Name of the bucket for which owner will be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceUpdateBucketOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucketServiceUpdateBucketOwnerRequest** | [**BucketServiceUpdateBucketOwnerRequest**](BucketServiceUpdateBucketOwnerRequest.md) |  | 

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


## BucketServiceUpdateBucketQuota

> map[string]interface{} BucketServiceUpdateBucketQuota(ctx, bucketName).BucketServiceUpdateBucketQuotaRequest(bucketServiceUpdateBucketQuotaRequest).Execute()

Updates the quota for the given bucket



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
    bucketName := "bucketName_example" // string | Name of the bucket for which the quota is to be updated.
    bucketServiceUpdateBucketQuotaRequest := *openapiclient.NewBucketServiceUpdateBucketQuotaRequest() // BucketServiceUpdateBucketQuotaRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceUpdateBucketQuota(context.Background(), bucketName).BucketServiceUpdateBucketQuotaRequest(bucketServiceUpdateBucketQuotaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceUpdateBucketQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceUpdateBucketQuota`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceUpdateBucketQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Name of the bucket for which the quota is to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceUpdateBucketQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucketServiceUpdateBucketQuotaRequest** | [**BucketServiceUpdateBucketQuotaRequest**](BucketServiceUpdateBucketQuotaRequest.md) |  | 

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


## BucketServiceUpdateBucketTags

> map[string]interface{} BucketServiceUpdateBucketTags(ctx, bucketName).BucketServiceUpdateBucketTagsRequest(bucketServiceUpdateBucketTagsRequest).Execute()

Updates the provided tags for the specified bucket. Note that the operation will over write the existing tags with the new values.



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
    bucketName := "bucketName_example" // string | Bucket name for which specified tags will be updated.
    bucketServiceUpdateBucketTagsRequest := *openapiclient.NewBucketServiceUpdateBucketTagsRequest() // BucketServiceUpdateBucketTagsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketApi.BucketServiceUpdateBucketTags(context.Background(), bucketName).BucketServiceUpdateBucketTagsRequest(bucketServiceUpdateBucketTagsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketApi.BucketServiceUpdateBucketTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketServiceUpdateBucketTags`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketApi.BucketServiceUpdateBucketTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | Bucket name for which specified tags will be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketServiceUpdateBucketTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucketServiceUpdateBucketTagsRequest** | [**BucketServiceUpdateBucketTagsRequest**](BucketServiceUpdateBucketTagsRequest.md) |  | 

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

