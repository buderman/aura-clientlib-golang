# LimitRateOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bandwidth** | **int32** | The maximum bandwidth limit, as specified by the respective unit. | 
**Unit** | Pointer to **string** | Bandwidth unit, either the default &#x60;BYTES_PER_SECOND&#x60;, &#x60;KB_PER_SECOND&#x60;, or &#x60;MB_PER_SECOND&#x60;. | [optional] [default to "BYTES_PER_SECOND"]

## Methods

### NewLimitRateOptions

`func NewLimitRateOptions(bandwidth int32, ) *LimitRateOptions`

NewLimitRateOptions instantiates a new LimitRateOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLimitRateOptionsWithDefaults

`func NewLimitRateOptionsWithDefaults() *LimitRateOptions`

NewLimitRateOptionsWithDefaults instantiates a new LimitRateOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidth

`func (o *LimitRateOptions) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *LimitRateOptions) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *LimitRateOptions) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.


### GetUnit

`func (o *LimitRateOptions) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *LimitRateOptions) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *LimitRateOptions) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *LimitRateOptions) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


