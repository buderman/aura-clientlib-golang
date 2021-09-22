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

// GeoipConfigReadSummary GeoIP config summary read schema.
type GeoipConfigReadSummary struct {
	// The unique identifier for a GeoIP configuration.
	GeoipConfigId *int32 `json:"geoipConfigId,omitempty"`
	Type TypeType `json:"type"`
	// A unique name for the entity.
	Name string `json:"name"`
}

// NewGeoipConfigReadSummary instantiates a new GeoipConfigReadSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoipConfigReadSummary(type_ TypeType, name string) *GeoipConfigReadSummary {
	this := GeoipConfigReadSummary{}
	this.Type = type_
	this.Name = name
	return &this
}

// NewGeoipConfigReadSummaryWithDefaults instantiates a new GeoipConfigReadSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoipConfigReadSummaryWithDefaults() *GeoipConfigReadSummary {
	this := GeoipConfigReadSummary{}
	return &this
}

// GetGeoipConfigId returns the GeoipConfigId field value if set, zero value otherwise.
func (o *GeoipConfigReadSummary) GetGeoipConfigId() int32 {
	if o == nil || o.GeoipConfigId == nil {
		var ret int32
		return ret
	}
	return *o.GeoipConfigId
}

// GetGeoipConfigIdOk returns a tuple with the GeoipConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoipConfigReadSummary) GetGeoipConfigIdOk() (*int32, bool) {
	if o == nil || o.GeoipConfigId == nil {
		return nil, false
	}
	return o.GeoipConfigId, true
}

// HasGeoipConfigId returns a boolean if a field has been set.
func (o *GeoipConfigReadSummary) HasGeoipConfigId() bool {
	if o != nil && o.GeoipConfigId != nil {
		return true
	}

	return false
}

// SetGeoipConfigId gets a reference to the given int32 and assigns it to the GeoipConfigId field.
func (o *GeoipConfigReadSummary) SetGeoipConfigId(v int32) {
	o.GeoipConfigId = &v
}

// GetType returns the Type field value
func (o *GeoipConfigReadSummary) GetType() TypeType {
	if o == nil {
		var ret TypeType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GeoipConfigReadSummary) GetTypeOk() (*TypeType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GeoipConfigReadSummary) SetType(v TypeType) {
	o.Type = v
}

// GetName returns the Name field value
func (o *GeoipConfigReadSummary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GeoipConfigReadSummary) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GeoipConfigReadSummary) SetName(v string) {
	o.Name = v
}

func (o GeoipConfigReadSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GeoipConfigId != nil {
		toSerialize["geoipConfigId"] = o.GeoipConfigId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableGeoipConfigReadSummary struct {
	value *GeoipConfigReadSummary
	isSet bool
}

func (v NullableGeoipConfigReadSummary) Get() *GeoipConfigReadSummary {
	return v.value
}

func (v *NullableGeoipConfigReadSummary) Set(val *GeoipConfigReadSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoipConfigReadSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoipConfigReadSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoipConfigReadSummary(val *GeoipConfigReadSummary) *NullableGeoipConfigReadSummary {
	return &NullableGeoipConfigReadSummary{value: val, isSet: true}
}

func (v NullableGeoipConfigReadSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoipConfigReadSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


