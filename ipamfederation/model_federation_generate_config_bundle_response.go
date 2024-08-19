/*
IPAM Federation API

The DDI IPAM Federation application enables a SaaS administrator to manage multiple IPAM systems from one central control point CSP.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipamfederation

import (
	"encoding/json"
)

// checks if the FederationGenerateConfigBundleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FederationGenerateConfigBundleResponse{}

// FederationGenerateConfigBundleResponse The response format to retrieve __ConfigBundle__ objects.
type FederationGenerateConfigBundleResponse struct {
	Config               *FederationConfigObject `json:"config,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FederationGenerateConfigBundleResponse FederationGenerateConfigBundleResponse

// NewFederationGenerateConfigBundleResponse instantiates a new FederationGenerateConfigBundleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFederationGenerateConfigBundleResponse() *FederationGenerateConfigBundleResponse {
	this := FederationGenerateConfigBundleResponse{}
	return &this
}

// NewFederationGenerateConfigBundleResponseWithDefaults instantiates a new FederationGenerateConfigBundleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFederationGenerateConfigBundleResponseWithDefaults() *FederationGenerateConfigBundleResponse {
	this := FederationGenerateConfigBundleResponse{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *FederationGenerateConfigBundleResponse) GetConfig() FederationConfigObject {
	if o == nil || IsNil(o.Config) {
		var ret FederationConfigObject
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationGenerateConfigBundleResponse) GetConfigOk() (*FederationConfigObject, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *FederationGenerateConfigBundleResponse) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given FederationConfigObject and assigns it to the Config field.
func (o *FederationGenerateConfigBundleResponse) SetConfig(v FederationConfigObject) {
	o.Config = &v
}

func (o FederationGenerateConfigBundleResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FederationGenerateConfigBundleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FederationGenerateConfigBundleResponse) UnmarshalJSON(data []byte) (err error) {
	varFederationGenerateConfigBundleResponse := _FederationGenerateConfigBundleResponse{}

	err = json.Unmarshal(data, &varFederationGenerateConfigBundleResponse)

	if err != nil {
		return err
	}

	*o = FederationGenerateConfigBundleResponse(varFederationGenerateConfigBundleResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "config")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFederationGenerateConfigBundleResponse struct {
	value *FederationGenerateConfigBundleResponse
	isSet bool
}

func (v NullableFederationGenerateConfigBundleResponse) Get() *FederationGenerateConfigBundleResponse {
	return v.value
}

func (v *NullableFederationGenerateConfigBundleResponse) Set(val *FederationGenerateConfigBundleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFederationGenerateConfigBundleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFederationGenerateConfigBundleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFederationGenerateConfigBundleResponse(val *FederationGenerateConfigBundleResponse) *NullableFederationGenerateConfigBundleResponse {
	return &NullableFederationGenerateConfigBundleResponse{value: val, isSet: true}
}

func (v NullableFederationGenerateConfigBundleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFederationGenerateConfigBundleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}