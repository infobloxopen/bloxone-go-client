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

// checks if the ConfigExternalPrimary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigExternalPrimary{}

// ConfigExternalPrimary External DNS primary.
type ConfigExternalPrimary struct {
	// Optional. Required only if _type_ is _server_. IP Address of nameserver.
	Address *string `json:"address,omitempty"`
	// Optional. Required only if _type_ is _server_. FQDN of nameserver.
	Fqdn *string `json:"fqdn,omitempty"`
	// The resource identifier.
	Nsg *string `json:"nsg,omitempty"`
	// FQDN of nameserver in punycode.
	ProtocolFqdn *string `json:"protocol_fqdn,omitempty"`
	// Optional. If enabled, secondaries will use the configured TSIG key when requesting a zone transfer from this primary.
	TsigEnabled *bool          `json:"tsig_enabled,omitempty"`
	TsigKey     *ConfigTSIGKey `json:"tsig_key,omitempty"`
	// Allowed values: * _nsg_, * _primary_.
	Type string `json:"type"`
}

// NewConfigExternalPrimary instantiates a new ConfigExternalPrimary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigExternalPrimary(type_ string) *ConfigExternalPrimary {
	this := ConfigExternalPrimary{}
	this.Type = type_
	return &this
}

// NewConfigExternalPrimaryWithDefaults instantiates a new ConfigExternalPrimary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigExternalPrimaryWithDefaults() *ConfigExternalPrimary {
	this := ConfigExternalPrimary{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ConfigExternalPrimary) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExternalPrimary) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ConfigExternalPrimary) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ConfigExternalPrimary) SetAddress(v string) {
	o.Address = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *ConfigExternalPrimary) GetFqdn() string {
	if o == nil || IsNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExternalPrimary) GetFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqdn) {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *ConfigExternalPrimary) HasFqdn() bool {
	if o != nil && !IsNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *ConfigExternalPrimary) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetNsg returns the Nsg field value if set, zero value otherwise.
func (o *ConfigExternalPrimary) GetNsg() string {
	if o == nil || IsNil(o.Nsg) {
		var ret string
		return ret
	}
	return *o.Nsg
}

// GetNsgOk returns a tuple with the Nsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExternalPrimary) GetNsgOk() (*string, bool) {
	if o == nil || IsNil(o.Nsg) {
		return nil, false
	}
	return o.Nsg, true
}

// HasNsg returns a boolean if a field has been set.
func (o *ConfigExternalPrimary) HasNsg() bool {
	if o != nil && !IsNil(o.Nsg) {
		return true
	}

	return false
}

// SetNsg gets a reference to the given string and assigns it to the Nsg field.
func (o *ConfigExternalPrimary) SetNsg(v string) {
	o.Nsg = &v
}

// GetProtocolFqdn returns the ProtocolFqdn field value if set, zero value otherwise.
func (o *ConfigExternalPrimary) GetProtocolFqdn() string {
	if o == nil || IsNil(o.ProtocolFqdn) {
		var ret string
		return ret
	}
	return *o.ProtocolFqdn
}

// GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExternalPrimary) GetProtocolFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolFqdn) {
		return nil, false
	}
	return o.ProtocolFqdn, true
}

// HasProtocolFqdn returns a boolean if a field has been set.
func (o *ConfigExternalPrimary) HasProtocolFqdn() bool {
	if o != nil && !IsNil(o.ProtocolFqdn) {
		return true
	}

	return false
}

// SetProtocolFqdn gets a reference to the given string and assigns it to the ProtocolFqdn field.
func (o *ConfigExternalPrimary) SetProtocolFqdn(v string) {
	o.ProtocolFqdn = &v
}

// GetTsigEnabled returns the TsigEnabled field value if set, zero value otherwise.
func (o *ConfigExternalPrimary) GetTsigEnabled() bool {
	if o == nil || IsNil(o.TsigEnabled) {
		var ret bool
		return ret
	}
	return *o.TsigEnabled
}

// GetTsigEnabledOk returns a tuple with the TsigEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExternalPrimary) GetTsigEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TsigEnabled) {
		return nil, false
	}
	return o.TsigEnabled, true
}

// HasTsigEnabled returns a boolean if a field has been set.
func (o *ConfigExternalPrimary) HasTsigEnabled() bool {
	if o != nil && !IsNil(o.TsigEnabled) {
		return true
	}

	return false
}

// SetTsigEnabled gets a reference to the given bool and assigns it to the TsigEnabled field.
func (o *ConfigExternalPrimary) SetTsigEnabled(v bool) {
	o.TsigEnabled = &v
}

// GetTsigKey returns the TsigKey field value if set, zero value otherwise.
func (o *ConfigExternalPrimary) GetTsigKey() ConfigTSIGKey {
	if o == nil || IsNil(o.TsigKey) {
		var ret ConfigTSIGKey
		return ret
	}
	return *o.TsigKey
}

// GetTsigKeyOk returns a tuple with the TsigKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExternalPrimary) GetTsigKeyOk() (*ConfigTSIGKey, bool) {
	if o == nil || IsNil(o.TsigKey) {
		return nil, false
	}
	return o.TsigKey, true
}

// HasTsigKey returns a boolean if a field has been set.
func (o *ConfigExternalPrimary) HasTsigKey() bool {
	if o != nil && !IsNil(o.TsigKey) {
		return true
	}

	return false
}

// SetTsigKey gets a reference to the given ConfigTSIGKey and assigns it to the TsigKey field.
func (o *ConfigExternalPrimary) SetTsigKey(v ConfigTSIGKey) {
	o.TsigKey = &v
}

// GetType returns the Type field value
func (o *ConfigExternalPrimary) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConfigExternalPrimary) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConfigExternalPrimary) SetType(v string) {
	o.Type = v
}

func (o ConfigExternalPrimary) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigExternalPrimary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !IsNil(o.Nsg) {
		toSerialize["nsg"] = o.Nsg
	}
	if !IsNil(o.ProtocolFqdn) {
		toSerialize["protocol_fqdn"] = o.ProtocolFqdn
	}
	if !IsNil(o.TsigEnabled) {
		toSerialize["tsig_enabled"] = o.TsigEnabled
	}
	if !IsNil(o.TsigKey) {
		toSerialize["tsig_key"] = o.TsigKey
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableConfigExternalPrimary struct {
	value *ConfigExternalPrimary
	isSet bool
}

func (v NullableConfigExternalPrimary) Get() *ConfigExternalPrimary {
	return v.value
}

func (v *NullableConfigExternalPrimary) Set(val *ConfigExternalPrimary) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigExternalPrimary) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigExternalPrimary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigExternalPrimary(val *ConfigExternalPrimary) *NullableConfigExternalPrimary {
	return &NullableConfigExternalPrimary{value: val, isSet: true}
}

func (v NullableConfigExternalPrimary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigExternalPrimary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
