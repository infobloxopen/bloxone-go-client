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

// checks if the ServiceV2UpdateMaintenanceWindowRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceV2UpdateMaintenanceWindowRequest{}

// ServiceV2UpdateMaintenanceWindowRequest struct for ServiceV2UpdateMaintenanceWindowRequest
type ServiceV2UpdateMaintenanceWindowRequest struct {
	Id                   *string                           `json:"id,omitempty"`
	Payload              *ServiceV2UpdateMaintenanceWindow `json:"payload,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceV2UpdateMaintenanceWindowRequest ServiceV2UpdateMaintenanceWindowRequest

// NewServiceV2UpdateMaintenanceWindowRequest instantiates a new ServiceV2UpdateMaintenanceWindowRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceV2UpdateMaintenanceWindowRequest() *ServiceV2UpdateMaintenanceWindowRequest {
	this := ServiceV2UpdateMaintenanceWindowRequest{}
	return &this
}

// NewServiceV2UpdateMaintenanceWindowRequestWithDefaults instantiates a new ServiceV2UpdateMaintenanceWindowRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceV2UpdateMaintenanceWindowRequestWithDefaults() *ServiceV2UpdateMaintenanceWindowRequest {
	this := ServiceV2UpdateMaintenanceWindowRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceV2UpdateMaintenanceWindowRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceV2UpdateMaintenanceWindowRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceV2UpdateMaintenanceWindowRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceV2UpdateMaintenanceWindowRequest) SetId(v string) {
	o.Id = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *ServiceV2UpdateMaintenanceWindowRequest) GetPayload() ServiceV2UpdateMaintenanceWindow {
	if o == nil || IsNil(o.Payload) {
		var ret ServiceV2UpdateMaintenanceWindow
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceV2UpdateMaintenanceWindowRequest) GetPayloadOk() (*ServiceV2UpdateMaintenanceWindow, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *ServiceV2UpdateMaintenanceWindowRequest) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given ServiceV2UpdateMaintenanceWindow and assigns it to the Payload field.
func (o *ServiceV2UpdateMaintenanceWindowRequest) SetPayload(v ServiceV2UpdateMaintenanceWindow) {
	o.Payload = &v
}

func (o ServiceV2UpdateMaintenanceWindowRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceV2UpdateMaintenanceWindowRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceV2UpdateMaintenanceWindowRequest) UnmarshalJSON(data []byte) (err error) {
	varServiceV2UpdateMaintenanceWindowRequest := _ServiceV2UpdateMaintenanceWindowRequest{}

	err = json.Unmarshal(data, &varServiceV2UpdateMaintenanceWindowRequest)

	if err != nil {
		return err
	}

	*o = ServiceV2UpdateMaintenanceWindowRequest(varServiceV2UpdateMaintenanceWindowRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "payload")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceV2UpdateMaintenanceWindowRequest struct {
	value *ServiceV2UpdateMaintenanceWindowRequest
	isSet bool
}

func (v NullableServiceV2UpdateMaintenanceWindowRequest) Get() *ServiceV2UpdateMaintenanceWindowRequest {
	return v.value
}

func (v *NullableServiceV2UpdateMaintenanceWindowRequest) Set(val *ServiceV2UpdateMaintenanceWindowRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceV2UpdateMaintenanceWindowRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceV2UpdateMaintenanceWindowRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceV2UpdateMaintenanceWindowRequest(val *ServiceV2UpdateMaintenanceWindowRequest) *NullableServiceV2UpdateMaintenanceWindowRequest {
	return &NullableServiceV2UpdateMaintenanceWindowRequest{value: val, isSet: true}
}

func (v NullableServiceV2UpdateMaintenanceWindowRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceV2UpdateMaintenanceWindowRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
