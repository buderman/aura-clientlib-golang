# Site

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name for this site. | 
**AbbreviatedName** | **string** | A shortened name for this site. | 

## Methods

### NewSite

`func NewSite(name string, abbreviatedName string, ) *Site`

NewSite instantiates a new Site object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteWithDefaults

`func NewSiteWithDefaults() *Site`

NewSiteWithDefaults instantiates a new Site object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Site) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Site) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Site) SetName(v string)`

SetName sets Name field to given value.


### GetAbbreviatedName

`func (o *Site) GetAbbreviatedName() string`

GetAbbreviatedName returns the AbbreviatedName field if non-nil, zero value otherwise.

### GetAbbreviatedNameOk

`func (o *Site) GetAbbreviatedNameOk() (*string, bool)`

GetAbbreviatedNameOk returns a tuple with the AbbreviatedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbbreviatedName

`func (o *Site) SetAbbreviatedName(v string)`

SetAbbreviatedName sets AbbreviatedName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


