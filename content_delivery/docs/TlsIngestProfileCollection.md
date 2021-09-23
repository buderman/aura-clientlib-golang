# TlsIngestProfileCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TlsDeliveryProfiles** | Pointer to [**[]TlsIngestProfile3**](TlsIngestProfile3.md) |  | [optional] 
**Page** | [**PageType**](PageType.md) |  | 

## Methods

### NewTlsIngestProfileCollection

`func NewTlsIngestProfileCollection(page PageType, ) *TlsIngestProfileCollection`

NewTlsIngestProfileCollection instantiates a new TlsIngestProfileCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsIngestProfileCollectionWithDefaults

`func NewTlsIngestProfileCollectionWithDefaults() *TlsIngestProfileCollection`

NewTlsIngestProfileCollectionWithDefaults instantiates a new TlsIngestProfileCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTlsDeliveryProfiles

`func (o *TlsIngestProfileCollection) GetTlsDeliveryProfiles() []TlsIngestProfile3`

GetTlsDeliveryProfiles returns the TlsDeliveryProfiles field if non-nil, zero value otherwise.

### GetTlsDeliveryProfilesOk

`func (o *TlsIngestProfileCollection) GetTlsDeliveryProfilesOk() (*[]TlsIngestProfile3, bool)`

GetTlsDeliveryProfilesOk returns a tuple with the TlsDeliveryProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsDeliveryProfiles

`func (o *TlsIngestProfileCollection) SetTlsDeliveryProfiles(v []TlsIngestProfile3)`

SetTlsDeliveryProfiles sets TlsDeliveryProfiles field to given value.

### HasTlsDeliveryProfiles

`func (o *TlsIngestProfileCollection) HasTlsDeliveryProfiles() bool`

HasTlsDeliveryProfiles returns a boolean if a field has been set.

### GetPage

`func (o *TlsIngestProfileCollection) GetPage() PageType`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *TlsIngestProfileCollection) GetPageOk() (*PageType, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *TlsIngestProfileCollection) SetPage(v PageType)`

SetPage sets Page field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


