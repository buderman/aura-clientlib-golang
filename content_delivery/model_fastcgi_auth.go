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

// FastcgiAuth Parameters associated with FastCGI Authentication.
type FastcgiAuth struct {
	// The name of the behavior, `fastcgiAuth` in this case.
	Name string `json:"name"`
	Options FastcgiAuthOptions `json:"options"`
}

// NewFastcgiAuth instantiates a new FastcgiAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFastcgiAuth(name string, options FastcgiAuthOptions) *FastcgiAuth {
	this := FastcgiAuth{}
	this.Name = name
	this.Options = options
	return &this
}

// NewFastcgiAuthWithDefaults instantiates a new FastcgiAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFastcgiAuthWithDefaults() *FastcgiAuth {
	this := FastcgiAuth{}
	return &this
}

// GetName returns the Name field value
func (o *FastcgiAuth) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FastcgiAuth) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FastcgiAuth) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *FastcgiAuth) GetOptions() FastcgiAuthOptions {
	if o == nil {
		var ret FastcgiAuthOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *FastcgiAuth) GetOptionsOk() (*FastcgiAuthOptions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *FastcgiAuth) SetOptions(v FastcgiAuthOptions) {
	o.Options = v
}

func (o FastcgiAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableFastcgiAuth struct {
	value *FastcgiAuth
	isSet bool
}

func (v NullableFastcgiAuth) Get() *FastcgiAuth {
	return v.value
}

func (v *NullableFastcgiAuth) Set(val *FastcgiAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableFastcgiAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableFastcgiAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFastcgiAuth(val *FastcgiAuth) *NullableFastcgiAuth {
	return &NullableFastcgiAuth{value: val, isSet: true}
}

func (v NullableFastcgiAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFastcgiAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


