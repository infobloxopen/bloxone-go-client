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

// checks if the FederationCreateOverlappingBlockResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FederationCreateOverlappingBlockResponse{}

// FederationCreateOverlappingBlockResponse The response format to create the __OverlappingBlock__ object.
type FederationCreateOverlappingBlockResponse struct {
	// The created OverlappingBlock object.
	Result               *FederationOverlappingBlock `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FederationCreateOverlappingBlockResponse FederationCreateOverlappingBlockResponse

// NewFederationCreateOverlappingBlockResponse instantiates a new FederationCreateOverlappingBlockResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFederationCreateOverlappingBlockResponse() *FederationCreateOverlappingBlockResponse {
	this := FederationCreateOverlappingBlockResponse{}
	return &this
}

// NewFederationCreateOverlappingBlockResponseWithDefaults instantiates a new FederationCreateOverlappingBlockResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFederationCreateOverlappingBlockResponseWithDefaults() *FederationCreateOverlappingBlockResponse {
	this := FederationCreateOverlappingBlockResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *FederationCreateOverlappingBlockResponse) GetResult() FederationOverlappingBlock {
	if o == nil || IsNil(o.Result) {
		var ret FederationOverlappingBlock
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationCreateOverlappingBlockResponse) GetResultOk() (*FederationOverlappingBlock, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *FederationCreateOverlappingBlockResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given FederationOverlappingBlock and assigns it to the Result field.
func (o *FederationCreateOverlappingBlockResponse) SetResult(v FederationOverlappingBlock) {
	o.Result = &v
}

func (o FederationCreateOverlappingBlockResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FederationCreateOverlappingBlockResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FederationCreateOverlappingBlockResponse) UnmarshalJSON(data []byte) (err error) {
	varFederationCreateOverlappingBlockResponse := _FederationCreateOverlappingBlockResponse{}

	err = json.Unmarshal(data, &varFederationCreateOverlappingBlockResponse)

	if err != nil {
		return err
	}

	*o = FederationCreateOverlappingBlockResponse(varFederationCreateOverlappingBlockResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFederationCreateOverlappingBlockResponse struct {
	value *FederationCreateOverlappingBlockResponse
	isSet bool
}

func (v NullableFederationCreateOverlappingBlockResponse) Get() *FederationCreateOverlappingBlockResponse {
	return v.value
}

func (v *NullableFederationCreateOverlappingBlockResponse) Set(val *FederationCreateOverlappingBlockResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFederationCreateOverlappingBlockResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFederationCreateOverlappingBlockResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFederationCreateOverlappingBlockResponse(val *FederationCreateOverlappingBlockResponse) *NullableFederationCreateOverlappingBlockResponse {
	return &NullableFederationCreateOverlappingBlockResponse{value: val, isSet: true}
}

func (v NullableFederationCreateOverlappingBlockResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFederationCreateOverlappingBlockResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
