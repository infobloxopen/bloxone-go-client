# ConfigView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddEdnsOptionInOutgoingQuery** | Pointer to **bool** | _add_edns_option_in_outgoing_query_ adds client IP, MAC address and view name into outgoing recursive query. Defaults to _false_. | [optional] 
**Comment** | Pointer to **string** | Optional. Comment for view. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The timestamp when the object has been created. | [optional] [readonly] 
**CustomRootNs** | Pointer to [**[]ConfigRootNS**](ConfigRootNS.md) | Optional. List of custom root nameservers. The order does not matter.  Error if empty while _custom_root_ns_enabled_ is _true_. Error if there are duplicate items in the list.  Defaults to empty. | [optional] 
**CustomRootNsEnabled** | Pointer to **bool** | Optional. _true_ to use custom root nameservers instead of the default ones.  The _custom_root_ns_ is validated when enabled.  Defaults to _false_. | [optional] 
**Disabled** | Pointer to **bool** | Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration. | [optional] 
**DnssecEnableValidation** | Pointer to **bool** | Optional. _true_ to perform DNSSEC validation. Ignored if _dnssec_enabled_ is _false_.  Defaults to _true_. | [optional] 
**DnssecEnabled** | Pointer to **bool** | Optional. Master toggle for all DNSSEC processing. Other _dnssec_*_ configuration is unused if this is disabled.  Defaults to _true_. | [optional] 
**DnssecRootKeys** | Pointer to [**[]ConfigTrustAnchor**](ConfigTrustAnchor.md) | DNSSEC root keys. The root keys are not configurable.  A default list is provided by cloud management and included here for config generation. | [optional] [readonly] 
**DnssecTrustAnchors** | Pointer to [**[]ConfigTrustAnchor**](ConfigTrustAnchor.md) | Optional. DNSSEC trust anchors.  Error if there are list items with duplicate (_zone_, _sep_, _algorithm_) combinations.  Defaults to empty. | [optional] 
**DnssecValidateExpiry** | Pointer to **bool** | Optional. _true_ to reject expired DNSSEC keys. Ignored if either _dnssec_enabled_ or _dnssec_enable_validation_ is _false_.  Defaults to _true_. | [optional] 
**DtcConfig** | Pointer to [**ConfigDTCConfig**](ConfigDTCConfig.md) |  | [optional] 
**EcsEnabled** | Pointer to **bool** | Optional. _true_ to enable EDNS client subnet for recursive queries. Other _ecs_*_ fields are ignored if this field is not enabled.  Defaults to _false-. | [optional] 
**EcsForwarding** | Pointer to **bool** | Optional. _true_ to enable ECS options in outbound queries. This functionality has additional overhead so it is disabled by default.  Defaults to _false_. | [optional] 
**EcsPrefixV4** | Pointer to **int64** | Optional. Maximum scope length for v4 ECS.  Unsigned integer, min 1 max 24  Defaults to 24. | [optional] 
**EcsPrefixV6** | Pointer to **int64** | Optional. Maximum scope length for v6 ECS.  Unsigned integer, min 1 max 56  Defaults to 56. | [optional] 
**EcsZones** | Pointer to [**[]ConfigECSZone**](ConfigECSZone.md) | Optional. List of zones where ECS queries may be sent.  Error if empty while _ecs_enabled_ is _true_. Error if there are duplicate FQDNs in the list.  Defaults to empty. | [optional] 
**EdnsUdpSize** | Pointer to **int64** | Optional. _edns_udp_size_ represents the edns UDP size. The size a querying DNS server advertises to the DNS server it’s sending a query to.  Defaults to 1232 bytes. | [optional] 
**FilterAaaaAcl** | Pointer to [**[]ConfigACLItem**](ConfigACLItem.md) | Optional. Specifies a list of client addresses for which AAAA filtering is to be applied.  Defaults to _empty_. | [optional] 
**FilterAaaaOnV4** | Pointer to **string** | _filter_aaaa_on_v4_ allows named to omit some IPv6 addresses when responding to IPv4 clients.  Allowed values: * _yes_, * _no_, * _break_dnssec_.  Defaults to _no_ | [optional] 
**Forwarders** | Pointer to [**[]ConfigForwarder**](ConfigForwarder.md) | Optional. List of forwarders.  Error if empty while _forwarders_only_ or _use_root_forwarders_for_local_resolution_with_b1td_ is _true_. Error if there are items in the list with duplicate addresses.  Defaults to empty. | [optional] 
**ForwardersOnly** | Pointer to **bool** | Optional. _true_ to only forward.  Defaults to _false_. | [optional] 
**GssTsigEnabled** | Pointer to **bool** | _gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.  Defaults to _false_. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InheritanceSources** | Pointer to [**ConfigViewInheritance**](ConfigViewInheritance.md) |  | [optional] 
**IpSpaces** | Pointer to **[]string** | The resource identifier. | [optional] 
**LameTtl** | Pointer to **int64** | Optional. Unused in the current on-prem DNS server implementation.  Unsigned integer, min 0 max 3600 (1h).  Defaults to 600. | [optional] 
**MatchClientsAcl** | Pointer to [**[]ConfigACLItem**](ConfigACLItem.md) | Optional. Specifies which clients have access to the view.  Defaults to empty. | [optional] 
**MatchDestinationsAcl** | Pointer to [**[]ConfigACLItem**](ConfigACLItem.md) | Optional. Specifies which destination addresses have access to the view.  Defaults to empty. | [optional] 
**MatchRecursiveOnly** | Pointer to **bool** | Optional. If _true_ only recursive queries from matching clients access the view.  Defaults to _false_. | [optional] 
**MaxCacheTtl** | Pointer to **int64** | Optional. Seconds to cache positive responses.  Unsigned integer, min 1 max 604800 (7d).  Defaults to 604800 (7d). | [optional] 
**MaxNegativeTtl** | Pointer to **int64** | Optional. Seconds to cache negative responses.  Unsigned integer, min 1 max 604800 (7d).  Defaults to 10800 (3h). | [optional] 
**MaxUdpSize** | Pointer to **int64** | Optional. _max_udp_size_ represents maximum UDP payload size. The maximum number of bytes a responding DNS server will send to a UDP datagram.  Defaults to 1232 bytes. | [optional] 
**MinimalResponses** | Pointer to **bool** | Optional. When enabled, the DNS server will only add records to the authority and additional data sections when they are required.  Defaults to _false_. | [optional] 
**Name** | **string** | Name of view. | 
**Notify** | Pointer to **bool** | _notify_ all external secondary DNS servers.  Defaults to _false_. | [optional] 
**QueryAcl** | Pointer to [**[]ConfigACLItem**](ConfigACLItem.md) | Optional. Clients must match this ACL to make authoritative queries. Also used for recursive queries if that ACL is unset.  Defaults to empty. | [optional] 
**RecursionAcl** | Pointer to [**[]ConfigACLItem**](ConfigACLItem.md) | Optional. Clients must match this ACL to make recursive queries. If this ACL is empty, then the _query_acl_ will be used instead.  Defaults to empty. | [optional] 
**RecursionEnabled** | Pointer to **bool** | Optional. _true_ to allow recursive DNS queries.  Defaults to _true_. | [optional] 
**SortList** | Pointer to [**[]ConfigSortListItem**](ConfigSortListItem.md) | Optional. Specifies a sorted network list for A/AAAA records in DNS query response.  Defaults to _empty_. | [optional] 
**SynthesizeAddressRecordsFromHttps** | Pointer to **bool** | _synthesize_address_records_from_https_ enables/disables creation of A/AAAA records from HTTPS RR Defaults to _false_. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Tagging specifics. | [optional] 
**TransferAcl** | Pointer to [**[]ConfigACLItem**](ConfigACLItem.md) | Optional. Clients must match this ACL to receive zone transfers.  Defaults to empty. | [optional] 
**UpdateAcl** | Pointer to [**[]ConfigACLItem**](ConfigACLItem.md) | Optional. Specifies which hosts are allowed to issue Dynamic DNS updates for authoritative zones of _primary_type_ _cloud_.  Defaults to empty. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The timestamp when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 
**UseForwardersForSubzones** | Pointer to **bool** | Optional. Use default forwarders to resolve queries for subzones.  Defaults to _true_. | [optional] 
**UseRootForwardersForLocalResolutionWithB1td** | Pointer to **bool** | _use_root_forwarders_for_local_resolution_with_b1td_ allows DNS recursive queries sent to root forwarders for local resolution when deployed alongside BloxOne Thread Defense. Defaults to _false_. | [optional] 
**ZoneAuthority** | Pointer to [**ConfigZoneAuthority**](ConfigZoneAuthority.md) |  | [optional] 

## Methods

### NewConfigView

`func NewConfigView(name string, ) *ConfigView`

NewConfigView instantiates a new ConfigView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigViewWithDefaults

`func NewConfigViewWithDefaults() *ConfigView`

NewConfigViewWithDefaults instantiates a new ConfigView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddEdnsOptionInOutgoingQuery

`func (o *ConfigView) GetAddEdnsOptionInOutgoingQuery() bool`

GetAddEdnsOptionInOutgoingQuery returns the AddEdnsOptionInOutgoingQuery field if non-nil, zero value otherwise.

### GetAddEdnsOptionInOutgoingQueryOk

`func (o *ConfigView) GetAddEdnsOptionInOutgoingQueryOk() (*bool, bool)`

GetAddEdnsOptionInOutgoingQueryOk returns a tuple with the AddEdnsOptionInOutgoingQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddEdnsOptionInOutgoingQuery

`func (o *ConfigView) SetAddEdnsOptionInOutgoingQuery(v bool)`

SetAddEdnsOptionInOutgoingQuery sets AddEdnsOptionInOutgoingQuery field to given value.

### HasAddEdnsOptionInOutgoingQuery

`func (o *ConfigView) HasAddEdnsOptionInOutgoingQuery() bool`

HasAddEdnsOptionInOutgoingQuery returns a boolean if a field has been set.

### GetComment

`func (o *ConfigView) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ConfigView) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ConfigView) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ConfigView) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ConfigView) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConfigView) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConfigView) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ConfigView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomRootNs

`func (o *ConfigView) GetCustomRootNs() []ConfigRootNS`

GetCustomRootNs returns the CustomRootNs field if non-nil, zero value otherwise.

### GetCustomRootNsOk

`func (o *ConfigView) GetCustomRootNsOk() (*[]ConfigRootNS, bool)`

GetCustomRootNsOk returns a tuple with the CustomRootNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRootNs

`func (o *ConfigView) SetCustomRootNs(v []ConfigRootNS)`

SetCustomRootNs sets CustomRootNs field to given value.

### HasCustomRootNs

`func (o *ConfigView) HasCustomRootNs() bool`

HasCustomRootNs returns a boolean if a field has been set.

### GetCustomRootNsEnabled

`func (o *ConfigView) GetCustomRootNsEnabled() bool`

GetCustomRootNsEnabled returns the CustomRootNsEnabled field if non-nil, zero value otherwise.

### GetCustomRootNsEnabledOk

`func (o *ConfigView) GetCustomRootNsEnabledOk() (*bool, bool)`

GetCustomRootNsEnabledOk returns a tuple with the CustomRootNsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRootNsEnabled

`func (o *ConfigView) SetCustomRootNsEnabled(v bool)`

SetCustomRootNsEnabled sets CustomRootNsEnabled field to given value.

### HasCustomRootNsEnabled

`func (o *ConfigView) HasCustomRootNsEnabled() bool`

HasCustomRootNsEnabled returns a boolean if a field has been set.

### GetDisabled

`func (o *ConfigView) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ConfigView) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ConfigView) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ConfigView) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDnssecEnableValidation

`func (o *ConfigView) GetDnssecEnableValidation() bool`

GetDnssecEnableValidation returns the DnssecEnableValidation field if non-nil, zero value otherwise.

### GetDnssecEnableValidationOk

`func (o *ConfigView) GetDnssecEnableValidationOk() (*bool, bool)`

GetDnssecEnableValidationOk returns a tuple with the DnssecEnableValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecEnableValidation

`func (o *ConfigView) SetDnssecEnableValidation(v bool)`

SetDnssecEnableValidation sets DnssecEnableValidation field to given value.

### HasDnssecEnableValidation

`func (o *ConfigView) HasDnssecEnableValidation() bool`

HasDnssecEnableValidation returns a boolean if a field has been set.

### GetDnssecEnabled

`func (o *ConfigView) GetDnssecEnabled() bool`

GetDnssecEnabled returns the DnssecEnabled field if non-nil, zero value otherwise.

### GetDnssecEnabledOk

`func (o *ConfigView) GetDnssecEnabledOk() (*bool, bool)`

GetDnssecEnabledOk returns a tuple with the DnssecEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecEnabled

`func (o *ConfigView) SetDnssecEnabled(v bool)`

SetDnssecEnabled sets DnssecEnabled field to given value.

### HasDnssecEnabled

`func (o *ConfigView) HasDnssecEnabled() bool`

HasDnssecEnabled returns a boolean if a field has been set.

### GetDnssecRootKeys

`func (o *ConfigView) GetDnssecRootKeys() []ConfigTrustAnchor`

GetDnssecRootKeys returns the DnssecRootKeys field if non-nil, zero value otherwise.

### GetDnssecRootKeysOk

`func (o *ConfigView) GetDnssecRootKeysOk() (*[]ConfigTrustAnchor, bool)`

GetDnssecRootKeysOk returns a tuple with the DnssecRootKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecRootKeys

`func (o *ConfigView) SetDnssecRootKeys(v []ConfigTrustAnchor)`

SetDnssecRootKeys sets DnssecRootKeys field to given value.

### HasDnssecRootKeys

`func (o *ConfigView) HasDnssecRootKeys() bool`

HasDnssecRootKeys returns a boolean if a field has been set.

### GetDnssecTrustAnchors

`func (o *ConfigView) GetDnssecTrustAnchors() []ConfigTrustAnchor`

GetDnssecTrustAnchors returns the DnssecTrustAnchors field if non-nil, zero value otherwise.

### GetDnssecTrustAnchorsOk

`func (o *ConfigView) GetDnssecTrustAnchorsOk() (*[]ConfigTrustAnchor, bool)`

GetDnssecTrustAnchorsOk returns a tuple with the DnssecTrustAnchors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecTrustAnchors

`func (o *ConfigView) SetDnssecTrustAnchors(v []ConfigTrustAnchor)`

SetDnssecTrustAnchors sets DnssecTrustAnchors field to given value.

### HasDnssecTrustAnchors

`func (o *ConfigView) HasDnssecTrustAnchors() bool`

HasDnssecTrustAnchors returns a boolean if a field has been set.

### GetDnssecValidateExpiry

`func (o *ConfigView) GetDnssecValidateExpiry() bool`

GetDnssecValidateExpiry returns the DnssecValidateExpiry field if non-nil, zero value otherwise.

### GetDnssecValidateExpiryOk

`func (o *ConfigView) GetDnssecValidateExpiryOk() (*bool, bool)`

GetDnssecValidateExpiryOk returns a tuple with the DnssecValidateExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecValidateExpiry

`func (o *ConfigView) SetDnssecValidateExpiry(v bool)`

SetDnssecValidateExpiry sets DnssecValidateExpiry field to given value.

### HasDnssecValidateExpiry

`func (o *ConfigView) HasDnssecValidateExpiry() bool`

HasDnssecValidateExpiry returns a boolean if a field has been set.

### GetDtcConfig

`func (o *ConfigView) GetDtcConfig() ConfigDTCConfig`

GetDtcConfig returns the DtcConfig field if non-nil, zero value otherwise.

### GetDtcConfigOk

`func (o *ConfigView) GetDtcConfigOk() (*ConfigDTCConfig, bool)`

GetDtcConfigOk returns a tuple with the DtcConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcConfig

`func (o *ConfigView) SetDtcConfig(v ConfigDTCConfig)`

SetDtcConfig sets DtcConfig field to given value.

### HasDtcConfig

`func (o *ConfigView) HasDtcConfig() bool`

HasDtcConfig returns a boolean if a field has been set.

### GetEcsEnabled

`func (o *ConfigView) GetEcsEnabled() bool`

GetEcsEnabled returns the EcsEnabled field if non-nil, zero value otherwise.

### GetEcsEnabledOk

`func (o *ConfigView) GetEcsEnabledOk() (*bool, bool)`

GetEcsEnabledOk returns a tuple with the EcsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsEnabled

`func (o *ConfigView) SetEcsEnabled(v bool)`

SetEcsEnabled sets EcsEnabled field to given value.

### HasEcsEnabled

`func (o *ConfigView) HasEcsEnabled() bool`

HasEcsEnabled returns a boolean if a field has been set.

### GetEcsForwarding

`func (o *ConfigView) GetEcsForwarding() bool`

GetEcsForwarding returns the EcsForwarding field if non-nil, zero value otherwise.

### GetEcsForwardingOk

`func (o *ConfigView) GetEcsForwardingOk() (*bool, bool)`

GetEcsForwardingOk returns a tuple with the EcsForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsForwarding

`func (o *ConfigView) SetEcsForwarding(v bool)`

SetEcsForwarding sets EcsForwarding field to given value.

### HasEcsForwarding

`func (o *ConfigView) HasEcsForwarding() bool`

HasEcsForwarding returns a boolean if a field has been set.

### GetEcsPrefixV4

`func (o *ConfigView) GetEcsPrefixV4() int64`

GetEcsPrefixV4 returns the EcsPrefixV4 field if non-nil, zero value otherwise.

### GetEcsPrefixV4Ok

`func (o *ConfigView) GetEcsPrefixV4Ok() (*int64, bool)`

GetEcsPrefixV4Ok returns a tuple with the EcsPrefixV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsPrefixV4

`func (o *ConfigView) SetEcsPrefixV4(v int64)`

SetEcsPrefixV4 sets EcsPrefixV4 field to given value.

### HasEcsPrefixV4

`func (o *ConfigView) HasEcsPrefixV4() bool`

HasEcsPrefixV4 returns a boolean if a field has been set.

### GetEcsPrefixV6

`func (o *ConfigView) GetEcsPrefixV6() int64`

GetEcsPrefixV6 returns the EcsPrefixV6 field if non-nil, zero value otherwise.

### GetEcsPrefixV6Ok

`func (o *ConfigView) GetEcsPrefixV6Ok() (*int64, bool)`

GetEcsPrefixV6Ok returns a tuple with the EcsPrefixV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsPrefixV6

`func (o *ConfigView) SetEcsPrefixV6(v int64)`

SetEcsPrefixV6 sets EcsPrefixV6 field to given value.

### HasEcsPrefixV6

`func (o *ConfigView) HasEcsPrefixV6() bool`

HasEcsPrefixV6 returns a boolean if a field has been set.

### GetEcsZones

`func (o *ConfigView) GetEcsZones() []ConfigECSZone`

GetEcsZones returns the EcsZones field if non-nil, zero value otherwise.

### GetEcsZonesOk

`func (o *ConfigView) GetEcsZonesOk() (*[]ConfigECSZone, bool)`

GetEcsZonesOk returns a tuple with the EcsZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsZones

`func (o *ConfigView) SetEcsZones(v []ConfigECSZone)`

SetEcsZones sets EcsZones field to given value.

### HasEcsZones

`func (o *ConfigView) HasEcsZones() bool`

HasEcsZones returns a boolean if a field has been set.

### GetEdnsUdpSize

`func (o *ConfigView) GetEdnsUdpSize() int64`

GetEdnsUdpSize returns the EdnsUdpSize field if non-nil, zero value otherwise.

### GetEdnsUdpSizeOk

`func (o *ConfigView) GetEdnsUdpSizeOk() (*int64, bool)`

GetEdnsUdpSizeOk returns a tuple with the EdnsUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdnsUdpSize

`func (o *ConfigView) SetEdnsUdpSize(v int64)`

SetEdnsUdpSize sets EdnsUdpSize field to given value.

### HasEdnsUdpSize

`func (o *ConfigView) HasEdnsUdpSize() bool`

HasEdnsUdpSize returns a boolean if a field has been set.

### GetFilterAaaaAcl

`func (o *ConfigView) GetFilterAaaaAcl() []ConfigACLItem`

GetFilterAaaaAcl returns the FilterAaaaAcl field if non-nil, zero value otherwise.

### GetFilterAaaaAclOk

`func (o *ConfigView) GetFilterAaaaAclOk() (*[]ConfigACLItem, bool)`

GetFilterAaaaAclOk returns a tuple with the FilterAaaaAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaaAcl

`func (o *ConfigView) SetFilterAaaaAcl(v []ConfigACLItem)`

SetFilterAaaaAcl sets FilterAaaaAcl field to given value.

### HasFilterAaaaAcl

`func (o *ConfigView) HasFilterAaaaAcl() bool`

HasFilterAaaaAcl returns a boolean if a field has been set.

### GetFilterAaaaOnV4

`func (o *ConfigView) GetFilterAaaaOnV4() string`

GetFilterAaaaOnV4 returns the FilterAaaaOnV4 field if non-nil, zero value otherwise.

### GetFilterAaaaOnV4Ok

`func (o *ConfigView) GetFilterAaaaOnV4Ok() (*string, bool)`

GetFilterAaaaOnV4Ok returns a tuple with the FilterAaaaOnV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaaOnV4

`func (o *ConfigView) SetFilterAaaaOnV4(v string)`

SetFilterAaaaOnV4 sets FilterAaaaOnV4 field to given value.

### HasFilterAaaaOnV4

`func (o *ConfigView) HasFilterAaaaOnV4() bool`

HasFilterAaaaOnV4 returns a boolean if a field has been set.

### GetForwarders

`func (o *ConfigView) GetForwarders() []ConfigForwarder`

GetForwarders returns the Forwarders field if non-nil, zero value otherwise.

### GetForwardersOk

`func (o *ConfigView) GetForwardersOk() (*[]ConfigForwarder, bool)`

GetForwardersOk returns a tuple with the Forwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwarders

`func (o *ConfigView) SetForwarders(v []ConfigForwarder)`

SetForwarders sets Forwarders field to given value.

### HasForwarders

`func (o *ConfigView) HasForwarders() bool`

HasForwarders returns a boolean if a field has been set.

### GetForwardersOnly

`func (o *ConfigView) GetForwardersOnly() bool`

GetForwardersOnly returns the ForwardersOnly field if non-nil, zero value otherwise.

### GetForwardersOnlyOk

`func (o *ConfigView) GetForwardersOnlyOk() (*bool, bool)`

GetForwardersOnlyOk returns a tuple with the ForwardersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardersOnly

`func (o *ConfigView) SetForwardersOnly(v bool)`

SetForwardersOnly sets ForwardersOnly field to given value.

### HasForwardersOnly

`func (o *ConfigView) HasForwardersOnly() bool`

HasForwardersOnly returns a boolean if a field has been set.

### GetGssTsigEnabled

`func (o *ConfigView) GetGssTsigEnabled() bool`

GetGssTsigEnabled returns the GssTsigEnabled field if non-nil, zero value otherwise.

### GetGssTsigEnabledOk

`func (o *ConfigView) GetGssTsigEnabledOk() (*bool, bool)`

GetGssTsigEnabledOk returns a tuple with the GssTsigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigEnabled

`func (o *ConfigView) SetGssTsigEnabled(v bool)`

SetGssTsigEnabled sets GssTsigEnabled field to given value.

### HasGssTsigEnabled

`func (o *ConfigView) HasGssTsigEnabled() bool`

HasGssTsigEnabled returns a boolean if a field has been set.

### GetId

`func (o *ConfigView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritanceSources

`func (o *ConfigView) GetInheritanceSources() ConfigViewInheritance`

GetInheritanceSources returns the InheritanceSources field if non-nil, zero value otherwise.

### GetInheritanceSourcesOk

`func (o *ConfigView) GetInheritanceSourcesOk() (*ConfigViewInheritance, bool)`

GetInheritanceSourcesOk returns a tuple with the InheritanceSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceSources

`func (o *ConfigView) SetInheritanceSources(v ConfigViewInheritance)`

SetInheritanceSources sets InheritanceSources field to given value.

### HasInheritanceSources

`func (o *ConfigView) HasInheritanceSources() bool`

HasInheritanceSources returns a boolean if a field has been set.

### GetIpSpaces

`func (o *ConfigView) GetIpSpaces() []string`

GetIpSpaces returns the IpSpaces field if non-nil, zero value otherwise.

### GetIpSpacesOk

`func (o *ConfigView) GetIpSpacesOk() (*[]string, bool)`

GetIpSpacesOk returns a tuple with the IpSpaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSpaces

`func (o *ConfigView) SetIpSpaces(v []string)`

SetIpSpaces sets IpSpaces field to given value.

### HasIpSpaces

`func (o *ConfigView) HasIpSpaces() bool`

HasIpSpaces returns a boolean if a field has been set.

### GetLameTtl

`func (o *ConfigView) GetLameTtl() int64`

GetLameTtl returns the LameTtl field if non-nil, zero value otherwise.

### GetLameTtlOk

`func (o *ConfigView) GetLameTtlOk() (*int64, bool)`

GetLameTtlOk returns a tuple with the LameTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLameTtl

`func (o *ConfigView) SetLameTtl(v int64)`

SetLameTtl sets LameTtl field to given value.

### HasLameTtl

`func (o *ConfigView) HasLameTtl() bool`

HasLameTtl returns a boolean if a field has been set.

### GetMatchClientsAcl

`func (o *ConfigView) GetMatchClientsAcl() []ConfigACLItem`

GetMatchClientsAcl returns the MatchClientsAcl field if non-nil, zero value otherwise.

### GetMatchClientsAclOk

`func (o *ConfigView) GetMatchClientsAclOk() (*[]ConfigACLItem, bool)`

GetMatchClientsAclOk returns a tuple with the MatchClientsAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchClientsAcl

`func (o *ConfigView) SetMatchClientsAcl(v []ConfigACLItem)`

SetMatchClientsAcl sets MatchClientsAcl field to given value.

### HasMatchClientsAcl

`func (o *ConfigView) HasMatchClientsAcl() bool`

HasMatchClientsAcl returns a boolean if a field has been set.

### GetMatchDestinationsAcl

`func (o *ConfigView) GetMatchDestinationsAcl() []ConfigACLItem`

GetMatchDestinationsAcl returns the MatchDestinationsAcl field if non-nil, zero value otherwise.

### GetMatchDestinationsAclOk

`func (o *ConfigView) GetMatchDestinationsAclOk() (*[]ConfigACLItem, bool)`

GetMatchDestinationsAclOk returns a tuple with the MatchDestinationsAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchDestinationsAcl

`func (o *ConfigView) SetMatchDestinationsAcl(v []ConfigACLItem)`

SetMatchDestinationsAcl sets MatchDestinationsAcl field to given value.

### HasMatchDestinationsAcl

`func (o *ConfigView) HasMatchDestinationsAcl() bool`

HasMatchDestinationsAcl returns a boolean if a field has been set.

### GetMatchRecursiveOnly

`func (o *ConfigView) GetMatchRecursiveOnly() bool`

GetMatchRecursiveOnly returns the MatchRecursiveOnly field if non-nil, zero value otherwise.

### GetMatchRecursiveOnlyOk

`func (o *ConfigView) GetMatchRecursiveOnlyOk() (*bool, bool)`

GetMatchRecursiveOnlyOk returns a tuple with the MatchRecursiveOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchRecursiveOnly

`func (o *ConfigView) SetMatchRecursiveOnly(v bool)`

SetMatchRecursiveOnly sets MatchRecursiveOnly field to given value.

### HasMatchRecursiveOnly

`func (o *ConfigView) HasMatchRecursiveOnly() bool`

HasMatchRecursiveOnly returns a boolean if a field has been set.

### GetMaxCacheTtl

`func (o *ConfigView) GetMaxCacheTtl() int64`

GetMaxCacheTtl returns the MaxCacheTtl field if non-nil, zero value otherwise.

### GetMaxCacheTtlOk

`func (o *ConfigView) GetMaxCacheTtlOk() (*int64, bool)`

GetMaxCacheTtlOk returns a tuple with the MaxCacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheTtl

`func (o *ConfigView) SetMaxCacheTtl(v int64)`

SetMaxCacheTtl sets MaxCacheTtl field to given value.

### HasMaxCacheTtl

`func (o *ConfigView) HasMaxCacheTtl() bool`

HasMaxCacheTtl returns a boolean if a field has been set.

### GetMaxNegativeTtl

`func (o *ConfigView) GetMaxNegativeTtl() int64`

GetMaxNegativeTtl returns the MaxNegativeTtl field if non-nil, zero value otherwise.

### GetMaxNegativeTtlOk

`func (o *ConfigView) GetMaxNegativeTtlOk() (*int64, bool)`

GetMaxNegativeTtlOk returns a tuple with the MaxNegativeTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNegativeTtl

`func (o *ConfigView) SetMaxNegativeTtl(v int64)`

SetMaxNegativeTtl sets MaxNegativeTtl field to given value.

### HasMaxNegativeTtl

`func (o *ConfigView) HasMaxNegativeTtl() bool`

HasMaxNegativeTtl returns a boolean if a field has been set.

### GetMaxUdpSize

`func (o *ConfigView) GetMaxUdpSize() int64`

GetMaxUdpSize returns the MaxUdpSize field if non-nil, zero value otherwise.

### GetMaxUdpSizeOk

`func (o *ConfigView) GetMaxUdpSizeOk() (*int64, bool)`

GetMaxUdpSizeOk returns a tuple with the MaxUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUdpSize

`func (o *ConfigView) SetMaxUdpSize(v int64)`

SetMaxUdpSize sets MaxUdpSize field to given value.

### HasMaxUdpSize

`func (o *ConfigView) HasMaxUdpSize() bool`

HasMaxUdpSize returns a boolean if a field has been set.

### GetMinimalResponses

`func (o *ConfigView) GetMinimalResponses() bool`

GetMinimalResponses returns the MinimalResponses field if non-nil, zero value otherwise.

### GetMinimalResponsesOk

`func (o *ConfigView) GetMinimalResponsesOk() (*bool, bool)`

GetMinimalResponsesOk returns a tuple with the MinimalResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimalResponses

`func (o *ConfigView) SetMinimalResponses(v bool)`

SetMinimalResponses sets MinimalResponses field to given value.

### HasMinimalResponses

`func (o *ConfigView) HasMinimalResponses() bool`

HasMinimalResponses returns a boolean if a field has been set.

### GetName

`func (o *ConfigView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigView) SetName(v string)`

SetName sets Name field to given value.


### GetNotify

`func (o *ConfigView) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *ConfigView) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *ConfigView) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *ConfigView) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetQueryAcl

`func (o *ConfigView) GetQueryAcl() []ConfigACLItem`

GetQueryAcl returns the QueryAcl field if non-nil, zero value otherwise.

### GetQueryAclOk

`func (o *ConfigView) GetQueryAclOk() (*[]ConfigACLItem, bool)`

GetQueryAclOk returns a tuple with the QueryAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryAcl

`func (o *ConfigView) SetQueryAcl(v []ConfigACLItem)`

SetQueryAcl sets QueryAcl field to given value.

### HasQueryAcl

`func (o *ConfigView) HasQueryAcl() bool`

HasQueryAcl returns a boolean if a field has been set.

### GetRecursionAcl

`func (o *ConfigView) GetRecursionAcl() []ConfigACLItem`

GetRecursionAcl returns the RecursionAcl field if non-nil, zero value otherwise.

### GetRecursionAclOk

`func (o *ConfigView) GetRecursionAclOk() (*[]ConfigACLItem, bool)`

GetRecursionAclOk returns a tuple with the RecursionAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursionAcl

`func (o *ConfigView) SetRecursionAcl(v []ConfigACLItem)`

SetRecursionAcl sets RecursionAcl field to given value.

### HasRecursionAcl

`func (o *ConfigView) HasRecursionAcl() bool`

HasRecursionAcl returns a boolean if a field has been set.

### GetRecursionEnabled

`func (o *ConfigView) GetRecursionEnabled() bool`

GetRecursionEnabled returns the RecursionEnabled field if non-nil, zero value otherwise.

### GetRecursionEnabledOk

`func (o *ConfigView) GetRecursionEnabledOk() (*bool, bool)`

GetRecursionEnabledOk returns a tuple with the RecursionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursionEnabled

`func (o *ConfigView) SetRecursionEnabled(v bool)`

SetRecursionEnabled sets RecursionEnabled field to given value.

### HasRecursionEnabled

`func (o *ConfigView) HasRecursionEnabled() bool`

HasRecursionEnabled returns a boolean if a field has been set.

### GetSortList

`func (o *ConfigView) GetSortList() []ConfigSortListItem`

GetSortList returns the SortList field if non-nil, zero value otherwise.

### GetSortListOk

`func (o *ConfigView) GetSortListOk() (*[]ConfigSortListItem, bool)`

GetSortListOk returns a tuple with the SortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortList

`func (o *ConfigView) SetSortList(v []ConfigSortListItem)`

SetSortList sets SortList field to given value.

### HasSortList

`func (o *ConfigView) HasSortList() bool`

HasSortList returns a boolean if a field has been set.

### GetSynthesizeAddressRecordsFromHttps

`func (o *ConfigView) GetSynthesizeAddressRecordsFromHttps() bool`

GetSynthesizeAddressRecordsFromHttps returns the SynthesizeAddressRecordsFromHttps field if non-nil, zero value otherwise.

### GetSynthesizeAddressRecordsFromHttpsOk

`func (o *ConfigView) GetSynthesizeAddressRecordsFromHttpsOk() (*bool, bool)`

GetSynthesizeAddressRecordsFromHttpsOk returns a tuple with the SynthesizeAddressRecordsFromHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthesizeAddressRecordsFromHttps

`func (o *ConfigView) SetSynthesizeAddressRecordsFromHttps(v bool)`

SetSynthesizeAddressRecordsFromHttps sets SynthesizeAddressRecordsFromHttps field to given value.

### HasSynthesizeAddressRecordsFromHttps

`func (o *ConfigView) HasSynthesizeAddressRecordsFromHttps() bool`

HasSynthesizeAddressRecordsFromHttps returns a boolean if a field has been set.

### GetTags

`func (o *ConfigView) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConfigView) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConfigView) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConfigView) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTransferAcl

`func (o *ConfigView) GetTransferAcl() []ConfigACLItem`

GetTransferAcl returns the TransferAcl field if non-nil, zero value otherwise.

### GetTransferAclOk

`func (o *ConfigView) GetTransferAclOk() (*[]ConfigACLItem, bool)`

GetTransferAclOk returns a tuple with the TransferAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAcl

`func (o *ConfigView) SetTransferAcl(v []ConfigACLItem)`

SetTransferAcl sets TransferAcl field to given value.

### HasTransferAcl

`func (o *ConfigView) HasTransferAcl() bool`

HasTransferAcl returns a boolean if a field has been set.

### GetUpdateAcl

`func (o *ConfigView) GetUpdateAcl() []ConfigACLItem`

GetUpdateAcl returns the UpdateAcl field if non-nil, zero value otherwise.

### GetUpdateAclOk

`func (o *ConfigView) GetUpdateAclOk() (*[]ConfigACLItem, bool)`

GetUpdateAclOk returns a tuple with the UpdateAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAcl

`func (o *ConfigView) SetUpdateAcl(v []ConfigACLItem)`

SetUpdateAcl sets UpdateAcl field to given value.

### HasUpdateAcl

`func (o *ConfigView) HasUpdateAcl() bool`

HasUpdateAcl returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ConfigView) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ConfigView) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ConfigView) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ConfigView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUseForwardersForSubzones

`func (o *ConfigView) GetUseForwardersForSubzones() bool`

GetUseForwardersForSubzones returns the UseForwardersForSubzones field if non-nil, zero value otherwise.

### GetUseForwardersForSubzonesOk

`func (o *ConfigView) GetUseForwardersForSubzonesOk() (*bool, bool)`

GetUseForwardersForSubzonesOk returns a tuple with the UseForwardersForSubzones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForwardersForSubzones

`func (o *ConfigView) SetUseForwardersForSubzones(v bool)`

SetUseForwardersForSubzones sets UseForwardersForSubzones field to given value.

### HasUseForwardersForSubzones

`func (o *ConfigView) HasUseForwardersForSubzones() bool`

HasUseForwardersForSubzones returns a boolean if a field has been set.

### GetUseRootForwardersForLocalResolutionWithB1td

`func (o *ConfigView) GetUseRootForwardersForLocalResolutionWithB1td() bool`

GetUseRootForwardersForLocalResolutionWithB1td returns the UseRootForwardersForLocalResolutionWithB1td field if non-nil, zero value otherwise.

### GetUseRootForwardersForLocalResolutionWithB1tdOk

`func (o *ConfigView) GetUseRootForwardersForLocalResolutionWithB1tdOk() (*bool, bool)`

GetUseRootForwardersForLocalResolutionWithB1tdOk returns a tuple with the UseRootForwardersForLocalResolutionWithB1td field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRootForwardersForLocalResolutionWithB1td

`func (o *ConfigView) SetUseRootForwardersForLocalResolutionWithB1td(v bool)`

SetUseRootForwardersForLocalResolutionWithB1td sets UseRootForwardersForLocalResolutionWithB1td field to given value.

### HasUseRootForwardersForLocalResolutionWithB1td

`func (o *ConfigView) HasUseRootForwardersForLocalResolutionWithB1td() bool`

HasUseRootForwardersForLocalResolutionWithB1td returns a boolean if a field has been set.

### GetZoneAuthority

`func (o *ConfigView) GetZoneAuthority() ConfigZoneAuthority`

GetZoneAuthority returns the ZoneAuthority field if non-nil, zero value otherwise.

### GetZoneAuthorityOk

`func (o *ConfigView) GetZoneAuthorityOk() (*ConfigZoneAuthority, bool)`

GetZoneAuthorityOk returns a tuple with the ZoneAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAuthority

`func (o *ConfigView) SetZoneAuthority(v ConfigZoneAuthority)`

SetZoneAuthority sets ZoneAuthority field to given value.

### HasZoneAuthority

`func (o *ConfigView) HasZoneAuthority() bool`

HasZoneAuthority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


