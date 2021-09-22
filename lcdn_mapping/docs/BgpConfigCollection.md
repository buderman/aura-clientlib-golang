# BgpConfigCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpConfigs** | [**[]BgpConfig3**](BgpConfig3.md) |  | 
**Page** | [**PageType**](PageType.md) |  | 

## Methods

### NewBgpConfigCollection

`func NewBgpConfigCollection(bgpConfigs []BgpConfig3, page PageType, ) *BgpConfigCollection`

NewBgpConfigCollection instantiates a new BgpConfigCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpConfigCollectionWithDefaults

`func NewBgpConfigCollectionWithDefaults() *BgpConfigCollection`

NewBgpConfigCollectionWithDefaults instantiates a new BgpConfigCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpConfigs

`func (o *BgpConfigCollection) GetBgpConfigs() []BgpConfig3`

GetBgpConfigs returns the BgpConfigs field if non-nil, zero value otherwise.

### GetBgpConfigsOk

`func (o *BgpConfigCollection) GetBgpConfigsOk() (*[]BgpConfig3, bool)`

GetBgpConfigsOk returns a tuple with the BgpConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpConfigs

`func (o *BgpConfigCollection) SetBgpConfigs(v []BgpConfig3)`

SetBgpConfigs sets BgpConfigs field to given value.


### GetPage

`func (o *BgpConfigCollection) GetPage() PageType`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *BgpConfigCollection) GetPageOk() (*PageType, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *BgpConfigCollection) SetPage(v PageType)`

SetPage sets Page field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


