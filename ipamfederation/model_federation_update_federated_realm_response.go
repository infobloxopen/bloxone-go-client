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

// checks if the FederationUpdateFederatedRealmResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FederationUpdateFederatedRealmResponse{}

// FederationUpdateFederatedRealmResponse The response format to update the __FederatedRealm__ object.
type FederationUpdateFederatedRealmResponse struct {
	// The FederatedRealm object.
	Result               *FederationFederatedRealm `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FederationUpdateFederatedRealmResponse FederationUpdateFederatedRealmResponse

// NewFederationUpdateFederatedRealmResponse instantiates a new FederationUpdateFederatedRealmResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFederationUpdateFederatedRealmResponse() *FederationUpdateFederatedRealmResponse {
	this := FederationUpdateFederatedRealmResponse{}
	return &this
}

// NewFederationUpdateFederatedRealmResponseWithDefaults instantiates a new FederationUpdateFederatedRealmResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFederationUpdateFederatedRealmResponseWithDefaults() *FederationUpdateFederatedRealmResponse {
	this := FederationUpdateFederatedRealmResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *FederationUpdateFederatedRealmResponse) GetResult() FederationFederatedRealm {
	if o == nil || IsNil(o.Result) {
		var ret FederationFederatedRealm
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationUpdateFederatedRealmResponse) GetResultOk() (*FederationFederatedRealm, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *FederationUpdateFederatedRealmResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given FederationFederatedRealm and assigns it to the Result field.
func (o *FederationUpdateFederatedRealmResponse) SetResult(v FederationFederatedRealm) {
	o.Result = &v
}

func (o FederationUpdateFederatedRealmResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FederationUpdateFederatedRealmResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FederationUpdateFederatedRealmResponse) UnmarshalJSON(data []byte) (err error) {
	varFederationUpdateFederatedRealmResponse := _FederationUpdateFederatedRealmResponse{}

	err = json.Unmarshal(data, &varFederationUpdateFederatedRealmResponse)

	if err != nil {
		return err
	}

	*o = FederationUpdateFederatedRealmResponse(varFederationUpdateFederatedRealmResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFederationUpdateFederatedRealmResponse struct {
	value *FederationUpdateFederatedRealmResponse
	isSet bool
}

func (v NullableFederationUpdateFederatedRealmResponse) Get() *FederationUpdateFederatedRealmResponse {
	return v.value
}

func (v *NullableFederationUpdateFederatedRealmResponse) Set(val *FederationUpdateFederatedRealmResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFederationUpdateFederatedRealmResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFederationUpdateFederatedRealmResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFederationUpdateFederatedRealmResponse(val *FederationUpdateFederatedRealmResponse) *NullableFederationUpdateFederatedRealmResponse {
	return &NullableFederationUpdateFederatedRealmResponse{value: val, isSet: true}
}

func (v NullableFederationUpdateFederatedRealmResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFederationUpdateFederatedRealmResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}