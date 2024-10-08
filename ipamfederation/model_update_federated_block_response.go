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

// checks if the UpdateFederatedBlockResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFederatedBlockResponse{}

// UpdateFederatedBlockResponse The response format to update the __FederatedBlock__ object.
type UpdateFederatedBlockResponse struct {
	// The FederatedBlock object.
	Result               *FederatedBlock `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateFederatedBlockResponse UpdateFederatedBlockResponse

// NewUpdateFederatedBlockResponse instantiates a new UpdateFederatedBlockResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFederatedBlockResponse() *UpdateFederatedBlockResponse {
	this := UpdateFederatedBlockResponse{}
	return &this
}

// NewUpdateFederatedBlockResponseWithDefaults instantiates a new UpdateFederatedBlockResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFederatedBlockResponseWithDefaults() *UpdateFederatedBlockResponse {
	this := UpdateFederatedBlockResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateFederatedBlockResponse) GetResult() FederatedBlock {
	if o == nil || IsNil(o.Result) {
		var ret FederatedBlock
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFederatedBlockResponse) GetResultOk() (*FederatedBlock, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateFederatedBlockResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given FederatedBlock and assigns it to the Result field.
func (o *UpdateFederatedBlockResponse) SetResult(v FederatedBlock) {
	o.Result = &v
}

func (o UpdateFederatedBlockResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFederatedBlockResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateFederatedBlockResponse) UnmarshalJSON(data []byte) (err error) {
	varUpdateFederatedBlockResponse := _UpdateFederatedBlockResponse{}

	err = json.Unmarshal(data, &varUpdateFederatedBlockResponse)

	if err != nil {
		return err
	}

	*o = UpdateFederatedBlockResponse(varUpdateFederatedBlockResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateFederatedBlockResponse struct {
	value *UpdateFederatedBlockResponse
	isSet bool
}

func (v NullableUpdateFederatedBlockResponse) Get() *UpdateFederatedBlockResponse {
	return v.value
}

func (v *NullableUpdateFederatedBlockResponse) Set(val *UpdateFederatedBlockResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFederatedBlockResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFederatedBlockResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFederatedBlockResponse(val *UpdateFederatedBlockResponse) *NullableUpdateFederatedBlockResponse {
	return &NullableUpdateFederatedBlockResponse{value: val, isSet: true}
}

func (v NullableUpdateFederatedBlockResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFederatedBlockResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
