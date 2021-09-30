# RequestrouterInstanceReadDetailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | The unique fully qualified domain name for a node. | [optional] 
**NodeId** | **int32** | The unique identifier for a Request Router node. | 
**ServiceLabels** | Pointer to [**ServiceLabelsType**](ServiceLabelsType.md) |  | [optional] 

## Methods

### NewRequestrouterInstanceReadDetailed

`func NewRequestrouterInstanceReadDetailed(nodeId int32, ) *RequestrouterInstanceReadDetailed`

NewRequestrouterInstanceReadDetailed instantiates a new RequestrouterInstanceReadDetailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestrouterInstanceReadDetailedWithDefaults

`func NewRequestrouterInstanceReadDetailedWithDefaults() *RequestrouterInstanceReadDetailed`

NewRequestrouterInstanceReadDetailedWithDefaults instantiates a new RequestrouterInstanceReadDetailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *RequestrouterInstanceReadDetailed) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *RequestrouterInstanceReadDetailed) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *RequestrouterInstanceReadDetailed) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *RequestrouterInstanceReadDetailed) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetNodeId

`func (o *RequestrouterInstanceReadDetailed) GetNodeId() int32`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *RequestrouterInstanceReadDetailed) GetNodeIdOk() (*int32, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *RequestrouterInstanceReadDetailed) SetNodeId(v int32)`

SetNodeId sets NodeId field to given value.


### GetServiceLabels

`func (o *RequestrouterInstanceReadDetailed) GetServiceLabels() ServiceLabelsType`

GetServiceLabels returns the ServiceLabels field if non-nil, zero value otherwise.

### GetServiceLabelsOk

`func (o *RequestrouterInstanceReadDetailed) GetServiceLabelsOk() (*ServiceLabelsType, bool)`

GetServiceLabelsOk returns a tuple with the ServiceLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLabels

`func (o *RequestrouterInstanceReadDetailed) SetServiceLabels(v ServiceLabelsType)`

SetServiceLabels sets ServiceLabels field to given value.

### HasServiceLabels

`func (o *RequestrouterInstanceReadDetailed) HasServiceLabels() bool`

HasServiceLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


