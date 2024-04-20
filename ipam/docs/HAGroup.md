# HAGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnycastConfigId** | Pointer to **string** | The resource identifier. | [optional] 
**Comment** | Pointer to **string** | The description for the HA group. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**Hosts** | [**[]HAGroupHost**](HAGroupHost.md) | The list of hosts. | 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**IpSpace** | Pointer to **string** | The resource identifier. | [optional] 
**Mode** | Pointer to **string** | The mode of the HA group.  Valid values are: * _active-active_: Both on-prem hosts remain active. * _active-passive_: One on-prem host remains active and one remains passive. When the active on-prem host is down, the passive on-prem host takes over. * _advanced-active-passive_: One on-prem host may be part of multiple HA groups. When the active on-prem host is down, the passive on-prem host takes over. | [optional] 
**Name** | **string** | The name of the HA group. Must contain 1 to 256 characters. Can include UTF-8. | 
**Status** | Pointer to **string** | Status of the HA group. This field is set when the _collect_stats_ is set to _true_ in the _GET_ _/dhcp/ha_group_ request. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | The tags for the HA group. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 

## Methods

### NewHAGroup

`func NewHAGroup(hosts []HAGroupHost, name string, ) *HAGroup`

NewHAGroup instantiates a new HAGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHAGroupWithDefaults

`func NewHAGroupWithDefaults() *HAGroup`

NewHAGroupWithDefaults instantiates a new HAGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnycastConfigId

`func (o *HAGroup) GetAnycastConfigId() string`

GetAnycastConfigId returns the AnycastConfigId field if non-nil, zero value otherwise.

### GetAnycastConfigIdOk

`func (o *HAGroup) GetAnycastConfigIdOk() (*string, bool)`

GetAnycastConfigIdOk returns a tuple with the AnycastConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnycastConfigId

`func (o *HAGroup) SetAnycastConfigId(v string)`

SetAnycastConfigId sets AnycastConfigId field to given value.

### HasAnycastConfigId

`func (o *HAGroup) HasAnycastConfigId() bool`

HasAnycastConfigId returns a boolean if a field has been set.

### GetComment

`func (o *HAGroup) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *HAGroup) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *HAGroup) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *HAGroup) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *HAGroup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HAGroup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HAGroup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HAGroup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHosts

`func (o *HAGroup) GetHosts() []HAGroupHost`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *HAGroup) GetHostsOk() (*[]HAGroupHost, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *HAGroup) SetHosts(v []HAGroupHost)`

SetHosts sets Hosts field to given value.


### GetId

`func (o *HAGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HAGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HAGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HAGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpSpace

`func (o *HAGroup) GetIpSpace() string`

GetIpSpace returns the IpSpace field if non-nil, zero value otherwise.

### GetIpSpaceOk

`func (o *HAGroup) GetIpSpaceOk() (*string, bool)`

GetIpSpaceOk returns a tuple with the IpSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSpace

`func (o *HAGroup) SetIpSpace(v string)`

SetIpSpace sets IpSpace field to given value.

### HasIpSpace

`func (o *HAGroup) HasIpSpace() bool`

HasIpSpace returns a boolean if a field has been set.

### GetMode

`func (o *HAGroup) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *HAGroup) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *HAGroup) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *HAGroup) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *HAGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HAGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HAGroup) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *HAGroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HAGroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HAGroup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HAGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *HAGroup) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HAGroup) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HAGroup) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *HAGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *HAGroup) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HAGroup) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HAGroup) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HAGroup) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


