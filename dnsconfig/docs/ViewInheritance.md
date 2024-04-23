# ViewInheritance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddEdnsOptionInOutgoingQuery** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**CustomRootNsBlock** | Pointer to [**InheritedCustomRootNSBlock**](InheritedCustomRootNSBlock.md) |  | [optional] 
**DnssecValidationBlock** | Pointer to [**InheritedDNSSECValidationBlock**](InheritedDNSSECValidationBlock.md) |  | [optional] 
**DtcConfig** | Pointer to [**InheritedDtcConfig**](InheritedDtcConfig.md) |  | [optional] 
**EcsBlock** | Pointer to [**InheritedECSBlock**](InheritedECSBlock.md) |  | [optional] 
**EdnsUdpSize** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**FilterAaaaAcl** | Pointer to [**InheritedACLItems**](InheritedACLItems.md) |  | [optional] 
**FilterAaaaOnV4** | Pointer to [**Inheritance2InheritedString**](Inheritance2InheritedString.md) |  | [optional] 
**ForwardersBlock** | Pointer to [**InheritedForwardersBlock**](InheritedForwardersBlock.md) |  | [optional] 
**GssTsigEnabled** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**LameTtl** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**MatchRecursiveOnly** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**MaxCacheTtl** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**MaxNegativeTtl** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**MaxUdpSize** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**MinimalResponses** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**Notify** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**QueryAcl** | Pointer to [**InheritedACLItems**](InheritedACLItems.md) |  | [optional] 
**RecursionAcl** | Pointer to [**InheritedACLItems**](InheritedACLItems.md) |  | [optional] 
**RecursionEnabled** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**SortList** | Pointer to [**InheritedSortListItems**](InheritedSortListItems.md) |  | [optional] 
**SynthesizeAddressRecordsFromHttps** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**TransferAcl** | Pointer to [**InheritedACLItems**](InheritedACLItems.md) |  | [optional] 
**UpdateAcl** | Pointer to [**InheritedACLItems**](InheritedACLItems.md) |  | [optional] 
**UseForwardersForSubzones** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**ZoneAuthority** | Pointer to [**InheritedZoneAuthority**](InheritedZoneAuthority.md) |  | [optional] 

## Methods

### NewViewInheritance

`func NewViewInheritance() *ViewInheritance`

NewViewInheritance instantiates a new ViewInheritance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewInheritanceWithDefaults

`func NewViewInheritanceWithDefaults() *ViewInheritance`

NewViewInheritanceWithDefaults instantiates a new ViewInheritance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddEdnsOptionInOutgoingQuery

`func (o *ViewInheritance) GetAddEdnsOptionInOutgoingQuery() Inheritance2InheritedBool`

GetAddEdnsOptionInOutgoingQuery returns the AddEdnsOptionInOutgoingQuery field if non-nil, zero value otherwise.

### GetAddEdnsOptionInOutgoingQueryOk

`func (o *ViewInheritance) GetAddEdnsOptionInOutgoingQueryOk() (*Inheritance2InheritedBool, bool)`

GetAddEdnsOptionInOutgoingQueryOk returns a tuple with the AddEdnsOptionInOutgoingQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddEdnsOptionInOutgoingQuery

`func (o *ViewInheritance) SetAddEdnsOptionInOutgoingQuery(v Inheritance2InheritedBool)`

SetAddEdnsOptionInOutgoingQuery sets AddEdnsOptionInOutgoingQuery field to given value.

### HasAddEdnsOptionInOutgoingQuery

`func (o *ViewInheritance) HasAddEdnsOptionInOutgoingQuery() bool`

HasAddEdnsOptionInOutgoingQuery returns a boolean if a field has been set.

### GetCustomRootNsBlock

`func (o *ViewInheritance) GetCustomRootNsBlock() InheritedCustomRootNSBlock`

GetCustomRootNsBlock returns the CustomRootNsBlock field if non-nil, zero value otherwise.

### GetCustomRootNsBlockOk

`func (o *ViewInheritance) GetCustomRootNsBlockOk() (*InheritedCustomRootNSBlock, bool)`

GetCustomRootNsBlockOk returns a tuple with the CustomRootNsBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRootNsBlock

`func (o *ViewInheritance) SetCustomRootNsBlock(v InheritedCustomRootNSBlock)`

SetCustomRootNsBlock sets CustomRootNsBlock field to given value.

### HasCustomRootNsBlock

`func (o *ViewInheritance) HasCustomRootNsBlock() bool`

HasCustomRootNsBlock returns a boolean if a field has been set.

### GetDnssecValidationBlock

`func (o *ViewInheritance) GetDnssecValidationBlock() InheritedDNSSECValidationBlock`

GetDnssecValidationBlock returns the DnssecValidationBlock field if non-nil, zero value otherwise.

### GetDnssecValidationBlockOk

`func (o *ViewInheritance) GetDnssecValidationBlockOk() (*InheritedDNSSECValidationBlock, bool)`

GetDnssecValidationBlockOk returns a tuple with the DnssecValidationBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecValidationBlock

`func (o *ViewInheritance) SetDnssecValidationBlock(v InheritedDNSSECValidationBlock)`

SetDnssecValidationBlock sets DnssecValidationBlock field to given value.

### HasDnssecValidationBlock

`func (o *ViewInheritance) HasDnssecValidationBlock() bool`

HasDnssecValidationBlock returns a boolean if a field has been set.

### GetDtcConfig

`func (o *ViewInheritance) GetDtcConfig() InheritedDtcConfig`

GetDtcConfig returns the DtcConfig field if non-nil, zero value otherwise.

### GetDtcConfigOk

`func (o *ViewInheritance) GetDtcConfigOk() (*InheritedDtcConfig, bool)`

GetDtcConfigOk returns a tuple with the DtcConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcConfig

`func (o *ViewInheritance) SetDtcConfig(v InheritedDtcConfig)`

SetDtcConfig sets DtcConfig field to given value.

### HasDtcConfig

`func (o *ViewInheritance) HasDtcConfig() bool`

HasDtcConfig returns a boolean if a field has been set.

### GetEcsBlock

`func (o *ViewInheritance) GetEcsBlock() InheritedECSBlock`

GetEcsBlock returns the EcsBlock field if non-nil, zero value otherwise.

### GetEcsBlockOk

`func (o *ViewInheritance) GetEcsBlockOk() (*InheritedECSBlock, bool)`

GetEcsBlockOk returns a tuple with the EcsBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsBlock

`func (o *ViewInheritance) SetEcsBlock(v InheritedECSBlock)`

SetEcsBlock sets EcsBlock field to given value.

### HasEcsBlock

`func (o *ViewInheritance) HasEcsBlock() bool`

HasEcsBlock returns a boolean if a field has been set.

### GetEdnsUdpSize

`func (o *ViewInheritance) GetEdnsUdpSize() Inheritance2InheritedUInt32`

GetEdnsUdpSize returns the EdnsUdpSize field if non-nil, zero value otherwise.

### GetEdnsUdpSizeOk

`func (o *ViewInheritance) GetEdnsUdpSizeOk() (*Inheritance2InheritedUInt32, bool)`

GetEdnsUdpSizeOk returns a tuple with the EdnsUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdnsUdpSize

`func (o *ViewInheritance) SetEdnsUdpSize(v Inheritance2InheritedUInt32)`

SetEdnsUdpSize sets EdnsUdpSize field to given value.

### HasEdnsUdpSize

`func (o *ViewInheritance) HasEdnsUdpSize() bool`

HasEdnsUdpSize returns a boolean if a field has been set.

### GetFilterAaaaAcl

`func (o *ViewInheritance) GetFilterAaaaAcl() InheritedACLItems`

GetFilterAaaaAcl returns the FilterAaaaAcl field if non-nil, zero value otherwise.

### GetFilterAaaaAclOk

`func (o *ViewInheritance) GetFilterAaaaAclOk() (*InheritedACLItems, bool)`

GetFilterAaaaAclOk returns a tuple with the FilterAaaaAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaaAcl

`func (o *ViewInheritance) SetFilterAaaaAcl(v InheritedACLItems)`

SetFilterAaaaAcl sets FilterAaaaAcl field to given value.

### HasFilterAaaaAcl

`func (o *ViewInheritance) HasFilterAaaaAcl() bool`

HasFilterAaaaAcl returns a boolean if a field has been set.

### GetFilterAaaaOnV4

`func (o *ViewInheritance) GetFilterAaaaOnV4() Inheritance2InheritedString`

GetFilterAaaaOnV4 returns the FilterAaaaOnV4 field if non-nil, zero value otherwise.

### GetFilterAaaaOnV4Ok

`func (o *ViewInheritance) GetFilterAaaaOnV4Ok() (*Inheritance2InheritedString, bool)`

GetFilterAaaaOnV4Ok returns a tuple with the FilterAaaaOnV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaaOnV4

`func (o *ViewInheritance) SetFilterAaaaOnV4(v Inheritance2InheritedString)`

SetFilterAaaaOnV4 sets FilterAaaaOnV4 field to given value.

### HasFilterAaaaOnV4

`func (o *ViewInheritance) HasFilterAaaaOnV4() bool`

HasFilterAaaaOnV4 returns a boolean if a field has been set.

### GetForwardersBlock

`func (o *ViewInheritance) GetForwardersBlock() InheritedForwardersBlock`

GetForwardersBlock returns the ForwardersBlock field if non-nil, zero value otherwise.

### GetForwardersBlockOk

`func (o *ViewInheritance) GetForwardersBlockOk() (*InheritedForwardersBlock, bool)`

GetForwardersBlockOk returns a tuple with the ForwardersBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardersBlock

`func (o *ViewInheritance) SetForwardersBlock(v InheritedForwardersBlock)`

SetForwardersBlock sets ForwardersBlock field to given value.

### HasForwardersBlock

`func (o *ViewInheritance) HasForwardersBlock() bool`

HasForwardersBlock returns a boolean if a field has been set.

### GetGssTsigEnabled

`func (o *ViewInheritance) GetGssTsigEnabled() Inheritance2InheritedBool`

GetGssTsigEnabled returns the GssTsigEnabled field if non-nil, zero value otherwise.

### GetGssTsigEnabledOk

`func (o *ViewInheritance) GetGssTsigEnabledOk() (*Inheritance2InheritedBool, bool)`

GetGssTsigEnabledOk returns a tuple with the GssTsigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigEnabled

`func (o *ViewInheritance) SetGssTsigEnabled(v Inheritance2InheritedBool)`

SetGssTsigEnabled sets GssTsigEnabled field to given value.

### HasGssTsigEnabled

`func (o *ViewInheritance) HasGssTsigEnabled() bool`

HasGssTsigEnabled returns a boolean if a field has been set.

### GetLameTtl

`func (o *ViewInheritance) GetLameTtl() Inheritance2InheritedUInt32`

GetLameTtl returns the LameTtl field if non-nil, zero value otherwise.

### GetLameTtlOk

`func (o *ViewInheritance) GetLameTtlOk() (*Inheritance2InheritedUInt32, bool)`

GetLameTtlOk returns a tuple with the LameTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLameTtl

`func (o *ViewInheritance) SetLameTtl(v Inheritance2InheritedUInt32)`

SetLameTtl sets LameTtl field to given value.

### HasLameTtl

`func (o *ViewInheritance) HasLameTtl() bool`

HasLameTtl returns a boolean if a field has been set.

### GetMatchRecursiveOnly

`func (o *ViewInheritance) GetMatchRecursiveOnly() Inheritance2InheritedBool`

GetMatchRecursiveOnly returns the MatchRecursiveOnly field if non-nil, zero value otherwise.

### GetMatchRecursiveOnlyOk

`func (o *ViewInheritance) GetMatchRecursiveOnlyOk() (*Inheritance2InheritedBool, bool)`

GetMatchRecursiveOnlyOk returns a tuple with the MatchRecursiveOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchRecursiveOnly

`func (o *ViewInheritance) SetMatchRecursiveOnly(v Inheritance2InheritedBool)`

SetMatchRecursiveOnly sets MatchRecursiveOnly field to given value.

### HasMatchRecursiveOnly

`func (o *ViewInheritance) HasMatchRecursiveOnly() bool`

HasMatchRecursiveOnly returns a boolean if a field has been set.

### GetMaxCacheTtl

`func (o *ViewInheritance) GetMaxCacheTtl() Inheritance2InheritedUInt32`

GetMaxCacheTtl returns the MaxCacheTtl field if non-nil, zero value otherwise.

### GetMaxCacheTtlOk

`func (o *ViewInheritance) GetMaxCacheTtlOk() (*Inheritance2InheritedUInt32, bool)`

GetMaxCacheTtlOk returns a tuple with the MaxCacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheTtl

`func (o *ViewInheritance) SetMaxCacheTtl(v Inheritance2InheritedUInt32)`

SetMaxCacheTtl sets MaxCacheTtl field to given value.

### HasMaxCacheTtl

`func (o *ViewInheritance) HasMaxCacheTtl() bool`

HasMaxCacheTtl returns a boolean if a field has been set.

### GetMaxNegativeTtl

`func (o *ViewInheritance) GetMaxNegativeTtl() Inheritance2InheritedUInt32`

GetMaxNegativeTtl returns the MaxNegativeTtl field if non-nil, zero value otherwise.

### GetMaxNegativeTtlOk

`func (o *ViewInheritance) GetMaxNegativeTtlOk() (*Inheritance2InheritedUInt32, bool)`

GetMaxNegativeTtlOk returns a tuple with the MaxNegativeTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNegativeTtl

`func (o *ViewInheritance) SetMaxNegativeTtl(v Inheritance2InheritedUInt32)`

SetMaxNegativeTtl sets MaxNegativeTtl field to given value.

### HasMaxNegativeTtl

`func (o *ViewInheritance) HasMaxNegativeTtl() bool`

HasMaxNegativeTtl returns a boolean if a field has been set.

### GetMaxUdpSize

`func (o *ViewInheritance) GetMaxUdpSize() Inheritance2InheritedUInt32`

GetMaxUdpSize returns the MaxUdpSize field if non-nil, zero value otherwise.

### GetMaxUdpSizeOk

`func (o *ViewInheritance) GetMaxUdpSizeOk() (*Inheritance2InheritedUInt32, bool)`

GetMaxUdpSizeOk returns a tuple with the MaxUdpSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUdpSize

`func (o *ViewInheritance) SetMaxUdpSize(v Inheritance2InheritedUInt32)`

SetMaxUdpSize sets MaxUdpSize field to given value.

### HasMaxUdpSize

`func (o *ViewInheritance) HasMaxUdpSize() bool`

HasMaxUdpSize returns a boolean if a field has been set.

### GetMinimalResponses

`func (o *ViewInheritance) GetMinimalResponses() Inheritance2InheritedBool`

GetMinimalResponses returns the MinimalResponses field if non-nil, zero value otherwise.

### GetMinimalResponsesOk

`func (o *ViewInheritance) GetMinimalResponsesOk() (*Inheritance2InheritedBool, bool)`

GetMinimalResponsesOk returns a tuple with the MinimalResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimalResponses

`func (o *ViewInheritance) SetMinimalResponses(v Inheritance2InheritedBool)`

SetMinimalResponses sets MinimalResponses field to given value.

### HasMinimalResponses

`func (o *ViewInheritance) HasMinimalResponses() bool`

HasMinimalResponses returns a boolean if a field has been set.

### GetNotify

`func (o *ViewInheritance) GetNotify() Inheritance2InheritedBool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *ViewInheritance) GetNotifyOk() (*Inheritance2InheritedBool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *ViewInheritance) SetNotify(v Inheritance2InheritedBool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *ViewInheritance) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetQueryAcl

`func (o *ViewInheritance) GetQueryAcl() InheritedACLItems`

GetQueryAcl returns the QueryAcl field if non-nil, zero value otherwise.

### GetQueryAclOk

`func (o *ViewInheritance) GetQueryAclOk() (*InheritedACLItems, bool)`

GetQueryAclOk returns a tuple with the QueryAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryAcl

`func (o *ViewInheritance) SetQueryAcl(v InheritedACLItems)`

SetQueryAcl sets QueryAcl field to given value.

### HasQueryAcl

`func (o *ViewInheritance) HasQueryAcl() bool`

HasQueryAcl returns a boolean if a field has been set.

### GetRecursionAcl

`func (o *ViewInheritance) GetRecursionAcl() InheritedACLItems`

GetRecursionAcl returns the RecursionAcl field if non-nil, zero value otherwise.

### GetRecursionAclOk

`func (o *ViewInheritance) GetRecursionAclOk() (*InheritedACLItems, bool)`

GetRecursionAclOk returns a tuple with the RecursionAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursionAcl

`func (o *ViewInheritance) SetRecursionAcl(v InheritedACLItems)`

SetRecursionAcl sets RecursionAcl field to given value.

### HasRecursionAcl

`func (o *ViewInheritance) HasRecursionAcl() bool`

HasRecursionAcl returns a boolean if a field has been set.

### GetRecursionEnabled

`func (o *ViewInheritance) GetRecursionEnabled() Inheritance2InheritedBool`

GetRecursionEnabled returns the RecursionEnabled field if non-nil, zero value otherwise.

### GetRecursionEnabledOk

`func (o *ViewInheritance) GetRecursionEnabledOk() (*Inheritance2InheritedBool, bool)`

GetRecursionEnabledOk returns a tuple with the RecursionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursionEnabled

`func (o *ViewInheritance) SetRecursionEnabled(v Inheritance2InheritedBool)`

SetRecursionEnabled sets RecursionEnabled field to given value.

### HasRecursionEnabled

`func (o *ViewInheritance) HasRecursionEnabled() bool`

HasRecursionEnabled returns a boolean if a field has been set.

### GetSortList

`func (o *ViewInheritance) GetSortList() InheritedSortListItems`

GetSortList returns the SortList field if non-nil, zero value otherwise.

### GetSortListOk

`func (o *ViewInheritance) GetSortListOk() (*InheritedSortListItems, bool)`

GetSortListOk returns a tuple with the SortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortList

`func (o *ViewInheritance) SetSortList(v InheritedSortListItems)`

SetSortList sets SortList field to given value.

### HasSortList

`func (o *ViewInheritance) HasSortList() bool`

HasSortList returns a boolean if a field has been set.

### GetSynthesizeAddressRecordsFromHttps

`func (o *ViewInheritance) GetSynthesizeAddressRecordsFromHttps() Inheritance2InheritedBool`

GetSynthesizeAddressRecordsFromHttps returns the SynthesizeAddressRecordsFromHttps field if non-nil, zero value otherwise.

### GetSynthesizeAddressRecordsFromHttpsOk

`func (o *ViewInheritance) GetSynthesizeAddressRecordsFromHttpsOk() (*Inheritance2InheritedBool, bool)`

GetSynthesizeAddressRecordsFromHttpsOk returns a tuple with the SynthesizeAddressRecordsFromHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthesizeAddressRecordsFromHttps

`func (o *ViewInheritance) SetSynthesizeAddressRecordsFromHttps(v Inheritance2InheritedBool)`

SetSynthesizeAddressRecordsFromHttps sets SynthesizeAddressRecordsFromHttps field to given value.

### HasSynthesizeAddressRecordsFromHttps

`func (o *ViewInheritance) HasSynthesizeAddressRecordsFromHttps() bool`

HasSynthesizeAddressRecordsFromHttps returns a boolean if a field has been set.

### GetTransferAcl

`func (o *ViewInheritance) GetTransferAcl() InheritedACLItems`

GetTransferAcl returns the TransferAcl field if non-nil, zero value otherwise.

### GetTransferAclOk

`func (o *ViewInheritance) GetTransferAclOk() (*InheritedACLItems, bool)`

GetTransferAclOk returns a tuple with the TransferAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAcl

`func (o *ViewInheritance) SetTransferAcl(v InheritedACLItems)`

SetTransferAcl sets TransferAcl field to given value.

### HasTransferAcl

`func (o *ViewInheritance) HasTransferAcl() bool`

HasTransferAcl returns a boolean if a field has been set.

### GetUpdateAcl

`func (o *ViewInheritance) GetUpdateAcl() InheritedACLItems`

GetUpdateAcl returns the UpdateAcl field if non-nil, zero value otherwise.

### GetUpdateAclOk

`func (o *ViewInheritance) GetUpdateAclOk() (*InheritedACLItems, bool)`

GetUpdateAclOk returns a tuple with the UpdateAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAcl

`func (o *ViewInheritance) SetUpdateAcl(v InheritedACLItems)`

SetUpdateAcl sets UpdateAcl field to given value.

### HasUpdateAcl

`func (o *ViewInheritance) HasUpdateAcl() bool`

HasUpdateAcl returns a boolean if a field has been set.

### GetUseForwardersForSubzones

`func (o *ViewInheritance) GetUseForwardersForSubzones() Inheritance2InheritedBool`

GetUseForwardersForSubzones returns the UseForwardersForSubzones field if non-nil, zero value otherwise.

### GetUseForwardersForSubzonesOk

`func (o *ViewInheritance) GetUseForwardersForSubzonesOk() (*Inheritance2InheritedBool, bool)`

GetUseForwardersForSubzonesOk returns a tuple with the UseForwardersForSubzones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForwardersForSubzones

`func (o *ViewInheritance) SetUseForwardersForSubzones(v Inheritance2InheritedBool)`

SetUseForwardersForSubzones sets UseForwardersForSubzones field to given value.

### HasUseForwardersForSubzones

`func (o *ViewInheritance) HasUseForwardersForSubzones() bool`

HasUseForwardersForSubzones returns a boolean if a field has been set.

### GetZoneAuthority

`func (o *ViewInheritance) GetZoneAuthority() InheritedZoneAuthority`

GetZoneAuthority returns the ZoneAuthority field if non-nil, zero value otherwise.

### GetZoneAuthorityOk

`func (o *ViewInheritance) GetZoneAuthorityOk() (*InheritedZoneAuthority, bool)`

GetZoneAuthorityOk returns a tuple with the ZoneAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAuthority

`func (o *ViewInheritance) SetZoneAuthority(v InheritedZoneAuthority)`

SetZoneAuthority sets ZoneAuthority field to given value.

### HasZoneAuthority

`func (o *ViewInheritance) HasZoneAuthority() bool`

HasZoneAuthority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


