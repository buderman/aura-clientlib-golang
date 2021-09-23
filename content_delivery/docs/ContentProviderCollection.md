# ContentProviderCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | [**PageType**](PageType.md) |  | 
**ContentProviders** | [**[]ContentProvider2**](ContentProvider2.md) |  | 

## Methods

### NewContentProviderCollection

`func NewContentProviderCollection(page PageType, contentProviders []ContentProvider2, ) *ContentProviderCollection`

NewContentProviderCollection instantiates a new ContentProviderCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentProviderCollectionWithDefaults

`func NewContentProviderCollectionWithDefaults() *ContentProviderCollection`

NewContentProviderCollectionWithDefaults instantiates a new ContentProviderCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ContentProviderCollection) GetPage() PageType`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ContentProviderCollection) GetPageOk() (*PageType, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ContentProviderCollection) SetPage(v PageType)`

SetPage sets Page field to given value.


### GetContentProviders

`func (o *ContentProviderCollection) GetContentProviders() []ContentProvider2`

GetContentProviders returns the ContentProviders field if non-nil, zero value otherwise.

### GetContentProvidersOk

`func (o *ContentProviderCollection) GetContentProvidersOk() (*[]ContentProvider2, bool)`

GetContentProvidersOk returns a tuple with the ContentProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentProviders

`func (o *ContentProviderCollection) SetContentProviders(v []ContentProvider2)`

SetContentProviders sets ContentProviders field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


