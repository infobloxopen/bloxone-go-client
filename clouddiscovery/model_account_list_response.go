/*
Discovery Configuration API V2

The Discovery configuration service is a BloxOne Service that provides configuration for accessing and syncing the Cloud assets   Base Paths:  1. provider: **_/api/cloud_discovery/v2/_**

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clouddiscovery

import (
	"encoding/json"
)

// checks if the AccountListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountListResponse{}

// AccountListResponse The Account object List response format.
type AccountListResponse struct {
	Page                 *ApiPageInfo `json:"page,omitempty"`
	Results              []Account    `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountListResponse AccountListResponse

// NewAccountListResponse instantiates a new AccountListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountListResponse() *AccountListResponse {
	this := AccountListResponse{}
	return &this
}

// NewAccountListResponseWithDefaults instantiates a new AccountListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountListResponseWithDefaults() *AccountListResponse {
	this := AccountListResponse{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *AccountListResponse) GetPage() ApiPageInfo {
	if o == nil || IsNil(o.Page) {
		var ret ApiPageInfo
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountListResponse) GetPageOk() (*ApiPageInfo, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *AccountListResponse) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given ApiPageInfo and assigns it to the Page field.
func (o *AccountListResponse) SetPage(v ApiPageInfo) {
	o.Page = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *AccountListResponse) GetResults() []Account {
	if o == nil || IsNil(o.Results) {
		var ret []Account
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountListResponse) GetResultsOk() ([]Account, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *AccountListResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Account and assigns it to the Results field.
func (o *AccountListResponse) SetResults(v []Account) {
	o.Results = v
}

func (o AccountListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountListResponse) UnmarshalJSON(data []byte) (err error) {
	varAccountListResponse := _AccountListResponse{}

	err = json.Unmarshal(data, &varAccountListResponse)

	if err != nil {
		return err
	}

	*o = AccountListResponse(varAccountListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "page")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountListResponse struct {
	value *AccountListResponse
	isSet bool
}

func (v NullableAccountListResponse) Get() *AccountListResponse {
	return v.value
}

func (v *NullableAccountListResponse) Set(val *AccountListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountListResponse(val *AccountListResponse) *NullableAccountListResponse {
	return &NullableAccountListResponse{value: val, isSet: true}
}

func (v NullableAccountListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
