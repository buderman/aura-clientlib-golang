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

// Site1 Site collection.
type Site1 struct {
	Sites []Site2 `json:"sites"`
	Page PageType `json:"page"`
}

// NewSite1 instantiates a new Site1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSite1(sites []Site2, page PageType) *Site1 {
	this := Site1{}
	this.Sites = sites
	this.Page = page
	return &this
}

// NewSite1WithDefaults instantiates a new Site1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSite1WithDefaults() *Site1 {
	this := Site1{}
	return &this
}

// GetSites returns the Sites field value
func (o *Site1) GetSites() []Site2 {
	if o == nil {
		var ret []Site2
		return ret
	}

	return o.Sites
}

// GetSitesOk returns a tuple with the Sites field value
// and a boolean to check if the value has been set.
func (o *Site1) GetSitesOk() (*[]Site2, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sites, true
}

// SetSites sets field value
func (o *Site1) SetSites(v []Site2) {
	o.Sites = v
}

// GetPage returns the Page field value
func (o *Site1) GetPage() PageType {
	if o == nil {
		var ret PageType
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *Site1) GetPageOk() (*PageType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *Site1) SetPage(v PageType) {
	o.Page = v
}

func (o Site1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sites"] = o.Sites
	}
	if true {
		toSerialize["page"] = o.Page
	}
	return json.Marshal(toSerialize)
}

type NullableSite1 struct {
	value *Site1
	isSet bool
}

func (v NullableSite1) Get() *Site1 {
	return v.value
}

func (v *NullableSite1) Set(val *Site1) {
	v.value = val
	v.isSet = true
}

func (v NullableSite1) IsSet() bool {
	return v.isSet
}

func (v *NullableSite1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSite1(val *Site1) *NullableSite1 {
	return &NullableSite1{value: val, isSet: true}
}

func (v NullableSite1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSite1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


