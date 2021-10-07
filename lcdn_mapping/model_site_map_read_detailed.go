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

// SiteMapReadDetailed Site map object descriptions.
type SiteMapReadDetailed struct {
	// The unique identifier for an entity.
	MapId int32 `json:"mapId"`
	// A unique name for the entity.
	Name string `json:"name"`
	// A description for the entity.
	Description *string `json:"description,omitempty"`
}

// NewSiteMapReadDetailed instantiates a new SiteMapReadDetailed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteMapReadDetailed(mapId int32, name string) *SiteMapReadDetailed {
	this := SiteMapReadDetailed{}
	this.MapId = mapId
	this.Name = name
	return &this
}

// NewSiteMapReadDetailedWithDefaults instantiates a new SiteMapReadDetailed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteMapReadDetailedWithDefaults() *SiteMapReadDetailed {
	this := SiteMapReadDetailed{}
	return &this
}

// GetMapId returns the MapId field value
func (o *SiteMapReadDetailed) GetMapId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MapId
}

// GetMapIdOk returns a tuple with the MapId field value
// and a boolean to check if the value has been set.
func (o *SiteMapReadDetailed) GetMapIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MapId, true
}

// SetMapId sets field value
func (o *SiteMapReadDetailed) SetMapId(v int32) {
	o.MapId = v
}

// GetName returns the Name field value
func (o *SiteMapReadDetailed) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SiteMapReadDetailed) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SiteMapReadDetailed) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SiteMapReadDetailed) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteMapReadDetailed) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SiteMapReadDetailed) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SiteMapReadDetailed) SetDescription(v string) {
	o.Description = &v
}

func (o SiteMapReadDetailed) MarshalJSON() ([]byte, error) {
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

type NullableSiteMapReadDetailed struct {
	value *SiteMapReadDetailed
	isSet bool
}

func (v NullableSiteMapReadDetailed) Get() *SiteMapReadDetailed {
	return v.value
}

func (v *NullableSiteMapReadDetailed) Set(val *SiteMapReadDetailed) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteMapReadDetailed) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteMapReadDetailed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteMapReadDetailed(val *SiteMapReadDetailed) *NullableSiteMapReadDetailed {
	return &NullableSiteMapReadDetailed{value: val, isSet: true}
}

func (v NullableSiteMapReadDetailed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteMapReadDetailed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


