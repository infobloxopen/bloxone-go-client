# IPSpace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsmConfig** | Pointer to [**ASMConfig**](ASMConfig.md) |  | [optional] 
**AsmScopeFlag** | Pointer to **int64** | The number of times the automated scope management usage limits have been exceeded for any of the subnets in this IP space. | [optional] [readonly] 
**Comment** | Pointer to **string** | The description for the IP space. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**DdnsClientUpdate** | Pointer to **string** | Controls who does the DDNS updates.  Valid values are: * _client_: DHCP server updates DNS if requested by client. * _server_: DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates. * _ignore_: DHCP server always updates DNS, even if the client says not to. * _over_client_update_: Same as _server_. DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates. * _over_no_update_: DHCP server updates DNS even if the client requests that no updates be done. If the client requests to do the update, DHCP server allows it.  Defaults to _client_. | [optional] [default to "client"]
**DdnsConflictResolutionMode** | Pointer to **string** | The mode used for resolving conflicts while performing DDNS updates.  Valid values are: * _check_with_dhcid_: It includes adding a DHCID record and checking that record via conflict detection as per RFC 4703. * _no_check_with_dhcid_: This will ignore conflict detection but add a DHCID record when creating/updating an entry. * _check_exists_with_dhcid_: This will check if there is an existing DHCID record but does not verify the value of the record matches the update. This will also update the DHCID record for the entry. * _no_check_without_dhcid_: This ignores conflict detection and will not add a DHCID record when creating/updating a DDNS entry.  Defaults to _check_with_dhcid_. | [optional] [default to "check_with_dhcid"]
**DdnsDomain** | Pointer to **string** | The domain suffix for DDNS updates. FQDN, may be empty.  Defaults to empty. | [optional] 
**DdnsGenerateName** | Pointer to **bool** | Indicates if DDNS needs to generate a hostname when not supplied by the client.  Defaults to _false_. | [optional] [default to false]
**DdnsGeneratedPrefix** | Pointer to **string** | The prefix used in the generation of an FQDN.  When generating a name, DHCP server will construct the name in the format: [ddns-generated-prefix]-[address-text].[ddns-qualifying-suffix]. where address-text is simply the lease IP address converted to a hyphenated string.  Defaults to \&quot;myhost\&quot;. | [optional] [default to "myhost"]
**DdnsSendUpdates** | Pointer to **bool** | Determines if DDNS updates are enabled at the IP space level. Defaults to _true_. | [optional] [default to true]
**DdnsTtlPercent** | Pointer to **float32** | DDNS TTL value - to be calculated as a simple percentage of the lease&#39;s lifetime, using the parameter&#39;s value as the percentage. It is specified as a percentage (e.g. 25, 75). Defaults to unspecified. | [optional] 
**DdnsUpdateOnRenew** | Pointer to **bool** | Instructs the DHCP server to always update the DNS information when a lease is renewed even if its DNS information has not changed.  Defaults to _false_. | [optional] [default to false]
**DdnsUseConflictResolution** | Pointer to **bool** | When true, DHCP server will apply conflict resolution, as described in RFC 4703, when attempting to fulfill the update request.  When false, DHCP server will simply attempt to update the DNS entries per the request, regardless of whether or not they conflict with existing entries owned by other DHCP4 clients.  Defaults to _true_. | [optional] [default to true]
**DhcpConfig** | Pointer to [**DHCPConfig**](DHCPConfig.md) |  | [optional] 
**DhcpOptions** | Pointer to [**[]OptionItem**](OptionItem.md) | The list of IPv4 DHCP options for IP space. May be either a specific option or a group of options. | [optional] 
**DhcpOptionsV6** | Pointer to [**[]OptionItem**](OptionItem.md) | The list of IPv6 DHCP options for IP space. May be either a specific option or a group of options. | [optional] 
**HeaderOptionFilename** | Pointer to **string** | The configuration for header option filename field. | [optional] 
**HeaderOptionServerAddress** | Pointer to **string** | The configuration for header option server address field. | [optional] 
**HeaderOptionServerName** | Pointer to **string** | The configuration for header option server name field. | [optional] 
**HostnameRewriteChar** | Pointer to **string** | The character to replace non-matching characters with, when hostname rewrite is enabled.  Any single ASCII character or no character if the invalid characters should be removed without replacement.  Defaults to \&quot;-\&quot;. | [optional] [default to "-"]
**HostnameRewriteEnabled** | Pointer to **bool** | Indicates if client supplied hostnames will be rewritten prior to DDNS update by replacing every character that does not match _hostname_rewrite_regex_ by _hostname_rewrite_char_.  Defaults to _false_. | [optional] [default to false]
**HostnameRewriteRegex** | Pointer to **string** | The regex bracket expression to match valid characters.  Must begin with \&quot;[\&quot; and end with \&quot;]\&quot; and be a compilable POSIX regex.  Defaults to \&quot;[^a-zA-Z0-9_.]\&quot;. | [optional] [default to "[^a-zA-Z0-9_.]"]
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InheritanceSources** | Pointer to [**IPSpaceInheritance**](IPSpaceInheritance.md) |  | [optional] 
**Name** | **string** | The name of the IP space. Must contain 1 to 256 characters. Can include UTF-8. | 
**Tags** | Pointer to **map[string]interface{}** | The tags for the IP space in JSON format. | [optional] 
**Threshold** | Pointer to [**UtilizationThreshold**](UtilizationThreshold.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 
**Utilization** | Pointer to [**Utilization**](Utilization.md) |  | [optional] 
**UtilizationV6** | Pointer to [**UtilizationV6**](UtilizationV6.md) |  | [optional] 
**VendorSpecificOptionOptionSpace** | Pointer to **string** | The resource identifier. | [optional] 

## Methods

### NewIPSpace

`func NewIPSpace(name string, ) *IPSpace`

NewIPSpace instantiates a new IPSpace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSpaceWithDefaults

`func NewIPSpaceWithDefaults() *IPSpace`

NewIPSpaceWithDefaults instantiates a new IPSpace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsmConfig

`func (o *IPSpace) GetAsmConfig() ASMConfig`

GetAsmConfig returns the AsmConfig field if non-nil, zero value otherwise.

### GetAsmConfigOk

`func (o *IPSpace) GetAsmConfigOk() (*ASMConfig, bool)`

GetAsmConfigOk returns a tuple with the AsmConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmConfig

`func (o *IPSpace) SetAsmConfig(v ASMConfig)`

SetAsmConfig sets AsmConfig field to given value.

### HasAsmConfig

`func (o *IPSpace) HasAsmConfig() bool`

HasAsmConfig returns a boolean if a field has been set.

### GetAsmScopeFlag

`func (o *IPSpace) GetAsmScopeFlag() int64`

GetAsmScopeFlag returns the AsmScopeFlag field if non-nil, zero value otherwise.

### GetAsmScopeFlagOk

`func (o *IPSpace) GetAsmScopeFlagOk() (*int64, bool)`

GetAsmScopeFlagOk returns a tuple with the AsmScopeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmScopeFlag

`func (o *IPSpace) SetAsmScopeFlag(v int64)`

SetAsmScopeFlag sets AsmScopeFlag field to given value.

### HasAsmScopeFlag

`func (o *IPSpace) HasAsmScopeFlag() bool`

HasAsmScopeFlag returns a boolean if a field has been set.

### GetComment

`func (o *IPSpace) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IPSpace) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IPSpace) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IPSpace) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IPSpace) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IPSpace) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IPSpace) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IPSpace) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDdnsClientUpdate

`func (o *IPSpace) GetDdnsClientUpdate() string`

GetDdnsClientUpdate returns the DdnsClientUpdate field if non-nil, zero value otherwise.

### GetDdnsClientUpdateOk

`func (o *IPSpace) GetDdnsClientUpdateOk() (*string, bool)`

GetDdnsClientUpdateOk returns a tuple with the DdnsClientUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsClientUpdate

`func (o *IPSpace) SetDdnsClientUpdate(v string)`

SetDdnsClientUpdate sets DdnsClientUpdate field to given value.

### HasDdnsClientUpdate

`func (o *IPSpace) HasDdnsClientUpdate() bool`

HasDdnsClientUpdate returns a boolean if a field has been set.

### GetDdnsConflictResolutionMode

`func (o *IPSpace) GetDdnsConflictResolutionMode() string`

GetDdnsConflictResolutionMode returns the DdnsConflictResolutionMode field if non-nil, zero value otherwise.

### GetDdnsConflictResolutionModeOk

`func (o *IPSpace) GetDdnsConflictResolutionModeOk() (*string, bool)`

GetDdnsConflictResolutionModeOk returns a tuple with the DdnsConflictResolutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsConflictResolutionMode

`func (o *IPSpace) SetDdnsConflictResolutionMode(v string)`

SetDdnsConflictResolutionMode sets DdnsConflictResolutionMode field to given value.

### HasDdnsConflictResolutionMode

`func (o *IPSpace) HasDdnsConflictResolutionMode() bool`

HasDdnsConflictResolutionMode returns a boolean if a field has been set.

### GetDdnsDomain

`func (o *IPSpace) GetDdnsDomain() string`

GetDdnsDomain returns the DdnsDomain field if non-nil, zero value otherwise.

### GetDdnsDomainOk

`func (o *IPSpace) GetDdnsDomainOk() (*string, bool)`

GetDdnsDomainOk returns a tuple with the DdnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDomain

`func (o *IPSpace) SetDdnsDomain(v string)`

SetDdnsDomain sets DdnsDomain field to given value.

### HasDdnsDomain

`func (o *IPSpace) HasDdnsDomain() bool`

HasDdnsDomain returns a boolean if a field has been set.

### GetDdnsGenerateName

`func (o *IPSpace) GetDdnsGenerateName() bool`

GetDdnsGenerateName returns the DdnsGenerateName field if non-nil, zero value otherwise.

### GetDdnsGenerateNameOk

`func (o *IPSpace) GetDdnsGenerateNameOk() (*bool, bool)`

GetDdnsGenerateNameOk returns a tuple with the DdnsGenerateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGenerateName

`func (o *IPSpace) SetDdnsGenerateName(v bool)`

SetDdnsGenerateName sets DdnsGenerateName field to given value.

### HasDdnsGenerateName

`func (o *IPSpace) HasDdnsGenerateName() bool`

HasDdnsGenerateName returns a boolean if a field has been set.

### GetDdnsGeneratedPrefix

`func (o *IPSpace) GetDdnsGeneratedPrefix() string`

GetDdnsGeneratedPrefix returns the DdnsGeneratedPrefix field if non-nil, zero value otherwise.

### GetDdnsGeneratedPrefixOk

`func (o *IPSpace) GetDdnsGeneratedPrefixOk() (*string, bool)`

GetDdnsGeneratedPrefixOk returns a tuple with the DdnsGeneratedPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsGeneratedPrefix

`func (o *IPSpace) SetDdnsGeneratedPrefix(v string)`

SetDdnsGeneratedPrefix sets DdnsGeneratedPrefix field to given value.

### HasDdnsGeneratedPrefix

`func (o *IPSpace) HasDdnsGeneratedPrefix() bool`

HasDdnsGeneratedPrefix returns a boolean if a field has been set.

### GetDdnsSendUpdates

`func (o *IPSpace) GetDdnsSendUpdates() bool`

GetDdnsSendUpdates returns the DdnsSendUpdates field if non-nil, zero value otherwise.

### GetDdnsSendUpdatesOk

`func (o *IPSpace) GetDdnsSendUpdatesOk() (*bool, bool)`

GetDdnsSendUpdatesOk returns a tuple with the DdnsSendUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsSendUpdates

`func (o *IPSpace) SetDdnsSendUpdates(v bool)`

SetDdnsSendUpdates sets DdnsSendUpdates field to given value.

### HasDdnsSendUpdates

`func (o *IPSpace) HasDdnsSendUpdates() bool`

HasDdnsSendUpdates returns a boolean if a field has been set.

### GetDdnsTtlPercent

`func (o *IPSpace) GetDdnsTtlPercent() float32`

GetDdnsTtlPercent returns the DdnsTtlPercent field if non-nil, zero value otherwise.

### GetDdnsTtlPercentOk

`func (o *IPSpace) GetDdnsTtlPercentOk() (*float32, bool)`

GetDdnsTtlPercentOk returns a tuple with the DdnsTtlPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsTtlPercent

`func (o *IPSpace) SetDdnsTtlPercent(v float32)`

SetDdnsTtlPercent sets DdnsTtlPercent field to given value.

### HasDdnsTtlPercent

`func (o *IPSpace) HasDdnsTtlPercent() bool`

HasDdnsTtlPercent returns a boolean if a field has been set.

### GetDdnsUpdateOnRenew

`func (o *IPSpace) GetDdnsUpdateOnRenew() bool`

GetDdnsUpdateOnRenew returns the DdnsUpdateOnRenew field if non-nil, zero value otherwise.

### GetDdnsUpdateOnRenewOk

`func (o *IPSpace) GetDdnsUpdateOnRenewOk() (*bool, bool)`

GetDdnsUpdateOnRenewOk returns a tuple with the DdnsUpdateOnRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUpdateOnRenew

`func (o *IPSpace) SetDdnsUpdateOnRenew(v bool)`

SetDdnsUpdateOnRenew sets DdnsUpdateOnRenew field to given value.

### HasDdnsUpdateOnRenew

`func (o *IPSpace) HasDdnsUpdateOnRenew() bool`

HasDdnsUpdateOnRenew returns a boolean if a field has been set.

### GetDdnsUseConflictResolution

`func (o *IPSpace) GetDdnsUseConflictResolution() bool`

GetDdnsUseConflictResolution returns the DdnsUseConflictResolution field if non-nil, zero value otherwise.

### GetDdnsUseConflictResolutionOk

`func (o *IPSpace) GetDdnsUseConflictResolutionOk() (*bool, bool)`

GetDdnsUseConflictResolutionOk returns a tuple with the DdnsUseConflictResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUseConflictResolution

`func (o *IPSpace) SetDdnsUseConflictResolution(v bool)`

SetDdnsUseConflictResolution sets DdnsUseConflictResolution field to given value.

### HasDdnsUseConflictResolution

`func (o *IPSpace) HasDdnsUseConflictResolution() bool`

HasDdnsUseConflictResolution returns a boolean if a field has been set.

### GetDhcpConfig

`func (o *IPSpace) GetDhcpConfig() DHCPConfig`

GetDhcpConfig returns the DhcpConfig field if non-nil, zero value otherwise.

### GetDhcpConfigOk

`func (o *IPSpace) GetDhcpConfigOk() (*DHCPConfig, bool)`

GetDhcpConfigOk returns a tuple with the DhcpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpConfig

`func (o *IPSpace) SetDhcpConfig(v DHCPConfig)`

SetDhcpConfig sets DhcpConfig field to given value.

### HasDhcpConfig

`func (o *IPSpace) HasDhcpConfig() bool`

HasDhcpConfig returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *IPSpace) GetDhcpOptions() []OptionItem`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *IPSpace) GetDhcpOptionsOk() (*[]OptionItem, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *IPSpace) SetDhcpOptions(v []OptionItem)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *IPSpace) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetDhcpOptionsV6

`func (o *IPSpace) GetDhcpOptionsV6() []OptionItem`

GetDhcpOptionsV6 returns the DhcpOptionsV6 field if non-nil, zero value otherwise.

### GetDhcpOptionsV6Ok

`func (o *IPSpace) GetDhcpOptionsV6Ok() (*[]OptionItem, bool)`

GetDhcpOptionsV6Ok returns a tuple with the DhcpOptionsV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptionsV6

`func (o *IPSpace) SetDhcpOptionsV6(v []OptionItem)`

SetDhcpOptionsV6 sets DhcpOptionsV6 field to given value.

### HasDhcpOptionsV6

`func (o *IPSpace) HasDhcpOptionsV6() bool`

HasDhcpOptionsV6 returns a boolean if a field has been set.

### GetHeaderOptionFilename

`func (o *IPSpace) GetHeaderOptionFilename() string`

GetHeaderOptionFilename returns the HeaderOptionFilename field if non-nil, zero value otherwise.

### GetHeaderOptionFilenameOk

`func (o *IPSpace) GetHeaderOptionFilenameOk() (*string, bool)`

GetHeaderOptionFilenameOk returns a tuple with the HeaderOptionFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionFilename

`func (o *IPSpace) SetHeaderOptionFilename(v string)`

SetHeaderOptionFilename sets HeaderOptionFilename field to given value.

### HasHeaderOptionFilename

`func (o *IPSpace) HasHeaderOptionFilename() bool`

HasHeaderOptionFilename returns a boolean if a field has been set.

### GetHeaderOptionServerAddress

`func (o *IPSpace) GetHeaderOptionServerAddress() string`

GetHeaderOptionServerAddress returns the HeaderOptionServerAddress field if non-nil, zero value otherwise.

### GetHeaderOptionServerAddressOk

`func (o *IPSpace) GetHeaderOptionServerAddressOk() (*string, bool)`

GetHeaderOptionServerAddressOk returns a tuple with the HeaderOptionServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerAddress

`func (o *IPSpace) SetHeaderOptionServerAddress(v string)`

SetHeaderOptionServerAddress sets HeaderOptionServerAddress field to given value.

### HasHeaderOptionServerAddress

`func (o *IPSpace) HasHeaderOptionServerAddress() bool`

HasHeaderOptionServerAddress returns a boolean if a field has been set.

### GetHeaderOptionServerName

`func (o *IPSpace) GetHeaderOptionServerName() string`

GetHeaderOptionServerName returns the HeaderOptionServerName field if non-nil, zero value otherwise.

### GetHeaderOptionServerNameOk

`func (o *IPSpace) GetHeaderOptionServerNameOk() (*string, bool)`

GetHeaderOptionServerNameOk returns a tuple with the HeaderOptionServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerName

`func (o *IPSpace) SetHeaderOptionServerName(v string)`

SetHeaderOptionServerName sets HeaderOptionServerName field to given value.

### HasHeaderOptionServerName

`func (o *IPSpace) HasHeaderOptionServerName() bool`

HasHeaderOptionServerName returns a boolean if a field has been set.

### GetHostnameRewriteChar

`func (o *IPSpace) GetHostnameRewriteChar() string`

GetHostnameRewriteChar returns the HostnameRewriteChar field if non-nil, zero value otherwise.

### GetHostnameRewriteCharOk

`func (o *IPSpace) GetHostnameRewriteCharOk() (*string, bool)`

GetHostnameRewriteCharOk returns a tuple with the HostnameRewriteChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameRewriteChar

`func (o *IPSpace) SetHostnameRewriteChar(v string)`

SetHostnameRewriteChar sets HostnameRewriteChar field to given value.

### HasHostnameRewriteChar

`func (o *IPSpace) HasHostnameRewriteChar() bool`

HasHostnameRewriteChar returns a boolean if a field has been set.

### GetHostnameRewriteEnabled

`func (o *IPSpace) GetHostnameRewriteEnabled() bool`

GetHostnameRewriteEnabled returns the HostnameRewriteEnabled field if non-nil, zero value otherwise.

### GetHostnameRewriteEnabledOk

`func (o *IPSpace) GetHostnameRewriteEnabledOk() (*bool, bool)`

GetHostnameRewriteEnabledOk returns a tuple with the HostnameRewriteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameRewriteEnabled

`func (o *IPSpace) SetHostnameRewriteEnabled(v bool)`

SetHostnameRewriteEnabled sets HostnameRewriteEnabled field to given value.

### HasHostnameRewriteEnabled

`func (o *IPSpace) HasHostnameRewriteEnabled() bool`

HasHostnameRewriteEnabled returns a boolean if a field has been set.

### GetHostnameRewriteRegex

`func (o *IPSpace) GetHostnameRewriteRegex() string`

GetHostnameRewriteRegex returns the HostnameRewriteRegex field if non-nil, zero value otherwise.

### GetHostnameRewriteRegexOk

`func (o *IPSpace) GetHostnameRewriteRegexOk() (*string, bool)`

GetHostnameRewriteRegexOk returns a tuple with the HostnameRewriteRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameRewriteRegex

`func (o *IPSpace) SetHostnameRewriteRegex(v string)`

SetHostnameRewriteRegex sets HostnameRewriteRegex field to given value.

### HasHostnameRewriteRegex

`func (o *IPSpace) HasHostnameRewriteRegex() bool`

HasHostnameRewriteRegex returns a boolean if a field has been set.

### GetId

`func (o *IPSpace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPSpace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPSpace) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IPSpace) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritanceSources

`func (o *IPSpace) GetInheritanceSources() IPSpaceInheritance`

GetInheritanceSources returns the InheritanceSources field if non-nil, zero value otherwise.

### GetInheritanceSourcesOk

`func (o *IPSpace) GetInheritanceSourcesOk() (*IPSpaceInheritance, bool)`

GetInheritanceSourcesOk returns a tuple with the InheritanceSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceSources

`func (o *IPSpace) SetInheritanceSources(v IPSpaceInheritance)`

SetInheritanceSources sets InheritanceSources field to given value.

### HasInheritanceSources

`func (o *IPSpace) HasInheritanceSources() bool`

HasInheritanceSources returns a boolean if a field has been set.

### GetName

`func (o *IPSpace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IPSpace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IPSpace) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *IPSpace) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IPSpace) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IPSpace) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *IPSpace) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThreshold

`func (o *IPSpace) GetThreshold() UtilizationThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *IPSpace) GetThresholdOk() (*UtilizationThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *IPSpace) SetThreshold(v UtilizationThreshold)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *IPSpace) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IPSpace) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IPSpace) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IPSpace) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IPSpace) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUtilization

`func (o *IPSpace) GetUtilization() Utilization`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *IPSpace) GetUtilizationOk() (*Utilization, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *IPSpace) SetUtilization(v Utilization)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *IPSpace) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.

### GetUtilizationV6

`func (o *IPSpace) GetUtilizationV6() UtilizationV6`

GetUtilizationV6 returns the UtilizationV6 field if non-nil, zero value otherwise.

### GetUtilizationV6Ok

`func (o *IPSpace) GetUtilizationV6Ok() (*UtilizationV6, bool)`

GetUtilizationV6Ok returns a tuple with the UtilizationV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationV6

`func (o *IPSpace) SetUtilizationV6(v UtilizationV6)`

SetUtilizationV6 sets UtilizationV6 field to given value.

### HasUtilizationV6

`func (o *IPSpace) HasUtilizationV6() bool`

HasUtilizationV6 returns a boolean if a field has been set.

### GetVendorSpecificOptionOptionSpace

`func (o *IPSpace) GetVendorSpecificOptionOptionSpace() string`

GetVendorSpecificOptionOptionSpace returns the VendorSpecificOptionOptionSpace field if non-nil, zero value otherwise.

### GetVendorSpecificOptionOptionSpaceOk

`func (o *IPSpace) GetVendorSpecificOptionOptionSpaceOk() (*string, bool)`

GetVendorSpecificOptionOptionSpaceOk returns a tuple with the VendorSpecificOptionOptionSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSpecificOptionOptionSpace

`func (o *IPSpace) SetVendorSpecificOptionOptionSpace(v string)`

SetVendorSpecificOptionOptionSpace sets VendorSpecificOptionOptionSpace field to given value.

### HasVendorSpecificOptionOptionSpace

`func (o *IPSpace) HasVendorSpecificOptionOptionSpace() bool`

HasVendorSpecificOptionOptionSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


