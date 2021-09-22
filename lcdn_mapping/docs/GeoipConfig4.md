# GeoipConfig4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeoIpConfigId** | **int32** | The unique identifier for a GeoIP configuration. | 
**Type** | [**TypeType**](TypeType.md) |  | 
**Name** | **string** | A unique name for the entity. | 
**Description** | Pointer to **string** | A description for the entity. | [optional] 

## Methods

### NewGeoipConfig4

`func NewGeoipConfig4(geoIpConfigId int32, type_ TypeType, name string, ) *GeoipConfig4`

NewGeoipConfig4 instantiates a new GeoipConfig4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoipConfig4WithDefaults

`func NewGeoipConfig4WithDefaults() *GeoipConfig4`

NewGeoipConfig4WithDefaults instantiates a new GeoipConfig4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeoIpConfigId

`func (o *GeoipConfig4) GetGeoIpConfigId() int32`

GetGeoIpConfigId returns the GeoIpConfigId field if non-nil, zero value otherwise.

### GetGeoIpConfigIdOk

`func (o *GeoipConfig4) GetGeoIpConfigIdOk() (*int32, bool)`

GetGeoIpConfigIdOk returns a tuple with the GeoIpConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoIpConfigId

`func (o *GeoipConfig4) SetGeoIpConfigId(v int32)`

SetGeoIpConfigId sets GeoIpConfigId field to given value.


### GetType

`func (o *GeoipConfig4) GetType() TypeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GeoipConfig4) GetTypeOk() (*TypeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GeoipConfig4) SetType(v TypeType)`

SetType sets Type field to given value.


### GetName

`func (o *GeoipConfig4) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GeoipConfig4) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GeoipConfig4) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GeoipConfig4) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GeoipConfig4) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GeoipConfig4) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GeoipConfig4) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


