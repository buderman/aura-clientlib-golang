/*
lcdn-mapping

Aura LCDN Mapping API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lcdn_mapping

import (
	"encoding/json"
)

// SiteMap2 Site map summary read schema.
type SiteMap2 struct {
	// The unique identifier for an entity.
	MapId int32 `json:"mapId"`
	// A unique name for the entity.
	Name string `json:"name"`
}

// NewSiteMap2 instantiates a new SiteMap2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteMap2(mapId int32, name string) *SiteMap2 {
	this := SiteMap2{}
	this.MapId = mapId
	this.Name = name
	return &this
}

// NewSiteMap2WithDefaults instantiates a new SiteMap2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteMap2WithDefaults() *SiteMap2 {
	this := SiteMap2{}
	return &this
}

// GetMapId returns the MapId field value
func (o *SiteMap2) GetMapId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MapId
}

// GetMapIdOk returns a tuple with the MapId field value
// and a boolean to check if the value has been set.
func (o *SiteMap2) GetMapIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MapId, true
}

// SetMapId sets field value
func (o *SiteMap2) SetMapId(v int32) {
	o.MapId = v
}

// GetName returns the Name field value
func (o *SiteMap2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SiteMap2) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SiteMap2) SetName(v string) {
	o.Name = v
}

func (o SiteMap2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mapId"] = o.MapId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableSiteMap2 struct {
	value *SiteMap2
	isSet bool
}

func (v NullableSiteMap2) Get() *SiteMap2 {
	return v.value
}

func (v *NullableSiteMap2) Set(val *SiteMap2) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteMap2) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteMap2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteMap2(val *SiteMap2) *NullableSiteMap2 {
	return &NullableSiteMap2{value: val, isSet: true}
}

func (v NullableSiteMap2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteMap2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


