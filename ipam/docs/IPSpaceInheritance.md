# IPSpaceInheritance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsmConfig** | Pointer to [**InheritedASMConfig**](InheritedASMConfig.md) |  | [optional] 
**DdnsClientUpdate** | Pointer to [**InheritanceInheritedString**](InheritanceInheritedString.md) |  | [optional] 
**DdnsConflictResolutionMode** | Pointer to [**InheritanceInheritedString**](InheritanceInheritedString.md) |  | [optional] 
**DdnsEnabled** | Pointer to [**InheritanceInheritedBool**](InheritanceInheritedBool.md) |  | [optional] 
**DdnsHostnameBlock** | Pointer to [**InheritedDDNSHostnameBlock**](InheritedDDNSHostnameBlock.md) |  | [optional] 
**DdnsTtlPercent** | Pointer to [**InheritanceInheritedFloat**](InheritanceInheritedFloat.md) |  | [optional] 
**DdnsUpdateBlock** | Pointer to [**InheritedDDNSUpdateBlock**](InheritedDDNSUpdateBlock.md) |  | [optional] 
**DdnsUpdateOnRenew** | Pointer to [**InheritanceInheritedBool**](InheritanceInheritedBool.md) |  | [optional] 
**DdnsUseConflictResolution** | Pointer to [**InheritanceInheritedBool**](InheritanceInheritedBool.md) |  | [optional] 
**DhcpConfig** | Pointer to [**InheritedDHCPConfig**](InheritedDHCPConfig.md) |  | [optional] 
**DhcpOptions** | Pointer to [**InheritedDHCPOptionList**](InheritedDHCPOptionList.md) |  | [optional] 
**DhcpOptionsV6** | Pointer to [**InheritedDHCPOptionList**](InheritedDHCPOptionList.md) |  | [optional] 
**HeaderOptionFilename** | Pointer to [**InheritanceInheritedString**](InheritanceInheritedString.md) |  | [optional] 
**HeaderOptionServerAddress** | Pointer to [**InheritanceInheritedString**](InheritanceInheritedString.md) |  | [optional] 
**HeaderOptionServerName** | Pointer to [**InheritanceInheritedString**](InheritanceInheritedString.md) |  | [optional] 
**HostnameRewriteBlock** | Pointer to [**InheritedHostnameRewriteBlock**](InheritedHostnameRewriteBlock.md) |  | [optional] 
**VendorSpecificOptionOptionSpace** | Pointer to [**InheritanceInheritedIdentifier**](InheritanceInheritedIdentifier.md) |  | [optional] 

## Methods

### NewIPSpaceInheritance

`func NewIPSpaceInheritance() *IPSpaceInheritance`

NewIPSpaceInheritance instantiates a new IPSpaceInheritance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSpaceInheritanceWithDefaults

`func NewIPSpaceInheritanceWithDefaults() *IPSpaceInheritance`

NewIPSpaceInheritanceWithDefaults instantiates a new IPSpaceInheritance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsmConfig

`func (o *IPSpaceInheritance) GetAsmConfig() InheritedASMConfig`

GetAsmConfig returns the AsmConfig field if non-nil, zero value otherwise.

### GetAsmConfigOk

`func (o *IPSpaceInheritance) GetAsmConfigOk() (*InheritedASMConfig, bool)`

GetAsmConfigOk returns a tuple with the AsmConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmConfig

`func (o *IPSpaceInheritance) SetAsmConfig(v InheritedASMConfig)`

SetAsmConfig sets AsmConfig field to given value.

### HasAsmConfig

`func (o *IPSpaceInheritance) HasAsmConfig() bool`

HasAsmConfig returns a boolean if a field has been set.

### GetDdnsClientUpdate

`func (o *IPSpaceInheritance) GetDdnsClientUpdate() InheritanceInheritedString`

GetDdnsClientUpdate returns the DdnsClientUpdate field if non-nil, zero value otherwise.

### GetDdnsClientUpdateOk

`func (o *IPSpaceInheritance) GetDdnsClientUpdateOk() (*InheritanceInheritedString, bool)`

GetDdnsClientUpdateOk returns a tuple with the DdnsClientUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsClientUpdate

`func (o *IPSpaceInheritance) SetDdnsClientUpdate(v InheritanceInheritedString)`

SetDdnsClientUpdate sets DdnsClientUpdate field to given value.

### HasDdnsClientUpdate

`func (o *IPSpaceInheritance) HasDdnsClientUpdate() bool`

HasDdnsClientUpdate returns a boolean if a field has been set.

### GetDdnsConflictResolutionMode

`func (o *IPSpaceInheritance) GetDdnsConflictResolutionMode() InheritanceInheritedString`

GetDdnsConflictResolutionMode returns the DdnsConflictResolutionMode field if non-nil, zero value otherwise.

### GetDdnsConflictResolutionModeOk

`func (o *IPSpaceInheritance) GetDdnsConflictResolutionModeOk() (*InheritanceInheritedString, bool)`

GetDdnsConflictResolutionModeOk returns a tuple with the DdnsConflictResolutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsConflictResolutionMode

`func (o *IPSpaceInheritance) SetDdnsConflictResolutionMode(v InheritanceInheritedString)`

SetDdnsConflictResolutionMode sets DdnsConflictResolutionMode field to given value.

### HasDdnsConflictResolutionMode

`func (o *IPSpaceInheritance) HasDdnsConflictResolutionMode() bool`

HasDdnsConflictResolutionMode returns a boolean if a field has been set.

### GetDdnsEnabled

`func (o *IPSpaceInheritance) GetDdnsEnabled() InheritanceInheritedBool`

GetDdnsEnabled returns the DdnsEnabled field if non-nil, zero value otherwise.

### GetDdnsEnabledOk

`func (o *IPSpaceInheritance) GetDdnsEnabledOk() (*InheritanceInheritedBool, bool)`

GetDdnsEnabledOk returns a tuple with the DdnsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsEnabled

`func (o *IPSpaceInheritance) SetDdnsEnabled(v InheritanceInheritedBool)`

SetDdnsEnabled sets DdnsEnabled field to given value.

### HasDdnsEnabled

`func (o *IPSpaceInheritance) HasDdnsEnabled() bool`

HasDdnsEnabled returns a boolean if a field has been set.

### GetDdnsHostnameBlock

`func (o *IPSpaceInheritance) GetDdnsHostnameBlock() InheritedDDNSHostnameBlock`

GetDdnsHostnameBlock returns the DdnsHostnameBlock field if non-nil, zero value otherwise.

### GetDdnsHostnameBlockOk

`func (o *IPSpaceInheritance) GetDdnsHostnameBlockOk() (*InheritedDDNSHostnameBlock, bool)`

GetDdnsHostnameBlockOk returns a tuple with the DdnsHostnameBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsHostnameBlock

`func (o *IPSpaceInheritance) SetDdnsHostnameBlock(v InheritedDDNSHostnameBlock)`

SetDdnsHostnameBlock sets DdnsHostnameBlock field to given value.

### HasDdnsHostnameBlock

`func (o *IPSpaceInheritance) HasDdnsHostnameBlock() bool`

HasDdnsHostnameBlock returns a boolean if a field has been set.

### GetDdnsTtlPercent

`func (o *IPSpaceInheritance) GetDdnsTtlPercent() InheritanceInheritedFloat`

GetDdnsTtlPercent returns the DdnsTtlPercent field if non-nil, zero value otherwise.

### GetDdnsTtlPercentOk

`func (o *IPSpaceInheritance) GetDdnsTtlPercentOk() (*InheritanceInheritedFloat, bool)`

GetDdnsTtlPercentOk returns a tuple with the DdnsTtlPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsTtlPercent

`func (o *IPSpaceInheritance) SetDdnsTtlPercent(v InheritanceInheritedFloat)`

SetDdnsTtlPercent sets DdnsTtlPercent field to given value.

### HasDdnsTtlPercent

`func (o *IPSpaceInheritance) HasDdnsTtlPercent() bool`

HasDdnsTtlPercent returns a boolean if a field has been set.

### GetDdnsUpdateBlock

`func (o *IPSpaceInheritance) GetDdnsUpdateBlock() InheritedDDNSUpdateBlock`

GetDdnsUpdateBlock returns the DdnsUpdateBlock field if non-nil, zero value otherwise.

### GetDdnsUpdateBlockOk

`func (o *IPSpaceInheritance) GetDdnsUpdateBlockOk() (*InheritedDDNSUpdateBlock, bool)`

GetDdnsUpdateBlockOk returns a tuple with the DdnsUpdateBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUpdateBlock

`func (o *IPSpaceInheritance) SetDdnsUpdateBlock(v InheritedDDNSUpdateBlock)`

SetDdnsUpdateBlock sets DdnsUpdateBlock field to given value.

### HasDdnsUpdateBlock

`func (o *IPSpaceInheritance) HasDdnsUpdateBlock() bool`

HasDdnsUpdateBlock returns a boolean if a field has been set.

### GetDdnsUpdateOnRenew

`func (o *IPSpaceInheritance) GetDdnsUpdateOnRenew() InheritanceInheritedBool`

GetDdnsUpdateOnRenew returns the DdnsUpdateOnRenew field if non-nil, zero value otherwise.

### GetDdnsUpdateOnRenewOk

`func (o *IPSpaceInheritance) GetDdnsUpdateOnRenewOk() (*InheritanceInheritedBool, bool)`

GetDdnsUpdateOnRenewOk returns a tuple with the DdnsUpdateOnRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUpdateOnRenew

`func (o *IPSpaceInheritance) SetDdnsUpdateOnRenew(v InheritanceInheritedBool)`

SetDdnsUpdateOnRenew sets DdnsUpdateOnRenew field to given value.

### HasDdnsUpdateOnRenew

`func (o *IPSpaceInheritance) HasDdnsUpdateOnRenew() bool`

HasDdnsUpdateOnRenew returns a boolean if a field has been set.

### GetDdnsUseConflictResolution

`func (o *IPSpaceInheritance) GetDdnsUseConflictResolution() InheritanceInheritedBool`

GetDdnsUseConflictResolution returns the DdnsUseConflictResolution field if non-nil, zero value otherwise.

### GetDdnsUseConflictResolutionOk

`func (o *IPSpaceInheritance) GetDdnsUseConflictResolutionOk() (*InheritanceInheritedBool, bool)`

GetDdnsUseConflictResolutionOk returns a tuple with the DdnsUseConflictResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUseConflictResolution

`func (o *IPSpaceInheritance) SetDdnsUseConflictResolution(v InheritanceInheritedBool)`

SetDdnsUseConflictResolution sets DdnsUseConflictResolution field to given value.

### HasDdnsUseConflictResolution

`func (o *IPSpaceInheritance) HasDdnsUseConflictResolution() bool`

HasDdnsUseConflictResolution returns a boolean if a field has been set.

### GetDhcpConfig

`func (o *IPSpaceInheritance) GetDhcpConfig() InheritedDHCPConfig`

GetDhcpConfig returns the DhcpConfig field if non-nil, zero value otherwise.

### GetDhcpConfigOk

`func (o *IPSpaceInheritance) GetDhcpConfigOk() (*InheritedDHCPConfig, bool)`

GetDhcpConfigOk returns a tuple with the DhcpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpConfig

`func (o *IPSpaceInheritance) SetDhcpConfig(v InheritedDHCPConfig)`

SetDhcpConfig sets DhcpConfig field to given value.

### HasDhcpConfig

`func (o *IPSpaceInheritance) HasDhcpConfig() bool`

HasDhcpConfig returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *IPSpaceInheritance) GetDhcpOptions() InheritedDHCPOptionList`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *IPSpaceInheritance) GetDhcpOptionsOk() (*InheritedDHCPOptionList, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *IPSpaceInheritance) SetDhcpOptions(v InheritedDHCPOptionList)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *IPSpaceInheritance) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetDhcpOptionsV6

`func (o *IPSpaceInheritance) GetDhcpOptionsV6() InheritedDHCPOptionList`

GetDhcpOptionsV6 returns the DhcpOptionsV6 field if non-nil, zero value otherwise.

### GetDhcpOptionsV6Ok

`func (o *IPSpaceInheritance) GetDhcpOptionsV6Ok() (*InheritedDHCPOptionList, bool)`

GetDhcpOptionsV6Ok returns a tuple with the DhcpOptionsV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptionsV6

`func (o *IPSpaceInheritance) SetDhcpOptionsV6(v InheritedDHCPOptionList)`

SetDhcpOptionsV6 sets DhcpOptionsV6 field to given value.

### HasDhcpOptionsV6

`func (o *IPSpaceInheritance) HasDhcpOptionsV6() bool`

HasDhcpOptionsV6 returns a boolean if a field has been set.

### GetHeaderOptionFilename

`func (o *IPSpaceInheritance) GetHeaderOptionFilename() InheritanceInheritedString`

GetHeaderOptionFilename returns the HeaderOptionFilename field if non-nil, zero value otherwise.

### GetHeaderOptionFilenameOk

`func (o *IPSpaceInheritance) GetHeaderOptionFilenameOk() (*InheritanceInheritedString, bool)`

GetHeaderOptionFilenameOk returns a tuple with the HeaderOptionFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionFilename

`func (o *IPSpaceInheritance) SetHeaderOptionFilename(v InheritanceInheritedString)`

SetHeaderOptionFilename sets HeaderOptionFilename field to given value.

### HasHeaderOptionFilename

`func (o *IPSpaceInheritance) HasHeaderOptionFilename() bool`

HasHeaderOptionFilename returns a boolean if a field has been set.

### GetHeaderOptionServerAddress

`func (o *IPSpaceInheritance) GetHeaderOptionServerAddress() InheritanceInheritedString`

GetHeaderOptionServerAddress returns the HeaderOptionServerAddress field if non-nil, zero value otherwise.

### GetHeaderOptionServerAddressOk

`func (o *IPSpaceInheritance) GetHeaderOptionServerAddressOk() (*InheritanceInheritedString, bool)`

GetHeaderOptionServerAddressOk returns a tuple with the HeaderOptionServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerAddress

`func (o *IPSpaceInheritance) SetHeaderOptionServerAddress(v InheritanceInheritedString)`

SetHeaderOptionServerAddress sets HeaderOptionServerAddress field to given value.

### HasHeaderOptionServerAddress

`func (o *IPSpaceInheritance) HasHeaderOptionServerAddress() bool`

HasHeaderOptionServerAddress returns a boolean if a field has been set.

### GetHeaderOptionServerName

`func (o *IPSpaceInheritance) GetHeaderOptionServerName() InheritanceInheritedString`

GetHeaderOptionServerName returns the HeaderOptionServerName field if non-nil, zero value otherwise.

### GetHeaderOptionServerNameOk

`func (o *IPSpaceInheritance) GetHeaderOptionServerNameOk() (*InheritanceInheritedString, bool)`

GetHeaderOptionServerNameOk returns a tuple with the HeaderOptionServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerName

`func (o *IPSpaceInheritance) SetHeaderOptionServerName(v InheritanceInheritedString)`

SetHeaderOptionServerName sets HeaderOptionServerName field to given value.

### HasHeaderOptionServerName

`func (o *IPSpaceInheritance) HasHeaderOptionServerName() bool`

HasHeaderOptionServerName returns a boolean if a field has been set.

### GetHostnameRewriteBlock

`func (o *IPSpaceInheritance) GetHostnameRewriteBlock() InheritedHostnameRewriteBlock`

GetHostnameRewriteBlock returns the HostnameRewriteBlock field if non-nil, zero value otherwise.

### GetHostnameRewriteBlockOk

`func (o *IPSpaceInheritance) GetHostnameRewriteBlockOk() (*InheritedHostnameRewriteBlock, bool)`

GetHostnameRewriteBlockOk returns a tuple with the HostnameRewriteBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameRewriteBlock

`func (o *IPSpaceInheritance) SetHostnameRewriteBlock(v InheritedHostnameRewriteBlock)`

SetHostnameRewriteBlock sets HostnameRewriteBlock field to given value.

### HasHostnameRewriteBlock

`func (o *IPSpaceInheritance) HasHostnameRewriteBlock() bool`

HasHostnameRewriteBlock returns a boolean if a field has been set.

### GetVendorSpecificOptionOptionSpace

`func (o *IPSpaceInheritance) GetVendorSpecificOptionOptionSpace() InheritanceInheritedIdentifier`

GetVendorSpecificOptionOptionSpace returns the VendorSpecificOptionOptionSpace field if non-nil, zero value otherwise.

### GetVendorSpecificOptionOptionSpaceOk

`func (o *IPSpaceInheritance) GetVendorSpecificOptionOptionSpaceOk() (*InheritanceInheritedIdentifier, bool)`

GetVendorSpecificOptionOptionSpaceOk returns a tuple with the VendorSpecificOptionOptionSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSpecificOptionOptionSpace

`func (o *IPSpaceInheritance) SetVendorSpecificOptionOptionSpace(v InheritanceInheritedIdentifier)`

SetVendorSpecificOptionOptionSpace sets VendorSpecificOptionOptionSpace field to given value.

### HasVendorSpecificOptionOptionSpace

`func (o *IPSpaceInheritance) HasVendorSpecificOptionOptionSpace() bool`

HasVendorSpecificOptionOptionSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


