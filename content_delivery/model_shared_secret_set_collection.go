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

// SharedSecretSetCollection A list of shared secret sets.
type SharedSecretSetCollection struct {
	SharedSecretSets []SharedSecretSet3 `json:"sharedSecretSets"`
	Page PageType `json:"page"`
}

// NewSharedSecretSetCollection instantiates a new SharedSecretSetCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedSecretSetCollection(sharedSecretSets []SharedSecretSet3, page PageType) *SharedSecretSetCollection {
	this := SharedSecretSetCollection{}
	this.SharedSecretSets = sharedSecretSets
	this.Page = page
	return &this
}

// NewSharedSecretSetCollectionWithDefaults instantiates a new SharedSecretSetCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedSecretSetCollectionWithDefaults() *SharedSecretSetCollection {
	this := SharedSecretSetCollection{}
	return &this
}

// GetSharedSecretSets returns the SharedSecretSets field value
func (o *SharedSecretSetCollection) GetSharedSecretSets() []SharedSecretSet3 {
	if o == nil {
		var ret []SharedSecretSet3
		return ret
	}

	return o.SharedSecretSets
}

// GetSharedSecretSetsOk returns a tuple with the SharedSecretSets field value
// and a boolean to check if the value has been set.
func (o *SharedSecretSetCollection) GetSharedSecretSetsOk() (*[]SharedSecretSet3, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SharedSecretSets, true
}

// SetSharedSecretSets sets field value
func (o *SharedSecretSetCollection) SetSharedSecretSets(v []SharedSecretSet3) {
	o.SharedSecretSets = v
}

// GetPage returns the Page field value
func (o *SharedSecretSetCollection) GetPage() PageType {
	if o == nil {
		var ret PageType
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *SharedSecretSetCollection) GetPageOk() (*PageType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *SharedSecretSetCollection) SetPage(v PageType) {
	o.Page = v
}

func (o SharedSecretSetCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sharedSecretSets"] = o.SharedSecretSets
	}
	if true {
		toSerialize["page"] = o.Page
	}
	return json.Marshal(toSerialize)
}

type NullableSharedSecretSetCollection struct {
	value *SharedSecretSetCollection
	isSet bool
}

func (v NullableSharedSecretSetCollection) Get() *SharedSecretSetCollection {
	return v.value
}

func (v *NullableSharedSecretSetCollection) Set(val *SharedSecretSetCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedSecretSetCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedSecretSetCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedSecretSetCollection(val *SharedSecretSetCollection) *NullableSharedSecretSetCollection {
	return &NullableSharedSecretSetCollection{value: val, isSet: true}
}

func (v NullableSharedSecretSetCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedSecretSetCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


