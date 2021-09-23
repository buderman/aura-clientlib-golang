# CriterionOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UriFilterType** | **string** | The type of the filter to match URIs, either &#x60;EXACT_MATCH&#x60;, &#x60;PREFIX_MATCH&#x60;, &#x60;REGULAR_EXPRESSION_MATCH&#x60;, or &#x60;CASE_INSENSITIVE_REGULAR_EXPRESSION_MATCH&#x60;. | 
**UriFilterValue** | **string** | Value of the filter used to match incoming URIs. | 

## Methods

### NewCriterionOptions

`func NewCriterionOptions(uriFilterType string, uriFilterValue string, ) *CriterionOptions`

NewCriterionOptions instantiates a new CriterionOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCriterionOptionsWithDefaults

`func NewCriterionOptionsWithDefaults() *CriterionOptions`

NewCriterionOptionsWithDefaults instantiates a new CriterionOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUriFilterType

`func (o *CriterionOptions) GetUriFilterType() string`

GetUriFilterType returns the UriFilterType field if non-nil, zero value otherwise.

### GetUriFilterTypeOk

`func (o *CriterionOptions) GetUriFilterTypeOk() (*string, bool)`

GetUriFilterTypeOk returns a tuple with the UriFilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUriFilterType

`func (o *CriterionOptions) SetUriFilterType(v string)`

SetUriFilterType sets UriFilterType field to given value.


### GetUriFilterValue

`func (o *CriterionOptions) GetUriFilterValue() string`

GetUriFilterValue returns the UriFilterValue field if non-nil, zero value otherwise.

### GetUriFilterValueOk

`func (o *CriterionOptions) GetUriFilterValueOk() (*string, bool)`

GetUriFilterValueOk returns a tuple with the UriFilterValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUriFilterValue

`func (o *CriterionOptions) SetUriFilterValue(v string)`

SetUriFilterValue sets UriFilterValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


