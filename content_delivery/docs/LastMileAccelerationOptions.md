# LastMileAccelerationOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompressionLevel** | Pointer to **int32** | GZIP compression level, 6 by default. | [optional] [default to 6]
**ContentTypes** | Pointer to **[]string** | List of acceptable MIME-Types for which LMA is enabled. | [optional] 
**Enable** | Pointer to **bool** | Feature enable/disable, or &#x60;null&#x60; to inherit settings from the global HyperCache configuration. | [optional] 
**UaRestrict** | Pointer to **string** | Regular expression match on the &#x60;User-Agent&#x60; header to disable LMA for particular clients. | [optional] 

## Methods

### NewLastMileAccelerationOptions

`func NewLastMileAccelerationOptions() *LastMileAccelerationOptions`

NewLastMileAccelerationOptions instantiates a new LastMileAccelerationOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastMileAccelerationOptionsWithDefaults

`func NewLastMileAccelerationOptionsWithDefaults() *LastMileAccelerationOptions`

NewLastMileAccelerationOptionsWithDefaults instantiates a new LastMileAccelerationOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompressionLevel

`func (o *LastMileAccelerationOptions) GetCompressionLevel() int32`

GetCompressionLevel returns the CompressionLevel field if non-nil, zero value otherwise.

### GetCompressionLevelOk

`func (o *LastMileAccelerationOptions) GetCompressionLevelOk() (*int32, bool)`

GetCompressionLevelOk returns a tuple with the CompressionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionLevel

`func (o *LastMileAccelerationOptions) SetCompressionLevel(v int32)`

SetCompressionLevel sets CompressionLevel field to given value.

### HasCompressionLevel

`func (o *LastMileAccelerationOptions) HasCompressionLevel() bool`

HasCompressionLevel returns a boolean if a field has been set.

### GetContentTypes

`func (o *LastMileAccelerationOptions) GetContentTypes() []string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *LastMileAccelerationOptions) GetContentTypesOk() (*[]string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *LastMileAccelerationOptions) SetContentTypes(v []string)`

SetContentTypes sets ContentTypes field to given value.

### HasContentTypes

`func (o *LastMileAccelerationOptions) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### GetEnable

`func (o *LastMileAccelerationOptions) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *LastMileAccelerationOptions) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *LastMileAccelerationOptions) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *LastMileAccelerationOptions) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetUaRestrict

`func (o *LastMileAccelerationOptions) GetUaRestrict() string`

GetUaRestrict returns the UaRestrict field if non-nil, zero value otherwise.

### GetUaRestrictOk

`func (o *LastMileAccelerationOptions) GetUaRestrictOk() (*string, bool)`

GetUaRestrictOk returns a tuple with the UaRestrict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUaRestrict

`func (o *LastMileAccelerationOptions) SetUaRestrict(v string)`

SetUaRestrict sets UaRestrict field to given value.

### HasUaRestrict

`func (o *LastMileAccelerationOptions) HasUaRestrict() bool`

HasUaRestrict returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


