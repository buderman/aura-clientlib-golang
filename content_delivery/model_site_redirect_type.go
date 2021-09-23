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

// SiteRedirectType Site redirect type for HyperCache to make redirect decisions based on client IP or site information.
type SiteRedirectType struct {
	// The name of the behavior, `siteRedirectType` in this case.
	Name string `json:"name"`
	Options SiteRedirectTypeOptions `json:"options"`
}

// NewSiteRedirectType instantiates a new SiteRedirectType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteRedirectType(name string, options SiteRedirectTypeOptions) *SiteRedirectType {
	this := SiteRedirectType{}
	this.Name = name
	this.Options = options
	return &this
}

// NewSiteRedirectTypeWithDefaults instantiates a new SiteRedirectType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteRedirectTypeWithDefaults() *SiteRedirectType {
	this := SiteRedirectType{}
	return &this
}

// GetName returns the Name field value
func (o *SiteRedirectType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SiteRedirectType) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SiteRedirectType) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *SiteRedirectType) GetOptions() SiteRedirectTypeOptions {
	if o == nil {
		var ret SiteRedirectTypeOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *SiteRedirectType) GetOptionsOk() (*SiteRedirectTypeOptions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *SiteRedirectType) SetOptions(v SiteRedirectTypeOptions) {
	o.Options = v
}

func (o SiteRedirectType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableSiteRedirectType struct {
	value *SiteRedirectType
	isSet bool
}

func (v NullableSiteRedirectType) Get() *SiteRedirectType {
	return v.value
}

func (v *NullableSiteRedirectType) Set(val *SiteRedirectType) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteRedirectType) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteRedirectType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteRedirectType(val *SiteRedirectType) *NullableSiteRedirectType {
	return &NullableSiteRedirectType{value: val, isSet: true}
}

func (v NullableSiteRedirectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteRedirectType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


