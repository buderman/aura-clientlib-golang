# X509CertificateSummaryTypeIssuedTo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | **string** | The organization of the certificate owner. | 
**Serial** | **string** | A serial number that uniquely identifies the certificate. | 
**CommonName** | **string** | A common name (CN) that identifies the host and domain name associated with the certificate. | 
**OrganizationalUnit** | **string** | The organizational unit of the certificate owner. | 

## Methods

### NewX509CertificateSummaryTypeIssuedTo

`func NewX509CertificateSummaryTypeIssuedTo(organization string, serial string, commonName string, organizationalUnit string, ) *X509CertificateSummaryTypeIssuedTo`

NewX509CertificateSummaryTypeIssuedTo instantiates a new X509CertificateSummaryTypeIssuedTo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewX509CertificateSummaryTypeIssuedToWithDefaults

`func NewX509CertificateSummaryTypeIssuedToWithDefaults() *X509CertificateSummaryTypeIssuedTo`

NewX509CertificateSummaryTypeIssuedToWithDefaults instantiates a new X509CertificateSummaryTypeIssuedTo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *X509CertificateSummaryTypeIssuedTo) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *X509CertificateSummaryTypeIssuedTo) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *X509CertificateSummaryTypeIssuedTo) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetSerial

`func (o *X509CertificateSummaryTypeIssuedTo) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *X509CertificateSummaryTypeIssuedTo) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *X509CertificateSummaryTypeIssuedTo) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetCommonName

`func (o *X509CertificateSummaryTypeIssuedTo) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *X509CertificateSummaryTypeIssuedTo) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *X509CertificateSummaryTypeIssuedTo) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### GetOrganizationalUnit

`func (o *X509CertificateSummaryTypeIssuedTo) GetOrganizationalUnit() string`

GetOrganizationalUnit returns the OrganizationalUnit field if non-nil, zero value otherwise.

### GetOrganizationalUnitOk

`func (o *X509CertificateSummaryTypeIssuedTo) GetOrganizationalUnitOk() (*string, bool)`

GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnit

`func (o *X509CertificateSummaryTypeIssuedTo) SetOrganizationalUnit(v string)`

SetOrganizationalUnit sets OrganizationalUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


