# IpamsvcAddressBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The address field in form “a.b.c.d/n” where the “/n” may be omitted. In this case, the CIDR value must be defined in the _cidr_ field. When reading, the _address_ field is always in the form “a.b.c.d”. | 
**AsmConfig** | Pointer to [**IpamsvcASMConfig**](IpamsvcASMConfig.md) |  | [optional] 
**AsmScopeFlag** | Pointer to **int64** | Incremented by 1 if the IP address usage limits for automated scope management are exceeded for any subnets in the address block. | [optional] [readonly] 
**Cidr** | Pointer to **int64** | The CIDR of the address block. This is required, if _address_ does not specify it in its input. | [optional] 
**Comment** | Pointer to **string** | The description for the address block. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**DdnsClientUpdate** | Pointer to **string** | Controls who does the DDNS updates.  Valid values are: * _client_: DHCP server updates DNS if requested by client. * _server_: DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates. * _ignore_: DHCP server always updates DNS, even if the client says not to. * _over_client_update_: Same as _server_. DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates. * _over_no_update_: DHCP server updates DNS even if the client requests that no updates be done. If the client requests to do the update, DHCP server allows it.  Defaults to _client_. | [optional] 
**DdnsConflictResolutionMode** | Pointer to **string** | The mode used for resolving conflicts while performing DDNS updates.  Valid values are: * _check_with_dhcid_: It includes adding a DHCID record and checking that record via conflict detection as per RFC 4703. * _no_check_with_dhcid_: This will ignore conflict detection but add a DHCID record when creating/updating an entry. * _check_exists_with_dhcid_: This will check if there is an existing DHCID record but does not verify the value of the record matches the update. This will also update the DHCID record for the entry. * _no_check_without_dhcid_: This ignores conflict detection and will not add a DHCID record when creating/updating a DDNS entry.  Defaults to _check_with_dhcid_. | [optional] 
**DdnsDomain** | Pointer to **string** | The domain suffix for DDNS updates. FQDN, may be empty.  Defaults to empty. | [optional] 
**DdnsGenerateName** | Pointer to **bool** | Indicates if DDNS needs to generate a hostname when not supplied by the client.  Defaults to _false_. | [optional] 
**DdnsGeneratedPrefix** | Pointer to **string** | The prefix used in the generation of an FQDN.  When generating a name, DHCP server will construct the name in the format: [ddns-generated-prefix]-[address-text].[ddns-qualifying-suffix]. where address-text is simply the lease IP address converted to a hyphenated string.  Defaults to \&quot;myhost\&quot;. | [optional] 
**DdnsSendUpdates** | Pointer to **bool** | Determines if DDNS updates are enabled at the address block level. Defaults to _true_. | [optional] 
**DdnsTtlPercent** | Pointer to **float32** | DDNS TTL value - to be calculated as a simple percentage of the lease&#39;s lifetime, using the parameter&#39;s value as the percentage. It is specified as a percentage (e.g. 25, 75). Defaults to unspecified. | [optional] 
**DdnsUpdateOnRenew** | Pointer to **bool** | Instructs the DHCP server to always update the DNS information when a lease is renewed even if its DNS information has not changed.  Defaults to _false_. | [optional] 
**DdnsUseConflictResolution** | Pointer to **bool** | When true, DHCP server will apply conflict resolution, as described in RFC 4703, when attempting to fulfill the update request.  When false, DHCP server will simply attempt to update the DNS entries per the request, regardless of whether or not they conflict with existing entries owned by other DHCP4 clients.  Defaults to _true_. | [optional] 
**DhcpConfig** | Pointer to [**IpamsvcDHCPConfig**](IpamsvcDHCPConfig.md) |  | [optional] 
**DhcpOptions** | Pointer to [**[]IpamsvcOptionItem**](IpamsvcOptionItem.md) | The list of DHCP options for the address block. May be either a specific option or a group of options. | [optional] 
**DhcpUtilization** | Pointer to [**IpamsvcDHCPUtilization**](IpamsvcDHCPUtilization.md) |  | [optional] 
**DiscoveryAttrs** | Pointer to **map[string]interface{}** | The discovery attributes for this address block in JSON format. | [optional] 
**DiscoveryMetadata** | Pointer to **map[string]interface{}** | The discovery metadata for this address block in JSON format. | [optional] 
**HeaderOptionFilename** | Pointer to **string** | The configuration for header option filename field. | [optional] 
**HeaderOptionServerAddress** | Pointer to **string** | The configuration for header option server address field. | [optional] 
**HeaderOptionServerName** | Pointer to **string** | The configuration for header option server name field. | [optional] 
**HostnameRewriteChar** | Pointer to **string** | The character to replace non-matching characters with, when hostname rewrite is enabled.  Any single ASCII character or no character if the invalid characters should be removed without replacement.  Defaults to \&quot;-\&quot;. | [optional] 
**HostnameRewriteEnabled** | Pointer to **bool** | Indicates if client supplied hostnames will be rewritten prior to DDNS update by replacing every character that does not match _hostname_rewrite_regex_ by _hostname_rewrite_char_.  Defaults to _false_. | [optional] 
**HostnameRewriteRegex** | Pointer to **string** | The regex bracket expression to match valid characters.  Must begin with \&quot;[\&quot; and end with \&quot;]\&quot; and be a compilable POSIX regex.  Defaults to \&quot;[^a-zA-Z0-9_.]\&quot;. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InheritanceParent** | Pointer to **string** | The resource identifier. | [optional] 
**InheritanceSources** | Pointer to [**IpamsvcDHCPInheritance**](IpamsvcDHCPInheritance.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the address block. May contain 1 to 256 characters. Can include UTF-8. | [optional] 
**Parent** | Pointer to **string** | The resource identifier. | [optional] 
**Protocol** | Pointer to **string** | The type of protocol of address block (_ip4_ or _ip6_). | [optional] [readonly] 
**Space** | **string** | The resource identifier. | 
**Tags** | Pointer to **map[string]interface{}** | The tags for the address block in JSON format. | [optional] 
**Threshold** | Pointer to [**IpamsvcUtilizationThreshold**](IpamsvcUtilizationThreshold.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 
**Usage** | Pointer to **[]string** | The usage is a combination of indicators, each tracking a specific associated use. Listed below are usage indicators with their meaning:  usage indicator        | description  ---------------------- | --------------------------------  _IPAM_                 |  AddressBlock is managed in BloxOne DDI.  _DISCOVERED_           |  AddressBlock is discovered by some network discovery probe like Network Insight or NetMRI in NIOS. | [optional] [readonly] 
**Utilization** | Pointer to [**IpamsvcUtilization**](IpamsvcUtilization.md) |  | [optional] 
**UtilizationV6** | Pointer to [**IpamsvcUtilizationV6**](IpamsvcUtilizationV6.md) |  | [optional] 

## Methods

### NewIpamsvcAddressBlock

`func NewIpamsvcAddressBlock(address string, space string, ) *IpamsvcAddressBlock`

NewIpamsvcAddressBlock instantiates a new IpamsvcAddressBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcAddressBlockWithDefaults

`func NewIpamsvcAddressBlockWithDefaults() *IpamsvcAddressBlock`

NewIpamsvcAddressBlockWithDefaults instantiates a new IpamsvcAddressBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *IpamsvcAddressBlock) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IpamsvcAddressBlock) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IpamsvcAddressBlock) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAsmConfig

`func (o *IpamsvcAddressBlock) GetAsmConfig() IpamsvcASMConfig`

GetAsmConfig returns the AsmConfig field if non-nil, zero value otherwise.

### GetAsmConfigOk

`func (o *IpamsvcAddressBlock) GetAsmConfigOk() (*IpamsvcASMConfig, bool)`

GetAsmConfigOk returns a tuple with the AsmConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmConfig

`func (o *IpamsvcAddressBlock) SetAsmConfig(v IpamsvcASMConfig)`

SetAsmConfig sets AsmConfig field to given value.

### HasAsmConfig

`func (o *IpamsvcAddressBlock) HasAsmConfig() bool`

HasAsmConfig returns a boolean if a field has been set.

### GetAsmScopeFlag

`func (o *IpamsvcAddressBlock) GetAsmScopeFlag() int64`

GetAsmScopeFlag returns the AsmScopeFlag field if non-nil, zero value otherwise.

### GetAsmScopeFlagOk

`func (o *IpamsvcAddressBlock) GetAsmScopeFlagOk() (*int64, bool)`

GetAsmScopeFlagOk returns a tuple with the AsmScopeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmScopeFlag

`func (o *IpamsvcAddressBlock) SetAsmScopeFlag(v int64)`

SetAsmScopeFlag sets AsmScopeFlag field to given value.

### HasAsmScopeFlag

`func (o *IpamsvcAddressBlock) HasAsmScopeFlag() bool`

HasAsmScopeFlag returns a boolean if a field has been set.

### GetCidr

`func (o *IpamsvcAddressBlock) GetCidr() int64`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *IpamsvcAddressBlock) GetCidrOk() (*int64, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *IpamsvcAddressBlock) SetCidr(v int64)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *IpamsvcAddressBlock) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetComment

`func (o *IpamsvcAddressBlock) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IpamsvcAddressBlock) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IpamsvcAddressBlock) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IpamsvcAddressBlock) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IpamsvcAddressBlock) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IpamsvcAddressBlock) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IpamsvcAddressBlock) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IpamsvcAddressBlock) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDdnsClientUpdate

`func (o *IpamsvcAddressBlock) GetDdnsClientUpdate() string`

GetDdnsClientUpdate returns the DdnsClientUpdate field if non-nil, zero value otherwise.

### GetDdnsClientUpdateOk

`func (o *IpamsvcAddressBlock) GetDdnsClientUpdateOk() (*string, bool)`

GetDdnsClientUpdateOk returns a tuple with the DdnsClientUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsClientUpdate

`func (o *IpamsvcAddressBlock) SetDdnsClientUpdate(v string)`

SetDdnsClientUpdate sets DdnsClientUpdate field to given value.

### HasDdnsClientUpdate

`func (o *IpamsvcAddressBlock) HasDdnsClientUpdate() bool`

HasDdnsClientUpdate returns a boolean if a field has been set.

### GetDdnsConflictResolutionMode

`func (o *IpamsvcAddressBlock) GetDdnsConflictResolutionMode() string`

GetDdnsConflictResolutionMode returns the DdnsConflictResolutionMode field if non-nil, zero value otherwise.

### GetDdnsConflictResolutionModeOk

`func (o *IpamsvcAddressBlock) GetDdnsConflictResolutionModeOk() (*string, bool)`

GetDdnsConflictResolutionModeOk returns a tuple with the DdnsConflictResolutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsConflictResolutionMode

`func (o *IpamsvcAddressBlock) SetDdnsConflictResolutionMode(v string)`

SetDdnsConflictResolutionMode sets DdnsConflictResolutionMode field to given value.

### HasDdnsConflictResolutionMode

`func (o *IpamsvcAddressBlock) HasDdnsConflictResolutionMode() bool`

HasDdnsConflictResolutionMode returns a boolean if a field has been set.

### GetDdnsDomain

`func (o *IpamsvcAddressBlock) GetDdnsDomain() string`

GetDdnsDomain returns the DdnsDomain field if non-nil, zero value otherwise.

### GetDdnsDomainOk

`func (o *IpamsvcAddressBlock) GetDdnsDomainOk() (*string, bool)`

GetDdnsDomainOk returns a tuple with the DdnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDomain

`func (o *IpamsvcAddressBlock) SetDdnsDomain(v string)`

SetDdnsDomain sets DdnsDomain field to given value.

### HasDdnsDomain

`func (o *IpamsvcAddressBlock) HasDdnsDomain() bool`

HasDdnsDomain returns a boolean if a field has been set.

### GetDdnsGenerateName

`func (o *IpamsvcAddressBlock) GetDdnsGenerateName() bool`

GetDdnsGenerateName returns the DdnsGenerateName field if non-nil, zero value otherwise.

### GetDdnsGenerateNameOk

`func (o *IpamsvcAddressBlock) GetDdnsGenerateNameOk() (*bool, bool)`

GetDdnsGenerateNameOk returns a tuple with the DdnsGenerateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGenerateName

`func (o *IpamsvcAddressBlock) SetDdnsGenerateName(v bool)`

SetDdnsGenerateName sets DdnsGenerateName field to given value.

### HasDdnsGenerateName

`func (o *IpamsvcAddressBlock) HasDdnsGenerateName() bool`

HasDdnsGenerateName returns a boolean if a field has been set.

### GetDdnsGeneratedPrefix

`func (o *IpamsvcAddressBlock) GetDdnsGeneratedPrefix() string`

GetDdnsGeneratedPrefix returns the DdnsGeneratedPrefix field if non-nil, zero value otherwise.

### GetDdnsGeneratedPrefixOk

`func (o *IpamsvcAddressBlock) GetDdnsGeneratedPrefixOk() (*string, bool)`

GetDdnsGeneratedPrefixOk returns a tuple with the DdnsGeneratedPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGeneratedPrefix

`func (o *IpamsvcAddressBlock) SetDdnsGeneratedPrefix(v string)`

SetDdnsGeneratedPrefix sets DdnsGeneratedPrefix field to given value.

### HasDdnsGeneratedPrefix

`func (o *IpamsvcAddressBlock) HasDdnsGeneratedPrefix() bool`

HasDdnsGeneratedPrefix returns a boolean if a field has been set.

### GetDdnsSendUpdates

`func (o *IpamsvcAddressBlock) GetDdnsSendUpdates() bool`

GetDdnsSendUpdates returns the DdnsSendUpdates field if non-nil, zero value otherwise.

### GetDdnsSendUpdatesOk

`func (o *IpamsvcAddressBlock) GetDdnsSendUpdatesOk() (*bool, bool)`

GetDdnsSendUpdatesOk returns a tuple with the DdnsSendUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsSendUpdates

`func (o *IpamsvcAddressBlock) SetDdnsSendUpdates(v bool)`

SetDdnsSendUpdates sets DdnsSendUpdates field to given value.

### HasDdnsSendUpdates

`func (o *IpamsvcAddressBlock) HasDdnsSendUpdates() bool`

HasDdnsSendUpdates returns a boolean if a field has been set.

### GetDdnsTtlPercent

`func (o *IpamsvcAddressBlock) GetDdnsTtlPercent() float32`

GetDdnsTtlPercent returns the DdnsTtlPercent field if non-nil, zero value otherwise.

### GetDdnsTtlPercentOk

`func (o *IpamsvcAddressBlock) GetDdnsTtlPercentOk() (*float32, bool)`

GetDdnsTtlPercentOk returns a tuple with the DdnsTtlPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsTtlPercent

`func (o *IpamsvcAddressBlock) SetDdnsTtlPercent(v float32)`

SetDdnsTtlPercent sets DdnsTtlPercent field to given value.

### HasDdnsTtlPercent

`func (o *IpamsvcAddressBlock) HasDdnsTtlPercent() bool`

HasDdnsTtlPercent returns a boolean if a field has been set.

### GetDdnsUpdateOnRenew

`func (o *IpamsvcAddressBlock) GetDdnsUpdateOnRenew() bool`

GetDdnsUpdateOnRenew returns the DdnsUpdateOnRenew field if non-nil, zero value otherwise.

### GetDdnsUpdateOnRenewOk

`func (o *IpamsvcAddressBlock) GetDdnsUpdateOnRenewOk() (*bool, bool)`

GetDdnsUpdateOnRenewOk returns a tuple with the DdnsUpdateOnRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUpdateOnRenew

`func (o *IpamsvcAddressBlock) SetDdnsUpdateOnRenew(v bool)`

SetDdnsUpdateOnRenew sets DdnsUpdateOnRenew field to given value.

### HasDdnsUpdateOnRenew

`func (o *IpamsvcAddressBlock) HasDdnsUpdateOnRenew() bool`

HasDdnsUpdateOnRenew returns a boolean if a field has been set.

### GetDdnsUseConflictResolution

`func (o *IpamsvcAddressBlock) GetDdnsUseConflictResolution() bool`

GetDdnsUseConflictResolution returns the DdnsUseConflictResolution field if non-nil, zero value otherwise.

### GetDdnsUseConflictResolutionOk

`func (o *IpamsvcAddressBlock) GetDdnsUseConflictResolutionOk() (*bool, bool)`

GetDdnsUseConflictResolutionOk returns a tuple with the DdnsUseConflictResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUseConflictResolution

`func (o *IpamsvcAddressBlock) SetDdnsUseConflictResolution(v bool)`

SetDdnsUseConflictResolution sets DdnsUseConflictResolution field to given value.

### HasDdnsUseConflictResolution

`func (o *IpamsvcAddressBlock) HasDdnsUseConflictResolution() bool`

HasDdnsUseConflictResolution returns a boolean if a field has been set.

### GetDhcpConfig

`func (o *IpamsvcAddressBlock) GetDhcpConfig() IpamsvcDHCPConfig`

GetDhcpConfig returns the DhcpConfig field if non-nil, zero value otherwise.

### GetDhcpConfigOk

`func (o *IpamsvcAddressBlock) GetDhcpConfigOk() (*IpamsvcDHCPConfig, bool)`

GetDhcpConfigOk returns a tuple with the DhcpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpConfig

`func (o *IpamsvcAddressBlock) SetDhcpConfig(v IpamsvcDHCPConfig)`

SetDhcpConfig sets DhcpConfig field to given value.

### HasDhcpConfig

`func (o *IpamsvcAddressBlock) HasDhcpConfig() bool`

HasDhcpConfig returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *IpamsvcAddressBlock) GetDhcpOptions() []IpamsvcOptionItem`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *IpamsvcAddressBlock) GetDhcpOptionsOk() (*[]IpamsvcOptionItem, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *IpamsvcAddressBlock) SetDhcpOptions(v []IpamsvcOptionItem)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *IpamsvcAddressBlock) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetDhcpUtilization

`func (o *IpamsvcAddressBlock) GetDhcpUtilization() IpamsvcDHCPUtilization`

GetDhcpUtilization returns the DhcpUtilization field if non-nil, zero value otherwise.

### GetDhcpUtilizationOk

`func (o *IpamsvcAddressBlock) GetDhcpUtilizationOk() (*IpamsvcDHCPUtilization, bool)`

GetDhcpUtilizationOk returns a tuple with the DhcpUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilization

`func (o *IpamsvcAddressBlock) SetDhcpUtilization(v IpamsvcDHCPUtilization)`

SetDhcpUtilization sets DhcpUtilization field to given value.

### HasDhcpUtilization

`func (o *IpamsvcAddressBlock) HasDhcpUtilization() bool`

HasDhcpUtilization returns a boolean if a field has been set.

### GetDiscoveryAttrs

`func (o *IpamsvcAddressBlock) GetDiscoveryAttrs() map[string]interface{}`

GetDiscoveryAttrs returns the DiscoveryAttrs field if non-nil, zero value otherwise.

### GetDiscoveryAttrsOk

`func (o *IpamsvcAddressBlock) GetDiscoveryAttrsOk() (*map[string]interface{}, bool)`

GetDiscoveryAttrsOk returns a tuple with the DiscoveryAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryAttrs

`func (o *IpamsvcAddressBlock) SetDiscoveryAttrs(v map[string]interface{})`

SetDiscoveryAttrs sets DiscoveryAttrs field to given value.

### HasDiscoveryAttrs

`func (o *IpamsvcAddressBlock) HasDiscoveryAttrs() bool`

HasDiscoveryAttrs returns a boolean if a field has been set.

### GetDiscoveryMetadata

`func (o *IpamsvcAddressBlock) GetDiscoveryMetadata() map[string]interface{}`

GetDiscoveryMetadata returns the DiscoveryMetadata field if non-nil, zero value otherwise.

### GetDiscoveryMetadataOk

`func (o *IpamsvcAddressBlock) GetDiscoveryMetadataOk() (*map[string]interface{}, bool)`

GetDiscoveryMetadataOk returns a tuple with the DiscoveryMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryMetadata

`func (o *IpamsvcAddressBlock) SetDiscoveryMetadata(v map[string]interface{})`

SetDiscoveryMetadata sets DiscoveryMetadata field to given value.

### HasDiscoveryMetadata

`func (o *IpamsvcAddressBlock) HasDiscoveryMetadata() bool`

HasDiscoveryMetadata returns a boolean if a field has been set.

### GetHeaderOptionFilename

`func (o *IpamsvcAddressBlock) GetHeaderOptionFilename() string`

GetHeaderOptionFilename returns the HeaderOptionFilename field if non-nil, zero value otherwise.

### GetHeaderOptionFilenameOk

`func (o *IpamsvcAddressBlock) GetHeaderOptionFilenameOk() (*string, bool)`

GetHeaderOptionFilenameOk returns a tuple with the HeaderOptionFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionFilename

`func (o *IpamsvcAddressBlock) SetHeaderOptionFilename(v string)`

SetHeaderOptionFilename sets HeaderOptionFilename field to given value.

### HasHeaderOptionFilename

`func (o *IpamsvcAddressBlock) HasHeaderOptionFilename() bool`

HasHeaderOptionFilename returns a boolean if a field has been set.

### GetHeaderOptionServerAddress

`func (o *IpamsvcAddressBlock) GetHeaderOptionServerAddress() string`

GetHeaderOptionServerAddress returns the HeaderOptionServerAddress field if non-nil, zero value otherwise.

### GetHeaderOptionServerAddressOk

`func (o *IpamsvcAddressBlock) GetHeaderOptionServerAddressOk() (*string, bool)`

GetHeaderOptionServerAddressOk returns a tuple with the HeaderOptionServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerAddress

`func (o *IpamsvcAddressBlock) SetHeaderOptionServerAddress(v string)`

SetHeaderOptionServerAddress sets HeaderOptionServerAddress field to given value.

### HasHeaderOptionServerAddress

`func (o *IpamsvcAddressBlock) HasHeaderOptionServerAddress() bool`

HasHeaderOptionServerAddress returns a boolean if a field has been set.

### GetHeaderOptionServerName

`func (o *IpamsvcAddressBlock) GetHeaderOptionServerName() string`

GetHeaderOptionServerName returns the HeaderOptionServerName field if non-nil, zero value otherwise.

### GetHeaderOptionServerNameOk

`func (o *IpamsvcAddressBlock) GetHeaderOptionServerNameOk() (*string, bool)`

GetHeaderOptionServerNameOk returns a tuple with the HeaderOptionServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerName

`func (o *IpamsvcAddressBlock) SetHeaderOptionServerName(v string)`

SetHeaderOptionServerName sets HeaderOptionServerName field to given value.

### HasHeaderOptionServerName

`func (o *IpamsvcAddressBlock) HasHeaderOptionServerName() bool`

HasHeaderOptionServerName returns a boolean if a field has been set.

### GetHostnameRewriteChar

`func (o *IpamsvcAddressBlock) GetHostnameRewriteChar() string`

GetHostnameRewriteChar returns the HostnameRewriteChar field if non-nil, zero value otherwise.

### GetHostnameRewriteCharOk

`func (o *IpamsvcAddressBlock) GetHostnameRewriteCharOk() (*string, bool)`

GetHostnameRewriteCharOk returns a tuple with the HostnameRewriteChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameRewriteChar

`func (o *IpamsvcAddressBlock) SetHostnameRewriteChar(v string)`

SetHostnameRewriteChar sets HostnameRewriteChar field to given value.

### HasHostnameRewriteChar

`func (o *IpamsvcAddressBlock) HasHostnameRewriteChar() bool`

HasHostnameRewriteChar returns a boolean if a field has been set.

### GetHostnameRewriteEnabled

`func (o *IpamsvcAddressBlock) GetHostnameRewriteEnabled() bool`

GetHostnameRewriteEnabled returns the HostnameRewriteEnabled field if non-nil, zero value otherwise.

### GetHostnameRewriteEnabledOk

`func (o *IpamsvcAddressBlock) GetHostnameRewriteEnabledOk() (*bool, bool)`

GetHostnameRewriteEnabledOk returns a tuple with the HostnameRewriteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameRewriteEnabled

`func (o *IpamsvcAddressBlock) SetHostnameRewriteEnabled(v bool)`

SetHostnameRewriteEnabled sets HostnameRewriteEnabled field to given value.

### HasHostnameRewriteEnabled

`func (o *IpamsvcAddressBlock) HasHostnameRewriteEnabled() bool`

HasHostnameRewriteEnabled returns a boolean if a field has been set.

### GetHostnameRewriteRegex

`func (o *IpamsvcAddressBlock) GetHostnameRewriteRegex() string`

GetHostnameRewriteRegex returns the HostnameRewriteRegex field if non-nil, zero value otherwise.

### GetHostnameRewriteRegexOk

`func (o *IpamsvcAddressBlock) GetHostnameRewriteRegexOk() (*string, bool)`

GetHostnameRewriteRegexOk returns a tuple with the HostnameRewriteRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameRewriteRegex

`func (o *IpamsvcAddressBlock) SetHostnameRewriteRegex(v string)`

SetHostnameRewriteRegex sets HostnameRewriteRegex field to given value.

### HasHostnameRewriteRegex

`func (o *IpamsvcAddressBlock) HasHostnameRewriteRegex() bool`

HasHostnameRewriteRegex returns a boolean if a field has been set.

### GetId

`func (o *IpamsvcAddressBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpamsvcAddressBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpamsvcAddressBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpamsvcAddressBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritanceParent

`func (o *IpamsvcAddressBlock) GetInheritanceParent() string`

GetInheritanceParent returns the InheritanceParent field if non-nil, zero value otherwise.

### GetInheritanceParentOk

`func (o *IpamsvcAddressBlock) GetInheritanceParentOk() (*string, bool)`

GetInheritanceParentOk returns a tuple with the InheritanceParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceParent

`func (o *IpamsvcAddressBlock) SetInheritanceParent(v string)`

SetInheritanceParent sets InheritanceParent field to given value.

### HasInheritanceParent

`func (o *IpamsvcAddressBlock) HasInheritanceParent() bool`

HasInheritanceParent returns a boolean if a field has been set.

### GetInheritanceSources

`func (o *IpamsvcAddressBlock) GetInheritanceSources() IpamsvcDHCPInheritance`

GetInheritanceSources returns the InheritanceSources field if non-nil, zero value otherwise.

### GetInheritanceSourcesOk

`func (o *IpamsvcAddressBlock) GetInheritanceSourcesOk() (*IpamsvcDHCPInheritance, bool)`

GetInheritanceSourcesOk returns a tuple with the InheritanceSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceSources

`func (o *IpamsvcAddressBlock) SetInheritanceSources(v IpamsvcDHCPInheritance)`

SetInheritanceSources sets InheritanceSources field to given value.

### HasInheritanceSources

`func (o *IpamsvcAddressBlock) HasInheritanceSources() bool`

HasInheritanceSources returns a boolean if a field has been set.

### GetName

`func (o *IpamsvcAddressBlock) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpamsvcAddressBlock) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpamsvcAddressBlock) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IpamsvcAddressBlock) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *IpamsvcAddressBlock) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IpamsvcAddressBlock) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IpamsvcAddressBlock) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IpamsvcAddressBlock) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProtocol

`func (o *IpamsvcAddressBlock) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IpamsvcAddressBlock) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IpamsvcAddressBlock) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IpamsvcAddressBlock) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSpace

`func (o *IpamsvcAddressBlock) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *IpamsvcAddressBlock) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *IpamsvcAddressBlock) SetSpace(v string)`

SetSpace sets Space field to given value.


### GetTags

`func (o *IpamsvcAddressBlock) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IpamsvcAddressBlock) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IpamsvcAddressBlock) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *IpamsvcAddressBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThreshold

`func (o *IpamsvcAddressBlock) GetThreshold() IpamsvcUtilizationThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *IpamsvcAddressBlock) GetThresholdOk() (*IpamsvcUtilizationThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *IpamsvcAddressBlock) SetThreshold(v IpamsvcUtilizationThreshold)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *IpamsvcAddressBlock) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IpamsvcAddressBlock) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IpamsvcAddressBlock) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IpamsvcAddressBlock) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IpamsvcAddressBlock) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsage

`func (o *IpamsvcAddressBlock) GetUsage() []string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *IpamsvcAddressBlock) GetUsageOk() (*[]string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *IpamsvcAddressBlock) SetUsage(v []string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *IpamsvcAddressBlock) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetUtilization

`func (o *IpamsvcAddressBlock) GetUtilization() IpamsvcUtilization`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *IpamsvcAddressBlock) GetUtilizationOk() (*IpamsvcUtilization, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *IpamsvcAddressBlock) SetUtilization(v IpamsvcUtilization)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *IpamsvcAddressBlock) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.

### GetUtilizationV6

`func (o *IpamsvcAddressBlock) GetUtilizationV6() IpamsvcUtilizationV6`

GetUtilizationV6 returns the UtilizationV6 field if non-nil, zero value otherwise.

### GetUtilizationV6Ok

`func (o *IpamsvcAddressBlock) GetUtilizationV6Ok() (*IpamsvcUtilizationV6, bool)`

GetUtilizationV6Ok returns a tuple with the UtilizationV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationV6

`func (o *IpamsvcAddressBlock) SetUtilizationV6(v IpamsvcUtilizationV6)`

SetUtilizationV6 sets UtilizationV6 field to given value.

### HasUtilizationV6

`func (o *IpamsvcAddressBlock) HasUtilizationV6() bool`

HasUtilizationV6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


