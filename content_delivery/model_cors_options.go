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

// CorsOptions This behavior's set of configuration options.
type CorsOptions struct {
	// `Access-Control-Allow-Methods` method list. Enter a list of allowed methods or leave empty to use value of `Access-Control-Request-Method` header.
	AccessControlAllowMethods *[]string `json:"accessControlAllowMethods,omitempty"`
	AccessControlAllowOrigin *CorsOptionsAccessControlAllowOrigin `json:"accessControlAllowOrigin,omitempty"`
	// `Access-Control-Max-Age` header value. Enter number of seconds for which client may cache preflight response or leave empty to use default of one day.
	AccessControlMaxAge *int32 `json:"accessControlMaxAge,omitempty"`
	// Enables preflight requests.
	EnablePreflight *bool `json:"enablePreflight,omitempty"`
	// `Access-Control-Expose-Headers` header value.
	AccessControlExposeHeaders *[]string `json:"accessControlExposeHeaders,omitempty"`
	// `Access-Control-Allow-Credentials` header value.
	AccessControlAllowCredentials *bool `json:"accessControlAllowCredentials,omitempty"`
	// `Access-Control-Allow-Headers` header list. Enter a list of allowed headers or leave empty to use value of `Access-Control-Request-Headers` header.
	AccessControlAllowHeaders *[]string `json:"accessControlAllowHeaders,omitempty"`
}

// NewCorsOptions instantiates a new CorsOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCorsOptions() *CorsOptions {
	this := CorsOptions{}
	return &this
}

// NewCorsOptionsWithDefaults instantiates a new CorsOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCorsOptionsWithDefaults() *CorsOptions {
	this := CorsOptions{}
	return &this
}

// GetAccessControlAllowMethods returns the AccessControlAllowMethods field value if set, zero value otherwise.
func (o *CorsOptions) GetAccessControlAllowMethods() []string {
	if o == nil || o.AccessControlAllowMethods == nil {
		var ret []string
		return ret
	}
	return *o.AccessControlAllowMethods
}

// GetAccessControlAllowMethodsOk returns a tuple with the AccessControlAllowMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorsOptions) GetAccessControlAllowMethodsOk() (*[]string, bool) {
	if o == nil || o.AccessControlAllowMethods == nil {
		return nil, false
	}
	return o.AccessControlAllowMethods, true
}

// HasAccessControlAllowMethods returns a boolean if a field has been set.
func (o *CorsOptions) HasAccessControlAllowMethods() bool {
	if o != nil && o.AccessControlAllowMethods != nil {
		return true
	}

	return false
}

// SetAccessControlAllowMethods gets a reference to the given []string and assigns it to the AccessControlAllowMethods field.
func (o *CorsOptions) SetAccessControlAllowMethods(v []string) {
	o.AccessControlAllowMethods = &v
}

// GetAccessControlAllowOrigin returns the AccessControlAllowOrigin field value if set, zero value otherwise.
func (o *CorsOptions) GetAccessControlAllowOrigin() CorsOptionsAccessControlAllowOrigin {
	if o == nil || o.AccessControlAllowOrigin == nil {
		var ret CorsOptionsAccessControlAllowOrigin
		return ret
	}
	return *o.AccessControlAllowOrigin
}

// GetAccessControlAllowOriginOk returns a tuple with the AccessControlAllowOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorsOptions) GetAccessControlAllowOriginOk() (*CorsOptionsAccessControlAllowOrigin, bool) {
	if o == nil || o.AccessControlAllowOrigin == nil {
		return nil, false
	}
	return o.AccessControlAllowOrigin, true
}

// HasAccessControlAllowOrigin returns a boolean if a field has been set.
func (o *CorsOptions) HasAccessControlAllowOrigin() bool {
	if o != nil && o.AccessControlAllowOrigin != nil {
		return true
	}

	return false
}

// SetAccessControlAllowOrigin gets a reference to the given CorsOptionsAccessControlAllowOrigin and assigns it to the AccessControlAllowOrigin field.
func (o *CorsOptions) SetAccessControlAllowOrigin(v CorsOptionsAccessControlAllowOrigin) {
	o.AccessControlAllowOrigin = &v
}

// GetAccessControlMaxAge returns the AccessControlMaxAge field value if set, zero value otherwise.
func (o *CorsOptions) GetAccessControlMaxAge() int32 {
	if o == nil || o.AccessControlMaxAge == nil {
		var ret int32
		return ret
	}
	return *o.AccessControlMaxAge
}

// GetAccessControlMaxAgeOk returns a tuple with the AccessControlMaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorsOptions) GetAccessControlMaxAgeOk() (*int32, bool) {
	if o == nil || o.AccessControlMaxAge == nil {
		return nil, false
	}
	return o.AccessControlMaxAge, true
}

// HasAccessControlMaxAge returns a boolean if a field has been set.
func (o *CorsOptions) HasAccessControlMaxAge() bool {
	if o != nil && o.AccessControlMaxAge != nil {
		return true
	}

	return false
}

// SetAccessControlMaxAge gets a reference to the given int32 and assigns it to the AccessControlMaxAge field.
func (o *CorsOptions) SetAccessControlMaxAge(v int32) {
	o.AccessControlMaxAge = &v
}

// GetEnablePreflight returns the EnablePreflight field value if set, zero value otherwise.
func (o *CorsOptions) GetEnablePreflight() bool {
	if o == nil || o.EnablePreflight == nil {
		var ret bool
		return ret
	}
	return *o.EnablePreflight
}

// GetEnablePreflightOk returns a tuple with the EnablePreflight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorsOptions) GetEnablePreflightOk() (*bool, bool) {
	if o == nil || o.EnablePreflight == nil {
		return nil, false
	}
	return o.EnablePreflight, true
}

// HasEnablePreflight returns a boolean if a field has been set.
func (o *CorsOptions) HasEnablePreflight() bool {
	if o != nil && o.EnablePreflight != nil {
		return true
	}

	return false
}

// SetEnablePreflight gets a reference to the given bool and assigns it to the EnablePreflight field.
func (o *CorsOptions) SetEnablePreflight(v bool) {
	o.EnablePreflight = &v
}

// GetAccessControlExposeHeaders returns the AccessControlExposeHeaders field value if set, zero value otherwise.
func (o *CorsOptions) GetAccessControlExposeHeaders() []string {
	if o == nil || o.AccessControlExposeHeaders == nil {
		var ret []string
		return ret
	}
	return *o.AccessControlExposeHeaders
}

// GetAccessControlExposeHeadersOk returns a tuple with the AccessControlExposeHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorsOptions) GetAccessControlExposeHeadersOk() (*[]string, bool) {
	if o == nil || o.AccessControlExposeHeaders == nil {
		return nil, false
	}
	return o.AccessControlExposeHeaders, true
}

// HasAccessControlExposeHeaders returns a boolean if a field has been set.
func (o *CorsOptions) HasAccessControlExposeHeaders() bool {
	if o != nil && o.AccessControlExposeHeaders != nil {
		return true
	}

	return false
}

// SetAccessControlExposeHeaders gets a reference to the given []string and assigns it to the AccessControlExposeHeaders field.
func (o *CorsOptions) SetAccessControlExposeHeaders(v []string) {
	o.AccessControlExposeHeaders = &v
}

// GetAccessControlAllowCredentials returns the AccessControlAllowCredentials field value if set, zero value otherwise.
func (o *CorsOptions) GetAccessControlAllowCredentials() bool {
	if o == nil || o.AccessControlAllowCredentials == nil {
		var ret bool
		return ret
	}
	return *o.AccessControlAllowCredentials
}

// GetAccessControlAllowCredentialsOk returns a tuple with the AccessControlAllowCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorsOptions) GetAccessControlAllowCredentialsOk() (*bool, bool) {
	if o == nil || o.AccessControlAllowCredentials == nil {
		return nil, false
	}
	return o.AccessControlAllowCredentials, true
}

// HasAccessControlAllowCredentials returns a boolean if a field has been set.
func (o *CorsOptions) HasAccessControlAllowCredentials() bool {
	if o != nil && o.AccessControlAllowCredentials != nil {
		return true
	}

	return false
}

// SetAccessControlAllowCredentials gets a reference to the given bool and assigns it to the AccessControlAllowCredentials field.
func (o *CorsOptions) SetAccessControlAllowCredentials(v bool) {
	o.AccessControlAllowCredentials = &v
}

// GetAccessControlAllowHeaders returns the AccessControlAllowHeaders field value if set, zero value otherwise.
func (o *CorsOptions) GetAccessControlAllowHeaders() []string {
	if o == nil || o.AccessControlAllowHeaders == nil {
		var ret []string
		return ret
	}
	return *o.AccessControlAllowHeaders
}

// GetAccessControlAllowHeadersOk returns a tuple with the AccessControlAllowHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorsOptions) GetAccessControlAllowHeadersOk() (*[]string, bool) {
	if o == nil || o.AccessControlAllowHeaders == nil {
		return nil, false
	}
	return o.AccessControlAllowHeaders, true
}

// HasAccessControlAllowHeaders returns a boolean if a field has been set.
func (o *CorsOptions) HasAccessControlAllowHeaders() bool {
	if o != nil && o.AccessControlAllowHeaders != nil {
		return true
	}

	return false
}

// SetAccessControlAllowHeaders gets a reference to the given []string and assigns it to the AccessControlAllowHeaders field.
func (o *CorsOptions) SetAccessControlAllowHeaders(v []string) {
	o.AccessControlAllowHeaders = &v
}

func (o CorsOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessControlAllowMethods != nil {
		toSerialize["accessControlAllowMethods"] = o.AccessControlAllowMethods
	}
	if o.AccessControlAllowOrigin != nil {
		toSerialize["accessControlAllowOrigin"] = o.AccessControlAllowOrigin
	}
	if o.AccessControlMaxAge != nil {
		toSerialize["accessControlMaxAge"] = o.AccessControlMaxAge
	}
	if o.EnablePreflight != nil {
		toSerialize["enablePreflight"] = o.EnablePreflight
	}
	if o.AccessControlExposeHeaders != nil {
		toSerialize["accessControlExposeHeaders"] = o.AccessControlExposeHeaders
	}
	if o.AccessControlAllowCredentials != nil {
		toSerialize["accessControlAllowCredentials"] = o.AccessControlAllowCredentials
	}
	if o.AccessControlAllowHeaders != nil {
		toSerialize["accessControlAllowHeaders"] = o.AccessControlAllowHeaders
	}
	return json.Marshal(toSerialize)
}

type NullableCorsOptions struct {
	value *CorsOptions
	isSet bool
}

func (v NullableCorsOptions) Get() *CorsOptions {
	return v.value
}

func (v *NullableCorsOptions) Set(val *CorsOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableCorsOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableCorsOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCorsOptions(val *CorsOptions) *NullableCorsOptions {
	return &NullableCorsOptions{value: val, isSet: true}
}

func (v NullableCorsOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCorsOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


