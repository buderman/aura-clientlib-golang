# \DefaultApi

All URIs are relative to *https://aura.akamai.com:443/api/content-delivery/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCdnPrefix**](DefaultApi.md#CreateCdnPrefix) | **Post** /content-providers/{contentProviderId}/cdn-prefixes | cdn prefixes
[**CreateOrigin**](DefaultApi.md#CreateOrigin) | **Post** /content-providers/{contentProviderId}/origins | origins
[**CreateSharedSecretSet**](DefaultApi.md#CreateSharedSecretSet) | **Post** /content-providers/{contentProviderId}/shared-secret-sets | shared secret sets
[**CreateTlsDeliveryProfile**](DefaultApi.md#CreateTlsDeliveryProfile) | **Post** /content-providers/{contentProviderId}/tls-delivery-profiles | tls delivery profiles
[**CreateTlsIngestProfiles**](DefaultApi.md#CreateTlsIngestProfiles) | **Post** /content-providers/{contentProviderId}/tls-ingest-profiles | tls ingest profiles
[**DeleteCdnPrefix**](DefaultApi.md#DeleteCdnPrefix) | **Delete** /content-providers/{contentProviderId}/cdn-prefixes/{cdnPrefixId} | cdn prefix
[**DeleteOrigin**](DefaultApi.md#DeleteOrigin) | **Delete** /content-providers/{contentProviderId}/origins/{originId} | origin
[**DeleteSharedSecretSet**](DefaultApi.md#DeleteSharedSecretSet) | **Delete** /content-providers/{contentProviderId}/shared-secret-sets/{sharedSecretSetId} | shared secret set
[**DeleteTlsDeliveryProfile**](DefaultApi.md#DeleteTlsDeliveryProfile) | **Delete** /content-providers/{contentProviderId}/tls-delivery-profiles/{tlsDeliveryProfileId} | tls delivery profile
[**DeleteTlsIngestProfile**](DefaultApi.md#DeleteTlsIngestProfile) | **Delete** /content-providers/{contentProviderId}/tls-ingest-profiles/{tlsIngestProfileId} | tls ingest profile
[**ListCdnPrefixes**](DefaultApi.md#ListCdnPrefixes) | **Get** /content-providers/{contentProviderId}/cdn-prefixes | cdn prefixes
[**ListContentProviders**](DefaultApi.md#ListContentProviders) | **Get** /content-providers | content providers
[**ListOrigins**](DefaultApi.md#ListOrigins) | **Get** /content-providers/{contentProviderId}/origins | origins
[**ListSharedSecretSets**](DefaultApi.md#ListSharedSecretSets) | **Get** /content-providers/{contentProviderId}/shared-secret-sets | shared secret sets
[**ListTlsDeliveryProfileCdnPrefixes**](DefaultApi.md#ListTlsDeliveryProfileCdnPrefixes) | **Get** /content-providers/{contentProviderId}/tls-delivery-profiles/{tlsDeliveryProfileId}/cdn-prefixes | cdn prefixes per tls
[**ListTlsDeliveryProfiles**](DefaultApi.md#ListTlsDeliveryProfiles) | **Get** /content-providers/{contentProviderId}/tls-delivery-profiles | tls delivery profiles
[**ListTlsIngestProfileOrigins**](DefaultApi.md#ListTlsIngestProfileOrigins) | **Get** /content-providers/{contentProviderId}/tls-ingest-profiles/{tlsIngestProfileId}/origins | origin per tls ingest profile
[**ListTlsIngestProfiles**](DefaultApi.md#ListTlsIngestProfiles) | **Get** /content-providers/{contentProviderId}/tls-ingest-profiles | tls ingest profiles
[**ReadCdnPrefix**](DefaultApi.md#ReadCdnPrefix) | **Get** /content-providers/{contentProviderId}/cdn-prefixes/{cdnPrefixId} | cdn prefix
[**ReadCdnPrefixRules**](DefaultApi.md#ReadCdnPrefixRules) | **Get** /content-providers/{contentProviderId}/cdn-prefixes/{cdnPrefixId}/rules | rules
[**ReadContentProvider**](DefaultApi.md#ReadContentProvider) | **Get** /content-providers/{contentProviderId} | content provider
[**ReadOrigin**](DefaultApi.md#ReadOrigin) | **Get** /content-providers/{contentProviderId}/origins/{originId} | origin
[**ReadSharedSecretSet**](DefaultApi.md#ReadSharedSecretSet) | **Get** /content-providers/{contentProviderId}/shared-secret-sets/{sharedSecretSetId} | shared secret set
[**ReadSharedSecretSetSecrets**](DefaultApi.md#ReadSharedSecretSetSecrets) | **Get** /content-providers/{contentProviderId}/shared-secret-sets/{sharedSecretSetId}/secrets | secrets
[**ReadTlsDeliveryProfile**](DefaultApi.md#ReadTlsDeliveryProfile) | **Get** /content-providers/{contentProviderId}/tls-delivery-profiles/{tlsDeliveryProfileId} | tls delivery profile
[**ReadTlsIngestProfile**](DefaultApi.md#ReadTlsIngestProfile) | **Get** /content-providers/{contentProviderId}/tls-ingest-profiles/{tlsIngestProfileId} | tls ingest profile
[**UpdateCdnPrefix**](DefaultApi.md#UpdateCdnPrefix) | **Put** /content-providers/{contentProviderId}/cdn-prefixes/{cdnPrefixId} | cdn prefix
[**UpdateCdnPrefixRules**](DefaultApi.md#UpdateCdnPrefixRules) | **Put** /content-providers/{contentProviderId}/cdn-prefixes/{cdnPrefixId}/rules | rules
[**UpdateOrigin**](DefaultApi.md#UpdateOrigin) | **Put** /content-providers/{contentProviderId}/origins/{originId} | origin
[**UpdateSharedSecretSetSecrets**](DefaultApi.md#UpdateSharedSecretSetSecrets) | **Put** /content-providers/{contentProviderId}/shared-secret-sets/{sharedSecretSetId}/secrets | secrets
[**UpdateTlsDeliveryProfile**](DefaultApi.md#UpdateTlsDeliveryProfile) | **Put** /content-providers/{contentProviderId}/tls-delivery-profiles/{tlsDeliveryProfileId} | tls delivery profile
[**UpdateTlsIngestProfile**](DefaultApi.md#UpdateTlsIngestProfile) | **Put** /content-providers/{contentProviderId}/tls-ingest-profiles/{tlsIngestProfileId} | tls ingest profile



## CreateCdnPrefix

> CdnPrefixReadDetailed CreateCdnPrefix(ctx, contentProviderId).CdnPrefixCreate(cdnPrefixCreate).Execute()

cdn prefixes



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    cdnPrefixCreate := *openapiclient.NewCdnPrefixCreate(int32(123), "Prefix_example") // CdnPrefixCreate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateCdnPrefix(context.Background(), contentProviderId).CdnPrefixCreate(cdnPrefixCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCdnPrefix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCdnPrefix`: CdnPrefixReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCdnPrefix`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCdnPrefixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cdnPrefixCreate** | [**CdnPrefixCreate**](CdnPrefixCreate.md) |  | 

### Return type

[**CdnPrefixReadDetailed**](CdnPrefixReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrigin

> OriginReadDetailed CreateOrigin(ctx, contentProviderId).OriginCreate(originCreate).Execute()

origins



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    originCreate := *openapiclient.NewOriginCreate("Hostname_example", int32(123)) // OriginCreate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateOrigin(context.Background(), contentProviderId).OriginCreate(originCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrigin`: OriginReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateOrigin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOriginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **originCreate** | [**OriginCreate**](OriginCreate.md) |  | 

### Return type

[**OriginReadDetailed**](OriginReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSharedSecretSet

> SharedSecretSetReadDetailed CreateSharedSecretSet(ctx, contentProviderId).SharedSecretSetCreate(sharedSecretSetCreate).Execute()

shared secret sets



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    sharedSecretSetCreate := *openapiclient.NewSharedSecretSetCreate(int32(123), "Name_example") // SharedSecretSetCreate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSharedSecretSet(context.Background(), contentProviderId).SharedSecretSetCreate(sharedSecretSetCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSharedSecretSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSharedSecretSet`: SharedSecretSetReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSharedSecretSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSharedSecretSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sharedSecretSetCreate** | [**SharedSecretSetCreate**](SharedSecretSetCreate.md) |  | 

### Return type

[**SharedSecretSetReadDetailed**](SharedSecretSetReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTlsDeliveryProfile

> TlsDeliveryProfileReadDetailed CreateTlsDeliveryProfile(ctx, contentProviderId).TlsDeliveryProfileCreate(tlsDeliveryProfileCreate).Execute()

tls delivery profiles



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    tlsDeliveryProfileCreate := *openapiclient.NewTlsDeliveryProfileCreate("Name_example", int32(123)) // TlsDeliveryProfileCreate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateTlsDeliveryProfile(context.Background(), contentProviderId).TlsDeliveryProfileCreate(tlsDeliveryProfileCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTlsDeliveryProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTlsDeliveryProfile`: TlsDeliveryProfileReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTlsDeliveryProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTlsDeliveryProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tlsDeliveryProfileCreate** | [**TlsDeliveryProfileCreate**](TlsDeliveryProfileCreate.md) |  | 

### Return type

[**TlsDeliveryProfileReadDetailed**](TlsDeliveryProfileReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTlsIngestProfiles

> TlsIngestProfileReadDetailed CreateTlsIngestProfiles(ctx, contentProviderId).TlsIngestProfileCreate(tlsIngestProfileCreate).Execute()

tls ingest profiles



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    tlsIngestProfileCreate := *openapiclient.NewTlsIngestProfileCreate("Name_example", int32(123)) // TlsIngestProfileCreate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateTlsIngestProfiles(context.Background(), contentProviderId).TlsIngestProfileCreate(tlsIngestProfileCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTlsIngestProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTlsIngestProfiles`: TlsIngestProfileReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTlsIngestProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTlsIngestProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tlsIngestProfileCreate** | [**TlsIngestProfileCreate**](TlsIngestProfileCreate.md) |  | 

### Return type

[**TlsIngestProfileReadDetailed**](TlsIngestProfileReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCdnPrefix

> DeleteCdnPrefix(ctx, contentProviderId, cdnPrefixId).Execute()

cdn prefix



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    cdnPrefixId := int32(56) // int32 | Unique identifier for each CDN prefix. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteCdnPrefix(context.Background(), contentProviderId, cdnPrefixId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteCdnPrefix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 
**cdnPrefixId** | **int32** | Unique identifier for each CDN prefix.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCdnPrefixRequest struct via the builder pattern


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


## DeleteOrigin

> DeleteOrigin(ctx, contentProviderId, originId).Execute()

origin



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    originId := int32(56) // int32 | Unique identifier for each origin. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteOrigin(context.Background(), contentProviderId, originId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 
**originId** | **int32** | Unique identifier for each origin.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOriginRequest struct via the builder pattern


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


## DeleteSharedSecretSet

> DeleteSharedSecretSet(ctx, contentProviderId, sharedSecretSetId).Execute()

shared secret set



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    sharedSecretSetId := int32(56) // int32 | Unique identifier for each shared secret set. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSharedSecretSet(context.Background(), contentProviderId, sharedSecretSetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSharedSecretSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 
**sharedSecretSetId** | **int32** | Unique identifier for each shared secret set.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSharedSecretSetRequest struct via the builder pattern


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


## DeleteTlsDeliveryProfile

> DeleteTlsDeliveryProfile(ctx, contentProviderId, tlsDeliveryProfileId).Execute()

tls delivery profile



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    tlsDeliveryProfileId := int32(56) // int32 | Unique identifier for each TLS delivery profile. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteTlsDeliveryProfile(context.Background(), contentProviderId, tlsDeliveryProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTlsDeliveryProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 
**tlsDeliveryProfileId** | **int32** | Unique identifier for each TLS delivery profile.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTlsDeliveryProfileRequest struct via the builder pattern


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


## DeleteTlsIngestProfile

> DeleteTlsIngestProfile(ctx, contentProviderId, tlsIngestProfileId).Execute()

tls ingest profile



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    tlsIngestProfileId := int32(56) // int32 | Unique identifier for each TLS ingest profile. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteTlsIngestProfile(context.Background(), contentProviderId, tlsIngestProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTlsIngestProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 
**tlsIngestProfileId** | **int32** | Unique identifier for each TLS ingest profile.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTlsIngestProfileRequest struct via the builder pattern


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


## ListCdnPrefixes

> CdnPrefixCollection ListCdnPrefixes(ctx, contentProviderId).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()

cdn prefixes



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    view := "view_example" // string | Results display either `summary` or `detailed` objects.  (optional)
    filter := "filter_example" // string | Specifies a query to filter down the set of results. The accompanying example specifies more than one query, where `;` and `,` delimiters clarify AND and OR logic. (Optional parentheses clarify logical groupings.) The `==` and `!=` operators test equivalency. The `=lt=`, `=le=`, `=gt=`, and `=ge=` operators compare numbers. The `=contains=` and `=excludes=` operators match elements within an array.  (optional)
    sortBy := "sortBy_example" // string | Specifies an object field name to sort results on. Provide a comma-separated list of field names for additional sort criteria. Prefix field names with a dash for descending sort order. For example, `name,-tlsIngestProfileId` sorts primarily by name, then IDs in descending order.  (optional)
    pageNumber := int32(56) // int32 | Specifies an offset within the full set of results, with page numbers starting at 1.  Use this in conjunction with `pageSize`.  (optional)
    pageSize := int32(56) // int32 | Specifies the number of results on each page.  Use this in conjunction with `pageNumber`.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListCdnPrefixes(context.Background(), contentProviderId).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCdnPrefixes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCdnPrefixes`: CdnPrefixCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCdnPrefixes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCdnPrefixesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **string** | Results display either &#x60;summary&#x60; or &#x60;detailed&#x60; objects.  | 
 **filter** | **string** | Specifies a query to filter down the set of results. The accompanying example specifies more than one query, where &#x60;;&#x60; and &#x60;,&#x60; delimiters clarify AND and OR logic. (Optional parentheses clarify logical groupings.) The &#x60;&#x3D;&#x3D;&#x60; and &#x60;!&#x3D;&#x60; operators test equivalency. The &#x60;&#x3D;lt&#x3D;&#x60;, &#x60;&#x3D;le&#x3D;&#x60;, &#x60;&#x3D;gt&#x3D;&#x60;, and &#x60;&#x3D;ge&#x3D;&#x60; operators compare numbers. The &#x60;&#x3D;contains&#x3D;&#x60; and &#x60;&#x3D;excludes&#x3D;&#x60; operators match elements within an array.  | 
 **sortBy** | **string** | Specifies an object field name to sort results on. Provide a comma-separated list of field names for additional sort criteria. Prefix field names with a dash for descending sort order. For example, &#x60;name,-tlsIngestProfileId&#x60; sorts primarily by name, then IDs in descending order.  | 
 **pageNumber** | **int32** | Specifies an offset within the full set of results, with page numbers starting at 1.  Use this in conjunction with &#x60;pageSize&#x60;.  | 
 **pageSize** | **int32** | Specifies the number of results on each page.  Use this in conjunction with &#x60;pageNumber&#x60;.  | 

### Return type

[**CdnPrefixCollection**](CdnPrefixCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContentProviders

> ContentProviderCollection ListContentProviders(ctx).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()

content providers



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
    view := "view_example" // string | Results display either `summary` or `detailed` objects.  (optional)
    filter := "filter_example" // string | Specifies a query to filter down the set of results. The accompanying example specifies more than one query, where `;` and `,` delimiters clarify AND and OR logic. (Optional parentheses clarify logical groupings.) The `==` and `!=` operators test equivalency. The `=lt=`, `=le=`, `=gt=`, and `=ge=` operators compare numbers. The `=contains=` and `=excludes=` operators match elements within an array.  (optional)
    sortBy := "sortBy_example" // string | Specifies an object field name to sort results on. Provide a comma-separated list of field names for additional sort criteria. Prefix field names with a dash for descending sort order. For example, `name,-tlsIngestProfileId` sorts primarily by name, then IDs in descending order.  (optional)
    pageNumber := int32(56) // int32 | Specifies an offset within the full set of results, with page numbers starting at 1.  Use this in conjunction with `pageSize`.  (optional)
    pageSize := int32(56) // int32 | Specifies the number of results on each page.  Use this in conjunction with `pageNumber`.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListContentProviders(context.Background()).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListContentProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContentProviders`: ContentProviderCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListContentProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListContentProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **string** | Results display either &#x60;summary&#x60; or &#x60;detailed&#x60; objects.  | 
 **filter** | **string** | Specifies a query to filter down the set of results. The accompanying example specifies more than one query, where &#x60;;&#x60; and &#x60;,&#x60; delimiters clarify AND and OR logic. (Optional parentheses clarify logical groupings.) The &#x60;&#x3D;&#x3D;&#x60; and &#x60;!&#x3D;&#x60; operators test equivalency. The &#x60;&#x3D;lt&#x3D;&#x60;, &#x60;&#x3D;le&#x3D;&#x60;, &#x60;&#x3D;gt&#x3D;&#x60;, and &#x60;&#x3D;ge&#x3D;&#x60; operators compare numbers. The &#x60;&#x3D;contains&#x3D;&#x60; and &#x60;&#x3D;excludes&#x3D;&#x60; operators match elements within an array.  | 
 **sortBy** | **string** | Specifies an object field name to sort results on. Provide a comma-separated list of field names for additional sort criteria. Prefix field names with a dash for descending sort order. For example, &#x60;name,-tlsIngestProfileId&#x60; sorts primarily by name, then IDs in descending order.  | 
 **pageNumber** | **int32** | Specifies an offset within the full set of results, with page numbers starting at 1.  Use this in conjunction with &#x60;pageSize&#x60;.  | 
 **pageSize** | **int32** | Specifies the number of results on each page.  Use this in conjunction with &#x60;pageNumber&#x60;.  | 

### Return type

[**ContentProviderCollection**](ContentProviderCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrigins

> OriginCollection ListOrigins(ctx, contentProviderId).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()

origins



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    view := "view_example" // string | Results display either `summary` or `detailed` objects.  (optional)
    filter := "filter_example" // string | Specifies a query to filter down the set of results. The accompanying example specifies more than one query, where `;` and `,` delimiters clarify AND and OR logic. (Optional parentheses clarify logical groupings.) The `==` and `!=` operators test equivalency. The `=lt=`, `=le=`, `=gt=`, and `=ge=` operators compare numbers. The `=contains=` and `=excludes=` operators match elements within an array.  (optional)
    sortBy := "sortBy_example" // string | Specifies an object field name to sort results on. Provide a comma-separated list of field names for additional sort criteria. Prefix field names with a dash for descending sort order. For example, `name,-tlsIngestProfileId` sorts primarily by name, then IDs in descending order.  (optional)
    pageNumber := int32(56) // int32 | Specifies an offset within the full set of results, with page numbers starting at 1.  Use this in conjunction with `pageSize`.  (optional)
    pageSize := int32(56) // int32 | Specifies the number of results on each page.  Use this in conjunction with `pageNumber`.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListOrigins(context.Background(), contentProviderId).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListOrigins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrigins`: OriginCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListOrigins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOriginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **string** | Results display either &#x60;summary&#x60; or &#x60;detailed&#x60; objects.  | 
 **filter** | **string** | Specifies a query to filter down the set of results. The accompanying example specifies more than one query, where &#x60;;&#x60; and &#x60;,&#x60; delimiters clarify AND and OR logic. (Optional parentheses clarify logical groupings.) The &#x60;&#x3D;&#x3D;&#x60; and &#x60;!&#x3D;&#x60; operators test equivalency. The &#x60;&#x3D;lt&#x3D;&#x60;, &#x60;&#x3D;le&#x3D;&#x60;, &#x60;&#x3D;gt&#x3D;&#x60;, and &#x60;&#x3D;ge&#x3D;&#x60; operators compare numbers. The &#x60;&#x3D;contains&#x3D;&#x60; and &#x60;&#x3D;excludes&#x3D;&#x60; operators match elements within an array.  | 
 **sortBy** | **string** | Specifies an object field name to sort results on. Provide a comma-separated list of field names for additional sort criteria. Prefix field names with a dash for descending sort order. For example, &#x60;name,-tlsIngestProfileId&#x60; sorts primarily by name, then IDs in descending order.  | 
 **pageNumber** | **int32** | Specifies an offset within the full set of results, with page numbers starting at 1.  Use this in conjunction with &#x60;pageSize&#x60;.  | 
 **pageSize** | **int32** | Specifies the number of results on each page.  Use this in conjunction with &#x60;pageNumber&#x60;.  | 

### Return type

[**OriginCollection**](OriginCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSharedSecretSets

> SharedSecretSetCollection ListSharedSecretSets(ctx, contentProviderId).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()

shared secret sets



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    view := "view_example" // string | Results display either `summary` or `detailed` objects.  (optional)
    filter := "filter_example" // string | Specifies a query to filter down the set of results. The accompanying example specifies more than one query, where `;` and `,` delimiters clarify AND and OR logic. (Optional parentheses clarify logical groupings.) The `==` and `!=` operators test equivalency. The `=lt=`, `=le=`, `=gt=`, and `=ge=` operators compare numbers. The `=contains=` and `=excludes=` operators match elements within an array.  (optional)
    sortBy := "sortBy_example" // string | Specifies an object field name to sort results on. Provide a comma-separated list of field names for additional sort criteria. Prefix field names with a dash for descending sort order. For example, `name,-tlsIngestProfileId` sorts primarily by name, then IDs in descending order.  (optional)
    pageNumber := int32(56) // int32 | Specifies an offset within the full set of results, with page numbers starting at 1.  Use this in conjunction with `pageSize`.  (optional)
    pageSize := int32(56) // int32 | Specifies the number of results on each page.  Use this in conjunction with `pageNumber`.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSharedSecretSets(context.Background(), contentProviderId).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSharedSecretSets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSharedSecretSets`: SharedSecretSetCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSharedSecretSets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSharedSecretSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **string** | Results display either &#x60;summary&#x60; or &#x60;detailed&#x60; objects.  | 
 **filter** | **string** | Specifies a query to filter down the set of results. The accompanying example specifies more than one query, where &#x60;;&#x60; and &#x60;,&#x60; delimiters clarify AND and OR logic. (Optional parentheses clarify logical groupings.) The &#x60;&#x3D;&#x3D;&#x60; and &#x60;!&#x3D;&#x60; operators test equivalency. The &#x60;&#x3D;lt&#x3D;&#x60;, &#x60;&#x3D;le&#x3D;&#x60;, &#x60;&#x3D;gt&#x3D;&#x60;, and &#x60;&#x3D;ge&#x3D;&#x60; operators compare numbers. The &#x60;&#x3D;contains&#x3D;&#x60; and &#x60;&#x3D;excludes&#x3D;&#x60; operators match elements within an array.  | 
 **sortBy** | **string** | Specifies an object field name to sort results on. Provide a comma-separated list of field names for additional sort criteria. Prefix field names with a dash for descending sort order. For example, &#x60;name,-tlsIngestProfileId&#x60; sorts primarily by name, then IDs in descending order.  | 
 **pageNumber** | **int32** | Specifies an offset within the full set of results, with page numbers starting at 1.  Use this in conjunction with &#x60;pageSize&#x60;.  | 
 **pageSize** | **int32** | Specifies the number of results on each page.  Use this in conjunction with &#x60;pageNumber&#x60;.  | 

### Return type

[**SharedSecretSetCollection**](SharedSecretSetCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTlsDeliveryProfileCdnPrefixes

> CdnPrefixCollection ListTlsDeliveryProfileCdnPrefixes(ctx, contentProviderId, tlsDeliveryProfileId).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()

cdn prefixes per tls



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    tlsDeliveryProfileId := int32(56) // int32 | Unique identifier for each TLS delivery profile. 
    view := "view_example" // string | Results display either `summary` or `detailed` objects.  (optional)
    filter := "filter_example" // string | Specifies a query to filter down the set of results. The accompanying example specifies more than one query, where `;` and `,` delimiters clarify AND and OR logic. (Optional parentheses clarify logical groupings.) The `==` and `!=` operators test equivalency. The `=lt=`, `=le=`, `=gt=`, and `=ge=` operators compare numbers. The `=contains=` and `=excludes=` operators match elements within an array.  (optional)
    sortBy := "sortBy_example" // string | Specifies an object field name to sort results on. Provide a comma-separated list of field names for additional sort criteria. Prefix field names with a dash for descending sort order. For example, `name,-tlsIngestProfileId` sorts primarily by name, then IDs in descending order.  (optional)
    pageNumber := int32(56) // int32 | Specifies an offset within the full set of results, with page numbers starting at 1.  Use this in conjunction with `pageSize`.  (optional)
    pageSize := int32(56) // int32 | Specifies the number of results on each page.  Use this in conjunction with `pageNumber`.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListTlsDeliveryProfileCdnPrefixes(context.Background(), contentProviderId, tlsDeliveryProfileId).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListTlsDeliveryProfileCdnPrefixes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTlsDeliveryProfileCdnPrefixes`: CdnPrefixCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListTlsDeliveryProfileCdnPrefixes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 
**tlsDeliveryProfileId** | **int32** | Unique identifier for each TLS delivery profile.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTlsDeliveryProfileCdnPrefixesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **view** | **string** | Results display either &#x60;summary&#x60; or &#x60;detailed&#x60; objects.  | 
 **filter** | **string** | Specifies a query to filter down the set of results. The accompanying example specifies more than one query, where &#x60;;&#x60; and &#x60;,&#x60; delimiters clarify AND and OR logic. (Optional parentheses clarify logical groupings.) The &#x60;&#x3D;&#x3D;&#x60; and &#x60;!&#x3D;&#x60; operators test equivalency. The &#x60;&#x3D;lt&#x3D;&#x60;, &#x60;&#x3D;le&#x3D;&#x60;, &#x60;&#x3D;gt&#x3D;&#x60;, and &#x60;&#x3D;ge&#x3D;&#x60; operators compare numbers. The &#x60;&#x3D;contains&#x3D;&#x60; and &#x60;&#x3D;excludes&#x3D;&#x60; operators match elements within an array.  | 
 **sortBy** | **string** | Specifies an object field name to sort results on. Provide a comma-separated list of field names for additional sort criteria. Prefix field names with a dash for descending sort order. For example, &#x60;name,-tlsIngestProfileId&#x60; sorts primarily by name, then IDs in descending order.  | 
 **pageNumber** | **int32** | Specifies an offset within the full set of results, with page numbers starting at 1.  Use this in conjunction with &#x60;pageSize&#x60;.  | 
 **pageSize** | **int32** | Specifies the number of results on each page.  Use this in conjunction with &#x60;pageNumber&#x60;.  | 

### Return type

[**CdnPrefixCollection**](CdnPrefixCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTlsDeliveryProfiles

> TlsDeliveryProfileCollection ListTlsDeliveryProfiles(ctx, contentProviderId).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()

tls delivery profiles



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    view := "view_example" // string | Results display either `summary` or `detailed` objects.  (optional)
    filter := "filter_example" // string | Specifies a query to filter down the set of results. The accompanying example specifies more than one query, where `;` and `,` delimiters clarify AND and OR logic. (Optional parentheses clarify logical groupings.) The `==` and `!=` operators test equivalency. The `=lt=`, `=le=`, `=gt=`, and `=ge=` operators compare numbers. The `=contains=` and `=excludes=` operators match elements within an array.  (optional)
    sortBy := "sortBy_example" // string | Specifies an object field name to sort results on. Provide a comma-separated list of field names for additional sort criteria. Prefix field names with a dash for descending sort order. For example, `name,-tlsIngestProfileId` sorts primarily by name, then IDs in descending order.  (optional)
    pageNumber := int32(56) // int32 | Specifies an offset within the full set of results, with page numbers starting at 1.  Use this in conjunction with `pageSize`.  (optional)
    pageSize := int32(56) // int32 | Specifies the number of results on each page.  Use this in conjunction with `pageNumber`.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListTlsDeliveryProfiles(context.Background(), contentProviderId).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListTlsDeliveryProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTlsDeliveryProfiles`: TlsDeliveryProfileCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListTlsDeliveryProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTlsDeliveryProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **string** | Results display either &#x60;summary&#x60; or &#x60;detailed&#x60; objects.  | 
 **filter** | **string** | Specifies a query to filter down the set of results. The accompanying example specifies more than one query, where &#x60;;&#x60; and &#x60;,&#x60; delimiters clarify AND and OR logic. (Optional parentheses clarify logical groupings.) The &#x60;&#x3D;&#x3D;&#x60; and &#x60;!&#x3D;&#x60; operators test equivalency. The &#x60;&#x3D;lt&#x3D;&#x60;, &#x60;&#x3D;le&#x3D;&#x60;, &#x60;&#x3D;gt&#x3D;&#x60;, and &#x60;&#x3D;ge&#x3D;&#x60; operators compare numbers. The &#x60;&#x3D;contains&#x3D;&#x60; and &#x60;&#x3D;excludes&#x3D;&#x60; operators match elements within an array.  | 
 **sortBy** | **string** | Specifies an object field name to sort results on. Provide a comma-separated list of field names for additional sort criteria. Prefix field names with a dash for descending sort order. For example, &#x60;name,-tlsIngestProfileId&#x60; sorts primarily by name, then IDs in descending order.  | 
 **pageNumber** | **int32** | Specifies an offset within the full set of results, with page numbers starting at 1.  Use this in conjunction with &#x60;pageSize&#x60;.  | 
 **pageSize** | **int32** | Specifies the number of results on each page.  Use this in conjunction with &#x60;pageNumber&#x60;.  | 

### Return type

[**TlsDeliveryProfileCollection**](TlsDeliveryProfileCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTlsIngestProfileOrigins

> OriginCollection ListTlsIngestProfileOrigins(ctx, contentProviderId, tlsIngestProfileId).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()

origin per tls ingest profile



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    tlsIngestProfileId := int32(56) // int32 | Unique identifier for each TLS ingest profile. 
    view := "view_example" // string | Results display either `summary` or `detailed` objects.  (optional)
    filter := "filter_example" // string | Specifies a query to filter down the set of results. The accompanying example specifies more than one query, where `;` and `,` delimiters clarify AND and OR logic. (Optional parentheses clarify logical groupings.) The `==` and `!=` operators test equivalency. The `=lt=`, `=le=`, `=gt=`, and `=ge=` operators compare numbers. The `=contains=` and `=excludes=` operators match elements within an array.  (optional)
    sortBy := "sortBy_example" // string | Specifies an object field name to sort results on. Provide a comma-separated list of field names for additional sort criteria. Prefix field names with a dash for descending sort order. For example, `name,-tlsIngestProfileId` sorts primarily by name, then IDs in descending order.  (optional)
    pageNumber := int32(56) // int32 | Specifies an offset within the full set of results, with page numbers starting at 1.  Use this in conjunction with `pageSize`.  (optional)
    pageSize := int32(56) // int32 | Specifies the number of results on each page.  Use this in conjunction with `pageNumber`.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListTlsIngestProfileOrigins(context.Background(), contentProviderId, tlsIngestProfileId).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListTlsIngestProfileOrigins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTlsIngestProfileOrigins`: OriginCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListTlsIngestProfileOrigins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 
**tlsIngestProfileId** | **int32** | Unique identifier for each TLS ingest profile.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTlsIngestProfileOriginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **view** | **string** | Results display either &#x60;summary&#x60; or &#x60;detailed&#x60; objects.  | 
 **filter** | **string** | Specifies a query to filter down the set of results. The accompanying example specifies more than one query, where &#x60;;&#x60; and &#x60;,&#x60; delimiters clarify AND and OR logic. (Optional parentheses clarify logical groupings.) The &#x60;&#x3D;&#x3D;&#x60; and &#x60;!&#x3D;&#x60; operators test equivalency. The &#x60;&#x3D;lt&#x3D;&#x60;, &#x60;&#x3D;le&#x3D;&#x60;, &#x60;&#x3D;gt&#x3D;&#x60;, and &#x60;&#x3D;ge&#x3D;&#x60; operators compare numbers. The &#x60;&#x3D;contains&#x3D;&#x60; and &#x60;&#x3D;excludes&#x3D;&#x60; operators match elements within an array.  | 
 **sortBy** | **string** | Specifies an object field name to sort results on. Provide a comma-separated list of field names for additional sort criteria. Prefix field names with a dash for descending sort order. For example, &#x60;name,-tlsIngestProfileId&#x60; sorts primarily by name, then IDs in descending order.  | 
 **pageNumber** | **int32** | Specifies an offset within the full set of results, with page numbers starting at 1.  Use this in conjunction with &#x60;pageSize&#x60;.  | 
 **pageSize** | **int32** | Specifies the number of results on each page.  Use this in conjunction with &#x60;pageNumber&#x60;.  | 

### Return type

[**OriginCollection**](OriginCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTlsIngestProfiles

> TlsIngestProfileCollection ListTlsIngestProfiles(ctx, contentProviderId).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()

tls ingest profiles



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    view := "view_example" // string | Results display either `summary` or `detailed` objects.  (optional)
    filter := "filter_example" // string | Specifies a query to filter down the set of results. The accompanying example specifies more than one query, where `;` and `,` delimiters clarify AND and OR logic. (Optional parentheses clarify logical groupings.) The `==` and `!=` operators test equivalency. The `=lt=`, `=le=`, `=gt=`, and `=ge=` operators compare numbers. The `=contains=` and `=excludes=` operators match elements within an array.  (optional)
    sortBy := "sortBy_example" // string | Specifies an object field name to sort results on. Provide a comma-separated list of field names for additional sort criteria. Prefix field names with a dash for descending sort order. For example, `name,-tlsIngestProfileId` sorts primarily by name, then IDs in descending order.  (optional)
    pageNumber := int32(56) // int32 | Specifies an offset within the full set of results, with page numbers starting at 1.  Use this in conjunction with `pageSize`.  (optional)
    pageSize := int32(56) // int32 | Specifies the number of results on each page.  Use this in conjunction with `pageNumber`.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListTlsIngestProfiles(context.Background(), contentProviderId).View(view).Filter(filter).SortBy(sortBy).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListTlsIngestProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTlsIngestProfiles`: TlsIngestProfileCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListTlsIngestProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTlsIngestProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **string** | Results display either &#x60;summary&#x60; or &#x60;detailed&#x60; objects.  | 
 **filter** | **string** | Specifies a query to filter down the set of results. The accompanying example specifies more than one query, where &#x60;;&#x60; and &#x60;,&#x60; delimiters clarify AND and OR logic. (Optional parentheses clarify logical groupings.) The &#x60;&#x3D;&#x3D;&#x60; and &#x60;!&#x3D;&#x60; operators test equivalency. The &#x60;&#x3D;lt&#x3D;&#x60;, &#x60;&#x3D;le&#x3D;&#x60;, &#x60;&#x3D;gt&#x3D;&#x60;, and &#x60;&#x3D;ge&#x3D;&#x60; operators compare numbers. The &#x60;&#x3D;contains&#x3D;&#x60; and &#x60;&#x3D;excludes&#x3D;&#x60; operators match elements within an array.  | 
 **sortBy** | **string** | Specifies an object field name to sort results on. Provide a comma-separated list of field names for additional sort criteria. Prefix field names with a dash for descending sort order. For example, &#x60;name,-tlsIngestProfileId&#x60; sorts primarily by name, then IDs in descending order.  | 
 **pageNumber** | **int32** | Specifies an offset within the full set of results, with page numbers starting at 1.  Use this in conjunction with &#x60;pageSize&#x60;.  | 
 **pageSize** | **int32** | Specifies the number of results on each page.  Use this in conjunction with &#x60;pageNumber&#x60;.  | 

### Return type

[**TlsIngestProfileCollection**](TlsIngestProfileCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCdnPrefix

> CdnPrefixReadDetailed ReadCdnPrefix(ctx, contentProviderId, cdnPrefixId).Execute()

cdn prefix



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    cdnPrefixId := int32(56) // int32 | Unique identifier for each CDN prefix. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ReadCdnPrefix(context.Background(), contentProviderId, cdnPrefixId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReadCdnPrefix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadCdnPrefix`: CdnPrefixReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReadCdnPrefix`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 
**cdnPrefixId** | **int32** | Unique identifier for each CDN prefix.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadCdnPrefixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CdnPrefixReadDetailed**](CdnPrefixReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCdnPrefixRules

> CdnPrefixRules ReadCdnPrefixRules(ctx, contentProviderId, cdnPrefixId).Execute()

rules



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    cdnPrefixId := int32(56) // int32 | Unique identifier for each CDN prefix. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ReadCdnPrefixRules(context.Background(), contentProviderId, cdnPrefixId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReadCdnPrefixRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadCdnPrefixRules`: CdnPrefixRules
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReadCdnPrefixRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 
**cdnPrefixId** | **int32** | Unique identifier for each CDN prefix.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadCdnPrefixRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CdnPrefixRules**](CdnPrefixRules.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadContentProvider

> ContentProviderReadDetailed ReadContentProvider(ctx, contentProviderId).Execute()

content provider



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
    contentProviderId := int32(56) // int32 | Unique identifier for each Content Provider. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ReadContentProvider(context.Background(), contentProviderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReadContentProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadContentProvider`: ContentProviderReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReadContentProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each Content Provider.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadContentProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContentProviderReadDetailed**](ContentProviderReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOrigin

> OriginReadDetailed ReadOrigin(ctx, contentProviderId, originId).Execute()

origin



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    originId := int32(56) // int32 | Unique identifier for each origin. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ReadOrigin(context.Background(), contentProviderId, originId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReadOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOrigin`: OriginReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReadOrigin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 
**originId** | **int32** | Unique identifier for each origin.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOriginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OriginReadDetailed**](OriginReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadSharedSecretSet

> SharedSecretSetReadDetailed ReadSharedSecretSet(ctx, contentProviderId, sharedSecretSetId).Execute()

shared secret set



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    sharedSecretSetId := int32(56) // int32 | Unique identifier for each shared secret set. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ReadSharedSecretSet(context.Background(), contentProviderId, sharedSecretSetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReadSharedSecretSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadSharedSecretSet`: SharedSecretSetReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReadSharedSecretSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 
**sharedSecretSetId** | **int32** | Unique identifier for each shared secret set.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadSharedSecretSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SharedSecretSetReadDetailed**](SharedSecretSetReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadSharedSecretSetSecrets

> SharedSecretSetSecrets ReadSharedSecretSetSecrets(ctx, contentProviderId, sharedSecretSetId).Execute()

secrets



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    sharedSecretSetId := int32(56) // int32 | Unique identifier for each shared secret set. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ReadSharedSecretSetSecrets(context.Background(), contentProviderId, sharedSecretSetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReadSharedSecretSetSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadSharedSecretSetSecrets`: SharedSecretSetSecrets
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReadSharedSecretSetSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 
**sharedSecretSetId** | **int32** | Unique identifier for each shared secret set.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadSharedSecretSetSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SharedSecretSetSecrets**](SharedSecretSetSecrets.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadTlsDeliveryProfile

> TlsDeliveryProfileReadDetailed ReadTlsDeliveryProfile(ctx, contentProviderId, tlsDeliveryProfileId).Execute()

tls delivery profile



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    tlsDeliveryProfileId := int32(56) // int32 | Unique identifier for each TLS delivery profile. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ReadTlsDeliveryProfile(context.Background(), contentProviderId, tlsDeliveryProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReadTlsDeliveryProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadTlsDeliveryProfile`: TlsDeliveryProfileReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReadTlsDeliveryProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 
**tlsDeliveryProfileId** | **int32** | Unique identifier for each TLS delivery profile.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadTlsDeliveryProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TlsDeliveryProfileReadDetailed**](TlsDeliveryProfileReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadTlsIngestProfile

> TlsIngestProfileReadDetailed ReadTlsIngestProfile(ctx, contentProviderId, tlsIngestProfileId).Execute()

tls ingest profile



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    tlsIngestProfileId := int32(56) // int32 | Unique identifier for each TLS ingest profile. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ReadTlsIngestProfile(context.Background(), contentProviderId, tlsIngestProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReadTlsIngestProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadTlsIngestProfile`: TlsIngestProfileReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReadTlsIngestProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 
**tlsIngestProfileId** | **int32** | Unique identifier for each TLS ingest profile.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadTlsIngestProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TlsIngestProfileReadDetailed**](TlsIngestProfileReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCdnPrefix

> CdnPrefixReadDetailed UpdateCdnPrefix(ctx, contentProviderId, cdnPrefixId).CdnPrefixUpdate(cdnPrefixUpdate).Execute()

cdn prefix



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    cdnPrefixId := int32(56) // int32 | Unique identifier for each CDN prefix. 
    cdnPrefixUpdate := *openapiclient.NewCdnPrefixUpdate(int32(123), "Prefix_example", int32(123)) // CdnPrefixUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateCdnPrefix(context.Background(), contentProviderId, cdnPrefixId).CdnPrefixUpdate(cdnPrefixUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateCdnPrefix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCdnPrefix`: CdnPrefixReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateCdnPrefix`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 
**cdnPrefixId** | **int32** | Unique identifier for each CDN prefix.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCdnPrefixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cdnPrefixUpdate** | [**CdnPrefixUpdate**](CdnPrefixUpdate.md) |  | 

### Return type

[**CdnPrefixReadDetailed**](CdnPrefixReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCdnPrefixRules

> CdnPrefixRules UpdateCdnPrefixRules(ctx, contentProviderId, cdnPrefixId).CdnPrefixRules(cdnPrefixRules).Execute()

rules



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    cdnPrefixId := int32(56) // int32 | Unique identifier for each CDN prefix. 
    cdnPrefixRules := *openapiclient.NewCdnPrefixRules() // CdnPrefixRules | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateCdnPrefixRules(context.Background(), contentProviderId, cdnPrefixId).CdnPrefixRules(cdnPrefixRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateCdnPrefixRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCdnPrefixRules`: CdnPrefixRules
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateCdnPrefixRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 
**cdnPrefixId** | **int32** | Unique identifier for each CDN prefix.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCdnPrefixRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cdnPrefixRules** | [**CdnPrefixRules**](CdnPrefixRules.md) |  | 

### Return type

[**CdnPrefixRules**](CdnPrefixRules.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrigin

> OriginReadDetailed UpdateOrigin(ctx, contentProviderId, originId).OriginUpdate(originUpdate).Execute()

origin



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    originId := int32(56) // int32 | Unique identifier for each origin. 
    originUpdate := *openapiclient.NewOriginUpdate("Hostname_example", int32(123), int32(123)) // OriginUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateOrigin(context.Background(), contentProviderId, originId).OriginUpdate(originUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrigin`: OriginReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateOrigin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 
**originId** | **int32** | Unique identifier for each origin.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOriginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **originUpdate** | [**OriginUpdate**](OriginUpdate.md) |  | 

### Return type

[**OriginReadDetailed**](OriginReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSharedSecretSetSecrets

> SharedSecretSetSecrets UpdateSharedSecretSetSecrets(ctx, contentProviderId, sharedSecretSetId).SharedSecretSetSecrets(sharedSecretSetSecrets).Execute()

secrets



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    sharedSecretSetId := int32(56) // int32 | Unique identifier for each shared secret set. 
    sharedSecretSetSecrets := *openapiclient.NewSharedSecretSetSecrets(map[string]string{"key": "Inner_example"}) // SharedSecretSetSecrets | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSharedSecretSetSecrets(context.Background(), contentProviderId, sharedSecretSetId).SharedSecretSetSecrets(sharedSecretSetSecrets).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSharedSecretSetSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSharedSecretSetSecrets`: SharedSecretSetSecrets
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSharedSecretSetSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 
**sharedSecretSetId** | **int32** | Unique identifier for each shared secret set.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSharedSecretSetSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sharedSecretSetSecrets** | [**SharedSecretSetSecrets**](SharedSecretSetSecrets.md) |  | 

### Return type

[**SharedSecretSetSecrets**](SharedSecretSetSecrets.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTlsDeliveryProfile

> TlsDeliveryProfileReadDetailed UpdateTlsDeliveryProfile(ctx, contentProviderId, tlsDeliveryProfileId).TlsDeliveryProfileUpdate(tlsDeliveryProfileUpdate).Execute()

tls delivery profile



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    tlsDeliveryProfileId := int32(56) // int32 | Unique identifier for each TLS delivery profile. 
    tlsDeliveryProfileUpdate := *openapiclient.NewTlsDeliveryProfileUpdate(int32(123), "Name_example", int32(123)) // TlsDeliveryProfileUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateTlsDeliveryProfile(context.Background(), contentProviderId, tlsDeliveryProfileId).TlsDeliveryProfileUpdate(tlsDeliveryProfileUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateTlsDeliveryProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTlsDeliveryProfile`: TlsDeliveryProfileReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateTlsDeliveryProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 
**tlsDeliveryProfileId** | **int32** | Unique identifier for each TLS delivery profile.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTlsDeliveryProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tlsDeliveryProfileUpdate** | [**TlsDeliveryProfileUpdate**](TlsDeliveryProfileUpdate.md) |  | 

### Return type

[**TlsDeliveryProfileReadDetailed**](TlsDeliveryProfileReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTlsIngestProfile

> TlsIngestProfileReadDetailed UpdateTlsIngestProfile(ctx, contentProviderId, tlsIngestProfileId).TlsIngestProfileUpdate(tlsIngestProfileUpdate).Execute()

tls ingest profile



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
    contentProviderId := int32(56) // int32 | Unique identifier for each content provider. 
    tlsIngestProfileId := int32(56) // int32 | Unique identifier for each TLS ingest profile. 
    tlsIngestProfileUpdate := *openapiclient.NewTlsIngestProfileUpdate(int32(123), int32(123), "Name_example") // TlsIngestProfileUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateTlsIngestProfile(context.Background(), contentProviderId, tlsIngestProfileId).TlsIngestProfileUpdate(tlsIngestProfileUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateTlsIngestProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTlsIngestProfile`: TlsIngestProfileReadDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateTlsIngestProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentProviderId** | **int32** | Unique identifier for each content provider.  | 
**tlsIngestProfileId** | **int32** | Unique identifier for each TLS ingest profile.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTlsIngestProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tlsIngestProfileUpdate** | [**TlsIngestProfileUpdate**](TlsIngestProfileUpdate.md) |  | 

### Return type

[**TlsIngestProfileReadDetailed**](TlsIngestProfileReadDetailed.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

