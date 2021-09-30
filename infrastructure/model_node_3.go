/*
infrastructure

Infrastructure API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infrastructure

import (
	"encoding/json"
)

// Node3 A node represents an aggregation of relevant parts that include a list of interfaces, each containing a list of addresses and routes, and a list of attribute values.
type Node3 struct {
	BootState *BootStateType `json:"bootState,omitempty"`
	// List of interfaces for the node.  Bond interfaces `lag0`, `lag1`, `lag2`, and `lag3` must exist.  If not supplied on a POST, they will be added.  A minimum of one and maximum of 32 ethernet interfaces must exist.  One of the ethernet interfaces must contain the management address, and this interface may not be part of a bond.
	Interfaces []InterfaceType `json:"interfaces"`
	// The unique fully qualified domain name (FQDN) for this node.
	Hostname string `json:"hostname"`
	// The unique identifier for a site.
	SiteId int32 `json:"siteId"`
	// Array of mappings of attributeType to value for this node.
	Attributes *[]map[string]interface{} `json:"attributes,omitempty"`
	AdministrativeState *AdministrativeStateType `json:"administrativeState,omitempty"`
	// List of DNS Server IP addresses.
	DnsServers []string `json:"dnsServers"`
}

// NewNode3 instantiates a new Node3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNode3(interfaces []InterfaceType, hostname string, siteId int32, dnsServers []string) *Node3 {
	this := Node3{}
	this.Interfaces = interfaces
	this.Hostname = hostname
	this.SiteId = siteId
	this.DnsServers = dnsServers
	return &this
}

// NewNode3WithDefaults instantiates a new Node3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNode3WithDefaults() *Node3 {
	this := Node3{}
	return &this
}

// GetBootState returns the BootState field value if set, zero value otherwise.
func (o *Node3) GetBootState() BootStateType {
	if o == nil || o.BootState == nil {
		var ret BootStateType
		return ret
	}
	return *o.BootState
}

// GetBootStateOk returns a tuple with the BootState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node3) GetBootStateOk() (*BootStateType, bool) {
	if o == nil || o.BootState == nil {
		return nil, false
	}
	return o.BootState, true
}

// HasBootState returns a boolean if a field has been set.
func (o *Node3) HasBootState() bool {
	if o != nil && o.BootState != nil {
		return true
	}

	return false
}

// SetBootState gets a reference to the given BootStateType and assigns it to the BootState field.
func (o *Node3) SetBootState(v BootStateType) {
	o.BootState = &v
}

// GetInterfaces returns the Interfaces field value
func (o *Node3) GetInterfaces() []InterfaceType {
	if o == nil {
		var ret []InterfaceType
		return ret
	}

	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value
// and a boolean to check if the value has been set.
func (o *Node3) GetInterfacesOk() (*[]InterfaceType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Interfaces, true
}

// SetInterfaces sets field value
func (o *Node3) SetInterfaces(v []InterfaceType) {
	o.Interfaces = v
}

// GetHostname returns the Hostname field value
func (o *Node3) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *Node3) GetHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *Node3) SetHostname(v string) {
	o.Hostname = v
}

// GetSiteId returns the SiteId field value
func (o *Node3) GetSiteId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *Node3) GetSiteIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *Node3) SetSiteId(v int32) {
	o.SiteId = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Node3) GetAttributes() []map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node3) GetAttributesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Node3) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []map[string]interface{} and assigns it to the Attributes field.
func (o *Node3) SetAttributes(v []map[string]interface{}) {
	o.Attributes = &v
}

// GetAdministrativeState returns the AdministrativeState field value if set, zero value otherwise.
func (o *Node3) GetAdministrativeState() AdministrativeStateType {
	if o == nil || o.AdministrativeState == nil {
		var ret AdministrativeStateType
		return ret
	}
	return *o.AdministrativeState
}

// GetAdministrativeStateOk returns a tuple with the AdministrativeState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node3) GetAdministrativeStateOk() (*AdministrativeStateType, bool) {
	if o == nil || o.AdministrativeState == nil {
		return nil, false
	}
	return o.AdministrativeState, true
}

// HasAdministrativeState returns a boolean if a field has been set.
func (o *Node3) HasAdministrativeState() bool {
	if o != nil && o.AdministrativeState != nil {
		return true
	}

	return false
}

// SetAdministrativeState gets a reference to the given AdministrativeStateType and assigns it to the AdministrativeState field.
func (o *Node3) SetAdministrativeState(v AdministrativeStateType) {
	o.AdministrativeState = &v
}

// GetDnsServers returns the DnsServers field value
func (o *Node3) GetDnsServers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DnsServers
}

// GetDnsServersOk returns a tuple with the DnsServers field value
// and a boolean to check if the value has been set.
func (o *Node3) GetDnsServersOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DnsServers, true
}

// SetDnsServers sets field value
func (o *Node3) SetDnsServers(v []string) {
	o.DnsServers = v
}

func (o Node3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BootState != nil {
		toSerialize["bootState"] = o.BootState
	}
	if true {
		toSerialize["interfaces"] = o.Interfaces
	}
	if true {
		toSerialize["hostname"] = o.Hostname
	}
	if true {
		toSerialize["siteId"] = o.SiteId
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.AdministrativeState != nil {
		toSerialize["administrativeState"] = o.AdministrativeState
	}
	if true {
		toSerialize["dnsServers"] = o.DnsServers
	}
	return json.Marshal(toSerialize)
}

type NullableNode3 struct {
	value *Node3
	isSet bool
}

func (v NullableNode3) Get() *Node3 {
	return v.value
}

func (v *NullableNode3) Set(val *Node3) {
	v.value = val
	v.isSet = true
}

func (v NullableNode3) IsSet() bool {
	return v.isSet
}

func (v *NullableNode3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNode3(val *Node3) *NullableNode3 {
	return &NullableNode3{value: val, isSet: true}
}

func (v NullableNode3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNode3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


