# ClientServingLabelType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | **string** | The node&#39;s IP address to which the HyperCache instance is bound. Specify this IP address for serving client-facing traffic. | 
**IpAddressTagId** | Pointer to **int32** | The unique identifier for an IP address tag. Use the IP address tag to route traffic on a per-IP address basis and to provide service differentiation based on HyperCache IP addresses. | [optional] 

## Methods

### NewClientServingLabelType

`func NewClientServingLabelType(ipAddress string, ) *ClientServingLabelType`

NewClientServingLabelType instantiates a new ClientServingLabelType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientServingLabelTypeWithDefaults

`func NewClientServingLabelTypeWithDefaults() *ClientServingLabelType`

NewClientServingLabelTypeWithDefaults instantiates a new ClientServingLabelType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *ClientServingLabelType) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ClientServingLabelType) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ClientServingLabelType) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetIpAddressTagId

`func (o *ClientServingLabelType) GetIpAddressTagId() int32`

GetIpAddressTagId returns the IpAddressTagId field if non-nil, zero value otherwise.

### GetIpAddressTagIdOk

`func (o *ClientServingLabelType) GetIpAddressTagIdOk() (*int32, bool)`

GetIpAddressTagIdOk returns a tuple with the IpAddressTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressTagId

`func (o *ClientServingLabelType) SetIpAddressTagId(v int32)`

SetIpAddressTagId sets IpAddressTagId field to given value.

### HasIpAddressTagId

`func (o *ClientServingLabelType) HasIpAddressTagId() bool`

HasIpAddressTagId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


