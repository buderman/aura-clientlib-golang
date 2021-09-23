# TlsIngestProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientTlsCertificate** | Pointer to **string** | An OpenSSL PEM-formatted certificate associated with the trusted CA that signed the origin&#39;s certificate. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TlsIngestProfileId** | **int32** |  | 
**CaCertificate** | Pointer to **string** | An OpenSSL PEM-formatted certificate associated with the trusted CA that signed the origin&#39;s certificate. | [optional] 
**CaCertificateDetails** | Pointer to [**X509CertificateSummaryType**](X509CertificateSummaryType.md) |  | [optional] 
**ContentProviderId** | **int32** |  | 
**Ciphers** | Pointer to **string** | An OpenSSL style [ciphers string](https://www.openssl.org/docs/apps/ciphers.html#CIPHER-STRINGS). | [optional] 
**CrlDetails** | Pointer to [**X509CrlSummaryType**](X509CrlSummaryType.md) |  | [optional] 
**TlsProtocols** | Pointer to **[]string** | TLS protocols supported, either &#x60;TLSV1&#x60;, &#x60;TLSV1_1&#x60;, or &#x60;TLSV1_2&#x60;. Defaults to &#x60;TLSV1_2&#x60;.  The use of &#x60;SSLV2&#x60; or &#x60;SSLV3&#x60; are discouraged for security reasons. | [optional] [default to ["TLSV1_2"]]
**CipherSuites** | Pointer to **string** | An OpenSSL style [ciphers string](https://www.openssl.org/docs/apps/ciphers.html#CIPHER-STRINGS). | [optional] 
**ClientTlsCertificateKey** | Pointer to **string** | An OpenSSL PEM-formatted certificate associated with the trusted CA that signed the origin&#39;s certificate. | [optional] 
**ClientTlsCertificateDetails** | Pointer to [**X509CertificateSummaryType**](X509CertificateSummaryType.md) |  | [optional] 
**EnableClientAuth** | Pointer to **bool** | Allows you to enable or disable client authorization. Enabling this feature causes HyperCache to authenticate client certificates using the CA certificate for each TLS connection. | [optional] [default to false]
**Crl** | Pointer to **string** | Certificate Revocation List. An OpenSSL PEM-formatted list of revoked client certificates. | [optional] 
**Name** | **string** |  | 

## Methods

### NewTlsIngestProfile

`func NewTlsIngestProfile(tlsIngestProfileId int32, contentProviderId int32, name string, ) *TlsIngestProfile`

NewTlsIngestProfile instantiates a new TlsIngestProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsIngestProfileWithDefaults

`func NewTlsIngestProfileWithDefaults() *TlsIngestProfile`

NewTlsIngestProfileWithDefaults instantiates a new TlsIngestProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientTlsCertificate

`func (o *TlsIngestProfile) GetClientTlsCertificate() string`

GetClientTlsCertificate returns the ClientTlsCertificate field if non-nil, zero value otherwise.

### GetClientTlsCertificateOk

`func (o *TlsIngestProfile) GetClientTlsCertificateOk() (*string, bool)`

GetClientTlsCertificateOk returns a tuple with the ClientTlsCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTlsCertificate

`func (o *TlsIngestProfile) SetClientTlsCertificate(v string)`

SetClientTlsCertificate sets ClientTlsCertificate field to given value.

### HasClientTlsCertificate

`func (o *TlsIngestProfile) HasClientTlsCertificate() bool`

HasClientTlsCertificate returns a boolean if a field has been set.

### GetDescription

`func (o *TlsIngestProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TlsIngestProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TlsIngestProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TlsIngestProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTlsIngestProfileId

`func (o *TlsIngestProfile) GetTlsIngestProfileId() int32`

GetTlsIngestProfileId returns the TlsIngestProfileId field if non-nil, zero value otherwise.

### GetTlsIngestProfileIdOk

`func (o *TlsIngestProfile) GetTlsIngestProfileIdOk() (*int32, bool)`

GetTlsIngestProfileIdOk returns a tuple with the TlsIngestProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsIngestProfileId

`func (o *TlsIngestProfile) SetTlsIngestProfileId(v int32)`

SetTlsIngestProfileId sets TlsIngestProfileId field to given value.


### GetCaCertificate

`func (o *TlsIngestProfile) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *TlsIngestProfile) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *TlsIngestProfile) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *TlsIngestProfile) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### GetCaCertificateDetails

`func (o *TlsIngestProfile) GetCaCertificateDetails() X509CertificateSummaryType`

GetCaCertificateDetails returns the CaCertificateDetails field if non-nil, zero value otherwise.

### GetCaCertificateDetailsOk

`func (o *TlsIngestProfile) GetCaCertificateDetailsOk() (*X509CertificateSummaryType, bool)`

GetCaCertificateDetailsOk returns a tuple with the CaCertificateDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificateDetails

`func (o *TlsIngestProfile) SetCaCertificateDetails(v X509CertificateSummaryType)`

SetCaCertificateDetails sets CaCertificateDetails field to given value.

### HasCaCertificateDetails

`func (o *TlsIngestProfile) HasCaCertificateDetails() bool`

HasCaCertificateDetails returns a boolean if a field has been set.

### GetContentProviderId

`func (o *TlsIngestProfile) GetContentProviderId() int32`

GetContentProviderId returns the ContentProviderId field if non-nil, zero value otherwise.

### GetContentProviderIdOk

`func (o *TlsIngestProfile) GetContentProviderIdOk() (*int32, bool)`

GetContentProviderIdOk returns a tuple with the ContentProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentProviderId

`func (o *TlsIngestProfile) SetContentProviderId(v int32)`

SetContentProviderId sets ContentProviderId field to given value.


### GetCiphers

`func (o *TlsIngestProfile) GetCiphers() string`

GetCiphers returns the Ciphers field if non-nil, zero value otherwise.

### GetCiphersOk

`func (o *TlsIngestProfile) GetCiphersOk() (*string, bool)`

GetCiphersOk returns a tuple with the Ciphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphers

`func (o *TlsIngestProfile) SetCiphers(v string)`

SetCiphers sets Ciphers field to given value.

### HasCiphers

`func (o *TlsIngestProfile) HasCiphers() bool`

HasCiphers returns a boolean if a field has been set.

### GetCrlDetails

`func (o *TlsIngestProfile) GetCrlDetails() X509CrlSummaryType`

GetCrlDetails returns the CrlDetails field if non-nil, zero value otherwise.

### GetCrlDetailsOk

`func (o *TlsIngestProfile) GetCrlDetailsOk() (*X509CrlSummaryType, bool)`

GetCrlDetailsOk returns a tuple with the CrlDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlDetails

`func (o *TlsIngestProfile) SetCrlDetails(v X509CrlSummaryType)`

SetCrlDetails sets CrlDetails field to given value.

### HasCrlDetails

`func (o *TlsIngestProfile) HasCrlDetails() bool`

HasCrlDetails returns a boolean if a field has been set.

### GetTlsProtocols

`func (o *TlsIngestProfile) GetTlsProtocols() []string`

GetTlsProtocols returns the TlsProtocols field if non-nil, zero value otherwise.

### GetTlsProtocolsOk

`func (o *TlsIngestProfile) GetTlsProtocolsOk() (*[]string, bool)`

GetTlsProtocolsOk returns a tuple with the TlsProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsProtocols

`func (o *TlsIngestProfile) SetTlsProtocols(v []string)`

SetTlsProtocols sets TlsProtocols field to given value.

### HasTlsProtocols

`func (o *TlsIngestProfile) HasTlsProtocols() bool`

HasTlsProtocols returns a boolean if a field has been set.

### GetCipherSuites

`func (o *TlsIngestProfile) GetCipherSuites() string`

GetCipherSuites returns the CipherSuites field if non-nil, zero value otherwise.

### GetCipherSuitesOk

`func (o *TlsIngestProfile) GetCipherSuitesOk() (*string, bool)`

GetCipherSuitesOk returns a tuple with the CipherSuites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherSuites

`func (o *TlsIngestProfile) SetCipherSuites(v string)`

SetCipherSuites sets CipherSuites field to given value.

### HasCipherSuites

`func (o *TlsIngestProfile) HasCipherSuites() bool`

HasCipherSuites returns a boolean if a field has been set.

### GetClientTlsCertificateKey

`func (o *TlsIngestProfile) GetClientTlsCertificateKey() string`

GetClientTlsCertificateKey returns the ClientTlsCertificateKey field if non-nil, zero value otherwise.

### GetClientTlsCertificateKeyOk

`func (o *TlsIngestProfile) GetClientTlsCertificateKeyOk() (*string, bool)`

GetClientTlsCertificateKeyOk returns a tuple with the ClientTlsCertificateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTlsCertificateKey

`func (o *TlsIngestProfile) SetClientTlsCertificateKey(v string)`

SetClientTlsCertificateKey sets ClientTlsCertificateKey field to given value.

### HasClientTlsCertificateKey

`func (o *TlsIngestProfile) HasClientTlsCertificateKey() bool`

HasClientTlsCertificateKey returns a boolean if a field has been set.

### GetClientTlsCertificateDetails

`func (o *TlsIngestProfile) GetClientTlsCertificateDetails() X509CertificateSummaryType`

GetClientTlsCertificateDetails returns the ClientTlsCertificateDetails field if non-nil, zero value otherwise.

### GetClientTlsCertificateDetailsOk

`func (o *TlsIngestProfile) GetClientTlsCertificateDetailsOk() (*X509CertificateSummaryType, bool)`

GetClientTlsCertificateDetailsOk returns a tuple with the ClientTlsCertificateDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTlsCertificateDetails

`func (o *TlsIngestProfile) SetClientTlsCertificateDetails(v X509CertificateSummaryType)`

SetClientTlsCertificateDetails sets ClientTlsCertificateDetails field to given value.

### HasClientTlsCertificateDetails

`func (o *TlsIngestProfile) HasClientTlsCertificateDetails() bool`

HasClientTlsCertificateDetails returns a boolean if a field has been set.

### GetEnableClientAuth

`func (o *TlsIngestProfile) GetEnableClientAuth() bool`

GetEnableClientAuth returns the EnableClientAuth field if non-nil, zero value otherwise.

### GetEnableClientAuthOk

`func (o *TlsIngestProfile) GetEnableClientAuthOk() (*bool, bool)`

GetEnableClientAuthOk returns a tuple with the EnableClientAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientAuth

`func (o *TlsIngestProfile) SetEnableClientAuth(v bool)`

SetEnableClientAuth sets EnableClientAuth field to given value.

### HasEnableClientAuth

`func (o *TlsIngestProfile) HasEnableClientAuth() bool`

HasEnableClientAuth returns a boolean if a field has been set.

### GetCrl

`func (o *TlsIngestProfile) GetCrl() string`

GetCrl returns the Crl field if non-nil, zero value otherwise.

### GetCrlOk

`func (o *TlsIngestProfile) GetCrlOk() (*string, bool)`

GetCrlOk returns a tuple with the Crl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrl

`func (o *TlsIngestProfile) SetCrl(v string)`

SetCrl sets Crl field to given value.

### HasCrl

`func (o *TlsIngestProfile) HasCrl() bool`

HasCrl returns a boolean if a field has been set.

### GetName

`func (o *TlsIngestProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TlsIngestProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TlsIngestProfile) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


