# SiteUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | **int32** | The unique identifier for a site. | 
**Name** | **string** | The unique name for this site. | 
**AbbreviatedName** | **string** | A shortened name for this site. | 

## Methods

### NewSiteUpdate

`func NewSiteUpdate(siteId int32, name string, abbreviatedName string, ) *SiteUpdate`

NewSiteUpdate instantiates a new SiteUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteUpdateWithDefaults

`func NewSiteUpdateWithDefaults() *SiteUpdate`

NewSiteUpdateWithDefaults instantiates a new SiteUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *SiteUpdate) GetSiteId() int32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SiteUpdate) GetSiteIdOk() (*int32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SiteUpdate) SetSiteId(v int32)`

SetSiteId sets SiteId field to given value.


### GetName

`func (o *SiteUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetAbbreviatedName

`func (o *SiteUpdate) GetAbbreviatedName() string`

GetAbbreviatedName returns the AbbreviatedName field if non-nil, zero value otherwise.

### GetAbbreviatedNameOk

`func (o *SiteUpdate) GetAbbreviatedNameOk() (*string, bool)`

GetAbbreviatedNameOk returns a tuple with the AbbreviatedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbbreviatedName

`func (o *SiteUpdate) SetAbbreviatedName(v string)`

SetAbbreviatedName sets AbbreviatedName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


