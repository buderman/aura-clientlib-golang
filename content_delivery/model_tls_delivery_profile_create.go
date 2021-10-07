/*
content-delivery

Aura LCDN Content Delivery API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package content_delivery

import (
	"encoding/json"
)

// TlsDeliveryProfileCreate TLS delivery profile create
type TlsDeliveryProfileCreate struct {
	// An OpenSSL PEM-formatted certificate associated with the trusted CA that signed the origin's certificate.
	TlsCertificate *string `json:"tlsCertificate,omitempty"`
	TlsCertificateAndKeySecretId *int32 `json:"tlsCertificateAndKeySecretId,omitempty"`
	Name string `json:"name"`
	// An OpenSSL style [ciphers string](https://www.openssl.org/docs/apps/ciphers.html#CIPHER-STRINGS).
	Ciphers *string `json:"ciphers,omitempty"`
	// An OpenSSL PEM-formatted certificate associated with the trusted CA that signed the origin's certificate.
	CaCertificate *string `json:"caCertificate,omitempty"`
	CaCertificateDetails *X509CertificateSummaryType `json:"caCertificateDetails,omitempty"`
	ContentProviderId int32 `json:"contentProviderId"`
	CaCertificateSecretId *int32 `json:"caCertificateSecretId,omitempty"`
	// Sets the verification depth in the client certificates chain.
	TlsVerificationDepth *int32 `json:"tlsVerificationDepth,omitempty"`
	TlsBufferSize *TlsBufferSizeType `json:"tlsBufferSize,omitempty"`
	// An OpenSSL PEM-formatted certificate associated with the trusted CA that signed the origin's certificate.
	TlsCertificateKey *string `json:"tlsCertificateKey,omitempty"`
	EcdhCurve *EcdhCurveType `json:"ecdhCurve,omitempty"`
	TlsCertificateDetails *X509CertificateSummaryType `json:"tlsCertificateDetails,omitempty"`
	// TLS protocols supported, either `TLSV1`, `TLSV1_1`, `TLSV1_2`, or `TLSV1_3`.  Defaults to `TLSV1_2`, `TLSV1_3`.  The use of `SSLV2` or `SSLV3` are discouraged for security reasons.
	TlsProtocols *[]string `json:"tlsProtocols,omitempty"`
	// An OpenSSL style [ciphers string](https://www.openssl.org/docs/apps/ciphers.html#CIPHER-STRINGS).
	CipherSuites *string `json:"cipherSuites,omitempty"`
	// Allows you to enable or disable client authorization. Enabling this feature causes HyperCache to authenticate client certificates using the CA certificate for each TLS connection.
	EnableClientAuth *bool `json:"enableClientAuth,omitempty"`
	// Certificate Revocation List. An OpenSSL PEM-formatted list of revoked client certificates.
	Crl *string `json:"crl,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewTlsDeliveryProfileCreate instantiates a new TlsDeliveryProfileCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsDeliveryProfileCreate(name string, contentProviderId int32) *TlsDeliveryProfileCreate {
	this := TlsDeliveryProfileCreate{}
	this.Name = name
	this.ContentProviderId = contentProviderId
	var enableClientAuth bool = false
	this.EnableClientAuth = &enableClientAuth
	return &this
}

// NewTlsDeliveryProfileCreateWithDefaults instantiates a new TlsDeliveryProfileCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsDeliveryProfileCreateWithDefaults() *TlsDeliveryProfileCreate {
	this := TlsDeliveryProfileCreate{}
	var enableClientAuth bool = false
	this.EnableClientAuth = &enableClientAuth
	return &this
}

// GetTlsCertificate returns the TlsCertificate field value if set, zero value otherwise.
func (o *TlsDeliveryProfileCreate) GetTlsCertificate() string {
	if o == nil || o.TlsCertificate == nil {
		var ret string
		return ret
	}
	return *o.TlsCertificate
}

// GetTlsCertificateOk returns a tuple with the TlsCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsDeliveryProfileCreate) GetTlsCertificateOk() (*string, bool) {
	if o == nil || o.TlsCertificate == nil {
		return nil, false
	}
	return o.TlsCertificate, true
}

// HasTlsCertificate returns a boolean if a field has been set.
func (o *TlsDeliveryProfileCreate) HasTlsCertificate() bool {
	if o != nil && o.TlsCertificate != nil {
		return true
	}

	return false
}

// SetTlsCertificate gets a reference to the given string and assigns it to the TlsCertificate field.
func (o *TlsDeliveryProfileCreate) SetTlsCertificate(v string) {
	o.TlsCertificate = &v
}

// GetTlsCertificateAndKeySecretId returns the TlsCertificateAndKeySecretId field value if set, zero value otherwise.
func (o *TlsDeliveryProfileCreate) GetTlsCertificateAndKeySecretId() int32 {
	if o == nil || o.TlsCertificateAndKeySecretId == nil {
		var ret int32
		return ret
	}
	return *o.TlsCertificateAndKeySecretId
}

// GetTlsCertificateAndKeySecretIdOk returns a tuple with the TlsCertificateAndKeySecretId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsDeliveryProfileCreate) GetTlsCertificateAndKeySecretIdOk() (*int32, bool) {
	if o == nil || o.TlsCertificateAndKeySecretId == nil {
		return nil, false
	}
	return o.TlsCertificateAndKeySecretId, true
}

// HasTlsCertificateAndKeySecretId returns a boolean if a field has been set.
func (o *TlsDeliveryProfileCreate) HasTlsCertificateAndKeySecretId() bool {
	if o != nil && o.TlsCertificateAndKeySecretId != nil {
		return true
	}

	return false
}

// SetTlsCertificateAndKeySecretId gets a reference to the given int32 and assigns it to the TlsCertificateAndKeySecretId field.
func (o *TlsDeliveryProfileCreate) SetTlsCertificateAndKeySecretId(v int32) {
	o.TlsCertificateAndKeySecretId = &v
}

// GetName returns the Name field value
func (o *TlsDeliveryProfileCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TlsDeliveryProfileCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TlsDeliveryProfileCreate) SetName(v string) {
	o.Name = v
}

// GetCiphers returns the Ciphers field value if set, zero value otherwise.
func (o *TlsDeliveryProfileCreate) GetCiphers() string {
	if o == nil || o.Ciphers == nil {
		var ret string
		return ret
	}
	return *o.Ciphers
}

// GetCiphersOk returns a tuple with the Ciphers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsDeliveryProfileCreate) GetCiphersOk() (*string, bool) {
	if o == nil || o.Ciphers == nil {
		return nil, false
	}
	return o.Ciphers, true
}

// HasCiphers returns a boolean if a field has been set.
func (o *TlsDeliveryProfileCreate) HasCiphers() bool {
	if o != nil && o.Ciphers != nil {
		return true
	}

	return false
}

// SetCiphers gets a reference to the given string and assigns it to the Ciphers field.
func (o *TlsDeliveryProfileCreate) SetCiphers(v string) {
	o.Ciphers = &v
}

// GetCaCertificate returns the CaCertificate field value if set, zero value otherwise.
func (o *TlsDeliveryProfileCreate) GetCaCertificate() string {
	if o == nil || o.CaCertificate == nil {
		var ret string
		return ret
	}
	return *o.CaCertificate
}

// GetCaCertificateOk returns a tuple with the CaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsDeliveryProfileCreate) GetCaCertificateOk() (*string, bool) {
	if o == nil || o.CaCertificate == nil {
		return nil, false
	}
	return o.CaCertificate, true
}

// HasCaCertificate returns a boolean if a field has been set.
func (o *TlsDeliveryProfileCreate) HasCaCertificate() bool {
	if o != nil && o.CaCertificate != nil {
		return true
	}

	return false
}

// SetCaCertificate gets a reference to the given string and assigns it to the CaCertificate field.
func (o *TlsDeliveryProfileCreate) SetCaCertificate(v string) {
	o.CaCertificate = &v
}

// GetCaCertificateDetails returns the CaCertificateDetails field value if set, zero value otherwise.
func (o *TlsDeliveryProfileCreate) GetCaCertificateDetails() X509CertificateSummaryType {
	if o == nil || o.CaCertificateDetails == nil {
		var ret X509CertificateSummaryType
		return ret
	}
	return *o.CaCertificateDetails
}

// GetCaCertificateDetailsOk returns a tuple with the CaCertificateDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsDeliveryProfileCreate) GetCaCertificateDetailsOk() (*X509CertificateSummaryType, bool) {
	if o == nil || o.CaCertificateDetails == nil {
		return nil, false
	}
	return o.CaCertificateDetails, true
}

// HasCaCertificateDetails returns a boolean if a field has been set.
func (o *TlsDeliveryProfileCreate) HasCaCertificateDetails() bool {
	if o != nil && o.CaCertificateDetails != nil {
		return true
	}

	return false
}

// SetCaCertificateDetails gets a reference to the given X509CertificateSummaryType and assigns it to the CaCertificateDetails field.
func (o *TlsDeliveryProfileCreate) SetCaCertificateDetails(v X509CertificateSummaryType) {
	o.CaCertificateDetails = &v
}

// GetContentProviderId returns the ContentProviderId field value
func (o *TlsDeliveryProfileCreate) GetContentProviderId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ContentProviderId
}

// GetContentProviderIdOk returns a tuple with the ContentProviderId field value
// and a boolean to check if the value has been set.
func (o *TlsDeliveryProfileCreate) GetContentProviderIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContentProviderId, true
}

// SetContentProviderId sets field value
func (o *TlsDeliveryProfileCreate) SetContentProviderId(v int32) {
	o.ContentProviderId = v
}

// GetCaCertificateSecretId returns the CaCertificateSecretId field value if set, zero value otherwise.
func (o *TlsDeliveryProfileCreate) GetCaCertificateSecretId() int32 {
	if o == nil || o.CaCertificateSecretId == nil {
		var ret int32
		return ret
	}
	return *o.CaCertificateSecretId
}

// GetCaCertificateSecretIdOk returns a tuple with the CaCertificateSecretId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsDeliveryProfileCreate) GetCaCertificateSecretIdOk() (*int32, bool) {
	if o == nil || o.CaCertificateSecretId == nil {
		return nil, false
	}
	return o.CaCertificateSecretId, true
}

// HasCaCertificateSecretId returns a boolean if a field has been set.
func (o *TlsDeliveryProfileCreate) HasCaCertificateSecretId() bool {
	if o != nil && o.CaCertificateSecretId != nil {
		return true
	}

	return false
}

// SetCaCertificateSecretId gets a reference to the given int32 and assigns it to the CaCertificateSecretId field.
func (o *TlsDeliveryProfileCreate) SetCaCertificateSecretId(v int32) {
	o.CaCertificateSecretId = &v
}

// GetTlsVerificationDepth returns the TlsVerificationDepth field value if set, zero value otherwise.
func (o *TlsDeliveryProfileCreate) GetTlsVerificationDepth() int32 {
	if o == nil || o.TlsVerificationDepth == nil {
		var ret int32
		return ret
	}
	return *o.TlsVerificationDepth
}

// GetTlsVerificationDepthOk returns a tuple with the TlsVerificationDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsDeliveryProfileCreate) GetTlsVerificationDepthOk() (*int32, bool) {
	if o == nil || o.TlsVerificationDepth == nil {
		return nil, false
	}
	return o.TlsVerificationDepth, true
}

// HasTlsVerificationDepth returns a boolean if a field has been set.
func (o *TlsDeliveryProfileCreate) HasTlsVerificationDepth() bool {
	if o != nil && o.TlsVerificationDepth != nil {
		return true
	}

	return false
}

// SetTlsVerificationDepth gets a reference to the given int32 and assigns it to the TlsVerificationDepth field.
func (o *TlsDeliveryProfileCreate) SetTlsVerificationDepth(v int32) {
	o.TlsVerificationDepth = &v
}

// GetTlsBufferSize returns the TlsBufferSize field value if set, zero value otherwise.
func (o *TlsDeliveryProfileCreate) GetTlsBufferSize() TlsBufferSizeType {
	if o == nil || o.TlsBufferSize == nil {
		var ret TlsBufferSizeType
		return ret
	}
	return *o.TlsBufferSize
}

// GetTlsBufferSizeOk returns a tuple with the TlsBufferSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsDeliveryProfileCreate) GetTlsBufferSizeOk() (*TlsBufferSizeType, bool) {
	if o == nil || o.TlsBufferSize == nil {
		return nil, false
	}
	return o.TlsBufferSize, true
}

// HasTlsBufferSize returns a boolean if a field has been set.
func (o *TlsDeliveryProfileCreate) HasTlsBufferSize() bool {
	if o != nil && o.TlsBufferSize != nil {
		return true
	}

	return false
}

// SetTlsBufferSize gets a reference to the given TlsBufferSizeType and assigns it to the TlsBufferSize field.
func (o *TlsDeliveryProfileCreate) SetTlsBufferSize(v TlsBufferSizeType) {
	o.TlsBufferSize = &v
}

// GetTlsCertificateKey returns the TlsCertificateKey field value if set, zero value otherwise.
func (o *TlsDeliveryProfileCreate) GetTlsCertificateKey() string {
	if o == nil || o.TlsCertificateKey == nil {
		var ret string
		return ret
	}
	return *o.TlsCertificateKey
}

// GetTlsCertificateKeyOk returns a tuple with the TlsCertificateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsDeliveryProfileCreate) GetTlsCertificateKeyOk() (*string, bool) {
	if o == nil || o.TlsCertificateKey == nil {
		return nil, false
	}
	return o.TlsCertificateKey, true
}

// HasTlsCertificateKey returns a boolean if a field has been set.
func (o *TlsDeliveryProfileCreate) HasTlsCertificateKey() bool {
	if o != nil && o.TlsCertificateKey != nil {
		return true
	}

	return false
}

// SetTlsCertificateKey gets a reference to the given string and assigns it to the TlsCertificateKey field.
func (o *TlsDeliveryProfileCreate) SetTlsCertificateKey(v string) {
	o.TlsCertificateKey = &v
}

// GetEcdhCurve returns the EcdhCurve field value if set, zero value otherwise.
func (o *TlsDeliveryProfileCreate) GetEcdhCurve() EcdhCurveType {
	if o == nil || o.EcdhCurve == nil {
		var ret EcdhCurveType
		return ret
	}
	return *o.EcdhCurve
}

// GetEcdhCurveOk returns a tuple with the EcdhCurve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsDeliveryProfileCreate) GetEcdhCurveOk() (*EcdhCurveType, bool) {
	if o == nil || o.EcdhCurve == nil {
		return nil, false
	}
	return o.EcdhCurve, true
}

// HasEcdhCurve returns a boolean if a field has been set.
func (o *TlsDeliveryProfileCreate) HasEcdhCurve() bool {
	if o != nil && o.EcdhCurve != nil {
		return true
	}

	return false
}

// SetEcdhCurve gets a reference to the given EcdhCurveType and assigns it to the EcdhCurve field.
func (o *TlsDeliveryProfileCreate) SetEcdhCurve(v EcdhCurveType) {
	o.EcdhCurve = &v
}

// GetTlsCertificateDetails returns the TlsCertificateDetails field value if set, zero value otherwise.
func (o *TlsDeliveryProfileCreate) GetTlsCertificateDetails() X509CertificateSummaryType {
	if o == nil || o.TlsCertificateDetails == nil {
		var ret X509CertificateSummaryType
		return ret
	}
	return *o.TlsCertificateDetails
}

// GetTlsCertificateDetailsOk returns a tuple with the TlsCertificateDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsDeliveryProfileCreate) GetTlsCertificateDetailsOk() (*X509CertificateSummaryType, bool) {
	if o == nil || o.TlsCertificateDetails == nil {
		return nil, false
	}
	return o.TlsCertificateDetails, true
}

// HasTlsCertificateDetails returns a boolean if a field has been set.
func (o *TlsDeliveryProfileCreate) HasTlsCertificateDetails() bool {
	if o != nil && o.TlsCertificateDetails != nil {
		return true
	}

	return false
}

// SetTlsCertificateDetails gets a reference to the given X509CertificateSummaryType and assigns it to the TlsCertificateDetails field.
func (o *TlsDeliveryProfileCreate) SetTlsCertificateDetails(v X509CertificateSummaryType) {
	o.TlsCertificateDetails = &v
}

// GetTlsProtocols returns the TlsProtocols field value if set, zero value otherwise.
func (o *TlsDeliveryProfileCreate) GetTlsProtocols() []string {
	if o == nil || o.TlsProtocols == nil {
		var ret []string
		return ret
	}
	return *o.TlsProtocols
}

// GetTlsProtocolsOk returns a tuple with the TlsProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsDeliveryProfileCreate) GetTlsProtocolsOk() (*[]string, bool) {
	if o == nil || o.TlsProtocols == nil {
		return nil, false
	}
	return o.TlsProtocols, true
}

// HasTlsProtocols returns a boolean if a field has been set.
func (o *TlsDeliveryProfileCreate) HasTlsProtocols() bool {
	if o != nil && o.TlsProtocols != nil {
		return true
	}

	return false
}

// SetTlsProtocols gets a reference to the given []string and assigns it to the TlsProtocols field.
func (o *TlsDeliveryProfileCreate) SetTlsProtocols(v []string) {
	o.TlsProtocols = &v
}

// GetCipherSuites returns the CipherSuites field value if set, zero value otherwise.
func (o *TlsDeliveryProfileCreate) GetCipherSuites() string {
	if o == nil || o.CipherSuites == nil {
		var ret string
		return ret
	}
	return *o.CipherSuites
}

// GetCipherSuitesOk returns a tuple with the CipherSuites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsDeliveryProfileCreate) GetCipherSuitesOk() (*string, bool) {
	if o == nil || o.CipherSuites == nil {
		return nil, false
	}
	return o.CipherSuites, true
}

// HasCipherSuites returns a boolean if a field has been set.
func (o *TlsDeliveryProfileCreate) HasCipherSuites() bool {
	if o != nil && o.CipherSuites != nil {
		return true
	}

	return false
}

// SetCipherSuites gets a reference to the given string and assigns it to the CipherSuites field.
func (o *TlsDeliveryProfileCreate) SetCipherSuites(v string) {
	o.CipherSuites = &v
}

// GetEnableClientAuth returns the EnableClientAuth field value if set, zero value otherwise.
func (o *TlsDeliveryProfileCreate) GetEnableClientAuth() bool {
	if o == nil || o.EnableClientAuth == nil {
		var ret bool
		return ret
	}
	return *o.EnableClientAuth
}

// GetEnableClientAuthOk returns a tuple with the EnableClientAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsDeliveryProfileCreate) GetEnableClientAuthOk() (*bool, bool) {
	if o == nil || o.EnableClientAuth == nil {
		return nil, false
	}
	return o.EnableClientAuth, true
}

// HasEnableClientAuth returns a boolean if a field has been set.
func (o *TlsDeliveryProfileCreate) HasEnableClientAuth() bool {
	if o != nil && o.EnableClientAuth != nil {
		return true
	}

	return false
}

// SetEnableClientAuth gets a reference to the given bool and assigns it to the EnableClientAuth field.
func (o *TlsDeliveryProfileCreate) SetEnableClientAuth(v bool) {
	o.EnableClientAuth = &v
}

// GetCrl returns the Crl field value if set, zero value otherwise.
func (o *TlsDeliveryProfileCreate) GetCrl() string {
	if o == nil || o.Crl == nil {
		var ret string
		return ret
	}
	return *o.Crl
}

// GetCrlOk returns a tuple with the Crl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsDeliveryProfileCreate) GetCrlOk() (*string, bool) {
	if o == nil || o.Crl == nil {
		return nil, false
	}
	return o.Crl, true
}

// HasCrl returns a boolean if a field has been set.
func (o *TlsDeliveryProfileCreate) HasCrl() bool {
	if o != nil && o.Crl != nil {
		return true
	}

	return false
}

// SetCrl gets a reference to the given string and assigns it to the Crl field.
func (o *TlsDeliveryProfileCreate) SetCrl(v string) {
	o.Crl = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TlsDeliveryProfileCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsDeliveryProfileCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TlsDeliveryProfileCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TlsDeliveryProfileCreate) SetDescription(v string) {
	o.Description = &v
}

func (o TlsDeliveryProfileCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TlsCertificate != nil {
		toSerialize["tlsCertificate"] = o.TlsCertificate
	}
	if o.TlsCertificateAndKeySecretId != nil {
		toSerialize["tlsCertificateAndKeySecretId"] = o.TlsCertificateAndKeySecretId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Ciphers != nil {
		toSerialize["ciphers"] = o.Ciphers
	}
	if o.CaCertificate != nil {
		toSerialize["caCertificate"] = o.CaCertificate
	}
	if o.CaCertificateDetails != nil {
		toSerialize["caCertificateDetails"] = o.CaCertificateDetails
	}
	if true {
		toSerialize["contentProviderId"] = o.ContentProviderId
	}
	if o.CaCertificateSecretId != nil {
		toSerialize["caCertificateSecretId"] = o.CaCertificateSecretId
	}
	if o.TlsVerificationDepth != nil {
		toSerialize["tlsVerificationDepth"] = o.TlsVerificationDepth
	}
	if o.TlsBufferSize != nil {
		toSerialize["tlsBufferSize"] = o.TlsBufferSize
	}
	if o.TlsCertificateKey != nil {
		toSerialize["tlsCertificateKey"] = o.TlsCertificateKey
	}
	if o.EcdhCurve != nil {
		toSerialize["ecdhCurve"] = o.EcdhCurve
	}
	if o.TlsCertificateDetails != nil {
		toSerialize["tlsCertificateDetails"] = o.TlsCertificateDetails
	}
	if o.TlsProtocols != nil {
		toSerialize["tlsProtocols"] = o.TlsProtocols
	}
	if o.CipherSuites != nil {
		toSerialize["cipherSuites"] = o.CipherSuites
	}
	if o.EnableClientAuth != nil {
		toSerialize["enableClientAuth"] = o.EnableClientAuth
	}
	if o.Crl != nil {
		toSerialize["crl"] = o.Crl
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableTlsDeliveryProfileCreate struct {
	value *TlsDeliveryProfileCreate
	isSet bool
}

func (v NullableTlsDeliveryProfileCreate) Get() *TlsDeliveryProfileCreate {
	return v.value
}

func (v *NullableTlsDeliveryProfileCreate) Set(val *TlsDeliveryProfileCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableTlsDeliveryProfileCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableTlsDeliveryProfileCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTlsDeliveryProfileCreate(val *TlsDeliveryProfileCreate) *NullableTlsDeliveryProfileCreate {
	return &NullableTlsDeliveryProfileCreate{value: val, isSet: true}
}

func (v NullableTlsDeliveryProfileCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTlsDeliveryProfileCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


