# View

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddEdnsOptionInOutgoingQuery** | Pointer to **bool** | _add_edns_option_in_outgoing_query_ adds client IP, MAC address and view name into outgoing recursive query. Defaults to _false_. | [optional] 
**Comment** | Pointer to **string** | Optional. Comment for view. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The timestamp when the object has been created. | [optional] [readonly] 
**CustomRootNs** | Pointer to [**[]RootNS**](RootNS.md) | Optional. List of custom root nameservers. The order does not matter.  Error if empty while _custom_root_ns_enabled_ is _true_. Error if there are duplicate items in the list.  Defaults to empty. | [optional] 
**CustomRootNsEnabled** | Pointer to **bool** | Optional. _true_ to use custom root nameservers instead of the default ones.  The _custom_root_ns_ is validated when enabled.  Defaults to _false_. | [optional] 
**Disabled** | Pointer to **bool** | Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration. | [optional] 
**DnssecEnableValidation** | Pointer to **bool** | Optional. _true_ to perform DNSSEC validation. Ignored if _dnssec_enabled_ is _false_.  Defaults to _true_. | [optional] 
**DnssecEnabled** | Pointer to **bool** | Optional. Master toggle for all DNSSEC processing. Other _dnssec_*_ configuration is unused if this is disabled.  Defaults to _true_. | [optional] 
**DnssecRootKeys** | Pointer to [**[]TrustAnchor**](TrustAnchor.md) | DNSSEC root keys. The root keys are not configurable.  A default list is provided by cloud management and included here for config generation. | [optional] [readonly] 
**DnssecTrustAnchors** | Pointer to [**[]TrustAnchor**](TrustAnchor.md) | Optional. DNSSEC trust anchors.  Error if there are list items with duplicate (_zone_, _sep_, _algorithm_) combinations.  Defaults to empty. | [optional] 
**DnssecValidateExpiry** | Pointer to **bool** | Optional. _true_ to reject expired DNSSEC keys. Ignored if either _dnssec_enabled_ or _dnssec_enable_validation_ is _false_.  Defaults to _true_. | [optional] 
**DtcConfig** | Pointer to [**DTCConfig**](DTCConfig.md) |  | [optional] 
**EcsEnabled** | Pointer to **bool** | Optional. _true_ to enable EDNS client subnet for recursive queries. Other _ecs_*_ fields are ignored if this field is not enabled.  Defaults to _false-. | [optional] 
**EcsForwarding** | Pointer to **bool** | Optional. _true_ to enable ECS options in outbound queries. This functionality has additional overhead so it is disabled by default.  Defaults to _false_. | [optional] 
**EcsPrefixV4** | Pointer to **int64** | Optional. Maximum scope length for v4 ECS.  Unsigned integer, min 1 max 24  Defaults to 24. | [optional] 
**EcsPrefixV6** | Pointer to **int64** | Optional. Maximum scope length for v6 ECS.  Unsigned integer, min 1 max 56  Defaults to 56. | [optional] 
**EcsZones** | Pointer to [**[]ECSZone**](ECSZone.md) | Optional. List of zones where ECS queries may be sent.  Error if empty while _ecs_enabled_ is _true_. Error if there are duplicate FQDNs in the list.  Defaults to empty. | [optional] 
**EdnsUdpSize** | Pointer to **int64** | Optional. _edns_udp_size_ represents the edns UDP size. The size a querying DNS server advertises to the DNS server it’s sending a query to.  Defaults to 1232 bytes. | [optional] 
**FilterAaaaAcl** | Pointer to [**[]ACLItem**](ACLItem.md) | Optional. Specifies a list of client addresses for which AAAA filtering is to be applied.  Defaults to _empty_. | [optional] 
**FilterAaaaOnV4** | Pointer to **string** | _filter_aaaa_on_v4_ allows named to omit some IPv6 addresses when responding to IPv4 clients.  Allowed values: * _yes_, * _no_, * _break_dnssec_.  Defaults to _no_ | [optional] 
**Forwarders** | Pointer to [**[]Forwarder**](Forwarder.md) | Optional. List of forwarders.  Error if empty while _forwarders_only_ or _use_root_forwarders_for_local_resolution_with_b1td_ is _true_. Error if there are items in the list with duplicate addresses.  Defaults to empty. | [optional] 
**ForwardersOnly** | Pointer to **bool** | Optional. _true_ to only forward.  Defaults to _false_. | [optional] 
**GssTsigEnabled** | Pointer to **bool** | _gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.  Defaults to _false_. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InheritanceSources** | Pointer to [**ViewInheritance**](ViewInheritance.md) |  | [optional] 
**IpSpaces** | Pointer to **[]string** | The resource identifier. | [optional] 
**LameTtl** | Pointer to **int64** | Optional. Unused in the current on-prem DNS server implementation.  Unsigned integer, min 0 max 3600 (1h).  Defaults to 600. | [optional] 
**MatchClientsAcl** | Pointer to [**[]ACLItem**](ACLItem.md) | Optional. Specifies which clients have access to the view.  Defaults to empty. | [optional] 
**MatchDestinationsAcl** | Pointer to [**[]ACLItem**](ACLItem.md) | Optional. Specifies which destination addresses have access to the view.  Defaults to empty. | [optional] 
**MatchRecursiveOnly** | Pointer to **bool** | Optional. If _true_ only recursive queries from matching clients access the view.  Defaults to _false_. | [optional] 
**MaxCacheTtl** | Pointer to **int64** | Optional. Seconds to cache positive responses.  Unsigned integer, min 1 max 604800 (7d).  Defaults to 604800 (7d). | [optional] 
**MaxNegativeTtl** | Pointer to **int64** | Optional. Seconds to cache negative responses.  Unsigned integer, min 1 max 604800 (7d).  Defaults to 10800 (3h). | [optional] 
**MaxUdpSize** | Pointer to **int64** | Optional. _max_udp_size_ represents maximum UDP payload size. The maximum number of bytes a responding DNS server will send to a UDP datagram.  Defaults to 1232 bytes. | [optional] 
**MinimalResponses** | Pointer to **bool** | Optional. When enabled, the DNS server will only add records to the authority and additional data sections when they are required.  Defaults to _false_. | [optional] 
**Name** | **string** | Name of view. | 
**Notify** | Pointer to **bool** | _notify_ all external secondary DNS servers.  Defaults to _false_. | [optional] 
**QueryAcl** | Pointer to [**[]ACLItem**](ACLItem.md) | Optional. Clients must match this ACL to make authoritative queries. Also used for recursive queries if that ACL is unset.  Defaults to empty. | [optional] 
**RecursionAcl** | Pointer to [**[]ACLItem**](ACLItem.md) | Optional. Clients must match this ACL to make recursive queries. If this ACL is empty, then the _query_acl_ will be used instead.  Defaults to empty. | [optional] 
**RecursionEnabled** | Pointer to **bool** | Optional. _true_ to allow recursive DNS queries.  Defaults to _true_. | [optional] 
**SortList** | Pointer to [**[]SortListItem**](SortListItem.md) | Optional. Specifies a sorted network list for A/AAAA records in DNS query response.  Defaults to _empty_. | [optional] 
**SynthesizeAddressRecordsFromHttps** | Pointer to **bool** | _synthesize_address_records_from_https_ enables/disables creation of A/AAAA records from HTTPS RR Defaults to _false_. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Tagging specifics. | [optional] 
**TransferAcl** | Pointer to [**[]ACLItem**](ACLItem.md) | Optional. Clients must match this ACL to receive zone transfers.  Defaults to empty. | [optional] 
**UpdateAcl** | Pointer to [**[]ACLItem**](ACLItem.md) | Optional. Specifies which hosts are allowed to issue Dynamic DNS updates for authoritative zones of _primary_type_ _cloud_.  Defaults to empty. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The timestamp when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 
**UseForwardersForSubzones** | Pointer to **bool** | Optional. Use default forwarders to resolve queries for subzones.  Defaults to _true_. | [optional] 
**UseRootForwardersForLocalResolutionWithB1td** | Pointer to **bool** | _use_root_forwarders_for_local_resolution_with_b1td_ allows DNS recursive queries sent to root forwarders for local resolution when deployed alongside BloxOne Thread Defense. Defaults to _false_. | [optional] 
**ZoneAuthority** | Pointer to [**ZoneAuthority**](ZoneAuthority.md) |  | [optional] 

## Methods

### NewView

`func NewView(name string, ) *View`

NewView instantiates a new View object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewWithDefaults

`func NewViewWithDefaults() *View`

NewViewWithDefaults instantiates a new View object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddEdnsOptionInOutgoingQuery

`func (o *View) GetAddEdnsOptionInOutgoingQuery() bool`

GetAddEdnsOptionInOutgoingQuery returns the AddEdnsOptionInOutgoingQuery field if non-nil, zero value otherwise.

### GetAddEdnsOptionInOutgoingQueryOk

`func (o *View) GetAddEdnsOptionInOutgoingQueryOk() (*bool, bool)`

GetAddEdnsOptionInOutgoingQueryOk returns a tuple with the AddEdnsOptionInOutgoingQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddEdnsOptionInOutgoingQuery

`func (o *View) SetAddEdnsOptionInOutgoingQuery(v bool)`

SetAddEdnsOptionInOutgoingQuery sets AddEdnsOptionInOutgoingQuery field to given value.

### HasAddEdnsOptionInOutgoingQuery

`func (o *View) HasAddEdnsOptionInOutgoingQuery() bool`

HasAddEdnsOptionInOutgoingQuery returns a boolean if a field has been set.

### GetComment

`func (o *View) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *View) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *View) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *View) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *View) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *View) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *View) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *View) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomRootNs

`func (o *View) GetCustomRootNs() []RootNS`

GetCustomRootNs returns the CustomRootNs field if non-nil, zero value otherwise.

### GetCustomRootNsOk

`func (o *View) GetCustomRootNsOk() (*[]RootNS, bool)`

GetCustomRootNsOk returns a tuple with the CustomRootNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRootNs

`func (o *View) SetCustomRootNs(v []RootNS)`

SetCustomRootNs sets CustomRootNs field to given value.

### HasCustomRootNs

`func (o *View) HasCustomRootNs() bool`

HasCustomRootNs returns a boolean if a field has been set.

### GetCustomRootNsEnabled

`func (o *View) GetCustomRootNsEnabled() bool`

GetCustomRootNsEnabled returns the CustomRootNsEnabled field if non-nil, zero value otherwise.

### GetCustomRootNsEnabledOk

`func (o *View) GetCustomRootNsEnabledOk() (*bool, bool)`

GetCustomRootNsEnabledOk returns a tuple with the CustomRootNsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRootNsEnabled

`func (o *View) SetCustomRootNsEnabled(v bool)`

SetCustomRootNsEnabled sets CustomRootNsEnabled field to given value.

### HasCustomRootNsEnabled

`func (o *View) HasCustomRootNsEnabled() bool`

HasCustomRootNsEnabled returns a boolean if a field has been set.

### GetDisabled

`func (o *View) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *View) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *View) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *View) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDnssecEnableValidation

`func (o *View) GetDnssecEnableValidation() bool`

GetDnssecEnableValidation returns the DnssecEnableValidation field if non-nil, zero value otherwise.

### GetDnssecEnableValidationOk

`func (o *View) GetDnssecEnableValidationOk() (*bool, bool)`

GetDnssecEnableValidationOk returns a tuple with the DnssecEnableValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecEnableValidation

`func (o *View) SetDnssecEnableValidation(v bool)`

SetDnssecEnableValidation sets DnssecEnableValidation field to given value.

### HasDnssecEnableValidation

`func (o *View) HasDnssecEnableValidation() bool`

HasDnssecEnableValidation returns a boolean if a field has been set.

### GetDnssecEnabled

`func (o *View) GetDnssecEnabled() bool`

GetDnssecEnabled returns the DnssecEnabled field if non-nil, zero value otherwise.

### GetDnssecEnabledOk

`func (o *View) GetDnssecEnabledOk() (*bool, bool)`

GetDnssecEnabledOk returns a tuple with the DnssecEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecEnabled

`func (o *View) SetDnssecEnabled(v bool)`

SetDnssecEnabled sets DnssecEnabled field to given value.

### HasDnssecEnabled

`func (o *View) HasDnssecEnabled() bool`

HasDnssecEnabled returns a boolean if a field has been set.

### GetDnssecRootKeys

`func (o *View) GetDnssecRootKeys() []TrustAnchor`

GetDnssecRootKeys returns the DnssecRootKeys field if non-nil, zero value otherwise.

### GetDnssecRootKeysOk

`func (o *View) GetDnssecRootKeysOk() (*[]TrustAnchor, bool)`

GetDnssecRootKeysOk returns a tuple with the DnssecRootKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecRootKeys

`func (o *View) SetDnssecRootKeys(v []TrustAnchor)`

SetDnssecRootKeys sets DnssecRootKeys field to given value.

### HasDnssecRootKeys

`func (o *View) HasDnssecRootKeys() bool`

HasDnssecRootKeys returns a boolean if a field has been set.

### GetDnssecTrustAnchors

`func (o *View) GetDnssecTrustAnchors() []TrustAnchor`

GetDnssecTrustAnchors returns the DnssecTrustAnchors field if non-nil, zero value otherwise.

### GetDnssecTrustAnchorsOk

`func (o *View) GetDnssecTrustAnchorsOk() (*[]TrustAnchor, bool)`

GetDnssecTrustAnchorsOk returns a tuple with the DnssecTrustAnchors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecTrustAnchors

`func (o *View) SetDnssecTrustAnchors(v []TrustAnchor)`

SetDnssecTrustAnchors sets DnssecTrustAnchors field to given value.

### HasDnssecTrustAnchors

`func (o *View) HasDnssecTrustAnchors() bool`

HasDnssecTrustAnchors returns a boolean if a field has been set.

### GetDnssecValidateExpiry

`func (o *View) GetDnssecValidateExpiry() bool`

GetDnssecValidateExpiry returns the DnssecValidateExpiry field if non-nil, zero value otherwise.

### GetDnssecValidateExpiryOk

`func (o *View) GetDnssecValidateExpiryOk() (*bool, bool)`

GetDnssecValidateExpiryOk returns a tuple with the DnssecValidateExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecValidateExpiry

`func (o *View) SetDnssecValidateExpiry(v bool)`

SetDnssecValidateExpiry sets DnssecValidateExpiry field to given value.

### HasDnssecValidateExpiry

`func (o *View) HasDnssecValidateExpiry() bool`

HasDnssecValidateExpiry returns a boolean if a field has been set.

### GetDtcConfig

`func (o *View) GetDtcConfig() DTCConfig`

GetDtcConfig returns the DtcConfig field if non-nil, zero value otherwise.

### GetDtcConfigOk

`func (o *View) GetDtcConfigOk() (*DTCConfig, bool)`

GetDtcConfigOk returns a tuple with the DtcConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcConfig

`func (o *View) SetDtcConfig(v DTCConfig)`

SetDtcConfig sets DtcConfig field to given value.

### HasDtcConfig

`func (o *View) HasDtcConfig() bool`

HasDtcConfig returns a boolean if a field has been set.

### GetEcsEnabled

`func (o *View) GetEcsEnabled() bool`

GetEcsEnabled returns the EcsEnabled field if non-nil, zero value otherwise.

### GetEcsEnabledOk

`func (o *View) GetEcsEnabledOk() (*bool, bool)`

GetEcsEnabledOk returns a tuple with the EcsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsEnabled

`func (o *View) SetEcsEnabled(v bool)`

SetEcsEnabled sets EcsEnabled field to given value.

### HasEcsEnabled

`func (o *View) HasEcsEnabled() bool`

HasEcsEnabled returns a boolean if a field has been set.

### GetEcsForwarding

`func (o *View) GetEcsForwarding() bool`

GetEcsForwarding returns the EcsForwarding field if non-nil, zero value otherwise.

### GetEcsForwardingOk

`func (o *View) GetEcsForwardingOk() (*bool, bool)`

GetEcsForwardingOk returns a tuple with the EcsForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsForwarding

`func (o *View) SetEcsForwarding(v bool)`

SetEcsForwarding sets EcsForwarding field to given value.

### HasEcsForwarding

`func (o *View) HasEcsForwarding() bool`

HasEcsForwarding returns a boolean if a field has been set.

### GetEcsPrefixV4

`func (o *View) GetEcsPrefixV4() int64`

GetEcsPrefixV4 returns the EcsPrefixV4 field if non-nil, zero value otherwise.

### GetEcsPrefixV4Ok

`func (o *View) GetEcsPrefixV4Ok() (*int64, bool)`

GetEcsPrefixV4Ok returns a tuple with the EcsPrefixV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsPrefixV4

`func (o *View) SetEcsPrefixV4(v int64)`

SetEcsPrefixV4 sets EcsPrefixV4 field to given value.

### HasEcsPrefixV4

`func (o *View) HasEcsPrefixV4() bool`

HasEcsPrefixV4 returns a boolean if a field has been set.

### GetEcsPrefixV6

`func (o *View) GetEcsPrefixV6() int64`

GetEcsPrefixV6 returns the EcsPrefixV6 field if non-nil, zero value otherwise.

### GetEcsPrefixV6Ok

`func (o *View) GetEcsPrefixV6Ok() (*int64, bool)`

GetEcsPrefixV6Ok returns a tuple with the EcsPrefixV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsPrefixV6

`func (o *View) SetEcsPrefixV6(v int64)`

SetEcsPrefixV6 sets EcsPrefixV6 field to given value.

### HasEcsPrefixV6

`func (o *View) HasEcsPrefixV6() bool`

HasEcsPrefixV6 returns a boolean if a field has been set.

### GetEcsZones

`func (o *View) GetEcsZones() []ECSZone`

GetEcsZones returns the EcsZones field if non-nil, zero value otherwise.

### GetEcsZonesOk

`func (o *View) GetEcsZonesOk() (*[]ECSZone, bool)`

GetEcsZonesOk returns a tuple with the EcsZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsZones

`func (o *View) SetEcsZones(v []ECSZone)`

SetEcsZones sets EcsZones field to given value.

### HasEcsZones

`func (o *View) HasEcsZones() bool`

HasEcsZones returns a boolean if a field has been set.

### GetEdnsUdpSize

`func (o *View) GetEdnsUdpSize() int64`

GetEdnsUdpSize returns the EdnsUdpSize field if non-nil, zero value otherwise.

### GetEdnsUdpSizeOk

`func (o *View) GetEdnsUdpSizeOk() (*int64, bool)`

GetEdnsUdpSizeOk returns a tuple with the EdnsUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdnsUdpSize

`func (o *View) SetEdnsUdpSize(v int64)`

SetEdnsUdpSize sets EdnsUdpSize field to given value.

### HasEdnsUdpSize

`func (o *View) HasEdnsUdpSize() bool`

HasEdnsUdpSize returns a boolean if a field has been set.

### GetFilterAaaaAcl

`func (o *View) GetFilterAaaaAcl() []ACLItem`

GetFilterAaaaAcl returns the FilterAaaaAcl field if non-nil, zero value otherwise.

### GetFilterAaaaAclOk

`func (o *View) GetFilterAaaaAclOk() (*[]ACLItem, bool)`

GetFilterAaaaAclOk returns a tuple with the FilterAaaaAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaaAcl

`func (o *View) SetFilterAaaaAcl(v []ACLItem)`

SetFilterAaaaAcl sets FilterAaaaAcl field to given value.

### HasFilterAaaaAcl

`func (o *View) HasFilterAaaaAcl() bool`

HasFilterAaaaAcl returns a boolean if a field has been set.

### GetFilterAaaaOnV4

`func (o *View) GetFilterAaaaOnV4() string`

GetFilterAaaaOnV4 returns the FilterAaaaOnV4 field if non-nil, zero value otherwise.

### GetFilterAaaaOnV4Ok

`func (o *View) GetFilterAaaaOnV4Ok() (*string, bool)`

GetFilterAaaaOnV4Ok returns a tuple with the FilterAaaaOnV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaaOnV4

`func (o *View) SetFilterAaaaOnV4(v string)`

SetFilterAaaaOnV4 sets FilterAaaaOnV4 field to given value.

### HasFilterAaaaOnV4

`func (o *View) HasFilterAaaaOnV4() bool`

HasFilterAaaaOnV4 returns a boolean if a field has been set.

### GetForwarders

`func (o *View) GetForwarders() []Forwarder`

GetForwarders returns the Forwarders field if non-nil, zero value otherwise.

### GetForwardersOk

`func (o *View) GetForwardersOk() (*[]Forwarder, bool)`

GetForwardersOk returns a tuple with the Forwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwarders

`func (o *View) SetForwarders(v []Forwarder)`

SetForwarders sets Forwarders field to given value.

### HasForwarders

`func (o *View) HasForwarders() bool`

HasForwarders returns a boolean if a field has been set.

### GetForwardersOnly

`func (o *View) GetForwardersOnly() bool`

GetForwardersOnly returns the ForwardersOnly field if non-nil, zero value otherwise.

### GetForwardersOnlyOk

`func (o *View) GetForwardersOnlyOk() (*bool, bool)`

GetForwardersOnlyOk returns a tuple with the ForwardersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardersOnly

`func (o *View) SetForwardersOnly(v bool)`

SetForwardersOnly sets ForwardersOnly field to given value.

### HasForwardersOnly

`func (o *View) HasForwardersOnly() bool`

HasForwardersOnly returns a boolean if a field has been set.

### GetGssTsigEnabled

`func (o *View) GetGssTsigEnabled() bool`

GetGssTsigEnabled returns the GssTsigEnabled field if non-nil, zero value otherwise.

### GetGssTsigEnabledOk

`func (o *View) GetGssTsigEnabledOk() (*bool, bool)`

GetGssTsigEnabledOk returns a tuple with the GssTsigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigEnabled

`func (o *View) SetGssTsigEnabled(v bool)`

SetGssTsigEnabled sets GssTsigEnabled field to given value.

### HasGssTsigEnabled

`func (o *View) HasGssTsigEnabled() bool`

HasGssTsigEnabled returns a boolean if a field has been set.

### GetId

`func (o *View) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *View) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *View) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *View) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritanceSources

`func (o *View) GetInheritanceSources() ViewInheritance`

GetInheritanceSources returns the InheritanceSources field if non-nil, zero value otherwise.

### GetInheritanceSourcesOk

`func (o *View) GetInheritanceSourcesOk() (*ViewInheritance, bool)`

GetInheritanceSourcesOk returns a tuple with the InheritanceSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceSources

`func (o *View) SetInheritanceSources(v ViewInheritance)`

SetInheritanceSources sets InheritanceSources field to given value.

### HasInheritanceSources

`func (o *View) HasInheritanceSources() bool`

HasInheritanceSources returns a boolean if a field has been set.

### GetIpSpaces

`func (o *View) GetIpSpaces() []string`

GetIpSpaces returns the IpSpaces field if non-nil, zero value otherwise.

### GetIpSpacesOk

`func (o *View) GetIpSpacesOk() (*[]string, bool)`

GetIpSpacesOk returns a tuple with the IpSpaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSpaces

`func (o *View) SetIpSpaces(v []string)`

SetIpSpaces sets IpSpaces field to given value.

### HasIpSpaces

`func (o *View) HasIpSpaces() bool`

HasIpSpaces returns a boolean if a field has been set.

### GetLameTtl

`func (o *View) GetLameTtl() int64`

GetLameTtl returns the LameTtl field if non-nil, zero value otherwise.

### GetLameTtlOk

`func (o *View) GetLameTtlOk() (*int64, bool)`

GetLameTtlOk returns a tuple with the LameTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLameTtl

`func (o *View) SetLameTtl(v int64)`

SetLameTtl sets LameTtl field to given value.

### HasLameTtl

`func (o *View) HasLameTtl() bool`

HasLameTtl returns a boolean if a field has been set.

### GetMatchClientsAcl

`func (o *View) GetMatchClientsAcl() []ACLItem`

GetMatchClientsAcl returns the MatchClientsAcl field if non-nil, zero value otherwise.

### GetMatchClientsAclOk

`func (o *View) GetMatchClientsAclOk() (*[]ACLItem, bool)`

GetMatchClientsAclOk returns a tuple with the MatchClientsAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchClientsAcl

`func (o *View) SetMatchClientsAcl(v []ACLItem)`

SetMatchClientsAcl sets MatchClientsAcl field to given value.

### HasMatchClientsAcl

`func (o *View) HasMatchClientsAcl() bool`

HasMatchClientsAcl returns a boolean if a field has been set.

### GetMatchDestinationsAcl

`func (o *View) GetMatchDestinationsAcl() []ACLItem`

GetMatchDestinationsAcl returns the MatchDestinationsAcl field if non-nil, zero value otherwise.

### GetMatchDestinationsAclOk

`func (o *View) GetMatchDestinationsAclOk() (*[]ACLItem, bool)`

GetMatchDestinationsAclOk returns a tuple with the MatchDestinationsAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchDestinationsAcl

`func (o *View) SetMatchDestinationsAcl(v []ACLItem)`

SetMatchDestinationsAcl sets MatchDestinationsAcl field to given value.

### HasMatchDestinationsAcl

`func (o *View) HasMatchDestinationsAcl() bool`

HasMatchDestinationsAcl returns a boolean if a field has been set.

### GetMatchRecursiveOnly

`func (o *View) GetMatchRecursiveOnly() bool`

GetMatchRecursiveOnly returns the MatchRecursiveOnly field if non-nil, zero value otherwise.

### GetMatchRecursiveOnlyOk

`func (o *View) GetMatchRecursiveOnlyOk() (*bool, bool)`

GetMatchRecursiveOnlyOk returns a tuple with the MatchRecursiveOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchRecursiveOnly

`func (o *View) SetMatchRecursiveOnly(v bool)`

SetMatchRecursiveOnly sets MatchRecursiveOnly field to given value.

### HasMatchRecursiveOnly

`func (o *View) HasMatchRecursiveOnly() bool`

HasMatchRecursiveOnly returns a boolean if a field has been set.

### GetMaxCacheTtl

`func (o *View) GetMaxCacheTtl() int64`

GetMaxCacheTtl returns the MaxCacheTtl field if non-nil, zero value otherwise.

### GetMaxCacheTtlOk

`func (o *View) GetMaxCacheTtlOk() (*int64, bool)`

GetMaxCacheTtlOk returns a tuple with the MaxCacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheTtl

`func (o *View) SetMaxCacheTtl(v int64)`

SetMaxCacheTtl sets MaxCacheTtl field to given value.

### HasMaxCacheTtl

`func (o *View) HasMaxCacheTtl() bool`

HasMaxCacheTtl returns a boolean if a field has been set.

### GetMaxNegativeTtl

`func (o *View) GetMaxNegativeTtl() int64`

GetMaxNegativeTtl returns the MaxNegativeTtl field if non-nil, zero value otherwise.

### GetMaxNegativeTtlOk

`func (o *View) GetMaxNegativeTtlOk() (*int64, bool)`

GetMaxNegativeTtlOk returns a tuple with the MaxNegativeTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNegativeTtl

`func (o *View) SetMaxNegativeTtl(v int64)`

SetMaxNegativeTtl sets MaxNegativeTtl field to given value.

### HasMaxNegativeTtl

`func (o *View) HasMaxNegativeTtl() bool`

HasMaxNegativeTtl returns a boolean if a field has been set.

### GetMaxUdpSize

`func (o *View) GetMaxUdpSize() int64`

GetMaxUdpSize returns the MaxUdpSize field if non-nil, zero value otherwise.

### GetMaxUdpSizeOk

`func (o *View) GetMaxUdpSizeOk() (*int64, bool)`

GetMaxUdpSizeOk returns a tuple with the MaxUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUdpSize

`func (o *View) SetMaxUdpSize(v int64)`

SetMaxUdpSize sets MaxUdpSize field to given value.

### HasMaxUdpSize

`func (o *View) HasMaxUdpSize() bool`

HasMaxUdpSize returns a boolean if a field has been set.

### GetMinimalResponses

`func (o *View) GetMinimalResponses() bool`

GetMinimalResponses returns the MinimalResponses field if non-nil, zero value otherwise.

### GetMinimalResponsesOk

`func (o *View) GetMinimalResponsesOk() (*bool, bool)`

GetMinimalResponsesOk returns a tuple with the MinimalResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimalResponses

`func (o *View) SetMinimalResponses(v bool)`

SetMinimalResponses sets MinimalResponses field to given value.

### HasMinimalResponses

`func (o *View) HasMinimalResponses() bool`

HasMinimalResponses returns a boolean if a field has been set.

### GetName

`func (o *View) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *View) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *View) SetName(v string)`

SetName sets Name field to given value.


### GetNotify

`func (o *View) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *View) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *View) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *View) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetQueryAcl

`func (o *View) GetQueryAcl() []ACLItem`

GetQueryAcl returns the QueryAcl field if non-nil, zero value otherwise.

### GetQueryAclOk

`func (o *View) GetQueryAclOk() (*[]ACLItem, bool)`

GetQueryAclOk returns a tuple with the QueryAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryAcl

`func (o *View) SetQueryAcl(v []ACLItem)`

SetQueryAcl sets QueryAcl field to given value.

### HasQueryAcl

`func (o *View) HasQueryAcl() bool`

HasQueryAcl returns a boolean if a field has been set.

### GetRecursionAcl

`func (o *View) GetRecursionAcl() []ACLItem`

GetRecursionAcl returns the RecursionAcl field if non-nil, zero value otherwise.

### GetRecursionAclOk

`func (o *View) GetRecursionAclOk() (*[]ACLItem, bool)`

GetRecursionAclOk returns a tuple with the RecursionAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursionAcl

`func (o *View) SetRecursionAcl(v []ACLItem)`

SetRecursionAcl sets RecursionAcl field to given value.

### HasRecursionAcl

`func (o *View) HasRecursionAcl() bool`

HasRecursionAcl returns a boolean if a field has been set.

### GetRecursionEnabled

`func (o *View) GetRecursionEnabled() bool`

GetRecursionEnabled returns the RecursionEnabled field if non-nil, zero value otherwise.

### GetRecursionEnabledOk

`func (o *View) GetRecursionEnabledOk() (*bool, bool)`

GetRecursionEnabledOk returns a tuple with the RecursionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursionEnabled

`func (o *View) SetRecursionEnabled(v bool)`

SetRecursionEnabled sets RecursionEnabled field to given value.

### HasRecursionEnabled

`func (o *View) HasRecursionEnabled() bool`

HasRecursionEnabled returns a boolean if a field has been set.

### GetSortList

`func (o *View) GetSortList() []SortListItem`

GetSortList returns the SortList field if non-nil, zero value otherwise.

### GetSortListOk

`func (o *View) GetSortListOk() (*[]SortListItem, bool)`

GetSortListOk returns a tuple with the SortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortList

`func (o *View) SetSortList(v []SortListItem)`

SetSortList sets SortList field to given value.

### HasSortList

`func (o *View) HasSortList() bool`

HasSortList returns a boolean if a field has been set.

### GetSynthesizeAddressRecordsFromHttps

`func (o *View) GetSynthesizeAddressRecordsFromHttps() bool`

GetSynthesizeAddressRecordsFromHttps returns the SynthesizeAddressRecordsFromHttps field if non-nil, zero value otherwise.

### GetSynthesizeAddressRecordsFromHttpsOk

`func (o *View) GetSynthesizeAddressRecordsFromHttpsOk() (*bool, bool)`

GetSynthesizeAddressRecordsFromHttpsOk returns a tuple with the SynthesizeAddressRecordsFromHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthesizeAddressRecordsFromHttps

`func (o *View) SetSynthesizeAddressRecordsFromHttps(v bool)`

SetSynthesizeAddressRecordsFromHttps sets SynthesizeAddressRecordsFromHttps field to given value.

### HasSynthesizeAddressRecordsFromHttps

`func (o *View) HasSynthesizeAddressRecordsFromHttps() bool`

HasSynthesizeAddressRecordsFromHttps returns a boolean if a field has been set.

### GetTags

`func (o *View) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *View) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *View) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *View) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTransferAcl

`func (o *View) GetTransferAcl() []ACLItem`

GetTransferAcl returns the TransferAcl field if non-nil, zero value otherwise.

### GetTransferAclOk

`func (o *View) GetTransferAclOk() (*[]ACLItem, bool)`

GetTransferAclOk returns a tuple with the TransferAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAcl

`func (o *View) SetTransferAcl(v []ACLItem)`

SetTransferAcl sets TransferAcl field to given value.

### HasTransferAcl

`func (o *View) HasTransferAcl() bool`

HasTransferAcl returns a boolean if a field has been set.

### GetUpdateAcl

`func (o *View) GetUpdateAcl() []ACLItem`

GetUpdateAcl returns the UpdateAcl field if non-nil, zero value otherwise.

### GetUpdateAclOk

`func (o *View) GetUpdateAclOk() (*[]ACLItem, bool)`

GetUpdateAclOk returns a tuple with the UpdateAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAcl

`func (o *View) SetUpdateAcl(v []ACLItem)`

SetUpdateAcl sets UpdateAcl field to given value.

### HasUpdateAcl

`func (o *View) HasUpdateAcl() bool`

HasUpdateAcl returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *View) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *View) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *View) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *View) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUseForwardersForSubzones

`func (o *View) GetUseForwardersForSubzones() bool`

GetUseForwardersForSubzones returns the UseForwardersForSubzones field if non-nil, zero value otherwise.

### GetUseForwardersForSubzonesOk

`func (o *View) GetUseForwardersForSubzonesOk() (*bool, bool)`

GetUseForwardersForSubzonesOk returns a tuple with the UseForwardersForSubzones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForwardersForSubzones

`func (o *View) SetUseForwardersForSubzones(v bool)`

SetUseForwardersForSubzones sets UseForwardersForSubzones field to given value.

### HasUseForwardersForSubzones

`func (o *View) HasUseForwardersForSubzones() bool`

HasUseForwardersForSubzones returns a boolean if a field has been set.

### GetUseRootForwardersForLocalResolutionWithB1td

`func (o *View) GetUseRootForwardersForLocalResolutionWithB1td() bool`

GetUseRootForwardersForLocalResolutionWithB1td returns the UseRootForwardersForLocalResolutionWithB1td field if non-nil, zero value otherwise.

### GetUseRootForwardersForLocalResolutionWithB1tdOk

`func (o *View) GetUseRootForwardersForLocalResolutionWithB1tdOk() (*bool, bool)`

GetUseRootForwardersForLocalResolutionWithB1tdOk returns a tuple with the UseRootForwardersForLocalResolutionWithB1td field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRootForwardersForLocalResolutionWithB1td

`func (o *View) SetUseRootForwardersForLocalResolutionWithB1td(v bool)`

SetUseRootForwardersForLocalResolutionWithB1td sets UseRootForwardersForLocalResolutionWithB1td field to given value.

### HasUseRootForwardersForLocalResolutionWithB1td

`func (o *View) HasUseRootForwardersForLocalResolutionWithB1td() bool`

HasUseRootForwardersForLocalResolutionWithB1td returns a boolean if a field has been set.

### GetZoneAuthority

`func (o *View) GetZoneAuthority() ZoneAuthority`

GetZoneAuthority returns the ZoneAuthority field if non-nil, zero value otherwise.

### GetZoneAuthorityOk

`func (o *View) GetZoneAuthorityOk() (*ZoneAuthority, bool)`

GetZoneAuthorityOk returns a tuple with the ZoneAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAuthority

`func (o *View) SetZoneAuthority(v ZoneAuthority)`

SetZoneAuthority sets ZoneAuthority field to given value.

### HasZoneAuthority

`func (o *View) HasZoneAuthority() bool`

HasZoneAuthority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


