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

// checks if the ConfigHostAssociatedServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigHostAssociatedServer{}

// ConfigHostAssociatedServer struct for ConfigHostAssociatedServer
type ConfigHostAssociatedServer struct {
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// DNS server name.
	Name *string `json:"name,omitempty"`
}

// NewConfigHostAssociatedServer instantiates a new ConfigHostAssociatedServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigHostAssociatedServer() *ConfigHostAssociatedServer {
	this := ConfigHostAssociatedServer{}
	return &this
}

// NewConfigHostAssociatedServerWithDefaults instantiates a new ConfigHostAssociatedServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigHostAssociatedServerWithDefaults() *ConfigHostAssociatedServer {
	this := ConfigHostAssociatedServer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConfigHostAssociatedServer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHostAssociatedServer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConfigHostAssociatedServer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConfigHostAssociatedServer) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConfigHostAssociatedServer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigHostAssociatedServer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConfigHostAssociatedServer) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConfigHostAssociatedServer) SetName(v string) {
	o.Name = &v
}

func (o ConfigHostAssociatedServer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigHostAssociatedServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableConfigHostAssociatedServer struct {
	value *ConfigHostAssociatedServer
	isSet bool
}

func (v NullableConfigHostAssociatedServer) Get() *ConfigHostAssociatedServer {
	return v.value
}

func (v *NullableConfigHostAssociatedServer) Set(val *ConfigHostAssociatedServer) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigHostAssociatedServer) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigHostAssociatedServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigHostAssociatedServer(val *ConfigHostAssociatedServer) *NullableConfigHostAssociatedServer {
	return &NullableConfigHostAssociatedServer{value: val, isSet: true}
}

func (v NullableConfigHostAssociatedServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigHostAssociatedServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
