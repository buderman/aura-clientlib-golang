# SharedSecretSetCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentProviderId** | **int32** |  | 
**Name** | **string** |  | 

## Methods

### NewSharedSecretSetCreate

`func NewSharedSecretSetCreate(contentProviderId int32, name string, ) *SharedSecretSetCreate`

NewSharedSecretSetCreate instantiates a new SharedSecretSetCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedSecretSetCreateWithDefaults

`func NewSharedSecretSetCreateWithDefaults() *SharedSecretSetCreate`

NewSharedSecretSetCreateWithDefaults instantiates a new SharedSecretSetCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentProviderId

`func (o *SharedSecretSetCreate) GetContentProviderId() int32`

GetContentProviderId returns the ContentProviderId field if non-nil, zero value otherwise.

### GetContentProviderIdOk

`func (o *SharedSecretSetCreate) GetContentProviderIdOk() (*int32, bool)`

GetContentProviderIdOk returns a tuple with the ContentProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentProviderId

`func (o *SharedSecretSetCreate) SetContentProviderId(v int32)`

SetContentProviderId sets ContentProviderId field to given value.


### GetName

`func (o *SharedSecretSetCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedSecretSetCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedSecretSetCreate) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


