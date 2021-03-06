/*
lcdn-deployment

Aura LCDN Deployment API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lcdn_deployment

import (
	"encoding/json"
)

// ServiceLabelsType1 Service label definitions for a HyperCache instance.
type ServiceLabelsType1 struct {
	// List of IP addresses that you specify on a HyperCache node for intra-site traffic. Currently limited to a single address.
	IntraSite *[]string `json:"intraSite,omitempty"`
	// List of IP addresses that you configure on a HyperCache node for inter-site traffic. Regardless of the number of IP addresses you assigned, you can specify only one IP address for inter-site traffic.
	InterSite *[]string `json:"interSite,omitempty"`
	// Allows you to group one or more HyperCache IP addresses by associating them with a client-serving label.
	ClientServing *[]ClientServingLabelType `json:"clientServing,omitempty"`
}

// NewServiceLabelsType1 instantiates a new ServiceLabelsType1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceLabelsType1() *ServiceLabelsType1 {
	this := ServiceLabelsType1{}
	return &this
}

// NewServiceLabelsType1WithDefaults instantiates a new ServiceLabelsType1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceLabelsType1WithDefaults() *ServiceLabelsType1 {
	this := ServiceLabelsType1{}
	return &this
}

// GetIntraSite returns the IntraSite field value if set, zero value otherwise.
func (o *ServiceLabelsType1) GetIntraSite() []string {
	if o == nil || o.IntraSite == nil {
		var ret []string
		return ret
	}
	return *o.IntraSite
}

// GetIntraSiteOk returns a tuple with the IntraSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLabelsType1) GetIntraSiteOk() (*[]string, bool) {
	if o == nil || o.IntraSite == nil {
		return nil, false
	}
	return o.IntraSite, true
}

// HasIntraSite returns a boolean if a field has been set.
func (o *ServiceLabelsType1) HasIntraSite() bool {
	if o != nil && o.IntraSite != nil {
		return true
	}

	return false
}

// SetIntraSite gets a reference to the given []string and assigns it to the IntraSite field.
func (o *ServiceLabelsType1) SetIntraSite(v []string) {
	o.IntraSite = &v
}

// GetInterSite returns the InterSite field value if set, zero value otherwise.
func (o *ServiceLabelsType1) GetInterSite() []string {
	if o == nil || o.InterSite == nil {
		var ret []string
		return ret
	}
	return *o.InterSite
}

// GetInterSiteOk returns a tuple with the InterSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLabelsType1) GetInterSiteOk() (*[]string, bool) {
	if o == nil || o.InterSite == nil {
		return nil, false
	}
	return o.InterSite, true
}

// HasInterSite returns a boolean if a field has been set.
func (o *ServiceLabelsType1) HasInterSite() bool {
	if o != nil && o.InterSite != nil {
		return true
	}

	return false
}

// SetInterSite gets a reference to the given []string and assigns it to the InterSite field.
func (o *ServiceLabelsType1) SetInterSite(v []string) {
	o.InterSite = &v
}

// GetClientServing returns the ClientServing field value if set, zero value otherwise.
func (o *ServiceLabelsType1) GetClientServing() []ClientServingLabelType {
	if o == nil || o.ClientServing == nil {
		var ret []ClientServingLabelType
		return ret
	}
	return *o.ClientServing
}

// GetClientServingOk returns a tuple with the ClientServing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLabelsType1) GetClientServingOk() (*[]ClientServingLabelType, bool) {
	if o == nil || o.ClientServing == nil {
		return nil, false
	}
	return o.ClientServing, true
}

// HasClientServing returns a boolean if a field has been set.
func (o *ServiceLabelsType1) HasClientServing() bool {
	if o != nil && o.ClientServing != nil {
		return true
	}

	return false
}

// SetClientServing gets a reference to the given []ClientServingLabelType and assigns it to the ClientServing field.
func (o *ServiceLabelsType1) SetClientServing(v []ClientServingLabelType) {
	o.ClientServing = &v
}

func (o ServiceLabelsType1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IntraSite != nil {
		toSerialize["intraSite"] = o.IntraSite
	}
	if o.InterSite != nil {
		toSerialize["interSite"] = o.InterSite
	}
	if o.ClientServing != nil {
		toSerialize["clientServing"] = o.ClientServing
	}
	return json.Marshal(toSerialize)
}

type NullableServiceLabelsType1 struct {
	value *ServiceLabelsType1
	isSet bool
}

func (v NullableServiceLabelsType1) Get() *ServiceLabelsType1 {
	return v.value
}

func (v *NullableServiceLabelsType1) Set(val *ServiceLabelsType1) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceLabelsType1) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceLabelsType1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceLabelsType1(val *ServiceLabelsType1) *NullableServiceLabelsType1 {
	return &NullableServiceLabelsType1{value: val, isSet: true}
}

func (v NullableServiceLabelsType1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceLabelsType1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


