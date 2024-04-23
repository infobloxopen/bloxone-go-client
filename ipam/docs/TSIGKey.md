# TSIGKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | TSIG key algorithm.  Valid values are:  * _hmac_sha256_  * _hmac_sha1_  * _hmac_sha224_  * _hmac_sha384_  * _hmac_sha512_ | [optional] 
**Comment** | Pointer to **string** | The description for the TSIG key. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**Key** | **string** | The resource identifier. | 
**Name** | Pointer to **string** | The TSIG key name, FQDN. | [optional] 
**ProtocolName** | Pointer to **string** | The TSIG key name in punycode. | [optional] [readonly] 
**Secret** | Pointer to **string** | The TSIG key secret, base64 string. | [optional] 

## Methods

### NewTSIGKey

`func NewTSIGKey(key string, ) *TSIGKey`

NewTSIGKey instantiates a new TSIGKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSIGKeyWithDefaults

`func NewTSIGKeyWithDefaults() *TSIGKey`

NewTSIGKeyWithDefaults instantiates a new TSIGKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *TSIGKey) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *TSIGKey) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *TSIGKey) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *TSIGKey) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetComment

`func (o *TSIGKey) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TSIGKey) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TSIGKey) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TSIGKey) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetKey

`func (o *TSIGKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TSIGKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TSIGKey) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *TSIGKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TSIGKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TSIGKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TSIGKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProtocolName

`func (o *TSIGKey) GetProtocolName() string`

GetProtocolName returns the ProtocolName field if non-nil, zero value otherwise.

### GetProtocolNameOk

`func (o *TSIGKey) GetProtocolNameOk() (*string, bool)`

GetProtocolNameOk returns a tuple with the ProtocolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolName

`func (o *TSIGKey) SetProtocolName(v string)`

SetProtocolName sets ProtocolName field to given value.

### HasProtocolName

`func (o *TSIGKey) HasProtocolName() bool`

HasProtocolName returns a boolean if a field has been set.

### GetSecret

`func (o *TSIGKey) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *TSIGKey) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *TSIGKey) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *TSIGKey) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


