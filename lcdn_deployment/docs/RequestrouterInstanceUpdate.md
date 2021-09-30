# RequestrouterInstanceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | **int32** | The unique identifier for a Request Router node. | 
**ServiceLabels** | Pointer to [**ServiceLabelsType**](ServiceLabelsType.md) |  | [optional] 

## Methods

### NewRequestrouterInstanceUpdate

`func NewRequestrouterInstanceUpdate(nodeId int32, ) *RequestrouterInstanceUpdate`

NewRequestrouterInstanceUpdate instantiates a new RequestrouterInstanceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestrouterInstanceUpdateWithDefaults

`func NewRequestrouterInstanceUpdateWithDefaults() *RequestrouterInstanceUpdate`

NewRequestrouterInstanceUpdateWithDefaults instantiates a new RequestrouterInstanceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *RequestrouterInstanceUpdate) GetNodeId() int32`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *RequestrouterInstanceUpdate) GetNodeIdOk() (*int32, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *RequestrouterInstanceUpdate) SetNodeId(v int32)`

SetNodeId sets NodeId field to given value.


### GetServiceLabels

`func (o *RequestrouterInstanceUpdate) GetServiceLabels() ServiceLabelsType`

GetServiceLabels returns the ServiceLabels field if non-nil, zero value otherwise.

### GetServiceLabelsOk

`func (o *RequestrouterInstanceUpdate) GetServiceLabelsOk() (*ServiceLabelsType, bool)`

GetServiceLabelsOk returns a tuple with the ServiceLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLabels

`func (o *RequestrouterInstanceUpdate) SetServiceLabels(v ServiceLabelsType)`

SetServiceLabels sets ServiceLabels field to given value.

### HasServiceLabels

`func (o *RequestrouterInstanceUpdate) HasServiceLabels() bool`

HasServiceLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


