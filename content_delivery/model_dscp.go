/*
content-delivery

Cotent Delivery API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package content_delivery

import (
	"encoding/json"
)

// Dscp DiffServ byte field in the outgoing IP packets, typically used to prioritize, or generally schedule, traffic in an IP routed network.
type Dscp struct {
	// The name of the behavior, `dscp` in this case.
	Name string `json:"name"`
	Options DscpOptions `json:"options"`
}

// NewDscp instantiates a new Dscp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDscp(name string, options DscpOptions) *Dscp {
	this := Dscp{}
	this.Name = name
	this.Options = options
	return &this
}

// NewDscpWithDefaults instantiates a new Dscp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDscpWithDefaults() *Dscp {
	this := Dscp{}
	return &this
}

// GetName returns the Name field value
func (o *Dscp) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Dscp) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Dscp) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *Dscp) GetOptions() DscpOptions {
	if o == nil {
		var ret DscpOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *Dscp) GetOptionsOk() (*DscpOptions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *Dscp) SetOptions(v DscpOptions) {
	o.Options = v
}

func (o Dscp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableDscp struct {
	value *Dscp
	isSet bool
}

func (v NullableDscp) Get() *Dscp {
	return v.value
}

func (v *NullableDscp) Set(val *Dscp) {
	v.value = val
	v.isSet = true
}

func (v NullableDscp) IsSet() bool {
	return v.isSet
}

func (v *NullableDscp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDscp(val *Dscp) *NullableDscp {
	return &NullableDscp{value: val, isSet: true}
}

func (v NullableDscp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDscp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

