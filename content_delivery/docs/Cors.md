# Cors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the behavior, &#x60;cors&#x60; in this case. | 
**Options** | [**CorsOptions**](CorsOptions.md) |  | 

## Methods

### NewCors

`func NewCors(name string, options CorsOptions, ) *Cors`

NewCors instantiates a new Cors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorsWithDefaults

`func NewCorsWithDefaults() *Cors`

NewCorsWithDefaults instantiates a new Cors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Cors) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cors) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cors) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *Cors) GetOptions() CorsOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Cors) GetOptionsOk() (*CorsOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Cors) SetOptions(v CorsOptions)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


