# TlsDeliveryProfileReadDetailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TlsDeliveryProfileId** | **int32** |  | 
**TlsCertificate** | Pointer to **string** | An OpenSSL PEM-formatted certificate associated with the trusted CA that signed the origin&#39;s certificate. | [optional] 
**TlsCertificateAndKeySecretId** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Ciphers** | Pointer to **string** | An OpenSSL style [ciphers string](https://www.openssl.org/docs/apps/ciphers.html#CIPHER-STRINGS). | [optional] 
**CaCertificate** | Pointer to **string** | An OpenSSL PEM-formatted certificate associated with the trusted CA that signed the origin&#39;s certificate. | [optional] 
**CaCertificateDetails** | Pointer to [**X509CertificateSummaryType**](X509CertificateSummaryType.md) |  | [optional] 
**ContentProviderId** | **int32** |  | 
**CaCertificateSecretId** | Pointer to **int32** |  | [optional] 
**TlsVerificationDepth** | Pointer to **int32** | Sets the verification depth in the client certificates chain. | [optional] 
**CrlDetails** | Pointer to [**X509CrlSummaryType**](X509CrlSummaryType.md) |  | [optional] 
**TlsCertificateKey** | Pointer to **string** | An OpenSSL PEM-formatted certificate associated with the trusted CA that signed the origin&#39;s certificate. | [optional] 
**EcdhCurve** | Pointer to [**EcdhCurveType**](EcdhCurveType.md) |  | [optional] 
**TlsCertificateDetails** | Pointer to [**X509CertificateSummaryType**](X509CertificateSummaryType.md) |  | [optional] 
**TlsProtocols** | Pointer to **[]string** | TLS protocols supported, either &#x60;TLSV1&#x60;, &#x60;TLSV1_1&#x60;, &#x60;TLSV1_2&#x60;, or &#x60;TLSV1_3&#x60;.  Defaults to &#x60;TLSV1_2&#x60;, &#x60;TLSV1_3&#x60;.  The use of &#x60;SSLV2&#x60; or &#x60;SSLV3&#x60; are discouraged for security reasons. | [optional] [default to ["TLSV1_2","TLSV1_3"]]
**CipherSuites** | Pointer to **string** | An OpenSSL style [ciphers string](https://www.openssl.org/docs/apps/ciphers.html#CIPHER-STRINGS). | [optional] 
**TlsBufferSize** | Pointer to [**TlsBufferSizeType**](TlsBufferSizeType.md) |  | [optional] 
**EnableClientAuth** | Pointer to **bool** | Allows you to enable or disable client authorization. Enabling this feature causes HyperCache to authenticate client certificates using the CA certificate for each TLS connection. | [optional] [default to false]
**Crl** | Pointer to **string** | Certificate Revocation List. An OpenSSL PEM-formatted list of revoked client certificates. | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewTlsDeliveryProfileReadDetailed

`func NewTlsDeliveryProfileReadDetailed(tlsDeliveryProfileId int32, name string, contentProviderId int32, ) *TlsDeliveryProfileReadDetailed`

NewTlsDeliveryProfileReadDetailed instantiates a new TlsDeliveryProfileReadDetailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsDeliveryProfileReadDetailedWithDefaults

`func NewTlsDeliveryProfileReadDetailedWithDefaults() *TlsDeliveryProfileReadDetailed`

NewTlsDeliveryProfileReadDetailedWithDefaults instantiates a new TlsDeliveryProfileReadDetailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTlsDeliveryProfileId

`func (o *TlsDeliveryProfileReadDetailed) GetTlsDeliveryProfileId() int32`

GetTlsDeliveryProfileId returns the TlsDeliveryProfileId field if non-nil, zero value otherwise.

### GetTlsDeliveryProfileIdOk

`func (o *TlsDeliveryProfileReadDetailed) GetTlsDeliveryProfileIdOk() (*int32, bool)`

GetTlsDeliveryProfileIdOk returns a tuple with the TlsDeliveryProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsDeliveryProfileId

`func (o *TlsDeliveryProfileReadDetailed) SetTlsDeliveryProfileId(v int32)`

SetTlsDeliveryProfileId sets TlsDeliveryProfileId field to given value.


### GetTlsCertificate

`func (o *TlsDeliveryProfileReadDetailed) GetTlsCertificate() string`

GetTlsCertificate returns the TlsCertificate field if non-nil, zero value otherwise.

### GetTlsCertificateOk

`func (o *TlsDeliveryProfileReadDetailed) GetTlsCertificateOk() (*string, bool)`

GetTlsCertificateOk returns a tuple with the TlsCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificate

`func (o *TlsDeliveryProfileReadDetailed) SetTlsCertificate(v string)`

SetTlsCertificate sets TlsCertificate field to given value.

### HasTlsCertificate

`func (o *TlsDeliveryProfileReadDetailed) HasTlsCertificate() bool`

HasTlsCertificate returns a boolean if a field has been set.

### GetTlsCertificateAndKeySecretId

`func (o *TlsDeliveryProfileReadDetailed) GetTlsCertificateAndKeySecretId() int32`

GetTlsCertificateAndKeySecretId returns the TlsCertificateAndKeySecretId field if non-nil, zero value otherwise.

### GetTlsCertificateAndKeySecretIdOk

`func (o *TlsDeliveryProfileReadDetailed) GetTlsCertificateAndKeySecretIdOk() (*int32, bool)`

GetTlsCertificateAndKeySecretIdOk returns a tuple with the TlsCertificateAndKeySecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificateAndKeySecretId

`func (o *TlsDeliveryProfileReadDetailed) SetTlsCertificateAndKeySecretId(v int32)`

SetTlsCertificateAndKeySecretId sets TlsCertificateAndKeySecretId field to given value.

### HasTlsCertificateAndKeySecretId

`func (o *TlsDeliveryProfileReadDetailed) HasTlsCertificateAndKeySecretId() bool`

HasTlsCertificateAndKeySecretId returns a boolean if a field has been set.

### GetName

`func (o *TlsDeliveryProfileReadDetailed) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TlsDeliveryProfileReadDetailed) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TlsDeliveryProfileReadDetailed) SetName(v string)`

SetName sets Name field to given value.


### GetCiphers

`func (o *TlsDeliveryProfileReadDetailed) GetCiphers() string`

GetCiphers returns the Ciphers field if non-nil, zero value otherwise.

### GetCiphersOk

`func (o *TlsDeliveryProfileReadDetailed) GetCiphersOk() (*string, bool)`

GetCiphersOk returns a tuple with the Ciphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphers

`func (o *TlsDeliveryProfileReadDetailed) SetCiphers(v string)`

SetCiphers sets Ciphers field to given value.

### HasCiphers

`func (o *TlsDeliveryProfileReadDetailed) HasCiphers() bool`

HasCiphers returns a boolean if a field has been set.

### GetCaCertificate

`func (o *TlsDeliveryProfileReadDetailed) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *TlsDeliveryProfileReadDetailed) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *TlsDeliveryProfileReadDetailed) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *TlsDeliveryProfileReadDetailed) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### GetCaCertificateDetails

`func (o *TlsDeliveryProfileReadDetailed) GetCaCertificateDetails() X509CertificateSummaryType`

GetCaCertificateDetails returns the CaCertificateDetails field if non-nil, zero value otherwise.

### GetCaCertificateDetailsOk

`func (o *TlsDeliveryProfileReadDetailed) GetCaCertificateDetailsOk() (*X509CertificateSummaryType, bool)`

GetCaCertificateDetailsOk returns a tuple with the CaCertificateDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificateDetails

`func (o *TlsDeliveryProfileReadDetailed) SetCaCertificateDetails(v X509CertificateSummaryType)`

SetCaCertificateDetails sets CaCertificateDetails field to given value.

### HasCaCertificateDetails

`func (o *TlsDeliveryProfileReadDetailed) HasCaCertificateDetails() bool`

HasCaCertificateDetails returns a boolean if a field has been set.

### GetContentProviderId

`func (o *TlsDeliveryProfileReadDetailed) GetContentProviderId() int32`

GetContentProviderId returns the ContentProviderId field if non-nil, zero value otherwise.

### GetContentProviderIdOk

`func (o *TlsDeliveryProfileReadDetailed) GetContentProviderIdOk() (*int32, bool)`

GetContentProviderIdOk returns a tuple with the ContentProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentProviderId

`func (o *TlsDeliveryProfileReadDetailed) SetContentProviderId(v int32)`

SetContentProviderId sets ContentProviderId field to given value.


### GetCaCertificateSecretId

`func (o *TlsDeliveryProfileReadDetailed) GetCaCertificateSecretId() int32`

GetCaCertificateSecretId returns the CaCertificateSecretId field if non-nil, zero value otherwise.

### GetCaCertificateSecretIdOk

`func (o *TlsDeliveryProfileReadDetailed) GetCaCertificateSecretIdOk() (*int32, bool)`

GetCaCertificateSecretIdOk returns a tuple with the CaCertificateSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificateSecretId

`func (o *TlsDeliveryProfileReadDetailed) SetCaCertificateSecretId(v int32)`

SetCaCertificateSecretId sets CaCertificateSecretId field to given value.

### HasCaCertificateSecretId

`func (o *TlsDeliveryProfileReadDetailed) HasCaCertificateSecretId() bool`

HasCaCertificateSecretId returns a boolean if a field has been set.

### GetTlsVerificationDepth

`func (o *TlsDeliveryProfileReadDetailed) GetTlsVerificationDepth() int32`

GetTlsVerificationDepth returns the TlsVerificationDepth field if non-nil, zero value otherwise.

### GetTlsVerificationDepthOk

`func (o *TlsDeliveryProfileReadDetailed) GetTlsVerificationDepthOk() (*int32, bool)`

GetTlsVerificationDepthOk returns a tuple with the TlsVerificationDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsVerificationDepth

`func (o *TlsDeliveryProfileReadDetailed) SetTlsVerificationDepth(v int32)`

SetTlsVerificationDepth sets TlsVerificationDepth field to given value.

### HasTlsVerificationDepth

`func (o *TlsDeliveryProfileReadDetailed) HasTlsVerificationDepth() bool`

HasTlsVerificationDepth returns a boolean if a field has been set.

### GetCrlDetails

`func (o *TlsDeliveryProfileReadDetailed) GetCrlDetails() X509CrlSummaryType`

GetCrlDetails returns the CrlDetails field if non-nil, zero value otherwise.

### GetCrlDetailsOk

`func (o *TlsDeliveryProfileReadDetailed) GetCrlDetailsOk() (*X509CrlSummaryType, bool)`

GetCrlDetailsOk returns a tuple with the CrlDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlDetails

`func (o *TlsDeliveryProfileReadDetailed) SetCrlDetails(v X509CrlSummaryType)`

SetCrlDetails sets CrlDetails field to given value.

### HasCrlDetails

`func (o *TlsDeliveryProfileReadDetailed) HasCrlDetails() bool`

HasCrlDetails returns a boolean if a field has been set.

### GetTlsCertificateKey

`func (o *TlsDeliveryProfileReadDetailed) GetTlsCertificateKey() string`

GetTlsCertificateKey returns the TlsCertificateKey field if non-nil, zero value otherwise.

### GetTlsCertificateKeyOk

`func (o *TlsDeliveryProfileReadDetailed) GetTlsCertificateKeyOk() (*string, bool)`

GetTlsCertificateKeyOk returns a tuple with the TlsCertificateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificateKey

`func (o *TlsDeliveryProfileReadDetailed) SetTlsCertificateKey(v string)`

SetTlsCertificateKey sets TlsCertificateKey field to given value.

### HasTlsCertificateKey

`func (o *TlsDeliveryProfileReadDetailed) HasTlsCertificateKey() bool`

HasTlsCertificateKey returns a boolean if a field has been set.

### GetEcdhCurve

`func (o *TlsDeliveryProfileReadDetailed) GetEcdhCurve() EcdhCurveType`

GetEcdhCurve returns the EcdhCurve field if non-nil, zero value otherwise.

### GetEcdhCurveOk

`func (o *TlsDeliveryProfileReadDetailed) GetEcdhCurveOk() (*EcdhCurveType, bool)`

GetEcdhCurveOk returns a tuple with the EcdhCurve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdhCurve

`func (o *TlsDeliveryProfileReadDetailed) SetEcdhCurve(v EcdhCurveType)`

SetEcdhCurve sets EcdhCurve field to given value.

### HasEcdhCurve

`func (o *TlsDeliveryProfileReadDetailed) HasEcdhCurve() bool`

HasEcdhCurve returns a boolean if a field has been set.

### GetTlsCertificateDetails

`func (o *TlsDeliveryProfileReadDetailed) GetTlsCertificateDetails() X509CertificateSummaryType`

GetTlsCertificateDetails returns the TlsCertificateDetails field if non-nil, zero value otherwise.

### GetTlsCertificateDetailsOk

`func (o *TlsDeliveryProfileReadDetailed) GetTlsCertificateDetailsOk() (*X509CertificateSummaryType, bool)`

GetTlsCertificateDetailsOk returns a tuple with the TlsCertificateDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificateDetails

`func (o *TlsDeliveryProfileReadDetailed) SetTlsCertificateDetails(v X509CertificateSummaryType)`

SetTlsCertificateDetails sets TlsCertificateDetails field to given value.

### HasTlsCertificateDetails

`func (o *TlsDeliveryProfileReadDetailed) HasTlsCertificateDetails() bool`

HasTlsCertificateDetails returns a boolean if a field has been set.

### GetTlsProtocols

`func (o *TlsDeliveryProfileReadDetailed) GetTlsProtocols() []string`

GetTlsProtocols returns the TlsProtocols field if non-nil, zero value otherwise.

### GetTlsProtocolsOk

`func (o *TlsDeliveryProfileReadDetailed) GetTlsProtocolsOk() (*[]string, bool)`

GetTlsProtocolsOk returns a tuple with the TlsProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsProtocols

`func (o *TlsDeliveryProfileReadDetailed) SetTlsProtocols(v []string)`

SetTlsProtocols sets TlsProtocols field to given value.

### HasTlsProtocols

`func (o *TlsDeliveryProfileReadDetailed) HasTlsProtocols() bool`

HasTlsProtocols returns a boolean if a field has been set.

### GetCipherSuites

`func (o *TlsDeliveryProfileReadDetailed) GetCipherSuites() string`

GetCipherSuites returns the CipherSuites field if non-nil, zero value otherwise.

### GetCipherSuitesOk

`func (o *TlsDeliveryProfileReadDetailed) GetCipherSuitesOk() (*string, bool)`

GetCipherSuitesOk returns a tuple with the CipherSuites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherSuites

`func (o *TlsDeliveryProfileReadDetailed) SetCipherSuites(v string)`

SetCipherSuites sets CipherSuites field to given value.

### HasCipherSuites

`func (o *TlsDeliveryProfileReadDetailed) HasCipherSuites() bool`

HasCipherSuites returns a boolean if a field has been set.

### GetTlsBufferSize

`func (o *TlsDeliveryProfileReadDetailed) GetTlsBufferSize() TlsBufferSizeType`

GetTlsBufferSize returns the TlsBufferSize field if non-nil, zero value otherwise.

### GetTlsBufferSizeOk

`func (o *TlsDeliveryProfileReadDetailed) GetTlsBufferSizeOk() (*TlsBufferSizeType, bool)`

GetTlsBufferSizeOk returns a tuple with the TlsBufferSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsBufferSize

`func (o *TlsDeliveryProfileReadDetailed) SetTlsBufferSize(v TlsBufferSizeType)`

SetTlsBufferSize sets TlsBufferSize field to given value.

### HasTlsBufferSize

`func (o *TlsDeliveryProfileReadDetailed) HasTlsBufferSize() bool`

HasTlsBufferSize returns a boolean if a field has been set.

### GetEnableClientAuth

`func (o *TlsDeliveryProfileReadDetailed) GetEnableClientAuth() bool`

GetEnableClientAuth returns the EnableClientAuth field if non-nil, zero value otherwise.

### GetEnableClientAuthOk

`func (o *TlsDeliveryProfileReadDetailed) GetEnableClientAuthOk() (*bool, bool)`

GetEnableClientAuthOk returns a tuple with the EnableClientAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientAuth

`func (o *TlsDeliveryProfileReadDetailed) SetEnableClientAuth(v bool)`

SetEnableClientAuth sets EnableClientAuth field to given value.

### HasEnableClientAuth

`func (o *TlsDeliveryProfileReadDetailed) HasEnableClientAuth() bool`

HasEnableClientAuth returns a boolean if a field has been set.

### GetCrl

`func (o *TlsDeliveryProfileReadDetailed) GetCrl() string`

GetCrl returns the Crl field if non-nil, zero value otherwise.

### GetCrlOk

`func (o *TlsDeliveryProfileReadDetailed) GetCrlOk() (*string, bool)`

GetCrlOk returns a tuple with the Crl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrl

`func (o *TlsDeliveryProfileReadDetailed) SetCrl(v string)`

SetCrl sets Crl field to given value.

### HasCrl

`func (o *TlsDeliveryProfileReadDetailed) HasCrl() bool`

HasCrl returns a boolean if a field has been set.

### GetDescription

`func (o *TlsDeliveryProfileReadDetailed) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TlsDeliveryProfileReadDetailed) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TlsDeliveryProfileReadDetailed) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TlsDeliveryProfileReadDetailed) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


