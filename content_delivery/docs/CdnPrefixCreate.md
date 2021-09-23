# CdnPrefixCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] [default to true]
**ContentProviderId** | **int32** |  | 
**SiteMapId** | Pointer to **int32** |  | [optional] 
**PrefixPrioritization** | Pointer to [**PrefixPrioritizationType**](PrefixPrioritizationType.md) |  | [optional] [default to PREFIXPRIORITIZATIONTYPE_HIGH]
**KeepaliveRequests** | Pointer to **int32** | Sets the maximum number of requests that can be served through one keep-alive connection. After the maximum number of requests are made, the connection is closed. | [optional] 
**AccessMapId** | Pointer to **int32** |  | [optional] 
**Prefix** | **string** | The CDN prefix registered on behalf of the content provider. For example, &#x60;cdn.example.com&#x60;. | 
**DnsTtl** | Pointer to **int32** | DNS response time-to-live (TTL) value for this CDN prefix, in seconds. Note that setting this field is only recommended for CDN operators. | [optional] 
**IpAddressTagId** | Pointer to **int32** |  | [optional] 

## Methods

### NewCdnPrefixCreate

`func NewCdnPrefixCreate(contentProviderId int32, prefix string, ) *CdnPrefixCreate`

NewCdnPrefixCreate instantiates a new CdnPrefixCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnPrefixCreateWithDefaults

`func NewCdnPrefixCreateWithDefaults() *CdnPrefixCreate`

NewCdnPrefixCreateWithDefaults instantiates a new CdnPrefixCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *CdnPrefixCreate) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *CdnPrefixCreate) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *CdnPrefixCreate) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *CdnPrefixCreate) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetContentProviderId

`func (o *CdnPrefixCreate) GetContentProviderId() int32`

GetContentProviderId returns the ContentProviderId field if non-nil, zero value otherwise.

### GetContentProviderIdOk

`func (o *CdnPrefixCreate) GetContentProviderIdOk() (*int32, bool)`

GetContentProviderIdOk returns a tuple with the ContentProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentProviderId

`func (o *CdnPrefixCreate) SetContentProviderId(v int32)`

SetContentProviderId sets ContentProviderId field to given value.


### GetSiteMapId

`func (o *CdnPrefixCreate) GetSiteMapId() int32`

GetSiteMapId returns the SiteMapId field if non-nil, zero value otherwise.

### GetSiteMapIdOk

`func (o *CdnPrefixCreate) GetSiteMapIdOk() (*int32, bool)`

GetSiteMapIdOk returns a tuple with the SiteMapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteMapId

`func (o *CdnPrefixCreate) SetSiteMapId(v int32)`

SetSiteMapId sets SiteMapId field to given value.

### HasSiteMapId

`func (o *CdnPrefixCreate) HasSiteMapId() bool`

HasSiteMapId returns a boolean if a field has been set.

### GetPrefixPrioritization

`func (o *CdnPrefixCreate) GetPrefixPrioritization() PrefixPrioritizationType`

GetPrefixPrioritization returns the PrefixPrioritization field if non-nil, zero value otherwise.

### GetPrefixPrioritizationOk

`func (o *CdnPrefixCreate) GetPrefixPrioritizationOk() (*PrefixPrioritizationType, bool)`

GetPrefixPrioritizationOk returns a tuple with the PrefixPrioritization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixPrioritization

`func (o *CdnPrefixCreate) SetPrefixPrioritization(v PrefixPrioritizationType)`

SetPrefixPrioritization sets PrefixPrioritization field to given value.

### HasPrefixPrioritization

`func (o *CdnPrefixCreate) HasPrefixPrioritization() bool`

HasPrefixPrioritization returns a boolean if a field has been set.

### GetKeepaliveRequests

`func (o *CdnPrefixCreate) GetKeepaliveRequests() int32`

GetKeepaliveRequests returns the KeepaliveRequests field if non-nil, zero value otherwise.

### GetKeepaliveRequestsOk

`func (o *CdnPrefixCreate) GetKeepaliveRequestsOk() (*int32, bool)`

GetKeepaliveRequestsOk returns a tuple with the KeepaliveRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveRequests

`func (o *CdnPrefixCreate) SetKeepaliveRequests(v int32)`

SetKeepaliveRequests sets KeepaliveRequests field to given value.

### HasKeepaliveRequests

`func (o *CdnPrefixCreate) HasKeepaliveRequests() bool`

HasKeepaliveRequests returns a boolean if a field has been set.

### GetAccessMapId

`func (o *CdnPrefixCreate) GetAccessMapId() int32`

GetAccessMapId returns the AccessMapId field if non-nil, zero value otherwise.

### GetAccessMapIdOk

`func (o *CdnPrefixCreate) GetAccessMapIdOk() (*int32, bool)`

GetAccessMapIdOk returns a tuple with the AccessMapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMapId

`func (o *CdnPrefixCreate) SetAccessMapId(v int32)`

SetAccessMapId sets AccessMapId field to given value.

### HasAccessMapId

`func (o *CdnPrefixCreate) HasAccessMapId() bool`

HasAccessMapId returns a boolean if a field has been set.

### GetPrefix

`func (o *CdnPrefixCreate) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *CdnPrefixCreate) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *CdnPrefixCreate) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetDnsTtl

`func (o *CdnPrefixCreate) GetDnsTtl() int32`

GetDnsTtl returns the DnsTtl field if non-nil, zero value otherwise.

### GetDnsTtlOk

`func (o *CdnPrefixCreate) GetDnsTtlOk() (*int32, bool)`

GetDnsTtlOk returns a tuple with the DnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTtl

`func (o *CdnPrefixCreate) SetDnsTtl(v int32)`

SetDnsTtl sets DnsTtl field to given value.

### HasDnsTtl

`func (o *CdnPrefixCreate) HasDnsTtl() bool`

HasDnsTtl returns a boolean if a field has been set.

### GetIpAddressTagId

`func (o *CdnPrefixCreate) GetIpAddressTagId() int32`

GetIpAddressTagId returns the IpAddressTagId field if non-nil, zero value otherwise.

### GetIpAddressTagIdOk

`func (o *CdnPrefixCreate) GetIpAddressTagIdOk() (*int32, bool)`

GetIpAddressTagIdOk returns a tuple with the IpAddressTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressTagId

`func (o *CdnPrefixCreate) SetIpAddressTagId(v int32)`

SetIpAddressTagId sets IpAddressTagId field to given value.

### HasIpAddressTagId

`func (o *CdnPrefixCreate) HasIpAddressTagId() bool`

HasIpAddressTagId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


