/*
infrastructure

Infrastructure API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infrastructure

import (
	"encoding/json"
)

// EthernetInterfaceType A node ethernet interface.
type EthernetInterfaceType struct {
	// Interface MAC address.
	MacAddress string `json:"macAddress"`
	// Interface name, such as `eth0`.
	Name string `json:"name"`
	Ip4Routes *[]Route4Type `json:"ip4Routes,omitempty"`
	Ip6Routes *[]Route6Type `json:"ip6Routes,omitempty"`
	Ip4Addresses *[]Ip4AddressType `json:"ip4Addresses,omitempty"`
	// Designates the master bonding interface for this interface.
	Master *string `json:"master,omitempty"`
	// Identifies the interface type. The only supported type is `ETHERNET`.
	Type string `json:"type"`
	Ip6Addresses *[]Ip6AddressType `json:"ip6Addresses,omitempty"`
}

// NewEthernetInterfaceType instantiates a new EthernetInterfaceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEthernetInterfaceType(macAddress string, name string, type_ string) *EthernetInterfaceType {
	this := EthernetInterfaceType{}
	this.MacAddress = macAddress
	this.Name = name
	this.Type = type_
	return &this
}

// NewEthernetInterfaceTypeWithDefaults instantiates a new EthernetInterfaceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEthernetInterfaceTypeWithDefaults() *EthernetInterfaceType {
	this := EthernetInterfaceType{}
	return &this
}

// GetMacAddress returns the MacAddress field value
func (o *EthernetInterfaceType) GetMacAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value
// and a boolean to check if the value has been set.
func (o *EthernetInterfaceType) GetMacAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MacAddress, true
}

// SetMacAddress sets field value
func (o *EthernetInterfaceType) SetMacAddress(v string) {
	o.MacAddress = v
}

// GetName returns the Name field value
func (o *EthernetInterfaceType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EthernetInterfaceType) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EthernetInterfaceType) SetName(v string) {
	o.Name = v
}

// GetIp4Routes returns the Ip4Routes field value if set, zero value otherwise.
func (o *EthernetInterfaceType) GetIp4Routes() []Route4Type {
	if o == nil || o.Ip4Routes == nil {
		var ret []Route4Type
		return ret
	}
	return *o.Ip4Routes
}

// GetIp4RoutesOk returns a tuple with the Ip4Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetInterfaceType) GetIp4RoutesOk() (*[]Route4Type, bool) {
	if o == nil || o.Ip4Routes == nil {
		return nil, false
	}
	return o.Ip4Routes, true
}

// HasIp4Routes returns a boolean if a field has been set.
func (o *EthernetInterfaceType) HasIp4Routes() bool {
	if o != nil && o.Ip4Routes != nil {
		return true
	}

	return false
}

// SetIp4Routes gets a reference to the given []Route4Type and assigns it to the Ip4Routes field.
func (o *EthernetInterfaceType) SetIp4Routes(v []Route4Type) {
	o.Ip4Routes = &v
}

// GetIp6Routes returns the Ip6Routes field value if set, zero value otherwise.
func (o *EthernetInterfaceType) GetIp6Routes() []Route6Type {
	if o == nil || o.Ip6Routes == nil {
		var ret []Route6Type
		return ret
	}
	return *o.Ip6Routes
}

// GetIp6RoutesOk returns a tuple with the Ip6Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetInterfaceType) GetIp6RoutesOk() (*[]Route6Type, bool) {
	if o == nil || o.Ip6Routes == nil {
		return nil, false
	}
	return o.Ip6Routes, true
}

// HasIp6Routes returns a boolean if a field has been set.
func (o *EthernetInterfaceType) HasIp6Routes() bool {
	if o != nil && o.Ip6Routes != nil {
		return true
	}

	return false
}

// SetIp6Routes gets a reference to the given []Route6Type and assigns it to the Ip6Routes field.
func (o *EthernetInterfaceType) SetIp6Routes(v []Route6Type) {
	o.Ip6Routes = &v
}

// GetIp4Addresses returns the Ip4Addresses field value if set, zero value otherwise.
func (o *EthernetInterfaceType) GetIp4Addresses() []Ip4AddressType {
	if o == nil || o.Ip4Addresses == nil {
		var ret []Ip4AddressType
		return ret
	}
	return *o.Ip4Addresses
}

// GetIp4AddressesOk returns a tuple with the Ip4Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetInterfaceType) GetIp4AddressesOk() (*[]Ip4AddressType, bool) {
	if o == nil || o.Ip4Addresses == nil {
		return nil, false
	}
	return o.Ip4Addresses, true
}

// HasIp4Addresses returns a boolean if a field has been set.
func (o *EthernetInterfaceType) HasIp4Addresses() bool {
	if o != nil && o.Ip4Addresses != nil {
		return true
	}

	return false
}

// SetIp4Addresses gets a reference to the given []Ip4AddressType and assigns it to the Ip4Addresses field.
func (o *EthernetInterfaceType) SetIp4Addresses(v []Ip4AddressType) {
	o.Ip4Addresses = &v
}

// GetMaster returns the Master field value if set, zero value otherwise.
func (o *EthernetInterfaceType) GetMaster() string {
	if o == nil || o.Master == nil {
		var ret string
		return ret
	}
	return *o.Master
}

// GetMasterOk returns a tuple with the Master field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetInterfaceType) GetMasterOk() (*string, bool) {
	if o == nil || o.Master == nil {
		return nil, false
	}
	return o.Master, true
}

// HasMaster returns a boolean if a field has been set.
func (o *EthernetInterfaceType) HasMaster() bool {
	if o != nil && o.Master != nil {
		return true
	}

	return false
}

// SetMaster gets a reference to the given string and assigns it to the Master field.
func (o *EthernetInterfaceType) SetMaster(v string) {
	o.Master = &v
}

// GetType returns the Type field value
func (o *EthernetInterfaceType) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EthernetInterfaceType) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EthernetInterfaceType) SetType(v string) {
	o.Type = v
}

// GetIp6Addresses returns the Ip6Addresses field value if set, zero value otherwise.
func (o *EthernetInterfaceType) GetIp6Addresses() []Ip6AddressType {
	if o == nil || o.Ip6Addresses == nil {
		var ret []Ip6AddressType
		return ret
	}
	return *o.Ip6Addresses
}

// GetIp6AddressesOk returns a tuple with the Ip6Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetInterfaceType) GetIp6AddressesOk() (*[]Ip6AddressType, bool) {
	if o == nil || o.Ip6Addresses == nil {
		return nil, false
	}
	return o.Ip6Addresses, true
}

// HasIp6Addresses returns a boolean if a field has been set.
func (o *EthernetInterfaceType) HasIp6Addresses() bool {
	if o != nil && o.Ip6Addresses != nil {
		return true
	}

	return false
}

// SetIp6Addresses gets a reference to the given []Ip6AddressType and assigns it to the Ip6Addresses field.
func (o *EthernetInterfaceType) SetIp6Addresses(v []Ip6AddressType) {
	o.Ip6Addresses = &v
}

func (o EthernetInterfaceType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["macAddress"] = o.MacAddress
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Ip4Routes != nil {
		toSerialize["ip4Routes"] = o.Ip4Routes
	}
	if o.Ip6Routes != nil {
		toSerialize["ip6Routes"] = o.Ip6Routes
	}
	if o.Ip4Addresses != nil {
		toSerialize["ip4Addresses"] = o.Ip4Addresses
	}
	if o.Master != nil {
		toSerialize["master"] = o.Master
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Ip6Addresses != nil {
		toSerialize["ip6Addresses"] = o.Ip6Addresses
	}
	return json.Marshal(toSerialize)
}

type NullableEthernetInterfaceType struct {
	value *EthernetInterfaceType
	isSet bool
}

func (v NullableEthernetInterfaceType) Get() *EthernetInterfaceType {
	return v.value
}

func (v *NullableEthernetInterfaceType) Set(val *EthernetInterfaceType) {
	v.value = val
	v.isSet = true
}

func (v NullableEthernetInterfaceType) IsSet() bool {
	return v.isSet
}

func (v *NullableEthernetInterfaceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEthernetInterfaceType(val *EthernetInterfaceType) *NullableEthernetInterfaceType {
	return &NullableEthernetInterfaceType{value: val, isSet: true}
}

func (v NullableEthernetInterfaceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEthernetInterfaceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


