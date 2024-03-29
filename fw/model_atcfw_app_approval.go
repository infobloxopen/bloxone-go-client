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

// checks if the AtcfwAppApproval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AtcfwAppApproval{}

// AtcfwAppApproval struct for AtcfwAppApproval
type AtcfwAppApproval struct {
	// The name of the application, should be unique and is used as Application Identifier
	AppName *string `json:"app_name,omitempty"`
	Status  *string `json:"status,omitempty"`
}

// NewAtcfwAppApproval instantiates a new AtcfwAppApproval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAtcfwAppApproval() *AtcfwAppApproval {
	this := AtcfwAppApproval{}
	return &this
}

// NewAtcfwAppApprovalWithDefaults instantiates a new AtcfwAppApproval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAtcfwAppApprovalWithDefaults() *AtcfwAppApproval {
	this := AtcfwAppApproval{}
	return &this
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *AtcfwAppApproval) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtcfwAppApproval) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *AtcfwAppApproval) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *AtcfwAppApproval) SetAppName(v string) {
	o.AppName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AtcfwAppApproval) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtcfwAppApproval) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AtcfwAppApproval) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AtcfwAppApproval) SetStatus(v string) {
	o.Status = &v
}

func (o AtcfwAppApproval) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AtcfwAppApproval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppName) {
		toSerialize["app_name"] = o.AppName
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableAtcfwAppApproval struct {
	value *AtcfwAppApproval
	isSet bool
}

func (v NullableAtcfwAppApproval) Get() *AtcfwAppApproval {
	return v.value
}

func (v *NullableAtcfwAppApproval) Set(val *AtcfwAppApproval) {
	v.value = val
	v.isSet = true
}

func (v NullableAtcfwAppApproval) IsSet() bool {
	return v.isSet
}

func (v *NullableAtcfwAppApproval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAtcfwAppApproval(val *AtcfwAppApproval) *NullableAtcfwAppApproval {
	return &NullableAtcfwAppApproval{value: val, isSet: true}
}

func (v NullableAtcfwAppApproval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAtcfwAppApproval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
