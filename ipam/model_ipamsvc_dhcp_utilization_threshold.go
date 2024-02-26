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

// checks if the IpamsvcDHCPUtilizationThreshold type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcDHCPUtilizationThreshold{}

// IpamsvcDHCPUtilizationThreshold A __DHCPUtilizationThreshold__ object represents threshold settings for DHCP utilization.
type IpamsvcDHCPUtilizationThreshold struct {
	// Indicates whether the DHCP utilization threshold is enabled or not.
	Enabled bool `json:"enabled"`
	// The high threshold value for DHCP utilization in percentage.
	High int64 `json:"high"`
	// The low threshold value for DHCP utilization in percentage.
	Low int64 `json:"low"`
}

type _IpamsvcDHCPUtilizationThreshold IpamsvcDHCPUtilizationThreshold

// NewIpamsvcDHCPUtilizationThreshold instantiates a new IpamsvcDHCPUtilizationThreshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcDHCPUtilizationThreshold(enabled bool, high int64, low int64) *IpamsvcDHCPUtilizationThreshold {
	this := IpamsvcDHCPUtilizationThreshold{}
	this.Enabled = enabled
	this.High = high
	this.Low = low
	return &this
}

// NewIpamsvcDHCPUtilizationThresholdWithDefaults instantiates a new IpamsvcDHCPUtilizationThreshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcDHCPUtilizationThresholdWithDefaults() *IpamsvcDHCPUtilizationThreshold {
	this := IpamsvcDHCPUtilizationThreshold{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *IpamsvcDHCPUtilizationThreshold) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *IpamsvcDHCPUtilizationThreshold) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *IpamsvcDHCPUtilizationThreshold) SetEnabled(v bool) {
	o.Enabled = v
}

// GetHigh returns the High field value
func (o *IpamsvcDHCPUtilizationThreshold) GetHigh() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.High
}

// GetHighOk returns a tuple with the High field value
// and a boolean to check if the value has been set.
func (o *IpamsvcDHCPUtilizationThreshold) GetHighOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.High, true
}

// SetHigh sets field value
func (o *IpamsvcDHCPUtilizationThreshold) SetHigh(v int64) {
	o.High = v
}

// GetLow returns the Low field value
func (o *IpamsvcDHCPUtilizationThreshold) GetLow() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Low
}

// GetLowOk returns a tuple with the Low field value
// and a boolean to check if the value has been set.
func (o *IpamsvcDHCPUtilizationThreshold) GetLowOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Low, true
}

// SetLow sets field value
func (o *IpamsvcDHCPUtilizationThreshold) SetLow(v int64) {
	o.Low = v
}

func (o IpamsvcDHCPUtilizationThreshold) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcDHCPUtilizationThreshold) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	toSerialize["high"] = o.High
	toSerialize["low"] = o.Low
	return toSerialize, nil
}

func (o *IpamsvcDHCPUtilizationThreshold) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"high",
		"low",
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

	varIpamsvcDHCPUtilizationThreshold := _IpamsvcDHCPUtilizationThreshold{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIpamsvcDHCPUtilizationThreshold)

	if err != nil {
		return err
	}

	*o = IpamsvcDHCPUtilizationThreshold(varIpamsvcDHCPUtilizationThreshold)

	return err
}

type NullableIpamsvcDHCPUtilizationThreshold struct {
	value *IpamsvcDHCPUtilizationThreshold
	isSet bool
}

func (v NullableIpamsvcDHCPUtilizationThreshold) Get() *IpamsvcDHCPUtilizationThreshold {
	return v.value
}

func (v *NullableIpamsvcDHCPUtilizationThreshold) Set(val *IpamsvcDHCPUtilizationThreshold) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcDHCPUtilizationThreshold) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcDHCPUtilizationThreshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcDHCPUtilizationThreshold(val *IpamsvcDHCPUtilizationThreshold) *NullableIpamsvcDHCPUtilizationThreshold {
	return &NullableIpamsvcDHCPUtilizationThreshold{value: val, isSet: true}
}

func (v NullableIpamsvcDHCPUtilizationThreshold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcDHCPUtilizationThreshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
