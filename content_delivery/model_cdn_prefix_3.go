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

// CdnPrefix3 CdnPrefix summary read
type CdnPrefix3 struct {
	Enable *bool `json:"enable,omitempty"`
	ContentProviderId *int32 `json:"contentProviderId,omitempty"`
	SiteMapId *int32 `json:"siteMapId,omitempty"`
	PrefixPrioritization *PrefixPrioritizationType `json:"prefixPrioritization,omitempty"`
	// Sets the maximum number of requests that can be served through one keep-alive connection. After the maximum number of requests are made, the connection is closed.
	KeepaliveRequests *int32 `json:"keepaliveRequests,omitempty"`
	AccessMapId *int32 `json:"accessMapId,omitempty"`
	// The CDN prefix registered on behalf of the content provider. For example, `cdn.example.com`.
	Prefix string `json:"prefix"`
	// DNS response time-to-live (TTL) value for this CDN prefix, in seconds. Note that setting this field is only recommended for CDN operators.
	DnsTtl *int32 `json:"dnsTtl,omitempty"`
	CdnPrefixId int32 `json:"cdnPrefixId"`
	IpAddressTagId *int32 `json:"ipAddressTagId,omitempty"`
}

// NewCdnPrefix3 instantiates a new CdnPrefix3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdnPrefix3(prefix string, cdnPrefixId int32) *CdnPrefix3 {
	this := CdnPrefix3{}
	var enable bool = true
	this.Enable = &enable
	var prefixPrioritization PrefixPrioritizationType = PREFIXPRIORITIZATIONTYPE_HIGH
	this.PrefixPrioritization = &prefixPrioritization
	this.Prefix = prefix
	this.CdnPrefixId = cdnPrefixId
	return &this
}

// NewCdnPrefix3WithDefaults instantiates a new CdnPrefix3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdnPrefix3WithDefaults() *CdnPrefix3 {
	this := CdnPrefix3{}
	var enable bool = true
	this.Enable = &enable
	var prefixPrioritization PrefixPrioritizationType = PREFIXPRIORITIZATIONTYPE_HIGH
	this.PrefixPrioritization = &prefixPrioritization
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *CdnPrefix3) GetEnable() bool {
	if o == nil || o.Enable == nil {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnPrefix3) GetEnableOk() (*bool, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *CdnPrefix3) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *CdnPrefix3) SetEnable(v bool) {
	o.Enable = &v
}

// GetContentProviderId returns the ContentProviderId field value if set, zero value otherwise.
func (o *CdnPrefix3) GetContentProviderId() int32 {
	if o == nil || o.ContentProviderId == nil {
		var ret int32
		return ret
	}
	return *o.ContentProviderId
}

// GetContentProviderIdOk returns a tuple with the ContentProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnPrefix3) GetContentProviderIdOk() (*int32, bool) {
	if o == nil || o.ContentProviderId == nil {
		return nil, false
	}
	return o.ContentProviderId, true
}

// HasContentProviderId returns a boolean if a field has been set.
func (o *CdnPrefix3) HasContentProviderId() bool {
	if o != nil && o.ContentProviderId != nil {
		return true
	}

	return false
}

// SetContentProviderId gets a reference to the given int32 and assigns it to the ContentProviderId field.
func (o *CdnPrefix3) SetContentProviderId(v int32) {
	o.ContentProviderId = &v
}

// GetSiteMapId returns the SiteMapId field value if set, zero value otherwise.
func (o *CdnPrefix3) GetSiteMapId() int32 {
	if o == nil || o.SiteMapId == nil {
		var ret int32
		return ret
	}
	return *o.SiteMapId
}

// GetSiteMapIdOk returns a tuple with the SiteMapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnPrefix3) GetSiteMapIdOk() (*int32, bool) {
	if o == nil || o.SiteMapId == nil {
		return nil, false
	}
	return o.SiteMapId, true
}

// HasSiteMapId returns a boolean if a field has been set.
func (o *CdnPrefix3) HasSiteMapId() bool {
	if o != nil && o.SiteMapId != nil {
		return true
	}

	return false
}

// SetSiteMapId gets a reference to the given int32 and assigns it to the SiteMapId field.
func (o *CdnPrefix3) SetSiteMapId(v int32) {
	o.SiteMapId = &v
}

// GetPrefixPrioritization returns the PrefixPrioritization field value if set, zero value otherwise.
func (o *CdnPrefix3) GetPrefixPrioritization() PrefixPrioritizationType {
	if o == nil || o.PrefixPrioritization == nil {
		var ret PrefixPrioritizationType
		return ret
	}
	return *o.PrefixPrioritization
}

// GetPrefixPrioritizationOk returns a tuple with the PrefixPrioritization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnPrefix3) GetPrefixPrioritizationOk() (*PrefixPrioritizationType, bool) {
	if o == nil || o.PrefixPrioritization == nil {
		return nil, false
	}
	return o.PrefixPrioritization, true
}

// HasPrefixPrioritization returns a boolean if a field has been set.
func (o *CdnPrefix3) HasPrefixPrioritization() bool {
	if o != nil && o.PrefixPrioritization != nil {
		return true
	}

	return false
}

// SetPrefixPrioritization gets a reference to the given PrefixPrioritizationType and assigns it to the PrefixPrioritization field.
func (o *CdnPrefix3) SetPrefixPrioritization(v PrefixPrioritizationType) {
	o.PrefixPrioritization = &v
}

// GetKeepaliveRequests returns the KeepaliveRequests field value if set, zero value otherwise.
func (o *CdnPrefix3) GetKeepaliveRequests() int32 {
	if o == nil || o.KeepaliveRequests == nil {
		var ret int32
		return ret
	}
	return *o.KeepaliveRequests
}

// GetKeepaliveRequestsOk returns a tuple with the KeepaliveRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnPrefix3) GetKeepaliveRequestsOk() (*int32, bool) {
	if o == nil || o.KeepaliveRequests == nil {
		return nil, false
	}
	return o.KeepaliveRequests, true
}

// HasKeepaliveRequests returns a boolean if a field has been set.
func (o *CdnPrefix3) HasKeepaliveRequests() bool {
	if o != nil && o.KeepaliveRequests != nil {
		return true
	}

	return false
}

// SetKeepaliveRequests gets a reference to the given int32 and assigns it to the KeepaliveRequests field.
func (o *CdnPrefix3) SetKeepaliveRequests(v int32) {
	o.KeepaliveRequests = &v
}

// GetAccessMapId returns the AccessMapId field value if set, zero value otherwise.
func (o *CdnPrefix3) GetAccessMapId() int32 {
	if o == nil || o.AccessMapId == nil {
		var ret int32
		return ret
	}
	return *o.AccessMapId
}

// GetAccessMapIdOk returns a tuple with the AccessMapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnPrefix3) GetAccessMapIdOk() (*int32, bool) {
	if o == nil || o.AccessMapId == nil {
		return nil, false
	}
	return o.AccessMapId, true
}

// HasAccessMapId returns a boolean if a field has been set.
func (o *CdnPrefix3) HasAccessMapId() bool {
	if o != nil && o.AccessMapId != nil {
		return true
	}

	return false
}

// SetAccessMapId gets a reference to the given int32 and assigns it to the AccessMapId field.
func (o *CdnPrefix3) SetAccessMapId(v int32) {
	o.AccessMapId = &v
}

// GetPrefix returns the Prefix field value
func (o *CdnPrefix3) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *CdnPrefix3) GetPrefixOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *CdnPrefix3) SetPrefix(v string) {
	o.Prefix = v
}

// GetDnsTtl returns the DnsTtl field value if set, zero value otherwise.
func (o *CdnPrefix3) GetDnsTtl() int32 {
	if o == nil || o.DnsTtl == nil {
		var ret int32
		return ret
	}
	return *o.DnsTtl
}

// GetDnsTtlOk returns a tuple with the DnsTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnPrefix3) GetDnsTtlOk() (*int32, bool) {
	if o == nil || o.DnsTtl == nil {
		return nil, false
	}
	return o.DnsTtl, true
}

// HasDnsTtl returns a boolean if a field has been set.
func (o *CdnPrefix3) HasDnsTtl() bool {
	if o != nil && o.DnsTtl != nil {
		return true
	}

	return false
}

// SetDnsTtl gets a reference to the given int32 and assigns it to the DnsTtl field.
func (o *CdnPrefix3) SetDnsTtl(v int32) {
	o.DnsTtl = &v
}

// GetCdnPrefixId returns the CdnPrefixId field value
func (o *CdnPrefix3) GetCdnPrefixId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CdnPrefixId
}

// GetCdnPrefixIdOk returns a tuple with the CdnPrefixId field value
// and a boolean to check if the value has been set.
func (o *CdnPrefix3) GetCdnPrefixIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CdnPrefixId, true
}

// SetCdnPrefixId sets field value
func (o *CdnPrefix3) SetCdnPrefixId(v int32) {
	o.CdnPrefixId = v
}

// GetIpAddressTagId returns the IpAddressTagId field value if set, zero value otherwise.
func (o *CdnPrefix3) GetIpAddressTagId() int32 {
	if o == nil || o.IpAddressTagId == nil {
		var ret int32
		return ret
	}
	return *o.IpAddressTagId
}

// GetIpAddressTagIdOk returns a tuple with the IpAddressTagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnPrefix3) GetIpAddressTagIdOk() (*int32, bool) {
	if o == nil || o.IpAddressTagId == nil {
		return nil, false
	}
	return o.IpAddressTagId, true
}

// HasIpAddressTagId returns a boolean if a field has been set.
func (o *CdnPrefix3) HasIpAddressTagId() bool {
	if o != nil && o.IpAddressTagId != nil {
		return true
	}

	return false
}

// SetIpAddressTagId gets a reference to the given int32 and assigns it to the IpAddressTagId field.
func (o *CdnPrefix3) SetIpAddressTagId(v int32) {
	o.IpAddressTagId = &v
}

func (o CdnPrefix3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enable != nil {
		toSerialize["enable"] = o.Enable
	}
	if o.ContentProviderId != nil {
		toSerialize["contentProviderId"] = o.ContentProviderId
	}
	if o.SiteMapId != nil {
		toSerialize["siteMapId"] = o.SiteMapId
	}
	if o.PrefixPrioritization != nil {
		toSerialize["prefixPrioritization"] = o.PrefixPrioritization
	}
	if o.KeepaliveRequests != nil {
		toSerialize["keepaliveRequests"] = o.KeepaliveRequests
	}
	if o.AccessMapId != nil {
		toSerialize["accessMapId"] = o.AccessMapId
	}
	if true {
		toSerialize["prefix"] = o.Prefix
	}
	if o.DnsTtl != nil {
		toSerialize["dnsTtl"] = o.DnsTtl
	}
	if true {
		toSerialize["cdnPrefixId"] = o.CdnPrefixId
	}
	if o.IpAddressTagId != nil {
		toSerialize["ipAddressTagId"] = o.IpAddressTagId
	}
	return json.Marshal(toSerialize)
}

type NullableCdnPrefix3 struct {
	value *CdnPrefix3
	isSet bool
}

func (v NullableCdnPrefix3) Get() *CdnPrefix3 {
	return v.value
}

func (v *NullableCdnPrefix3) Set(val *CdnPrefix3) {
	v.value = val
	v.isSet = true
}

func (v NullableCdnPrefix3) IsSet() bool {
	return v.isSet
}

func (v *NullableCdnPrefix3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdnPrefix3(val *CdnPrefix3) *NullableCdnPrefix3 {
	return &NullableCdnPrefix3{value: val, isSet: true}
}

func (v NullableCdnPrefix3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdnPrefix3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


