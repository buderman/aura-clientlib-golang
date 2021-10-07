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

// SiteMapCreate Site Map create schema.
type SiteMapCreate struct {
	// A unique name for the entity.
	Name string `json:"name"`
	// A description for the entity.
	Description *string `json:"description,omitempty"`
}

// NewSiteMapCreate instantiates a new SiteMapCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteMapCreate(name string) *SiteMapCreate {
	this := SiteMapCreate{}
	this.Name = name
	return &this
}

// NewSiteMapCreateWithDefaults instantiates a new SiteMapCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteMapCreateWithDefaults() *SiteMapCreate {
	this := SiteMapCreate{}
	return &this
}

// GetName returns the Name field value
func (o *SiteMapCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SiteMapCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SiteMapCreate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SiteMapCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteMapCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SiteMapCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SiteMapCreate) SetDescription(v string) {
	o.Description = &v
}

func (o SiteMapCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableSiteMapCreate struct {
	value *SiteMapCreate
	isSet bool
}

func (v NullableSiteMapCreate) Get() *SiteMapCreate {
	return v.value
}

func (v *NullableSiteMapCreate) Set(val *SiteMapCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteMapCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteMapCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteMapCreate(val *SiteMapCreate) *NullableSiteMapCreate {
	return &NullableSiteMapCreate{value: val, isSet: true}
}

func (v NullableSiteMapCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteMapCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


