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

// checks if the AtcfwApplicationFilterMultiResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AtcfwApplicationFilterMultiResponse{}

// AtcfwApplicationFilterMultiResponse The Application Filter list response.
type AtcfwApplicationFilterMultiResponse struct {
	// The list of Application Filter objects.
	Results []AtcfwApplicationFilter `json:"results,omitempty"`
}

// NewAtcfwApplicationFilterMultiResponse instantiates a new AtcfwApplicationFilterMultiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAtcfwApplicationFilterMultiResponse() *AtcfwApplicationFilterMultiResponse {
	this := AtcfwApplicationFilterMultiResponse{}
	return &this
}

// NewAtcfwApplicationFilterMultiResponseWithDefaults instantiates a new AtcfwApplicationFilterMultiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAtcfwApplicationFilterMultiResponseWithDefaults() *AtcfwApplicationFilterMultiResponse {
	this := AtcfwApplicationFilterMultiResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *AtcfwApplicationFilterMultiResponse) GetResults() []AtcfwApplicationFilter {
	if o == nil || IsNil(o.Results) {
		var ret []AtcfwApplicationFilter
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtcfwApplicationFilterMultiResponse) GetResultsOk() ([]AtcfwApplicationFilter, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *AtcfwApplicationFilterMultiResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []AtcfwApplicationFilter and assigns it to the Results field.
func (o *AtcfwApplicationFilterMultiResponse) SetResults(v []AtcfwApplicationFilter) {
	o.Results = v
}

func (o AtcfwApplicationFilterMultiResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AtcfwApplicationFilterMultiResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableAtcfwApplicationFilterMultiResponse struct {
	value *AtcfwApplicationFilterMultiResponse
	isSet bool
}

func (v NullableAtcfwApplicationFilterMultiResponse) Get() *AtcfwApplicationFilterMultiResponse {
	return v.value
}

func (v *NullableAtcfwApplicationFilterMultiResponse) Set(val *AtcfwApplicationFilterMultiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAtcfwApplicationFilterMultiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAtcfwApplicationFilterMultiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAtcfwApplicationFilterMultiResponse(val *AtcfwApplicationFilterMultiResponse) *NullableAtcfwApplicationFilterMultiResponse {
	return &NullableAtcfwApplicationFilterMultiResponse{value: val, isSet: true}
}

func (v NullableAtcfwApplicationFilterMultiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAtcfwApplicationFilterMultiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
