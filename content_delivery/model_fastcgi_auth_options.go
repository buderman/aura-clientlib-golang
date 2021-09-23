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

// FastcgiAuthOptions This behavior's set of configuration options.
type FastcgiAuthOptions struct {
	// Role set for the FastCGI Server, either `AUTHORIZER` or `RESPONDER`.
	Role *string `json:"role,omitempty"`
	// Enables FastCGI Authentication.
	Enable bool `json:"enable"`
	// FastCGI parameters in map form.
	Params *map[string]string `json:"params,omitempty"`
	// Destination FastCGI server server:port, maximum length is 262, 256 of which is for the hostname.
	Server string `json:"server"`
	// List of backup destination FastCGI Servers, expressed as `server:port`.
	BackupServers *[]string `json:"backupServers,omitempty"`
}

// NewFastcgiAuthOptions instantiates a new FastcgiAuthOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFastcgiAuthOptions(enable bool, server string) *FastcgiAuthOptions {
	this := FastcgiAuthOptions{}
	this.Enable = enable
	this.Server = server
	return &this
}

// NewFastcgiAuthOptionsWithDefaults instantiates a new FastcgiAuthOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFastcgiAuthOptionsWithDefaults() *FastcgiAuthOptions {
	this := FastcgiAuthOptions{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *FastcgiAuthOptions) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FastcgiAuthOptions) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *FastcgiAuthOptions) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *FastcgiAuthOptions) SetRole(v string) {
	o.Role = &v
}

// GetEnable returns the Enable field value
func (o *FastcgiAuthOptions) GetEnable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enable
}

// GetEnableOk returns a tuple with the Enable field value
// and a boolean to check if the value has been set.
func (o *FastcgiAuthOptions) GetEnableOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enable, true
}

// SetEnable sets field value
func (o *FastcgiAuthOptions) SetEnable(v bool) {
	o.Enable = v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *FastcgiAuthOptions) GetParams() map[string]string {
	if o == nil || o.Params == nil {
		var ret map[string]string
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FastcgiAuthOptions) GetParamsOk() (*map[string]string, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *FastcgiAuthOptions) HasParams() bool {
	if o != nil && o.Params != nil {
		return true
	}

	return false
}

// SetParams gets a reference to the given map[string]string and assigns it to the Params field.
func (o *FastcgiAuthOptions) SetParams(v map[string]string) {
	o.Params = &v
}

// GetServer returns the Server field value
func (o *FastcgiAuthOptions) GetServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *FastcgiAuthOptions) GetServerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *FastcgiAuthOptions) SetServer(v string) {
	o.Server = v
}

// GetBackupServers returns the BackupServers field value if set, zero value otherwise.
func (o *FastcgiAuthOptions) GetBackupServers() []string {
	if o == nil || o.BackupServers == nil {
		var ret []string
		return ret
	}
	return *o.BackupServers
}

// GetBackupServersOk returns a tuple with the BackupServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FastcgiAuthOptions) GetBackupServersOk() (*[]string, bool) {
	if o == nil || o.BackupServers == nil {
		return nil, false
	}
	return o.BackupServers, true
}

// HasBackupServers returns a boolean if a field has been set.
func (o *FastcgiAuthOptions) HasBackupServers() bool {
	if o != nil && o.BackupServers != nil {
		return true
	}

	return false
}

// SetBackupServers gets a reference to the given []string and assigns it to the BackupServers field.
func (o *FastcgiAuthOptions) SetBackupServers(v []string) {
	o.BackupServers = &v
}

func (o FastcgiAuthOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if true {
		toSerialize["enable"] = o.Enable
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	if true {
		toSerialize["server"] = o.Server
	}
	if o.BackupServers != nil {
		toSerialize["backupServers"] = o.BackupServers
	}
	return json.Marshal(toSerialize)
}

type NullableFastcgiAuthOptions struct {
	value *FastcgiAuthOptions
	isSet bool
}

func (v NullableFastcgiAuthOptions) Get() *FastcgiAuthOptions {
	return v.value
}

func (v *NullableFastcgiAuthOptions) Set(val *FastcgiAuthOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableFastcgiAuthOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableFastcgiAuthOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFastcgiAuthOptions(val *FastcgiAuthOptions) *NullableFastcgiAuthOptions {
	return &NullableFastcgiAuthOptions{value: val, isSet: true}
}

func (v NullableFastcgiAuthOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFastcgiAuthOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


