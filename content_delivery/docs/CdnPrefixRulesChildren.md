# CdnPrefixRulesChildren

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Behaviors** | [**[]BehaviorType**](BehaviorType.md) |  | 
**Criterion** | [**Criterion**](Criterion.md) |  | 
**Name** | **string** | Name of the rule. | 

## Methods

### NewCdnPrefixRulesChildren

`func NewCdnPrefixRulesChildren(behaviors []BehaviorType, criterion Criterion, name string, ) *CdnPrefixRulesChildren`

NewCdnPrefixRulesChildren instantiates a new CdnPrefixRulesChildren object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnPrefixRulesChildrenWithDefaults

`func NewCdnPrefixRulesChildrenWithDefaults() *CdnPrefixRulesChildren`

NewCdnPrefixRulesChildrenWithDefaults instantiates a new CdnPrefixRulesChildren object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBehaviors

`func (o *CdnPrefixRulesChildren) GetBehaviors() []BehaviorType`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *CdnPrefixRulesChildren) GetBehaviorsOk() (*[]BehaviorType, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *CdnPrefixRulesChildren) SetBehaviors(v []BehaviorType)`

SetBehaviors sets Behaviors field to given value.


### GetCriterion

`func (o *CdnPrefixRulesChildren) GetCriterion() Criterion`

GetCriterion returns the Criterion field if non-nil, zero value otherwise.

### GetCriterionOk

`func (o *CdnPrefixRulesChildren) GetCriterionOk() (*Criterion, bool)`

GetCriterionOk returns a tuple with the Criterion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriterion

`func (o *CdnPrefixRulesChildren) SetCriterion(v Criterion)`

SetCriterion sets Criterion field to given value.


### GetName

`func (o *CdnPrefixRulesChildren) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CdnPrefixRulesChildren) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CdnPrefixRulesChildren) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


