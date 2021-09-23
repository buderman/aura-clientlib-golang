# FastcgiAuthOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** | Role set for the FastCGI Server, either &#x60;AUTHORIZER&#x60; or &#x60;RESPONDER&#x60;. | [optional] 
**Enable** | **bool** | Enables FastCGI Authentication. | 
**Params** | Pointer to **map[string]string** | FastCGI parameters in map form. | [optional] 
**Server** | **string** | Destination FastCGI server server:port, maximum length is 262, 256 of which is for the hostname. | 
**BackupServers** | Pointer to **[]string** | List of backup destination FastCGI Servers, expressed as &#x60;server:port&#x60;. | [optional] 

## Methods

### NewFastcgiAuthOptions

`func NewFastcgiAuthOptions(enable bool, server string, ) *FastcgiAuthOptions`

NewFastcgiAuthOptions instantiates a new FastcgiAuthOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFastcgiAuthOptionsWithDefaults

`func NewFastcgiAuthOptionsWithDefaults() *FastcgiAuthOptions`

NewFastcgiAuthOptionsWithDefaults instantiates a new FastcgiAuthOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *FastcgiAuthOptions) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *FastcgiAuthOptions) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *FastcgiAuthOptions) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *FastcgiAuthOptions) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetEnable

`func (o *FastcgiAuthOptions) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *FastcgiAuthOptions) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *FastcgiAuthOptions) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetParams

`func (o *FastcgiAuthOptions) GetParams() map[string]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *FastcgiAuthOptions) GetParamsOk() (*map[string]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *FastcgiAuthOptions) SetParams(v map[string]string)`

SetParams sets Params field to given value.

### HasParams

`func (o *FastcgiAuthOptions) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetServer

`func (o *FastcgiAuthOptions) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *FastcgiAuthOptions) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *FastcgiAuthOptions) SetServer(v string)`

SetServer sets Server field to given value.


### GetBackupServers

`func (o *FastcgiAuthOptions) GetBackupServers() []string`

GetBackupServers returns the BackupServers field if non-nil, zero value otherwise.

### GetBackupServersOk

`func (o *FastcgiAuthOptions) GetBackupServersOk() (*[]string, bool)`

GetBackupServersOk returns a tuple with the BackupServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupServers

`func (o *FastcgiAuthOptions) SetBackupServers(v []string)`

SetBackupServers sets BackupServers field to given value.

### HasBackupServers

`func (o *FastcgiAuthOptions) HasBackupServers() bool`

HasBackupServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


