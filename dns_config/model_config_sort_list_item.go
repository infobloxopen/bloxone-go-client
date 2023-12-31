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

// checks if the ConfigSortListItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigSortListItem{}

// ConfigSortListItem Element in a SortList.
type ConfigSortListItem struct {
	// The resource identifier.
	Acl *string `json:"acl,omitempty"`
	// Type of element.  Allowed values:  * _any_,  * _ip_,  * _acl_,
	Element string `json:"element"`
	// Optional. The prioritized networks. If empty, the value of _source_ or networks from _acl_ is used.
	PrioritizedNetworks []string `json:"prioritized_networks,omitempty"`
	// Must be empty if _element_ is not _ip_.
	Source *string `json:"source,omitempty"`
}

// NewConfigSortListItem instantiates a new ConfigSortListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigSortListItem(element string) *ConfigSortListItem {
	this := ConfigSortListItem{}
	this.Element = element
	return &this
}

// NewConfigSortListItemWithDefaults instantiates a new ConfigSortListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigSortListItemWithDefaults() *ConfigSortListItem {
	this := ConfigSortListItem{}
	return &this
}

// GetAcl returns the Acl field value if set, zero value otherwise.
func (o *ConfigSortListItem) GetAcl() string {
	if o == nil || IsNil(o.Acl) {
		var ret string
		return ret
	}
	return *o.Acl
}

// GetAclOk returns a tuple with the Acl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSortListItem) GetAclOk() (*string, bool) {
	if o == nil || IsNil(o.Acl) {
		return nil, false
	}
	return o.Acl, true
}

// HasAcl returns a boolean if a field has been set.
func (o *ConfigSortListItem) HasAcl() bool {
	if o != nil && !IsNil(o.Acl) {
		return true
	}

	return false
}

// SetAcl gets a reference to the given string and assigns it to the Acl field.
func (o *ConfigSortListItem) SetAcl(v string) {
	o.Acl = &v
}

// GetElement returns the Element field value
func (o *ConfigSortListItem) GetElement() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Element
}

// GetElementOk returns a tuple with the Element field value
// and a boolean to check if the value has been set.
func (o *ConfigSortListItem) GetElementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Element, true
}

// SetElement sets field value
func (o *ConfigSortListItem) SetElement(v string) {
	o.Element = v
}

// GetPrioritizedNetworks returns the PrioritizedNetworks field value if set, zero value otherwise.
func (o *ConfigSortListItem) GetPrioritizedNetworks() []string {
	if o == nil || IsNil(o.PrioritizedNetworks) {
		var ret []string
		return ret
	}
	return o.PrioritizedNetworks
}

// GetPrioritizedNetworksOk returns a tuple with the PrioritizedNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSortListItem) GetPrioritizedNetworksOk() ([]string, bool) {
	if o == nil || IsNil(o.PrioritizedNetworks) {
		return nil, false
	}
	return o.PrioritizedNetworks, true
}

// HasPrioritizedNetworks returns a boolean if a field has been set.
func (o *ConfigSortListItem) HasPrioritizedNetworks() bool {
	if o != nil && !IsNil(o.PrioritizedNetworks) {
		return true
	}

	return false
}

// SetPrioritizedNetworks gets a reference to the given []string and assigns it to the PrioritizedNetworks field.
func (o *ConfigSortListItem) SetPrioritizedNetworks(v []string) {
	o.PrioritizedNetworks = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ConfigSortListItem) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSortListItem) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ConfigSortListItem) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ConfigSortListItem) SetSource(v string) {
	o.Source = &v
}

func (o ConfigSortListItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigSortListItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Acl) {
		toSerialize["acl"] = o.Acl
	}
	toSerialize["element"] = o.Element
	if !IsNil(o.PrioritizedNetworks) {
		toSerialize["prioritized_networks"] = o.PrioritizedNetworks
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableConfigSortListItem struct {
	value *ConfigSortListItem
	isSet bool
}

func (v NullableConfigSortListItem) Get() *ConfigSortListItem {
	return v.value
}

func (v *NullableConfigSortListItem) Set(val *ConfigSortListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigSortListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigSortListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigSortListItem(val *ConfigSortListItem) *NullableConfigSortListItem {
	return &NullableConfigSortListItem{value: val, isSet: true}
}

func (v NullableConfigSortListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigSortListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
