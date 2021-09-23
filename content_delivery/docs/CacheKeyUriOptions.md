# CacheKeyUriOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | The cache key URI pattern without query parameters. By default, all incoming query parameters are added to the cache key URI. By default, this value is also used for the purge key URI. | 

## Methods

### NewCacheKeyUriOptions

`func NewCacheKeyUriOptions(value string, ) *CacheKeyUriOptions`

NewCacheKeyUriOptions instantiates a new CacheKeyUriOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheKeyUriOptionsWithDefaults

`func NewCacheKeyUriOptionsWithDefaults() *CacheKeyUriOptions`

NewCacheKeyUriOptionsWithDefaults instantiates a new CacheKeyUriOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *CacheKeyUriOptions) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CacheKeyUriOptions) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CacheKeyUriOptions) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


