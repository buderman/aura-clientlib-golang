# X509CertificateSummaryTypeValidity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotAfter** | Pointer to **time.Time** | The time and date on which the certificate is no longer valid. | [optional] 
**NotBefore** | Pointer to **time.Time** | The time and date on which the certificate is valid. | [optional] 

## Methods

### NewX509CertificateSummaryTypeValidity

`func NewX509CertificateSummaryTypeValidity() *X509CertificateSummaryTypeValidity`

NewX509CertificateSummaryTypeValidity instantiates a new X509CertificateSummaryTypeValidity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewX509CertificateSummaryTypeValidityWithDefaults

`func NewX509CertificateSummaryTypeValidityWithDefaults() *X509CertificateSummaryTypeValidity`

NewX509CertificateSummaryTypeValidityWithDefaults instantiates a new X509CertificateSummaryTypeValidity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotAfter

`func (o *X509CertificateSummaryTypeValidity) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *X509CertificateSummaryTypeValidity) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *X509CertificateSummaryTypeValidity) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *X509CertificateSummaryTypeValidity) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetNotBefore

`func (o *X509CertificateSummaryTypeValidity) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *X509CertificateSummaryTypeValidity) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *X509CertificateSummaryTypeValidity) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *X509CertificateSummaryTypeValidity) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


