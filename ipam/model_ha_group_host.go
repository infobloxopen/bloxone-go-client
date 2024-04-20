/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the HAGroupHost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HAGroupHost{}

// HAGroupHost An __HAGroupHost__ object (_dhcp/ha_group_host_) represents an on-prem host belonging to an HA Group.
type HAGroupHost struct {
	// The address on which this host listens.
	Address *string `json:"address,omitempty"`
	// Last successful heartbeat received from its peer/s. This field is set when the _collect_stats_ is set to _true_ in the _GET_ _/dhcp/ha_group_ request.
	Heartbeats []HAGroupHeartbeats `json:"heartbeats,omitempty"`
	// The resource identifier.
	Host string `json:"host"`
	// The HA port.
	Port *int64 `json:"port,omitempty"`
	// The role of this host in the HA relationship: _active_ or _passive_.
	Role *string `json:"role,omitempty"`
	// The state of DHCP on the host. This field is set when the _collect_stats_ is set to _true_ in the _GET_ _/dhcp/ha_group_ request.
	State *string `json:"state,omitempty"`
}

type _HAGroupHost HAGroupHost

// NewHAGroupHost instantiates a new HAGroupHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHAGroupHost(host string) *HAGroupHost {
	this := HAGroupHost{}
	this.Host = host
	return &this
}

// NewHAGroupHostWithDefaults instantiates a new HAGroupHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHAGroupHostWithDefaults() *HAGroupHost {
	this := HAGroupHost{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *HAGroupHost) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HAGroupHost) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *HAGroupHost) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *HAGroupHost) SetAddress(v string) {
	o.Address = &v
}

// GetHeartbeats returns the Heartbeats field value if set, zero value otherwise.
func (o *HAGroupHost) GetHeartbeats() []HAGroupHeartbeats {
	if o == nil || IsNil(o.Heartbeats) {
		var ret []HAGroupHeartbeats
		return ret
	}
	return o.Heartbeats
}

// GetHeartbeatsOk returns a tuple with the Heartbeats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HAGroupHost) GetHeartbeatsOk() ([]HAGroupHeartbeats, bool) {
	if o == nil || IsNil(o.Heartbeats) {
		return nil, false
	}
	return o.Heartbeats, true
}

// HasHeartbeats returns a boolean if a field has been set.
func (o *HAGroupHost) HasHeartbeats() bool {
	if o != nil && !IsNil(o.Heartbeats) {
		return true
	}

	return false
}

// SetHeartbeats gets a reference to the given []HAGroupHeartbeats and assigns it to the Heartbeats field.
func (o *HAGroupHost) SetHeartbeats(v []HAGroupHeartbeats) {
	o.Heartbeats = v
}

// GetHost returns the Host field value
func (o *HAGroupHost) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *HAGroupHost) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *HAGroupHost) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *HAGroupHost) GetPort() int64 {
	if o == nil || IsNil(o.Port) {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HAGroupHost) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *HAGroupHost) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *HAGroupHost) SetPort(v int64) {
	o.Port = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *HAGroupHost) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HAGroupHost) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *HAGroupHost) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *HAGroupHost) SetRole(v string) {
	o.Role = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *HAGroupHost) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HAGroupHost) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *HAGroupHost) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *HAGroupHost) SetState(v string) {
	o.State = &v
}

func (o HAGroupHost) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HAGroupHost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Heartbeats) {
		toSerialize["heartbeats"] = o.Heartbeats
	}
	toSerialize["host"] = o.Host
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

func (o *HAGroupHost) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"host",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varHAGroupHost := _HAGroupHost{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHAGroupHost)

	if err != nil {
		return err
	}

	*o = HAGroupHost(varHAGroupHost)

	return err
}

type NullableHAGroupHost struct {
	value *HAGroupHost
	isSet bool
}

func (v NullableHAGroupHost) Get() *HAGroupHost {
	return v.value
}

func (v *NullableHAGroupHost) Set(val *HAGroupHost) {
	v.value = val
	v.isSet = true
}

func (v NullableHAGroupHost) IsSet() bool {
	return v.isSet
}

func (v *NullableHAGroupHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHAGroupHost(val *HAGroupHost) *NullableHAGroupHost {
	return &NullableHAGroupHost{value: val, isSet: true}
}

func (v NullableHAGroupHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHAGroupHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
