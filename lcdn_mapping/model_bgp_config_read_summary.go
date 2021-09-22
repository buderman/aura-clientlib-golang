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

// BgpConfigReadSummary BGP config summary read schema.
type BgpConfigReadSummary struct {
	// The unique identifier for a BGP configuration.
	BgpConfigId int32 `json:"bgpConfigId"`
	// A unique name for the entity.
	Name string `json:"name"`
}

// NewBgpConfigReadSummary instantiates a new BgpConfigReadSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgpConfigReadSummary(bgpConfigId int32, name string) *BgpConfigReadSummary {
	this := BgpConfigReadSummary{}
	this.BgpConfigId = bgpConfigId
	this.Name = name
	return &this
}

// NewBgpConfigReadSummaryWithDefaults instantiates a new BgpConfigReadSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgpConfigReadSummaryWithDefaults() *BgpConfigReadSummary {
	this := BgpConfigReadSummary{}
	return &this
}

// GetBgpConfigId returns the BgpConfigId field value
func (o *BgpConfigReadSummary) GetBgpConfigId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BgpConfigId
}

// GetBgpConfigIdOk returns a tuple with the BgpConfigId field value
// and a boolean to check if the value has been set.
func (o *BgpConfigReadSummary) GetBgpConfigIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BgpConfigId, true
}

// SetBgpConfigId sets field value
func (o *BgpConfigReadSummary) SetBgpConfigId(v int32) {
	o.BgpConfigId = v
}

// GetName returns the Name field value
func (o *BgpConfigReadSummary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BgpConfigReadSummary) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BgpConfigReadSummary) SetName(v string) {
	o.Name = v
}

func (o BgpConfigReadSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bgpConfigId"] = o.BgpConfigId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableBgpConfigReadSummary struct {
	value *BgpConfigReadSummary
	isSet bool
}

func (v NullableBgpConfigReadSummary) Get() *BgpConfigReadSummary {
	return v.value
}

func (v *NullableBgpConfigReadSummary) Set(val *BgpConfigReadSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpConfigReadSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpConfigReadSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpConfigReadSummary(val *BgpConfigReadSummary) *NullableBgpConfigReadSummary {
	return &NullableBgpConfigReadSummary{value: val, isSet: true}
}

func (v NullableBgpConfigReadSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpConfigReadSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

