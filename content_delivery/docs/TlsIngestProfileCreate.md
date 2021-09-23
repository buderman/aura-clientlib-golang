# TlsIngestProfileCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientTlsCertificate** | Pointer to **string** | An OpenSSL PEM-formatted certificate associated with the trusted CA that signed the origin&#39;s certificate. | [optional] 
**Name** | **string** |  | 
**CaCertificate** | Pointer to **string** | An OpenSSL PEM-formatted certificate associated with the trusted CA that signed the origin&#39;s certificate. | [optional] 
**CaCertificateDetails** | Pointer to [**X509CertificateSummaryType**](X509CertificateSummaryType.md) |  | [optional] 
**ContentProviderId** | **int32** |  | 
**Ciphers** | Pointer to **string** | An OpenSSL style [ciphers string](https://www.openssl.org/docs/apps/ciphers.html#CIPHER-STRINGS). | [optional] 
**TlsProtocols** | Pointer to **[]string** | TLS protocols supported, either &#x60;TLSV1&#x60;, &#x60;TLSV1_1&#x60;, or &#x60;TLSV1_2&#x60;. Defaults to &#x60;TLSV1_2&#x60;.  The use of &#x60;SSLV2&#x60; or &#x60;SSLV3&#x60; are discouraged for security reasons. | [optional] [default to ["TLSV1_2"]]
**CipherSuites** | Pointer to **string** | An OpenSSL style [ciphers string](https://www.openssl.org/docs/apps/ciphers.html#CIPHER-STRINGS). | [optional] 
**ClientTlsCertificateKey** | Pointer to **string** | An OpenSSL PEM-formatted certificate associated with the trusted CA that signed the origin&#39;s certificate. | [optional] 
**ClientTlsCertificateDetails** | Pointer to [**X509CertificateSummaryType**](X509CertificateSummaryType.md) |  | [optional] 
**EnableClientAuth** | Pointer to **bool** | Allows you to enable or disable client authorization. Enabling this feature causes HyperCache to authenticate client certificates using the CA certificate for each TLS connection. | [optional] [default to false]
**Crl** | Pointer to **string** | Certificate Revocation List. An OpenSSL PEM-formatted list of revoked client certificates. | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewTlsIngestProfileCreate

`func NewTlsIngestProfileCreate(name string, contentProviderId int32, ) *TlsIngestProfileCreate`

NewTlsIngestProfileCreate instantiates a new TlsIngestProfileCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsIngestProfileCreateWithDefaults

`func NewTlsIngestProfileCreateWithDefaults() *TlsIngestProfileCreate`

NewTlsIngestProfileCreateWithDefaults instantiates a new TlsIngestProfileCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientTlsCertificate

`func (o *TlsIngestProfileCreate) GetClientTlsCertificate() string`

GetClientTlsCertificate returns the ClientTlsCertificate field if non-nil, zero value otherwise.

### GetClientTlsCertificateOk

`func (o *TlsIngestProfileCreate) GetClientTlsCertificateOk() (*string, bool)`

GetClientTlsCertificateOk returns a tuple with the ClientTlsCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTlsCertificate

`func (o *TlsIngestProfileCreate) SetClientTlsCertificate(v string)`

SetClientTlsCertificate sets ClientTlsCertificate field to given value.

### HasClientTlsCertificate

`func (o *TlsIngestProfileCreate) HasClientTlsCertificate() bool`

HasClientTlsCertificate returns a boolean if a field has been set.

### GetName

`func (o *TlsIngestProfileCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TlsIngestProfileCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TlsIngestProfileCreate) SetName(v string)`

SetName sets Name field to given value.


### GetCaCertificate

`func (o *TlsIngestProfileCreate) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *TlsIngestProfileCreate) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *TlsIngestProfileCreate) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *TlsIngestProfileCreate) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### GetCaCertificateDetails

`func (o *TlsIngestProfileCreate) GetCaCertificateDetails() X509CertificateSummaryType`

GetCaCertificateDetails returns the CaCertificateDetails field if non-nil, zero value otherwise.

### GetCaCertificateDetailsOk

`func (o *TlsIngestProfileCreate) GetCaCertificateDetailsOk() (*X509CertificateSummaryType, bool)`

GetCaCertificateDetailsOk returns a tuple with the CaCertificateDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificateDetails

`func (o *TlsIngestProfileCreate) SetCaCertificateDetails(v X509CertificateSummaryType)`

SetCaCertificateDetails sets CaCertificateDetails field to given value.

### HasCaCertificateDetails

`func (o *TlsIngestProfileCreate) HasCaCertificateDetails() bool`

HasCaCertificateDetails returns a boolean if a field has been set.

### GetContentProviderId

`func (o *TlsIngestProfileCreate) GetContentProviderId() int32`

GetContentProviderId returns the ContentProviderId field if non-nil, zero value otherwise.

### GetContentProviderIdOk

`func (o *TlsIngestProfileCreate) GetContentProviderIdOk() (*int32, bool)`

GetContentProviderIdOk returns a tuple with the ContentProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentProviderId

`func (o *TlsIngestProfileCreate) SetContentProviderId(v int32)`

SetContentProviderId sets ContentProviderId field to given value.


### GetCiphers

`func (o *TlsIngestProfileCreate) GetCiphers() string`

GetCiphers returns the Ciphers field if non-nil, zero value otherwise.

### GetCiphersOk

`func (o *TlsIngestProfileCreate) GetCiphersOk() (*string, bool)`

GetCiphersOk returns a tuple with the Ciphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphers

`func (o *TlsIngestProfileCreate) SetCiphers(v string)`

SetCiphers sets Ciphers field to given value.

### HasCiphers

`func (o *TlsIngestProfileCreate) HasCiphers() bool`

HasCiphers returns a boolean if a field has been set.

### GetTlsProtocols

`func (o *TlsIngestProfileCreate) GetTlsProtocols() []string`

GetTlsProtocols returns the TlsProtocols field if non-nil, zero value otherwise.

### GetTlsProtocolsOk

`func (o *TlsIngestProfileCreate) GetTlsProtocolsOk() (*[]string, bool)`

GetTlsProtocolsOk returns a tuple with the TlsProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsProtocols

`func (o *TlsIngestProfileCreate) SetTlsProtocols(v []string)`

SetTlsProtocols sets TlsProtocols field to given value.

### HasTlsProtocols

`func (o *TlsIngestProfileCreate) HasTlsProtocols() bool`

HasTlsProtocols returns a boolean if a field has been set.

### GetCipherSuites

`func (o *TlsIngestProfileCreate) GetCipherSuites() string`

GetCipherSuites returns the CipherSuites field if non-nil, zero value otherwise.

### GetCipherSuitesOk

`func (o *TlsIngestProfileCreate) GetCipherSuitesOk() (*string, bool)`

GetCipherSuitesOk returns a tuple with the CipherSuites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherSuites

`func (o *TlsIngestProfileCreate) SetCipherSuites(v string)`

SetCipherSuites sets CipherSuites field to given value.

### HasCipherSuites

`func (o *TlsIngestProfileCreate) HasCipherSuites() bool`

HasCipherSuites returns a boolean if a field has been set.

### GetClientTlsCertificateKey

`func (o *TlsIngestProfileCreate) GetClientTlsCertificateKey() string`

GetClientTlsCertificateKey returns the ClientTlsCertificateKey field if non-nil, zero value otherwise.

### GetClientTlsCertificateKeyOk

`func (o *TlsIngestProfileCreate) GetClientTlsCertificateKeyOk() (*string, bool)`

GetClientTlsCertificateKeyOk returns a tuple with the ClientTlsCertificateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTlsCertificateKey

`func (o *TlsIngestProfileCreate) SetClientTlsCertificateKey(v string)`

SetClientTlsCertificateKey sets ClientTlsCertificateKey field to given value.

### HasClientTlsCertificateKey

`func (o *TlsIngestProfileCreate) HasClientTlsCertificateKey() bool`

HasClientTlsCertificateKey returns a boolean if a field has been set.

### GetClientTlsCertificateDetails

`func (o *TlsIngestProfileCreate) GetClientTlsCertificateDetails() X509CertificateSummaryType`

GetClientTlsCertificateDetails returns the ClientTlsCertificateDetails field if non-nil, zero value otherwise.

### GetClientTlsCertificateDetailsOk

`func (o *TlsIngestProfileCreate) GetClientTlsCertificateDetailsOk() (*X509CertificateSummaryType, bool)`

GetClientTlsCertificateDetailsOk returns a tuple with the ClientTlsCertificateDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTlsCertificateDetails

`func (o *TlsIngestProfileCreate) SetClientTlsCertificateDetails(v X509CertificateSummaryType)`

SetClientTlsCertificateDetails sets ClientTlsCertificateDetails field to given value.

### HasClientTlsCertificateDetails

`func (o *TlsIngestProfileCreate) HasClientTlsCertificateDetails() bool`

HasClientTlsCertificateDetails returns a boolean if a field has been set.

### GetEnableClientAuth

`func (o *TlsIngestProfileCreate) GetEnableClientAuth() bool`

GetEnableClientAuth returns the EnableClientAuth field if non-nil, zero value otherwise.

### GetEnableClientAuthOk

`func (o *TlsIngestProfileCreate) GetEnableClientAuthOk() (*bool, bool)`

GetEnableClientAuthOk returns a tuple with the EnableClientAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientAuth

`func (o *TlsIngestProfileCreate) SetEnableClientAuth(v bool)`

SetEnableClientAuth sets EnableClientAuth field to given value.

### HasEnableClientAuth

`func (o *TlsIngestProfileCreate) HasEnableClientAuth() bool`

HasEnableClientAuth returns a boolean if a field has been set.

### GetCrl

`func (o *TlsIngestProfileCreate) GetCrl() string`

GetCrl returns the Crl field if non-nil, zero value otherwise.

### GetCrlOk

`func (o *TlsIngestProfileCreate) GetCrlOk() (*string, bool)`

GetCrlOk returns a tuple with the Crl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrl

`func (o *TlsIngestProfileCreate) SetCrl(v string)`

SetCrl sets Crl field to given value.

### HasCrl

`func (o *TlsIngestProfileCreate) HasCrl() bool`

HasCrl returns a boolean if a field has been set.

### GetDescription

`func (o *TlsIngestProfileCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TlsIngestProfileCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TlsIngestProfileCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TlsIngestProfileCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


