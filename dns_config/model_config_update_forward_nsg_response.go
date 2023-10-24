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

// checks if the ConfigUpdateForwardNSGResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigUpdateForwardNSGResponse{}

// ConfigUpdateForwardNSGResponse The ForwardNSG object update response format.
type ConfigUpdateForwardNSGResponse struct {
	Result *ConfigForwardNSG `json:"result,omitempty"`
}

// NewConfigUpdateForwardNSGResponse instantiates a new ConfigUpdateForwardNSGResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigUpdateForwardNSGResponse() *ConfigUpdateForwardNSGResponse {
	this := ConfigUpdateForwardNSGResponse{}
	return &this
}

// NewConfigUpdateForwardNSGResponseWithDefaults instantiates a new ConfigUpdateForwardNSGResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigUpdateForwardNSGResponseWithDefaults() *ConfigUpdateForwardNSGResponse {
	this := ConfigUpdateForwardNSGResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ConfigUpdateForwardNSGResponse) GetResult() ConfigForwardNSG {
	if o == nil || IsNil(o.Result) {
		var ret ConfigForwardNSG
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigUpdateForwardNSGResponse) GetResultOk() (*ConfigForwardNSG, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ConfigUpdateForwardNSGResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ConfigForwardNSG and assigns it to the Result field.
func (o *ConfigUpdateForwardNSGResponse) SetResult(v ConfigForwardNSG) {
	o.Result = &v
}

func (o ConfigUpdateForwardNSGResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigUpdateForwardNSGResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableConfigUpdateForwardNSGResponse struct {
	value *ConfigUpdateForwardNSGResponse
	isSet bool
}

func (v NullableConfigUpdateForwardNSGResponse) Get() *ConfigUpdateForwardNSGResponse {
	return v.value
}

func (v *NullableConfigUpdateForwardNSGResponse) Set(val *ConfigUpdateForwardNSGResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigUpdateForwardNSGResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigUpdateForwardNSGResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigUpdateForwardNSGResponse(val *ConfigUpdateForwardNSGResponse) *NullableConfigUpdateForwardNSGResponse {
	return &NullableConfigUpdateForwardNSGResponse{value: val, isSet: true}
}

func (v NullableConfigUpdateForwardNSGResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigUpdateForwardNSGResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
