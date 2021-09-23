# Origin3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** |  | 
**CachingType** | Pointer to [**CachingType**](CachingType.md) |  | [optional] [default to CACHINGTYPE_OPTIMISTIC]
**HostnameOverride** | Pointer to **string** | A DNS hostname to use when connecting to the origin instead of that specified by the &#x60;hostname&#x60; field.  A port may be optionally specified by using hostname:port format. | [optional] 
**StoragePartitionId** | Pointer to **int32** |  | [optional] 
**EdgeHostType** | Pointer to [**EdgeHostType**](EdgeHostType.md) |  | [optional] [default to EDGEHOSTTYPE_HTTP]
**CacheKeyOriginHostnameOriginId** | Pointer to **int32** |  | [optional] 
**ContentProviderId** | Pointer to **int32** |  | [optional] 
**FastReroute** | Pointer to [**FastRerouteType**](FastRerouteType.md) |  | [optional] 
**EnableSiteRedirects** | Pointer to **bool** | Determines whether or not HyperCache redirects a client to a better site based on the rules in the site map. Enabling this to &#x60;true&#x60; allows the HyperCache to adhere to the LCDN site map configured using the request router service. If set to &#x60;false&#x60;, cache miss client requests to a CDN prefix that uses this origin server as a default result in a HyperCache request for the content directly to the origin server. | [optional] [default to true]
**InterSiteProtocol** | Pointer to [**ProtocolType**](ProtocolType.md) |  | [optional] [default to PROTOCOLTYPE_HTTP]
**EnableRequestIdExport** | Pointer to **bool** | Causes the Hypercache to send its internal request ID to the origin server when ingesting content. | [optional] [default to false]
**OriginId** | **int32** |  | 
**Enable** | Pointer to **bool** |  | [optional] [default to true]
**EnableAuthenticatedContent** | Pointer to **bool** | Causes the HyperCache to send a HEAD request to the origin server for each client request. If the response is a status code of 2XX, the content is served to the client. Otherwise, the origin&#39;s status code is returned. This typically occurs when the origin server authenticates the client request using a cookie, token, or other means. Default is &#x60;false&#x60;. | [optional] [default to false]
**DynamicHierarchy** | Pointer to [**DynamicHierarchyType**](DynamicHierarchyType.md) |  | [optional] 
**TlsIngestProfileId** | Pointer to **int32** |  | [optional] 
**IntraSiteProtocol** | Pointer to [**ProtocolType**](ProtocolType.md) |  | [optional] [default to PROTOCOLTYPE_HTTP]
**CacheableErrorResponseCodes** | Pointer to **[]string** | A list of HTTP status codes that the HPC caches. Each item is a single status code or range of status codes, where the starting code must be less than the ending code. For example, &#x60;209&#x60; and &#x60;310-320&#x60; are both valid. The following status codes are not allowed, either in single code, or as part of a range: 200, 203, 206, 300, 301, 410, 416. | [optional] 
**ResolvableHostnames** | Pointer to **[]string** | A list of virtual hosts, origin servers with resolvable hostnames. These resolvable hostnames are used by DNS to determine the IP addresses of the destination origin server. | [optional] 
**ErrorCacheMaxRetry** | Pointer to **int32** | The maximum number of retries in the case of an error response. | [optional] 
**EdgeHostname** | Pointer to **string** |  | [optional] 
**OriginTimeout** | Pointer to **int32** | Origin idle connection timeout in seconds, specifies the length of time the HPC waits for a response to a request from an origin. If no response is received from an origin server within the configured timeout value, four additional attempts are made to the same origin server. If no response is received for the final request within the configured timeout, then that connection is terminated, and a 504 response is sent to the requesting client. | [optional] [default to 300]
**ErrorCacheMaxAge** | Pointer to **int32** | The maximum amount of time the HTTP status codes is cached by the HPC. | [optional] 

## Methods

### NewOrigin3

`func NewOrigin3(hostname string, originId int32, ) *Origin3`

NewOrigin3 instantiates a new Origin3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrigin3WithDefaults

`func NewOrigin3WithDefaults() *Origin3`

NewOrigin3WithDefaults instantiates a new Origin3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *Origin3) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Origin3) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Origin3) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetCachingType

`func (o *Origin3) GetCachingType() CachingType`

GetCachingType returns the CachingType field if non-nil, zero value otherwise.

### GetCachingTypeOk

`func (o *Origin3) GetCachingTypeOk() (*CachingType, bool)`

GetCachingTypeOk returns a tuple with the CachingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachingType

`func (o *Origin3) SetCachingType(v CachingType)`

SetCachingType sets CachingType field to given value.

### HasCachingType

`func (o *Origin3) HasCachingType() bool`

HasCachingType returns a boolean if a field has been set.

### GetHostnameOverride

`func (o *Origin3) GetHostnameOverride() string`

GetHostnameOverride returns the HostnameOverride field if non-nil, zero value otherwise.

### GetHostnameOverrideOk

`func (o *Origin3) GetHostnameOverrideOk() (*string, bool)`

GetHostnameOverrideOk returns a tuple with the HostnameOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameOverride

`func (o *Origin3) SetHostnameOverride(v string)`

SetHostnameOverride sets HostnameOverride field to given value.

### HasHostnameOverride

`func (o *Origin3) HasHostnameOverride() bool`

HasHostnameOverride returns a boolean if a field has been set.

### GetStoragePartitionId

`func (o *Origin3) GetStoragePartitionId() int32`

GetStoragePartitionId returns the StoragePartitionId field if non-nil, zero value otherwise.

### GetStoragePartitionIdOk

`func (o *Origin3) GetStoragePartitionIdOk() (*int32, bool)`

GetStoragePartitionIdOk returns a tuple with the StoragePartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePartitionId

`func (o *Origin3) SetStoragePartitionId(v int32)`

SetStoragePartitionId sets StoragePartitionId field to given value.

### HasStoragePartitionId

`func (o *Origin3) HasStoragePartitionId() bool`

HasStoragePartitionId returns a boolean if a field has been set.

### GetEdgeHostType

`func (o *Origin3) GetEdgeHostType() EdgeHostType`

GetEdgeHostType returns the EdgeHostType field if non-nil, zero value otherwise.

### GetEdgeHostTypeOk

`func (o *Origin3) GetEdgeHostTypeOk() (*EdgeHostType, bool)`

GetEdgeHostTypeOk returns a tuple with the EdgeHostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeHostType

`func (o *Origin3) SetEdgeHostType(v EdgeHostType)`

SetEdgeHostType sets EdgeHostType field to given value.

### HasEdgeHostType

`func (o *Origin3) HasEdgeHostType() bool`

HasEdgeHostType returns a boolean if a field has been set.

### GetCacheKeyOriginHostnameOriginId

`func (o *Origin3) GetCacheKeyOriginHostnameOriginId() int32`

GetCacheKeyOriginHostnameOriginId returns the CacheKeyOriginHostnameOriginId field if non-nil, zero value otherwise.

### GetCacheKeyOriginHostnameOriginIdOk

`func (o *Origin3) GetCacheKeyOriginHostnameOriginIdOk() (*int32, bool)`

GetCacheKeyOriginHostnameOriginIdOk returns a tuple with the CacheKeyOriginHostnameOriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheKeyOriginHostnameOriginId

`func (o *Origin3) SetCacheKeyOriginHostnameOriginId(v int32)`

SetCacheKeyOriginHostnameOriginId sets CacheKeyOriginHostnameOriginId field to given value.

### HasCacheKeyOriginHostnameOriginId

`func (o *Origin3) HasCacheKeyOriginHostnameOriginId() bool`

HasCacheKeyOriginHostnameOriginId returns a boolean if a field has been set.

### GetContentProviderId

`func (o *Origin3) GetContentProviderId() int32`

GetContentProviderId returns the ContentProviderId field if non-nil, zero value otherwise.

### GetContentProviderIdOk

`func (o *Origin3) GetContentProviderIdOk() (*int32, bool)`

GetContentProviderIdOk returns a tuple with the ContentProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentProviderId

`func (o *Origin3) SetContentProviderId(v int32)`

SetContentProviderId sets ContentProviderId field to given value.

### HasContentProviderId

`func (o *Origin3) HasContentProviderId() bool`

HasContentProviderId returns a boolean if a field has been set.

### GetFastReroute

`func (o *Origin3) GetFastReroute() FastRerouteType`

GetFastReroute returns the FastReroute field if non-nil, zero value otherwise.

### GetFastRerouteOk

`func (o *Origin3) GetFastRerouteOk() (*FastRerouteType, bool)`

GetFastRerouteOk returns a tuple with the FastReroute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastReroute

`func (o *Origin3) SetFastReroute(v FastRerouteType)`

SetFastReroute sets FastReroute field to given value.

### HasFastReroute

`func (o *Origin3) HasFastReroute() bool`

HasFastReroute returns a boolean if a field has been set.

### GetEnableSiteRedirects

`func (o *Origin3) GetEnableSiteRedirects() bool`

GetEnableSiteRedirects returns the EnableSiteRedirects field if non-nil, zero value otherwise.

### GetEnableSiteRedirectsOk

`func (o *Origin3) GetEnableSiteRedirectsOk() (*bool, bool)`

GetEnableSiteRedirectsOk returns a tuple with the EnableSiteRedirects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSiteRedirects

`func (o *Origin3) SetEnableSiteRedirects(v bool)`

SetEnableSiteRedirects sets EnableSiteRedirects field to given value.

### HasEnableSiteRedirects

`func (o *Origin3) HasEnableSiteRedirects() bool`

HasEnableSiteRedirects returns a boolean if a field has been set.

### GetInterSiteProtocol

`func (o *Origin3) GetInterSiteProtocol() ProtocolType`

GetInterSiteProtocol returns the InterSiteProtocol field if non-nil, zero value otherwise.

### GetInterSiteProtocolOk

`func (o *Origin3) GetInterSiteProtocolOk() (*ProtocolType, bool)`

GetInterSiteProtocolOk returns a tuple with the InterSiteProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterSiteProtocol

`func (o *Origin3) SetInterSiteProtocol(v ProtocolType)`

SetInterSiteProtocol sets InterSiteProtocol field to given value.

### HasInterSiteProtocol

`func (o *Origin3) HasInterSiteProtocol() bool`

HasInterSiteProtocol returns a boolean if a field has been set.

### GetEnableRequestIdExport

`func (o *Origin3) GetEnableRequestIdExport() bool`

GetEnableRequestIdExport returns the EnableRequestIdExport field if non-nil, zero value otherwise.

### GetEnableRequestIdExportOk

`func (o *Origin3) GetEnableRequestIdExportOk() (*bool, bool)`

GetEnableRequestIdExportOk returns a tuple with the EnableRequestIdExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRequestIdExport

`func (o *Origin3) SetEnableRequestIdExport(v bool)`

SetEnableRequestIdExport sets EnableRequestIdExport field to given value.

### HasEnableRequestIdExport

`func (o *Origin3) HasEnableRequestIdExport() bool`

HasEnableRequestIdExport returns a boolean if a field has been set.

### GetOriginId

`func (o *Origin3) GetOriginId() int32`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *Origin3) GetOriginIdOk() (*int32, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *Origin3) SetOriginId(v int32)`

SetOriginId sets OriginId field to given value.


### GetEnable

`func (o *Origin3) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *Origin3) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *Origin3) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *Origin3) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableAuthenticatedContent

`func (o *Origin3) GetEnableAuthenticatedContent() bool`

GetEnableAuthenticatedContent returns the EnableAuthenticatedContent field if non-nil, zero value otherwise.

### GetEnableAuthenticatedContentOk

`func (o *Origin3) GetEnableAuthenticatedContentOk() (*bool, bool)`

GetEnableAuthenticatedContentOk returns a tuple with the EnableAuthenticatedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAuthenticatedContent

`func (o *Origin3) SetEnableAuthenticatedContent(v bool)`

SetEnableAuthenticatedContent sets EnableAuthenticatedContent field to given value.

### HasEnableAuthenticatedContent

`func (o *Origin3) HasEnableAuthenticatedContent() bool`

HasEnableAuthenticatedContent returns a boolean if a field has been set.

### GetDynamicHierarchy

`func (o *Origin3) GetDynamicHierarchy() DynamicHierarchyType`

GetDynamicHierarchy returns the DynamicHierarchy field if non-nil, zero value otherwise.

### GetDynamicHierarchyOk

`func (o *Origin3) GetDynamicHierarchyOk() (*DynamicHierarchyType, bool)`

GetDynamicHierarchyOk returns a tuple with the DynamicHierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicHierarchy

`func (o *Origin3) SetDynamicHierarchy(v DynamicHierarchyType)`

SetDynamicHierarchy sets DynamicHierarchy field to given value.

### HasDynamicHierarchy

`func (o *Origin3) HasDynamicHierarchy() bool`

HasDynamicHierarchy returns a boolean if a field has been set.

### GetTlsIngestProfileId

`func (o *Origin3) GetTlsIngestProfileId() int32`

GetTlsIngestProfileId returns the TlsIngestProfileId field if non-nil, zero value otherwise.

### GetTlsIngestProfileIdOk

`func (o *Origin3) GetTlsIngestProfileIdOk() (*int32, bool)`

GetTlsIngestProfileIdOk returns a tuple with the TlsIngestProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsIngestProfileId

`func (o *Origin3) SetTlsIngestProfileId(v int32)`

SetTlsIngestProfileId sets TlsIngestProfileId field to given value.

### HasTlsIngestProfileId

`func (o *Origin3) HasTlsIngestProfileId() bool`

HasTlsIngestProfileId returns a boolean if a field has been set.

### GetIntraSiteProtocol

`func (o *Origin3) GetIntraSiteProtocol() ProtocolType`

GetIntraSiteProtocol returns the IntraSiteProtocol field if non-nil, zero value otherwise.

### GetIntraSiteProtocolOk

`func (o *Origin3) GetIntraSiteProtocolOk() (*ProtocolType, bool)`

GetIntraSiteProtocolOk returns a tuple with the IntraSiteProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntraSiteProtocol

`func (o *Origin3) SetIntraSiteProtocol(v ProtocolType)`

SetIntraSiteProtocol sets IntraSiteProtocol field to given value.

### HasIntraSiteProtocol

`func (o *Origin3) HasIntraSiteProtocol() bool`

HasIntraSiteProtocol returns a boolean if a field has been set.

### GetCacheableErrorResponseCodes

`func (o *Origin3) GetCacheableErrorResponseCodes() []string`

GetCacheableErrorResponseCodes returns the CacheableErrorResponseCodes field if non-nil, zero value otherwise.

### GetCacheableErrorResponseCodesOk

`func (o *Origin3) GetCacheableErrorResponseCodesOk() (*[]string, bool)`

GetCacheableErrorResponseCodesOk returns a tuple with the CacheableErrorResponseCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheableErrorResponseCodes

`func (o *Origin3) SetCacheableErrorResponseCodes(v []string)`

SetCacheableErrorResponseCodes sets CacheableErrorResponseCodes field to given value.

### HasCacheableErrorResponseCodes

`func (o *Origin3) HasCacheableErrorResponseCodes() bool`

HasCacheableErrorResponseCodes returns a boolean if a field has been set.

### GetResolvableHostnames

`func (o *Origin3) GetResolvableHostnames() []string`

GetResolvableHostnames returns the ResolvableHostnames field if non-nil, zero value otherwise.

### GetResolvableHostnamesOk

`func (o *Origin3) GetResolvableHostnamesOk() (*[]string, bool)`

GetResolvableHostnamesOk returns a tuple with the ResolvableHostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvableHostnames

`func (o *Origin3) SetResolvableHostnames(v []string)`

SetResolvableHostnames sets ResolvableHostnames field to given value.

### HasResolvableHostnames

`func (o *Origin3) HasResolvableHostnames() bool`

HasResolvableHostnames returns a boolean if a field has been set.

### GetErrorCacheMaxRetry

`func (o *Origin3) GetErrorCacheMaxRetry() int32`

GetErrorCacheMaxRetry returns the ErrorCacheMaxRetry field if non-nil, zero value otherwise.

### GetErrorCacheMaxRetryOk

`func (o *Origin3) GetErrorCacheMaxRetryOk() (*int32, bool)`

GetErrorCacheMaxRetryOk returns a tuple with the ErrorCacheMaxRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCacheMaxRetry

`func (o *Origin3) SetErrorCacheMaxRetry(v int32)`

SetErrorCacheMaxRetry sets ErrorCacheMaxRetry field to given value.

### HasErrorCacheMaxRetry

`func (o *Origin3) HasErrorCacheMaxRetry() bool`

HasErrorCacheMaxRetry returns a boolean if a field has been set.

### GetEdgeHostname

`func (o *Origin3) GetEdgeHostname() string`

GetEdgeHostname returns the EdgeHostname field if non-nil, zero value otherwise.

### GetEdgeHostnameOk

`func (o *Origin3) GetEdgeHostnameOk() (*string, bool)`

GetEdgeHostnameOk returns a tuple with the EdgeHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeHostname

`func (o *Origin3) SetEdgeHostname(v string)`

SetEdgeHostname sets EdgeHostname field to given value.

### HasEdgeHostname

`func (o *Origin3) HasEdgeHostname() bool`

HasEdgeHostname returns a boolean if a field has been set.

### GetOriginTimeout

`func (o *Origin3) GetOriginTimeout() int32`

GetOriginTimeout returns the OriginTimeout field if non-nil, zero value otherwise.

### GetOriginTimeoutOk

`func (o *Origin3) GetOriginTimeoutOk() (*int32, bool)`

GetOriginTimeoutOk returns a tuple with the OriginTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTimeout

`func (o *Origin3) SetOriginTimeout(v int32)`

SetOriginTimeout sets OriginTimeout field to given value.

### HasOriginTimeout

`func (o *Origin3) HasOriginTimeout() bool`

HasOriginTimeout returns a boolean if a field has been set.

### GetErrorCacheMaxAge

`func (o *Origin3) GetErrorCacheMaxAge() int32`

GetErrorCacheMaxAge returns the ErrorCacheMaxAge field if non-nil, zero value otherwise.

### GetErrorCacheMaxAgeOk

`func (o *Origin3) GetErrorCacheMaxAgeOk() (*int32, bool)`

GetErrorCacheMaxAgeOk returns a tuple with the ErrorCacheMaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCacheMaxAge

`func (o *Origin3) SetErrorCacheMaxAge(v int32)`

SetErrorCacheMaxAge sets ErrorCacheMaxAge field to given value.

### HasErrorCacheMaxAge

`func (o *Origin3) HasErrorCacheMaxAge() bool`

HasErrorCacheMaxAge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


