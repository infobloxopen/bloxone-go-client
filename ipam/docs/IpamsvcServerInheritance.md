# IpamsvcServerInheritance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DdnsBlock** | Pointer to [**IpamsvcInheritedDDNSBlock**](IpamsvcInheritedDDNSBlock.md) |  | [optional] 
**DdnsClientUpdate** | Pointer to [**InheritanceInheritedString**](InheritanceInheritedString.md) |  | [optional] 
**DdnsConflictResolutionMode** | Pointer to [**InheritanceInheritedString**](InheritanceInheritedString.md) |  | [optional] 
**DdnsHostnameBlock** | Pointer to [**IpamsvcInheritedDDNSHostnameBlock**](IpamsvcInheritedDDNSHostnameBlock.md) |  | [optional] 
**DdnsTtlPercent** | Pointer to [**InheritanceInheritedFloat**](InheritanceInheritedFloat.md) |  | [optional] 
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

### NewIpamsvcServerInheritance

`func NewIpamsvcServerInheritance() *IpamsvcServerInheritance`

NewIpamsvcServerInheritance instantiates a new IpamsvcServerInheritance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcServerInheritanceWithDefaults

`func NewIpamsvcServerInheritanceWithDefaults() *IpamsvcServerInheritance`

NewIpamsvcServerInheritanceWithDefaults instantiates a new IpamsvcServerInheritance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDdnsBlock

`func (o *IpamsvcServerInheritance) GetDdnsBlock() IpamsvcInheritedDDNSBlock`

GetDdnsBlock returns the DdnsBlock field if non-nil, zero value otherwise.

### GetDdnsBlockOk

`func (o *IpamsvcServerInheritance) GetDdnsBlockOk() (*IpamsvcInheritedDDNSBlock, bool)`

GetDdnsBlockOk returns a tuple with the DdnsBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsBlock

`func (o *IpamsvcServerInheritance) SetDdnsBlock(v IpamsvcInheritedDDNSBlock)`

SetDdnsBlock sets DdnsBlock field to given value.

### HasDdnsBlock

`func (o *IpamsvcServerInheritance) HasDdnsBlock() bool`

HasDdnsBlock returns a boolean if a field has been set.

### GetDdnsClientUpdate

`func (o *IpamsvcServerInheritance) GetDdnsClientUpdate() InheritanceInheritedString`

GetDdnsClientUpdate returns the DdnsClientUpdate field if non-nil, zero value otherwise.

### GetDdnsClientUpdateOk

`func (o *IpamsvcServerInheritance) GetDdnsClientUpdateOk() (*InheritanceInheritedString, bool)`

GetDdnsClientUpdateOk returns a tuple with the DdnsClientUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsClientUpdate

`func (o *IpamsvcServerInheritance) SetDdnsClientUpdate(v InheritanceInheritedString)`

SetDdnsClientUpdate sets DdnsClientUpdate field to given value.

### HasDdnsClientUpdate

`func (o *IpamsvcServerInheritance) HasDdnsClientUpdate() bool`

HasDdnsClientUpdate returns a boolean if a field has been set.

### GetDdnsConflictResolutionMode

`func (o *IpamsvcServerInheritance) GetDdnsConflictResolutionMode() InheritanceInheritedString`

GetDdnsConflictResolutionMode returns the DdnsConflictResolutionMode field if non-nil, zero value otherwise.

### GetDdnsConflictResolutionModeOk

`func (o *IpamsvcServerInheritance) GetDdnsConflictResolutionModeOk() (*InheritanceInheritedString, bool)`

GetDdnsConflictResolutionModeOk returns a tuple with the DdnsConflictResolutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsConflictResolutionMode

`func (o *IpamsvcServerInheritance) SetDdnsConflictResolutionMode(v InheritanceInheritedString)`

SetDdnsConflictResolutionMode sets DdnsConflictResolutionMode field to given value.

### HasDdnsConflictResolutionMode

`func (o *IpamsvcServerInheritance) HasDdnsConflictResolutionMode() bool`

HasDdnsConflictResolutionMode returns a boolean if a field has been set.

### GetDdnsHostnameBlock

`func (o *IpamsvcServerInheritance) GetDdnsHostnameBlock() IpamsvcInheritedDDNSHostnameBlock`

GetDdnsHostnameBlock returns the DdnsHostnameBlock field if non-nil, zero value otherwise.

### GetDdnsHostnameBlockOk

`func (o *IpamsvcServerInheritance) GetDdnsHostnameBlockOk() (*IpamsvcInheritedDDNSHostnameBlock, bool)`

GetDdnsHostnameBlockOk returns a tuple with the DdnsHostnameBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsHostnameBlock

`func (o *IpamsvcServerInheritance) SetDdnsHostnameBlock(v IpamsvcInheritedDDNSHostnameBlock)`

SetDdnsHostnameBlock sets DdnsHostnameBlock field to given value.

### HasDdnsHostnameBlock

`func (o *IpamsvcServerInheritance) HasDdnsHostnameBlock() bool`

HasDdnsHostnameBlock returns a boolean if a field has been set.

### GetDdnsTtlPercent

`func (o *IpamsvcServerInheritance) GetDdnsTtlPercent() InheritanceInheritedFloat`

GetDdnsTtlPercent returns the DdnsTtlPercent field if non-nil, zero value otherwise.

### GetDdnsTtlPercentOk

`func (o *IpamsvcServerInheritance) GetDdnsTtlPercentOk() (*InheritanceInheritedFloat, bool)`

GetDdnsTtlPercentOk returns a tuple with the DdnsTtlPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsTtlPercent

`func (o *IpamsvcServerInheritance) SetDdnsTtlPercent(v InheritanceInheritedFloat)`

SetDdnsTtlPercent sets DdnsTtlPercent field to given value.

### HasDdnsTtlPercent

`func (o *IpamsvcServerInheritance) HasDdnsTtlPercent() bool`

HasDdnsTtlPercent returns a boolean if a field has been set.

### GetDdnsUpdateOnRenew

`func (o *IpamsvcServerInheritance) GetDdnsUpdateOnRenew() InheritanceInheritedBool`

GetDdnsUpdateOnRenew returns the DdnsUpdateOnRenew field if non-nil, zero value otherwise.

### GetDdnsUpdateOnRenewOk

`func (o *IpamsvcServerInheritance) GetDdnsUpdateOnRenewOk() (*InheritanceInheritedBool, bool)`

GetDdnsUpdateOnRenewOk returns a tuple with the DdnsUpdateOnRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUpdateOnRenew

`func (o *IpamsvcServerInheritance) SetDdnsUpdateOnRenew(v InheritanceInheritedBool)`

SetDdnsUpdateOnRenew sets DdnsUpdateOnRenew field to given value.

### HasDdnsUpdateOnRenew

`func (o *IpamsvcServerInheritance) HasDdnsUpdateOnRenew() bool`

HasDdnsUpdateOnRenew returns a boolean if a field has been set.

### GetDdnsUseConflictResolution

`func (o *IpamsvcServerInheritance) GetDdnsUseConflictResolution() InheritanceInheritedBool`

GetDdnsUseConflictResolution returns the DdnsUseConflictResolution field if non-nil, zero value otherwise.

### GetDdnsUseConflictResolutionOk

`func (o *IpamsvcServerInheritance) GetDdnsUseConflictResolutionOk() (*InheritanceInheritedBool, bool)`

GetDdnsUseConflictResolutionOk returns a tuple with the DdnsUseConflictResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUseConflictResolution

`func (o *IpamsvcServerInheritance) SetDdnsUseConflictResolution(v InheritanceInheritedBool)`

SetDdnsUseConflictResolution sets DdnsUseConflictResolution field to given value.

### HasDdnsUseConflictResolution

`func (o *IpamsvcServerInheritance) HasDdnsUseConflictResolution() bool`

HasDdnsUseConflictResolution returns a boolean if a field has been set.

### GetDhcpConfig

`func (o *IpamsvcServerInheritance) GetDhcpConfig() IpamsvcInheritedDHCPConfig`

GetDhcpConfig returns the DhcpConfig field if non-nil, zero value otherwise.

### GetDhcpConfigOk

`func (o *IpamsvcServerInheritance) GetDhcpConfigOk() (*IpamsvcInheritedDHCPConfig, bool)`

GetDhcpConfigOk returns a tuple with the DhcpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpConfig

`func (o *IpamsvcServerInheritance) SetDhcpConfig(v IpamsvcInheritedDHCPConfig)`

SetDhcpConfig sets DhcpConfig field to given value.

### HasDhcpConfig

`func (o *IpamsvcServerInheritance) HasDhcpConfig() bool`

HasDhcpConfig returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *IpamsvcServerInheritance) GetDhcpOptions() IpamsvcInheritedDHCPOptionList`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *IpamsvcServerInheritance) GetDhcpOptionsOk() (*IpamsvcInheritedDHCPOptionList, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *IpamsvcServerInheritance) SetDhcpOptions(v IpamsvcInheritedDHCPOptionList)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *IpamsvcServerInheritance) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetDhcpOptionsV6

`func (o *IpamsvcServerInheritance) GetDhcpOptionsV6() IpamsvcInheritedDHCPOptionList`

GetDhcpOptionsV6 returns the DhcpOptionsV6 field if non-nil, zero value otherwise.

### GetDhcpOptionsV6Ok

`func (o *IpamsvcServerInheritance) GetDhcpOptionsV6Ok() (*IpamsvcInheritedDHCPOptionList, bool)`

GetDhcpOptionsV6Ok returns a tuple with the DhcpOptionsV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptionsV6

`func (o *IpamsvcServerInheritance) SetDhcpOptionsV6(v IpamsvcInheritedDHCPOptionList)`

SetDhcpOptionsV6 sets DhcpOptionsV6 field to given value.

### HasDhcpOptionsV6

`func (o *IpamsvcServerInheritance) HasDhcpOptionsV6() bool`

HasDhcpOptionsV6 returns a boolean if a field has been set.

### GetHeaderOptionFilename

`func (o *IpamsvcServerInheritance) GetHeaderOptionFilename() InheritanceInheritedString`

GetHeaderOptionFilename returns the HeaderOptionFilename field if non-nil, zero value otherwise.

### GetHeaderOptionFilenameOk

`func (o *IpamsvcServerInheritance) GetHeaderOptionFilenameOk() (*InheritanceInheritedString, bool)`

GetHeaderOptionFilenameOk returns a tuple with the HeaderOptionFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionFilename

`func (o *IpamsvcServerInheritance) SetHeaderOptionFilename(v InheritanceInheritedString)`

SetHeaderOptionFilename sets HeaderOptionFilename field to given value.

### HasHeaderOptionFilename

`func (o *IpamsvcServerInheritance) HasHeaderOptionFilename() bool`

HasHeaderOptionFilename returns a boolean if a field has been set.

### GetHeaderOptionServerAddress

`func (o *IpamsvcServerInheritance) GetHeaderOptionServerAddress() InheritanceInheritedString`

GetHeaderOptionServerAddress returns the HeaderOptionServerAddress field if non-nil, zero value otherwise.

### GetHeaderOptionServerAddressOk

`func (o *IpamsvcServerInheritance) GetHeaderOptionServerAddressOk() (*InheritanceInheritedString, bool)`

GetHeaderOptionServerAddressOk returns a tuple with the HeaderOptionServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerAddress

`func (o *IpamsvcServerInheritance) SetHeaderOptionServerAddress(v InheritanceInheritedString)`

SetHeaderOptionServerAddress sets HeaderOptionServerAddress field to given value.

### HasHeaderOptionServerAddress

`func (o *IpamsvcServerInheritance) HasHeaderOptionServerAddress() bool`

HasHeaderOptionServerAddress returns a boolean if a field has been set.

### GetHeaderOptionServerName

`func (o *IpamsvcServerInheritance) GetHeaderOptionServerName() InheritanceInheritedString`

GetHeaderOptionServerName returns the HeaderOptionServerName field if non-nil, zero value otherwise.

### GetHeaderOptionServerNameOk

`func (o *IpamsvcServerInheritance) GetHeaderOptionServerNameOk() (*InheritanceInheritedString, bool)`

GetHeaderOptionServerNameOk returns a tuple with the HeaderOptionServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerName

`func (o *IpamsvcServerInheritance) SetHeaderOptionServerName(v InheritanceInheritedString)`

SetHeaderOptionServerName sets HeaderOptionServerName field to given value.

### HasHeaderOptionServerName

`func (o *IpamsvcServerInheritance) HasHeaderOptionServerName() bool`

HasHeaderOptionServerName returns a boolean if a field has been set.

### GetHostnameRewriteBlock

`func (o *IpamsvcServerInheritance) GetHostnameRewriteBlock() IpamsvcInheritedHostnameRewriteBlock`

GetHostnameRewriteBlock returns the HostnameRewriteBlock field if non-nil, zero value otherwise.

### GetHostnameRewriteBlockOk

`func (o *IpamsvcServerInheritance) GetHostnameRewriteBlockOk() (*IpamsvcInheritedHostnameRewriteBlock, bool)`

GetHostnameRewriteBlockOk returns a tuple with the HostnameRewriteBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameRewriteBlock

`func (o *IpamsvcServerInheritance) SetHostnameRewriteBlock(v IpamsvcInheritedHostnameRewriteBlock)`

SetHostnameRewriteBlock sets HostnameRewriteBlock field to given value.

### HasHostnameRewriteBlock

`func (o *IpamsvcServerInheritance) HasHostnameRewriteBlock() bool`

HasHostnameRewriteBlock returns a boolean if a field has been set.

### GetVendorSpecificOptionOptionSpace

`func (o *IpamsvcServerInheritance) GetVendorSpecificOptionOptionSpace() InheritanceInheritedIdentifier`

GetVendorSpecificOptionOptionSpace returns the VendorSpecificOptionOptionSpace field if non-nil, zero value otherwise.

### GetVendorSpecificOptionOptionSpaceOk

`func (o *IpamsvcServerInheritance) GetVendorSpecificOptionOptionSpaceOk() (*InheritanceInheritedIdentifier, bool)`

GetVendorSpecificOptionOptionSpaceOk returns a tuple with the VendorSpecificOptionOptionSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSpecificOptionOptionSpace

`func (o *IpamsvcServerInheritance) SetVendorSpecificOptionOptionSpace(v InheritanceInheritedIdentifier)`

SetVendorSpecificOptionOptionSpace sets VendorSpecificOptionOptionSpace field to given value.

### HasVendorSpecificOptionOptionSpace

`func (o *IpamsvcServerInheritance) HasVendorSpecificOptionOptionSpace() bool`

HasVendorSpecificOptionOptionSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


