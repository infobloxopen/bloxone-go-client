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

// checks if the IpamsvcASM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcASM{}

// IpamsvcASM The __ASM__ object is a synthetic object representing the suggestions from the Automated Scope Management suggestion engine for expanding subnet or range.
type IpamsvcASM struct {
	// The end IP address when adding to the back.
	BackEnd *string `json:"back_end,omitempty"`
	// The start IP address when adding to the back.
	BackStart *string `json:"back_start,omitempty"`
	// The end IP address when adding to the back.
	BothEnd *string `json:"both_end,omitempty"`
	// The start IP address when adding to both front and back.
	BothStart *string `json:"both_start,omitempty"`
	// The end IP address when adding to the front.
	FrontEnd *string `json:"front_end,omitempty"`
	// The start IP address when adding to the front.
	FrontStart *string `json:"front_start,omitempty"`
	// Calculated number of addresses to grow range; its value is determined by asm_config growth factor, type, and min_unused after the expansion.
	Growth *int32 `json:"growth,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// Either the value from the ASM configuration or -1 if the estimate is that utilization will not exceed the threshold.
	Lookahead *int32 `json:"lookahead,omitempty"`
	// The end IP address of the range.
	RangeEnd *string `json:"range_end,omitempty"`
	// The resource identifier.
	RangeId *string `json:"range_id,omitempty"`
	// The start IP address of the range.
	RangeStart *string `json:"range_start,omitempty"`
	// The suggested subnet expansion.
	SubnetAddress *string `json:"subnet_address,omitempty"`
	// The CIDR of the subnet.
	SubnetCidr *int64 `json:"subnet_cidr,omitempty"`
	// Indicates where the subnet may expand. As the subnet can only be expanded by one bit at a time, it can only grow in one of the two directions. It is set to _none_ if the subnet can't be expanded.  Valid values are: * _front_ * _back_ * _none_
	SubnetDirection *string `json:"subnet_direction,omitempty"`
	// The resource identifier.
	SubnetId string `json:"subnet_id"`
	// The resource identifier.
	SubnetRange *string `json:"subnet_range,omitempty"`
	// The suggested new range end in expanded subnet.
	SubnetRangeEnd *string `json:"subnet_range_end,omitempty"`
	// The suggested new range start in expanded subnet.
	SubnetRangeStart *string `json:"subnet_range_start,omitempty"`
	// Indicates if future notifications for this subnet should be suppressed.  Valid values are: * _no_ * _time_ * _permanent_  If set to _permanent_ notifications are suppressed until the administrator modifies the configuration for the subnet. If set to _time_ notifications are suppressed until the specified time. Defaults to _no_.
	Suppress *string `json:"suppress,omitempty"`
	// The time duration in days to suppress the notifications for this subnet.
	SuppressTime *int64 `json:"suppress_time,omitempty"`
	// The utilization threshold as per ASM configuration.
	ThresholdUtilization *int64 `json:"threshold_utilization,omitempty"`
	// The object to update.  Valid values are: * _range_ * _subnet_ * _none_
	Update *string `json:"update,omitempty"`
	// The utilization of DHCP addresses in the subnet.
	Utilization *int64 `json:"utilization,omitempty"`
}

type _IpamsvcASM IpamsvcASM

// NewIpamsvcASM instantiates a new IpamsvcASM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcASM(subnetId string) *IpamsvcASM {
	this := IpamsvcASM{}
	this.SubnetId = subnetId
	return &this
}

// NewIpamsvcASMWithDefaults instantiates a new IpamsvcASM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcASMWithDefaults() *IpamsvcASM {
	this := IpamsvcASM{}
	return &this
}

// GetBackEnd returns the BackEnd field value if set, zero value otherwise.
func (o *IpamsvcASM) GetBackEnd() string {
	if o == nil || IsNil(o.BackEnd) {
		var ret string
		return ret
	}
	return *o.BackEnd
}

// GetBackEndOk returns a tuple with the BackEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetBackEndOk() (*string, bool) {
	if o == nil || IsNil(o.BackEnd) {
		return nil, false
	}
	return o.BackEnd, true
}

// HasBackEnd returns a boolean if a field has been set.
func (o *IpamsvcASM) HasBackEnd() bool {
	if o != nil && !IsNil(o.BackEnd) {
		return true
	}

	return false
}

// SetBackEnd gets a reference to the given string and assigns it to the BackEnd field.
func (o *IpamsvcASM) SetBackEnd(v string) {
	o.BackEnd = &v
}

// GetBackStart returns the BackStart field value if set, zero value otherwise.
func (o *IpamsvcASM) GetBackStart() string {
	if o == nil || IsNil(o.BackStart) {
		var ret string
		return ret
	}
	return *o.BackStart
}

// GetBackStartOk returns a tuple with the BackStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetBackStartOk() (*string, bool) {
	if o == nil || IsNil(o.BackStart) {
		return nil, false
	}
	return o.BackStart, true
}

// HasBackStart returns a boolean if a field has been set.
func (o *IpamsvcASM) HasBackStart() bool {
	if o != nil && !IsNil(o.BackStart) {
		return true
	}

	return false
}

// SetBackStart gets a reference to the given string and assigns it to the BackStart field.
func (o *IpamsvcASM) SetBackStart(v string) {
	o.BackStart = &v
}

// GetBothEnd returns the BothEnd field value if set, zero value otherwise.
func (o *IpamsvcASM) GetBothEnd() string {
	if o == nil || IsNil(o.BothEnd) {
		var ret string
		return ret
	}
	return *o.BothEnd
}

// GetBothEndOk returns a tuple with the BothEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetBothEndOk() (*string, bool) {
	if o == nil || IsNil(o.BothEnd) {
		return nil, false
	}
	return o.BothEnd, true
}

// HasBothEnd returns a boolean if a field has been set.
func (o *IpamsvcASM) HasBothEnd() bool {
	if o != nil && !IsNil(o.BothEnd) {
		return true
	}

	return false
}

// SetBothEnd gets a reference to the given string and assigns it to the BothEnd field.
func (o *IpamsvcASM) SetBothEnd(v string) {
	o.BothEnd = &v
}

// GetBothStart returns the BothStart field value if set, zero value otherwise.
func (o *IpamsvcASM) GetBothStart() string {
	if o == nil || IsNil(o.BothStart) {
		var ret string
		return ret
	}
	return *o.BothStart
}

// GetBothStartOk returns a tuple with the BothStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetBothStartOk() (*string, bool) {
	if o == nil || IsNil(o.BothStart) {
		return nil, false
	}
	return o.BothStart, true
}

// HasBothStart returns a boolean if a field has been set.
func (o *IpamsvcASM) HasBothStart() bool {
	if o != nil && !IsNil(o.BothStart) {
		return true
	}

	return false
}

// SetBothStart gets a reference to the given string and assigns it to the BothStart field.
func (o *IpamsvcASM) SetBothStart(v string) {
	o.BothStart = &v
}

// GetFrontEnd returns the FrontEnd field value if set, zero value otherwise.
func (o *IpamsvcASM) GetFrontEnd() string {
	if o == nil || IsNil(o.FrontEnd) {
		var ret string
		return ret
	}
	return *o.FrontEnd
}

// GetFrontEndOk returns a tuple with the FrontEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetFrontEndOk() (*string, bool) {
	if o == nil || IsNil(o.FrontEnd) {
		return nil, false
	}
	return o.FrontEnd, true
}

// HasFrontEnd returns a boolean if a field has been set.
func (o *IpamsvcASM) HasFrontEnd() bool {
	if o != nil && !IsNil(o.FrontEnd) {
		return true
	}

	return false
}

// SetFrontEnd gets a reference to the given string and assigns it to the FrontEnd field.
func (o *IpamsvcASM) SetFrontEnd(v string) {
	o.FrontEnd = &v
}

// GetFrontStart returns the FrontStart field value if set, zero value otherwise.
func (o *IpamsvcASM) GetFrontStart() string {
	if o == nil || IsNil(o.FrontStart) {
		var ret string
		return ret
	}
	return *o.FrontStart
}

// GetFrontStartOk returns a tuple with the FrontStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetFrontStartOk() (*string, bool) {
	if o == nil || IsNil(o.FrontStart) {
		return nil, false
	}
	return o.FrontStart, true
}

// HasFrontStart returns a boolean if a field has been set.
func (o *IpamsvcASM) HasFrontStart() bool {
	if o != nil && !IsNil(o.FrontStart) {
		return true
	}

	return false
}

// SetFrontStart gets a reference to the given string and assigns it to the FrontStart field.
func (o *IpamsvcASM) SetFrontStart(v string) {
	o.FrontStart = &v
}

// GetGrowth returns the Growth field value if set, zero value otherwise.
func (o *IpamsvcASM) GetGrowth() int32 {
	if o == nil || IsNil(o.Growth) {
		var ret int32
		return ret
	}
	return *o.Growth
}

// GetGrowthOk returns a tuple with the Growth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetGrowthOk() (*int32, bool) {
	if o == nil || IsNil(o.Growth) {
		return nil, false
	}
	return o.Growth, true
}

// HasGrowth returns a boolean if a field has been set.
func (o *IpamsvcASM) HasGrowth() bool {
	if o != nil && !IsNil(o.Growth) {
		return true
	}

	return false
}

// SetGrowth gets a reference to the given int32 and assigns it to the Growth field.
func (o *IpamsvcASM) SetGrowth(v int32) {
	o.Growth = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IpamsvcASM) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IpamsvcASM) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IpamsvcASM) SetId(v string) {
	o.Id = &v
}

// GetLookahead returns the Lookahead field value if set, zero value otherwise.
func (o *IpamsvcASM) GetLookahead() int32 {
	if o == nil || IsNil(o.Lookahead) {
		var ret int32
		return ret
	}
	return *o.Lookahead
}

// GetLookaheadOk returns a tuple with the Lookahead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetLookaheadOk() (*int32, bool) {
	if o == nil || IsNil(o.Lookahead) {
		return nil, false
	}
	return o.Lookahead, true
}

// HasLookahead returns a boolean if a field has been set.
func (o *IpamsvcASM) HasLookahead() bool {
	if o != nil && !IsNil(o.Lookahead) {
		return true
	}

	return false
}

// SetLookahead gets a reference to the given int32 and assigns it to the Lookahead field.
func (o *IpamsvcASM) SetLookahead(v int32) {
	o.Lookahead = &v
}

// GetRangeEnd returns the RangeEnd field value if set, zero value otherwise.
func (o *IpamsvcASM) GetRangeEnd() string {
	if o == nil || IsNil(o.RangeEnd) {
		var ret string
		return ret
	}
	return *o.RangeEnd
}

// GetRangeEndOk returns a tuple with the RangeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetRangeEndOk() (*string, bool) {
	if o == nil || IsNil(o.RangeEnd) {
		return nil, false
	}
	return o.RangeEnd, true
}

// HasRangeEnd returns a boolean if a field has been set.
func (o *IpamsvcASM) HasRangeEnd() bool {
	if o != nil && !IsNil(o.RangeEnd) {
		return true
	}

	return false
}

// SetRangeEnd gets a reference to the given string and assigns it to the RangeEnd field.
func (o *IpamsvcASM) SetRangeEnd(v string) {
	o.RangeEnd = &v
}

// GetRangeId returns the RangeId field value if set, zero value otherwise.
func (o *IpamsvcASM) GetRangeId() string {
	if o == nil || IsNil(o.RangeId) {
		var ret string
		return ret
	}
	return *o.RangeId
}

// GetRangeIdOk returns a tuple with the RangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetRangeIdOk() (*string, bool) {
	if o == nil || IsNil(o.RangeId) {
		return nil, false
	}
	return o.RangeId, true
}

// HasRangeId returns a boolean if a field has been set.
func (o *IpamsvcASM) HasRangeId() bool {
	if o != nil && !IsNil(o.RangeId) {
		return true
	}

	return false
}

// SetRangeId gets a reference to the given string and assigns it to the RangeId field.
func (o *IpamsvcASM) SetRangeId(v string) {
	o.RangeId = &v
}

// GetRangeStart returns the RangeStart field value if set, zero value otherwise.
func (o *IpamsvcASM) GetRangeStart() string {
	if o == nil || IsNil(o.RangeStart) {
		var ret string
		return ret
	}
	return *o.RangeStart
}

// GetRangeStartOk returns a tuple with the RangeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetRangeStartOk() (*string, bool) {
	if o == nil || IsNil(o.RangeStart) {
		return nil, false
	}
	return o.RangeStart, true
}

// HasRangeStart returns a boolean if a field has been set.
func (o *IpamsvcASM) HasRangeStart() bool {
	if o != nil && !IsNil(o.RangeStart) {
		return true
	}

	return false
}

// SetRangeStart gets a reference to the given string and assigns it to the RangeStart field.
func (o *IpamsvcASM) SetRangeStart(v string) {
	o.RangeStart = &v
}

// GetSubnetAddress returns the SubnetAddress field value if set, zero value otherwise.
func (o *IpamsvcASM) GetSubnetAddress() string {
	if o == nil || IsNil(o.SubnetAddress) {
		var ret string
		return ret
	}
	return *o.SubnetAddress
}

// GetSubnetAddressOk returns a tuple with the SubnetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetSubnetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.SubnetAddress) {
		return nil, false
	}
	return o.SubnetAddress, true
}

// HasSubnetAddress returns a boolean if a field has been set.
func (o *IpamsvcASM) HasSubnetAddress() bool {
	if o != nil && !IsNil(o.SubnetAddress) {
		return true
	}

	return false
}

// SetSubnetAddress gets a reference to the given string and assigns it to the SubnetAddress field.
func (o *IpamsvcASM) SetSubnetAddress(v string) {
	o.SubnetAddress = &v
}

// GetSubnetCidr returns the SubnetCidr field value if set, zero value otherwise.
func (o *IpamsvcASM) GetSubnetCidr() int64 {
	if o == nil || IsNil(o.SubnetCidr) {
		var ret int64
		return ret
	}
	return *o.SubnetCidr
}

// GetSubnetCidrOk returns a tuple with the SubnetCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetSubnetCidrOk() (*int64, bool) {
	if o == nil || IsNil(o.SubnetCidr) {
		return nil, false
	}
	return o.SubnetCidr, true
}

// HasSubnetCidr returns a boolean if a field has been set.
func (o *IpamsvcASM) HasSubnetCidr() bool {
	if o != nil && !IsNil(o.SubnetCidr) {
		return true
	}

	return false
}

// SetSubnetCidr gets a reference to the given int64 and assigns it to the SubnetCidr field.
func (o *IpamsvcASM) SetSubnetCidr(v int64) {
	o.SubnetCidr = &v
}

// GetSubnetDirection returns the SubnetDirection field value if set, zero value otherwise.
func (o *IpamsvcASM) GetSubnetDirection() string {
	if o == nil || IsNil(o.SubnetDirection) {
		var ret string
		return ret
	}
	return *o.SubnetDirection
}

// GetSubnetDirectionOk returns a tuple with the SubnetDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetSubnetDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.SubnetDirection) {
		return nil, false
	}
	return o.SubnetDirection, true
}

// HasSubnetDirection returns a boolean if a field has been set.
func (o *IpamsvcASM) HasSubnetDirection() bool {
	if o != nil && !IsNil(o.SubnetDirection) {
		return true
	}

	return false
}

// SetSubnetDirection gets a reference to the given string and assigns it to the SubnetDirection field.
func (o *IpamsvcASM) SetSubnetDirection(v string) {
	o.SubnetDirection = &v
}

// GetSubnetId returns the SubnetId field value
func (o *IpamsvcASM) GetSubnetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubnetId
}

// GetSubnetIdOk returns a tuple with the SubnetId field value
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetSubnetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubnetId, true
}

// SetSubnetId sets field value
func (o *IpamsvcASM) SetSubnetId(v string) {
	o.SubnetId = v
}

// GetSubnetRange returns the SubnetRange field value if set, zero value otherwise.
func (o *IpamsvcASM) GetSubnetRange() string {
	if o == nil || IsNil(o.SubnetRange) {
		var ret string
		return ret
	}
	return *o.SubnetRange
}

// GetSubnetRangeOk returns a tuple with the SubnetRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetSubnetRangeOk() (*string, bool) {
	if o == nil || IsNil(o.SubnetRange) {
		return nil, false
	}
	return o.SubnetRange, true
}

// HasSubnetRange returns a boolean if a field has been set.
func (o *IpamsvcASM) HasSubnetRange() bool {
	if o != nil && !IsNil(o.SubnetRange) {
		return true
	}

	return false
}

// SetSubnetRange gets a reference to the given string and assigns it to the SubnetRange field.
func (o *IpamsvcASM) SetSubnetRange(v string) {
	o.SubnetRange = &v
}

// GetSubnetRangeEnd returns the SubnetRangeEnd field value if set, zero value otherwise.
func (o *IpamsvcASM) GetSubnetRangeEnd() string {
	if o == nil || IsNil(o.SubnetRangeEnd) {
		var ret string
		return ret
	}
	return *o.SubnetRangeEnd
}

// GetSubnetRangeEndOk returns a tuple with the SubnetRangeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetSubnetRangeEndOk() (*string, bool) {
	if o == nil || IsNil(o.SubnetRangeEnd) {
		return nil, false
	}
	return o.SubnetRangeEnd, true
}

// HasSubnetRangeEnd returns a boolean if a field has been set.
func (o *IpamsvcASM) HasSubnetRangeEnd() bool {
	if o != nil && !IsNil(o.SubnetRangeEnd) {
		return true
	}

	return false
}

// SetSubnetRangeEnd gets a reference to the given string and assigns it to the SubnetRangeEnd field.
func (o *IpamsvcASM) SetSubnetRangeEnd(v string) {
	o.SubnetRangeEnd = &v
}

// GetSubnetRangeStart returns the SubnetRangeStart field value if set, zero value otherwise.
func (o *IpamsvcASM) GetSubnetRangeStart() string {
	if o == nil || IsNil(o.SubnetRangeStart) {
		var ret string
		return ret
	}
	return *o.SubnetRangeStart
}

// GetSubnetRangeStartOk returns a tuple with the SubnetRangeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetSubnetRangeStartOk() (*string, bool) {
	if o == nil || IsNil(o.SubnetRangeStart) {
		return nil, false
	}
	return o.SubnetRangeStart, true
}

// HasSubnetRangeStart returns a boolean if a field has been set.
func (o *IpamsvcASM) HasSubnetRangeStart() bool {
	if o != nil && !IsNil(o.SubnetRangeStart) {
		return true
	}

	return false
}

// SetSubnetRangeStart gets a reference to the given string and assigns it to the SubnetRangeStart field.
func (o *IpamsvcASM) SetSubnetRangeStart(v string) {
	o.SubnetRangeStart = &v
}

// GetSuppress returns the Suppress field value if set, zero value otherwise.
func (o *IpamsvcASM) GetSuppress() string {
	if o == nil || IsNil(o.Suppress) {
		var ret string
		return ret
	}
	return *o.Suppress
}

// GetSuppressOk returns a tuple with the Suppress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetSuppressOk() (*string, bool) {
	if o == nil || IsNil(o.Suppress) {
		return nil, false
	}
	return o.Suppress, true
}

// HasSuppress returns a boolean if a field has been set.
func (o *IpamsvcASM) HasSuppress() bool {
	if o != nil && !IsNil(o.Suppress) {
		return true
	}

	return false
}

// SetSuppress gets a reference to the given string and assigns it to the Suppress field.
func (o *IpamsvcASM) SetSuppress(v string) {
	o.Suppress = &v
}

// GetSuppressTime returns the SuppressTime field value if set, zero value otherwise.
func (o *IpamsvcASM) GetSuppressTime() int64 {
	if o == nil || IsNil(o.SuppressTime) {
		var ret int64
		return ret
	}
	return *o.SuppressTime
}

// GetSuppressTimeOk returns a tuple with the SuppressTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetSuppressTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.SuppressTime) {
		return nil, false
	}
	return o.SuppressTime, true
}

// HasSuppressTime returns a boolean if a field has been set.
func (o *IpamsvcASM) HasSuppressTime() bool {
	if o != nil && !IsNil(o.SuppressTime) {
		return true
	}

	return false
}

// SetSuppressTime gets a reference to the given int64 and assigns it to the SuppressTime field.
func (o *IpamsvcASM) SetSuppressTime(v int64) {
	o.SuppressTime = &v
}

// GetThresholdUtilization returns the ThresholdUtilization field value if set, zero value otherwise.
func (o *IpamsvcASM) GetThresholdUtilization() int64 {
	if o == nil || IsNil(o.ThresholdUtilization) {
		var ret int64
		return ret
	}
	return *o.ThresholdUtilization
}

// GetThresholdUtilizationOk returns a tuple with the ThresholdUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetThresholdUtilizationOk() (*int64, bool) {
	if o == nil || IsNil(o.ThresholdUtilization) {
		return nil, false
	}
	return o.ThresholdUtilization, true
}

// HasThresholdUtilization returns a boolean if a field has been set.
func (o *IpamsvcASM) HasThresholdUtilization() bool {
	if o != nil && !IsNil(o.ThresholdUtilization) {
		return true
	}

	return false
}

// SetThresholdUtilization gets a reference to the given int64 and assigns it to the ThresholdUtilization field.
func (o *IpamsvcASM) SetThresholdUtilization(v int64) {
	o.ThresholdUtilization = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *IpamsvcASM) GetUpdate() string {
	if o == nil || IsNil(o.Update) {
		var ret string
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetUpdateOk() (*string, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *IpamsvcASM) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given string and assigns it to the Update field.
func (o *IpamsvcASM) SetUpdate(v string) {
	o.Update = &v
}

// GetUtilization returns the Utilization field value if set, zero value otherwise.
func (o *IpamsvcASM) GetUtilization() int64 {
	if o == nil || IsNil(o.Utilization) {
		var ret int64
		return ret
	}
	return *o.Utilization
}

// GetUtilizationOk returns a tuple with the Utilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcASM) GetUtilizationOk() (*int64, bool) {
	if o == nil || IsNil(o.Utilization) {
		return nil, false
	}
	return o.Utilization, true
}

// HasUtilization returns a boolean if a field has been set.
func (o *IpamsvcASM) HasUtilization() bool {
	if o != nil && !IsNil(o.Utilization) {
		return true
	}

	return false
}

// SetUtilization gets a reference to the given int64 and assigns it to the Utilization field.
func (o *IpamsvcASM) SetUtilization(v int64) {
	o.Utilization = &v
}

func (o IpamsvcASM) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcASM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BackEnd) {
		toSerialize["back_end"] = o.BackEnd
	}
	if !IsNil(o.BackStart) {
		toSerialize["back_start"] = o.BackStart
	}
	if !IsNil(o.BothEnd) {
		toSerialize["both_end"] = o.BothEnd
	}
	if !IsNil(o.BothStart) {
		toSerialize["both_start"] = o.BothStart
	}
	if !IsNil(o.FrontEnd) {
		toSerialize["front_end"] = o.FrontEnd
	}
	if !IsNil(o.FrontStart) {
		toSerialize["front_start"] = o.FrontStart
	}
	if !IsNil(o.Growth) {
		toSerialize["growth"] = o.Growth
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Lookahead) {
		toSerialize["lookahead"] = o.Lookahead
	}
	if !IsNil(o.RangeEnd) {
		toSerialize["range_end"] = o.RangeEnd
	}
	if !IsNil(o.RangeId) {
		toSerialize["range_id"] = o.RangeId
	}
	if !IsNil(o.RangeStart) {
		toSerialize["range_start"] = o.RangeStart
	}
	if !IsNil(o.SubnetAddress) {
		toSerialize["subnet_address"] = o.SubnetAddress
	}
	if !IsNil(o.SubnetCidr) {
		toSerialize["subnet_cidr"] = o.SubnetCidr
	}
	if !IsNil(o.SubnetDirection) {
		toSerialize["subnet_direction"] = o.SubnetDirection
	}
	toSerialize["subnet_id"] = o.SubnetId
	if !IsNil(o.SubnetRange) {
		toSerialize["subnet_range"] = o.SubnetRange
	}
	if !IsNil(o.SubnetRangeEnd) {
		toSerialize["subnet_range_end"] = o.SubnetRangeEnd
	}
	if !IsNil(o.SubnetRangeStart) {
		toSerialize["subnet_range_start"] = o.SubnetRangeStart
	}
	if !IsNil(o.Suppress) {
		toSerialize["suppress"] = o.Suppress
	}
	if !IsNil(o.SuppressTime) {
		toSerialize["suppress_time"] = o.SuppressTime
	}
	if !IsNil(o.ThresholdUtilization) {
		toSerialize["threshold_utilization"] = o.ThresholdUtilization
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.Utilization) {
		toSerialize["utilization"] = o.Utilization
	}
	return toSerialize, nil
}

func (o *IpamsvcASM) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subnet_id",
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

	varIpamsvcASM := _IpamsvcASM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIpamsvcASM)

	if err != nil {
		return err
	}

	*o = IpamsvcASM(varIpamsvcASM)

	return err
}

type NullableIpamsvcASM struct {
	value *IpamsvcASM
	isSet bool
}

func (v NullableIpamsvcASM) Get() *IpamsvcASM {
	return v.value
}

func (v *NullableIpamsvcASM) Set(val *IpamsvcASM) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcASM) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcASM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcASM(val *IpamsvcASM) *NullableIpamsvcASM {
	return &NullableIpamsvcASM{value: val, isSet: true}
}

func (v NullableIpamsvcASM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcASM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
