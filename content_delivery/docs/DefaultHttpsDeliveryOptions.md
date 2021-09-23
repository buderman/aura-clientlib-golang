# DefaultHttpsDeliveryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TlsDeliveryProfileId** | Pointer to **int32** |  | [optional] 
**Enable** | **bool** | Enables default HTTPS Delivery for this prefix. | 

## Methods

### NewDefaultHttpsDeliveryOptions

`func NewDefaultHttpsDeliveryOptions(enable bool, ) *DefaultHttpsDeliveryOptions`

NewDefaultHttpsDeliveryOptions instantiates a new DefaultHttpsDeliveryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultHttpsDeliveryOptionsWithDefaults

`func NewDefaultHttpsDeliveryOptionsWithDefaults() *DefaultHttpsDeliveryOptions`

NewDefaultHttpsDeliveryOptionsWithDefaults instantiates a new DefaultHttpsDeliveryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTlsDeliveryProfileId

`func (o *DefaultHttpsDeliveryOptions) GetTlsDeliveryProfileId() int32`

GetTlsDeliveryProfileId returns the TlsDeliveryProfileId field if non-nil, zero value otherwise.

### GetTlsDeliveryProfileIdOk

`func (o *DefaultHttpsDeliveryOptions) GetTlsDeliveryProfileIdOk() (*int32, bool)`

GetTlsDeliveryProfileIdOk returns a tuple with the TlsDeliveryProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsDeliveryProfileId

`func (o *DefaultHttpsDeliveryOptions) SetTlsDeliveryProfileId(v int32)`

SetTlsDeliveryProfileId sets TlsDeliveryProfileId field to given value.

### HasTlsDeliveryProfileId

`func (o *DefaultHttpsDeliveryOptions) HasTlsDeliveryProfileId() bool`

HasTlsDeliveryProfileId returns a boolean if a field has been set.

### GetEnable

`func (o *DefaultHttpsDeliveryOptions) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *DefaultHttpsDeliveryOptions) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *DefaultHttpsDeliveryOptions) SetEnable(v bool)`

SetEnable sets Enable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


