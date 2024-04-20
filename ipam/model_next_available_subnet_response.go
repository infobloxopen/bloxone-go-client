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

// checks if the NextAvailableSubnetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NextAvailableSubnetResponse{}

// NextAvailableSubnetResponse The Next Available Subnet object list response format.
type NextAvailableSubnetResponse struct {
	// The list of Next Available Subnet objects.
	Results []Subnet `json:"results,omitempty"`
}

// NewNextAvailableSubnetResponse instantiates a new NextAvailableSubnetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNextAvailableSubnetResponse() *NextAvailableSubnetResponse {
	this := NextAvailableSubnetResponse{}
	return &this
}

// NewNextAvailableSubnetResponseWithDefaults instantiates a new NextAvailableSubnetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNextAvailableSubnetResponseWithDefaults() *NextAvailableSubnetResponse {
	this := NextAvailableSubnetResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *NextAvailableSubnetResponse) GetResults() []Subnet {
	if o == nil || IsNil(o.Results) {
		var ret []Subnet
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NextAvailableSubnetResponse) GetResultsOk() ([]Subnet, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *NextAvailableSubnetResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Subnet and assigns it to the Results field.
func (o *NextAvailableSubnetResponse) SetResults(v []Subnet) {
	o.Results = v
}

func (o NextAvailableSubnetResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NextAvailableSubnetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableNextAvailableSubnetResponse struct {
	value *NextAvailableSubnetResponse
	isSet bool
}

func (v NullableNextAvailableSubnetResponse) Get() *NextAvailableSubnetResponse {
	return v.value
}

func (v *NullableNextAvailableSubnetResponse) Set(val *NextAvailableSubnetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNextAvailableSubnetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNextAvailableSubnetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNextAvailableSubnetResponse(val *NextAvailableSubnetResponse) *NullableNextAvailableSubnetResponse {
	return &NullableNextAvailableSubnetResponse{value: val, isSet: true}
}

func (v NullableNextAvailableSubnetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNextAvailableSubnetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
