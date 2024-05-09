/*
Schedule Software/Config Updates

Infoblox by default does automatic software updates when they become available. Updates are applied to all on-prem hosts, physical or virtual. However, you can override and schedule the software updates. You can also defer the updates to a later date and time. You can configure up to a total of 50 deferrals (scheduled and deferred software updates), which means you have the flexibility to create up to 50 update groups across different on-prem hosts by mapping with appropriate tags. Tags are be used to associate deferrals (scheduled or deferred) with a specific or group of onprem-hosts. Apart from software update deferrals, config update deferrals also can be configured using these overrides.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package upgradepolicy

import (
	"encoding/json"
)

// checks if the ServiceV2DeleteMaintenanceWindowResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceV2DeleteMaintenanceWindowResponse{}

// ServiceV2DeleteMaintenanceWindowResponse struct for ServiceV2DeleteMaintenanceWindowResponse
type ServiceV2DeleteMaintenanceWindowResponse struct {
	WindowType           *string `json:"window_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceV2DeleteMaintenanceWindowResponse ServiceV2DeleteMaintenanceWindowResponse

// NewServiceV2DeleteMaintenanceWindowResponse instantiates a new ServiceV2DeleteMaintenanceWindowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceV2DeleteMaintenanceWindowResponse() *ServiceV2DeleteMaintenanceWindowResponse {
	this := ServiceV2DeleteMaintenanceWindowResponse{}
	return &this
}

// NewServiceV2DeleteMaintenanceWindowResponseWithDefaults instantiates a new ServiceV2DeleteMaintenanceWindowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceV2DeleteMaintenanceWindowResponseWithDefaults() *ServiceV2DeleteMaintenanceWindowResponse {
	this := ServiceV2DeleteMaintenanceWindowResponse{}
	return &this
}

// GetWindowType returns the WindowType field value if set, zero value otherwise.
func (o *ServiceV2DeleteMaintenanceWindowResponse) GetWindowType() string {
	if o == nil || IsNil(o.WindowType) {
		var ret string
		return ret
	}
	return *o.WindowType
}

// GetWindowTypeOk returns a tuple with the WindowType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceV2DeleteMaintenanceWindowResponse) GetWindowTypeOk() (*string, bool) {
	if o == nil || IsNil(o.WindowType) {
		return nil, false
	}
	return o.WindowType, true
}

// HasWindowType returns a boolean if a field has been set.
func (o *ServiceV2DeleteMaintenanceWindowResponse) HasWindowType() bool {
	if o != nil && !IsNil(o.WindowType) {
		return true
	}

	return false
}

// SetWindowType gets a reference to the given string and assigns it to the WindowType field.
func (o *ServiceV2DeleteMaintenanceWindowResponse) SetWindowType(v string) {
	o.WindowType = &v
}

func (o ServiceV2DeleteMaintenanceWindowResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceV2DeleteMaintenanceWindowResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WindowType) {
		toSerialize["window_type"] = o.WindowType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceV2DeleteMaintenanceWindowResponse) UnmarshalJSON(data []byte) (err error) {
	varServiceV2DeleteMaintenanceWindowResponse := _ServiceV2DeleteMaintenanceWindowResponse{}

	err = json.Unmarshal(data, &varServiceV2DeleteMaintenanceWindowResponse)

	if err != nil {
		return err
	}

	*o = ServiceV2DeleteMaintenanceWindowResponse(varServiceV2DeleteMaintenanceWindowResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "window_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceV2DeleteMaintenanceWindowResponse struct {
	value *ServiceV2DeleteMaintenanceWindowResponse
	isSet bool
}

func (v NullableServiceV2DeleteMaintenanceWindowResponse) Get() *ServiceV2DeleteMaintenanceWindowResponse {
	return v.value
}

func (v *NullableServiceV2DeleteMaintenanceWindowResponse) Set(val *ServiceV2DeleteMaintenanceWindowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceV2DeleteMaintenanceWindowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceV2DeleteMaintenanceWindowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceV2DeleteMaintenanceWindowResponse(val *ServiceV2DeleteMaintenanceWindowResponse) *NullableServiceV2DeleteMaintenanceWindowResponse {
	return &NullableServiceV2DeleteMaintenanceWindowResponse{value: val, isSet: true}
}

func (v NullableServiceV2DeleteMaintenanceWindowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceV2DeleteMaintenanceWindowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
