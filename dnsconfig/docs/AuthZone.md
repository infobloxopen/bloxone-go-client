# AuthZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional. Comment for zone configuration. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**Disabled** | Pointer to **bool** | Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration. | [optional] 
**ExternalPrimaries** | Pointer to [**[]ExternalPrimary**](ExternalPrimary.md) | Optional. DNS primaries external to BloxOne DDI. Order is not significant. | [optional] 
**ExternalProviders** | Pointer to [**[]AuthZoneExternalProvider**](AuthZoneExternalProvider.md) | list of external providers for the auth zone. | [optional] [readonly] 
**ExternalSecondaries** | Pointer to [**[]ExternalSecondary**](ExternalSecondary.md) | DNS secondaries external to BloxOne DDI. Order is not significant. | [optional] 
**Fqdn** | Pointer to **string** | Zone FQDN. The FQDN supplied at creation will be converted to canonical form.  Read-only after creation. | [optional] 
**GssTsigEnabled** | Pointer to **bool** | _gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.  Defaults to _false_. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InheritanceAssignedHosts** | Pointer to [**[]Inheritance2AssignedHost**](Inheritance2AssignedHost.md) | The list of the inheritance assigned hosts of the object. | [optional] [readonly] 
**InheritanceSources** | Pointer to [**AuthZoneInheritance**](AuthZoneInheritance.md) |  | [optional] 
**InitialSoaSerial** | Pointer to **int64** | On-create-only. SOA serial is allowed to be set when the authoritative zone is created. | [optional] 
**InternalSecondaries** | Pointer to [**[]InternalSecondary**](InternalSecondary.md) | Optional. BloxOne DDI hosts acting as internal secondaries. Order is not significant. | [optional] 
**MappedSubnet** | Pointer to **string** | Reverse zone network address in the following format: \&quot;ip-address/cidr\&quot;. Defaults to empty. | [optional] [readonly] 
**Mapping** | Pointer to **string** | Zone mapping type. Allowed values:  * _forward_,  * _ipv4_reverse_.  * _ipv6_reverse_.  Defaults to forward. | [optional] [readonly] 
**Notify** | Pointer to **bool** | Also notify all external secondary DNS servers if enabled.  Defaults to _false_. | [optional] 
**Nsgs** | Pointer to **[]string** | The resource identifier. | [optional] 
**Parent** | Pointer to **string** | The resource identifier. | [optional] 
**PrimaryType** | Pointer to **string** | Primary type for an authoritative zone. Read only after creation. Allowed values:  * _external_: zone data owned by an external nameserver,  * _cloud_: zone data is owned by a BloxOne DDI host. | [optional] 
**ProtocolFqdn** | Pointer to **string** | Zone FQDN in punycode. | [optional] [readonly] 
**QueryAcl** | Pointer to [**[]ACLItem**](ACLItem.md) | Optional. Clients must match this ACL to make authoritative queries. Also used for recursive queries if that ACL is unset.  Defaults to empty. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Tagging specifics. | [optional] 
**TransferAcl** | Pointer to [**[]ACLItem**](ACLItem.md) | Optional. Clients must match this ACL to receive zone transfers. | [optional] 
**UpdateAcl** | Pointer to [**[]ACLItem**](ACLItem.md) | Optional. Specifies which hosts are allowed to submit Dynamic DNS updates for authoritative zones of _primary_type_ _cloud_.  Defaults to empty. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 
**UseForwardersForSubzones** | Pointer to **bool** | Optional. Use default forwarders to resolve queries for subzones.  Defaults to _true_. | [optional] 
**View** | Pointer to **string** | The resource identifier. | [optional] 
**Warnings** | Pointer to [**[]Warning**](Warning.md) | The list of an auth zone warnings. | [optional] [readonly] 
**ZoneAuthority** | Pointer to [**ZoneAuthority**](ZoneAuthority.md) |  | [optional] 

## Methods

### NewAuthZone

`func NewAuthZone() *AuthZone`

NewAuthZone instantiates a new AuthZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthZoneWithDefaults

`func NewAuthZoneWithDefaults() *AuthZone`

NewAuthZoneWithDefaults instantiates a new AuthZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *AuthZone) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AuthZone) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AuthZone) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AuthZone) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuthZone) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthZone) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthZone) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthZone) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisabled

`func (o *AuthZone) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AuthZone) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AuthZone) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *AuthZone) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetExternalPrimaries

`func (o *AuthZone) GetExternalPrimaries() []ExternalPrimary`

GetExternalPrimaries returns the ExternalPrimaries field if non-nil, zero value otherwise.

### GetExternalPrimariesOk

`func (o *AuthZone) GetExternalPrimariesOk() (*[]ExternalPrimary, bool)`

GetExternalPrimariesOk returns a tuple with the ExternalPrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPrimaries

`func (o *AuthZone) SetExternalPrimaries(v []ExternalPrimary)`

SetExternalPrimaries sets ExternalPrimaries field to given value.

### HasExternalPrimaries

`func (o *AuthZone) HasExternalPrimaries() bool`

HasExternalPrimaries returns a boolean if a field has been set.

### GetExternalProviders

`func (o *AuthZone) GetExternalProviders() []AuthZoneExternalProvider`

GetExternalProviders returns the ExternalProviders field if non-nil, zero value otherwise.

### GetExternalProvidersOk

`func (o *AuthZone) GetExternalProvidersOk() (*[]AuthZoneExternalProvider, bool)`

GetExternalProvidersOk returns a tuple with the ExternalProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProviders

`func (o *AuthZone) SetExternalProviders(v []AuthZoneExternalProvider)`

SetExternalProviders sets ExternalProviders field to given value.

### HasExternalProviders

`func (o *AuthZone) HasExternalProviders() bool`

HasExternalProviders returns a boolean if a field has been set.

### GetExternalSecondaries

`func (o *AuthZone) GetExternalSecondaries() []ExternalSecondary`

GetExternalSecondaries returns the ExternalSecondaries field if non-nil, zero value otherwise.

### GetExternalSecondariesOk

`func (o *AuthZone) GetExternalSecondariesOk() (*[]ExternalSecondary, bool)`

GetExternalSecondariesOk returns a tuple with the ExternalSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSecondaries

`func (o *AuthZone) SetExternalSecondaries(v []ExternalSecondary)`

SetExternalSecondaries sets ExternalSecondaries field to given value.

### HasExternalSecondaries

`func (o *AuthZone) HasExternalSecondaries() bool`

HasExternalSecondaries returns a boolean if a field has been set.

### GetFqdn

`func (o *AuthZone) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *AuthZone) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *AuthZone) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *AuthZone) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetGssTsigEnabled

`func (o *AuthZone) GetGssTsigEnabled() bool`

GetGssTsigEnabled returns the GssTsigEnabled field if non-nil, zero value otherwise.

### GetGssTsigEnabledOk

`func (o *AuthZone) GetGssTsigEnabledOk() (*bool, bool)`

GetGssTsigEnabledOk returns a tuple with the GssTsigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigEnabled

`func (o *AuthZone) SetGssTsigEnabled(v bool)`

SetGssTsigEnabled sets GssTsigEnabled field to given value.

### HasGssTsigEnabled

`func (o *AuthZone) HasGssTsigEnabled() bool`

HasGssTsigEnabled returns a boolean if a field has been set.

### GetId

`func (o *AuthZone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthZone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthZone) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritanceAssignedHosts

`func (o *AuthZone) GetInheritanceAssignedHosts() []Inheritance2AssignedHost`

GetInheritanceAssignedHosts returns the InheritanceAssignedHosts field if non-nil, zero value otherwise.

### GetInheritanceAssignedHostsOk

`func (o *AuthZone) GetInheritanceAssignedHostsOk() (*[]Inheritance2AssignedHost, bool)`

GetInheritanceAssignedHostsOk returns a tuple with the InheritanceAssignedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceAssignedHosts

`func (o *AuthZone) SetInheritanceAssignedHosts(v []Inheritance2AssignedHost)`

SetInheritanceAssignedHosts sets InheritanceAssignedHosts field to given value.

### HasInheritanceAssignedHosts

`func (o *AuthZone) HasInheritanceAssignedHosts() bool`

HasInheritanceAssignedHosts returns a boolean if a field has been set.

### GetInheritanceSources

`func (o *AuthZone) GetInheritanceSources() AuthZoneInheritance`

GetInheritanceSources returns the InheritanceSources field if non-nil, zero value otherwise.

### GetInheritanceSourcesOk

`func (o *AuthZone) GetInheritanceSourcesOk() (*AuthZoneInheritance, bool)`

GetInheritanceSourcesOk returns a tuple with the InheritanceSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceSources

`func (o *AuthZone) SetInheritanceSources(v AuthZoneInheritance)`

SetInheritanceSources sets InheritanceSources field to given value.

### HasInheritanceSources

`func (o *AuthZone) HasInheritanceSources() bool`

HasInheritanceSources returns a boolean if a field has been set.

### GetInitialSoaSerial

`func (o *AuthZone) GetInitialSoaSerial() int64`

GetInitialSoaSerial returns the InitialSoaSerial field if non-nil, zero value otherwise.

### GetInitialSoaSerialOk

`func (o *AuthZone) GetInitialSoaSerialOk() (*int64, bool)`

GetInitialSoaSerialOk returns a tuple with the InitialSoaSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialSoaSerial

`func (o *AuthZone) SetInitialSoaSerial(v int64)`

SetInitialSoaSerial sets InitialSoaSerial field to given value.

### HasInitialSoaSerial

`func (o *AuthZone) HasInitialSoaSerial() bool`

HasInitialSoaSerial returns a boolean if a field has been set.

### GetInternalSecondaries

`func (o *AuthZone) GetInternalSecondaries() []InternalSecondary`

GetInternalSecondaries returns the InternalSecondaries field if non-nil, zero value otherwise.

### GetInternalSecondariesOk

`func (o *AuthZone) GetInternalSecondariesOk() (*[]InternalSecondary, bool)`

GetInternalSecondariesOk returns a tuple with the InternalSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalSecondaries

`func (o *AuthZone) SetInternalSecondaries(v []InternalSecondary)`

SetInternalSecondaries sets InternalSecondaries field to given value.

### HasInternalSecondaries

`func (o *AuthZone) HasInternalSecondaries() bool`

HasInternalSecondaries returns a boolean if a field has been set.

### GetMappedSubnet

`func (o *AuthZone) GetMappedSubnet() string`

GetMappedSubnet returns the MappedSubnet field if non-nil, zero value otherwise.

### GetMappedSubnetOk

`func (o *AuthZone) GetMappedSubnetOk() (*string, bool)`

GetMappedSubnetOk returns a tuple with the MappedSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedSubnet

`func (o *AuthZone) SetMappedSubnet(v string)`

SetMappedSubnet sets MappedSubnet field to given value.

### HasMappedSubnet

`func (o *AuthZone) HasMappedSubnet() bool`

HasMappedSubnet returns a boolean if a field has been set.

### GetMapping

`func (o *AuthZone) GetMapping() string`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *AuthZone) GetMappingOk() (*string, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *AuthZone) SetMapping(v string)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *AuthZone) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### GetNotify

`func (o *AuthZone) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *AuthZone) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *AuthZone) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *AuthZone) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetNsgs

`func (o *AuthZone) GetNsgs() []string`

GetNsgs returns the Nsgs field if non-nil, zero value otherwise.

### GetNsgsOk

`func (o *AuthZone) GetNsgsOk() (*[]string, bool)`

GetNsgsOk returns a tuple with the Nsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsgs

`func (o *AuthZone) SetNsgs(v []string)`

SetNsgs sets Nsgs field to given value.

### HasNsgs

`func (o *AuthZone) HasNsgs() bool`

HasNsgs returns a boolean if a field has been set.

### GetParent

`func (o *AuthZone) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *AuthZone) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *AuthZone) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *AuthZone) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPrimaryType

`func (o *AuthZone) GetPrimaryType() string`

GetPrimaryType returns the PrimaryType field if non-nil, zero value otherwise.

### GetPrimaryTypeOk

`func (o *AuthZone) GetPrimaryTypeOk() (*string, bool)`

GetPrimaryTypeOk returns a tuple with the PrimaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryType

`func (o *AuthZone) SetPrimaryType(v string)`

SetPrimaryType sets PrimaryType field to given value.

### HasPrimaryType

`func (o *AuthZone) HasPrimaryType() bool`

HasPrimaryType returns a boolean if a field has been set.

### GetProtocolFqdn

`func (o *AuthZone) GetProtocolFqdn() string`

GetProtocolFqdn returns the ProtocolFqdn field if non-nil, zero value otherwise.

### GetProtocolFqdnOk

`func (o *AuthZone) GetProtocolFqdnOk() (*string, bool)`

GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolFqdn

`func (o *AuthZone) SetProtocolFqdn(v string)`

SetProtocolFqdn sets ProtocolFqdn field to given value.

### HasProtocolFqdn

`func (o *AuthZone) HasProtocolFqdn() bool`

HasProtocolFqdn returns a boolean if a field has been set.

### GetQueryAcl

`func (o *AuthZone) GetQueryAcl() []ACLItem`

GetQueryAcl returns the QueryAcl field if non-nil, zero value otherwise.

### GetQueryAclOk

`func (o *AuthZone) GetQueryAclOk() (*[]ACLItem, bool)`

GetQueryAclOk returns a tuple with the QueryAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryAcl

`func (o *AuthZone) SetQueryAcl(v []ACLItem)`

SetQueryAcl sets QueryAcl field to given value.

### HasQueryAcl

`func (o *AuthZone) HasQueryAcl() bool`

HasQueryAcl returns a boolean if a field has been set.

### GetTags

`func (o *AuthZone) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AuthZone) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AuthZone) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *AuthZone) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTransferAcl

`func (o *AuthZone) GetTransferAcl() []ACLItem`

GetTransferAcl returns the TransferAcl field if non-nil, zero value otherwise.

### GetTransferAclOk

`func (o *AuthZone) GetTransferAclOk() (*[]ACLItem, bool)`

GetTransferAclOk returns a tuple with the TransferAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAcl

`func (o *AuthZone) SetTransferAcl(v []ACLItem)`

SetTransferAcl sets TransferAcl field to given value.

### HasTransferAcl

`func (o *AuthZone) HasTransferAcl() bool`

HasTransferAcl returns a boolean if a field has been set.

### GetUpdateAcl

`func (o *AuthZone) GetUpdateAcl() []ACLItem`

GetUpdateAcl returns the UpdateAcl field if non-nil, zero value otherwise.

### GetUpdateAclOk

`func (o *AuthZone) GetUpdateAclOk() (*[]ACLItem, bool)`

GetUpdateAclOk returns a tuple with the UpdateAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAcl

`func (o *AuthZone) SetUpdateAcl(v []ACLItem)`

SetUpdateAcl sets UpdateAcl field to given value.

### HasUpdateAcl

`func (o *AuthZone) HasUpdateAcl() bool`

HasUpdateAcl returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuthZone) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuthZone) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuthZone) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuthZone) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUseForwardersForSubzones

`func (o *AuthZone) GetUseForwardersForSubzones() bool`

GetUseForwardersForSubzones returns the UseForwardersForSubzones field if non-nil, zero value otherwise.

### GetUseForwardersForSubzonesOk

`func (o *AuthZone) GetUseForwardersForSubzonesOk() (*bool, bool)`

GetUseForwardersForSubzonesOk returns a tuple with the UseForwardersForSubzones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForwardersForSubzones

`func (o *AuthZone) SetUseForwardersForSubzones(v bool)`

SetUseForwardersForSubzones sets UseForwardersForSubzones field to given value.

### HasUseForwardersForSubzones

`func (o *AuthZone) HasUseForwardersForSubzones() bool`

HasUseForwardersForSubzones returns a boolean if a field has been set.

### GetView

`func (o *AuthZone) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *AuthZone) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *AuthZone) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *AuthZone) HasView() bool`

HasView returns a boolean if a field has been set.

### GetWarnings

`func (o *AuthZone) GetWarnings() []Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *AuthZone) GetWarningsOk() (*[]Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *AuthZone) SetWarnings(v []Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *AuthZone) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetZoneAuthority

`func (o *AuthZone) GetZoneAuthority() ZoneAuthority`

GetZoneAuthority returns the ZoneAuthority field if non-nil, zero value otherwise.

### GetZoneAuthorityOk

`func (o *AuthZone) GetZoneAuthorityOk() (*ZoneAuthority, bool)`

GetZoneAuthorityOk returns a tuple with the ZoneAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAuthority

`func (o *AuthZone) SetZoneAuthority(v ZoneAuthority)`

SetZoneAuthority sets ZoneAuthority field to given value.

### HasZoneAuthority

`func (o *AuthZone) HasZoneAuthority() bool`

HasZoneAuthority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


