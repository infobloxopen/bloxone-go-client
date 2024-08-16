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

// checks if the FederationCreateFederatedRealmResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FederationCreateFederatedRealmResponse{}

// FederationCreateFederatedRealmResponse The response format to create the __FederatedRealm__ object.
type FederationCreateFederatedRealmResponse struct {
	// The created Federated Realm object.
	Result               *FederationFederatedRealm `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FederationCreateFederatedRealmResponse FederationCreateFederatedRealmResponse

// NewFederationCreateFederatedRealmResponse instantiates a new FederationCreateFederatedRealmResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFederationCreateFederatedRealmResponse() *FederationCreateFederatedRealmResponse {
	this := FederationCreateFederatedRealmResponse{}
	return &this
}

// NewFederationCreateFederatedRealmResponseWithDefaults instantiates a new FederationCreateFederatedRealmResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFederationCreateFederatedRealmResponseWithDefaults() *FederationCreateFederatedRealmResponse {
	this := FederationCreateFederatedRealmResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *FederationCreateFederatedRealmResponse) GetResult() FederationFederatedRealm {
	if o == nil || IsNil(o.Result) {
		var ret FederationFederatedRealm
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationCreateFederatedRealmResponse) GetResultOk() (*FederationFederatedRealm, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *FederationCreateFederatedRealmResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given FederationFederatedRealm and assigns it to the Result field.
func (o *FederationCreateFederatedRealmResponse) SetResult(v FederationFederatedRealm) {
	o.Result = &v
}

func (o FederationCreateFederatedRealmResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FederationCreateFederatedRealmResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FederationCreateFederatedRealmResponse) UnmarshalJSON(data []byte) (err error) {
	varFederationCreateFederatedRealmResponse := _FederationCreateFederatedRealmResponse{}

	err = json.Unmarshal(data, &varFederationCreateFederatedRealmResponse)

	if err != nil {
		return err
	}

	*o = FederationCreateFederatedRealmResponse(varFederationCreateFederatedRealmResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFederationCreateFederatedRealmResponse struct {
	value *FederationCreateFederatedRealmResponse
	isSet bool
}

func (v NullableFederationCreateFederatedRealmResponse) Get() *FederationCreateFederatedRealmResponse {
	return v.value
}

func (v *NullableFederationCreateFederatedRealmResponse) Set(val *FederationCreateFederatedRealmResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFederationCreateFederatedRealmResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFederationCreateFederatedRealmResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFederationCreateFederatedRealmResponse(val *FederationCreateFederatedRealmResponse) *NullableFederationCreateFederatedRealmResponse {
	return &NullableFederationCreateFederatedRealmResponse{value: val, isSet: true}
}

func (v NullableFederationCreateFederatedRealmResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFederationCreateFederatedRealmResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
