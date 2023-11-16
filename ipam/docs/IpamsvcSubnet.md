# IpamsvcSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The address of the subnet in the form “a.b.c.d/n” where the “/n” may be omitted. In this case, the CIDR value must be defined in the _cidr_ field. When reading, the _address_ field is always in the form “a.b.c.d”. | [optional] 
**AsmConfig** | Pointer to [**IpamsvcASMConfig**](IpamsvcASMConfig.md) |  | [optional] 
**AsmScopeFlag** | Pointer to **int64** | Set to 1 to indicate that the subnet may run out of addresses. | [optional] [readonly] 
**Cidr** | Pointer to **int64** | The CIDR of the subnet. This is required if _address_ does not include CIDR. | [optional] 
**Comment** | Pointer to **string** | The description for the subnet. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**DdnsClientUpdate** | Pointer to **string** | Controls who does the DDNS updates.  Valid values are: * _client_: DHCP server updates DNS if requested by client. * _server_: DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates. * _ignore_: DHCP server always updates DNS, even if the client says not to. * _over_client_update_: Same as _server_. DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates. * _over_no_update_: DHCP server updates DNS even if the client requests that no updates be done. If the client requests to do the update, DHCP server allows it.  Defaults to _client_. | [optional] 
**DdnsConflictResolutionMode** | Pointer to **string** | The mode used for resolving conflicts while performing DDNS updates.  Valid values are: * _check_with_dhcid_: It includes adding a DHCID record and checking that record via conflict detection as per RFC 4703. * _no_check_with_dhcid_: This will ignore conflict detection but add a DHCID record when creating/updating an entry. * _check_exists_with_dhcid_: This will check if there is an existing DHCID record but does not verify the value of the record matches the update. This will also update the DHCID record for the entry. * _no_check_without_dhcid_: This ignores conflict detection and will not add a DHCID record when creating/updating a DDNS entry.  Defaults to _check_with_dhcid_. | [optional] 
**DdnsDomain** | Pointer to **string** | The domain suffix for DDNS updates. FQDN, may be empty.  Defaults to empty. | [optional] 
**DdnsGenerateName** | Pointer to **bool** | Indicates if DDNS needs to generate a hostname when not supplied by the client.  Defaults to _false_. | [optional] 
**DdnsGeneratedPrefix** | Pointer to **string** | The prefix used in the generation of an FQDN.  When generating a name, DHCP server will construct the name in the format: [ddns-generated-prefix]-[address-text].[ddns-qualifying-suffix]. where address-text is simply the lease IP address converted to a hyphenated string.  Defaults to \&quot;myhost\&quot;. | [optional] 
**DdnsSendUpdates** | Pointer to **bool** | Determines if DDNS updates are enabled at the subnet level. Defaults to _true_. | [optional] 
**DdnsTtlPercent** | Pointer to **float32** | DDNS TTL value - to be calculated as a simple percentage of the lease&#39;s lifetime, using the parameter&#39;s value as the percentage. It is specified as a percentage (e.g. 25, 75). Defaults to unspecified. | [optional] 
**DdnsUpdateOnRenew** | Pointer to **bool** | Instructs the DHCP server to always update the DNS information when a lease is renewed even if its DNS information has not changed.  Defaults to _false_. | [optional] 
**DdnsUseConflictResolution** | Pointer to **bool** | When true, DHCP server will apply conflict resolution, as described in RFC 4703, when attempting to fulfill the update request.  When false, DHCP server will simply attempt to update the DNS entries per the request, regardless of whether or not they conflict with existing entries owned by other DHCP4 clients.  Defaults to _true_. | [optional] 
**DhcpConfig** | Pointer to [**IpamsvcDHCPConfig**](IpamsvcDHCPConfig.md) |  | [optional] 
**DhcpHost** | Pointer to **string** | The resource identifier. | [optional] 
**DhcpOptions** | Pointer to [**[]IpamsvcOptionItem**](IpamsvcOptionItem.md) | The DHCP options of the subnet. This can either be a specific option or a group of options. | [optional] 
**DhcpUtilization** | Pointer to [**IpamsvcDHCPUtilization**](IpamsvcDHCPUtilization.md) |  | [optional] 
**DisableDhcp** | Pointer to **bool** | Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.  Defaults to _false_. | [optional] 
**DiscoveryAttrs** | Pointer to **map[string]interface{}** | The discovery attributes for this subnet in JSON format. | [optional] [readonly] 
**DiscoveryMetadata** | Pointer to **map[string]interface{}** | The discovery metadata for this subnet in JSON format. | [optional] [readonly] 
**HeaderOptionFilename** | Pointer to **string** | The configuration for header option filename field. | [optional] 
**HeaderOptionServerAddress** | Pointer to **string** | The configuration for header option server address field. | [optional] 
**HeaderOptionServerName** | Pointer to **string** | The configuration for header option server name field. | [optional] 
**HostnameRewriteChar** | Pointer to **string** | The character to replace non-matching characters with, when hostname rewrite is enabled.  Any single ASCII character or no character if the invalid characters should be removed without replacement.  Defaults to \&quot;-\&quot;. | [optional] 
**HostnameRewriteEnabled** | Pointer to **bool** | Indicates if client supplied hostnames will be rewritten prior to DDNS update by replacing every character that does not match _hostname_rewrite_regex_ by _hostname_rewrite_char_.  Defaults to _false_. | [optional] 
**HostnameRewriteRegex** | Pointer to **string** | The regex bracket expression to match valid characters.  Must begin with \&quot;[\&quot; and end with \&quot;]\&quot; and be a compilable POSIX regex.  Defaults to \&quot;[^a-zA-Z0-9_.]\&quot;. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InheritanceAssignedHosts** | Pointer to [**[]InheritanceAssignedHost**](InheritanceAssignedHost.md) | The list of the inheritance assigned hosts of the object. | [optional] [readonly] 
**InheritanceParent** | Pointer to **string** | The resource identifier. | [optional] 
**InheritanceSources** | Pointer to [**IpamsvcDHCPInheritance**](IpamsvcDHCPInheritance.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the subnet. May contain 1 to 256 characters. Can include UTF-8. | [optional] 
**Parent** | Pointer to **string** | The resource identifier. | [optional] 
**Protocol** | Pointer to **string** | The type of protocol of the subnet (_ip4_ or _ip6_). | [optional] [readonly] 
**RebindTime** | Pointer to **int64** | The lease rebind time (T2) in seconds. | [optional] 
**RenewTime** | Pointer to **int64** | The lease renew time (T1) in seconds. | [optional] 
**Space** | Pointer to **string** | The resource identifier. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | The tags for the subnet in JSON format. | [optional] 
**Threshold** | Pointer to [**IpamsvcUtilizationThreshold**](IpamsvcUtilizationThreshold.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 
**Usage** | Pointer to **[]string** | The usage is a combination of indicators, each tracking a specific associated use. Listed below are usage indicators with their meaning:  usage indicator        | description  ---------------------- | --------------------------------  _IPAM_                 |  Subnet is managed in BloxOne DDI.  _DHCP_                 |  Subnet is served by a DHCP Host.  _DISCOVERED_           |  Subnet is discovered by some network discovery probe like Network Insight or NetMRI in NIOS. | [optional] [readonly] 
**Utilization** | Pointer to [**IpamsvcUtilization**](IpamsvcUtilization.md) |  | [optional] 
**UtilizationV6** | Pointer to [**IpamsvcUtilizationV6**](IpamsvcUtilizationV6.md) |  | [optional] 

## Methods

### NewIpamsvcSubnet

`func NewIpamsvcSubnet() *IpamsvcSubnet`

NewIpamsvcSubnet instantiates a new IpamsvcSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcSubnetWithDefaults

`func NewIpamsvcSubnetWithDefaults() *IpamsvcSubnet`

NewIpamsvcSubnetWithDefaults instantiates a new IpamsvcSubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *IpamsvcSubnet) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IpamsvcSubnet) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IpamsvcSubnet) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IpamsvcSubnet) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAsmConfig

`func (o *IpamsvcSubnet) GetAsmConfig() IpamsvcASMConfig`

GetAsmConfig returns the AsmConfig field if non-nil, zero value otherwise.

### GetAsmConfigOk

`func (o *IpamsvcSubnet) GetAsmConfigOk() (*IpamsvcASMConfig, bool)`

GetAsmConfigOk returns a tuple with the AsmConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmConfig

`func (o *IpamsvcSubnet) SetAsmConfig(v IpamsvcASMConfig)`

SetAsmConfig sets AsmConfig field to given value.

### HasAsmConfig

`func (o *IpamsvcSubnet) HasAsmConfig() bool`

HasAsmConfig returns a boolean if a field has been set.

### GetAsmScopeFlag

`func (o *IpamsvcSubnet) GetAsmScopeFlag() int64`

GetAsmScopeFlag returns the AsmScopeFlag field if non-nil, zero value otherwise.

### GetAsmScopeFlagOk

`func (o *IpamsvcSubnet) GetAsmScopeFlagOk() (*int64, bool)`

GetAsmScopeFlagOk returns a tuple with the AsmScopeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmScopeFlag

`func (o *IpamsvcSubnet) SetAsmScopeFlag(v int64)`

SetAsmScopeFlag sets AsmScopeFlag field to given value.

### HasAsmScopeFlag

`func (o *IpamsvcSubnet) HasAsmScopeFlag() bool`

HasAsmScopeFlag returns a boolean if a field has been set.

### GetCidr

`func (o *IpamsvcSubnet) GetCidr() int64`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *IpamsvcSubnet) GetCidrOk() (*int64, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *IpamsvcSubnet) SetCidr(v int64)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *IpamsvcSubnet) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetComment

`func (o *IpamsvcSubnet) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IpamsvcSubnet) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IpamsvcSubnet) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IpamsvcSubnet) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IpamsvcSubnet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IpamsvcSubnet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IpamsvcSubnet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IpamsvcSubnet) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDdnsClientUpdate

`func (o *IpamsvcSubnet) GetDdnsClientUpdate() string`

GetDdnsClientUpdate returns the DdnsClientUpdate field if non-nil, zero value otherwise.

### GetDdnsClientUpdateOk

`func (o *IpamsvcSubnet) GetDdnsClientUpdateOk() (*string, bool)`

GetDdnsClientUpdateOk returns a tuple with the DdnsClientUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsClientUpdate

`func (o *IpamsvcSubnet) SetDdnsClientUpdate(v string)`

SetDdnsClientUpdate sets DdnsClientUpdate field to given value.

### HasDdnsClientUpdate

`func (o *IpamsvcSubnet) HasDdnsClientUpdate() bool`

HasDdnsClientUpdate returns a boolean if a field has been set.

### GetDdnsConflictResolutionMode

`func (o *IpamsvcSubnet) GetDdnsConflictResolutionMode() string`

GetDdnsConflictResolutionMode returns the DdnsConflictResolutionMode field if non-nil, zero value otherwise.

### GetDdnsConflictResolutionModeOk

`func (o *IpamsvcSubnet) GetDdnsConflictResolutionModeOk() (*string, bool)`

GetDdnsConflictResolutionModeOk returns a tuple with the DdnsConflictResolutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsConflictResolutionMode

`func (o *IpamsvcSubnet) SetDdnsConflictResolutionMode(v string)`

SetDdnsConflictResolutionMode sets DdnsConflictResolutionMode field to given value.

### HasDdnsConflictResolutionMode

`func (o *IpamsvcSubnet) HasDdnsConflictResolutionMode() bool`

HasDdnsConflictResolutionMode returns a boolean if a field has been set.

### GetDdnsDomain

`func (o *IpamsvcSubnet) GetDdnsDomain() string`

GetDdnsDomain returns the DdnsDomain field if non-nil, zero value otherwise.

### GetDdnsDomainOk

`func (o *IpamsvcSubnet) GetDdnsDomainOk() (*string, bool)`

GetDdnsDomainOk returns a tuple with the DdnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDomain

`func (o *IpamsvcSubnet) SetDdnsDomain(v string)`

SetDdnsDomain sets DdnsDomain field to given value.

### HasDdnsDomain

`func (o *IpamsvcSubnet) HasDdnsDomain() bool`

HasDdnsDomain returns a boolean if a field has been set.

### GetDdnsGenerateName

`func (o *IpamsvcSubnet) GetDdnsGenerateName() bool`

GetDdnsGenerateName returns the DdnsGenerateName field if non-nil, zero value otherwise.

### GetDdnsGenerateNameOk

`func (o *IpamsvcSubnet) GetDdnsGenerateNameOk() (*bool, bool)`

GetDdnsGenerateNameOk returns a tuple with the DdnsGenerateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGenerateName

`func (o *IpamsvcSubnet) SetDdnsGenerateName(v bool)`

SetDdnsGenerateName sets DdnsGenerateName field to given value.

### HasDdnsGenerateName

`func (o *IpamsvcSubnet) HasDdnsGenerateName() bool`

HasDdnsGenerateName returns a boolean if a field has been set.

### GetDdnsGeneratedPrefix

`func (o *IpamsvcSubnet) GetDdnsGeneratedPrefix() string`

GetDdnsGeneratedPrefix returns the DdnsGeneratedPrefix field if non-nil, zero value otherwise.

### GetDdnsGeneratedPrefixOk

`func (o *IpamsvcSubnet) GetDdnsGeneratedPrefixOk() (*string, bool)`

GetDdnsGeneratedPrefixOk returns a tuple with the DdnsGeneratedPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGeneratedPrefix

`func (o *IpamsvcSubnet) SetDdnsGeneratedPrefix(v string)`

SetDdnsGeneratedPrefix sets DdnsGeneratedPrefix field to given value.

### HasDdnsGeneratedPrefix

`func (o *IpamsvcSubnet) HasDdnsGeneratedPrefix() bool`

HasDdnsGeneratedPrefix returns a boolean if a field has been set.

### GetDdnsSendUpdates

`func (o *IpamsvcSubnet) GetDdnsSendUpdates() bool`

GetDdnsSendUpdates returns the DdnsSendUpdates field if non-nil, zero value otherwise.

### GetDdnsSendUpdatesOk

`func (o *IpamsvcSubnet) GetDdnsSendUpdatesOk() (*bool, bool)`

GetDdnsSendUpdatesOk returns a tuple with the DdnsSendUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsSendUpdates

`func (o *IpamsvcSubnet) SetDdnsSendUpdates(v bool)`

SetDdnsSendUpdates sets DdnsSendUpdates field to given value.

### HasDdnsSendUpdates

`func (o *IpamsvcSubnet) HasDdnsSendUpdates() bool`

HasDdnsSendUpdates returns a boolean if a field has been set.

### GetDdnsTtlPercent

`func (o *IpamsvcSubnet) GetDdnsTtlPercent() float32`

GetDdnsTtlPercent returns the DdnsTtlPercent field if non-nil, zero value otherwise.

### GetDdnsTtlPercentOk

`func (o *IpamsvcSubnet) GetDdnsTtlPercentOk() (*float32, bool)`

GetDdnsTtlPercentOk returns a tuple with the DdnsTtlPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsTtlPercent

`func (o *IpamsvcSubnet) SetDdnsTtlPercent(v float32)`

SetDdnsTtlPercent sets DdnsTtlPercent field to given value.

### HasDdnsTtlPercent

`func (o *IpamsvcSubnet) HasDdnsTtlPercent() bool`

HasDdnsTtlPercent returns a boolean if a field has been set.

### GetDdnsUpdateOnRenew

`func (o *IpamsvcSubnet) GetDdnsUpdateOnRenew() bool`

GetDdnsUpdateOnRenew returns the DdnsUpdateOnRenew field if non-nil, zero value otherwise.

### GetDdnsUpdateOnRenewOk

`func (o *IpamsvcSubnet) GetDdnsUpdateOnRenewOk() (*bool, bool)`

GetDdnsUpdateOnRenewOk returns a tuple with the DdnsUpdateOnRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUpdateOnRenew

`func (o *IpamsvcSubnet) SetDdnsUpdateOnRenew(v bool)`

SetDdnsUpdateOnRenew sets DdnsUpdateOnRenew field to given value.

### HasDdnsUpdateOnRenew

`func (o *IpamsvcSubnet) HasDdnsUpdateOnRenew() bool`

HasDdnsUpdateOnRenew returns a boolean if a field has been set.

### GetDdnsUseConflictResolution

`func (o *IpamsvcSubnet) GetDdnsUseConflictResolution() bool`

GetDdnsUseConflictResolution returns the DdnsUseConflictResolution field if non-nil, zero value otherwise.

### GetDdnsUseConflictResolutionOk

`func (o *IpamsvcSubnet) GetDdnsUseConflictResolutionOk() (*bool, bool)`

GetDdnsUseConflictResolutionOk returns a tuple with the DdnsUseConflictResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUseConflictResolution

`func (o *IpamsvcSubnet) SetDdnsUseConflictResolution(v bool)`

SetDdnsUseConflictResolution sets DdnsUseConflictResolution field to given value.

### HasDdnsUseConflictResolution

`func (o *IpamsvcSubnet) HasDdnsUseConflictResolution() bool`

HasDdnsUseConflictResolution returns a boolean if a field has been set.

### GetDhcpConfig

`func (o *IpamsvcSubnet) GetDhcpConfig() IpamsvcDHCPConfig`

GetDhcpConfig returns the DhcpConfig field if non-nil, zero value otherwise.

### GetDhcpConfigOk

`func (o *IpamsvcSubnet) GetDhcpConfigOk() (*IpamsvcDHCPConfig, bool)`

GetDhcpConfigOk returns a tuple with the DhcpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpConfig

`func (o *IpamsvcSubnet) SetDhcpConfig(v IpamsvcDHCPConfig)`

SetDhcpConfig sets DhcpConfig field to given value.

### HasDhcpConfig

`func (o *IpamsvcSubnet) HasDhcpConfig() bool`

HasDhcpConfig returns a boolean if a field has been set.

### GetDhcpHost

`func (o *IpamsvcSubnet) GetDhcpHost() string`

GetDhcpHost returns the DhcpHost field if non-nil, zero value otherwise.

### GetDhcpHostOk

`func (o *IpamsvcSubnet) GetDhcpHostOk() (*string, bool)`

GetDhcpHostOk returns a tuple with the DhcpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpHost

`func (o *IpamsvcSubnet) SetDhcpHost(v string)`

SetDhcpHost sets DhcpHost field to given value.

### HasDhcpHost

`func (o *IpamsvcSubnet) HasDhcpHost() bool`

HasDhcpHost returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *IpamsvcSubnet) GetDhcpOptions() []IpamsvcOptionItem`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *IpamsvcSubnet) GetDhcpOptionsOk() (*[]IpamsvcOptionItem, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *IpamsvcSubnet) SetDhcpOptions(v []IpamsvcOptionItem)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *IpamsvcSubnet) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetDhcpUtilization

`func (o *IpamsvcSubnet) GetDhcpUtilization() IpamsvcDHCPUtilization`

GetDhcpUtilization returns the DhcpUtilization field if non-nil, zero value otherwise.

### GetDhcpUtilizationOk

`func (o *IpamsvcSubnet) GetDhcpUtilizationOk() (*IpamsvcDHCPUtilization, bool)`

GetDhcpUtilizationOk returns a tuple with the DhcpUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilization

`func (o *IpamsvcSubnet) SetDhcpUtilization(v IpamsvcDHCPUtilization)`

SetDhcpUtilization sets DhcpUtilization field to given value.

### HasDhcpUtilization

`func (o *IpamsvcSubnet) HasDhcpUtilization() bool`

HasDhcpUtilization returns a boolean if a field has been set.

### GetDisableDhcp

`func (o *IpamsvcSubnet) GetDisableDhcp() bool`

GetDisableDhcp returns the DisableDhcp field if non-nil, zero value otherwise.

### GetDisableDhcpOk

`func (o *IpamsvcSubnet) GetDisableDhcpOk() (*bool, bool)`

GetDisableDhcpOk returns a tuple with the DisableDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDhcp

`func (o *IpamsvcSubnet) SetDisableDhcp(v bool)`

SetDisableDhcp sets DisableDhcp field to given value.

### HasDisableDhcp

`func (o *IpamsvcSubnet) HasDisableDhcp() bool`

HasDisableDhcp returns a boolean if a field has been set.

### GetDiscoveryAttrs

`func (o *IpamsvcSubnet) GetDiscoveryAttrs() map[string]interface{}`

GetDiscoveryAttrs returns the DiscoveryAttrs field if non-nil, zero value otherwise.

### GetDiscoveryAttrsOk

`func (o *IpamsvcSubnet) GetDiscoveryAttrsOk() (*map[string]interface{}, bool)`

GetDiscoveryAttrsOk returns a tuple with the DiscoveryAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryAttrs

`func (o *IpamsvcSubnet) SetDiscoveryAttrs(v map[string]interface{})`

SetDiscoveryAttrs sets DiscoveryAttrs field to given value.

### HasDiscoveryAttrs

`func (o *IpamsvcSubnet) HasDiscoveryAttrs() bool`

HasDiscoveryAttrs returns a boolean if a field has been set.

### GetDiscoveryMetadata

`func (o *IpamsvcSubnet) GetDiscoveryMetadata() map[string]interface{}`

GetDiscoveryMetadata returns the DiscoveryMetadata field if non-nil, zero value otherwise.

### GetDiscoveryMetadataOk

`func (o *IpamsvcSubnet) GetDiscoveryMetadataOk() (*map[string]interface{}, bool)`

GetDiscoveryMetadataOk returns a tuple with the DiscoveryMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryMetadata

`func (o *IpamsvcSubnet) SetDiscoveryMetadata(v map[string]interface{})`

SetDiscoveryMetadata sets DiscoveryMetadata field to given value.

### HasDiscoveryMetadata

`func (o *IpamsvcSubnet) HasDiscoveryMetadata() bool`

HasDiscoveryMetadata returns a boolean if a field has been set.

### GetHeaderOptionFilename

`func (o *IpamsvcSubnet) GetHeaderOptionFilename() string`

GetHeaderOptionFilename returns the HeaderOptionFilename field if non-nil, zero value otherwise.

### GetHeaderOptionFilenameOk

`func (o *IpamsvcSubnet) GetHeaderOptionFilenameOk() (*string, bool)`

GetHeaderOptionFilenameOk returns a tuple with the HeaderOptionFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionFilename

`func (o *IpamsvcSubnet) SetHeaderOptionFilename(v string)`

SetHeaderOptionFilename sets HeaderOptionFilename field to given value.

### HasHeaderOptionFilename

`func (o *IpamsvcSubnet) HasHeaderOptionFilename() bool`

HasHeaderOptionFilename returns a boolean if a field has been set.

### GetHeaderOptionServerAddress

`func (o *IpamsvcSubnet) GetHeaderOptionServerAddress() string`

GetHeaderOptionServerAddress returns the HeaderOptionServerAddress field if non-nil, zero value otherwise.

### GetHeaderOptionServerAddressOk

`func (o *IpamsvcSubnet) GetHeaderOptionServerAddressOk() (*string, bool)`

GetHeaderOptionServerAddressOk returns a tuple with the HeaderOptionServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerAddress

`func (o *IpamsvcSubnet) SetHeaderOptionServerAddress(v string)`

SetHeaderOptionServerAddress sets HeaderOptionServerAddress field to given value.

### HasHeaderOptionServerAddress

`func (o *IpamsvcSubnet) HasHeaderOptionServerAddress() bool`

HasHeaderOptionServerAddress returns a boolean if a field has been set.

### GetHeaderOptionServerName

`func (o *IpamsvcSubnet) GetHeaderOptionServerName() string`

GetHeaderOptionServerName returns the HeaderOptionServerName field if non-nil, zero value otherwise.

### GetHeaderOptionServerNameOk

`func (o *IpamsvcSubnet) GetHeaderOptionServerNameOk() (*string, bool)`

GetHeaderOptionServerNameOk returns a tuple with the HeaderOptionServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerName

`func (o *IpamsvcSubnet) SetHeaderOptionServerName(v string)`

SetHeaderOptionServerName sets HeaderOptionServerName field to given value.

### HasHeaderOptionServerName

`func (o *IpamsvcSubnet) HasHeaderOptionServerName() bool`

HasHeaderOptionServerName returns a boolean if a field has been set.

### GetHostnameRewriteChar

`func (o *IpamsvcSubnet) GetHostnameRewriteChar() string`

GetHostnameRewriteChar returns the HostnameRewriteChar field if non-nil, zero value otherwise.

### GetHostnameRewriteCharOk

`func (o *IpamsvcSubnet) GetHostnameRewriteCharOk() (*string, bool)`

GetHostnameRewriteCharOk returns a tuple with the HostnameRewriteChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameRewriteChar

`func (o *IpamsvcSubnet) SetHostnameRewriteChar(v string)`

SetHostnameRewriteChar sets HostnameRewriteChar field to given value.

### HasHostnameRewriteChar

`func (o *IpamsvcSubnet) HasHostnameRewriteChar() bool`

HasHostnameRewriteChar returns a boolean if a field has been set.

### GetHostnameRewriteEnabled

`func (o *IpamsvcSubnet) GetHostnameRewriteEnabled() bool`

GetHostnameRewriteEnabled returns the HostnameRewriteEnabled field if non-nil, zero value otherwise.

### GetHostnameRewriteEnabledOk

`func (o *IpamsvcSubnet) GetHostnameRewriteEnabledOk() (*bool, bool)`

GetHostnameRewriteEnabledOk returns a tuple with the HostnameRewriteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameRewriteEnabled

`func (o *IpamsvcSubnet) SetHostnameRewriteEnabled(v bool)`

SetHostnameRewriteEnabled sets HostnameRewriteEnabled field to given value.

### HasHostnameRewriteEnabled

`func (o *IpamsvcSubnet) HasHostnameRewriteEnabled() bool`

HasHostnameRewriteEnabled returns a boolean if a field has been set.

### GetHostnameRewriteRegex

`func (o *IpamsvcSubnet) GetHostnameRewriteRegex() string`

GetHostnameRewriteRegex returns the HostnameRewriteRegex field if non-nil, zero value otherwise.

### GetHostnameRewriteRegexOk

`func (o *IpamsvcSubnet) GetHostnameRewriteRegexOk() (*string, bool)`

GetHostnameRewriteRegexOk returns a tuple with the HostnameRewriteRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameRewriteRegex

`func (o *IpamsvcSubnet) SetHostnameRewriteRegex(v string)`

SetHostnameRewriteRegex sets HostnameRewriteRegex field to given value.

### HasHostnameRewriteRegex

`func (o *IpamsvcSubnet) HasHostnameRewriteRegex() bool`

HasHostnameRewriteRegex returns a boolean if a field has been set.

### GetId

`func (o *IpamsvcSubnet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpamsvcSubnet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpamsvcSubnet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpamsvcSubnet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritanceAssignedHosts

`func (o *IpamsvcSubnet) GetInheritanceAssignedHosts() []InheritanceAssignedHost`

GetInheritanceAssignedHosts returns the InheritanceAssignedHosts field if non-nil, zero value otherwise.

### GetInheritanceAssignedHostsOk

`func (o *IpamsvcSubnet) GetInheritanceAssignedHostsOk() (*[]InheritanceAssignedHost, bool)`

GetInheritanceAssignedHostsOk returns a tuple with the InheritanceAssignedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceAssignedHosts

`func (o *IpamsvcSubnet) SetInheritanceAssignedHosts(v []InheritanceAssignedHost)`

SetInheritanceAssignedHosts sets InheritanceAssignedHosts field to given value.

### HasInheritanceAssignedHosts

`func (o *IpamsvcSubnet) HasInheritanceAssignedHosts() bool`

HasInheritanceAssignedHosts returns a boolean if a field has been set.

### GetInheritanceParent

`func (o *IpamsvcSubnet) GetInheritanceParent() string`

GetInheritanceParent returns the InheritanceParent field if non-nil, zero value otherwise.

### GetInheritanceParentOk

`func (o *IpamsvcSubnet) GetInheritanceParentOk() (*string, bool)`

GetInheritanceParentOk returns a tuple with the InheritanceParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceParent

`func (o *IpamsvcSubnet) SetInheritanceParent(v string)`

SetInheritanceParent sets InheritanceParent field to given value.

### HasInheritanceParent

`func (o *IpamsvcSubnet) HasInheritanceParent() bool`

HasInheritanceParent returns a boolean if a field has been set.

### GetInheritanceSources

`func (o *IpamsvcSubnet) GetInheritanceSources() IpamsvcDHCPInheritance`

GetInheritanceSources returns the InheritanceSources field if non-nil, zero value otherwise.

### GetInheritanceSourcesOk

`func (o *IpamsvcSubnet) GetInheritanceSourcesOk() (*IpamsvcDHCPInheritance, bool)`

GetInheritanceSourcesOk returns a tuple with the InheritanceSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceSources

`func (o *IpamsvcSubnet) SetInheritanceSources(v IpamsvcDHCPInheritance)`

SetInheritanceSources sets InheritanceSources field to given value.

### HasInheritanceSources

`func (o *IpamsvcSubnet) HasInheritanceSources() bool`

HasInheritanceSources returns a boolean if a field has been set.

### GetName

`func (o *IpamsvcSubnet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpamsvcSubnet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpamsvcSubnet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IpamsvcSubnet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *IpamsvcSubnet) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IpamsvcSubnet) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IpamsvcSubnet) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IpamsvcSubnet) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProtocol

`func (o *IpamsvcSubnet) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IpamsvcSubnet) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IpamsvcSubnet) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IpamsvcSubnet) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRebindTime

`func (o *IpamsvcSubnet) GetRebindTime() int64`

GetRebindTime returns the RebindTime field if non-nil, zero value otherwise.

### GetRebindTimeOk

`func (o *IpamsvcSubnet) GetRebindTimeOk() (*int64, bool)`

GetRebindTimeOk returns a tuple with the RebindTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebindTime

`func (o *IpamsvcSubnet) SetRebindTime(v int64)`

SetRebindTime sets RebindTime field to given value.

### HasRebindTime

`func (o *IpamsvcSubnet) HasRebindTime() bool`

HasRebindTime returns a boolean if a field has been set.

### GetRenewTime

`func (o *IpamsvcSubnet) GetRenewTime() int64`

GetRenewTime returns the RenewTime field if non-nil, zero value otherwise.

### GetRenewTimeOk

`func (o *IpamsvcSubnet) GetRenewTimeOk() (*int64, bool)`

GetRenewTimeOk returns a tuple with the RenewTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewTime

`func (o *IpamsvcSubnet) SetRenewTime(v int64)`

SetRenewTime sets RenewTime field to given value.

### HasRenewTime

`func (o *IpamsvcSubnet) HasRenewTime() bool`

HasRenewTime returns a boolean if a field has been set.

### GetSpace

`func (o *IpamsvcSubnet) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *IpamsvcSubnet) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *IpamsvcSubnet) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *IpamsvcSubnet) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetTags

`func (o *IpamsvcSubnet) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IpamsvcSubnet) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IpamsvcSubnet) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *IpamsvcSubnet) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThreshold

`func (o *IpamsvcSubnet) GetThreshold() IpamsvcUtilizationThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *IpamsvcSubnet) GetThresholdOk() (*IpamsvcUtilizationThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *IpamsvcSubnet) SetThreshold(v IpamsvcUtilizationThreshold)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *IpamsvcSubnet) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IpamsvcSubnet) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IpamsvcSubnet) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IpamsvcSubnet) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IpamsvcSubnet) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsage

`func (o *IpamsvcSubnet) GetUsage() []string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *IpamsvcSubnet) GetUsageOk() (*[]string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *IpamsvcSubnet) SetUsage(v []string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *IpamsvcSubnet) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetUtilization

`func (o *IpamsvcSubnet) GetUtilization() IpamsvcUtilization`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *IpamsvcSubnet) GetUtilizationOk() (*IpamsvcUtilization, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *IpamsvcSubnet) SetUtilization(v IpamsvcUtilization)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *IpamsvcSubnet) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.

### GetUtilizationV6

`func (o *IpamsvcSubnet) GetUtilizationV6() IpamsvcUtilizationV6`

GetUtilizationV6 returns the UtilizationV6 field if non-nil, zero value otherwise.

### GetUtilizationV6Ok

`func (o *IpamsvcSubnet) GetUtilizationV6Ok() (*IpamsvcUtilizationV6, bool)`

GetUtilizationV6Ok returns a tuple with the UtilizationV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationV6

`func (o *IpamsvcSubnet) SetUtilizationV6(v IpamsvcUtilizationV6)`

SetUtilizationV6 sets UtilizationV6 field to given value.

### HasUtilizationV6

`func (o *IpamsvcSubnet) HasUtilizationV6() bool`

HasUtilizationV6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


