# IpamsvcHAGroupHeartbeats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Peer** | Pointer to **string** | The name of the peer. | [optional] 
**SuccessfulHeartbeat** | Pointer to **string** | The timestamp as a string of the last successful heartbeat received from the peer above. | [optional] 

## Methods

### NewIpamsvcHAGroupHeartbeats

`func NewIpamsvcHAGroupHeartbeats() *IpamsvcHAGroupHeartbeats`

NewIpamsvcHAGroupHeartbeats instantiates a new IpamsvcHAGroupHeartbeats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcHAGroupHeartbeatsWithDefaults

`func NewIpamsvcHAGroupHeartbeatsWithDefaults() *IpamsvcHAGroupHeartbeats`

NewIpamsvcHAGroupHeartbeatsWithDefaults instantiates a new IpamsvcHAGroupHeartbeats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeer

`func (o *IpamsvcHAGroupHeartbeats) GetPeer() string`

GetPeer returns the Peer field if non-nil, zero value otherwise.

### GetPeerOk

`func (o *IpamsvcHAGroupHeartbeats) GetPeerOk() (*string, bool)`

GetPeerOk returns a tuple with the Peer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeer

`func (o *IpamsvcHAGroupHeartbeats) SetPeer(v string)`

SetPeer sets Peer field to given value.

### HasPeer

`func (o *IpamsvcHAGroupHeartbeats) HasPeer() bool`

HasPeer returns a boolean if a field has been set.

### GetSuccessfulHeartbeat

`func (o *IpamsvcHAGroupHeartbeats) GetSuccessfulHeartbeat() string`

GetSuccessfulHeartbeat returns the SuccessfulHeartbeat field if non-nil, zero value otherwise.

### GetSuccessfulHeartbeatOk

`func (o *IpamsvcHAGroupHeartbeats) GetSuccessfulHeartbeatOk() (*string, bool)`

GetSuccessfulHeartbeatOk returns a tuple with the SuccessfulHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulHeartbeat

`func (o *IpamsvcHAGroupHeartbeats) SetSuccessfulHeartbeat(v string)`

SetSuccessfulHeartbeat sets SuccessfulHeartbeat field to given value.

### HasSuccessfulHeartbeat

`func (o *IpamsvcHAGroupHeartbeats) HasSuccessfulHeartbeat() bool`

HasSuccessfulHeartbeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


