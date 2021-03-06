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

// TlsIngestProfileReadDetailed TLS ingest profile object schema definition.
type TlsIngestProfileReadDetailed struct {
	// An OpenSSL PEM-formatted certificate associated with the trusted CA that signed the origin's certificate.
	ClientTlsCertificate *string `json:"clientTlsCertificate,omitempty"`
	Description *string `json:"description,omitempty"`
	TlsIngestProfileId int32 `json:"tlsIngestProfileId"`
	// An OpenSSL PEM-formatted certificate associated with the trusted CA that signed the origin's certificate.
	CaCertificate *string `json:"caCertificate,omitempty"`
	CaCertificateDetails *X509CertificateSummaryType `json:"caCertificateDetails,omitempty"`
	ContentProviderId int32 `json:"contentProviderId"`
	// An OpenSSL style [ciphers string](https://www.openssl.org/docs/apps/ciphers.html#CIPHER-STRINGS).
	Ciphers *string `json:"ciphers,omitempty"`
	CrlDetails *X509CrlSummaryType `json:"crlDetails,omitempty"`
	// TLS protocols supported, either `TLSV1`, `TLSV1_1`, or `TLSV1_2`. Defaults to `TLSV1_2`.  The use of `SSLV2` or `SSLV3` are discouraged for security reasons.
	TlsProtocols *[]string `json:"tlsProtocols,omitempty"`
	// An OpenSSL style [ciphers string](https://www.openssl.org/docs/apps/ciphers.html#CIPHER-STRINGS).
	CipherSuites *string `json:"cipherSuites,omitempty"`
	// An OpenSSL PEM-formatted certificate associated with the trusted CA that signed the origin's certificate.
	ClientTlsCertificateKey *string `json:"clientTlsCertificateKey,omitempty"`
	ClientTlsCertificateDetails *X509CertificateSummaryType `json:"clientTlsCertificateDetails,omitempty"`
	// Allows you to enable or disable client authorization. Enabling this feature causes HyperCache to authenticate client certificates using the CA certificate for each TLS connection.
	EnableClientAuth *bool `json:"enableClientAuth,omitempty"`
	// Certificate Revocation List. An OpenSSL PEM-formatted list of revoked client certificates.
	Crl *string `json:"crl,omitempty"`
	Name string `json:"name"`
}

// NewTlsIngestProfileReadDetailed instantiates a new TlsIngestProfileReadDetailed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsIngestProfileReadDetailed(tlsIngestProfileId int32, contentProviderId int32, name string) *TlsIngestProfileReadDetailed {
	this := TlsIngestProfileReadDetailed{}
	this.TlsIngestProfileId = tlsIngestProfileId
	this.ContentProviderId = contentProviderId
	var enableClientAuth bool = false
	this.EnableClientAuth = &enableClientAuth
	this.Name = name
	return &this
}

// NewTlsIngestProfileReadDetailedWithDefaults instantiates a new TlsIngestProfileReadDetailed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsIngestProfileReadDetailedWithDefaults() *TlsIngestProfileReadDetailed {
	this := TlsIngestProfileReadDetailed{}
	var enableClientAuth bool = false
	this.EnableClientAuth = &enableClientAuth
	return &this
}

// GetClientTlsCertificate returns the ClientTlsCertificate field value if set, zero value otherwise.
func (o *TlsIngestProfileReadDetailed) GetClientTlsCertificate() string {
	if o == nil || o.ClientTlsCertificate == nil {
		var ret string
		return ret
	}
	return *o.ClientTlsCertificate
}

// GetClientTlsCertificateOk returns a tuple with the ClientTlsCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileReadDetailed) GetClientTlsCertificateOk() (*string, bool) {
	if o == nil || o.ClientTlsCertificate == nil {
		return nil, false
	}
	return o.ClientTlsCertificate, true
}

// HasClientTlsCertificate returns a boolean if a field has been set.
func (o *TlsIngestProfileReadDetailed) HasClientTlsCertificate() bool {
	if o != nil && o.ClientTlsCertificate != nil {
		return true
	}

	return false
}

// SetClientTlsCertificate gets a reference to the given string and assigns it to the ClientTlsCertificate field.
func (o *TlsIngestProfileReadDetailed) SetClientTlsCertificate(v string) {
	o.ClientTlsCertificate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TlsIngestProfileReadDetailed) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileReadDetailed) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TlsIngestProfileReadDetailed) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TlsIngestProfileReadDetailed) SetDescription(v string) {
	o.Description = &v
}

// GetTlsIngestProfileId returns the TlsIngestProfileId field value
func (o *TlsIngestProfileReadDetailed) GetTlsIngestProfileId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TlsIngestProfileId
}

// GetTlsIngestProfileIdOk returns a tuple with the TlsIngestProfileId field value
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileReadDetailed) GetTlsIngestProfileIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TlsIngestProfileId, true
}

// SetTlsIngestProfileId sets field value
func (o *TlsIngestProfileReadDetailed) SetTlsIngestProfileId(v int32) {
	o.TlsIngestProfileId = v
}

// GetCaCertificate returns the CaCertificate field value if set, zero value otherwise.
func (o *TlsIngestProfileReadDetailed) GetCaCertificate() string {
	if o == nil || o.CaCertificate == nil {
		var ret string
		return ret
	}
	return *o.CaCertificate
}

// GetCaCertificateOk returns a tuple with the CaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileReadDetailed) GetCaCertificateOk() (*string, bool) {
	if o == nil || o.CaCertificate == nil {
		return nil, false
	}
	return o.CaCertificate, true
}

// HasCaCertificate returns a boolean if a field has been set.
func (o *TlsIngestProfileReadDetailed) HasCaCertificate() bool {
	if o != nil && o.CaCertificate != nil {
		return true
	}

	return false
}

// SetCaCertificate gets a reference to the given string and assigns it to the CaCertificate field.
func (o *TlsIngestProfileReadDetailed) SetCaCertificate(v string) {
	o.CaCertificate = &v
}

// GetCaCertificateDetails returns the CaCertificateDetails field value if set, zero value otherwise.
func (o *TlsIngestProfileReadDetailed) GetCaCertificateDetails() X509CertificateSummaryType {
	if o == nil || o.CaCertificateDetails == nil {
		var ret X509CertificateSummaryType
		return ret
	}
	return *o.CaCertificateDetails
}

// GetCaCertificateDetailsOk returns a tuple with the CaCertificateDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileReadDetailed) GetCaCertificateDetailsOk() (*X509CertificateSummaryType, bool) {
	if o == nil || o.CaCertificateDetails == nil {
		return nil, false
	}
	return o.CaCertificateDetails, true
}

// HasCaCertificateDetails returns a boolean if a field has been set.
func (o *TlsIngestProfileReadDetailed) HasCaCertificateDetails() bool {
	if o != nil && o.CaCertificateDetails != nil {
		return true
	}

	return false
}

// SetCaCertificateDetails gets a reference to the given X509CertificateSummaryType and assigns it to the CaCertificateDetails field.
func (o *TlsIngestProfileReadDetailed) SetCaCertificateDetails(v X509CertificateSummaryType) {
	o.CaCertificateDetails = &v
}

// GetContentProviderId returns the ContentProviderId field value
func (o *TlsIngestProfileReadDetailed) GetContentProviderId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ContentProviderId
}

// GetContentProviderIdOk returns a tuple with the ContentProviderId field value
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileReadDetailed) GetContentProviderIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContentProviderId, true
}

// SetContentProviderId sets field value
func (o *TlsIngestProfileReadDetailed) SetContentProviderId(v int32) {
	o.ContentProviderId = v
}

// GetCiphers returns the Ciphers field value if set, zero value otherwise.
func (o *TlsIngestProfileReadDetailed) GetCiphers() string {
	if o == nil || o.Ciphers == nil {
		var ret string
		return ret
	}
	return *o.Ciphers
}

// GetCiphersOk returns a tuple with the Ciphers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileReadDetailed) GetCiphersOk() (*string, bool) {
	if o == nil || o.Ciphers == nil {
		return nil, false
	}
	return o.Ciphers, true
}

// HasCiphers returns a boolean if a field has been set.
func (o *TlsIngestProfileReadDetailed) HasCiphers() bool {
	if o != nil && o.Ciphers != nil {
		return true
	}

	return false
}

// SetCiphers gets a reference to the given string and assigns it to the Ciphers field.
func (o *TlsIngestProfileReadDetailed) SetCiphers(v string) {
	o.Ciphers = &v
}

// GetCrlDetails returns the CrlDetails field value if set, zero value otherwise.
func (o *TlsIngestProfileReadDetailed) GetCrlDetails() X509CrlSummaryType {
	if o == nil || o.CrlDetails == nil {
		var ret X509CrlSummaryType
		return ret
	}
	return *o.CrlDetails
}

// GetCrlDetailsOk returns a tuple with the CrlDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileReadDetailed) GetCrlDetailsOk() (*X509CrlSummaryType, bool) {
	if o == nil || o.CrlDetails == nil {
		return nil, false
	}
	return o.CrlDetails, true
}

// HasCrlDetails returns a boolean if a field has been set.
func (o *TlsIngestProfileReadDetailed) HasCrlDetails() bool {
	if o != nil && o.CrlDetails != nil {
		return true
	}

	return false
}

// SetCrlDetails gets a reference to the given X509CrlSummaryType and assigns it to the CrlDetails field.
func (o *TlsIngestProfileReadDetailed) SetCrlDetails(v X509CrlSummaryType) {
	o.CrlDetails = &v
}

// GetTlsProtocols returns the TlsProtocols field value if set, zero value otherwise.
func (o *TlsIngestProfileReadDetailed) GetTlsProtocols() []string {
	if o == nil || o.TlsProtocols == nil {
		var ret []string
		return ret
	}
	return *o.TlsProtocols
}

// GetTlsProtocolsOk returns a tuple with the TlsProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileReadDetailed) GetTlsProtocolsOk() (*[]string, bool) {
	if o == nil || o.TlsProtocols == nil {
		return nil, false
	}
	return o.TlsProtocols, true
}

// HasTlsProtocols returns a boolean if a field has been set.
func (o *TlsIngestProfileReadDetailed) HasTlsProtocols() bool {
	if o != nil && o.TlsProtocols != nil {
		return true
	}

	return false
}

// SetTlsProtocols gets a reference to the given []string and assigns it to the TlsProtocols field.
func (o *TlsIngestProfileReadDetailed) SetTlsProtocols(v []string) {
	o.TlsProtocols = &v
}

// GetCipherSuites returns the CipherSuites field value if set, zero value otherwise.
func (o *TlsIngestProfileReadDetailed) GetCipherSuites() string {
	if o == nil || o.CipherSuites == nil {
		var ret string
		return ret
	}
	return *o.CipherSuites
}

// GetCipherSuitesOk returns a tuple with the CipherSuites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileReadDetailed) GetCipherSuitesOk() (*string, bool) {
	if o == nil || o.CipherSuites == nil {
		return nil, false
	}
	return o.CipherSuites, true
}

// HasCipherSuites returns a boolean if a field has been set.
func (o *TlsIngestProfileReadDetailed) HasCipherSuites() bool {
	if o != nil && o.CipherSuites != nil {
		return true
	}

	return false
}

// SetCipherSuites gets a reference to the given string and assigns it to the CipherSuites field.
func (o *TlsIngestProfileReadDetailed) SetCipherSuites(v string) {
	o.CipherSuites = &v
}

// GetClientTlsCertificateKey returns the ClientTlsCertificateKey field value if set, zero value otherwise.
func (o *TlsIngestProfileReadDetailed) GetClientTlsCertificateKey() string {
	if o == nil || o.ClientTlsCertificateKey == nil {
		var ret string
		return ret
	}
	return *o.ClientTlsCertificateKey
}

// GetClientTlsCertificateKeyOk returns a tuple with the ClientTlsCertificateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileReadDetailed) GetClientTlsCertificateKeyOk() (*string, bool) {
	if o == nil || o.ClientTlsCertificateKey == nil {
		return nil, false
	}
	return o.ClientTlsCertificateKey, true
}

// HasClientTlsCertificateKey returns a boolean if a field has been set.
func (o *TlsIngestProfileReadDetailed) HasClientTlsCertificateKey() bool {
	if o != nil && o.ClientTlsCertificateKey != nil {
		return true
	}

	return false
}

// SetClientTlsCertificateKey gets a reference to the given string and assigns it to the ClientTlsCertificateKey field.
func (o *TlsIngestProfileReadDetailed) SetClientTlsCertificateKey(v string) {
	o.ClientTlsCertificateKey = &v
}

// GetClientTlsCertificateDetails returns the ClientTlsCertificateDetails field value if set, zero value otherwise.
func (o *TlsIngestProfileReadDetailed) GetClientTlsCertificateDetails() X509CertificateSummaryType {
	if o == nil || o.ClientTlsCertificateDetails == nil {
		var ret X509CertificateSummaryType
		return ret
	}
	return *o.ClientTlsCertificateDetails
}

// GetClientTlsCertificateDetailsOk returns a tuple with the ClientTlsCertificateDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileReadDetailed) GetClientTlsCertificateDetailsOk() (*X509CertificateSummaryType, bool) {
	if o == nil || o.ClientTlsCertificateDetails == nil {
		return nil, false
	}
	return o.ClientTlsCertificateDetails, true
}

// HasClientTlsCertificateDetails returns a boolean if a field has been set.
func (o *TlsIngestProfileReadDetailed) HasClientTlsCertificateDetails() bool {
	if o != nil && o.ClientTlsCertificateDetails != nil {
		return true
	}

	return false
}

// SetClientTlsCertificateDetails gets a reference to the given X509CertificateSummaryType and assigns it to the ClientTlsCertificateDetails field.
func (o *TlsIngestProfileReadDetailed) SetClientTlsCertificateDetails(v X509CertificateSummaryType) {
	o.ClientTlsCertificateDetails = &v
}

// GetEnableClientAuth returns the EnableClientAuth field value if set, zero value otherwise.
func (o *TlsIngestProfileReadDetailed) GetEnableClientAuth() bool {
	if o == nil || o.EnableClientAuth == nil {
		var ret bool
		return ret
	}
	return *o.EnableClientAuth
}

// GetEnableClientAuthOk returns a tuple with the EnableClientAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileReadDetailed) GetEnableClientAuthOk() (*bool, bool) {
	if o == nil || o.EnableClientAuth == nil {
		return nil, false
	}
	return o.EnableClientAuth, true
}

// HasEnableClientAuth returns a boolean if a field has been set.
func (o *TlsIngestProfileReadDetailed) HasEnableClientAuth() bool {
	if o != nil && o.EnableClientAuth != nil {
		return true
	}

	return false
}

// SetEnableClientAuth gets a reference to the given bool and assigns it to the EnableClientAuth field.
func (o *TlsIngestProfileReadDetailed) SetEnableClientAuth(v bool) {
	o.EnableClientAuth = &v
}

// GetCrl returns the Crl field value if set, zero value otherwise.
func (o *TlsIngestProfileReadDetailed) GetCrl() string {
	if o == nil || o.Crl == nil {
		var ret string
		return ret
	}
	return *o.Crl
}

// GetCrlOk returns a tuple with the Crl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileReadDetailed) GetCrlOk() (*string, bool) {
	if o == nil || o.Crl == nil {
		return nil, false
	}
	return o.Crl, true
}

// HasCrl returns a boolean if a field has been set.
func (o *TlsIngestProfileReadDetailed) HasCrl() bool {
	if o != nil && o.Crl != nil {
		return true
	}

	return false
}

// SetCrl gets a reference to the given string and assigns it to the Crl field.
func (o *TlsIngestProfileReadDetailed) SetCrl(v string) {
	o.Crl = &v
}

// GetName returns the Name field value
func (o *TlsIngestProfileReadDetailed) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileReadDetailed) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TlsIngestProfileReadDetailed) SetName(v string) {
	o.Name = v
}

func (o TlsIngestProfileReadDetailed) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientTlsCertificate != nil {
		toSerialize["clientTlsCertificate"] = o.ClientTlsCertificate
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["tlsIngestProfileId"] = o.TlsIngestProfileId
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
	if o.Ciphers != nil {
		toSerialize["ciphers"] = o.Ciphers
	}
	if o.CrlDetails != nil {
		toSerialize["crlDetails"] = o.CrlDetails
	}
	if o.TlsProtocols != nil {
		toSerialize["tlsProtocols"] = o.TlsProtocols
	}
	if o.CipherSuites != nil {
		toSerialize["cipherSuites"] = o.CipherSuites
	}
	if o.ClientTlsCertificateKey != nil {
		toSerialize["clientTlsCertificateKey"] = o.ClientTlsCertificateKey
	}
	if o.ClientTlsCertificateDetails != nil {
		toSerialize["clientTlsCertificateDetails"] = o.ClientTlsCertificateDetails
	}
	if o.EnableClientAuth != nil {
		toSerialize["enableClientAuth"] = o.EnableClientAuth
	}
	if o.Crl != nil {
		toSerialize["crl"] = o.Crl
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableTlsIngestProfileReadDetailed struct {
	value *TlsIngestProfileReadDetailed
	isSet bool
}

func (v NullableTlsIngestProfileReadDetailed) Get() *TlsIngestProfileReadDetailed {
	return v.value
}

func (v *NullableTlsIngestProfileReadDetailed) Set(val *TlsIngestProfileReadDetailed) {
	v.value = val
	v.isSet = true
}

func (v NullableTlsIngestProfileReadDetailed) IsSet() bool {
	return v.isSet
}

func (v *NullableTlsIngestProfileReadDetailed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTlsIngestProfileReadDetailed(val *TlsIngestProfileReadDetailed) *NullableTlsIngestProfileReadDetailed {
	return &NullableTlsIngestProfileReadDetailed{value: val, isSet: true}
}

func (v NullableTlsIngestProfileReadDetailed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTlsIngestProfileReadDetailed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


