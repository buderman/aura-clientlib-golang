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

// DefaultRamOnlyCaching Enables or disables caching of specific types of content to RAM only memory.
type DefaultRamOnlyCaching struct {
	// The name of the behavior, `ramOnlyCaching` in this case.
	Name string `json:"name"`
	Options DefaultRamOnlyCachingOptions `json:"options"`
}

// NewDefaultRamOnlyCaching instantiates a new DefaultRamOnlyCaching object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultRamOnlyCaching(name string, options DefaultRamOnlyCachingOptions) *DefaultRamOnlyCaching {
	this := DefaultRamOnlyCaching{}
	this.Name = name
	this.Options = options
	return &this
}

// NewDefaultRamOnlyCachingWithDefaults instantiates a new DefaultRamOnlyCaching object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultRamOnlyCachingWithDefaults() *DefaultRamOnlyCaching {
	this := DefaultRamOnlyCaching{}
	return &this
}

// GetName returns the Name field value
func (o *DefaultRamOnlyCaching) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DefaultRamOnlyCaching) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DefaultRamOnlyCaching) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *DefaultRamOnlyCaching) GetOptions() DefaultRamOnlyCachingOptions {
	if o == nil {
		var ret DefaultRamOnlyCachingOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *DefaultRamOnlyCaching) GetOptionsOk() (*DefaultRamOnlyCachingOptions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *DefaultRamOnlyCaching) SetOptions(v DefaultRamOnlyCachingOptions) {
	o.Options = v
}

func (o DefaultRamOnlyCaching) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableDefaultRamOnlyCaching struct {
	value *DefaultRamOnlyCaching
	isSet bool
}

func (v NullableDefaultRamOnlyCaching) Get() *DefaultRamOnlyCaching {
	return v.value
}

func (v *NullableDefaultRamOnlyCaching) Set(val *DefaultRamOnlyCaching) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultRamOnlyCaching) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultRamOnlyCaching) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultRamOnlyCaching(val *DefaultRamOnlyCaching) *NullableDefaultRamOnlyCaching {
	return &NullableDefaultRamOnlyCaching{value: val, isSet: true}
}

func (v NullableDefaultRamOnlyCaching) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultRamOnlyCaching) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


