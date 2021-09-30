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

// Node2 Node update schema.
type Node2 struct {
	BootState BootStateType `json:"bootState"`
	// List of interfaces for the node.  Bond interfaces `lag0`, `lag1`, `lag2`, and `lag3` must exist.  If not supplied on a POST, they will be added.  A minimum of one and maximum of 32 ethernet interfaces must exist.  One of the ethernet interfaces must contain the management address, and this interface may not be part of a bond.
	Interfaces []InterfaceType `json:"interfaces"`
	// The unique fully qualified domain name (FQDN) for this node.
	Hostname string `json:"hostname"`
	// The unique identifier for a node.
	NodeId int32 `json:"nodeId"`
	// The unique identifier for a site.
	SiteId int32 `json:"siteId"`
	// Array of mappings of attributeType to value for this node.
	Attributes *[]map[string]interface{} `json:"attributes,omitempty"`
	AdministrativeState AdministrativeStateType `json:"administrativeState"`
	// List of DNS Server IP addresses.
	DnsServers []string `json:"dnsServers"`
}

// NewNode2 instantiates a new Node2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNode2(bootState BootStateType, interfaces []InterfaceType, hostname string, nodeId int32, siteId int32, administrativeState AdministrativeStateType, dnsServers []string) *Node2 {
	this := Node2{}
	this.BootState = bootState
	this.Interfaces = interfaces
	this.Hostname = hostname
	this.NodeId = nodeId
	this.SiteId = siteId
	this.AdministrativeState = administrativeState
	this.DnsServers = dnsServers
	return &this
}

// NewNode2WithDefaults instantiates a new Node2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNode2WithDefaults() *Node2 {
	this := Node2{}
	return &this
}

// GetBootState returns the BootState field value
func (o *Node2) GetBootState() BootStateType {
	if o == nil {
		var ret BootStateType
		return ret
	}

	return o.BootState
}

// GetBootStateOk returns a tuple with the BootState field value
// and a boolean to check if the value has been set.
func (o *Node2) GetBootStateOk() (*BootStateType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BootState, true
}

// SetBootState sets field value
func (o *Node2) SetBootState(v BootStateType) {
	o.BootState = v
}

// GetInterfaces returns the Interfaces field value
func (o *Node2) GetInterfaces() []InterfaceType {
	if o == nil {
		var ret []InterfaceType
		return ret
	}

	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value
// and a boolean to check if the value has been set.
func (o *Node2) GetInterfacesOk() (*[]InterfaceType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Interfaces, true
}

// SetInterfaces sets field value
func (o *Node2) SetInterfaces(v []InterfaceType) {
	o.Interfaces = v
}

// GetHostname returns the Hostname field value
func (o *Node2) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *Node2) GetHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *Node2) SetHostname(v string) {
	o.Hostname = v
}

// GetNodeId returns the NodeId field value
func (o *Node2) GetNodeId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *Node2) GetNodeIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *Node2) SetNodeId(v int32) {
	o.NodeId = v
}

// GetSiteId returns the SiteId field value
func (o *Node2) GetSiteId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *Node2) GetSiteIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *Node2) SetSiteId(v int32) {
	o.SiteId = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Node2) GetAttributes() []map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node2) GetAttributesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Node2) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []map[string]interface{} and assigns it to the Attributes field.
func (o *Node2) SetAttributes(v []map[string]interface{}) {
	o.Attributes = &v
}

// GetAdministrativeState returns the AdministrativeState field value
func (o *Node2) GetAdministrativeState() AdministrativeStateType {
	if o == nil {
		var ret AdministrativeStateType
		return ret
	}

	return o.AdministrativeState
}

// GetAdministrativeStateOk returns a tuple with the AdministrativeState field value
// and a boolean to check if the value has been set.
func (o *Node2) GetAdministrativeStateOk() (*AdministrativeStateType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AdministrativeState, true
}

// SetAdministrativeState sets field value
func (o *Node2) SetAdministrativeState(v AdministrativeStateType) {
	o.AdministrativeState = v
}

// GetDnsServers returns the DnsServers field value
func (o *Node2) GetDnsServers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DnsServers
}

// GetDnsServersOk returns a tuple with the DnsServers field value
// and a boolean to check if the value has been set.
func (o *Node2) GetDnsServersOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DnsServers, true
}

// SetDnsServers sets field value
func (o *Node2) SetDnsServers(v []string) {
	o.DnsServers = v
}

func (o Node2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bootState"] = o.BootState
	}
	if true {
		toSerialize["interfaces"] = o.Interfaces
	}
	if true {
		toSerialize["hostname"] = o.Hostname
	}
	if true {
		toSerialize["nodeId"] = o.NodeId
	}
	if true {
		toSerialize["siteId"] = o.SiteId
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["administrativeState"] = o.AdministrativeState
	}
	if true {
		toSerialize["dnsServers"] = o.DnsServers
	}
	return json.Marshal(toSerialize)
}

type NullableNode2 struct {
	value *Node2
	isSet bool
}

func (v NullableNode2) Get() *Node2 {
	return v.value
}

func (v *NullableNode2) Set(val *Node2) {
	v.value = val
	v.isSet = true
}

func (v NullableNode2) IsSet() bool {
	return v.isSet
}

func (v *NullableNode2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNode2(val *Node2) *NullableNode2 {
	return &NullableNode2{value: val, isSet: true}
}

func (v NullableNode2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNode2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


