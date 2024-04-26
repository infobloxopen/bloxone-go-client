/*
Infrastructure Management API

The **Infrastructure Management API** provides a RESTful interface to manage Infrastructure Hosts and Services objects.  The following is a list of the different Services and their string types (the string types are to be used with the APIs for the `service_type` field):  | Service name | Service type |   | ------ | ------ |   | Access Authentication | authn |   | Anycast | anycast |   | Data Connector | cdc |   | DHCP | dhcp |   | DNS | dns |   | DNS Forwarding Proxy | dfp |   | NIOS Grid Connector | orpheus |   | MS AD Sync | msad |   | NTP | ntp |   | BGP | bgp |   | RIP | rip |   | OSPF | ospf |    ---   ### Hosts API  The Hosts API is used to manage the Infrastructure Host resources. These include various operations related to hosts such as viewing, creating, updating, replacing, disconnecting, and deleting Hosts. Management of Hosts is done from the Cloud Services Portal (CSP) by navigating to the Manage -> Infrastructure -> Hosts tab.  ---   ### Services API  The Services API is used to manage the Infrastructure Service resources (a.k.a. BloxOne applications). These include various operations related to hosts such as viewing, creating, updating, starting/stopping, configuring, and deleting Services. Management of Services is done from the Cloud Services Portal (CSP) by navigating to the Manage -> Infrastructure -> Services tab.  ---   ### Detail APIs  The Detail APIs are read-only APIs used to list all the Infrastructure resources (Hosts and Services). Each resource record returned also contains information about its other associated resources and the status information for itself and the associated resource(s) (i.e., Host/Service status).  ---

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package inframgmt

import (
	"encoding/json"
)

// checks if the PoolInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoolInfo{}

// PoolInfo struct for PoolInfo
type PoolInfo struct {
	// Composite Status of the Pool that this resource belongs to (`online`, `degraded`, `error`, `unavailable`).
	CompositeStatus *string `json:"composite_status,omitempty"`
	// The resource identifier.
	PoolId *string `json:"pool_id,omitempty"`
	// Name of the Pool that this resource belongs to.
	PoolName             *string `json:"pool_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PoolInfo PoolInfo

// NewPoolInfo instantiates a new PoolInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolInfo() *PoolInfo {
	this := PoolInfo{}
	return &this
}

// NewPoolInfoWithDefaults instantiates a new PoolInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolInfoWithDefaults() *PoolInfo {
	this := PoolInfo{}
	return &this
}

// GetCompositeStatus returns the CompositeStatus field value if set, zero value otherwise.
func (o *PoolInfo) GetCompositeStatus() string {
	if o == nil || IsNil(o.CompositeStatus) {
		var ret string
		return ret
	}
	return *o.CompositeStatus
}

// GetCompositeStatusOk returns a tuple with the CompositeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolInfo) GetCompositeStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CompositeStatus) {
		return nil, false
	}
	return o.CompositeStatus, true
}

// HasCompositeStatus returns a boolean if a field has been set.
func (o *PoolInfo) HasCompositeStatus() bool {
	if o != nil && !IsNil(o.CompositeStatus) {
		return true
	}

	return false
}

// SetCompositeStatus gets a reference to the given string and assigns it to the CompositeStatus field.
func (o *PoolInfo) SetCompositeStatus(v string) {
	o.CompositeStatus = &v
}

// GetPoolId returns the PoolId field value if set, zero value otherwise.
func (o *PoolInfo) GetPoolId() string {
	if o == nil || IsNil(o.PoolId) {
		var ret string
		return ret
	}
	return *o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolInfo) GetPoolIdOk() (*string, bool) {
	if o == nil || IsNil(o.PoolId) {
		return nil, false
	}
	return o.PoolId, true
}

// HasPoolId returns a boolean if a field has been set.
func (o *PoolInfo) HasPoolId() bool {
	if o != nil && !IsNil(o.PoolId) {
		return true
	}

	return false
}

// SetPoolId gets a reference to the given string and assigns it to the PoolId field.
func (o *PoolInfo) SetPoolId(v string) {
	o.PoolId = &v
}

// GetPoolName returns the PoolName field value if set, zero value otherwise.
func (o *PoolInfo) GetPoolName() string {
	if o == nil || IsNil(o.PoolName) {
		var ret string
		return ret
	}
	return *o.PoolName
}

// GetPoolNameOk returns a tuple with the PoolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolInfo) GetPoolNameOk() (*string, bool) {
	if o == nil || IsNil(o.PoolName) {
		return nil, false
	}
	return o.PoolName, true
}

// HasPoolName returns a boolean if a field has been set.
func (o *PoolInfo) HasPoolName() bool {
	if o != nil && !IsNil(o.PoolName) {
		return true
	}

	return false
}

// SetPoolName gets a reference to the given string and assigns it to the PoolName field.
func (o *PoolInfo) SetPoolName(v string) {
	o.PoolName = &v
}

func (o PoolInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompositeStatus) {
		toSerialize["composite_status"] = o.CompositeStatus
	}
	if !IsNil(o.PoolId) {
		toSerialize["pool_id"] = o.PoolId
	}
	if !IsNil(o.PoolName) {
		toSerialize["pool_name"] = o.PoolName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PoolInfo) UnmarshalJSON(data []byte) (err error) {
	varPoolInfo := _PoolInfo{}

	err = json.Unmarshal(data, &varPoolInfo)

	if err != nil {
		return err
	}

	*o = PoolInfo(varPoolInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "composite_status")
		delete(additionalProperties, "pool_id")
		delete(additionalProperties, "pool_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePoolInfo struct {
	value *PoolInfo
	isSet bool
}

func (v NullablePoolInfo) Get() *PoolInfo {
	return v.value
}

func (v *NullablePoolInfo) Set(val *PoolInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolInfo(val *PoolInfo) *NullablePoolInfo {
	return &NullablePoolInfo{value: val, isSet: true}
}

func (v NullablePoolInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
