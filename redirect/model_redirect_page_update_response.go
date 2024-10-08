/*
BloxOne Redirect API

You can configure BloxOne Threat Defense Cloud to redirect traffic to the Infoblox server that displays the default or customized redirect page. You can redirect traffic to a custom destination using custom redirects.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package redirect

import (
	"encoding/json"
)

// checks if the RedirectPageUpdateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RedirectPageUpdateResponse{}

// RedirectPageUpdateResponse The Redirect Page update response.
type RedirectPageUpdateResponse struct {
	// The Redirect Page object.
	Results              *RedirectPage `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RedirectPageUpdateResponse RedirectPageUpdateResponse

// NewRedirectPageUpdateResponse instantiates a new RedirectPageUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedirectPageUpdateResponse() *RedirectPageUpdateResponse {
	this := RedirectPageUpdateResponse{}
	return &this
}

// NewRedirectPageUpdateResponseWithDefaults instantiates a new RedirectPageUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedirectPageUpdateResponseWithDefaults() *RedirectPageUpdateResponse {
	this := RedirectPageUpdateResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *RedirectPageUpdateResponse) GetResults() RedirectPage {
	if o == nil || IsNil(o.Results) {
		var ret RedirectPage
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectPageUpdateResponse) GetResultsOk() (*RedirectPage, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *RedirectPageUpdateResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given RedirectPage and assigns it to the Results field.
func (o *RedirectPageUpdateResponse) SetResults(v RedirectPage) {
	o.Results = &v
}

func (o RedirectPageUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedirectPageUpdateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RedirectPageUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	varRedirectPageUpdateResponse := _RedirectPageUpdateResponse{}

	err = json.Unmarshal(data, &varRedirectPageUpdateResponse)

	if err != nil {
		return err
	}

	*o = RedirectPageUpdateResponse(varRedirectPageUpdateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRedirectPageUpdateResponse struct {
	value *RedirectPageUpdateResponse
	isSet bool
}

func (v NullableRedirectPageUpdateResponse) Get() *RedirectPageUpdateResponse {
	return v.value
}

func (v *NullableRedirectPageUpdateResponse) Set(val *RedirectPageUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRedirectPageUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRedirectPageUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedirectPageUpdateResponse(val *RedirectPageUpdateResponse) *NullableRedirectPageUpdateResponse {
	return &NullableRedirectPageUpdateResponse{value: val, isSet: true}
}

func (v NullableRedirectPageUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedirectPageUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
