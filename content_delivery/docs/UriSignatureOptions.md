# UriSignatureOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedSecretIndexParameter** | **string** | The name of the request parameter which contains the index for the shared secret. | 
**InputParameterPattern** | **string** | Specifies the set of input parameters and their order. The input parameter pattern can use parts of the incoming request, for example: &#x60;$url_signature_shared_secret,#uri?sid&#x3D;$arg_sid&amp;e&#x3D;$arg_e&#x60;. To represent an expiration as a query argument, or with a cookie incoming URL, use &#x60;$uri_sign_expire&#x60;. | 
**Algorithm** | **string** | Signature calculation algorithm, either &#x60;SHA256&#x60;, &#x60;SHA384&#x60;, &#x60;SHA512&#x60;, &#x60;HMAC_MD5&#x60;, &#x60;HMAC_SHA1&#x60;, &#x60;HMAC_SHA256&#x60;, &#x60;HMAC_SHA384&#x60;, &#x60;HMAC_SHA512&#x60;, &#x60;MD5&#x60;, or &#x60;SHA1&#x60;. | 
**ProtectedPathPattern** | Pointer to **string** | The name of the capture path in the URI filter above used as a cookie path string to define the protected file names or path. | [optional] 
**SharedSecretSetId** | **int32** |  | 
**TokenEncoding** | **string** | The encoding used for the token parameter, either &#x60;BASE64&#x60;, &#x60;HEX&#x60;, or &#x60;URL_SAFE_BASE64&#x60;. | 
**UriExpirationParameter** | **string** | The name of the request parameter which contains the expiration time for the URI being examined. | 
**ProtectedSubFiles** | Pointer to **bool** | Use cookies derived from master query arguments to protect files matching URI filter with no query arguments, such as child &#x60;.m3u8&#x60; and &#x60;.ts&#x60;. | [optional] 
**InputParameterCapitalization** | Pointer to **string** | Specifies the capitalization used for transforming the input parameter before signature calculation. Either &#x60;LOWERCASE&#x60;, &#x60;UPPERCASE&#x60;, or &#x60;NO_CHANGE&#x60;. | [optional] 
**TokenParameter** | **string** | The name of the request parameter which contains the signature token. | 

## Methods

### NewUriSignatureOptions

`func NewUriSignatureOptions(sharedSecretIndexParameter string, inputParameterPattern string, algorithm string, sharedSecretSetId int32, tokenEncoding string, uriExpirationParameter string, tokenParameter string, ) *UriSignatureOptions`

NewUriSignatureOptions instantiates a new UriSignatureOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUriSignatureOptionsWithDefaults

`func NewUriSignatureOptionsWithDefaults() *UriSignatureOptions`

NewUriSignatureOptionsWithDefaults instantiates a new UriSignatureOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedSecretIndexParameter

`func (o *UriSignatureOptions) GetSharedSecretIndexParameter() string`

GetSharedSecretIndexParameter returns the SharedSecretIndexParameter field if non-nil, zero value otherwise.

### GetSharedSecretIndexParameterOk

`func (o *UriSignatureOptions) GetSharedSecretIndexParameterOk() (*string, bool)`

GetSharedSecretIndexParameterOk returns a tuple with the SharedSecretIndexParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecretIndexParameter

`func (o *UriSignatureOptions) SetSharedSecretIndexParameter(v string)`

SetSharedSecretIndexParameter sets SharedSecretIndexParameter field to given value.


### GetInputParameterPattern

`func (o *UriSignatureOptions) GetInputParameterPattern() string`

GetInputParameterPattern returns the InputParameterPattern field if non-nil, zero value otherwise.

### GetInputParameterPatternOk

`func (o *UriSignatureOptions) GetInputParameterPatternOk() (*string, bool)`

GetInputParameterPatternOk returns a tuple with the InputParameterPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameterPattern

`func (o *UriSignatureOptions) SetInputParameterPattern(v string)`

SetInputParameterPattern sets InputParameterPattern field to given value.


### GetAlgorithm

`func (o *UriSignatureOptions) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *UriSignatureOptions) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *UriSignatureOptions) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetProtectedPathPattern

`func (o *UriSignatureOptions) GetProtectedPathPattern() string`

GetProtectedPathPattern returns the ProtectedPathPattern field if non-nil, zero value otherwise.

### GetProtectedPathPatternOk

`func (o *UriSignatureOptions) GetProtectedPathPatternOk() (*string, bool)`

GetProtectedPathPatternOk returns a tuple with the ProtectedPathPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedPathPattern

`func (o *UriSignatureOptions) SetProtectedPathPattern(v string)`

SetProtectedPathPattern sets ProtectedPathPattern field to given value.

### HasProtectedPathPattern

`func (o *UriSignatureOptions) HasProtectedPathPattern() bool`

HasProtectedPathPattern returns a boolean if a field has been set.

### GetSharedSecretSetId

`func (o *UriSignatureOptions) GetSharedSecretSetId() int32`

GetSharedSecretSetId returns the SharedSecretSetId field if non-nil, zero value otherwise.

### GetSharedSecretSetIdOk

`func (o *UriSignatureOptions) GetSharedSecretSetIdOk() (*int32, bool)`

GetSharedSecretSetIdOk returns a tuple with the SharedSecretSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecretSetId

`func (o *UriSignatureOptions) SetSharedSecretSetId(v int32)`

SetSharedSecretSetId sets SharedSecretSetId field to given value.


### GetTokenEncoding

`func (o *UriSignatureOptions) GetTokenEncoding() string`

GetTokenEncoding returns the TokenEncoding field if non-nil, zero value otherwise.

### GetTokenEncodingOk

`func (o *UriSignatureOptions) GetTokenEncodingOk() (*string, bool)`

GetTokenEncodingOk returns a tuple with the TokenEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEncoding

`func (o *UriSignatureOptions) SetTokenEncoding(v string)`

SetTokenEncoding sets TokenEncoding field to given value.


### GetUriExpirationParameter

`func (o *UriSignatureOptions) GetUriExpirationParameter() string`

GetUriExpirationParameter returns the UriExpirationParameter field if non-nil, zero value otherwise.

### GetUriExpirationParameterOk

`func (o *UriSignatureOptions) GetUriExpirationParameterOk() (*string, bool)`

GetUriExpirationParameterOk returns a tuple with the UriExpirationParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUriExpirationParameter

`func (o *UriSignatureOptions) SetUriExpirationParameter(v string)`

SetUriExpirationParameter sets UriExpirationParameter field to given value.


### GetProtectedSubFiles

`func (o *UriSignatureOptions) GetProtectedSubFiles() bool`

GetProtectedSubFiles returns the ProtectedSubFiles field if non-nil, zero value otherwise.

### GetProtectedSubFilesOk

`func (o *UriSignatureOptions) GetProtectedSubFilesOk() (*bool, bool)`

GetProtectedSubFilesOk returns a tuple with the ProtectedSubFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedSubFiles

`func (o *UriSignatureOptions) SetProtectedSubFiles(v bool)`

SetProtectedSubFiles sets ProtectedSubFiles field to given value.

### HasProtectedSubFiles

`func (o *UriSignatureOptions) HasProtectedSubFiles() bool`

HasProtectedSubFiles returns a boolean if a field has been set.

### GetInputParameterCapitalization

`func (o *UriSignatureOptions) GetInputParameterCapitalization() string`

GetInputParameterCapitalization returns the InputParameterCapitalization field if non-nil, zero value otherwise.

### GetInputParameterCapitalizationOk

`func (o *UriSignatureOptions) GetInputParameterCapitalizationOk() (*string, bool)`

GetInputParameterCapitalizationOk returns a tuple with the InputParameterCapitalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameterCapitalization

`func (o *UriSignatureOptions) SetInputParameterCapitalization(v string)`

SetInputParameterCapitalization sets InputParameterCapitalization field to given value.

### HasInputParameterCapitalization

`func (o *UriSignatureOptions) HasInputParameterCapitalization() bool`

HasInputParameterCapitalization returns a boolean if a field has been set.

### GetTokenParameter

`func (o *UriSignatureOptions) GetTokenParameter() string`

GetTokenParameter returns the TokenParameter field if non-nil, zero value otherwise.

### GetTokenParameterOk

`func (o *UriSignatureOptions) GetTokenParameterOk() (*string, bool)`

GetTokenParameterOk returns a tuple with the TokenParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenParameter

`func (o *UriSignatureOptions) SetTokenParameter(v string)`

SetTokenParameter sets TokenParameter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


