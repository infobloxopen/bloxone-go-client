# IpamsvcHostAssociationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpPktStats** | Pointer to [**IpamsvcDHCPPacketStats**](IpamsvcDHCPPacketStats.md) |  | [optional] 
**HaGroups** | Pointer to [**[]IpamsvcHAGroup**](IpamsvcHAGroup.md) | The list of HA groups. | [optional] 
**Host** | Pointer to [**IpamsvcHost**](IpamsvcHost.md) |  | [optional] 
**Subnets** | Pointer to [**[]IpamsvcSubnet**](IpamsvcSubnet.md) | The list of subnets. | [optional] 

## Methods

### NewIpamsvcHostAssociationsResponse

`func NewIpamsvcHostAssociationsResponse() *IpamsvcHostAssociationsResponse`

NewIpamsvcHostAssociationsResponse instantiates a new IpamsvcHostAssociationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcHostAssociationsResponseWithDefaults

`func NewIpamsvcHostAssociationsResponseWithDefaults() *IpamsvcHostAssociationsResponse`

NewIpamsvcHostAssociationsResponseWithDefaults instantiates a new IpamsvcHostAssociationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpPktStats

`func (o *IpamsvcHostAssociationsResponse) GetDhcpPktStats() IpamsvcDHCPPacketStats`

GetDhcpPktStats returns the DhcpPktStats field if non-nil, zero value otherwise.

### GetDhcpPktStatsOk

`func (o *IpamsvcHostAssociationsResponse) GetDhcpPktStatsOk() (*IpamsvcDHCPPacketStats, bool)`

GetDhcpPktStatsOk returns a tuple with the DhcpPktStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpPktStats

`func (o *IpamsvcHostAssociationsResponse) SetDhcpPktStats(v IpamsvcDHCPPacketStats)`

SetDhcpPktStats sets DhcpPktStats field to given value.

### HasDhcpPktStats

`func (o *IpamsvcHostAssociationsResponse) HasDhcpPktStats() bool`

HasDhcpPktStats returns a boolean if a field has been set.

### GetHaGroups

`func (o *IpamsvcHostAssociationsResponse) GetHaGroups() []IpamsvcHAGroup`

GetHaGroups returns the HaGroups field if non-nil, zero value otherwise.

### GetHaGroupsOk

`func (o *IpamsvcHostAssociationsResponse) GetHaGroupsOk() (*[]IpamsvcHAGroup, bool)`

GetHaGroupsOk returns a tuple with the HaGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaGroups

`func (o *IpamsvcHostAssociationsResponse) SetHaGroups(v []IpamsvcHAGroup)`

SetHaGroups sets HaGroups field to given value.

### HasHaGroups

`func (o *IpamsvcHostAssociationsResponse) HasHaGroups() bool`

HasHaGroups returns a boolean if a field has been set.

### GetHost

`func (o *IpamsvcHostAssociationsResponse) GetHost() IpamsvcHost`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *IpamsvcHostAssociationsResponse) GetHostOk() (*IpamsvcHost, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *IpamsvcHostAssociationsResponse) SetHost(v IpamsvcHost)`

SetHost sets Host field to given value.

### HasHost

`func (o *IpamsvcHostAssociationsResponse) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetSubnets

`func (o *IpamsvcHostAssociationsResponse) GetSubnets() []IpamsvcSubnet`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *IpamsvcHostAssociationsResponse) GetSubnetsOk() (*[]IpamsvcSubnet, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *IpamsvcHostAssociationsResponse) SetSubnets(v []IpamsvcSubnet)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *IpamsvcHostAssociationsResponse) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


