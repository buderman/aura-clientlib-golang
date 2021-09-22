# GeoipConfigCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**TypeType**](TypeType.md) |  | 
**Name** | **string** | A unique name for the entity. | 
**Description** | Pointer to **string** | A description for the entity. | [optional] 

## Methods

### NewGeoipConfigCreate

`func NewGeoipConfigCreate(type_ TypeType, name string, ) *GeoipConfigCreate`

NewGeoipConfigCreate instantiates a new GeoipConfigCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoipConfigCreateWithDefaults

`func NewGeoipConfigCreateWithDefaults() *GeoipConfigCreate`

NewGeoipConfigCreateWithDefaults instantiates a new GeoipConfigCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GeoipConfigCreate) GetType() TypeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GeoipConfigCreate) GetTypeOk() (*TypeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GeoipConfigCreate) SetType(v TypeType)`

SetType sets Type field to given value.


### GetName

`func (o *GeoipConfigCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GeoipConfigCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GeoipConfigCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GeoipConfigCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GeoipConfigCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GeoipConfigCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GeoipConfigCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


