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

// checks if the AtcfwNamedListItemsInsertOrUpdateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AtcfwNamedListItemsInsertOrUpdateResponse{}

// AtcfwNamedListItemsInsertOrUpdateResponse The Named List Items create or update response.
type AtcfwNamedListItemsInsertOrUpdateResponse struct {
	Success *AtcfwNamedListItemsInsertOrUpdateResponseSuccess `json:"success,omitempty"`
}

// NewAtcfwNamedListItemsInsertOrUpdateResponse instantiates a new AtcfwNamedListItemsInsertOrUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAtcfwNamedListItemsInsertOrUpdateResponse() *AtcfwNamedListItemsInsertOrUpdateResponse {
	this := AtcfwNamedListItemsInsertOrUpdateResponse{}
	return &this
}

// NewAtcfwNamedListItemsInsertOrUpdateResponseWithDefaults instantiates a new AtcfwNamedListItemsInsertOrUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAtcfwNamedListItemsInsertOrUpdateResponseWithDefaults() *AtcfwNamedListItemsInsertOrUpdateResponse {
	this := AtcfwNamedListItemsInsertOrUpdateResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *AtcfwNamedListItemsInsertOrUpdateResponse) GetSuccess() AtcfwNamedListItemsInsertOrUpdateResponseSuccess {
	if o == nil || IsNil(o.Success) {
		var ret AtcfwNamedListItemsInsertOrUpdateResponseSuccess
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtcfwNamedListItemsInsertOrUpdateResponse) GetSuccessOk() (*AtcfwNamedListItemsInsertOrUpdateResponseSuccess, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *AtcfwNamedListItemsInsertOrUpdateResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given AtcfwNamedListItemsInsertOrUpdateResponseSuccess and assigns it to the Success field.
func (o *AtcfwNamedListItemsInsertOrUpdateResponse) SetSuccess(v AtcfwNamedListItemsInsertOrUpdateResponseSuccess) {
	o.Success = &v
}

func (o AtcfwNamedListItemsInsertOrUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AtcfwNamedListItemsInsertOrUpdateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableAtcfwNamedListItemsInsertOrUpdateResponse struct {
	value *AtcfwNamedListItemsInsertOrUpdateResponse
	isSet bool
}

func (v NullableAtcfwNamedListItemsInsertOrUpdateResponse) Get() *AtcfwNamedListItemsInsertOrUpdateResponse {
	return v.value
}

func (v *NullableAtcfwNamedListItemsInsertOrUpdateResponse) Set(val *AtcfwNamedListItemsInsertOrUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAtcfwNamedListItemsInsertOrUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAtcfwNamedListItemsInsertOrUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAtcfwNamedListItemsInsertOrUpdateResponse(val *AtcfwNamedListItemsInsertOrUpdateResponse) *NullableAtcfwNamedListItemsInsertOrUpdateResponse {
	return &NullableAtcfwNamedListItemsInsertOrUpdateResponse{value: val, isSet: true}
}

func (v NullableAtcfwNamedListItemsInsertOrUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAtcfwNamedListItemsInsertOrUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
