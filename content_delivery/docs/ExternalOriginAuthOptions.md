# ExternalOriginAuthOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretKey** | **string** | External origin secret access key. | 
**Region** | **string** | External origin service region, such as &#x60;us-east-1&#x60; and &#x60;eu-central-1&#x60;. | 
**SignedHeaders** | Pointer to **[]string** | List of request headers included in authorization signature. | [optional] 
**AccessKey** | **string** | External Oorigin access key ID. | 
**Service** | **string** | External origin service. The only available service is &#x60;AMAZON_AWS_S3&#x60;. | 

## Methods

### NewExternalOriginAuthOptions

`func NewExternalOriginAuthOptions(secretKey string, region string, accessKey string, service string, ) *ExternalOriginAuthOptions`

NewExternalOriginAuthOptions instantiates a new ExternalOriginAuthOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalOriginAuthOptionsWithDefaults

`func NewExternalOriginAuthOptionsWithDefaults() *ExternalOriginAuthOptions`

NewExternalOriginAuthOptionsWithDefaults instantiates a new ExternalOriginAuthOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretKey

`func (o *ExternalOriginAuthOptions) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *ExternalOriginAuthOptions) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *ExternalOriginAuthOptions) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetRegion

`func (o *ExternalOriginAuthOptions) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ExternalOriginAuthOptions) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ExternalOriginAuthOptions) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetSignedHeaders

`func (o *ExternalOriginAuthOptions) GetSignedHeaders() []string`

GetSignedHeaders returns the SignedHeaders field if non-nil, zero value otherwise.

### GetSignedHeadersOk

`func (o *ExternalOriginAuthOptions) GetSignedHeadersOk() (*[]string, bool)`

GetSignedHeadersOk returns a tuple with the SignedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedHeaders

`func (o *ExternalOriginAuthOptions) SetSignedHeaders(v []string)`

SetSignedHeaders sets SignedHeaders field to given value.

### HasSignedHeaders

`func (o *ExternalOriginAuthOptions) HasSignedHeaders() bool`

HasSignedHeaders returns a boolean if a field has been set.

### GetAccessKey

`func (o *ExternalOriginAuthOptions) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *ExternalOriginAuthOptions) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *ExternalOriginAuthOptions) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetService

`func (o *ExternalOriginAuthOptions) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ExternalOriginAuthOptions) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ExternalOriginAuthOptions) SetService(v string)`

SetService sets Service field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


