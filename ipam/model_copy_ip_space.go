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

// checks if the CopyIPSpace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CopyIPSpace{}

// CopyIPSpace struct for CopyIPSpace
type CopyIPSpace struct {
	// The description for the copied IP space. May contain 0 to 1024 characters. Can include UTF-8.
	Comment *string `json:"comment,omitempty"`
	// Indicates whether dhcp options should be copied or not when _ipam/ip_space_ object is copied.  Defaults to _false_.
	CopyDhcpOptions *bool `json:"copy_dhcp_options,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// The name for the copied IP space. Must contain 1 to 256 characters. Can include UTF-8.
	Name string `json:"name"`
	// Indicates whether copying should skip an object in case of error and continue with next, or abort copying in case of error.  Defaults to _false_.
	SkipOnError *bool `json:"skip_on_error,omitempty"`
}

type _CopyIPSpace CopyIPSpace

// NewCopyIPSpace instantiates a new CopyIPSpace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCopyIPSpace(name string) *CopyIPSpace {
	this := CopyIPSpace{}
	this.Name = name
	return &this
}

// NewCopyIPSpaceWithDefaults instantiates a new CopyIPSpace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCopyIPSpaceWithDefaults() *CopyIPSpace {
	this := CopyIPSpace{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CopyIPSpace) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyIPSpace) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CopyIPSpace) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CopyIPSpace) SetComment(v string) {
	o.Comment = &v
}

// GetCopyDhcpOptions returns the CopyDhcpOptions field value if set, zero value otherwise.
func (o *CopyIPSpace) GetCopyDhcpOptions() bool {
	if o == nil || IsNil(o.CopyDhcpOptions) {
		var ret bool
		return ret
	}
	return *o.CopyDhcpOptions
}

// GetCopyDhcpOptionsOk returns a tuple with the CopyDhcpOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyIPSpace) GetCopyDhcpOptionsOk() (*bool, bool) {
	if o == nil || IsNil(o.CopyDhcpOptions) {
		return nil, false
	}
	return o.CopyDhcpOptions, true
}

// HasCopyDhcpOptions returns a boolean if a field has been set.
func (o *CopyIPSpace) HasCopyDhcpOptions() bool {
	if o != nil && !IsNil(o.CopyDhcpOptions) {
		return true
	}

	return false
}

// SetCopyDhcpOptions gets a reference to the given bool and assigns it to the CopyDhcpOptions field.
func (o *CopyIPSpace) SetCopyDhcpOptions(v bool) {
	o.CopyDhcpOptions = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CopyIPSpace) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyIPSpace) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CopyIPSpace) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CopyIPSpace) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *CopyIPSpace) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CopyIPSpace) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CopyIPSpace) SetName(v string) {
	o.Name = v
}

// GetSkipOnError returns the SkipOnError field value if set, zero value otherwise.
func (o *CopyIPSpace) GetSkipOnError() bool {
	if o == nil || IsNil(o.SkipOnError) {
		var ret bool
		return ret
	}
	return *o.SkipOnError
}

// GetSkipOnErrorOk returns a tuple with the SkipOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyIPSpace) GetSkipOnErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipOnError) {
		return nil, false
	}
	return o.SkipOnError, true
}

// HasSkipOnError returns a boolean if a field has been set.
func (o *CopyIPSpace) HasSkipOnError() bool {
	if o != nil && !IsNil(o.SkipOnError) {
		return true
	}

	return false
}

// SetSkipOnError gets a reference to the given bool and assigns it to the SkipOnError field.
func (o *CopyIPSpace) SetSkipOnError(v bool) {
	o.SkipOnError = &v
}

func (o CopyIPSpace) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CopyIPSpace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CopyDhcpOptions) {
		toSerialize["copy_dhcp_options"] = o.CopyDhcpOptions
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.SkipOnError) {
		toSerialize["skip_on_error"] = o.SkipOnError
	}
	return toSerialize, nil
}

func (o *CopyIPSpace) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCopyIPSpace := _CopyIPSpace{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCopyIPSpace)

	if err != nil {
		return err
	}

	*o = CopyIPSpace(varCopyIPSpace)

	return err
}

type NullableCopyIPSpace struct {
	value *CopyIPSpace
	isSet bool
}

func (v NullableCopyIPSpace) Get() *CopyIPSpace {
	return v.value
}

func (v *NullableCopyIPSpace) Set(val *CopyIPSpace) {
	v.value = val
	v.isSet = true
}

func (v NullableCopyIPSpace) IsSet() bool {
	return v.isSet
}

func (v *NullableCopyIPSpace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCopyIPSpace(val *CopyIPSpace) *NullableCopyIPSpace {
	return &NullableCopyIPSpace{value: val, isSet: true}
}

func (v NullableCopyIPSpace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCopyIPSpace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
