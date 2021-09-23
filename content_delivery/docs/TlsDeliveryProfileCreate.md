# TlsDeliveryProfileCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TlsCertificate** | Pointer to **string** | An OpenSSL PEM-formatted certificate associated with the trusted CA that signed the origin&#39;s certificate. | [optional] 
**TlsCertificateAndKeySecretId** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Ciphers** | Pointer to **string** | An OpenSSL style [ciphers string](https://www.openssl.org/docs/apps/ciphers.html#CIPHER-STRINGS). | [optional] 
**CaCertificate** | Pointer to **string** | An OpenSSL PEM-formatted certificate associated with the trusted CA that signed the origin&#39;s certificate. | [optional] 
**CaCertificateDetails** | Pointer to [**X509CertificateSummaryType**](X509CertificateSummaryType.md) |  | [optional] 
**ContentProviderId** | **int32** |  | 
**CaCertificateSecretId** | Pointer to **int32** |  | [optional] 
**TlsVerificationDepth** | Pointer to **int32** | Sets the verification depth in the client certificates chain. | [optional] 
**TlsBufferSize** | Pointer to [**TlsBufferSizeType**](TlsBufferSizeType.md) |  | [optional] 
**TlsCertificateKey** | Pointer to **string** | An OpenSSL PEM-formatted certificate associated with the trusted CA that signed the origin&#39;s certificate. | [optional] 
**EcdhCurve** | Pointer to [**EcdhCurveType**](EcdhCurveType.md) |  | [optional] 
**TlsCertificateDetails** | Pointer to [**X509CertificateSummaryType**](X509CertificateSummaryType.md) |  | [optional] 
**TlsProtocols** | Pointer to **[]string** | TLS protocols supported, either &#x60;TLSV1&#x60;, &#x60;TLSV1_1&#x60;, &#x60;TLSV1_2&#x60;, or &#x60;TLSV1_3&#x60;.  Defaults to &#x60;TLSV1_2&#x60;, &#x60;TLSV1_3&#x60;.  The use of &#x60;SSLV2&#x60; or &#x60;SSLV3&#x60; are discouraged for security reasons. | [optional] [default to ["TLSV1_2","TLSV1_3"]]
**CipherSuites** | Pointer to **string** | An OpenSSL style [ciphers string](https://www.openssl.org/docs/apps/ciphers.html#CIPHER-STRINGS). | [optional] 
**EnableClientAuth** | Pointer to **bool** | Allows you to enable or disable client authorization. Enabling this feature causes HyperCache to authenticate client certificates using the CA certificate for each TLS connection. | [optional] [default to false]
**Crl** | Pointer to **string** | Certificate Revocation List. An OpenSSL PEM-formatted list of revoked client certificates. | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewTlsDeliveryProfileCreate

`func NewTlsDeliveryProfileCreate(name string, contentProviderId int32, ) *TlsDeliveryProfileCreate`

NewTlsDeliveryProfileCreate instantiates a new TlsDeliveryProfileCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsDeliveryProfileCreateWithDefaults

`func NewTlsDeliveryProfileCreateWithDefaults() *TlsDeliveryProfileCreate`

NewTlsDeliveryProfileCreateWithDefaults instantiates a new TlsDeliveryProfileCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTlsCertificate

`func (o *TlsDeliveryProfileCreate) GetTlsCertificate() string`

GetTlsCertificate returns the TlsCertificate field if non-nil, zero value otherwise.

### GetTlsCertificateOk

`func (o *TlsDeliveryProfileCreate) GetTlsCertificateOk() (*string, bool)`

GetTlsCertificateOk returns a tuple with the TlsCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificate

`func (o *TlsDeliveryProfileCreate) SetTlsCertificate(v string)`

SetTlsCertificate sets TlsCertificate field to given value.

### HasTlsCertificate

`func (o *TlsDeliveryProfileCreate) HasTlsCertificate() bool`

HasTlsCertificate returns a boolean if a field has been set.

### GetTlsCertificateAndKeySecretId

`func (o *TlsDeliveryProfileCreate) GetTlsCertificateAndKeySecretId() int32`

GetTlsCertificateAndKeySecretId returns the TlsCertificateAndKeySecretId field if non-nil, zero value otherwise.

### GetTlsCertificateAndKeySecretIdOk

`func (o *TlsDeliveryProfileCreate) GetTlsCertificateAndKeySecretIdOk() (*int32, bool)`

GetTlsCertificateAndKeySecretIdOk returns a tuple with the TlsCertificateAndKeySecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificateAndKeySecretId

`func (o *TlsDeliveryProfileCreate) SetTlsCertificateAndKeySecretId(v int32)`

SetTlsCertificateAndKeySecretId sets TlsCertificateAndKeySecretId field to given value.

### HasTlsCertificateAndKeySecretId

`func (o *TlsDeliveryProfileCreate) HasTlsCertificateAndKeySecretId() bool`

HasTlsCertificateAndKeySecretId returns a boolean if a field has been set.

### GetName

`func (o *TlsDeliveryProfileCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TlsDeliveryProfileCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TlsDeliveryProfileCreate) SetName(v string)`

SetName sets Name field to given value.


### GetCiphers

`func (o *TlsDeliveryProfileCreate) GetCiphers() string`

GetCiphers returns the Ciphers field if non-nil, zero value otherwise.

### GetCiphersOk

`func (o *TlsDeliveryProfileCreate) GetCiphersOk() (*string, bool)`

GetCiphersOk returns a tuple with the Ciphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphers

`func (o *TlsDeliveryProfileCreate) SetCiphers(v string)`

SetCiphers sets Ciphers field to given value.

### HasCiphers

`func (o *TlsDeliveryProfileCreate) HasCiphers() bool`

HasCiphers returns a boolean if a field has been set.

### GetCaCertificate

`func (o *TlsDeliveryProfileCreate) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *TlsDeliveryProfileCreate) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *TlsDeliveryProfileCreate) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *TlsDeliveryProfileCreate) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### GetCaCertificateDetails

`func (o *TlsDeliveryProfileCreate) GetCaCertificateDetails() X509CertificateSummaryType`

GetCaCertificateDetails returns the CaCertificateDetails field if non-nil, zero value otherwise.

### GetCaCertificateDetailsOk

`func (o *TlsDeliveryProfileCreate) GetCaCertificateDetailsOk() (*X509CertificateSummaryType, bool)`

GetCaCertificateDetailsOk returns a tuple with the CaCertificateDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificateDetails

`func (o *TlsDeliveryProfileCreate) SetCaCertificateDetails(v X509CertificateSummaryType)`

SetCaCertificateDetails sets CaCertificateDetails field to given value.

### HasCaCertificateDetails

`func (o *TlsDeliveryProfileCreate) HasCaCertificateDetails() bool`

HasCaCertificateDetails returns a boolean if a field has been set.

### GetContentProviderId

`func (o *TlsDeliveryProfileCreate) GetContentProviderId() int32`

GetContentProviderId returns the ContentProviderId field if non-nil, zero value otherwise.

### GetContentProviderIdOk

`func (o *TlsDeliveryProfileCreate) GetContentProviderIdOk() (*int32, bool)`

GetContentProviderIdOk returns a tuple with the ContentProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentProviderId

`func (o *TlsDeliveryProfileCreate) SetContentProviderId(v int32)`

SetContentProviderId sets ContentProviderId field to given value.


### GetCaCertificateSecretId

`func (o *TlsDeliveryProfileCreate) GetCaCertificateSecretId() int32`

GetCaCertificateSecretId returns the CaCertificateSecretId field if non-nil, zero value otherwise.

### GetCaCertificateSecretIdOk

`func (o *TlsDeliveryProfileCreate) GetCaCertificateSecretIdOk() (*int32, bool)`

GetCaCertificateSecretIdOk returns a tuple with the CaCertificateSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificateSecretId

`func (o *TlsDeliveryProfileCreate) SetCaCertificateSecretId(v int32)`

SetCaCertificateSecretId sets CaCertificateSecretId field to given value.

### HasCaCertificateSecretId

`func (o *TlsDeliveryProfileCreate) HasCaCertificateSecretId() bool`

HasCaCertificateSecretId returns a boolean if a field has been set.

### GetTlsVerificationDepth

`func (o *TlsDeliveryProfileCreate) GetTlsVerificationDepth() int32`

GetTlsVerificationDepth returns the TlsVerificationDepth field if non-nil, zero value otherwise.

### GetTlsVerificationDepthOk

`func (o *TlsDeliveryProfileCreate) GetTlsVerificationDepthOk() (*int32, bool)`

GetTlsVerificationDepthOk returns a tuple with the TlsVerificationDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsVerificationDepth

`func (o *TlsDeliveryProfileCreate) SetTlsVerificationDepth(v int32)`

SetTlsVerificationDepth sets TlsVerificationDepth field to given value.

### HasTlsVerificationDepth

`func (o *TlsDeliveryProfileCreate) HasTlsVerificationDepth() bool`

HasTlsVerificationDepth returns a boolean if a field has been set.

### GetTlsBufferSize

`func (o *TlsDeliveryProfileCreate) GetTlsBufferSize() TlsBufferSizeType`

GetTlsBufferSize returns the TlsBufferSize field if non-nil, zero value otherwise.

### GetTlsBufferSizeOk

`func (o *TlsDeliveryProfileCreate) GetTlsBufferSizeOk() (*TlsBufferSizeType, bool)`

GetTlsBufferSizeOk returns a tuple with the TlsBufferSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsBufferSize

`func (o *TlsDeliveryProfileCreate) SetTlsBufferSize(v TlsBufferSizeType)`

SetTlsBufferSize sets TlsBufferSize field to given value.

### HasTlsBufferSize

`func (o *TlsDeliveryProfileCreate) HasTlsBufferSize() bool`

HasTlsBufferSize returns a boolean if a field has been set.

### GetTlsCertificateKey

`func (o *TlsDeliveryProfileCreate) GetTlsCertificateKey() string`

GetTlsCertificateKey returns the TlsCertificateKey field if non-nil, zero value otherwise.

### GetTlsCertificateKeyOk

`func (o *TlsDeliveryProfileCreate) GetTlsCertificateKeyOk() (*string, bool)`

GetTlsCertificateKeyOk returns a tuple with the TlsCertificateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificateKey

`func (o *TlsDeliveryProfileCreate) SetTlsCertificateKey(v string)`

SetTlsCertificateKey sets TlsCertificateKey field to given value.

### HasTlsCertificateKey

`func (o *TlsDeliveryProfileCreate) HasTlsCertificateKey() bool`

HasTlsCertificateKey returns a boolean if a field has been set.

### GetEcdhCurve

`func (o *TlsDeliveryProfileCreate) GetEcdhCurve() EcdhCurveType`

GetEcdhCurve returns the EcdhCurve field if non-nil, zero value otherwise.

### GetEcdhCurveOk

`func (o *TlsDeliveryProfileCreate) GetEcdhCurveOk() (*EcdhCurveType, bool)`

GetEcdhCurveOk returns a tuple with the EcdhCurve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdhCurve

`func (o *TlsDeliveryProfileCreate) SetEcdhCurve(v EcdhCurveType)`

SetEcdhCurve sets EcdhCurve field to given value.

### HasEcdhCurve

`func (o *TlsDeliveryProfileCreate) HasEcdhCurve() bool`

HasEcdhCurve returns a boolean if a field has been set.

### GetTlsCertificateDetails

`func (o *TlsDeliveryProfileCreate) GetTlsCertificateDetails() X509CertificateSummaryType`

GetTlsCertificateDetails returns the TlsCertificateDetails field if non-nil, zero value otherwise.

### GetTlsCertificateDetailsOk

`func (o *TlsDeliveryProfileCreate) GetTlsCertificateDetailsOk() (*X509CertificateSummaryType, bool)`

GetTlsCertificateDetailsOk returns a tuple with the TlsCertificateDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificateDetails

`func (o *TlsDeliveryProfileCreate) SetTlsCertificateDetails(v X509CertificateSummaryType)`

SetTlsCertificateDetails sets TlsCertificateDetails field to given value.

### HasTlsCertificateDetails

`func (o *TlsDeliveryProfileCreate) HasTlsCertificateDetails() bool`

HasTlsCertificateDetails returns a boolean if a field has been set.

### GetTlsProtocols

`func (o *TlsDeliveryProfileCreate) GetTlsProtocols() []string`

GetTlsProtocols returns the TlsProtocols field if non-nil, zero value otherwise.

### GetTlsProtocolsOk

`func (o *TlsDeliveryProfileCreate) GetTlsProtocolsOk() (*[]string, bool)`

GetTlsProtocolsOk returns a tuple with the TlsProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsProtocols

`func (o *TlsDeliveryProfileCreate) SetTlsProtocols(v []string)`

SetTlsProtocols sets TlsProtocols field to given value.

### HasTlsProtocols

`func (o *TlsDeliveryProfileCreate) HasTlsProtocols() bool`

HasTlsProtocols returns a boolean if a field has been set.

### GetCipherSuites

`func (o *TlsDeliveryProfileCreate) GetCipherSuites() string`

GetCipherSuites returns the CipherSuites field if non-nil, zero value otherwise.

### GetCipherSuitesOk

`func (o *TlsDeliveryProfileCreate) GetCipherSuitesOk() (*string, bool)`

GetCipherSuitesOk returns a tuple with the CipherSuites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherSuites

`func (o *TlsDeliveryProfileCreate) SetCipherSuites(v string)`

SetCipherSuites sets CipherSuites field to given value.

### HasCipherSuites

`func (o *TlsDeliveryProfileCreate) HasCipherSuites() bool`

HasCipherSuites returns a boolean if a field has been set.

### GetEnableClientAuth

`func (o *TlsDeliveryProfileCreate) GetEnableClientAuth() bool`

GetEnableClientAuth returns the EnableClientAuth field if non-nil, zero value otherwise.

### GetEnableClientAuthOk

`func (o *TlsDeliveryProfileCreate) GetEnableClientAuthOk() (*bool, bool)`

GetEnableClientAuthOk returns a tuple with the EnableClientAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientAuth

`func (o *TlsDeliveryProfileCreate) SetEnableClientAuth(v bool)`

SetEnableClientAuth sets EnableClientAuth field to given value.

### HasEnableClientAuth

`func (o *TlsDeliveryProfileCreate) HasEnableClientAuth() bool`

HasEnableClientAuth returns a boolean if a field has been set.

### GetCrl

`func (o *TlsDeliveryProfileCreate) GetCrl() string`

GetCrl returns the Crl field if non-nil, zero value otherwise.

### GetCrlOk

`func (o *TlsDeliveryProfileCreate) GetCrlOk() (*string, bool)`

GetCrlOk returns a tuple with the Crl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrl

`func (o *TlsDeliveryProfileCreate) SetCrl(v string)`

SetCrl sets Crl field to given value.

### HasCrl

`func (o *TlsDeliveryProfileCreate) HasCrl() bool`

HasCrl returns a boolean if a field has been set.

### GetDescription

`func (o *TlsDeliveryProfileCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TlsDeliveryProfileCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TlsDeliveryProfileCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TlsDeliveryProfileCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


