# Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultBehaviors** | Pointer to [**[]DefaultBehaviorType**](DefaultBehaviorType.md) |  | [optional] 
**Children** | Pointer to [**[]CdnPrefixRulesChildren**](CdnPrefixRulesChildren.md) | A list of child rules. Each child rule object consists of its name, the criterion to match URIs, and a list of behaviors that act upon that criterion. | [optional] 

## Methods

### NewRules

`func NewRules() *Rules`

NewRules instantiates a new Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesWithDefaults

`func NewRulesWithDefaults() *Rules`

NewRulesWithDefaults instantiates a new Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultBehaviors

`func (o *Rules) GetDefaultBehaviors() []DefaultBehaviorType`

GetDefaultBehaviors returns the DefaultBehaviors field if non-nil, zero value otherwise.

### GetDefaultBehaviorsOk

`func (o *Rules) GetDefaultBehaviorsOk() (*[]DefaultBehaviorType, bool)`

GetDefaultBehaviorsOk returns a tuple with the DefaultBehaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBehaviors

`func (o *Rules) SetDefaultBehaviors(v []DefaultBehaviorType)`

SetDefaultBehaviors sets DefaultBehaviors field to given value.

### HasDefaultBehaviors

`func (o *Rules) HasDefaultBehaviors() bool`

HasDefaultBehaviors returns a boolean if a field has been set.

### GetChildren

`func (o *Rules) GetChildren() []CdnPrefixRulesChildren`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Rules) GetChildrenOk() (*[]CdnPrefixRulesChildren, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Rules) SetChildren(v []CdnPrefixRulesChildren)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *Rules) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


