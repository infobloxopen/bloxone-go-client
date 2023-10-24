# ConfigKerberosKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | Encryption algorithm of the key in accordance with RFC 3961. | [optional] [readonly] 
**Domain** | Pointer to **string** | Kerberos realm of the principal. | [optional] [readonly] 
**Key** | **string** | The resource identifier. | 
**Principal** | Pointer to **string** | Kerberos principal associated with key. | [optional] [readonly] 
**UploadedAt** | Pointer to **string** | Upload time for the key. | [optional] [readonly] 
**Version** | Pointer to **int64** | The version number (KVNO) of the key. | [optional] [readonly] 

## Methods

### NewConfigKerberosKey

`func NewConfigKerberosKey(key string, ) *ConfigKerberosKey`

NewConfigKerberosKey instantiates a new ConfigKerberosKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigKerberosKeyWithDefaults

`func NewConfigKerberosKeyWithDefaults() *ConfigKerberosKey`

NewConfigKerberosKeyWithDefaults instantiates a new ConfigKerberosKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *ConfigKerberosKey) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *ConfigKerberosKey) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *ConfigKerberosKey) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *ConfigKerberosKey) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetDomain

`func (o *ConfigKerberosKey) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ConfigKerberosKey) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ConfigKerberosKey) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ConfigKerberosKey) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetKey

`func (o *ConfigKerberosKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ConfigKerberosKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ConfigKerberosKey) SetKey(v string)`

SetKey sets Key field to given value.


### GetPrincipal

`func (o *ConfigKerberosKey) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *ConfigKerberosKey) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *ConfigKerberosKey) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *ConfigKerberosKey) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetUploadedAt

`func (o *ConfigKerberosKey) GetUploadedAt() string`

GetUploadedAt returns the UploadedAt field if non-nil, zero value otherwise.

### GetUploadedAtOk

`func (o *ConfigKerberosKey) GetUploadedAtOk() (*string, bool)`

GetUploadedAtOk returns a tuple with the UploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedAt

`func (o *ConfigKerberosKey) SetUploadedAt(v string)`

SetUploadedAt sets UploadedAt field to given value.

### HasUploadedAt

`func (o *ConfigKerberosKey) HasUploadedAt() bool`

HasUploadedAt returns a boolean if a field has been set.

### GetVersion

`func (o *ConfigKerberosKey) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConfigKerberosKey) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConfigKerberosKey) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConfigKerberosKey) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


