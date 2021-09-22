# AccessMap1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MapId** | **int32** | The unique identifier for an entity. | 
**Name** | **string** | A unique name for the entity. | 
**Description** | Pointer to **string** | A description for the entity. | [optional] 

## Methods

### NewAccessMap1

`func NewAccessMap1(mapId int32, name string, ) *AccessMap1`

NewAccessMap1 instantiates a new AccessMap1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessMap1WithDefaults

`func NewAccessMap1WithDefaults() *AccessMap1`

NewAccessMap1WithDefaults instantiates a new AccessMap1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMapId

`func (o *AccessMap1) GetMapId() int32`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *AccessMap1) GetMapIdOk() (*int32, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *AccessMap1) SetMapId(v int32)`

SetMapId sets MapId field to given value.


### GetName

`func (o *AccessMap1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessMap1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessMap1) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AccessMap1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessMap1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessMap1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessMap1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


