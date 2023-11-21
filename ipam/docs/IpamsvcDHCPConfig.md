# IpamsvcDHCPConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbandonedReclaimTime** | Pointer to **int64** | The abandoned reclaim time in seconds for IPV4 clients. | [optional] [default to 3600]
**AbandonedReclaimTimeV6** | Pointer to **int64** | The abandoned reclaim time in seconds for IPV6 clients. | [optional] [default to 3600]
**AllowUnknown** | Pointer to **bool** | Disable to allow leases only for known IPv4 clients, those for which a fixed address is configured. | [optional] [default to true]
**AllowUnknownV6** | Pointer to **bool** | Disable to allow leases only for known IPV6 clients, those for which a fixed address is configured. | [optional] [default to true]
**Filters** | Pointer to **[]string** | The resource identifier. | [optional] 
**FiltersV6** | Pointer to **[]string** | The resource identifier. | [optional] 
**IgnoreClientUid** | Pointer to **bool** | Enable to ignore the client UID when issuing a DHCP lease. Use this option to prevent assigning two IP addresses for a client which does not have a UID during one phase of PXE boot but acquires one for the other phase. | [optional] [default to false]
**IgnoreList** | Pointer to [**[]IpamsvcIgnoreItem**](IpamsvcIgnoreItem.md) | The list of clients to ignore requests from. | [optional] 
**LeaseTime** | Pointer to **int64** | The lease duration in seconds. | [optional] [default to 3600]
**LeaseTimeV6** | Pointer to **int64** | The lease duration in seconds for IPV6 clients. | [optional] [default to 3600]

## Methods

### NewIpamsvcDHCPConfig

`func NewIpamsvcDHCPConfig() *IpamsvcDHCPConfig`

NewIpamsvcDHCPConfig instantiates a new IpamsvcDHCPConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcDHCPConfigWithDefaults

`func NewIpamsvcDHCPConfigWithDefaults() *IpamsvcDHCPConfig`

NewIpamsvcDHCPConfigWithDefaults instantiates a new IpamsvcDHCPConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbandonedReclaimTime

`func (o *IpamsvcDHCPConfig) GetAbandonedReclaimTime() int64`

GetAbandonedReclaimTime returns the AbandonedReclaimTime field if non-nil, zero value otherwise.

### GetAbandonedReclaimTimeOk

`func (o *IpamsvcDHCPConfig) GetAbandonedReclaimTimeOk() (*int64, bool)`

GetAbandonedReclaimTimeOk returns a tuple with the AbandonedReclaimTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandonedReclaimTime

`func (o *IpamsvcDHCPConfig) SetAbandonedReclaimTime(v int64)`

SetAbandonedReclaimTime sets AbandonedReclaimTime field to given value.

### HasAbandonedReclaimTime

`func (o *IpamsvcDHCPConfig) HasAbandonedReclaimTime() bool`

HasAbandonedReclaimTime returns a boolean if a field has been set.

### GetAbandonedReclaimTimeV6

`func (o *IpamsvcDHCPConfig) GetAbandonedReclaimTimeV6() int64`

GetAbandonedReclaimTimeV6 returns the AbandonedReclaimTimeV6 field if non-nil, zero value otherwise.

### GetAbandonedReclaimTimeV6Ok

`func (o *IpamsvcDHCPConfig) GetAbandonedReclaimTimeV6Ok() (*int64, bool)`

GetAbandonedReclaimTimeV6Ok returns a tuple with the AbandonedReclaimTimeV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandonedReclaimTimeV6

`func (o *IpamsvcDHCPConfig) SetAbandonedReclaimTimeV6(v int64)`

SetAbandonedReclaimTimeV6 sets AbandonedReclaimTimeV6 field to given value.

### HasAbandonedReclaimTimeV6

`func (o *IpamsvcDHCPConfig) HasAbandonedReclaimTimeV6() bool`

HasAbandonedReclaimTimeV6 returns a boolean if a field has been set.

### GetAllowUnknown

`func (o *IpamsvcDHCPConfig) GetAllowUnknown() bool`

GetAllowUnknown returns the AllowUnknown field if non-nil, zero value otherwise.

### GetAllowUnknownOk

`func (o *IpamsvcDHCPConfig) GetAllowUnknownOk() (*bool, bool)`

GetAllowUnknownOk returns a tuple with the AllowUnknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnknown

`func (o *IpamsvcDHCPConfig) SetAllowUnknown(v bool)`

SetAllowUnknown sets AllowUnknown field to given value.

### HasAllowUnknown

`func (o *IpamsvcDHCPConfig) HasAllowUnknown() bool`

HasAllowUnknown returns a boolean if a field has been set.

### GetAllowUnknownV6

`func (o *IpamsvcDHCPConfig) GetAllowUnknownV6() bool`

GetAllowUnknownV6 returns the AllowUnknownV6 field if non-nil, zero value otherwise.

### GetAllowUnknownV6Ok

`func (o *IpamsvcDHCPConfig) GetAllowUnknownV6Ok() (*bool, bool)`

GetAllowUnknownV6Ok returns a tuple with the AllowUnknownV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnknownV6

`func (o *IpamsvcDHCPConfig) SetAllowUnknownV6(v bool)`

SetAllowUnknownV6 sets AllowUnknownV6 field to given value.

### HasAllowUnknownV6

`func (o *IpamsvcDHCPConfig) HasAllowUnknownV6() bool`

HasAllowUnknownV6 returns a boolean if a field has been set.

### GetFilters

`func (o *IpamsvcDHCPConfig) GetFilters() []string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *IpamsvcDHCPConfig) GetFiltersOk() (*[]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *IpamsvcDHCPConfig) SetFilters(v []string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *IpamsvcDHCPConfig) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetFiltersV6

`func (o *IpamsvcDHCPConfig) GetFiltersV6() []string`

GetFiltersV6 returns the FiltersV6 field if non-nil, zero value otherwise.

### GetFiltersV6Ok

`func (o *IpamsvcDHCPConfig) GetFiltersV6Ok() (*[]string, bool)`

GetFiltersV6Ok returns a tuple with the FiltersV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltersV6

`func (o *IpamsvcDHCPConfig) SetFiltersV6(v []string)`

SetFiltersV6 sets FiltersV6 field to given value.

### HasFiltersV6

`func (o *IpamsvcDHCPConfig) HasFiltersV6() bool`

HasFiltersV6 returns a boolean if a field has been set.

### GetIgnoreClientUid

`func (o *IpamsvcDHCPConfig) GetIgnoreClientUid() bool`

GetIgnoreClientUid returns the IgnoreClientUid field if non-nil, zero value otherwise.

### GetIgnoreClientUidOk

`func (o *IpamsvcDHCPConfig) GetIgnoreClientUidOk() (*bool, bool)`

GetIgnoreClientUidOk returns a tuple with the IgnoreClientUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreClientUid

`func (o *IpamsvcDHCPConfig) SetIgnoreClientUid(v bool)`

SetIgnoreClientUid sets IgnoreClientUid field to given value.

### HasIgnoreClientUid

`func (o *IpamsvcDHCPConfig) HasIgnoreClientUid() bool`

HasIgnoreClientUid returns a boolean if a field has been set.

### GetIgnoreList

`func (o *IpamsvcDHCPConfig) GetIgnoreList() []IpamsvcIgnoreItem`

GetIgnoreList returns the IgnoreList field if non-nil, zero value otherwise.

### GetIgnoreListOk

`func (o *IpamsvcDHCPConfig) GetIgnoreListOk() (*[]IpamsvcIgnoreItem, bool)`

GetIgnoreListOk returns a tuple with the IgnoreList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreList

`func (o *IpamsvcDHCPConfig) SetIgnoreList(v []IpamsvcIgnoreItem)`

SetIgnoreList sets IgnoreList field to given value.

### HasIgnoreList

`func (o *IpamsvcDHCPConfig) HasIgnoreList() bool`

HasIgnoreList returns a boolean if a field has been set.

### GetLeaseTime

`func (o *IpamsvcDHCPConfig) GetLeaseTime() int64`

GetLeaseTime returns the LeaseTime field if non-nil, zero value otherwise.

### GetLeaseTimeOk

`func (o *IpamsvcDHCPConfig) GetLeaseTimeOk() (*int64, bool)`

GetLeaseTimeOk returns a tuple with the LeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTime

`func (o *IpamsvcDHCPConfig) SetLeaseTime(v int64)`

SetLeaseTime sets LeaseTime field to given value.

### HasLeaseTime

`func (o *IpamsvcDHCPConfig) HasLeaseTime() bool`

HasLeaseTime returns a boolean if a field has been set.

### GetLeaseTimeV6

`func (o *IpamsvcDHCPConfig) GetLeaseTimeV6() int64`

GetLeaseTimeV6 returns the LeaseTimeV6 field if non-nil, zero value otherwise.

### GetLeaseTimeV6Ok

`func (o *IpamsvcDHCPConfig) GetLeaseTimeV6Ok() (*int64, bool)`

GetLeaseTimeV6Ok returns a tuple with the LeaseTimeV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTimeV6

`func (o *IpamsvcDHCPConfig) SetLeaseTimeV6(v int64)`

SetLeaseTimeV6 sets LeaseTimeV6 field to given value.

### HasLeaseTimeV6

`func (o *IpamsvcDHCPConfig) HasLeaseTimeV6() bool`

HasLeaseTimeV6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


