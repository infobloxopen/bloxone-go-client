# ServerInheritance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DdnsBlock** | Pointer to [**InheritedDDNSBlock**](InheritedDDNSBlock.md) |  | [optional] 
**DdnsClientUpdate** | Pointer to [**InheritanceInheritedString**](InheritanceInheritedString.md) |  | [optional] 
**DdnsConflictResolutionMode** | Pointer to [**InheritanceInheritedString**](InheritanceInheritedString.md) |  | [optional] 
**DdnsHostnameBlock** | Pointer to [**InheritedDDNSHostnameBlock**](InheritedDDNSHostnameBlock.md) |  | [optional] 
**DdnsTtlPercent** | Pointer to [**InheritanceInheritedFloat**](InheritanceInheritedFloat.md) |  | [optional] 
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

### NewServerInheritance

`func NewServerInheritance() *ServerInheritance`

NewServerInheritance instantiates a new ServerInheritance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInheritanceWithDefaults

`func NewServerInheritanceWithDefaults() *ServerInheritance`

NewServerInheritanceWithDefaults instantiates a new ServerInheritance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDdnsBlock

`func (o *ServerInheritance) GetDdnsBlock() InheritedDDNSBlock`

GetDdnsBlock returns the DdnsBlock field if non-nil, zero value otherwise.

### GetDdnsBlockOk

`func (o *ServerInheritance) GetDdnsBlockOk() (*InheritedDDNSBlock, bool)`

GetDdnsBlockOk returns a tuple with the DdnsBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsBlock

`func (o *ServerInheritance) SetDdnsBlock(v InheritedDDNSBlock)`

SetDdnsBlock sets DdnsBlock field to given value.

### HasDdnsBlock

`func (o *ServerInheritance) HasDdnsBlock() bool`

HasDdnsBlock returns a boolean if a field has been set.

### GetDdnsClientUpdate

`func (o *ServerInheritance) GetDdnsClientUpdate() InheritanceInheritedString`

GetDdnsClientUpdate returns the DdnsClientUpdate field if non-nil, zero value otherwise.

### GetDdnsClientUpdateOk

`func (o *ServerInheritance) GetDdnsClientUpdateOk() (*InheritanceInheritedString, bool)`

GetDdnsClientUpdateOk returns a tuple with the DdnsClientUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsClientUpdate

`func (o *ServerInheritance) SetDdnsClientUpdate(v InheritanceInheritedString)`

SetDdnsClientUpdate sets DdnsClientUpdate field to given value.

### HasDdnsClientUpdate

`func (o *ServerInheritance) HasDdnsClientUpdate() bool`

HasDdnsClientUpdate returns a boolean if a field has been set.

### GetDdnsConflictResolutionMode

`func (o *ServerInheritance) GetDdnsConflictResolutionMode() InheritanceInheritedString`

GetDdnsConflictResolutionMode returns the DdnsConflictResolutionMode field if non-nil, zero value otherwise.

### GetDdnsConflictResolutionModeOk

`func (o *ServerInheritance) GetDdnsConflictResolutionModeOk() (*InheritanceInheritedString, bool)`

GetDdnsConflictResolutionModeOk returns a tuple with the DdnsConflictResolutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsConflictResolutionMode

`func (o *ServerInheritance) SetDdnsConflictResolutionMode(v InheritanceInheritedString)`

SetDdnsConflictResolutionMode sets DdnsConflictResolutionMode field to given value.

### HasDdnsConflictResolutionMode

`func (o *ServerInheritance) HasDdnsConflictResolutionMode() bool`

HasDdnsConflictResolutionMode returns a boolean if a field has been set.

### GetDdnsHostnameBlock

`func (o *ServerInheritance) GetDdnsHostnameBlock() InheritedDDNSHostnameBlock`

GetDdnsHostnameBlock returns the DdnsHostnameBlock field if non-nil, zero value otherwise.

### GetDdnsHostnameBlockOk

`func (o *ServerInheritance) GetDdnsHostnameBlockOk() (*InheritedDDNSHostnameBlock, bool)`

GetDdnsHostnameBlockOk returns a tuple with the DdnsHostnameBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsHostnameBlock

`func (o *ServerInheritance) SetDdnsHostnameBlock(v InheritedDDNSHostnameBlock)`

SetDdnsHostnameBlock sets DdnsHostnameBlock field to given value.

### HasDdnsHostnameBlock

`func (o *ServerInheritance) HasDdnsHostnameBlock() bool`

HasDdnsHostnameBlock returns a boolean if a field has been set.

### GetDdnsTtlPercent

`func (o *ServerInheritance) GetDdnsTtlPercent() InheritanceInheritedFloat`

GetDdnsTtlPercent returns the DdnsTtlPercent field if non-nil, zero value otherwise.

### GetDdnsTtlPercentOk

`func (o *ServerInheritance) GetDdnsTtlPercentOk() (*InheritanceInheritedFloat, bool)`

GetDdnsTtlPercentOk returns a tuple with the DdnsTtlPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsTtlPercent

`func (o *ServerInheritance) SetDdnsTtlPercent(v InheritanceInheritedFloat)`

SetDdnsTtlPercent sets DdnsTtlPercent field to given value.

### HasDdnsTtlPercent

`func (o *ServerInheritance) HasDdnsTtlPercent() bool`

HasDdnsTtlPercent returns a boolean if a field has been set.

### GetDdnsUpdateOnRenew

`func (o *ServerInheritance) GetDdnsUpdateOnRenew() InheritanceInheritedBool`

GetDdnsUpdateOnRenew returns the DdnsUpdateOnRenew field if non-nil, zero value otherwise.

### GetDdnsUpdateOnRenewOk

`func (o *ServerInheritance) GetDdnsUpdateOnRenewOk() (*InheritanceInheritedBool, bool)`

GetDdnsUpdateOnRenewOk returns a tuple with the DdnsUpdateOnRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUpdateOnRenew

`func (o *ServerInheritance) SetDdnsUpdateOnRenew(v InheritanceInheritedBool)`

SetDdnsUpdateOnRenew sets DdnsUpdateOnRenew field to given value.

### HasDdnsUpdateOnRenew

`func (o *ServerInheritance) HasDdnsUpdateOnRenew() bool`

HasDdnsUpdateOnRenew returns a boolean if a field has been set.

### GetDdnsUseConflictResolution

`func (o *ServerInheritance) GetDdnsUseConflictResolution() InheritanceInheritedBool`

GetDdnsUseConflictResolution returns the DdnsUseConflictResolution field if non-nil, zero value otherwise.

### GetDdnsUseConflictResolutionOk

`func (o *ServerInheritance) GetDdnsUseConflictResolutionOk() (*InheritanceInheritedBool, bool)`

GetDdnsUseConflictResolutionOk returns a tuple with the DdnsUseConflictResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsUseConflictResolution

`func (o *ServerInheritance) SetDdnsUseConflictResolution(v InheritanceInheritedBool)`

SetDdnsUseConflictResolution sets DdnsUseConflictResolution field to given value.

### HasDdnsUseConflictResolution

`func (o *ServerInheritance) HasDdnsUseConflictResolution() bool`

HasDdnsUseConflictResolution returns a boolean if a field has been set.

### GetDhcpConfig

`func (o *ServerInheritance) GetDhcpConfig() InheritedDHCPConfig`

GetDhcpConfig returns the DhcpConfig field if non-nil, zero value otherwise.

### GetDhcpConfigOk

`func (o *ServerInheritance) GetDhcpConfigOk() (*InheritedDHCPConfig, bool)`

GetDhcpConfigOk returns a tuple with the DhcpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpConfig

`func (o *ServerInheritance) SetDhcpConfig(v InheritedDHCPConfig)`

SetDhcpConfig sets DhcpConfig field to given value.

### HasDhcpConfig

`func (o *ServerInheritance) HasDhcpConfig() bool`

HasDhcpConfig returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *ServerInheritance) GetDhcpOptions() InheritedDHCPOptionList`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *ServerInheritance) GetDhcpOptionsOk() (*InheritedDHCPOptionList, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *ServerInheritance) SetDhcpOptions(v InheritedDHCPOptionList)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *ServerInheritance) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetDhcpOptionsV6

`func (o *ServerInheritance) GetDhcpOptionsV6() InheritedDHCPOptionList`

GetDhcpOptionsV6 returns the DhcpOptionsV6 field if non-nil, zero value otherwise.

### GetDhcpOptionsV6Ok

`func (o *ServerInheritance) GetDhcpOptionsV6Ok() (*InheritedDHCPOptionList, bool)`

GetDhcpOptionsV6Ok returns a tuple with the DhcpOptionsV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptionsV6

`func (o *ServerInheritance) SetDhcpOptionsV6(v InheritedDHCPOptionList)`

SetDhcpOptionsV6 sets DhcpOptionsV6 field to given value.

### HasDhcpOptionsV6

`func (o *ServerInheritance) HasDhcpOptionsV6() bool`

HasDhcpOptionsV6 returns a boolean if a field has been set.

### GetHeaderOptionFilename

`func (o *ServerInheritance) GetHeaderOptionFilename() InheritanceInheritedString`

GetHeaderOptionFilename returns the HeaderOptionFilename field if non-nil, zero value otherwise.

### GetHeaderOptionFilenameOk

`func (o *ServerInheritance) GetHeaderOptionFilenameOk() (*InheritanceInheritedString, bool)`

GetHeaderOptionFilenameOk returns a tuple with the HeaderOptionFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionFilename

`func (o *ServerInheritance) SetHeaderOptionFilename(v InheritanceInheritedString)`

SetHeaderOptionFilename sets HeaderOptionFilename field to given value.

### HasHeaderOptionFilename

`func (o *ServerInheritance) HasHeaderOptionFilename() bool`

HasHeaderOptionFilename returns a boolean if a field has been set.

### GetHeaderOptionServerAddress

`func (o *ServerInheritance) GetHeaderOptionServerAddress() InheritanceInheritedString`

GetHeaderOptionServerAddress returns the HeaderOptionServerAddress field if non-nil, zero value otherwise.

### GetHeaderOptionServerAddressOk

`func (o *ServerInheritance) GetHeaderOptionServerAddressOk() (*InheritanceInheritedString, bool)`

GetHeaderOptionServerAddressOk returns a tuple with the HeaderOptionServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerAddress

`func (o *ServerInheritance) SetHeaderOptionServerAddress(v InheritanceInheritedString)`

SetHeaderOptionServerAddress sets HeaderOptionServerAddress field to given value.

### HasHeaderOptionServerAddress

`func (o *ServerInheritance) HasHeaderOptionServerAddress() bool`

HasHeaderOptionServerAddress returns a boolean if a field has been set.

### GetHeaderOptionServerName

`func (o *ServerInheritance) GetHeaderOptionServerName() InheritanceInheritedString`

GetHeaderOptionServerName returns the HeaderOptionServerName field if non-nil, zero value otherwise.

### GetHeaderOptionServerNameOk

`func (o *ServerInheritance) GetHeaderOptionServerNameOk() (*InheritanceInheritedString, bool)`

GetHeaderOptionServerNameOk returns a tuple with the HeaderOptionServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerName

`func (o *ServerInheritance) SetHeaderOptionServerName(v InheritanceInheritedString)`

SetHeaderOptionServerName sets HeaderOptionServerName field to given value.

### HasHeaderOptionServerName

`func (o *ServerInheritance) HasHeaderOptionServerName() bool`

HasHeaderOptionServerName returns a boolean if a field has been set.

### GetHostnameRewriteBlock

`func (o *ServerInheritance) GetHostnameRewriteBlock() InheritedHostnameRewriteBlock`

GetHostnameRewriteBlock returns the HostnameRewriteBlock field if non-nil, zero value otherwise.

### GetHostnameRewriteBlockOk

`func (o *ServerInheritance) GetHostnameRewriteBlockOk() (*InheritedHostnameRewriteBlock, bool)`

GetHostnameRewriteBlockOk returns a tuple with the HostnameRewriteBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameRewriteBlock

`func (o *ServerInheritance) SetHostnameRewriteBlock(v InheritedHostnameRewriteBlock)`

SetHostnameRewriteBlock sets HostnameRewriteBlock field to given value.

### HasHostnameRewriteBlock

`func (o *ServerInheritance) HasHostnameRewriteBlock() bool`

HasHostnameRewriteBlock returns a boolean if a field has been set.

### GetVendorSpecificOptionOptionSpace

`func (o *ServerInheritance) GetVendorSpecificOptionOptionSpace() InheritanceInheritedIdentifier`

GetVendorSpecificOptionOptionSpace returns the VendorSpecificOptionOptionSpace field if non-nil, zero value otherwise.

### GetVendorSpecificOptionOptionSpaceOk

`func (o *ServerInheritance) GetVendorSpecificOptionOptionSpaceOk() (*InheritanceInheritedIdentifier, bool)`

GetVendorSpecificOptionOptionSpaceOk returns a tuple with the VendorSpecificOptionOptionSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSpecificOptionOptionSpace

`func (o *ServerInheritance) SetVendorSpecificOptionOptionSpace(v InheritanceInheritedIdentifier)`

SetVendorSpecificOptionOptionSpace sets VendorSpecificOptionOptionSpace field to given value.

### HasVendorSpecificOptionOptionSpace

`func (o *ServerInheritance) HasVendorSpecificOptionOptionSpace() bool`

HasVendorSpecificOptionOptionSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


