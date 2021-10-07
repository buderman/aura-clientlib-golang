/*
lcdn-deployment

Aura LCDN Deployment API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lcdn_deployment

import (
	"encoding/json"
)

// ClientServingLabelType An IP address that you specify for serving client-facing traffic. You can optionally associate this IP address with an IpAddressTag.
type ClientServingLabelType struct {
	// The node's IP address to which the HyperCache instance is bound. Specify this IP address for serving client-facing traffic.
	IpAddress string `json:"ipAddress"`
	// The unique identifier for an IP address tag. Use the IP address tag to route traffic on a per-IP address basis and to provide service differentiation based on HyperCache IP addresses.
	IpAddressTagId *int32 `json:"ipAddressTagId,omitempty"`
}

// NewClientServingLabelType instantiates a new ClientServingLabelType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientServingLabelType(ipAddress string) *ClientServingLabelType {
	this := ClientServingLabelType{}
	this.IpAddress = ipAddress
	return &this
}

// NewClientServingLabelTypeWithDefaults instantiates a new ClientServingLabelType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientServingLabelTypeWithDefaults() *ClientServingLabelType {
	this := ClientServingLabelType{}
	return &this
}

// GetIpAddress returns the IpAddress field value
func (o *ClientServingLabelType) GetIpAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value
// and a boolean to check if the value has been set.
func (o *ClientServingLabelType) GetIpAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IpAddress, true
}

// SetIpAddress sets field value
func (o *ClientServingLabelType) SetIpAddress(v string) {
	o.IpAddress = v
}

// GetIpAddressTagId returns the IpAddressTagId field value if set, zero value otherwise.
func (o *ClientServingLabelType) GetIpAddressTagId() int32 {
	if o == nil || o.IpAddressTagId == nil {
		var ret int32
		return ret
	}
	return *o.IpAddressTagId
}

// GetIpAddressTagIdOk returns a tuple with the IpAddressTagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientServingLabelType) GetIpAddressTagIdOk() (*int32, bool) {
	if o == nil || o.IpAddressTagId == nil {
		return nil, false
	}
	return o.IpAddressTagId, true
}

// HasIpAddressTagId returns a boolean if a field has been set.
func (o *ClientServingLabelType) HasIpAddressTagId() bool {
	if o != nil && o.IpAddressTagId != nil {
		return true
	}

	return false
}

// SetIpAddressTagId gets a reference to the given int32 and assigns it to the IpAddressTagId field.
func (o *ClientServingLabelType) SetIpAddressTagId(v int32) {
	o.IpAddressTagId = &v
}

func (o ClientServingLabelType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if o.IpAddressTagId != nil {
		toSerialize["ipAddressTagId"] = o.IpAddressTagId
	}
	return json.Marshal(toSerialize)
}

type NullableClientServingLabelType struct {
	value *ClientServingLabelType
	isSet bool
}

func (v NullableClientServingLabelType) Get() *ClientServingLabelType {
	return v.value
}

func (v *NullableClientServingLabelType) Set(val *ClientServingLabelType) {
	v.value = val
	v.isSet = true
}

func (v NullableClientServingLabelType) IsSet() bool {
	return v.isSet
}

func (v *NullableClientServingLabelType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientServingLabelType(val *ClientServingLabelType) *NullableClientServingLabelType {
	return &NullableClientServingLabelType{value: val, isSet: true}
}

func (v NullableClientServingLabelType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientServingLabelType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


