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

// Origin3 Origin summary read
type Origin3 struct {
	Hostname string `json:"hostname"`
	CachingType *CachingType `json:"cachingType,omitempty"`
	// A DNS hostname to use when connecting to the origin instead of that specified by the `hostname` field.  A port may be optionally specified by using hostname:port format.
	HostnameOverride *string `json:"hostnameOverride,omitempty"`
	StoragePartitionId *int32 `json:"storagePartitionId,omitempty"`
	EdgeHostType *EdgeHostType `json:"edgeHostType,omitempty"`
	CacheKeyOriginHostnameOriginId *int32 `json:"cacheKeyOriginHostnameOriginId,omitempty"`
	ContentProviderId *int32 `json:"contentProviderId,omitempty"`
	FastReroute *FastRerouteType `json:"fastReroute,omitempty"`
	// Determines whether or not HyperCache redirects a client to a better site based on the rules in the site map. Enabling this to `true` allows the HyperCache to adhere to the LCDN site map configured using the request router service. If set to `false`, cache miss client requests to a CDN prefix that uses this origin server as a default result in a HyperCache request for the content directly to the origin server.
	EnableSiteRedirects *bool `json:"enableSiteRedirects,omitempty"`
	InterSiteProtocol *ProtocolType `json:"interSiteProtocol,omitempty"`
	// Causes the Hypercache to send its internal request ID to the origin server when ingesting content.
	EnableRequestIdExport *bool `json:"enableRequestIdExport,omitempty"`
	OriginId int32 `json:"originId"`
	Enable *bool `json:"enable,omitempty"`
	// Causes the HyperCache to send a HEAD request to the origin server for each client request. If the response is a status code of 2XX, the content is served to the client. Otherwise, the origin's status code is returned. This typically occurs when the origin server authenticates the client request using a cookie, token, or other means. Default is `false`.
	EnableAuthenticatedContent *bool `json:"enableAuthenticatedContent,omitempty"`
	DynamicHierarchy *DynamicHierarchyType `json:"dynamicHierarchy,omitempty"`
	TlsIngestProfileId *int32 `json:"tlsIngestProfileId,omitempty"`
	IntraSiteProtocol *ProtocolType `json:"intraSiteProtocol,omitempty"`
	// A list of HTTP status codes that the HPC caches. Each item is a single status code or range of status codes, where the starting code must be less than the ending code. For example, `209` and `310-320` are both valid. The following status codes are not allowed, either in single code, or as part of a range: 200, 203, 206, 300, 301, 410, 416.
	CacheableErrorResponseCodes *[]string `json:"cacheableErrorResponseCodes,omitempty"`
	// A list of virtual hosts, origin servers with resolvable hostnames. These resolvable hostnames are used by DNS to determine the IP addresses of the destination origin server.
	ResolvableHostnames *[]string `json:"resolvableHostnames,omitempty"`
	// The maximum number of retries in the case of an error response.
	ErrorCacheMaxRetry *int32 `json:"errorCacheMaxRetry,omitempty"`
	EdgeHostname *string `json:"edgeHostname,omitempty"`
	// Origin idle connection timeout in seconds, specifies the length of time the HPC waits for a response to a request from an origin. If no response is received from an origin server within the configured timeout value, four additional attempts are made to the same origin server. If no response is received for the final request within the configured timeout, then that connection is terminated, and a 504 response is sent to the requesting client.
	OriginTimeout *int32 `json:"originTimeout,omitempty"`
	// The maximum amount of time the HTTP status codes is cached by the HPC.
	ErrorCacheMaxAge *int32 `json:"errorCacheMaxAge,omitempty"`
}

// NewOrigin3 instantiates a new Origin3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrigin3(hostname string, originId int32) *Origin3 {
	this := Origin3{}
	this.Hostname = hostname
	var cachingType CachingType = CACHINGTYPE_OPTIMISTIC
	this.CachingType = &cachingType
	var edgeHostType EdgeHostType = EDGEHOSTTYPE_HTTP
	this.EdgeHostType = &edgeHostType
	var enableSiteRedirects bool = true
	this.EnableSiteRedirects = &enableSiteRedirects
	var interSiteProtocol ProtocolType = PROTOCOLTYPE_HTTP
	this.InterSiteProtocol = &interSiteProtocol
	var enableRequestIdExport bool = false
	this.EnableRequestIdExport = &enableRequestIdExport
	this.OriginId = originId
	var enable bool = true
	this.Enable = &enable
	var enableAuthenticatedContent bool = false
	this.EnableAuthenticatedContent = &enableAuthenticatedContent
	var intraSiteProtocol ProtocolType = PROTOCOLTYPE_HTTP
	this.IntraSiteProtocol = &intraSiteProtocol
	var originTimeout int32 = 300
	this.OriginTimeout = &originTimeout
	return &this
}

// NewOrigin3WithDefaults instantiates a new Origin3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrigin3WithDefaults() *Origin3 {
	this := Origin3{}
	var cachingType CachingType = CACHINGTYPE_OPTIMISTIC
	this.CachingType = &cachingType
	var edgeHostType EdgeHostType = EDGEHOSTTYPE_HTTP
	this.EdgeHostType = &edgeHostType
	var enableSiteRedirects bool = true
	this.EnableSiteRedirects = &enableSiteRedirects
	var interSiteProtocol ProtocolType = PROTOCOLTYPE_HTTP
	this.InterSiteProtocol = &interSiteProtocol
	var enableRequestIdExport bool = false
	this.EnableRequestIdExport = &enableRequestIdExport
	var enable bool = true
	this.Enable = &enable
	var enableAuthenticatedContent bool = false
	this.EnableAuthenticatedContent = &enableAuthenticatedContent
	var intraSiteProtocol ProtocolType = PROTOCOLTYPE_HTTP
	this.IntraSiteProtocol = &intraSiteProtocol
	var originTimeout int32 = 300
	this.OriginTimeout = &originTimeout
	return &this
}

// GetHostname returns the Hostname field value
func (o *Origin3) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *Origin3) GetHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *Origin3) SetHostname(v string) {
	o.Hostname = v
}

// GetCachingType returns the CachingType field value if set, zero value otherwise.
func (o *Origin3) GetCachingType() CachingType {
	if o == nil || o.CachingType == nil {
		var ret CachingType
		return ret
	}
	return *o.CachingType
}

// GetCachingTypeOk returns a tuple with the CachingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Origin3) GetCachingTypeOk() (*CachingType, bool) {
	if o == nil || o.CachingType == nil {
		return nil, false
	}
	return o.CachingType, true
}

// HasCachingType returns a boolean if a field has been set.
func (o *Origin3) HasCachingType() bool {
	if o != nil && o.CachingType != nil {
		return true
	}

	return false
}

// SetCachingType gets a reference to the given CachingType and assigns it to the CachingType field.
func (o *Origin3) SetCachingType(v CachingType) {
	o.CachingType = &v
}

// GetHostnameOverride returns the HostnameOverride field value if set, zero value otherwise.
func (o *Origin3) GetHostnameOverride() string {
	if o == nil || o.HostnameOverride == nil {
		var ret string
		return ret
	}
	return *o.HostnameOverride
}

// GetHostnameOverrideOk returns a tuple with the HostnameOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Origin3) GetHostnameOverrideOk() (*string, bool) {
	if o == nil || o.HostnameOverride == nil {
		return nil, false
	}
	return o.HostnameOverride, true
}

// HasHostnameOverride returns a boolean if a field has been set.
func (o *Origin3) HasHostnameOverride() bool {
	if o != nil && o.HostnameOverride != nil {
		return true
	}

	return false
}

// SetHostnameOverride gets a reference to the given string and assigns it to the HostnameOverride field.
func (o *Origin3) SetHostnameOverride(v string) {
	o.HostnameOverride = &v
}

// GetStoragePartitionId returns the StoragePartitionId field value if set, zero value otherwise.
func (o *Origin3) GetStoragePartitionId() int32 {
	if o == nil || o.StoragePartitionId == nil {
		var ret int32
		return ret
	}
	return *o.StoragePartitionId
}

// GetStoragePartitionIdOk returns a tuple with the StoragePartitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Origin3) GetStoragePartitionIdOk() (*int32, bool) {
	if o == nil || o.StoragePartitionId == nil {
		return nil, false
	}
	return o.StoragePartitionId, true
}

// HasStoragePartitionId returns a boolean if a field has been set.
func (o *Origin3) HasStoragePartitionId() bool {
	if o != nil && o.StoragePartitionId != nil {
		return true
	}

	return false
}

// SetStoragePartitionId gets a reference to the given int32 and assigns it to the StoragePartitionId field.
func (o *Origin3) SetStoragePartitionId(v int32) {
	o.StoragePartitionId = &v
}

// GetEdgeHostType returns the EdgeHostType field value if set, zero value otherwise.
func (o *Origin3) GetEdgeHostType() EdgeHostType {
	if o == nil || o.EdgeHostType == nil {
		var ret EdgeHostType
		return ret
	}
	return *o.EdgeHostType
}

// GetEdgeHostTypeOk returns a tuple with the EdgeHostType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Origin3) GetEdgeHostTypeOk() (*EdgeHostType, bool) {
	if o == nil || o.EdgeHostType == nil {
		return nil, false
	}
	return o.EdgeHostType, true
}

// HasEdgeHostType returns a boolean if a field has been set.
func (o *Origin3) HasEdgeHostType() bool {
	if o != nil && o.EdgeHostType != nil {
		return true
	}

	return false
}

// SetEdgeHostType gets a reference to the given EdgeHostType and assigns it to the EdgeHostType field.
func (o *Origin3) SetEdgeHostType(v EdgeHostType) {
	o.EdgeHostType = &v
}

// GetCacheKeyOriginHostnameOriginId returns the CacheKeyOriginHostnameOriginId field value if set, zero value otherwise.
func (o *Origin3) GetCacheKeyOriginHostnameOriginId() int32 {
	if o == nil || o.CacheKeyOriginHostnameOriginId == nil {
		var ret int32
		return ret
	}
	return *o.CacheKeyOriginHostnameOriginId
}

// GetCacheKeyOriginHostnameOriginIdOk returns a tuple with the CacheKeyOriginHostnameOriginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Origin3) GetCacheKeyOriginHostnameOriginIdOk() (*int32, bool) {
	if o == nil || o.CacheKeyOriginHostnameOriginId == nil {
		return nil, false
	}
	return o.CacheKeyOriginHostnameOriginId, true
}

// HasCacheKeyOriginHostnameOriginId returns a boolean if a field has been set.
func (o *Origin3) HasCacheKeyOriginHostnameOriginId() bool {
	if o != nil && o.CacheKeyOriginHostnameOriginId != nil {
		return true
	}

	return false
}

// SetCacheKeyOriginHostnameOriginId gets a reference to the given int32 and assigns it to the CacheKeyOriginHostnameOriginId field.
func (o *Origin3) SetCacheKeyOriginHostnameOriginId(v int32) {
	o.CacheKeyOriginHostnameOriginId = &v
}

// GetContentProviderId returns the ContentProviderId field value if set, zero value otherwise.
func (o *Origin3) GetContentProviderId() int32 {
	if o == nil || o.ContentProviderId == nil {
		var ret int32
		return ret
	}
	return *o.ContentProviderId
}

// GetContentProviderIdOk returns a tuple with the ContentProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Origin3) GetContentProviderIdOk() (*int32, bool) {
	if o == nil || o.ContentProviderId == nil {
		return nil, false
	}
	return o.ContentProviderId, true
}

// HasContentProviderId returns a boolean if a field has been set.
func (o *Origin3) HasContentProviderId() bool {
	if o != nil && o.ContentProviderId != nil {
		return true
	}

	return false
}

// SetContentProviderId gets a reference to the given int32 and assigns it to the ContentProviderId field.
func (o *Origin3) SetContentProviderId(v int32) {
	o.ContentProviderId = &v
}

// GetFastReroute returns the FastReroute field value if set, zero value otherwise.
func (o *Origin3) GetFastReroute() FastRerouteType {
	if o == nil || o.FastReroute == nil {
		var ret FastRerouteType
		return ret
	}
	return *o.FastReroute
}

// GetFastRerouteOk returns a tuple with the FastReroute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Origin3) GetFastRerouteOk() (*FastRerouteType, bool) {
	if o == nil || o.FastReroute == nil {
		return nil, false
	}
	return o.FastReroute, true
}

// HasFastReroute returns a boolean if a field has been set.
func (o *Origin3) HasFastReroute() bool {
	if o != nil && o.FastReroute != nil {
		return true
	}

	return false
}

// SetFastReroute gets a reference to the given FastRerouteType and assigns it to the FastReroute field.
func (o *Origin3) SetFastReroute(v FastRerouteType) {
	o.FastReroute = &v
}

// GetEnableSiteRedirects returns the EnableSiteRedirects field value if set, zero value otherwise.
func (o *Origin3) GetEnableSiteRedirects() bool {
	if o == nil || o.EnableSiteRedirects == nil {
		var ret bool
		return ret
	}
	return *o.EnableSiteRedirects
}

// GetEnableSiteRedirectsOk returns a tuple with the EnableSiteRedirects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Origin3) GetEnableSiteRedirectsOk() (*bool, bool) {
	if o == nil || o.EnableSiteRedirects == nil {
		return nil, false
	}
	return o.EnableSiteRedirects, true
}

// HasEnableSiteRedirects returns a boolean if a field has been set.
func (o *Origin3) HasEnableSiteRedirects() bool {
	if o != nil && o.EnableSiteRedirects != nil {
		return true
	}

	return false
}

// SetEnableSiteRedirects gets a reference to the given bool and assigns it to the EnableSiteRedirects field.
func (o *Origin3) SetEnableSiteRedirects(v bool) {
	o.EnableSiteRedirects = &v
}

// GetInterSiteProtocol returns the InterSiteProtocol field value if set, zero value otherwise.
func (o *Origin3) GetInterSiteProtocol() ProtocolType {
	if o == nil || o.InterSiteProtocol == nil {
		var ret ProtocolType
		return ret
	}
	return *o.InterSiteProtocol
}

// GetInterSiteProtocolOk returns a tuple with the InterSiteProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Origin3) GetInterSiteProtocolOk() (*ProtocolType, bool) {
	if o == nil || o.InterSiteProtocol == nil {
		return nil, false
	}
	return o.InterSiteProtocol, true
}

// HasInterSiteProtocol returns a boolean if a field has been set.
func (o *Origin3) HasInterSiteProtocol() bool {
	if o != nil && o.InterSiteProtocol != nil {
		return true
	}

	return false
}

// SetInterSiteProtocol gets a reference to the given ProtocolType and assigns it to the InterSiteProtocol field.
func (o *Origin3) SetInterSiteProtocol(v ProtocolType) {
	o.InterSiteProtocol = &v
}

// GetEnableRequestIdExport returns the EnableRequestIdExport field value if set, zero value otherwise.
func (o *Origin3) GetEnableRequestIdExport() bool {
	if o == nil || o.EnableRequestIdExport == nil {
		var ret bool
		return ret
	}
	return *o.EnableRequestIdExport
}

// GetEnableRequestIdExportOk returns a tuple with the EnableRequestIdExport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Origin3) GetEnableRequestIdExportOk() (*bool, bool) {
	if o == nil || o.EnableRequestIdExport == nil {
		return nil, false
	}
	return o.EnableRequestIdExport, true
}

// HasEnableRequestIdExport returns a boolean if a field has been set.
func (o *Origin3) HasEnableRequestIdExport() bool {
	if o != nil && o.EnableRequestIdExport != nil {
		return true
	}

	return false
}

// SetEnableRequestIdExport gets a reference to the given bool and assigns it to the EnableRequestIdExport field.
func (o *Origin3) SetEnableRequestIdExport(v bool) {
	o.EnableRequestIdExport = &v
}

// GetOriginId returns the OriginId field value
func (o *Origin3) GetOriginId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OriginId
}

// GetOriginIdOk returns a tuple with the OriginId field value
// and a boolean to check if the value has been set.
func (o *Origin3) GetOriginIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OriginId, true
}

// SetOriginId sets field value
func (o *Origin3) SetOriginId(v int32) {
	o.OriginId = v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *Origin3) GetEnable() bool {
	if o == nil || o.Enable == nil {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Origin3) GetEnableOk() (*bool, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *Origin3) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *Origin3) SetEnable(v bool) {
	o.Enable = &v
}

// GetEnableAuthenticatedContent returns the EnableAuthenticatedContent field value if set, zero value otherwise.
func (o *Origin3) GetEnableAuthenticatedContent() bool {
	if o == nil || o.EnableAuthenticatedContent == nil {
		var ret bool
		return ret
	}
	return *o.EnableAuthenticatedContent
}

// GetEnableAuthenticatedContentOk returns a tuple with the EnableAuthenticatedContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Origin3) GetEnableAuthenticatedContentOk() (*bool, bool) {
	if o == nil || o.EnableAuthenticatedContent == nil {
		return nil, false
	}
	return o.EnableAuthenticatedContent, true
}

// HasEnableAuthenticatedContent returns a boolean if a field has been set.
func (o *Origin3) HasEnableAuthenticatedContent() bool {
	if o != nil && o.EnableAuthenticatedContent != nil {
		return true
	}

	return false
}

// SetEnableAuthenticatedContent gets a reference to the given bool and assigns it to the EnableAuthenticatedContent field.
func (o *Origin3) SetEnableAuthenticatedContent(v bool) {
	o.EnableAuthenticatedContent = &v
}

// GetDynamicHierarchy returns the DynamicHierarchy field value if set, zero value otherwise.
func (o *Origin3) GetDynamicHierarchy() DynamicHierarchyType {
	if o == nil || o.DynamicHierarchy == nil {
		var ret DynamicHierarchyType
		return ret
	}
	return *o.DynamicHierarchy
}

// GetDynamicHierarchyOk returns a tuple with the DynamicHierarchy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Origin3) GetDynamicHierarchyOk() (*DynamicHierarchyType, bool) {
	if o == nil || o.DynamicHierarchy == nil {
		return nil, false
	}
	return o.DynamicHierarchy, true
}

// HasDynamicHierarchy returns a boolean if a field has been set.
func (o *Origin3) HasDynamicHierarchy() bool {
	if o != nil && o.DynamicHierarchy != nil {
		return true
	}

	return false
}

// SetDynamicHierarchy gets a reference to the given DynamicHierarchyType and assigns it to the DynamicHierarchy field.
func (o *Origin3) SetDynamicHierarchy(v DynamicHierarchyType) {
	o.DynamicHierarchy = &v
}

// GetTlsIngestProfileId returns the TlsIngestProfileId field value if set, zero value otherwise.
func (o *Origin3) GetTlsIngestProfileId() int32 {
	if o == nil || o.TlsIngestProfileId == nil {
		var ret int32
		return ret
	}
	return *o.TlsIngestProfileId
}

// GetTlsIngestProfileIdOk returns a tuple with the TlsIngestProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Origin3) GetTlsIngestProfileIdOk() (*int32, bool) {
	if o == nil || o.TlsIngestProfileId == nil {
		return nil, false
	}
	return o.TlsIngestProfileId, true
}

// HasTlsIngestProfileId returns a boolean if a field has been set.
func (o *Origin3) HasTlsIngestProfileId() bool {
	if o != nil && o.TlsIngestProfileId != nil {
		return true
	}

	return false
}

// SetTlsIngestProfileId gets a reference to the given int32 and assigns it to the TlsIngestProfileId field.
func (o *Origin3) SetTlsIngestProfileId(v int32) {
	o.TlsIngestProfileId = &v
}

// GetIntraSiteProtocol returns the IntraSiteProtocol field value if set, zero value otherwise.
func (o *Origin3) GetIntraSiteProtocol() ProtocolType {
	if o == nil || o.IntraSiteProtocol == nil {
		var ret ProtocolType
		return ret
	}
	return *o.IntraSiteProtocol
}

// GetIntraSiteProtocolOk returns a tuple with the IntraSiteProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Origin3) GetIntraSiteProtocolOk() (*ProtocolType, bool) {
	if o == nil || o.IntraSiteProtocol == nil {
		return nil, false
	}
	return o.IntraSiteProtocol, true
}

// HasIntraSiteProtocol returns a boolean if a field has been set.
func (o *Origin3) HasIntraSiteProtocol() bool {
	if o != nil && o.IntraSiteProtocol != nil {
		return true
	}

	return false
}

// SetIntraSiteProtocol gets a reference to the given ProtocolType and assigns it to the IntraSiteProtocol field.
func (o *Origin3) SetIntraSiteProtocol(v ProtocolType) {
	o.IntraSiteProtocol = &v
}

// GetCacheableErrorResponseCodes returns the CacheableErrorResponseCodes field value if set, zero value otherwise.
func (o *Origin3) GetCacheableErrorResponseCodes() []string {
	if o == nil || o.CacheableErrorResponseCodes == nil {
		var ret []string
		return ret
	}
	return *o.CacheableErrorResponseCodes
}

// GetCacheableErrorResponseCodesOk returns a tuple with the CacheableErrorResponseCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Origin3) GetCacheableErrorResponseCodesOk() (*[]string, bool) {
	if o == nil || o.CacheableErrorResponseCodes == nil {
		return nil, false
	}
	return o.CacheableErrorResponseCodes, true
}

// HasCacheableErrorResponseCodes returns a boolean if a field has been set.
func (o *Origin3) HasCacheableErrorResponseCodes() bool {
	if o != nil && o.CacheableErrorResponseCodes != nil {
		return true
	}

	return false
}

// SetCacheableErrorResponseCodes gets a reference to the given []string and assigns it to the CacheableErrorResponseCodes field.
func (o *Origin3) SetCacheableErrorResponseCodes(v []string) {
	o.CacheableErrorResponseCodes = &v
}

// GetResolvableHostnames returns the ResolvableHostnames field value if set, zero value otherwise.
func (o *Origin3) GetResolvableHostnames() []string {
	if o == nil || o.ResolvableHostnames == nil {
		var ret []string
		return ret
	}
	return *o.ResolvableHostnames
}

// GetResolvableHostnamesOk returns a tuple with the ResolvableHostnames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Origin3) GetResolvableHostnamesOk() (*[]string, bool) {
	if o == nil || o.ResolvableHostnames == nil {
		return nil, false
	}
	return o.ResolvableHostnames, true
}

// HasResolvableHostnames returns a boolean if a field has been set.
func (o *Origin3) HasResolvableHostnames() bool {
	if o != nil && o.ResolvableHostnames != nil {
		return true
	}

	return false
}

// SetResolvableHostnames gets a reference to the given []string and assigns it to the ResolvableHostnames field.
func (o *Origin3) SetResolvableHostnames(v []string) {
	o.ResolvableHostnames = &v
}

// GetErrorCacheMaxRetry returns the ErrorCacheMaxRetry field value if set, zero value otherwise.
func (o *Origin3) GetErrorCacheMaxRetry() int32 {
	if o == nil || o.ErrorCacheMaxRetry == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCacheMaxRetry
}

// GetErrorCacheMaxRetryOk returns a tuple with the ErrorCacheMaxRetry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Origin3) GetErrorCacheMaxRetryOk() (*int32, bool) {
	if o == nil || o.ErrorCacheMaxRetry == nil {
		return nil, false
	}
	return o.ErrorCacheMaxRetry, true
}

// HasErrorCacheMaxRetry returns a boolean if a field has been set.
func (o *Origin3) HasErrorCacheMaxRetry() bool {
	if o != nil && o.ErrorCacheMaxRetry != nil {
		return true
	}

	return false
}

// SetErrorCacheMaxRetry gets a reference to the given int32 and assigns it to the ErrorCacheMaxRetry field.
func (o *Origin3) SetErrorCacheMaxRetry(v int32) {
	o.ErrorCacheMaxRetry = &v
}

// GetEdgeHostname returns the EdgeHostname field value if set, zero value otherwise.
func (o *Origin3) GetEdgeHostname() string {
	if o == nil || o.EdgeHostname == nil {
		var ret string
		return ret
	}
	return *o.EdgeHostname
}

// GetEdgeHostnameOk returns a tuple with the EdgeHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Origin3) GetEdgeHostnameOk() (*string, bool) {
	if o == nil || o.EdgeHostname == nil {
		return nil, false
	}
	return o.EdgeHostname, true
}

// HasEdgeHostname returns a boolean if a field has been set.
func (o *Origin3) HasEdgeHostname() bool {
	if o != nil && o.EdgeHostname != nil {
		return true
	}

	return false
}

// SetEdgeHostname gets a reference to the given string and assigns it to the EdgeHostname field.
func (o *Origin3) SetEdgeHostname(v string) {
	o.EdgeHostname = &v
}

// GetOriginTimeout returns the OriginTimeout field value if set, zero value otherwise.
func (o *Origin3) GetOriginTimeout() int32 {
	if o == nil || o.OriginTimeout == nil {
		var ret int32
		return ret
	}
	return *o.OriginTimeout
}

// GetOriginTimeoutOk returns a tuple with the OriginTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Origin3) GetOriginTimeoutOk() (*int32, bool) {
	if o == nil || o.OriginTimeout == nil {
		return nil, false
	}
	return o.OriginTimeout, true
}

// HasOriginTimeout returns a boolean if a field has been set.
func (o *Origin3) HasOriginTimeout() bool {
	if o != nil && o.OriginTimeout != nil {
		return true
	}

	return false
}

// SetOriginTimeout gets a reference to the given int32 and assigns it to the OriginTimeout field.
func (o *Origin3) SetOriginTimeout(v int32) {
	o.OriginTimeout = &v
}

// GetErrorCacheMaxAge returns the ErrorCacheMaxAge field value if set, zero value otherwise.
func (o *Origin3) GetErrorCacheMaxAge() int32 {
	if o == nil || o.ErrorCacheMaxAge == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCacheMaxAge
}

// GetErrorCacheMaxAgeOk returns a tuple with the ErrorCacheMaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Origin3) GetErrorCacheMaxAgeOk() (*int32, bool) {
	if o == nil || o.ErrorCacheMaxAge == nil {
		return nil, false
	}
	return o.ErrorCacheMaxAge, true
}

// HasErrorCacheMaxAge returns a boolean if a field has been set.
func (o *Origin3) HasErrorCacheMaxAge() bool {
	if o != nil && o.ErrorCacheMaxAge != nil {
		return true
	}

	return false
}

// SetErrorCacheMaxAge gets a reference to the given int32 and assigns it to the ErrorCacheMaxAge field.
func (o *Origin3) SetErrorCacheMaxAge(v int32) {
	o.ErrorCacheMaxAge = &v
}

func (o Origin3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hostname"] = o.Hostname
	}
	if o.CachingType != nil {
		toSerialize["cachingType"] = o.CachingType
	}
	if o.HostnameOverride != nil {
		toSerialize["hostnameOverride"] = o.HostnameOverride
	}
	if o.StoragePartitionId != nil {
		toSerialize["storagePartitionId"] = o.StoragePartitionId
	}
	if o.EdgeHostType != nil {
		toSerialize["edgeHostType"] = o.EdgeHostType
	}
	if o.CacheKeyOriginHostnameOriginId != nil {
		toSerialize["cacheKeyOriginHostnameOriginId"] = o.CacheKeyOriginHostnameOriginId
	}
	if o.ContentProviderId != nil {
		toSerialize["contentProviderId"] = o.ContentProviderId
	}
	if o.FastReroute != nil {
		toSerialize["fastReroute"] = o.FastReroute
	}
	if o.EnableSiteRedirects != nil {
		toSerialize["enableSiteRedirects"] = o.EnableSiteRedirects
	}
	if o.InterSiteProtocol != nil {
		toSerialize["interSiteProtocol"] = o.InterSiteProtocol
	}
	if o.EnableRequestIdExport != nil {
		toSerialize["enableRequestIdExport"] = o.EnableRequestIdExport
	}
	if true {
		toSerialize["originId"] = o.OriginId
	}
	if o.Enable != nil {
		toSerialize["enable"] = o.Enable
	}
	if o.EnableAuthenticatedContent != nil {
		toSerialize["enableAuthenticatedContent"] = o.EnableAuthenticatedContent
	}
	if o.DynamicHierarchy != nil {
		toSerialize["dynamicHierarchy"] = o.DynamicHierarchy
	}
	if o.TlsIngestProfileId != nil {
		toSerialize["tlsIngestProfileId"] = o.TlsIngestProfileId
	}
	if o.IntraSiteProtocol != nil {
		toSerialize["intraSiteProtocol"] = o.IntraSiteProtocol
	}
	if o.CacheableErrorResponseCodes != nil {
		toSerialize["cacheableErrorResponseCodes"] = o.CacheableErrorResponseCodes
	}
	if o.ResolvableHostnames != nil {
		toSerialize["resolvableHostnames"] = o.ResolvableHostnames
	}
	if o.ErrorCacheMaxRetry != nil {
		toSerialize["errorCacheMaxRetry"] = o.ErrorCacheMaxRetry
	}
	if o.EdgeHostname != nil {
		toSerialize["edgeHostname"] = o.EdgeHostname
	}
	if o.OriginTimeout != nil {
		toSerialize["originTimeout"] = o.OriginTimeout
	}
	if o.ErrorCacheMaxAge != nil {
		toSerialize["errorCacheMaxAge"] = o.ErrorCacheMaxAge
	}
	return json.Marshal(toSerialize)
}

type NullableOrigin3 struct {
	value *Origin3
	isSet bool
}

func (v NullableOrigin3) Get() *Origin3 {
	return v.value
}

func (v *NullableOrigin3) Set(val *Origin3) {
	v.value = val
	v.isSet = true
}

func (v NullableOrigin3) IsSet() bool {
	return v.isSet
}

func (v *NullableOrigin3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrigin3(val *Origin3) *NullableOrigin3 {
	return &NullableOrigin3{value: val, isSet: true}
}

func (v NullableOrigin3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrigin3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


