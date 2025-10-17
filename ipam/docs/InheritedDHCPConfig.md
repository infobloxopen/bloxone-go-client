# InheritedDHCPConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbandonedReclaimTime** | Pointer to [**InheritanceInheritedUInt32**](InheritanceInheritedUInt32.md) | The inheritance configuration for _abandoned_reclaim_time_ field from _DHCPConfig_ object. | [optional] 
**AbandonedReclaimTimeV6** | Pointer to [**InheritanceInheritedUInt32**](InheritanceInheritedUInt32.md) | The inheritance configuration for _abandoned_reclaim_time_v6_ field from _DHCPConfig_ object. | [optional] 
**AllowUnknown** | Pointer to [**InheritanceInheritedBool**](InheritanceInheritedBool.md) | The inheritance configuration for _allow_unknown_ field from _DHCPConfig_ object. | [optional] 
**AllowUnknownV6** | Pointer to [**InheritanceInheritedBool**](InheritanceInheritedBool.md) | The inheritance configuration for _allow_unknown_v6_ field from _DHCPConfig_ object. | [optional] 
**AuthoritativeDhcp** | Pointer to [**InheritanceInheritedBool**](InheritanceInheritedBool.md) | The inheritance configuration for _authoritative_dhcp_ field from _DHCPConfig_ object. | [optional] 
**EchoClientId** | Pointer to [**InheritanceInheritedBool**](InheritanceInheritedBool.md) | The inheritance configuration for _echo_client_id_ field from _DHCPConfig_ object. | [optional] 
**Filters** | Pointer to [**InheritedDHCPConfigFilterList**](InheritedDHCPConfigFilterList.md) | The inheritance configuration for filters field from _DHCPConfig_ object. | [optional] 
**FiltersV6** | Pointer to [**InheritedDHCPConfigFilterList**](InheritedDHCPConfigFilterList.md) | The inheritance configuration for _filters_v6_ field from _DHCPConfig_ object. | [optional] 
**HoldReclaimedTime** | Pointer to [**InheritanceInheritedUInt32**](InheritanceInheritedUInt32.md) | The inheritance configuration for _hold_reclaimed_time_ field from _DHCPConfig_ object. | [optional] 
**HoldReclaimedTimeV6** | Pointer to [**InheritanceInheritedUInt32**](InheritanceInheritedUInt32.md) | The inheritance configuration for _hold_reclaimed_time_v6_ field from _DHCPConfig_ object. | [optional] 
**IgnoreClientUid** | Pointer to [**InheritanceInheritedBool**](InheritanceInheritedBool.md) | The inheritance configuration for _ignore_client_uid_ field from _DHCPConfig_ object. | [optional] 
**IgnoreList** | Pointer to [**InheritedDHCPConfigIgnoreItemList**](InheritedDHCPConfigIgnoreItemList.md) | The inheritance configuration for _ignore_list_ field from _DHCPConfig_ object. | [optional] 
**LeaseTime** | Pointer to [**InheritanceInheritedUInt32**](InheritanceInheritedUInt32.md) | The inheritance configuration for _lease_time_ field from _DHCPConfig_ object. | [optional] 
**LeaseTimeV6** | Pointer to [**InheritanceInheritedUInt32**](InheritanceInheritedUInt32.md) | The inheritance configuration for _lease_time_v6_ field from _DHCPConfig_ object. | [optional] 

## Methods

### NewInheritedDHCPConfig

`func NewInheritedDHCPConfig() *InheritedDHCPConfig`

NewInheritedDHCPConfig instantiates a new InheritedDHCPConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInheritedDHCPConfigWithDefaults

`func NewInheritedDHCPConfigWithDefaults() *InheritedDHCPConfig`

NewInheritedDHCPConfigWithDefaults instantiates a new InheritedDHCPConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbandonedReclaimTime

`func (o *InheritedDHCPConfig) GetAbandonedReclaimTime() InheritanceInheritedUInt32`

GetAbandonedReclaimTime returns the AbandonedReclaimTime field if non-nil, zero value otherwise.

### GetAbandonedReclaimTimeOk

`func (o *InheritedDHCPConfig) GetAbandonedReclaimTimeOk() (*InheritanceInheritedUInt32, bool)`

GetAbandonedReclaimTimeOk returns a tuple with the AbandonedReclaimTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandonedReclaimTime

`func (o *InheritedDHCPConfig) SetAbandonedReclaimTime(v InheritanceInheritedUInt32)`

SetAbandonedReclaimTime sets AbandonedReclaimTime field to given value.

### HasAbandonedReclaimTime

`func (o *InheritedDHCPConfig) HasAbandonedReclaimTime() bool`

HasAbandonedReclaimTime returns a boolean if a field has been set.

### GetAbandonedReclaimTimeV6

`func (o *InheritedDHCPConfig) GetAbandonedReclaimTimeV6() InheritanceInheritedUInt32`

GetAbandonedReclaimTimeV6 returns the AbandonedReclaimTimeV6 field if non-nil, zero value otherwise.

### GetAbandonedReclaimTimeV6Ok

`func (o *InheritedDHCPConfig) GetAbandonedReclaimTimeV6Ok() (*InheritanceInheritedUInt32, bool)`

GetAbandonedReclaimTimeV6Ok returns a tuple with the AbandonedReclaimTimeV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandonedReclaimTimeV6

`func (o *InheritedDHCPConfig) SetAbandonedReclaimTimeV6(v InheritanceInheritedUInt32)`

SetAbandonedReclaimTimeV6 sets AbandonedReclaimTimeV6 field to given value.

### HasAbandonedReclaimTimeV6

`func (o *InheritedDHCPConfig) HasAbandonedReclaimTimeV6() bool`

HasAbandonedReclaimTimeV6 returns a boolean if a field has been set.

### GetAllowUnknown

`func (o *InheritedDHCPConfig) GetAllowUnknown() InheritanceInheritedBool`

GetAllowUnknown returns the AllowUnknown field if non-nil, zero value otherwise.

### GetAllowUnknownOk

`func (o *InheritedDHCPConfig) GetAllowUnknownOk() (*InheritanceInheritedBool, bool)`

GetAllowUnknownOk returns a tuple with the AllowUnknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnknown

`func (o *InheritedDHCPConfig) SetAllowUnknown(v InheritanceInheritedBool)`

SetAllowUnknown sets AllowUnknown field to given value.

### HasAllowUnknown

`func (o *InheritedDHCPConfig) HasAllowUnknown() bool`

HasAllowUnknown returns a boolean if a field has been set.

### GetAllowUnknownV6

`func (o *InheritedDHCPConfig) GetAllowUnknownV6() InheritanceInheritedBool`

GetAllowUnknownV6 returns the AllowUnknownV6 field if non-nil, zero value otherwise.

### GetAllowUnknownV6Ok

`func (o *InheritedDHCPConfig) GetAllowUnknownV6Ok() (*InheritanceInheritedBool, bool)`

GetAllowUnknownV6Ok returns a tuple with the AllowUnknownV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnknownV6

`func (o *InheritedDHCPConfig) SetAllowUnknownV6(v InheritanceInheritedBool)`

SetAllowUnknownV6 sets AllowUnknownV6 field to given value.

### HasAllowUnknownV6

`func (o *InheritedDHCPConfig) HasAllowUnknownV6() bool`

HasAllowUnknownV6 returns a boolean if a field has been set.

### GetAuthoritativeDhcp

`func (o *InheritedDHCPConfig) GetAuthoritativeDhcp() InheritanceInheritedBool`

GetAuthoritativeDhcp returns the AuthoritativeDhcp field if non-nil, zero value otherwise.

### GetAuthoritativeDhcpOk

`func (o *InheritedDHCPConfig) GetAuthoritativeDhcpOk() (*InheritanceInheritedBool, bool)`

GetAuthoritativeDhcpOk returns a tuple with the AuthoritativeDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoritativeDhcp

`func (o *InheritedDHCPConfig) SetAuthoritativeDhcp(v InheritanceInheritedBool)`

SetAuthoritativeDhcp sets AuthoritativeDhcp field to given value.

### HasAuthoritativeDhcp

`func (o *InheritedDHCPConfig) HasAuthoritativeDhcp() bool`

HasAuthoritativeDhcp returns a boolean if a field has been set.

### GetEchoClientId

`func (o *InheritedDHCPConfig) GetEchoClientId() InheritanceInheritedBool`

GetEchoClientId returns the EchoClientId field if non-nil, zero value otherwise.

### GetEchoClientIdOk

`func (o *InheritedDHCPConfig) GetEchoClientIdOk() (*InheritanceInheritedBool, bool)`

GetEchoClientIdOk returns a tuple with the EchoClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEchoClientId

`func (o *InheritedDHCPConfig) SetEchoClientId(v InheritanceInheritedBool)`

SetEchoClientId sets EchoClientId field to given value.

### HasEchoClientId

`func (o *InheritedDHCPConfig) HasEchoClientId() bool`

HasEchoClientId returns a boolean if a field has been set.

### GetFilters

`func (o *InheritedDHCPConfig) GetFilters() InheritedDHCPConfigFilterList`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *InheritedDHCPConfig) GetFiltersOk() (*InheritedDHCPConfigFilterList, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *InheritedDHCPConfig) SetFilters(v InheritedDHCPConfigFilterList)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *InheritedDHCPConfig) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetFiltersV6

`func (o *InheritedDHCPConfig) GetFiltersV6() InheritedDHCPConfigFilterList`

GetFiltersV6 returns the FiltersV6 field if non-nil, zero value otherwise.

### GetFiltersV6Ok

`func (o *InheritedDHCPConfig) GetFiltersV6Ok() (*InheritedDHCPConfigFilterList, bool)`

GetFiltersV6Ok returns a tuple with the FiltersV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltersV6

`func (o *InheritedDHCPConfig) SetFiltersV6(v InheritedDHCPConfigFilterList)`

SetFiltersV6 sets FiltersV6 field to given value.

### HasFiltersV6

`func (o *InheritedDHCPConfig) HasFiltersV6() bool`

HasFiltersV6 returns a boolean if a field has been set.

### GetHoldReclaimedTime

`func (o *InheritedDHCPConfig) GetHoldReclaimedTime() InheritanceInheritedUInt32`

GetHoldReclaimedTime returns the HoldReclaimedTime field if non-nil, zero value otherwise.

### GetHoldReclaimedTimeOk

`func (o *InheritedDHCPConfig) GetHoldReclaimedTimeOk() (*InheritanceInheritedUInt32, bool)`

GetHoldReclaimedTimeOk returns a tuple with the HoldReclaimedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldReclaimedTime

`func (o *InheritedDHCPConfig) SetHoldReclaimedTime(v InheritanceInheritedUInt32)`

SetHoldReclaimedTime sets HoldReclaimedTime field to given value.

### HasHoldReclaimedTime

`func (o *InheritedDHCPConfig) HasHoldReclaimedTime() bool`

HasHoldReclaimedTime returns a boolean if a field has been set.

### GetHoldReclaimedTimeV6

`func (o *InheritedDHCPConfig) GetHoldReclaimedTimeV6() InheritanceInheritedUInt32`

GetHoldReclaimedTimeV6 returns the HoldReclaimedTimeV6 field if non-nil, zero value otherwise.

### GetHoldReclaimedTimeV6Ok

`func (o *InheritedDHCPConfig) GetHoldReclaimedTimeV6Ok() (*InheritanceInheritedUInt32, bool)`

GetHoldReclaimedTimeV6Ok returns a tuple with the HoldReclaimedTimeV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldReclaimedTimeV6

`func (o *InheritedDHCPConfig) SetHoldReclaimedTimeV6(v InheritanceInheritedUInt32)`

SetHoldReclaimedTimeV6 sets HoldReclaimedTimeV6 field to given value.

### HasHoldReclaimedTimeV6

`func (o *InheritedDHCPConfig) HasHoldReclaimedTimeV6() bool`

HasHoldReclaimedTimeV6 returns a boolean if a field has been set.

### GetIgnoreClientUid

`func (o *InheritedDHCPConfig) GetIgnoreClientUid() InheritanceInheritedBool`

GetIgnoreClientUid returns the IgnoreClientUid field if non-nil, zero value otherwise.

### GetIgnoreClientUidOk

`func (o *InheritedDHCPConfig) GetIgnoreClientUidOk() (*InheritanceInheritedBool, bool)`

GetIgnoreClientUidOk returns a tuple with the IgnoreClientUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreClientUid

`func (o *InheritedDHCPConfig) SetIgnoreClientUid(v InheritanceInheritedBool)`

SetIgnoreClientUid sets IgnoreClientUid field to given value.

### HasIgnoreClientUid

`func (o *InheritedDHCPConfig) HasIgnoreClientUid() bool`

HasIgnoreClientUid returns a boolean if a field has been set.

### GetIgnoreList

`func (o *InheritedDHCPConfig) GetIgnoreList() InheritedDHCPConfigIgnoreItemList`

GetIgnoreList returns the IgnoreList field if non-nil, zero value otherwise.

### GetIgnoreListOk

`func (o *InheritedDHCPConfig) GetIgnoreListOk() (*InheritedDHCPConfigIgnoreItemList, bool)`

GetIgnoreListOk returns a tuple with the IgnoreList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreList

`func (o *InheritedDHCPConfig) SetIgnoreList(v InheritedDHCPConfigIgnoreItemList)`

SetIgnoreList sets IgnoreList field to given value.

### HasIgnoreList

`func (o *InheritedDHCPConfig) HasIgnoreList() bool`

HasIgnoreList returns a boolean if a field has been set.

### GetLeaseTime

`func (o *InheritedDHCPConfig) GetLeaseTime() InheritanceInheritedUInt32`

GetLeaseTime returns the LeaseTime field if non-nil, zero value otherwise.

### GetLeaseTimeOk

`func (o *InheritedDHCPConfig) GetLeaseTimeOk() (*InheritanceInheritedUInt32, bool)`

GetLeaseTimeOk returns a tuple with the LeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTime

`func (o *InheritedDHCPConfig) SetLeaseTime(v InheritanceInheritedUInt32)`

SetLeaseTime sets LeaseTime field to given value.

### HasLeaseTime

`func (o *InheritedDHCPConfig) HasLeaseTime() bool`

HasLeaseTime returns a boolean if a field has been set.

### GetLeaseTimeV6

`func (o *InheritedDHCPConfig) GetLeaseTimeV6() InheritanceInheritedUInt32`

GetLeaseTimeV6 returns the LeaseTimeV6 field if non-nil, zero value otherwise.

### GetLeaseTimeV6Ok

`func (o *InheritedDHCPConfig) GetLeaseTimeV6Ok() (*InheritanceInheritedUInt32, bool)`

GetLeaseTimeV6Ok returns a tuple with the LeaseTimeV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTimeV6

`func (o *InheritedDHCPConfig) SetLeaseTimeV6(v InheritanceInheritedUInt32)`

SetLeaseTimeV6 sets LeaseTimeV6 field to given value.

### HasLeaseTimeV6

`func (o *InheritedDHCPConfig) HasLeaseTimeV6() bool`

HasLeaseTimeV6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


