/*
infrastructure

Aura Infrastructure API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infrastructure

import (
	"encoding/json"
)

// Site4 Site detailed read schema.
type Site4 struct {
	// The unique identifier for a site.
	SiteId int32 `json:"siteId"`
	// The unique name for this site.
	Name string `json:"name"`
	// A shortened name for this site.
	AbbreviatedName string `json:"abbreviatedName"`
}

// NewSite4 instantiates a new Site4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSite4(siteId int32, name string, abbreviatedName string) *Site4 {
	this := Site4{}
	this.SiteId = siteId
	this.Name = name
	this.AbbreviatedName = abbreviatedName
	return &this
}

// NewSite4WithDefaults instantiates a new Site4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSite4WithDefaults() *Site4 {
	this := Site4{}
	return &this
}

// GetSiteId returns the SiteId field value
func (o *Site4) GetSiteId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *Site4) GetSiteIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *Site4) SetSiteId(v int32) {
	o.SiteId = v
}

// GetName returns the Name field value
func (o *Site4) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Site4) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Site4) SetName(v string) {
	o.Name = v
}

// GetAbbreviatedName returns the AbbreviatedName field value
func (o *Site4) GetAbbreviatedName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AbbreviatedName
}

// GetAbbreviatedNameOk returns a tuple with the AbbreviatedName field value
// and a boolean to check if the value has been set.
func (o *Site4) GetAbbreviatedNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AbbreviatedName, true
}

// SetAbbreviatedName sets field value
func (o *Site4) SetAbbreviatedName(v string) {
	o.AbbreviatedName = v
}

func (o Site4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["siteId"] = o.SiteId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["abbreviatedName"] = o.AbbreviatedName
	}
	return json.Marshal(toSerialize)
}

type NullableSite4 struct {
	value *Site4
	isSet bool
}

func (v NullableSite4) Get() *Site4 {
	return v.value
}

func (v *NullableSite4) Set(val *Site4) {
	v.value = val
	v.isSet = true
}

func (v NullableSite4) IsSet() bool {
	return v.isSet
}

func (v *NullableSite4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSite4(val *Site4) *NullableSite4 {
	return &NullableSite4{value: val, isSet: true}
}

func (v NullableSite4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSite4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


