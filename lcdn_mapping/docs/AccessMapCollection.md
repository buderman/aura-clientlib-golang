# AccessMapCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Maps** | [**[]AccessMap**](AccessMap.md) |  | 
**Page** | [**PageType**](PageType.md) |  | 

## Methods

### NewAccessMapCollection

`func NewAccessMapCollection(maps []AccessMap, page PageType, ) *AccessMapCollection`

NewAccessMapCollection instantiates a new AccessMapCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessMapCollectionWithDefaults

`func NewAccessMapCollectionWithDefaults() *AccessMapCollection`

NewAccessMapCollectionWithDefaults instantiates a new AccessMapCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaps

`func (o *AccessMapCollection) GetMaps() []AccessMap`

GetMaps returns the Maps field if non-nil, zero value otherwise.

### GetMapsOk

`func (o *AccessMapCollection) GetMapsOk() (*[]AccessMap, bool)`

GetMapsOk returns a tuple with the Maps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaps

`func (o *AccessMapCollection) SetMaps(v []AccessMap)`

SetMaps sets Maps field to given value.


### GetPage

`func (o *AccessMapCollection) GetPage() PageType`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *AccessMapCollection) GetPageOk() (*PageType, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *AccessMapCollection) SetPage(v PageType)`

SetPage sets Page field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


