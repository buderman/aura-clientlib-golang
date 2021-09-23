# SharedSecretSetSecrets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secrets** | **map[string]string** | One or more unique key/index string pairs to be added as members of the secret set. | 

## Methods

### NewSharedSecretSetSecrets

`func NewSharedSecretSetSecrets(secrets map[string]string, ) *SharedSecretSetSecrets`

NewSharedSecretSetSecrets instantiates a new SharedSecretSetSecrets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedSecretSetSecretsWithDefaults

`func NewSharedSecretSetSecretsWithDefaults() *SharedSecretSetSecrets`

NewSharedSecretSetSecretsWithDefaults instantiates a new SharedSecretSetSecrets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecrets

`func (o *SharedSecretSetSecrets) GetSecrets() map[string]string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *SharedSecretSetSecrets) GetSecretsOk() (*map[string]string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *SharedSecretSetSecrets) SetSecrets(v map[string]string)`

SetSecrets sets Secrets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


