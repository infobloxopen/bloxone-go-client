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

// checks if the IpamsvcHost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcHost{}

// IpamsvcHost A DHCP __Host__ (_dhcp/host_) object associates a DHCP Config Profile with an on-prem host.
type IpamsvcHost struct {
	// The primary IP address of the on-prem host.
	Address *string `json:"address,omitempty"`
	// Anycast address configured to the host. Order is not significant.
	AnycastAddresses []string                     `json:"anycast_addresses,omitempty"`
	AssociatedServer *IpamsvcHostAssociatedServer `json:"associated_server,omitempty"`
	// The description for the on-prem host.
	Comment *string `json:"comment,omitempty"`
	// Current dhcp application version of the host.
	CurrentVersion *string `json:"current_version,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// The resource identifier.
	IpSpace *string `json:"ip_space,omitempty"`
	// The display name of the on-prem host.
	Name *string `json:"name,omitempty"`
	// The on-prem host ID.
	Ophid *string `json:"ophid,omitempty"`
	// External provider identifier.
	ProviderId *string `json:"provider_id,omitempty"`
	// The resource identifier.
	Server *string `json:"server,omitempty"`
	// The tags of the on-prem host in JSON format.
	Tags map[string]interface{} `json:"tags,omitempty"`
	// Defines the type of host. Allowed values:  * _bloxone_ddi_: host type is BloxOne DDI,  * _microsoft_azure_: host type is Microsoft Azure,  * _amazon_web_service_: host type is Amazon Web Services.  * _microsoft_active_directory_: host type is Microsoft Active Directory.
	Type *string `json:"type,omitempty"`
}

// NewIpamsvcHost instantiates a new IpamsvcHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcHost() *IpamsvcHost {
	this := IpamsvcHost{}
	return &this
}

// NewIpamsvcHostWithDefaults instantiates a new IpamsvcHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcHostWithDefaults() *IpamsvcHost {
	this := IpamsvcHost{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *IpamsvcHost) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHost) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *IpamsvcHost) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *IpamsvcHost) SetAddress(v string) {
	o.Address = &v
}

// GetAnycastAddresses returns the AnycastAddresses field value if set, zero value otherwise.
func (o *IpamsvcHost) GetAnycastAddresses() []string {
	if o == nil || IsNil(o.AnycastAddresses) {
		var ret []string
		return ret
	}
	return o.AnycastAddresses
}

// GetAnycastAddressesOk returns a tuple with the AnycastAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHost) GetAnycastAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.AnycastAddresses) {
		return nil, false
	}
	return o.AnycastAddresses, true
}

// HasAnycastAddresses returns a boolean if a field has been set.
func (o *IpamsvcHost) HasAnycastAddresses() bool {
	if o != nil && !IsNil(o.AnycastAddresses) {
		return true
	}

	return false
}

// SetAnycastAddresses gets a reference to the given []string and assigns it to the AnycastAddresses field.
func (o *IpamsvcHost) SetAnycastAddresses(v []string) {
	o.AnycastAddresses = v
}

// GetAssociatedServer returns the AssociatedServer field value if set, zero value otherwise.
func (o *IpamsvcHost) GetAssociatedServer() IpamsvcHostAssociatedServer {
	if o == nil || IsNil(o.AssociatedServer) {
		var ret IpamsvcHostAssociatedServer
		return ret
	}
	return *o.AssociatedServer
}

// GetAssociatedServerOk returns a tuple with the AssociatedServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHost) GetAssociatedServerOk() (*IpamsvcHostAssociatedServer, bool) {
	if o == nil || IsNil(o.AssociatedServer) {
		return nil, false
	}
	return o.AssociatedServer, true
}

// HasAssociatedServer returns a boolean if a field has been set.
func (o *IpamsvcHost) HasAssociatedServer() bool {
	if o != nil && !IsNil(o.AssociatedServer) {
		return true
	}

	return false
}

// SetAssociatedServer gets a reference to the given IpamsvcHostAssociatedServer and assigns it to the AssociatedServer field.
func (o *IpamsvcHost) SetAssociatedServer(v IpamsvcHostAssociatedServer) {
	o.AssociatedServer = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *IpamsvcHost) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHost) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *IpamsvcHost) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *IpamsvcHost) SetComment(v string) {
	o.Comment = &v
}

// GetCurrentVersion returns the CurrentVersion field value if set, zero value otherwise.
func (o *IpamsvcHost) GetCurrentVersion() string {
	if o == nil || IsNil(o.CurrentVersion) {
		var ret string
		return ret
	}
	return *o.CurrentVersion
}

// GetCurrentVersionOk returns a tuple with the CurrentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHost) GetCurrentVersionOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentVersion) {
		return nil, false
	}
	return o.CurrentVersion, true
}

// HasCurrentVersion returns a boolean if a field has been set.
func (o *IpamsvcHost) HasCurrentVersion() bool {
	if o != nil && !IsNil(o.CurrentVersion) {
		return true
	}

	return false
}

// SetCurrentVersion gets a reference to the given string and assigns it to the CurrentVersion field.
func (o *IpamsvcHost) SetCurrentVersion(v string) {
	o.CurrentVersion = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IpamsvcHost) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHost) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IpamsvcHost) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IpamsvcHost) SetId(v string) {
	o.Id = &v
}

// GetIpSpace returns the IpSpace field value if set, zero value otherwise.
func (o *IpamsvcHost) GetIpSpace() string {
	if o == nil || IsNil(o.IpSpace) {
		var ret string
		return ret
	}
	return *o.IpSpace
}

// GetIpSpaceOk returns a tuple with the IpSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHost) GetIpSpaceOk() (*string, bool) {
	if o == nil || IsNil(o.IpSpace) {
		return nil, false
	}
	return o.IpSpace, true
}

// HasIpSpace returns a boolean if a field has been set.
func (o *IpamsvcHost) HasIpSpace() bool {
	if o != nil && !IsNil(o.IpSpace) {
		return true
	}

	return false
}

// SetIpSpace gets a reference to the given string and assigns it to the IpSpace field.
func (o *IpamsvcHost) SetIpSpace(v string) {
	o.IpSpace = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IpamsvcHost) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHost) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IpamsvcHost) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IpamsvcHost) SetName(v string) {
	o.Name = &v
}

// GetOphid returns the Ophid field value if set, zero value otherwise.
func (o *IpamsvcHost) GetOphid() string {
	if o == nil || IsNil(o.Ophid) {
		var ret string
		return ret
	}
	return *o.Ophid
}

// GetOphidOk returns a tuple with the Ophid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHost) GetOphidOk() (*string, bool) {
	if o == nil || IsNil(o.Ophid) {
		return nil, false
	}
	return o.Ophid, true
}

// HasOphid returns a boolean if a field has been set.
func (o *IpamsvcHost) HasOphid() bool {
	if o != nil && !IsNil(o.Ophid) {
		return true
	}

	return false
}

// SetOphid gets a reference to the given string and assigns it to the Ophid field.
func (o *IpamsvcHost) SetOphid(v string) {
	o.Ophid = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *IpamsvcHost) GetProviderId() string {
	if o == nil || IsNil(o.ProviderId) {
		var ret string
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHost) GetProviderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *IpamsvcHost) HasProviderId() bool {
	if o != nil && !IsNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given string and assigns it to the ProviderId field.
func (o *IpamsvcHost) SetProviderId(v string) {
	o.ProviderId = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *IpamsvcHost) GetServer() string {
	if o == nil || IsNil(o.Server) {
		var ret string
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHost) GetServerOk() (*string, bool) {
	if o == nil || IsNil(o.Server) {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *IpamsvcHost) HasServer() bool {
	if o != nil && !IsNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given string and assigns it to the Server field.
func (o *IpamsvcHost) SetServer(v string) {
	o.Server = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IpamsvcHost) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHost) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IpamsvcHost) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *IpamsvcHost) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IpamsvcHost) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHost) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IpamsvcHost) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IpamsvcHost) SetType(v string) {
	o.Type = &v
}

func (o IpamsvcHost) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcHost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.AnycastAddresses) {
		toSerialize["anycast_addresses"] = o.AnycastAddresses
	}
	if !IsNil(o.AssociatedServer) {
		toSerialize["associated_server"] = o.AssociatedServer
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CurrentVersion) {
		toSerialize["current_version"] = o.CurrentVersion
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IpSpace) {
		toSerialize["ip_space"] = o.IpSpace
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Ophid) {
		toSerialize["ophid"] = o.Ophid
	}
	if !IsNil(o.ProviderId) {
		toSerialize["provider_id"] = o.ProviderId
	}
	if !IsNil(o.Server) {
		toSerialize["server"] = o.Server
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableIpamsvcHost struct {
	value *IpamsvcHost
	isSet bool
}

func (v NullableIpamsvcHost) Get() *IpamsvcHost {
	return v.value
}

func (v *NullableIpamsvcHost) Set(val *IpamsvcHost) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcHost) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcHost(val *IpamsvcHost) *NullableIpamsvcHost {
	return &NullableIpamsvcHost{value: val, isSet: true}
}

func (v NullableIpamsvcHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
