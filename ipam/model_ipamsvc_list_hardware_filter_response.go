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

// checks if the IpamsvcListHardwareFilterResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcListHardwareFilterResponse{}

// IpamsvcListHardwareFilterResponse The response format to retrieve __HardwareFilter__ objects.
type IpamsvcListHardwareFilterResponse struct {
	// The list of HardwareFilter objects.
	Results []IpamsvcHardwareFilter `json:"results,omitempty"`
}

// NewIpamsvcListHardwareFilterResponse instantiates a new IpamsvcListHardwareFilterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcListHardwareFilterResponse() *IpamsvcListHardwareFilterResponse {
	this := IpamsvcListHardwareFilterResponse{}
	return &this
}

// NewIpamsvcListHardwareFilterResponseWithDefaults instantiates a new IpamsvcListHardwareFilterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcListHardwareFilterResponseWithDefaults() *IpamsvcListHardwareFilterResponse {
	this := IpamsvcListHardwareFilterResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *IpamsvcListHardwareFilterResponse) GetResults() []IpamsvcHardwareFilter {
	if o == nil || IsNil(o.Results) {
		var ret []IpamsvcHardwareFilter
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcListHardwareFilterResponse) GetResultsOk() ([]IpamsvcHardwareFilter, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *IpamsvcListHardwareFilterResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []IpamsvcHardwareFilter and assigns it to the Results field.
func (o *IpamsvcListHardwareFilterResponse) SetResults(v []IpamsvcHardwareFilter) {
	o.Results = v
}

func (o IpamsvcListHardwareFilterResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcListHardwareFilterResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableIpamsvcListHardwareFilterResponse struct {
	value *IpamsvcListHardwareFilterResponse
	isSet bool
}

func (v NullableIpamsvcListHardwareFilterResponse) Get() *IpamsvcListHardwareFilterResponse {
	return v.value
}

func (v *NullableIpamsvcListHardwareFilterResponse) Set(val *IpamsvcListHardwareFilterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcListHardwareFilterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcListHardwareFilterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcListHardwareFilterResponse(val *IpamsvcListHardwareFilterResponse) *NullableIpamsvcListHardwareFilterResponse {
	return &NullableIpamsvcListHardwareFilterResponse{value: val, isSet: true}
}

func (v NullableIpamsvcListHardwareFilterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcListHardwareFilterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
