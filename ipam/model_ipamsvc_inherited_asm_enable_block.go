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

// checks if the IpamsvcInheritedAsmEnableBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcInheritedAsmEnableBlock{}

// IpamsvcInheritedAsmEnableBlock The inheritance block for ASM fields: _enable_, _enable_notification_, _reenable_date_.
type IpamsvcInheritedAsmEnableBlock struct {
	// The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.
	Action *string `json:"action,omitempty"`
	// The human-readable display name for the object referred to by _source_.
	DisplayName *string `json:"display_name,omitempty"`
	// The resource identifier.
	Source *string `json:"source,omitempty"`
	Value *IpamsvcAsmEnableBlock `json:"value,omitempty"`
}

// NewIpamsvcInheritedAsmEnableBlock instantiates a new IpamsvcInheritedAsmEnableBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcInheritedAsmEnableBlock() *IpamsvcInheritedAsmEnableBlock {
	this := IpamsvcInheritedAsmEnableBlock{}
	return &this
}

// NewIpamsvcInheritedAsmEnableBlockWithDefaults instantiates a new IpamsvcInheritedAsmEnableBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcInheritedAsmEnableBlockWithDefaults() *IpamsvcInheritedAsmEnableBlock {
	this := IpamsvcInheritedAsmEnableBlock{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *IpamsvcInheritedAsmEnableBlock) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcInheritedAsmEnableBlock) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *IpamsvcInheritedAsmEnableBlock) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *IpamsvcInheritedAsmEnableBlock) SetAction(v string) {
	o.Action = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *IpamsvcInheritedAsmEnableBlock) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcInheritedAsmEnableBlock) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *IpamsvcInheritedAsmEnableBlock) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *IpamsvcInheritedAsmEnableBlock) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *IpamsvcInheritedAsmEnableBlock) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcInheritedAsmEnableBlock) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *IpamsvcInheritedAsmEnableBlock) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *IpamsvcInheritedAsmEnableBlock) SetSource(v string) {
	o.Source = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IpamsvcInheritedAsmEnableBlock) GetValue() IpamsvcAsmEnableBlock {
	if o == nil || IsNil(o.Value) {
		var ret IpamsvcAsmEnableBlock
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcInheritedAsmEnableBlock) GetValueOk() (*IpamsvcAsmEnableBlock, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IpamsvcInheritedAsmEnableBlock) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given IpamsvcAsmEnableBlock and assigns it to the Value field.
func (o *IpamsvcInheritedAsmEnableBlock) SetValue(v IpamsvcAsmEnableBlock) {
	o.Value = &v
}

func (o IpamsvcInheritedAsmEnableBlock) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcInheritedAsmEnableBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableIpamsvcInheritedAsmEnableBlock struct {
	value *IpamsvcInheritedAsmEnableBlock
	isSet bool
}

func (v NullableIpamsvcInheritedAsmEnableBlock) Get() *IpamsvcInheritedAsmEnableBlock {
	return v.value
}

func (v *NullableIpamsvcInheritedAsmEnableBlock) Set(val *IpamsvcInheritedAsmEnableBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcInheritedAsmEnableBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcInheritedAsmEnableBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcInheritedAsmEnableBlock(val *IpamsvcInheritedAsmEnableBlock) *NullableIpamsvcInheritedAsmEnableBlock {
	return &NullableIpamsvcInheritedAsmEnableBlock{value: val, isSet: true}
}

func (v NullableIpamsvcInheritedAsmEnableBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcInheritedAsmEnableBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


