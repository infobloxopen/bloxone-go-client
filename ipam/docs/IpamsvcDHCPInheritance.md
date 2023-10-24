# IpamsvcDHCPInheritance

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
**HeaderOptionFilename** | Pointer to [**InheritanceInheritedString**](InheritanceInheritedString.md) |  | [optional] 
**HeaderOptionServerAddress** | Pointer to [**InheritanceInheritedString**](InheritanceInheritedString.md) |  | [optional] 
**HeaderOptionServerName** | Pointer to [**InheritanceInheritedString**](InheritanceInheritedString.md) |  | [optional] 
**HostnameRewriteBlock** | Pointer to [**IpamsvcInheritedHostnameRewriteBlock**](IpamsvcInheritedHostnameRewriteBlock.md) |  | [optional] 

## Methods

### NewIpamsvcDHCPInheritance

`func NewIpamsvcDHCPInheritance() *IpamsvcDHCPInheritance`

NewIpamsvcDHCPInheritance instantiates a new IpamsvcDHCPInheritance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcDHCPInheritanceWithDefaults

`func NewIpamsvcDHCPInheritanceWithDefaults() *IpamsvcDHCPInheritance`

NewIpamsvcDHCPInheritanceWithDefaults instantiates a new IpamsvcDHCPInheritance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsmConfig

`func (o *IpamsvcDHCPInheritance) GetAsmConfig() IpamsvcInheritedASMConfig`

GetAsmConfig returns the AsmConfig field if non-nil, zero value otherwise.

### GetAsmConfigOk

`func (o *IpamsvcDHCPInheritance) GetAsmConfigOk() (*IpamsvcInheritedASMConfig, bool)`

GetAsmConfigOk returns a tuple with the AsmConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmConfig

`func (o *IpamsvcDHCPInheritance) SetAsmConfig(v IpamsvcInheritedASMConfig)`

SetAsmConfig sets AsmConfig field to given value.

### HasAsmConfig

`func (o *IpamsvcDHCPInheritance) HasAsmConfig() bool`

HasAsmConfig returns a boolean if a field has been set.

### GetDdnsClientUpdate

`func (o *IpamsvcDHCPInheritance) GetDdnsClientUpdate() InheritanceInheritedString`

GetDdnsClientUpdate returns the DdnsClientUpdate field if non-nil, zero value otherwise.

### GetDdnsClientUpdateOk

`func (o *IpamsvcDHCPInheritance) GetDdnsClientUpdateOk() (*InheritanceInheritedString, bool)`

GetDdnsClientUpdateOk returns a tuple with the DdnsClientUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsClientUpdate

`func (o *IpamsvcDHCPInheritance) SetDdnsClientUpdate(v InheritanceInheritedString)`

SetDdnsClientUpdate sets DdnsClientUpdate field to given value.

### HasDdnsClientUpdate

`func (o *IpamsvcDHCPInheritance) HasDdnsClientUpdate() bool`

HasDdnsClientUpdate returns a boolean if a field has been set.

### GetDdnsConflictResolutionMode

`func (o *IpamsvcDHCPInheritance) GetDdnsConflictResolutionMode() InheritanceInheritedString`

GetDdnsConflictResolutionMode returns the DdnsConflictResolutionMode field if non-nil, zero value otherwise.

### GetDdnsConflictResolutionModeOk

`func (o *IpamsvcDHCPInheritance) GetDdnsConflictResolutionModeOk() (*InheritanceInheritedString, bool)`

GetDdnsConflictResolutionModeOk returns a tuple with the DdnsConflictResolutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsConflictResolutionMode

`func (o *IpamsvcDHCPInheritance) SetDdnsConflictResolutionMode(v InheritanceInheritedString)`

SetDdnsConflictResolutionMode sets DdnsConflictResolutionMode field to given value.

### HasDdnsConflictResolutionMode

`func (o *IpamsvcDHCPInheritance) HasDdnsConflictResolutionMode() bool`

HasDdnsConflictResolutionMode returns a boolean if a field has been set.

### GetDdnsEnabled

`func (o *IpamsvcDHCPInheritance) GetDdnsEnabled() InheritanceInheritedBool`

GetDdnsEnabled returns the DdnsEnabled field if non-nil, zero value otherwise.

### GetDdnsEnabledOk

`func (o *IpamsvcDHCPInheritance) GetDdnsEnabledOk() (*InheritanceInheritedBool, bool)`

GetDdnsEnabledOk returns a tuple with the DdnsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsEnabled

`func (o *IpamsvcDHCPInheritance) SetDdnsEnabled(v InheritanceInheritedBool)`

SetDdnsEnabled sets DdnsEnabled field to given value.

### HasDdnsEnabled

`func (o *IpamsvcDHCPInheritance) HasDdnsEnabled() bool`

HasDdnsEnabled returns a boolean if a field has been set.

### GetDdnsHostnameBlock

`func (o *IpamsvcDHCPInheritance) GetDdnsHostnameBlock() IpamsvcInheritedDDNSHostnameBlock`

GetDdnsHostnameBlock returns the DdnsHostnameBlock field if non-nil, zero value otherwise.

### GetDdnsHostnameBlockOk

`func (o *IpamsvcDHCPInheritance) GetDdnsHostnameBlockOk() (*IpamsvcInheritedDDNSHostnameBlock, bool)`

GetDdnsHostnameBlockOk returns a tuple with the DdnsHostnameBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsHostnameBlock

`func (o *IpamsvcDHCPInheritance) SetDdnsHostnameBlock(v IpamsvcInheritedDDNSHostnameBlock)`

SetDdnsHostnameBlock sets DdnsHostnameBlock field to given value.

### HasDdnsHostnameBlock

`func (o *IpamsvcDHCPInheritance) HasDdnsHostnameBlock() bool`

HasDdnsHostnameBlock returns a boolean if a field has been set.

### GetDdnsTtlPercent

`func (o *IpamsvcDHCPInheritance) GetDdnsTtlPercent() InheritanceInheritedFloat`

GetDdnsTtlPercent returns the DdnsTtlPercent field if non-nil, zero value otherwise.

### GetDdnsTtlPercentOk

`func (o *IpamsvcDHCPInheritance) GetDdnsTtlPercentOk() (*InheritanceInheritedFloat, bool)`

GetDdnsTtlPercentOk returns a tuple with the DdnsTtlPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsTtlPercent

`func (o *IpamsvcDHCPInheritance) SetDdnsTtlPercent(v InheritanceInheritedFloat)`

SetDdnsTtlPercent sets DdnsTtlPercent field to given value.

### HasDdnsTtlPercent

`func (o *IpamsvcDHCPInheritance) HasDdnsTtlPercent() bool`

HasDdnsTtlPercent returns a boolean if a field has been set.

### GetDdnsUpdateBlock

`func (o *IpamsvcDHCPInheritance) GetDdnsUpdateBlock() IpamsvcInheritedDDNSUpdateBlock`

GetDdnsUpdateBlock returns the DdnsUpdateBlock field if non-nil, zero value otherwise.

### GetDdnsUpdateBlockOk

`func (o *IpamsvcDHCPInheritance) GetDdnsUpdateBlockOk() (*IpamsvcInheritedDDNSUpdateBlock, bool)`

GetDdnsUpdateBlockOk returns a tuple with the DdnsUpdateBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUpdateBlock

`func (o *IpamsvcDHCPInheritance) SetDdnsUpdateBlock(v IpamsvcInheritedDDNSUpdateBlock)`

SetDdnsUpdateBlock sets DdnsUpdateBlock field to given value.

### HasDdnsUpdateBlock

`func (o *IpamsvcDHCPInheritance) HasDdnsUpdateBlock() bool`

HasDdnsUpdateBlock returns a boolean if a field has been set.

### GetDdnsUpdateOnRenew

`func (o *IpamsvcDHCPInheritance) GetDdnsUpdateOnRenew() InheritanceInheritedBool`

GetDdnsUpdateOnRenew returns the DdnsUpdateOnRenew field if non-nil, zero value otherwise.

### GetDdnsUpdateOnRenewOk

`func (o *IpamsvcDHCPInheritance) GetDdnsUpdateOnRenewOk() (*InheritanceInheritedBool, bool)`

GetDdnsUpdateOnRenewOk returns a tuple with the DdnsUpdateOnRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUpdateOnRenew

`func (o *IpamsvcDHCPInheritance) SetDdnsUpdateOnRenew(v InheritanceInheritedBool)`

SetDdnsUpdateOnRenew sets DdnsUpdateOnRenew field to given value.

### HasDdnsUpdateOnRenew

`func (o *IpamsvcDHCPInheritance) HasDdnsUpdateOnRenew() bool`

HasDdnsUpdateOnRenew returns a boolean if a field has been set.

### GetDdnsUseConflictResolution

`func (o *IpamsvcDHCPInheritance) GetDdnsUseConflictResolution() InheritanceInheritedBool`

GetDdnsUseConflictResolution returns the DdnsUseConflictResolution field if non-nil, zero value otherwise.

### GetDdnsUseConflictResolutionOk

`func (o *IpamsvcDHCPInheritance) GetDdnsUseConflictResolutionOk() (*InheritanceInheritedBool, bool)`

GetDdnsUseConflictResolutionOk returns a tuple with the DdnsUseConflictResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUseConflictResolution

`func (o *IpamsvcDHCPInheritance) SetDdnsUseConflictResolution(v InheritanceInheritedBool)`

SetDdnsUseConflictResolution sets DdnsUseConflictResolution field to given value.

### HasDdnsUseConflictResolution

`func (o *IpamsvcDHCPInheritance) HasDdnsUseConflictResolution() bool`

HasDdnsUseConflictResolution returns a boolean if a field has been set.

### GetDhcpConfig

`func (o *IpamsvcDHCPInheritance) GetDhcpConfig() IpamsvcInheritedDHCPConfig`

GetDhcpConfig returns the DhcpConfig field if non-nil, zero value otherwise.

### GetDhcpConfigOk

`func (o *IpamsvcDHCPInheritance) GetDhcpConfigOk() (*IpamsvcInheritedDHCPConfig, bool)`

GetDhcpConfigOk returns a tuple with the DhcpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpConfig

`func (o *IpamsvcDHCPInheritance) SetDhcpConfig(v IpamsvcInheritedDHCPConfig)`

SetDhcpConfig sets DhcpConfig field to given value.

### HasDhcpConfig

`func (o *IpamsvcDHCPInheritance) HasDhcpConfig() bool`

HasDhcpConfig returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *IpamsvcDHCPInheritance) GetDhcpOptions() IpamsvcInheritedDHCPOptionList`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *IpamsvcDHCPInheritance) GetDhcpOptionsOk() (*IpamsvcInheritedDHCPOptionList, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *IpamsvcDHCPInheritance) SetDhcpOptions(v IpamsvcInheritedDHCPOptionList)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *IpamsvcDHCPInheritance) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetHeaderOptionFilename

`func (o *IpamsvcDHCPInheritance) GetHeaderOptionFilename() InheritanceInheritedString`

GetHeaderOptionFilename returns the HeaderOptionFilename field if non-nil, zero value otherwise.

### GetHeaderOptionFilenameOk

`func (o *IpamsvcDHCPInheritance) GetHeaderOptionFilenameOk() (*InheritanceInheritedString, bool)`

GetHeaderOptionFilenameOk returns a tuple with the HeaderOptionFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionFilename

`func (o *IpamsvcDHCPInheritance) SetHeaderOptionFilename(v InheritanceInheritedString)`

SetHeaderOptionFilename sets HeaderOptionFilename field to given value.

### HasHeaderOptionFilename

`func (o *IpamsvcDHCPInheritance) HasHeaderOptionFilename() bool`

HasHeaderOptionFilename returns a boolean if a field has been set.

### GetHeaderOptionServerAddress

`func (o *IpamsvcDHCPInheritance) GetHeaderOptionServerAddress() InheritanceInheritedString`

GetHeaderOptionServerAddress returns the HeaderOptionServerAddress field if non-nil, zero value otherwise.

### GetHeaderOptionServerAddressOk

`func (o *IpamsvcDHCPInheritance) GetHeaderOptionServerAddressOk() (*InheritanceInheritedString, bool)`

GetHeaderOptionServerAddressOk returns a tuple with the HeaderOptionServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerAddress

`func (o *IpamsvcDHCPInheritance) SetHeaderOptionServerAddress(v InheritanceInheritedString)`

SetHeaderOptionServerAddress sets HeaderOptionServerAddress field to given value.

### HasHeaderOptionServerAddress

`func (o *IpamsvcDHCPInheritance) HasHeaderOptionServerAddress() bool`

HasHeaderOptionServerAddress returns a boolean if a field has been set.

### GetHeaderOptionServerName

`func (o *IpamsvcDHCPInheritance) GetHeaderOptionServerName() InheritanceInheritedString`

GetHeaderOptionServerName returns the HeaderOptionServerName field if non-nil, zero value otherwise.

### GetHeaderOptionServerNameOk

`func (o *IpamsvcDHCPInheritance) GetHeaderOptionServerNameOk() (*InheritanceInheritedString, bool)`

GetHeaderOptionServerNameOk returns a tuple with the HeaderOptionServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerName

`func (o *IpamsvcDHCPInheritance) SetHeaderOptionServerName(v InheritanceInheritedString)`

SetHeaderOptionServerName sets HeaderOptionServerName field to given value.

### HasHeaderOptionServerName

`func (o *IpamsvcDHCPInheritance) HasHeaderOptionServerName() bool`

HasHeaderOptionServerName returns a boolean if a field has been set.

### GetHostnameRewriteBlock

`func (o *IpamsvcDHCPInheritance) GetHostnameRewriteBlock() IpamsvcInheritedHostnameRewriteBlock`

GetHostnameRewriteBlock returns the HostnameRewriteBlock field if non-nil, zero value otherwise.

### GetHostnameRewriteBlockOk

`func (o *IpamsvcDHCPInheritance) GetHostnameRewriteBlockOk() (*IpamsvcInheritedHostnameRewriteBlock, bool)`

GetHostnameRewriteBlockOk returns a tuple with the HostnameRewriteBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameRewriteBlock

`func (o *IpamsvcDHCPInheritance) SetHostnameRewriteBlock(v IpamsvcInheritedHostnameRewriteBlock)`

SetHostnameRewriteBlock sets HostnameRewriteBlock field to given value.

### HasHostnameRewriteBlock

`func (o *IpamsvcDHCPInheritance) HasHostnameRewriteBlock() bool`

HasHostnameRewriteBlock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


