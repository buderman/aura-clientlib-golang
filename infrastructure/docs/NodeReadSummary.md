# NodeReadSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | The unique fully qualified domain name (FQDN) for this node. | 
**NodeId** | **int32** | The unique identifier for a node. | 

## Methods

### NewNodeReadSummary

`func NewNodeReadSummary(hostname string, nodeId int32, ) *NodeReadSummary`

NewNodeReadSummary instantiates a new NodeReadSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeReadSummaryWithDefaults

`func NewNodeReadSummaryWithDefaults() *NodeReadSummary`

NewNodeReadSummaryWithDefaults instantiates a new NodeReadSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *NodeReadSummary) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NodeReadSummary) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NodeReadSummary) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetNodeId

`func (o *NodeReadSummary) GetNodeId() int32`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NodeReadSummary) GetNodeIdOk() (*int32, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NodeReadSummary) SetNodeId(v int32)`

SetNodeId sets NodeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


