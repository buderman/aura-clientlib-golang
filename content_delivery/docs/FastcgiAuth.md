# FastcgiAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the behavior, &#x60;fastcgiAuth&#x60; in this case. | 
**Options** | [**FastcgiAuthOptions**](FastcgiAuthOptions.md) |  | 

## Methods

### NewFastcgiAuth

`func NewFastcgiAuth(name string, options FastcgiAuthOptions, ) *FastcgiAuth`

NewFastcgiAuth instantiates a new FastcgiAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFastcgiAuthWithDefaults

`func NewFastcgiAuthWithDefaults() *FastcgiAuth`

NewFastcgiAuthWithDefaults instantiates a new FastcgiAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FastcgiAuth) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FastcgiAuth) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FastcgiAuth) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *FastcgiAuth) GetOptions() FastcgiAuthOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FastcgiAuth) GetOptionsOk() (*FastcgiAuthOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FastcgiAuth) SetOptions(v FastcgiAuthOptions)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


