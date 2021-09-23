/*
content-delivery

Cotent Delivery API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package content_delivery

import (
	"encoding/json"
)

// ChunkSize Chunk size for the asset, used for tuning the HyperCache internal chunk size.
type ChunkSize struct {
	// The name of the behavior, `chunkSize` in this case.
	Name string `json:"name"`
	Options ChunkSizeOptions `json:"options"`
}

// NewChunkSize instantiates a new ChunkSize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChunkSize(name string, options ChunkSizeOptions) *ChunkSize {
	this := ChunkSize{}
	this.Name = name
	this.Options = options
	return &this
}

// NewChunkSizeWithDefaults instantiates a new ChunkSize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChunkSizeWithDefaults() *ChunkSize {
	this := ChunkSize{}
	return &this
}

// GetName returns the Name field value
func (o *ChunkSize) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ChunkSize) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ChunkSize) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *ChunkSize) GetOptions() ChunkSizeOptions {
	if o == nil {
		var ret ChunkSizeOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ChunkSize) GetOptionsOk() (*ChunkSizeOptions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *ChunkSize) SetOptions(v ChunkSizeOptions) {
	o.Options = v
}

func (o ChunkSize) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableChunkSize struct {
	value *ChunkSize
	isSet bool
}

func (v NullableChunkSize) Get() *ChunkSize {
	return v.value
}

func (v *NullableChunkSize) Set(val *ChunkSize) {
	v.value = val
	v.isSet = true
}

func (v NullableChunkSize) IsSet() bool {
	return v.isSet
}

func (v *NullableChunkSize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChunkSize(val *ChunkSize) *NullableChunkSize {
	return &NullableChunkSize{value: val, isSet: true}
}

func (v NullableChunkSize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChunkSize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

