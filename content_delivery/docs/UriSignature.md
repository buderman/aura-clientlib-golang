# UriSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the behavior, &#x60;uriSignature&#x60; in this case. | 
**Options** | [**UriSignatureOptions**](UriSignatureOptions.md) |  | 

## Methods

### NewUriSignature

`func NewUriSignature(name string, options UriSignatureOptions, ) *UriSignature`

NewUriSignature instantiates a new UriSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUriSignatureWithDefaults

`func NewUriSignatureWithDefaults() *UriSignature`

NewUriSignatureWithDefaults instantiates a new UriSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UriSignature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UriSignature) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UriSignature) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *UriSignature) GetOptions() UriSignatureOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UriSignature) GetOptionsOk() (*UriSignatureOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UriSignature) SetOptions(v UriSignatureOptions)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


