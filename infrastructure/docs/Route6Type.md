# Route6Type

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | **string** | An IPv6 address. | 
**Metric** | Pointer to **int32** | A value up to 9999 that represents the cost metric for this static route. The route with the lowest metric is chosen among multiple routes that most closely match the destination address of a packet being forwarded. | [optional] 
**SourceIpAddress** | Pointer to **string** | The IPv6 source address for this IPv6 route. | [optional] 
**NextHop** | **string** | The address of the next gateway to which packets should be forwarded along the path to their final destination. | 

## Methods

### NewRoute6Type

`func NewRoute6Type(destination string, nextHop string, ) *Route6Type`

NewRoute6Type instantiates a new Route6Type object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoute6TypeWithDefaults

`func NewRoute6TypeWithDefaults() *Route6Type`

NewRoute6TypeWithDefaults instantiates a new Route6Type object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *Route6Type) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Route6Type) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Route6Type) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetMetric

`func (o *Route6Type) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *Route6Type) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *Route6Type) SetMetric(v int32)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *Route6Type) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetSourceIpAddress

`func (o *Route6Type) GetSourceIpAddress() string`

GetSourceIpAddress returns the SourceIpAddress field if non-nil, zero value otherwise.

### GetSourceIpAddressOk

`func (o *Route6Type) GetSourceIpAddressOk() (*string, bool)`

GetSourceIpAddressOk returns a tuple with the SourceIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIpAddress

`func (o *Route6Type) SetSourceIpAddress(v string)`

SetSourceIpAddress sets SourceIpAddress field to given value.

### HasSourceIpAddress

`func (o *Route6Type) HasSourceIpAddress() bool`

HasSourceIpAddress returns a boolean if a field has been set.

### GetNextHop

`func (o *Route6Type) GetNextHop() string`

GetNextHop returns the NextHop field if non-nil, zero value otherwise.

### GetNextHopOk

`func (o *Route6Type) GetNextHopOk() (*string, bool)`

GetNextHopOk returns a tuple with the NextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHop

`func (o *Route6Type) SetNextHop(v string)`

SetNextHop sets NextHop field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


