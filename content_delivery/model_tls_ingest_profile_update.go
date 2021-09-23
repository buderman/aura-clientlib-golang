/*
content-delivery

Cotent Delivery API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package content_delivery

import (
	"encoding/json"
)

// TlsIngestProfileUpdate TLS ingest profile object schema definition.
type TlsIngestProfileUpdate struct {
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

// NewTlsIngestProfileUpdate instantiates a new TlsIngestProfileUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsIngestProfileUpdate(tlsIngestProfileId int32, contentProviderId int32, name string) *TlsIngestProfileUpdate {
	this := TlsIngestProfileUpdate{}
	this.TlsIngestProfileId = tlsIngestProfileId
	this.ContentProviderId = contentProviderId
	var enableClientAuth bool = false
	this.EnableClientAuth = &enableClientAuth
	this.Name = name
	return &this
}

// NewTlsIngestProfileUpdateWithDefaults instantiates a new TlsIngestProfileUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsIngestProfileUpdateWithDefaults() *TlsIngestProfileUpdate {
	this := TlsIngestProfileUpdate{}
	var enableClientAuth bool = false
	this.EnableClientAuth = &enableClientAuth
	return &this
}

// GetClientTlsCertificate returns the ClientTlsCertificate field value if set, zero value otherwise.
func (o *TlsIngestProfileUpdate) GetClientTlsCertificate() string {
	if o == nil || o.ClientTlsCertificate == nil {
		var ret string
		return ret
	}
	return *o.ClientTlsCertificate
}

// GetClientTlsCertificateOk returns a tuple with the ClientTlsCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileUpdate) GetClientTlsCertificateOk() (*string, bool) {
	if o == nil || o.ClientTlsCertificate == nil {
		return nil, false
	}
	return o.ClientTlsCertificate, true
}

// HasClientTlsCertificate returns a boolean if a field has been set.
func (o *TlsIngestProfileUpdate) HasClientTlsCertificate() bool {
	if o != nil && o.ClientTlsCertificate != nil {
		return true
	}

	return false
}

// SetClientTlsCertificate gets a reference to the given string and assigns it to the ClientTlsCertificate field.
func (o *TlsIngestProfileUpdate) SetClientTlsCertificate(v string) {
	o.ClientTlsCertificate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TlsIngestProfileUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TlsIngestProfileUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TlsIngestProfileUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetTlsIngestProfileId returns the TlsIngestProfileId field value
func (o *TlsIngestProfileUpdate) GetTlsIngestProfileId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TlsIngestProfileId
}

// GetTlsIngestProfileIdOk returns a tuple with the TlsIngestProfileId field value
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileUpdate) GetTlsIngestProfileIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TlsIngestProfileId, true
}

// SetTlsIngestProfileId sets field value
func (o *TlsIngestProfileUpdate) SetTlsIngestProfileId(v int32) {
	o.TlsIngestProfileId = v
}

// GetCaCertificate returns the CaCertificate field value if set, zero value otherwise.
func (o *TlsIngestProfileUpdate) GetCaCertificate() string {
	if o == nil || o.CaCertificate == nil {
		var ret string
		return ret
	}
	return *o.CaCertificate
}

// GetCaCertificateOk returns a tuple with the CaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileUpdate) GetCaCertificateOk() (*string, bool) {
	if o == nil || o.CaCertificate == nil {
		return nil, false
	}
	return o.CaCertificate, true
}

// HasCaCertificate returns a boolean if a field has been set.
func (o *TlsIngestProfileUpdate) HasCaCertificate() bool {
	if o != nil && o.CaCertificate != nil {
		return true
	}

	return false
}

// SetCaCertificate gets a reference to the given string and assigns it to the CaCertificate field.
func (o *TlsIngestProfileUpdate) SetCaCertificate(v string) {
	o.CaCertificate = &v
}

// GetCaCertificateDetails returns the CaCertificateDetails field value if set, zero value otherwise.
func (o *TlsIngestProfileUpdate) GetCaCertificateDetails() X509CertificateSummaryType {
	if o == nil || o.CaCertificateDetails == nil {
		var ret X509CertificateSummaryType
		return ret
	}
	return *o.CaCertificateDetails
}

// GetCaCertificateDetailsOk returns a tuple with the CaCertificateDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileUpdate) GetCaCertificateDetailsOk() (*X509CertificateSummaryType, bool) {
	if o == nil || o.CaCertificateDetails == nil {
		return nil, false
	}
	return o.CaCertificateDetails, true
}

// HasCaCertificateDetails returns a boolean if a field has been set.
func (o *TlsIngestProfileUpdate) HasCaCertificateDetails() bool {
	if o != nil && o.CaCertificateDetails != nil {
		return true
	}

	return false
}

// SetCaCertificateDetails gets a reference to the given X509CertificateSummaryType and assigns it to the CaCertificateDetails field.
func (o *TlsIngestProfileUpdate) SetCaCertificateDetails(v X509CertificateSummaryType) {
	o.CaCertificateDetails = &v
}

// GetContentProviderId returns the ContentProviderId field value
func (o *TlsIngestProfileUpdate) GetContentProviderId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ContentProviderId
}

// GetContentProviderIdOk returns a tuple with the ContentProviderId field value
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileUpdate) GetContentProviderIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContentProviderId, true
}

// SetContentProviderId sets field value
func (o *TlsIngestProfileUpdate) SetContentProviderId(v int32) {
	o.ContentProviderId = v
}

// GetCiphers returns the Ciphers field value if set, zero value otherwise.
func (o *TlsIngestProfileUpdate) GetCiphers() string {
	if o == nil || o.Ciphers == nil {
		var ret string
		return ret
	}
	return *o.Ciphers
}

// GetCiphersOk returns a tuple with the Ciphers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileUpdate) GetCiphersOk() (*string, bool) {
	if o == nil || o.Ciphers == nil {
		return nil, false
	}
	return o.Ciphers, true
}

// HasCiphers returns a boolean if a field has been set.
func (o *TlsIngestProfileUpdate) HasCiphers() bool {
	if o != nil && o.Ciphers != nil {
		return true
	}

	return false
}

// SetCiphers gets a reference to the given string and assigns it to the Ciphers field.
func (o *TlsIngestProfileUpdate) SetCiphers(v string) {
	o.Ciphers = &v
}

// GetCrlDetails returns the CrlDetails field value if set, zero value otherwise.
func (o *TlsIngestProfileUpdate) GetCrlDetails() X509CrlSummaryType {
	if o == nil || o.CrlDetails == nil {
		var ret X509CrlSummaryType
		return ret
	}
	return *o.CrlDetails
}

// GetCrlDetailsOk returns a tuple with the CrlDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileUpdate) GetCrlDetailsOk() (*X509CrlSummaryType, bool) {
	if o == nil || o.CrlDetails == nil {
		return nil, false
	}
	return o.CrlDetails, true
}

// HasCrlDetails returns a boolean if a field has been set.
func (o *TlsIngestProfileUpdate) HasCrlDetails() bool {
	if o != nil && o.CrlDetails != nil {
		return true
	}

	return false
}

// SetCrlDetails gets a reference to the given X509CrlSummaryType and assigns it to the CrlDetails field.
func (o *TlsIngestProfileUpdate) SetCrlDetails(v X509CrlSummaryType) {
	o.CrlDetails = &v
}

// GetTlsProtocols returns the TlsProtocols field value if set, zero value otherwise.
func (o *TlsIngestProfileUpdate) GetTlsProtocols() []string {
	if o == nil || o.TlsProtocols == nil {
		var ret []string
		return ret
	}
	return *o.TlsProtocols
}

// GetTlsProtocolsOk returns a tuple with the TlsProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileUpdate) GetTlsProtocolsOk() (*[]string, bool) {
	if o == nil || o.TlsProtocols == nil {
		return nil, false
	}
	return o.TlsProtocols, true
}

// HasTlsProtocols returns a boolean if a field has been set.
func (o *TlsIngestProfileUpdate) HasTlsProtocols() bool {
	if o != nil && o.TlsProtocols != nil {
		return true
	}

	return false
}

// SetTlsProtocols gets a reference to the given []string and assigns it to the TlsProtocols field.
func (o *TlsIngestProfileUpdate) SetTlsProtocols(v []string) {
	o.TlsProtocols = &v
}

// GetCipherSuites returns the CipherSuites field value if set, zero value otherwise.
func (o *TlsIngestProfileUpdate) GetCipherSuites() string {
	if o == nil || o.CipherSuites == nil {
		var ret string
		return ret
	}
	return *o.CipherSuites
}

// GetCipherSuitesOk returns a tuple with the CipherSuites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileUpdate) GetCipherSuitesOk() (*string, bool) {
	if o == nil || o.CipherSuites == nil {
		return nil, false
	}
	return o.CipherSuites, true
}

// HasCipherSuites returns a boolean if a field has been set.
func (o *TlsIngestProfileUpdate) HasCipherSuites() bool {
	if o != nil && o.CipherSuites != nil {
		return true
	}

	return false
}

// SetCipherSuites gets a reference to the given string and assigns it to the CipherSuites field.
func (o *TlsIngestProfileUpdate) SetCipherSuites(v string) {
	o.CipherSuites = &v
}

// GetClientTlsCertificateKey returns the ClientTlsCertificateKey field value if set, zero value otherwise.
func (o *TlsIngestProfileUpdate) GetClientTlsCertificateKey() string {
	if o == nil || o.ClientTlsCertificateKey == nil {
		var ret string
		return ret
	}
	return *o.ClientTlsCertificateKey
}

// GetClientTlsCertificateKeyOk returns a tuple with the ClientTlsCertificateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileUpdate) GetClientTlsCertificateKeyOk() (*string, bool) {
	if o == nil || o.ClientTlsCertificateKey == nil {
		return nil, false
	}
	return o.ClientTlsCertificateKey, true
}

// HasClientTlsCertificateKey returns a boolean if a field has been set.
func (o *TlsIngestProfileUpdate) HasClientTlsCertificateKey() bool {
	if o != nil && o.ClientTlsCertificateKey != nil {
		return true
	}

	return false
}

// SetClientTlsCertificateKey gets a reference to the given string and assigns it to the ClientTlsCertificateKey field.
func (o *TlsIngestProfileUpdate) SetClientTlsCertificateKey(v string) {
	o.ClientTlsCertificateKey = &v
}

// GetClientTlsCertificateDetails returns the ClientTlsCertificateDetails field value if set, zero value otherwise.
func (o *TlsIngestProfileUpdate) GetClientTlsCertificateDetails() X509CertificateSummaryType {
	if o == nil || o.ClientTlsCertificateDetails == nil {
		var ret X509CertificateSummaryType
		return ret
	}
	return *o.ClientTlsCertificateDetails
}

// GetClientTlsCertificateDetailsOk returns a tuple with the ClientTlsCertificateDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileUpdate) GetClientTlsCertificateDetailsOk() (*X509CertificateSummaryType, bool) {
	if o == nil || o.ClientTlsCertificateDetails == nil {
		return nil, false
	}
	return o.ClientTlsCertificateDetails, true
}

// HasClientTlsCertificateDetails returns a boolean if a field has been set.
func (o *TlsIngestProfileUpdate) HasClientTlsCertificateDetails() bool {
	if o != nil && o.ClientTlsCertificateDetails != nil {
		return true
	}

	return false
}

// SetClientTlsCertificateDetails gets a reference to the given X509CertificateSummaryType and assigns it to the ClientTlsCertificateDetails field.
func (o *TlsIngestProfileUpdate) SetClientTlsCertificateDetails(v X509CertificateSummaryType) {
	o.ClientTlsCertificateDetails = &v
}

// GetEnableClientAuth returns the EnableClientAuth field value if set, zero value otherwise.
func (o *TlsIngestProfileUpdate) GetEnableClientAuth() bool {
	if o == nil || o.EnableClientAuth == nil {
		var ret bool
		return ret
	}
	return *o.EnableClientAuth
}

// GetEnableClientAuthOk returns a tuple with the EnableClientAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileUpdate) GetEnableClientAuthOk() (*bool, bool) {
	if o == nil || o.EnableClientAuth == nil {
		return nil, false
	}
	return o.EnableClientAuth, true
}

// HasEnableClientAuth returns a boolean if a field has been set.
func (o *TlsIngestProfileUpdate) HasEnableClientAuth() bool {
	if o != nil && o.EnableClientAuth != nil {
		return true
	}

	return false
}

// SetEnableClientAuth gets a reference to the given bool and assigns it to the EnableClientAuth field.
func (o *TlsIngestProfileUpdate) SetEnableClientAuth(v bool) {
	o.EnableClientAuth = &v
}

// GetCrl returns the Crl field value if set, zero value otherwise.
func (o *TlsIngestProfileUpdate) GetCrl() string {
	if o == nil || o.Crl == nil {
		var ret string
		return ret
	}
	return *o.Crl
}

// GetCrlOk returns a tuple with the Crl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileUpdate) GetCrlOk() (*string, bool) {
	if o == nil || o.Crl == nil {
		return nil, false
	}
	return o.Crl, true
}

// HasCrl returns a boolean if a field has been set.
func (o *TlsIngestProfileUpdate) HasCrl() bool {
	if o != nil && o.Crl != nil {
		return true
	}

	return false
}

// SetCrl gets a reference to the given string and assigns it to the Crl field.
func (o *TlsIngestProfileUpdate) SetCrl(v string) {
	o.Crl = &v
}

// GetName returns the Name field value
func (o *TlsIngestProfileUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TlsIngestProfileUpdate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TlsIngestProfileUpdate) SetName(v string) {
	o.Name = v
}

func (o TlsIngestProfileUpdate) MarshalJSON() ([]byte, error) {
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

type NullableTlsIngestProfileUpdate struct {
	value *TlsIngestProfileUpdate
	isSet bool
}

func (v NullableTlsIngestProfileUpdate) Get() *TlsIngestProfileUpdate {
	return v.value
}

func (v *NullableTlsIngestProfileUpdate) Set(val *TlsIngestProfileUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableTlsIngestProfileUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableTlsIngestProfileUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTlsIngestProfileUpdate(val *TlsIngestProfileUpdate) *NullableTlsIngestProfileUpdate {
	return &NullableTlsIngestProfileUpdate{value: val, isSet: true}
}

func (v NullableTlsIngestProfileUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTlsIngestProfileUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

