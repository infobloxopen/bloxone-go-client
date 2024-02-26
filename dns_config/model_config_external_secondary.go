/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.   

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns_config

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ConfigExternalSecondary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigExternalSecondary{}

// ConfigExternalSecondary External DNS secondary.
type ConfigExternalSecondary struct {
	// IP Address of nameserver.
	Address string `json:"address"`
	// FQDN of nameserver.
	Fqdn string `json:"fqdn"`
	// FQDN of nameserver in punycode.
	ProtocolFqdn *string `json:"protocol_fqdn,omitempty"`
	// If enabled, the NS record and glue record will NOT be automatically generated according to secondaries nameserver assignment.  Default: _false_
	Stealth *bool `json:"stealth,omitempty"`
	// If enabled, secondaries will use the configured TSIG key when requesting a zone transfer.  Default: _false_
	TsigEnabled *bool `json:"tsig_enabled,omitempty"`
	TsigKey *ConfigTSIGKey `json:"tsig_key,omitempty"`
}

type _ConfigExternalSecondary ConfigExternalSecondary

// NewConfigExternalSecondary instantiates a new ConfigExternalSecondary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigExternalSecondary(address string, fqdn string) *ConfigExternalSecondary {
	this := ConfigExternalSecondary{}
	this.Address = address
	this.Fqdn = fqdn
	return &this
}

// NewConfigExternalSecondaryWithDefaults instantiates a new ConfigExternalSecondary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigExternalSecondaryWithDefaults() *ConfigExternalSecondary {
	this := ConfigExternalSecondary{}
	return &this
}

// GetAddress returns the Address field value
func (o *ConfigExternalSecondary) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ConfigExternalSecondary) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ConfigExternalSecondary) SetAddress(v string) {
	o.Address = v
}

// GetFqdn returns the Fqdn field value
func (o *ConfigExternalSecondary) GetFqdn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value
// and a boolean to check if the value has been set.
func (o *ConfigExternalSecondary) GetFqdnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fqdn, true
}

// SetFqdn sets field value
func (o *ConfigExternalSecondary) SetFqdn(v string) {
	o.Fqdn = v
}

// GetProtocolFqdn returns the ProtocolFqdn field value if set, zero value otherwise.
func (o *ConfigExternalSecondary) GetProtocolFqdn() string {
	if o == nil || IsNil(o.ProtocolFqdn) {
		var ret string
		return ret
	}
	return *o.ProtocolFqdn
}

// GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExternalSecondary) GetProtocolFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolFqdn) {
		return nil, false
	}
	return o.ProtocolFqdn, true
}

// HasProtocolFqdn returns a boolean if a field has been set.
func (o *ConfigExternalSecondary) HasProtocolFqdn() bool {
	if o != nil && !IsNil(o.ProtocolFqdn) {
		return true
	}

	return false
}

// SetProtocolFqdn gets a reference to the given string and assigns it to the ProtocolFqdn field.
func (o *ConfigExternalSecondary) SetProtocolFqdn(v string) {
	o.ProtocolFqdn = &v
}

// GetStealth returns the Stealth field value if set, zero value otherwise.
func (o *ConfigExternalSecondary) GetStealth() bool {
	if o == nil || IsNil(o.Stealth) {
		var ret bool
		return ret
	}
	return *o.Stealth
}

// GetStealthOk returns a tuple with the Stealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExternalSecondary) GetStealthOk() (*bool, bool) {
	if o == nil || IsNil(o.Stealth) {
		return nil, false
	}
	return o.Stealth, true
}

// HasStealth returns a boolean if a field has been set.
func (o *ConfigExternalSecondary) HasStealth() bool {
	if o != nil && !IsNil(o.Stealth) {
		return true
	}

	return false
}

// SetStealth gets a reference to the given bool and assigns it to the Stealth field.
func (o *ConfigExternalSecondary) SetStealth(v bool) {
	o.Stealth = &v
}

// GetTsigEnabled returns the TsigEnabled field value if set, zero value otherwise.
func (o *ConfigExternalSecondary) GetTsigEnabled() bool {
	if o == nil || IsNil(o.TsigEnabled) {
		var ret bool
		return ret
	}
	return *o.TsigEnabled
}

// GetTsigEnabledOk returns a tuple with the TsigEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExternalSecondary) GetTsigEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TsigEnabled) {
		return nil, false
	}
	return o.TsigEnabled, true
}

// HasTsigEnabled returns a boolean if a field has been set.
func (o *ConfigExternalSecondary) HasTsigEnabled() bool {
	if o != nil && !IsNil(o.TsigEnabled) {
		return true
	}

	return false
}

// SetTsigEnabled gets a reference to the given bool and assigns it to the TsigEnabled field.
func (o *ConfigExternalSecondary) SetTsigEnabled(v bool) {
	o.TsigEnabled = &v
}

// GetTsigKey returns the TsigKey field value if set, zero value otherwise.
func (o *ConfigExternalSecondary) GetTsigKey() ConfigTSIGKey {
	if o == nil || IsNil(o.TsigKey) {
		var ret ConfigTSIGKey
		return ret
	}
	return *o.TsigKey
}

// GetTsigKeyOk returns a tuple with the TsigKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigExternalSecondary) GetTsigKeyOk() (*ConfigTSIGKey, bool) {
	if o == nil || IsNil(o.TsigKey) {
		return nil, false
	}
	return o.TsigKey, true
}

// HasTsigKey returns a boolean if a field has been set.
func (o *ConfigExternalSecondary) HasTsigKey() bool {
	if o != nil && !IsNil(o.TsigKey) {
		return true
	}

	return false
}

// SetTsigKey gets a reference to the given ConfigTSIGKey and assigns it to the TsigKey field.
func (o *ConfigExternalSecondary) SetTsigKey(v ConfigTSIGKey) {
	o.TsigKey = &v
}

func (o ConfigExternalSecondary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigExternalSecondary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["fqdn"] = o.Fqdn
	if !IsNil(o.ProtocolFqdn) {
		toSerialize["protocol_fqdn"] = o.ProtocolFqdn
	}
	if !IsNil(o.Stealth) {
		toSerialize["stealth"] = o.Stealth
	}
	if !IsNil(o.TsigEnabled) {
		toSerialize["tsig_enabled"] = o.TsigEnabled
	}
	if !IsNil(o.TsigKey) {
		toSerialize["tsig_key"] = o.TsigKey
	}
	return toSerialize, nil
}

func (o *ConfigExternalSecondary) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"fqdn",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varConfigExternalSecondary := _ConfigExternalSecondary{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConfigExternalSecondary)

	if err != nil {
		return err
	}

	*o = ConfigExternalSecondary(varConfigExternalSecondary)

	return err
}

type NullableConfigExternalSecondary struct {
	value *ConfigExternalSecondary
	isSet bool
}

func (v NullableConfigExternalSecondary) Get() *ConfigExternalSecondary {
	return v.value
}

func (v *NullableConfigExternalSecondary) Set(val *ConfigExternalSecondary) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigExternalSecondary) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigExternalSecondary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigExternalSecondary(val *ConfigExternalSecondary) *NullableConfigExternalSecondary {
	return &NullableConfigExternalSecondary{value: val, isSet: true}
}

func (v NullableConfigExternalSecondary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigExternalSecondary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


