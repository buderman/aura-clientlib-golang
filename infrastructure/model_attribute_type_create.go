/*
infrastructure

Aura Infrastructure API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infrastructure

import (
	"encoding/json"
)

// AttributeTypeCreate A CDN Operator has the option to create a global set of node attribute types that provide additional node information to CDN operators, IT, or support organizations.
type AttributeTypeCreate struct {
	// The unique name for an attribute type.
	Name string `json:"name"`
}

// NewAttributeTypeCreate instantiates a new AttributeTypeCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributeTypeCreate(name string) *AttributeTypeCreate {
	this := AttributeTypeCreate{}
	this.Name = name
	return &this
}

// NewAttributeTypeCreateWithDefaults instantiates a new AttributeTypeCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeTypeCreateWithDefaults() *AttributeTypeCreate {
	this := AttributeTypeCreate{}
	return &this
}

// GetName returns the Name field value
func (o *AttributeTypeCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AttributeTypeCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AttributeTypeCreate) SetName(v string) {
	o.Name = v
}

func (o AttributeTypeCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableAttributeTypeCreate struct {
	value *AttributeTypeCreate
	isSet bool
}

func (v NullableAttributeTypeCreate) Get() *AttributeTypeCreate {
	return v.value
}

func (v *NullableAttributeTypeCreate) Set(val *AttributeTypeCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeTypeCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeTypeCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeTypeCreate(val *AttributeTypeCreate) *NullableAttributeTypeCreate {
	return &NullableAttributeTypeCreate{value: val, isSet: true}
}

func (v NullableAttributeTypeCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeTypeCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


