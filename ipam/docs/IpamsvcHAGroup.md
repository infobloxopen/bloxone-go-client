# IpamsvcHAGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnycastConfigId** | Pointer to **string** | The resource identifier. | [optional] 
**Comment** | Pointer to **string** | The description for the HA group. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**Hosts** | [**[]IpamsvcHAGroupHost**](IpamsvcHAGroupHost.md) | The list of hosts. | 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**IpSpace** | Pointer to **string** | The resource identifier. | [optional] 
**Mode** | Pointer to **string** | The mode of the HA group.  Valid values are: * _active-active_: Both on-prem hosts remain active. * _active-passive_: One on-prem host remains active and one remains passive. When the active on-prem host is down, the passive on-prem host takes over. * _advanced-active-passive_: One on-prem host may be part of multiple HA groups. When the active on-prem host is down, the passive on-prem host takes over. | [optional] 
**Name** | **string** | The name of the HA group. Must contain 1 to 256 characters. Can include UTF-8. | 
**Status** | Pointer to **string** | Status of the HA group. This field is set when the _collect_stats_ is set to _true_ in the _GET_ _/dhcp/ha_group_ request. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | The tags for the HA group. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 

## Methods

### NewIpamsvcHAGroup

`func NewIpamsvcHAGroup(hosts []IpamsvcHAGroupHost, name string, ) *IpamsvcHAGroup`

NewIpamsvcHAGroup instantiates a new IpamsvcHAGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcHAGroupWithDefaults

`func NewIpamsvcHAGroupWithDefaults() *IpamsvcHAGroup`

NewIpamsvcHAGroupWithDefaults instantiates a new IpamsvcHAGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnycastConfigId

`func (o *IpamsvcHAGroup) GetAnycastConfigId() string`

GetAnycastConfigId returns the AnycastConfigId field if non-nil, zero value otherwise.

### GetAnycastConfigIdOk

`func (o *IpamsvcHAGroup) GetAnycastConfigIdOk() (*string, bool)`

GetAnycastConfigIdOk returns a tuple with the AnycastConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnycastConfigId

`func (o *IpamsvcHAGroup) SetAnycastConfigId(v string)`

SetAnycastConfigId sets AnycastConfigId field to given value.

### HasAnycastConfigId

`func (o *IpamsvcHAGroup) HasAnycastConfigId() bool`

HasAnycastConfigId returns a boolean if a field has been set.

### GetComment

`func (o *IpamsvcHAGroup) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IpamsvcHAGroup) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IpamsvcHAGroup) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IpamsvcHAGroup) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IpamsvcHAGroup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IpamsvcHAGroup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IpamsvcHAGroup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IpamsvcHAGroup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHosts

`func (o *IpamsvcHAGroup) GetHosts() []IpamsvcHAGroupHost`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *IpamsvcHAGroup) GetHostsOk() (*[]IpamsvcHAGroupHost, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *IpamsvcHAGroup) SetHosts(v []IpamsvcHAGroupHost)`

SetHosts sets Hosts field to given value.


### GetId

`func (o *IpamsvcHAGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpamsvcHAGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpamsvcHAGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpamsvcHAGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpSpace

`func (o *IpamsvcHAGroup) GetIpSpace() string`

GetIpSpace returns the IpSpace field if non-nil, zero value otherwise.

### GetIpSpaceOk

`func (o *IpamsvcHAGroup) GetIpSpaceOk() (*string, bool)`

GetIpSpaceOk returns a tuple with the IpSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSpace

`func (o *IpamsvcHAGroup) SetIpSpace(v string)`

SetIpSpace sets IpSpace field to given value.

### HasIpSpace

`func (o *IpamsvcHAGroup) HasIpSpace() bool`

HasIpSpace returns a boolean if a field has been set.

### GetMode

`func (o *IpamsvcHAGroup) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *IpamsvcHAGroup) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *IpamsvcHAGroup) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *IpamsvcHAGroup) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *IpamsvcHAGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpamsvcHAGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpamsvcHAGroup) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *IpamsvcHAGroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IpamsvcHAGroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IpamsvcHAGroup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IpamsvcHAGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *IpamsvcHAGroup) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IpamsvcHAGroup) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IpamsvcHAGroup) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *IpamsvcHAGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IpamsvcHAGroup) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IpamsvcHAGroup) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IpamsvcHAGroup) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IpamsvcHAGroup) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


