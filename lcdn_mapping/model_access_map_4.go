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

// AccessMap4 Access Map collection.
type AccessMap4 struct {
	Maps []AccessMap `json:"maps"`
	Page PageType `json:"page"`
}

// NewAccessMap4 instantiates a new AccessMap4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessMap4(maps []AccessMap, page PageType) *AccessMap4 {
	this := AccessMap4{}
	this.Maps = maps
	this.Page = page
	return &this
}

// NewAccessMap4WithDefaults instantiates a new AccessMap4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessMap4WithDefaults() *AccessMap4 {
	this := AccessMap4{}
	return &this
}

// GetMaps returns the Maps field value
func (o *AccessMap4) GetMaps() []AccessMap {
	if o == nil {
		var ret []AccessMap
		return ret
	}

	return o.Maps
}

// GetMapsOk returns a tuple with the Maps field value
// and a boolean to check if the value has been set.
func (o *AccessMap4) GetMapsOk() (*[]AccessMap, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Maps, true
}

// SetMaps sets field value
func (o *AccessMap4) SetMaps(v []AccessMap) {
	o.Maps = v
}

// GetPage returns the Page field value
func (o *AccessMap4) GetPage() PageType {
	if o == nil {
		var ret PageType
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *AccessMap4) GetPageOk() (*PageType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *AccessMap4) SetPage(v PageType) {
	o.Page = v
}

func (o AccessMap4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["maps"] = o.Maps
	}
	if true {
		toSerialize["page"] = o.Page
	}
	return json.Marshal(toSerialize)
}

type NullableAccessMap4 struct {
	value *AccessMap4
	isSet bool
}

func (v NullableAccessMap4) Get() *AccessMap4 {
	return v.value
}

func (v *NullableAccessMap4) Set(val *AccessMap4) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessMap4) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessMap4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessMap4(val *AccessMap4) *NullableAccessMap4 {
	return &NullableAccessMap4{value: val, isSet: true}
}

func (v NullableAccessMap4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessMap4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


