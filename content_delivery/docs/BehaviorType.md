# BehaviorType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the behavior, &#x60;udpSyslogAccessLog&#x60; in this case. | 
**Options** | [**UdpSyslogAccessLogOptions**](UdpSyslogAccessLogOptions.md) |  | 

## Methods

### NewBehaviorType

`func NewBehaviorType(name string, options UdpSyslogAccessLogOptions, ) *BehaviorType`

NewBehaviorType instantiates a new BehaviorType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorTypeWithDefaults

`func NewBehaviorTypeWithDefaults() *BehaviorType`

NewBehaviorTypeWithDefaults instantiates a new BehaviorType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BehaviorType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BehaviorType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BehaviorType) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *BehaviorType) GetOptions() UdpSyslogAccessLogOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *BehaviorType) GetOptionsOk() (*UdpSyslogAccessLogOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *BehaviorType) SetOptions(v UdpSyslogAccessLogOptions)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


