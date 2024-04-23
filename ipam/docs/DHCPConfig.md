# DHCPConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbandonedReclaimTime** | Pointer to **int64** | The abandoned reclaim time in seconds for IPV4 clients. | [optional] [default to 3600]
**AbandonedReclaimTimeV6** | Pointer to **int64** | The abandoned reclaim time in seconds for IPV6 clients. | [optional] [default to 3600]
**AllowUnknown** | Pointer to **bool** | Disable to allow leases only for known IPv4 clients, those for which a fixed address is configured. | [optional] [default to true]
**AllowUnknownV6** | Pointer to **bool** | Disable to allow leases only for known IPV6 clients, those for which a fixed address is configured. | [optional] [default to true]
**EchoClientId** | Pointer to **bool** | Enable/disable to include/exclude the client id when responding to discover or request. | [optional] [default to false]
**Filters** | Pointer to **[]string** | The resource identifier. | [optional] 
**FiltersV6** | Pointer to **[]string** | The resource identifier. | [optional] 
**IgnoreClientUid** | Pointer to **bool** | Enable to ignore the client UID when issuing a DHCP lease. Use this option to prevent assigning two IP addresses for a client which does not have a UID during one phase of PXE boot but acquires one for the other phase. | [optional] [default to false]
**IgnoreList** | Pointer to [**[]IgnoreItem**](IgnoreItem.md) | The list of clients to ignore requests from. | [optional] 
**LeaseTime** | Pointer to **int64** | The lease duration in seconds. | [optional] [default to 3600]
**LeaseTimeV6** | Pointer to **int64** | The lease duration in seconds for IPV6 clients. | [optional] [default to 3600]

## Methods

### NewDHCPConfig

`func NewDHCPConfig() *DHCPConfig`

NewDHCPConfig instantiates a new DHCPConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDHCPConfigWithDefaults

`func NewDHCPConfigWithDefaults() *DHCPConfig`

NewDHCPConfigWithDefaults instantiates a new DHCPConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbandonedReclaimTime

`func (o *DHCPConfig) GetAbandonedReclaimTime() int64`

GetAbandonedReclaimTime returns the AbandonedReclaimTime field if non-nil, zero value otherwise.

### GetAbandonedReclaimTimeOk

`func (o *DHCPConfig) GetAbandonedReclaimTimeOk() (*int64, bool)`

GetAbandonedReclaimTimeOk returns a tuple with the AbandonedReclaimTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandonedReclaimTime

`func (o *DHCPConfig) SetAbandonedReclaimTime(v int64)`

SetAbandonedReclaimTime sets AbandonedReclaimTime field to given value.

### HasAbandonedReclaimTime

`func (o *DHCPConfig) HasAbandonedReclaimTime() bool`

HasAbandonedReclaimTime returns a boolean if a field has been set.

### GetAbandonedReclaimTimeV6

`func (o *DHCPConfig) GetAbandonedReclaimTimeV6() int64`

GetAbandonedReclaimTimeV6 returns the AbandonedReclaimTimeV6 field if non-nil, zero value otherwise.

### GetAbandonedReclaimTimeV6Ok

`func (o *DHCPConfig) GetAbandonedReclaimTimeV6Ok() (*int64, bool)`

GetAbandonedReclaimTimeV6Ok returns a tuple with the AbandonedReclaimTimeV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandonedReclaimTimeV6

`func (o *DHCPConfig) SetAbandonedReclaimTimeV6(v int64)`

SetAbandonedReclaimTimeV6 sets AbandonedReclaimTimeV6 field to given value.

### HasAbandonedReclaimTimeV6

`func (o *DHCPConfig) HasAbandonedReclaimTimeV6() bool`

HasAbandonedReclaimTimeV6 returns a boolean if a field has been set.

### GetAllowUnknown

`func (o *DHCPConfig) GetAllowUnknown() bool`

GetAllowUnknown returns the AllowUnknown field if non-nil, zero value otherwise.

### GetAllowUnknownOk

`func (o *DHCPConfig) GetAllowUnknownOk() (*bool, bool)`

GetAllowUnknownOk returns a tuple with the AllowUnknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnknown

`func (o *DHCPConfig) SetAllowUnknown(v bool)`

SetAllowUnknown sets AllowUnknown field to given value.

### HasAllowUnknown

`func (o *DHCPConfig) HasAllowUnknown() bool`

HasAllowUnknown returns a boolean if a field has been set.

### GetAllowUnknownV6

`func (o *DHCPConfig) GetAllowUnknownV6() bool`

GetAllowUnknownV6 returns the AllowUnknownV6 field if non-nil, zero value otherwise.

### GetAllowUnknownV6Ok

`func (o *DHCPConfig) GetAllowUnknownV6Ok() (*bool, bool)`

GetAllowUnknownV6Ok returns a tuple with the AllowUnknownV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnknownV6

`func (o *DHCPConfig) SetAllowUnknownV6(v bool)`

SetAllowUnknownV6 sets AllowUnknownV6 field to given value.

### HasAllowUnknownV6

`func (o *DHCPConfig) HasAllowUnknownV6() bool`

HasAllowUnknownV6 returns a boolean if a field has been set.

### GetEchoClientId

`func (o *DHCPConfig) GetEchoClientId() bool`

GetEchoClientId returns the EchoClientId field if non-nil, zero value otherwise.

### GetEchoClientIdOk

`func (o *DHCPConfig) GetEchoClientIdOk() (*bool, bool)`

GetEchoClientIdOk returns a tuple with the EchoClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEchoClientId

`func (o *DHCPConfig) SetEchoClientId(v bool)`

SetEchoClientId sets EchoClientId field to given value.

### HasEchoClientId

`func (o *DHCPConfig) HasEchoClientId() bool`

HasEchoClientId returns a boolean if a field has been set.

### GetFilters

`func (o *DHCPConfig) GetFilters() []string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DHCPConfig) GetFiltersOk() (*[]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DHCPConfig) SetFilters(v []string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *DHCPConfig) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetFiltersV6

`func (o *DHCPConfig) GetFiltersV6() []string`

GetFiltersV6 returns the FiltersV6 field if non-nil, zero value otherwise.

### GetFiltersV6Ok

`func (o *DHCPConfig) GetFiltersV6Ok() (*[]string, bool)`

GetFiltersV6Ok returns a tuple with the FiltersV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltersV6

`func (o *DHCPConfig) SetFiltersV6(v []string)`

SetFiltersV6 sets FiltersV6 field to given value.

### HasFiltersV6

`func (o *DHCPConfig) HasFiltersV6() bool`

HasFiltersV6 returns a boolean if a field has been set.

### GetIgnoreClientUid

`func (o *DHCPConfig) GetIgnoreClientUid() bool`

GetIgnoreClientUid returns the IgnoreClientUid field if non-nil, zero value otherwise.

### GetIgnoreClientUidOk

`func (o *DHCPConfig) GetIgnoreClientUidOk() (*bool, bool)`

GetIgnoreClientUidOk returns a tuple with the IgnoreClientUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreClientUid

`func (o *DHCPConfig) SetIgnoreClientUid(v bool)`

SetIgnoreClientUid sets IgnoreClientUid field to given value.

### HasIgnoreClientUid

`func (o *DHCPConfig) HasIgnoreClientUid() bool`

HasIgnoreClientUid returns a boolean if a field has been set.

### GetIgnoreList

`func (o *DHCPConfig) GetIgnoreList() []IgnoreItem`

GetIgnoreList returns the IgnoreList field if non-nil, zero value otherwise.

### GetIgnoreListOk

`func (o *DHCPConfig) GetIgnoreListOk() (*[]IgnoreItem, bool)`

GetIgnoreListOk returns a tuple with the IgnoreList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreList

`func (o *DHCPConfig) SetIgnoreList(v []IgnoreItem)`

SetIgnoreList sets IgnoreList field to given value.

### HasIgnoreList

`func (o *DHCPConfig) HasIgnoreList() bool`

HasIgnoreList returns a boolean if a field has been set.

### GetLeaseTime

`func (o *DHCPConfig) GetLeaseTime() int64`

GetLeaseTime returns the LeaseTime field if non-nil, zero value otherwise.

### GetLeaseTimeOk

`func (o *DHCPConfig) GetLeaseTimeOk() (*int64, bool)`

GetLeaseTimeOk returns a tuple with the LeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTime

`func (o *DHCPConfig) SetLeaseTime(v int64)`

SetLeaseTime sets LeaseTime field to given value.

### HasLeaseTime

`func (o *DHCPConfig) HasLeaseTime() bool`

HasLeaseTime returns a boolean if a field has been set.

### GetLeaseTimeV6

`func (o *DHCPConfig) GetLeaseTimeV6() int64`

GetLeaseTimeV6 returns the LeaseTimeV6 field if non-nil, zero value otherwise.

### GetLeaseTimeV6Ok

`func (o *DHCPConfig) GetLeaseTimeV6Ok() (*int64, bool)`

GetLeaseTimeV6Ok returns a tuple with the LeaseTimeV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTimeV6

`func (o *DHCPConfig) SetLeaseTimeV6(v int64)`

SetLeaseTimeV6 sets LeaseTimeV6 field to given value.

### HasLeaseTimeV6

`func (o *DHCPConfig) HasLeaseTimeV6() bool`

HasLeaseTimeV6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


