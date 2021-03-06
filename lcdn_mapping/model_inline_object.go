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

// InlineObject struct for InlineObject
type InlineObject struct {
	Metadata *SiteMapCreate `json:"metadata,omitempty"`
	Config interface{} `json:"config,omitempty"`
}

// NewInlineObject instantiates a new InlineObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject() *InlineObject {
	this := InlineObject{}
	return &this
}

// NewInlineObjectWithDefaults instantiates a new InlineObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObjectWithDefaults() *InlineObject {
	this := InlineObject{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *InlineObject) GetMetadata() SiteMapCreate {
	if o == nil || o.Metadata == nil {
		var ret SiteMapCreate
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject) GetMetadataOk() (*SiteMapCreate, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *InlineObject) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given SiteMapCreate and assigns it to the Metadata field.
func (o *InlineObject) SetMetadata(v SiteMapCreate) {
	o.Metadata = &v
}

// GetConfig returns the Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject) GetConfig() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject) GetConfigOk() (*interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return &o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *InlineObject) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given interface{} and assigns it to the Config field.
func (o *InlineObject) SetConfig(v interface{}) {
	o.Config = v
}

func (o InlineObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject struct {
	value *InlineObject
	isSet bool
}

func (v NullableInlineObject) Get() *InlineObject {
	return v.value
}

func (v *NullableInlineObject) Set(val *InlineObject) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject(val *InlineObject) *NullableInlineObject {
	return &NullableInlineObject{value: val, isSet: true}
}

func (v NullableInlineObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


