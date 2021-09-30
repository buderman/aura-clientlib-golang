# EthernetInterfaceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MacAddress** | **string** | Interface MAC address. | 
**Name** | **string** | Interface name, such as &#x60;eth0&#x60;. | 
**Ip4Routes** | Pointer to [**[]Route4Type**](Route4Type.md) |  | [optional] 
**Ip6Routes** | Pointer to [**[]Route6Type**](Route6Type.md) |  | [optional] 
**Ip4Addresses** | Pointer to [**[]Ip4AddressType**](Ip4AddressType.md) |  | [optional] 
**Master** | Pointer to **string** | Designates the master bonding interface for this interface. | [optional] [readonly] 
**Type** | **string** | Identifies the interface type. The only supported type is &#x60;ETHERNET&#x60;. | 
**Ip6Addresses** | Pointer to [**[]Ip6AddressType**](Ip6AddressType.md) |  | [optional] 

## Methods

### NewEthernetInterfaceType

`func NewEthernetInterfaceType(macAddress string, name string, type_ string, ) *EthernetInterfaceType`

NewEthernetInterfaceType instantiates a new EthernetInterfaceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEthernetInterfaceTypeWithDefaults

`func NewEthernetInterfaceTypeWithDefaults() *EthernetInterfaceType`

NewEthernetInterfaceTypeWithDefaults instantiates a new EthernetInterfaceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacAddress

`func (o *EthernetInterfaceType) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *EthernetInterfaceType) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *EthernetInterfaceType) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.


### GetName

`func (o *EthernetInterfaceType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EthernetInterfaceType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EthernetInterfaceType) SetName(v string)`

SetName sets Name field to given value.


### GetIp4Routes

`func (o *EthernetInterfaceType) GetIp4Routes() []Route4Type`

GetIp4Routes returns the Ip4Routes field if non-nil, zero value otherwise.

### GetIp4RoutesOk

`func (o *EthernetInterfaceType) GetIp4RoutesOk() (*[]Route4Type, bool)`

GetIp4RoutesOk returns a tuple with the Ip4Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp4Routes

`func (o *EthernetInterfaceType) SetIp4Routes(v []Route4Type)`

SetIp4Routes sets Ip4Routes field to given value.

### HasIp4Routes

`func (o *EthernetInterfaceType) HasIp4Routes() bool`

HasIp4Routes returns a boolean if a field has been set.

### GetIp6Routes

`func (o *EthernetInterfaceType) GetIp6Routes() []Route6Type`

GetIp6Routes returns the Ip6Routes field if non-nil, zero value otherwise.

### GetIp6RoutesOk

`func (o *EthernetInterfaceType) GetIp6RoutesOk() (*[]Route6Type, bool)`

GetIp6RoutesOk returns a tuple with the Ip6Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp6Routes

`func (o *EthernetInterfaceType) SetIp6Routes(v []Route6Type)`

SetIp6Routes sets Ip6Routes field to given value.

### HasIp6Routes

`func (o *EthernetInterfaceType) HasIp6Routes() bool`

HasIp6Routes returns a boolean if a field has been set.

### GetIp4Addresses

`func (o *EthernetInterfaceType) GetIp4Addresses() []Ip4AddressType`

GetIp4Addresses returns the Ip4Addresses field if non-nil, zero value otherwise.

### GetIp4AddressesOk

`func (o *EthernetInterfaceType) GetIp4AddressesOk() (*[]Ip4AddressType, bool)`

GetIp4AddressesOk returns a tuple with the Ip4Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp4Addresses

`func (o *EthernetInterfaceType) SetIp4Addresses(v []Ip4AddressType)`

SetIp4Addresses sets Ip4Addresses field to given value.

### HasIp4Addresses

`func (o *EthernetInterfaceType) HasIp4Addresses() bool`

HasIp4Addresses returns a boolean if a field has been set.

### GetMaster

`func (o *EthernetInterfaceType) GetMaster() string`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *EthernetInterfaceType) GetMasterOk() (*string, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *EthernetInterfaceType) SetMaster(v string)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *EthernetInterfaceType) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetType

`func (o *EthernetInterfaceType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EthernetInterfaceType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EthernetInterfaceType) SetType(v string)`

SetType sets Type field to given value.


### GetIp6Addresses

`func (o *EthernetInterfaceType) GetIp6Addresses() []Ip6AddressType`

GetIp6Addresses returns the Ip6Addresses field if non-nil, zero value otherwise.

### GetIp6AddressesOk

`func (o *EthernetInterfaceType) GetIp6AddressesOk() (*[]Ip6AddressType, bool)`

GetIp6AddressesOk returns a tuple with the Ip6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp6Addresses

`func (o *EthernetInterfaceType) SetIp6Addresses(v []Ip6AddressType)`

SetIp6Addresses sets Ip6Addresses field to given value.

### HasIp6Addresses

`func (o *EthernetInterfaceType) HasIp6Addresses() bool`

HasIp6Addresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


