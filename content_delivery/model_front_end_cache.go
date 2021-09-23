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

// FrontEndCache Enable or disable caching of specific types of content to front-end memory. This stores entire objects in HPC nodes, not chunks.
type FrontEndCache struct {
	// The name of the behavior, `frontEndCache` in this case.
	Name string `json:"name"`
	Options DefaultFrontEndCacheOptions `json:"options"`
}

// NewFrontEndCache instantiates a new FrontEndCache object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrontEndCache(name string, options DefaultFrontEndCacheOptions) *FrontEndCache {
	this := FrontEndCache{}
	this.Name = name
	this.Options = options
	return &this
}

// NewFrontEndCacheWithDefaults instantiates a new FrontEndCache object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrontEndCacheWithDefaults() *FrontEndCache {
	this := FrontEndCache{}
	return &this
}

// GetName returns the Name field value
func (o *FrontEndCache) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FrontEndCache) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FrontEndCache) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *FrontEndCache) GetOptions() DefaultFrontEndCacheOptions {
	if o == nil {
		var ret DefaultFrontEndCacheOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *FrontEndCache) GetOptionsOk() (*DefaultFrontEndCacheOptions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *FrontEndCache) SetOptions(v DefaultFrontEndCacheOptions) {
	o.Options = v
}

func (o FrontEndCache) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableFrontEndCache struct {
	value *FrontEndCache
	isSet bool
}

func (v NullableFrontEndCache) Get() *FrontEndCache {
	return v.value
}

func (v *NullableFrontEndCache) Set(val *FrontEndCache) {
	v.value = val
	v.isSet = true
}

func (v NullableFrontEndCache) IsSet() bool {
	return v.isSet
}

func (v *NullableFrontEndCache) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrontEndCache(val *FrontEndCache) *NullableFrontEndCache {
	return &NullableFrontEndCache{value: val, isSet: true}
}

func (v NullableFrontEndCache) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrontEndCache) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


