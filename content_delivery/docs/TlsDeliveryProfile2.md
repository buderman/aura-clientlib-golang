# TlsDeliveryProfile2

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

### NewTlsDeliveryProfile2

`func NewTlsDeliveryProfile2(name string, contentProviderId int32, ) *TlsDeliveryProfile2`

NewTlsDeliveryProfile2 instantiates a new TlsDeliveryProfile2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsDeliveryProfile2WithDefaults

`func NewTlsDeliveryProfile2WithDefaults() *TlsDeliveryProfile2`

NewTlsDeliveryProfile2WithDefaults instantiates a new TlsDeliveryProfile2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTlsCertificate

`func (o *TlsDeliveryProfile2) GetTlsCertificate() string`

GetTlsCertificate returns the TlsCertificate field if non-nil, zero value otherwise.

### GetTlsCertificateOk

`func (o *TlsDeliveryProfile2) GetTlsCertificateOk() (*string, bool)`

GetTlsCertificateOk returns a tuple with the TlsCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificate

`func (o *TlsDeliveryProfile2) SetTlsCertificate(v string)`

SetTlsCertificate sets TlsCertificate field to given value.

### HasTlsCertificate

`func (o *TlsDeliveryProfile2) HasTlsCertificate() bool`

HasTlsCertificate returns a boolean if a field has been set.

### GetTlsCertificateAndKeySecretId

`func (o *TlsDeliveryProfile2) GetTlsCertificateAndKeySecretId() int32`

GetTlsCertificateAndKeySecretId returns the TlsCertificateAndKeySecretId field if non-nil, zero value otherwise.

### GetTlsCertificateAndKeySecretIdOk

`func (o *TlsDeliveryProfile2) GetTlsCertificateAndKeySecretIdOk() (*int32, bool)`

GetTlsCertificateAndKeySecretIdOk returns a tuple with the TlsCertificateAndKeySecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificateAndKeySecretId

`func (o *TlsDeliveryProfile2) SetTlsCertificateAndKeySecretId(v int32)`

SetTlsCertificateAndKeySecretId sets TlsCertificateAndKeySecretId field to given value.

### HasTlsCertificateAndKeySecretId

`func (o *TlsDeliveryProfile2) HasTlsCertificateAndKeySecretId() bool`

HasTlsCertificateAndKeySecretId returns a boolean if a field has been set.

### GetName

`func (o *TlsDeliveryProfile2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TlsDeliveryProfile2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TlsDeliveryProfile2) SetName(v string)`

SetName sets Name field to given value.


### GetCiphers

`func (o *TlsDeliveryProfile2) GetCiphers() string`

GetCiphers returns the Ciphers field if non-nil, zero value otherwise.

### GetCiphersOk

`func (o *TlsDeliveryProfile2) GetCiphersOk() (*string, bool)`

GetCiphersOk returns a tuple with the Ciphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphers

`func (o *TlsDeliveryProfile2) SetCiphers(v string)`

SetCiphers sets Ciphers field to given value.

### HasCiphers

`func (o *TlsDeliveryProfile2) HasCiphers() bool`

HasCiphers returns a boolean if a field has been set.

### GetCaCertificate

`func (o *TlsDeliveryProfile2) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *TlsDeliveryProfile2) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *TlsDeliveryProfile2) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *TlsDeliveryProfile2) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### GetCaCertificateDetails

`func (o *TlsDeliveryProfile2) GetCaCertificateDetails() X509CertificateSummaryType`

GetCaCertificateDetails returns the CaCertificateDetails field if non-nil, zero value otherwise.

### GetCaCertificateDetailsOk

`func (o *TlsDeliveryProfile2) GetCaCertificateDetailsOk() (*X509CertificateSummaryType, bool)`

GetCaCertificateDetailsOk returns a tuple with the CaCertificateDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificateDetails

`func (o *TlsDeliveryProfile2) SetCaCertificateDetails(v X509CertificateSummaryType)`

SetCaCertificateDetails sets CaCertificateDetails field to given value.

### HasCaCertificateDetails

`func (o *TlsDeliveryProfile2) HasCaCertificateDetails() bool`

HasCaCertificateDetails returns a boolean if a field has been set.

### GetContentProviderId

`func (o *TlsDeliveryProfile2) GetContentProviderId() int32`

GetContentProviderId returns the ContentProviderId field if non-nil, zero value otherwise.

### GetContentProviderIdOk

`func (o *TlsDeliveryProfile2) GetContentProviderIdOk() (*int32, bool)`

GetContentProviderIdOk returns a tuple with the ContentProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentProviderId

`func (o *TlsDeliveryProfile2) SetContentProviderId(v int32)`

SetContentProviderId sets ContentProviderId field to given value.


### GetCaCertificateSecretId

`func (o *TlsDeliveryProfile2) GetCaCertificateSecretId() int32`

GetCaCertificateSecretId returns the CaCertificateSecretId field if non-nil, zero value otherwise.

### GetCaCertificateSecretIdOk

`func (o *TlsDeliveryProfile2) GetCaCertificateSecretIdOk() (*int32, bool)`

GetCaCertificateSecretIdOk returns a tuple with the CaCertificateSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificateSecretId

`func (o *TlsDeliveryProfile2) SetCaCertificateSecretId(v int32)`

SetCaCertificateSecretId sets CaCertificateSecretId field to given value.

### HasCaCertificateSecretId

`func (o *TlsDeliveryProfile2) HasCaCertificateSecretId() bool`

HasCaCertificateSecretId returns a boolean if a field has been set.

### GetTlsVerificationDepth

`func (o *TlsDeliveryProfile2) GetTlsVerificationDepth() int32`

GetTlsVerificationDepth returns the TlsVerificationDepth field if non-nil, zero value otherwise.

### GetTlsVerificationDepthOk

`func (o *TlsDeliveryProfile2) GetTlsVerificationDepthOk() (*int32, bool)`

GetTlsVerificationDepthOk returns a tuple with the TlsVerificationDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsVerificationDepth

`func (o *TlsDeliveryProfile2) SetTlsVerificationDepth(v int32)`

SetTlsVerificationDepth sets TlsVerificationDepth field to given value.

### HasTlsVerificationDepth

`func (o *TlsDeliveryProfile2) HasTlsVerificationDepth() bool`

HasTlsVerificationDepth returns a boolean if a field has been set.

### GetTlsBufferSize

`func (o *TlsDeliveryProfile2) GetTlsBufferSize() TlsBufferSizeType`

GetTlsBufferSize returns the TlsBufferSize field if non-nil, zero value otherwise.

### GetTlsBufferSizeOk

`func (o *TlsDeliveryProfile2) GetTlsBufferSizeOk() (*TlsBufferSizeType, bool)`

GetTlsBufferSizeOk returns a tuple with the TlsBufferSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsBufferSize

`func (o *TlsDeliveryProfile2) SetTlsBufferSize(v TlsBufferSizeType)`

SetTlsBufferSize sets TlsBufferSize field to given value.

### HasTlsBufferSize

`func (o *TlsDeliveryProfile2) HasTlsBufferSize() bool`

HasTlsBufferSize returns a boolean if a field has been set.

### GetTlsCertificateKey

`func (o *TlsDeliveryProfile2) GetTlsCertificateKey() string`

GetTlsCertificateKey returns the TlsCertificateKey field if non-nil, zero value otherwise.

### GetTlsCertificateKeyOk

`func (o *TlsDeliveryProfile2) GetTlsCertificateKeyOk() (*string, bool)`

GetTlsCertificateKeyOk returns a tuple with the TlsCertificateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificateKey

`func (o *TlsDeliveryProfile2) SetTlsCertificateKey(v string)`

SetTlsCertificateKey sets TlsCertificateKey field to given value.

### HasTlsCertificateKey

`func (o *TlsDeliveryProfile2) HasTlsCertificateKey() bool`

HasTlsCertificateKey returns a boolean if a field has been set.

### GetEcdhCurve

`func (o *TlsDeliveryProfile2) GetEcdhCurve() EcdhCurveType`

GetEcdhCurve returns the EcdhCurve field if non-nil, zero value otherwise.

### GetEcdhCurveOk

`func (o *TlsDeliveryProfile2) GetEcdhCurveOk() (*EcdhCurveType, bool)`

GetEcdhCurveOk returns a tuple with the EcdhCurve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdhCurve

`func (o *TlsDeliveryProfile2) SetEcdhCurve(v EcdhCurveType)`

SetEcdhCurve sets EcdhCurve field to given value.

### HasEcdhCurve

`func (o *TlsDeliveryProfile2) HasEcdhCurve() bool`

HasEcdhCurve returns a boolean if a field has been set.

### GetTlsCertificateDetails

`func (o *TlsDeliveryProfile2) GetTlsCertificateDetails() X509CertificateSummaryType`

GetTlsCertificateDetails returns the TlsCertificateDetails field if non-nil, zero value otherwise.

### GetTlsCertificateDetailsOk

`func (o *TlsDeliveryProfile2) GetTlsCertificateDetailsOk() (*X509CertificateSummaryType, bool)`

GetTlsCertificateDetailsOk returns a tuple with the TlsCertificateDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificateDetails

`func (o *TlsDeliveryProfile2) SetTlsCertificateDetails(v X509CertificateSummaryType)`

SetTlsCertificateDetails sets TlsCertificateDetails field to given value.

### HasTlsCertificateDetails

`func (o *TlsDeliveryProfile2) HasTlsCertificateDetails() bool`

HasTlsCertificateDetails returns a boolean if a field has been set.

### GetTlsProtocols

`func (o *TlsDeliveryProfile2) GetTlsProtocols() []string`

GetTlsProtocols returns the TlsProtocols field if non-nil, zero value otherwise.

### GetTlsProtocolsOk

`func (o *TlsDeliveryProfile2) GetTlsProtocolsOk() (*[]string, bool)`

GetTlsProtocolsOk returns a tuple with the TlsProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsProtocols

`func (o *TlsDeliveryProfile2) SetTlsProtocols(v []string)`

SetTlsProtocols sets TlsProtocols field to given value.

### HasTlsProtocols

`func (o *TlsDeliveryProfile2) HasTlsProtocols() bool`

HasTlsProtocols returns a boolean if a field has been set.

### GetCipherSuites

`func (o *TlsDeliveryProfile2) GetCipherSuites() string`

GetCipherSuites returns the CipherSuites field if non-nil, zero value otherwise.

### GetCipherSuitesOk

`func (o *TlsDeliveryProfile2) GetCipherSuitesOk() (*string, bool)`

GetCipherSuitesOk returns a tuple with the CipherSuites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherSuites

`func (o *TlsDeliveryProfile2) SetCipherSuites(v string)`

SetCipherSuites sets CipherSuites field to given value.

### HasCipherSuites

`func (o *TlsDeliveryProfile2) HasCipherSuites() bool`

HasCipherSuites returns a boolean if a field has been set.

### GetEnableClientAuth

`func (o *TlsDeliveryProfile2) GetEnableClientAuth() bool`

GetEnableClientAuth returns the EnableClientAuth field if non-nil, zero value otherwise.

### GetEnableClientAuthOk

`func (o *TlsDeliveryProfile2) GetEnableClientAuthOk() (*bool, bool)`

GetEnableClientAuthOk returns a tuple with the EnableClientAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientAuth

`func (o *TlsDeliveryProfile2) SetEnableClientAuth(v bool)`

SetEnableClientAuth sets EnableClientAuth field to given value.

### HasEnableClientAuth

`func (o *TlsDeliveryProfile2) HasEnableClientAuth() bool`

HasEnableClientAuth returns a boolean if a field has been set.

### GetCrl

`func (o *TlsDeliveryProfile2) GetCrl() string`

GetCrl returns the Crl field if non-nil, zero value otherwise.

### GetCrlOk

`func (o *TlsDeliveryProfile2) GetCrlOk() (*string, bool)`

GetCrlOk returns a tuple with the Crl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrl

`func (o *TlsDeliveryProfile2) SetCrl(v string)`

SetCrl sets Crl field to given value.

### HasCrl

`func (o *TlsDeliveryProfile2) HasCrl() bool`

HasCrl returns a boolean if a field has been set.

### GetDescription

`func (o *TlsDeliveryProfile2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TlsDeliveryProfile2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TlsDeliveryProfile2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TlsDeliveryProfile2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


