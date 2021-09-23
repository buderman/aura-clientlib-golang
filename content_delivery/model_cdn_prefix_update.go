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

// CdnPrefixUpdate A CDN prefix corresponds to the fully qualified domain name (FQDN) at the beginning of a URI, effectively carving out a region of the URI name space to which a set of rules and policies are applied.
type CdnPrefixUpdate struct {
	Enable *bool `json:"enable,omitempty"`
	ContentProviderId int32 `json:"contentProviderId"`
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

// NewCdnPrefixUpdate instantiates a new CdnPrefixUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdnPrefixUpdate(contentProviderId int32, prefix string, cdnPrefixId int32) *CdnPrefixUpdate {
	this := CdnPrefixUpdate{}
	var enable bool = true
	this.Enable = &enable
	this.ContentProviderId = contentProviderId
	var prefixPrioritization PrefixPrioritizationType = PREFIXPRIORITIZATIONTYPE_HIGH
	this.PrefixPrioritization = &prefixPrioritization
	this.Prefix = prefix
	this.CdnPrefixId = cdnPrefixId
	return &this
}

// NewCdnPrefixUpdateWithDefaults instantiates a new CdnPrefixUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdnPrefixUpdateWithDefaults() *CdnPrefixUpdate {
	this := CdnPrefixUpdate{}
	var enable bool = true
	this.Enable = &enable
	var prefixPrioritization PrefixPrioritizationType = PREFIXPRIORITIZATIONTYPE_HIGH
	this.PrefixPrioritization = &prefixPrioritization
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *CdnPrefixUpdate) GetEnable() bool {
	if o == nil || o.Enable == nil {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnPrefixUpdate) GetEnableOk() (*bool, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *CdnPrefixUpdate) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *CdnPrefixUpdate) SetEnable(v bool) {
	o.Enable = &v
}

// GetContentProviderId returns the ContentProviderId field value
func (o *CdnPrefixUpdate) GetContentProviderId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ContentProviderId
}

// GetContentProviderIdOk returns a tuple with the ContentProviderId field value
// and a boolean to check if the value has been set.
func (o *CdnPrefixUpdate) GetContentProviderIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContentProviderId, true
}

// SetContentProviderId sets field value
func (o *CdnPrefixUpdate) SetContentProviderId(v int32) {
	o.ContentProviderId = v
}

// GetSiteMapId returns the SiteMapId field value if set, zero value otherwise.
func (o *CdnPrefixUpdate) GetSiteMapId() int32 {
	if o == nil || o.SiteMapId == nil {
		var ret int32
		return ret
	}
	return *o.SiteMapId
}

// GetSiteMapIdOk returns a tuple with the SiteMapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnPrefixUpdate) GetSiteMapIdOk() (*int32, bool) {
	if o == nil || o.SiteMapId == nil {
		return nil, false
	}
	return o.SiteMapId, true
}

// HasSiteMapId returns a boolean if a field has been set.
func (o *CdnPrefixUpdate) HasSiteMapId() bool {
	if o != nil && o.SiteMapId != nil {
		return true
	}

	return false
}

// SetSiteMapId gets a reference to the given int32 and assigns it to the SiteMapId field.
func (o *CdnPrefixUpdate) SetSiteMapId(v int32) {
	o.SiteMapId = &v
}

// GetPrefixPrioritization returns the PrefixPrioritization field value if set, zero value otherwise.
func (o *CdnPrefixUpdate) GetPrefixPrioritization() PrefixPrioritizationType {
	if o == nil || o.PrefixPrioritization == nil {
		var ret PrefixPrioritizationType
		return ret
	}
	return *o.PrefixPrioritization
}

// GetPrefixPrioritizationOk returns a tuple with the PrefixPrioritization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnPrefixUpdate) GetPrefixPrioritizationOk() (*PrefixPrioritizationType, bool) {
	if o == nil || o.PrefixPrioritization == nil {
		return nil, false
	}
	return o.PrefixPrioritization, true
}

// HasPrefixPrioritization returns a boolean if a field has been set.
func (o *CdnPrefixUpdate) HasPrefixPrioritization() bool {
	if o != nil && o.PrefixPrioritization != nil {
		return true
	}

	return false
}

// SetPrefixPrioritization gets a reference to the given PrefixPrioritizationType and assigns it to the PrefixPrioritization field.
func (o *CdnPrefixUpdate) SetPrefixPrioritization(v PrefixPrioritizationType) {
	o.PrefixPrioritization = &v
}

// GetKeepaliveRequests returns the KeepaliveRequests field value if set, zero value otherwise.
func (o *CdnPrefixUpdate) GetKeepaliveRequests() int32 {
	if o == nil || o.KeepaliveRequests == nil {
		var ret int32
		return ret
	}
	return *o.KeepaliveRequests
}

// GetKeepaliveRequestsOk returns a tuple with the KeepaliveRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnPrefixUpdate) GetKeepaliveRequestsOk() (*int32, bool) {
	if o == nil || o.KeepaliveRequests == nil {
		return nil, false
	}
	return o.KeepaliveRequests, true
}

// HasKeepaliveRequests returns a boolean if a field has been set.
func (o *CdnPrefixUpdate) HasKeepaliveRequests() bool {
	if o != nil && o.KeepaliveRequests != nil {
		return true
	}

	return false
}

// SetKeepaliveRequests gets a reference to the given int32 and assigns it to the KeepaliveRequests field.
func (o *CdnPrefixUpdate) SetKeepaliveRequests(v int32) {
	o.KeepaliveRequests = &v
}

// GetAccessMapId returns the AccessMapId field value if set, zero value otherwise.
func (o *CdnPrefixUpdate) GetAccessMapId() int32 {
	if o == nil || o.AccessMapId == nil {
		var ret int32
		return ret
	}
	return *o.AccessMapId
}

// GetAccessMapIdOk returns a tuple with the AccessMapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnPrefixUpdate) GetAccessMapIdOk() (*int32, bool) {
	if o == nil || o.AccessMapId == nil {
		return nil, false
	}
	return o.AccessMapId, true
}

// HasAccessMapId returns a boolean if a field has been set.
func (o *CdnPrefixUpdate) HasAccessMapId() bool {
	if o != nil && o.AccessMapId != nil {
		return true
	}

	return false
}

// SetAccessMapId gets a reference to the given int32 and assigns it to the AccessMapId field.
func (o *CdnPrefixUpdate) SetAccessMapId(v int32) {
	o.AccessMapId = &v
}

// GetPrefix returns the Prefix field value
func (o *CdnPrefixUpdate) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *CdnPrefixUpdate) GetPrefixOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *CdnPrefixUpdate) SetPrefix(v string) {
	o.Prefix = v
}

// GetDnsTtl returns the DnsTtl field value if set, zero value otherwise.
func (o *CdnPrefixUpdate) GetDnsTtl() int32 {
	if o == nil || o.DnsTtl == nil {
		var ret int32
		return ret
	}
	return *o.DnsTtl
}

// GetDnsTtlOk returns a tuple with the DnsTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnPrefixUpdate) GetDnsTtlOk() (*int32, bool) {
	if o == nil || o.DnsTtl == nil {
		return nil, false
	}
	return o.DnsTtl, true
}

// HasDnsTtl returns a boolean if a field has been set.
func (o *CdnPrefixUpdate) HasDnsTtl() bool {
	if o != nil && o.DnsTtl != nil {
		return true
	}

	return false
}

// SetDnsTtl gets a reference to the given int32 and assigns it to the DnsTtl field.
func (o *CdnPrefixUpdate) SetDnsTtl(v int32) {
	o.DnsTtl = &v
}

// GetCdnPrefixId returns the CdnPrefixId field value
func (o *CdnPrefixUpdate) GetCdnPrefixId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CdnPrefixId
}

// GetCdnPrefixIdOk returns a tuple with the CdnPrefixId field value
// and a boolean to check if the value has been set.
func (o *CdnPrefixUpdate) GetCdnPrefixIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CdnPrefixId, true
}

// SetCdnPrefixId sets field value
func (o *CdnPrefixUpdate) SetCdnPrefixId(v int32) {
	o.CdnPrefixId = v
}

// GetIpAddressTagId returns the IpAddressTagId field value if set, zero value otherwise.
func (o *CdnPrefixUpdate) GetIpAddressTagId() int32 {
	if o == nil || o.IpAddressTagId == nil {
		var ret int32
		return ret
	}
	return *o.IpAddressTagId
}

// GetIpAddressTagIdOk returns a tuple with the IpAddressTagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnPrefixUpdate) GetIpAddressTagIdOk() (*int32, bool) {
	if o == nil || o.IpAddressTagId == nil {
		return nil, false
	}
	return o.IpAddressTagId, true
}

// HasIpAddressTagId returns a boolean if a field has been set.
func (o *CdnPrefixUpdate) HasIpAddressTagId() bool {
	if o != nil && o.IpAddressTagId != nil {
		return true
	}

	return false
}

// SetIpAddressTagId gets a reference to the given int32 and assigns it to the IpAddressTagId field.
func (o *CdnPrefixUpdate) SetIpAddressTagId(v int32) {
	o.IpAddressTagId = &v
}

func (o CdnPrefixUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enable != nil {
		toSerialize["enable"] = o.Enable
	}
	if true {
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

type NullableCdnPrefixUpdate struct {
	value *CdnPrefixUpdate
	isSet bool
}

func (v NullableCdnPrefixUpdate) Get() *CdnPrefixUpdate {
	return v.value
}

func (v *NullableCdnPrefixUpdate) Set(val *CdnPrefixUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableCdnPrefixUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableCdnPrefixUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdnPrefixUpdate(val *CdnPrefixUpdate) *NullableCdnPrefixUpdate {
	return &NullableCdnPrefixUpdate{value: val, isSet: true}
}

func (v NullableCdnPrefixUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdnPrefixUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


