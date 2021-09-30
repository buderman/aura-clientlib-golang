# ServiceLabelsType1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntraSite** | Pointer to **[]string** | List of IP addresses that you specify on a HyperCache node for intra-site traffic. Currently limited to a single address. | [optional] 
**InterSite** | Pointer to **[]string** | List of IP addresses that you configure on a HyperCache node for inter-site traffic. Regardless of the number of IP addresses you assigned, you can specify only one IP address for inter-site traffic. | [optional] 
**ClientServing** | Pointer to [**[]ClientServingLabelType**](ClientServingLabelType.md) | Allows you to group one or more HyperCache IP addresses by associating them with a client-serving label. | [optional] 

## Methods

### NewServiceLabelsType1

`func NewServiceLabelsType1() *ServiceLabelsType1`

NewServiceLabelsType1 instantiates a new ServiceLabelsType1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceLabelsType1WithDefaults

`func NewServiceLabelsType1WithDefaults() *ServiceLabelsType1`

NewServiceLabelsType1WithDefaults instantiates a new ServiceLabelsType1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntraSite

`func (o *ServiceLabelsType1) GetIntraSite() []string`

GetIntraSite returns the IntraSite field if non-nil, zero value otherwise.

### GetIntraSiteOk

`func (o *ServiceLabelsType1) GetIntraSiteOk() (*[]string, bool)`

GetIntraSiteOk returns a tuple with the IntraSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntraSite

`func (o *ServiceLabelsType1) SetIntraSite(v []string)`

SetIntraSite sets IntraSite field to given value.

### HasIntraSite

`func (o *ServiceLabelsType1) HasIntraSite() bool`

HasIntraSite returns a boolean if a field has been set.

### GetInterSite

`func (o *ServiceLabelsType1) GetInterSite() []string`

GetInterSite returns the InterSite field if non-nil, zero value otherwise.

### GetInterSiteOk

`func (o *ServiceLabelsType1) GetInterSiteOk() (*[]string, bool)`

GetInterSiteOk returns a tuple with the InterSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterSite

`func (o *ServiceLabelsType1) SetInterSite(v []string)`

SetInterSite sets InterSite field to given value.

### HasInterSite

`func (o *ServiceLabelsType1) HasInterSite() bool`

HasInterSite returns a boolean if a field has been set.

### GetClientServing

`func (o *ServiceLabelsType1) GetClientServing() []ClientServingLabelType`

GetClientServing returns the ClientServing field if non-nil, zero value otherwise.

### GetClientServingOk

`func (o *ServiceLabelsType1) GetClientServingOk() (*[]ClientServingLabelType, bool)`

GetClientServingOk returns a tuple with the ClientServing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientServing

`func (o *ServiceLabelsType1) SetClientServing(v []ClientServingLabelType)`

SetClientServing sets ClientServing field to given value.

### HasClientServing

`func (o *ServiceLabelsType1) HasClientServing() bool`

HasClientServing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


