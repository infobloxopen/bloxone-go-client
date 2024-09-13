# SubAccountListRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | Pointer to **string** |  | [optional] 
**CredentialId** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to **string** | atlas.api.field_selection | [optional] 
**ProviderCredentialsConfig** | Pointer to [**SubAccountProvCredConfig**](SubAccountProvCredConfig.md) |  | [optional] 
**ProviderType** | Pointer to **string** |  | [optional] 

## Methods

### NewSubAccountListRequestV2

`func NewSubAccountListRequestV2() *SubAccountListRequestV2`

NewSubAccountListRequestV2 instantiates a new SubAccountListRequestV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubAccountListRequestV2WithDefaults

`func NewSubAccountListRequestV2WithDefaults() *SubAccountListRequestV2`

NewSubAccountListRequestV2WithDefaults instantiates a new SubAccountListRequestV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *SubAccountListRequestV2) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *SubAccountListRequestV2) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *SubAccountListRequestV2) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *SubAccountListRequestV2) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetCredentialId

`func (o *SubAccountListRequestV2) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *SubAccountListRequestV2) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *SubAccountListRequestV2) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.

### HasCredentialId

`func (o *SubAccountListRequestV2) HasCredentialId() bool`

HasCredentialId returns a boolean if a field has been set.

### GetFields

`func (o *SubAccountListRequestV2) GetFields() string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SubAccountListRequestV2) GetFieldsOk() (*string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SubAccountListRequestV2) SetFields(v string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *SubAccountListRequestV2) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetProviderCredentialsConfig

`func (o *SubAccountListRequestV2) GetProviderCredentialsConfig() SubAccountProvCredConfig`

GetProviderCredentialsConfig returns the ProviderCredentialsConfig field if non-nil, zero value otherwise.

### GetProviderCredentialsConfigOk

`func (o *SubAccountListRequestV2) GetProviderCredentialsConfigOk() (*SubAccountProvCredConfig, bool)`

GetProviderCredentialsConfigOk returns a tuple with the ProviderCredentialsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCredentialsConfig

`func (o *SubAccountListRequestV2) SetProviderCredentialsConfig(v SubAccountProvCredConfig)`

SetProviderCredentialsConfig sets ProviderCredentialsConfig field to given value.

### HasProviderCredentialsConfig

`func (o *SubAccountListRequestV2) HasProviderCredentialsConfig() bool`

HasProviderCredentialsConfig returns a boolean if a field has been set.

### GetProviderType

`func (o *SubAccountListRequestV2) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *SubAccountListRequestV2) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *SubAccountListRequestV2) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *SubAccountListRequestV2) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


