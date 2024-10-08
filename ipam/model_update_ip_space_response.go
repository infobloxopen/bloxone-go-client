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

// checks if the UpdateIPSpaceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateIPSpaceResponse{}

// UpdateIPSpaceResponse The response format to update the __IPSpace__ object.
type UpdateIPSpaceResponse struct {
	// The IPSpace object.
	Result               *IPSpace `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateIPSpaceResponse UpdateIPSpaceResponse

// NewUpdateIPSpaceResponse instantiates a new UpdateIPSpaceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateIPSpaceResponse() *UpdateIPSpaceResponse {
	this := UpdateIPSpaceResponse{}
	return &this
}

// NewUpdateIPSpaceResponseWithDefaults instantiates a new UpdateIPSpaceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateIPSpaceResponseWithDefaults() *UpdateIPSpaceResponse {
	this := UpdateIPSpaceResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateIPSpaceResponse) GetResult() IPSpace {
	if o == nil || IsNil(o.Result) {
		var ret IPSpace
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIPSpaceResponse) GetResultOk() (*IPSpace, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateIPSpaceResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given IPSpace and assigns it to the Result field.
func (o *UpdateIPSpaceResponse) SetResult(v IPSpace) {
	o.Result = &v
}

func (o UpdateIPSpaceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateIPSpaceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateIPSpaceResponse) UnmarshalJSON(data []byte) (err error) {
	varUpdateIPSpaceResponse := _UpdateIPSpaceResponse{}

	err = json.Unmarshal(data, &varUpdateIPSpaceResponse)

	if err != nil {
		return err
	}

	*o = UpdateIPSpaceResponse(varUpdateIPSpaceResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateIPSpaceResponse struct {
	value *UpdateIPSpaceResponse
	isSet bool
}

func (v NullableUpdateIPSpaceResponse) Get() *UpdateIPSpaceResponse {
	return v.value
}

func (v *NullableUpdateIPSpaceResponse) Set(val *UpdateIPSpaceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateIPSpaceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateIPSpaceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateIPSpaceResponse(val *UpdateIPSpaceResponse) *NullableUpdateIPSpaceResponse {
	return &NullableUpdateIPSpaceResponse{value: val, isSet: true}
}

func (v NullableUpdateIPSpaceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateIPSpaceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
