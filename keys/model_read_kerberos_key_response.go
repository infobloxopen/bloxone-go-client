/*
DDI Keys API

The DDI Keys application is a BloxOne DDI service for managing TSIG keys and GSS-TSIG (Kerberos) keys which are used by other BloxOne DDI applications. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package keys

import (
	"encoding/json"
)

// checks if the ReadKerberosKeyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadKerberosKeyResponse{}

// ReadKerberosKeyResponse The response format to retrieve the __KerberosKey__ resource extracted from the uploaded keytab file.
type ReadKerberosKeyResponse struct {
	// The __KerberosKey__ object.
	Result               *KerberosKey `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReadKerberosKeyResponse ReadKerberosKeyResponse

// NewReadKerberosKeyResponse instantiates a new ReadKerberosKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadKerberosKeyResponse() *ReadKerberosKeyResponse {
	this := ReadKerberosKeyResponse{}
	return &this
}

// NewReadKerberosKeyResponseWithDefaults instantiates a new ReadKerberosKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadKerberosKeyResponseWithDefaults() *ReadKerberosKeyResponse {
	this := ReadKerberosKeyResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ReadKerberosKeyResponse) GetResult() KerberosKey {
	if o == nil || IsNil(o.Result) {
		var ret KerberosKey
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadKerberosKeyResponse) GetResultOk() (*KerberosKey, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ReadKerberosKeyResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given KerberosKey and assigns it to the Result field.
func (o *ReadKerberosKeyResponse) SetResult(v KerberosKey) {
	o.Result = &v
}

func (o ReadKerberosKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadKerberosKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReadKerberosKeyResponse) UnmarshalJSON(data []byte) (err error) {
	varReadKerberosKeyResponse := _ReadKerberosKeyResponse{}

	err = json.Unmarshal(data, &varReadKerberosKeyResponse)

	if err != nil {
		return err
	}

	*o = ReadKerberosKeyResponse(varReadKerberosKeyResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReadKerberosKeyResponse struct {
	value *ReadKerberosKeyResponse
	isSet bool
}

func (v NullableReadKerberosKeyResponse) Get() *ReadKerberosKeyResponse {
	return v.value
}

func (v *NullableReadKerberosKeyResponse) Set(val *ReadKerberosKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadKerberosKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadKerberosKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadKerberosKeyResponse(val *ReadKerberosKeyResponse) *NullableReadKerberosKeyResponse {
	return &NullableReadKerberosKeyResponse{value: val, isSet: true}
}

func (v NullableReadKerberosKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadKerberosKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
