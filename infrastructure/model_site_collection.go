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

// SiteCollection Site collection.
type SiteCollection struct {
	Sites []Site2 `json:"sites"`
	Page PageType `json:"page"`
}

// NewSiteCollection instantiates a new SiteCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteCollection(sites []Site2, page PageType) *SiteCollection {
	this := SiteCollection{}
	this.Sites = sites
	this.Page = page
	return &this
}

// NewSiteCollectionWithDefaults instantiates a new SiteCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteCollectionWithDefaults() *SiteCollection {
	this := SiteCollection{}
	return &this
}

// GetSites returns the Sites field value
func (o *SiteCollection) GetSites() []Site2 {
	if o == nil {
		var ret []Site2
		return ret
	}

	return o.Sites
}

// GetSitesOk returns a tuple with the Sites field value
// and a boolean to check if the value has been set.
func (o *SiteCollection) GetSitesOk() (*[]Site2, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sites, true
}

// SetSites sets field value
func (o *SiteCollection) SetSites(v []Site2) {
	o.Sites = v
}

// GetPage returns the Page field value
func (o *SiteCollection) GetPage() PageType {
	if o == nil {
		var ret PageType
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *SiteCollection) GetPageOk() (*PageType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *SiteCollection) SetPage(v PageType) {
	o.Page = v
}

func (o SiteCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sites"] = o.Sites
	}
	if true {
		toSerialize["page"] = o.Page
	}
	return json.Marshal(toSerialize)
}

type NullableSiteCollection struct {
	value *SiteCollection
	isSet bool
}

func (v NullableSiteCollection) Get() *SiteCollection {
	return v.value
}

func (v *NullableSiteCollection) Set(val *SiteCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteCollection(val *SiteCollection) *NullableSiteCollection {
	return &NullableSiteCollection{value: val, isSet: true}
}

func (v NullableSiteCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


