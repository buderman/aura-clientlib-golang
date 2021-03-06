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

// Route4Type A single IPv4 route.
type Route4Type struct {
	// An IPv4 address.
	Destination string `json:"destination"`
	// A value up to 9999 that represents the cost metric for this static route. The route with the lowest metric is chosen among multiple routes that most closely match the destination address of a packet being forwarded.
	Metric *int32 `json:"metric,omitempty"`
	// The IPv4 source address for this IPv4 route.
	SourceIpAddress *string `json:"sourceIpAddress,omitempty"`
	// The address of the next gateway to which packets should be forwarded along the path to their final destination.
	NextHop string `json:"nextHop"`
}

// NewRoute4Type instantiates a new Route4Type object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoute4Type(destination string, nextHop string) *Route4Type {
	this := Route4Type{}
	this.Destination = destination
	this.NextHop = nextHop
	return &this
}

// NewRoute4TypeWithDefaults instantiates a new Route4Type object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoute4TypeWithDefaults() *Route4Type {
	this := Route4Type{}
	return &this
}

// GetDestination returns the Destination field value
func (o *Route4Type) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *Route4Type) GetDestinationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *Route4Type) SetDestination(v string) {
	o.Destination = v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *Route4Type) GetMetric() int32 {
	if o == nil || o.Metric == nil {
		var ret int32
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route4Type) GetMetricOk() (*int32, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *Route4Type) HasMetric() bool {
	if o != nil && o.Metric != nil {
		return true
	}

	return false
}

// SetMetric gets a reference to the given int32 and assigns it to the Metric field.
func (o *Route4Type) SetMetric(v int32) {
	o.Metric = &v
}

// GetSourceIpAddress returns the SourceIpAddress field value if set, zero value otherwise.
func (o *Route4Type) GetSourceIpAddress() string {
	if o == nil || o.SourceIpAddress == nil {
		var ret string
		return ret
	}
	return *o.SourceIpAddress
}

// GetSourceIpAddressOk returns a tuple with the SourceIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route4Type) GetSourceIpAddressOk() (*string, bool) {
	if o == nil || o.SourceIpAddress == nil {
		return nil, false
	}
	return o.SourceIpAddress, true
}

// HasSourceIpAddress returns a boolean if a field has been set.
func (o *Route4Type) HasSourceIpAddress() bool {
	if o != nil && o.SourceIpAddress != nil {
		return true
	}

	return false
}

// SetSourceIpAddress gets a reference to the given string and assigns it to the SourceIpAddress field.
func (o *Route4Type) SetSourceIpAddress(v string) {
	o.SourceIpAddress = &v
}

// GetNextHop returns the NextHop field value
func (o *Route4Type) GetNextHop() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextHop
}

// GetNextHopOk returns a tuple with the NextHop field value
// and a boolean to check if the value has been set.
func (o *Route4Type) GetNextHopOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NextHop, true
}

// SetNextHop sets field value
func (o *Route4Type) SetNextHop(v string) {
	o.NextHop = v
}

func (o Route4Type) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["destination"] = o.Destination
	}
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}
	if o.SourceIpAddress != nil {
		toSerialize["sourceIpAddress"] = o.SourceIpAddress
	}
	if true {
		toSerialize["nextHop"] = o.NextHop
	}
	return json.Marshal(toSerialize)
}

type NullableRoute4Type struct {
	value *Route4Type
	isSet bool
}

func (v NullableRoute4Type) Get() *Route4Type {
	return v.value
}

func (v *NullableRoute4Type) Set(val *Route4Type) {
	v.value = val
	v.isSet = true
}

func (v NullableRoute4Type) IsSet() bool {
	return v.isSet
}

func (v *NullableRoute4Type) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoute4Type(val *Route4Type) *NullableRoute4Type {
	return &NullableRoute4Type{value: val, isSet: true}
}

func (v NullableRoute4Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoute4Type) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


