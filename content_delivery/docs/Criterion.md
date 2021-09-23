# Criterion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the criterion, &#x60;uriFilter&#x60; in this case. | 
**Options** | [**CriterionOptions**](CriterionOptions.md) |  | 

## Methods

### NewCriterion

`func NewCriterion(name string, options CriterionOptions, ) *Criterion`

NewCriterion instantiates a new Criterion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCriterionWithDefaults

`func NewCriterionWithDefaults() *Criterion`

NewCriterionWithDefaults instantiates a new Criterion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Criterion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Criterion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Criterion) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *Criterion) GetOptions() CriterionOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Criterion) GetOptionsOk() (*CriterionOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Criterion) SetOptions(v CriterionOptions)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


