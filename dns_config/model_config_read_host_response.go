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

// checks if the ConfigReadHostResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigReadHostResponse{}

// ConfigReadHostResponse The DNS Host object read response format.
type ConfigReadHostResponse struct {
	Result *ConfigHost `json:"result,omitempty"`
}

// NewConfigReadHostResponse instantiates a new ConfigReadHostResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigReadHostResponse() *ConfigReadHostResponse {
	this := ConfigReadHostResponse{}
	return &this
}

// NewConfigReadHostResponseWithDefaults instantiates a new ConfigReadHostResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigReadHostResponseWithDefaults() *ConfigReadHostResponse {
	this := ConfigReadHostResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ConfigReadHostResponse) GetResult() ConfigHost {
	if o == nil || IsNil(o.Result) {
		var ret ConfigHost
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigReadHostResponse) GetResultOk() (*ConfigHost, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ConfigReadHostResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ConfigHost and assigns it to the Result field.
func (o *ConfigReadHostResponse) SetResult(v ConfigHost) {
	o.Result = &v
}

func (o ConfigReadHostResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigReadHostResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableConfigReadHostResponse struct {
	value *ConfigReadHostResponse
	isSet bool
}

func (v NullableConfigReadHostResponse) Get() *ConfigReadHostResponse {
	return v.value
}

func (v *NullableConfigReadHostResponse) Set(val *ConfigReadHostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigReadHostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigReadHostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigReadHostResponse(val *ConfigReadHostResponse) *NullableConfigReadHostResponse {
	return &NullableConfigReadHostResponse{value: val, isSet: true}
}

func (v NullableConfigReadHostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigReadHostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
