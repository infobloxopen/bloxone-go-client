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

// checks if the AtcfwListSeverityLevels type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AtcfwListSeverityLevels{}

// AtcfwListSeverityLevels The Payload for Patch Operation to update Threat/Confidence Levels and Tags in TI List
type AtcfwListSeverityLevels struct {
	// The confidence level for a TI list. The possible values are [LOW\", \"MEDIUM\", \"HIGH\"]
	ConfidenceLevel *string `json:"confidence_level,omitempty"`
	// The Named List object identifier.
	Id *int32 `json:"id,omitempty"`
	// Enables tag support for resource where tags attribute contains user-defined key value pairs
	Tags map[string]interface{} `json:"tags,omitempty"`
	// The threat level for a TI list. The possible values are [\"INFO\", \"LOW\", \"MEDIUM\", \"HIGH\"]
	ThreatLevel *string `json:"threat_level,omitempty"`
}

// NewAtcfwListSeverityLevels instantiates a new AtcfwListSeverityLevels object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAtcfwListSeverityLevels() *AtcfwListSeverityLevels {
	this := AtcfwListSeverityLevels{}
	return &this
}

// NewAtcfwListSeverityLevelsWithDefaults instantiates a new AtcfwListSeverityLevels object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAtcfwListSeverityLevelsWithDefaults() *AtcfwListSeverityLevels {
	this := AtcfwListSeverityLevels{}
	return &this
}

// GetConfidenceLevel returns the ConfidenceLevel field value if set, zero value otherwise.
func (o *AtcfwListSeverityLevels) GetConfidenceLevel() string {
	if o == nil || IsNil(o.ConfidenceLevel) {
		var ret string
		return ret
	}
	return *o.ConfidenceLevel
}

// GetConfidenceLevelOk returns a tuple with the ConfidenceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtcfwListSeverityLevels) GetConfidenceLevelOk() (*string, bool) {
	if o == nil || IsNil(o.ConfidenceLevel) {
		return nil, false
	}
	return o.ConfidenceLevel, true
}

// HasConfidenceLevel returns a boolean if a field has been set.
func (o *AtcfwListSeverityLevels) HasConfidenceLevel() bool {
	if o != nil && !IsNil(o.ConfidenceLevel) {
		return true
	}

	return false
}

// SetConfidenceLevel gets a reference to the given string and assigns it to the ConfidenceLevel field.
func (o *AtcfwListSeverityLevels) SetConfidenceLevel(v string) {
	o.ConfidenceLevel = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AtcfwListSeverityLevels) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtcfwListSeverityLevels) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AtcfwListSeverityLevels) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AtcfwListSeverityLevels) SetId(v int32) {
	o.Id = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AtcfwListSeverityLevels) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtcfwListSeverityLevels) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AtcfwListSeverityLevels) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *AtcfwListSeverityLevels) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetThreatLevel returns the ThreatLevel field value if set, zero value otherwise.
func (o *AtcfwListSeverityLevels) GetThreatLevel() string {
	if o == nil || IsNil(o.ThreatLevel) {
		var ret string
		return ret
	}
	return *o.ThreatLevel
}

// GetThreatLevelOk returns a tuple with the ThreatLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtcfwListSeverityLevels) GetThreatLevelOk() (*string, bool) {
	if o == nil || IsNil(o.ThreatLevel) {
		return nil, false
	}
	return o.ThreatLevel, true
}

// HasThreatLevel returns a boolean if a field has been set.
func (o *AtcfwListSeverityLevels) HasThreatLevel() bool {
	if o != nil && !IsNil(o.ThreatLevel) {
		return true
	}

	return false
}

// SetThreatLevel gets a reference to the given string and assigns it to the ThreatLevel field.
func (o *AtcfwListSeverityLevels) SetThreatLevel(v string) {
	o.ThreatLevel = &v
}

func (o AtcfwListSeverityLevels) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AtcfwListSeverityLevels) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConfidenceLevel) {
		toSerialize["confidence_level"] = o.ConfidenceLevel
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.ThreatLevel) {
		toSerialize["threat_level"] = o.ThreatLevel
	}
	return toSerialize, nil
}

type NullableAtcfwListSeverityLevels struct {
	value *AtcfwListSeverityLevels
	isSet bool
}

func (v NullableAtcfwListSeverityLevels) Get() *AtcfwListSeverityLevels {
	return v.value
}

func (v *NullableAtcfwListSeverityLevels) Set(val *AtcfwListSeverityLevels) {
	v.value = val
	v.isSet = true
}

func (v NullableAtcfwListSeverityLevels) IsSet() bool {
	return v.isSet
}

func (v *NullableAtcfwListSeverityLevels) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAtcfwListSeverityLevels(val *AtcfwListSeverityLevels) *NullableAtcfwListSeverityLevels {
	return &NullableAtcfwListSeverityLevels{value: val, isSet: true}
}

func (v NullableAtcfwListSeverityLevels) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAtcfwListSeverityLevels) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
