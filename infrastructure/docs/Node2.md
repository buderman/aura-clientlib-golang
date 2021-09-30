# Node2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootState** | [**BootStateType**](BootStateType.md) |  | 
**Interfaces** | [**[]InterfaceType**](InterfaceType.md) | List of interfaces for the node.  Bond interfaces &#x60;lag0&#x60;, &#x60;lag1&#x60;, &#x60;lag2&#x60;, and &#x60;lag3&#x60; must exist.  If not supplied on a POST, they will be added.  A minimum of one and maximum of 32 ethernet interfaces must exist.  One of the ethernet interfaces must contain the management address, and this interface may not be part of a bond. | 
**Hostname** | **string** | The unique fully qualified domain name (FQDN) for this node. | 
**NodeId** | **int32** | The unique identifier for a node. | 
**SiteId** | **int32** | The unique identifier for a site. | 
**Attributes** | Pointer to **[]map[string]interface{}** | Array of mappings of attributeType to value for this node. | [optional] 
**AdministrativeState** | [**AdministrativeStateType**](AdministrativeStateType.md) |  | 
**DnsServers** | **[]string** | List of DNS Server IP addresses. | 

## Methods

### NewNode2

`func NewNode2(bootState BootStateType, interfaces []InterfaceType, hostname string, nodeId int32, siteId int32, administrativeState AdministrativeStateType, dnsServers []string, ) *Node2`

NewNode2 instantiates a new Node2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNode2WithDefaults

`func NewNode2WithDefaults() *Node2`

NewNode2WithDefaults instantiates a new Node2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootState

`func (o *Node2) GetBootState() BootStateType`

GetBootState returns the BootState field if non-nil, zero value otherwise.

### GetBootStateOk

`func (o *Node2) GetBootStateOk() (*BootStateType, bool)`

GetBootStateOk returns a tuple with the BootState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootState

`func (o *Node2) SetBootState(v BootStateType)`

SetBootState sets BootState field to given value.


### GetInterfaces

`func (o *Node2) GetInterfaces() []InterfaceType`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *Node2) GetInterfacesOk() (*[]InterfaceType, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *Node2) SetInterfaces(v []InterfaceType)`

SetInterfaces sets Interfaces field to given value.


### GetHostname

`func (o *Node2) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Node2) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Node2) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetNodeId

`func (o *Node2) GetNodeId() int32`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Node2) GetNodeIdOk() (*int32, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Node2) SetNodeId(v int32)`

SetNodeId sets NodeId field to given value.


### GetSiteId

`func (o *Node2) GetSiteId() int32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Node2) GetSiteIdOk() (*int32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Node2) SetSiteId(v int32)`

SetSiteId sets SiteId field to given value.


### GetAttributes

`func (o *Node2) GetAttributes() []map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Node2) GetAttributesOk() (*[]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Node2) SetAttributes(v []map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Node2) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetAdministrativeState

`func (o *Node2) GetAdministrativeState() AdministrativeStateType`

GetAdministrativeState returns the AdministrativeState field if non-nil, zero value otherwise.

### GetAdministrativeStateOk

`func (o *Node2) GetAdministrativeStateOk() (*AdministrativeStateType, bool)`

GetAdministrativeStateOk returns a tuple with the AdministrativeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeState

`func (o *Node2) SetAdministrativeState(v AdministrativeStateType)`

SetAdministrativeState sets AdministrativeState field to given value.


### GetDnsServers

`func (o *Node2) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *Node2) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *Node2) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


