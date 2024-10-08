/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the HAGroupHeartbeats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HAGroupHeartbeats{}

// HAGroupHeartbeats struct for HAGroupHeartbeats
type HAGroupHeartbeats struct {
	// The name of the peer.
	Peer *string `json:"peer,omitempty"`
	// The timestamp as a string of the last successful heartbeat received from the peer above.
	SuccessfulHeartbeat *string `json:"successful_heartbeat,omitempty"`
	// The timestamp as a string of the last successful DHCPv6 heartbeat received from the peer above.
	SuccessfulHeartbeatV6 *string `json:"successful_heartbeat_v6,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _HAGroupHeartbeats HAGroupHeartbeats

// NewHAGroupHeartbeats instantiates a new HAGroupHeartbeats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHAGroupHeartbeats() *HAGroupHeartbeats {
	this := HAGroupHeartbeats{}
	return &this
}

// NewHAGroupHeartbeatsWithDefaults instantiates a new HAGroupHeartbeats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHAGroupHeartbeatsWithDefaults() *HAGroupHeartbeats {
	this := HAGroupHeartbeats{}
	return &this
}

// GetPeer returns the Peer field value if set, zero value otherwise.
func (o *HAGroupHeartbeats) GetPeer() string {
	if o == nil || IsNil(o.Peer) {
		var ret string
		return ret
	}
	return *o.Peer
}

// GetPeerOk returns a tuple with the Peer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HAGroupHeartbeats) GetPeerOk() (*string, bool) {
	if o == nil || IsNil(o.Peer) {
		return nil, false
	}
	return o.Peer, true
}

// HasPeer returns a boolean if a field has been set.
func (o *HAGroupHeartbeats) HasPeer() bool {
	if o != nil && !IsNil(o.Peer) {
		return true
	}

	return false
}

// SetPeer gets a reference to the given string and assigns it to the Peer field.
func (o *HAGroupHeartbeats) SetPeer(v string) {
	o.Peer = &v
}

// GetSuccessfulHeartbeat returns the SuccessfulHeartbeat field value if set, zero value otherwise.
func (o *HAGroupHeartbeats) GetSuccessfulHeartbeat() string {
	if o == nil || IsNil(o.SuccessfulHeartbeat) {
		var ret string
		return ret
	}
	return *o.SuccessfulHeartbeat
}

// GetSuccessfulHeartbeatOk returns a tuple with the SuccessfulHeartbeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HAGroupHeartbeats) GetSuccessfulHeartbeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuccessfulHeartbeat) {
		return nil, false
	}
	return o.SuccessfulHeartbeat, true
}

// HasSuccessfulHeartbeat returns a boolean if a field has been set.
func (o *HAGroupHeartbeats) HasSuccessfulHeartbeat() bool {
	if o != nil && !IsNil(o.SuccessfulHeartbeat) {
		return true
	}

	return false
}

// SetSuccessfulHeartbeat gets a reference to the given string and assigns it to the SuccessfulHeartbeat field.
func (o *HAGroupHeartbeats) SetSuccessfulHeartbeat(v string) {
	o.SuccessfulHeartbeat = &v
}

// GetSuccessfulHeartbeatV6 returns the SuccessfulHeartbeatV6 field value if set, zero value otherwise.
func (o *HAGroupHeartbeats) GetSuccessfulHeartbeatV6() string {
	if o == nil || IsNil(o.SuccessfulHeartbeatV6) {
		var ret string
		return ret
	}
	return *o.SuccessfulHeartbeatV6
}

// GetSuccessfulHeartbeatV6Ok returns a tuple with the SuccessfulHeartbeatV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HAGroupHeartbeats) GetSuccessfulHeartbeatV6Ok() (*string, bool) {
	if o == nil || IsNil(o.SuccessfulHeartbeatV6) {
		return nil, false
	}
	return o.SuccessfulHeartbeatV6, true
}

// HasSuccessfulHeartbeatV6 returns a boolean if a field has been set.
func (o *HAGroupHeartbeats) HasSuccessfulHeartbeatV6() bool {
	if o != nil && !IsNil(o.SuccessfulHeartbeatV6) {
		return true
	}

	return false
}

// SetSuccessfulHeartbeatV6 gets a reference to the given string and assigns it to the SuccessfulHeartbeatV6 field.
func (o *HAGroupHeartbeats) SetSuccessfulHeartbeatV6(v string) {
	o.SuccessfulHeartbeatV6 = &v
}

func (o HAGroupHeartbeats) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HAGroupHeartbeats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Peer) {
		toSerialize["peer"] = o.Peer
	}
	if !IsNil(o.SuccessfulHeartbeat) {
		toSerialize["successful_heartbeat"] = o.SuccessfulHeartbeat
	}
	if !IsNil(o.SuccessfulHeartbeatV6) {
		toSerialize["successful_heartbeat_v6"] = o.SuccessfulHeartbeatV6
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HAGroupHeartbeats) UnmarshalJSON(data []byte) (err error) {
	varHAGroupHeartbeats := _HAGroupHeartbeats{}

	err = json.Unmarshal(data, &varHAGroupHeartbeats)

	if err != nil {
		return err
	}

	*o = HAGroupHeartbeats(varHAGroupHeartbeats)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "peer")
		delete(additionalProperties, "successful_heartbeat")
		delete(additionalProperties, "successful_heartbeat_v6")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHAGroupHeartbeats struct {
	value *HAGroupHeartbeats
	isSet bool
}

func (v NullableHAGroupHeartbeats) Get() *HAGroupHeartbeats {
	return v.value
}

func (v *NullableHAGroupHeartbeats) Set(val *HAGroupHeartbeats) {
	v.value = val
	v.isSet = true
}

func (v NullableHAGroupHeartbeats) IsSet() bool {
	return v.isSet
}

func (v *NullableHAGroupHeartbeats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHAGroupHeartbeats(val *HAGroupHeartbeats) *NullableHAGroupHeartbeats {
	return &NullableHAGroupHeartbeats{value: val, isSet: true}
}

func (v NullableHAGroupHeartbeats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHAGroupHeartbeats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
