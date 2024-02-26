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

// checks if the ConfigListACLResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigListACLResponse{}

// ConfigListACLResponse The ACL object list response format.
type ConfigListACLResponse struct {
	// List of ACL objects.
	Results []ConfigACL `json:"results,omitempty"`
}

// NewConfigListACLResponse instantiates a new ConfigListACLResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigListACLResponse() *ConfigListACLResponse {
	this := ConfigListACLResponse{}
	return &this
}

// NewConfigListACLResponseWithDefaults instantiates a new ConfigListACLResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigListACLResponseWithDefaults() *ConfigListACLResponse {
	this := ConfigListACLResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ConfigListACLResponse) GetResults() []ConfigACL {
	if o == nil || IsNil(o.Results) {
		var ret []ConfigACL
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigListACLResponse) GetResultsOk() ([]ConfigACL, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ConfigListACLResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ConfigACL and assigns it to the Results field.
func (o *ConfigListACLResponse) SetResults(v []ConfigACL) {
	o.Results = v
}

func (o ConfigListACLResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigListACLResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableConfigListACLResponse struct {
	value *ConfigListACLResponse
	isSet bool
}

func (v NullableConfigListACLResponse) Get() *ConfigListACLResponse {
	return v.value
}

func (v *NullableConfigListACLResponse) Set(val *ConfigListACLResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigListACLResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigListACLResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigListACLResponse(val *ConfigListACLResponse) *NullableConfigListACLResponse {
	return &NullableConfigListACLResponse{value: val, isSet: true}
}

func (v NullableConfigListACLResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigListACLResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


