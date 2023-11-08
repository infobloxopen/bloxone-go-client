# IpamsvcGlobal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsmConfig** | Pointer to [**IpamsvcASMConfig**](IpamsvcASMConfig.md) |  | [optional] 
**ClientPrincipal** | Pointer to **string** | The Kerberos principal name. It uses the typical Kerberos notation: &lt;SERVICE-NAME&gt;/&lt;server-domain-name&gt;@&lt;REALM&gt;.  Defaults to empty. | [optional] 
**DdnsClientUpdate** | Pointer to **string** | The global configuration to control who does the DDNS updates.  Valid values are: * _client_: DHCP server updates DNS if requested by client. * _server_: DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates. * _ignore_: DHCP server always updates DNS, even if the client says not to. * _over_client_update_: Same as _server_. DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates. * _over_no_update_: DHCP server updates DNS even if the client requests that no updates be done. If the client requests to do the update, DHCP server allows it.  Defaults to _client_. | [optional] 
**DdnsConflictResolutionMode** | Pointer to **string** | The mode used for resolving conflicts while performing DDNS updates.  Valid values are: * _check_with_dhcid_: It includes adding a DHCID record and checking that record via conflict detection as per RFC 4703. * _no_check_with_dhcid_: This will ignore conflict detection but add a DHCID record when creating/updating an entry. * _check_exists_with_dhcid_: This will check if there is an existing DHCID record but does not verify the value of the record matches the update. This will also update the DHCID record for the entry. * _no_check_without_dhcid_: This ignores conflict detection and will not add a DHCID record when creating/updating a DDNS entry.  Defaults to _check_with_dhcid_. | [optional] 
**DdnsDomain** | Pointer to **string** | The domain suffix for DDNS updates. FQDN, may be empty.  Must be specified if _ddns_enabled_ is _true_.  Defaults to empty. | [optional] 
**DdnsEnabled** | Pointer to **bool** | Indicates if DDNS updates should be performed for leases.  All other ddns_* configuration fields are ignored when this flag is unset.  At a minimum, _ddns_domain_ and _ddns_zones_ must be configured to enable DDNS.  Defaults to _false_. | [optional] 
**DdnsGenerateName** | Pointer to **bool** | Indicates if DDNS needs to generate a hostname when not supplied by the client.  Defaults to _false_. | [optional] 
**DdnsGeneratedPrefix** | Pointer to **string** | The prefix used in the generation of an FQDN.  When generating a name, DHCP server will construct the name in the format: [ddns-generated-prefix]-[address-text].[ddns-qualifying-suffix]. where address-text is simply the lease IP address converted to a hyphenated string.  Defaults to \&quot;myhost\&quot;. | [optional] 
**DdnsSendUpdates** | Pointer to **bool** | Determines if DDNS updates are enabled at the global level. Defaults to _true_. | [optional] 
**DdnsTtlPercent** | Pointer to **float32** | DDNS TTL value - to be calculated as a simple percentage of the lease&#39;s lifetime, using the parameter&#39;s value as the percentage. It is specified as a percentage (e.g. 25, 75). Defaults to unspecified. | [optional] 
**DdnsUpdateOnRenew** | Pointer to **bool** | Instructs the DHCP server to always update the DNS information when a lease is renewed even if its DNS information has not changed.  Defaults to _false_. | [optional] 
**DdnsUseConflictResolution** | Pointer to **bool** | When true, DHCP server will apply conflict resolution, as described in RFC 4703, when attempting to fulfill the update request.  When false, DHCP server will simply attempt to update the DNS entries per the request, regardless of whether or not they conflict with existing entries owned by other DHCP4 clients.  Defaults to _true_. | [optional] 
**DdnsZones** | Pointer to [**[]IpamsvcDDNSZone**](IpamsvcDDNSZone.md) | DNS zones that DDNS updates can be sent to. There is no resolver fallback. The target zone must be explicitly configured for the update to be performed.  Updates are sent to the closest enclosing zone.  Error if _ddns_enabled_ is _true_ and the _ddns_domain_ does not have a corresponding entry in _ddns_zones_.  Error if there are items with duplicate zone in the list.  Defaults to empty list. | [optional] 
**DhcpConfig** | Pointer to [**IpamsvcDHCPConfig**](IpamsvcDHCPConfig.md) |  | [optional] 
**DhcpOptions** | Pointer to [**[]IpamsvcOptionItem**](IpamsvcOptionItem.md) | The list of DHCP options or group of options for IPv4. An option list is ordered and may include both option groups and specific options. Multiple occurrences of the same option or group is not an error. The last occurrence of an option in the list will be used.  Error if the graph of referenced groups contains cycles.  Defaults to empty list. | [optional] 
**DhcpOptionsV6** | Pointer to [**[]IpamsvcOptionItem**](IpamsvcOptionItem.md) | The list of DHCP options or group of options for IPv6. An option list is ordered and may include both option groups and specific options. Multiple occurrences of the same option or group is not an error. The last occurrence of an option in the list will be used.  Error if the graph of referenced groups contains cycles.  Defaults to empty list. | [optional] 
**DhcpThreshold** | Pointer to [**IpamsvcDHCPUtilizationThreshold**](IpamsvcDHCPUtilizationThreshold.md) |  | [optional] 
**GssTsigFallback** | Pointer to **bool** | The behavior when GSS-TSIG should be used (a matching external DNS server is configured) but no GSS-TSIG key is available. If configured to _false_ (the default) this DNS server is skipped, if configured to _true_ the DNS server is ignored and the DNS update is sent with the configured DHCP-DDNS protection e.g. TSIG key or without any protection when none was configured.  Defaults to _false_. | [optional] 
**HeaderOptionFilename** | Pointer to **string** | The configuration for header option filename field. | [optional] 
**HeaderOptionServerAddress** | Pointer to **string** | The configuration for header option server address field. | [optional] 
**HeaderOptionServerName** | Pointer to **string** | The configuration for header option server name field. | [optional] 
**HostnameRewriteChar** | Pointer to **string** | The character to replace non-matching characters with, when hostname rewrite is enabled in global configuration.  Any single ASCII character or no character if the invalid characters should be removed without replacement.  Defaults to \&quot;-\&quot;. | [optional] 
**HostnameRewriteEnabled** | Pointer to **bool** | The global configuration to indicate if the hostnames supplied by the client will be rewritten prior to DDNS update by replacing every character that does not match _hostname_rewrite_regex_ by _hostname_rewrite_char_.  Defaults to _false_. | [optional] 
**HostnameRewriteRegex** | Pointer to **string** | The regex bracket expression to match valid characters when hostname rewrite is enabled in global configuration.  Must begin with \&quot;[\&quot; and end with \&quot;]\&quot; and be a compilable POSIX regex.  Defaults to \&quot;[^a-zA-Z0-9_.]\&quot;. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**KerberosKdc** | Pointer to **string** | Address of Kerberos Key Distribution Center.  Defaults to empty. | [optional] 
**KerberosKeys** | Pointer to [**[]IpamsvcKerberosKey**](IpamsvcKerberosKey.md) | _kerberos_keys_ contains a list of keys for GSS-TSIG signed dynamic updates.  Defaults to empty. | [optional] 
**KerberosRekeyInterval** | Pointer to **int64** | Time interval (in seconds) the keys for each configured external DNS server are checked for rekeying, i.e. a new key is created to replace the current usable one when its age is greater than the _kerberos_rekey_interval_ value.  Defaults to 120 seconds. | [optional] 
**KerberosRetryInterval** | Pointer to **int64** | Time interval (in seconds) to retry to create a key if any error occurred previously for any configured external DNS server.  Defaults to 30 seconds. | [optional] 
**KerberosTkeyLifetime** | Pointer to **int64** | Lifetime (in seconds) of GSS-TSIG keys in the TKEY protocol.  Defaults to 160 seconds. | [optional] 
**KerberosTkeyProtocol** | Pointer to **string** | Determines which protocol is used to establish the security context with the external DNS servers, TCP or UDP.  Defaults to _tcp_. | [optional] 
**PreferOption12** | Pointer to **bool** | When enabled, DHCP Server will prefer option 12 over option 81 in the incoming client request.  Defaults to _false_. | [optional] 
**RemoveSuffixOption81** | Pointer to **bool** | When enabled, DHCP Server will remove the suffix from the option 81 in the incoming client request.  Defaults to _false_. | [optional] 
**ServerPrincipal** | Pointer to **string** | The Kerberos principal name of the external DNS server that will receive updates.  Defaults to empty. | [optional] 
**VendorSpecificOptionOptionSpace** | Pointer to **string** | The resource identifier. | [optional] 

## Methods

### NewIpamsvcGlobal

`func NewIpamsvcGlobal() *IpamsvcGlobal`

NewIpamsvcGlobal instantiates a new IpamsvcGlobal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcGlobalWithDefaults

`func NewIpamsvcGlobalWithDefaults() *IpamsvcGlobal`

NewIpamsvcGlobalWithDefaults instantiates a new IpamsvcGlobal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsmConfig

`func (o *IpamsvcGlobal) GetAsmConfig() IpamsvcASMConfig`

GetAsmConfig returns the AsmConfig field if non-nil, zero value otherwise.

### GetAsmConfigOk

`func (o *IpamsvcGlobal) GetAsmConfigOk() (*IpamsvcASMConfig, bool)`

GetAsmConfigOk returns a tuple with the AsmConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmConfig

`func (o *IpamsvcGlobal) SetAsmConfig(v IpamsvcASMConfig)`

SetAsmConfig sets AsmConfig field to given value.

### HasAsmConfig

`func (o *IpamsvcGlobal) HasAsmConfig() bool`

HasAsmConfig returns a boolean if a field has been set.

### GetClientPrincipal

`func (o *IpamsvcGlobal) GetClientPrincipal() string`

GetClientPrincipal returns the ClientPrincipal field if non-nil, zero value otherwise.

### GetClientPrincipalOk

`func (o *IpamsvcGlobal) GetClientPrincipalOk() (*string, bool)`

GetClientPrincipalOk returns a tuple with the ClientPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrincipal

`func (o *IpamsvcGlobal) SetClientPrincipal(v string)`

SetClientPrincipal sets ClientPrincipal field to given value.

### HasClientPrincipal

`func (o *IpamsvcGlobal) HasClientPrincipal() bool`

HasClientPrincipal returns a boolean if a field has been set.

### GetDdnsClientUpdate

`func (o *IpamsvcGlobal) GetDdnsClientUpdate() string`

GetDdnsClientUpdate returns the DdnsClientUpdate field if non-nil, zero value otherwise.

### GetDdnsClientUpdateOk

`func (o *IpamsvcGlobal) GetDdnsClientUpdateOk() (*string, bool)`

GetDdnsClientUpdateOk returns a tuple with the DdnsClientUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsClientUpdate

`func (o *IpamsvcGlobal) SetDdnsClientUpdate(v string)`

SetDdnsClientUpdate sets DdnsClientUpdate field to given value.

### HasDdnsClientUpdate

`func (o *IpamsvcGlobal) HasDdnsClientUpdate() bool`

HasDdnsClientUpdate returns a boolean if a field has been set.

### GetDdnsConflictResolutionMode

`func (o *IpamsvcGlobal) GetDdnsConflictResolutionMode() string`

GetDdnsConflictResolutionMode returns the DdnsConflictResolutionMode field if non-nil, zero value otherwise.

### GetDdnsConflictResolutionModeOk

`func (o *IpamsvcGlobal) GetDdnsConflictResolutionModeOk() (*string, bool)`

GetDdnsConflictResolutionModeOk returns a tuple with the DdnsConflictResolutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsConflictResolutionMode

`func (o *IpamsvcGlobal) SetDdnsConflictResolutionMode(v string)`

SetDdnsConflictResolutionMode sets DdnsConflictResolutionMode field to given value.

### HasDdnsConflictResolutionMode

`func (o *IpamsvcGlobal) HasDdnsConflictResolutionMode() bool`

HasDdnsConflictResolutionMode returns a boolean if a field has been set.

### GetDdnsDomain

`func (o *IpamsvcGlobal) GetDdnsDomain() string`

GetDdnsDomain returns the DdnsDomain field if non-nil, zero value otherwise.

### GetDdnsDomainOk

`func (o *IpamsvcGlobal) GetDdnsDomainOk() (*string, bool)`

GetDdnsDomainOk returns a tuple with the DdnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDomain

`func (o *IpamsvcGlobal) SetDdnsDomain(v string)`

SetDdnsDomain sets DdnsDomain field to given value.

### HasDdnsDomain

`func (o *IpamsvcGlobal) HasDdnsDomain() bool`

HasDdnsDomain returns a boolean if a field has been set.

### GetDdnsEnabled

`func (o *IpamsvcGlobal) GetDdnsEnabled() bool`

GetDdnsEnabled returns the DdnsEnabled field if non-nil, zero value otherwise.

### GetDdnsEnabledOk

`func (o *IpamsvcGlobal) GetDdnsEnabledOk() (*bool, bool)`

GetDdnsEnabledOk returns a tuple with the DdnsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsEnabled

`func (o *IpamsvcGlobal) SetDdnsEnabled(v bool)`

SetDdnsEnabled sets DdnsEnabled field to given value.

### HasDdnsEnabled

`func (o *IpamsvcGlobal) HasDdnsEnabled() bool`

HasDdnsEnabled returns a boolean if a field has been set.

### GetDdnsGenerateName

`func (o *IpamsvcGlobal) GetDdnsGenerateName() bool`

GetDdnsGenerateName returns the DdnsGenerateName field if non-nil, zero value otherwise.

### GetDdnsGenerateNameOk

`func (o *IpamsvcGlobal) GetDdnsGenerateNameOk() (*bool, bool)`

GetDdnsGenerateNameOk returns a tuple with the DdnsGenerateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGenerateName

`func (o *IpamsvcGlobal) SetDdnsGenerateName(v bool)`

SetDdnsGenerateName sets DdnsGenerateName field to given value.

### HasDdnsGenerateName

`func (o *IpamsvcGlobal) HasDdnsGenerateName() bool`

HasDdnsGenerateName returns a boolean if a field has been set.

### GetDdnsGeneratedPrefix

`func (o *IpamsvcGlobal) GetDdnsGeneratedPrefix() string`

GetDdnsGeneratedPrefix returns the DdnsGeneratedPrefix field if non-nil, zero value otherwise.

### GetDdnsGeneratedPrefixOk

`func (o *IpamsvcGlobal) GetDdnsGeneratedPrefixOk() (*string, bool)`

GetDdnsGeneratedPrefixOk returns a tuple with the DdnsGeneratedPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGeneratedPrefix

`func (o *IpamsvcGlobal) SetDdnsGeneratedPrefix(v string)`

SetDdnsGeneratedPrefix sets DdnsGeneratedPrefix field to given value.

### HasDdnsGeneratedPrefix

`func (o *IpamsvcGlobal) HasDdnsGeneratedPrefix() bool`

HasDdnsGeneratedPrefix returns a boolean if a field has been set.

### GetDdnsSendUpdates

`func (o *IpamsvcGlobal) GetDdnsSendUpdates() bool`

GetDdnsSendUpdates returns the DdnsSendUpdates field if non-nil, zero value otherwise.

### GetDdnsSendUpdatesOk

`func (o *IpamsvcGlobal) GetDdnsSendUpdatesOk() (*bool, bool)`

GetDdnsSendUpdatesOk returns a tuple with the DdnsSendUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsSendUpdates

`func (o *IpamsvcGlobal) SetDdnsSendUpdates(v bool)`

SetDdnsSendUpdates sets DdnsSendUpdates field to given value.

### HasDdnsSendUpdates

`func (o *IpamsvcGlobal) HasDdnsSendUpdates() bool`

HasDdnsSendUpdates returns a boolean if a field has been set.

### GetDdnsTtlPercent

`func (o *IpamsvcGlobal) GetDdnsTtlPercent() float32`

GetDdnsTtlPercent returns the DdnsTtlPercent field if non-nil, zero value otherwise.

### GetDdnsTtlPercentOk

`func (o *IpamsvcGlobal) GetDdnsTtlPercentOk() (*float32, bool)`

GetDdnsTtlPercentOk returns a tuple with the DdnsTtlPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsTtlPercent

`func (o *IpamsvcGlobal) SetDdnsTtlPercent(v float32)`

SetDdnsTtlPercent sets DdnsTtlPercent field to given value.

### HasDdnsTtlPercent

`func (o *IpamsvcGlobal) HasDdnsTtlPercent() bool`

HasDdnsTtlPercent returns a boolean if a field has been set.

### GetDdnsUpdateOnRenew

`func (o *IpamsvcGlobal) GetDdnsUpdateOnRenew() bool`

GetDdnsUpdateOnRenew returns the DdnsUpdateOnRenew field if non-nil, zero value otherwise.

### GetDdnsUpdateOnRenewOk

`func (o *IpamsvcGlobal) GetDdnsUpdateOnRenewOk() (*bool, bool)`

GetDdnsUpdateOnRenewOk returns a tuple with the DdnsUpdateOnRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUpdateOnRenew

`func (o *IpamsvcGlobal) SetDdnsUpdateOnRenew(v bool)`

SetDdnsUpdateOnRenew sets DdnsUpdateOnRenew field to given value.

### HasDdnsUpdateOnRenew

`func (o *IpamsvcGlobal) HasDdnsUpdateOnRenew() bool`

HasDdnsUpdateOnRenew returns a boolean if a field has been set.

### GetDdnsUseConflictResolution

`func (o *IpamsvcGlobal) GetDdnsUseConflictResolution() bool`

GetDdnsUseConflictResolution returns the DdnsUseConflictResolution field if non-nil, zero value otherwise.

### GetDdnsUseConflictResolutionOk

`func (o *IpamsvcGlobal) GetDdnsUseConflictResolutionOk() (*bool, bool)`

GetDdnsUseConflictResolutionOk returns a tuple with the DdnsUseConflictResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUseConflictResolution

`func (o *IpamsvcGlobal) SetDdnsUseConflictResolution(v bool)`

SetDdnsUseConflictResolution sets DdnsUseConflictResolution field to given value.

### HasDdnsUseConflictResolution

`func (o *IpamsvcGlobal) HasDdnsUseConflictResolution() bool`

HasDdnsUseConflictResolution returns a boolean if a field has been set.

### GetDdnsZones

`func (o *IpamsvcGlobal) GetDdnsZones() []IpamsvcDDNSZone`

GetDdnsZones returns the DdnsZones field if non-nil, zero value otherwise.

### GetDdnsZonesOk

`func (o *IpamsvcGlobal) GetDdnsZonesOk() (*[]IpamsvcDDNSZone, bool)`

GetDdnsZonesOk returns a tuple with the DdnsZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsZones

`func (o *IpamsvcGlobal) SetDdnsZones(v []IpamsvcDDNSZone)`

SetDdnsZones sets DdnsZones field to given value.

### HasDdnsZones

`func (o *IpamsvcGlobal) HasDdnsZones() bool`

HasDdnsZones returns a boolean if a field has been set.

### GetDhcpConfig

`func (o *IpamsvcGlobal) GetDhcpConfig() IpamsvcDHCPConfig`

GetDhcpConfig returns the DhcpConfig field if non-nil, zero value otherwise.

### GetDhcpConfigOk

`func (o *IpamsvcGlobal) GetDhcpConfigOk() (*IpamsvcDHCPConfig, bool)`

GetDhcpConfigOk returns a tuple with the DhcpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpConfig

`func (o *IpamsvcGlobal) SetDhcpConfig(v IpamsvcDHCPConfig)`

SetDhcpConfig sets DhcpConfig field to given value.

### HasDhcpConfig

`func (o *IpamsvcGlobal) HasDhcpConfig() bool`

HasDhcpConfig returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *IpamsvcGlobal) GetDhcpOptions() []IpamsvcOptionItem`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *IpamsvcGlobal) GetDhcpOptionsOk() (*[]IpamsvcOptionItem, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *IpamsvcGlobal) SetDhcpOptions(v []IpamsvcOptionItem)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *IpamsvcGlobal) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetDhcpOptionsV6

`func (o *IpamsvcGlobal) GetDhcpOptionsV6() []IpamsvcOptionItem`

GetDhcpOptionsV6 returns the DhcpOptionsV6 field if non-nil, zero value otherwise.

### GetDhcpOptionsV6Ok

`func (o *IpamsvcGlobal) GetDhcpOptionsV6Ok() (*[]IpamsvcOptionItem, bool)`

GetDhcpOptionsV6Ok returns a tuple with the DhcpOptionsV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptionsV6

`func (o *IpamsvcGlobal) SetDhcpOptionsV6(v []IpamsvcOptionItem)`

SetDhcpOptionsV6 sets DhcpOptionsV6 field to given value.

### HasDhcpOptionsV6

`func (o *IpamsvcGlobal) HasDhcpOptionsV6() bool`

HasDhcpOptionsV6 returns a boolean if a field has been set.

### GetDhcpThreshold

`func (o *IpamsvcGlobal) GetDhcpThreshold() IpamsvcDHCPUtilizationThreshold`

GetDhcpThreshold returns the DhcpThreshold field if non-nil, zero value otherwise.

### GetDhcpThresholdOk

`func (o *IpamsvcGlobal) GetDhcpThresholdOk() (*IpamsvcDHCPUtilizationThreshold, bool)`

GetDhcpThresholdOk returns a tuple with the DhcpThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpThreshold

`func (o *IpamsvcGlobal) SetDhcpThreshold(v IpamsvcDHCPUtilizationThreshold)`

SetDhcpThreshold sets DhcpThreshold field to given value.

### HasDhcpThreshold

`func (o *IpamsvcGlobal) HasDhcpThreshold() bool`

HasDhcpThreshold returns a boolean if a field has been set.

### GetGssTsigFallback

`func (o *IpamsvcGlobal) GetGssTsigFallback() bool`

GetGssTsigFallback returns the GssTsigFallback field if non-nil, zero value otherwise.

### GetGssTsigFallbackOk

`func (o *IpamsvcGlobal) GetGssTsigFallbackOk() (*bool, bool)`

GetGssTsigFallbackOk returns a tuple with the GssTsigFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigFallback

`func (o *IpamsvcGlobal) SetGssTsigFallback(v bool)`

SetGssTsigFallback sets GssTsigFallback field to given value.

### HasGssTsigFallback

`func (o *IpamsvcGlobal) HasGssTsigFallback() bool`

HasGssTsigFallback returns a boolean if a field has been set.

### GetHeaderOptionFilename

`func (o *IpamsvcGlobal) GetHeaderOptionFilename() string`

GetHeaderOptionFilename returns the HeaderOptionFilename field if non-nil, zero value otherwise.

### GetHeaderOptionFilenameOk

`func (o *IpamsvcGlobal) GetHeaderOptionFilenameOk() (*string, bool)`

GetHeaderOptionFilenameOk returns a tuple with the HeaderOptionFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionFilename

`func (o *IpamsvcGlobal) SetHeaderOptionFilename(v string)`

SetHeaderOptionFilename sets HeaderOptionFilename field to given value.

### HasHeaderOptionFilename

`func (o *IpamsvcGlobal) HasHeaderOptionFilename() bool`

HasHeaderOptionFilename returns a boolean if a field has been set.

### GetHeaderOptionServerAddress

`func (o *IpamsvcGlobal) GetHeaderOptionServerAddress() string`

GetHeaderOptionServerAddress returns the HeaderOptionServerAddress field if non-nil, zero value otherwise.

### GetHeaderOptionServerAddressOk

`func (o *IpamsvcGlobal) GetHeaderOptionServerAddressOk() (*string, bool)`

GetHeaderOptionServerAddressOk returns a tuple with the HeaderOptionServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerAddress

`func (o *IpamsvcGlobal) SetHeaderOptionServerAddress(v string)`

SetHeaderOptionServerAddress sets HeaderOptionServerAddress field to given value.

### HasHeaderOptionServerAddress

`func (o *IpamsvcGlobal) HasHeaderOptionServerAddress() bool`

HasHeaderOptionServerAddress returns a boolean if a field has been set.

### GetHeaderOptionServerName

`func (o *IpamsvcGlobal) GetHeaderOptionServerName() string`

GetHeaderOptionServerName returns the HeaderOptionServerName field if non-nil, zero value otherwise.

### GetHeaderOptionServerNameOk

`func (o *IpamsvcGlobal) GetHeaderOptionServerNameOk() (*string, bool)`

GetHeaderOptionServerNameOk returns a tuple with the HeaderOptionServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerName

`func (o *IpamsvcGlobal) SetHeaderOptionServerName(v string)`

SetHeaderOptionServerName sets HeaderOptionServerName field to given value.

### HasHeaderOptionServerName

`func (o *IpamsvcGlobal) HasHeaderOptionServerName() bool`

HasHeaderOptionServerName returns a boolean if a field has been set.

### GetHostnameRewriteChar

`func (o *IpamsvcGlobal) GetHostnameRewriteChar() string`

GetHostnameRewriteChar returns the HostnameRewriteChar field if non-nil, zero value otherwise.

### GetHostnameRewriteCharOk

`func (o *IpamsvcGlobal) GetHostnameRewriteCharOk() (*string, bool)`

GetHostnameRewriteCharOk returns a tuple with the HostnameRewriteChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameRewriteChar

`func (o *IpamsvcGlobal) SetHostnameRewriteChar(v string)`

SetHostnameRewriteChar sets HostnameRewriteChar field to given value.

### HasHostnameRewriteChar

`func (o *IpamsvcGlobal) HasHostnameRewriteChar() bool`

HasHostnameRewriteChar returns a boolean if a field has been set.

### GetHostnameRewriteEnabled

`func (o *IpamsvcGlobal) GetHostnameRewriteEnabled() bool`

GetHostnameRewriteEnabled returns the HostnameRewriteEnabled field if non-nil, zero value otherwise.

### GetHostnameRewriteEnabledOk

`func (o *IpamsvcGlobal) GetHostnameRewriteEnabledOk() (*bool, bool)`

GetHostnameRewriteEnabledOk returns a tuple with the HostnameRewriteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameRewriteEnabled

`func (o *IpamsvcGlobal) SetHostnameRewriteEnabled(v bool)`

SetHostnameRewriteEnabled sets HostnameRewriteEnabled field to given value.

### HasHostnameRewriteEnabled

`func (o *IpamsvcGlobal) HasHostnameRewriteEnabled() bool`

HasHostnameRewriteEnabled returns a boolean if a field has been set.

### GetHostnameRewriteRegex

`func (o *IpamsvcGlobal) GetHostnameRewriteRegex() string`

GetHostnameRewriteRegex returns the HostnameRewriteRegex field if non-nil, zero value otherwise.

### GetHostnameRewriteRegexOk

`func (o *IpamsvcGlobal) GetHostnameRewriteRegexOk() (*string, bool)`

GetHostnameRewriteRegexOk returns a tuple with the HostnameRewriteRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameRewriteRegex

`func (o *IpamsvcGlobal) SetHostnameRewriteRegex(v string)`

SetHostnameRewriteRegex sets HostnameRewriteRegex field to given value.

### HasHostnameRewriteRegex

`func (o *IpamsvcGlobal) HasHostnameRewriteRegex() bool`

HasHostnameRewriteRegex returns a boolean if a field has been set.

### GetId

`func (o *IpamsvcGlobal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpamsvcGlobal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpamsvcGlobal) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpamsvcGlobal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKerberosKdc

`func (o *IpamsvcGlobal) GetKerberosKdc() string`

GetKerberosKdc returns the KerberosKdc field if non-nil, zero value otherwise.

### GetKerberosKdcOk

`func (o *IpamsvcGlobal) GetKerberosKdcOk() (*string, bool)`

GetKerberosKdcOk returns a tuple with the KerberosKdc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosKdc

`func (o *IpamsvcGlobal) SetKerberosKdc(v string)`

SetKerberosKdc sets KerberosKdc field to given value.

### HasKerberosKdc

`func (o *IpamsvcGlobal) HasKerberosKdc() bool`

HasKerberosKdc returns a boolean if a field has been set.

### GetKerberosKeys

`func (o *IpamsvcGlobal) GetKerberosKeys() []IpamsvcKerberosKey`

GetKerberosKeys returns the KerberosKeys field if non-nil, zero value otherwise.

### GetKerberosKeysOk

`func (o *IpamsvcGlobal) GetKerberosKeysOk() (*[]IpamsvcKerberosKey, bool)`

GetKerberosKeysOk returns a tuple with the KerberosKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosKeys

`func (o *IpamsvcGlobal) SetKerberosKeys(v []IpamsvcKerberosKey)`

SetKerberosKeys sets KerberosKeys field to given value.

### HasKerberosKeys

`func (o *IpamsvcGlobal) HasKerberosKeys() bool`

HasKerberosKeys returns a boolean if a field has been set.

### GetKerberosRekeyInterval

`func (o *IpamsvcGlobal) GetKerberosRekeyInterval() int64`

GetKerberosRekeyInterval returns the KerberosRekeyInterval field if non-nil, zero value otherwise.

### GetKerberosRekeyIntervalOk

`func (o *IpamsvcGlobal) GetKerberosRekeyIntervalOk() (*int64, bool)`

GetKerberosRekeyIntervalOk returns a tuple with the KerberosRekeyInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosRekeyInterval

`func (o *IpamsvcGlobal) SetKerberosRekeyInterval(v int64)`

SetKerberosRekeyInterval sets KerberosRekeyInterval field to given value.

### HasKerberosRekeyInterval

`func (o *IpamsvcGlobal) HasKerberosRekeyInterval() bool`

HasKerberosRekeyInterval returns a boolean if a field has been set.

### GetKerberosRetryInterval

`func (o *IpamsvcGlobal) GetKerberosRetryInterval() int64`

GetKerberosRetryInterval returns the KerberosRetryInterval field if non-nil, zero value otherwise.

### GetKerberosRetryIntervalOk

`func (o *IpamsvcGlobal) GetKerberosRetryIntervalOk() (*int64, bool)`

GetKerberosRetryIntervalOk returns a tuple with the KerberosRetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosRetryInterval

`func (o *IpamsvcGlobal) SetKerberosRetryInterval(v int64)`

SetKerberosRetryInterval sets KerberosRetryInterval field to given value.

### HasKerberosRetryInterval

`func (o *IpamsvcGlobal) HasKerberosRetryInterval() bool`

HasKerberosRetryInterval returns a boolean if a field has been set.

### GetKerberosTkeyLifetime

`func (o *IpamsvcGlobal) GetKerberosTkeyLifetime() int64`

GetKerberosTkeyLifetime returns the KerberosTkeyLifetime field if non-nil, zero value otherwise.

### GetKerberosTkeyLifetimeOk

`func (o *IpamsvcGlobal) GetKerberosTkeyLifetimeOk() (*int64, bool)`

GetKerberosTkeyLifetimeOk returns a tuple with the KerberosTkeyLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosTkeyLifetime

`func (o *IpamsvcGlobal) SetKerberosTkeyLifetime(v int64)`

SetKerberosTkeyLifetime sets KerberosTkeyLifetime field to given value.

### HasKerberosTkeyLifetime

`func (o *IpamsvcGlobal) HasKerberosTkeyLifetime() bool`

HasKerberosTkeyLifetime returns a boolean if a field has been set.

### GetKerberosTkeyProtocol

`func (o *IpamsvcGlobal) GetKerberosTkeyProtocol() string`

GetKerberosTkeyProtocol returns the KerberosTkeyProtocol field if non-nil, zero value otherwise.

### GetKerberosTkeyProtocolOk

`func (o *IpamsvcGlobal) GetKerberosTkeyProtocolOk() (*string, bool)`

GetKerberosTkeyProtocolOk returns a tuple with the KerberosTkeyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosTkeyProtocol

`func (o *IpamsvcGlobal) SetKerberosTkeyProtocol(v string)`

SetKerberosTkeyProtocol sets KerberosTkeyProtocol field to given value.

### HasKerberosTkeyProtocol

`func (o *IpamsvcGlobal) HasKerberosTkeyProtocol() bool`

HasKerberosTkeyProtocol returns a boolean if a field has been set.

### GetPreferOption12

`func (o *IpamsvcGlobal) GetPreferOption12() bool`

GetPreferOption12 returns the PreferOption12 field if non-nil, zero value otherwise.

### GetPreferOption12Ok

`func (o *IpamsvcGlobal) GetPreferOption12Ok() (*bool, bool)`

GetPreferOption12Ok returns a tuple with the PreferOption12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferOption12

`func (o *IpamsvcGlobal) SetPreferOption12(v bool)`

SetPreferOption12 sets PreferOption12 field to given value.

### HasPreferOption12

`func (o *IpamsvcGlobal) HasPreferOption12() bool`

HasPreferOption12 returns a boolean if a field has been set.

### GetRemoveSuffixOption81

`func (o *IpamsvcGlobal) GetRemoveSuffixOption81() bool`

GetRemoveSuffixOption81 returns the RemoveSuffixOption81 field if non-nil, zero value otherwise.

### GetRemoveSuffixOption81Ok

`func (o *IpamsvcGlobal) GetRemoveSuffixOption81Ok() (*bool, bool)`

GetRemoveSuffixOption81Ok returns a tuple with the RemoveSuffixOption81 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveSuffixOption81

`func (o *IpamsvcGlobal) SetRemoveSuffixOption81(v bool)`

SetRemoveSuffixOption81 sets RemoveSuffixOption81 field to given value.

### HasRemoveSuffixOption81

`func (o *IpamsvcGlobal) HasRemoveSuffixOption81() bool`

HasRemoveSuffixOption81 returns a boolean if a field has been set.

### GetServerPrincipal

`func (o *IpamsvcGlobal) GetServerPrincipal() string`

GetServerPrincipal returns the ServerPrincipal field if non-nil, zero value otherwise.

### GetServerPrincipalOk

`func (o *IpamsvcGlobal) GetServerPrincipalOk() (*string, bool)`

GetServerPrincipalOk returns a tuple with the ServerPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPrincipal

`func (o *IpamsvcGlobal) SetServerPrincipal(v string)`

SetServerPrincipal sets ServerPrincipal field to given value.

### HasServerPrincipal

`func (o *IpamsvcGlobal) HasServerPrincipal() bool`

HasServerPrincipal returns a boolean if a field has been set.

### GetVendorSpecificOptionOptionSpace

`func (o *IpamsvcGlobal) GetVendorSpecificOptionOptionSpace() string`

GetVendorSpecificOptionOptionSpace returns the VendorSpecificOptionOptionSpace field if non-nil, zero value otherwise.

### GetVendorSpecificOptionOptionSpaceOk

`func (o *IpamsvcGlobal) GetVendorSpecificOptionOptionSpaceOk() (*string, bool)`

GetVendorSpecificOptionOptionSpaceOk returns a tuple with the VendorSpecificOptionOptionSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSpecificOptionOptionSpace

`func (o *IpamsvcGlobal) SetVendorSpecificOptionOptionSpace(v string)`

SetVendorSpecificOptionOptionSpace sets VendorSpecificOptionOptionSpace field to given value.

### HasVendorSpecificOptionOptionSpace

`func (o *IpamsvcGlobal) HasVendorSpecificOptionOptionSpace() bool`

HasVendorSpecificOptionOptionSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


