/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the IpamsvcHostName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcHostName{}

// IpamsvcHostName The __HostName__ object represents a name associated with the __Host__ object.
type IpamsvcHostName struct {
	// When _true_, the name is treated as an alias.
	Alias *bool `json:"alias,omitempty"`
	// A name for the host.
	Name string `json:"name"`
	// When _true_, the name field is treated as primary name. There must be one and only one primary name in the list of host names. The primary name will be treated as the canonical name for all the aliases. PTR record will be generated only for the primary name.
	PrimaryName *bool `json:"primary_name,omitempty"`
	// The resource identifier.
	Zone string `json:"zone"`
}

type _IpamsvcHostName IpamsvcHostName

// NewIpamsvcHostName instantiates a new IpamsvcHostName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcHostName(name string, zone string) *IpamsvcHostName {
	this := IpamsvcHostName{}
	this.Name = name
	this.Zone = zone
	return &this
}

// NewIpamsvcHostNameWithDefaults instantiates a new IpamsvcHostName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcHostNameWithDefaults() *IpamsvcHostName {
	this := IpamsvcHostName{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *IpamsvcHostName) GetAlias() bool {
	if o == nil || IsNil(o.Alias) {
		var ret bool
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHostName) GetAliasOk() (*bool, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *IpamsvcHostName) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given bool and assigns it to the Alias field.
func (o *IpamsvcHostName) SetAlias(v bool) {
	o.Alias = &v
}

// GetName returns the Name field value
func (o *IpamsvcHostName) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IpamsvcHostName) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IpamsvcHostName) SetName(v string) {
	o.Name = v
}

// GetPrimaryName returns the PrimaryName field value if set, zero value otherwise.
func (o *IpamsvcHostName) GetPrimaryName() bool {
	if o == nil || IsNil(o.PrimaryName) {
		var ret bool
		return ret
	}
	return *o.PrimaryName
}

// GetPrimaryNameOk returns a tuple with the PrimaryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHostName) GetPrimaryNameOk() (*bool, bool) {
	if o == nil || IsNil(o.PrimaryName) {
		return nil, false
	}
	return o.PrimaryName, true
}

// HasPrimaryName returns a boolean if a field has been set.
func (o *IpamsvcHostName) HasPrimaryName() bool {
	if o != nil && !IsNil(o.PrimaryName) {
		return true
	}

	return false
}

// SetPrimaryName gets a reference to the given bool and assigns it to the PrimaryName field.
func (o *IpamsvcHostName) SetPrimaryName(v bool) {
	o.PrimaryName = &v
}

// GetZone returns the Zone field value
func (o *IpamsvcHostName) GetZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Zone
}

// GetZoneOk returns a tuple with the Zone field value
// and a boolean to check if the value has been set.
func (o *IpamsvcHostName) GetZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Zone, true
}

// SetZone sets field value
func (o *IpamsvcHostName) SetZone(v string) {
	o.Zone = v
}

func (o IpamsvcHostName) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcHostName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.PrimaryName) {
		toSerialize["primary_name"] = o.PrimaryName
	}
	toSerialize["zone"] = o.Zone
	return toSerialize, nil
}

func (o *IpamsvcHostName) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"zone",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varIpamsvcHostName := _IpamsvcHostName{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIpamsvcHostName)

	if err != nil {
		return err
	}

	*o = IpamsvcHostName(varIpamsvcHostName)

	return err
}

type NullableIpamsvcHostName struct {
	value *IpamsvcHostName
	isSet bool
}

func (v NullableIpamsvcHostName) Get() *IpamsvcHostName {
	return v.value
}

func (v *NullableIpamsvcHostName) Set(val *IpamsvcHostName) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcHostName) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcHostName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcHostName(val *IpamsvcHostName) *NullableIpamsvcHostName {
	return &NullableIpamsvcHostName{value: val, isSet: true}
}

func (v NullableIpamsvcHostName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcHostName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
