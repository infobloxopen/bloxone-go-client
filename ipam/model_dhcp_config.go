/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the DHCPConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DHCPConfig{}

// DHCPConfig A DHCP Config object (_dhcp/dhcp_config_) represents a shared DHCP configuration that controls how leases are issued.
type DHCPConfig struct {
	// The abandoned reclaim time in seconds for IPV4 clients.
	AbandonedReclaimTime *int64 `json:"abandoned_reclaim_time,omitempty"`
	// The abandoned reclaim time in seconds for IPV6 clients.
	AbandonedReclaimTimeV6 *int64 `json:"abandoned_reclaim_time_v6,omitempty"`
	// Disable to allow leases only for known IPv4 clients, those for which a fixed address is configured.
	AllowUnknown *bool `json:"allow_unknown,omitempty"`
	// Disable to allow leases only for known IPV6 clients, those for which a fixed address is configured.
	AllowUnknownV6 *bool `json:"allow_unknown_v6,omitempty"`
	// Enable/disable to include/exclude the client id when responding to discover or request.
	EchoClientId *bool `json:"echo_client_id,omitempty"`
	// The resource identifier.
	Filters []string `json:"filters,omitempty"`
	// The resource identifier.
	FiltersV6 []string `json:"filters_v6,omitempty"`
	// Enable to ignore the client UID when issuing a DHCP lease. Use this option to prevent assigning two IP addresses for a client which does not have a UID during one phase of PXE boot but acquires one for the other phase.
	IgnoreClientUid *bool `json:"ignore_client_uid,omitempty"`
	// The list of clients to ignore requests from.
	IgnoreList []IgnoreItem `json:"ignore_list,omitempty"`
	// The lease duration in seconds.
	LeaseTime *int64 `json:"lease_time,omitempty"`
	// The lease duration in seconds for IPV6 clients.
	LeaseTimeV6 *int64 `json:"lease_time_v6,omitempty"`
}

// NewDHCPConfig instantiates a new DHCPConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDHCPConfig() *DHCPConfig {
	this := DHCPConfig{}
	var abandonedReclaimTime int64 = 3600
	this.AbandonedReclaimTime = &abandonedReclaimTime
	var abandonedReclaimTimeV6 int64 = 3600
	this.AbandonedReclaimTimeV6 = &abandonedReclaimTimeV6
	var allowUnknown bool = true
	this.AllowUnknown = &allowUnknown
	var allowUnknownV6 bool = true
	this.AllowUnknownV6 = &allowUnknownV6
	var echoClientId bool = false
	this.EchoClientId = &echoClientId
	var ignoreClientUid bool = false
	this.IgnoreClientUid = &ignoreClientUid
	var leaseTime int64 = 3600
	this.LeaseTime = &leaseTime
	var leaseTimeV6 int64 = 3600
	this.LeaseTimeV6 = &leaseTimeV6
	return &this
}

// NewDHCPConfigWithDefaults instantiates a new DHCPConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDHCPConfigWithDefaults() *DHCPConfig {
	this := DHCPConfig{}
	var abandonedReclaimTime int64 = 3600
	this.AbandonedReclaimTime = &abandonedReclaimTime
	var abandonedReclaimTimeV6 int64 = 3600
	this.AbandonedReclaimTimeV6 = &abandonedReclaimTimeV6
	var allowUnknown bool = true
	this.AllowUnknown = &allowUnknown
	var allowUnknownV6 bool = true
	this.AllowUnknownV6 = &allowUnknownV6
	var echoClientId bool = false
	this.EchoClientId = &echoClientId
	var ignoreClientUid bool = false
	this.IgnoreClientUid = &ignoreClientUid
	var leaseTime int64 = 3600
	this.LeaseTime = &leaseTime
	var leaseTimeV6 int64 = 3600
	this.LeaseTimeV6 = &leaseTimeV6
	return &this
}

// GetAbandonedReclaimTime returns the AbandonedReclaimTime field value if set, zero value otherwise.
func (o *DHCPConfig) GetAbandonedReclaimTime() int64 {
	if o == nil || IsNil(o.AbandonedReclaimTime) {
		var ret int64
		return ret
	}
	return *o.AbandonedReclaimTime
}

// GetAbandonedReclaimTimeOk returns a tuple with the AbandonedReclaimTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPConfig) GetAbandonedReclaimTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.AbandonedReclaimTime) {
		return nil, false
	}
	return o.AbandonedReclaimTime, true
}

// HasAbandonedReclaimTime returns a boolean if a field has been set.
func (o *DHCPConfig) HasAbandonedReclaimTime() bool {
	if o != nil && !IsNil(o.AbandonedReclaimTime) {
		return true
	}

	return false
}

// SetAbandonedReclaimTime gets a reference to the given int64 and assigns it to the AbandonedReclaimTime field.
func (o *DHCPConfig) SetAbandonedReclaimTime(v int64) {
	o.AbandonedReclaimTime = &v
}

// GetAbandonedReclaimTimeV6 returns the AbandonedReclaimTimeV6 field value if set, zero value otherwise.
func (o *DHCPConfig) GetAbandonedReclaimTimeV6() int64 {
	if o == nil || IsNil(o.AbandonedReclaimTimeV6) {
		var ret int64
		return ret
	}
	return *o.AbandonedReclaimTimeV6
}

// GetAbandonedReclaimTimeV6Ok returns a tuple with the AbandonedReclaimTimeV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPConfig) GetAbandonedReclaimTimeV6Ok() (*int64, bool) {
	if o == nil || IsNil(o.AbandonedReclaimTimeV6) {
		return nil, false
	}
	return o.AbandonedReclaimTimeV6, true
}

// HasAbandonedReclaimTimeV6 returns a boolean if a field has been set.
func (o *DHCPConfig) HasAbandonedReclaimTimeV6() bool {
	if o != nil && !IsNil(o.AbandonedReclaimTimeV6) {
		return true
	}

	return false
}

// SetAbandonedReclaimTimeV6 gets a reference to the given int64 and assigns it to the AbandonedReclaimTimeV6 field.
func (o *DHCPConfig) SetAbandonedReclaimTimeV6(v int64) {
	o.AbandonedReclaimTimeV6 = &v
}

// GetAllowUnknown returns the AllowUnknown field value if set, zero value otherwise.
func (o *DHCPConfig) GetAllowUnknown() bool {
	if o == nil || IsNil(o.AllowUnknown) {
		var ret bool
		return ret
	}
	return *o.AllowUnknown
}

// GetAllowUnknownOk returns a tuple with the AllowUnknown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPConfig) GetAllowUnknownOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowUnknown) {
		return nil, false
	}
	return o.AllowUnknown, true
}

// HasAllowUnknown returns a boolean if a field has been set.
func (o *DHCPConfig) HasAllowUnknown() bool {
	if o != nil && !IsNil(o.AllowUnknown) {
		return true
	}

	return false
}

// SetAllowUnknown gets a reference to the given bool and assigns it to the AllowUnknown field.
func (o *DHCPConfig) SetAllowUnknown(v bool) {
	o.AllowUnknown = &v
}

// GetAllowUnknownV6 returns the AllowUnknownV6 field value if set, zero value otherwise.
func (o *DHCPConfig) GetAllowUnknownV6() bool {
	if o == nil || IsNil(o.AllowUnknownV6) {
		var ret bool
		return ret
	}
	return *o.AllowUnknownV6
}

// GetAllowUnknownV6Ok returns a tuple with the AllowUnknownV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPConfig) GetAllowUnknownV6Ok() (*bool, bool) {
	if o == nil || IsNil(o.AllowUnknownV6) {
		return nil, false
	}
	return o.AllowUnknownV6, true
}

// HasAllowUnknownV6 returns a boolean if a field has been set.
func (o *DHCPConfig) HasAllowUnknownV6() bool {
	if o != nil && !IsNil(o.AllowUnknownV6) {
		return true
	}

	return false
}

// SetAllowUnknownV6 gets a reference to the given bool and assigns it to the AllowUnknownV6 field.
func (o *DHCPConfig) SetAllowUnknownV6(v bool) {
	o.AllowUnknownV6 = &v
}

// GetEchoClientId returns the EchoClientId field value if set, zero value otherwise.
func (o *DHCPConfig) GetEchoClientId() bool {
	if o == nil || IsNil(o.EchoClientId) {
		var ret bool
		return ret
	}
	return *o.EchoClientId
}

// GetEchoClientIdOk returns a tuple with the EchoClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPConfig) GetEchoClientIdOk() (*bool, bool) {
	if o == nil || IsNil(o.EchoClientId) {
		return nil, false
	}
	return o.EchoClientId, true
}

// HasEchoClientId returns a boolean if a field has been set.
func (o *DHCPConfig) HasEchoClientId() bool {
	if o != nil && !IsNil(o.EchoClientId) {
		return true
	}

	return false
}

// SetEchoClientId gets a reference to the given bool and assigns it to the EchoClientId field.
func (o *DHCPConfig) SetEchoClientId(v bool) {
	o.EchoClientId = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *DHCPConfig) GetFilters() []string {
	if o == nil || IsNil(o.Filters) {
		var ret []string
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPConfig) GetFiltersOk() ([]string, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *DHCPConfig) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []string and assigns it to the Filters field.
func (o *DHCPConfig) SetFilters(v []string) {
	o.Filters = v
}

// GetFiltersV6 returns the FiltersV6 field value if set, zero value otherwise.
func (o *DHCPConfig) GetFiltersV6() []string {
	if o == nil || IsNil(o.FiltersV6) {
		var ret []string
		return ret
	}
	return o.FiltersV6
}

// GetFiltersV6Ok returns a tuple with the FiltersV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPConfig) GetFiltersV6Ok() ([]string, bool) {
	if o == nil || IsNil(o.FiltersV6) {
		return nil, false
	}
	return o.FiltersV6, true
}

// HasFiltersV6 returns a boolean if a field has been set.
func (o *DHCPConfig) HasFiltersV6() bool {
	if o != nil && !IsNil(o.FiltersV6) {
		return true
	}

	return false
}

// SetFiltersV6 gets a reference to the given []string and assigns it to the FiltersV6 field.
func (o *DHCPConfig) SetFiltersV6(v []string) {
	o.FiltersV6 = v
}

// GetIgnoreClientUid returns the IgnoreClientUid field value if set, zero value otherwise.
func (o *DHCPConfig) GetIgnoreClientUid() bool {
	if o == nil || IsNil(o.IgnoreClientUid) {
		var ret bool
		return ret
	}
	return *o.IgnoreClientUid
}

// GetIgnoreClientUidOk returns a tuple with the IgnoreClientUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPConfig) GetIgnoreClientUidOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreClientUid) {
		return nil, false
	}
	return o.IgnoreClientUid, true
}

// HasIgnoreClientUid returns a boolean if a field has been set.
func (o *DHCPConfig) HasIgnoreClientUid() bool {
	if o != nil && !IsNil(o.IgnoreClientUid) {
		return true
	}

	return false
}

// SetIgnoreClientUid gets a reference to the given bool and assigns it to the IgnoreClientUid field.
func (o *DHCPConfig) SetIgnoreClientUid(v bool) {
	o.IgnoreClientUid = &v
}

// GetIgnoreList returns the IgnoreList field value if set, zero value otherwise.
func (o *DHCPConfig) GetIgnoreList() []IgnoreItem {
	if o == nil || IsNil(o.IgnoreList) {
		var ret []IgnoreItem
		return ret
	}
	return o.IgnoreList
}

// GetIgnoreListOk returns a tuple with the IgnoreList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPConfig) GetIgnoreListOk() ([]IgnoreItem, bool) {
	if o == nil || IsNil(o.IgnoreList) {
		return nil, false
	}
	return o.IgnoreList, true
}

// HasIgnoreList returns a boolean if a field has been set.
func (o *DHCPConfig) HasIgnoreList() bool {
	if o != nil && !IsNil(o.IgnoreList) {
		return true
	}

	return false
}

// SetIgnoreList gets a reference to the given []IgnoreItem and assigns it to the IgnoreList field.
func (o *DHCPConfig) SetIgnoreList(v []IgnoreItem) {
	o.IgnoreList = v
}

// GetLeaseTime returns the LeaseTime field value if set, zero value otherwise.
func (o *DHCPConfig) GetLeaseTime() int64 {
	if o == nil || IsNil(o.LeaseTime) {
		var ret int64
		return ret
	}
	return *o.LeaseTime
}

// GetLeaseTimeOk returns a tuple with the LeaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPConfig) GetLeaseTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.LeaseTime) {
		return nil, false
	}
	return o.LeaseTime, true
}

// HasLeaseTime returns a boolean if a field has been set.
func (o *DHCPConfig) HasLeaseTime() bool {
	if o != nil && !IsNil(o.LeaseTime) {
		return true
	}

	return false
}

// SetLeaseTime gets a reference to the given int64 and assigns it to the LeaseTime field.
func (o *DHCPConfig) SetLeaseTime(v int64) {
	o.LeaseTime = &v
}

// GetLeaseTimeV6 returns the LeaseTimeV6 field value if set, zero value otherwise.
func (o *DHCPConfig) GetLeaseTimeV6() int64 {
	if o == nil || IsNil(o.LeaseTimeV6) {
		var ret int64
		return ret
	}
	return *o.LeaseTimeV6
}

// GetLeaseTimeV6Ok returns a tuple with the LeaseTimeV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPConfig) GetLeaseTimeV6Ok() (*int64, bool) {
	if o == nil || IsNil(o.LeaseTimeV6) {
		return nil, false
	}
	return o.LeaseTimeV6, true
}

// HasLeaseTimeV6 returns a boolean if a field has been set.
func (o *DHCPConfig) HasLeaseTimeV6() bool {
	if o != nil && !IsNil(o.LeaseTimeV6) {
		return true
	}

	return false
}

// SetLeaseTimeV6 gets a reference to the given int64 and assigns it to the LeaseTimeV6 field.
func (o *DHCPConfig) SetLeaseTimeV6(v int64) {
	o.LeaseTimeV6 = &v
}

func (o DHCPConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DHCPConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AbandonedReclaimTime) {
		toSerialize["abandoned_reclaim_time"] = o.AbandonedReclaimTime
	}
	if !IsNil(o.AbandonedReclaimTimeV6) {
		toSerialize["abandoned_reclaim_time_v6"] = o.AbandonedReclaimTimeV6
	}
	if !IsNil(o.AllowUnknown) {
		toSerialize["allow_unknown"] = o.AllowUnknown
	}
	if !IsNil(o.AllowUnknownV6) {
		toSerialize["allow_unknown_v6"] = o.AllowUnknownV6
	}
	if !IsNil(o.EchoClientId) {
		toSerialize["echo_client_id"] = o.EchoClientId
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.FiltersV6) {
		toSerialize["filters_v6"] = o.FiltersV6
	}
	if !IsNil(o.IgnoreClientUid) {
		toSerialize["ignore_client_uid"] = o.IgnoreClientUid
	}
	if !IsNil(o.IgnoreList) {
		toSerialize["ignore_list"] = o.IgnoreList
	}
	if !IsNil(o.LeaseTime) {
		toSerialize["lease_time"] = o.LeaseTime
	}
	if !IsNil(o.LeaseTimeV6) {
		toSerialize["lease_time_v6"] = o.LeaseTimeV6
	}
	return toSerialize, nil
}

type NullableDHCPConfig struct {
	value *DHCPConfig
	isSet bool
}

func (v NullableDHCPConfig) Get() *DHCPConfig {
	return v.value
}

func (v *NullableDHCPConfig) Set(val *DHCPConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDHCPConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDHCPConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDHCPConfig(val *DHCPConfig) *NullableDHCPConfig {
	return &NullableDHCPConfig{value: val, isSet: true}
}

func (v NullableDHCPConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDHCPConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
