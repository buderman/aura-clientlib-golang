# BgpConfigReadSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpConfigId** | **int32** | The unique identifier for a BGP configuration. | 
**Name** | **string** | A unique name for the entity. | 

## Methods

### NewBgpConfigReadSummary

`func NewBgpConfigReadSummary(bgpConfigId int32, name string, ) *BgpConfigReadSummary`

NewBgpConfigReadSummary instantiates a new BgpConfigReadSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpConfigReadSummaryWithDefaults

`func NewBgpConfigReadSummaryWithDefaults() *BgpConfigReadSummary`

NewBgpConfigReadSummaryWithDefaults instantiates a new BgpConfigReadSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpConfigId

`func (o *BgpConfigReadSummary) GetBgpConfigId() int32`

GetBgpConfigId returns the BgpConfigId field if non-nil, zero value otherwise.

### GetBgpConfigIdOk

`func (o *BgpConfigReadSummary) GetBgpConfigIdOk() (*int32, bool)`

GetBgpConfigIdOk returns a tuple with the BgpConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpConfigId

`func (o *BgpConfigReadSummary) SetBgpConfigId(v int32)`

SetBgpConfigId sets BgpConfigId field to given value.


### GetName

`func (o *BgpConfigReadSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BgpConfigReadSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BgpConfigReadSummary) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


