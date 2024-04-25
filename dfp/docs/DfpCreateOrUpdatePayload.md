# DfpCreateOrUpdatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForwardingPolicy** | Pointer to **string** | The type of DNS resolver as Forwarding Policy. It can hold values as ib_cloud_first, external_first or external_only The default value is ib_cloud_first. If empty string is sent then ib_cloud_first will be considered. | [optional] 
**Host** | Pointer to [**[]DfpHost**](DfpHost.md) | host information. For internal Use only. | [optional] 
**Id** | Pointer to **int32** | The DNS Forwarding Proxy object identifier. | [optional] [readonly] 
**InternalDomainLists** | Pointer to **[]int32** | The list of internal domain list ids associated with this DFP (or resolvers) | [optional] 
**Name** | Pointer to **string** | The name of the DNS Forwarding Proxy. | [optional] 
**PopRegionId** | Pointer to **int32** | Point of Presence (PoP) region | [optional] 
**ResolversAll** | Pointer to [**[]Resolver**](Resolver.md) | The DNS forwarding proxy additional resolvers used for fallback and local resolution. This field replaces resolvers and default_resolvers fields which are deprecated. Either deprecated fields or new field can be used, both can not be used at same time. | [optional] 
**ServiceId** | Pointer to **string** | The DNS Forwarding Proxy Service ID object identifier. | [optional] 
**ServiceName** | Pointer to **string** | The name of the DNS Forwarding Proxy Service. | [optional] 
**SiteId** | Pointer to **string** | The DNS Forwarding Proxy site identifier that is appended to DNS queries originating from this DNS Forwarding Proxy and subsequently used for policy lookup purposes. | [optional] 

## Methods

### NewDfpCreateOrUpdatePayload

`func NewDfpCreateOrUpdatePayload() *DfpCreateOrUpdatePayload`

NewDfpCreateOrUpdatePayload instantiates a new DfpCreateOrUpdatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfpCreateOrUpdatePayloadWithDefaults

`func NewDfpCreateOrUpdatePayloadWithDefaults() *DfpCreateOrUpdatePayload`

NewDfpCreateOrUpdatePayloadWithDefaults instantiates a new DfpCreateOrUpdatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForwardingPolicy

`func (o *DfpCreateOrUpdatePayload) GetForwardingPolicy() string`

GetForwardingPolicy returns the ForwardingPolicy field if non-nil, zero value otherwise.

### GetForwardingPolicyOk

`func (o *DfpCreateOrUpdatePayload) GetForwardingPolicyOk() (*string, bool)`

GetForwardingPolicyOk returns a tuple with the ForwardingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingPolicy

`func (o *DfpCreateOrUpdatePayload) SetForwardingPolicy(v string)`

SetForwardingPolicy sets ForwardingPolicy field to given value.

### HasForwardingPolicy

`func (o *DfpCreateOrUpdatePayload) HasForwardingPolicy() bool`

HasForwardingPolicy returns a boolean if a field has been set.

### GetHost

`func (o *DfpCreateOrUpdatePayload) GetHost() []DfpHost`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DfpCreateOrUpdatePayload) GetHostOk() (*[]DfpHost, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DfpCreateOrUpdatePayload) SetHost(v []DfpHost)`

SetHost sets Host field to given value.

### HasHost

`func (o *DfpCreateOrUpdatePayload) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *DfpCreateOrUpdatePayload) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfpCreateOrUpdatePayload) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfpCreateOrUpdatePayload) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DfpCreateOrUpdatePayload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalDomainLists

`func (o *DfpCreateOrUpdatePayload) GetInternalDomainLists() []int32`

GetInternalDomainLists returns the InternalDomainLists field if non-nil, zero value otherwise.

### GetInternalDomainListsOk

`func (o *DfpCreateOrUpdatePayload) GetInternalDomainListsOk() (*[]int32, bool)`

GetInternalDomainListsOk returns a tuple with the InternalDomainLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDomainLists

`func (o *DfpCreateOrUpdatePayload) SetInternalDomainLists(v []int32)`

SetInternalDomainLists sets InternalDomainLists field to given value.

### HasInternalDomainLists

`func (o *DfpCreateOrUpdatePayload) HasInternalDomainLists() bool`

HasInternalDomainLists returns a boolean if a field has been set.

### GetName

`func (o *DfpCreateOrUpdatePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfpCreateOrUpdatePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfpCreateOrUpdatePayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfpCreateOrUpdatePayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPopRegionId

`func (o *DfpCreateOrUpdatePayload) GetPopRegionId() int32`

GetPopRegionId returns the PopRegionId field if non-nil, zero value otherwise.

### GetPopRegionIdOk

`func (o *DfpCreateOrUpdatePayload) GetPopRegionIdOk() (*int32, bool)`

GetPopRegionIdOk returns a tuple with the PopRegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopRegionId

`func (o *DfpCreateOrUpdatePayload) SetPopRegionId(v int32)`

SetPopRegionId sets PopRegionId field to given value.

### HasPopRegionId

`func (o *DfpCreateOrUpdatePayload) HasPopRegionId() bool`

HasPopRegionId returns a boolean if a field has been set.

### GetResolversAll

`func (o *DfpCreateOrUpdatePayload) GetResolversAll() []Resolver`

GetResolversAll returns the ResolversAll field if non-nil, zero value otherwise.

### GetResolversAllOk

`func (o *DfpCreateOrUpdatePayload) GetResolversAllOk() (*[]Resolver, bool)`

GetResolversAllOk returns a tuple with the ResolversAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolversAll

`func (o *DfpCreateOrUpdatePayload) SetResolversAll(v []Resolver)`

SetResolversAll sets ResolversAll field to given value.

### HasResolversAll

`func (o *DfpCreateOrUpdatePayload) HasResolversAll() bool`

HasResolversAll returns a boolean if a field has been set.

### GetServiceId

`func (o *DfpCreateOrUpdatePayload) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DfpCreateOrUpdatePayload) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DfpCreateOrUpdatePayload) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *DfpCreateOrUpdatePayload) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *DfpCreateOrUpdatePayload) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *DfpCreateOrUpdatePayload) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *DfpCreateOrUpdatePayload) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *DfpCreateOrUpdatePayload) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSiteId

`func (o *DfpCreateOrUpdatePayload) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *DfpCreateOrUpdatePayload) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *DfpCreateOrUpdatePayload) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *DfpCreateOrUpdatePayload) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


