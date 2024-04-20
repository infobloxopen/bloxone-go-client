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

// checks if the ListFixedAddressResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFixedAddressResponse{}

// ListFixedAddressResponse The response format to retrieve __FixedAddress__ objects.
type ListFixedAddressResponse struct {
	// The list of FixedAddress objects.
	Results []FixedAddress `json:"results,omitempty"`
}

// NewListFixedAddressResponse instantiates a new ListFixedAddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFixedAddressResponse() *ListFixedAddressResponse {
	this := ListFixedAddressResponse{}
	return &this
}

// NewListFixedAddressResponseWithDefaults instantiates a new ListFixedAddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFixedAddressResponseWithDefaults() *ListFixedAddressResponse {
	this := ListFixedAddressResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListFixedAddressResponse) GetResults() []FixedAddress {
	if o == nil || IsNil(o.Results) {
		var ret []FixedAddress
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFixedAddressResponse) GetResultsOk() ([]FixedAddress, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListFixedAddressResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []FixedAddress and assigns it to the Results field.
func (o *ListFixedAddressResponse) SetResults(v []FixedAddress) {
	o.Results = v
}

func (o ListFixedAddressResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListFixedAddressResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableListFixedAddressResponse struct {
	value *ListFixedAddressResponse
	isSet bool
}

func (v NullableListFixedAddressResponse) Get() *ListFixedAddressResponse {
	return v.value
}

func (v *NullableListFixedAddressResponse) Set(val *ListFixedAddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListFixedAddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListFixedAddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFixedAddressResponse(val *ListFixedAddressResponse) *NullableListFixedAddressResponse {
	return &NullableListFixedAddressResponse{value: val, isSet: true}
}

func (v NullableListFixedAddressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFixedAddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
