# AtcdfpDfpCreateOrUpdatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultResolvers** | Pointer to **[]string** | The list of default DNS resolvers that will be used in case if the BloxOne Cloud is unreachable. Deprecated DO NOT USE. Use resolvers_all. | [optional] 
**ForwardingPolicy** | Pointer to **string** | The type of DNS resolver as Forwarding Policy. It can hold values as ib_cloud_first, external_first or external_only The default value is ib_cloud_first. If empty string is sent then ib_cloud_first will be considered. | [optional] 
**Host** | Pointer to [**[]AtcdfpDfpHost**](AtcdfpDfpHost.md) | host information. For internal Use only. | [optional] 
**Id** | Pointer to **int32** | The DNS Forwarding Proxy object identifier. | [optional] [readonly] 
**InternalDomainLists** | Pointer to **[]int32** | The list of internal domain list ids associated with this DFP (or resolvers) | [optional] 
**Name** | Pointer to **string** | The name of the DNS Forwarding Proxy. | [optional] 
**PopRegionId** | Pointer to **int32** | Point of Presence (PoP) region | [optional] 
**Resolvers** | Pointer to **[]string** | The list of internal or local DNS servers&#39; IPv4 or IPv6 addresses that are used as DNS resolvers. Deprecated DO NOT USE. Use resolvers_all. | [optional] 
**ResolversAll** | Pointer to [**[]AtcdfpResolver**](AtcdfpResolver.md) | The DNS forwarding proxy additional resolvers used for fallback and local resolution. This field replaces resolvers and default_resolvers fields which are deprecated. Either deprecated fields or new field can be used, both can not be used at same time. | [optional] 
**ServiceId** | Pointer to **string** | The DNS Forwarding Proxy Service ID object identifier. | [optional] 
**ServiceName** | Pointer to **string** | The name of the DNS Forwarding Proxy Service. | [optional] 
**SiteId** | Pointer to **string** | The DNS Forwarding Proxy site identifier that is appended to DNS queries originating from this DNS Forwarding Proxy and subsequently used for policy lookup purposes. | [optional] 

## Methods

### NewAtcdfpDfpCreateOrUpdatePayload

`func NewAtcdfpDfpCreateOrUpdatePayload() *AtcdfpDfpCreateOrUpdatePayload`

NewAtcdfpDfpCreateOrUpdatePayload instantiates a new AtcdfpDfpCreateOrUpdatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtcdfpDfpCreateOrUpdatePayloadWithDefaults

`func NewAtcdfpDfpCreateOrUpdatePayloadWithDefaults() *AtcdfpDfpCreateOrUpdatePayload`

NewAtcdfpDfpCreateOrUpdatePayloadWithDefaults instantiates a new AtcdfpDfpCreateOrUpdatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultResolvers

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetDefaultResolvers() []string`

GetDefaultResolvers returns the DefaultResolvers field if non-nil, zero value otherwise.

### GetDefaultResolversOk

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetDefaultResolversOk() (*[]string, bool)`

GetDefaultResolversOk returns a tuple with the DefaultResolvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultResolvers

`func (o *AtcdfpDfpCreateOrUpdatePayload) SetDefaultResolvers(v []string)`

SetDefaultResolvers sets DefaultResolvers field to given value.

### HasDefaultResolvers

`func (o *AtcdfpDfpCreateOrUpdatePayload) HasDefaultResolvers() bool`

HasDefaultResolvers returns a boolean if a field has been set.

### GetForwardingPolicy

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetForwardingPolicy() string`

GetForwardingPolicy returns the ForwardingPolicy field if non-nil, zero value otherwise.

### GetForwardingPolicyOk

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetForwardingPolicyOk() (*string, bool)`

GetForwardingPolicyOk returns a tuple with the ForwardingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingPolicy

`func (o *AtcdfpDfpCreateOrUpdatePayload) SetForwardingPolicy(v string)`

SetForwardingPolicy sets ForwardingPolicy field to given value.

### HasForwardingPolicy

`func (o *AtcdfpDfpCreateOrUpdatePayload) HasForwardingPolicy() bool`

HasForwardingPolicy returns a boolean if a field has been set.

### GetHost

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetHost() []AtcdfpDfpHost`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetHostOk() (*[]AtcdfpDfpHost, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AtcdfpDfpCreateOrUpdatePayload) SetHost(v []AtcdfpDfpHost)`

SetHost sets Host field to given value.

### HasHost

`func (o *AtcdfpDfpCreateOrUpdatePayload) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AtcdfpDfpCreateOrUpdatePayload) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AtcdfpDfpCreateOrUpdatePayload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalDomainLists

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetInternalDomainLists() []int32`

GetInternalDomainLists returns the InternalDomainLists field if non-nil, zero value otherwise.

### GetInternalDomainListsOk

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetInternalDomainListsOk() (*[]int32, bool)`

GetInternalDomainListsOk returns a tuple with the InternalDomainLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDomainLists

`func (o *AtcdfpDfpCreateOrUpdatePayload) SetInternalDomainLists(v []int32)`

SetInternalDomainLists sets InternalDomainLists field to given value.

### HasInternalDomainLists

`func (o *AtcdfpDfpCreateOrUpdatePayload) HasInternalDomainLists() bool`

HasInternalDomainLists returns a boolean if a field has been set.

### GetName

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AtcdfpDfpCreateOrUpdatePayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AtcdfpDfpCreateOrUpdatePayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPopRegionId

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetPopRegionId() int32`

GetPopRegionId returns the PopRegionId field if non-nil, zero value otherwise.

### GetPopRegionIdOk

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetPopRegionIdOk() (*int32, bool)`

GetPopRegionIdOk returns a tuple with the PopRegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopRegionId

`func (o *AtcdfpDfpCreateOrUpdatePayload) SetPopRegionId(v int32)`

SetPopRegionId sets PopRegionId field to given value.

### HasPopRegionId

`func (o *AtcdfpDfpCreateOrUpdatePayload) HasPopRegionId() bool`

HasPopRegionId returns a boolean if a field has been set.

### GetResolvers

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetResolvers() []string`

GetResolvers returns the Resolvers field if non-nil, zero value otherwise.

### GetResolversOk

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetResolversOk() (*[]string, bool)`

GetResolversOk returns a tuple with the Resolvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvers

`func (o *AtcdfpDfpCreateOrUpdatePayload) SetResolvers(v []string)`

SetResolvers sets Resolvers field to given value.

### HasResolvers

`func (o *AtcdfpDfpCreateOrUpdatePayload) HasResolvers() bool`

HasResolvers returns a boolean if a field has been set.

### GetResolversAll

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetResolversAll() []AtcdfpResolver`

GetResolversAll returns the ResolversAll field if non-nil, zero value otherwise.

### GetResolversAllOk

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetResolversAllOk() (*[]AtcdfpResolver, bool)`

GetResolversAllOk returns a tuple with the ResolversAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolversAll

`func (o *AtcdfpDfpCreateOrUpdatePayload) SetResolversAll(v []AtcdfpResolver)`

SetResolversAll sets ResolversAll field to given value.

### HasResolversAll

`func (o *AtcdfpDfpCreateOrUpdatePayload) HasResolversAll() bool`

HasResolversAll returns a boolean if a field has been set.

### GetServiceId

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *AtcdfpDfpCreateOrUpdatePayload) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *AtcdfpDfpCreateOrUpdatePayload) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *AtcdfpDfpCreateOrUpdatePayload) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *AtcdfpDfpCreateOrUpdatePayload) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSiteId

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *AtcdfpDfpCreateOrUpdatePayload) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *AtcdfpDfpCreateOrUpdatePayload) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *AtcdfpDfpCreateOrUpdatePayload) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


