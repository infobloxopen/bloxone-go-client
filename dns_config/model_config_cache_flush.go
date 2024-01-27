/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns_config

import (
	"encoding/json"
)

// checks if the ConfigCacheFlush type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigCacheFlush{}

// ConfigCacheFlush The _dns/cache_flush_ API removes entries from the DNS cache on the on-prem host. The command will be forwarded to the on-prem host and executed asynchronously. The on-prem host must be available and running DNS for this to succeed.
type ConfigCacheFlush struct {
	// Optional. If _true_, all names below the given FQDN will also be removed from cache.  Defaults to _true_.
	FlushSubdomains *bool `json:"flush_subdomains,omitempty"`
	// Optional. The FQDN to remove.  Defaults to '.'
	Fqdn *string `json:"fqdn,omitempty"`
	// The host to alter. Either _ophid_ or _service_id_ should be provided.
	Ophid *string `json:"ophid,omitempty"`
	// Service Id. Either _ophid_ or _service_id_ should be provided.
	ServiceId *string `json:"service_id,omitempty"`
	// Optional. The time in seconds the command is valid for. Command is executed on the onprem host only if it takes less than this time for the command to be transmitted to the host. Otherwise the onprem host discards this command.  Defaults to 120 (2 min).
	Ttl *int64 `json:"ttl,omitempty"`
	// Optional, If provided, flushes the server's cache for a view.
	ViewName *string `json:"view_name,omitempty"`
}

// NewConfigCacheFlush instantiates a new ConfigCacheFlush object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigCacheFlush() *ConfigCacheFlush {
	this := ConfigCacheFlush{}
	return &this
}

// NewConfigCacheFlushWithDefaults instantiates a new ConfigCacheFlush object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigCacheFlushWithDefaults() *ConfigCacheFlush {
	this := ConfigCacheFlush{}
	return &this
}

// GetFlushSubdomains returns the FlushSubdomains field value if set, zero value otherwise.
func (o *ConfigCacheFlush) GetFlushSubdomains() bool {
	if o == nil || IsNil(o.FlushSubdomains) {
		var ret bool
		return ret
	}
	return *o.FlushSubdomains
}

// GetFlushSubdomainsOk returns a tuple with the FlushSubdomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigCacheFlush) GetFlushSubdomainsOk() (*bool, bool) {
	if o == nil || IsNil(o.FlushSubdomains) {
		return nil, false
	}
	return o.FlushSubdomains, true
}

// HasFlushSubdomains returns a boolean if a field has been set.
func (o *ConfigCacheFlush) HasFlushSubdomains() bool {
	if o != nil && !IsNil(o.FlushSubdomains) {
		return true
	}

	return false
}

// SetFlushSubdomains gets a reference to the given bool and assigns it to the FlushSubdomains field.
func (o *ConfigCacheFlush) SetFlushSubdomains(v bool) {
	o.FlushSubdomains = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *ConfigCacheFlush) GetFqdn() string {
	if o == nil || IsNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigCacheFlush) GetFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqdn) {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *ConfigCacheFlush) HasFqdn() bool {
	if o != nil && !IsNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *ConfigCacheFlush) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetOphid returns the Ophid field value if set, zero value otherwise.
func (o *ConfigCacheFlush) GetOphid() string {
	if o == nil || IsNil(o.Ophid) {
		var ret string
		return ret
	}
	return *o.Ophid
}

// GetOphidOk returns a tuple with the Ophid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigCacheFlush) GetOphidOk() (*string, bool) {
	if o == nil || IsNil(o.Ophid) {
		return nil, false
	}
	return o.Ophid, true
}

// HasOphid returns a boolean if a field has been set.
func (o *ConfigCacheFlush) HasOphid() bool {
	if o != nil && !IsNil(o.Ophid) {
		return true
	}

	return false
}

// SetOphid gets a reference to the given string and assigns it to the Ophid field.
func (o *ConfigCacheFlush) SetOphid(v string) {
	o.Ophid = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *ConfigCacheFlush) GetServiceId() string {
	if o == nil || IsNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigCacheFlush) GetServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *ConfigCacheFlush) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *ConfigCacheFlush) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *ConfigCacheFlush) GetTtl() int64 {
	if o == nil || IsNil(o.Ttl) {
		var ret int64
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigCacheFlush) GetTtlOk() (*int64, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *ConfigCacheFlush) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int64 and assigns it to the Ttl field.
func (o *ConfigCacheFlush) SetTtl(v int64) {
	o.Ttl = &v
}

// GetViewName returns the ViewName field value if set, zero value otherwise.
func (o *ConfigCacheFlush) GetViewName() string {
	if o == nil || IsNil(o.ViewName) {
		var ret string
		return ret
	}
	return *o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigCacheFlush) GetViewNameOk() (*string, bool) {
	if o == nil || IsNil(o.ViewName) {
		return nil, false
	}
	return o.ViewName, true
}

// HasViewName returns a boolean if a field has been set.
func (o *ConfigCacheFlush) HasViewName() bool {
	if o != nil && !IsNil(o.ViewName) {
		return true
	}

	return false
}

// SetViewName gets a reference to the given string and assigns it to the ViewName field.
func (o *ConfigCacheFlush) SetViewName(v string) {
	o.ViewName = &v
}

func (o ConfigCacheFlush) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigCacheFlush) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FlushSubdomains) {
		toSerialize["flush_subdomains"] = o.FlushSubdomains
	}
	if !IsNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !IsNil(o.Ophid) {
		toSerialize["ophid"] = o.Ophid
	}
	if !IsNil(o.ServiceId) {
		toSerialize["service_id"] = o.ServiceId
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.ViewName) {
		toSerialize["view_name"] = o.ViewName
	}
	return toSerialize, nil
}

type NullableConfigCacheFlush struct {
	value *ConfigCacheFlush
	isSet bool
}

func (v NullableConfigCacheFlush) Get() *ConfigCacheFlush {
	return v.value
}

func (v *NullableConfigCacheFlush) Set(val *ConfigCacheFlush) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigCacheFlush) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigCacheFlush) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigCacheFlush(val *ConfigCacheFlush) *NullableConfigCacheFlush {
	return &NullableConfigCacheFlush{value: val, isSet: true}
}

func (v NullableConfigCacheFlush) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigCacheFlush) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
