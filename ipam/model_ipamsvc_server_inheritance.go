/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the IpamsvcServerInheritance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcServerInheritance{}

// IpamsvcServerInheritance The inheritance configuration specifies how and which fields _Server_ object (DHCP Config Profile) inherits from _Global_ parent.
type IpamsvcServerInheritance struct {
	DdnsBlock                       *IpamsvcInheritedDDNSBlock            `json:"ddns_block,omitempty"`
	DdnsClientUpdate                *InheritanceInheritedString           `json:"ddns_client_update,omitempty"`
	DdnsConflictResolutionMode      *InheritanceInheritedString           `json:"ddns_conflict_resolution_mode,omitempty"`
	DdnsHostnameBlock               *IpamsvcInheritedDDNSHostnameBlock    `json:"ddns_hostname_block,omitempty"`
	DdnsTtlPercent                  *InheritanceInheritedFloat            `json:"ddns_ttl_percent,omitempty"`
	DdnsUpdateOnRenew               *InheritanceInheritedBool             `json:"ddns_update_on_renew,omitempty"`
	DdnsUseConflictResolution       *InheritanceInheritedBool             `json:"ddns_use_conflict_resolution,omitempty"`
	DhcpConfig                      *IpamsvcInheritedDHCPConfig           `json:"dhcp_config,omitempty"`
	DhcpOptions                     *IpamsvcInheritedDHCPOptionList       `json:"dhcp_options,omitempty"`
	DhcpOptionsV6                   *IpamsvcInheritedDHCPOptionList       `json:"dhcp_options_v6,omitempty"`
	HeaderOptionFilename            *InheritanceInheritedString           `json:"header_option_filename,omitempty"`
	HeaderOptionServerAddress       *InheritanceInheritedString           `json:"header_option_server_address,omitempty"`
	HeaderOptionServerName          *InheritanceInheritedString           `json:"header_option_server_name,omitempty"`
	HostnameRewriteBlock            *IpamsvcInheritedHostnameRewriteBlock `json:"hostname_rewrite_block,omitempty"`
	VendorSpecificOptionOptionSpace *InheritanceInheritedIdentifier       `json:"vendor_specific_option_option_space,omitempty"`
}

// NewIpamsvcServerInheritance instantiates a new IpamsvcServerInheritance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcServerInheritance() *IpamsvcServerInheritance {
	this := IpamsvcServerInheritance{}
	return &this
}

// NewIpamsvcServerInheritanceWithDefaults instantiates a new IpamsvcServerInheritance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcServerInheritanceWithDefaults() *IpamsvcServerInheritance {
	this := IpamsvcServerInheritance{}
	return &this
}

// GetDdnsBlock returns the DdnsBlock field value if set, zero value otherwise.
func (o *IpamsvcServerInheritance) GetDdnsBlock() IpamsvcInheritedDDNSBlock {
	if o == nil || IsNil(o.DdnsBlock) {
		var ret IpamsvcInheritedDDNSBlock
		return ret
	}
	return *o.DdnsBlock
}

// GetDdnsBlockOk returns a tuple with the DdnsBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcServerInheritance) GetDdnsBlockOk() (*IpamsvcInheritedDDNSBlock, bool) {
	if o == nil || IsNil(o.DdnsBlock) {
		return nil, false
	}
	return o.DdnsBlock, true
}

// HasDdnsBlock returns a boolean if a field has been set.
func (o *IpamsvcServerInheritance) HasDdnsBlock() bool {
	if o != nil && !IsNil(o.DdnsBlock) {
		return true
	}

	return false
}

// SetDdnsBlock gets a reference to the given IpamsvcInheritedDDNSBlock and assigns it to the DdnsBlock field.
func (o *IpamsvcServerInheritance) SetDdnsBlock(v IpamsvcInheritedDDNSBlock) {
	o.DdnsBlock = &v
}

// GetDdnsClientUpdate returns the DdnsClientUpdate field value if set, zero value otherwise.
func (o *IpamsvcServerInheritance) GetDdnsClientUpdate() InheritanceInheritedString {
	if o == nil || IsNil(o.DdnsClientUpdate) {
		var ret InheritanceInheritedString
		return ret
	}
	return *o.DdnsClientUpdate
}

// GetDdnsClientUpdateOk returns a tuple with the DdnsClientUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcServerInheritance) GetDdnsClientUpdateOk() (*InheritanceInheritedString, bool) {
	if o == nil || IsNil(o.DdnsClientUpdate) {
		return nil, false
	}
	return o.DdnsClientUpdate, true
}

// HasDdnsClientUpdate returns a boolean if a field has been set.
func (o *IpamsvcServerInheritance) HasDdnsClientUpdate() bool {
	if o != nil && !IsNil(o.DdnsClientUpdate) {
		return true
	}

	return false
}

// SetDdnsClientUpdate gets a reference to the given InheritanceInheritedString and assigns it to the DdnsClientUpdate field.
func (o *IpamsvcServerInheritance) SetDdnsClientUpdate(v InheritanceInheritedString) {
	o.DdnsClientUpdate = &v
}

// GetDdnsConflictResolutionMode returns the DdnsConflictResolutionMode field value if set, zero value otherwise.
func (o *IpamsvcServerInheritance) GetDdnsConflictResolutionMode() InheritanceInheritedString {
	if o == nil || IsNil(o.DdnsConflictResolutionMode) {
		var ret InheritanceInheritedString
		return ret
	}
	return *o.DdnsConflictResolutionMode
}

// GetDdnsConflictResolutionModeOk returns a tuple with the DdnsConflictResolutionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcServerInheritance) GetDdnsConflictResolutionModeOk() (*InheritanceInheritedString, bool) {
	if o == nil || IsNil(o.DdnsConflictResolutionMode) {
		return nil, false
	}
	return o.DdnsConflictResolutionMode, true
}

// HasDdnsConflictResolutionMode returns a boolean if a field has been set.
func (o *IpamsvcServerInheritance) HasDdnsConflictResolutionMode() bool {
	if o != nil && !IsNil(o.DdnsConflictResolutionMode) {
		return true
	}

	return false
}

// SetDdnsConflictResolutionMode gets a reference to the given InheritanceInheritedString and assigns it to the DdnsConflictResolutionMode field.
func (o *IpamsvcServerInheritance) SetDdnsConflictResolutionMode(v InheritanceInheritedString) {
	o.DdnsConflictResolutionMode = &v
}

// GetDdnsHostnameBlock returns the DdnsHostnameBlock field value if set, zero value otherwise.
func (o *IpamsvcServerInheritance) GetDdnsHostnameBlock() IpamsvcInheritedDDNSHostnameBlock {
	if o == nil || IsNil(o.DdnsHostnameBlock) {
		var ret IpamsvcInheritedDDNSHostnameBlock
		return ret
	}
	return *o.DdnsHostnameBlock
}

// GetDdnsHostnameBlockOk returns a tuple with the DdnsHostnameBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcServerInheritance) GetDdnsHostnameBlockOk() (*IpamsvcInheritedDDNSHostnameBlock, bool) {
	if o == nil || IsNil(o.DdnsHostnameBlock) {
		return nil, false
	}
	return o.DdnsHostnameBlock, true
}

// HasDdnsHostnameBlock returns a boolean if a field has been set.
func (o *IpamsvcServerInheritance) HasDdnsHostnameBlock() bool {
	if o != nil && !IsNil(o.DdnsHostnameBlock) {
		return true
	}

	return false
}

// SetDdnsHostnameBlock gets a reference to the given IpamsvcInheritedDDNSHostnameBlock and assigns it to the DdnsHostnameBlock field.
func (o *IpamsvcServerInheritance) SetDdnsHostnameBlock(v IpamsvcInheritedDDNSHostnameBlock) {
	o.DdnsHostnameBlock = &v
}

// GetDdnsTtlPercent returns the DdnsTtlPercent field value if set, zero value otherwise.
func (o *IpamsvcServerInheritance) GetDdnsTtlPercent() InheritanceInheritedFloat {
	if o == nil || IsNil(o.DdnsTtlPercent) {
		var ret InheritanceInheritedFloat
		return ret
	}
	return *o.DdnsTtlPercent
}

// GetDdnsTtlPercentOk returns a tuple with the DdnsTtlPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcServerInheritance) GetDdnsTtlPercentOk() (*InheritanceInheritedFloat, bool) {
	if o == nil || IsNil(o.DdnsTtlPercent) {
		return nil, false
	}
	return o.DdnsTtlPercent, true
}

// HasDdnsTtlPercent returns a boolean if a field has been set.
func (o *IpamsvcServerInheritance) HasDdnsTtlPercent() bool {
	if o != nil && !IsNil(o.DdnsTtlPercent) {
		return true
	}

	return false
}

// SetDdnsTtlPercent gets a reference to the given InheritanceInheritedFloat and assigns it to the DdnsTtlPercent field.
func (o *IpamsvcServerInheritance) SetDdnsTtlPercent(v InheritanceInheritedFloat) {
	o.DdnsTtlPercent = &v
}

// GetDdnsUpdateOnRenew returns the DdnsUpdateOnRenew field value if set, zero value otherwise.
func (o *IpamsvcServerInheritance) GetDdnsUpdateOnRenew() InheritanceInheritedBool {
	if o == nil || IsNil(o.DdnsUpdateOnRenew) {
		var ret InheritanceInheritedBool
		return ret
	}
	return *o.DdnsUpdateOnRenew
}

// GetDdnsUpdateOnRenewOk returns a tuple with the DdnsUpdateOnRenew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcServerInheritance) GetDdnsUpdateOnRenewOk() (*InheritanceInheritedBool, bool) {
	if o == nil || IsNil(o.DdnsUpdateOnRenew) {
		return nil, false
	}
	return o.DdnsUpdateOnRenew, true
}

// HasDdnsUpdateOnRenew returns a boolean if a field has been set.
func (o *IpamsvcServerInheritance) HasDdnsUpdateOnRenew() bool {
	if o != nil && !IsNil(o.DdnsUpdateOnRenew) {
		return true
	}

	return false
}

// SetDdnsUpdateOnRenew gets a reference to the given InheritanceInheritedBool and assigns it to the DdnsUpdateOnRenew field.
func (o *IpamsvcServerInheritance) SetDdnsUpdateOnRenew(v InheritanceInheritedBool) {
	o.DdnsUpdateOnRenew = &v
}

// GetDdnsUseConflictResolution returns the DdnsUseConflictResolution field value if set, zero value otherwise.
func (o *IpamsvcServerInheritance) GetDdnsUseConflictResolution() InheritanceInheritedBool {
	if o == nil || IsNil(o.DdnsUseConflictResolution) {
		var ret InheritanceInheritedBool
		return ret
	}
	return *o.DdnsUseConflictResolution
}

// GetDdnsUseConflictResolutionOk returns a tuple with the DdnsUseConflictResolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcServerInheritance) GetDdnsUseConflictResolutionOk() (*InheritanceInheritedBool, bool) {
	if o == nil || IsNil(o.DdnsUseConflictResolution) {
		return nil, false
	}
	return o.DdnsUseConflictResolution, true
}

// HasDdnsUseConflictResolution returns a boolean if a field has been set.
func (o *IpamsvcServerInheritance) HasDdnsUseConflictResolution() bool {
	if o != nil && !IsNil(o.DdnsUseConflictResolution) {
		return true
	}

	return false
}

// SetDdnsUseConflictResolution gets a reference to the given InheritanceInheritedBool and assigns it to the DdnsUseConflictResolution field.
func (o *IpamsvcServerInheritance) SetDdnsUseConflictResolution(v InheritanceInheritedBool) {
	o.DdnsUseConflictResolution = &v
}

// GetDhcpConfig returns the DhcpConfig field value if set, zero value otherwise.
func (o *IpamsvcServerInheritance) GetDhcpConfig() IpamsvcInheritedDHCPConfig {
	if o == nil || IsNil(o.DhcpConfig) {
		var ret IpamsvcInheritedDHCPConfig
		return ret
	}
	return *o.DhcpConfig
}

// GetDhcpConfigOk returns a tuple with the DhcpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcServerInheritance) GetDhcpConfigOk() (*IpamsvcInheritedDHCPConfig, bool) {
	if o == nil || IsNil(o.DhcpConfig) {
		return nil, false
	}
	return o.DhcpConfig, true
}

// HasDhcpConfig returns a boolean if a field has been set.
func (o *IpamsvcServerInheritance) HasDhcpConfig() bool {
	if o != nil && !IsNil(o.DhcpConfig) {
		return true
	}

	return false
}

// SetDhcpConfig gets a reference to the given IpamsvcInheritedDHCPConfig and assigns it to the DhcpConfig field.
func (o *IpamsvcServerInheritance) SetDhcpConfig(v IpamsvcInheritedDHCPConfig) {
	o.DhcpConfig = &v
}

// GetDhcpOptions returns the DhcpOptions field value if set, zero value otherwise.
func (o *IpamsvcServerInheritance) GetDhcpOptions() IpamsvcInheritedDHCPOptionList {
	if o == nil || IsNil(o.DhcpOptions) {
		var ret IpamsvcInheritedDHCPOptionList
		return ret
	}
	return *o.DhcpOptions
}

// GetDhcpOptionsOk returns a tuple with the DhcpOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcServerInheritance) GetDhcpOptionsOk() (*IpamsvcInheritedDHCPOptionList, bool) {
	if o == nil || IsNil(o.DhcpOptions) {
		return nil, false
	}
	return o.DhcpOptions, true
}

// HasDhcpOptions returns a boolean if a field has been set.
func (o *IpamsvcServerInheritance) HasDhcpOptions() bool {
	if o != nil && !IsNil(o.DhcpOptions) {
		return true
	}

	return false
}

// SetDhcpOptions gets a reference to the given IpamsvcInheritedDHCPOptionList and assigns it to the DhcpOptions field.
func (o *IpamsvcServerInheritance) SetDhcpOptions(v IpamsvcInheritedDHCPOptionList) {
	o.DhcpOptions = &v
}

// GetDhcpOptionsV6 returns the DhcpOptionsV6 field value if set, zero value otherwise.
func (o *IpamsvcServerInheritance) GetDhcpOptionsV6() IpamsvcInheritedDHCPOptionList {
	if o == nil || IsNil(o.DhcpOptionsV6) {
		var ret IpamsvcInheritedDHCPOptionList
		return ret
	}
	return *o.DhcpOptionsV6
}

// GetDhcpOptionsV6Ok returns a tuple with the DhcpOptionsV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcServerInheritance) GetDhcpOptionsV6Ok() (*IpamsvcInheritedDHCPOptionList, bool) {
	if o == nil || IsNil(o.DhcpOptionsV6) {
		return nil, false
	}
	return o.DhcpOptionsV6, true
}

// HasDhcpOptionsV6 returns a boolean if a field has been set.
func (o *IpamsvcServerInheritance) HasDhcpOptionsV6() bool {
	if o != nil && !IsNil(o.DhcpOptionsV6) {
		return true
	}

	return false
}

// SetDhcpOptionsV6 gets a reference to the given IpamsvcInheritedDHCPOptionList and assigns it to the DhcpOptionsV6 field.
func (o *IpamsvcServerInheritance) SetDhcpOptionsV6(v IpamsvcInheritedDHCPOptionList) {
	o.DhcpOptionsV6 = &v
}

// GetHeaderOptionFilename returns the HeaderOptionFilename field value if set, zero value otherwise.
func (o *IpamsvcServerInheritance) GetHeaderOptionFilename() InheritanceInheritedString {
	if o == nil || IsNil(o.HeaderOptionFilename) {
		var ret InheritanceInheritedString
		return ret
	}
	return *o.HeaderOptionFilename
}

// GetHeaderOptionFilenameOk returns a tuple with the HeaderOptionFilename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcServerInheritance) GetHeaderOptionFilenameOk() (*InheritanceInheritedString, bool) {
	if o == nil || IsNil(o.HeaderOptionFilename) {
		return nil, false
	}
	return o.HeaderOptionFilename, true
}

// HasHeaderOptionFilename returns a boolean if a field has been set.
func (o *IpamsvcServerInheritance) HasHeaderOptionFilename() bool {
	if o != nil && !IsNil(o.HeaderOptionFilename) {
		return true
	}

	return false
}

// SetHeaderOptionFilename gets a reference to the given InheritanceInheritedString and assigns it to the HeaderOptionFilename field.
func (o *IpamsvcServerInheritance) SetHeaderOptionFilename(v InheritanceInheritedString) {
	o.HeaderOptionFilename = &v
}

// GetHeaderOptionServerAddress returns the HeaderOptionServerAddress field value if set, zero value otherwise.
func (o *IpamsvcServerInheritance) GetHeaderOptionServerAddress() InheritanceInheritedString {
	if o == nil || IsNil(o.HeaderOptionServerAddress) {
		var ret InheritanceInheritedString
		return ret
	}
	return *o.HeaderOptionServerAddress
}

// GetHeaderOptionServerAddressOk returns a tuple with the HeaderOptionServerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcServerInheritance) GetHeaderOptionServerAddressOk() (*InheritanceInheritedString, bool) {
	if o == nil || IsNil(o.HeaderOptionServerAddress) {
		return nil, false
	}
	return o.HeaderOptionServerAddress, true
}

// HasHeaderOptionServerAddress returns a boolean if a field has been set.
func (o *IpamsvcServerInheritance) HasHeaderOptionServerAddress() bool {
	if o != nil && !IsNil(o.HeaderOptionServerAddress) {
		return true
	}

	return false
}

// SetHeaderOptionServerAddress gets a reference to the given InheritanceInheritedString and assigns it to the HeaderOptionServerAddress field.
func (o *IpamsvcServerInheritance) SetHeaderOptionServerAddress(v InheritanceInheritedString) {
	o.HeaderOptionServerAddress = &v
}

// GetHeaderOptionServerName returns the HeaderOptionServerName field value if set, zero value otherwise.
func (o *IpamsvcServerInheritance) GetHeaderOptionServerName() InheritanceInheritedString {
	if o == nil || IsNil(o.HeaderOptionServerName) {
		var ret InheritanceInheritedString
		return ret
	}
	return *o.HeaderOptionServerName
}

// GetHeaderOptionServerNameOk returns a tuple with the HeaderOptionServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcServerInheritance) GetHeaderOptionServerNameOk() (*InheritanceInheritedString, bool) {
	if o == nil || IsNil(o.HeaderOptionServerName) {
		return nil, false
	}
	return o.HeaderOptionServerName, true
}

// HasHeaderOptionServerName returns a boolean if a field has been set.
func (o *IpamsvcServerInheritance) HasHeaderOptionServerName() bool {
	if o != nil && !IsNil(o.HeaderOptionServerName) {
		return true
	}

	return false
}

// SetHeaderOptionServerName gets a reference to the given InheritanceInheritedString and assigns it to the HeaderOptionServerName field.
func (o *IpamsvcServerInheritance) SetHeaderOptionServerName(v InheritanceInheritedString) {
	o.HeaderOptionServerName = &v
}

// GetHostnameRewriteBlock returns the HostnameRewriteBlock field value if set, zero value otherwise.
func (o *IpamsvcServerInheritance) GetHostnameRewriteBlock() IpamsvcInheritedHostnameRewriteBlock {
	if o == nil || IsNil(o.HostnameRewriteBlock) {
		var ret IpamsvcInheritedHostnameRewriteBlock
		return ret
	}
	return *o.HostnameRewriteBlock
}

// GetHostnameRewriteBlockOk returns a tuple with the HostnameRewriteBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcServerInheritance) GetHostnameRewriteBlockOk() (*IpamsvcInheritedHostnameRewriteBlock, bool) {
	if o == nil || IsNil(o.HostnameRewriteBlock) {
		return nil, false
	}
	return o.HostnameRewriteBlock, true
}

// HasHostnameRewriteBlock returns a boolean if a field has been set.
func (o *IpamsvcServerInheritance) HasHostnameRewriteBlock() bool {
	if o != nil && !IsNil(o.HostnameRewriteBlock) {
		return true
	}

	return false
}

// SetHostnameRewriteBlock gets a reference to the given IpamsvcInheritedHostnameRewriteBlock and assigns it to the HostnameRewriteBlock field.
func (o *IpamsvcServerInheritance) SetHostnameRewriteBlock(v IpamsvcInheritedHostnameRewriteBlock) {
	o.HostnameRewriteBlock = &v
}

// GetVendorSpecificOptionOptionSpace returns the VendorSpecificOptionOptionSpace field value if set, zero value otherwise.
func (o *IpamsvcServerInheritance) GetVendorSpecificOptionOptionSpace() InheritanceInheritedIdentifier {
	if o == nil || IsNil(o.VendorSpecificOptionOptionSpace) {
		var ret InheritanceInheritedIdentifier
		return ret
	}
	return *o.VendorSpecificOptionOptionSpace
}

// GetVendorSpecificOptionOptionSpaceOk returns a tuple with the VendorSpecificOptionOptionSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcServerInheritance) GetVendorSpecificOptionOptionSpaceOk() (*InheritanceInheritedIdentifier, bool) {
	if o == nil || IsNil(o.VendorSpecificOptionOptionSpace) {
		return nil, false
	}
	return o.VendorSpecificOptionOptionSpace, true
}

// HasVendorSpecificOptionOptionSpace returns a boolean if a field has been set.
func (o *IpamsvcServerInheritance) HasVendorSpecificOptionOptionSpace() bool {
	if o != nil && !IsNil(o.VendorSpecificOptionOptionSpace) {
		return true
	}

	return false
}

// SetVendorSpecificOptionOptionSpace gets a reference to the given InheritanceInheritedIdentifier and assigns it to the VendorSpecificOptionOptionSpace field.
func (o *IpamsvcServerInheritance) SetVendorSpecificOptionOptionSpace(v InheritanceInheritedIdentifier) {
	o.VendorSpecificOptionOptionSpace = &v
}

func (o IpamsvcServerInheritance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcServerInheritance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DdnsBlock) {
		toSerialize["ddns_block"] = o.DdnsBlock
	}
	if !IsNil(o.DdnsClientUpdate) {
		toSerialize["ddns_client_update"] = o.DdnsClientUpdate
	}
	if !IsNil(o.DdnsConflictResolutionMode) {
		toSerialize["ddns_conflict_resolution_mode"] = o.DdnsConflictResolutionMode
	}
	if !IsNil(o.DdnsHostnameBlock) {
		toSerialize["ddns_hostname_block"] = o.DdnsHostnameBlock
	}
	if !IsNil(o.DdnsTtlPercent) {
		toSerialize["ddns_ttl_percent"] = o.DdnsTtlPercent
	}
	if !IsNil(o.DdnsUpdateOnRenew) {
		toSerialize["ddns_update_on_renew"] = o.DdnsUpdateOnRenew
	}
	if !IsNil(o.DdnsUseConflictResolution) {
		toSerialize["ddns_use_conflict_resolution"] = o.DdnsUseConflictResolution
	}
	if !IsNil(o.DhcpConfig) {
		toSerialize["dhcp_config"] = o.DhcpConfig
	}
	if !IsNil(o.DhcpOptions) {
		toSerialize["dhcp_options"] = o.DhcpOptions
	}
	if !IsNil(o.DhcpOptionsV6) {
		toSerialize["dhcp_options_v6"] = o.DhcpOptionsV6
	}
	if !IsNil(o.HeaderOptionFilename) {
		toSerialize["header_option_filename"] = o.HeaderOptionFilename
	}
	if !IsNil(o.HeaderOptionServerAddress) {
		toSerialize["header_option_server_address"] = o.HeaderOptionServerAddress
	}
	if !IsNil(o.HeaderOptionServerName) {
		toSerialize["header_option_server_name"] = o.HeaderOptionServerName
	}
	if !IsNil(o.HostnameRewriteBlock) {
		toSerialize["hostname_rewrite_block"] = o.HostnameRewriteBlock
	}
	if !IsNil(o.VendorSpecificOptionOptionSpace) {
		toSerialize["vendor_specific_option_option_space"] = o.VendorSpecificOptionOptionSpace
	}
	return toSerialize, nil
}

type NullableIpamsvcServerInheritance struct {
	value *IpamsvcServerInheritance
	isSet bool
}

func (v NullableIpamsvcServerInheritance) Get() *IpamsvcServerInheritance {
	return v.value
}

func (v *NullableIpamsvcServerInheritance) Set(val *IpamsvcServerInheritance) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcServerInheritance) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcServerInheritance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcServerInheritance(val *IpamsvcServerInheritance) *NullableIpamsvcServerInheritance {
	return &NullableIpamsvcServerInheritance{value: val, isSet: true}
}

func (v NullableIpamsvcServerInheritance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcServerInheritance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
