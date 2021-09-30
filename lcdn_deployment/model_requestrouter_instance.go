/*
lcdn-deployment

LCDN Deployment API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lcdn_deployment

import (
	"encoding/json"
)

// RequestrouterInstance Request Router instance update schema.
type RequestrouterInstance struct {
	// The unique identifier for a Request Router node.
	NodeId int32 `json:"nodeId"`
	ServiceLabels *ServiceLabelsType `json:"serviceLabels,omitempty"`
}

// NewRequestrouterInstance instantiates a new RequestrouterInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestrouterInstance(nodeId int32) *RequestrouterInstance {
	this := RequestrouterInstance{}
	this.NodeId = nodeId
	return &this
}

// NewRequestrouterInstanceWithDefaults instantiates a new RequestrouterInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestrouterInstanceWithDefaults() *RequestrouterInstance {
	this := RequestrouterInstance{}
	return &this
}

// GetNodeId returns the NodeId field value
func (o *RequestrouterInstance) GetNodeId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *RequestrouterInstance) GetNodeIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *RequestrouterInstance) SetNodeId(v int32) {
	o.NodeId = v
}

// GetServiceLabels returns the ServiceLabels field value if set, zero value otherwise.
func (o *RequestrouterInstance) GetServiceLabels() ServiceLabelsType {
	if o == nil || o.ServiceLabels == nil {
		var ret ServiceLabelsType
		return ret
	}
	return *o.ServiceLabels
}

// GetServiceLabelsOk returns a tuple with the ServiceLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestrouterInstance) GetServiceLabelsOk() (*ServiceLabelsType, bool) {
	if o == nil || o.ServiceLabels == nil {
		return nil, false
	}
	return o.ServiceLabels, true
}

// HasServiceLabels returns a boolean if a field has been set.
func (o *RequestrouterInstance) HasServiceLabels() bool {
	if o != nil && o.ServiceLabels != nil {
		return true
	}

	return false
}

// SetServiceLabels gets a reference to the given ServiceLabelsType and assigns it to the ServiceLabels field.
func (o *RequestrouterInstance) SetServiceLabels(v ServiceLabelsType) {
	o.ServiceLabels = &v
}

func (o RequestrouterInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ServiceLabels != nil {
		toSerialize["serviceLabels"] = o.ServiceLabels
	}
	return json.Marshal(toSerialize)
}

type NullableRequestrouterInstance struct {
	value *RequestrouterInstance
	isSet bool
}

func (v NullableRequestrouterInstance) Get() *RequestrouterInstance {
	return v.value
}

func (v *NullableRequestrouterInstance) Set(val *RequestrouterInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestrouterInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestrouterInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestrouterInstance(val *RequestrouterInstance) *NullableRequestrouterInstance {
	return &NullableRequestrouterInstance{value: val, isSet: true}
}

func (v NullableRequestrouterInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestrouterInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


