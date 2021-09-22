# GeoipConfigReadDetailed

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

### NewGeoipConfigReadDetailed

`func NewGeoipConfigReadDetailed(name string, geoIpConfigId int32, type_ TypeType, ) *GeoipConfigReadDetailed`

NewGeoipConfigReadDetailed instantiates a new GeoipConfigReadDetailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoipConfigReadDetailedWithDefaults

`func NewGeoipConfigReadDetailedWithDefaults() *GeoipConfigReadDetailed`

NewGeoipConfigReadDetailedWithDefaults instantiates a new GeoipConfigReadDetailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GeoipConfigReadDetailed) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GeoipConfigReadDetailed) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GeoipConfigReadDetailed) SetName(v string)`

SetName sets Name field to given value.


### GetGeoIpConfigId

`func (o *GeoipConfigReadDetailed) GetGeoIpConfigId() int32`

GetGeoIpConfigId returns the GeoIpConfigId field if non-nil, zero value otherwise.

### GetGeoIpConfigIdOk

`func (o *GeoipConfigReadDetailed) GetGeoIpConfigIdOk() (*int32, bool)`

GetGeoIpConfigIdOk returns a tuple with the GeoIpConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoIpConfigId

`func (o *GeoipConfigReadDetailed) SetGeoIpConfigId(v int32)`

SetGeoIpConfigId sets GeoIpConfigId field to given value.


### GetLastUpdated

`func (o *GeoipConfigReadDetailed) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GeoipConfigReadDetailed) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GeoipConfigReadDetailed) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GeoipConfigReadDetailed) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDateCreated

`func (o *GeoipConfigReadDetailed) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GeoipConfigReadDetailed) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GeoipConfigReadDetailed) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GeoipConfigReadDetailed) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetType

`func (o *GeoipConfigReadDetailed) GetType() TypeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GeoipConfigReadDetailed) GetTypeOk() (*TypeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GeoipConfigReadDetailed) SetType(v TypeType)`

SetType sets Type field to given value.


### GetDescription

`func (o *GeoipConfigReadDetailed) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GeoipConfigReadDetailed) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GeoipConfigReadDetailed) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GeoipConfigReadDetailed) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


