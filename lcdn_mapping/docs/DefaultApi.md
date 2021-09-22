# \DefaultApi

All URIs are relative to *https://aura.akamai.com:443/api/lcdn-mapping/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessMap**](DefaultApi.md#CreateAccessMap) | **Post** /access-maps | AccessMap collection
[**CreateBgpConfig**](DefaultApi.md#CreateBgpConfig) | **Post** /bgp-configs | BgpConfig collection
[**CreateGeoIpConfig**](DefaultApi.md#CreateGeoIpConfig) | **Post** /geoip-configs | GeoIpConfig collection
[**CreateSiteMap**](DefaultApi.md#CreateSiteMap) | **Post** /site-maps | SiteMap collection
[**DeleteAccessMap**](DefaultApi.md#DeleteAccessMap) | **Delete** /access-maps/{mapId} | AccessMap Instance.
[**DeleteBgpConfig**](DefaultApi.md#DeleteBgpConfig) | **Delete** /bgp-configs/{bgpConfigId} | BgpConfig Instance.
[**DeleteGeoIpConfig**](DefaultApi.md#DeleteGeoIpConfig) | **Delete** /geoip-configs/{geoIpConfigId} | GeoIpConfig Instance.
[**DeleteSiteMap**](DefaultApi.md#DeleteSiteMap) | **Delete** /site-maps/{mapId} | SiteMap Instance.
[**ListAccessMaps**](DefaultApi.md#ListAccessMaps) | **Get** /access-maps | AccessMap collection
[**ListBgpConfigs**](DefaultApi.md#ListBgpConfigs) | **Get** /bgp-configs | BgpConfig collection
[**ListGeoIpConfigs**](DefaultApi.md#ListGeoIpConfigs) | **Get** /geoip-configs | GeoIpConfig collection
[**ListSiteMaps**](DefaultApi.md#ListSiteMaps) | **Get** /site-maps | SiteMap collection
[**ReadAccessMap**](DefaultApi.md#ReadAccessMap) | **Get** /access-maps/{mapId} | AccessMap Instance.
[**ReadAccessMapContent**](DefaultApi.md#ReadAccessMapContent) | **Get** /access-maps/{mapId}/map | AccessMap Instance Map Content.
[**ReadBgpConfig**](DefaultApi.md#ReadBgpConfig) | **Get** /bgp-configs/{bgpConfigId} | BgpConfig Instance.
[**ReadBgpConfigContent**](DefaultApi.md#ReadBgpConfigContent) | **Get** /bgp-configs/{bgpConfigId}/bgp-config | BgpConfig Instance BGP configuration.
[**ReadGeoIpConfig**](DefaultApi.md#ReadGeoIpConfig) | **Get** /geoip-configs/{geoIpConfigId} | GeoIpConfig Instance.
[**ReadGeoIpConfigContent**](DefaultApi.md#ReadGeoIpConfigContent) | **Get** /geoip-configs/{geoIpConfigId}/geoip-config | GeoIpConfig Instance GeoIP configuration.
[**ReadSiteMap**](DefaultApi.md#ReadSiteMap) | **Get** /site-maps/{mapId} | SiteMap Instance.
[**ReadSiteMapContent**](DefaultApi.md#ReadSiteMapContent) | **Get** /site-maps/{mapId}/map | SiteMap Instance Map Content.
[**UpdateAccessMap**](DefaultApi.md#UpdateAccessMap) | **Put** /access-maps/{mapId} | AccessMap Instance.
[**UpdateAccessMapContent**](DefaultApi.md#UpdateAccessMapContent) | **Put** /access-maps/{mapId}/map | AccessMap Instance Map Content.
[**UpdateBgpConfig**](DefaultApi.md#UpdateBgpConfig) | **Put** /bgp-configs/{bgpConfigId} | BgpConfig Instance.
[**UpdateBgpConfigContent**](DefaultApi.md#UpdateBgpConfigContent) | **Put** /bgp-configs/{bgpConfigId}/bgp-config | BgpConfig Instance BGP configuration.
[**UpdateGeoIpConfig**](DefaultApi.md#UpdateGeoIpConfig) | **Put** /geoip-configs/{geoIpConfigId} | GeoIpConfig Instance.
[**UpdateGeoIpConfigContent**](DefaultApi.md#UpdateGeoIpConfigContent) | **Put** /geoip-configs/{geoIpConfigId}/geoip-config | GeoIpConfig Instance GeoIP configuration.
[**UpdateSiteMap**](DefaultApi.md#UpdateSiteMap) | **Put** /site-maps/{mapId} | SiteMap Instance.
[**UpdateSiteMapContent**](DefaultApi.md#UpdateSiteMapContent) | **Put** /site-maps/{mapId}/map | SiteMap Instance Map Content.
[**ValidateAccessMap**](DefaultApi.md#ValidateAccessMap) | **Post** /access-maps/validate | /validate
[**ValidateSiteMap**](DefaultApi.md#ValidateSiteMap) | **Post** /site-maps/validate | /validate



## CreateAccessMap

> AccessMapReadDetailed CreateAccessMap(ctx).Metadata(metadata).Config(config).Execute()

AccessMap collection



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
    metadata := *openapiclient.NewAccessMapCreate("Name_example") // AccessMapCreate |  (optional)
    config := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateAccessMap(context.Background()).Metadata(metadata).Config(config).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAccessMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccessMap`: AccessMapReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAccessMap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metadata** | [**AccessMapCreate**](AccessMapCreate.md) |  | 
 **config** | [**interface{}**](interface{}.md) |  | 

### Return type

[**AccessMapReadDetailed**](AccessMapReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBgpConfig

> BgpConfigReadDetailed CreateBgpConfig(ctx).Metadata(metadata).Config(config).Execute()

BgpConfig collection



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
    metadata := *openapiclient.NewBgpConfigCreate("Name_example") // BgpConfigCreate |  (optional)
    config := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateBgpConfig(context.Background()).Metadata(metadata).Config(config).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateBgpConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBgpConfig`: BgpConfigReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateBgpConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBgpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metadata** | [**BgpConfigCreate**](BgpConfigCreate.md) |  | 
 **config** | [**interface{}**](interface{}.md) |  | 

### Return type

[**BgpConfigReadDetailed**](BgpConfigReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGeoIpConfig

> GeoipConfigReadDetailed CreateGeoIpConfig(ctx).Metadata(metadata).Config(config).Execute()

GeoIpConfig collection



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
    metadata := *openapiclient.NewGeoipConfigCreate(openapiclient.typeType("IPV4"), "Name_example") // GeoipConfigCreate |  (optional)
    config := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateGeoIpConfig(context.Background()).Metadata(metadata).Config(config).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateGeoIpConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGeoIpConfig`: GeoipConfigReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateGeoIpConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGeoIpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metadata** | [**GeoipConfigCreate**](GeoipConfigCreate.md) |  | 
 **config** | [**interface{}**](interface{}.md) |  | 

### Return type

[**GeoipConfigReadDetailed**](GeoipConfigReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSiteMap

> SiteMapReadDetailed CreateSiteMap(ctx).Metadata(metadata).Config(config).Execute()

SiteMap collection



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
    metadata := *openapiclient.NewSiteMapCreate("Name_example") // SiteMapCreate |  (optional)
    config := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSiteMap(context.Background()).Metadata(metadata).Config(config).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSiteMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSiteMap`: SiteMapReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSiteMap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metadata** | [**SiteMapCreate**](SiteMapCreate.md) |  | 
 **config** | [**interface{}**](interface{}.md) |  | 

### Return type

[**SiteMapReadDetailed**](SiteMapReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessMap

> DeleteAccessMap(ctx, mapId).Execute()

AccessMap Instance.



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
    mapId := int32(56) // int32 | Uniquely identifies an access map instance. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteAccessMap(context.Background(), mapId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAccessMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mapId** | **int32** | Uniquely identifies an access map instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessMapRequest struct via the builder pattern


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


## DeleteBgpConfig

> DeleteBgpConfig(ctx, bgpConfigId).Execute()

BgpConfig Instance.



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
    bgpConfigId := int32(56) // int32 | Uniquely identifies a BGP configuration instance. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteBgpConfig(context.Background(), bgpConfigId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteBgpConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bgpConfigId** | **int32** | Uniquely identifies a BGP configuration instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBgpConfigRequest struct via the builder pattern


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


## DeleteGeoIpConfig

> DeleteGeoIpConfig(ctx, geoIpConfigId).Execute()

GeoIpConfig Instance.



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
    geoIpConfigId := int32(56) // int32 | Uniquely identifies a GeoIP configuration instance. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteGeoIpConfig(context.Background(), geoIpConfigId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteGeoIpConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geoIpConfigId** | **int32** | Uniquely identifies a GeoIP configuration instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGeoIpConfigRequest struct via the builder pattern


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


## DeleteSiteMap

> DeleteSiteMap(ctx, mapId).Execute()

SiteMap Instance.



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
    mapId := int32(56) // int32 | Uniquely identifies a site map instance. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSiteMap(context.Background(), mapId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSiteMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mapId** | **int32** | Uniquely identifies a site map instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteMapRequest struct via the builder pattern


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


## ListAccessMaps

> AccessMapCollection ListAccessMaps(ctx).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()

AccessMap collection



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
    view := "view_example" // string | The `view` parameter may be used to request that either the summary or detailed view be returned. (optional)
    filter := "filter_example" // string | The `filter` parameter may be used to filter the returned collection. (optional)
    sortBy := "sortBy_example" // string | The `sortBy` parameter may be used to specify the sort order of the returned collection. (optional)
    pageNumber := int32(56) // int32 | The `pageNumber` parameter may be used to specify an offset into the results.  Useful to use in conjunction with `pageSize`.  Page number offset is 1-based. (optional)
    pageSize := int32(56) // int32 | The `pageSize` parameter may be used to specify the returned results.  Useful to use in conjunction with `pageNumber`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListAccessMaps(context.Background()).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAccessMaps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccessMaps`: AccessMapCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAccessMaps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccessMapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **string** | The &#x60;view&#x60; parameter may be used to request that either the summary or detailed view be returned. | 
 **filter** | **string** | The &#x60;filter&#x60; parameter may be used to filter the returned collection. | 
 **sortBy** | **string** | The &#x60;sortBy&#x60; parameter may be used to specify the sort order of the returned collection. | 
 **pageNumber** | **int32** | The &#x60;pageNumber&#x60; parameter may be used to specify an offset into the results.  Useful to use in conjunction with &#x60;pageSize&#x60;.  Page number offset is 1-based. | 
 **pageSize** | **int32** | The &#x60;pageSize&#x60; parameter may be used to specify the returned results.  Useful to use in conjunction with &#x60;pageNumber&#x60;. | 

### Return type

[**AccessMapCollection**](AccessMapCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBgpConfigs

> BgpConfigCollection ListBgpConfigs(ctx).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()

BgpConfig collection



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
    view := "view_example" // string | The `view` parameter may be used to request that either the summary or detailed view be returned. (optional)
    filter := "filter_example" // string | The `filter` parameter may be used to filter the returned collection. (optional)
    sortBy := "sortBy_example" // string | The `sortBy` parameter may be used to specify the sort order of the returned collection. (optional)
    pageNumber := int32(56) // int32 | The `pageNumber` parameter may be used to specify an offset into the results.  Useful to use in conjunction with `pageSize`.  Page number offset is 1-based. (optional)
    pageSize := int32(56) // int32 | The `pageSize` parameter may be used to specify the returned results.  Useful to use in conjunction with `pageNumber`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListBgpConfigs(context.Background()).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListBgpConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBgpConfigs`: BgpConfigCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListBgpConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBgpConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **string** | The &#x60;view&#x60; parameter may be used to request that either the summary or detailed view be returned. | 
 **filter** | **string** | The &#x60;filter&#x60; parameter may be used to filter the returned collection. | 
 **sortBy** | **string** | The &#x60;sortBy&#x60; parameter may be used to specify the sort order of the returned collection. | 
 **pageNumber** | **int32** | The &#x60;pageNumber&#x60; parameter may be used to specify an offset into the results.  Useful to use in conjunction with &#x60;pageSize&#x60;.  Page number offset is 1-based. | 
 **pageSize** | **int32** | The &#x60;pageSize&#x60; parameter may be used to specify the returned results.  Useful to use in conjunction with &#x60;pageNumber&#x60;. | 

### Return type

[**BgpConfigCollection**](BgpConfigCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGeoIpConfigs

> GeoipConfigCollection ListGeoIpConfigs(ctx).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()

GeoIpConfig collection



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
    view := "view_example" // string | The `view` parameter may be used to request that either the summary or detailed view be returned. (optional)
    filter := "filter_example" // string | The `filter` parameter may be used to filter the returned collection. (optional)
    sortBy := "sortBy_example" // string | The `sortBy` parameter may be used to specify the sort order of the returned collection. (optional)
    pageNumber := int32(56) // int32 | The `pageNumber` parameter may be used to specify an offset into the results.  Useful to use in conjunction with `pageSize`.  Page number offset is 1-based. (optional)
    pageSize := int32(56) // int32 | The `pageSize` parameter may be used to specify the returned results.  Useful to use in conjunction with `pageNumber`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListGeoIpConfigs(context.Background()).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListGeoIpConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGeoIpConfigs`: GeoipConfigCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListGeoIpConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGeoIpConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **string** | The &#x60;view&#x60; parameter may be used to request that either the summary or detailed view be returned. | 
 **filter** | **string** | The &#x60;filter&#x60; parameter may be used to filter the returned collection. | 
 **sortBy** | **string** | The &#x60;sortBy&#x60; parameter may be used to specify the sort order of the returned collection. | 
 **pageNumber** | **int32** | The &#x60;pageNumber&#x60; parameter may be used to specify an offset into the results.  Useful to use in conjunction with &#x60;pageSize&#x60;.  Page number offset is 1-based. | 
 **pageSize** | **int32** | The &#x60;pageSize&#x60; parameter may be used to specify the returned results.  Useful to use in conjunction with &#x60;pageNumber&#x60;. | 

### Return type

[**GeoipConfigCollection**](GeoipConfigCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteMaps

> SiteMapCollection ListSiteMaps(ctx).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()

SiteMap collection



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
    view := "view_example" // string | The `view` parameter may be used to request that either the summary or detailed view be returned. (optional)
    filter := "filter_example" // string | The `filter` parameter may be used to filter the returned collection. (optional)
    sortBy := "sortBy_example" // string | The `sortBy` parameter may be used to specify the sort order of the returned collection. (optional)
    pageNumber := int32(56) // int32 | The `pageNumber` parameter may be used to specify an offset into the results.  Useful to use in conjunction with `pageSize`.  Page number offset is 1-based. (optional)
    pageSize := int32(56) // int32 | The `pageSize` parameter may be used to specify the returned results.  Useful to use in conjunction with `pageNumber`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSiteMaps(context.Background()).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSiteMaps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSiteMaps`: SiteMapCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSiteMaps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSiteMapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **string** | The &#x60;view&#x60; parameter may be used to request that either the summary or detailed view be returned. | 
 **filter** | **string** | The &#x60;filter&#x60; parameter may be used to filter the returned collection. | 
 **sortBy** | **string** | The &#x60;sortBy&#x60; parameter may be used to specify the sort order of the returned collection. | 
 **pageNumber** | **int32** | The &#x60;pageNumber&#x60; parameter may be used to specify an offset into the results.  Useful to use in conjunction with &#x60;pageSize&#x60;.  Page number offset is 1-based. | 
 **pageSize** | **int32** | The &#x60;pageSize&#x60; parameter may be used to specify the returned results.  Useful to use in conjunction with &#x60;pageNumber&#x60;. | 

### Return type

[**SiteMapCollection**](SiteMapCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAccessMap

> AccessMapReadDetailed ReadAccessMap(ctx, mapId).Execute()

AccessMap Instance.



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
    mapId := int32(56) // int32 | Uniquely identifies an access map instance. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ReadAccessMap(context.Background(), mapId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReadAccessMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAccessMap`: AccessMapReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReadAccessMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mapId** | **int32** | Uniquely identifies an access map instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAccessMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessMapReadDetailed**](AccessMapReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAccessMapContent

> ReadAccessMapContent(ctx, mapId).Execute()

AccessMap Instance Map Content.



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
    mapId := int32(56) // int32 | Uniquely identifies an access map instance. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ReadAccessMapContent(context.Background(), mapId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReadAccessMapContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mapId** | **int32** | Uniquely identifies an access map instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAccessMapContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadBgpConfig

> BgpConfigReadDetailed ReadBgpConfig(ctx, bgpConfigId).Execute()

BgpConfig Instance.



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
    bgpConfigId := int32(56) // int32 | Uniquely identifies a BGP configuration instance. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ReadBgpConfig(context.Background(), bgpConfigId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReadBgpConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadBgpConfig`: BgpConfigReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReadBgpConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bgpConfigId** | **int32** | Uniquely identifies a BGP configuration instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadBgpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BgpConfigReadDetailed**](BgpConfigReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadBgpConfigContent

> ReadBgpConfigContent(ctx, bgpConfigId).Execute()

BgpConfig Instance BGP configuration.



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
    bgpConfigId := int32(56) // int32 | Uniquely identifies a BGP configuration instance. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ReadBgpConfigContent(context.Background(), bgpConfigId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReadBgpConfigContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bgpConfigId** | **int32** | Uniquely identifies a BGP configuration instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadBgpConfigContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadGeoIpConfig

> GeoipConfigReadDetailed ReadGeoIpConfig(ctx, geoIpConfigId).Execute()

GeoIpConfig Instance.



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
    geoIpConfigId := int32(56) // int32 | Uniquely identifies a GeoIP configuration instance. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ReadGeoIpConfig(context.Background(), geoIpConfigId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReadGeoIpConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadGeoIpConfig`: GeoipConfigReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReadGeoIpConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geoIpConfigId** | **int32** | Uniquely identifies a GeoIP configuration instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadGeoIpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GeoipConfigReadDetailed**](GeoipConfigReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadGeoIpConfigContent

> ReadGeoIpConfigContent(ctx, geoIpConfigId).Execute()

GeoIpConfig Instance GeoIP configuration.



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
    geoIpConfigId := int32(56) // int32 | Uniquely identifies a GeoIP configuration instance. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ReadGeoIpConfigContent(context.Background(), geoIpConfigId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReadGeoIpConfigContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geoIpConfigId** | **int32** | Uniquely identifies a GeoIP configuration instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadGeoIpConfigContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadSiteMap

> SiteMapReadDetailed ReadSiteMap(ctx, mapId).Execute()

SiteMap Instance.



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
    mapId := int32(56) // int32 | Uniquely identifies a site map instance. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ReadSiteMap(context.Background(), mapId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReadSiteMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadSiteMap`: SiteMapReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReadSiteMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mapId** | **int32** | Uniquely identifies a site map instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadSiteMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SiteMapReadDetailed**](SiteMapReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadSiteMapContent

> ReadSiteMapContent(ctx, mapId).Execute()

SiteMap Instance Map Content.



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
    mapId := int32(56) // int32 | Uniquely identifies a site map instance. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ReadSiteMapContent(context.Background(), mapId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReadSiteMapContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mapId** | **int32** | Uniquely identifies a site map instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadSiteMapContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessMap

> AccessMapReadDetailed UpdateAccessMap(ctx, mapId).AccessMapUpdate(accessMapUpdate).Execute()

AccessMap Instance.



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
    mapId := int32(56) // int32 | Uniquely identifies an access map instance. 
    accessMapUpdate := *openapiclient.NewAccessMapUpdate(int32(123), "Name_example") // AccessMapUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateAccessMap(context.Background(), mapId).AccessMapUpdate(accessMapUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAccessMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccessMap`: AccessMapReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateAccessMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mapId** | **int32** | Uniquely identifies an access map instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessMapUpdate** | [**AccessMapUpdate**](AccessMapUpdate.md) |  | 

### Return type

[**AccessMapReadDetailed**](AccessMapReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessMapContent

> UpdateAccessMapContent(ctx, mapId).Body(body).Execute()

AccessMap Instance Map Content.



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
    mapId := int32(56) // int32 | Uniquely identifies an access map instance. 
    body := "{"$ref":"examples/site-map.read.json"}" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateAccessMapContent(context.Background(), mapId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAccessMapContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mapId** | **int32** | Uniquely identifies an access map instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessMapContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBgpConfig

> BgpConfigReadDetailed UpdateBgpConfig(ctx, bgpConfigId).BgpConfigUpdate(bgpConfigUpdate).Execute()

BgpConfig Instance.



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
    bgpConfigId := int32(56) // int32 | Uniquely identifies a BGP configuration instance. 
    bgpConfigUpdate := *openapiclient.NewBgpConfigUpdate(int32(123), "Name_example") // BgpConfigUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateBgpConfig(context.Background(), bgpConfigId).BgpConfigUpdate(bgpConfigUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateBgpConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBgpConfig`: BgpConfigReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateBgpConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bgpConfigId** | **int32** | Uniquely identifies a BGP configuration instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBgpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bgpConfigUpdate** | [**BgpConfigUpdate**](BgpConfigUpdate.md) |  | 

### Return type

[**BgpConfigReadDetailed**](BgpConfigReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBgpConfigContent

> UpdateBgpConfigContent(ctx, bgpConfigId).Body(body).Execute()

BgpConfig Instance BGP configuration.



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
    bgpConfigId := int32(56) // int32 | Uniquely identifies a BGP configuration instance. 
    body := "{"$ref":"examples/bgp-config.config"}" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateBgpConfigContent(context.Background(), bgpConfigId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateBgpConfigContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bgpConfigId** | **int32** | Uniquely identifies a BGP configuration instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBgpConfigContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGeoIpConfig

> GeoipConfigReadDetailed UpdateGeoIpConfig(ctx, geoIpConfigId).GeoipConfigUpdate(geoipConfigUpdate).Execute()

GeoIpConfig Instance.



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
    geoIpConfigId := int32(56) // int32 | Uniquely identifies a GeoIP configuration instance. 
    geoipConfigUpdate := *openapiclient.NewGeoipConfigUpdate(int32(123), openapiclient.typeType("IPV4"), "Name_example") // GeoipConfigUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateGeoIpConfig(context.Background(), geoIpConfigId).GeoipConfigUpdate(geoipConfigUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateGeoIpConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGeoIpConfig`: GeoipConfigReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateGeoIpConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geoIpConfigId** | **int32** | Uniquely identifies a GeoIP configuration instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGeoIpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **geoipConfigUpdate** | [**GeoipConfigUpdate**](GeoipConfigUpdate.md) |  | 

### Return type

[**GeoipConfigReadDetailed**](GeoipConfigReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGeoIpConfigContent

> UpdateGeoIpConfigContent(ctx, geoIpConfigId).Body(body).Execute()

GeoIpConfig Instance GeoIP configuration.



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
    geoIpConfigId := int32(56) // int32 | Uniquely identifies a GeoIP configuration instance. 
    body := "{"$ref":"examples/geoip-config.config"}" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateGeoIpConfigContent(context.Background(), geoIpConfigId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateGeoIpConfigContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geoIpConfigId** | **int32** | Uniquely identifies a GeoIP configuration instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGeoIpConfigContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: text/csv
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteMap

> SiteMapReadDetailed UpdateSiteMap(ctx, mapId).SiteMapUpdate(siteMapUpdate).Execute()

SiteMap Instance.



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
    mapId := int32(56) // int32 | Uniquely identifies a site map instance. 
    siteMapUpdate := *openapiclient.NewSiteMapUpdate(int32(123), "Name_example") // SiteMapUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSiteMap(context.Background(), mapId).SiteMapUpdate(siteMapUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSiteMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSiteMap`: SiteMapReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSiteMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mapId** | **int32** | Uniquely identifies a site map instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteMapUpdate** | [**SiteMapUpdate**](SiteMapUpdate.md) |  | 

### Return type

[**SiteMapReadDetailed**](SiteMapReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteMapContent

> UpdateSiteMapContent(ctx, mapId).Body(body).Execute()

SiteMap Instance Map Content.



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
    mapId := int32(56) // int32 | Uniquely identifies a site map instance. 
    body := "{"$ref":"examples/site-map.read.json"}" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSiteMapContent(context.Background(), mapId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSiteMapContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mapId** | **int32** | Uniquely identifies a site map instance.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteMapContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateAccessMap

> ValidateAccessMap(ctx).Metadata(metadata).Config(config).Execute()

/validate



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
    metadata := *openapiclient.NewAccessMapCreate("Name_example") // AccessMapCreate |  (optional)
    config := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ValidateAccessMap(context.Background()).Metadata(metadata).Config(config).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ValidateAccessMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAccessMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metadata** | [**AccessMapCreate**](AccessMapCreate.md) |  | 
 **config** | [**interface{}**](interface{}.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateSiteMap

> ValidateSiteMap(ctx).Metadata(metadata).Config(config).Execute()

/validate



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
    metadata := *openapiclient.NewSiteMapCreate("Name_example") // SiteMapCreate |  (optional)
    config := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ValidateSiteMap(context.Background()).Metadata(metadata).Config(config).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ValidateSiteMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateSiteMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metadata** | [**SiteMapCreate**](SiteMapCreate.md) |  | 
 **config** | [**interface{}**](interface{}.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

