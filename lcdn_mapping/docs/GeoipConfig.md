# GeoipConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name for the entity. | 
**GeoIpConfigId** | **int32** | The unique identifier for a GeoIP configuration. | 
**LastUpdated** | Pointer to **time.Time** | Timestamp of when the configuration was last updated. | [optional] [readonly] 
**DateCreated** | Pointer to **time.Time** | Timestamp of when the configuration was created. | [optional] [readonly] 
**Type** | [**TypeType**](TypeType.md) |  | 
**Description** | Pointer to **string** | A description for the entity. | [optional] 

## Methods

### NewGeoipConfig

`func NewGeoipConfig(name string, geoIpConfigId int32, type_ TypeType, ) *GeoipConfig`

NewGeoipConfig instantiates a new GeoipConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoipConfigWithDefaults

`func NewGeoipConfigWithDefaults() *GeoipConfig`

NewGeoipConfigWithDefaults instantiates a new GeoipConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GeoipConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GeoipConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GeoipConfig) SetName(v string)`

SetName sets Name field to given value.


### GetGeoIpConfigId

`func (o *GeoipConfig) GetGeoIpConfigId() int32`

GetGeoIpConfigId returns the GeoIpConfigId field if non-nil, zero value otherwise.

### GetGeoIpConfigIdOk

`func (o *GeoipConfig) GetGeoIpConfigIdOk() (*int32, bool)`

GetGeoIpConfigIdOk returns a tuple with the GeoIpConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoIpConfigId

`func (o *GeoipConfig) SetGeoIpConfigId(v int32)`

SetGeoIpConfigId sets GeoIpConfigId field to given value.


### GetLastUpdated

`func (o *GeoipConfig) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GeoipConfig) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GeoipConfig) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GeoipConfig) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDateCreated

`func (o *GeoipConfig) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GeoipConfig) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GeoipConfig) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GeoipConfig) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetType

`func (o *GeoipConfig) GetType() TypeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GeoipConfig) GetTypeOk() (*TypeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GeoipConfig) SetType(v TypeType)`

SetType sets Type field to given value.


### GetDescription

`func (o *GeoipConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GeoipConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GeoipConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GeoipConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


