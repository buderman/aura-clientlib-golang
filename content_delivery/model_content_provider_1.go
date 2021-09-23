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

// ContentProvider1 ContentProvider collection
type ContentProvider1 struct {
	Page PageType `json:"page"`
	ContentProviders []ContentProvider2 `json:"contentProviders"`
}

// NewContentProvider1 instantiates a new ContentProvider1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentProvider1(page PageType, contentProviders []ContentProvider2) *ContentProvider1 {
	this := ContentProvider1{}
	this.Page = page
	this.ContentProviders = contentProviders
	return &this
}

// NewContentProvider1WithDefaults instantiates a new ContentProvider1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentProvider1WithDefaults() *ContentProvider1 {
	this := ContentProvider1{}
	return &this
}

// GetPage returns the Page field value
func (o *ContentProvider1) GetPage() PageType {
	if o == nil {
		var ret PageType
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *ContentProvider1) GetPageOk() (*PageType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *ContentProvider1) SetPage(v PageType) {
	o.Page = v
}

// GetContentProviders returns the ContentProviders field value
func (o *ContentProvider1) GetContentProviders() []ContentProvider2 {
	if o == nil {
		var ret []ContentProvider2
		return ret
	}

	return o.ContentProviders
}

// GetContentProvidersOk returns a tuple with the ContentProviders field value
// and a boolean to check if the value has been set.
func (o *ContentProvider1) GetContentProvidersOk() (*[]ContentProvider2, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContentProviders, true
}

// SetContentProviders sets field value
func (o *ContentProvider1) SetContentProviders(v []ContentProvider2) {
	o.ContentProviders = v
}

func (o ContentProvider1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["contentProviders"] = o.ContentProviders
	}
	return json.Marshal(toSerialize)
}

type NullableContentProvider1 struct {
	value *ContentProvider1
	isSet bool
}

func (v NullableContentProvider1) Get() *ContentProvider1 {
	return v.value
}

func (v *NullableContentProvider1) Set(val *ContentProvider1) {
	v.value = val
	v.isSet = true
}

func (v NullableContentProvider1) IsSet() bool {
	return v.isSet
}

func (v *NullableContentProvider1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentProvider1(val *ContentProvider1) *NullableContentProvider1 {
	return &NullableContentProvider1{value: val, isSet: true}
}

func (v NullableContentProvider1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentProvider1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

