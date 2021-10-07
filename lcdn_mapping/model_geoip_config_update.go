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

// GeoipConfigUpdate GeoIP config update schema.
type GeoipConfigUpdate struct {
	// The unique identifier for a GeoIP configuration.
	GeoIpConfigId int32 `json:"geoIpConfigId"`
	Type TypeType `json:"type"`
	// A unique name for the entity.
	Name string `json:"name"`
	// A description for the entity.
	Description *string `json:"description,omitempty"`
}

// NewGeoipConfigUpdate instantiates a new GeoipConfigUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoipConfigUpdate(geoIpConfigId int32, type_ TypeType, name string) *GeoipConfigUpdate {
	this := GeoipConfigUpdate{}
	this.GeoIpConfigId = geoIpConfigId
	this.Type = type_
	this.Name = name
	return &this
}

// NewGeoipConfigUpdateWithDefaults instantiates a new GeoipConfigUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoipConfigUpdateWithDefaults() *GeoipConfigUpdate {
	this := GeoipConfigUpdate{}
	return &this
}

// GetGeoIpConfigId returns the GeoIpConfigId field value
func (o *GeoipConfigUpdate) GetGeoIpConfigId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GeoIpConfigId
}

// GetGeoIpConfigIdOk returns a tuple with the GeoIpConfigId field value
// and a boolean to check if the value has been set.
func (o *GeoipConfigUpdate) GetGeoIpConfigIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GeoIpConfigId, true
}

// SetGeoIpConfigId sets field value
func (o *GeoipConfigUpdate) SetGeoIpConfigId(v int32) {
	o.GeoIpConfigId = v
}

// GetType returns the Type field value
func (o *GeoipConfigUpdate) GetType() TypeType {
	if o == nil {
		var ret TypeType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GeoipConfigUpdate) GetTypeOk() (*TypeType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GeoipConfigUpdate) SetType(v TypeType) {
	o.Type = v
}

// GetName returns the Name field value
func (o *GeoipConfigUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GeoipConfigUpdate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GeoipConfigUpdate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GeoipConfigUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoipConfigUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GeoipConfigUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GeoipConfigUpdate) SetDescription(v string) {
	o.Description = &v
}

func (o GeoipConfigUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["geoIpConfigId"] = o.GeoIpConfigId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableGeoipConfigUpdate struct {
	value *GeoipConfigUpdate
	isSet bool
}

func (v NullableGeoipConfigUpdate) Get() *GeoipConfigUpdate {
	return v.value
}

func (v *NullableGeoipConfigUpdate) Set(val *GeoipConfigUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoipConfigUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoipConfigUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoipConfigUpdate(val *GeoipConfigUpdate) *NullableGeoipConfigUpdate {
	return &NullableGeoipConfigUpdate{value: val, isSet: true}
}

func (v NullableGeoipConfigUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoipConfigUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


