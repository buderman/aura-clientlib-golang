# Dscp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the behavior, &#x60;dscp&#x60; in this case. | 
**Options** | [**DscpOptions**](DscpOptions.md) |  | 

## Methods

### NewDscp

`func NewDscp(name string, options DscpOptions, ) *Dscp`

NewDscp instantiates a new Dscp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDscpWithDefaults

`func NewDscpWithDefaults() *Dscp`

NewDscpWithDefaults instantiates a new Dscp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Dscp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dscp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dscp) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *Dscp) GetOptions() DscpOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Dscp) GetOptionsOk() (*DscpOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Dscp) SetOptions(v DscpOptions)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


