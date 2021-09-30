# BondInterfaceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MacAddress** | Pointer to **string** | Interface MAC address. | [optional] 
**Name** | **string** | Interface name, such as &#x60;lag0&#x60;.  Service nodes may use lag0 - lag3 for service bonding interfaces.  For management, a single bonding interface with the name lag_mgmt0 may be used.  Once created, this interface may not be deleted. | 
**Ip4Routes** | Pointer to [**[]Route4Type**](Route4Type.md) |  | [optional] 
**Ip6Routes** | Pointer to [**[]Route6Type**](Route6Type.md) |  | [optional] 
**SlaveInterfaces** | Pointer to **[]string** | List of slave interfaces for this bond. | [optional] 
**Ip4Addresses** | Pointer to [**[]Ip4AddressType**](Ip4AddressType.md) |  | [optional] 
**Type** | **string** | Identifies the interface type. The only supported type is &#x60;BOND&#x60;. | 
**Ip6Addresses** | Pointer to [**[]Ip6AddressType**](Ip6AddressType.md) |  | [optional] 

## Methods

### NewBondInterfaceType

`func NewBondInterfaceType(name string, type_ string, ) *BondInterfaceType`

NewBondInterfaceType instantiates a new BondInterfaceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBondInterfaceTypeWithDefaults

`func NewBondInterfaceTypeWithDefaults() *BondInterfaceType`

NewBondInterfaceTypeWithDefaults instantiates a new BondInterfaceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacAddress

`func (o *BondInterfaceType) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *BondInterfaceType) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *BondInterfaceType) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *BondInterfaceType) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetName

`func (o *BondInterfaceType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BondInterfaceType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BondInterfaceType) SetName(v string)`

SetName sets Name field to given value.


### GetIp4Routes

`func (o *BondInterfaceType) GetIp4Routes() []Route4Type`

GetIp4Routes returns the Ip4Routes field if non-nil, zero value otherwise.

### GetIp4RoutesOk

`func (o *BondInterfaceType) GetIp4RoutesOk() (*[]Route4Type, bool)`

GetIp4RoutesOk returns a tuple with the Ip4Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp4Routes

`func (o *BondInterfaceType) SetIp4Routes(v []Route4Type)`

SetIp4Routes sets Ip4Routes field to given value.

### HasIp4Routes

`func (o *BondInterfaceType) HasIp4Routes() bool`

HasIp4Routes returns a boolean if a field has been set.

### GetIp6Routes

`func (o *BondInterfaceType) GetIp6Routes() []Route6Type`

GetIp6Routes returns the Ip6Routes field if non-nil, zero value otherwise.

### GetIp6RoutesOk

`func (o *BondInterfaceType) GetIp6RoutesOk() (*[]Route6Type, bool)`

GetIp6RoutesOk returns a tuple with the Ip6Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp6Routes

`func (o *BondInterfaceType) SetIp6Routes(v []Route6Type)`

SetIp6Routes sets Ip6Routes field to given value.

### HasIp6Routes

`func (o *BondInterfaceType) HasIp6Routes() bool`

HasIp6Routes returns a boolean if a field has been set.

### GetSlaveInterfaces

`func (o *BondInterfaceType) GetSlaveInterfaces() []string`

GetSlaveInterfaces returns the SlaveInterfaces field if non-nil, zero value otherwise.

### GetSlaveInterfacesOk

`func (o *BondInterfaceType) GetSlaveInterfacesOk() (*[]string, bool)`

GetSlaveInterfacesOk returns a tuple with the SlaveInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaveInterfaces

`func (o *BondInterfaceType) SetSlaveInterfaces(v []string)`

SetSlaveInterfaces sets SlaveInterfaces field to given value.

### HasSlaveInterfaces

`func (o *BondInterfaceType) HasSlaveInterfaces() bool`

HasSlaveInterfaces returns a boolean if a field has been set.

### GetIp4Addresses

`func (o *BondInterfaceType) GetIp4Addresses() []Ip4AddressType`

GetIp4Addresses returns the Ip4Addresses field if non-nil, zero value otherwise.

### GetIp4AddressesOk

`func (o *BondInterfaceType) GetIp4AddressesOk() (*[]Ip4AddressType, bool)`

GetIp4AddressesOk returns a tuple with the Ip4Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp4Addresses

`func (o *BondInterfaceType) SetIp4Addresses(v []Ip4AddressType)`

SetIp4Addresses sets Ip4Addresses field to given value.

### HasIp4Addresses

`func (o *BondInterfaceType) HasIp4Addresses() bool`

HasIp4Addresses returns a boolean if a field has been set.

### GetType

`func (o *BondInterfaceType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BondInterfaceType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BondInterfaceType) SetType(v string)`

SetType sets Type field to given value.


### GetIp6Addresses

`func (o *BondInterfaceType) GetIp6Addresses() []Ip6AddressType`

GetIp6Addresses returns the Ip6Addresses field if non-nil, zero value otherwise.

### GetIp6AddressesOk

`func (o *BondInterfaceType) GetIp6AddressesOk() (*[]Ip6AddressType, bool)`

GetIp6AddressesOk returns a tuple with the Ip6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp6Addresses

`func (o *BondInterfaceType) SetIp6Addresses(v []Ip6AddressType)`

SetIp6Addresses sets Ip6Addresses field to given value.

### HasIp6Addresses

`func (o *BondInterfaceType) HasIp6Addresses() bool`

HasIp6Addresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


