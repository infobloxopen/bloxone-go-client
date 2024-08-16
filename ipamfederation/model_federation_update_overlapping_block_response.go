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

// checks if the FederationUpdateOverlappingBlockResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FederationUpdateOverlappingBlockResponse{}

// FederationUpdateOverlappingBlockResponse The response format to update the __OverlappingBlock__ object.
type FederationUpdateOverlappingBlockResponse struct {
	// The OverlappingBlock object.
	Result               *FederationOverlappingBlock `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FederationUpdateOverlappingBlockResponse FederationUpdateOverlappingBlockResponse

// NewFederationUpdateOverlappingBlockResponse instantiates a new FederationUpdateOverlappingBlockResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFederationUpdateOverlappingBlockResponse() *FederationUpdateOverlappingBlockResponse {
	this := FederationUpdateOverlappingBlockResponse{}
	return &this
}

// NewFederationUpdateOverlappingBlockResponseWithDefaults instantiates a new FederationUpdateOverlappingBlockResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFederationUpdateOverlappingBlockResponseWithDefaults() *FederationUpdateOverlappingBlockResponse {
	this := FederationUpdateOverlappingBlockResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *FederationUpdateOverlappingBlockResponse) GetResult() FederationOverlappingBlock {
	if o == nil || IsNil(o.Result) {
		var ret FederationOverlappingBlock
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationUpdateOverlappingBlockResponse) GetResultOk() (*FederationOverlappingBlock, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *FederationUpdateOverlappingBlockResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given FederationOverlappingBlock and assigns it to the Result field.
func (o *FederationUpdateOverlappingBlockResponse) SetResult(v FederationOverlappingBlock) {
	o.Result = &v
}

func (o FederationUpdateOverlappingBlockResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FederationUpdateOverlappingBlockResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FederationUpdateOverlappingBlockResponse) UnmarshalJSON(data []byte) (err error) {
	varFederationUpdateOverlappingBlockResponse := _FederationUpdateOverlappingBlockResponse{}

	err = json.Unmarshal(data, &varFederationUpdateOverlappingBlockResponse)

	if err != nil {
		return err
	}

	*o = FederationUpdateOverlappingBlockResponse(varFederationUpdateOverlappingBlockResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFederationUpdateOverlappingBlockResponse struct {
	value *FederationUpdateOverlappingBlockResponse
	isSet bool
}

func (v NullableFederationUpdateOverlappingBlockResponse) Get() *FederationUpdateOverlappingBlockResponse {
	return v.value
}

func (v *NullableFederationUpdateOverlappingBlockResponse) Set(val *FederationUpdateOverlappingBlockResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFederationUpdateOverlappingBlockResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFederationUpdateOverlappingBlockResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFederationUpdateOverlappingBlockResponse(val *FederationUpdateOverlappingBlockResponse) *NullableFederationUpdateOverlappingBlockResponse {
	return &NullableFederationUpdateOverlappingBlockResponse{value: val, isSet: true}
}

func (v NullableFederationUpdateOverlappingBlockResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFederationUpdateOverlappingBlockResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
