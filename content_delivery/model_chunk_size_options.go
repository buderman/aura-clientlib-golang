/*
content-delivery

Aura LCDN Content Delivery API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package content_delivery

import (
	"encoding/json"
)

// ChunkSizeOptions This behavior's set of configuration options.
type ChunkSizeOptions struct {
	// Size of the chunk for this content, in bytes.
	Value int32 `json:"value"`
}

// NewChunkSizeOptions instantiates a new ChunkSizeOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChunkSizeOptions(value int32) *ChunkSizeOptions {
	this := ChunkSizeOptions{}
	this.Value = value
	return &this
}

// NewChunkSizeOptionsWithDefaults instantiates a new ChunkSizeOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChunkSizeOptionsWithDefaults() *ChunkSizeOptions {
	this := ChunkSizeOptions{}
	return &this
}

// GetValue returns the Value field value
func (o *ChunkSizeOptions) GetValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ChunkSizeOptions) GetValueOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ChunkSizeOptions) SetValue(v int32) {
	o.Value = v
}

func (o ChunkSizeOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableChunkSizeOptions struct {
	value *ChunkSizeOptions
	isSet bool
}

func (v NullableChunkSizeOptions) Get() *ChunkSizeOptions {
	return v.value
}

func (v *NullableChunkSizeOptions) Set(val *ChunkSizeOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableChunkSizeOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableChunkSizeOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChunkSizeOptions(val *ChunkSizeOptions) *NullableChunkSizeOptions {
	return &NullableChunkSizeOptions{value: val, isSet: true}
}

func (v NullableChunkSizeOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChunkSizeOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


