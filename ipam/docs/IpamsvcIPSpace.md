# IpamsvcIPSpace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsmConfig** | Pointer to [**IpamsvcASMConfig**](IpamsvcASMConfig.md) |  | [optional] 
**AsmScopeFlag** | Pointer to **int64** | The number of times the automated scope management usage limits have been exceeded for any of the subnets in this IP space. | [optional] [readonly] 
**Comment** | Pointer to **string** | The description for the IP space. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**DdnsClientUpdate** | Pointer to **string** | Controls who does the DDNS updates.  Valid values are: * _client_: DHCP server updates DNS if requested by client. * _server_: DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates. * _ignore_: DHCP server always updates DNS, even if the client says not to. * _over_client_update_: Same as _server_. DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates. * _over_no_update_: DHCP server updates DNS even if the client requests that no updates be done. If the client requests to do the update, DHCP server allows it.  Defaults to _client_. | [optional] 
**DdnsConflictResolutionMode** | Pointer to **string** | The mode used for resolving conflicts while performing DDNS updates.  Valid values are: * _check_with_dhcid_: It includes adding a DHCID record and checking that record via conflict detection as per RFC 4703. * _no_check_with_dhcid_: This will ignore conflict detection but add a DHCID record when creating/updating an entry. * _check_exists_with_dhcid_: This will check if there is an existing DHCID record but does not verify the value of the record matches the update. This will also update the DHCID record for the entry. * _no_check_without_dhcid_: This ignores conflict detection and will not add a DHCID record when creating/updating a DDNS entry.  Defaults to _check_with_dhcid_. | [optional] 
**DdnsDomain** | Pointer to **string** | The domain suffix for DDNS updates. FQDN, may be empty.  Defaults to empty. | [optional] 
**DdnsGenerateName** | Pointer to **bool** | Indicates if DDNS needs to generate a hostname when not supplied by the client.  Defaults to _false_. | [optional] 
**DdnsGeneratedPrefix** | Pointer to **string** | The prefix used in the generation of an FQDN.  When generating a name, DHCP server will construct the name in the format: [ddns-generated-prefix]-[address-text].[ddns-qualifying-suffix]. where address-text is simply the lease IP address converted to a hyphenated string.  Defaults to \&quot;myhost\&quot;. | [optional] 
**DdnsSendUpdates** | Pointer to **bool** | Determines if DDNS updates are enabled at the IP space level. Defaults to _true_. | [optional] 
**DdnsTtlPercent** | Pointer to **float32** | DDNS TTL value - to be calculated as a simple percentage of the lease&#39;s lifetime, using the parameter&#39;s value as the percentage. It is specified as a percentage (e.g. 25, 75). Defaults to unspecified. | [optional] 
**DdnsUpdateOnRenew** | Pointer to **bool** | Instructs the DHCP server to always update the DNS information when a lease is renewed even if its DNS information has not changed.  Defaults to _false_. | [optional] 
**DdnsUseConflictResolution** | Pointer to **bool** | When true, DHCP server will apply conflict resolution, as described in RFC 4703, when attempting to fulfill the update request.  When false, DHCP server will simply attempt to update the DNS entries per the request, regardless of whether or not they conflict with existing entries owned by other DHCP4 clients.  Defaults to _true_. | [optional] 
**DhcpConfig** | Pointer to [**IpamsvcDHCPConfig**](IpamsvcDHCPConfig.md) |  | [optional] 
**DhcpOptions** | Pointer to [**[]IpamsvcOptionItem**](IpamsvcOptionItem.md) | The list of IPv4 DHCP options for IP space. May be either a specific option or a group of options. | [optional] 
**DhcpOptionsV6** | Pointer to [**[]IpamsvcOptionItem**](IpamsvcOptionItem.md) | The list of IPv6 DHCP options for IP space. May be either a specific option or a group of options. | [optional] 
**HeaderOptionFilename** | Pointer to **string** | The configuration for header option filename field. | [optional] 
**HeaderOptionServerAddress** | Pointer to **string** | The configuration for header option server address field. | [optional] 
**HeaderOptionServerName** | Pointer to **string** | The configuration for header option server name field. | [optional] 
**HostnameRewriteChar** | Pointer to **string** | The character to replace non-matching characters with, when hostname rewrite is enabled.  Any single ASCII character or no character if the invalid characters should be removed without replacement.  Defaults to \&quot;-\&quot;. | [optional] 
**HostnameRewriteEnabled** | Pointer to **bool** | Indicates if client supplied hostnames will be rewritten prior to DDNS update by replacing every character that does not match _hostname_rewrite_regex_ by _hostname_rewrite_char_.  Defaults to _false_. | [optional] 
**HostnameRewriteRegex** | Pointer to **string** | The regex bracket expression to match valid characters.  Must begin with \&quot;[\&quot; and end with \&quot;]\&quot; and be a compilable POSIX regex.  Defaults to \&quot;[^a-zA-Z0-9_.]\&quot;. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InheritanceSources** | Pointer to [**IpamsvcIPSpaceInheritance**](IpamsvcIPSpaceInheritance.md) |  | [optional] 
**Name** | **string** | The name of the IP space. Must contain 1 to 256 characters. Can include UTF-8. | 
**Tags** | Pointer to **map[string]interface{}** | The tags for the IP space in JSON format. | [optional] 
**Threshold** | Pointer to [**IpamsvcUtilizationThreshold**](IpamsvcUtilizationThreshold.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 
**Utilization** | Pointer to [**IpamsvcUtilization**](IpamsvcUtilization.md) |  | [optional] 
**UtilizationV6** | Pointer to [**IpamsvcUtilizationV6**](IpamsvcUtilizationV6.md) |  | [optional] 
**VendorSpecificOptionOptionSpace** | Pointer to **string** | The resource identifier. | [optional] 

## Methods

### NewIpamsvcIPSpace

`func NewIpamsvcIPSpace(name string, ) *IpamsvcIPSpace`

NewIpamsvcIPSpace instantiates a new IpamsvcIPSpace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcIPSpaceWithDefaults

`func NewIpamsvcIPSpaceWithDefaults() *IpamsvcIPSpace`

NewIpamsvcIPSpaceWithDefaults instantiates a new IpamsvcIPSpace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsmConfig

`func (o *IpamsvcIPSpace) GetAsmConfig() IpamsvcASMConfig`

GetAsmConfig returns the AsmConfig field if non-nil, zero value otherwise.

### GetAsmConfigOk

`func (o *IpamsvcIPSpace) GetAsmConfigOk() (*IpamsvcASMConfig, bool)`

GetAsmConfigOk returns a tuple with the AsmConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmConfig

`func (o *IpamsvcIPSpace) SetAsmConfig(v IpamsvcASMConfig)`

SetAsmConfig sets AsmConfig field to given value.

### HasAsmConfig

`func (o *IpamsvcIPSpace) HasAsmConfig() bool`

HasAsmConfig returns a boolean if a field has been set.

### GetAsmScopeFlag

`func (o *IpamsvcIPSpace) GetAsmScopeFlag() int64`

GetAsmScopeFlag returns the AsmScopeFlag field if non-nil, zero value otherwise.

### GetAsmScopeFlagOk

`func (o *IpamsvcIPSpace) GetAsmScopeFlagOk() (*int64, bool)`

GetAsmScopeFlagOk returns a tuple with the AsmScopeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmScopeFlag

`func (o *IpamsvcIPSpace) SetAsmScopeFlag(v int64)`

SetAsmScopeFlag sets AsmScopeFlag field to given value.

### HasAsmScopeFlag

`func (o *IpamsvcIPSpace) HasAsmScopeFlag() bool`

HasAsmScopeFlag returns a boolean if a field has been set.

### GetComment

`func (o *IpamsvcIPSpace) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IpamsvcIPSpace) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IpamsvcIPSpace) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IpamsvcIPSpace) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IpamsvcIPSpace) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IpamsvcIPSpace) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IpamsvcIPSpace) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IpamsvcIPSpace) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDdnsClientUpdate

`func (o *IpamsvcIPSpace) GetDdnsClientUpdate() string`

GetDdnsClientUpdate returns the DdnsClientUpdate field if non-nil, zero value otherwise.

### GetDdnsClientUpdateOk

`func (o *IpamsvcIPSpace) GetDdnsClientUpdateOk() (*string, bool)`

GetDdnsClientUpdateOk returns a tuple with the DdnsClientUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsClientUpdate

`func (o *IpamsvcIPSpace) SetDdnsClientUpdate(v string)`

SetDdnsClientUpdate sets DdnsClientUpdate field to given value.

### HasDdnsClientUpdate

`func (o *IpamsvcIPSpace) HasDdnsClientUpdate() bool`

HasDdnsClientUpdate returns a boolean if a field has been set.

### GetDdnsConflictResolutionMode

`func (o *IpamsvcIPSpace) GetDdnsConflictResolutionMode() string`

GetDdnsConflictResolutionMode returns the DdnsConflictResolutionMode field if non-nil, zero value otherwise.

### GetDdnsConflictResolutionModeOk

`func (o *IpamsvcIPSpace) GetDdnsConflictResolutionModeOk() (*string, bool)`

GetDdnsConflictResolutionModeOk returns a tuple with the DdnsConflictResolutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsConflictResolutionMode

`func (o *IpamsvcIPSpace) SetDdnsConflictResolutionMode(v string)`

SetDdnsConflictResolutionMode sets DdnsConflictResolutionMode field to given value.

### HasDdnsConflictResolutionMode

`func (o *IpamsvcIPSpace) HasDdnsConflictResolutionMode() bool`

HasDdnsConflictResolutionMode returns a boolean if a field has been set.

### GetDdnsDomain

`func (o *IpamsvcIPSpace) GetDdnsDomain() string`

GetDdnsDomain returns the DdnsDomain field if non-nil, zero value otherwise.

### GetDdnsDomainOk

`func (o *IpamsvcIPSpace) GetDdnsDomainOk() (*string, bool)`

GetDdnsDomainOk returns a tuple with the DdnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDomain

`func (o *IpamsvcIPSpace) SetDdnsDomain(v string)`

SetDdnsDomain sets DdnsDomain field to given value.

### HasDdnsDomain

`func (o *IpamsvcIPSpace) HasDdnsDomain() bool`

HasDdnsDomain returns a boolean if a field has been set.

### GetDdnsGenerateName

`func (o *IpamsvcIPSpace) GetDdnsGenerateName() bool`

GetDdnsGenerateName returns the DdnsGenerateName field if non-nil, zero value otherwise.

### GetDdnsGenerateNameOk

`func (o *IpamsvcIPSpace) GetDdnsGenerateNameOk() (*bool, bool)`

GetDdnsGenerateNameOk returns a tuple with the DdnsGenerateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGenerateName

`func (o *IpamsvcIPSpace) SetDdnsGenerateName(v bool)`

SetDdnsGenerateName sets DdnsGenerateName field to given value.

### HasDdnsGenerateName

`func (o *IpamsvcIPSpace) HasDdnsGenerateName() bool`

HasDdnsGenerateName returns a boolean if a field has been set.

### GetDdnsGeneratedPrefix

`func (o *IpamsvcIPSpace) GetDdnsGeneratedPrefix() string`

GetDdnsGeneratedPrefix returns the DdnsGeneratedPrefix field if non-nil, zero value otherwise.

### GetDdnsGeneratedPrefixOk

`func (o *IpamsvcIPSpace) GetDdnsGeneratedPrefixOk() (*string, bool)`

GetDdnsGeneratedPrefixOk returns a tuple with the DdnsGeneratedPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGeneratedPrefix

`func (o *IpamsvcIPSpace) SetDdnsGeneratedPrefix(v string)`

SetDdnsGeneratedPrefix sets DdnsGeneratedPrefix field to given value.

### HasDdnsGeneratedPrefix

`func (o *IpamsvcIPSpace) HasDdnsGeneratedPrefix() bool`

HasDdnsGeneratedPrefix returns a boolean if a field has been set.

### GetDdnsSendUpdates

`func (o *IpamsvcIPSpace) GetDdnsSendUpdates() bool`

GetDdnsSendUpdates returns the DdnsSendUpdates field if non-nil, zero value otherwise.

### GetDdnsSendUpdatesOk

`func (o *IpamsvcIPSpace) GetDdnsSendUpdatesOk() (*bool, bool)`

GetDdnsSendUpdatesOk returns a tuple with the DdnsSendUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsSendUpdates

`func (o *IpamsvcIPSpace) SetDdnsSendUpdates(v bool)`

SetDdnsSendUpdates sets DdnsSendUpdates field to given value.

### HasDdnsSendUpdates

`func (o *IpamsvcIPSpace) HasDdnsSendUpdates() bool`

HasDdnsSendUpdates returns a boolean if a field has been set.

### GetDdnsTtlPercent

`func (o *IpamsvcIPSpace) GetDdnsTtlPercent() float32`

GetDdnsTtlPercent returns the DdnsTtlPercent field if non-nil, zero value otherwise.

### GetDdnsTtlPercentOk

`func (o *IpamsvcIPSpace) GetDdnsTtlPercentOk() (*float32, bool)`

GetDdnsTtlPercentOk returns a tuple with the DdnsTtlPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsTtlPercent

`func (o *IpamsvcIPSpace) SetDdnsTtlPercent(v float32)`

SetDdnsTtlPercent sets DdnsTtlPercent field to given value.

### HasDdnsTtlPercent

`func (o *IpamsvcIPSpace) HasDdnsTtlPercent() bool`

HasDdnsTtlPercent returns a boolean if a field has been set.

### GetDdnsUpdateOnRenew

`func (o *IpamsvcIPSpace) GetDdnsUpdateOnRenew() bool`

GetDdnsUpdateOnRenew returns the DdnsUpdateOnRenew field if non-nil, zero value otherwise.

### GetDdnsUpdateOnRenewOk

`func (o *IpamsvcIPSpace) GetDdnsUpdateOnRenewOk() (*bool, bool)`

GetDdnsUpdateOnRenewOk returns a tuple with the DdnsUpdateOnRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUpdateOnRenew

`func (o *IpamsvcIPSpace) SetDdnsUpdateOnRenew(v bool)`

SetDdnsUpdateOnRenew sets DdnsUpdateOnRenew field to given value.

### HasDdnsUpdateOnRenew

`func (o *IpamsvcIPSpace) HasDdnsUpdateOnRenew() bool`

HasDdnsUpdateOnRenew returns a boolean if a field has been set.

### GetDdnsUseConflictResolution

`func (o *IpamsvcIPSpace) GetDdnsUseConflictResolution() bool`

GetDdnsUseConflictResolution returns the DdnsUseConflictResolution field if non-nil, zero value otherwise.

### GetDdnsUseConflictResolutionOk

`func (o *IpamsvcIPSpace) GetDdnsUseConflictResolutionOk() (*bool, bool)`

GetDdnsUseConflictResolutionOk returns a tuple with the DdnsUseConflictResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUseConflictResolution

`func (o *IpamsvcIPSpace) SetDdnsUseConflictResolution(v bool)`

SetDdnsUseConflictResolution sets DdnsUseConflictResolution field to given value.

### HasDdnsUseConflictResolution

`func (o *IpamsvcIPSpace) HasDdnsUseConflictResolution() bool`

HasDdnsUseConflictResolution returns a boolean if a field has been set.

### GetDhcpConfig

`func (o *IpamsvcIPSpace) GetDhcpConfig() IpamsvcDHCPConfig`

GetDhcpConfig returns the DhcpConfig field if non-nil, zero value otherwise.

### GetDhcpConfigOk

`func (o *IpamsvcIPSpace) GetDhcpConfigOk() (*IpamsvcDHCPConfig, bool)`

GetDhcpConfigOk returns a tuple with the DhcpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpConfig

`func (o *IpamsvcIPSpace) SetDhcpConfig(v IpamsvcDHCPConfig)`

SetDhcpConfig sets DhcpConfig field to given value.

### HasDhcpConfig

`func (o *IpamsvcIPSpace) HasDhcpConfig() bool`

HasDhcpConfig returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *IpamsvcIPSpace) GetDhcpOptions() []IpamsvcOptionItem`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *IpamsvcIPSpace) GetDhcpOptionsOk() (*[]IpamsvcOptionItem, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *IpamsvcIPSpace) SetDhcpOptions(v []IpamsvcOptionItem)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *IpamsvcIPSpace) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetDhcpOptionsV6

`func (o *IpamsvcIPSpace) GetDhcpOptionsV6() []IpamsvcOptionItem`

GetDhcpOptionsV6 returns the DhcpOptionsV6 field if non-nil, zero value otherwise.

### GetDhcpOptionsV6Ok

`func (o *IpamsvcIPSpace) GetDhcpOptionsV6Ok() (*[]IpamsvcOptionItem, bool)`

GetDhcpOptionsV6Ok returns a tuple with the DhcpOptionsV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptionsV6

`func (o *IpamsvcIPSpace) SetDhcpOptionsV6(v []IpamsvcOptionItem)`

SetDhcpOptionsV6 sets DhcpOptionsV6 field to given value.

### HasDhcpOptionsV6

`func (o *IpamsvcIPSpace) HasDhcpOptionsV6() bool`

HasDhcpOptionsV6 returns a boolean if a field has been set.

### GetHeaderOptionFilename

`func (o *IpamsvcIPSpace) GetHeaderOptionFilename() string`

GetHeaderOptionFilename returns the HeaderOptionFilename field if non-nil, zero value otherwise.

### GetHeaderOptionFilenameOk

`func (o *IpamsvcIPSpace) GetHeaderOptionFilenameOk() (*string, bool)`

GetHeaderOptionFilenameOk returns a tuple with the HeaderOptionFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionFilename

`func (o *IpamsvcIPSpace) SetHeaderOptionFilename(v string)`

SetHeaderOptionFilename sets HeaderOptionFilename field to given value.

### HasHeaderOptionFilename

`func (o *IpamsvcIPSpace) HasHeaderOptionFilename() bool`

HasHeaderOptionFilename returns a boolean if a field has been set.

### GetHeaderOptionServerAddress

`func (o *IpamsvcIPSpace) GetHeaderOptionServerAddress() string`

GetHeaderOptionServerAddress returns the HeaderOptionServerAddress field if non-nil, zero value otherwise.

### GetHeaderOptionServerAddressOk

`func (o *IpamsvcIPSpace) GetHeaderOptionServerAddressOk() (*string, bool)`

GetHeaderOptionServerAddressOk returns a tuple with the HeaderOptionServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerAddress

`func (o *IpamsvcIPSpace) SetHeaderOptionServerAddress(v string)`

SetHeaderOptionServerAddress sets HeaderOptionServerAddress field to given value.

### HasHeaderOptionServerAddress

`func (o *IpamsvcIPSpace) HasHeaderOptionServerAddress() bool`

HasHeaderOptionServerAddress returns a boolean if a field has been set.

### GetHeaderOptionServerName

`func (o *IpamsvcIPSpace) GetHeaderOptionServerName() string`

GetHeaderOptionServerName returns the HeaderOptionServerName field if non-nil, zero value otherwise.

### GetHeaderOptionServerNameOk

`func (o *IpamsvcIPSpace) GetHeaderOptionServerNameOk() (*string, bool)`

GetHeaderOptionServerNameOk returns a tuple with the HeaderOptionServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerName

`func (o *IpamsvcIPSpace) SetHeaderOptionServerName(v string)`

SetHeaderOptionServerName sets HeaderOptionServerName field to given value.

### HasHeaderOptionServerName

`func (o *IpamsvcIPSpace) HasHeaderOptionServerName() bool`

HasHeaderOptionServerName returns a boolean if a field has been set.

### GetHostnameRewriteChar

`func (o *IpamsvcIPSpace) GetHostnameRewriteChar() string`

GetHostnameRewriteChar returns the HostnameRewriteChar field if non-nil, zero value otherwise.

### GetHostnameRewriteCharOk

`func (o *IpamsvcIPSpace) GetHostnameRewriteCharOk() (*string, bool)`

GetHostnameRewriteCharOk returns a tuple with the HostnameRewriteChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameRewriteChar

`func (o *IpamsvcIPSpace) SetHostnameRewriteChar(v string)`

SetHostnameRewriteChar sets HostnameRewriteChar field to given value.

### HasHostnameRewriteChar

`func (o *IpamsvcIPSpace) HasHostnameRewriteChar() bool`

HasHostnameRewriteChar returns a boolean if a field has been set.

### GetHostnameRewriteEnabled

`func (o *IpamsvcIPSpace) GetHostnameRewriteEnabled() bool`

GetHostnameRewriteEnabled returns the HostnameRewriteEnabled field if non-nil, zero value otherwise.

### GetHostnameRewriteEnabledOk

`func (o *IpamsvcIPSpace) GetHostnameRewriteEnabledOk() (*bool, bool)`

GetHostnameRewriteEnabledOk returns a tuple with the HostnameRewriteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameRewriteEnabled

`func (o *IpamsvcIPSpace) SetHostnameRewriteEnabled(v bool)`

SetHostnameRewriteEnabled sets HostnameRewriteEnabled field to given value.

### HasHostnameRewriteEnabled

`func (o *IpamsvcIPSpace) HasHostnameRewriteEnabled() bool`

HasHostnameRewriteEnabled returns a boolean if a field has been set.

### GetHostnameRewriteRegex

`func (o *IpamsvcIPSpace) GetHostnameRewriteRegex() string`

GetHostnameRewriteRegex returns the HostnameRewriteRegex field if non-nil, zero value otherwise.

### GetHostnameRewriteRegexOk

`func (o *IpamsvcIPSpace) GetHostnameRewriteRegexOk() (*string, bool)`

GetHostnameRewriteRegexOk returns a tuple with the HostnameRewriteRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameRewriteRegex

`func (o *IpamsvcIPSpace) SetHostnameRewriteRegex(v string)`

SetHostnameRewriteRegex sets HostnameRewriteRegex field to given value.

### HasHostnameRewriteRegex

`func (o *IpamsvcIPSpace) HasHostnameRewriteRegex() bool`

HasHostnameRewriteRegex returns a boolean if a field has been set.

### GetId

`func (o *IpamsvcIPSpace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpamsvcIPSpace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpamsvcIPSpace) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpamsvcIPSpace) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritanceSources

`func (o *IpamsvcIPSpace) GetInheritanceSources() IpamsvcIPSpaceInheritance`

GetInheritanceSources returns the InheritanceSources field if non-nil, zero value otherwise.

### GetInheritanceSourcesOk

`func (o *IpamsvcIPSpace) GetInheritanceSourcesOk() (*IpamsvcIPSpaceInheritance, bool)`

GetInheritanceSourcesOk returns a tuple with the InheritanceSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceSources

`func (o *IpamsvcIPSpace) SetInheritanceSources(v IpamsvcIPSpaceInheritance)`

SetInheritanceSources sets InheritanceSources field to given value.

### HasInheritanceSources

`func (o *IpamsvcIPSpace) HasInheritanceSources() bool`

HasInheritanceSources returns a boolean if a field has been set.

### GetName

`func (o *IpamsvcIPSpace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpamsvcIPSpace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpamsvcIPSpace) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *IpamsvcIPSpace) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IpamsvcIPSpace) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IpamsvcIPSpace) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *IpamsvcIPSpace) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThreshold

`func (o *IpamsvcIPSpace) GetThreshold() IpamsvcUtilizationThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *IpamsvcIPSpace) GetThresholdOk() (*IpamsvcUtilizationThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *IpamsvcIPSpace) SetThreshold(v IpamsvcUtilizationThreshold)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *IpamsvcIPSpace) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IpamsvcIPSpace) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IpamsvcIPSpace) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IpamsvcIPSpace) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IpamsvcIPSpace) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUtilization

`func (o *IpamsvcIPSpace) GetUtilization() IpamsvcUtilization`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *IpamsvcIPSpace) GetUtilizationOk() (*IpamsvcUtilization, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *IpamsvcIPSpace) SetUtilization(v IpamsvcUtilization)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *IpamsvcIPSpace) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.

### GetUtilizationV6

`func (o *IpamsvcIPSpace) GetUtilizationV6() IpamsvcUtilizationV6`

GetUtilizationV6 returns the UtilizationV6 field if non-nil, zero value otherwise.

### GetUtilizationV6Ok

`func (o *IpamsvcIPSpace) GetUtilizationV6Ok() (*IpamsvcUtilizationV6, bool)`

GetUtilizationV6Ok returns a tuple with the UtilizationV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationV6

`func (o *IpamsvcIPSpace) SetUtilizationV6(v IpamsvcUtilizationV6)`

SetUtilizationV6 sets UtilizationV6 field to given value.

### HasUtilizationV6

`func (o *IpamsvcIPSpace) HasUtilizationV6() bool`

HasUtilizationV6 returns a boolean if a field has been set.

### GetVendorSpecificOptionOptionSpace

`func (o *IpamsvcIPSpace) GetVendorSpecificOptionOptionSpace() string`

GetVendorSpecificOptionOptionSpace returns the VendorSpecificOptionOptionSpace field if non-nil, zero value otherwise.

### GetVendorSpecificOptionOptionSpaceOk

`func (o *IpamsvcIPSpace) GetVendorSpecificOptionOptionSpaceOk() (*string, bool)`

GetVendorSpecificOptionOptionSpaceOk returns a tuple with the VendorSpecificOptionOptionSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSpecificOptionOptionSpace

`func (o *IpamsvcIPSpace) SetVendorSpecificOptionOptionSpace(v string)`

SetVendorSpecificOptionOptionSpace sets VendorSpecificOptionOptionSpace field to given value.

### HasVendorSpecificOptionOptionSpace

`func (o *IpamsvcIPSpace) HasVendorSpecificOptionOptionSpace() bool`

HasVendorSpecificOptionOptionSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


