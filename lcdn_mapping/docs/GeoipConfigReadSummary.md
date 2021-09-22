# GeoipConfigReadSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeoipConfigId** | Pointer to **int32** | The unique identifier for a GeoIP configuration. | [optional] 
**Type** | [**TypeType**](TypeType.md) |  | 
**Name** | **string** | A unique name for the entity. | 

## Methods

### NewGeoipConfigReadSummary

`func NewGeoipConfigReadSummary(type_ TypeType, name string, ) *GeoipConfigReadSummary`

NewGeoipConfigReadSummary instantiates a new GeoipConfigReadSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoipConfigReadSummaryWithDefaults

`func NewGeoipConfigReadSummaryWithDefaults() *GeoipConfigReadSummary`

NewGeoipConfigReadSummaryWithDefaults instantiates a new GeoipConfigReadSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeoipConfigId

`func (o *GeoipConfigReadSummary) GetGeoipConfigId() int32`

GetGeoipConfigId returns the GeoipConfigId field if non-nil, zero value otherwise.

### GetGeoipConfigIdOk

`func (o *GeoipConfigReadSummary) GetGeoipConfigIdOk() (*int32, bool)`

GetGeoipConfigIdOk returns a tuple with the GeoipConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoipConfigId

`func (o *GeoipConfigReadSummary) SetGeoipConfigId(v int32)`

SetGeoipConfigId sets GeoipConfigId field to given value.

### HasGeoipConfigId

`func (o *GeoipConfigReadSummary) HasGeoipConfigId() bool`

HasGeoipConfigId returns a boolean if a field has been set.

### GetType

`func (o *GeoipConfigReadSummary) GetType() TypeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GeoipConfigReadSummary) GetTypeOk() (*TypeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GeoipConfigReadSummary) SetType(v TypeType)`

SetType sets Type field to given value.


### GetName

`func (o *GeoipConfigReadSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GeoipConfigReadSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GeoipConfigReadSummary) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


