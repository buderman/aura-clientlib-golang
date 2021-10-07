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

// Rules A rule object which is used to respond to different kinds of requests. A rule consists of criteria that identify which requests to process, and the behaviors to apply to those requests.
type Rules struct {
	DefaultBehaviors *[]DefaultBehaviorType `json:"defaultBehaviors,omitempty"`
	// A list of child rules. Each child rule object consists of its name, the criterion to match URIs, and a list of behaviors that act upon that criterion.
	Children *[]CdnPrefixRulesChildren `json:"children,omitempty"`
}

// NewRules instantiates a new Rules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRules() *Rules {
	this := Rules{}
	return &this
}

// NewRulesWithDefaults instantiates a new Rules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRulesWithDefaults() *Rules {
	this := Rules{}
	return &this
}

// GetDefaultBehaviors returns the DefaultBehaviors field value if set, zero value otherwise.
func (o *Rules) GetDefaultBehaviors() []DefaultBehaviorType {
	if o == nil || o.DefaultBehaviors == nil {
		var ret []DefaultBehaviorType
		return ret
	}
	return *o.DefaultBehaviors
}

// GetDefaultBehaviorsOk returns a tuple with the DefaultBehaviors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rules) GetDefaultBehaviorsOk() (*[]DefaultBehaviorType, bool) {
	if o == nil || o.DefaultBehaviors == nil {
		return nil, false
	}
	return o.DefaultBehaviors, true
}

// HasDefaultBehaviors returns a boolean if a field has been set.
func (o *Rules) HasDefaultBehaviors() bool {
	if o != nil && o.DefaultBehaviors != nil {
		return true
	}

	return false
}

// SetDefaultBehaviors gets a reference to the given []DefaultBehaviorType and assigns it to the DefaultBehaviors field.
func (o *Rules) SetDefaultBehaviors(v []DefaultBehaviorType) {
	o.DefaultBehaviors = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *Rules) GetChildren() []CdnPrefixRulesChildren {
	if o == nil || o.Children == nil {
		var ret []CdnPrefixRulesChildren
		return ret
	}
	return *o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rules) GetChildrenOk() (*[]CdnPrefixRulesChildren, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *Rules) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []CdnPrefixRulesChildren and assigns it to the Children field.
func (o *Rules) SetChildren(v []CdnPrefixRulesChildren) {
	o.Children = &v
}

func (o Rules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultBehaviors != nil {
		toSerialize["defaultBehaviors"] = o.DefaultBehaviors
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	return json.Marshal(toSerialize)
}

type NullableRules struct {
	value *Rules
	isSet bool
}

func (v NullableRules) Get() *Rules {
	return v.value
}

func (v *NullableRules) Set(val *Rules) {
	v.value = val
	v.isSet = true
}

func (v NullableRules) IsSet() bool {
	return v.isSet
}

func (v *NullableRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRules(val *Rules) *NullableRules {
	return &NullableRules{value: val, isSet: true}
}

func (v NullableRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


