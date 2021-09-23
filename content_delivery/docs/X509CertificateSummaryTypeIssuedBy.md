# X509CertificateSummaryTypeIssuedBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | **string** | The organization of the certificate authorizer. | 
**CommonName** | **string** | The entity (user/web server/domain/hostname) associated with the certificate authorizer. | 
**OrganizationalUnit** | **string** | The organizational unit of the certificate issuer. | 

## Methods

### NewX509CertificateSummaryTypeIssuedBy

`func NewX509CertificateSummaryTypeIssuedBy(organization string, commonName string, organizationalUnit string, ) *X509CertificateSummaryTypeIssuedBy`

NewX509CertificateSummaryTypeIssuedBy instantiates a new X509CertificateSummaryTypeIssuedBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewX509CertificateSummaryTypeIssuedByWithDefaults

`func NewX509CertificateSummaryTypeIssuedByWithDefaults() *X509CertificateSummaryTypeIssuedBy`

NewX509CertificateSummaryTypeIssuedByWithDefaults instantiates a new X509CertificateSummaryTypeIssuedBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *X509CertificateSummaryTypeIssuedBy) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *X509CertificateSummaryTypeIssuedBy) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *X509CertificateSummaryTypeIssuedBy) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetCommonName

`func (o *X509CertificateSummaryTypeIssuedBy) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *X509CertificateSummaryTypeIssuedBy) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *X509CertificateSummaryTypeIssuedBy) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### GetOrganizationalUnit

`func (o *X509CertificateSummaryTypeIssuedBy) GetOrganizationalUnit() string`

GetOrganizationalUnit returns the OrganizationalUnit field if non-nil, zero value otherwise.

### GetOrganizationalUnitOk

`func (o *X509CertificateSummaryTypeIssuedBy) GetOrganizationalUnitOk() (*string, bool)`

GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnit

`func (o *X509CertificateSummaryTypeIssuedBy) SetOrganizationalUnit(v string)`

SetOrganizationalUnit sets OrganizationalUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


