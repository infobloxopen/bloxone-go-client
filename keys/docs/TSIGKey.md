# TSIGKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | The TSIG key algorithm.  Valid values are: * _hmac_sha1_ * _hmac_sha224_ * _hmac_sha256_ * _hmac_sha384_ * _hmac_sha512_  Defaults to _hmac_sha256_. | [optional] 
**Comment** | Pointer to **string** | The description for the TSIG key. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Name** | **string** | The TSIG key name in the absolute domain name format. | 
**ProtocolName** | Pointer to **string** | The TSIG key name supplied during a create/update operation that is converted to canonical form in punycode. | [optional] [readonly] 
**Secret** | **string** | The TSIG key secret as a Base64 encoded string. | 
**Tags** | Pointer to **map[string]interface{}** | The tags for the TSIG key in JSON format. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 

## Methods

### NewTSIGKey

`func NewTSIGKey(name string, secret string, ) *TSIGKey`

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

### GetCreatedAt

`func (o *TSIGKey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TSIGKey) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TSIGKey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TSIGKey) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *TSIGKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TSIGKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TSIGKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TSIGKey) HasId() bool`

HasId returns a boolean if a field has been set.

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


### GetTags

`func (o *TSIGKey) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TSIGKey) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TSIGKey) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *TSIGKey) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TSIGKey) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TSIGKey) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TSIGKey) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TSIGKey) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


