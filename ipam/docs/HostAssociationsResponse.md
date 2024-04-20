# HostAssociationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpPktStats** | Pointer to [**DHCPPacketStats**](DHCPPacketStats.md) |  | [optional] 
**HaGroups** | Pointer to [**[]HAGroup**](HAGroup.md) | The list of HA groups. | [optional] 
**Host** | Pointer to [**Host**](Host.md) |  | [optional] 
**Subnets** | Pointer to [**[]Subnet**](Subnet.md) | The list of subnets. | [optional] 

## Methods

### NewHostAssociationsResponse

`func NewHostAssociationsResponse() *HostAssociationsResponse`

NewHostAssociationsResponse instantiates a new HostAssociationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostAssociationsResponseWithDefaults

`func NewHostAssociationsResponseWithDefaults() *HostAssociationsResponse`

NewHostAssociationsResponseWithDefaults instantiates a new HostAssociationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpPktStats

`func (o *HostAssociationsResponse) GetDhcpPktStats() DHCPPacketStats`

GetDhcpPktStats returns the DhcpPktStats field if non-nil, zero value otherwise.

### GetDhcpPktStatsOk

`func (o *HostAssociationsResponse) GetDhcpPktStatsOk() (*DHCPPacketStats, bool)`

GetDhcpPktStatsOk returns a tuple with the DhcpPktStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpPktStats

`func (o *HostAssociationsResponse) SetDhcpPktStats(v DHCPPacketStats)`

SetDhcpPktStats sets DhcpPktStats field to given value.

### HasDhcpPktStats

`func (o *HostAssociationsResponse) HasDhcpPktStats() bool`

HasDhcpPktStats returns a boolean if a field has been set.

### GetHaGroups

`func (o *HostAssociationsResponse) GetHaGroups() []HAGroup`

GetHaGroups returns the HaGroups field if non-nil, zero value otherwise.

### GetHaGroupsOk

`func (o *HostAssociationsResponse) GetHaGroupsOk() (*[]HAGroup, bool)`

GetHaGroupsOk returns a tuple with the HaGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaGroups

`func (o *HostAssociationsResponse) SetHaGroups(v []HAGroup)`

SetHaGroups sets HaGroups field to given value.

### HasHaGroups

`func (o *HostAssociationsResponse) HasHaGroups() bool`

HasHaGroups returns a boolean if a field has been set.

### GetHost

`func (o *HostAssociationsResponse) GetHost() Host`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HostAssociationsResponse) GetHostOk() (*Host, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HostAssociationsResponse) SetHost(v Host)`

SetHost sets Host field to given value.

### HasHost

`func (o *HostAssociationsResponse) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetSubnets

`func (o *HostAssociationsResponse) GetSubnets() []Subnet`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *HostAssociationsResponse) GetSubnetsOk() (*[]Subnet, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *HostAssociationsResponse) SetSubnets(v []Subnet)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *HostAssociationsResponse) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


