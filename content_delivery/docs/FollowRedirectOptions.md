# FollowRedirectOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | When enabled, pass on HTTP 302 redirects to the client. | 

## Methods

### NewFollowRedirectOptions

`func NewFollowRedirectOptions(enable bool, ) *FollowRedirectOptions`

NewFollowRedirectOptions instantiates a new FollowRedirectOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFollowRedirectOptionsWithDefaults

`func NewFollowRedirectOptionsWithDefaults() *FollowRedirectOptions`

NewFollowRedirectOptionsWithDefaults instantiates a new FollowRedirectOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *FollowRedirectOptions) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *FollowRedirectOptions) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *FollowRedirectOptions) SetEnable(v bool)`

SetEnable sets Enable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

