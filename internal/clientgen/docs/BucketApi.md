# \BucketApi

All URIs are relative to *https://objectscale.local:4443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BucketServiceCreateBucket**](BucketApi.md#BucketServiceCreateBucket) | **Post** /object/bucket | Creates a bucket in which users can create objects
[**BucketServiceGetBuckets**](BucketApi.md#BucketServiceGetBuckets) | **Get** /object/bucket | Gets the list of buckets for the specified namespace



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

