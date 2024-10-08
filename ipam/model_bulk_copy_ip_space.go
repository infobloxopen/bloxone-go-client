/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
	"fmt"
)

// checks if the BulkCopyIPSpace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkCopyIPSpace{}

// BulkCopyIPSpace struct for BulkCopyIPSpace
type BulkCopyIPSpace struct {
	// Indicates whether dhcp options for IPv4 should be copied or not when objects (_ipam/address_block_ and _ipam/subnet_ only) are copied.  Defaults to _false_.
	CopyDhcpOptions *bool `json:"copy_dhcp_options,omitempty"`
	// The resource identifier.
	CopyObjects []string `json:"copy_objects"`
	// Indicates whether child objects should be copied or not.  Defaults to _false_.
	Recursive *bool `json:"recursive,omitempty"`
	// Indicates whether the child objects are going to retain their compartment_id, or inherit from the object to copy into.  Defaults to false
	RetainChildCompartment *bool `json:"retain_child_compartment,omitempty"`
	// Indicates whether copying should skip object in case of error and continue with next, or abort copying in case of error.  Defaults to _false_.
	SkipOnError *bool `json:"skip_on_error,omitempty"`
	// The resource identifier.
	Target               string `json:"target"`
	AdditionalProperties map[string]interface{}
}

type _BulkCopyIPSpace BulkCopyIPSpace

// NewBulkCopyIPSpace instantiates a new BulkCopyIPSpace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkCopyIPSpace(copyObjects []string, target string) *BulkCopyIPSpace {
	this := BulkCopyIPSpace{}
	this.CopyObjects = copyObjects
	this.Target = target
	return &this
}

// NewBulkCopyIPSpaceWithDefaults instantiates a new BulkCopyIPSpace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkCopyIPSpaceWithDefaults() *BulkCopyIPSpace {
	this := BulkCopyIPSpace{}
	return &this
}

// GetCopyDhcpOptions returns the CopyDhcpOptions field value if set, zero value otherwise.
func (o *BulkCopyIPSpace) GetCopyDhcpOptions() bool {
	if o == nil || IsNil(o.CopyDhcpOptions) {
		var ret bool
		return ret
	}
	return *o.CopyDhcpOptions
}

// GetCopyDhcpOptionsOk returns a tuple with the CopyDhcpOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCopyIPSpace) GetCopyDhcpOptionsOk() (*bool, bool) {
	if o == nil || IsNil(o.CopyDhcpOptions) {
		return nil, false
	}
	return o.CopyDhcpOptions, true
}

// HasCopyDhcpOptions returns a boolean if a field has been set.
func (o *BulkCopyIPSpace) HasCopyDhcpOptions() bool {
	if o != nil && !IsNil(o.CopyDhcpOptions) {
		return true
	}

	return false
}

// SetCopyDhcpOptions gets a reference to the given bool and assigns it to the CopyDhcpOptions field.
func (o *BulkCopyIPSpace) SetCopyDhcpOptions(v bool) {
	o.CopyDhcpOptions = &v
}

// GetCopyObjects returns the CopyObjects field value
func (o *BulkCopyIPSpace) GetCopyObjects() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CopyObjects
}

// GetCopyObjectsOk returns a tuple with the CopyObjects field value
// and a boolean to check if the value has been set.
func (o *BulkCopyIPSpace) GetCopyObjectsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CopyObjects, true
}

// SetCopyObjects sets field value
func (o *BulkCopyIPSpace) SetCopyObjects(v []string) {
	o.CopyObjects = v
}

// GetRecursive returns the Recursive field value if set, zero value otherwise.
func (o *BulkCopyIPSpace) GetRecursive() bool {
	if o == nil || IsNil(o.Recursive) {
		var ret bool
		return ret
	}
	return *o.Recursive
}

// GetRecursiveOk returns a tuple with the Recursive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCopyIPSpace) GetRecursiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Recursive) {
		return nil, false
	}
	return o.Recursive, true
}

// HasRecursive returns a boolean if a field has been set.
func (o *BulkCopyIPSpace) HasRecursive() bool {
	if o != nil && !IsNil(o.Recursive) {
		return true
	}

	return false
}

// SetRecursive gets a reference to the given bool and assigns it to the Recursive field.
func (o *BulkCopyIPSpace) SetRecursive(v bool) {
	o.Recursive = &v
}

// GetRetainChildCompartment returns the RetainChildCompartment field value if set, zero value otherwise.
func (o *BulkCopyIPSpace) GetRetainChildCompartment() bool {
	if o == nil || IsNil(o.RetainChildCompartment) {
		var ret bool
		return ret
	}
	return *o.RetainChildCompartment
}

// GetRetainChildCompartmentOk returns a tuple with the RetainChildCompartment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCopyIPSpace) GetRetainChildCompartmentOk() (*bool, bool) {
	if o == nil || IsNil(o.RetainChildCompartment) {
		return nil, false
	}
	return o.RetainChildCompartment, true
}

// HasRetainChildCompartment returns a boolean if a field has been set.
func (o *BulkCopyIPSpace) HasRetainChildCompartment() bool {
	if o != nil && !IsNil(o.RetainChildCompartment) {
		return true
	}

	return false
}

// SetRetainChildCompartment gets a reference to the given bool and assigns it to the RetainChildCompartment field.
func (o *BulkCopyIPSpace) SetRetainChildCompartment(v bool) {
	o.RetainChildCompartment = &v
}

// GetSkipOnError returns the SkipOnError field value if set, zero value otherwise.
func (o *BulkCopyIPSpace) GetSkipOnError() bool {
	if o == nil || IsNil(o.SkipOnError) {
		var ret bool
		return ret
	}
	return *o.SkipOnError
}

// GetSkipOnErrorOk returns a tuple with the SkipOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCopyIPSpace) GetSkipOnErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipOnError) {
		return nil, false
	}
	return o.SkipOnError, true
}

// HasSkipOnError returns a boolean if a field has been set.
func (o *BulkCopyIPSpace) HasSkipOnError() bool {
	if o != nil && !IsNil(o.SkipOnError) {
		return true
	}

	return false
}

// SetSkipOnError gets a reference to the given bool and assigns it to the SkipOnError field.
func (o *BulkCopyIPSpace) SetSkipOnError(v bool) {
	o.SkipOnError = &v
}

// GetTarget returns the Target field value
func (o *BulkCopyIPSpace) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *BulkCopyIPSpace) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *BulkCopyIPSpace) SetTarget(v string) {
	o.Target = v
}

func (o BulkCopyIPSpace) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkCopyIPSpace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CopyDhcpOptions) {
		toSerialize["copy_dhcp_options"] = o.CopyDhcpOptions
	}
	toSerialize["copy_objects"] = o.CopyObjects
	if !IsNil(o.Recursive) {
		toSerialize["recursive"] = o.Recursive
	}
	if !IsNil(o.RetainChildCompartment) {
		toSerialize["retain_child_compartment"] = o.RetainChildCompartment
	}
	if !IsNil(o.SkipOnError) {
		toSerialize["skip_on_error"] = o.SkipOnError
	}
	toSerialize["target"] = o.Target

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkCopyIPSpace) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"copy_objects",
		"target",
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

	varBulkCopyIPSpace := _BulkCopyIPSpace{}

	err = json.Unmarshal(data, &varBulkCopyIPSpace)

	if err != nil {
		return err
	}

	*o = BulkCopyIPSpace(varBulkCopyIPSpace)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "copy_dhcp_options")
		delete(additionalProperties, "copy_objects")
		delete(additionalProperties, "recursive")
		delete(additionalProperties, "retain_child_compartment")
		delete(additionalProperties, "skip_on_error")
		delete(additionalProperties, "target")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkCopyIPSpace struct {
	value *BulkCopyIPSpace
	isSet bool
}

func (v NullableBulkCopyIPSpace) Get() *BulkCopyIPSpace {
	return v.value
}

func (v *NullableBulkCopyIPSpace) Set(val *BulkCopyIPSpace) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkCopyIPSpace) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkCopyIPSpace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkCopyIPSpace(val *BulkCopyIPSpace) *NullableBulkCopyIPSpace {
	return &NullableBulkCopyIPSpace{value: val, isSet: true}
}

func (v NullableBulkCopyIPSpace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkCopyIPSpace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
