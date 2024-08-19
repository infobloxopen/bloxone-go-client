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

// checks if the FederationHealthCheckConfigResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FederationHealthCheckConfigResponse{}

// FederationHealthCheckConfigResponse struct for FederationHealthCheckConfigResponse
type FederationHealthCheckConfigResponse struct {
	Message              *string `json:"message,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FederationHealthCheckConfigResponse FederationHealthCheckConfigResponse

// NewFederationHealthCheckConfigResponse instantiates a new FederationHealthCheckConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFederationHealthCheckConfigResponse() *FederationHealthCheckConfigResponse {
	this := FederationHealthCheckConfigResponse{}
	return &this
}

// NewFederationHealthCheckConfigResponseWithDefaults instantiates a new FederationHealthCheckConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFederationHealthCheckConfigResponseWithDefaults() *FederationHealthCheckConfigResponse {
	this := FederationHealthCheckConfigResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *FederationHealthCheckConfigResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationHealthCheckConfigResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *FederationHealthCheckConfigResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *FederationHealthCheckConfigResponse) SetMessage(v string) {
	o.Message = &v
}

func (o FederationHealthCheckConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FederationHealthCheckConfigResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FederationHealthCheckConfigResponse) UnmarshalJSON(data []byte) (err error) {
	varFederationHealthCheckConfigResponse := _FederationHealthCheckConfigResponse{}

	err = json.Unmarshal(data, &varFederationHealthCheckConfigResponse)

	if err != nil {
		return err
	}

	*o = FederationHealthCheckConfigResponse(varFederationHealthCheckConfigResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFederationHealthCheckConfigResponse struct {
	value *FederationHealthCheckConfigResponse
	isSet bool
}

func (v NullableFederationHealthCheckConfigResponse) Get() *FederationHealthCheckConfigResponse {
	return v.value
}

func (v *NullableFederationHealthCheckConfigResponse) Set(val *FederationHealthCheckConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFederationHealthCheckConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFederationHealthCheckConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFederationHealthCheckConfigResponse(val *FederationHealthCheckConfigResponse) *NullableFederationHealthCheckConfigResponse {
	return &NullableFederationHealthCheckConfigResponse{value: val, isSet: true}
}

func (v NullableFederationHealthCheckConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFederationHealthCheckConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}