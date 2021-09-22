# BgpConfigUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpConfigId** | **int32** | The unique identifier for a BGP configuration. | 
**Name** | **string** | A unique name for the entity. | 
**Description** | Pointer to **string** | A description for the entity. | [optional] 

## Methods

### NewBgpConfigUpdate

`func NewBgpConfigUpdate(bgpConfigId int32, name string, ) *BgpConfigUpdate`

NewBgpConfigUpdate instantiates a new BgpConfigUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpConfigUpdateWithDefaults

`func NewBgpConfigUpdateWithDefaults() *BgpConfigUpdate`

NewBgpConfigUpdateWithDefaults instantiates a new BgpConfigUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpConfigId

`func (o *BgpConfigUpdate) GetBgpConfigId() int32`

GetBgpConfigId returns the BgpConfigId field if non-nil, zero value otherwise.

### GetBgpConfigIdOk

`func (o *BgpConfigUpdate) GetBgpConfigIdOk() (*int32, bool)`

GetBgpConfigIdOk returns a tuple with the BgpConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpConfigId

`func (o *BgpConfigUpdate) SetBgpConfigId(v int32)`

SetBgpConfigId sets BgpConfigId field to given value.


### GetName

`func (o *BgpConfigUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BgpConfigUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BgpConfigUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BgpConfigUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BgpConfigUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BgpConfigUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BgpConfigUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


