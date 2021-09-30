# NodeCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootState** | Pointer to [**BootStateType**](BootStateType.md) |  | [optional] 
**Interfaces** | [**[]InterfaceType**](InterfaceType.md) | List of interfaces for the node.  Bond interfaces &#x60;lag0&#x60;, &#x60;lag1&#x60;, &#x60;lag2&#x60;, and &#x60;lag3&#x60; must exist.  If not supplied on a POST, they will be added.  A minimum of one and maximum of 32 ethernet interfaces must exist.  One of the ethernet interfaces must contain the management address, and this interface may not be part of a bond. | 
**Hostname** | **string** | The unique fully qualified domain name (FQDN) for this node. | 
**SiteId** | **int32** | The unique identifier for a site. | 
**Attributes** | Pointer to **[]map[string]interface{}** | Array of mappings of attributeType to value for this node. | [optional] 
**AdministrativeState** | Pointer to [**AdministrativeStateType**](AdministrativeStateType.md) |  | [optional] 
**DnsServers** | **[]string** | List of DNS Server IP addresses. | 

## Methods

### NewNodeCreate

`func NewNodeCreate(interfaces []InterfaceType, hostname string, siteId int32, dnsServers []string, ) *NodeCreate`

NewNodeCreate instantiates a new NodeCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeCreateWithDefaults

`func NewNodeCreateWithDefaults() *NodeCreate`

NewNodeCreateWithDefaults instantiates a new NodeCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootState

`func (o *NodeCreate) GetBootState() BootStateType`

GetBootState returns the BootState field if non-nil, zero value otherwise.

### GetBootStateOk

`func (o *NodeCreate) GetBootStateOk() (*BootStateType, bool)`

GetBootStateOk returns a tuple with the BootState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootState

`func (o *NodeCreate) SetBootState(v BootStateType)`

SetBootState sets BootState field to given value.

### HasBootState

`func (o *NodeCreate) HasBootState() bool`

HasBootState returns a boolean if a field has been set.

### GetInterfaces

`func (o *NodeCreate) GetInterfaces() []InterfaceType`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *NodeCreate) GetInterfacesOk() (*[]InterfaceType, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *NodeCreate) SetInterfaces(v []InterfaceType)`

SetInterfaces sets Interfaces field to given value.


### GetHostname

`func (o *NodeCreate) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NodeCreate) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NodeCreate) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetSiteId

`func (o *NodeCreate) GetSiteId() int32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *NodeCreate) GetSiteIdOk() (*int32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *NodeCreate) SetSiteId(v int32)`

SetSiteId sets SiteId field to given value.


### GetAttributes

`func (o *NodeCreate) GetAttributes() []map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NodeCreate) GetAttributesOk() (*[]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NodeCreate) SetAttributes(v []map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *NodeCreate) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetAdministrativeState

`func (o *NodeCreate) GetAdministrativeState() AdministrativeStateType`

GetAdministrativeState returns the AdministrativeState field if non-nil, zero value otherwise.

### GetAdministrativeStateOk

`func (o *NodeCreate) GetAdministrativeStateOk() (*AdministrativeStateType, bool)`

GetAdministrativeStateOk returns a tuple with the AdministrativeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeState

`func (o *NodeCreate) SetAdministrativeState(v AdministrativeStateType)`

SetAdministrativeState sets AdministrativeState field to given value.

### HasAdministrativeState

`func (o *NodeCreate) HasAdministrativeState() bool`

HasAdministrativeState returns a boolean if a field has been set.

### GetDnsServers

`func (o *NodeCreate) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *NodeCreate) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *NodeCreate) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


