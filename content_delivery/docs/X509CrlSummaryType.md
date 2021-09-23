# X509CrlSummaryType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevokedCertificates** | **[]map[string]interface{}** |  | 
**NextUpdate** | **time.Time** | The date by which the next CRL will be issued. | 
**IssuedBy** | [**X509CrlSummaryTypeIssuedBy**](X509CrlSummaryTypeIssuedBy.md) |  | 
**ThisUpdate** | **time.Time** | The issue date of the current CRL. | 

## Methods

### NewX509CrlSummaryType

`func NewX509CrlSummaryType(revokedCertificates []map[string]interface{}, nextUpdate time.Time, issuedBy X509CrlSummaryTypeIssuedBy, thisUpdate time.Time, ) *X509CrlSummaryType`

NewX509CrlSummaryType instantiates a new X509CrlSummaryType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewX509CrlSummaryTypeWithDefaults

`func NewX509CrlSummaryTypeWithDefaults() *X509CrlSummaryType`

NewX509CrlSummaryTypeWithDefaults instantiates a new X509CrlSummaryType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevokedCertificates

`func (o *X509CrlSummaryType) GetRevokedCertificates() []map[string]interface{}`

GetRevokedCertificates returns the RevokedCertificates field if non-nil, zero value otherwise.

### GetRevokedCertificatesOk

`func (o *X509CrlSummaryType) GetRevokedCertificatesOk() (*[]map[string]interface{}, bool)`

GetRevokedCertificatesOk returns a tuple with the RevokedCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedCertificates

`func (o *X509CrlSummaryType) SetRevokedCertificates(v []map[string]interface{})`

SetRevokedCertificates sets RevokedCertificates field to given value.


### GetNextUpdate

`func (o *X509CrlSummaryType) GetNextUpdate() time.Time`

GetNextUpdate returns the NextUpdate field if non-nil, zero value otherwise.

### GetNextUpdateOk

`func (o *X509CrlSummaryType) GetNextUpdateOk() (*time.Time, bool)`

GetNextUpdateOk returns a tuple with the NextUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUpdate

`func (o *X509CrlSummaryType) SetNextUpdate(v time.Time)`

SetNextUpdate sets NextUpdate field to given value.


### GetIssuedBy

`func (o *X509CrlSummaryType) GetIssuedBy() X509CrlSummaryTypeIssuedBy`

GetIssuedBy returns the IssuedBy field if non-nil, zero value otherwise.

### GetIssuedByOk

`func (o *X509CrlSummaryType) GetIssuedByOk() (*X509CrlSummaryTypeIssuedBy, bool)`

GetIssuedByOk returns a tuple with the IssuedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedBy

`func (o *X509CrlSummaryType) SetIssuedBy(v X509CrlSummaryTypeIssuedBy)`

SetIssuedBy sets IssuedBy field to given value.


### GetThisUpdate

`func (o *X509CrlSummaryType) GetThisUpdate() time.Time`

GetThisUpdate returns the ThisUpdate field if non-nil, zero value otherwise.

### GetThisUpdateOk

`func (o *X509CrlSummaryType) GetThisUpdateOk() (*time.Time, bool)`

GetThisUpdateOk returns a tuple with the ThisUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThisUpdate

`func (o *X509CrlSummaryType) SetThisUpdate(v time.Time)`

SetThisUpdate sets ThisUpdate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


