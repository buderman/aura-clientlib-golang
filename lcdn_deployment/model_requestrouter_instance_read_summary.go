/*
lcdn-deployment

Aura LCDN Deployment API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lcdn_deployment

import (
	"encoding/json"
)

// RequestrouterInstanceReadSummary Request Router instance summary read schema.
type RequestrouterInstanceReadSummary struct {
	// The unique fully qualified domain name for a node.
	Hostname string `json:"hostname"`
	// The unique identifier for a Request Router node.
	NodeId int32 `json:"nodeId"`
}

// NewRequestrouterInstanceReadSummary instantiates a new RequestrouterInstanceReadSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestrouterInstanceReadSummary(hostname string, nodeId int32) *RequestrouterInstanceReadSummary {
	this := RequestrouterInstanceReadSummary{}
	this.Hostname = hostname
	this.NodeId = nodeId
	return &this
}

// NewRequestrouterInstanceReadSummaryWithDefaults instantiates a new RequestrouterInstanceReadSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestrouterInstanceReadSummaryWithDefaults() *RequestrouterInstanceReadSummary {
	this := RequestrouterInstanceReadSummary{}
	return &this
}

// GetHostname returns the Hostname field value
func (o *RequestrouterInstanceReadSummary) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *RequestrouterInstanceReadSummary) GetHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *RequestrouterInstanceReadSummary) SetHostname(v string) {
	o.Hostname = v
}

// GetNodeId returns the NodeId field value
func (o *RequestrouterInstanceReadSummary) GetNodeId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *RequestrouterInstanceReadSummary) GetNodeIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *RequestrouterInstanceReadSummary) SetNodeId(v int32) {
	o.NodeId = v
}

func (o RequestrouterInstanceReadSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hostname"] = o.Hostname
	}
	if true {
		toSerialize["nodeId"] = o.NodeId
	}
	return json.Marshal(toSerialize)
}

type NullableRequestrouterInstanceReadSummary struct {
	value *RequestrouterInstanceReadSummary
	isSet bool
}

func (v NullableRequestrouterInstanceReadSummary) Get() *RequestrouterInstanceReadSummary {
	return v.value
}

func (v *NullableRequestrouterInstanceReadSummary) Set(val *RequestrouterInstanceReadSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestrouterInstanceReadSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestrouterInstanceReadSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestrouterInstanceReadSummary(val *RequestrouterInstanceReadSummary) *NullableRequestrouterInstanceReadSummary {
	return &NullableRequestrouterInstanceReadSummary{value: val, isSet: true}
}

func (v NullableRequestrouterInstanceReadSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestrouterInstanceReadSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


