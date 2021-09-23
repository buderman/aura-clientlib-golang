# CdnPrefix3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] [default to true]
**ContentProviderId** | Pointer to **int32** |  | [optional] 
**SiteMapId** | Pointer to **int32** |  | [optional] 
**PrefixPrioritization** | Pointer to [**PrefixPrioritizationType**](PrefixPrioritizationType.md) |  | [optional] [default to PREFIXPRIORITIZATIONTYPE_HIGH]
**KeepaliveRequests** | Pointer to **int32** | Sets the maximum number of requests that can be served through one keep-alive connection. After the maximum number of requests are made, the connection is closed. | [optional] 
**AccessMapId** | Pointer to **int32** |  | [optional] 
**Prefix** | **string** | The CDN prefix registered on behalf of the content provider. For example, &#x60;cdn.example.com&#x60;. | 
**DnsTtl** | Pointer to **int32** | DNS response time-to-live (TTL) value for this CDN prefix, in seconds. Note that setting this field is only recommended for CDN operators. | [optional] 
**CdnPrefixId** | **int32** |  | 
**IpAddressTagId** | Pointer to **int32** |  | [optional] 

## Methods

### NewCdnPrefix3

`func NewCdnPrefix3(prefix string, cdnPrefixId int32, ) *CdnPrefix3`

NewCdnPrefix3 instantiates a new CdnPrefix3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnPrefix3WithDefaults

`func NewCdnPrefix3WithDefaults() *CdnPrefix3`

NewCdnPrefix3WithDefaults instantiates a new CdnPrefix3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *CdnPrefix3) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *CdnPrefix3) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *CdnPrefix3) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *CdnPrefix3) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetContentProviderId

`func (o *CdnPrefix3) GetContentProviderId() int32`

GetContentProviderId returns the ContentProviderId field if non-nil, zero value otherwise.

### GetContentProviderIdOk

`func (o *CdnPrefix3) GetContentProviderIdOk() (*int32, bool)`

GetContentProviderIdOk returns a tuple with the ContentProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentProviderId

`func (o *CdnPrefix3) SetContentProviderId(v int32)`

SetContentProviderId sets ContentProviderId field to given value.

### HasContentProviderId

`func (o *CdnPrefix3) HasContentProviderId() bool`

HasContentProviderId returns a boolean if a field has been set.

### GetSiteMapId

`func (o *CdnPrefix3) GetSiteMapId() int32`

GetSiteMapId returns the SiteMapId field if non-nil, zero value otherwise.

### GetSiteMapIdOk

`func (o *CdnPrefix3) GetSiteMapIdOk() (*int32, bool)`

GetSiteMapIdOk returns a tuple with the SiteMapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteMapId

`func (o *CdnPrefix3) SetSiteMapId(v int32)`

SetSiteMapId sets SiteMapId field to given value.

### HasSiteMapId

`func (o *CdnPrefix3) HasSiteMapId() bool`

HasSiteMapId returns a boolean if a field has been set.

### GetPrefixPrioritization

`func (o *CdnPrefix3) GetPrefixPrioritization() PrefixPrioritizationType`

GetPrefixPrioritization returns the PrefixPrioritization field if non-nil, zero value otherwise.

### GetPrefixPrioritizationOk

`func (o *CdnPrefix3) GetPrefixPrioritizationOk() (*PrefixPrioritizationType, bool)`

GetPrefixPrioritizationOk returns a tuple with the PrefixPrioritization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixPrioritization

`func (o *CdnPrefix3) SetPrefixPrioritization(v PrefixPrioritizationType)`

SetPrefixPrioritization sets PrefixPrioritization field to given value.

### HasPrefixPrioritization

`func (o *CdnPrefix3) HasPrefixPrioritization() bool`

HasPrefixPrioritization returns a boolean if a field has been set.

### GetKeepaliveRequests

`func (o *CdnPrefix3) GetKeepaliveRequests() int32`

GetKeepaliveRequests returns the KeepaliveRequests field if non-nil, zero value otherwise.

### GetKeepaliveRequestsOk

`func (o *CdnPrefix3) GetKeepaliveRequestsOk() (*int32, bool)`

GetKeepaliveRequestsOk returns a tuple with the KeepaliveRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveRequests

`func (o *CdnPrefix3) SetKeepaliveRequests(v int32)`

SetKeepaliveRequests sets KeepaliveRequests field to given value.

### HasKeepaliveRequests

`func (o *CdnPrefix3) HasKeepaliveRequests() bool`

HasKeepaliveRequests returns a boolean if a field has been set.

### GetAccessMapId

`func (o *CdnPrefix3) GetAccessMapId() int32`

GetAccessMapId returns the AccessMapId field if non-nil, zero value otherwise.

### GetAccessMapIdOk

`func (o *CdnPrefix3) GetAccessMapIdOk() (*int32, bool)`

GetAccessMapIdOk returns a tuple with the AccessMapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMapId

`func (o *CdnPrefix3) SetAccessMapId(v int32)`

SetAccessMapId sets AccessMapId field to given value.

### HasAccessMapId

`func (o *CdnPrefix3) HasAccessMapId() bool`

HasAccessMapId returns a boolean if a field has been set.

### GetPrefix

`func (o *CdnPrefix3) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *CdnPrefix3) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *CdnPrefix3) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetDnsTtl

`func (o *CdnPrefix3) GetDnsTtl() int32`

GetDnsTtl returns the DnsTtl field if non-nil, zero value otherwise.

### GetDnsTtlOk

`func (o *CdnPrefix3) GetDnsTtlOk() (*int32, bool)`

GetDnsTtlOk returns a tuple with the DnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTtl

`func (o *CdnPrefix3) SetDnsTtl(v int32)`

SetDnsTtl sets DnsTtl field to given value.

### HasDnsTtl

`func (o *CdnPrefix3) HasDnsTtl() bool`

HasDnsTtl returns a boolean if a field has been set.

### GetCdnPrefixId

`func (o *CdnPrefix3) GetCdnPrefixId() int32`

GetCdnPrefixId returns the CdnPrefixId field if non-nil, zero value otherwise.

### GetCdnPrefixIdOk

`func (o *CdnPrefix3) GetCdnPrefixIdOk() (*int32, bool)`

GetCdnPrefixIdOk returns a tuple with the CdnPrefixId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnPrefixId

`func (o *CdnPrefix3) SetCdnPrefixId(v int32)`

SetCdnPrefixId sets CdnPrefixId field to given value.


### GetIpAddressTagId

`func (o *CdnPrefix3) GetIpAddressTagId() int32`

GetIpAddressTagId returns the IpAddressTagId field if non-nil, zero value otherwise.

### GetIpAddressTagIdOk

`func (o *CdnPrefix3) GetIpAddressTagIdOk() (*int32, bool)`

GetIpAddressTagIdOk returns a tuple with the IpAddressTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressTagId

`func (o *CdnPrefix3) SetIpAddressTagId(v int32)`

SetIpAddressTagId sets IpAddressTagId field to given value.

### HasIpAddressTagId

`func (o *CdnPrefix3) HasIpAddressTagId() bool`

HasIpAddressTagId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


