# AttributeTypeUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeTypeId** | **int32** | The unique identifier for an attribute type. | 
**Name** | **string** | The unique name for an attribute type. | 

## Methods

### NewAttributeTypeUpdate

`func NewAttributeTypeUpdate(attributeTypeId int32, name string, ) *AttributeTypeUpdate`

NewAttributeTypeUpdate instantiates a new AttributeTypeUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeTypeUpdateWithDefaults

`func NewAttributeTypeUpdateWithDefaults() *AttributeTypeUpdate`

NewAttributeTypeUpdateWithDefaults instantiates a new AttributeTypeUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeTypeId

`func (o *AttributeTypeUpdate) GetAttributeTypeId() int32`

GetAttributeTypeId returns the AttributeTypeId field if non-nil, zero value otherwise.

### GetAttributeTypeIdOk

`func (o *AttributeTypeUpdate) GetAttributeTypeIdOk() (*int32, bool)`

GetAttributeTypeIdOk returns a tuple with the AttributeTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeTypeId

`func (o *AttributeTypeUpdate) SetAttributeTypeId(v int32)`

SetAttributeTypeId sets AttributeTypeId field to given value.


### GetName

`func (o *AttributeTypeUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttributeTypeUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttributeTypeUpdate) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


