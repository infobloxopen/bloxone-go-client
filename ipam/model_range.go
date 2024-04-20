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
	"time"
)

// checks if the Range type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Range{}

// Range A __Range__ object (_ipam/range_) represents a set of contiguous IP addresses in the same IP space with no gap, expressed as a (start, end) pair within a given subnet that are grouped together for administrative purpose and protocol management. The start and end values are not required to align with CIDR boundaries.
type Range struct {
	// The description for the range. May contain 0 to 1024 characters. Can include UTF-8.
	Comment *string `json:"comment,omitempty"`
	// Time when the object has been created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The resource identifier.
	DhcpHost *string `json:"dhcp_host,omitempty"`
	// The list of DHCP options. May be either a specific option or a group of options.
	DhcpOptions []OptionItem `json:"dhcp_options,omitempty"`
	// Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.  Defaults to _false_.
	DisableDhcp *bool `json:"disable_dhcp,omitempty"`
	// The end IP address of the range.
	End string `json:"end"`
	// The list of all exclusion ranges in the scope of the range.
	ExclusionRanges []ExclusionRange `json:"exclusion_ranges,omitempty"`
	// The list of all allow/deny filters of the range.
	Filters []AccessFilter `json:"filters,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// The list of the inheritance assigned hosts of the object.
	InheritanceAssignedHosts []InheritanceAssignedHost `json:"inheritance_assigned_hosts,omitempty"`
	// The resource identifier.
	InheritanceParent  *string                 `json:"inheritance_parent,omitempty"`
	InheritanceSources *DHCPOptionsInheritance `json:"inheritance_sources,omitempty"`
	// The name of the range. May contain 1 to 256 characters. Can include UTF-8.
	Name *string `json:"name,omitempty"`
	// The resource identifier.
	Parent *string `json:"parent,omitempty"`
	// The type of protocol (_ip4_ or _ip6_).
	Protocol *string `json:"protocol,omitempty"`
	// The resource identifier.
	Space *string `json:"space,omitempty"`
	// The start IP address of the range.
	Start string `json:"start"`
	// The tags for the range in JSON format.
	Tags      map[string]interface{} `json:"tags,omitempty"`
	Threshold *UtilizationThreshold  `json:"threshold,omitempty"`
	// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
	UpdatedAt     *time.Time     `json:"updated_at,omitempty"`
	Utilization   *Utilization   `json:"utilization,omitempty"`
	UtilizationV6 *UtilizationV6 `json:"utilization_v6,omitempty"`
}

type _Range Range

// NewRange instantiates a new Range object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRange(end string, start string) *Range {
	this := Range{}
	this.End = end
	this.Start = start
	return &this
}

// NewRangeWithDefaults instantiates a new Range object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRangeWithDefaults() *Range {
	this := Range{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *Range) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *Range) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *Range) SetComment(v string) {
	o.Comment = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Range) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Range) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Range) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDhcpHost returns the DhcpHost field value if set, zero value otherwise.
func (o *Range) GetDhcpHost() string {
	if o == nil || IsNil(o.DhcpHost) {
		var ret string
		return ret
	}
	return *o.DhcpHost
}

// GetDhcpHostOk returns a tuple with the DhcpHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetDhcpHostOk() (*string, bool) {
	if o == nil || IsNil(o.DhcpHost) {
		return nil, false
	}
	return o.DhcpHost, true
}

// HasDhcpHost returns a boolean if a field has been set.
func (o *Range) HasDhcpHost() bool {
	if o != nil && !IsNil(o.DhcpHost) {
		return true
	}

	return false
}

// SetDhcpHost gets a reference to the given string and assigns it to the DhcpHost field.
func (o *Range) SetDhcpHost(v string) {
	o.DhcpHost = &v
}

// GetDhcpOptions returns the DhcpOptions field value if set, zero value otherwise.
func (o *Range) GetDhcpOptions() []OptionItem {
	if o == nil || IsNil(o.DhcpOptions) {
		var ret []OptionItem
		return ret
	}
	return o.DhcpOptions
}

// GetDhcpOptionsOk returns a tuple with the DhcpOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetDhcpOptionsOk() ([]OptionItem, bool) {
	if o == nil || IsNil(o.DhcpOptions) {
		return nil, false
	}
	return o.DhcpOptions, true
}

// HasDhcpOptions returns a boolean if a field has been set.
func (o *Range) HasDhcpOptions() bool {
	if o != nil && !IsNil(o.DhcpOptions) {
		return true
	}

	return false
}

// SetDhcpOptions gets a reference to the given []OptionItem and assigns it to the DhcpOptions field.
func (o *Range) SetDhcpOptions(v []OptionItem) {
	o.DhcpOptions = v
}

// GetDisableDhcp returns the DisableDhcp field value if set, zero value otherwise.
func (o *Range) GetDisableDhcp() bool {
	if o == nil || IsNil(o.DisableDhcp) {
		var ret bool
		return ret
	}
	return *o.DisableDhcp
}

// GetDisableDhcpOk returns a tuple with the DisableDhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetDisableDhcpOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableDhcp) {
		return nil, false
	}
	return o.DisableDhcp, true
}

// HasDisableDhcp returns a boolean if a field has been set.
func (o *Range) HasDisableDhcp() bool {
	if o != nil && !IsNil(o.DisableDhcp) {
		return true
	}

	return false
}

// SetDisableDhcp gets a reference to the given bool and assigns it to the DisableDhcp field.
func (o *Range) SetDisableDhcp(v bool) {
	o.DisableDhcp = &v
}

// GetEnd returns the End field value
func (o *Range) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *Range) GetEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *Range) SetEnd(v string) {
	o.End = v
}

// GetExclusionRanges returns the ExclusionRanges field value if set, zero value otherwise.
func (o *Range) GetExclusionRanges() []ExclusionRange {
	if o == nil || IsNil(o.ExclusionRanges) {
		var ret []ExclusionRange
		return ret
	}
	return o.ExclusionRanges
}

// GetExclusionRangesOk returns a tuple with the ExclusionRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetExclusionRangesOk() ([]ExclusionRange, bool) {
	if o == nil || IsNil(o.ExclusionRanges) {
		return nil, false
	}
	return o.ExclusionRanges, true
}

// HasExclusionRanges returns a boolean if a field has been set.
func (o *Range) HasExclusionRanges() bool {
	if o != nil && !IsNil(o.ExclusionRanges) {
		return true
	}

	return false
}

// SetExclusionRanges gets a reference to the given []ExclusionRange and assigns it to the ExclusionRanges field.
func (o *Range) SetExclusionRanges(v []ExclusionRange) {
	o.ExclusionRanges = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *Range) GetFilters() []AccessFilter {
	if o == nil || IsNil(o.Filters) {
		var ret []AccessFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetFiltersOk() ([]AccessFilter, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *Range) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []AccessFilter and assigns it to the Filters field.
func (o *Range) SetFilters(v []AccessFilter) {
	o.Filters = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Range) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Range) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Range) SetId(v string) {
	o.Id = &v
}

// GetInheritanceAssignedHosts returns the InheritanceAssignedHosts field value if set, zero value otherwise.
func (o *Range) GetInheritanceAssignedHosts() []InheritanceAssignedHost {
	if o == nil || IsNil(o.InheritanceAssignedHosts) {
		var ret []InheritanceAssignedHost
		return ret
	}
	return o.InheritanceAssignedHosts
}

// GetInheritanceAssignedHostsOk returns a tuple with the InheritanceAssignedHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetInheritanceAssignedHostsOk() ([]InheritanceAssignedHost, bool) {
	if o == nil || IsNil(o.InheritanceAssignedHosts) {
		return nil, false
	}
	return o.InheritanceAssignedHosts, true
}

// HasInheritanceAssignedHosts returns a boolean if a field has been set.
func (o *Range) HasInheritanceAssignedHosts() bool {
	if o != nil && !IsNil(o.InheritanceAssignedHosts) {
		return true
	}

	return false
}

// SetInheritanceAssignedHosts gets a reference to the given []InheritanceAssignedHost and assigns it to the InheritanceAssignedHosts field.
func (o *Range) SetInheritanceAssignedHosts(v []InheritanceAssignedHost) {
	o.InheritanceAssignedHosts = v
}

// GetInheritanceParent returns the InheritanceParent field value if set, zero value otherwise.
func (o *Range) GetInheritanceParent() string {
	if o == nil || IsNil(o.InheritanceParent) {
		var ret string
		return ret
	}
	return *o.InheritanceParent
}

// GetInheritanceParentOk returns a tuple with the InheritanceParent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetInheritanceParentOk() (*string, bool) {
	if o == nil || IsNil(o.InheritanceParent) {
		return nil, false
	}
	return o.InheritanceParent, true
}

// HasInheritanceParent returns a boolean if a field has been set.
func (o *Range) HasInheritanceParent() bool {
	if o != nil && !IsNil(o.InheritanceParent) {
		return true
	}

	return false
}

// SetInheritanceParent gets a reference to the given string and assigns it to the InheritanceParent field.
func (o *Range) SetInheritanceParent(v string) {
	o.InheritanceParent = &v
}

// GetInheritanceSources returns the InheritanceSources field value if set, zero value otherwise.
func (o *Range) GetInheritanceSources() DHCPOptionsInheritance {
	if o == nil || IsNil(o.InheritanceSources) {
		var ret DHCPOptionsInheritance
		return ret
	}
	return *o.InheritanceSources
}

// GetInheritanceSourcesOk returns a tuple with the InheritanceSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetInheritanceSourcesOk() (*DHCPOptionsInheritance, bool) {
	if o == nil || IsNil(o.InheritanceSources) {
		return nil, false
	}
	return o.InheritanceSources, true
}

// HasInheritanceSources returns a boolean if a field has been set.
func (o *Range) HasInheritanceSources() bool {
	if o != nil && !IsNil(o.InheritanceSources) {
		return true
	}

	return false
}

// SetInheritanceSources gets a reference to the given DHCPOptionsInheritance and assigns it to the InheritanceSources field.
func (o *Range) SetInheritanceSources(v DHCPOptionsInheritance) {
	o.InheritanceSources = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Range) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Range) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Range) SetName(v string) {
	o.Name = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *Range) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *Range) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *Range) SetParent(v string) {
	o.Parent = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *Range) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *Range) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *Range) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSpace returns the Space field value if set, zero value otherwise.
func (o *Range) GetSpace() string {
	if o == nil || IsNil(o.Space) {
		var ret string
		return ret
	}
	return *o.Space
}

// GetSpaceOk returns a tuple with the Space field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetSpaceOk() (*string, bool) {
	if o == nil || IsNil(o.Space) {
		return nil, false
	}
	return o.Space, true
}

// HasSpace returns a boolean if a field has been set.
func (o *Range) HasSpace() bool {
	if o != nil && !IsNil(o.Space) {
		return true
	}

	return false
}

// SetSpace gets a reference to the given string and assigns it to the Space field.
func (o *Range) SetSpace(v string) {
	o.Space = &v
}

// GetStart returns the Start field value
func (o *Range) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *Range) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *Range) SetStart(v string) {
	o.Start = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Range) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Range) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *Range) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *Range) GetThreshold() UtilizationThreshold {
	if o == nil || IsNil(o.Threshold) {
		var ret UtilizationThreshold
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetThresholdOk() (*UtilizationThreshold, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *Range) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given UtilizationThreshold and assigns it to the Threshold field.
func (o *Range) SetThreshold(v UtilizationThreshold) {
	o.Threshold = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Range) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Range) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Range) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUtilization returns the Utilization field value if set, zero value otherwise.
func (o *Range) GetUtilization() Utilization {
	if o == nil || IsNil(o.Utilization) {
		var ret Utilization
		return ret
	}
	return *o.Utilization
}

// GetUtilizationOk returns a tuple with the Utilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetUtilizationOk() (*Utilization, bool) {
	if o == nil || IsNil(o.Utilization) {
		return nil, false
	}
	return o.Utilization, true
}

// HasUtilization returns a boolean if a field has been set.
func (o *Range) HasUtilization() bool {
	if o != nil && !IsNil(o.Utilization) {
		return true
	}

	return false
}

// SetUtilization gets a reference to the given Utilization and assigns it to the Utilization field.
func (o *Range) SetUtilization(v Utilization) {
	o.Utilization = &v
}

// GetUtilizationV6 returns the UtilizationV6 field value if set, zero value otherwise.
func (o *Range) GetUtilizationV6() UtilizationV6 {
	if o == nil || IsNil(o.UtilizationV6) {
		var ret UtilizationV6
		return ret
	}
	return *o.UtilizationV6
}

// GetUtilizationV6Ok returns a tuple with the UtilizationV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetUtilizationV6Ok() (*UtilizationV6, bool) {
	if o == nil || IsNil(o.UtilizationV6) {
		return nil, false
	}
	return o.UtilizationV6, true
}

// HasUtilizationV6 returns a boolean if a field has been set.
func (o *Range) HasUtilizationV6() bool {
	if o != nil && !IsNil(o.UtilizationV6) {
		return true
	}

	return false
}

// SetUtilizationV6 gets a reference to the given UtilizationV6 and assigns it to the UtilizationV6 field.
func (o *Range) SetUtilizationV6(v UtilizationV6) {
	o.UtilizationV6 = &v
}

func (o Range) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Range) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DhcpHost) {
		toSerialize["dhcp_host"] = o.DhcpHost
	}
	if !IsNil(o.DhcpOptions) {
		toSerialize["dhcp_options"] = o.DhcpOptions
	}
	if !IsNil(o.DisableDhcp) {
		toSerialize["disable_dhcp"] = o.DisableDhcp
	}
	toSerialize["end"] = o.End
	if !IsNil(o.ExclusionRanges) {
		toSerialize["exclusion_ranges"] = o.ExclusionRanges
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InheritanceAssignedHosts) {
		toSerialize["inheritance_assigned_hosts"] = o.InheritanceAssignedHosts
	}
	if !IsNil(o.InheritanceParent) {
		toSerialize["inheritance_parent"] = o.InheritanceParent
	}
	if !IsNil(o.InheritanceSources) {
		toSerialize["inheritance_sources"] = o.InheritanceSources
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Space) {
		toSerialize["space"] = o.Space
	}
	toSerialize["start"] = o.Start
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Utilization) {
		toSerialize["utilization"] = o.Utilization
	}
	if !IsNil(o.UtilizationV6) {
		toSerialize["utilization_v6"] = o.UtilizationV6
	}
	return toSerialize, nil
}

func (o *Range) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"end",
		"start",
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

	varRange := _Range{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRange)

	if err != nil {
		return err
	}

	*o = Range(varRange)

	return err
}

type NullableRange struct {
	value *Range
	isSet bool
}

func (v NullableRange) Get() *Range {
	return v.value
}

func (v *NullableRange) Set(val *Range) {
	v.value = val
	v.isSet = true
}

func (v NullableRange) IsSet() bool {
	return v.isSet
}

func (v *NullableRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRange(val *Range) *NullableRange {
	return &NullableRange{value: val, isSet: true}
}

func (v NullableRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
