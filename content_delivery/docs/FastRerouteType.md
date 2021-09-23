# FastRerouteType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginTimeout** | Pointer to **float32** | Value used by the HyperCache when it makes a request directly to the origin server. It specifies the number of seconds you want the HyperCache to wait before sending the second request to an alternate origin IP address. The number of seconds configured can be expressed to the nearest thousandth of a second, for example: 4, 3.0, 1.05, 2.005. Note that the configured value for this field must be at least 0.5 seconds less than the &#x60;intraCdnTimeout&#x60;. | [optional] [default to 1.0]
**Enable** | Pointer to **bool** | Enables Fast Reroute functionality. | [optional] [default to false]
**IntraCdnTimeout** | Pointer to **float32** | Used by the HyperCache when it makes a request to an origin through a parent site or a peer node within the same site. It specifies the number of seconds you want the HyperCache in one site to wait before sending the second request to an alternate site. The number of seconds configured can be expressed to the nearest thousandth of a second, for example: 4, 3.0, 1.05, 2.005. Note that the configured value for this field must be at least 0.5 seconds greater than the &#x60;originTimeout&#x60;. | [optional] [default to 1.5]

## Methods

### NewFastRerouteType

`func NewFastRerouteType() *FastRerouteType`

NewFastRerouteType instantiates a new FastRerouteType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFastRerouteTypeWithDefaults

`func NewFastRerouteTypeWithDefaults() *FastRerouteType`

NewFastRerouteTypeWithDefaults instantiates a new FastRerouteType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginTimeout

`func (o *FastRerouteType) GetOriginTimeout() float32`

GetOriginTimeout returns the OriginTimeout field if non-nil, zero value otherwise.

### GetOriginTimeoutOk

`func (o *FastRerouteType) GetOriginTimeoutOk() (*float32, bool)`

GetOriginTimeoutOk returns a tuple with the OriginTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTimeout

`func (o *FastRerouteType) SetOriginTimeout(v float32)`

SetOriginTimeout sets OriginTimeout field to given value.

### HasOriginTimeout

`func (o *FastRerouteType) HasOriginTimeout() bool`

HasOriginTimeout returns a boolean if a field has been set.

### GetEnable

`func (o *FastRerouteType) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *FastRerouteType) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *FastRerouteType) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *FastRerouteType) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetIntraCdnTimeout

`func (o *FastRerouteType) GetIntraCdnTimeout() float32`

GetIntraCdnTimeout returns the IntraCdnTimeout field if non-nil, zero value otherwise.

### GetIntraCdnTimeoutOk

`func (o *FastRerouteType) GetIntraCdnTimeoutOk() (*float32, bool)`

GetIntraCdnTimeoutOk returns a tuple with the IntraCdnTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntraCdnTimeout

`func (o *FastRerouteType) SetIntraCdnTimeout(v float32)`

SetIntraCdnTimeout sets IntraCdnTimeout field to given value.

### HasIntraCdnTimeout

`func (o *FastRerouteType) HasIntraCdnTimeout() bool`

HasIntraCdnTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


