# DdidnsrickettsSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountScheduleId** | Pointer to **string** | Account Schedule ID. | [optional] [readonly] 
**Accounts** | Pointer to [**[]DdidnsrickettsAccount**](DdidnsrickettsAccount.md) |  | [optional] 
**CloudCredentialId** | **string** | Cloud Credential ID. | 
**CreatedAt** | Pointer to **time.Time** | Timestamp when the object has been created. | [optional] [readonly] 
**CredentialConfig** | Pointer to [**DdidnsrickettsCredentialConfig**](DdidnsrickettsCredentialConfig.md) | Credential configuration. Ex.: &#39;{    \&quot;access_identifier\&quot;: \&quot;arn:aws:iam::1234:role/access_for_discovery\&quot;,    \&quot;region\&quot;: \&quot;us-east-1\&quot;,    \&quot;enclave\&quot;: \&quot;commercial/gov\&quot;  }&#39;. | [optional] 
**DeletedAt** | Pointer to **time.Time** | Timestamp when the object has been deleted. | [optional] [readonly] 
**Id** | Pointer to **string** | Auto-generated unique source config ID. Format BloxID. | [optional] [readonly] 
**RestrictedToAccounts** | Pointer to **[]string** | Provider account IDs such as accountID/ SubscriptionID to be restricted for a given source_config. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp when the object has been updated. | [optional] [readonly] 

## Methods

### NewDdidnsrickettsSourceConfig

`func NewDdidnsrickettsSourceConfig(cloudCredentialId string, ) *DdidnsrickettsSourceConfig`

NewDdidnsrickettsSourceConfig instantiates a new DdidnsrickettsSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdidnsrickettsSourceConfigWithDefaults

`func NewDdidnsrickettsSourceConfigWithDefaults() *DdidnsrickettsSourceConfig`

NewDdidnsrickettsSourceConfigWithDefaults instantiates a new DdidnsrickettsSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountScheduleId

`func (o *DdidnsrickettsSourceConfig) GetAccountScheduleId() string`

GetAccountScheduleId returns the AccountScheduleId field if non-nil, zero value otherwise.

### GetAccountScheduleIdOk

`func (o *DdidnsrickettsSourceConfig) GetAccountScheduleIdOk() (*string, bool)`

GetAccountScheduleIdOk returns a tuple with the AccountScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountScheduleId

`func (o *DdidnsrickettsSourceConfig) SetAccountScheduleId(v string)`

SetAccountScheduleId sets AccountScheduleId field to given value.

### HasAccountScheduleId

`func (o *DdidnsrickettsSourceConfig) HasAccountScheduleId() bool`

HasAccountScheduleId returns a boolean if a field has been set.

### GetAccounts

`func (o *DdidnsrickettsSourceConfig) GetAccounts() []DdidnsrickettsAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *DdidnsrickettsSourceConfig) GetAccountsOk() (*[]DdidnsrickettsAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *DdidnsrickettsSourceConfig) SetAccounts(v []DdidnsrickettsAccount)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *DdidnsrickettsSourceConfig) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetCloudCredentialId

`func (o *DdidnsrickettsSourceConfig) GetCloudCredentialId() string`

GetCloudCredentialId returns the CloudCredentialId field if non-nil, zero value otherwise.

### GetCloudCredentialIdOk

`func (o *DdidnsrickettsSourceConfig) GetCloudCredentialIdOk() (*string, bool)`

GetCloudCredentialIdOk returns a tuple with the CloudCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentialId

`func (o *DdidnsrickettsSourceConfig) SetCloudCredentialId(v string)`

SetCloudCredentialId sets CloudCredentialId field to given value.


### GetCreatedAt

`func (o *DdidnsrickettsSourceConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DdidnsrickettsSourceConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DdidnsrickettsSourceConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DdidnsrickettsSourceConfig) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCredentialConfig

`func (o *DdidnsrickettsSourceConfig) GetCredentialConfig() DdidnsrickettsCredentialConfig`

GetCredentialConfig returns the CredentialConfig field if non-nil, zero value otherwise.

### GetCredentialConfigOk

`func (o *DdidnsrickettsSourceConfig) GetCredentialConfigOk() (*DdidnsrickettsCredentialConfig, bool)`

GetCredentialConfigOk returns a tuple with the CredentialConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialConfig

`func (o *DdidnsrickettsSourceConfig) SetCredentialConfig(v DdidnsrickettsCredentialConfig)`

SetCredentialConfig sets CredentialConfig field to given value.

### HasCredentialConfig

`func (o *DdidnsrickettsSourceConfig) HasCredentialConfig() bool`

HasCredentialConfig returns a boolean if a field has been set.

### GetDeletedAt

`func (o *DdidnsrickettsSourceConfig) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DdidnsrickettsSourceConfig) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DdidnsrickettsSourceConfig) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DdidnsrickettsSourceConfig) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *DdidnsrickettsSourceConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DdidnsrickettsSourceConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DdidnsrickettsSourceConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DdidnsrickettsSourceConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRestrictedToAccounts

`func (o *DdidnsrickettsSourceConfig) GetRestrictedToAccounts() []string`

GetRestrictedToAccounts returns the RestrictedToAccounts field if non-nil, zero value otherwise.

### GetRestrictedToAccountsOk

`func (o *DdidnsrickettsSourceConfig) GetRestrictedToAccountsOk() (*[]string, bool)`

GetRestrictedToAccountsOk returns a tuple with the RestrictedToAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedToAccounts

`func (o *DdidnsrickettsSourceConfig) SetRestrictedToAccounts(v []string)`

SetRestrictedToAccounts sets RestrictedToAccounts field to given value.

### HasRestrictedToAccounts

`func (o *DdidnsrickettsSourceConfig) HasRestrictedToAccounts() bool`

HasRestrictedToAccounts returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DdidnsrickettsSourceConfig) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DdidnsrickettsSourceConfig) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DdidnsrickettsSourceConfig) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DdidnsrickettsSourceConfig) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


