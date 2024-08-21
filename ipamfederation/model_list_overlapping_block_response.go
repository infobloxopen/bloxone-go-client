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

// checks if the ListOverlappingBlockResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOverlappingBlockResponse{}

// ListOverlappingBlockResponse The response format to retrieve __OverlappingBlock__ objects.
type ListOverlappingBlockResponse struct {
	// A list of OverlappingBlock objects.
	Results              []OverlappingBlock `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListOverlappingBlockResponse ListOverlappingBlockResponse

// NewListOverlappingBlockResponse instantiates a new ListOverlappingBlockResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOverlappingBlockResponse() *ListOverlappingBlockResponse {
	this := ListOverlappingBlockResponse{}
	return &this
}

// NewListOverlappingBlockResponseWithDefaults instantiates a new ListOverlappingBlockResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOverlappingBlockResponseWithDefaults() *ListOverlappingBlockResponse {
	this := ListOverlappingBlockResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListOverlappingBlockResponse) GetResults() []OverlappingBlock {
	if o == nil || IsNil(o.Results) {
		var ret []OverlappingBlock
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOverlappingBlockResponse) GetResultsOk() ([]OverlappingBlock, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListOverlappingBlockResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []OverlappingBlock and assigns it to the Results field.
func (o *ListOverlappingBlockResponse) SetResults(v []OverlappingBlock) {
	o.Results = v
}

func (o ListOverlappingBlockResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOverlappingBlockResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListOverlappingBlockResponse) UnmarshalJSON(data []byte) (err error) {
	varListOverlappingBlockResponse := _ListOverlappingBlockResponse{}

	err = json.Unmarshal(data, &varListOverlappingBlockResponse)

	if err != nil {
		return err
	}

	*o = ListOverlappingBlockResponse(varListOverlappingBlockResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListOverlappingBlockResponse struct {
	value *ListOverlappingBlockResponse
	isSet bool
}

func (v NullableListOverlappingBlockResponse) Get() *ListOverlappingBlockResponse {
	return v.value
}

func (v *NullableListOverlappingBlockResponse) Set(val *ListOverlappingBlockResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListOverlappingBlockResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListOverlappingBlockResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOverlappingBlockResponse(val *ListOverlappingBlockResponse) *NullableListOverlappingBlockResponse {
	return &NullableListOverlappingBlockResponse{value: val, isSet: true}
}

func (v NullableListOverlappingBlockResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOverlappingBlockResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}