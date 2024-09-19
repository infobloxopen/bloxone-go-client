# DiscoveryConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountPreference** | **string** | Account preference. For ex.: single, multiple, auto-discover-multiple. | 
**AdditionalConfig** | Pointer to [**AdditionalConfig**](AdditionalConfig.md) | Additional configuration. Ex.: &#39;{    \&quot;excluded_object_types\&quot;: [],    \&quot;exclusion_account_list\&quot;: [],    \&quot;zone_forwarding\&quot;: \&quot;true\&quot; or \&quot;false\&quot; }&#39;. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp when the object has been created. | [optional] [readonly] 
**CredentialPreference** | Pointer to [**CredentialPreference**](CredentialPreference.md) | Credential preference. Ex.: &#39;{    \&quot;type\&quot;: \&quot;static\&quot; or \&quot;delegated\&quot;,    \&quot;access_identifier_type\&quot;: \&quot;role_arn\&quot; or \&quot;tenant_id\&quot; or \&quot;project_id\&quot;  }&#39;. | [optional] 
**DeletedAt** | Pointer to **time.Time** | Timestamp when the object has been deleted. | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the discovery config. Optional. | [optional] 
**DesiredState** | Pointer to **string** | Desired state. Default is \&quot;enabled\&quot;. | [optional] 
**DestinationTypesEnabled** | Pointer to **[]string** | Destinations types enabled: Ex.: DNS, IPAM and ACCOUNT. | [optional] 
**Destinations** | Pointer to [**[]Destination**](Destination.md) | Destinations. | [optional] 
**Id** | Pointer to **string** | Auto-generated unique discovery config ID. Format BloxID. | [optional] [readonly] 
**LastSync** | Pointer to **time.Time** | Last sync timestamp. | [optional] [readonly] 
**Name** | **string** | Name of the discovery config. | 
**ProviderType** | **string** | Provider type. Ex.: Amazon Web Services, Google Cloud Platform, Microsoft Azure. | 
**SourceConfigs** | Pointer to [**[]SourceConfig**](SourceConfig.md) | Source configs. | [optional] 
**Status** | Pointer to **string** | Status of the sync operation. In single account case, Its the combined status of account &amp; all the destinations statuses In auto discover case, Its the status of the account discovery only. | [optional] [readonly] 
**StatusMessage** | Pointer to **string** | Aggregate status message of the sync operation. | [optional] [readonly] 
**SyncInterval** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Tagging specifics. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp when the object has been updated. | [optional] [readonly] 

## Methods

### NewDiscoveryConfig

`func NewDiscoveryConfig(accountPreference string, name string, providerType string, ) *DiscoveryConfig`

NewDiscoveryConfig instantiates a new DiscoveryConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryConfigWithDefaults

`func NewDiscoveryConfigWithDefaults() *DiscoveryConfig`

NewDiscoveryConfigWithDefaults instantiates a new DiscoveryConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountPreference

`func (o *DiscoveryConfig) GetAccountPreference() string`

GetAccountPreference returns the AccountPreference field if non-nil, zero value otherwise.

### GetAccountPreferenceOk

`func (o *DiscoveryConfig) GetAccountPreferenceOk() (*string, bool)`

GetAccountPreferenceOk returns a tuple with the AccountPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPreference

`func (o *DiscoveryConfig) SetAccountPreference(v string)`

SetAccountPreference sets AccountPreference field to given value.


### GetAdditionalConfig

`func (o *DiscoveryConfig) GetAdditionalConfig() AdditionalConfig`

GetAdditionalConfig returns the AdditionalConfig field if non-nil, zero value otherwise.

### GetAdditionalConfigOk

`func (o *DiscoveryConfig) GetAdditionalConfigOk() (*AdditionalConfig, bool)`

GetAdditionalConfigOk returns a tuple with the AdditionalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalConfig

`func (o *DiscoveryConfig) SetAdditionalConfig(v AdditionalConfig)`

SetAdditionalConfig sets AdditionalConfig field to given value.

### HasAdditionalConfig

`func (o *DiscoveryConfig) HasAdditionalConfig() bool`

HasAdditionalConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DiscoveryConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DiscoveryConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DiscoveryConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DiscoveryConfig) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCredentialPreference

`func (o *DiscoveryConfig) GetCredentialPreference() CredentialPreference`

GetCredentialPreference returns the CredentialPreference field if non-nil, zero value otherwise.

### GetCredentialPreferenceOk

`func (o *DiscoveryConfig) GetCredentialPreferenceOk() (*CredentialPreference, bool)`

GetCredentialPreferenceOk returns a tuple with the CredentialPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialPreference

`func (o *DiscoveryConfig) SetCredentialPreference(v CredentialPreference)`

SetCredentialPreference sets CredentialPreference field to given value.

### HasCredentialPreference

`func (o *DiscoveryConfig) HasCredentialPreference() bool`

HasCredentialPreference returns a boolean if a field has been set.

### GetDeletedAt

`func (o *DiscoveryConfig) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DiscoveryConfig) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DiscoveryConfig) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DiscoveryConfig) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *DiscoveryConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiscoveryConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiscoveryConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DiscoveryConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDesiredState

`func (o *DiscoveryConfig) GetDesiredState() string`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *DiscoveryConfig) GetDesiredStateOk() (*string, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *DiscoveryConfig) SetDesiredState(v string)`

SetDesiredState sets DesiredState field to given value.

### HasDesiredState

`func (o *DiscoveryConfig) HasDesiredState() bool`

HasDesiredState returns a boolean if a field has been set.

### GetDestinationTypesEnabled

`func (o *DiscoveryConfig) GetDestinationTypesEnabled() []string`

GetDestinationTypesEnabled returns the DestinationTypesEnabled field if non-nil, zero value otherwise.

### GetDestinationTypesEnabledOk

`func (o *DiscoveryConfig) GetDestinationTypesEnabledOk() (*[]string, bool)`

GetDestinationTypesEnabledOk returns a tuple with the DestinationTypesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTypesEnabled

`func (o *DiscoveryConfig) SetDestinationTypesEnabled(v []string)`

SetDestinationTypesEnabled sets DestinationTypesEnabled field to given value.

### HasDestinationTypesEnabled

`func (o *DiscoveryConfig) HasDestinationTypesEnabled() bool`

HasDestinationTypesEnabled returns a boolean if a field has been set.

### GetDestinations

`func (o *DiscoveryConfig) GetDestinations() []Destination`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *DiscoveryConfig) GetDestinationsOk() (*[]Destination, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *DiscoveryConfig) SetDestinations(v []Destination)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *DiscoveryConfig) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetId

`func (o *DiscoveryConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiscoveryConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiscoveryConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DiscoveryConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastSync

`func (o *DiscoveryConfig) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *DiscoveryConfig) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *DiscoveryConfig) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *DiscoveryConfig) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetName

`func (o *DiscoveryConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiscoveryConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiscoveryConfig) SetName(v string)`

SetName sets Name field to given value.


### GetProviderType

`func (o *DiscoveryConfig) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *DiscoveryConfig) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *DiscoveryConfig) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### GetSourceConfigs

`func (o *DiscoveryConfig) GetSourceConfigs() []SourceConfig`

GetSourceConfigs returns the SourceConfigs field if non-nil, zero value otherwise.

### GetSourceConfigsOk

`func (o *DiscoveryConfig) GetSourceConfigsOk() (*[]SourceConfig, bool)`

GetSourceConfigsOk returns a tuple with the SourceConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConfigs

`func (o *DiscoveryConfig) SetSourceConfigs(v []SourceConfig)`

SetSourceConfigs sets SourceConfigs field to given value.

### HasSourceConfigs

`func (o *DiscoveryConfig) HasSourceConfigs() bool`

HasSourceConfigs returns a boolean if a field has been set.

### GetStatus

`func (o *DiscoveryConfig) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DiscoveryConfig) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DiscoveryConfig) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DiscoveryConfig) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *DiscoveryConfig) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *DiscoveryConfig) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *DiscoveryConfig) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *DiscoveryConfig) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetSyncInterval

`func (o *DiscoveryConfig) GetSyncInterval() string`

GetSyncInterval returns the SyncInterval field if non-nil, zero value otherwise.

### GetSyncIntervalOk

`func (o *DiscoveryConfig) GetSyncIntervalOk() (*string, bool)`

GetSyncIntervalOk returns a tuple with the SyncInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncInterval

`func (o *DiscoveryConfig) SetSyncInterval(v string)`

SetSyncInterval sets SyncInterval field to given value.

### HasSyncInterval

`func (o *DiscoveryConfig) HasSyncInterval() bool`

HasSyncInterval returns a boolean if a field has been set.

### GetTags

`func (o *DiscoveryConfig) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DiscoveryConfig) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DiscoveryConfig) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *DiscoveryConfig) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DiscoveryConfig) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DiscoveryConfig) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DiscoveryConfig) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DiscoveryConfig) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


