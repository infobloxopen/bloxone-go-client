# JoinToken

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

### NewJoinToken

`func NewJoinToken() *JoinToken`

NewJoinToken instantiates a new JoinToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJoinTokenWithDefaults

`func NewJoinTokenWithDefaults() *JoinToken`

NewJoinTokenWithDefaults instantiates a new JoinToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedAt

`func (o *JoinToken) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *JoinToken) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *JoinToken) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *JoinToken) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *JoinToken) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JoinToken) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JoinToken) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JoinToken) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiresAt

`func (o *JoinToken) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *JoinToken) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *JoinToken) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *JoinToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *JoinToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JoinToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JoinToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JoinToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *JoinToken) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *JoinToken) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *JoinToken) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *JoinToken) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetName

`func (o *JoinToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JoinToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JoinToken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JoinToken) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *JoinToken) GetStatus() JoinTokenJoinTokenStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JoinToken) GetStatusOk() (*JoinTokenJoinTokenStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JoinToken) SetStatus(v JoinTokenJoinTokenStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JoinToken) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *JoinToken) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *JoinToken) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *JoinToken) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *JoinToken) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTokenId

`func (o *JoinToken) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *JoinToken) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *JoinToken) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *JoinToken) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetUseCounter

`func (o *JoinToken) GetUseCounter() int64`

GetUseCounter returns the UseCounter field if non-nil, zero value otherwise.

### GetUseCounterOk

`func (o *JoinToken) GetUseCounterOk() (*int64, bool)`

GetUseCounterOk returns a tuple with the UseCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCounter

`func (o *JoinToken) SetUseCounter(v int64)`

SetUseCounter sets UseCounter field to given value.

### HasUseCounter

`func (o *JoinToken) HasUseCounter() bool`

HasUseCounter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


