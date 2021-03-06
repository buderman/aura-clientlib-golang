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

// GeoipConfigCreate GeoIP config create schema.
type GeoipConfigCreate struct {
	Type TypeType `json:"type"`
	// A unique name for the entity.
	Name string `json:"name"`
	// A description for the entity.
	Description *string `json:"description,omitempty"`
}

// NewGeoipConfigCreate instantiates a new GeoipConfigCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoipConfigCreate(type_ TypeType, name string) *GeoipConfigCreate {
	this := GeoipConfigCreate{}
	this.Type = type_
	this.Name = name
	return &this
}

// NewGeoipConfigCreateWithDefaults instantiates a new GeoipConfigCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoipConfigCreateWithDefaults() *GeoipConfigCreate {
	this := GeoipConfigCreate{}
	return &this
}

// GetType returns the Type field value
func (o *GeoipConfigCreate) GetType() TypeType {
	if o == nil {
		var ret TypeType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GeoipConfigCreate) GetTypeOk() (*TypeType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GeoipConfigCreate) SetType(v TypeType) {
	o.Type = v
}

// GetName returns the Name field value
func (o *GeoipConfigCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GeoipConfigCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GeoipConfigCreate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GeoipConfigCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoipConfigCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GeoipConfigCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GeoipConfigCreate) SetDescription(v string) {
	o.Description = &v
}

func (o GeoipConfigCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableGeoipConfigCreate struct {
	value *GeoipConfigCreate
	isSet bool
}

func (v NullableGeoipConfigCreate) Get() *GeoipConfigCreate {
	return v.value
}

func (v *NullableGeoipConfigCreate) Set(val *GeoipConfigCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoipConfigCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoipConfigCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoipConfigCreate(val *GeoipConfigCreate) *NullableGeoipConfigCreate {
	return &NullableGeoipConfigCreate{value: val, isSet: true}
}

func (v NullableGeoipConfigCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoipConfigCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


