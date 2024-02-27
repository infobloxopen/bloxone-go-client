/*
BloxOne FW API

BloxOne Threat Defense Cloud is an extension of the BloxOne Suite that provides visibility into infected and compromised off-premises devices, roaming users, remote sites, and branch offices. You can subscribe to BloxOne Cloud and use its functionality to mitigate and control malware as well as provide unprecedented insight into your network security posture and enable timely action. BloxOne Cloud also offers unified policy management, reporting, and threat analytics across the entire spectrum. Using automated and high-quality threat intelligence feeds and unique behavioral analytics, it automatically stops device communications with C&Cs/botnets and prevents DNS based data exfiltration.  The mission-critical DNS infrastructure can become a vulnerable component in your network when it is inadequately protected by traditional security solutions and consequently used as an attack surface. Compromised DNS services can result in catastrophic network and system failures. To fully protect your network in today’s cyber security threat environment, Infoblox sets a new DNS security standard by offering scalable, enterprise-grade, and integrated protection for your DNS infrastructure.  Through the Infoblox Cloud Services Portal, you can view the status of your subscription and threat intelligence feeds, manage your network scope and roaming end users, and learn more about threats on your networks through the Infoblox Threat Lookup tool and predefined reports.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fw

import (
	"encoding/json"
)

// checks if the NamedListItemsInsertOrReplaceNamedListItems400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NamedListItemsInsertOrReplaceNamedListItems400Response{}

// NamedListItemsInsertOrReplaceNamedListItems400Response struct for NamedListItemsInsertOrReplaceNamedListItems400Response
type NamedListItemsInsertOrReplaceNamedListItems400Response struct {
	Error *NamedListItemsInsertOrReplaceNamedListItems400ResponseError `json:"error,omitempty"`
}

// NewNamedListItemsInsertOrReplaceNamedListItems400Response instantiates a new NamedListItemsInsertOrReplaceNamedListItems400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamedListItemsInsertOrReplaceNamedListItems400Response() *NamedListItemsInsertOrReplaceNamedListItems400Response {
	this := NamedListItemsInsertOrReplaceNamedListItems400Response{}
	return &this
}

// NewNamedListItemsInsertOrReplaceNamedListItems400ResponseWithDefaults instantiates a new NamedListItemsInsertOrReplaceNamedListItems400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamedListItemsInsertOrReplaceNamedListItems400ResponseWithDefaults() *NamedListItemsInsertOrReplaceNamedListItems400Response {
	this := NamedListItemsInsertOrReplaceNamedListItems400Response{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *NamedListItemsInsertOrReplaceNamedListItems400Response) GetError() NamedListItemsInsertOrReplaceNamedListItems400ResponseError {
	if o == nil || IsNil(o.Error) {
		var ret NamedListItemsInsertOrReplaceNamedListItems400ResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamedListItemsInsertOrReplaceNamedListItems400Response) GetErrorOk() (*NamedListItemsInsertOrReplaceNamedListItems400ResponseError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *NamedListItemsInsertOrReplaceNamedListItems400Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given NamedListItemsInsertOrReplaceNamedListItems400ResponseError and assigns it to the Error field.
func (o *NamedListItemsInsertOrReplaceNamedListItems400Response) SetError(v NamedListItemsInsertOrReplaceNamedListItems400ResponseError) {
	o.Error = &v
}

func (o NamedListItemsInsertOrReplaceNamedListItems400Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NamedListItemsInsertOrReplaceNamedListItems400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableNamedListItemsInsertOrReplaceNamedListItems400Response struct {
	value *NamedListItemsInsertOrReplaceNamedListItems400Response
	isSet bool
}

func (v NullableNamedListItemsInsertOrReplaceNamedListItems400Response) Get() *NamedListItemsInsertOrReplaceNamedListItems400Response {
	return v.value
}

func (v *NullableNamedListItemsInsertOrReplaceNamedListItems400Response) Set(val *NamedListItemsInsertOrReplaceNamedListItems400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableNamedListItemsInsertOrReplaceNamedListItems400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableNamedListItemsInsertOrReplaceNamedListItems400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamedListItemsInsertOrReplaceNamedListItems400Response(val *NamedListItemsInsertOrReplaceNamedListItems400Response) *NullableNamedListItemsInsertOrReplaceNamedListItems400Response {
	return &NullableNamedListItemsInsertOrReplaceNamedListItems400Response{value: val, isSet: true}
}

func (v NullableNamedListItemsInsertOrReplaceNamedListItems400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamedListItemsInsertOrReplaceNamedListItems400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
