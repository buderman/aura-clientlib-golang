# \DefaultApi

All URIs are relative to *https://aura.akamai.com:443/api/lcdn-deployment/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeployHyperCacheInstance**](DefaultApi.md#DeployHyperCacheInstance) | **Put** /hypercache/instances/{nodeId} | HyperCache Instance.
[**DeployRequestRouterInstance**](DefaultApi.md#DeployRequestRouterInstance) | **Put** /request-router/instances/{nodeId} | RequestRouter Instance.
[**ListHyperCacheInstances**](DefaultApi.md#ListHyperCacheInstances) | **Get** /hypercache/instances | HyperCache Instances
[**ListRequestRouterInstances**](DefaultApi.md#ListRequestRouterInstances) | **Get** /request-router/instances | RequestRouter Instances
[**ReadHyperCacheInstance**](DefaultApi.md#ReadHyperCacheInstance) | **Get** /hypercache/instances/{nodeId} | HyperCache Instance.
[**ReadRequestRouterInstance**](DefaultApi.md#ReadRequestRouterInstance) | **Get** /request-router/instances/{nodeId} | RequestRouter Instance.
[**UndeployHyperCacheInstance**](DefaultApi.md#UndeployHyperCacheInstance) | **Delete** /hypercache/instances/{nodeId} | HyperCache Instance.
[**UndeployRequestRouterInstance**](DefaultApi.md#UndeployRequestRouterInstance) | **Delete** /request-router/instances/{nodeId} | RequestRouter Instance.



## DeployHyperCacheInstance

> HypercacheInstanceReadDetailed DeployHyperCacheInstance(ctx, nodeId).HypercacheInstanceUpdate(hypercacheInstanceUpdate).Execute()

HyperCache Instance.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    nodeId := int32(56) // int32 | Uniquely identifies a HyperCache instance. 
    hypercacheInstanceUpdate := *openapiclient.NewHypercacheInstanceUpdate(int32(123), "License_example") // HypercacheInstanceUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeployHyperCacheInstance(context.Background(), nodeId).HypercacheInstanceUpdate(hypercacheInstanceUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeployHyperCacheInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeployHyperCacheInstance`: HypercacheInstanceReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeployHyperCacheInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **int32** | Uniquely identifies a HyperCache instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployHyperCacheInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hypercacheInstanceUpdate** | [**HypercacheInstanceUpdate**](HypercacheInstanceUpdate.md) |  | 

### Return type

[**HypercacheInstanceReadDetailed**](HypercacheInstanceReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeployRequestRouterInstance

> RequestrouterInstanceReadDetailed DeployRequestRouterInstance(ctx, nodeId).RequestrouterInstanceUpdate(requestrouterInstanceUpdate).Execute()

RequestRouter Instance.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    nodeId := int32(56) // int32 | Uniquely identifies a RequestRouter instance. 
    requestrouterInstanceUpdate := *openapiclient.NewRequestrouterInstanceUpdate(int32(123)) // RequestrouterInstanceUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeployRequestRouterInstance(context.Background(), nodeId).RequestrouterInstanceUpdate(requestrouterInstanceUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeployRequestRouterInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeployRequestRouterInstance`: RequestrouterInstanceReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeployRequestRouterInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **int32** | Uniquely identifies a RequestRouter instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployRequestRouterInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestrouterInstanceUpdate** | [**RequestrouterInstanceUpdate**](RequestrouterInstanceUpdate.md) |  | 

### Return type

[**RequestrouterInstanceReadDetailed**](RequestrouterInstanceReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHyperCacheInstances

> HypercacheInstanceCollection ListHyperCacheInstances(ctx).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()

HyperCache Instances



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    view := "view_example" // string | The view parameter may be used to request that either the summary or detailed view be returned. (optional)
    filter := "filter_example" // string | The filter parameter may be used to filter the returned collection. (optional)
    sortBy := "sortBy_example" // string | The sortBy parameter may be used to specify the sort order of the returned collection. (optional)
    pageNumber := int32(56) // int32 | The pageNumber parameter may be used to specify an offset into the results.  Useful to use in conjunction with pageSize.  Page number offset is 1-based. (optional)
    pageSize := int32(56) // int32 | The pageSize parameter may be used to specify the page size of the returned results.  Useful to use in conjunction with pageNumber. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListHyperCacheInstances(context.Background()).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListHyperCacheInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHyperCacheInstances`: HypercacheInstanceCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListHyperCacheInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHyperCacheInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **string** | The view parameter may be used to request that either the summary or detailed view be returned. | 
 **filter** | **string** | The filter parameter may be used to filter the returned collection. | 
 **sortBy** | **string** | The sortBy parameter may be used to specify the sort order of the returned collection. | 
 **pageNumber** | **int32** | The pageNumber parameter may be used to specify an offset into the results.  Useful to use in conjunction with pageSize.  Page number offset is 1-based. | 
 **pageSize** | **int32** | The pageSize parameter may be used to specify the page size of the returned results.  Useful to use in conjunction with pageNumber. | 

### Return type

[**HypercacheInstanceCollection**](HypercacheInstanceCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRequestRouterInstances

> RequestrouterInstanceCollection ListRequestRouterInstances(ctx).Execute()

RequestRouter Instances



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListRequestRouterInstances(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRequestRouterInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRequestRouterInstances`: RequestrouterInstanceCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRequestRouterInstances`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRequestRouterInstancesRequest struct via the builder pattern


### Return type

[**RequestrouterInstanceCollection**](RequestrouterInstanceCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadHyperCacheInstance

> HypercacheInstanceReadDetailed ReadHyperCacheInstance(ctx, nodeId).Execute()

HyperCache Instance.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    nodeId := int32(56) // int32 | Uniquely identifies a HyperCache instance. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ReadHyperCacheInstance(context.Background(), nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReadHyperCacheInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadHyperCacheInstance`: HypercacheInstanceReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReadHyperCacheInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **int32** | Uniquely identifies a HyperCache instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadHyperCacheInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HypercacheInstanceReadDetailed**](HypercacheInstanceReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadRequestRouterInstance

> RequestrouterInstanceReadDetailed ReadRequestRouterInstance(ctx, nodeId).Execute()

RequestRouter Instance.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    nodeId := int32(56) // int32 | Uniquely identifies a RequestRouter instance. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ReadRequestRouterInstance(context.Background(), nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReadRequestRouterInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadRequestRouterInstance`: RequestrouterInstanceReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReadRequestRouterInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **int32** | Uniquely identifies a RequestRouter instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadRequestRouterInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestrouterInstanceReadDetailed**](RequestrouterInstanceReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UndeployHyperCacheInstance

> UndeployHyperCacheInstance(ctx, nodeId).Execute()

HyperCache Instance.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    nodeId := int32(56) // int32 | Uniquely identifies a HyperCache instance. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UndeployHyperCacheInstance(context.Background(), nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UndeployHyperCacheInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **int32** | Uniquely identifies a HyperCache instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUndeployHyperCacheInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UndeployRequestRouterInstance

> UndeployRequestRouterInstance(ctx, nodeId).Execute()

RequestRouter Instance.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    nodeId := int32(56) // int32 | Uniquely identifies a RequestRouter instance. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UndeployRequestRouterInstance(context.Background(), nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UndeployRequestRouterInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **int32** | Uniquely identifies a RequestRouter instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUndeployRequestRouterInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

