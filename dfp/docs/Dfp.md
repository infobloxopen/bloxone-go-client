# Dfp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **time.Time** | The time when this DNS Forwarding Proxy object was created. | [optional] [readonly] 
**ElbIpList** | Pointer to **[]string** | The list of internal or local DNS servers&#39; IPv4 or IPv6 addresses that are used as ELB IPs. | [optional] [readonly] 
**ForwardingPolicy** | Pointer to **string** |  | [optional] 
**Host** | Pointer to [**[]DfpHost**](DfpHost.md) | host information. For internal Use only. | [optional] 
**Id** | Pointer to **int32** | The DNS Forwarding Proxy object identifier. | [optional] [readonly] 
**InternalDomainLists** | Pointer to **[]int32** | The list of internal domains list IDs that are associated with this DFP | [optional] 
**Name** | Pointer to **string** | The name of the DNS Forwarding Proxy. | [optional] [readonly] 
**NetAddrPolicyIds** | Pointer to [**[]NetAddrPolicyAssignment**](NetAddrPolicyAssignment.md) | List of network-address-scoped security policy assignments | [optional] 
**Ophid** | Pointer to **string** | The On-Prem Host identifier. | [optional] [readonly] 
**PolicyId** | Pointer to **int32** | The identifier of the security policy with which the DNS Forwarding Proxy is associated. | [optional] [readonly] 
**PopRegionId** | Pointer to **int32** | Point of Presence (PoP) region | [optional] [readonly] 
**ResolversAll** | Pointer to [**[]Resolver**](Resolver.md) |  | [optional] 
**ServiceId** | Pointer to **string** | The On-Prem Application Service identifier. For internal Use only | [optional] [readonly] 
**ServiceName** | Pointer to **string** | The On-Prem Application Service name. For internal Use only | [optional] [readonly] 
**SiteId** | Pointer to **string** | The DNS Forwarding Proxy site identifier that is appended to DNS queries originating from this DNS Forwarding Proxy and subsequently used for policy lookup purposes. | [optional] [readonly] 
**UpdatedTime** | Pointer to **time.Time** | The time when this DNS Forwarding Proxy object was last updated. | [optional] [readonly] 

## Methods

### NewDfp

`func NewDfp() *Dfp`

NewDfp instantiates a new Dfp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfpWithDefaults

`func NewDfpWithDefaults() *Dfp`

NewDfpWithDefaults instantiates a new Dfp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *Dfp) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Dfp) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Dfp) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Dfp) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetElbIpList

`func (o *Dfp) GetElbIpList() []string`

GetElbIpList returns the ElbIpList field if non-nil, zero value otherwise.

### GetElbIpListOk

`func (o *Dfp) GetElbIpListOk() (*[]string, bool)`

GetElbIpListOk returns a tuple with the ElbIpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElbIpList

`func (o *Dfp) SetElbIpList(v []string)`

SetElbIpList sets ElbIpList field to given value.

### HasElbIpList

`func (o *Dfp) HasElbIpList() bool`

HasElbIpList returns a boolean if a field has been set.

### GetForwardingPolicy

`func (o *Dfp) GetForwardingPolicy() string`

GetForwardingPolicy returns the ForwardingPolicy field if non-nil, zero value otherwise.

### GetForwardingPolicyOk

`func (o *Dfp) GetForwardingPolicyOk() (*string, bool)`

GetForwardingPolicyOk returns a tuple with the ForwardingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingPolicy

`func (o *Dfp) SetForwardingPolicy(v string)`

SetForwardingPolicy sets ForwardingPolicy field to given value.

### HasForwardingPolicy

`func (o *Dfp) HasForwardingPolicy() bool`

HasForwardingPolicy returns a boolean if a field has been set.

### GetHost

`func (o *Dfp) GetHost() []DfpHost`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Dfp) GetHostOk() (*[]DfpHost, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Dfp) SetHost(v []DfpHost)`

SetHost sets Host field to given value.

### HasHost

`func (o *Dfp) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *Dfp) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Dfp) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Dfp) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Dfp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalDomainLists

`func (o *Dfp) GetInternalDomainLists() []int32`

GetInternalDomainLists returns the InternalDomainLists field if non-nil, zero value otherwise.

### GetInternalDomainListsOk

`func (o *Dfp) GetInternalDomainListsOk() (*[]int32, bool)`

GetInternalDomainListsOk returns a tuple with the InternalDomainLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDomainLists

`func (o *Dfp) SetInternalDomainLists(v []int32)`

SetInternalDomainLists sets InternalDomainLists field to given value.

### HasInternalDomainLists

`func (o *Dfp) HasInternalDomainLists() bool`

HasInternalDomainLists returns a boolean if a field has been set.

### GetName

`func (o *Dfp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dfp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dfp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dfp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetAddrPolicyIds

`func (o *Dfp) GetNetAddrPolicyIds() []NetAddrPolicyAssignment`

GetNetAddrPolicyIds returns the NetAddrPolicyIds field if non-nil, zero value otherwise.

### GetNetAddrPolicyIdsOk

`func (o *Dfp) GetNetAddrPolicyIdsOk() (*[]NetAddrPolicyAssignment, bool)`

GetNetAddrPolicyIdsOk returns a tuple with the NetAddrPolicyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAddrPolicyIds

`func (o *Dfp) SetNetAddrPolicyIds(v []NetAddrPolicyAssignment)`

SetNetAddrPolicyIds sets NetAddrPolicyIds field to given value.

### HasNetAddrPolicyIds

`func (o *Dfp) HasNetAddrPolicyIds() bool`

HasNetAddrPolicyIds returns a boolean if a field has been set.

### GetOphid

`func (o *Dfp) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *Dfp) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *Dfp) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *Dfp) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetPolicyId

`func (o *Dfp) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *Dfp) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *Dfp) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *Dfp) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPopRegionId

`func (o *Dfp) GetPopRegionId() int32`

GetPopRegionId returns the PopRegionId field if non-nil, zero value otherwise.

### GetPopRegionIdOk

`func (o *Dfp) GetPopRegionIdOk() (*int32, bool)`

GetPopRegionIdOk returns a tuple with the PopRegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopRegionId

`func (o *Dfp) SetPopRegionId(v int32)`

SetPopRegionId sets PopRegionId field to given value.

### HasPopRegionId

`func (o *Dfp) HasPopRegionId() bool`

HasPopRegionId returns a boolean if a field has been set.

### GetResolversAll

`func (o *Dfp) GetResolversAll() []Resolver`

GetResolversAll returns the ResolversAll field if non-nil, zero value otherwise.

### GetResolversAllOk

`func (o *Dfp) GetResolversAllOk() (*[]Resolver, bool)`

GetResolversAllOk returns a tuple with the ResolversAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolversAll

`func (o *Dfp) SetResolversAll(v []Resolver)`

SetResolversAll sets ResolversAll field to given value.

### HasResolversAll

`func (o *Dfp) HasResolversAll() bool`

HasResolversAll returns a boolean if a field has been set.

### GetServiceId

`func (o *Dfp) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *Dfp) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *Dfp) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *Dfp) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *Dfp) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *Dfp) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *Dfp) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *Dfp) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSiteId

`func (o *Dfp) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Dfp) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Dfp) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Dfp) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *Dfp) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *Dfp) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *Dfp) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *Dfp) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


