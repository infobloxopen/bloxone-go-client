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

// checks if the AtcfwAppApprovalsUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AtcfwAppApprovalsUpdateRequest{}

// AtcfwAppApprovalsUpdateRequest struct for AtcfwAppApprovalsUpdateRequest
type AtcfwAppApprovalsUpdateRequest struct {
	Fields            *ProtobufFieldMask               `json:"fields,omitempty"`
	InsertedApprovals []AtcfwAppApproval               `json:"inserted_approvals,omitempty"`
	RemovedApprovals  []AtcfwAppApprovalRemovalRequest `json:"removed_approvals,omitempty"`
}

// NewAtcfwAppApprovalsUpdateRequest instantiates a new AtcfwAppApprovalsUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAtcfwAppApprovalsUpdateRequest() *AtcfwAppApprovalsUpdateRequest {
	this := AtcfwAppApprovalsUpdateRequest{}
	return &this
}

// NewAtcfwAppApprovalsUpdateRequestWithDefaults instantiates a new AtcfwAppApprovalsUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAtcfwAppApprovalsUpdateRequestWithDefaults() *AtcfwAppApprovalsUpdateRequest {
	this := AtcfwAppApprovalsUpdateRequest{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *AtcfwAppApprovalsUpdateRequest) GetFields() ProtobufFieldMask {
	if o == nil || IsNil(o.Fields) {
		var ret ProtobufFieldMask
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtcfwAppApprovalsUpdateRequest) GetFieldsOk() (*ProtobufFieldMask, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *AtcfwAppApprovalsUpdateRequest) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given ProtobufFieldMask and assigns it to the Fields field.
func (o *AtcfwAppApprovalsUpdateRequest) SetFields(v ProtobufFieldMask) {
	o.Fields = &v
}

// GetInsertedApprovals returns the InsertedApprovals field value if set, zero value otherwise.
func (o *AtcfwAppApprovalsUpdateRequest) GetInsertedApprovals() []AtcfwAppApproval {
	if o == nil || IsNil(o.InsertedApprovals) {
		var ret []AtcfwAppApproval
		return ret
	}
	return o.InsertedApprovals
}

// GetInsertedApprovalsOk returns a tuple with the InsertedApprovals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtcfwAppApprovalsUpdateRequest) GetInsertedApprovalsOk() ([]AtcfwAppApproval, bool) {
	if o == nil || IsNil(o.InsertedApprovals) {
		return nil, false
	}
	return o.InsertedApprovals, true
}

// HasInsertedApprovals returns a boolean if a field has been set.
func (o *AtcfwAppApprovalsUpdateRequest) HasInsertedApprovals() bool {
	if o != nil && !IsNil(o.InsertedApprovals) {
		return true
	}

	return false
}

// SetInsertedApprovals gets a reference to the given []AtcfwAppApproval and assigns it to the InsertedApprovals field.
func (o *AtcfwAppApprovalsUpdateRequest) SetInsertedApprovals(v []AtcfwAppApproval) {
	o.InsertedApprovals = v
}

// GetRemovedApprovals returns the RemovedApprovals field value if set, zero value otherwise.
func (o *AtcfwAppApprovalsUpdateRequest) GetRemovedApprovals() []AtcfwAppApprovalRemovalRequest {
	if o == nil || IsNil(o.RemovedApprovals) {
		var ret []AtcfwAppApprovalRemovalRequest
		return ret
	}
	return o.RemovedApprovals
}

// GetRemovedApprovalsOk returns a tuple with the RemovedApprovals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtcfwAppApprovalsUpdateRequest) GetRemovedApprovalsOk() ([]AtcfwAppApprovalRemovalRequest, bool) {
	if o == nil || IsNil(o.RemovedApprovals) {
		return nil, false
	}
	return o.RemovedApprovals, true
}

// HasRemovedApprovals returns a boolean if a field has been set.
func (o *AtcfwAppApprovalsUpdateRequest) HasRemovedApprovals() bool {
	if o != nil && !IsNil(o.RemovedApprovals) {
		return true
	}

	return false
}

// SetRemovedApprovals gets a reference to the given []AtcfwAppApprovalRemovalRequest and assigns it to the RemovedApprovals field.
func (o *AtcfwAppApprovalsUpdateRequest) SetRemovedApprovals(v []AtcfwAppApprovalRemovalRequest) {
	o.RemovedApprovals = v
}

func (o AtcfwAppApprovalsUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AtcfwAppApprovalsUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.InsertedApprovals) {
		toSerialize["inserted_approvals"] = o.InsertedApprovals
	}
	if !IsNil(o.RemovedApprovals) {
		toSerialize["removed_approvals"] = o.RemovedApprovals
	}
	return toSerialize, nil
}

type NullableAtcfwAppApprovalsUpdateRequest struct {
	value *AtcfwAppApprovalsUpdateRequest
	isSet bool
}

func (v NullableAtcfwAppApprovalsUpdateRequest) Get() *AtcfwAppApprovalsUpdateRequest {
	return v.value
}

func (v *NullableAtcfwAppApprovalsUpdateRequest) Set(val *AtcfwAppApprovalsUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAtcfwAppApprovalsUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAtcfwAppApprovalsUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAtcfwAppApprovalsUpdateRequest(val *AtcfwAppApprovalsUpdateRequest) *NullableAtcfwAppApprovalsUpdateRequest {
	return &NullableAtcfwAppApprovalsUpdateRequest{value: val, isSet: true}
}

func (v NullableAtcfwAppApprovalsUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAtcfwAppApprovalsUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
