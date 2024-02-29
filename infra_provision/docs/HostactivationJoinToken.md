# HostactivationJoinToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**LastUsedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**JoinTokenJoinTokenStatus**](JoinTokenJoinTokenStatus.md) |  | [optional] [default to JOINTOKENJOINTOKENSTATUS_UNKNOWN]
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**TokenId** | Pointer to **string** | first half of the token. | [optional] [readonly] 
**UseCounter** | Pointer to **int64** |  | [optional] [readonly] 

## Methods

### NewHostactivationJoinToken

`func NewHostactivationJoinToken() *HostactivationJoinToken`

NewHostactivationJoinToken instantiates a new HostactivationJoinToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostactivationJoinTokenWithDefaults

`func NewHostactivationJoinTokenWithDefaults() *HostactivationJoinToken`

NewHostactivationJoinTokenWithDefaults instantiates a new HostactivationJoinToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedAt

`func (o *HostactivationJoinToken) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *HostactivationJoinToken) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *HostactivationJoinToken) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *HostactivationJoinToken) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *HostactivationJoinToken) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HostactivationJoinToken) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HostactivationJoinToken) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HostactivationJoinToken) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiresAt

`func (o *HostactivationJoinToken) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *HostactivationJoinToken) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *HostactivationJoinToken) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *HostactivationJoinToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *HostactivationJoinToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostactivationJoinToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostactivationJoinToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HostactivationJoinToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *HostactivationJoinToken) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *HostactivationJoinToken) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *HostactivationJoinToken) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *HostactivationJoinToken) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetName

`func (o *HostactivationJoinToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HostactivationJoinToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HostactivationJoinToken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HostactivationJoinToken) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *HostactivationJoinToken) GetStatus() JoinTokenJoinTokenStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HostactivationJoinToken) GetStatusOk() (*JoinTokenJoinTokenStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HostactivationJoinToken) SetStatus(v JoinTokenJoinTokenStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HostactivationJoinToken) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *HostactivationJoinToken) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HostactivationJoinToken) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HostactivationJoinToken) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *HostactivationJoinToken) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTokenId

`func (o *HostactivationJoinToken) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *HostactivationJoinToken) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *HostactivationJoinToken) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *HostactivationJoinToken) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetUseCounter

`func (o *HostactivationJoinToken) GetUseCounter() int64`

GetUseCounter returns the UseCounter field if non-nil, zero value otherwise.

### GetUseCounterOk

`func (o *HostactivationJoinToken) GetUseCounterOk() (*int64, bool)`

GetUseCounterOk returns a tuple with the UseCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCounter

`func (o *HostactivationJoinToken) SetUseCounter(v int64)`

SetUseCounter sets UseCounter field to given value.

### HasUseCounter

`func (o *HostactivationJoinToken) HasUseCounter() bool`

HasUseCounter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


