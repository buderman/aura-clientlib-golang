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

// AttributeType Attribute types collection.
type AttributeType struct {
	AttributeTypes []AttributeType3 `json:"attributeTypes"`
	Page PageType `json:"page"`
}

// NewAttributeType instantiates a new AttributeType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributeType(attributeTypes []AttributeType3, page PageType) *AttributeType {
	this := AttributeType{}
	this.AttributeTypes = attributeTypes
	this.Page = page
	return &this
}

// NewAttributeTypeWithDefaults instantiates a new AttributeType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeTypeWithDefaults() *AttributeType {
	this := AttributeType{}
	return &this
}

// GetAttributeTypes returns the AttributeTypes field value
func (o *AttributeType) GetAttributeTypes() []AttributeType3 {
	if o == nil {
		var ret []AttributeType3
		return ret
	}

	return o.AttributeTypes
}

// GetAttributeTypesOk returns a tuple with the AttributeTypes field value
// and a boolean to check if the value has been set.
func (o *AttributeType) GetAttributeTypesOk() (*[]AttributeType3, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AttributeTypes, true
}

// SetAttributeTypes sets field value
func (o *AttributeType) SetAttributeTypes(v []AttributeType3) {
	o.AttributeTypes = v
}

// GetPage returns the Page field value
func (o *AttributeType) GetPage() PageType {
	if o == nil {
		var ret PageType
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *AttributeType) GetPageOk() (*PageType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *AttributeType) SetPage(v PageType) {
	o.Page = v
}

func (o AttributeType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attributeTypes"] = o.AttributeTypes
	}
	if true {
		toSerialize["page"] = o.Page
	}
	return json.Marshal(toSerialize)
}

type NullableAttributeType struct {
	value *AttributeType
	isSet bool
}

func (v NullableAttributeType) Get() *AttributeType {
	return v.value
}

func (v *NullableAttributeType) Set(val *AttributeType) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeType) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeType(val *AttributeType) *NullableAttributeType {
	return &NullableAttributeType{value: val, isSet: true}
}

func (v NullableAttributeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


