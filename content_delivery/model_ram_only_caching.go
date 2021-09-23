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

// RamOnlyCaching Enables or disables caching of specific types of content to RAM only memory.
type RamOnlyCaching struct {
	// The name of the behavior, `ramOnlyCaching` in this case.
	Name string `json:"name"`
	Options DefaultRamOnlyCachingOptions `json:"options"`
}

// NewRamOnlyCaching instantiates a new RamOnlyCaching object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRamOnlyCaching(name string, options DefaultRamOnlyCachingOptions) *RamOnlyCaching {
	this := RamOnlyCaching{}
	this.Name = name
	this.Options = options
	return &this
}

// NewRamOnlyCachingWithDefaults instantiates a new RamOnlyCaching object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRamOnlyCachingWithDefaults() *RamOnlyCaching {
	this := RamOnlyCaching{}
	return &this
}

// GetName returns the Name field value
func (o *RamOnlyCaching) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RamOnlyCaching) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RamOnlyCaching) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *RamOnlyCaching) GetOptions() DefaultRamOnlyCachingOptions {
	if o == nil {
		var ret DefaultRamOnlyCachingOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *RamOnlyCaching) GetOptionsOk() (*DefaultRamOnlyCachingOptions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *RamOnlyCaching) SetOptions(v DefaultRamOnlyCachingOptions) {
	o.Options = v
}

func (o RamOnlyCaching) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableRamOnlyCaching struct {
	value *RamOnlyCaching
	isSet bool
}

func (v NullableRamOnlyCaching) Get() *RamOnlyCaching {
	return v.value
}

func (v *NullableRamOnlyCaching) Set(val *RamOnlyCaching) {
	v.value = val
	v.isSet = true
}

func (v NullableRamOnlyCaching) IsSet() bool {
	return v.isSet
}

func (v *NullableRamOnlyCaching) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRamOnlyCaching(val *RamOnlyCaching) *NullableRamOnlyCaching {
	return &NullableRamOnlyCaching{value: val, isSet: true}
}

func (v NullableRamOnlyCaching) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRamOnlyCaching) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

