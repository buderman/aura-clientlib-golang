# OriginCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | [**PageType**](PageType.md) |  | 
**Origins** | [**[]Origin3**](Origin3.md) |  | 

## Methods

### NewOriginCollection

`func NewOriginCollection(page PageType, origins []Origin3, ) *OriginCollection`

NewOriginCollection instantiates a new OriginCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginCollectionWithDefaults

`func NewOriginCollectionWithDefaults() *OriginCollection`

NewOriginCollectionWithDefaults instantiates a new OriginCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *OriginCollection) GetPage() PageType`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *OriginCollection) GetPageOk() (*PageType, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *OriginCollection) SetPage(v PageType)`

SetPage sets Page field to given value.


### GetOrigins

`func (o *OriginCollection) GetOrigins() []Origin3`

GetOrigins returns the Origins field if non-nil, zero value otherwise.

### GetOriginsOk

`func (o *OriginCollection) GetOriginsOk() (*[]Origin3, bool)`

GetOriginsOk returns a tuple with the Origins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigins

`func (o *OriginCollection) SetOrigins(v []Origin3)`

SetOrigins sets Origins field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


