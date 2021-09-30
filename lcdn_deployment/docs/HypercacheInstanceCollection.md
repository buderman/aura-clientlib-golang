# HypercacheInstanceCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | [**[]HypercacheInstance**](HypercacheInstance.md) |  | 
**Page** | [**PageType**](PageType.md) |  | 

## Methods

### NewHypercacheInstanceCollection

`func NewHypercacheInstanceCollection(instances []HypercacheInstance, page PageType, ) *HypercacheInstanceCollection`

NewHypercacheInstanceCollection instantiates a new HypercacheInstanceCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypercacheInstanceCollectionWithDefaults

`func NewHypercacheInstanceCollectionWithDefaults() *HypercacheInstanceCollection`

NewHypercacheInstanceCollectionWithDefaults instantiates a new HypercacheInstanceCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *HypercacheInstanceCollection) GetInstances() []HypercacheInstance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *HypercacheInstanceCollection) GetInstancesOk() (*[]HypercacheInstance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *HypercacheInstanceCollection) SetInstances(v []HypercacheInstance)`

SetInstances sets Instances field to given value.


### GetPage

`func (o *HypercacheInstanceCollection) GetPage() PageType`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *HypercacheInstanceCollection) GetPageOk() (*PageType, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *HypercacheInstanceCollection) SetPage(v PageType)`

SetPage sets Page field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


