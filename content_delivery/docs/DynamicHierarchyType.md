# DynamicHierarchyType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HealthCheckValidResponseCodes** | Pointer to **[]string** | A list of valid HTTP response codes that are to be considered a valid HTTP response. Each item is either a single valid response code or a range of valid response codes. | [optional] 
**HealthCheckInterval** | Pointer to **int32** | A numeric value that represents the number of seconds between successive HTTP GET health check requests. The minimum acceptable value is 1, and the default is 2. | [optional] 
**HealthCheckUrl** | Pointer to **string** | URL the health check request uses to determine the health of the origin server. The HyperCache periodically sends an HTTP GET request for this URL to check the health of the origin server. Any response received is interpreted as a healthy origin, including error responses.  This URI may either be a complete URL using the &#39;http&#39; scheme, or just a path with a leading &#39;/&#39;. | [optional] 

## Methods

### NewDynamicHierarchyType

`func NewDynamicHierarchyType() *DynamicHierarchyType`

NewDynamicHierarchyType instantiates a new DynamicHierarchyType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicHierarchyTypeWithDefaults

`func NewDynamicHierarchyTypeWithDefaults() *DynamicHierarchyType`

NewDynamicHierarchyTypeWithDefaults instantiates a new DynamicHierarchyType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealthCheckValidResponseCodes

`func (o *DynamicHierarchyType) GetHealthCheckValidResponseCodes() []string`

GetHealthCheckValidResponseCodes returns the HealthCheckValidResponseCodes field if non-nil, zero value otherwise.

### GetHealthCheckValidResponseCodesOk

`func (o *DynamicHierarchyType) GetHealthCheckValidResponseCodesOk() (*[]string, bool)`

GetHealthCheckValidResponseCodesOk returns a tuple with the HealthCheckValidResponseCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckValidResponseCodes

`func (o *DynamicHierarchyType) SetHealthCheckValidResponseCodes(v []string)`

SetHealthCheckValidResponseCodes sets HealthCheckValidResponseCodes field to given value.

### HasHealthCheckValidResponseCodes

`func (o *DynamicHierarchyType) HasHealthCheckValidResponseCodes() bool`

HasHealthCheckValidResponseCodes returns a boolean if a field has been set.

### GetHealthCheckInterval

`func (o *DynamicHierarchyType) GetHealthCheckInterval() int32`

GetHealthCheckInterval returns the HealthCheckInterval field if non-nil, zero value otherwise.

### GetHealthCheckIntervalOk

`func (o *DynamicHierarchyType) GetHealthCheckIntervalOk() (*int32, bool)`

GetHealthCheckIntervalOk returns a tuple with the HealthCheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckInterval

`func (o *DynamicHierarchyType) SetHealthCheckInterval(v int32)`

SetHealthCheckInterval sets HealthCheckInterval field to given value.

### HasHealthCheckInterval

`func (o *DynamicHierarchyType) HasHealthCheckInterval() bool`

HasHealthCheckInterval returns a boolean if a field has been set.

### GetHealthCheckUrl

`func (o *DynamicHierarchyType) GetHealthCheckUrl() string`

GetHealthCheckUrl returns the HealthCheckUrl field if non-nil, zero value otherwise.

### GetHealthCheckUrlOk

`func (o *DynamicHierarchyType) GetHealthCheckUrlOk() (*string, bool)`

GetHealthCheckUrlOk returns a tuple with the HealthCheckUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckUrl

`func (o *DynamicHierarchyType) SetHealthCheckUrl(v string)`

SetHealthCheckUrl sets HealthCheckUrl field to given value.

### HasHealthCheckUrl

`func (o *DynamicHierarchyType) HasHealthCheckUrl() bool`

HasHealthCheckUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


