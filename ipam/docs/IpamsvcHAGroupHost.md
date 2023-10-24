# IpamsvcHAGroupHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The address on which this host listens. | [optional] 
**Heartbeats** | Pointer to [**[]IpamsvcHAGroupHeartbeats**](IpamsvcHAGroupHeartbeats.md) | Last successful heartbeat received from its peer/s. This field is set when the _collect_stats_ is set to _true_ in the _GET_ _/dhcp/ha_group_ request. | [optional] 
**Host** | **string** | The resource identifier. | 
**Port** | Pointer to **int64** | The HA port. | [optional] [readonly] 
**Role** | Pointer to **string** | The role of this host in the HA relationship: _active_ or _passive_. | [optional] 
**State** | Pointer to **string** | The state of DHCP on the host. This field is set when the _collect_stats_ is set to _true_ in the _GET_ _/dhcp/ha_group_ request. | [optional] 

## Methods

### NewIpamsvcHAGroupHost

`func NewIpamsvcHAGroupHost(host string, ) *IpamsvcHAGroupHost`

NewIpamsvcHAGroupHost instantiates a new IpamsvcHAGroupHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcHAGroupHostWithDefaults

`func NewIpamsvcHAGroupHostWithDefaults() *IpamsvcHAGroupHost`

NewIpamsvcHAGroupHostWithDefaults instantiates a new IpamsvcHAGroupHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *IpamsvcHAGroupHost) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IpamsvcHAGroupHost) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IpamsvcHAGroupHost) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IpamsvcHAGroupHost) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetHeartbeats

`func (o *IpamsvcHAGroupHost) GetHeartbeats() []IpamsvcHAGroupHeartbeats`

GetHeartbeats returns the Heartbeats field if non-nil, zero value otherwise.

### GetHeartbeatsOk

`func (o *IpamsvcHAGroupHost) GetHeartbeatsOk() (*[]IpamsvcHAGroupHeartbeats, bool)`

GetHeartbeatsOk returns a tuple with the Heartbeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeats

`func (o *IpamsvcHAGroupHost) SetHeartbeats(v []IpamsvcHAGroupHeartbeats)`

SetHeartbeats sets Heartbeats field to given value.

### HasHeartbeats

`func (o *IpamsvcHAGroupHost) HasHeartbeats() bool`

HasHeartbeats returns a boolean if a field has been set.

### GetHost

`func (o *IpamsvcHAGroupHost) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *IpamsvcHAGroupHost) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *IpamsvcHAGroupHost) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *IpamsvcHAGroupHost) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IpamsvcHAGroupHost) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IpamsvcHAGroupHost) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *IpamsvcHAGroupHost) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRole

`func (o *IpamsvcHAGroupHost) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *IpamsvcHAGroupHost) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *IpamsvcHAGroupHost) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *IpamsvcHAGroupHost) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetState

`func (o *IpamsvcHAGroupHost) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IpamsvcHAGroupHost) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IpamsvcHAGroupHost) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IpamsvcHAGroupHost) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


