# TlsIngestProfileReadDetailed

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

### NewTlsIngestProfileReadDetailed

`func NewTlsIngestProfileReadDetailed(tlsIngestProfileId int32, contentProviderId int32, name string, ) *TlsIngestProfileReadDetailed`

NewTlsIngestProfileReadDetailed instantiates a new TlsIngestProfileReadDetailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsIngestProfileReadDetailedWithDefaults

`func NewTlsIngestProfileReadDetailedWithDefaults() *TlsIngestProfileReadDetailed`

NewTlsIngestProfileReadDetailedWithDefaults instantiates a new TlsIngestProfileReadDetailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientTlsCertificate

`func (o *TlsIngestProfileReadDetailed) GetClientTlsCertificate() string`

GetClientTlsCertificate returns the ClientTlsCertificate field if non-nil, zero value otherwise.

### GetClientTlsCertificateOk

`func (o *TlsIngestProfileReadDetailed) GetClientTlsCertificateOk() (*string, bool)`

GetClientTlsCertificateOk returns a tuple with the ClientTlsCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTlsCertificate

`func (o *TlsIngestProfileReadDetailed) SetClientTlsCertificate(v string)`

SetClientTlsCertificate sets ClientTlsCertificate field to given value.

### HasClientTlsCertificate

`func (o *TlsIngestProfileReadDetailed) HasClientTlsCertificate() bool`

HasClientTlsCertificate returns a boolean if a field has been set.

### GetDescription

`func (o *TlsIngestProfileReadDetailed) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TlsIngestProfileReadDetailed) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TlsIngestProfileReadDetailed) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TlsIngestProfileReadDetailed) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTlsIngestProfileId

`func (o *TlsIngestProfileReadDetailed) GetTlsIngestProfileId() int32`

GetTlsIngestProfileId returns the TlsIngestProfileId field if non-nil, zero value otherwise.

### GetTlsIngestProfileIdOk

`func (o *TlsIngestProfileReadDetailed) GetTlsIngestProfileIdOk() (*int32, bool)`

GetTlsIngestProfileIdOk returns a tuple with the TlsIngestProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsIngestProfileId

`func (o *TlsIngestProfileReadDetailed) SetTlsIngestProfileId(v int32)`

SetTlsIngestProfileId sets TlsIngestProfileId field to given value.


### GetCaCertificate

`func (o *TlsIngestProfileReadDetailed) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *TlsIngestProfileReadDetailed) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *TlsIngestProfileReadDetailed) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *TlsIngestProfileReadDetailed) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### GetCaCertificateDetails

`func (o *TlsIngestProfileReadDetailed) GetCaCertificateDetails() X509CertificateSummaryType`

GetCaCertificateDetails returns the CaCertificateDetails field if non-nil, zero value otherwise.

### GetCaCertificateDetailsOk

`func (o *TlsIngestProfileReadDetailed) GetCaCertificateDetailsOk() (*X509CertificateSummaryType, bool)`

GetCaCertificateDetailsOk returns a tuple with the CaCertificateDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificateDetails

`func (o *TlsIngestProfileReadDetailed) SetCaCertificateDetails(v X509CertificateSummaryType)`

SetCaCertificateDetails sets CaCertificateDetails field to given value.

### HasCaCertificateDetails

`func (o *TlsIngestProfileReadDetailed) HasCaCertificateDetails() bool`

HasCaCertificateDetails returns a boolean if a field has been set.

### GetContentProviderId

`func (o *TlsIngestProfileReadDetailed) GetContentProviderId() int32`

GetContentProviderId returns the ContentProviderId field if non-nil, zero value otherwise.

### GetContentProviderIdOk

`func (o *TlsIngestProfileReadDetailed) GetContentProviderIdOk() (*int32, bool)`

GetContentProviderIdOk returns a tuple with the ContentProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentProviderId

`func (o *TlsIngestProfileReadDetailed) SetContentProviderId(v int32)`

SetContentProviderId sets ContentProviderId field to given value.


### GetCiphers

`func (o *TlsIngestProfileReadDetailed) GetCiphers() string`

GetCiphers returns the Ciphers field if non-nil, zero value otherwise.

### GetCiphersOk

`func (o *TlsIngestProfileReadDetailed) GetCiphersOk() (*string, bool)`

GetCiphersOk returns a tuple with the Ciphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphers

`func (o *TlsIngestProfileReadDetailed) SetCiphers(v string)`

SetCiphers sets Ciphers field to given value.

### HasCiphers

`func (o *TlsIngestProfileReadDetailed) HasCiphers() bool`

HasCiphers returns a boolean if a field has been set.

### GetCrlDetails

`func (o *TlsIngestProfileReadDetailed) GetCrlDetails() X509CrlSummaryType`

GetCrlDetails returns the CrlDetails field if non-nil, zero value otherwise.

### GetCrlDetailsOk

`func (o *TlsIngestProfileReadDetailed) GetCrlDetailsOk() (*X509CrlSummaryType, bool)`

GetCrlDetailsOk returns a tuple with the CrlDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlDetails

`func (o *TlsIngestProfileReadDetailed) SetCrlDetails(v X509CrlSummaryType)`

SetCrlDetails sets CrlDetails field to given value.

### HasCrlDetails

`func (o *TlsIngestProfileReadDetailed) HasCrlDetails() bool`

HasCrlDetails returns a boolean if a field has been set.

### GetTlsProtocols

`func (o *TlsIngestProfileReadDetailed) GetTlsProtocols() []string`

GetTlsProtocols returns the TlsProtocols field if non-nil, zero value otherwise.

### GetTlsProtocolsOk

`func (o *TlsIngestProfileReadDetailed) GetTlsProtocolsOk() (*[]string, bool)`

GetTlsProtocolsOk returns a tuple with the TlsProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsProtocols

`func (o *TlsIngestProfileReadDetailed) SetTlsProtocols(v []string)`

SetTlsProtocols sets TlsProtocols field to given value.

### HasTlsProtocols

`func (o *TlsIngestProfileReadDetailed) HasTlsProtocols() bool`

HasTlsProtocols returns a boolean if a field has been set.

### GetCipherSuites

`func (o *TlsIngestProfileReadDetailed) GetCipherSuites() string`

GetCipherSuites returns the CipherSuites field if non-nil, zero value otherwise.

### GetCipherSuitesOk

`func (o *TlsIngestProfileReadDetailed) GetCipherSuitesOk() (*string, bool)`

GetCipherSuitesOk returns a tuple with the CipherSuites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherSuites

`func (o *TlsIngestProfileReadDetailed) SetCipherSuites(v string)`

SetCipherSuites sets CipherSuites field to given value.

### HasCipherSuites

`func (o *TlsIngestProfileReadDetailed) HasCipherSuites() bool`

HasCipherSuites returns a boolean if a field has been set.

### GetClientTlsCertificateKey

`func (o *TlsIngestProfileReadDetailed) GetClientTlsCertificateKey() string`

GetClientTlsCertificateKey returns the ClientTlsCertificateKey field if non-nil, zero value otherwise.

### GetClientTlsCertificateKeyOk

`func (o *TlsIngestProfileReadDetailed) GetClientTlsCertificateKeyOk() (*string, bool)`

GetClientTlsCertificateKeyOk returns a tuple with the ClientTlsCertificateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTlsCertificateKey

`func (o *TlsIngestProfileReadDetailed) SetClientTlsCertificateKey(v string)`

SetClientTlsCertificateKey sets ClientTlsCertificateKey field to given value.

### HasClientTlsCertificateKey

`func (o *TlsIngestProfileReadDetailed) HasClientTlsCertificateKey() bool`

HasClientTlsCertificateKey returns a boolean if a field has been set.

### GetClientTlsCertificateDetails

`func (o *TlsIngestProfileReadDetailed) GetClientTlsCertificateDetails() X509CertificateSummaryType`

GetClientTlsCertificateDetails returns the ClientTlsCertificateDetails field if non-nil, zero value otherwise.

### GetClientTlsCertificateDetailsOk

`func (o *TlsIngestProfileReadDetailed) GetClientTlsCertificateDetailsOk() (*X509CertificateSummaryType, bool)`

GetClientTlsCertificateDetailsOk returns a tuple with the ClientTlsCertificateDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTlsCertificateDetails

`func (o *TlsIngestProfileReadDetailed) SetClientTlsCertificateDetails(v X509CertificateSummaryType)`

SetClientTlsCertificateDetails sets ClientTlsCertificateDetails field to given value.

### HasClientTlsCertificateDetails

`func (o *TlsIngestProfileReadDetailed) HasClientTlsCertificateDetails() bool`

HasClientTlsCertificateDetails returns a boolean if a field has been set.

### GetEnableClientAuth

`func (o *TlsIngestProfileReadDetailed) GetEnableClientAuth() bool`

GetEnableClientAuth returns the EnableClientAuth field if non-nil, zero value otherwise.

### GetEnableClientAuthOk

`func (o *TlsIngestProfileReadDetailed) GetEnableClientAuthOk() (*bool, bool)`

GetEnableClientAuthOk returns a tuple with the EnableClientAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientAuth

`func (o *TlsIngestProfileReadDetailed) SetEnableClientAuth(v bool)`

SetEnableClientAuth sets EnableClientAuth field to given value.

### HasEnableClientAuth

`func (o *TlsIngestProfileReadDetailed) HasEnableClientAuth() bool`

HasEnableClientAuth returns a boolean if a field has been set.

### GetCrl

`func (o *TlsIngestProfileReadDetailed) GetCrl() string`

GetCrl returns the Crl field if non-nil, zero value otherwise.

### GetCrlOk

`func (o *TlsIngestProfileReadDetailed) GetCrlOk() (*string, bool)`

GetCrlOk returns a tuple with the Crl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrl

`func (o *TlsIngestProfileReadDetailed) SetCrl(v string)`

SetCrl sets Crl field to given value.

### HasCrl

`func (o *TlsIngestProfileReadDetailed) HasCrl() bool`

HasCrl returns a boolean if a field has been set.

### GetName

`func (o *TlsIngestProfileReadDetailed) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TlsIngestProfileReadDetailed) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TlsIngestProfileReadDetailed) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


