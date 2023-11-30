# KerberosKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | Encryption algorithm of the key in accordance with RFC 3961. | [optional] [readonly] 
**Comment** | Pointer to **string** | The description for Kerberos key. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**Domain** | Pointer to **string** | Kerberos realm of the principal. | [optional] [readonly] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Principal** | Pointer to **string** | Kerberos principal associated with key. | [optional] [readonly] 
**Tags** | Pointer to **map[string]interface{}** | The tags for the Kerberos key in JSON format. | [optional] 
**UploadedAt** | Pointer to **string** | Upload time for the key. | [optional] [readonly] 
**Version** | Pointer to **int64** | The version number (KVNO) of the key. | [optional] [readonly] 

## Methods

### NewKerberosKey

`func NewKerberosKey() *KerberosKey`

NewKerberosKey instantiates a new KerberosKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKerberosKeyWithDefaults

`func NewKerberosKeyWithDefaults() *KerberosKey`

NewKerberosKeyWithDefaults instantiates a new KerberosKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *KerberosKey) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *KerberosKey) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *KerberosKey) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *KerberosKey) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetComment

`func (o *KerberosKey) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *KerberosKey) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *KerberosKey) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *KerberosKey) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDomain

`func (o *KerberosKey) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *KerberosKey) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *KerberosKey) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *KerberosKey) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetId

`func (o *KerberosKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KerberosKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KerberosKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KerberosKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrincipal

`func (o *KerberosKey) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *KerberosKey) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *KerberosKey) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *KerberosKey) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetTags

`func (o *KerberosKey) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *KerberosKey) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *KerberosKey) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *KerberosKey) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUploadedAt

`func (o *KerberosKey) GetUploadedAt() string`

GetUploadedAt returns the UploadedAt field if non-nil, zero value otherwise.

### GetUploadedAtOk

`func (o *KerberosKey) GetUploadedAtOk() (*string, bool)`

GetUploadedAtOk returns a tuple with the UploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedAt

`func (o *KerberosKey) SetUploadedAt(v string)`

SetUploadedAt sets UploadedAt field to given value.

### HasUploadedAt

`func (o *KerberosKey) HasUploadedAt() bool`

HasUploadedAt returns a boolean if a field has been set.

### GetVersion

`func (o *KerberosKey) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KerberosKey) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KerberosKey) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KerberosKey) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


