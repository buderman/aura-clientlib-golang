# ServiceLabelsType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientServing** | Pointer to **[]string** | List of IP addresses you can specify for a node. The maximum length is 1 and the minimum length is 2048. | [optional] 

## Methods

### NewServiceLabelsType

`func NewServiceLabelsType() *ServiceLabelsType`

NewServiceLabelsType instantiates a new ServiceLabelsType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceLabelsTypeWithDefaults

`func NewServiceLabelsTypeWithDefaults() *ServiceLabelsType`

NewServiceLabelsTypeWithDefaults instantiates a new ServiceLabelsType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientServing

`func (o *ServiceLabelsType) GetClientServing() []string`

GetClientServing returns the ClientServing field if non-nil, zero value otherwise.

### GetClientServingOk

`func (o *ServiceLabelsType) GetClientServingOk() (*[]string, bool)`

GetClientServingOk returns a tuple with the ClientServing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientServing

`func (o *ServiceLabelsType) SetClientServing(v []string)`

SetClientServing sets ClientServing field to given value.

### HasClientServing

`func (o *ServiceLabelsType) HasClientServing() bool`

HasClientServing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


