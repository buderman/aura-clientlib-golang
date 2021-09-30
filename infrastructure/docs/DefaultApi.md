# \DefaultApi

All URIs are relative to *https://aura.akamai.com:443/api/infrastructure/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAttributeType**](DefaultApi.md#CreateAttributeType) | **Post** /attribute-types | Attribute types
[**CreateNode**](DefaultApi.md#CreateNode) | **Post** /nodes | Nodes
[**CreateSite**](DefaultApi.md#CreateSite) | **Post** /sites | Sites
[**DeleteAttributeType**](DefaultApi.md#DeleteAttributeType) | **Delete** /attribute-types/{attributeTypeId} | AttributeType Instance.
[**DeleteNode**](DefaultApi.md#DeleteNode) | **Delete** /nodes/{nodeId} | Node Instance.
[**DeleteSite**](DefaultApi.md#DeleteSite) | **Delete** /sites/{siteId} | Site Instance.
[**ListAttributeTypes**](DefaultApi.md#ListAttributeTypes) | **Get** /attribute-types | Attribute types
[**ListNodes**](DefaultApi.md#ListNodes) | **Get** /nodes | Nodes
[**ListSiteNodes**](DefaultApi.md#ListSiteNodes) | **Get** /sites/{siteId}/nodes | nodes per site
[**ListSites**](DefaultApi.md#ListSites) | **Get** /sites | Sites
[**ReadAttributeType**](DefaultApi.md#ReadAttributeType) | **Get** /attribute-types/{attributeTypeId} | AttributeType Instance.
[**ReadAttributeTypeSettings**](DefaultApi.md#ReadAttributeTypeSettings) | **Get** /attribute-types/settings | AttributeType Settings
[**ReadNode**](DefaultApi.md#ReadNode) | **Get** /nodes/{nodeId} | Node Instance.
[**ReadNodeBootMedium**](DefaultApi.md#ReadNodeBootMedium) | **Get** /nodes/boot-medium | NodeBootMedium
[**ReadSite**](DefaultApi.md#ReadSite) | **Get** /sites/{siteId} | Site Instance.
[**UpdateAttributeType**](DefaultApi.md#UpdateAttributeType) | **Put** /attribute-types/{attributeTypeId} | AttributeType Instance.
[**UpdateAttributeTypeSettings**](DefaultApi.md#UpdateAttributeTypeSettings) | **Put** /attribute-types/settings | AttributeType Settings
[**UpdateNode**](DefaultApi.md#UpdateNode) | **Put** /nodes/{nodeId} | Node Instance.
[**UpdateSite**](DefaultApi.md#UpdateSite) | **Put** /sites/{siteId} | Site Instance.



## CreateAttributeType

> AttributeTypeRead CreateAttributeType(ctx).AttributeTypeCreate(attributeTypeCreate).Execute()

Attribute types



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
    attributeTypeCreate := *openapiclient.NewAttributeTypeCreate("Name_example") // AttributeTypeCreate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateAttributeType(context.Background()).AttributeTypeCreate(attributeTypeCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAttributeType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAttributeType`: AttributeTypeRead
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAttributeType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAttributeTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributeTypeCreate** | [**AttributeTypeCreate**](AttributeTypeCreate.md) |  | 

### Return type

[**AttributeTypeRead**](AttributeTypeRead.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNode

> NodeReadDetailed CreateNode(ctx).NodeCreate(nodeCreate).Execute()

Nodes



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
    nodeCreate := *openapiclient.NewNodeCreate([]openapiclient.InterfaceType{openapiclient.interfaceType{BondInterfaceType: openapiclient.NewBondInterfaceType("Name_example", "Type_example")}}, "Hostname_example", int32(123), []string{"DnsServers_example"}) // NodeCreate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateNode(context.Background()).NodeCreate(nodeCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNode`: NodeReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateNode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeCreate** | [**NodeCreate**](NodeCreate.md) |  | 

### Return type

[**NodeReadDetailed**](NodeReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSite

> SiteReadDetailed CreateSite(ctx).SiteCreate(siteCreate).Execute()

Sites



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
    siteCreate := *openapiclient.NewSiteCreate("Name_example", "AbbreviatedName_example") // SiteCreate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSite(context.Background()).SiteCreate(siteCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSite`: SiteReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siteCreate** | [**SiteCreate**](SiteCreate.md) |  | 

### Return type

[**SiteReadDetailed**](SiteReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAttributeType

> DeleteAttributeType(ctx, attributeTypeId).Execute()

AttributeType Instance.



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
    attributeTypeId := int32(56) // int32 | Uniquely identifies an attribute type. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteAttributeType(context.Background(), attributeTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAttributeType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeTypeId** | **int32** | Uniquely identifies an attribute type.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttributeTypeRequest struct via the builder pattern


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


## DeleteNode

> DeleteNode(ctx, nodeId).Execute()

Node Instance.



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
    nodeId := int32(56) // int32 | Uniquely identifies a node. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteNode(context.Background(), nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **int32** | Uniquely identifies a node.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodeRequest struct via the builder pattern


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


## DeleteSite

> DeleteSite(ctx, siteId).Execute()

Site Instance.



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
    siteId := int32(56) // int32 | Uniquely identifies a site. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSite(context.Background(), siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int32** | Uniquely identifies a site.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteRequest struct via the builder pattern


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


## ListAttributeTypes

> AttributeTypeCollection ListAttributeTypes(ctx).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()

Attribute types



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
    resp, r, err := api_client.DefaultApi.ListAttributeTypes(context.Background()).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAttributeTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAttributeTypes`: AttributeTypeCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAttributeTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAttributeTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **string** | The view parameter may be used to request that either the summary or detailed view be returned. | 
 **filter** | **string** | The filter parameter may be used to filter the returned collection. | 
 **sortBy** | **string** | The sortBy parameter may be used to specify the sort order of the returned collection. | 
 **pageNumber** | **int32** | The pageNumber parameter may be used to specify an offset into the results.  Useful to use in conjunction with pageSize.  Page number offset is 1-based. | 
 **pageSize** | **int32** | The pageSize parameter may be used to specify the page size of the returned results.  Useful to use in conjunction with pageNumber. | 

### Return type

[**AttributeTypeCollection**](AttributeTypeCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNodes

> NodeCollection ListNodes(ctx).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()

Nodes



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
    resp, r, err := api_client.DefaultApi.ListNodes(context.Background()).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNodes`: NodeCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **string** | The view parameter may be used to request that either the summary or detailed view be returned. | 
 **filter** | **string** | The filter parameter may be used to filter the returned collection. | 
 **sortBy** | **string** | The sortBy parameter may be used to specify the sort order of the returned collection. | 
 **pageNumber** | **int32** | The pageNumber parameter may be used to specify an offset into the results.  Useful to use in conjunction with pageSize.  Page number offset is 1-based. | 
 **pageSize** | **int32** | The pageSize parameter may be used to specify the page size of the returned results.  Useful to use in conjunction with pageNumber. | 

### Return type

[**NodeCollection**](NodeCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteNodes

> NodeCollection ListSiteNodes(ctx, siteId).Execute()

nodes per site



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
    siteId := int32(56) // int32 | Uniquely identifies a site. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSiteNodes(context.Background(), siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSiteNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSiteNodes`: NodeCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSiteNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int32** | Uniquely identifies a site.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NodeCollection**](NodeCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSites

> SiteCollection ListSites(ctx).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()

Sites



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
    resp, r, err := api_client.DefaultApi.ListSites(context.Background()).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSites`: SiteCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **string** | The view parameter may be used to request that either the summary or detailed view be returned. | 
 **filter** | **string** | The filter parameter may be used to filter the returned collection. | 
 **sortBy** | **string** | The sortBy parameter may be used to specify the sort order of the returned collection. | 
 **pageNumber** | **int32** | The pageNumber parameter may be used to specify an offset into the results.  Useful to use in conjunction with pageSize.  Page number offset is 1-based. | 
 **pageSize** | **int32** | The pageSize parameter may be used to specify the page size of the returned results.  Useful to use in conjunction with pageNumber. | 

### Return type

[**SiteCollection**](SiteCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAttributeType

> AttributeTypeRead ReadAttributeType(ctx, attributeTypeId).Execute()

AttributeType Instance.



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
    attributeTypeId := int32(56) // int32 | Uniquely identifies an attribute type. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ReadAttributeType(context.Background(), attributeTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReadAttributeType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAttributeType`: AttributeTypeRead
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReadAttributeType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeTypeId** | **int32** | Uniquely identifies an attribute type.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAttributeTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AttributeTypeRead**](AttributeTypeRead.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAttributeTypeSettings

> AttributeTypeSettings ReadAttributeTypeSettings(ctx).Execute()

AttributeType Settings



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
    resp, r, err := api_client.DefaultApi.ReadAttributeTypeSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReadAttributeTypeSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAttributeTypeSettings`: AttributeTypeSettings
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReadAttributeTypeSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReadAttributeTypeSettingsRequest struct via the builder pattern


### Return type

[**AttributeTypeSettings**](AttributeTypeSettings.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNode

> NodeReadDetailed ReadNode(ctx, nodeId).Execute()

Node Instance.



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
    nodeId := int32(56) // int32 | Uniquely identifies a node. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ReadNode(context.Background(), nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReadNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadNode`: NodeReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReadNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **int32** | Uniquely identifies a node.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NodeReadDetailed**](NodeReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNodeBootMedium

> ReadNodeBootMedium(ctx).Format(format).Execute()

NodeBootMedium



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
    format := "format_example" // string | `usb` or `iso`. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ReadNodeBootMedium(context.Background()).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReadNodeBootMedium``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadNodeBootMediumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | &#x60;usb&#x60; or &#x60;iso&#x60;.  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadSite

> SiteReadDetailed ReadSite(ctx, siteId).Execute()

Site Instance.



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
    siteId := int32(56) // int32 | Uniquely identifies a site. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ReadSite(context.Background(), siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReadSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadSite`: SiteReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReadSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int32** | Uniquely identifies a site.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SiteReadDetailed**](SiteReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAttributeType

> AttributeTypeRead UpdateAttributeType(ctx, attributeTypeId).AttributeTypeUpdate(attributeTypeUpdate).Execute()

AttributeType Instance.



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
    attributeTypeId := int32(56) // int32 | Uniquely identifies an attribute type. 
    attributeTypeUpdate := *openapiclient.NewAttributeTypeUpdate(int32(123), "Name_example") // AttributeTypeUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateAttributeType(context.Background(), attributeTypeId).AttributeTypeUpdate(attributeTypeUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAttributeType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAttributeType`: AttributeTypeRead
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateAttributeType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeTypeId** | **int32** | Uniquely identifies an attribute type.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAttributeTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attributeTypeUpdate** | [**AttributeTypeUpdate**](AttributeTypeUpdate.md) |  | 

### Return type

[**AttributeTypeRead**](AttributeTypeRead.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAttributeTypeSettings

> AttributeTypeSettings UpdateAttributeTypeSettings(ctx).AttributeTypeSettings(attributeTypeSettings).Execute()

AttributeType Settings



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
    attributeTypeSettings := *openapiclient.NewAttributeTypeSettings() // AttributeTypeSettings | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateAttributeTypeSettings(context.Background()).AttributeTypeSettings(attributeTypeSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAttributeTypeSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAttributeTypeSettings`: AttributeTypeSettings
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateAttributeTypeSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAttributeTypeSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributeTypeSettings** | [**AttributeTypeSettings**](AttributeTypeSettings.md) |  | 

### Return type

[**AttributeTypeSettings**](AttributeTypeSettings.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNode

> NodeReadDetailed UpdateNode(ctx, nodeId).NodeUpdate(nodeUpdate).Execute()

Node Instance.



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
    nodeId := int32(56) // int32 | Uniquely identifies a node. 
    nodeUpdate := *openapiclient.NewNodeUpdate(openapiclient.bootStateType("BOOT"), []openapiclient.InterfaceType{openapiclient.interfaceType{BondInterfaceType: openapiclient.NewBondInterfaceType("Name_example", "Type_example")}}, "Hostname_example", int32(123), int32(123), openapiclient.administrativeStateType("ACTIVE"), []string{"DnsServers_example"}) // NodeUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateNode(context.Background(), nodeId).NodeUpdate(nodeUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNode`: NodeReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **int32** | Uniquely identifies a node.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodeUpdate** | [**NodeUpdate**](NodeUpdate.md) |  | 

### Return type

[**NodeReadDetailed**](NodeReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSite

> SiteReadDetailed UpdateSite(ctx, siteId).SiteUpdate(siteUpdate).Execute()

Site Instance.



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
    siteId := int32(56) // int32 | Uniquely identifies a site. 
    siteUpdate := *openapiclient.NewSiteUpdate(int32(123), "Name_example", "AbbreviatedName_example") // SiteUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSite(context.Background(), siteId).SiteUpdate(siteUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSite`: SiteReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int32** | Uniquely identifies a site.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteUpdate** | [**SiteUpdate**](SiteUpdate.md) |  | 

### Return type

[**SiteReadDetailed**](SiteReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

