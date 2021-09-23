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

// CacheControlOverride Ignore `Cache-Control`, overriding `Cache-Control` at the HyperCache with the provided value when caching contents. This override value does not modify any received `Cache-Control` headers or insert any missing `Cache-Control` headers. `Origin` headers remain intact when sent to the client.
type CacheControlOverride struct {
	// The name of the behavior, `cacheControlOverride` in this case.
	Name string `json:"name"`
	Options CacheControlOverrideOptions `json:"options"`
}

// NewCacheControlOverride instantiates a new CacheControlOverride object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCacheControlOverride(name string, options CacheControlOverrideOptions) *CacheControlOverride {
	this := CacheControlOverride{}
	this.Name = name
	this.Options = options
	return &this
}

// NewCacheControlOverrideWithDefaults instantiates a new CacheControlOverride object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCacheControlOverrideWithDefaults() *CacheControlOverride {
	this := CacheControlOverride{}
	return &this
}

// GetName returns the Name field value
func (o *CacheControlOverride) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CacheControlOverride) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CacheControlOverride) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *CacheControlOverride) GetOptions() CacheControlOverrideOptions {
	if o == nil {
		var ret CacheControlOverrideOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *CacheControlOverride) GetOptionsOk() (*CacheControlOverrideOptions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *CacheControlOverride) SetOptions(v CacheControlOverrideOptions) {
	o.Options = v
}

func (o CacheControlOverride) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableCacheControlOverride struct {
	value *CacheControlOverride
	isSet bool
}

func (v NullableCacheControlOverride) Get() *CacheControlOverride {
	return v.value
}

func (v *NullableCacheControlOverride) Set(val *CacheControlOverride) {
	v.value = val
	v.isSet = true
}

func (v NullableCacheControlOverride) IsSet() bool {
	return v.isSet
}

func (v *NullableCacheControlOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCacheControlOverride(val *CacheControlOverride) *NullableCacheControlOverride {
	return &NullableCacheControlOverride{value: val, isSet: true}
}

func (v NullableCacheControlOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCacheControlOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


