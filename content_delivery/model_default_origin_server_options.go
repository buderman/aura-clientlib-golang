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

// DefaultOriginServerOptions This behavior's set of configuration options.
type DefaultOriginServerOptions struct {
	OriginId *int32 `json:"originId,omitempty"`
}

// NewDefaultOriginServerOptions instantiates a new DefaultOriginServerOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultOriginServerOptions() *DefaultOriginServerOptions {
	this := DefaultOriginServerOptions{}
	return &this
}

// NewDefaultOriginServerOptionsWithDefaults instantiates a new DefaultOriginServerOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultOriginServerOptionsWithDefaults() *DefaultOriginServerOptions {
	this := DefaultOriginServerOptions{}
	return &this
}

// GetOriginId returns the OriginId field value if set, zero value otherwise.
func (o *DefaultOriginServerOptions) GetOriginId() int32 {
	if o == nil || o.OriginId == nil {
		var ret int32
		return ret
	}
	return *o.OriginId
}

// GetOriginIdOk returns a tuple with the OriginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultOriginServerOptions) GetOriginIdOk() (*int32, bool) {
	if o == nil || o.OriginId == nil {
		return nil, false
	}
	return o.OriginId, true
}

// HasOriginId returns a boolean if a field has been set.
func (o *DefaultOriginServerOptions) HasOriginId() bool {
	if o != nil && o.OriginId != nil {
		return true
	}

	return false
}

// SetOriginId gets a reference to the given int32 and assigns it to the OriginId field.
func (o *DefaultOriginServerOptions) SetOriginId(v int32) {
	o.OriginId = &v
}

func (o DefaultOriginServerOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OriginId != nil {
		toSerialize["originId"] = o.OriginId
	}
	return json.Marshal(toSerialize)
}

type NullableDefaultOriginServerOptions struct {
	value *DefaultOriginServerOptions
	isSet bool
}

func (v NullableDefaultOriginServerOptions) Get() *DefaultOriginServerOptions {
	return v.value
}

func (v *NullableDefaultOriginServerOptions) Set(val *DefaultOriginServerOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultOriginServerOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultOriginServerOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultOriginServerOptions(val *DefaultOriginServerOptions) *NullableDefaultOriginServerOptions {
	return &NullableDefaultOriginServerOptions{value: val, isSet: true}
}

func (v NullableDefaultOriginServerOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultOriginServerOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


