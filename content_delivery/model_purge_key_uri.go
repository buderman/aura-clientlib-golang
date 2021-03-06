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

// PurgeKeyUri URI pattern used to identify assets to be purged using the Purge API. By default this value is the same as cache key URI.
type PurgeKeyUri struct {
	// The name of the behavior, `purgeKeyUri` in this case.
	Name string `json:"name"`
	Options PurgeKeyUriOptions `json:"options"`
}

// NewPurgeKeyUri instantiates a new PurgeKeyUri object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurgeKeyUri(name string, options PurgeKeyUriOptions) *PurgeKeyUri {
	this := PurgeKeyUri{}
	this.Name = name
	this.Options = options
	return &this
}

// NewPurgeKeyUriWithDefaults instantiates a new PurgeKeyUri object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurgeKeyUriWithDefaults() *PurgeKeyUri {
	this := PurgeKeyUri{}
	return &this
}

// GetName returns the Name field value
func (o *PurgeKeyUri) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PurgeKeyUri) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PurgeKeyUri) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *PurgeKeyUri) GetOptions() PurgeKeyUriOptions {
	if o == nil {
		var ret PurgeKeyUriOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *PurgeKeyUri) GetOptionsOk() (*PurgeKeyUriOptions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *PurgeKeyUri) SetOptions(v PurgeKeyUriOptions) {
	o.Options = v
}

func (o PurgeKeyUri) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullablePurgeKeyUri struct {
	value *PurgeKeyUri
	isSet bool
}

func (v NullablePurgeKeyUri) Get() *PurgeKeyUri {
	return v.value
}

func (v *NullablePurgeKeyUri) Set(val *PurgeKeyUri) {
	v.value = val
	v.isSet = true
}

func (v NullablePurgeKeyUri) IsSet() bool {
	return v.isSet
}

func (v *NullablePurgeKeyUri) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurgeKeyUri(val *PurgeKeyUri) *NullablePurgeKeyUri {
	return &NullablePurgeKeyUri{value: val, isSet: true}
}

func (v NullablePurgeKeyUri) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurgeKeyUri) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


