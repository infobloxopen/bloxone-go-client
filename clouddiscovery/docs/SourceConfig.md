# SourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountScheduleId** | Pointer to **string** | Account Schedule ID. | [optional] [readonly] 
**Accounts** | Pointer to [**[]Account**](Account.md) |  | [optional] 
**CloudCredentialId** | Pointer to **string** | Cloud Credential ID. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp when the object has been created. | [optional] [readonly] 
**CredentialConfig** | Pointer to [**CredentialConfig**](CredentialConfig.md) | Credential configuration. Ex.: &#39;{    \&quot;access_identifier\&quot;: \&quot;arn:aws:iam::1234:role/access_for_discovery\&quot;,    \&quot;region\&quot;: \&quot;us-east-1\&quot;,    \&quot;enclave\&quot;: \&quot;commercial/gov\&quot;  }&#39;. | [optional] 
**DeletedAt** | Pointer to **time.Time** | Timestamp when the object has been deleted. | [optional] [readonly] 
**Id** | Pointer to **string** | Auto-generated unique source config ID. Format BloxID. | [optional] [readonly] 
**RestrictedToAccounts** | Pointer to **[]string** | Provider account IDs such as accountID/ SubscriptionID to be restricted for a given source_config. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp when the object has been updated. | [optional] [readonly] 

## Methods

### NewSourceConfig

`func NewSourceConfig() *SourceConfig`

NewSourceConfig instantiates a new SourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceConfigWithDefaults

`func NewSourceConfigWithDefaults() *SourceConfig`

NewSourceConfigWithDefaults instantiates a new SourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountScheduleId

`func (o *SourceConfig) GetAccountScheduleId() string`

GetAccountScheduleId returns the AccountScheduleId field if non-nil, zero value otherwise.

### GetAccountScheduleIdOk

`func (o *SourceConfig) GetAccountScheduleIdOk() (*string, bool)`

GetAccountScheduleIdOk returns a tuple with the AccountScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountScheduleId

`func (o *SourceConfig) SetAccountScheduleId(v string)`

SetAccountScheduleId sets AccountScheduleId field to given value.

### HasAccountScheduleId

`func (o *SourceConfig) HasAccountScheduleId() bool`

HasAccountScheduleId returns a boolean if a field has been set.

### GetAccounts

`func (o *SourceConfig) GetAccounts() []Account`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *SourceConfig) GetAccountsOk() (*[]Account, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *SourceConfig) SetAccounts(v []Account)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *SourceConfig) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetCloudCredentialId

`func (o *SourceConfig) GetCloudCredentialId() string`

GetCloudCredentialId returns the CloudCredentialId field if non-nil, zero value otherwise.

### GetCloudCredentialIdOk

`func (o *SourceConfig) GetCloudCredentialIdOk() (*string, bool)`

GetCloudCredentialIdOk returns a tuple with the CloudCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentialId

`func (o *SourceConfig) SetCloudCredentialId(v string)`

SetCloudCredentialId sets CloudCredentialId field to given value.

### HasCloudCredentialId

`func (o *SourceConfig) HasCloudCredentialId() bool`

HasCloudCredentialId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SourceConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SourceConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SourceConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SourceConfig) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCredentialConfig

`func (o *SourceConfig) GetCredentialConfig() CredentialConfig`

GetCredentialConfig returns the CredentialConfig field if non-nil, zero value otherwise.

### GetCredentialConfigOk

`func (o *SourceConfig) GetCredentialConfigOk() (*CredentialConfig, bool)`

GetCredentialConfigOk returns a tuple with the CredentialConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialConfig

`func (o *SourceConfig) SetCredentialConfig(v CredentialConfig)`

SetCredentialConfig sets CredentialConfig field to given value.

### HasCredentialConfig

`func (o *SourceConfig) HasCredentialConfig() bool`

HasCredentialConfig returns a boolean if a field has been set.

### GetDeletedAt

`func (o *SourceConfig) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *SourceConfig) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *SourceConfig) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *SourceConfig) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *SourceConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SourceConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SourceConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SourceConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRestrictedToAccounts

`func (o *SourceConfig) GetRestrictedToAccounts() []string`

GetRestrictedToAccounts returns the RestrictedToAccounts field if non-nil, zero value otherwise.

### GetRestrictedToAccountsOk

`func (o *SourceConfig) GetRestrictedToAccountsOk() (*[]string, bool)`

GetRestrictedToAccountsOk returns a tuple with the RestrictedToAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedToAccounts

`func (o *SourceConfig) SetRestrictedToAccounts(v []string)`

SetRestrictedToAccounts sets RestrictedToAccounts field to given value.

### HasRestrictedToAccounts

`func (o *SourceConfig) HasRestrictedToAccounts() bool`

HasRestrictedToAccounts returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SourceConfig) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SourceConfig) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SourceConfig) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SourceConfig) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


