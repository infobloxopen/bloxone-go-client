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

// checks if the FederationGetNextAvailableDelegationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FederationGetNextAvailableDelegationRequest{}

// FederationGetNextAvailableDelegationRequest struct for FederationGetNextAvailableDelegationRequest
type FederationGetNextAvailableDelegationRequest struct {
	// The arguments which will be used as parameters while creating __Delegation__ object.
	Arguments *FederationNextAvailableDelegationArguments `json:"arguments,omitempty"`
	// The properties, which will be used to fill the created __Delegation__ object.
	Properties           *FederationNextAvailableDelegationProperties `json:"properties,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FederationGetNextAvailableDelegationRequest FederationGetNextAvailableDelegationRequest

// NewFederationGetNextAvailableDelegationRequest instantiates a new FederationGetNextAvailableDelegationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFederationGetNextAvailableDelegationRequest() *FederationGetNextAvailableDelegationRequest {
	this := FederationGetNextAvailableDelegationRequest{}
	return &this
}

// NewFederationGetNextAvailableDelegationRequestWithDefaults instantiates a new FederationGetNextAvailableDelegationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFederationGetNextAvailableDelegationRequestWithDefaults() *FederationGetNextAvailableDelegationRequest {
	this := FederationGetNextAvailableDelegationRequest{}
	return &this
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *FederationGetNextAvailableDelegationRequest) GetArguments() FederationNextAvailableDelegationArguments {
	if o == nil || IsNil(o.Arguments) {
		var ret FederationNextAvailableDelegationArguments
		return ret
	}
	return *o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationGetNextAvailableDelegationRequest) GetArgumentsOk() (*FederationNextAvailableDelegationArguments, bool) {
	if o == nil || IsNil(o.Arguments) {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *FederationGetNextAvailableDelegationRequest) HasArguments() bool {
	if o != nil && !IsNil(o.Arguments) {
		return true
	}

	return false
}

// SetArguments gets a reference to the given FederationNextAvailableDelegationArguments and assigns it to the Arguments field.
func (o *FederationGetNextAvailableDelegationRequest) SetArguments(v FederationNextAvailableDelegationArguments) {
	o.Arguments = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *FederationGetNextAvailableDelegationRequest) GetProperties() FederationNextAvailableDelegationProperties {
	if o == nil || IsNil(o.Properties) {
		var ret FederationNextAvailableDelegationProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationGetNextAvailableDelegationRequest) GetPropertiesOk() (*FederationNextAvailableDelegationProperties, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *FederationGetNextAvailableDelegationRequest) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given FederationNextAvailableDelegationProperties and assigns it to the Properties field.
func (o *FederationGetNextAvailableDelegationRequest) SetProperties(v FederationNextAvailableDelegationProperties) {
	o.Properties = &v
}

func (o FederationGetNextAvailableDelegationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FederationGetNextAvailableDelegationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Arguments) {
		toSerialize["arguments"] = o.Arguments
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FederationGetNextAvailableDelegationRequest) UnmarshalJSON(data []byte) (err error) {
	varFederationGetNextAvailableDelegationRequest := _FederationGetNextAvailableDelegationRequest{}

	err = json.Unmarshal(data, &varFederationGetNextAvailableDelegationRequest)

	if err != nil {
		return err
	}

	*o = FederationGetNextAvailableDelegationRequest(varFederationGetNextAvailableDelegationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "arguments")
		delete(additionalProperties, "properties")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFederationGetNextAvailableDelegationRequest struct {
	value *FederationGetNextAvailableDelegationRequest
	isSet bool
}

func (v NullableFederationGetNextAvailableDelegationRequest) Get() *FederationGetNextAvailableDelegationRequest {
	return v.value
}

func (v *NullableFederationGetNextAvailableDelegationRequest) Set(val *FederationGetNextAvailableDelegationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFederationGetNextAvailableDelegationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFederationGetNextAvailableDelegationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFederationGetNextAvailableDelegationRequest(val *FederationGetNextAvailableDelegationRequest) *NullableFederationGetNextAvailableDelegationRequest {
	return &NullableFederationGetNextAvailableDelegationRequest{value: val, isSet: true}
}

func (v NullableFederationGetNextAvailableDelegationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFederationGetNextAvailableDelegationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
