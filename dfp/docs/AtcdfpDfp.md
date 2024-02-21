# AtcdfpDfp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **time.Time** | The time when this DNS Forwarding Proxy object was created. | [optional] [readonly] 
**DefaultResolvers** | Pointer to **[]string** | The list of default DNS resolvers that will be used in case if the BloxOne Cloud is unreachable.  Deprecated DO NOT USE. Use resolvers_all. | [optional] [readonly] 
**ElbIpList** | Pointer to **[]string** | The list of internal or local DNS servers&#39; IPv4 or IPv6 addresses that are used as ELB IPs. | [optional] [readonly] 
**ForwardingPolicy** | Pointer to **string** |  | [optional] 
**Host** | Pointer to [**[]AtcdfpDfpHost**](AtcdfpDfpHost.md) | host information. For internal Use only. | [optional] 
**Id** | Pointer to **int32** | The DNS Forwarding Proxy object identifier. | [optional] [readonly] 
**InternalDomainLists** | Pointer to **[]int32** | The list of internal domains list IDs that are associated with this DFP | [optional] 
**Name** | Pointer to **string** | The name of the DNS Forwarding Proxy. | [optional] [readonly] 
**NetAddrPolicyIds** | Pointer to [**[]AtcdfpNetAddrPolicyAssignment**](AtcdfpNetAddrPolicyAssignment.md) | List of network-address-scoped security policy assignments | [optional] 
**Ophid** | Pointer to **string** | The On-Prem Host identifier. | [optional] [readonly] 
**PolicyId** | Pointer to **int32** | The identifier of the security policy with which the DNS Forwarding Proxy is associated. | [optional] [readonly] 
**PopRegionId** | Pointer to **int32** | Point of Presence (PoP) region | [optional] [readonly] 
**Resolvers** | Pointer to **[]string** | The list of internal or local DNS servers&#39; IPv4 or IPv6 addresses that are used as DNS resolvers. Deprecated DO NOT USE. Use resolvers_all. | [optional] [readonly] 
**ResolversAll** | Pointer to [**[]AtcdfpResolver**](AtcdfpResolver.md) |  | [optional] 
**ServiceId** | Pointer to **string** | The On-Prem Application Service identifier. For internal Use only | [optional] [readonly] 
**ServiceName** | Pointer to **string** | The On-Prem Application Service name. For internal Use only | [optional] [readonly] 
**SiteId** | Pointer to **string** | The DNS Forwarding Proxy site identifier that is appended to DNS queries originating from this DNS Forwarding Proxy and subsequently used for policy lookup purposes. | [optional] [readonly] 
**UpdatedTime** | Pointer to **time.Time** | The time when this DNS Forwarding Proxy object was last updated. | [optional] [readonly] 

## Methods

### NewAtcdfpDfp

`func NewAtcdfpDfp() *AtcdfpDfp`

NewAtcdfpDfp instantiates a new AtcdfpDfp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtcdfpDfpWithDefaults

`func NewAtcdfpDfpWithDefaults() *AtcdfpDfp`

NewAtcdfpDfpWithDefaults instantiates a new AtcdfpDfp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *AtcdfpDfp) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AtcdfpDfp) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AtcdfpDfp) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *AtcdfpDfp) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDefaultResolvers

`func (o *AtcdfpDfp) GetDefaultResolvers() []string`

GetDefaultResolvers returns the DefaultResolvers field if non-nil, zero value otherwise.

### GetDefaultResolversOk

`func (o *AtcdfpDfp) GetDefaultResolversOk() (*[]string, bool)`

GetDefaultResolversOk returns a tuple with the DefaultResolvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultResolvers

`func (o *AtcdfpDfp) SetDefaultResolvers(v []string)`

SetDefaultResolvers sets DefaultResolvers field to given value.

### HasDefaultResolvers

`func (o *AtcdfpDfp) HasDefaultResolvers() bool`

HasDefaultResolvers returns a boolean if a field has been set.

### GetElbIpList

`func (o *AtcdfpDfp) GetElbIpList() []string`

GetElbIpList returns the ElbIpList field if non-nil, zero value otherwise.

### GetElbIpListOk

`func (o *AtcdfpDfp) GetElbIpListOk() (*[]string, bool)`

GetElbIpListOk returns a tuple with the ElbIpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElbIpList

`func (o *AtcdfpDfp) SetElbIpList(v []string)`

SetElbIpList sets ElbIpList field to given value.

### HasElbIpList

`func (o *AtcdfpDfp) HasElbIpList() bool`

HasElbIpList returns a boolean if a field has been set.

### GetForwardingPolicy

`func (o *AtcdfpDfp) GetForwardingPolicy() string`

GetForwardingPolicy returns the ForwardingPolicy field if non-nil, zero value otherwise.

### GetForwardingPolicyOk

`func (o *AtcdfpDfp) GetForwardingPolicyOk() (*string, bool)`

GetForwardingPolicyOk returns a tuple with the ForwardingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingPolicy

`func (o *AtcdfpDfp) SetForwardingPolicy(v string)`

SetForwardingPolicy sets ForwardingPolicy field to given value.

### HasForwardingPolicy

`func (o *AtcdfpDfp) HasForwardingPolicy() bool`

HasForwardingPolicy returns a boolean if a field has been set.

### GetHost

`func (o *AtcdfpDfp) GetHost() []AtcdfpDfpHost`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AtcdfpDfp) GetHostOk() (*[]AtcdfpDfpHost, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AtcdfpDfp) SetHost(v []AtcdfpDfpHost)`

SetHost sets Host field to given value.

### HasHost

`func (o *AtcdfpDfp) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *AtcdfpDfp) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AtcdfpDfp) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AtcdfpDfp) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AtcdfpDfp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalDomainLists

`func (o *AtcdfpDfp) GetInternalDomainLists() []int32`

GetInternalDomainLists returns the InternalDomainLists field if non-nil, zero value otherwise.

### GetInternalDomainListsOk

`func (o *AtcdfpDfp) GetInternalDomainListsOk() (*[]int32, bool)`

GetInternalDomainListsOk returns a tuple with the InternalDomainLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDomainLists

`func (o *AtcdfpDfp) SetInternalDomainLists(v []int32)`

SetInternalDomainLists sets InternalDomainLists field to given value.

### HasInternalDomainLists

`func (o *AtcdfpDfp) HasInternalDomainLists() bool`

HasInternalDomainLists returns a boolean if a field has been set.

### GetName

`func (o *AtcdfpDfp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AtcdfpDfp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AtcdfpDfp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AtcdfpDfp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetAddrPolicyIds

`func (o *AtcdfpDfp) GetNetAddrPolicyIds() []AtcdfpNetAddrPolicyAssignment`

GetNetAddrPolicyIds returns the NetAddrPolicyIds field if non-nil, zero value otherwise.

### GetNetAddrPolicyIdsOk

`func (o *AtcdfpDfp) GetNetAddrPolicyIdsOk() (*[]AtcdfpNetAddrPolicyAssignment, bool)`

GetNetAddrPolicyIdsOk returns a tuple with the NetAddrPolicyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAddrPolicyIds

`func (o *AtcdfpDfp) SetNetAddrPolicyIds(v []AtcdfpNetAddrPolicyAssignment)`

SetNetAddrPolicyIds sets NetAddrPolicyIds field to given value.

### HasNetAddrPolicyIds

`func (o *AtcdfpDfp) HasNetAddrPolicyIds() bool`

HasNetAddrPolicyIds returns a boolean if a field has been set.

### GetOphid

`func (o *AtcdfpDfp) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *AtcdfpDfp) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *AtcdfpDfp) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *AtcdfpDfp) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetPolicyId

`func (o *AtcdfpDfp) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *AtcdfpDfp) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *AtcdfpDfp) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *AtcdfpDfp) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPopRegionId

`func (o *AtcdfpDfp) GetPopRegionId() int32`

GetPopRegionId returns the PopRegionId field if non-nil, zero value otherwise.

### GetPopRegionIdOk

`func (o *AtcdfpDfp) GetPopRegionIdOk() (*int32, bool)`

GetPopRegionIdOk returns a tuple with the PopRegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopRegionId

`func (o *AtcdfpDfp) SetPopRegionId(v int32)`

SetPopRegionId sets PopRegionId field to given value.

### HasPopRegionId

`func (o *AtcdfpDfp) HasPopRegionId() bool`

HasPopRegionId returns a boolean if a field has been set.

### GetResolvers

`func (o *AtcdfpDfp) GetResolvers() []string`

GetResolvers returns the Resolvers field if non-nil, zero value otherwise.

### GetResolversOk

`func (o *AtcdfpDfp) GetResolversOk() (*[]string, bool)`

GetResolversOk returns a tuple with the Resolvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvers

`func (o *AtcdfpDfp) SetResolvers(v []string)`

SetResolvers sets Resolvers field to given value.

### HasResolvers

`func (o *AtcdfpDfp) HasResolvers() bool`

HasResolvers returns a boolean if a field has been set.

### GetResolversAll

`func (o *AtcdfpDfp) GetResolversAll() []AtcdfpResolver`

GetResolversAll returns the ResolversAll field if non-nil, zero value otherwise.

### GetResolversAllOk

`func (o *AtcdfpDfp) GetResolversAllOk() (*[]AtcdfpResolver, bool)`

GetResolversAllOk returns a tuple with the ResolversAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolversAll

`func (o *AtcdfpDfp) SetResolversAll(v []AtcdfpResolver)`

SetResolversAll sets ResolversAll field to given value.

### HasResolversAll

`func (o *AtcdfpDfp) HasResolversAll() bool`

HasResolversAll returns a boolean if a field has been set.

### GetServiceId

`func (o *AtcdfpDfp) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *AtcdfpDfp) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *AtcdfpDfp) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *AtcdfpDfp) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *AtcdfpDfp) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *AtcdfpDfp) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *AtcdfpDfp) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *AtcdfpDfp) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSiteId

`func (o *AtcdfpDfp) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *AtcdfpDfp) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *AtcdfpDfp) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *AtcdfpDfp) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *AtcdfpDfp) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *AtcdfpDfp) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *AtcdfpDfp) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *AtcdfpDfp) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


