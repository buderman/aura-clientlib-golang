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

// CdnPrefixRules A rule object which is used to respond to different kinds of requests. A rule consists of criteria that identify which requests to process, and the behaviors to apply to those requests.
type CdnPrefixRules struct {
	DefaultBehaviors *[]DefaultBehaviorType `json:"defaultBehaviors,omitempty"`
	// A list of child rules. Each child rule object consists of its name, the criterion to match URIs, and a list of behaviors that act upon that criterion.
	Children *[]CdnPrefixRulesChildren `json:"children,omitempty"`
}

// NewCdnPrefixRules instantiates a new CdnPrefixRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdnPrefixRules() *CdnPrefixRules {
	this := CdnPrefixRules{}
	return &this
}

// NewCdnPrefixRulesWithDefaults instantiates a new CdnPrefixRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdnPrefixRulesWithDefaults() *CdnPrefixRules {
	this := CdnPrefixRules{}
	return &this
}

// GetDefaultBehaviors returns the DefaultBehaviors field value if set, zero value otherwise.
func (o *CdnPrefixRules) GetDefaultBehaviors() []DefaultBehaviorType {
	if o == nil || o.DefaultBehaviors == nil {
		var ret []DefaultBehaviorType
		return ret
	}
	return *o.DefaultBehaviors
}

// GetDefaultBehaviorsOk returns a tuple with the DefaultBehaviors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnPrefixRules) GetDefaultBehaviorsOk() (*[]DefaultBehaviorType, bool) {
	if o == nil || o.DefaultBehaviors == nil {
		return nil, false
	}
	return o.DefaultBehaviors, true
}

// HasDefaultBehaviors returns a boolean if a field has been set.
func (o *CdnPrefixRules) HasDefaultBehaviors() bool {
	if o != nil && o.DefaultBehaviors != nil {
		return true
	}

	return false
}

// SetDefaultBehaviors gets a reference to the given []DefaultBehaviorType and assigns it to the DefaultBehaviors field.
func (o *CdnPrefixRules) SetDefaultBehaviors(v []DefaultBehaviorType) {
	o.DefaultBehaviors = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *CdnPrefixRules) GetChildren() []CdnPrefixRulesChildren {
	if o == nil || o.Children == nil {
		var ret []CdnPrefixRulesChildren
		return ret
	}
	return *o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnPrefixRules) GetChildrenOk() (*[]CdnPrefixRulesChildren, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *CdnPrefixRules) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []CdnPrefixRulesChildren and assigns it to the Children field.
func (o *CdnPrefixRules) SetChildren(v []CdnPrefixRulesChildren) {
	o.Children = &v
}

func (o CdnPrefixRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultBehaviors != nil {
		toSerialize["defaultBehaviors"] = o.DefaultBehaviors
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	return json.Marshal(toSerialize)
}

type NullableCdnPrefixRules struct {
	value *CdnPrefixRules
	isSet bool
}

func (v NullableCdnPrefixRules) Get() *CdnPrefixRules {
	return v.value
}

func (v *NullableCdnPrefixRules) Set(val *CdnPrefixRules) {
	v.value = val
	v.isSet = true
}

func (v NullableCdnPrefixRules) IsSet() bool {
	return v.isSet
}

func (v *NullableCdnPrefixRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdnPrefixRules(val *CdnPrefixRules) *NullableCdnPrefixRules {
	return &NullableCdnPrefixRules{value: val, isSet: true}
}

func (v NullableCdnPrefixRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdnPrefixRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


