# ContentProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentProviderId** | **int32** |  | 
**Enable** | Pointer to **bool** |  | [optional] [default to true]
**ServiceProviderId** | **int32** |  | 
**Account** | **string** |  | 
**Name** | **string** |  | 

## Methods

### NewContentProvider

`func NewContentProvider(contentProviderId int32, serviceProviderId int32, account string, name string, ) *ContentProvider`

NewContentProvider instantiates a new ContentProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentProviderWithDefaults

`func NewContentProviderWithDefaults() *ContentProvider`

NewContentProviderWithDefaults instantiates a new ContentProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentProviderId

`func (o *ContentProvider) GetContentProviderId() int32`

GetContentProviderId returns the ContentProviderId field if non-nil, zero value otherwise.

### GetContentProviderIdOk

`func (o *ContentProvider) GetContentProviderIdOk() (*int32, bool)`

GetContentProviderIdOk returns a tuple with the ContentProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentProviderId

`func (o *ContentProvider) SetContentProviderId(v int32)`

SetContentProviderId sets ContentProviderId field to given value.


### GetEnable

`func (o *ContentProvider) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ContentProvider) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ContentProvider) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *ContentProvider) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetServiceProviderId

`func (o *ContentProvider) GetServiceProviderId() int32`

GetServiceProviderId returns the ServiceProviderId field if non-nil, zero value otherwise.

### GetServiceProviderIdOk

`func (o *ContentProvider) GetServiceProviderIdOk() (*int32, bool)`

GetServiceProviderIdOk returns a tuple with the ServiceProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProviderId

`func (o *ContentProvider) SetServiceProviderId(v int32)`

SetServiceProviderId sets ServiceProviderId field to given value.


### GetAccount

`func (o *ContentProvider) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ContentProvider) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ContentProvider) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetName

`func (o *ContentProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentProvider) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


