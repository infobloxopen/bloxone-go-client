# DdidnsrickettsDiscoveryConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountPreference** | Pointer to **string** | Account preference. For ex.: single, multiple, auto-discover-multiple. | [optional] 
**AdditionalConfig** | Pointer to [**DdidnsrickettsAdditionalConfig**](DdidnsrickettsAdditionalConfig.md) | Additional configuration. Ex.: &#39;{    \&quot;excluded_object_types\&quot;: [],    \&quot;exclusion_account_list\&quot;: [],    \&quot;zone_forwarding\&quot;: \&quot;true\&quot; or \&quot;false\&quot; }&#39;. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp when the object has been created. | [optional] [readonly] 
**CredentialPreference** | Pointer to [**DdidnsrickettsCredentialPreference**](DdidnsrickettsCredentialPreference.md) | Credential preference. Ex.: &#39;{    \&quot;type\&quot;: \&quot;static\&quot; or \&quot;delegated\&quot;,    \&quot;access_identifier_type\&quot;: \&quot;role_arn\&quot; or \&quot;tenant_id\&quot; or \&quot;project_id\&quot;  }&#39;. | [optional] 
**DeletedAt** | Pointer to **time.Time** | Timestamp when the object has been deleted. | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the discovery config. Optional. | [optional] 
**DesiredState** | Pointer to **string** | Desired state. Default is \&quot;enabled\&quot;. | [optional] 
**DestinationTypesEnabled** | Pointer to **[]string** | Destinations types enabled: Ex.: DNS, IPAM and ACCOUNT. | [optional] 
**Destinations** | Pointer to [**[]DdidnsrickettsDestination**](DdidnsrickettsDestination.md) | Destinations. | [optional] 
**Id** | Pointer to **string** | Auto-generated unique discovery config ID. Format BloxID. | [optional] [readonly] 
**LastSync** | Pointer to **time.Time** | Last sync timestamp. | [optional] 
**Name** | **string** | Name of the discovery config. | 
**ProviderType** | Pointer to **string** | Provider type. Ex.: Amazon Web Services, Google Cloud Platform, Microsoft Azure. | [optional] 
**SourceConfigs** | Pointer to [**[]DdidnsrickettsSourceConfig**](DdidnsrickettsSourceConfig.md) | Source configs. | [optional] 
**Status** | Pointer to **string** | Status of the sync operation. In single account case, Its the combined status of account &amp; all the destinations statuses In auto discover case, Its the status of the account discovery only. | [optional] 
**StatusMessage** | Pointer to **string** | Aggregate status message of the sync operation. | [optional] 
**SyncInterval** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]map[string]interface{}** | Tagging specifics. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp when the object has been updated. | [optional] [readonly] 

## Methods

### NewDdidnsrickettsDiscoveryConfig

`func NewDdidnsrickettsDiscoveryConfig(name string, ) *DdidnsrickettsDiscoveryConfig`

NewDdidnsrickettsDiscoveryConfig instantiates a new DdidnsrickettsDiscoveryConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdidnsrickettsDiscoveryConfigWithDefaults

`func NewDdidnsrickettsDiscoveryConfigWithDefaults() *DdidnsrickettsDiscoveryConfig`

NewDdidnsrickettsDiscoveryConfigWithDefaults instantiates a new DdidnsrickettsDiscoveryConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountPreference

`func (o *DdidnsrickettsDiscoveryConfig) GetAccountPreference() string`

GetAccountPreference returns the AccountPreference field if non-nil, zero value otherwise.

### GetAccountPreferenceOk

`func (o *DdidnsrickettsDiscoveryConfig) GetAccountPreferenceOk() (*string, bool)`

GetAccountPreferenceOk returns a tuple with the AccountPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPreference

`func (o *DdidnsrickettsDiscoveryConfig) SetAccountPreference(v string)`

SetAccountPreference sets AccountPreference field to given value.

### HasAccountPreference

`func (o *DdidnsrickettsDiscoveryConfig) HasAccountPreference() bool`

HasAccountPreference returns a boolean if a field has been set.

### GetAdditionalConfig

`func (o *DdidnsrickettsDiscoveryConfig) GetAdditionalConfig() DdidnsrickettsAdditionalConfig`

GetAdditionalConfig returns the AdditionalConfig field if non-nil, zero value otherwise.

### GetAdditionalConfigOk

`func (o *DdidnsrickettsDiscoveryConfig) GetAdditionalConfigOk() (*DdidnsrickettsAdditionalConfig, bool)`

GetAdditionalConfigOk returns a tuple with the AdditionalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalConfig

`func (o *DdidnsrickettsDiscoveryConfig) SetAdditionalConfig(v DdidnsrickettsAdditionalConfig)`

SetAdditionalConfig sets AdditionalConfig field to given value.

### HasAdditionalConfig

`func (o *DdidnsrickettsDiscoveryConfig) HasAdditionalConfig() bool`

HasAdditionalConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DdidnsrickettsDiscoveryConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DdidnsrickettsDiscoveryConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DdidnsrickettsDiscoveryConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DdidnsrickettsDiscoveryConfig) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCredentialPreference

`func (o *DdidnsrickettsDiscoveryConfig) GetCredentialPreference() DdidnsrickettsCredentialPreference`

GetCredentialPreference returns the CredentialPreference field if non-nil, zero value otherwise.

### GetCredentialPreferenceOk

`func (o *DdidnsrickettsDiscoveryConfig) GetCredentialPreferenceOk() (*DdidnsrickettsCredentialPreference, bool)`

GetCredentialPreferenceOk returns a tuple with the CredentialPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialPreference

`func (o *DdidnsrickettsDiscoveryConfig) SetCredentialPreference(v DdidnsrickettsCredentialPreference)`

SetCredentialPreference sets CredentialPreference field to given value.

### HasCredentialPreference

`func (o *DdidnsrickettsDiscoveryConfig) HasCredentialPreference() bool`

HasCredentialPreference returns a boolean if a field has been set.

### GetDeletedAt

`func (o *DdidnsrickettsDiscoveryConfig) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DdidnsrickettsDiscoveryConfig) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DdidnsrickettsDiscoveryConfig) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DdidnsrickettsDiscoveryConfig) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *DdidnsrickettsDiscoveryConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DdidnsrickettsDiscoveryConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DdidnsrickettsDiscoveryConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DdidnsrickettsDiscoveryConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDesiredState

`func (o *DdidnsrickettsDiscoveryConfig) GetDesiredState() string`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *DdidnsrickettsDiscoveryConfig) GetDesiredStateOk() (*string, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *DdidnsrickettsDiscoveryConfig) SetDesiredState(v string)`

SetDesiredState sets DesiredState field to given value.

### HasDesiredState

`func (o *DdidnsrickettsDiscoveryConfig) HasDesiredState() bool`

HasDesiredState returns a boolean if a field has been set.

### GetDestinationTypesEnabled

`func (o *DdidnsrickettsDiscoveryConfig) GetDestinationTypesEnabled() []string`

GetDestinationTypesEnabled returns the DestinationTypesEnabled field if non-nil, zero value otherwise.

### GetDestinationTypesEnabledOk

`func (o *DdidnsrickettsDiscoveryConfig) GetDestinationTypesEnabledOk() (*[]string, bool)`

GetDestinationTypesEnabledOk returns a tuple with the DestinationTypesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTypesEnabled

`func (o *DdidnsrickettsDiscoveryConfig) SetDestinationTypesEnabled(v []string)`

SetDestinationTypesEnabled sets DestinationTypesEnabled field to given value.

### HasDestinationTypesEnabled

`func (o *DdidnsrickettsDiscoveryConfig) HasDestinationTypesEnabled() bool`

HasDestinationTypesEnabled returns a boolean if a field has been set.

### GetDestinations

`func (o *DdidnsrickettsDiscoveryConfig) GetDestinations() []DdidnsrickettsDestination`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *DdidnsrickettsDiscoveryConfig) GetDestinationsOk() (*[]DdidnsrickettsDestination, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *DdidnsrickettsDiscoveryConfig) SetDestinations(v []DdidnsrickettsDestination)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *DdidnsrickettsDiscoveryConfig) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetId

`func (o *DdidnsrickettsDiscoveryConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DdidnsrickettsDiscoveryConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DdidnsrickettsDiscoveryConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DdidnsrickettsDiscoveryConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastSync

`func (o *DdidnsrickettsDiscoveryConfig) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *DdidnsrickettsDiscoveryConfig) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *DdidnsrickettsDiscoveryConfig) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *DdidnsrickettsDiscoveryConfig) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetName

`func (o *DdidnsrickettsDiscoveryConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DdidnsrickettsDiscoveryConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DdidnsrickettsDiscoveryConfig) SetName(v string)`

SetName sets Name field to given value.


### GetProviderType

`func (o *DdidnsrickettsDiscoveryConfig) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *DdidnsrickettsDiscoveryConfig) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *DdidnsrickettsDiscoveryConfig) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *DdidnsrickettsDiscoveryConfig) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetSourceConfigs

`func (o *DdidnsrickettsDiscoveryConfig) GetSourceConfigs() []DdidnsrickettsSourceConfig`

GetSourceConfigs returns the SourceConfigs field if non-nil, zero value otherwise.

### GetSourceConfigsOk

`func (o *DdidnsrickettsDiscoveryConfig) GetSourceConfigsOk() (*[]DdidnsrickettsSourceConfig, bool)`

GetSourceConfigsOk returns a tuple with the SourceConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConfigs

`func (o *DdidnsrickettsDiscoveryConfig) SetSourceConfigs(v []DdidnsrickettsSourceConfig)`

SetSourceConfigs sets SourceConfigs field to given value.

### HasSourceConfigs

`func (o *DdidnsrickettsDiscoveryConfig) HasSourceConfigs() bool`

HasSourceConfigs returns a boolean if a field has been set.

### GetStatus

`func (o *DdidnsrickettsDiscoveryConfig) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DdidnsrickettsDiscoveryConfig) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DdidnsrickettsDiscoveryConfig) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DdidnsrickettsDiscoveryConfig) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *DdidnsrickettsDiscoveryConfig) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *DdidnsrickettsDiscoveryConfig) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *DdidnsrickettsDiscoveryConfig) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *DdidnsrickettsDiscoveryConfig) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetSyncInterval

`func (o *DdidnsrickettsDiscoveryConfig) GetSyncInterval() string`

GetSyncInterval returns the SyncInterval field if non-nil, zero value otherwise.

### GetSyncIntervalOk

`func (o *DdidnsrickettsDiscoveryConfig) GetSyncIntervalOk() (*string, bool)`

GetSyncIntervalOk returns a tuple with the SyncInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncInterval

`func (o *DdidnsrickettsDiscoveryConfig) SetSyncInterval(v string)`

SetSyncInterval sets SyncInterval field to given value.

### HasSyncInterval

`func (o *DdidnsrickettsDiscoveryConfig) HasSyncInterval() bool`

HasSyncInterval returns a boolean if a field has been set.

### GetTags

`func (o *DdidnsrickettsDiscoveryConfig) GetTags() map[string]map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DdidnsrickettsDiscoveryConfig) GetTagsOk() (*map[string]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DdidnsrickettsDiscoveryConfig) SetTags(v map[string]map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *DdidnsrickettsDiscoveryConfig) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DdidnsrickettsDiscoveryConfig) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DdidnsrickettsDiscoveryConfig) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DdidnsrickettsDiscoveryConfig) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DdidnsrickettsDiscoveryConfig) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


