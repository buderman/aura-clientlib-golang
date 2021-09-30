# RequestrouterInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | **int32** | The unique identifier for a Request Router node. | 
**ServiceLabels** | Pointer to [**ServiceLabelsType**](ServiceLabelsType.md) |  | [optional] 

## Methods

### NewRequestrouterInstance

`func NewRequestrouterInstance(nodeId int32, ) *RequestrouterInstance`

NewRequestrouterInstance instantiates a new RequestrouterInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestrouterInstanceWithDefaults

`func NewRequestrouterInstanceWithDefaults() *RequestrouterInstance`

NewRequestrouterInstanceWithDefaults instantiates a new RequestrouterInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *RequestrouterInstance) GetNodeId() int32`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *RequestrouterInstance) GetNodeIdOk() (*int32, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *RequestrouterInstance) SetNodeId(v int32)`

SetNodeId sets NodeId field to given value.


### GetServiceLabels

`func (o *RequestrouterInstance) GetServiceLabels() ServiceLabelsType`

GetServiceLabels returns the ServiceLabels field if non-nil, zero value otherwise.

### GetServiceLabelsOk

`func (o *RequestrouterInstance) GetServiceLabelsOk() (*ServiceLabelsType, bool)`

GetServiceLabelsOk returns a tuple with the ServiceLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLabels

`func (o *RequestrouterInstance) SetServiceLabels(v ServiceLabelsType)`

SetServiceLabels sets ServiceLabels field to given value.

### HasServiceLabels

`func (o *RequestrouterInstance) HasServiceLabels() bool`

HasServiceLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


