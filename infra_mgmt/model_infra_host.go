/*
Infrastructure Management API

The **Infrastructure Management API** provides a RESTful interface to manage Infrastructure Hosts and Services objects.  The following is a list of the different Services and their string types (the string types are to be used with the APIs for the `service_type` field):  | Service name | Service type |   | ------ | ------ |   | Access Authentication | authn |   | Anycast | anycast |   | Data Connector | cdc |   | DHCP | dhcp |   | DNS | dns |   | DNS Forwarding Proxy | dfp |   | NIOS Grid Connector | orpheus |   | MS AD Sync | msad |   | NTP | ntp |   | BGP | bgp |   | RIP | rip |   | OSPF | ospf |    ---   ### Hosts API  The Hosts API is used to manage the Infrastructure Host resources. These include various operations related to hosts such as viewing, creating, updating, replacing, disconnecting, and deleting Hosts. Management of Hosts is done from the Cloud Services Portal (CSP) by navigating to the Manage -> Infrastructure -> Hosts tab.  ---   ### Services API  The Services API is used to manage the Infrastructure Service resources (a.k.a. BloxOne applications). These include various operations related to hosts such as viewing, creating, updating, starting/stopping, configuring, and deleting Services. Management of Services is done from the Cloud Services Portal (CSP) by navigating to the Manage -> Infrastructure -> Services tab.  ---   ### Detail APIs  The Detail APIs are read-only APIs used to list all the Infrastructure resources (Hosts and Services). Each resource record returned also contains information about its other associated resources and the status information for itself and the associated resource(s) (i.e., Host/Service status).  ---

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infra_mgmt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the InfraHost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InfraHost{}

// InfraHost Infrastructure Service
type InfraHost struct {
	// The list of Host-specific configurations for each Service deployed on this Host.
	Configs []InfraServiceHostConfig `json:"configs,omitempty"`
	// Represents the connectivity monitor properties of a Host, to enable/disable connectivity monitoring for redundant network interfaces.  The \"endpoint_type\" is: - `\"csp\"` for enabling monitoring - `\"\"` for disabling monitoring (default)  Note: Currently, all fields except \"endpoint_type\" are read-only, and will be overridden to default values in case they are edited.  Example: ``` {   \"connectivity_monitor\": {     \"cost\":1000000,     \"endpoint_type\":\"csp\",     \"endpoint\":\"http://csp.infoblox.com\",     \"interval\":15,     \"failure_threshold\":1,     \"success_threshold\":2   } } ```
	ConnectivityMonitor map[string]interface{} `json:"connectivity_monitor,omitempty"`
	// The timestamp of creation of Host.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The creator of the Host (internal use only).
	CreatedBy *string `json:"created_by,omitempty"`
	// The description of the Host (optional).
	Description *string `json:"description,omitempty"`
	// The name of the Host (unique).
	DisplayName string `json:"display_name"`
	// The sub-type of a specific Host type.  Example: For Host type BloxOne Appliance, sub-type could be \"B105\" or \"VEP1425\"
	HostSubtype *string `json:"host_subtype,omitempty"`
	// The type of Host.  Should be one of: 1. NIOS , 2. NIOS HA, 3. BloxOne VM , 4. BloxOne Appliance, 5. BloxOne Container, 6. CNIOS
	HostType *string `json:"host_type,omitempty"`
	// The version of the Host platform services.
	HostVersion *string `json:"host_version,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// The IP address of the Host.
	IpAddress *string `json:"ip_address,omitempty"`
	// The IP Space of the Host.
	IpSpace *string `json:"ip_space,omitempty"`
	// The legacy Host object identifier.
	LegacyId *string `json:"legacy_id,omitempty"`
	// The resource identifier.
	LocationId *string `json:"location_id,omitempty"`
	// The MAC address of the Host.
	MacAddress      *string `json:"mac_address,omitempty"`
	MaintenanceMode *string `json:"maintenance_mode,omitempty"`
	// The NAT IP address of the Host.
	NatIp *string `json:"nat_ip,omitempty"`
	// The CSP cluster identifier (internal use only).
	NoaCluster *string `json:"noa_cluster,omitempty"`
	// The unique On-Prem Host ID generated by the On-Prem device and assigned to the Host once it is registered and logged into the Infoblox Cloud.
	Ophid *string `json:"ophid,omitempty"`
	// The resource identifier.
	PoolId *string `json:"pool_id,omitempty"`
	// The unique serial number of the Host.
	SerialNumber *string `json:"serial_number,omitempty"`
	// Tags associated with this Host.
	Tags map[string]interface{} `json:"tags,omitempty"`
	// The timezone of the Host.
	Timezone *string `json:"timezone,omitempty"`
	// The timestamp of the latest update on Host.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type _InfraHost InfraHost

// NewInfraHost instantiates a new InfraHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfraHost(displayName string) *InfraHost {
	this := InfraHost{}
	this.DisplayName = displayName
	return &this
}

// NewInfraHostWithDefaults instantiates a new InfraHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfraHostWithDefaults() *InfraHost {
	this := InfraHost{}
	return &this
}

// GetConfigs returns the Configs field value if set, zero value otherwise.
func (o *InfraHost) GetConfigs() []InfraServiceHostConfig {
	if o == nil || IsNil(o.Configs) {
		var ret []InfraServiceHostConfig
		return ret
	}
	return o.Configs
}

// GetConfigsOk returns a tuple with the Configs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHost) GetConfigsOk() ([]InfraServiceHostConfig, bool) {
	if o == nil || IsNil(o.Configs) {
		return nil, false
	}
	return o.Configs, true
}

// HasConfigs returns a boolean if a field has been set.
func (o *InfraHost) HasConfigs() bool {
	if o != nil && !IsNil(o.Configs) {
		return true
	}

	return false
}

// SetConfigs gets a reference to the given []InfraServiceHostConfig and assigns it to the Configs field.
func (o *InfraHost) SetConfigs(v []InfraServiceHostConfig) {
	o.Configs = v
}

// GetConnectivityMonitor returns the ConnectivityMonitor field value if set, zero value otherwise.
func (o *InfraHost) GetConnectivityMonitor() map[string]interface{} {
	if o == nil || IsNil(o.ConnectivityMonitor) {
		var ret map[string]interface{}
		return ret
	}
	return o.ConnectivityMonitor
}

// GetConnectivityMonitorOk returns a tuple with the ConnectivityMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHost) GetConnectivityMonitorOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ConnectivityMonitor) {
		return map[string]interface{}{}, false
	}
	return o.ConnectivityMonitor, true
}

// HasConnectivityMonitor returns a boolean if a field has been set.
func (o *InfraHost) HasConnectivityMonitor() bool {
	if o != nil && !IsNil(o.ConnectivityMonitor) {
		return true
	}

	return false
}

// SetConnectivityMonitor gets a reference to the given map[string]interface{} and assigns it to the ConnectivityMonitor field.
func (o *InfraHost) SetConnectivityMonitor(v map[string]interface{}) {
	o.ConnectivityMonitor = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InfraHost) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHost) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InfraHost) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *InfraHost) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *InfraHost) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHost) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *InfraHost) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *InfraHost) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InfraHost) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHost) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InfraHost) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InfraHost) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value
func (o *InfraHost) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *InfraHost) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *InfraHost) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetHostSubtype returns the HostSubtype field value if set, zero value otherwise.
func (o *InfraHost) GetHostSubtype() string {
	if o == nil || IsNil(o.HostSubtype) {
		var ret string
		return ret
	}
	return *o.HostSubtype
}

// GetHostSubtypeOk returns a tuple with the HostSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHost) GetHostSubtypeOk() (*string, bool) {
	if o == nil || IsNil(o.HostSubtype) {
		return nil, false
	}
	return o.HostSubtype, true
}

// HasHostSubtype returns a boolean if a field has been set.
func (o *InfraHost) HasHostSubtype() bool {
	if o != nil && !IsNil(o.HostSubtype) {
		return true
	}

	return false
}

// SetHostSubtype gets a reference to the given string and assigns it to the HostSubtype field.
func (o *InfraHost) SetHostSubtype(v string) {
	o.HostSubtype = &v
}

// GetHostType returns the HostType field value if set, zero value otherwise.
func (o *InfraHost) GetHostType() string {
	if o == nil || IsNil(o.HostType) {
		var ret string
		return ret
	}
	return *o.HostType
}

// GetHostTypeOk returns a tuple with the HostType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHost) GetHostTypeOk() (*string, bool) {
	if o == nil || IsNil(o.HostType) {
		return nil, false
	}
	return o.HostType, true
}

// HasHostType returns a boolean if a field has been set.
func (o *InfraHost) HasHostType() bool {
	if o != nil && !IsNil(o.HostType) {
		return true
	}

	return false
}

// SetHostType gets a reference to the given string and assigns it to the HostType field.
func (o *InfraHost) SetHostType(v string) {
	o.HostType = &v
}

// GetHostVersion returns the HostVersion field value if set, zero value otherwise.
func (o *InfraHost) GetHostVersion() string {
	if o == nil || IsNil(o.HostVersion) {
		var ret string
		return ret
	}
	return *o.HostVersion
}

// GetHostVersionOk returns a tuple with the HostVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHost) GetHostVersionOk() (*string, bool) {
	if o == nil || IsNil(o.HostVersion) {
		return nil, false
	}
	return o.HostVersion, true
}

// HasHostVersion returns a boolean if a field has been set.
func (o *InfraHost) HasHostVersion() bool {
	if o != nil && !IsNil(o.HostVersion) {
		return true
	}

	return false
}

// SetHostVersion gets a reference to the given string and assigns it to the HostVersion field.
func (o *InfraHost) SetHostVersion(v string) {
	o.HostVersion = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InfraHost) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHost) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InfraHost) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InfraHost) SetId(v string) {
	o.Id = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *InfraHost) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHost) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *InfraHost) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *InfraHost) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetIpSpace returns the IpSpace field value if set, zero value otherwise.
func (o *InfraHost) GetIpSpace() string {
	if o == nil || IsNil(o.IpSpace) {
		var ret string
		return ret
	}
	return *o.IpSpace
}

// GetIpSpaceOk returns a tuple with the IpSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHost) GetIpSpaceOk() (*string, bool) {
	if o == nil || IsNil(o.IpSpace) {
		return nil, false
	}
	return o.IpSpace, true
}

// HasIpSpace returns a boolean if a field has been set.
func (o *InfraHost) HasIpSpace() bool {
	if o != nil && !IsNil(o.IpSpace) {
		return true
	}

	return false
}

// SetIpSpace gets a reference to the given string and assigns it to the IpSpace field.
func (o *InfraHost) SetIpSpace(v string) {
	o.IpSpace = &v
}

// GetLegacyId returns the LegacyId field value if set, zero value otherwise.
func (o *InfraHost) GetLegacyId() string {
	if o == nil || IsNil(o.LegacyId) {
		var ret string
		return ret
	}
	return *o.LegacyId
}

// GetLegacyIdOk returns a tuple with the LegacyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHost) GetLegacyIdOk() (*string, bool) {
	if o == nil || IsNil(o.LegacyId) {
		return nil, false
	}
	return o.LegacyId, true
}

// HasLegacyId returns a boolean if a field has been set.
func (o *InfraHost) HasLegacyId() bool {
	if o != nil && !IsNil(o.LegacyId) {
		return true
	}

	return false
}

// SetLegacyId gets a reference to the given string and assigns it to the LegacyId field.
func (o *InfraHost) SetLegacyId(v string) {
	o.LegacyId = &v
}

// GetLocationId returns the LocationId field value if set, zero value otherwise.
func (o *InfraHost) GetLocationId() string {
	if o == nil || IsNil(o.LocationId) {
		var ret string
		return ret
	}
	return *o.LocationId
}

// GetLocationIdOk returns a tuple with the LocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHost) GetLocationIdOk() (*string, bool) {
	if o == nil || IsNil(o.LocationId) {
		return nil, false
	}
	return o.LocationId, true
}

// HasLocationId returns a boolean if a field has been set.
func (o *InfraHost) HasLocationId() bool {
	if o != nil && !IsNil(o.LocationId) {
		return true
	}

	return false
}

// SetLocationId gets a reference to the given string and assigns it to the LocationId field.
func (o *InfraHost) SetLocationId(v string) {
	o.LocationId = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *InfraHost) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress) {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHost) GetMacAddressOk() (*string, bool) {
	if o == nil || IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *InfraHost) HasMacAddress() bool {
	if o != nil && !IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *InfraHost) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetMaintenanceMode returns the MaintenanceMode field value if set, zero value otherwise.
func (o *InfraHost) GetMaintenanceMode() string {
	if o == nil || IsNil(o.MaintenanceMode) {
		var ret string
		return ret
	}
	return *o.MaintenanceMode
}

// GetMaintenanceModeOk returns a tuple with the MaintenanceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHost) GetMaintenanceModeOk() (*string, bool) {
	if o == nil || IsNil(o.MaintenanceMode) {
		return nil, false
	}
	return o.MaintenanceMode, true
}

// HasMaintenanceMode returns a boolean if a field has been set.
func (o *InfraHost) HasMaintenanceMode() bool {
	if o != nil && !IsNil(o.MaintenanceMode) {
		return true
	}

	return false
}

// SetMaintenanceMode gets a reference to the given string and assigns it to the MaintenanceMode field.
func (o *InfraHost) SetMaintenanceMode(v string) {
	o.MaintenanceMode = &v
}

// GetNatIp returns the NatIp field value if set, zero value otherwise.
func (o *InfraHost) GetNatIp() string {
	if o == nil || IsNil(o.NatIp) {
		var ret string
		return ret
	}
	return *o.NatIp
}

// GetNatIpOk returns a tuple with the NatIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHost) GetNatIpOk() (*string, bool) {
	if o == nil || IsNil(o.NatIp) {
		return nil, false
	}
	return o.NatIp, true
}

// HasNatIp returns a boolean if a field has been set.
func (o *InfraHost) HasNatIp() bool {
	if o != nil && !IsNil(o.NatIp) {
		return true
	}

	return false
}

// SetNatIp gets a reference to the given string and assigns it to the NatIp field.
func (o *InfraHost) SetNatIp(v string) {
	o.NatIp = &v
}

// GetNoaCluster returns the NoaCluster field value if set, zero value otherwise.
func (o *InfraHost) GetNoaCluster() string {
	if o == nil || IsNil(o.NoaCluster) {
		var ret string
		return ret
	}
	return *o.NoaCluster
}

// GetNoaClusterOk returns a tuple with the NoaCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHost) GetNoaClusterOk() (*string, bool) {
	if o == nil || IsNil(o.NoaCluster) {
		return nil, false
	}
	return o.NoaCluster, true
}

// HasNoaCluster returns a boolean if a field has been set.
func (o *InfraHost) HasNoaCluster() bool {
	if o != nil && !IsNil(o.NoaCluster) {
		return true
	}

	return false
}

// SetNoaCluster gets a reference to the given string and assigns it to the NoaCluster field.
func (o *InfraHost) SetNoaCluster(v string) {
	o.NoaCluster = &v
}

// GetOphid returns the Ophid field value if set, zero value otherwise.
func (o *InfraHost) GetOphid() string {
	if o == nil || IsNil(o.Ophid) {
		var ret string
		return ret
	}
	return *o.Ophid
}

// GetOphidOk returns a tuple with the Ophid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHost) GetOphidOk() (*string, bool) {
	if o == nil || IsNil(o.Ophid) {
		return nil, false
	}
	return o.Ophid, true
}

// HasOphid returns a boolean if a field has been set.
func (o *InfraHost) HasOphid() bool {
	if o != nil && !IsNil(o.Ophid) {
		return true
	}

	return false
}

// SetOphid gets a reference to the given string and assigns it to the Ophid field.
func (o *InfraHost) SetOphid(v string) {
	o.Ophid = &v
}

// GetPoolId returns the PoolId field value if set, zero value otherwise.
func (o *InfraHost) GetPoolId() string {
	if o == nil || IsNil(o.PoolId) {
		var ret string
		return ret
	}
	return *o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHost) GetPoolIdOk() (*string, bool) {
	if o == nil || IsNil(o.PoolId) {
		return nil, false
	}
	return o.PoolId, true
}

// HasPoolId returns a boolean if a field has been set.
func (o *InfraHost) HasPoolId() bool {
	if o != nil && !IsNil(o.PoolId) {
		return true
	}

	return false
}

// SetPoolId gets a reference to the given string and assigns it to the PoolId field.
func (o *InfraHost) SetPoolId(v string) {
	o.PoolId = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *InfraHost) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHost) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *InfraHost) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *InfraHost) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InfraHost) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHost) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InfraHost) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *InfraHost) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *InfraHost) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHost) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *InfraHost) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *InfraHost) SetTimezone(v string) {
	o.Timezone = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *InfraHost) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraHost) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *InfraHost) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *InfraHost) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o InfraHost) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InfraHost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Configs) {
		toSerialize["configs"] = o.Configs
	}
	if !IsNil(o.ConnectivityMonitor) {
		toSerialize["connectivity_monitor"] = o.ConnectivityMonitor
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["display_name"] = o.DisplayName
	if !IsNil(o.HostSubtype) {
		toSerialize["host_subtype"] = o.HostSubtype
	}
	if !IsNil(o.HostType) {
		toSerialize["host_type"] = o.HostType
	}
	if !IsNil(o.HostVersion) {
		toSerialize["host_version"] = o.HostVersion
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IpAddress) {
		toSerialize["ip_address"] = o.IpAddress
	}
	if !IsNil(o.IpSpace) {
		toSerialize["ip_space"] = o.IpSpace
	}
	if !IsNil(o.LegacyId) {
		toSerialize["legacy_id"] = o.LegacyId
	}
	if !IsNil(o.LocationId) {
		toSerialize["location_id"] = o.LocationId
	}
	if !IsNil(o.MacAddress) {
		toSerialize["mac_address"] = o.MacAddress
	}
	if !IsNil(o.MaintenanceMode) {
		toSerialize["maintenance_mode"] = o.MaintenanceMode
	}
	if !IsNil(o.NatIp) {
		toSerialize["nat_ip"] = o.NatIp
	}
	if !IsNil(o.NoaCluster) {
		toSerialize["noa_cluster"] = o.NoaCluster
	}
	if !IsNil(o.Ophid) {
		toSerialize["ophid"] = o.Ophid
	}
	if !IsNil(o.PoolId) {
		toSerialize["pool_id"] = o.PoolId
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["serial_number"] = o.SerialNumber
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

func (o *InfraHost) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"display_name",
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

	varInfraHost := _InfraHost{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInfraHost)

	if err != nil {
		return err
	}

	*o = InfraHost(varInfraHost)

	return err
}

type NullableInfraHost struct {
	value *InfraHost
	isSet bool
}

func (v NullableInfraHost) Get() *InfraHost {
	return v.value
}

func (v *NullableInfraHost) Set(val *InfraHost) {
	v.value = val
	v.isSet = true
}

func (v NullableInfraHost) IsSet() bool {
	return v.isSet
}

func (v *NullableInfraHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfraHost(val *InfraHost) *NullableInfraHost {
	return &NullableInfraHost{value: val, isSet: true}
}

func (v NullableInfraHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfraHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
