# HttpDelivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the behavior, &#x60;httpDelivery&#x60; in this case. | 
**Options** | [**HttpDeliveryOptions**](HttpDeliveryOptions.md) |  | 

## Methods

### NewHttpDelivery

`func NewHttpDelivery(name string, options HttpDeliveryOptions, ) *HttpDelivery`

NewHttpDelivery instantiates a new HttpDelivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpDeliveryWithDefaults

`func NewHttpDeliveryWithDefaults() *HttpDelivery`

NewHttpDeliveryWithDefaults instantiates a new HttpDelivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HttpDelivery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HttpDelivery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HttpDelivery) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *HttpDelivery) GetOptions() HttpDeliveryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *HttpDelivery) GetOptionsOk() (*HttpDeliveryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *HttpDelivery) SetOptions(v HttpDeliveryOptions)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


