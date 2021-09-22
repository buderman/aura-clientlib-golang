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

// AccessMapReadSummary Access Map summary read schema.
type AccessMapReadSummary struct {
	// The unique identifier for an entity.
	MapId int32 `json:"mapId"`
	// A unique name for the entity.
	Name string `json:"name"`
}

// NewAccessMapReadSummary instantiates a new AccessMapReadSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessMapReadSummary(mapId int32, name string) *AccessMapReadSummary {
	this := AccessMapReadSummary{}
	this.MapId = mapId
	this.Name = name
	return &this
}

// NewAccessMapReadSummaryWithDefaults instantiates a new AccessMapReadSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessMapReadSummaryWithDefaults() *AccessMapReadSummary {
	this := AccessMapReadSummary{}
	return &this
}

// GetMapId returns the MapId field value
func (o *AccessMapReadSummary) GetMapId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MapId
}

// GetMapIdOk returns a tuple with the MapId field value
// and a boolean to check if the value has been set.
func (o *AccessMapReadSummary) GetMapIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MapId, true
}

// SetMapId sets field value
func (o *AccessMapReadSummary) SetMapId(v int32) {
	o.MapId = v
}

// GetName returns the Name field value
func (o *AccessMapReadSummary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccessMapReadSummary) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccessMapReadSummary) SetName(v string) {
	o.Name = v
}

func (o AccessMapReadSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mapId"] = o.MapId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableAccessMapReadSummary struct {
	value *AccessMapReadSummary
	isSet bool
}

func (v NullableAccessMapReadSummary) Get() *AccessMapReadSummary {
	return v.value
}

func (v *NullableAccessMapReadSummary) Set(val *AccessMapReadSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessMapReadSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessMapReadSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessMapReadSummary(val *AccessMapReadSummary) *NullableAccessMapReadSummary {
	return &NullableAccessMapReadSummary{value: val, isSet: true}
}

func (v NullableAccessMapReadSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessMapReadSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


