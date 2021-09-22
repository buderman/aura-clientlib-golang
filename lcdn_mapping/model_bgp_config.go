/*
lcdn-mapping

LCDN Mapping API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lcdn_mapping

import (
	"encoding/json"
)

// BgpConfig BGP config update schema.
type BgpConfig struct {
	// The unique identifier for a BGP configuration.
	BgpConfigId int32 `json:"bgpConfigId"`
	// A unique name for the entity.
	Name string `json:"name"`
	// A description for the entity.
	Description *string `json:"description,omitempty"`
}

// NewBgpConfig instantiates a new BgpConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgpConfig(bgpConfigId int32, name string) *BgpConfig {
	this := BgpConfig{}
	this.BgpConfigId = bgpConfigId
	this.Name = name
	return &this
}

// NewBgpConfigWithDefaults instantiates a new BgpConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgpConfigWithDefaults() *BgpConfig {
	this := BgpConfig{}
	return &this
}

// GetBgpConfigId returns the BgpConfigId field value
func (o *BgpConfig) GetBgpConfigId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BgpConfigId
}

// GetBgpConfigIdOk returns a tuple with the BgpConfigId field value
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetBgpConfigIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BgpConfigId, true
}

// SetBgpConfigId sets field value
func (o *BgpConfig) SetBgpConfigId(v int32) {
	o.BgpConfigId = v
}

// GetName returns the Name field value
func (o *BgpConfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BgpConfig) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BgpConfig) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BgpConfig) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BgpConfig) SetDescription(v string) {
	o.Description = &v
}

func (o BgpConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bgpConfigId"] = o.BgpConfigId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableBgpConfig struct {
	value *BgpConfig
	isSet bool
}

func (v NullableBgpConfig) Get() *BgpConfig {
	return v.value
}

func (v *NullableBgpConfig) Set(val *BgpConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpConfig(val *BgpConfig) *NullableBgpConfig {
	return &NullableBgpConfig{value: val, isSet: true}
}

func (v NullableBgpConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

