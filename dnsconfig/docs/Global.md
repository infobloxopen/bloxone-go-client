# Global

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddEdnsOptionInOutgoingQuery** | Pointer to **bool** | _add_edns_option_in_outgoing_query_ adds client IP, MAC address and view name into outgoing recursive query. Defaults to _false_. | [optional] 
**CustomRootNs** | Pointer to [**[]RootNS**](RootNS.md) | Optional. List of custom root nameservers. The order does not matter.  Error if empty while _custom_root_ns_enabled_ is _true_. Error if there are duplicate items in the list.  Defaults to empty. | [optional] 
**CustomRootNsEnabled** | Pointer to **bool** | Optional. _true_ to use custom root nameservers instead of the default ones.  The _custom_root_ns_ field is validated when enabled.  Defaults to false. | [optional] 
**DnssecEnableValidation** | Pointer to **bool** | Optional. _true_ to perform DNSSEC validation. Ignored if _dnssec_enabled_ is _false_.  Defaults to _true_. | [optional] 
**DnssecEnabled** | Pointer to **bool** | Optional. Master toggle for all DNSSEC processing. Other _dnssec_*_ configuration is unused if this is disabled.  Defaults to _true_. | [optional] 
**DnssecRootKeys** | Pointer to [**[]TrustAnchor**](TrustAnchor.md) | DNSSEC root keys. The root keys are not configurable.  A default list is provided by cloud management and included here for config generation. | [optional] [readonly] 
**DnssecTrustAnchors** | Pointer to [**[]TrustAnchor**](TrustAnchor.md) | Optional. DNSSEC trust anchors.  Error if there are list items with duplicate (_zone_, _sep_, _algorithm_) combinations.  Defaults to empty. | [optional] 
**DnssecValidateExpiry** | Pointer to **bool** | Optional. _true_ to reject expired DNSSEC keys. Ignored if either _dnssec_enabled_ or _dnssec_enable_validation_ is _false_.  Defaults to _true_. | [optional] 
**DtcConfig** | Pointer to [**DTCConfig**](DTCConfig.md) |  | [optional] 
**EcsEnabled** | Pointer to **bool** | Optional. _true_ to enable EDNS client subnet for recursive queries. Other _ecs_*_ fields are ignored if this field is not enabled.  Defaults to _false_. | [optional] 
**EcsForwarding** | Pointer to **bool** | Optional. _true_ to enable ECS options in outbound queries. This functionality has additional overhead so it is disabled by default.  Defaults to _false_. | [optional] 
**EcsPrefixV4** | Pointer to **int64** | Optional. Maximum scope length for v4 ECS.  Unsigned integer, min 1 max 24.  Defaults to 24. | [optional] 
**EcsPrefixV6** | Pointer to **int64** | Optional. Maximum scope length for v6 ECS.  Unsigned integer, min 1 max 56.  Defaults to 56. | [optional] 
**EcsZones** | Pointer to [**[]ECSZone**](ECSZone.md) | Optional. List of zones where ECS queries may be sent.  Error if empty while _ecs_enabled_ is true. Error if there are duplicate FQDNs in the list.  Defaults to empty. | [optional] 
**EdnsUdpSize** | Pointer to **int64** | Optional. _edns_udp_size_ represents the edns UDP size. The size a querying DNS server advertises to the DNS server it’s sending a query to.  Defaults to 1232 bytes. | [optional] 
**FilterAaaaAcl** | Pointer to [**[]ACLItem**](ACLItem.md) | Optional. Specifies a list of client addresses for which AAAA filtering is to be applied.  Defaults to _empty_. | [optional] 
**FilterAaaaOnV4** | Pointer to **string** | _filter_aaaa_on_v4_ allows named to omit some IPv6 addresses when responding to IPv4 clients.  Allowed values: * _yes_, * _no_, * _break_dnssec_.  Defaults to _no_ | [optional] 
**Forwarders** | Pointer to [**[]Forwarder**](Forwarder.md) | Optional. List of forwarders.  Error if empty while _forwarders_only_ or _use_root_forwarders_for_local_resolution_with_b1td_ is _true_. Error if there are items in the list with duplicate addresses.  Defaults to empty. | [optional] 
**ForwardersOnly** | Pointer to **bool** | Optional. _true_ to only forward.  Defaults to _false_. | [optional] 
**GssTsigEnabled** | Pointer to **bool** | _gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.  Defaults to _false_. | [optional] 
**Id** | **string** | The resource identifier. | [readonly] 
**KerberosKeys** | Pointer to [**[]KerberosKey**](KerberosKey.md) | _kerberos_keys_ contains a list of keys for GSS-TSIG signed dynamic updates.  Defaults to empty. | [optional] 
**LameTtl** | Pointer to **int64** | Optional. Unused in the current on-prem DNS server implementation.  Unsigned integer, min 0 max 3600 (1h).  Defaults to 600. | [optional] 
**LogQueryResponse** | Pointer to **bool** | Optional. Control DNS query/response logging functionality.  Defaults to _true_. | [optional] 
**MatchRecursiveOnly** | Pointer to **bool** | Optional. If _true_ only recursive queries from matching clients access the view.  Defaults to _false_. | [optional] 
**MaxCacheTtl** | Pointer to **int64** | Optional. Seconds to cache positive responses.  Unsigned integer, min 1 max 604800 (7d).  Defaults to 604800 (7d). | [optional] 
**MaxNegativeTtl** | Pointer to **int64** | Optional. Seconds to cache negative responses.  Unsigned integer, min 1 max 604800 (7d).  Defaults to 10800 (3h). | [optional] 
**MaxUdpSize** | Pointer to **int64** | Optional. _max_udp_size_ represents maximum UDP payload size. The maximum number of bytes a responding DNS server will send to a UDP datagram.  Defaults to 1232 bytes. | [optional] 
**MinimalResponses** | Pointer to **bool** | Optional. When enabled, the DNS server will only add records to the authority and additional data sections when they are required.  Defaults to _false_. | [optional] 
**Notify** | Pointer to **bool** | _notify_ all external secondary DNS servers.  Defaults to _false_. | [optional] 
**QueryAcl** | Pointer to [**[]ACLItem**](ACLItem.md) | Optional. Clients must match this ACL to make authoritative queries. Also used for recursive queries if that ACL is unset.  Defaults to empty. | [optional] 
**QueryPort** | Pointer to **int64** | Optional. Source port for outbound DNS queries. When set to 0 the port is unspecified and the implementation may randomize it using any available ports.  Defaults to 0. | [optional] 
**RecursionAcl** | Pointer to [**[]ACLItem**](ACLItem.md) | Optional. Clients must match this ACL to make recursive queries. If this ACL is empty, then the _query_acl_ field will be used instead.  Defaults to empty. | [optional] 
**RecursionEnabled** | Pointer to **bool** | Optional. _true_ to allow recursive DNS queries.  Defaults to _true_. | [optional] 
**RecursiveClients** | Pointer to **int64** | Optional. Defines the number of simultaneous recursive lookups the server will perform on behalf of its clients.  Defaults to 1000. | [optional] 
**ResolverQueryTimeout** | Pointer to **int64** | Optional. Seconds before a recursive query times out.  Unsigned integer, min 10 max 30.  Defaults to 10. | [optional] 
**SecondaryAxfrQueryLimit** | Pointer to **int64** | Optional. Maximum concurrent inbound AXFRs. When set to 0 a host-dependent default will be used.  Defaults to 0. | [optional] 
**SecondarySoaQueryLimit** | Pointer to **int64** | Optional. Maximum concurrent outbound SOA queries. When set to 0 a host-dependent default will be used.  Defaults to 0. | [optional] 
**SortList** | Pointer to [**[]SortListItem**](SortListItem.md) | Optional. Specifies a sorted network list for A/AAAA records in DNS query response.  Defaults to _empty_. | [optional] 
**SynthesizeAddressRecordsFromHttps** | Pointer to **bool** | _synthesize_address_records_from_https_ enables/disables creation of A/AAAA records from HTTPS RR Defaults to _false_. | [optional] 
**TransferAcl** | Pointer to [**[]ACLItem**](ACLItem.md) | Optional. Clients must match this ACL to receive zone transfers.  Defaults to \&quot;deny any\&quot;. | [optional] 
**UpdateAcl** | Pointer to [**[]ACLItem**](ACLItem.md) | Optional. Specifies which hosts are allowed to issue Dynamic DNS updates for authoritative zones of _primary_type_ _cloud_.  Defaults to empty. | [optional] 
**UseForwardersForSubzones** | Pointer to **bool** | Optional. Use default forwarders to resolve queries for subzones.  Defaults to _true_. | [optional] 
**UseRootForwardersForLocalResolutionWithB1td** | Pointer to **bool** | _use_root_forwarders_for_local_resolution_with_b1td_ allows DNS recursive queries sent to root forwarders for local resolution when deployed alongside BloxOne Thread Defense. Defaults to _false_. | [optional] 
**ZoneAuthority** | Pointer to [**ZoneAuthority**](ZoneAuthority.md) |  | [optional] 

## Methods

### NewGlobal

`func NewGlobal(id string, ) *Global`

NewGlobal instantiates a new Global object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalWithDefaults

`func NewGlobalWithDefaults() *Global`

NewGlobalWithDefaults instantiates a new Global object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddEdnsOptionInOutgoingQuery

`func (o *Global) GetAddEdnsOptionInOutgoingQuery() bool`

GetAddEdnsOptionInOutgoingQuery returns the AddEdnsOptionInOutgoingQuery field if non-nil, zero value otherwise.

### GetAddEdnsOptionInOutgoingQueryOk

`func (o *Global) GetAddEdnsOptionInOutgoingQueryOk() (*bool, bool)`

GetAddEdnsOptionInOutgoingQueryOk returns a tuple with the AddEdnsOptionInOutgoingQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddEdnsOptionInOutgoingQuery

`func (o *Global) SetAddEdnsOptionInOutgoingQuery(v bool)`

SetAddEdnsOptionInOutgoingQuery sets AddEdnsOptionInOutgoingQuery field to given value.

### HasAddEdnsOptionInOutgoingQuery

`func (o *Global) HasAddEdnsOptionInOutgoingQuery() bool`

HasAddEdnsOptionInOutgoingQuery returns a boolean if a field has been set.

### GetCustomRootNs

`func (o *Global) GetCustomRootNs() []RootNS`

GetCustomRootNs returns the CustomRootNs field if non-nil, zero value otherwise.

### GetCustomRootNsOk

`func (o *Global) GetCustomRootNsOk() (*[]RootNS, bool)`

GetCustomRootNsOk returns a tuple with the CustomRootNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRootNs

`func (o *Global) SetCustomRootNs(v []RootNS)`

SetCustomRootNs sets CustomRootNs field to given value.

### HasCustomRootNs

`func (o *Global) HasCustomRootNs() bool`

HasCustomRootNs returns a boolean if a field has been set.

### GetCustomRootNsEnabled

`func (o *Global) GetCustomRootNsEnabled() bool`

GetCustomRootNsEnabled returns the CustomRootNsEnabled field if non-nil, zero value otherwise.

### GetCustomRootNsEnabledOk

`func (o *Global) GetCustomRootNsEnabledOk() (*bool, bool)`

GetCustomRootNsEnabledOk returns a tuple with the CustomRootNsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRootNsEnabled

`func (o *Global) SetCustomRootNsEnabled(v bool)`

SetCustomRootNsEnabled sets CustomRootNsEnabled field to given value.

### HasCustomRootNsEnabled

`func (o *Global) HasCustomRootNsEnabled() bool`

HasCustomRootNsEnabled returns a boolean if a field has been set.

### GetDnssecEnableValidation

`func (o *Global) GetDnssecEnableValidation() bool`

GetDnssecEnableValidation returns the DnssecEnableValidation field if non-nil, zero value otherwise.

### GetDnssecEnableValidationOk

`func (o *Global) GetDnssecEnableValidationOk() (*bool, bool)`

GetDnssecEnableValidationOk returns a tuple with the DnssecEnableValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecEnableValidation

`func (o *Global) SetDnssecEnableValidation(v bool)`

SetDnssecEnableValidation sets DnssecEnableValidation field to given value.

### HasDnssecEnableValidation

`func (o *Global) HasDnssecEnableValidation() bool`

HasDnssecEnableValidation returns a boolean if a field has been set.

### GetDnssecEnabled

`func (o *Global) GetDnssecEnabled() bool`

GetDnssecEnabled returns the DnssecEnabled field if non-nil, zero value otherwise.

### GetDnssecEnabledOk

`func (o *Global) GetDnssecEnabledOk() (*bool, bool)`

GetDnssecEnabledOk returns a tuple with the DnssecEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecEnabled

`func (o *Global) SetDnssecEnabled(v bool)`

SetDnssecEnabled sets DnssecEnabled field to given value.

### HasDnssecEnabled

`func (o *Global) HasDnssecEnabled() bool`

HasDnssecEnabled returns a boolean if a field has been set.

### GetDnssecRootKeys

`func (o *Global) GetDnssecRootKeys() []TrustAnchor`

GetDnssecRootKeys returns the DnssecRootKeys field if non-nil, zero value otherwise.

### GetDnssecRootKeysOk

`func (o *Global) GetDnssecRootKeysOk() (*[]TrustAnchor, bool)`

GetDnssecRootKeysOk returns a tuple with the DnssecRootKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecRootKeys

`func (o *Global) SetDnssecRootKeys(v []TrustAnchor)`

SetDnssecRootKeys sets DnssecRootKeys field to given value.

### HasDnssecRootKeys

`func (o *Global) HasDnssecRootKeys() bool`

HasDnssecRootKeys returns a boolean if a field has been set.

### GetDnssecTrustAnchors

`func (o *Global) GetDnssecTrustAnchors() []TrustAnchor`

GetDnssecTrustAnchors returns the DnssecTrustAnchors field if non-nil, zero value otherwise.

### GetDnssecTrustAnchorsOk

`func (o *Global) GetDnssecTrustAnchorsOk() (*[]TrustAnchor, bool)`

GetDnssecTrustAnchorsOk returns a tuple with the DnssecTrustAnchors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecTrustAnchors

`func (o *Global) SetDnssecTrustAnchors(v []TrustAnchor)`

SetDnssecTrustAnchors sets DnssecTrustAnchors field to given value.

### HasDnssecTrustAnchors

`func (o *Global) HasDnssecTrustAnchors() bool`

HasDnssecTrustAnchors returns a boolean if a field has been set.

### GetDnssecValidateExpiry

`func (o *Global) GetDnssecValidateExpiry() bool`

GetDnssecValidateExpiry returns the DnssecValidateExpiry field if non-nil, zero value otherwise.

### GetDnssecValidateExpiryOk

`func (o *Global) GetDnssecValidateExpiryOk() (*bool, bool)`

GetDnssecValidateExpiryOk returns a tuple with the DnssecValidateExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecValidateExpiry

`func (o *Global) SetDnssecValidateExpiry(v bool)`

SetDnssecValidateExpiry sets DnssecValidateExpiry field to given value.

### HasDnssecValidateExpiry

`func (o *Global) HasDnssecValidateExpiry() bool`

HasDnssecValidateExpiry returns a boolean if a field has been set.

### GetDtcConfig

`func (o *Global) GetDtcConfig() DTCConfig`

GetDtcConfig returns the DtcConfig field if non-nil, zero value otherwise.

### GetDtcConfigOk

`func (o *Global) GetDtcConfigOk() (*DTCConfig, bool)`

GetDtcConfigOk returns a tuple with the DtcConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcConfig

`func (o *Global) SetDtcConfig(v DTCConfig)`

SetDtcConfig sets DtcConfig field to given value.

### HasDtcConfig

`func (o *Global) HasDtcConfig() bool`

HasDtcConfig returns a boolean if a field has been set.

### GetEcsEnabled

`func (o *Global) GetEcsEnabled() bool`

GetEcsEnabled returns the EcsEnabled field if non-nil, zero value otherwise.

### GetEcsEnabledOk

`func (o *Global) GetEcsEnabledOk() (*bool, bool)`

GetEcsEnabledOk returns a tuple with the EcsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsEnabled

`func (o *Global) SetEcsEnabled(v bool)`

SetEcsEnabled sets EcsEnabled field to given value.

### HasEcsEnabled

`func (o *Global) HasEcsEnabled() bool`

HasEcsEnabled returns a boolean if a field has been set.

### GetEcsForwarding

`func (o *Global) GetEcsForwarding() bool`

GetEcsForwarding returns the EcsForwarding field if non-nil, zero value otherwise.

### GetEcsForwardingOk

`func (o *Global) GetEcsForwardingOk() (*bool, bool)`

GetEcsForwardingOk returns a tuple with the EcsForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsForwarding

`func (o *Global) SetEcsForwarding(v bool)`

SetEcsForwarding sets EcsForwarding field to given value.

### HasEcsForwarding

`func (o *Global) HasEcsForwarding() bool`

HasEcsForwarding returns a boolean if a field has been set.

### GetEcsPrefixV4

`func (o *Global) GetEcsPrefixV4() int64`

GetEcsPrefixV4 returns the EcsPrefixV4 field if non-nil, zero value otherwise.

### GetEcsPrefixV4Ok

`func (o *Global) GetEcsPrefixV4Ok() (*int64, bool)`

GetEcsPrefixV4Ok returns a tuple with the EcsPrefixV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsPrefixV4

`func (o *Global) SetEcsPrefixV4(v int64)`

SetEcsPrefixV4 sets EcsPrefixV4 field to given value.

### HasEcsPrefixV4

`func (o *Global) HasEcsPrefixV4() bool`

HasEcsPrefixV4 returns a boolean if a field has been set.

### GetEcsPrefixV6

`func (o *Global) GetEcsPrefixV6() int64`

GetEcsPrefixV6 returns the EcsPrefixV6 field if non-nil, zero value otherwise.

### GetEcsPrefixV6Ok

`func (o *Global) GetEcsPrefixV6Ok() (*int64, bool)`

GetEcsPrefixV6Ok returns a tuple with the EcsPrefixV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsPrefixV6

`func (o *Global) SetEcsPrefixV6(v int64)`

SetEcsPrefixV6 sets EcsPrefixV6 field to given value.

### HasEcsPrefixV6

`func (o *Global) HasEcsPrefixV6() bool`

HasEcsPrefixV6 returns a boolean if a field has been set.

### GetEcsZones

`func (o *Global) GetEcsZones() []ECSZone`

GetEcsZones returns the EcsZones field if non-nil, zero value otherwise.

### GetEcsZonesOk

`func (o *Global) GetEcsZonesOk() (*[]ECSZone, bool)`

GetEcsZonesOk returns a tuple with the EcsZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsZones

`func (o *Global) SetEcsZones(v []ECSZone)`

SetEcsZones sets EcsZones field to given value.

### HasEcsZones

`func (o *Global) HasEcsZones() bool`

HasEcsZones returns a boolean if a field has been set.

### GetEdnsUdpSize

`func (o *Global) GetEdnsUdpSize() int64`

GetEdnsUdpSize returns the EdnsUdpSize field if non-nil, zero value otherwise.

### GetEdnsUdpSizeOk

`func (o *Global) GetEdnsUdpSizeOk() (*int64, bool)`

GetEdnsUdpSizeOk returns a tuple with the EdnsUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdnsUdpSize

`func (o *Global) SetEdnsUdpSize(v int64)`

SetEdnsUdpSize sets EdnsUdpSize field to given value.

### HasEdnsUdpSize

`func (o *Global) HasEdnsUdpSize() bool`

HasEdnsUdpSize returns a boolean if a field has been set.

### GetFilterAaaaAcl

`func (o *Global) GetFilterAaaaAcl() []ACLItem`

GetFilterAaaaAcl returns the FilterAaaaAcl field if non-nil, zero value otherwise.

### GetFilterAaaaAclOk

`func (o *Global) GetFilterAaaaAclOk() (*[]ACLItem, bool)`

GetFilterAaaaAclOk returns a tuple with the FilterAaaaAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaaAcl

`func (o *Global) SetFilterAaaaAcl(v []ACLItem)`

SetFilterAaaaAcl sets FilterAaaaAcl field to given value.

### HasFilterAaaaAcl

`func (o *Global) HasFilterAaaaAcl() bool`

HasFilterAaaaAcl returns a boolean if a field has been set.

### GetFilterAaaaOnV4

`func (o *Global) GetFilterAaaaOnV4() string`

GetFilterAaaaOnV4 returns the FilterAaaaOnV4 field if non-nil, zero value otherwise.

### GetFilterAaaaOnV4Ok

`func (o *Global) GetFilterAaaaOnV4Ok() (*string, bool)`

GetFilterAaaaOnV4Ok returns a tuple with the FilterAaaaOnV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaaOnV4

`func (o *Global) SetFilterAaaaOnV4(v string)`

SetFilterAaaaOnV4 sets FilterAaaaOnV4 field to given value.

### HasFilterAaaaOnV4

`func (o *Global) HasFilterAaaaOnV4() bool`

HasFilterAaaaOnV4 returns a boolean if a field has been set.

### GetForwarders

`func (o *Global) GetForwarders() []Forwarder`

GetForwarders returns the Forwarders field if non-nil, zero value otherwise.

### GetForwardersOk

`func (o *Global) GetForwardersOk() (*[]Forwarder, bool)`

GetForwardersOk returns a tuple with the Forwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwarders

`func (o *Global) SetForwarders(v []Forwarder)`

SetForwarders sets Forwarders field to given value.

### HasForwarders

`func (o *Global) HasForwarders() bool`

HasForwarders returns a boolean if a field has been set.

### GetForwardersOnly

`func (o *Global) GetForwardersOnly() bool`

GetForwardersOnly returns the ForwardersOnly field if non-nil, zero value otherwise.

### GetForwardersOnlyOk

`func (o *Global) GetForwardersOnlyOk() (*bool, bool)`

GetForwardersOnlyOk returns a tuple with the ForwardersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardersOnly

`func (o *Global) SetForwardersOnly(v bool)`

SetForwardersOnly sets ForwardersOnly field to given value.

### HasForwardersOnly

`func (o *Global) HasForwardersOnly() bool`

HasForwardersOnly returns a boolean if a field has been set.

### GetGssTsigEnabled

`func (o *Global) GetGssTsigEnabled() bool`

GetGssTsigEnabled returns the GssTsigEnabled field if non-nil, zero value otherwise.

### GetGssTsigEnabledOk

`func (o *Global) GetGssTsigEnabledOk() (*bool, bool)`

GetGssTsigEnabledOk returns a tuple with the GssTsigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigEnabled

`func (o *Global) SetGssTsigEnabled(v bool)`

SetGssTsigEnabled sets GssTsigEnabled field to given value.

### HasGssTsigEnabled

`func (o *Global) HasGssTsigEnabled() bool`

HasGssTsigEnabled returns a boolean if a field has been set.

### GetId

`func (o *Global) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Global) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Global) SetId(v string)`

SetId sets Id field to given value.


### GetKerberosKeys

`func (o *Global) GetKerberosKeys() []KerberosKey`

GetKerberosKeys returns the KerberosKeys field if non-nil, zero value otherwise.

### GetKerberosKeysOk

`func (o *Global) GetKerberosKeysOk() (*[]KerberosKey, bool)`

GetKerberosKeysOk returns a tuple with the KerberosKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosKeys

`func (o *Global) SetKerberosKeys(v []KerberosKey)`

SetKerberosKeys sets KerberosKeys field to given value.

### HasKerberosKeys

`func (o *Global) HasKerberosKeys() bool`

HasKerberosKeys returns a boolean if a field has been set.

### GetLameTtl

`func (o *Global) GetLameTtl() int64`

GetLameTtl returns the LameTtl field if non-nil, zero value otherwise.

### GetLameTtlOk

`func (o *Global) GetLameTtlOk() (*int64, bool)`

GetLameTtlOk returns a tuple with the LameTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLameTtl

`func (o *Global) SetLameTtl(v int64)`

SetLameTtl sets LameTtl field to given value.

### HasLameTtl

`func (o *Global) HasLameTtl() bool`

HasLameTtl returns a boolean if a field has been set.

### GetLogQueryResponse

`func (o *Global) GetLogQueryResponse() bool`

GetLogQueryResponse returns the LogQueryResponse field if non-nil, zero value otherwise.

### GetLogQueryResponseOk

`func (o *Global) GetLogQueryResponseOk() (*bool, bool)`

GetLogQueryResponseOk returns a tuple with the LogQueryResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogQueryResponse

`func (o *Global) SetLogQueryResponse(v bool)`

SetLogQueryResponse sets LogQueryResponse field to given value.

### HasLogQueryResponse

`func (o *Global) HasLogQueryResponse() bool`

HasLogQueryResponse returns a boolean if a field has been set.

### GetMatchRecursiveOnly

`func (o *Global) GetMatchRecursiveOnly() bool`

GetMatchRecursiveOnly returns the MatchRecursiveOnly field if non-nil, zero value otherwise.

### GetMatchRecursiveOnlyOk

`func (o *Global) GetMatchRecursiveOnlyOk() (*bool, bool)`

GetMatchRecursiveOnlyOk returns a tuple with the MatchRecursiveOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchRecursiveOnly

`func (o *Global) SetMatchRecursiveOnly(v bool)`

SetMatchRecursiveOnly sets MatchRecursiveOnly field to given value.

### HasMatchRecursiveOnly

`func (o *Global) HasMatchRecursiveOnly() bool`

HasMatchRecursiveOnly returns a boolean if a field has been set.

### GetMaxCacheTtl

`func (o *Global) GetMaxCacheTtl() int64`

GetMaxCacheTtl returns the MaxCacheTtl field if non-nil, zero value otherwise.

### GetMaxCacheTtlOk

`func (o *Global) GetMaxCacheTtlOk() (*int64, bool)`

GetMaxCacheTtlOk returns a tuple with the MaxCacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheTtl

`func (o *Global) SetMaxCacheTtl(v int64)`

SetMaxCacheTtl sets MaxCacheTtl field to given value.

### HasMaxCacheTtl

`func (o *Global) HasMaxCacheTtl() bool`

HasMaxCacheTtl returns a boolean if a field has been set.

### GetMaxNegativeTtl

`func (o *Global) GetMaxNegativeTtl() int64`

GetMaxNegativeTtl returns the MaxNegativeTtl field if non-nil, zero value otherwise.

### GetMaxNegativeTtlOk

`func (o *Global) GetMaxNegativeTtlOk() (*int64, bool)`

GetMaxNegativeTtlOk returns a tuple with the MaxNegativeTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNegativeTtl

`func (o *Global) SetMaxNegativeTtl(v int64)`

SetMaxNegativeTtl sets MaxNegativeTtl field to given value.

### HasMaxNegativeTtl

`func (o *Global) HasMaxNegativeTtl() bool`

HasMaxNegativeTtl returns a boolean if a field has been set.

### GetMaxUdpSize

`func (o *Global) GetMaxUdpSize() int64`

GetMaxUdpSize returns the MaxUdpSize field if non-nil, zero value otherwise.

### GetMaxUdpSizeOk

`func (o *Global) GetMaxUdpSizeOk() (*int64, bool)`

GetMaxUdpSizeOk returns a tuple with the MaxUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUdpSize

`func (o *Global) SetMaxUdpSize(v int64)`

SetMaxUdpSize sets MaxUdpSize field to given value.

### HasMaxUdpSize

`func (o *Global) HasMaxUdpSize() bool`

HasMaxUdpSize returns a boolean if a field has been set.

### GetMinimalResponses

`func (o *Global) GetMinimalResponses() bool`

GetMinimalResponses returns the MinimalResponses field if non-nil, zero value otherwise.

### GetMinimalResponsesOk

`func (o *Global) GetMinimalResponsesOk() (*bool, bool)`

GetMinimalResponsesOk returns a tuple with the MinimalResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimalResponses

`func (o *Global) SetMinimalResponses(v bool)`

SetMinimalResponses sets MinimalResponses field to given value.

### HasMinimalResponses

`func (o *Global) HasMinimalResponses() bool`

HasMinimalResponses returns a boolean if a field has been set.

### GetNotify

`func (o *Global) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *Global) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *Global) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *Global) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetQueryAcl

`func (o *Global) GetQueryAcl() []ACLItem`

GetQueryAcl returns the QueryAcl field if non-nil, zero value otherwise.

### GetQueryAclOk

`func (o *Global) GetQueryAclOk() (*[]ACLItem, bool)`

GetQueryAclOk returns a tuple with the QueryAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryAcl

`func (o *Global) SetQueryAcl(v []ACLItem)`

SetQueryAcl sets QueryAcl field to given value.

### HasQueryAcl

`func (o *Global) HasQueryAcl() bool`

HasQueryAcl returns a boolean if a field has been set.

### GetQueryPort

`func (o *Global) GetQueryPort() int64`

GetQueryPort returns the QueryPort field if non-nil, zero value otherwise.

### GetQueryPortOk

`func (o *Global) GetQueryPortOk() (*int64, bool)`

GetQueryPortOk returns a tuple with the QueryPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryPort

`func (o *Global) SetQueryPort(v int64)`

SetQueryPort sets QueryPort field to given value.

### HasQueryPort

`func (o *Global) HasQueryPort() bool`

HasQueryPort returns a boolean if a field has been set.

### GetRecursionAcl

`func (o *Global) GetRecursionAcl() []ACLItem`

GetRecursionAcl returns the RecursionAcl field if non-nil, zero value otherwise.

### GetRecursionAclOk

`func (o *Global) GetRecursionAclOk() (*[]ACLItem, bool)`

GetRecursionAclOk returns a tuple with the RecursionAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursionAcl

`func (o *Global) SetRecursionAcl(v []ACLItem)`

SetRecursionAcl sets RecursionAcl field to given value.

### HasRecursionAcl

`func (o *Global) HasRecursionAcl() bool`

HasRecursionAcl returns a boolean if a field has been set.

### GetRecursionEnabled

`func (o *Global) GetRecursionEnabled() bool`

GetRecursionEnabled returns the RecursionEnabled field if non-nil, zero value otherwise.

### GetRecursionEnabledOk

`func (o *Global) GetRecursionEnabledOk() (*bool, bool)`

GetRecursionEnabledOk returns a tuple with the RecursionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursionEnabled

`func (o *Global) SetRecursionEnabled(v bool)`

SetRecursionEnabled sets RecursionEnabled field to given value.

### HasRecursionEnabled

`func (o *Global) HasRecursionEnabled() bool`

HasRecursionEnabled returns a boolean if a field has been set.

### GetRecursiveClients

`func (o *Global) GetRecursiveClients() int64`

GetRecursiveClients returns the RecursiveClients field if non-nil, zero value otherwise.

### GetRecursiveClientsOk

`func (o *Global) GetRecursiveClientsOk() (*int64, bool)`

GetRecursiveClientsOk returns a tuple with the RecursiveClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursiveClients

`func (o *Global) SetRecursiveClients(v int64)`

SetRecursiveClients sets RecursiveClients field to given value.

### HasRecursiveClients

`func (o *Global) HasRecursiveClients() bool`

HasRecursiveClients returns a boolean if a field has been set.

### GetResolverQueryTimeout

`func (o *Global) GetResolverQueryTimeout() int64`

GetResolverQueryTimeout returns the ResolverQueryTimeout field if non-nil, zero value otherwise.

### GetResolverQueryTimeoutOk

`func (o *Global) GetResolverQueryTimeoutOk() (*int64, bool)`

GetResolverQueryTimeoutOk returns a tuple with the ResolverQueryTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolverQueryTimeout

`func (o *Global) SetResolverQueryTimeout(v int64)`

SetResolverQueryTimeout sets ResolverQueryTimeout field to given value.

### HasResolverQueryTimeout

`func (o *Global) HasResolverQueryTimeout() bool`

HasResolverQueryTimeout returns a boolean if a field has been set.

### GetSecondaryAxfrQueryLimit

`func (o *Global) GetSecondaryAxfrQueryLimit() int64`

GetSecondaryAxfrQueryLimit returns the SecondaryAxfrQueryLimit field if non-nil, zero value otherwise.

### GetSecondaryAxfrQueryLimitOk

`func (o *Global) GetSecondaryAxfrQueryLimitOk() (*int64, bool)`

GetSecondaryAxfrQueryLimitOk returns a tuple with the SecondaryAxfrQueryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAxfrQueryLimit

`func (o *Global) SetSecondaryAxfrQueryLimit(v int64)`

SetSecondaryAxfrQueryLimit sets SecondaryAxfrQueryLimit field to given value.

### HasSecondaryAxfrQueryLimit

`func (o *Global) HasSecondaryAxfrQueryLimit() bool`

HasSecondaryAxfrQueryLimit returns a boolean if a field has been set.

### GetSecondarySoaQueryLimit

`func (o *Global) GetSecondarySoaQueryLimit() int64`

GetSecondarySoaQueryLimit returns the SecondarySoaQueryLimit field if non-nil, zero value otherwise.

### GetSecondarySoaQueryLimitOk

`func (o *Global) GetSecondarySoaQueryLimitOk() (*int64, bool)`

GetSecondarySoaQueryLimitOk returns a tuple with the SecondarySoaQueryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySoaQueryLimit

`func (o *Global) SetSecondarySoaQueryLimit(v int64)`

SetSecondarySoaQueryLimit sets SecondarySoaQueryLimit field to given value.

### HasSecondarySoaQueryLimit

`func (o *Global) HasSecondarySoaQueryLimit() bool`

HasSecondarySoaQueryLimit returns a boolean if a field has been set.

### GetSortList

`func (o *Global) GetSortList() []SortListItem`

GetSortList returns the SortList field if non-nil, zero value otherwise.

### GetSortListOk

`func (o *Global) GetSortListOk() (*[]SortListItem, bool)`

GetSortListOk returns a tuple with the SortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortList

`func (o *Global) SetSortList(v []SortListItem)`

SetSortList sets SortList field to given value.

### HasSortList

`func (o *Global) HasSortList() bool`

HasSortList returns a boolean if a field has been set.

### GetSynthesizeAddressRecordsFromHttps

`func (o *Global) GetSynthesizeAddressRecordsFromHttps() bool`

GetSynthesizeAddressRecordsFromHttps returns the SynthesizeAddressRecordsFromHttps field if non-nil, zero value otherwise.

### GetSynthesizeAddressRecordsFromHttpsOk

`func (o *Global) GetSynthesizeAddressRecordsFromHttpsOk() (*bool, bool)`

GetSynthesizeAddressRecordsFromHttpsOk returns a tuple with the SynthesizeAddressRecordsFromHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthesizeAddressRecordsFromHttps

`func (o *Global) SetSynthesizeAddressRecordsFromHttps(v bool)`

SetSynthesizeAddressRecordsFromHttps sets SynthesizeAddressRecordsFromHttps field to given value.

### HasSynthesizeAddressRecordsFromHttps

`func (o *Global) HasSynthesizeAddressRecordsFromHttps() bool`

HasSynthesizeAddressRecordsFromHttps returns a boolean if a field has been set.

### GetTransferAcl

`func (o *Global) GetTransferAcl() []ACLItem`

GetTransferAcl returns the TransferAcl field if non-nil, zero value otherwise.

### GetTransferAclOk

`func (o *Global) GetTransferAclOk() (*[]ACLItem, bool)`

GetTransferAclOk returns a tuple with the TransferAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAcl

`func (o *Global) SetTransferAcl(v []ACLItem)`

SetTransferAcl sets TransferAcl field to given value.

### HasTransferAcl

`func (o *Global) HasTransferAcl() bool`

HasTransferAcl returns a boolean if a field has been set.

### GetUpdateAcl

`func (o *Global) GetUpdateAcl() []ACLItem`

GetUpdateAcl returns the UpdateAcl field if non-nil, zero value otherwise.

### GetUpdateAclOk

`func (o *Global) GetUpdateAclOk() (*[]ACLItem, bool)`

GetUpdateAclOk returns a tuple with the UpdateAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAcl

`func (o *Global) SetUpdateAcl(v []ACLItem)`

SetUpdateAcl sets UpdateAcl field to given value.

### HasUpdateAcl

`func (o *Global) HasUpdateAcl() bool`

HasUpdateAcl returns a boolean if a field has been set.

### GetUseForwardersForSubzones

`func (o *Global) GetUseForwardersForSubzones() bool`

GetUseForwardersForSubzones returns the UseForwardersForSubzones field if non-nil, zero value otherwise.

### GetUseForwardersForSubzonesOk

`func (o *Global) GetUseForwardersForSubzonesOk() (*bool, bool)`

GetUseForwardersForSubzonesOk returns a tuple with the UseForwardersForSubzones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForwardersForSubzones

`func (o *Global) SetUseForwardersForSubzones(v bool)`

SetUseForwardersForSubzones sets UseForwardersForSubzones field to given value.

### HasUseForwardersForSubzones

`func (o *Global) HasUseForwardersForSubzones() bool`

HasUseForwardersForSubzones returns a boolean if a field has been set.

### GetUseRootForwardersForLocalResolutionWithB1td

`func (o *Global) GetUseRootForwardersForLocalResolutionWithB1td() bool`

GetUseRootForwardersForLocalResolutionWithB1td returns the UseRootForwardersForLocalResolutionWithB1td field if non-nil, zero value otherwise.

### GetUseRootForwardersForLocalResolutionWithB1tdOk

`func (o *Global) GetUseRootForwardersForLocalResolutionWithB1tdOk() (*bool, bool)`

GetUseRootForwardersForLocalResolutionWithB1tdOk returns a tuple with the UseRootForwardersForLocalResolutionWithB1td field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRootForwardersForLocalResolutionWithB1td

`func (o *Global) SetUseRootForwardersForLocalResolutionWithB1td(v bool)`

SetUseRootForwardersForLocalResolutionWithB1td sets UseRootForwardersForLocalResolutionWithB1td field to given value.

### HasUseRootForwardersForLocalResolutionWithB1td

`func (o *Global) HasUseRootForwardersForLocalResolutionWithB1td() bool`

HasUseRootForwardersForLocalResolutionWithB1td returns a boolean if a field has been set.

### GetZoneAuthority

`func (o *Global) GetZoneAuthority() ZoneAuthority`

GetZoneAuthority returns the ZoneAuthority field if non-nil, zero value otherwise.

### GetZoneAuthorityOk

`func (o *Global) GetZoneAuthorityOk() (*ZoneAuthority, bool)`

GetZoneAuthorityOk returns a tuple with the ZoneAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAuthority

`func (o *Global) SetZoneAuthority(v ZoneAuthority)`

SetZoneAuthority sets ZoneAuthority field to given value.

### HasZoneAuthority

`func (o *Global) HasZoneAuthority() bool`

HasZoneAuthority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


