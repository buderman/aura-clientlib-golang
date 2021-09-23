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

// CdnPrefixRulesChildren struct for CdnPrefixRulesChildren
type CdnPrefixRulesChildren struct {
	Behaviors []BehaviorType `json:"behaviors"`
	Criterion Criterion `json:"criterion"`
	// Name of the rule.
	Name string `json:"name"`
}

// NewCdnPrefixRulesChildren instantiates a new CdnPrefixRulesChildren object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdnPrefixRulesChildren(behaviors []BehaviorType, criterion Criterion, name string) *CdnPrefixRulesChildren {
	this := CdnPrefixRulesChildren{}
	this.Behaviors = behaviors
	this.Criterion = criterion
	this.Name = name
	return &this
}

// NewCdnPrefixRulesChildrenWithDefaults instantiates a new CdnPrefixRulesChildren object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdnPrefixRulesChildrenWithDefaults() *CdnPrefixRulesChildren {
	this := CdnPrefixRulesChildren{}
	return &this
}

// GetBehaviors returns the Behaviors field value
func (o *CdnPrefixRulesChildren) GetBehaviors() []BehaviorType {
	if o == nil {
		var ret []BehaviorType
		return ret
	}

	return o.Behaviors
}

// GetBehaviorsOk returns a tuple with the Behaviors field value
// and a boolean to check if the value has been set.
func (o *CdnPrefixRulesChildren) GetBehaviorsOk() (*[]BehaviorType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Behaviors, true
}

// SetBehaviors sets field value
func (o *CdnPrefixRulesChildren) SetBehaviors(v []BehaviorType) {
	o.Behaviors = v
}

// GetCriterion returns the Criterion field value
func (o *CdnPrefixRulesChildren) GetCriterion() Criterion {
	if o == nil {
		var ret Criterion
		return ret
	}

	return o.Criterion
}

// GetCriterionOk returns a tuple with the Criterion field value
// and a boolean to check if the value has been set.
func (o *CdnPrefixRulesChildren) GetCriterionOk() (*Criterion, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Criterion, true
}

// SetCriterion sets field value
func (o *CdnPrefixRulesChildren) SetCriterion(v Criterion) {
	o.Criterion = v
}

// GetName returns the Name field value
func (o *CdnPrefixRulesChildren) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CdnPrefixRulesChildren) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CdnPrefixRulesChildren) SetName(v string) {
	o.Name = v
}

func (o CdnPrefixRulesChildren) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["behaviors"] = o.Behaviors
	}
	if true {
		toSerialize["criterion"] = o.Criterion
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableCdnPrefixRulesChildren struct {
	value *CdnPrefixRulesChildren
	isSet bool
}

func (v NullableCdnPrefixRulesChildren) Get() *CdnPrefixRulesChildren {
	return v.value
}

func (v *NullableCdnPrefixRulesChildren) Set(val *CdnPrefixRulesChildren) {
	v.value = val
	v.isSet = true
}

func (v NullableCdnPrefixRulesChildren) IsSet() bool {
	return v.isSet
}

func (v *NullableCdnPrefixRulesChildren) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdnPrefixRulesChildren(val *CdnPrefixRulesChildren) *NullableCdnPrefixRulesChildren {
	return &NullableCdnPrefixRulesChildren{value: val, isSet: true}
}

func (v NullableCdnPrefixRulesChildren) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdnPrefixRulesChildren) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


