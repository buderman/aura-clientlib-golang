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

// SiteMap1 Site map object descriptions.
type SiteMap1 struct {
	// The unique identifier for an entity.
	MapId int32 `json:"mapId"`
	// A unique name for the entity.
	Name string `json:"name"`
	// A description for the entity.
	Description *string `json:"description,omitempty"`
}

// NewSiteMap1 instantiates a new SiteMap1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteMap1(mapId int32, name string) *SiteMap1 {
	this := SiteMap1{}
	this.MapId = mapId
	this.Name = name
	return &this
}

// NewSiteMap1WithDefaults instantiates a new SiteMap1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteMap1WithDefaults() *SiteMap1 {
	this := SiteMap1{}
	return &this
}

// GetMapId returns the MapId field value
func (o *SiteMap1) GetMapId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MapId
}

// GetMapIdOk returns a tuple with the MapId field value
// and a boolean to check if the value has been set.
func (o *SiteMap1) GetMapIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MapId, true
}

// SetMapId sets field value
func (o *SiteMap1) SetMapId(v int32) {
	o.MapId = v
}

// GetName returns the Name field value
func (o *SiteMap1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SiteMap1) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SiteMap1) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SiteMap1) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteMap1) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SiteMap1) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SiteMap1) SetDescription(v string) {
	o.Description = &v
}

func (o SiteMap1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mapId"] = o.MapId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableSiteMap1 struct {
	value *SiteMap1
	isSet bool
}

func (v NullableSiteMap1) Get() *SiteMap1 {
	return v.value
}

func (v *NullableSiteMap1) Set(val *SiteMap1) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteMap1) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteMap1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteMap1(val *SiteMap1) *NullableSiteMap1 {
	return &NullableSiteMap1{value: val, isSet: true}
}

func (v NullableSiteMap1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteMap1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


