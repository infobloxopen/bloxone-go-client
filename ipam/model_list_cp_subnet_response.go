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

// checks if the ListCPSubnetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCPSubnetResponse{}

// ListCPSubnetResponse The response format to retrieve subnets associated with config profile.
type ListCPSubnetResponse struct {
	Results              []CPSubnet `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListCPSubnetResponse ListCPSubnetResponse

// NewListCPSubnetResponse instantiates a new ListCPSubnetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCPSubnetResponse() *ListCPSubnetResponse {
	this := ListCPSubnetResponse{}
	return &this
}

// NewListCPSubnetResponseWithDefaults instantiates a new ListCPSubnetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCPSubnetResponseWithDefaults() *ListCPSubnetResponse {
	this := ListCPSubnetResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListCPSubnetResponse) GetResults() []CPSubnet {
	if o == nil || IsNil(o.Results) {
		var ret []CPSubnet
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCPSubnetResponse) GetResultsOk() ([]CPSubnet, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListCPSubnetResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []CPSubnet and assigns it to the Results field.
func (o *ListCPSubnetResponse) SetResults(v []CPSubnet) {
	o.Results = v
}

func (o ListCPSubnetResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCPSubnetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListCPSubnetResponse) UnmarshalJSON(data []byte) (err error) {
	varListCPSubnetResponse := _ListCPSubnetResponse{}

	err = json.Unmarshal(data, &varListCPSubnetResponse)

	if err != nil {
		return err
	}

	*o = ListCPSubnetResponse(varListCPSubnetResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListCPSubnetResponse struct {
	value *ListCPSubnetResponse
	isSet bool
}

func (v NullableListCPSubnetResponse) Get() *ListCPSubnetResponse {
	return v.value
}

func (v *NullableListCPSubnetResponse) Set(val *ListCPSubnetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListCPSubnetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListCPSubnetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCPSubnetResponse(val *ListCPSubnetResponse) *NullableListCPSubnetResponse {
	return &NullableListCPSubnetResponse{value: val, isSet: true}
}

func (v NullableListCPSubnetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCPSubnetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
