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

// DefaultStoragePartition Specifies HyperCache partitions to reserve hard disk space for certain content types.
type DefaultStoragePartition struct {
	// The name of the behavior, `storagePartition` in this case.
	Name string `json:"name"`
	Options DefaultStoragePartitionOptions `json:"options"`
}

// NewDefaultStoragePartition instantiates a new DefaultStoragePartition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultStoragePartition(name string, options DefaultStoragePartitionOptions) *DefaultStoragePartition {
	this := DefaultStoragePartition{}
	this.Name = name
	this.Options = options
	return &this
}

// NewDefaultStoragePartitionWithDefaults instantiates a new DefaultStoragePartition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultStoragePartitionWithDefaults() *DefaultStoragePartition {
	this := DefaultStoragePartition{}
	return &this
}

// GetName returns the Name field value
func (o *DefaultStoragePartition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DefaultStoragePartition) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DefaultStoragePartition) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *DefaultStoragePartition) GetOptions() DefaultStoragePartitionOptions {
	if o == nil {
		var ret DefaultStoragePartitionOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *DefaultStoragePartition) GetOptionsOk() (*DefaultStoragePartitionOptions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *DefaultStoragePartition) SetOptions(v DefaultStoragePartitionOptions) {
	o.Options = v
}

func (o DefaultStoragePartition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableDefaultStoragePartition struct {
	value *DefaultStoragePartition
	isSet bool
}

func (v NullableDefaultStoragePartition) Get() *DefaultStoragePartition {
	return v.value
}

func (v *NullableDefaultStoragePartition) Set(val *DefaultStoragePartition) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultStoragePartition) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultStoragePartition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultStoragePartition(val *DefaultStoragePartition) *NullableDefaultStoragePartition {
	return &NullableDefaultStoragePartition{value: val, isSet: true}
}

func (v NullableDefaultStoragePartition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultStoragePartition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

