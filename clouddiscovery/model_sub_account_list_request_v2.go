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

// checks if the SubAccountListRequestV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubAccountListRequestV2{}

// SubAccountListRequestV2 struct for SubAccountListRequestV2
type SubAccountListRequestV2 struct {
	AccessType   *string `json:"access_type,omitempty"`
	CredentialId *string `json:"credential_id,omitempty"`
	// atlas.api.field_selection
	Fields                    *string                   `json:"fields,omitempty"`
	ProviderCredentialsConfig *SubAccountProvCredConfig `json:"provider_credentials_config,omitempty"`
	ProviderType              *string                   `json:"provider_type,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _SubAccountListRequestV2 SubAccountListRequestV2

// NewSubAccountListRequestV2 instantiates a new SubAccountListRequestV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubAccountListRequestV2() *SubAccountListRequestV2 {
	this := SubAccountListRequestV2{}
	return &this
}

// NewSubAccountListRequestV2WithDefaults instantiates a new SubAccountListRequestV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubAccountListRequestV2WithDefaults() *SubAccountListRequestV2 {
	this := SubAccountListRequestV2{}
	return &this
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *SubAccountListRequestV2) GetAccessType() string {
	if o == nil || IsNil(o.AccessType) {
		var ret string
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountListRequestV2) GetAccessTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *SubAccountListRequestV2) HasAccessType() bool {
	if o != nil && !IsNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given string and assigns it to the AccessType field.
func (o *SubAccountListRequestV2) SetAccessType(v string) {
	o.AccessType = &v
}

// GetCredentialId returns the CredentialId field value if set, zero value otherwise.
func (o *SubAccountListRequestV2) GetCredentialId() string {
	if o == nil || IsNil(o.CredentialId) {
		var ret string
		return ret
	}
	return *o.CredentialId
}

// GetCredentialIdOk returns a tuple with the CredentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountListRequestV2) GetCredentialIdOk() (*string, bool) {
	if o == nil || IsNil(o.CredentialId) {
		return nil, false
	}
	return o.CredentialId, true
}

// HasCredentialId returns a boolean if a field has been set.
func (o *SubAccountListRequestV2) HasCredentialId() bool {
	if o != nil && !IsNil(o.CredentialId) {
		return true
	}

	return false
}

// SetCredentialId gets a reference to the given string and assigns it to the CredentialId field.
func (o *SubAccountListRequestV2) SetCredentialId(v string) {
	o.CredentialId = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *SubAccountListRequestV2) GetFields() string {
	if o == nil || IsNil(o.Fields) {
		var ret string
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountListRequestV2) GetFieldsOk() (*string, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *SubAccountListRequestV2) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given string and assigns it to the Fields field.
func (o *SubAccountListRequestV2) SetFields(v string) {
	o.Fields = &v
}

// GetProviderCredentialsConfig returns the ProviderCredentialsConfig field value if set, zero value otherwise.
func (o *SubAccountListRequestV2) GetProviderCredentialsConfig() SubAccountProvCredConfig {
	if o == nil || IsNil(o.ProviderCredentialsConfig) {
		var ret SubAccountProvCredConfig
		return ret
	}
	return *o.ProviderCredentialsConfig
}

// GetProviderCredentialsConfigOk returns a tuple with the ProviderCredentialsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountListRequestV2) GetProviderCredentialsConfigOk() (*SubAccountProvCredConfig, bool) {
	if o == nil || IsNil(o.ProviderCredentialsConfig) {
		return nil, false
	}
	return o.ProviderCredentialsConfig, true
}

// HasProviderCredentialsConfig returns a boolean if a field has been set.
func (o *SubAccountListRequestV2) HasProviderCredentialsConfig() bool {
	if o != nil && !IsNil(o.ProviderCredentialsConfig) {
		return true
	}

	return false
}

// SetProviderCredentialsConfig gets a reference to the given SubAccountProvCredConfig and assigns it to the ProviderCredentialsConfig field.
func (o *SubAccountListRequestV2) SetProviderCredentialsConfig(v SubAccountProvCredConfig) {
	o.ProviderCredentialsConfig = &v
}

// GetProviderType returns the ProviderType field value if set, zero value otherwise.
func (o *SubAccountListRequestV2) GetProviderType() string {
	if o == nil || IsNil(o.ProviderType) {
		var ret string
		return ret
	}
	return *o.ProviderType
}

// GetProviderTypeOk returns a tuple with the ProviderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountListRequestV2) GetProviderTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderType) {
		return nil, false
	}
	return o.ProviderType, true
}

// HasProviderType returns a boolean if a field has been set.
func (o *SubAccountListRequestV2) HasProviderType() bool {
	if o != nil && !IsNil(o.ProviderType) {
		return true
	}

	return false
}

// SetProviderType gets a reference to the given string and assigns it to the ProviderType field.
func (o *SubAccountListRequestV2) SetProviderType(v string) {
	o.ProviderType = &v
}

func (o SubAccountListRequestV2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubAccountListRequestV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessType) {
		toSerialize["access_type"] = o.AccessType
	}
	if !IsNil(o.CredentialId) {
		toSerialize["credential_id"] = o.CredentialId
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.ProviderCredentialsConfig) {
		toSerialize["provider_credentials_config"] = o.ProviderCredentialsConfig
	}
	if !IsNil(o.ProviderType) {
		toSerialize["provider_type"] = o.ProviderType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubAccountListRequestV2) UnmarshalJSON(data []byte) (err error) {
	varSubAccountListRequestV2 := _SubAccountListRequestV2{}

	err = json.Unmarshal(data, &varSubAccountListRequestV2)

	if err != nil {
		return err
	}

	*o = SubAccountListRequestV2(varSubAccountListRequestV2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "access_type")
		delete(additionalProperties, "credential_id")
		delete(additionalProperties, "fields")
		delete(additionalProperties, "provider_credentials_config")
		delete(additionalProperties, "provider_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubAccountListRequestV2 struct {
	value *SubAccountListRequestV2
	isSet bool
}

func (v NullableSubAccountListRequestV2) Get() *SubAccountListRequestV2 {
	return v.value
}

func (v *NullableSubAccountListRequestV2) Set(val *SubAccountListRequestV2) {
	v.value = val
	v.isSet = true
}

func (v NullableSubAccountListRequestV2) IsSet() bool {
	return v.isSet
}

func (v *NullableSubAccountListRequestV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubAccountListRequestV2(val *SubAccountListRequestV2) *NullableSubAccountListRequestV2 {
	return &NullableSubAccountListRequestV2{value: val, isSet: true}
}

func (v NullableSubAccountListRequestV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubAccountListRequestV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
