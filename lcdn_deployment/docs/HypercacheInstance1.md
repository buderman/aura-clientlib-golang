# HypercacheInstance1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientServingBandwidthLimitMbps** | Pointer to **int32** | A value that determines the maximum client-serving interface bandwidth on a node, in megabits per second. Request Router compares the bandwidth usage for each node with the maximum configured limit to help determine node availability. | [optional] 
**CachingMemoryOverrideMB** | Pointer to **int32** | A value that you assign to the HyperCache node to override the amount of caching memory the HyperCache node uses. Specify the value in megabytes. Recommended caching memory is 1891 megabytes out of 3783 megabytes of total memory. Improperly configuring the HyperCache caching memory can negatively affect the CDN. You should not have to override the automatically calculated memory. However, if you choose to override the caching memory, we strongly recommend that you first contact [Akamai Technical Support](https://control.akamai.com/apps/support-ui/#/contact-support). | [optional] 
**NodeId** | **int32** | The unique identifier for a HyperCache node. | 
**License** | **string** | The license for a particular HyperCache instance. The minimum length is 1 and the maximum length is 32767. | 
**ServiceLabels** | Pointer to [**ServiceLabelsType1**](ServiceLabelsType1.md) |  | [optional] 

## Methods

### NewHypercacheInstance1

`func NewHypercacheInstance1(nodeId int32, license string, ) *HypercacheInstance1`

NewHypercacheInstance1 instantiates a new HypercacheInstance1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypercacheInstance1WithDefaults

`func NewHypercacheInstance1WithDefaults() *HypercacheInstance1`

NewHypercacheInstance1WithDefaults instantiates a new HypercacheInstance1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientServingBandwidthLimitMbps

`func (o *HypercacheInstance1) GetClientServingBandwidthLimitMbps() int32`

GetClientServingBandwidthLimitMbps returns the ClientServingBandwidthLimitMbps field if non-nil, zero value otherwise.

### GetClientServingBandwidthLimitMbpsOk

`func (o *HypercacheInstance1) GetClientServingBandwidthLimitMbpsOk() (*int32, bool)`

GetClientServingBandwidthLimitMbpsOk returns a tuple with the ClientServingBandwidthLimitMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientServingBandwidthLimitMbps

`func (o *HypercacheInstance1) SetClientServingBandwidthLimitMbps(v int32)`

SetClientServingBandwidthLimitMbps sets ClientServingBandwidthLimitMbps field to given value.

### HasClientServingBandwidthLimitMbps

`func (o *HypercacheInstance1) HasClientServingBandwidthLimitMbps() bool`

HasClientServingBandwidthLimitMbps returns a boolean if a field has been set.

### GetCachingMemoryOverrideMB

`func (o *HypercacheInstance1) GetCachingMemoryOverrideMB() int32`

GetCachingMemoryOverrideMB returns the CachingMemoryOverrideMB field if non-nil, zero value otherwise.

### GetCachingMemoryOverrideMBOk

`func (o *HypercacheInstance1) GetCachingMemoryOverrideMBOk() (*int32, bool)`

GetCachingMemoryOverrideMBOk returns a tuple with the CachingMemoryOverrideMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachingMemoryOverrideMB

`func (o *HypercacheInstance1) SetCachingMemoryOverrideMB(v int32)`

SetCachingMemoryOverrideMB sets CachingMemoryOverrideMB field to given value.

### HasCachingMemoryOverrideMB

`func (o *HypercacheInstance1) HasCachingMemoryOverrideMB() bool`

HasCachingMemoryOverrideMB returns a boolean if a field has been set.

### GetNodeId

`func (o *HypercacheInstance1) GetNodeId() int32`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *HypercacheInstance1) GetNodeIdOk() (*int32, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *HypercacheInstance1) SetNodeId(v int32)`

SetNodeId sets NodeId field to given value.


### GetLicense

`func (o *HypercacheInstance1) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *HypercacheInstance1) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *HypercacheInstance1) SetLicense(v string)`

SetLicense sets License field to given value.


### GetServiceLabels

`func (o *HypercacheInstance1) GetServiceLabels() ServiceLabelsType1`

GetServiceLabels returns the ServiceLabels field if non-nil, zero value otherwise.

### GetServiceLabelsOk

`func (o *HypercacheInstance1) GetServiceLabelsOk() (*ServiceLabelsType1, bool)`

GetServiceLabelsOk returns a tuple with the ServiceLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLabels

`func (o *HypercacheInstance1) SetServiceLabels(v ServiceLabelsType1)`

SetServiceLabels sets ServiceLabels field to given value.

### HasServiceLabels

`func (o *HypercacheInstance1) HasServiceLabels() bool`

HasServiceLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


