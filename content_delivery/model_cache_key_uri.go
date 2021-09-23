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

// CacheKeyUri URI pattern used by HyperCache for generating a cache key.
type CacheKeyUri struct {
	// The name of the behavior, `cacheKeyUri` in this case.
	Name string `json:"name"`
	Options CacheKeyUriOptions `json:"options"`
}

// NewCacheKeyUri instantiates a new CacheKeyUri object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCacheKeyUri(name string, options CacheKeyUriOptions) *CacheKeyUri {
	this := CacheKeyUri{}
	this.Name = name
	this.Options = options
	return &this
}

// NewCacheKeyUriWithDefaults instantiates a new CacheKeyUri object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCacheKeyUriWithDefaults() *CacheKeyUri {
	this := CacheKeyUri{}
	return &this
}

// GetName returns the Name field value
func (o *CacheKeyUri) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CacheKeyUri) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CacheKeyUri) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *CacheKeyUri) GetOptions() CacheKeyUriOptions {
	if o == nil {
		var ret CacheKeyUriOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *CacheKeyUri) GetOptionsOk() (*CacheKeyUriOptions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *CacheKeyUri) SetOptions(v CacheKeyUriOptions) {
	o.Options = v
}

func (o CacheKeyUri) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableCacheKeyUri struct {
	value *CacheKeyUri
	isSet bool
}

func (v NullableCacheKeyUri) Get() *CacheKeyUri {
	return v.value
}

func (v *NullableCacheKeyUri) Set(val *CacheKeyUri) {
	v.value = val
	v.isSet = true
}

func (v NullableCacheKeyUri) IsSet() bool {
	return v.isSet
}

func (v *NullableCacheKeyUri) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCacheKeyUri(val *CacheKeyUri) *NullableCacheKeyUri {
	return &NullableCacheKeyUri{value: val, isSet: true}
}

func (v NullableCacheKeyUri) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCacheKeyUri) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


