# SiteCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sites** | [**[]Site2**](Site2.md) |  | 
**Page** | [**PageType**](PageType.md) |  | 

## Methods

### NewSiteCollection

`func NewSiteCollection(sites []Site2, page PageType, ) *SiteCollection`

NewSiteCollection instantiates a new SiteCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteCollectionWithDefaults

`func NewSiteCollectionWithDefaults() *SiteCollection`

NewSiteCollectionWithDefaults instantiates a new SiteCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSites

`func (o *SiteCollection) GetSites() []Site2`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *SiteCollection) GetSitesOk() (*[]Site2, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *SiteCollection) SetSites(v []Site2)`

SetSites sets Sites field to given value.


### GetPage

`func (o *SiteCollection) GetPage() PageType`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SiteCollection) GetPageOk() (*PageType, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SiteCollection) SetPage(v PageType)`

SetPage sets Page field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


