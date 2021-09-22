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

// AccessMap2 Access Map create schema.
type AccessMap2 struct {
	// A unique name for the entity.
	Name string `json:"name"`
	// A description for the entity.
	Description *string `json:"description,omitempty"`
}

// NewAccessMap2 instantiates a new AccessMap2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessMap2(name string) *AccessMap2 {
	this := AccessMap2{}
	this.Name = name
	return &this
}

// NewAccessMap2WithDefaults instantiates a new AccessMap2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessMap2WithDefaults() *AccessMap2 {
	this := AccessMap2{}
	return &this
}

// GetName returns the Name field value
func (o *AccessMap2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccessMap2) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccessMap2) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AccessMap2) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessMap2) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AccessMap2) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AccessMap2) SetDescription(v string) {
	o.Description = &v
}

func (o AccessMap2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableAccessMap2 struct {
	value *AccessMap2
	isSet bool
}

func (v NullableAccessMap2) Get() *AccessMap2 {
	return v.value
}

func (v *NullableAccessMap2) Set(val *AccessMap2) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessMap2) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessMap2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessMap2(val *AccessMap2) *NullableAccessMap2 {
	return &NullableAccessMap2{value: val, isSet: true}
}

func (v NullableAccessMap2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessMap2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


