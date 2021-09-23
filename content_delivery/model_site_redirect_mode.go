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

// SiteRedirectMode Enables site redirect mode for HyperCache to redirect requests to optimal sites if needed.
type SiteRedirectMode struct {
	// The name of the behavior, `siteRedirectMode` in this case.
	Name string `json:"name"`
	Options SiteRedirectModeOptions `json:"options"`
}

// NewSiteRedirectMode instantiates a new SiteRedirectMode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteRedirectMode(name string, options SiteRedirectModeOptions) *SiteRedirectMode {
	this := SiteRedirectMode{}
	this.Name = name
	this.Options = options
	return &this
}

// NewSiteRedirectModeWithDefaults instantiates a new SiteRedirectMode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteRedirectModeWithDefaults() *SiteRedirectMode {
	this := SiteRedirectMode{}
	return &this
}

// GetName returns the Name field value
func (o *SiteRedirectMode) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SiteRedirectMode) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SiteRedirectMode) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *SiteRedirectMode) GetOptions() SiteRedirectModeOptions {
	if o == nil {
		var ret SiteRedirectModeOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *SiteRedirectMode) GetOptionsOk() (*SiteRedirectModeOptions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *SiteRedirectMode) SetOptions(v SiteRedirectModeOptions) {
	o.Options = v
}

func (o SiteRedirectMode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableSiteRedirectMode struct {
	value *SiteRedirectMode
	isSet bool
}

func (v NullableSiteRedirectMode) Get() *SiteRedirectMode {
	return v.value
}

func (v *NullableSiteRedirectMode) Set(val *SiteRedirectMode) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteRedirectMode) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteRedirectMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteRedirectMode(val *SiteRedirectMode) *NullableSiteRedirectMode {
	return &NullableSiteRedirectMode{value: val, isSet: true}
}

func (v NullableSiteRedirectMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteRedirectMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

