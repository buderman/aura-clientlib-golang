# HttpDeliveryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | [**DeliveryOverrideType**](DeliveryOverrideType.md) |  | [default to DELIVERYOVERRIDETYPE_DEFAULT]

## Methods

### NewHttpDeliveryOptions

`func NewHttpDeliveryOptions(value DeliveryOverrideType, ) *HttpDeliveryOptions`

NewHttpDeliveryOptions instantiates a new HttpDeliveryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpDeliveryOptionsWithDefaults

`func NewHttpDeliveryOptionsWithDefaults() *HttpDeliveryOptions`

NewHttpDeliveryOptionsWithDefaults instantiates a new HttpDeliveryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *HttpDeliveryOptions) GetValue() DeliveryOverrideType`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HttpDeliveryOptions) GetValueOk() (*DeliveryOverrideType, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HttpDeliveryOptions) SetValue(v DeliveryOverrideType)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


