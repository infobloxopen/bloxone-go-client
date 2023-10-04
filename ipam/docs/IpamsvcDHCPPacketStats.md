# IpamsvcDHCPPacketStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpPktReceived** | Pointer to **string** | The number of DHCP packets received. | [optional] [readonly] 
**DhcpPktReceivedV6** | Pointer to **string** | The number of DHCP V6 packets received. | [optional] [readonly] 
**DhcpPktSent** | Pointer to **string** | The number of DHCP packets sent. | [optional] [readonly] 
**DhcpPktSentV6** | Pointer to **string** | The number of DHCP V6 packets sent. | [optional] [readonly] 
**DhcpReqReceived** | Pointer to **string** | The number of DHCP requests received. | [optional] [readonly] 
**DhcpReqReceivedV6** | Pointer to **string** | The number of DHCP V6 requests received. | [optional] [readonly] 

## Methods

### NewIpamsvcDHCPPacketStats

`func NewIpamsvcDHCPPacketStats() *IpamsvcDHCPPacketStats`

NewIpamsvcDHCPPacketStats instantiates a new IpamsvcDHCPPacketStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcDHCPPacketStatsWithDefaults

`func NewIpamsvcDHCPPacketStatsWithDefaults() *IpamsvcDHCPPacketStats`

NewIpamsvcDHCPPacketStatsWithDefaults instantiates a new IpamsvcDHCPPacketStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpPktReceived

`func (o *IpamsvcDHCPPacketStats) GetDhcpPktReceived() string`

GetDhcpPktReceived returns the DhcpPktReceived field if non-nil, zero value otherwise.

### GetDhcpPktReceivedOk

`func (o *IpamsvcDHCPPacketStats) GetDhcpPktReceivedOk() (*string, bool)`

GetDhcpPktReceivedOk returns a tuple with the DhcpPktReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpPktReceived

`func (o *IpamsvcDHCPPacketStats) SetDhcpPktReceived(v string)`

SetDhcpPktReceived sets DhcpPktReceived field to given value.

### HasDhcpPktReceived

`func (o *IpamsvcDHCPPacketStats) HasDhcpPktReceived() bool`

HasDhcpPktReceived returns a boolean if a field has been set.

### GetDhcpPktReceivedV6

`func (o *IpamsvcDHCPPacketStats) GetDhcpPktReceivedV6() string`

GetDhcpPktReceivedV6 returns the DhcpPktReceivedV6 field if non-nil, zero value otherwise.

### GetDhcpPktReceivedV6Ok

`func (o *IpamsvcDHCPPacketStats) GetDhcpPktReceivedV6Ok() (*string, bool)`

GetDhcpPktReceivedV6Ok returns a tuple with the DhcpPktReceivedV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpPktReceivedV6

`func (o *IpamsvcDHCPPacketStats) SetDhcpPktReceivedV6(v string)`

SetDhcpPktReceivedV6 sets DhcpPktReceivedV6 field to given value.

### HasDhcpPktReceivedV6

`func (o *IpamsvcDHCPPacketStats) HasDhcpPktReceivedV6() bool`

HasDhcpPktReceivedV6 returns a boolean if a field has been set.

### GetDhcpPktSent

`func (o *IpamsvcDHCPPacketStats) GetDhcpPktSent() string`

GetDhcpPktSent returns the DhcpPktSent field if non-nil, zero value otherwise.

### GetDhcpPktSentOk

`func (o *IpamsvcDHCPPacketStats) GetDhcpPktSentOk() (*string, bool)`

GetDhcpPktSentOk returns a tuple with the DhcpPktSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpPktSent

`func (o *IpamsvcDHCPPacketStats) SetDhcpPktSent(v string)`

SetDhcpPktSent sets DhcpPktSent field to given value.

### HasDhcpPktSent

`func (o *IpamsvcDHCPPacketStats) HasDhcpPktSent() bool`

HasDhcpPktSent returns a boolean if a field has been set.

### GetDhcpPktSentV6

`func (o *IpamsvcDHCPPacketStats) GetDhcpPktSentV6() string`

GetDhcpPktSentV6 returns the DhcpPktSentV6 field if non-nil, zero value otherwise.

### GetDhcpPktSentV6Ok

`func (o *IpamsvcDHCPPacketStats) GetDhcpPktSentV6Ok() (*string, bool)`

GetDhcpPktSentV6Ok returns a tuple with the DhcpPktSentV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpPktSentV6

`func (o *IpamsvcDHCPPacketStats) SetDhcpPktSentV6(v string)`

SetDhcpPktSentV6 sets DhcpPktSentV6 field to given value.

### HasDhcpPktSentV6

`func (o *IpamsvcDHCPPacketStats) HasDhcpPktSentV6() bool`

HasDhcpPktSentV6 returns a boolean if a field has been set.

### GetDhcpReqReceived

`func (o *IpamsvcDHCPPacketStats) GetDhcpReqReceived() string`

GetDhcpReqReceived returns the DhcpReqReceived field if non-nil, zero value otherwise.

### GetDhcpReqReceivedOk

`func (o *IpamsvcDHCPPacketStats) GetDhcpReqReceivedOk() (*string, bool)`

GetDhcpReqReceivedOk returns a tuple with the DhcpReqReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpReqReceived

`func (o *IpamsvcDHCPPacketStats) SetDhcpReqReceived(v string)`

SetDhcpReqReceived sets DhcpReqReceived field to given value.

### HasDhcpReqReceived

`func (o *IpamsvcDHCPPacketStats) HasDhcpReqReceived() bool`

HasDhcpReqReceived returns a boolean if a field has been set.

### GetDhcpReqReceivedV6

`func (o *IpamsvcDHCPPacketStats) GetDhcpReqReceivedV6() string`

GetDhcpReqReceivedV6 returns the DhcpReqReceivedV6 field if non-nil, zero value otherwise.

### GetDhcpReqReceivedV6Ok

`func (o *IpamsvcDHCPPacketStats) GetDhcpReqReceivedV6Ok() (*string, bool)`

GetDhcpReqReceivedV6Ok returns a tuple with the DhcpReqReceivedV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpReqReceivedV6

`func (o *IpamsvcDHCPPacketStats) SetDhcpReqReceivedV6(v string)`

SetDhcpReqReceivedV6 sets DhcpReqReceivedV6 field to given value.

### HasDhcpReqReceivedV6

`func (o *IpamsvcDHCPPacketStats) HasDhcpReqReceivedV6() bool`

HasDhcpReqReceivedV6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


