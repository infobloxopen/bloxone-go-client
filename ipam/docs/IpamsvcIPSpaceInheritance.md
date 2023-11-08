# IpamsvcIPSpaceInheritance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsmConfig** | Pointer to [**IpamsvcInheritedASMConfig**](IpamsvcInheritedASMConfig.md) |  | [optional] 
**DdnsClientUpdate** | Pointer to [**InheritanceInheritedString**](InheritanceInheritedString.md) |  | [optional] 
**DdnsConflictResolutionMode** | Pointer to [**InheritanceInheritedString**](InheritanceInheritedString.md) |  | [optional] 
**DdnsEnabled** | Pointer to [**InheritanceInheritedBool**](InheritanceInheritedBool.md) |  | [optional] 
**DdnsHostnameBlock** | Pointer to [**IpamsvcInheritedDDNSHostnameBlock**](IpamsvcInheritedDDNSHostnameBlock.md) |  | [optional] 
**DdnsTtlPercent** | Pointer to [**InheritanceInheritedFloat**](InheritanceInheritedFloat.md) |  | [optional] 
**DdnsUpdateBlock** | Pointer to [**IpamsvcInheritedDDNSUpdateBlock**](IpamsvcInheritedDDNSUpdateBlock.md) |  | [optional] 
**DdnsUpdateOnRenew** | Pointer to [**InheritanceInheritedBool**](InheritanceInheritedBool.md) |  | [optional] 
**DdnsUseConflictResolution** | Pointer to [**InheritanceInheritedBool**](InheritanceInheritedBool.md) |  | [optional] 
**DhcpConfig** | Pointer to [**IpamsvcInheritedDHCPConfig**](IpamsvcInheritedDHCPConfig.md) |  | [optional] 
**DhcpOptions** | Pointer to [**IpamsvcInheritedDHCPOptionList**](IpamsvcInheritedDHCPOptionList.md) |  | [optional] 
**DhcpOptionsV6** | Pointer to [**IpamsvcInheritedDHCPOptionList**](IpamsvcInheritedDHCPOptionList.md) |  | [optional] 
**HeaderOptionFilename** | Pointer to [**InheritanceInheritedString**](InheritanceInheritedString.md) |  | [optional] 
**HeaderOptionServerAddress** | Pointer to [**InheritanceInheritedString**](InheritanceInheritedString.md) |  | [optional] 
**HeaderOptionServerName** | Pointer to [**InheritanceInheritedString**](InheritanceInheritedString.md) |  | [optional] 
**HostnameRewriteBlock** | Pointer to [**IpamsvcInheritedHostnameRewriteBlock**](IpamsvcInheritedHostnameRewriteBlock.md) |  | [optional] 
**VendorSpecificOptionOptionSpace** | Pointer to [**InheritanceInheritedIdentifier**](InheritanceInheritedIdentifier.md) |  | [optional] 

## Methods

### NewIpamsvcIPSpaceInheritance

`func NewIpamsvcIPSpaceInheritance() *IpamsvcIPSpaceInheritance`

NewIpamsvcIPSpaceInheritance instantiates a new IpamsvcIPSpaceInheritance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcIPSpaceInheritanceWithDefaults

`func NewIpamsvcIPSpaceInheritanceWithDefaults() *IpamsvcIPSpaceInheritance`

NewIpamsvcIPSpaceInheritanceWithDefaults instantiates a new IpamsvcIPSpaceInheritance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsmConfig

`func (o *IpamsvcIPSpaceInheritance) GetAsmConfig() IpamsvcInheritedASMConfig`

GetAsmConfig returns the AsmConfig field if non-nil, zero value otherwise.

### GetAsmConfigOk

`func (o *IpamsvcIPSpaceInheritance) GetAsmConfigOk() (*IpamsvcInheritedASMConfig, bool)`

GetAsmConfigOk returns a tuple with the AsmConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmConfig

`func (o *IpamsvcIPSpaceInheritance) SetAsmConfig(v IpamsvcInheritedASMConfig)`

SetAsmConfig sets AsmConfig field to given value.

### HasAsmConfig

`func (o *IpamsvcIPSpaceInheritance) HasAsmConfig() bool`

HasAsmConfig returns a boolean if a field has been set.

### GetDdnsClientUpdate

`func (o *IpamsvcIPSpaceInheritance) GetDdnsClientUpdate() InheritanceInheritedString`

GetDdnsClientUpdate returns the DdnsClientUpdate field if non-nil, zero value otherwise.

### GetDdnsClientUpdateOk

`func (o *IpamsvcIPSpaceInheritance) GetDdnsClientUpdateOk() (*InheritanceInheritedString, bool)`

GetDdnsClientUpdateOk returns a tuple with the DdnsClientUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsClientUpdate

`func (o *IpamsvcIPSpaceInheritance) SetDdnsClientUpdate(v InheritanceInheritedString)`

SetDdnsClientUpdate sets DdnsClientUpdate field to given value.

### HasDdnsClientUpdate

`func (o *IpamsvcIPSpaceInheritance) HasDdnsClientUpdate() bool`

HasDdnsClientUpdate returns a boolean if a field has been set.

### GetDdnsConflictResolutionMode

`func (o *IpamsvcIPSpaceInheritance) GetDdnsConflictResolutionMode() InheritanceInheritedString`

GetDdnsConflictResolutionMode returns the DdnsConflictResolutionMode field if non-nil, zero value otherwise.

### GetDdnsConflictResolutionModeOk

`func (o *IpamsvcIPSpaceInheritance) GetDdnsConflictResolutionModeOk() (*InheritanceInheritedString, bool)`

GetDdnsConflictResolutionModeOk returns a tuple with the DdnsConflictResolutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsConflictResolutionMode

`func (o *IpamsvcIPSpaceInheritance) SetDdnsConflictResolutionMode(v InheritanceInheritedString)`

SetDdnsConflictResolutionMode sets DdnsConflictResolutionMode field to given value.

### HasDdnsConflictResolutionMode

`func (o *IpamsvcIPSpaceInheritance) HasDdnsConflictResolutionMode() bool`

HasDdnsConflictResolutionMode returns a boolean if a field has been set.

### GetDdnsEnabled

`func (o *IpamsvcIPSpaceInheritance) GetDdnsEnabled() InheritanceInheritedBool`

GetDdnsEnabled returns the DdnsEnabled field if non-nil, zero value otherwise.

### GetDdnsEnabledOk

`func (o *IpamsvcIPSpaceInheritance) GetDdnsEnabledOk() (*InheritanceInheritedBool, bool)`

GetDdnsEnabledOk returns a tuple with the DdnsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsEnabled

`func (o *IpamsvcIPSpaceInheritance) SetDdnsEnabled(v InheritanceInheritedBool)`

SetDdnsEnabled sets DdnsEnabled field to given value.

### HasDdnsEnabled

`func (o *IpamsvcIPSpaceInheritance) HasDdnsEnabled() bool`

HasDdnsEnabled returns a boolean if a field has been set.

### GetDdnsHostnameBlock

`func (o *IpamsvcIPSpaceInheritance) GetDdnsHostnameBlock() IpamsvcInheritedDDNSHostnameBlock`

GetDdnsHostnameBlock returns the DdnsHostnameBlock field if non-nil, zero value otherwise.

### GetDdnsHostnameBlockOk

`func (o *IpamsvcIPSpaceInheritance) GetDdnsHostnameBlockOk() (*IpamsvcInheritedDDNSHostnameBlock, bool)`

GetDdnsHostnameBlockOk returns a tuple with the DdnsHostnameBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsHostnameBlock

`func (o *IpamsvcIPSpaceInheritance) SetDdnsHostnameBlock(v IpamsvcInheritedDDNSHostnameBlock)`

SetDdnsHostnameBlock sets DdnsHostnameBlock field to given value.

### HasDdnsHostnameBlock

`func (o *IpamsvcIPSpaceInheritance) HasDdnsHostnameBlock() bool`

HasDdnsHostnameBlock returns a boolean if a field has been set.

### GetDdnsTtlPercent

`func (o *IpamsvcIPSpaceInheritance) GetDdnsTtlPercent() InheritanceInheritedFloat`

GetDdnsTtlPercent returns the DdnsTtlPercent field if non-nil, zero value otherwise.

### GetDdnsTtlPercentOk

`func (o *IpamsvcIPSpaceInheritance) GetDdnsTtlPercentOk() (*InheritanceInheritedFloat, bool)`

GetDdnsTtlPercentOk returns a tuple with the DdnsTtlPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsTtlPercent

`func (o *IpamsvcIPSpaceInheritance) SetDdnsTtlPercent(v InheritanceInheritedFloat)`

SetDdnsTtlPercent sets DdnsTtlPercent field to given value.

### HasDdnsTtlPercent

`func (o *IpamsvcIPSpaceInheritance) HasDdnsTtlPercent() bool`

HasDdnsTtlPercent returns a boolean if a field has been set.

### GetDdnsUpdateBlock

`func (o *IpamsvcIPSpaceInheritance) GetDdnsUpdateBlock() IpamsvcInheritedDDNSUpdateBlock`

GetDdnsUpdateBlock returns the DdnsUpdateBlock field if non-nil, zero value otherwise.

### GetDdnsUpdateBlockOk

`func (o *IpamsvcIPSpaceInheritance) GetDdnsUpdateBlockOk() (*IpamsvcInheritedDDNSUpdateBlock, bool)`

GetDdnsUpdateBlockOk returns a tuple with the DdnsUpdateBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUpdateBlock

`func (o *IpamsvcIPSpaceInheritance) SetDdnsUpdateBlock(v IpamsvcInheritedDDNSUpdateBlock)`

SetDdnsUpdateBlock sets DdnsUpdateBlock field to given value.

### HasDdnsUpdateBlock

`func (o *IpamsvcIPSpaceInheritance) HasDdnsUpdateBlock() bool`

HasDdnsUpdateBlock returns a boolean if a field has been set.

### GetDdnsUpdateOnRenew

`func (o *IpamsvcIPSpaceInheritance) GetDdnsUpdateOnRenew() InheritanceInheritedBool`

GetDdnsUpdateOnRenew returns the DdnsUpdateOnRenew field if non-nil, zero value otherwise.

### GetDdnsUpdateOnRenewOk

`func (o *IpamsvcIPSpaceInheritance) GetDdnsUpdateOnRenewOk() (*InheritanceInheritedBool, bool)`

GetDdnsUpdateOnRenewOk returns a tuple with the DdnsUpdateOnRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUpdateOnRenew

`func (o *IpamsvcIPSpaceInheritance) SetDdnsUpdateOnRenew(v InheritanceInheritedBool)`

SetDdnsUpdateOnRenew sets DdnsUpdateOnRenew field to given value.

### HasDdnsUpdateOnRenew

`func (o *IpamsvcIPSpaceInheritance) HasDdnsUpdateOnRenew() bool`

HasDdnsUpdateOnRenew returns a boolean if a field has been set.

### GetDdnsUseConflictResolution

`func (o *IpamsvcIPSpaceInheritance) GetDdnsUseConflictResolution() InheritanceInheritedBool`

GetDdnsUseConflictResolution returns the DdnsUseConflictResolution field if non-nil, zero value otherwise.

### GetDdnsUseConflictResolutionOk

`func (o *IpamsvcIPSpaceInheritance) GetDdnsUseConflictResolutionOk() (*InheritanceInheritedBool, bool)`

GetDdnsUseConflictResolutionOk returns a tuple with the DdnsUseConflictResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUseConflictResolution

`func (o *IpamsvcIPSpaceInheritance) SetDdnsUseConflictResolution(v InheritanceInheritedBool)`

SetDdnsUseConflictResolution sets DdnsUseConflictResolution field to given value.

### HasDdnsUseConflictResolution

`func (o *IpamsvcIPSpaceInheritance) HasDdnsUseConflictResolution() bool`

HasDdnsUseConflictResolution returns a boolean if a field has been set.

### GetDhcpConfig

`func (o *IpamsvcIPSpaceInheritance) GetDhcpConfig() IpamsvcInheritedDHCPConfig`

GetDhcpConfig returns the DhcpConfig field if non-nil, zero value otherwise.

### GetDhcpConfigOk

`func (o *IpamsvcIPSpaceInheritance) GetDhcpConfigOk() (*IpamsvcInheritedDHCPConfig, bool)`

GetDhcpConfigOk returns a tuple with the DhcpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpConfig

`func (o *IpamsvcIPSpaceInheritance) SetDhcpConfig(v IpamsvcInheritedDHCPConfig)`

SetDhcpConfig sets DhcpConfig field to given value.

### HasDhcpConfig

`func (o *IpamsvcIPSpaceInheritance) HasDhcpConfig() bool`

HasDhcpConfig returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *IpamsvcIPSpaceInheritance) GetDhcpOptions() IpamsvcInheritedDHCPOptionList`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *IpamsvcIPSpaceInheritance) GetDhcpOptionsOk() (*IpamsvcInheritedDHCPOptionList, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *IpamsvcIPSpaceInheritance) SetDhcpOptions(v IpamsvcInheritedDHCPOptionList)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *IpamsvcIPSpaceInheritance) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetDhcpOptionsV6

`func (o *IpamsvcIPSpaceInheritance) GetDhcpOptionsV6() IpamsvcInheritedDHCPOptionList`

GetDhcpOptionsV6 returns the DhcpOptionsV6 field if non-nil, zero value otherwise.

### GetDhcpOptionsV6Ok

`func (o *IpamsvcIPSpaceInheritance) GetDhcpOptionsV6Ok() (*IpamsvcInheritedDHCPOptionList, bool)`

GetDhcpOptionsV6Ok returns a tuple with the DhcpOptionsV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptionsV6

`func (o *IpamsvcIPSpaceInheritance) SetDhcpOptionsV6(v IpamsvcInheritedDHCPOptionList)`

SetDhcpOptionsV6 sets DhcpOptionsV6 field to given value.

### HasDhcpOptionsV6

`func (o *IpamsvcIPSpaceInheritance) HasDhcpOptionsV6() bool`

HasDhcpOptionsV6 returns a boolean if a field has been set.

### GetHeaderOptionFilename

`func (o *IpamsvcIPSpaceInheritance) GetHeaderOptionFilename() InheritanceInheritedString`

GetHeaderOptionFilename returns the HeaderOptionFilename field if non-nil, zero value otherwise.

### GetHeaderOptionFilenameOk

`func (o *IpamsvcIPSpaceInheritance) GetHeaderOptionFilenameOk() (*InheritanceInheritedString, bool)`

GetHeaderOptionFilenameOk returns a tuple with the HeaderOptionFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionFilename

`func (o *IpamsvcIPSpaceInheritance) SetHeaderOptionFilename(v InheritanceInheritedString)`

SetHeaderOptionFilename sets HeaderOptionFilename field to given value.

### HasHeaderOptionFilename

`func (o *IpamsvcIPSpaceInheritance) HasHeaderOptionFilename() bool`

HasHeaderOptionFilename returns a boolean if a field has been set.

### GetHeaderOptionServerAddress

`func (o *IpamsvcIPSpaceInheritance) GetHeaderOptionServerAddress() InheritanceInheritedString`

GetHeaderOptionServerAddress returns the HeaderOptionServerAddress field if non-nil, zero value otherwise.

### GetHeaderOptionServerAddressOk

`func (o *IpamsvcIPSpaceInheritance) GetHeaderOptionServerAddressOk() (*InheritanceInheritedString, bool)`

GetHeaderOptionServerAddressOk returns a tuple with the HeaderOptionServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerAddress

`func (o *IpamsvcIPSpaceInheritance) SetHeaderOptionServerAddress(v InheritanceInheritedString)`

SetHeaderOptionServerAddress sets HeaderOptionServerAddress field to given value.

### HasHeaderOptionServerAddress

`func (o *IpamsvcIPSpaceInheritance) HasHeaderOptionServerAddress() bool`

HasHeaderOptionServerAddress returns a boolean if a field has been set.

### GetHeaderOptionServerName

`func (o *IpamsvcIPSpaceInheritance) GetHeaderOptionServerName() InheritanceInheritedString`

GetHeaderOptionServerName returns the HeaderOptionServerName field if non-nil, zero value otherwise.

### GetHeaderOptionServerNameOk

`func (o *IpamsvcIPSpaceInheritance) GetHeaderOptionServerNameOk() (*InheritanceInheritedString, bool)`

GetHeaderOptionServerNameOk returns a tuple with the HeaderOptionServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerName

`func (o *IpamsvcIPSpaceInheritance) SetHeaderOptionServerName(v InheritanceInheritedString)`

SetHeaderOptionServerName sets HeaderOptionServerName field to given value.

### HasHeaderOptionServerName

`func (o *IpamsvcIPSpaceInheritance) HasHeaderOptionServerName() bool`

HasHeaderOptionServerName returns a boolean if a field has been set.

### GetHostnameRewriteBlock

`func (o *IpamsvcIPSpaceInheritance) GetHostnameRewriteBlock() IpamsvcInheritedHostnameRewriteBlock`

GetHostnameRewriteBlock returns the HostnameRewriteBlock field if non-nil, zero value otherwise.

### GetHostnameRewriteBlockOk

`func (o *IpamsvcIPSpaceInheritance) GetHostnameRewriteBlockOk() (*IpamsvcInheritedHostnameRewriteBlock, bool)`

GetHostnameRewriteBlockOk returns a tuple with the HostnameRewriteBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameRewriteBlock

`func (o *IpamsvcIPSpaceInheritance) SetHostnameRewriteBlock(v IpamsvcInheritedHostnameRewriteBlock)`

SetHostnameRewriteBlock sets HostnameRewriteBlock field to given value.

### HasHostnameRewriteBlock

`func (o *IpamsvcIPSpaceInheritance) HasHostnameRewriteBlock() bool`

HasHostnameRewriteBlock returns a boolean if a field has been set.

### GetVendorSpecificOptionOptionSpace

`func (o *IpamsvcIPSpaceInheritance) GetVendorSpecificOptionOptionSpace() InheritanceInheritedIdentifier`

GetVendorSpecificOptionOptionSpace returns the VendorSpecificOptionOptionSpace field if non-nil, zero value otherwise.

### GetVendorSpecificOptionOptionSpaceOk

`func (o *IpamsvcIPSpaceInheritance) GetVendorSpecificOptionOptionSpaceOk() (*InheritanceInheritedIdentifier, bool)`

GetVendorSpecificOptionOptionSpaceOk returns a tuple with the VendorSpecificOptionOptionSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSpecificOptionOptionSpace

`func (o *IpamsvcIPSpaceInheritance) SetVendorSpecificOptionOptionSpace(v InheritanceInheritedIdentifier)`

SetVendorSpecificOptionOptionSpace sets VendorSpecificOptionOptionSpace field to given value.

### HasVendorSpecificOptionOptionSpace

`func (o *IpamsvcIPSpaceInheritance) HasVendorSpecificOptionOptionSpace() bool`

HasVendorSpecificOptionOptionSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


