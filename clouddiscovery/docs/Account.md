# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompositeStatus** | Pointer to **string** |  | [optional] 
**CompositeStatusMessage** | Pointer to **string** | Status message of the sync operation. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp when the object has been created. | [optional] [readonly] 
**DeletedAt** | Pointer to **time.Time** | Timestamp when the object has been deleted. | [optional] [readonly] 
**DhcpServerId** | Pointer to **string** |  | [optional] [readonly] 
**DnsServerId** | Pointer to **string** | DNS Server ID. | [optional] [readonly] 
**Id** | Pointer to **string** | Auto-generated unique source account ID. Format BloxID. | [optional] [readonly] 
**LastSuccessfulSync** | Pointer to **time.Time** | Last successful sync timestamp. | [optional] [readonly] 
**LastSync** | Pointer to **time.Time** | Last sync timestamp. | [optional] [readonly] 
**Name** | **string** | Name of the source account. | 
**ParentId** | Pointer to **string** | Parent ID. | [optional] 
**PercentComplete** | Pointer to **int32** | Sync progress as a percentage. | [optional] [readonly] 
**ProviderAccountId** | Pointer to **string** |  | [optional] 
**ScheduleId** | Pointer to **string** | Schedule ID. | [optional] [readonly] 
**State** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the sync operation. | [optional] [readonly] 
**StatusMessage** | Pointer to **string** | Status message of the sync operation. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp when the object has been updated. | [optional] [readonly] 

## Methods

### NewAccount

`func NewAccount(name string, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompositeStatus

`func (o *Account) GetCompositeStatus() string`

GetCompositeStatus returns the CompositeStatus field if non-nil, zero value otherwise.

### GetCompositeStatusOk

`func (o *Account) GetCompositeStatusOk() (*string, bool)`

GetCompositeStatusOk returns a tuple with the CompositeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeStatus

`func (o *Account) SetCompositeStatus(v string)`

SetCompositeStatus sets CompositeStatus field to given value.

### HasCompositeStatus

`func (o *Account) HasCompositeStatus() bool`

HasCompositeStatus returns a boolean if a field has been set.

### GetCompositeStatusMessage

`func (o *Account) GetCompositeStatusMessage() string`

GetCompositeStatusMessage returns the CompositeStatusMessage field if non-nil, zero value otherwise.

### GetCompositeStatusMessageOk

`func (o *Account) GetCompositeStatusMessageOk() (*string, bool)`

GetCompositeStatusMessageOk returns a tuple with the CompositeStatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeStatusMessage

`func (o *Account) SetCompositeStatusMessage(v string)`

SetCompositeStatusMessage sets CompositeStatusMessage field to given value.

### HasCompositeStatusMessage

`func (o *Account) HasCompositeStatusMessage() bool`

HasCompositeStatusMessage returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Account) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Account) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Account) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Account) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Account) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Account) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Account) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Account) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDhcpServerId

`func (o *Account) GetDhcpServerId() string`

GetDhcpServerId returns the DhcpServerId field if non-nil, zero value otherwise.

### GetDhcpServerIdOk

`func (o *Account) GetDhcpServerIdOk() (*string, bool)`

GetDhcpServerIdOk returns a tuple with the DhcpServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServerId

`func (o *Account) SetDhcpServerId(v string)`

SetDhcpServerId sets DhcpServerId field to given value.

### HasDhcpServerId

`func (o *Account) HasDhcpServerId() bool`

HasDhcpServerId returns a boolean if a field has been set.

### GetDnsServerId

`func (o *Account) GetDnsServerId() string`

GetDnsServerId returns the DnsServerId field if non-nil, zero value otherwise.

### GetDnsServerIdOk

`func (o *Account) GetDnsServerIdOk() (*string, bool)`

GetDnsServerIdOk returns a tuple with the DnsServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServerId

`func (o *Account) SetDnsServerId(v string)`

SetDnsServerId sets DnsServerId field to given value.

### HasDnsServerId

`func (o *Account) HasDnsServerId() bool`

HasDnsServerId returns a boolean if a field has been set.

### GetId

`func (o *Account) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Account) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastSuccessfulSync

`func (o *Account) GetLastSuccessfulSync() time.Time`

GetLastSuccessfulSync returns the LastSuccessfulSync field if non-nil, zero value otherwise.

### GetLastSuccessfulSyncOk

`func (o *Account) GetLastSuccessfulSyncOk() (*time.Time, bool)`

GetLastSuccessfulSyncOk returns a tuple with the LastSuccessfulSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulSync

`func (o *Account) SetLastSuccessfulSync(v time.Time)`

SetLastSuccessfulSync sets LastSuccessfulSync field to given value.

### HasLastSuccessfulSync

`func (o *Account) HasLastSuccessfulSync() bool`

HasLastSuccessfulSync returns a boolean if a field has been set.

### GetLastSync

`func (o *Account) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *Account) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *Account) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *Account) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetName

`func (o *Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Account) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *Account) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Account) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Account) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Account) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPercentComplete

`func (o *Account) GetPercentComplete() int32`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *Account) GetPercentCompleteOk() (*int32, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *Account) SetPercentComplete(v int32)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *Account) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### GetProviderAccountId

`func (o *Account) GetProviderAccountId() string`

GetProviderAccountId returns the ProviderAccountId field if non-nil, zero value otherwise.

### GetProviderAccountIdOk

`func (o *Account) GetProviderAccountIdOk() (*string, bool)`

GetProviderAccountIdOk returns a tuple with the ProviderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccountId

`func (o *Account) SetProviderAccountId(v string)`

SetProviderAccountId sets ProviderAccountId field to given value.

### HasProviderAccountId

`func (o *Account) HasProviderAccountId() bool`

HasProviderAccountId returns a boolean if a field has been set.

### GetScheduleId

`func (o *Account) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *Account) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *Account) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *Account) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### GetState

`func (o *Account) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Account) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Account) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Account) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *Account) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Account) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Account) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Account) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *Account) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *Account) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *Account) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *Account) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Account) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Account) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Account) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Account) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


