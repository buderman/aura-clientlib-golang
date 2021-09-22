# GeoipConfig1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpConfigs** | Pointer to [**[]GeoipConfig2**](GeoipConfig2.md) |  | [optional] 
**Page** | [**PageType**](PageType.md) |  | 

## Methods

### NewGeoipConfig1

`func NewGeoipConfig1(page PageType, ) *GeoipConfig1`

NewGeoipConfig1 instantiates a new GeoipConfig1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoipConfig1WithDefaults

`func NewGeoipConfig1WithDefaults() *GeoipConfig1`

NewGeoipConfig1WithDefaults instantiates a new GeoipConfig1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpConfigs

`func (o *GeoipConfig1) GetBgpConfigs() []GeoipConfig2`

GetBgpConfigs returns the BgpConfigs field if non-nil, zero value otherwise.

### GetBgpConfigsOk

`func (o *GeoipConfig1) GetBgpConfigsOk() (*[]GeoipConfig2, bool)`

GetBgpConfigsOk returns a tuple with the BgpConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpConfigs

`func (o *GeoipConfig1) SetBgpConfigs(v []GeoipConfig2)`

SetBgpConfigs sets BgpConfigs field to given value.

### HasBgpConfigs

`func (o *GeoipConfig1) HasBgpConfigs() bool`

HasBgpConfigs returns a boolean if a field has been set.

### GetPage

`func (o *GeoipConfig1) GetPage() PageType`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GeoipConfig1) GetPageOk() (*PageType, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GeoipConfig1) SetPage(v PageType)`

SetPage sets Page field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


