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

// checks if the DHCPOptionsInheritance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DHCPOptionsInheritance{}

// DHCPOptionsInheritance The inheritance configuration that specifies how the _dhcp_options_ field is inherited from the parent object.
type DHCPOptionsInheritance struct {
	// The inheritance configuration for the _dhcp_options_ field.
	DhcpOptions          *InheritedDHCPOptionList `json:"dhcp_options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DHCPOptionsInheritance DHCPOptionsInheritance

// NewDHCPOptionsInheritance instantiates a new DHCPOptionsInheritance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDHCPOptionsInheritance() *DHCPOptionsInheritance {
	this := DHCPOptionsInheritance{}
	return &this
}

// NewDHCPOptionsInheritanceWithDefaults instantiates a new DHCPOptionsInheritance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDHCPOptionsInheritanceWithDefaults() *DHCPOptionsInheritance {
	this := DHCPOptionsInheritance{}
	return &this
}

// GetDhcpOptions returns the DhcpOptions field value if set, zero value otherwise.
func (o *DHCPOptionsInheritance) GetDhcpOptions() InheritedDHCPOptionList {
	if o == nil || IsNil(o.DhcpOptions) {
		var ret InheritedDHCPOptionList
		return ret
	}
	return *o.DhcpOptions
}

// GetDhcpOptionsOk returns a tuple with the DhcpOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPOptionsInheritance) GetDhcpOptionsOk() (*InheritedDHCPOptionList, bool) {
	if o == nil || IsNil(o.DhcpOptions) {
		return nil, false
	}
	return o.DhcpOptions, true
}

// HasDhcpOptions returns a boolean if a field has been set.
func (o *DHCPOptionsInheritance) HasDhcpOptions() bool {
	if o != nil && !IsNil(o.DhcpOptions) {
		return true
	}

	return false
}

// SetDhcpOptions gets a reference to the given InheritedDHCPOptionList and assigns it to the DhcpOptions field.
func (o *DHCPOptionsInheritance) SetDhcpOptions(v InheritedDHCPOptionList) {
	o.DhcpOptions = &v
}

func (o DHCPOptionsInheritance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DHCPOptionsInheritance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DhcpOptions) {
		toSerialize["dhcp_options"] = o.DhcpOptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DHCPOptionsInheritance) UnmarshalJSON(data []byte) (err error) {
	varDHCPOptionsInheritance := _DHCPOptionsInheritance{}

	err = json.Unmarshal(data, &varDHCPOptionsInheritance)

	if err != nil {
		return err
	}

	*o = DHCPOptionsInheritance(varDHCPOptionsInheritance)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dhcp_options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDHCPOptionsInheritance struct {
	value *DHCPOptionsInheritance
	isSet bool
}

func (v NullableDHCPOptionsInheritance) Get() *DHCPOptionsInheritance {
	return v.value
}

func (v *NullableDHCPOptionsInheritance) Set(val *DHCPOptionsInheritance) {
	v.value = val
	v.isSet = true
}

func (v NullableDHCPOptionsInheritance) IsSet() bool {
	return v.isSet
}

func (v *NullableDHCPOptionsInheritance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDHCPOptionsInheritance(val *DHCPOptionsInheritance) *NullableDHCPOptionsInheritance {
	return &NullableDHCPOptionsInheritance{value: val, isSet: true}
}

func (v NullableDHCPOptionsInheritance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDHCPOptionsInheritance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
