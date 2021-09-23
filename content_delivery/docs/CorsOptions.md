# CorsOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessControlAllowMethods** | Pointer to **[]string** | &#x60;Access-Control-Allow-Methods&#x60; method list. Enter a list of allowed methods or leave empty to use value of &#x60;Access-Control-Request-Method&#x60; header. | [optional] 
**AccessControlAllowOrigin** | Pointer to [**CorsOptionsAccessControlAllowOrigin**](CorsOptionsAccessControlAllowOrigin.md) |  | [optional] 
**AccessControlMaxAge** | Pointer to **int32** | &#x60;Access-Control-Max-Age&#x60; header value. Enter number of seconds for which client may cache preflight response or leave empty to use default of one day. | [optional] 
**EnablePreflight** | Pointer to **bool** | Enables preflight requests. | [optional] 
**AccessControlExposeHeaders** | Pointer to **[]string** | &#x60;Access-Control-Expose-Headers&#x60; header value. | [optional] 
**AccessControlAllowCredentials** | Pointer to **bool** | &#x60;Access-Control-Allow-Credentials&#x60; header value. | [optional] 
**AccessControlAllowHeaders** | Pointer to **[]string** | &#x60;Access-Control-Allow-Headers&#x60; header list. Enter a list of allowed headers or leave empty to use value of &#x60;Access-Control-Request-Headers&#x60; header. | [optional] 

## Methods

### NewCorsOptions

`func NewCorsOptions() *CorsOptions`

NewCorsOptions instantiates a new CorsOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorsOptionsWithDefaults

`func NewCorsOptionsWithDefaults() *CorsOptions`

NewCorsOptionsWithDefaults instantiates a new CorsOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessControlAllowMethods

`func (o *CorsOptions) GetAccessControlAllowMethods() []string`

GetAccessControlAllowMethods returns the AccessControlAllowMethods field if non-nil, zero value otherwise.

### GetAccessControlAllowMethodsOk

`func (o *CorsOptions) GetAccessControlAllowMethodsOk() (*[]string, bool)`

GetAccessControlAllowMethodsOk returns a tuple with the AccessControlAllowMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlAllowMethods

`func (o *CorsOptions) SetAccessControlAllowMethods(v []string)`

SetAccessControlAllowMethods sets AccessControlAllowMethods field to given value.

### HasAccessControlAllowMethods

`func (o *CorsOptions) HasAccessControlAllowMethods() bool`

HasAccessControlAllowMethods returns a boolean if a field has been set.

### GetAccessControlAllowOrigin

`func (o *CorsOptions) GetAccessControlAllowOrigin() CorsOptionsAccessControlAllowOrigin`

GetAccessControlAllowOrigin returns the AccessControlAllowOrigin field if non-nil, zero value otherwise.

### GetAccessControlAllowOriginOk

`func (o *CorsOptions) GetAccessControlAllowOriginOk() (*CorsOptionsAccessControlAllowOrigin, bool)`

GetAccessControlAllowOriginOk returns a tuple with the AccessControlAllowOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlAllowOrigin

`func (o *CorsOptions) SetAccessControlAllowOrigin(v CorsOptionsAccessControlAllowOrigin)`

SetAccessControlAllowOrigin sets AccessControlAllowOrigin field to given value.

### HasAccessControlAllowOrigin

`func (o *CorsOptions) HasAccessControlAllowOrigin() bool`

HasAccessControlAllowOrigin returns a boolean if a field has been set.

### GetAccessControlMaxAge

`func (o *CorsOptions) GetAccessControlMaxAge() int32`

GetAccessControlMaxAge returns the AccessControlMaxAge field if non-nil, zero value otherwise.

### GetAccessControlMaxAgeOk

`func (o *CorsOptions) GetAccessControlMaxAgeOk() (*int32, bool)`

GetAccessControlMaxAgeOk returns a tuple with the AccessControlMaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlMaxAge

`func (o *CorsOptions) SetAccessControlMaxAge(v int32)`

SetAccessControlMaxAge sets AccessControlMaxAge field to given value.

### HasAccessControlMaxAge

`func (o *CorsOptions) HasAccessControlMaxAge() bool`

HasAccessControlMaxAge returns a boolean if a field has been set.

### GetEnablePreflight

`func (o *CorsOptions) GetEnablePreflight() bool`

GetEnablePreflight returns the EnablePreflight field if non-nil, zero value otherwise.

### GetEnablePreflightOk

`func (o *CorsOptions) GetEnablePreflightOk() (*bool, bool)`

GetEnablePreflightOk returns a tuple with the EnablePreflight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePreflight

`func (o *CorsOptions) SetEnablePreflight(v bool)`

SetEnablePreflight sets EnablePreflight field to given value.

### HasEnablePreflight

`func (o *CorsOptions) HasEnablePreflight() bool`

HasEnablePreflight returns a boolean if a field has been set.

### GetAccessControlExposeHeaders

`func (o *CorsOptions) GetAccessControlExposeHeaders() []string`

GetAccessControlExposeHeaders returns the AccessControlExposeHeaders field if non-nil, zero value otherwise.

### GetAccessControlExposeHeadersOk

`func (o *CorsOptions) GetAccessControlExposeHeadersOk() (*[]string, bool)`

GetAccessControlExposeHeadersOk returns a tuple with the AccessControlExposeHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlExposeHeaders

`func (o *CorsOptions) SetAccessControlExposeHeaders(v []string)`

SetAccessControlExposeHeaders sets AccessControlExposeHeaders field to given value.

### HasAccessControlExposeHeaders

`func (o *CorsOptions) HasAccessControlExposeHeaders() bool`

HasAccessControlExposeHeaders returns a boolean if a field has been set.

### GetAccessControlAllowCredentials

`func (o *CorsOptions) GetAccessControlAllowCredentials() bool`

GetAccessControlAllowCredentials returns the AccessControlAllowCredentials field if non-nil, zero value otherwise.

### GetAccessControlAllowCredentialsOk

`func (o *CorsOptions) GetAccessControlAllowCredentialsOk() (*bool, bool)`

GetAccessControlAllowCredentialsOk returns a tuple with the AccessControlAllowCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlAllowCredentials

`func (o *CorsOptions) SetAccessControlAllowCredentials(v bool)`

SetAccessControlAllowCredentials sets AccessControlAllowCredentials field to given value.

### HasAccessControlAllowCredentials

`func (o *CorsOptions) HasAccessControlAllowCredentials() bool`

HasAccessControlAllowCredentials returns a boolean if a field has been set.

### GetAccessControlAllowHeaders

`func (o *CorsOptions) GetAccessControlAllowHeaders() []string`

GetAccessControlAllowHeaders returns the AccessControlAllowHeaders field if non-nil, zero value otherwise.

### GetAccessControlAllowHeadersOk

`func (o *CorsOptions) GetAccessControlAllowHeadersOk() (*[]string, bool)`

GetAccessControlAllowHeadersOk returns a tuple with the AccessControlAllowHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlAllowHeaders

`func (o *CorsOptions) SetAccessControlAllowHeaders(v []string)`

SetAccessControlAllowHeaders sets AccessControlAllowHeaders field to given value.

### HasAccessControlAllowHeaders

`func (o *CorsOptions) HasAccessControlAllowHeaders() bool`

HasAccessControlAllowHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


