# HAGroupHeartbeats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Peer** | Pointer to **string** | The name of the peer. | [optional] 
**SuccessfulHeartbeat** | Pointer to **string** | The timestamp as a string of the last successful heartbeat received from the peer above. | [optional] 
**SuccessfulHeartbeatV6** | Pointer to **string** | The timestamp as a string of the last successful DHCPv6 heartbeat received from the peer above. | [optional] 

## Methods

### NewHAGroupHeartbeats

`func NewHAGroupHeartbeats() *HAGroupHeartbeats`

NewHAGroupHeartbeats instantiates a new HAGroupHeartbeats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHAGroupHeartbeatsWithDefaults

`func NewHAGroupHeartbeatsWithDefaults() *HAGroupHeartbeats`

NewHAGroupHeartbeatsWithDefaults instantiates a new HAGroupHeartbeats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeer

`func (o *HAGroupHeartbeats) GetPeer() string`

GetPeer returns the Peer field if non-nil, zero value otherwise.

### GetPeerOk

`func (o *HAGroupHeartbeats) GetPeerOk() (*string, bool)`

GetPeerOk returns a tuple with the Peer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeer

`func (o *HAGroupHeartbeats) SetPeer(v string)`

SetPeer sets Peer field to given value.

### HasPeer

`func (o *HAGroupHeartbeats) HasPeer() bool`

HasPeer returns a boolean if a field has been set.

### GetSuccessfulHeartbeat

`func (o *HAGroupHeartbeats) GetSuccessfulHeartbeat() string`

GetSuccessfulHeartbeat returns the SuccessfulHeartbeat field if non-nil, zero value otherwise.

### GetSuccessfulHeartbeatOk

`func (o *HAGroupHeartbeats) GetSuccessfulHeartbeatOk() (*string, bool)`

GetSuccessfulHeartbeatOk returns a tuple with the SuccessfulHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulHeartbeat

`func (o *HAGroupHeartbeats) SetSuccessfulHeartbeat(v string)`

SetSuccessfulHeartbeat sets SuccessfulHeartbeat field to given value.

### HasSuccessfulHeartbeat

`func (o *HAGroupHeartbeats) HasSuccessfulHeartbeat() bool`

HasSuccessfulHeartbeat returns a boolean if a field has been set.

### GetSuccessfulHeartbeatV6

`func (o *HAGroupHeartbeats) GetSuccessfulHeartbeatV6() string`

GetSuccessfulHeartbeatV6 returns the SuccessfulHeartbeatV6 field if non-nil, zero value otherwise.

### GetSuccessfulHeartbeatV6Ok

`func (o *HAGroupHeartbeats) GetSuccessfulHeartbeatV6Ok() (*string, bool)`

GetSuccessfulHeartbeatV6Ok returns a tuple with the SuccessfulHeartbeatV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulHeartbeatV6

`func (o *HAGroupHeartbeats) SetSuccessfulHeartbeatV6(v string)`

SetSuccessfulHeartbeatV6 sets SuccessfulHeartbeatV6 field to given value.

### HasSuccessfulHeartbeatV6

`func (o *HAGroupHeartbeats) HasSuccessfulHeartbeatV6() bool`

HasSuccessfulHeartbeatV6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


