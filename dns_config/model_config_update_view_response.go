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

// checks if the ConfigUpdateViewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigUpdateViewResponse{}

// ConfigUpdateViewResponse The View object update response format.
type ConfigUpdateViewResponse struct {
	Result *ConfigView `json:"result,omitempty"`
}

// NewConfigUpdateViewResponse instantiates a new ConfigUpdateViewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigUpdateViewResponse() *ConfigUpdateViewResponse {
	this := ConfigUpdateViewResponse{}
	return &this
}

// NewConfigUpdateViewResponseWithDefaults instantiates a new ConfigUpdateViewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigUpdateViewResponseWithDefaults() *ConfigUpdateViewResponse {
	this := ConfigUpdateViewResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ConfigUpdateViewResponse) GetResult() ConfigView {
	if o == nil || IsNil(o.Result) {
		var ret ConfigView
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigUpdateViewResponse) GetResultOk() (*ConfigView, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ConfigUpdateViewResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ConfigView and assigns it to the Result field.
func (o *ConfigUpdateViewResponse) SetResult(v ConfigView) {
	o.Result = &v
}

func (o ConfigUpdateViewResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigUpdateViewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableConfigUpdateViewResponse struct {
	value *ConfigUpdateViewResponse
	isSet bool
}

func (v NullableConfigUpdateViewResponse) Get() *ConfigUpdateViewResponse {
	return v.value
}

func (v *NullableConfigUpdateViewResponse) Set(val *ConfigUpdateViewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigUpdateViewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigUpdateViewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigUpdateViewResponse(val *ConfigUpdateViewResponse) *NullableConfigUpdateViewResponse {
	return &NullableConfigUpdateViewResponse{value: val, isSet: true}
}

func (v NullableConfigUpdateViewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigUpdateViewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


