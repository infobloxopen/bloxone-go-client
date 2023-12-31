/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns_config

import (
	"encoding/json"
)

// checks if the ConfigZoneAuthorityMNameBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigZoneAuthorityMNameBlock{}

// ConfigZoneAuthorityMNameBlock Block for fields: _mname_, _protocol_mname_, _use_default_mname_.
type ConfigZoneAuthorityMNameBlock struct {
	// Defaults to empty.
	Mname *string `json:"mname,omitempty"`
	// Optional. Master name server in punycode.  Defaults to empty.
	ProtocolMname *string `json:"protocol_mname,omitempty"`
	// Optional. Use default value for master name server.  Defaults to true.
	UseDefaultMname *bool `json:"use_default_mname,omitempty"`
}

// NewConfigZoneAuthorityMNameBlock instantiates a new ConfigZoneAuthorityMNameBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigZoneAuthorityMNameBlock() *ConfigZoneAuthorityMNameBlock {
	this := ConfigZoneAuthorityMNameBlock{}
	return &this
}

// NewConfigZoneAuthorityMNameBlockWithDefaults instantiates a new ConfigZoneAuthorityMNameBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigZoneAuthorityMNameBlockWithDefaults() *ConfigZoneAuthorityMNameBlock {
	this := ConfigZoneAuthorityMNameBlock{}
	return &this
}

// GetMname returns the Mname field value if set, zero value otherwise.
func (o *ConfigZoneAuthorityMNameBlock) GetMname() string {
	if o == nil || IsNil(o.Mname) {
		var ret string
		return ret
	}
	return *o.Mname
}

// GetMnameOk returns a tuple with the Mname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigZoneAuthorityMNameBlock) GetMnameOk() (*string, bool) {
	if o == nil || IsNil(o.Mname) {
		return nil, false
	}
	return o.Mname, true
}

// HasMname returns a boolean if a field has been set.
func (o *ConfigZoneAuthorityMNameBlock) HasMname() bool {
	if o != nil && !IsNil(o.Mname) {
		return true
	}

	return false
}

// SetMname gets a reference to the given string and assigns it to the Mname field.
func (o *ConfigZoneAuthorityMNameBlock) SetMname(v string) {
	o.Mname = &v
}

// GetProtocolMname returns the ProtocolMname field value if set, zero value otherwise.
func (o *ConfigZoneAuthorityMNameBlock) GetProtocolMname() string {
	if o == nil || IsNil(o.ProtocolMname) {
		var ret string
		return ret
	}
	return *o.ProtocolMname
}

// GetProtocolMnameOk returns a tuple with the ProtocolMname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigZoneAuthorityMNameBlock) GetProtocolMnameOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolMname) {
		return nil, false
	}
	return o.ProtocolMname, true
}

// HasProtocolMname returns a boolean if a field has been set.
func (o *ConfigZoneAuthorityMNameBlock) HasProtocolMname() bool {
	if o != nil && !IsNil(o.ProtocolMname) {
		return true
	}

	return false
}

// SetProtocolMname gets a reference to the given string and assigns it to the ProtocolMname field.
func (o *ConfigZoneAuthorityMNameBlock) SetProtocolMname(v string) {
	o.ProtocolMname = &v
}

// GetUseDefaultMname returns the UseDefaultMname field value if set, zero value otherwise.
func (o *ConfigZoneAuthorityMNameBlock) GetUseDefaultMname() bool {
	if o == nil || IsNil(o.UseDefaultMname) {
		var ret bool
		return ret
	}
	return *o.UseDefaultMname
}

// GetUseDefaultMnameOk returns a tuple with the UseDefaultMname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigZoneAuthorityMNameBlock) GetUseDefaultMnameOk() (*bool, bool) {
	if o == nil || IsNil(o.UseDefaultMname) {
		return nil, false
	}
	return o.UseDefaultMname, true
}

// HasUseDefaultMname returns a boolean if a field has been set.
func (o *ConfigZoneAuthorityMNameBlock) HasUseDefaultMname() bool {
	if o != nil && !IsNil(o.UseDefaultMname) {
		return true
	}

	return false
}

// SetUseDefaultMname gets a reference to the given bool and assigns it to the UseDefaultMname field.
func (o *ConfigZoneAuthorityMNameBlock) SetUseDefaultMname(v bool) {
	o.UseDefaultMname = &v
}

func (o ConfigZoneAuthorityMNameBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigZoneAuthorityMNameBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mname) {
		toSerialize["mname"] = o.Mname
	}
	if !IsNil(o.ProtocolMname) {
		toSerialize["protocol_mname"] = o.ProtocolMname
	}
	if !IsNil(o.UseDefaultMname) {
		toSerialize["use_default_mname"] = o.UseDefaultMname
	}
	return toSerialize, nil
}

type NullableConfigZoneAuthorityMNameBlock struct {
	value *ConfigZoneAuthorityMNameBlock
	isSet bool
}

func (v NullableConfigZoneAuthorityMNameBlock) Get() *ConfigZoneAuthorityMNameBlock {
	return v.value
}

func (v *NullableConfigZoneAuthorityMNameBlock) Set(val *ConfigZoneAuthorityMNameBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigZoneAuthorityMNameBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigZoneAuthorityMNameBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigZoneAuthorityMNameBlock(val *ConfigZoneAuthorityMNameBlock) *NullableConfigZoneAuthorityMNameBlock {
	return &NullableConfigZoneAuthorityMNameBlock{value: val, isSet: true}
}

func (v NullableConfigZoneAuthorityMNameBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigZoneAuthorityMNameBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
