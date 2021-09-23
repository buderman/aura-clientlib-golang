# OriginServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the behavior, &#x60;originServer&#x60; in this case. | 
**Options** | [**DefaultOriginServerOptions**](DefaultOriginServerOptions.md) |  | 

## Methods

### NewOriginServer

`func NewOriginServer(name string, options DefaultOriginServerOptions, ) *OriginServer`

NewOriginServer instantiates a new OriginServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginServerWithDefaults

`func NewOriginServerWithDefaults() *OriginServer`

NewOriginServerWithDefaults instantiates a new OriginServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OriginServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OriginServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OriginServer) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *OriginServer) GetOptions() DefaultOriginServerOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *OriginServer) GetOptionsOk() (*DefaultOriginServerOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *OriginServer) SetOptions(v DefaultOriginServerOptions)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


