# Ip4AddressType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Management** | Pointer to **bool** | Indicates whether this is the management address.  There may only be one management address. | [optional] [default to false]
**IpAddress** | **string** | An IPv4 address. | 

## Methods

### NewIp4AddressType

`func NewIp4AddressType(ipAddress string, ) *Ip4AddressType`

NewIp4AddressType instantiates a new Ip4AddressType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIp4AddressTypeWithDefaults

`func NewIp4AddressTypeWithDefaults() *Ip4AddressType`

NewIp4AddressTypeWithDefaults instantiates a new Ip4AddressType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagement

`func (o *Ip4AddressType) GetManagement() bool`

GetManagement returns the Management field if non-nil, zero value otherwise.

### GetManagementOk

`func (o *Ip4AddressType) GetManagementOk() (*bool, bool)`

GetManagementOk returns a tuple with the Management field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagement

`func (o *Ip4AddressType) SetManagement(v bool)`

SetManagement sets Management field to given value.

### HasManagement

`func (o *Ip4AddressType) HasManagement() bool`

HasManagement returns a boolean if a field has been set.

### GetIpAddress

`func (o *Ip4AddressType) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Ip4AddressType) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Ip4AddressType) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


