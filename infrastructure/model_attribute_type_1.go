/*
infrastructure

Infrastructure API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infrastructure

import (
	"encoding/json"
)

// AttributeType1 Attribute type update schema.
type AttributeType1 struct {
	// The unique identifier for an attribute type.
	AttributeTypeId int32 `json:"attributeTypeId"`
	// The unique name for an attribute type.
	Name string `json:"name"`
}

// NewAttributeType1 instantiates a new AttributeType1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributeType1(attributeTypeId int32, name string) *AttributeType1 {
	this := AttributeType1{}
	this.AttributeTypeId = attributeTypeId
	this.Name = name
	return &this
}

// NewAttributeType1WithDefaults instantiates a new AttributeType1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeType1WithDefaults() *AttributeType1 {
	this := AttributeType1{}
	return &this
}

// GetAttributeTypeId returns the AttributeTypeId field value
func (o *AttributeType1) GetAttributeTypeId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AttributeTypeId
}

// GetAttributeTypeIdOk returns a tuple with the AttributeTypeId field value
// and a boolean to check if the value has been set.
func (o *AttributeType1) GetAttributeTypeIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AttributeTypeId, true
}

// SetAttributeTypeId sets field value
func (o *AttributeType1) SetAttributeTypeId(v int32) {
	o.AttributeTypeId = v
}

// GetName returns the Name field value
func (o *AttributeType1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AttributeType1) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AttributeType1) SetName(v string) {
	o.Name = v
}

func (o AttributeType1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attributeTypeId"] = o.AttributeTypeId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableAttributeType1 struct {
	value *AttributeType1
	isSet bool
}

func (v NullableAttributeType1) Get() *AttributeType1 {
	return v.value
}

func (v *NullableAttributeType1) Set(val *AttributeType1) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeType1) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeType1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeType1(val *AttributeType1) *NullableAttributeType1 {
	return &NullableAttributeType1{value: val, isSet: true}
}

func (v NullableAttributeType1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeType1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


