# DscpOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **int32** | DiffServ byte field, including the last two ECN bits. When configuring this field, the ECN bits must be set to 0 (zero), and only the upper 6 bits should be set. | 

## Methods

### NewDscpOptions

`func NewDscpOptions(value int32, ) *DscpOptions`

NewDscpOptions instantiates a new DscpOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDscpOptionsWithDefaults

`func NewDscpOptionsWithDefaults() *DscpOptions`

NewDscpOptionsWithDefaults instantiates a new DscpOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *DscpOptions) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DscpOptions) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DscpOptions) SetValue(v int32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


