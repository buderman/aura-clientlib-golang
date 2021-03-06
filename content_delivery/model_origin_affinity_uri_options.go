/*
content-delivery

Aura LCDN Content Delivery API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package content_delivery

import (
	"encoding/json"
)

// OriginAffinityUriOptions This behavior's set of configuration options.
type OriginAffinityUriOptions struct {
	// A URI prefix pattern used to identify client requests to which the HPC applies origin affinity.
	Value string `json:"value"`
}

// NewOriginAffinityUriOptions instantiates a new OriginAffinityUriOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginAffinityUriOptions(value string) *OriginAffinityUriOptions {
	this := OriginAffinityUriOptions{}
	this.Value = value
	return &this
}

// NewOriginAffinityUriOptionsWithDefaults instantiates a new OriginAffinityUriOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginAffinityUriOptionsWithDefaults() *OriginAffinityUriOptions {
	this := OriginAffinityUriOptions{}
	return &this
}

// GetValue returns the Value field value
func (o *OriginAffinityUriOptions) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *OriginAffinityUriOptions) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *OriginAffinityUriOptions) SetValue(v string) {
	o.Value = v
}

func (o OriginAffinityUriOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableOriginAffinityUriOptions struct {
	value *OriginAffinityUriOptions
	isSet bool
}

func (v NullableOriginAffinityUriOptions) Get() *OriginAffinityUriOptions {
	return v.value
}

func (v *NullableOriginAffinityUriOptions) Set(val *OriginAffinityUriOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginAffinityUriOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginAffinityUriOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginAffinityUriOptions(val *OriginAffinityUriOptions) *NullableOriginAffinityUriOptions {
	return &NullableOriginAffinityUriOptions{value: val, isSet: true}
}

func (v NullableOriginAffinityUriOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginAffinityUriOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


