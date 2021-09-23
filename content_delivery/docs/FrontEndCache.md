# FrontEndCache

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the behavior, &#x60;frontEndCache&#x60; in this case. | 
**Options** | [**DefaultFrontEndCacheOptions**](DefaultFrontEndCacheOptions.md) |  | 

## Methods

### NewFrontEndCache

`func NewFrontEndCache(name string, options DefaultFrontEndCacheOptions, ) *FrontEndCache`

NewFrontEndCache instantiates a new FrontEndCache object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrontEndCacheWithDefaults

`func NewFrontEndCacheWithDefaults() *FrontEndCache`

NewFrontEndCacheWithDefaults instantiates a new FrontEndCache object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FrontEndCache) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FrontEndCache) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FrontEndCache) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *FrontEndCache) GetOptions() DefaultFrontEndCacheOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FrontEndCache) GetOptionsOk() (*DefaultFrontEndCacheOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FrontEndCache) SetOptions(v DefaultFrontEndCacheOptions)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


