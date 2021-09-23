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

// X509CrlSummaryTypeIssuedBy struct for X509CrlSummaryTypeIssuedBy
type X509CrlSummaryTypeIssuedBy struct {
	Organization string `json:"organization"`
	CommonName string `json:"commonName"`
	OrganizationalUnit string `json:"organizationalUnit"`
}

// NewX509CrlSummaryTypeIssuedBy instantiates a new X509CrlSummaryTypeIssuedBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewX509CrlSummaryTypeIssuedBy(organization string, commonName string, organizationalUnit string) *X509CrlSummaryTypeIssuedBy {
	this := X509CrlSummaryTypeIssuedBy{}
	this.Organization = organization
	this.CommonName = commonName
	this.OrganizationalUnit = organizationalUnit
	return &this
}

// NewX509CrlSummaryTypeIssuedByWithDefaults instantiates a new X509CrlSummaryTypeIssuedBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewX509CrlSummaryTypeIssuedByWithDefaults() *X509CrlSummaryTypeIssuedBy {
	this := X509CrlSummaryTypeIssuedBy{}
	return &this
}

// GetOrganization returns the Organization field value
func (o *X509CrlSummaryTypeIssuedBy) GetOrganization() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *X509CrlSummaryTypeIssuedBy) GetOrganizationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *X509CrlSummaryTypeIssuedBy) SetOrganization(v string) {
	o.Organization = v
}

// GetCommonName returns the CommonName field value
func (o *X509CrlSummaryTypeIssuedBy) GetCommonName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value
// and a boolean to check if the value has been set.
func (o *X509CrlSummaryTypeIssuedBy) GetCommonNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CommonName, true
}

// SetCommonName sets field value
func (o *X509CrlSummaryTypeIssuedBy) SetCommonName(v string) {
	o.CommonName = v
}

// GetOrganizationalUnit returns the OrganizationalUnit field value
func (o *X509CrlSummaryTypeIssuedBy) GetOrganizationalUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationalUnit
}

// GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field value
// and a boolean to check if the value has been set.
func (o *X509CrlSummaryTypeIssuedBy) GetOrganizationalUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OrganizationalUnit, true
}

// SetOrganizationalUnit sets field value
func (o *X509CrlSummaryTypeIssuedBy) SetOrganizationalUnit(v string) {
	o.OrganizationalUnit = v
}

func (o X509CrlSummaryTypeIssuedBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["organization"] = o.Organization
	}
	if true {
		toSerialize["commonName"] = o.CommonName
	}
	if true {
		toSerialize["organizationalUnit"] = o.OrganizationalUnit
	}
	return json.Marshal(toSerialize)
}

type NullableX509CrlSummaryTypeIssuedBy struct {
	value *X509CrlSummaryTypeIssuedBy
	isSet bool
}

func (v NullableX509CrlSummaryTypeIssuedBy) Get() *X509CrlSummaryTypeIssuedBy {
	return v.value
}

func (v *NullableX509CrlSummaryTypeIssuedBy) Set(val *X509CrlSummaryTypeIssuedBy) {
	v.value = val
	v.isSet = true
}

func (v NullableX509CrlSummaryTypeIssuedBy) IsSet() bool {
	return v.isSet
}

func (v *NullableX509CrlSummaryTypeIssuedBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableX509CrlSummaryTypeIssuedBy(val *X509CrlSummaryTypeIssuedBy) *NullableX509CrlSummaryTypeIssuedBy {
	return &NullableX509CrlSummaryTypeIssuedBy{value: val, isSet: true}
}

func (v NullableX509CrlSummaryTypeIssuedBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableX509CrlSummaryTypeIssuedBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


