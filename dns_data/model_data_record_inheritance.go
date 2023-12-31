/*
DNS Data API

The DNS Data is a BloxOne DDI service providing primary authoritative zone support. DNS Data is authoritative for all DNS resource records and is acting as a primary DNS server. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns_data

import (
	"encoding/json"
)

// checks if the DataRecordInheritance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataRecordInheritance{}

// DataRecordInheritance The inheritance configuration specifies how the _Record_ object inherits the _ttl_ field.
type DataRecordInheritance struct {
	Ttl *Inheritance2InheritedUInt32 `json:"ttl,omitempty"`
}

// NewDataRecordInheritance instantiates a new DataRecordInheritance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataRecordInheritance() *DataRecordInheritance {
	this := DataRecordInheritance{}
	return &this
}

// NewDataRecordInheritanceWithDefaults instantiates a new DataRecordInheritance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataRecordInheritanceWithDefaults() *DataRecordInheritance {
	this := DataRecordInheritance{}
	return &this
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *DataRecordInheritance) GetTtl() Inheritance2InheritedUInt32 {
	if o == nil || IsNil(o.Ttl) {
		var ret Inheritance2InheritedUInt32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRecordInheritance) GetTtlOk() (*Inheritance2InheritedUInt32, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *DataRecordInheritance) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given Inheritance2InheritedUInt32 and assigns it to the Ttl field.
func (o *DataRecordInheritance) SetTtl(v Inheritance2InheritedUInt32) {
	o.Ttl = &v
}

func (o DataRecordInheritance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataRecordInheritance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	return toSerialize, nil
}

type NullableDataRecordInheritance struct {
	value *DataRecordInheritance
	isSet bool
}

func (v NullableDataRecordInheritance) Get() *DataRecordInheritance {
	return v.value
}

func (v *NullableDataRecordInheritance) Set(val *DataRecordInheritance) {
	v.value = val
	v.isSet = true
}

func (v NullableDataRecordInheritance) IsSet() bool {
	return v.isSet
}

func (v *NullableDataRecordInheritance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataRecordInheritance(val *DataRecordInheritance) *NullableDataRecordInheritance {
	return &NullableDataRecordInheritance{value: val, isSet: true}
}

func (v NullableDataRecordInheritance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataRecordInheritance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
