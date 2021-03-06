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

// ExternalOriginAuthOptions This behavior's set of configuration options.
type ExternalOriginAuthOptions struct {
	// External origin secret access key.
	SecretKey string `json:"secretKey"`
	// External origin service region, such as `us-east-1` and `eu-central-1`.
	Region string `json:"region"`
	// List of request headers included in authorization signature.
	SignedHeaders *[]string `json:"signedHeaders,omitempty"`
	// External Oorigin access key ID.
	AccessKey string `json:"accessKey"`
	// External origin service. The only available service is `AMAZON_AWS_S3`.
	Service string `json:"service"`
}

// NewExternalOriginAuthOptions instantiates a new ExternalOriginAuthOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalOriginAuthOptions(secretKey string, region string, accessKey string, service string) *ExternalOriginAuthOptions {
	this := ExternalOriginAuthOptions{}
	this.SecretKey = secretKey
	this.Region = region
	this.AccessKey = accessKey
	this.Service = service
	return &this
}

// NewExternalOriginAuthOptionsWithDefaults instantiates a new ExternalOriginAuthOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalOriginAuthOptionsWithDefaults() *ExternalOriginAuthOptions {
	this := ExternalOriginAuthOptions{}
	return &this
}

// GetSecretKey returns the SecretKey field value
func (o *ExternalOriginAuthOptions) GetSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value
// and a boolean to check if the value has been set.
func (o *ExternalOriginAuthOptions) GetSecretKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SecretKey, true
}

// SetSecretKey sets field value
func (o *ExternalOriginAuthOptions) SetSecretKey(v string) {
	o.SecretKey = v
}

// GetRegion returns the Region field value
func (o *ExternalOriginAuthOptions) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *ExternalOriginAuthOptions) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *ExternalOriginAuthOptions) SetRegion(v string) {
	o.Region = v
}

// GetSignedHeaders returns the SignedHeaders field value if set, zero value otherwise.
func (o *ExternalOriginAuthOptions) GetSignedHeaders() []string {
	if o == nil || o.SignedHeaders == nil {
		var ret []string
		return ret
	}
	return *o.SignedHeaders
}

// GetSignedHeadersOk returns a tuple with the SignedHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalOriginAuthOptions) GetSignedHeadersOk() (*[]string, bool) {
	if o == nil || o.SignedHeaders == nil {
		return nil, false
	}
	return o.SignedHeaders, true
}

// HasSignedHeaders returns a boolean if a field has been set.
func (o *ExternalOriginAuthOptions) HasSignedHeaders() bool {
	if o != nil && o.SignedHeaders != nil {
		return true
	}

	return false
}

// SetSignedHeaders gets a reference to the given []string and assigns it to the SignedHeaders field.
func (o *ExternalOriginAuthOptions) SetSignedHeaders(v []string) {
	o.SignedHeaders = &v
}

// GetAccessKey returns the AccessKey field value
func (o *ExternalOriginAuthOptions) GetAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value
// and a boolean to check if the value has been set.
func (o *ExternalOriginAuthOptions) GetAccessKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessKey, true
}

// SetAccessKey sets field value
func (o *ExternalOriginAuthOptions) SetAccessKey(v string) {
	o.AccessKey = v
}

// GetService returns the Service field value
func (o *ExternalOriginAuthOptions) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *ExternalOriginAuthOptions) GetServiceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *ExternalOriginAuthOptions) SetService(v string) {
	o.Service = v
}

func (o ExternalOriginAuthOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["secretKey"] = o.SecretKey
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if o.SignedHeaders != nil {
		toSerialize["signedHeaders"] = o.SignedHeaders
	}
	if true {
		toSerialize["accessKey"] = o.AccessKey
	}
	if true {
		toSerialize["service"] = o.Service
	}
	return json.Marshal(toSerialize)
}

type NullableExternalOriginAuthOptions struct {
	value *ExternalOriginAuthOptions
	isSet bool
}

func (v NullableExternalOriginAuthOptions) Get() *ExternalOriginAuthOptions {
	return v.value
}

func (v *NullableExternalOriginAuthOptions) Set(val *ExternalOriginAuthOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalOriginAuthOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalOriginAuthOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalOriginAuthOptions(val *ExternalOriginAuthOptions) *NullableExternalOriginAuthOptions {
	return &NullableExternalOriginAuthOptions{value: val, isSet: true}
}

func (v NullableExternalOriginAuthOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalOriginAuthOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


