# Server

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddEdnsOptionInOutgoingQuery** | Pointer to **bool** | _add_edns_option_in_outgoing_query_ adds client IP, MAC address and view name into outgoing recursive query. Defaults to _false_. | [optional] 
**AutoSortViews** | Pointer to **bool** | Optional. Controls manual/automatic views ordering.  Defaults to _true_. | [optional] 
**Comment** | Pointer to **string** | Optional. Comment for configuration. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**CustomRootNs** | Pointer to [**[]RootNS**](RootNS.md) | Optional. List of custom root nameservers. The order does not matter.  Error if empty while _custom_root_ns_enabled_ is _true_. Error if there are duplicate items in the list.  Defaults to empty. | [optional] 
**CustomRootNsEnabled** | Pointer to **bool** | Optional. _true_ to use custom root nameservers instead of the default ones.  The _custom_root_ns_ is validated when enabled.  Defaults to _false_. | [optional] 
**DnssecEnableValidation** | Pointer to **bool** | Optional. _true_ to perform DNSSEC validation. Ignored if _dnssec_enabled_ is _false_.  Defaults to _true_. | [optional] 
**DnssecEnabled** | Pointer to **bool** | Optional. Master toggle for all DNSSEC processing. Other _dnssec_*_ configuration is unused if this is disabled.  Defaults to _true_. | [optional] 
**DnssecRootKeys** | Pointer to [**[]TrustAnchor**](TrustAnchor.md) | DNSSEC root keys. The root keys are not configurable.  A default list is provided by cloud management and included here for config generation. | [optional] [readonly] 
**DnssecTrustAnchors** | Pointer to [**[]TrustAnchor**](TrustAnchor.md) | Optional. DNSSEC trust anchors.  Error if there are list items with duplicate (_zone_, _sep_, _algorithm_) combinations.  Defaults to empty. | [optional] 
**DnssecValidateExpiry** | Pointer to **bool** | Optional. _true_ to reject expired DNSSEC keys. Ignored if either _dnssec_enabled_ or _dnssec_enable_validation_ is _false_.  Defaults to _true_. | [optional] 
**EcsEnabled** | Pointer to **bool** | Optional. _true_ to enable EDNS client subnet for recursive queries. Other _ecs_*_ fields are ignored if this field is not enabled.  Defaults to _false_. | [optional] 
**EcsForwarding** | Pointer to **bool** | Optional. _true_ to enable ECS options in outbound queries. This functionality has additional overhead so it is disabled by default.  Defaults to _false_. | [optional] 
**EcsPrefixV4** | Pointer to **int64** | Optional. Maximum scope length for v4 ECS.  Unsigned integer, min 1 max 24  Defaults to 24. | [optional] 
**EcsPrefixV6** | Pointer to **int64** | Optional. Maximum scope length for v6 ECS.  Unsigned integer, min 1 max 56  Defaults to 56. | [optional] 
**EcsZones** | Pointer to [**[]ECSZone**](ECSZone.md) | Optional. List of zones where ECS queries may be sent.  Error if empty while _ecs_enabled_ is _true_. Error if there are duplicate FQDNs in the list.  Defaults to empty. | [optional] 
**FilterAaaaAcl** | Pointer to [**[]ACLItem**](ACLItem.md) | Optional. Specifies a list of client addresses for which AAAA filtering is to be applied.  Defaults to _empty_. | [optional] 
**FilterAaaaOnV4** | Pointer to **string** | _filter_aaaa_on_v4_ allows named to omit some IPv6 addresses when responding to IPv4 clients.  Allowed values: * _yes_, * _no_, * _break_dnssec_.  Defaults to _no_ | [optional] 
**Forwarders** | Pointer to [**[]Forwarder**](Forwarder.md) | Optional. List of forwarders.  Error if empty while _forwarders_only_ or _use_root_forwarders_for_local_resolution_with_b1td_ is _true_. Error if there are items in the list with duplicate addresses.  Defaults to empty. | [optional] 
**ForwardersOnly** | Pointer to **bool** | Optional. _true_ to only forward.  Defaults to _false_. | [optional] 
**GssTsigEnabled** | Pointer to **bool** | _gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.  Defaults to _false_. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InheritanceSources** | Pointer to [**ServerInheritance**](ServerInheritance.md) |  | [optional] 
**KerberosKeys** | Pointer to [**[]KerberosKey**](KerberosKey.md) | _kerberos_keys_ contains a list of keys for GSS-TSIG signed dynamic updates.  Defaults to empty. | [optional] 
**LameTtl** | Pointer to **int64** | Optional. Unused in the current on-prem DNS server implementation.  Unsigned integer, min 0 max 3600 (1h).  Defaults to 600. | [optional] 
**LogQueryResponse** | Pointer to **bool** | Optional. Control DNS query/response logging functionality.  Defaults to _true_. | [optional] 
**MatchRecursiveOnly** | Pointer to **bool** | Optional. If _true_ only recursive queries from matching clients access the view.  Defaults to _false_. | [optional] 
**MaxCacheTtl** | Pointer to **int64** | Optional. Seconds to cache positive responses.  Unsigned integer, min 1 max 604800 (7d).  Defaults to 604800 (7d). | [optional] 
**MaxNegativeTtl** | Pointer to **int64** | Optional. Seconds to cache negative responses.  Unsigned integer, min 1 max 604800 (7d).  Defaults to 10800 (3h). | [optional] 
**MinimalResponses** | Pointer to **bool** | Optional. When enabled, the DNS server will only add records to the authority and additional data sections when they are required.  Defaults to _false_. | [optional] 
**Name** | **string** | Name of configuration. | 
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
**Tags** | Pointer to **map[string]interface{}** | Tagging specifics. | [optional] 
**TransferAcl** | Pointer to [**[]ACLItem**](ACLItem.md) | Optional. Clients must match this ACL to receive zone transfers.  Defaults to empty. | [optional] 
**UpdateAcl** | Pointer to [**[]ACLItem**](ACLItem.md) | Optional. Specifies which hosts are allowed to issue Dynamic DNS updates for authoritative zones of _primary_type_ _cloud_.  Defaults to empty. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 
**UseForwardersForSubzones** | Pointer to **bool** | Optional. Use default forwarders to resolve queries for subzones.  Defaults to _true_. | [optional] 
**UseRootForwardersForLocalResolutionWithB1td** | Pointer to **bool** | _use_root_forwarders_for_local_resolution_with_b1td_ allows DNS recursive queries sent to root forwarders for local resolution when deployed alongside BloxOne Thread Defense. Defaults to _false_. | [optional] 
**Views** | Pointer to [**[]DisplayView**](DisplayView.md) | Optional. Ordered list of _dns/display_view_ objects served by any of _dns/host_ assigned to a particular DNS Config Profile. Automatically determined. Allows re-ordering only. | [optional] 

## Methods

### NewServer

`func NewServer(name string, ) *Server`

NewServer instantiates a new Server object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerWithDefaults

`func NewServerWithDefaults() *Server`

NewServerWithDefaults instantiates a new Server object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddEdnsOptionInOutgoingQuery

`func (o *Server) GetAddEdnsOptionInOutgoingQuery() bool`

GetAddEdnsOptionInOutgoingQuery returns the AddEdnsOptionInOutgoingQuery field if non-nil, zero value otherwise.

### GetAddEdnsOptionInOutgoingQueryOk

`func (o *Server) GetAddEdnsOptionInOutgoingQueryOk() (*bool, bool)`

GetAddEdnsOptionInOutgoingQueryOk returns a tuple with the AddEdnsOptionInOutgoingQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddEdnsOptionInOutgoingQuery

`func (o *Server) SetAddEdnsOptionInOutgoingQuery(v bool)`

SetAddEdnsOptionInOutgoingQuery sets AddEdnsOptionInOutgoingQuery field to given value.

### HasAddEdnsOptionInOutgoingQuery

`func (o *Server) HasAddEdnsOptionInOutgoingQuery() bool`

HasAddEdnsOptionInOutgoingQuery returns a boolean if a field has been set.

### GetAutoSortViews

`func (o *Server) GetAutoSortViews() bool`

GetAutoSortViews returns the AutoSortViews field if non-nil, zero value otherwise.

### GetAutoSortViewsOk

`func (o *Server) GetAutoSortViewsOk() (*bool, bool)`

GetAutoSortViewsOk returns a tuple with the AutoSortViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSortViews

`func (o *Server) SetAutoSortViews(v bool)`

SetAutoSortViews sets AutoSortViews field to given value.

### HasAutoSortViews

`func (o *Server) HasAutoSortViews() bool`

HasAutoSortViews returns a boolean if a field has been set.

### GetComment

`func (o *Server) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Server) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Server) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Server) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Server) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Server) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Server) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Server) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomRootNs

`func (o *Server) GetCustomRootNs() []RootNS`

GetCustomRootNs returns the CustomRootNs field if non-nil, zero value otherwise.

### GetCustomRootNsOk

`func (o *Server) GetCustomRootNsOk() (*[]RootNS, bool)`

GetCustomRootNsOk returns a tuple with the CustomRootNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRootNs

`func (o *Server) SetCustomRootNs(v []RootNS)`

SetCustomRootNs sets CustomRootNs field to given value.

### HasCustomRootNs

`func (o *Server) HasCustomRootNs() bool`

HasCustomRootNs returns a boolean if a field has been set.

### GetCustomRootNsEnabled

`func (o *Server) GetCustomRootNsEnabled() bool`

GetCustomRootNsEnabled returns the CustomRootNsEnabled field if non-nil, zero value otherwise.

### GetCustomRootNsEnabledOk

`func (o *Server) GetCustomRootNsEnabledOk() (*bool, bool)`

GetCustomRootNsEnabledOk returns a tuple with the CustomRootNsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRootNsEnabled

`func (o *Server) SetCustomRootNsEnabled(v bool)`

SetCustomRootNsEnabled sets CustomRootNsEnabled field to given value.

### HasCustomRootNsEnabled

`func (o *Server) HasCustomRootNsEnabled() bool`

HasCustomRootNsEnabled returns a boolean if a field has been set.

### GetDnssecEnableValidation

`func (o *Server) GetDnssecEnableValidation() bool`

GetDnssecEnableValidation returns the DnssecEnableValidation field if non-nil, zero value otherwise.

### GetDnssecEnableValidationOk

`func (o *Server) GetDnssecEnableValidationOk() (*bool, bool)`

GetDnssecEnableValidationOk returns a tuple with the DnssecEnableValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecEnableValidation

`func (o *Server) SetDnssecEnableValidation(v bool)`

SetDnssecEnableValidation sets DnssecEnableValidation field to given value.

### HasDnssecEnableValidation

`func (o *Server) HasDnssecEnableValidation() bool`

HasDnssecEnableValidation returns a boolean if a field has been set.

### GetDnssecEnabled

`func (o *Server) GetDnssecEnabled() bool`

GetDnssecEnabled returns the DnssecEnabled field if non-nil, zero value otherwise.

### GetDnssecEnabledOk

`func (o *Server) GetDnssecEnabledOk() (*bool, bool)`

GetDnssecEnabledOk returns a tuple with the DnssecEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecEnabled

`func (o *Server) SetDnssecEnabled(v bool)`

SetDnssecEnabled sets DnssecEnabled field to given value.

### HasDnssecEnabled

`func (o *Server) HasDnssecEnabled() bool`

HasDnssecEnabled returns a boolean if a field has been set.

### GetDnssecRootKeys

`func (o *Server) GetDnssecRootKeys() []TrustAnchor`

GetDnssecRootKeys returns the DnssecRootKeys field if non-nil, zero value otherwise.

### GetDnssecRootKeysOk

`func (o *Server) GetDnssecRootKeysOk() (*[]TrustAnchor, bool)`

GetDnssecRootKeysOk returns a tuple with the DnssecRootKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecRootKeys

`func (o *Server) SetDnssecRootKeys(v []TrustAnchor)`

SetDnssecRootKeys sets DnssecRootKeys field to given value.

### HasDnssecRootKeys

`func (o *Server) HasDnssecRootKeys() bool`

HasDnssecRootKeys returns a boolean if a field has been set.

### GetDnssecTrustAnchors

`func (o *Server) GetDnssecTrustAnchors() []TrustAnchor`

GetDnssecTrustAnchors returns the DnssecTrustAnchors field if non-nil, zero value otherwise.

### GetDnssecTrustAnchorsOk

`func (o *Server) GetDnssecTrustAnchorsOk() (*[]TrustAnchor, bool)`

GetDnssecTrustAnchorsOk returns a tuple with the DnssecTrustAnchors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecTrustAnchors

`func (o *Server) SetDnssecTrustAnchors(v []TrustAnchor)`

SetDnssecTrustAnchors sets DnssecTrustAnchors field to given value.

### HasDnssecTrustAnchors

`func (o *Server) HasDnssecTrustAnchors() bool`

HasDnssecTrustAnchors returns a boolean if a field has been set.

### GetDnssecValidateExpiry

`func (o *Server) GetDnssecValidateExpiry() bool`

GetDnssecValidateExpiry returns the DnssecValidateExpiry field if non-nil, zero value otherwise.

### GetDnssecValidateExpiryOk

`func (o *Server) GetDnssecValidateExpiryOk() (*bool, bool)`

GetDnssecValidateExpiryOk returns a tuple with the DnssecValidateExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecValidateExpiry

`func (o *Server) SetDnssecValidateExpiry(v bool)`

SetDnssecValidateExpiry sets DnssecValidateExpiry field to given value.

### HasDnssecValidateExpiry

`func (o *Server) HasDnssecValidateExpiry() bool`

HasDnssecValidateExpiry returns a boolean if a field has been set.

### GetEcsEnabled

`func (o *Server) GetEcsEnabled() bool`

GetEcsEnabled returns the EcsEnabled field if non-nil, zero value otherwise.

### GetEcsEnabledOk

`func (o *Server) GetEcsEnabledOk() (*bool, bool)`

GetEcsEnabledOk returns a tuple with the EcsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsEnabled

`func (o *Server) SetEcsEnabled(v bool)`

SetEcsEnabled sets EcsEnabled field to given value.

### HasEcsEnabled

`func (o *Server) HasEcsEnabled() bool`

HasEcsEnabled returns a boolean if a field has been set.

### GetEcsForwarding

`func (o *Server) GetEcsForwarding() bool`

GetEcsForwarding returns the EcsForwarding field if non-nil, zero value otherwise.

### GetEcsForwardingOk

`func (o *Server) GetEcsForwardingOk() (*bool, bool)`

GetEcsForwardingOk returns a tuple with the EcsForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsForwarding

`func (o *Server) SetEcsForwarding(v bool)`

SetEcsForwarding sets EcsForwarding field to given value.

### HasEcsForwarding

`func (o *Server) HasEcsForwarding() bool`

HasEcsForwarding returns a boolean if a field has been set.

### GetEcsPrefixV4

`func (o *Server) GetEcsPrefixV4() int64`

GetEcsPrefixV4 returns the EcsPrefixV4 field if non-nil, zero value otherwise.

### GetEcsPrefixV4Ok

`func (o *Server) GetEcsPrefixV4Ok() (*int64, bool)`

GetEcsPrefixV4Ok returns a tuple with the EcsPrefixV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsPrefixV4

`func (o *Server) SetEcsPrefixV4(v int64)`

SetEcsPrefixV4 sets EcsPrefixV4 field to given value.

### HasEcsPrefixV4

`func (o *Server) HasEcsPrefixV4() bool`

HasEcsPrefixV4 returns a boolean if a field has been set.

### GetEcsPrefixV6

`func (o *Server) GetEcsPrefixV6() int64`

GetEcsPrefixV6 returns the EcsPrefixV6 field if non-nil, zero value otherwise.

### GetEcsPrefixV6Ok

`func (o *Server) GetEcsPrefixV6Ok() (*int64, bool)`

GetEcsPrefixV6Ok returns a tuple with the EcsPrefixV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsPrefixV6

`func (o *Server) SetEcsPrefixV6(v int64)`

SetEcsPrefixV6 sets EcsPrefixV6 field to given value.

### HasEcsPrefixV6

`func (o *Server) HasEcsPrefixV6() bool`

HasEcsPrefixV6 returns a boolean if a field has been set.

### GetEcsZones

`func (o *Server) GetEcsZones() []ECSZone`

GetEcsZones returns the EcsZones field if non-nil, zero value otherwise.

### GetEcsZonesOk

`func (o *Server) GetEcsZonesOk() (*[]ECSZone, bool)`

GetEcsZonesOk returns a tuple with the EcsZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsZones

`func (o *Server) SetEcsZones(v []ECSZone)`

SetEcsZones sets EcsZones field to given value.

### HasEcsZones

`func (o *Server) HasEcsZones() bool`

HasEcsZones returns a boolean if a field has been set.

### GetFilterAaaaAcl

`func (o *Server) GetFilterAaaaAcl() []ACLItem`

GetFilterAaaaAcl returns the FilterAaaaAcl field if non-nil, zero value otherwise.

### GetFilterAaaaAclOk

`func (o *Server) GetFilterAaaaAclOk() (*[]ACLItem, bool)`

GetFilterAaaaAclOk returns a tuple with the FilterAaaaAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaaAcl

`func (o *Server) SetFilterAaaaAcl(v []ACLItem)`

SetFilterAaaaAcl sets FilterAaaaAcl field to given value.

### HasFilterAaaaAcl

`func (o *Server) HasFilterAaaaAcl() bool`

HasFilterAaaaAcl returns a boolean if a field has been set.

### GetFilterAaaaOnV4

`func (o *Server) GetFilterAaaaOnV4() string`

GetFilterAaaaOnV4 returns the FilterAaaaOnV4 field if non-nil, zero value otherwise.

### GetFilterAaaaOnV4Ok

`func (o *Server) GetFilterAaaaOnV4Ok() (*string, bool)`

GetFilterAaaaOnV4Ok returns a tuple with the FilterAaaaOnV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaaOnV4

`func (o *Server) SetFilterAaaaOnV4(v string)`

SetFilterAaaaOnV4 sets FilterAaaaOnV4 field to given value.

### HasFilterAaaaOnV4

`func (o *Server) HasFilterAaaaOnV4() bool`

HasFilterAaaaOnV4 returns a boolean if a field has been set.

### GetForwarders

`func (o *Server) GetForwarders() []Forwarder`

GetForwarders returns the Forwarders field if non-nil, zero value otherwise.

### GetForwardersOk

`func (o *Server) GetForwardersOk() (*[]Forwarder, bool)`

GetForwardersOk returns a tuple with the Forwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwarders

`func (o *Server) SetForwarders(v []Forwarder)`

SetForwarders sets Forwarders field to given value.

### HasForwarders

`func (o *Server) HasForwarders() bool`

HasForwarders returns a boolean if a field has been set.

### GetForwardersOnly

`func (o *Server) GetForwardersOnly() bool`

GetForwardersOnly returns the ForwardersOnly field if non-nil, zero value otherwise.

### GetForwardersOnlyOk

`func (o *Server) GetForwardersOnlyOk() (*bool, bool)`

GetForwardersOnlyOk returns a tuple with the ForwardersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardersOnly

`func (o *Server) SetForwardersOnly(v bool)`

SetForwardersOnly sets ForwardersOnly field to given value.

### HasForwardersOnly

`func (o *Server) HasForwardersOnly() bool`

HasForwardersOnly returns a boolean if a field has been set.

### GetGssTsigEnabled

`func (o *Server) GetGssTsigEnabled() bool`

GetGssTsigEnabled returns the GssTsigEnabled field if non-nil, zero value otherwise.

### GetGssTsigEnabledOk

`func (o *Server) GetGssTsigEnabledOk() (*bool, bool)`

GetGssTsigEnabledOk returns a tuple with the GssTsigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigEnabled

`func (o *Server) SetGssTsigEnabled(v bool)`

SetGssTsigEnabled sets GssTsigEnabled field to given value.

### HasGssTsigEnabled

`func (o *Server) HasGssTsigEnabled() bool`

HasGssTsigEnabled returns a boolean if a field has been set.

### GetId

`func (o *Server) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Server) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Server) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Server) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritanceSources

`func (o *Server) GetInheritanceSources() ServerInheritance`

GetInheritanceSources returns the InheritanceSources field if non-nil, zero value otherwise.

### GetInheritanceSourcesOk

`func (o *Server) GetInheritanceSourcesOk() (*ServerInheritance, bool)`

GetInheritanceSourcesOk returns a tuple with the InheritanceSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceSources

`func (o *Server) SetInheritanceSources(v ServerInheritance)`

SetInheritanceSources sets InheritanceSources field to given value.

### HasInheritanceSources

`func (o *Server) HasInheritanceSources() bool`

HasInheritanceSources returns a boolean if a field has been set.

### GetKerberosKeys

`func (o *Server) GetKerberosKeys() []KerberosKey`

GetKerberosKeys returns the KerberosKeys field if non-nil, zero value otherwise.

### GetKerberosKeysOk

`func (o *Server) GetKerberosKeysOk() (*[]KerberosKey, bool)`

GetKerberosKeysOk returns a tuple with the KerberosKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosKeys

`func (o *Server) SetKerberosKeys(v []KerberosKey)`

SetKerberosKeys sets KerberosKeys field to given value.

### HasKerberosKeys

`func (o *Server) HasKerberosKeys() bool`

HasKerberosKeys returns a boolean if a field has been set.

### GetLameTtl

`func (o *Server) GetLameTtl() int64`

GetLameTtl returns the LameTtl field if non-nil, zero value otherwise.

### GetLameTtlOk

`func (o *Server) GetLameTtlOk() (*int64, bool)`

GetLameTtlOk returns a tuple with the LameTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLameTtl

`func (o *Server) SetLameTtl(v int64)`

SetLameTtl sets LameTtl field to given value.

### HasLameTtl

`func (o *Server) HasLameTtl() bool`

HasLameTtl returns a boolean if a field has been set.

### GetLogQueryResponse

`func (o *Server) GetLogQueryResponse() bool`

GetLogQueryResponse returns the LogQueryResponse field if non-nil, zero value otherwise.

### GetLogQueryResponseOk

`func (o *Server) GetLogQueryResponseOk() (*bool, bool)`

GetLogQueryResponseOk returns a tuple with the LogQueryResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogQueryResponse

`func (o *Server) SetLogQueryResponse(v bool)`

SetLogQueryResponse sets LogQueryResponse field to given value.

### HasLogQueryResponse

`func (o *Server) HasLogQueryResponse() bool`

HasLogQueryResponse returns a boolean if a field has been set.

### GetMatchRecursiveOnly

`func (o *Server) GetMatchRecursiveOnly() bool`

GetMatchRecursiveOnly returns the MatchRecursiveOnly field if non-nil, zero value otherwise.

### GetMatchRecursiveOnlyOk

`func (o *Server) GetMatchRecursiveOnlyOk() (*bool, bool)`

GetMatchRecursiveOnlyOk returns a tuple with the MatchRecursiveOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchRecursiveOnly

`func (o *Server) SetMatchRecursiveOnly(v bool)`

SetMatchRecursiveOnly sets MatchRecursiveOnly field to given value.

### HasMatchRecursiveOnly

`func (o *Server) HasMatchRecursiveOnly() bool`

HasMatchRecursiveOnly returns a boolean if a field has been set.

### GetMaxCacheTtl

`func (o *Server) GetMaxCacheTtl() int64`

GetMaxCacheTtl returns the MaxCacheTtl field if non-nil, zero value otherwise.

### GetMaxCacheTtlOk

`func (o *Server) GetMaxCacheTtlOk() (*int64, bool)`

GetMaxCacheTtlOk returns a tuple with the MaxCacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheTtl

`func (o *Server) SetMaxCacheTtl(v int64)`

SetMaxCacheTtl sets MaxCacheTtl field to given value.

### HasMaxCacheTtl

`func (o *Server) HasMaxCacheTtl() bool`

HasMaxCacheTtl returns a boolean if a field has been set.

### GetMaxNegativeTtl

`func (o *Server) GetMaxNegativeTtl() int64`

GetMaxNegativeTtl returns the MaxNegativeTtl field if non-nil, zero value otherwise.

### GetMaxNegativeTtlOk

`func (o *Server) GetMaxNegativeTtlOk() (*int64, bool)`

GetMaxNegativeTtlOk returns a tuple with the MaxNegativeTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNegativeTtl

`func (o *Server) SetMaxNegativeTtl(v int64)`

SetMaxNegativeTtl sets MaxNegativeTtl field to given value.

### HasMaxNegativeTtl

`func (o *Server) HasMaxNegativeTtl() bool`

HasMaxNegativeTtl returns a boolean if a field has been set.

### GetMinimalResponses

`func (o *Server) GetMinimalResponses() bool`

GetMinimalResponses returns the MinimalResponses field if non-nil, zero value otherwise.

### GetMinimalResponsesOk

`func (o *Server) GetMinimalResponsesOk() (*bool, bool)`

GetMinimalResponsesOk returns a tuple with the MinimalResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimalResponses

`func (o *Server) SetMinimalResponses(v bool)`

SetMinimalResponses sets MinimalResponses field to given value.

### HasMinimalResponses

`func (o *Server) HasMinimalResponses() bool`

HasMinimalResponses returns a boolean if a field has been set.

### GetName

`func (o *Server) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Server) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Server) SetName(v string)`

SetName sets Name field to given value.


### GetNotify

`func (o *Server) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *Server) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *Server) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *Server) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetQueryAcl

`func (o *Server) GetQueryAcl() []ACLItem`

GetQueryAcl returns the QueryAcl field if non-nil, zero value otherwise.

### GetQueryAclOk

`func (o *Server) GetQueryAclOk() (*[]ACLItem, bool)`

GetQueryAclOk returns a tuple with the QueryAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryAcl

`func (o *Server) SetQueryAcl(v []ACLItem)`

SetQueryAcl sets QueryAcl field to given value.

### HasQueryAcl

`func (o *Server) HasQueryAcl() bool`

HasQueryAcl returns a boolean if a field has been set.

### GetQueryPort

`func (o *Server) GetQueryPort() int64`

GetQueryPort returns the QueryPort field if non-nil, zero value otherwise.

### GetQueryPortOk

`func (o *Server) GetQueryPortOk() (*int64, bool)`

GetQueryPortOk returns a tuple with the QueryPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryPort

`func (o *Server) SetQueryPort(v int64)`

SetQueryPort sets QueryPort field to given value.

### HasQueryPort

`func (o *Server) HasQueryPort() bool`

HasQueryPort returns a boolean if a field has been set.

### GetRecursionAcl

`func (o *Server) GetRecursionAcl() []ACLItem`

GetRecursionAcl returns the RecursionAcl field if non-nil, zero value otherwise.

### GetRecursionAclOk

`func (o *Server) GetRecursionAclOk() (*[]ACLItem, bool)`

GetRecursionAclOk returns a tuple with the RecursionAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursionAcl

`func (o *Server) SetRecursionAcl(v []ACLItem)`

SetRecursionAcl sets RecursionAcl field to given value.

### HasRecursionAcl

`func (o *Server) HasRecursionAcl() bool`

HasRecursionAcl returns a boolean if a field has been set.

### GetRecursionEnabled

`func (o *Server) GetRecursionEnabled() bool`

GetRecursionEnabled returns the RecursionEnabled field if non-nil, zero value otherwise.

### GetRecursionEnabledOk

`func (o *Server) GetRecursionEnabledOk() (*bool, bool)`

GetRecursionEnabledOk returns a tuple with the RecursionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursionEnabled

`func (o *Server) SetRecursionEnabled(v bool)`

SetRecursionEnabled sets RecursionEnabled field to given value.

### HasRecursionEnabled

`func (o *Server) HasRecursionEnabled() bool`

HasRecursionEnabled returns a boolean if a field has been set.

### GetRecursiveClients

`func (o *Server) GetRecursiveClients() int64`

GetRecursiveClients returns the RecursiveClients field if non-nil, zero value otherwise.

### GetRecursiveClientsOk

`func (o *Server) GetRecursiveClientsOk() (*int64, bool)`

GetRecursiveClientsOk returns a tuple with the RecursiveClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursiveClients

`func (o *Server) SetRecursiveClients(v int64)`

SetRecursiveClients sets RecursiveClients field to given value.

### HasRecursiveClients

`func (o *Server) HasRecursiveClients() bool`

HasRecursiveClients returns a boolean if a field has been set.

### GetResolverQueryTimeout

`func (o *Server) GetResolverQueryTimeout() int64`

GetResolverQueryTimeout returns the ResolverQueryTimeout field if non-nil, zero value otherwise.

### GetResolverQueryTimeoutOk

`func (o *Server) GetResolverQueryTimeoutOk() (*int64, bool)`

GetResolverQueryTimeoutOk returns a tuple with the ResolverQueryTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolverQueryTimeout

`func (o *Server) SetResolverQueryTimeout(v int64)`

SetResolverQueryTimeout sets ResolverQueryTimeout field to given value.

### HasResolverQueryTimeout

`func (o *Server) HasResolverQueryTimeout() bool`

HasResolverQueryTimeout returns a boolean if a field has been set.

### GetSecondaryAxfrQueryLimit

`func (o *Server) GetSecondaryAxfrQueryLimit() int64`

GetSecondaryAxfrQueryLimit returns the SecondaryAxfrQueryLimit field if non-nil, zero value otherwise.

### GetSecondaryAxfrQueryLimitOk

`func (o *Server) GetSecondaryAxfrQueryLimitOk() (*int64, bool)`

GetSecondaryAxfrQueryLimitOk returns a tuple with the SecondaryAxfrQueryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAxfrQueryLimit

`func (o *Server) SetSecondaryAxfrQueryLimit(v int64)`

SetSecondaryAxfrQueryLimit sets SecondaryAxfrQueryLimit field to given value.

### HasSecondaryAxfrQueryLimit

`func (o *Server) HasSecondaryAxfrQueryLimit() bool`

HasSecondaryAxfrQueryLimit returns a boolean if a field has been set.

### GetSecondarySoaQueryLimit

`func (o *Server) GetSecondarySoaQueryLimit() int64`

GetSecondarySoaQueryLimit returns the SecondarySoaQueryLimit field if non-nil, zero value otherwise.

### GetSecondarySoaQueryLimitOk

`func (o *Server) GetSecondarySoaQueryLimitOk() (*int64, bool)`

GetSecondarySoaQueryLimitOk returns a tuple with the SecondarySoaQueryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySoaQueryLimit

`func (o *Server) SetSecondarySoaQueryLimit(v int64)`

SetSecondarySoaQueryLimit sets SecondarySoaQueryLimit field to given value.

### HasSecondarySoaQueryLimit

`func (o *Server) HasSecondarySoaQueryLimit() bool`

HasSecondarySoaQueryLimit returns a boolean if a field has been set.

### GetSortList

`func (o *Server) GetSortList() []SortListItem`

GetSortList returns the SortList field if non-nil, zero value otherwise.

### GetSortListOk

`func (o *Server) GetSortListOk() (*[]SortListItem, bool)`

GetSortListOk returns a tuple with the SortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortList

`func (o *Server) SetSortList(v []SortListItem)`

SetSortList sets SortList field to given value.

### HasSortList

`func (o *Server) HasSortList() bool`

HasSortList returns a boolean if a field has been set.

### GetSynthesizeAddressRecordsFromHttps

`func (o *Server) GetSynthesizeAddressRecordsFromHttps() bool`

GetSynthesizeAddressRecordsFromHttps returns the SynthesizeAddressRecordsFromHttps field if non-nil, zero value otherwise.

### GetSynthesizeAddressRecordsFromHttpsOk

`func (o *Server) GetSynthesizeAddressRecordsFromHttpsOk() (*bool, bool)`

GetSynthesizeAddressRecordsFromHttpsOk returns a tuple with the SynthesizeAddressRecordsFromHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthesizeAddressRecordsFromHttps

`func (o *Server) SetSynthesizeAddressRecordsFromHttps(v bool)`

SetSynthesizeAddressRecordsFromHttps sets SynthesizeAddressRecordsFromHttps field to given value.

### HasSynthesizeAddressRecordsFromHttps

`func (o *Server) HasSynthesizeAddressRecordsFromHttps() bool`

HasSynthesizeAddressRecordsFromHttps returns a boolean if a field has been set.

### GetTags

`func (o *Server) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Server) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Server) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Server) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTransferAcl

`func (o *Server) GetTransferAcl() []ACLItem`

GetTransferAcl returns the TransferAcl field if non-nil, zero value otherwise.

### GetTransferAclOk

`func (o *Server) GetTransferAclOk() (*[]ACLItem, bool)`

GetTransferAclOk returns a tuple with the TransferAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAcl

`func (o *Server) SetTransferAcl(v []ACLItem)`

SetTransferAcl sets TransferAcl field to given value.

### HasTransferAcl

`func (o *Server) HasTransferAcl() bool`

HasTransferAcl returns a boolean if a field has been set.

### GetUpdateAcl

`func (o *Server) GetUpdateAcl() []ACLItem`

GetUpdateAcl returns the UpdateAcl field if non-nil, zero value otherwise.

### GetUpdateAclOk

`func (o *Server) GetUpdateAclOk() (*[]ACLItem, bool)`

GetUpdateAclOk returns a tuple with the UpdateAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAcl

`func (o *Server) SetUpdateAcl(v []ACLItem)`

SetUpdateAcl sets UpdateAcl field to given value.

### HasUpdateAcl

`func (o *Server) HasUpdateAcl() bool`

HasUpdateAcl returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Server) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Server) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Server) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Server) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUseForwardersForSubzones

`func (o *Server) GetUseForwardersForSubzones() bool`

GetUseForwardersForSubzones returns the UseForwardersForSubzones field if non-nil, zero value otherwise.

### GetUseForwardersForSubzonesOk

`func (o *Server) GetUseForwardersForSubzonesOk() (*bool, bool)`

GetUseForwardersForSubzonesOk returns a tuple with the UseForwardersForSubzones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForwardersForSubzones

`func (o *Server) SetUseForwardersForSubzones(v bool)`

SetUseForwardersForSubzones sets UseForwardersForSubzones field to given value.

### HasUseForwardersForSubzones

`func (o *Server) HasUseForwardersForSubzones() bool`

HasUseForwardersForSubzones returns a boolean if a field has been set.

### GetUseRootForwardersForLocalResolutionWithB1td

`func (o *Server) GetUseRootForwardersForLocalResolutionWithB1td() bool`

GetUseRootForwardersForLocalResolutionWithB1td returns the UseRootForwardersForLocalResolutionWithB1td field if non-nil, zero value otherwise.

### GetUseRootForwardersForLocalResolutionWithB1tdOk

`func (o *Server) GetUseRootForwardersForLocalResolutionWithB1tdOk() (*bool, bool)`

GetUseRootForwardersForLocalResolutionWithB1tdOk returns a tuple with the UseRootForwardersForLocalResolutionWithB1td field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRootForwardersForLocalResolutionWithB1td

`func (o *Server) SetUseRootForwardersForLocalResolutionWithB1td(v bool)`

SetUseRootForwardersForLocalResolutionWithB1td sets UseRootForwardersForLocalResolutionWithB1td field to given value.

### HasUseRootForwardersForLocalResolutionWithB1td

`func (o *Server) HasUseRootForwardersForLocalResolutionWithB1td() bool`

HasUseRootForwardersForLocalResolutionWithB1td returns a boolean if a field has been set.

### GetViews

`func (o *Server) GetViews() []DisplayView`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *Server) GetViewsOk() (*[]DisplayView, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *Server) SetViews(v []DisplayView)`

SetViews sets Views field to given value.

### HasViews

`func (o *Server) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


