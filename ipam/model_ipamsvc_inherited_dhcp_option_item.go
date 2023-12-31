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

// checks if the IpamsvcInheritedDHCPOptionItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcInheritedDHCPOptionItem{}

// IpamsvcInheritedDHCPOptionItem A wrapper of item (_dhcp/option_item_) in a list of Inherited DHCP options. It contains extra fields not covered by OptionItem.
type IpamsvcInheritedDHCPOptionItem struct {
	Option *IpamsvcOptionItem `json:"option,omitempty"`
	// The resource identifier.
	OverridingGroup *string `json:"overriding_group,omitempty"`
}

// NewIpamsvcInheritedDHCPOptionItem instantiates a new IpamsvcInheritedDHCPOptionItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcInheritedDHCPOptionItem() *IpamsvcInheritedDHCPOptionItem {
	this := IpamsvcInheritedDHCPOptionItem{}
	return &this
}

// NewIpamsvcInheritedDHCPOptionItemWithDefaults instantiates a new IpamsvcInheritedDHCPOptionItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcInheritedDHCPOptionItemWithDefaults() *IpamsvcInheritedDHCPOptionItem {
	this := IpamsvcInheritedDHCPOptionItem{}
	return &this
}

// GetOption returns the Option field value if set, zero value otherwise.
func (o *IpamsvcInheritedDHCPOptionItem) GetOption() IpamsvcOptionItem {
	if o == nil || IsNil(o.Option) {
		var ret IpamsvcOptionItem
		return ret
	}
	return *o.Option
}

// GetOptionOk returns a tuple with the Option field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcInheritedDHCPOptionItem) GetOptionOk() (*IpamsvcOptionItem, bool) {
	if o == nil || IsNil(o.Option) {
		return nil, false
	}
	return o.Option, true
}

// HasOption returns a boolean if a field has been set.
func (o *IpamsvcInheritedDHCPOptionItem) HasOption() bool {
	if o != nil && !IsNil(o.Option) {
		return true
	}

	return false
}

// SetOption gets a reference to the given IpamsvcOptionItem and assigns it to the Option field.
func (o *IpamsvcInheritedDHCPOptionItem) SetOption(v IpamsvcOptionItem) {
	o.Option = &v
}

// GetOverridingGroup returns the OverridingGroup field value if set, zero value otherwise.
func (o *IpamsvcInheritedDHCPOptionItem) GetOverridingGroup() string {
	if o == nil || IsNil(o.OverridingGroup) {
		var ret string
		return ret
	}
	return *o.OverridingGroup
}

// GetOverridingGroupOk returns a tuple with the OverridingGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcInheritedDHCPOptionItem) GetOverridingGroupOk() (*string, bool) {
	if o == nil || IsNil(o.OverridingGroup) {
		return nil, false
	}
	return o.OverridingGroup, true
}

// HasOverridingGroup returns a boolean if a field has been set.
func (o *IpamsvcInheritedDHCPOptionItem) HasOverridingGroup() bool {
	if o != nil && !IsNil(o.OverridingGroup) {
		return true
	}

	return false
}

// SetOverridingGroup gets a reference to the given string and assigns it to the OverridingGroup field.
func (o *IpamsvcInheritedDHCPOptionItem) SetOverridingGroup(v string) {
	o.OverridingGroup = &v
}

func (o IpamsvcInheritedDHCPOptionItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcInheritedDHCPOptionItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Option) {
		toSerialize["option"] = o.Option
	}
	if !IsNil(o.OverridingGroup) {
		toSerialize["overriding_group"] = o.OverridingGroup
	}
	return toSerialize, nil
}

type NullableIpamsvcInheritedDHCPOptionItem struct {
	value *IpamsvcInheritedDHCPOptionItem
	isSet bool
}

func (v NullableIpamsvcInheritedDHCPOptionItem) Get() *IpamsvcInheritedDHCPOptionItem {
	return v.value
}

func (v *NullableIpamsvcInheritedDHCPOptionItem) Set(val *IpamsvcInheritedDHCPOptionItem) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcInheritedDHCPOptionItem) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcInheritedDHCPOptionItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcInheritedDHCPOptionItem(val *IpamsvcInheritedDHCPOptionItem) *NullableIpamsvcInheritedDHCPOptionItem {
	return &NullableIpamsvcInheritedDHCPOptionItem{value: val, isSet: true}
}

func (v NullableIpamsvcInheritedDHCPOptionItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcInheritedDHCPOptionItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
