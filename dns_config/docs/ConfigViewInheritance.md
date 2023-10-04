# ConfigViewInheritance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddEdnsOptionInOutgoingQuery** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**CustomRootNsBlock** | Pointer to [**ConfigInheritedCustomRootNSBlock**](ConfigInheritedCustomRootNSBlock.md) |  | [optional] 
**DnssecValidationBlock** | Pointer to [**ConfigInheritedDNSSECValidationBlock**](ConfigInheritedDNSSECValidationBlock.md) |  | [optional] 
**EcsBlock** | Pointer to [**ConfigInheritedECSBlock**](ConfigInheritedECSBlock.md) |  | [optional] 
**EdnsUdpSize** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**FilterAaaaAcl** | Pointer to [**ConfigInheritedACLItems**](ConfigInheritedACLItems.md) |  | [optional] 
**FilterAaaaOnV4** | Pointer to [**Inheritance2InheritedString**](Inheritance2InheritedString.md) |  | [optional] 
**ForwardersBlock** | Pointer to [**ConfigInheritedForwardersBlock**](ConfigInheritedForwardersBlock.md) |  | [optional] 
**GssTsigEnabled** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**LameTtl** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**MatchRecursiveOnly** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**MaxCacheTtl** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**MaxNegativeTtl** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**MaxUdpSize** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**MinimalResponses** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**Notify** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**QueryAcl** | Pointer to [**ConfigInheritedACLItems**](ConfigInheritedACLItems.md) |  | [optional] 
**RecursionAcl** | Pointer to [**ConfigInheritedACLItems**](ConfigInheritedACLItems.md) |  | [optional] 
**RecursionEnabled** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**SortList** | Pointer to [**ConfigInheritedSortListItems**](ConfigInheritedSortListItems.md) |  | [optional] 
**SynthesizeAddressRecordsFromHttps** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**TransferAcl** | Pointer to [**ConfigInheritedACLItems**](ConfigInheritedACLItems.md) |  | [optional] 
**UpdateAcl** | Pointer to [**ConfigInheritedACLItems**](ConfigInheritedACLItems.md) |  | [optional] 
**UseForwardersForSubzones** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**ZoneAuthority** | Pointer to [**ConfigInheritedZoneAuthority**](ConfigInheritedZoneAuthority.md) |  | [optional] 

## Methods

### NewConfigViewInheritance

`func NewConfigViewInheritance() *ConfigViewInheritance`

NewConfigViewInheritance instantiates a new ConfigViewInheritance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigViewInheritanceWithDefaults

`func NewConfigViewInheritanceWithDefaults() *ConfigViewInheritance`

NewConfigViewInheritanceWithDefaults instantiates a new ConfigViewInheritance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddEdnsOptionInOutgoingQuery

`func (o *ConfigViewInheritance) GetAddEdnsOptionInOutgoingQuery() Inheritance2InheritedBool`

GetAddEdnsOptionInOutgoingQuery returns the AddEdnsOptionInOutgoingQuery field if non-nil, zero value otherwise.

### GetAddEdnsOptionInOutgoingQueryOk

`func (o *ConfigViewInheritance) GetAddEdnsOptionInOutgoingQueryOk() (*Inheritance2InheritedBool, bool)`

GetAddEdnsOptionInOutgoingQueryOk returns a tuple with the AddEdnsOptionInOutgoingQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddEdnsOptionInOutgoingQuery

`func (o *ConfigViewInheritance) SetAddEdnsOptionInOutgoingQuery(v Inheritance2InheritedBool)`

SetAddEdnsOptionInOutgoingQuery sets AddEdnsOptionInOutgoingQuery field to given value.

### HasAddEdnsOptionInOutgoingQuery

`func (o *ConfigViewInheritance) HasAddEdnsOptionInOutgoingQuery() bool`

HasAddEdnsOptionInOutgoingQuery returns a boolean if a field has been set.

### GetCustomRootNsBlock

`func (o *ConfigViewInheritance) GetCustomRootNsBlock() ConfigInheritedCustomRootNSBlock`

GetCustomRootNsBlock returns the CustomRootNsBlock field if non-nil, zero value otherwise.

### GetCustomRootNsBlockOk

`func (o *ConfigViewInheritance) GetCustomRootNsBlockOk() (*ConfigInheritedCustomRootNSBlock, bool)`

GetCustomRootNsBlockOk returns a tuple with the CustomRootNsBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRootNsBlock

`func (o *ConfigViewInheritance) SetCustomRootNsBlock(v ConfigInheritedCustomRootNSBlock)`

SetCustomRootNsBlock sets CustomRootNsBlock field to given value.

### HasCustomRootNsBlock

`func (o *ConfigViewInheritance) HasCustomRootNsBlock() bool`

HasCustomRootNsBlock returns a boolean if a field has been set.

### GetDnssecValidationBlock

`func (o *ConfigViewInheritance) GetDnssecValidationBlock() ConfigInheritedDNSSECValidationBlock`

GetDnssecValidationBlock returns the DnssecValidationBlock field if non-nil, zero value otherwise.

### GetDnssecValidationBlockOk

`func (o *ConfigViewInheritance) GetDnssecValidationBlockOk() (*ConfigInheritedDNSSECValidationBlock, bool)`

GetDnssecValidationBlockOk returns a tuple with the DnssecValidationBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecValidationBlock

`func (o *ConfigViewInheritance) SetDnssecValidationBlock(v ConfigInheritedDNSSECValidationBlock)`

SetDnssecValidationBlock sets DnssecValidationBlock field to given value.

### HasDnssecValidationBlock

`func (o *ConfigViewInheritance) HasDnssecValidationBlock() bool`

HasDnssecValidationBlock returns a boolean if a field has been set.

### GetEcsBlock

`func (o *ConfigViewInheritance) GetEcsBlock() ConfigInheritedECSBlock`

GetEcsBlock returns the EcsBlock field if non-nil, zero value otherwise.

### GetEcsBlockOk

`func (o *ConfigViewInheritance) GetEcsBlockOk() (*ConfigInheritedECSBlock, bool)`

GetEcsBlockOk returns a tuple with the EcsBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsBlock

`func (o *ConfigViewInheritance) SetEcsBlock(v ConfigInheritedECSBlock)`

SetEcsBlock sets EcsBlock field to given value.

### HasEcsBlock

`func (o *ConfigViewInheritance) HasEcsBlock() bool`

HasEcsBlock returns a boolean if a field has been set.

### GetEdnsUdpSize

`func (o *ConfigViewInheritance) GetEdnsUdpSize() Inheritance2InheritedUInt32`

GetEdnsUdpSize returns the EdnsUdpSize field if non-nil, zero value otherwise.

### GetEdnsUdpSizeOk

`func (o *ConfigViewInheritance) GetEdnsUdpSizeOk() (*Inheritance2InheritedUInt32, bool)`

GetEdnsUdpSizeOk returns a tuple with the EdnsUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdnsUdpSize

`func (o *ConfigViewInheritance) SetEdnsUdpSize(v Inheritance2InheritedUInt32)`

SetEdnsUdpSize sets EdnsUdpSize field to given value.

### HasEdnsUdpSize

`func (o *ConfigViewInheritance) HasEdnsUdpSize() bool`

HasEdnsUdpSize returns a boolean if a field has been set.

### GetFilterAaaaAcl

`func (o *ConfigViewInheritance) GetFilterAaaaAcl() ConfigInheritedACLItems`

GetFilterAaaaAcl returns the FilterAaaaAcl field if non-nil, zero value otherwise.

### GetFilterAaaaAclOk

`func (o *ConfigViewInheritance) GetFilterAaaaAclOk() (*ConfigInheritedACLItems, bool)`

GetFilterAaaaAclOk returns a tuple with the FilterAaaaAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaaAcl

`func (o *ConfigViewInheritance) SetFilterAaaaAcl(v ConfigInheritedACLItems)`

SetFilterAaaaAcl sets FilterAaaaAcl field to given value.

### HasFilterAaaaAcl

`func (o *ConfigViewInheritance) HasFilterAaaaAcl() bool`

HasFilterAaaaAcl returns a boolean if a field has been set.

### GetFilterAaaaOnV4

`func (o *ConfigViewInheritance) GetFilterAaaaOnV4() Inheritance2InheritedString`

GetFilterAaaaOnV4 returns the FilterAaaaOnV4 field if non-nil, zero value otherwise.

### GetFilterAaaaOnV4Ok

`func (o *ConfigViewInheritance) GetFilterAaaaOnV4Ok() (*Inheritance2InheritedString, bool)`

GetFilterAaaaOnV4Ok returns a tuple with the FilterAaaaOnV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaaOnV4

`func (o *ConfigViewInheritance) SetFilterAaaaOnV4(v Inheritance2InheritedString)`

SetFilterAaaaOnV4 sets FilterAaaaOnV4 field to given value.

### HasFilterAaaaOnV4

`func (o *ConfigViewInheritance) HasFilterAaaaOnV4() bool`

HasFilterAaaaOnV4 returns a boolean if a field has been set.

### GetForwardersBlock

`func (o *ConfigViewInheritance) GetForwardersBlock() ConfigInheritedForwardersBlock`

GetForwardersBlock returns the ForwardersBlock field if non-nil, zero value otherwise.

### GetForwardersBlockOk

`func (o *ConfigViewInheritance) GetForwardersBlockOk() (*ConfigInheritedForwardersBlock, bool)`

GetForwardersBlockOk returns a tuple with the ForwardersBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardersBlock

`func (o *ConfigViewInheritance) SetForwardersBlock(v ConfigInheritedForwardersBlock)`

SetForwardersBlock sets ForwardersBlock field to given value.

### HasForwardersBlock

`func (o *ConfigViewInheritance) HasForwardersBlock() bool`

HasForwardersBlock returns a boolean if a field has been set.

### GetGssTsigEnabled

`func (o *ConfigViewInheritance) GetGssTsigEnabled() Inheritance2InheritedBool`

GetGssTsigEnabled returns the GssTsigEnabled field if non-nil, zero value otherwise.

### GetGssTsigEnabledOk

`func (o *ConfigViewInheritance) GetGssTsigEnabledOk() (*Inheritance2InheritedBool, bool)`

GetGssTsigEnabledOk returns a tuple with the GssTsigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigEnabled

`func (o *ConfigViewInheritance) SetGssTsigEnabled(v Inheritance2InheritedBool)`

SetGssTsigEnabled sets GssTsigEnabled field to given value.

### HasGssTsigEnabled

`func (o *ConfigViewInheritance) HasGssTsigEnabled() bool`

HasGssTsigEnabled returns a boolean if a field has been set.

### GetLameTtl

`func (o *ConfigViewInheritance) GetLameTtl() Inheritance2InheritedUInt32`

GetLameTtl returns the LameTtl field if non-nil, zero value otherwise.

### GetLameTtlOk

`func (o *ConfigViewInheritance) GetLameTtlOk() (*Inheritance2InheritedUInt32, bool)`

GetLameTtlOk returns a tuple with the LameTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLameTtl

`func (o *ConfigViewInheritance) SetLameTtl(v Inheritance2InheritedUInt32)`

SetLameTtl sets LameTtl field to given value.

### HasLameTtl

`func (o *ConfigViewInheritance) HasLameTtl() bool`

HasLameTtl returns a boolean if a field has been set.

### GetMatchRecursiveOnly

`func (o *ConfigViewInheritance) GetMatchRecursiveOnly() Inheritance2InheritedBool`

GetMatchRecursiveOnly returns the MatchRecursiveOnly field if non-nil, zero value otherwise.

### GetMatchRecursiveOnlyOk

`func (o *ConfigViewInheritance) GetMatchRecursiveOnlyOk() (*Inheritance2InheritedBool, bool)`

GetMatchRecursiveOnlyOk returns a tuple with the MatchRecursiveOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchRecursiveOnly

`func (o *ConfigViewInheritance) SetMatchRecursiveOnly(v Inheritance2InheritedBool)`

SetMatchRecursiveOnly sets MatchRecursiveOnly field to given value.

### HasMatchRecursiveOnly

`func (o *ConfigViewInheritance) HasMatchRecursiveOnly() bool`

HasMatchRecursiveOnly returns a boolean if a field has been set.

### GetMaxCacheTtl

`func (o *ConfigViewInheritance) GetMaxCacheTtl() Inheritance2InheritedUInt32`

GetMaxCacheTtl returns the MaxCacheTtl field if non-nil, zero value otherwise.

### GetMaxCacheTtlOk

`func (o *ConfigViewInheritance) GetMaxCacheTtlOk() (*Inheritance2InheritedUInt32, bool)`

GetMaxCacheTtlOk returns a tuple with the MaxCacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheTtl

`func (o *ConfigViewInheritance) SetMaxCacheTtl(v Inheritance2InheritedUInt32)`

SetMaxCacheTtl sets MaxCacheTtl field to given value.

### HasMaxCacheTtl

`func (o *ConfigViewInheritance) HasMaxCacheTtl() bool`

HasMaxCacheTtl returns a boolean if a field has been set.

### GetMaxNegativeTtl

`func (o *ConfigViewInheritance) GetMaxNegativeTtl() Inheritance2InheritedUInt32`

GetMaxNegativeTtl returns the MaxNegativeTtl field if non-nil, zero value otherwise.

### GetMaxNegativeTtlOk

`func (o *ConfigViewInheritance) GetMaxNegativeTtlOk() (*Inheritance2InheritedUInt32, bool)`

GetMaxNegativeTtlOk returns a tuple with the MaxNegativeTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNegativeTtl

`func (o *ConfigViewInheritance) SetMaxNegativeTtl(v Inheritance2InheritedUInt32)`

SetMaxNegativeTtl sets MaxNegativeTtl field to given value.

### HasMaxNegativeTtl

`func (o *ConfigViewInheritance) HasMaxNegativeTtl() bool`

HasMaxNegativeTtl returns a boolean if a field has been set.

### GetMaxUdpSize

`func (o *ConfigViewInheritance) GetMaxUdpSize() Inheritance2InheritedUInt32`

GetMaxUdpSize returns the MaxUdpSize field if non-nil, zero value otherwise.

### GetMaxUdpSizeOk

`func (o *ConfigViewInheritance) GetMaxUdpSizeOk() (*Inheritance2InheritedUInt32, bool)`

GetMaxUdpSizeOk returns a tuple with the MaxUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUdpSize

`func (o *ConfigViewInheritance) SetMaxUdpSize(v Inheritance2InheritedUInt32)`

SetMaxUdpSize sets MaxUdpSize field to given value.

### HasMaxUdpSize

`func (o *ConfigViewInheritance) HasMaxUdpSize() bool`

HasMaxUdpSize returns a boolean if a field has been set.

### GetMinimalResponses

`func (o *ConfigViewInheritance) GetMinimalResponses() Inheritance2InheritedBool`

GetMinimalResponses returns the MinimalResponses field if non-nil, zero value otherwise.

### GetMinimalResponsesOk

`func (o *ConfigViewInheritance) GetMinimalResponsesOk() (*Inheritance2InheritedBool, bool)`

GetMinimalResponsesOk returns a tuple with the MinimalResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimalResponses

`func (o *ConfigViewInheritance) SetMinimalResponses(v Inheritance2InheritedBool)`

SetMinimalResponses sets MinimalResponses field to given value.

### HasMinimalResponses

`func (o *ConfigViewInheritance) HasMinimalResponses() bool`

HasMinimalResponses returns a boolean if a field has been set.

### GetNotify

`func (o *ConfigViewInheritance) GetNotify() Inheritance2InheritedBool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *ConfigViewInheritance) GetNotifyOk() (*Inheritance2InheritedBool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *ConfigViewInheritance) SetNotify(v Inheritance2InheritedBool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *ConfigViewInheritance) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetQueryAcl

`func (o *ConfigViewInheritance) GetQueryAcl() ConfigInheritedACLItems`

GetQueryAcl returns the QueryAcl field if non-nil, zero value otherwise.

### GetQueryAclOk

`func (o *ConfigViewInheritance) GetQueryAclOk() (*ConfigInheritedACLItems, bool)`

GetQueryAclOk returns a tuple with the QueryAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryAcl

`func (o *ConfigViewInheritance) SetQueryAcl(v ConfigInheritedACLItems)`

SetQueryAcl sets QueryAcl field to given value.

### HasQueryAcl

`func (o *ConfigViewInheritance) HasQueryAcl() bool`

HasQueryAcl returns a boolean if a field has been set.

### GetRecursionAcl

`func (o *ConfigViewInheritance) GetRecursionAcl() ConfigInheritedACLItems`

GetRecursionAcl returns the RecursionAcl field if non-nil, zero value otherwise.

### GetRecursionAclOk

`func (o *ConfigViewInheritance) GetRecursionAclOk() (*ConfigInheritedACLItems, bool)`

GetRecursionAclOk returns a tuple with the RecursionAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursionAcl

`func (o *ConfigViewInheritance) SetRecursionAcl(v ConfigInheritedACLItems)`

SetRecursionAcl sets RecursionAcl field to given value.

### HasRecursionAcl

`func (o *ConfigViewInheritance) HasRecursionAcl() bool`

HasRecursionAcl returns a boolean if a field has been set.

### GetRecursionEnabled

`func (o *ConfigViewInheritance) GetRecursionEnabled() Inheritance2InheritedBool`

GetRecursionEnabled returns the RecursionEnabled field if non-nil, zero value otherwise.

### GetRecursionEnabledOk

`func (o *ConfigViewInheritance) GetRecursionEnabledOk() (*Inheritance2InheritedBool, bool)`

GetRecursionEnabledOk returns a tuple with the RecursionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursionEnabled

`func (o *ConfigViewInheritance) SetRecursionEnabled(v Inheritance2InheritedBool)`

SetRecursionEnabled sets RecursionEnabled field to given value.

### HasRecursionEnabled

`func (o *ConfigViewInheritance) HasRecursionEnabled() bool`

HasRecursionEnabled returns a boolean if a field has been set.

### GetSortList

`func (o *ConfigViewInheritance) GetSortList() ConfigInheritedSortListItems`

GetSortList returns the SortList field if non-nil, zero value otherwise.

### GetSortListOk

`func (o *ConfigViewInheritance) GetSortListOk() (*ConfigInheritedSortListItems, bool)`

GetSortListOk returns a tuple with the SortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortList

`func (o *ConfigViewInheritance) SetSortList(v ConfigInheritedSortListItems)`

SetSortList sets SortList field to given value.

### HasSortList

`func (o *ConfigViewInheritance) HasSortList() bool`

HasSortList returns a boolean if a field has been set.

### GetSynthesizeAddressRecordsFromHttps

`func (o *ConfigViewInheritance) GetSynthesizeAddressRecordsFromHttps() Inheritance2InheritedBool`

GetSynthesizeAddressRecordsFromHttps returns the SynthesizeAddressRecordsFromHttps field if non-nil, zero value otherwise.

### GetSynthesizeAddressRecordsFromHttpsOk

`func (o *ConfigViewInheritance) GetSynthesizeAddressRecordsFromHttpsOk() (*Inheritance2InheritedBool, bool)`

GetSynthesizeAddressRecordsFromHttpsOk returns a tuple with the SynthesizeAddressRecordsFromHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthesizeAddressRecordsFromHttps

`func (o *ConfigViewInheritance) SetSynthesizeAddressRecordsFromHttps(v Inheritance2InheritedBool)`

SetSynthesizeAddressRecordsFromHttps sets SynthesizeAddressRecordsFromHttps field to given value.

### HasSynthesizeAddressRecordsFromHttps

`func (o *ConfigViewInheritance) HasSynthesizeAddressRecordsFromHttps() bool`

HasSynthesizeAddressRecordsFromHttps returns a boolean if a field has been set.

### GetTransferAcl

`func (o *ConfigViewInheritance) GetTransferAcl() ConfigInheritedACLItems`

GetTransferAcl returns the TransferAcl field if non-nil, zero value otherwise.

### GetTransferAclOk

`func (o *ConfigViewInheritance) GetTransferAclOk() (*ConfigInheritedACLItems, bool)`

GetTransferAclOk returns a tuple with the TransferAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAcl

`func (o *ConfigViewInheritance) SetTransferAcl(v ConfigInheritedACLItems)`

SetTransferAcl sets TransferAcl field to given value.

### HasTransferAcl

`func (o *ConfigViewInheritance) HasTransferAcl() bool`

HasTransferAcl returns a boolean if a field has been set.

### GetUpdateAcl

`func (o *ConfigViewInheritance) GetUpdateAcl() ConfigInheritedACLItems`

GetUpdateAcl returns the UpdateAcl field if non-nil, zero value otherwise.

### GetUpdateAclOk

`func (o *ConfigViewInheritance) GetUpdateAclOk() (*ConfigInheritedACLItems, bool)`

GetUpdateAclOk returns a tuple with the UpdateAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAcl

`func (o *ConfigViewInheritance) SetUpdateAcl(v ConfigInheritedACLItems)`

SetUpdateAcl sets UpdateAcl field to given value.

### HasUpdateAcl

`func (o *ConfigViewInheritance) HasUpdateAcl() bool`

HasUpdateAcl returns a boolean if a field has been set.

### GetUseForwardersForSubzones

`func (o *ConfigViewInheritance) GetUseForwardersForSubzones() Inheritance2InheritedBool`

GetUseForwardersForSubzones returns the UseForwardersForSubzones field if non-nil, zero value otherwise.

### GetUseForwardersForSubzonesOk

`func (o *ConfigViewInheritance) GetUseForwardersForSubzonesOk() (*Inheritance2InheritedBool, bool)`

GetUseForwardersForSubzonesOk returns a tuple with the UseForwardersForSubzones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForwardersForSubzones

`func (o *ConfigViewInheritance) SetUseForwardersForSubzones(v Inheritance2InheritedBool)`

SetUseForwardersForSubzones sets UseForwardersForSubzones field to given value.

### HasUseForwardersForSubzones

`func (o *ConfigViewInheritance) HasUseForwardersForSubzones() bool`

HasUseForwardersForSubzones returns a boolean if a field has been set.

### GetZoneAuthority

`func (o *ConfigViewInheritance) GetZoneAuthority() ConfigInheritedZoneAuthority`

GetZoneAuthority returns the ZoneAuthority field if non-nil, zero value otherwise.

### GetZoneAuthorityOk

`func (o *ConfigViewInheritance) GetZoneAuthorityOk() (*ConfigInheritedZoneAuthority, bool)`

GetZoneAuthorityOk returns a tuple with the ZoneAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAuthority

`func (o *ConfigViewInheritance) SetZoneAuthority(v ConfigInheritedZoneAuthority)`

SetZoneAuthority sets ZoneAuthority field to given value.

### HasZoneAuthority

`func (o *ConfigViewInheritance) HasZoneAuthority() bool`

HasZoneAuthority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


