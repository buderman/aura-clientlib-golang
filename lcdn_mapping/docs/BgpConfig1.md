# BgpConfig1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpConfigId** | **int32** | The unique identifier for a BGP configuration. | 
**DateCreated** | Pointer to **time.Time** | Timestamp of when the configuration was created. | [optional] [readonly] 
**Name** | **string** | A unique name for the entity. | 
**LastUpdated** | Pointer to **time.Time** | Timestamp of when the configuration was last updated. | [optional] [readonly] 
**Description** | Pointer to **string** | A description for the entity. | [optional] 

## Methods

### NewBgpConfig1

`func NewBgpConfig1(bgpConfigId int32, name string, ) *BgpConfig1`

NewBgpConfig1 instantiates a new BgpConfig1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpConfig1WithDefaults

`func NewBgpConfig1WithDefaults() *BgpConfig1`

NewBgpConfig1WithDefaults instantiates a new BgpConfig1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpConfigId

`func (o *BgpConfig1) GetBgpConfigId() int32`

GetBgpConfigId returns the BgpConfigId field if non-nil, zero value otherwise.

### GetBgpConfigIdOk

`func (o *BgpConfig1) GetBgpConfigIdOk() (*int32, bool)`

GetBgpConfigIdOk returns a tuple with the BgpConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpConfigId

`func (o *BgpConfig1) SetBgpConfigId(v int32)`

SetBgpConfigId sets BgpConfigId field to given value.


### GetDateCreated

`func (o *BgpConfig1) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *BgpConfig1) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *BgpConfig1) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *BgpConfig1) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetName

`func (o *BgpConfig1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BgpConfig1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BgpConfig1) SetName(v string)`

SetName sets Name field to given value.


### GetLastUpdated

`func (o *BgpConfig1) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *BgpConfig1) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *BgpConfig1) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *BgpConfig1) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDescription

`func (o *BgpConfig1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BgpConfig1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BgpConfig1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BgpConfig1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


