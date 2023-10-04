# ConfigGlobal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddEdnsOptionInOutgoingQuery** | Pointer to **bool** | _add_edns_option_in_outgoing_query_ adds client IP, MAC address and view name into outgoing recursive query. Defaults to _false_. | [optional] 
**CustomRootNs** | Pointer to [**[]ConfigRootNS**](ConfigRootNS.md) | Optional. List of custom root nameservers. The order does not matter.  Error if empty while _custom_root_ns_enabled_ is _true_. Error if there are duplicate items in the list.  Defaults to empty. | [optional] 
**CustomRootNsEnabled** | Pointer to **bool** | Optional. _true_ to use custom root nameservers instead of the default ones.  The _custom_root_ns_ field is validated when enabled.  Defaults to false. | [optional] 
**DnssecEnableValidation** | Pointer to **bool** | Optional. _true_ to perform DNSSEC validation. Ignored if _dnssec_enabled_ is _false_.  Defaults to _true_. | [optional] 
**DnssecEnabled** | Pointer to **bool** | Optional. Master toggle for all DNSSEC processing. Other _dnssec_*_ configuration is unused if this is disabled.  Defaults to _true_. | [optional] 
**DnssecRootKeys** | Pointer to [**[]ConfigTrustAnchor**](ConfigTrustAnchor.md) | DNSSEC root keys. The root keys are not configurable.  A default list is provided by cloud management and included here for config generation. | [optional] [readonly] 
**DnssecTrustAnchors** | Pointer to [**[]ConfigTrustAnchor**](ConfigTrustAnchor.md) | Optional. DNSSEC trust anchors.  Error if there are list items with duplicate (_zone_, _sep_, _algorithm_) combinations.  Defaults to empty. | [optional] 
**DnssecValidateExpiry** | Pointer to **bool** | Optional. _true_ to reject expired DNSSEC keys. Ignored if either _dnssec_enabled_ or _dnssec_enable_validation_ is _false_.  Defaults to _true_. | [optional] 
**EcsEnabled** | Pointer to **bool** | Optional. _true_ to enable EDNS client subnet for recursive queries. Other _ecs_*_ fields are ignored if this field is not enabled.  Defaults to _false_. | [optional] 
**EcsForwarding** | Pointer to **bool** | Optional. _true_ to enable ECS options in outbound queries. This functionality has additional overhead so it is disabled by default.  Defaults to _false_. | [optional] 
**EcsPrefixV4** | Pointer to **int64** | Optional. Maximum scope length for v4 ECS.  Unsigned integer, min 1 max 24.  Defaults to 24. | [optional] 
**EcsPrefixV6** | Pointer to **int64** | Optional. Maximum scope length for v6 ECS.  Unsigned integer, min 1 max 56.  Defaults to 56. | [optional] 
**EcsZones** | Pointer to [**[]ConfigECSZone**](ConfigECSZone.md) | Optional. List of zones where ECS queries may be sent.  Error if empty while _ecs_enabled_ is true. Error if there are duplicate FQDNs in the list.  Defaults to empty. | [optional] 
**EdnsUdpSize** | Pointer to **int64** | Optional. _edns_udp_size_ represents the edns UDP size. The size a querying DNS server advertises to the DNS server it’s sending a query to.  Defaults to 1232 bytes. | [optional] 
**FilterAaaaAcl** | Pointer to [**[]ConfigACLItem**](ConfigACLItem.md) | Optional. Specifies a list of client addresses for which AAAA filtering is to be applied.  Defaults to _empty_. | [optional] 
**FilterAaaaOnV4** | Pointer to **string** | _filter_aaaa_on_v4_ allows named to omit some IPv6 addresses when responding to IPv4 clients.  Allowed values: * _yes_, * _no_, * _break_dnssec_.  Defaults to _no_ | [optional] 
**Forwarders** | Pointer to [**[]ConfigForwarder**](ConfigForwarder.md) | Optional. List of forwarders.  Error if empty while _forwarders_only_ or _use_root_forwarders_for_local_resolution_with_b1td_ is _true_. Error if there are items in the list with duplicate addresses.  Defaults to empty. | [optional] 
**ForwardersOnly** | Pointer to **bool** | Optional. _true_ to only forward.  Defaults to _false_. | [optional] 
**GssTsigEnabled** | Pointer to **bool** | _gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.  Defaults to _false_. | [optional] 
**Id** | **string** | The resource identifier. | [readonly] 
**KerberosKeys** | Pointer to [**[]ConfigKerberosKey**](ConfigKerberosKey.md) | _kerberos_keys_ contains a list of keys for GSS-TSIG signed dynamic updates.  Defaults to empty. | [optional] 
**LameTtl** | Pointer to **int64** | Optional. Unused in the current on-prem DNS server implementation.  Unsigned integer, min 0 max 3600 (1h).  Defaults to 600. | [optional] 
**LogQueryResponse** | Pointer to **bool** | Optional. Control DNS query/response logging functionality.  Defaults to _true_. | [optional] 
**MatchRecursiveOnly** | Pointer to **bool** | Optional. If _true_ only recursive queries from matching clients access the view.  Defaults to _false_. | [optional] 
**MaxCacheTtl** | Pointer to **int64** | Optional. Seconds to cache positive responses.  Unsigned integer, min 1 max 604800 (7d).  Defaults to 604800 (7d). | [optional] 
**MaxNegativeTtl** | Pointer to **int64** | Optional. Seconds to cache negative responses.  Unsigned integer, min 1 max 604800 (7d).  Defaults to 10800 (3h). | [optional] 
**MaxUdpSize** | Pointer to **int64** | Optional. _max_udp_size_ represents maximum UDP payload size. The maximum number of bytes a responding DNS server will send to a UDP datagram.  Defaults to 1232 bytes. | [optional] 
**MinimalResponses** | Pointer to **bool** | Optional. When enabled, the DNS server will only add records to the authority and additional data sections when they are required.  Defaults to _false_. | [optional] 
**Notify** | Pointer to **bool** | _notify_ all external secondary DNS servers.  Defaults to _false_. | [optional] 
**QueryAcl** | Pointer to [**[]ConfigACLItem**](ConfigACLItem.md) | Optional. Clients must match this ACL to make authoritative queries. Also used for recursive queries if that ACL is unset.  Defaults to empty. | [optional] 
**QueryPort** | Pointer to **int64** | Optional. Source port for outbound DNS queries. When set to 0 the port is unspecified and the implementation may randomize it using any available ports.  Defaults to 0. | [optional] 
**RecursionAcl** | Pointer to [**[]ConfigACLItem**](ConfigACLItem.md) | Optional. Clients must match this ACL to make recursive queries. If this ACL is empty, then the _query_acl_ field will be used instead.  Defaults to empty. | [optional] 
**RecursionEnabled** | Pointer to **bool** | Optional. _true_ to allow recursive DNS queries.  Defaults to _true_. | [optional] 
**RecursiveClients** | Pointer to **int64** | Optional. Defines the number of simultaneous recursive lookups the server will perform on behalf of its clients.  Defaults to 1000. | [optional] 
**ResolverQueryTimeout** | Pointer to **int64** | Optional. Seconds before a recursive query times out.  Unsigned integer, min 10 max 30.  Defaults to 10. | [optional] 
**SecondaryAxfrQueryLimit** | Pointer to **int64** | Optional. Maximum concurrent inbound AXFRs. When set to 0 a host-dependent default will be used.  Defaults to 0. | [optional] 
**SecondarySoaQueryLimit** | Pointer to **int64** | Optional. Maximum concurrent outbound SOA queries. When set to 0 a host-dependent default will be used.  Defaults to 0. | [optional] 
**SortList** | Pointer to [**[]ConfigSortListItem**](ConfigSortListItem.md) | Optional. Specifies a sorted network list for A/AAAA records in DNS query response.  Defaults to _empty_. | [optional] 
**SynthesizeAddressRecordsFromHttps** | Pointer to **bool** | _synthesize_address_records_from_https_ enables/disables creation of A/AAAA records from HTTPS RR Defaults to _false_. | [optional] 
**TransferAcl** | Pointer to [**[]ConfigACLItem**](ConfigACLItem.md) | Optional. Clients must match this ACL to receive zone transfers.  Defaults to \&quot;deny any\&quot;. | [optional] 
**UpdateAcl** | Pointer to [**[]ConfigACLItem**](ConfigACLItem.md) | Optional. Specifies which hosts are allowed to issue Dynamic DNS updates for authoritative zones of _primary_type_ _cloud_.  Defaults to empty. | [optional] 
**UseForwardersForSubzones** | Pointer to **bool** | Optional. Use default forwarders to resolve queries for subzones.  Defaults to _true_. | [optional] 
**UseRootForwardersForLocalResolutionWithB1td** | Pointer to **bool** | _use_root_forwarders_for_local_resolution_with_b1td_ allows DNS recursive queries sent to root forwarders for local resolution when deployed alongside BloxOne Thread Defense. Defaults to _false_. | [optional] 
**ZoneAuthority** | Pointer to [**ConfigZoneAuthority**](ConfigZoneAuthority.md) |  | [optional] 

## Methods

### NewConfigGlobal

`func NewConfigGlobal(id string, ) *ConfigGlobal`

NewConfigGlobal instantiates a new ConfigGlobal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigGlobalWithDefaults

`func NewConfigGlobalWithDefaults() *ConfigGlobal`

NewConfigGlobalWithDefaults instantiates a new ConfigGlobal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddEdnsOptionInOutgoingQuery

`func (o *ConfigGlobal) GetAddEdnsOptionInOutgoingQuery() bool`

GetAddEdnsOptionInOutgoingQuery returns the AddEdnsOptionInOutgoingQuery field if non-nil, zero value otherwise.

### GetAddEdnsOptionInOutgoingQueryOk

`func (o *ConfigGlobal) GetAddEdnsOptionInOutgoingQueryOk() (*bool, bool)`

GetAddEdnsOptionInOutgoingQueryOk returns a tuple with the AddEdnsOptionInOutgoingQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddEdnsOptionInOutgoingQuery

`func (o *ConfigGlobal) SetAddEdnsOptionInOutgoingQuery(v bool)`

SetAddEdnsOptionInOutgoingQuery sets AddEdnsOptionInOutgoingQuery field to given value.

### HasAddEdnsOptionInOutgoingQuery

`func (o *ConfigGlobal) HasAddEdnsOptionInOutgoingQuery() bool`

HasAddEdnsOptionInOutgoingQuery returns a boolean if a field has been set.

### GetCustomRootNs

`func (o *ConfigGlobal) GetCustomRootNs() []ConfigRootNS`

GetCustomRootNs returns the CustomRootNs field if non-nil, zero value otherwise.

### GetCustomRootNsOk

`func (o *ConfigGlobal) GetCustomRootNsOk() (*[]ConfigRootNS, bool)`

GetCustomRootNsOk returns a tuple with the CustomRootNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRootNs

`func (o *ConfigGlobal) SetCustomRootNs(v []ConfigRootNS)`

SetCustomRootNs sets CustomRootNs field to given value.

### HasCustomRootNs

`func (o *ConfigGlobal) HasCustomRootNs() bool`

HasCustomRootNs returns a boolean if a field has been set.

### GetCustomRootNsEnabled

`func (o *ConfigGlobal) GetCustomRootNsEnabled() bool`

GetCustomRootNsEnabled returns the CustomRootNsEnabled field if non-nil, zero value otherwise.

### GetCustomRootNsEnabledOk

`func (o *ConfigGlobal) GetCustomRootNsEnabledOk() (*bool, bool)`

GetCustomRootNsEnabledOk returns a tuple with the CustomRootNsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRootNsEnabled

`func (o *ConfigGlobal) SetCustomRootNsEnabled(v bool)`

SetCustomRootNsEnabled sets CustomRootNsEnabled field to given value.

### HasCustomRootNsEnabled

`func (o *ConfigGlobal) HasCustomRootNsEnabled() bool`

HasCustomRootNsEnabled returns a boolean if a field has been set.

### GetDnssecEnableValidation

`func (o *ConfigGlobal) GetDnssecEnableValidation() bool`

GetDnssecEnableValidation returns the DnssecEnableValidation field if non-nil, zero value otherwise.

### GetDnssecEnableValidationOk

`func (o *ConfigGlobal) GetDnssecEnableValidationOk() (*bool, bool)`

GetDnssecEnableValidationOk returns a tuple with the DnssecEnableValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecEnableValidation

`func (o *ConfigGlobal) SetDnssecEnableValidation(v bool)`

SetDnssecEnableValidation sets DnssecEnableValidation field to given value.

### HasDnssecEnableValidation

`func (o *ConfigGlobal) HasDnssecEnableValidation() bool`

HasDnssecEnableValidation returns a boolean if a field has been set.

### GetDnssecEnabled

`func (o *ConfigGlobal) GetDnssecEnabled() bool`

GetDnssecEnabled returns the DnssecEnabled field if non-nil, zero value otherwise.

### GetDnssecEnabledOk

`func (o *ConfigGlobal) GetDnssecEnabledOk() (*bool, bool)`

GetDnssecEnabledOk returns a tuple with the DnssecEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecEnabled

`func (o *ConfigGlobal) SetDnssecEnabled(v bool)`

SetDnssecEnabled sets DnssecEnabled field to given value.

### HasDnssecEnabled

`func (o *ConfigGlobal) HasDnssecEnabled() bool`

HasDnssecEnabled returns a boolean if a field has been set.

### GetDnssecRootKeys

`func (o *ConfigGlobal) GetDnssecRootKeys() []ConfigTrustAnchor`

GetDnssecRootKeys returns the DnssecRootKeys field if non-nil, zero value otherwise.

### GetDnssecRootKeysOk

`func (o *ConfigGlobal) GetDnssecRootKeysOk() (*[]ConfigTrustAnchor, bool)`

GetDnssecRootKeysOk returns a tuple with the DnssecRootKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecRootKeys

`func (o *ConfigGlobal) SetDnssecRootKeys(v []ConfigTrustAnchor)`

SetDnssecRootKeys sets DnssecRootKeys field to given value.

### HasDnssecRootKeys

`func (o *ConfigGlobal) HasDnssecRootKeys() bool`

HasDnssecRootKeys returns a boolean if a field has been set.

### GetDnssecTrustAnchors

`func (o *ConfigGlobal) GetDnssecTrustAnchors() []ConfigTrustAnchor`

GetDnssecTrustAnchors returns the DnssecTrustAnchors field if non-nil, zero value otherwise.

### GetDnssecTrustAnchorsOk

`func (o *ConfigGlobal) GetDnssecTrustAnchorsOk() (*[]ConfigTrustAnchor, bool)`

GetDnssecTrustAnchorsOk returns a tuple with the DnssecTrustAnchors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecTrustAnchors

`func (o *ConfigGlobal) SetDnssecTrustAnchors(v []ConfigTrustAnchor)`

SetDnssecTrustAnchors sets DnssecTrustAnchors field to given value.

### HasDnssecTrustAnchors

`func (o *ConfigGlobal) HasDnssecTrustAnchors() bool`

HasDnssecTrustAnchors returns a boolean if a field has been set.

### GetDnssecValidateExpiry

`func (o *ConfigGlobal) GetDnssecValidateExpiry() bool`

GetDnssecValidateExpiry returns the DnssecValidateExpiry field if non-nil, zero value otherwise.

### GetDnssecValidateExpiryOk

`func (o *ConfigGlobal) GetDnssecValidateExpiryOk() (*bool, bool)`

GetDnssecValidateExpiryOk returns a tuple with the DnssecValidateExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecValidateExpiry

`func (o *ConfigGlobal) SetDnssecValidateExpiry(v bool)`

SetDnssecValidateExpiry sets DnssecValidateExpiry field to given value.

### HasDnssecValidateExpiry

`func (o *ConfigGlobal) HasDnssecValidateExpiry() bool`

HasDnssecValidateExpiry returns a boolean if a field has been set.

### GetEcsEnabled

`func (o *ConfigGlobal) GetEcsEnabled() bool`

GetEcsEnabled returns the EcsEnabled field if non-nil, zero value otherwise.

### GetEcsEnabledOk

`func (o *ConfigGlobal) GetEcsEnabledOk() (*bool, bool)`

GetEcsEnabledOk returns a tuple with the EcsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsEnabled

`func (o *ConfigGlobal) SetEcsEnabled(v bool)`

SetEcsEnabled sets EcsEnabled field to given value.

### HasEcsEnabled

`func (o *ConfigGlobal) HasEcsEnabled() bool`

HasEcsEnabled returns a boolean if a field has been set.

### GetEcsForwarding

`func (o *ConfigGlobal) GetEcsForwarding() bool`

GetEcsForwarding returns the EcsForwarding field if non-nil, zero value otherwise.

### GetEcsForwardingOk

`func (o *ConfigGlobal) GetEcsForwardingOk() (*bool, bool)`

GetEcsForwardingOk returns a tuple with the EcsForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsForwarding

`func (o *ConfigGlobal) SetEcsForwarding(v bool)`

SetEcsForwarding sets EcsForwarding field to given value.

### HasEcsForwarding

`func (o *ConfigGlobal) HasEcsForwarding() bool`

HasEcsForwarding returns a boolean if a field has been set.

### GetEcsPrefixV4

`func (o *ConfigGlobal) GetEcsPrefixV4() int64`

GetEcsPrefixV4 returns the EcsPrefixV4 field if non-nil, zero value otherwise.

### GetEcsPrefixV4Ok

`func (o *ConfigGlobal) GetEcsPrefixV4Ok() (*int64, bool)`

GetEcsPrefixV4Ok returns a tuple with the EcsPrefixV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsPrefixV4

`func (o *ConfigGlobal) SetEcsPrefixV4(v int64)`

SetEcsPrefixV4 sets EcsPrefixV4 field to given value.

### HasEcsPrefixV4

`func (o *ConfigGlobal) HasEcsPrefixV4() bool`

HasEcsPrefixV4 returns a boolean if a field has been set.

### GetEcsPrefixV6

`func (o *ConfigGlobal) GetEcsPrefixV6() int64`

GetEcsPrefixV6 returns the EcsPrefixV6 field if non-nil, zero value otherwise.

### GetEcsPrefixV6Ok

`func (o *ConfigGlobal) GetEcsPrefixV6Ok() (*int64, bool)`

GetEcsPrefixV6Ok returns a tuple with the EcsPrefixV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsPrefixV6

`func (o *ConfigGlobal) SetEcsPrefixV6(v int64)`

SetEcsPrefixV6 sets EcsPrefixV6 field to given value.

### HasEcsPrefixV6

`func (o *ConfigGlobal) HasEcsPrefixV6() bool`

HasEcsPrefixV6 returns a boolean if a field has been set.

### GetEcsZones

`func (o *ConfigGlobal) GetEcsZones() []ConfigECSZone`

GetEcsZones returns the EcsZones field if non-nil, zero value otherwise.

### GetEcsZonesOk

`func (o *ConfigGlobal) GetEcsZonesOk() (*[]ConfigECSZone, bool)`

GetEcsZonesOk returns a tuple with the EcsZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsZones

`func (o *ConfigGlobal) SetEcsZones(v []ConfigECSZone)`

SetEcsZones sets EcsZones field to given value.

### HasEcsZones

`func (o *ConfigGlobal) HasEcsZones() bool`

HasEcsZones returns a boolean if a field has been set.

### GetEdnsUdpSize

`func (o *ConfigGlobal) GetEdnsUdpSize() int64`

GetEdnsUdpSize returns the EdnsUdpSize field if non-nil, zero value otherwise.

### GetEdnsUdpSizeOk

`func (o *ConfigGlobal) GetEdnsUdpSizeOk() (*int64, bool)`

GetEdnsUdpSizeOk returns a tuple with the EdnsUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdnsUdpSize

`func (o *ConfigGlobal) SetEdnsUdpSize(v int64)`

SetEdnsUdpSize sets EdnsUdpSize field to given value.

### HasEdnsUdpSize

`func (o *ConfigGlobal) HasEdnsUdpSize() bool`

HasEdnsUdpSize returns a boolean if a field has been set.

### GetFilterAaaaAcl

`func (o *ConfigGlobal) GetFilterAaaaAcl() []ConfigACLItem`

GetFilterAaaaAcl returns the FilterAaaaAcl field if non-nil, zero value otherwise.

### GetFilterAaaaAclOk

`func (o *ConfigGlobal) GetFilterAaaaAclOk() (*[]ConfigACLItem, bool)`

GetFilterAaaaAclOk returns a tuple with the FilterAaaaAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaaAcl

`func (o *ConfigGlobal) SetFilterAaaaAcl(v []ConfigACLItem)`

SetFilterAaaaAcl sets FilterAaaaAcl field to given value.

### HasFilterAaaaAcl

`func (o *ConfigGlobal) HasFilterAaaaAcl() bool`

HasFilterAaaaAcl returns a boolean if a field has been set.

### GetFilterAaaaOnV4

`func (o *ConfigGlobal) GetFilterAaaaOnV4() string`

GetFilterAaaaOnV4 returns the FilterAaaaOnV4 field if non-nil, zero value otherwise.

### GetFilterAaaaOnV4Ok

`func (o *ConfigGlobal) GetFilterAaaaOnV4Ok() (*string, bool)`

GetFilterAaaaOnV4Ok returns a tuple with the FilterAaaaOnV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaaOnV4

`func (o *ConfigGlobal) SetFilterAaaaOnV4(v string)`

SetFilterAaaaOnV4 sets FilterAaaaOnV4 field to given value.

### HasFilterAaaaOnV4

`func (o *ConfigGlobal) HasFilterAaaaOnV4() bool`

HasFilterAaaaOnV4 returns a boolean if a field has been set.

### GetForwarders

`func (o *ConfigGlobal) GetForwarders() []ConfigForwarder`

GetForwarders returns the Forwarders field if non-nil, zero value otherwise.

### GetForwardersOk

`func (o *ConfigGlobal) GetForwardersOk() (*[]ConfigForwarder, bool)`

GetForwardersOk returns a tuple with the Forwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwarders

`func (o *ConfigGlobal) SetForwarders(v []ConfigForwarder)`

SetForwarders sets Forwarders field to given value.

### HasForwarders

`func (o *ConfigGlobal) HasForwarders() bool`

HasForwarders returns a boolean if a field has been set.

### GetForwardersOnly

`func (o *ConfigGlobal) GetForwardersOnly() bool`

GetForwardersOnly returns the ForwardersOnly field if non-nil, zero value otherwise.

### GetForwardersOnlyOk

`func (o *ConfigGlobal) GetForwardersOnlyOk() (*bool, bool)`

GetForwardersOnlyOk returns a tuple with the ForwardersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardersOnly

`func (o *ConfigGlobal) SetForwardersOnly(v bool)`

SetForwardersOnly sets ForwardersOnly field to given value.

### HasForwardersOnly

`func (o *ConfigGlobal) HasForwardersOnly() bool`

HasForwardersOnly returns a boolean if a field has been set.

### GetGssTsigEnabled

`func (o *ConfigGlobal) GetGssTsigEnabled() bool`

GetGssTsigEnabled returns the GssTsigEnabled field if non-nil, zero value otherwise.

### GetGssTsigEnabledOk

`func (o *ConfigGlobal) GetGssTsigEnabledOk() (*bool, bool)`

GetGssTsigEnabledOk returns a tuple with the GssTsigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigEnabled

`func (o *ConfigGlobal) SetGssTsigEnabled(v bool)`

SetGssTsigEnabled sets GssTsigEnabled field to given value.

### HasGssTsigEnabled

`func (o *ConfigGlobal) HasGssTsigEnabled() bool`

HasGssTsigEnabled returns a boolean if a field has been set.

### GetId

`func (o *ConfigGlobal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigGlobal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigGlobal) SetId(v string)`

SetId sets Id field to given value.


### GetKerberosKeys

`func (o *ConfigGlobal) GetKerberosKeys() []ConfigKerberosKey`

GetKerberosKeys returns the KerberosKeys field if non-nil, zero value otherwise.

### GetKerberosKeysOk

`func (o *ConfigGlobal) GetKerberosKeysOk() (*[]ConfigKerberosKey, bool)`

GetKerberosKeysOk returns a tuple with the KerberosKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosKeys

`func (o *ConfigGlobal) SetKerberosKeys(v []ConfigKerberosKey)`

SetKerberosKeys sets KerberosKeys field to given value.

### HasKerberosKeys

`func (o *ConfigGlobal) HasKerberosKeys() bool`

HasKerberosKeys returns a boolean if a field has been set.

### GetLameTtl

`func (o *ConfigGlobal) GetLameTtl() int64`

GetLameTtl returns the LameTtl field if non-nil, zero value otherwise.

### GetLameTtlOk

`func (o *ConfigGlobal) GetLameTtlOk() (*int64, bool)`

GetLameTtlOk returns a tuple with the LameTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLameTtl

`func (o *ConfigGlobal) SetLameTtl(v int64)`

SetLameTtl sets LameTtl field to given value.

### HasLameTtl

`func (o *ConfigGlobal) HasLameTtl() bool`

HasLameTtl returns a boolean if a field has been set.

### GetLogQueryResponse

`func (o *ConfigGlobal) GetLogQueryResponse() bool`

GetLogQueryResponse returns the LogQueryResponse field if non-nil, zero value otherwise.

### GetLogQueryResponseOk

`func (o *ConfigGlobal) GetLogQueryResponseOk() (*bool, bool)`

GetLogQueryResponseOk returns a tuple with the LogQueryResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogQueryResponse

`func (o *ConfigGlobal) SetLogQueryResponse(v bool)`

SetLogQueryResponse sets LogQueryResponse field to given value.

### HasLogQueryResponse

`func (o *ConfigGlobal) HasLogQueryResponse() bool`

HasLogQueryResponse returns a boolean if a field has been set.

### GetMatchRecursiveOnly

`func (o *ConfigGlobal) GetMatchRecursiveOnly() bool`

GetMatchRecursiveOnly returns the MatchRecursiveOnly field if non-nil, zero value otherwise.

### GetMatchRecursiveOnlyOk

`func (o *ConfigGlobal) GetMatchRecursiveOnlyOk() (*bool, bool)`

GetMatchRecursiveOnlyOk returns a tuple with the MatchRecursiveOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchRecursiveOnly

`func (o *ConfigGlobal) SetMatchRecursiveOnly(v bool)`

SetMatchRecursiveOnly sets MatchRecursiveOnly field to given value.

### HasMatchRecursiveOnly

`func (o *ConfigGlobal) HasMatchRecursiveOnly() bool`

HasMatchRecursiveOnly returns a boolean if a field has been set.

### GetMaxCacheTtl

`func (o *ConfigGlobal) GetMaxCacheTtl() int64`

GetMaxCacheTtl returns the MaxCacheTtl field if non-nil, zero value otherwise.

### GetMaxCacheTtlOk

`func (o *ConfigGlobal) GetMaxCacheTtlOk() (*int64, bool)`

GetMaxCacheTtlOk returns a tuple with the MaxCacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheTtl

`func (o *ConfigGlobal) SetMaxCacheTtl(v int64)`

SetMaxCacheTtl sets MaxCacheTtl field to given value.

### HasMaxCacheTtl

`func (o *ConfigGlobal) HasMaxCacheTtl() bool`

HasMaxCacheTtl returns a boolean if a field has been set.

### GetMaxNegativeTtl

`func (o *ConfigGlobal) GetMaxNegativeTtl() int64`

GetMaxNegativeTtl returns the MaxNegativeTtl field if non-nil, zero value otherwise.

### GetMaxNegativeTtlOk

`func (o *ConfigGlobal) GetMaxNegativeTtlOk() (*int64, bool)`

GetMaxNegativeTtlOk returns a tuple with the MaxNegativeTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNegativeTtl

`func (o *ConfigGlobal) SetMaxNegativeTtl(v int64)`

SetMaxNegativeTtl sets MaxNegativeTtl field to given value.

### HasMaxNegativeTtl

`func (o *ConfigGlobal) HasMaxNegativeTtl() bool`

HasMaxNegativeTtl returns a boolean if a field has been set.

### GetMaxUdpSize

`func (o *ConfigGlobal) GetMaxUdpSize() int64`

GetMaxUdpSize returns the MaxUdpSize field if non-nil, zero value otherwise.

### GetMaxUdpSizeOk

`func (o *ConfigGlobal) GetMaxUdpSizeOk() (*int64, bool)`

GetMaxUdpSizeOk returns a tuple with the MaxUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUdpSize

`func (o *ConfigGlobal) SetMaxUdpSize(v int64)`

SetMaxUdpSize sets MaxUdpSize field to given value.

### HasMaxUdpSize

`func (o *ConfigGlobal) HasMaxUdpSize() bool`

HasMaxUdpSize returns a boolean if a field has been set.

### GetMinimalResponses

`func (o *ConfigGlobal) GetMinimalResponses() bool`

GetMinimalResponses returns the MinimalResponses field if non-nil, zero value otherwise.

### GetMinimalResponsesOk

`func (o *ConfigGlobal) GetMinimalResponsesOk() (*bool, bool)`

GetMinimalResponsesOk returns a tuple with the MinimalResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimalResponses

`func (o *ConfigGlobal) SetMinimalResponses(v bool)`

SetMinimalResponses sets MinimalResponses field to given value.

### HasMinimalResponses

`func (o *ConfigGlobal) HasMinimalResponses() bool`

HasMinimalResponses returns a boolean if a field has been set.

### GetNotify

`func (o *ConfigGlobal) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *ConfigGlobal) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *ConfigGlobal) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *ConfigGlobal) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetQueryAcl

`func (o *ConfigGlobal) GetQueryAcl() []ConfigACLItem`

GetQueryAcl returns the QueryAcl field if non-nil, zero value otherwise.

### GetQueryAclOk

`func (o *ConfigGlobal) GetQueryAclOk() (*[]ConfigACLItem, bool)`

GetQueryAclOk returns a tuple with the QueryAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryAcl

`func (o *ConfigGlobal) SetQueryAcl(v []ConfigACLItem)`

SetQueryAcl sets QueryAcl field to given value.

### HasQueryAcl

`func (o *ConfigGlobal) HasQueryAcl() bool`

HasQueryAcl returns a boolean if a field has been set.

### GetQueryPort

`func (o *ConfigGlobal) GetQueryPort() int64`

GetQueryPort returns the QueryPort field if non-nil, zero value otherwise.

### GetQueryPortOk

`func (o *ConfigGlobal) GetQueryPortOk() (*int64, bool)`

GetQueryPortOk returns a tuple with the QueryPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryPort

`func (o *ConfigGlobal) SetQueryPort(v int64)`

SetQueryPort sets QueryPort field to given value.

### HasQueryPort

`func (o *ConfigGlobal) HasQueryPort() bool`

HasQueryPort returns a boolean if a field has been set.

### GetRecursionAcl

`func (o *ConfigGlobal) GetRecursionAcl() []ConfigACLItem`

GetRecursionAcl returns the RecursionAcl field if non-nil, zero value otherwise.

### GetRecursionAclOk

`func (o *ConfigGlobal) GetRecursionAclOk() (*[]ConfigACLItem, bool)`

GetRecursionAclOk returns a tuple with the RecursionAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursionAcl

`func (o *ConfigGlobal) SetRecursionAcl(v []ConfigACLItem)`

SetRecursionAcl sets RecursionAcl field to given value.

### HasRecursionAcl

`func (o *ConfigGlobal) HasRecursionAcl() bool`

HasRecursionAcl returns a boolean if a field has been set.

### GetRecursionEnabled

`func (o *ConfigGlobal) GetRecursionEnabled() bool`

GetRecursionEnabled returns the RecursionEnabled field if non-nil, zero value otherwise.

### GetRecursionEnabledOk

`func (o *ConfigGlobal) GetRecursionEnabledOk() (*bool, bool)`

GetRecursionEnabledOk returns a tuple with the RecursionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursionEnabled

`func (o *ConfigGlobal) SetRecursionEnabled(v bool)`

SetRecursionEnabled sets RecursionEnabled field to given value.

### HasRecursionEnabled

`func (o *ConfigGlobal) HasRecursionEnabled() bool`

HasRecursionEnabled returns a boolean if a field has been set.

### GetRecursiveClients

`func (o *ConfigGlobal) GetRecursiveClients() int64`

GetRecursiveClients returns the RecursiveClients field if non-nil, zero value otherwise.

### GetRecursiveClientsOk

`func (o *ConfigGlobal) GetRecursiveClientsOk() (*int64, bool)`

GetRecursiveClientsOk returns a tuple with the RecursiveClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursiveClients

`func (o *ConfigGlobal) SetRecursiveClients(v int64)`

SetRecursiveClients sets RecursiveClients field to given value.

### HasRecursiveClients

`func (o *ConfigGlobal) HasRecursiveClients() bool`

HasRecursiveClients returns a boolean if a field has been set.

### GetResolverQueryTimeout

`func (o *ConfigGlobal) GetResolverQueryTimeout() int64`

GetResolverQueryTimeout returns the ResolverQueryTimeout field if non-nil, zero value otherwise.

### GetResolverQueryTimeoutOk

`func (o *ConfigGlobal) GetResolverQueryTimeoutOk() (*int64, bool)`

GetResolverQueryTimeoutOk returns a tuple with the ResolverQueryTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolverQueryTimeout

`func (o *ConfigGlobal) SetResolverQueryTimeout(v int64)`

SetResolverQueryTimeout sets ResolverQueryTimeout field to given value.

### HasResolverQueryTimeout

`func (o *ConfigGlobal) HasResolverQueryTimeout() bool`

HasResolverQueryTimeout returns a boolean if a field has been set.

### GetSecondaryAxfrQueryLimit

`func (o *ConfigGlobal) GetSecondaryAxfrQueryLimit() int64`

GetSecondaryAxfrQueryLimit returns the SecondaryAxfrQueryLimit field if non-nil, zero value otherwise.

### GetSecondaryAxfrQueryLimitOk

`func (o *ConfigGlobal) GetSecondaryAxfrQueryLimitOk() (*int64, bool)`

GetSecondaryAxfrQueryLimitOk returns a tuple with the SecondaryAxfrQueryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAxfrQueryLimit

`func (o *ConfigGlobal) SetSecondaryAxfrQueryLimit(v int64)`

SetSecondaryAxfrQueryLimit sets SecondaryAxfrQueryLimit field to given value.

### HasSecondaryAxfrQueryLimit

`func (o *ConfigGlobal) HasSecondaryAxfrQueryLimit() bool`

HasSecondaryAxfrQueryLimit returns a boolean if a field has been set.

### GetSecondarySoaQueryLimit

`func (o *ConfigGlobal) GetSecondarySoaQueryLimit() int64`

GetSecondarySoaQueryLimit returns the SecondarySoaQueryLimit field if non-nil, zero value otherwise.

### GetSecondarySoaQueryLimitOk

`func (o *ConfigGlobal) GetSecondarySoaQueryLimitOk() (*int64, bool)`

GetSecondarySoaQueryLimitOk returns a tuple with the SecondarySoaQueryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySoaQueryLimit

`func (o *ConfigGlobal) SetSecondarySoaQueryLimit(v int64)`

SetSecondarySoaQueryLimit sets SecondarySoaQueryLimit field to given value.

### HasSecondarySoaQueryLimit

`func (o *ConfigGlobal) HasSecondarySoaQueryLimit() bool`

HasSecondarySoaQueryLimit returns a boolean if a field has been set.

### GetSortList

`func (o *ConfigGlobal) GetSortList() []ConfigSortListItem`

GetSortList returns the SortList field if non-nil, zero value otherwise.

### GetSortListOk

`func (o *ConfigGlobal) GetSortListOk() (*[]ConfigSortListItem, bool)`

GetSortListOk returns a tuple with the SortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortList

`func (o *ConfigGlobal) SetSortList(v []ConfigSortListItem)`

SetSortList sets SortList field to given value.

### HasSortList

`func (o *ConfigGlobal) HasSortList() bool`

HasSortList returns a boolean if a field has been set.

### GetSynthesizeAddressRecordsFromHttps

`func (o *ConfigGlobal) GetSynthesizeAddressRecordsFromHttps() bool`

GetSynthesizeAddressRecordsFromHttps returns the SynthesizeAddressRecordsFromHttps field if non-nil, zero value otherwise.

### GetSynthesizeAddressRecordsFromHttpsOk

`func (o *ConfigGlobal) GetSynthesizeAddressRecordsFromHttpsOk() (*bool, bool)`

GetSynthesizeAddressRecordsFromHttpsOk returns a tuple with the SynthesizeAddressRecordsFromHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthesizeAddressRecordsFromHttps

`func (o *ConfigGlobal) SetSynthesizeAddressRecordsFromHttps(v bool)`

SetSynthesizeAddressRecordsFromHttps sets SynthesizeAddressRecordsFromHttps field to given value.

### HasSynthesizeAddressRecordsFromHttps

`func (o *ConfigGlobal) HasSynthesizeAddressRecordsFromHttps() bool`

HasSynthesizeAddressRecordsFromHttps returns a boolean if a field has been set.

### GetTransferAcl

`func (o *ConfigGlobal) GetTransferAcl() []ConfigACLItem`

GetTransferAcl returns the TransferAcl field if non-nil, zero value otherwise.

### GetTransferAclOk

`func (o *ConfigGlobal) GetTransferAclOk() (*[]ConfigACLItem, bool)`

GetTransferAclOk returns a tuple with the TransferAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAcl

`func (o *ConfigGlobal) SetTransferAcl(v []ConfigACLItem)`

SetTransferAcl sets TransferAcl field to given value.

### HasTransferAcl

`func (o *ConfigGlobal) HasTransferAcl() bool`

HasTransferAcl returns a boolean if a field has been set.

### GetUpdateAcl

`func (o *ConfigGlobal) GetUpdateAcl() []ConfigACLItem`

GetUpdateAcl returns the UpdateAcl field if non-nil, zero value otherwise.

### GetUpdateAclOk

`func (o *ConfigGlobal) GetUpdateAclOk() (*[]ConfigACLItem, bool)`

GetUpdateAclOk returns a tuple with the UpdateAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAcl

`func (o *ConfigGlobal) SetUpdateAcl(v []ConfigACLItem)`

SetUpdateAcl sets UpdateAcl field to given value.

### HasUpdateAcl

`func (o *ConfigGlobal) HasUpdateAcl() bool`

HasUpdateAcl returns a boolean if a field has been set.

### GetUseForwardersForSubzones

`func (o *ConfigGlobal) GetUseForwardersForSubzones() bool`

GetUseForwardersForSubzones returns the UseForwardersForSubzones field if non-nil, zero value otherwise.

### GetUseForwardersForSubzonesOk

`func (o *ConfigGlobal) GetUseForwardersForSubzonesOk() (*bool, bool)`

GetUseForwardersForSubzonesOk returns a tuple with the UseForwardersForSubzones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForwardersForSubzones

`func (o *ConfigGlobal) SetUseForwardersForSubzones(v bool)`

SetUseForwardersForSubzones sets UseForwardersForSubzones field to given value.

### HasUseForwardersForSubzones

`func (o *ConfigGlobal) HasUseForwardersForSubzones() bool`

HasUseForwardersForSubzones returns a boolean if a field has been set.

### GetUseRootForwardersForLocalResolutionWithB1td

`func (o *ConfigGlobal) GetUseRootForwardersForLocalResolutionWithB1td() bool`

GetUseRootForwardersForLocalResolutionWithB1td returns the UseRootForwardersForLocalResolutionWithB1td field if non-nil, zero value otherwise.

### GetUseRootForwardersForLocalResolutionWithB1tdOk

`func (o *ConfigGlobal) GetUseRootForwardersForLocalResolutionWithB1tdOk() (*bool, bool)`

GetUseRootForwardersForLocalResolutionWithB1tdOk returns a tuple with the UseRootForwardersForLocalResolutionWithB1td field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRootForwardersForLocalResolutionWithB1td

`func (o *ConfigGlobal) SetUseRootForwardersForLocalResolutionWithB1td(v bool)`

SetUseRootForwardersForLocalResolutionWithB1td sets UseRootForwardersForLocalResolutionWithB1td field to given value.

### HasUseRootForwardersForLocalResolutionWithB1td

`func (o *ConfigGlobal) HasUseRootForwardersForLocalResolutionWithB1td() bool`

HasUseRootForwardersForLocalResolutionWithB1td returns a boolean if a field has been set.

### GetZoneAuthority

`func (o *ConfigGlobal) GetZoneAuthority() ConfigZoneAuthority`

GetZoneAuthority returns the ZoneAuthority field if non-nil, zero value otherwise.

### GetZoneAuthorityOk

`func (o *ConfigGlobal) GetZoneAuthorityOk() (*ConfigZoneAuthority, bool)`

GetZoneAuthorityOk returns a tuple with the ZoneAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAuthority

`func (o *ConfigGlobal) SetZoneAuthority(v ConfigZoneAuthority)`

SetZoneAuthority sets ZoneAuthority field to given value.

### HasZoneAuthority

`func (o *ConfigGlobal) HasZoneAuthority() bool`

HasZoneAuthority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


