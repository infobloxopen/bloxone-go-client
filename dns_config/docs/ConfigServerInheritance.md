# ConfigServerInheritance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddEdnsOptionInOutgoingQuery** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**CustomRootNsBlock** | Pointer to [**ConfigInheritedCustomRootNSBlock**](ConfigInheritedCustomRootNSBlock.md) |  | [optional] 
**DnssecValidationBlock** | Pointer to [**ConfigInheritedDNSSECValidationBlock**](ConfigInheritedDNSSECValidationBlock.md) |  | [optional] 
**EcsBlock** | Pointer to [**ConfigInheritedECSBlock**](ConfigInheritedECSBlock.md) |  | [optional] 
**FilterAaaaAcl** | Pointer to [**ConfigInheritedACLItems**](ConfigInheritedACLItems.md) |  | [optional] 
**FilterAaaaOnV4** | Pointer to [**Inheritance2InheritedString**](Inheritance2InheritedString.md) |  | [optional] 
**ForwardersBlock** | Pointer to [**ConfigInheritedForwardersBlock**](ConfigInheritedForwardersBlock.md) |  | [optional] 
**GssTsigEnabled** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**KerberosKeys** | Pointer to [**ConfigInheritedKerberosKeys**](ConfigInheritedKerberosKeys.md) |  | [optional] 
**LameTtl** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**LogQueryResponse** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**MatchRecursiveOnly** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**MaxCacheTtl** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**MaxNegativeTtl** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**MinimalResponses** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**Notify** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**QueryAcl** | Pointer to [**ConfigInheritedACLItems**](ConfigInheritedACLItems.md) |  | [optional] 
**QueryPort** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**RecursionAcl** | Pointer to [**ConfigInheritedACLItems**](ConfigInheritedACLItems.md) |  | [optional] 
**RecursionEnabled** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**RecursiveClients** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**ResolverQueryTimeout** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**SecondaryAxfrQueryLimit** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**SecondarySoaQueryLimit** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**SortList** | Pointer to [**ConfigInheritedSortListItems**](ConfigInheritedSortListItems.md) |  | [optional] 
**SynthesizeAddressRecordsFromHttps** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**TransferAcl** | Pointer to [**ConfigInheritedACLItems**](ConfigInheritedACLItems.md) |  | [optional] 
**UpdateAcl** | Pointer to [**ConfigInheritedACLItems**](ConfigInheritedACLItems.md) |  | [optional] 
**UseForwardersForSubzones** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 

## Methods

### NewConfigServerInheritance

`func NewConfigServerInheritance() *ConfigServerInheritance`

NewConfigServerInheritance instantiates a new ConfigServerInheritance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigServerInheritanceWithDefaults

`func NewConfigServerInheritanceWithDefaults() *ConfigServerInheritance`

NewConfigServerInheritanceWithDefaults instantiates a new ConfigServerInheritance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddEdnsOptionInOutgoingQuery

`func (o *ConfigServerInheritance) GetAddEdnsOptionInOutgoingQuery() Inheritance2InheritedBool`

GetAddEdnsOptionInOutgoingQuery returns the AddEdnsOptionInOutgoingQuery field if non-nil, zero value otherwise.

### GetAddEdnsOptionInOutgoingQueryOk

`func (o *ConfigServerInheritance) GetAddEdnsOptionInOutgoingQueryOk() (*Inheritance2InheritedBool, bool)`

GetAddEdnsOptionInOutgoingQueryOk returns a tuple with the AddEdnsOptionInOutgoingQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddEdnsOptionInOutgoingQuery

`func (o *ConfigServerInheritance) SetAddEdnsOptionInOutgoingQuery(v Inheritance2InheritedBool)`

SetAddEdnsOptionInOutgoingQuery sets AddEdnsOptionInOutgoingQuery field to given value.

### HasAddEdnsOptionInOutgoingQuery

`func (o *ConfigServerInheritance) HasAddEdnsOptionInOutgoingQuery() bool`

HasAddEdnsOptionInOutgoingQuery returns a boolean if a field has been set.

### GetCustomRootNsBlock

`func (o *ConfigServerInheritance) GetCustomRootNsBlock() ConfigInheritedCustomRootNSBlock`

GetCustomRootNsBlock returns the CustomRootNsBlock field if non-nil, zero value otherwise.

### GetCustomRootNsBlockOk

`func (o *ConfigServerInheritance) GetCustomRootNsBlockOk() (*ConfigInheritedCustomRootNSBlock, bool)`

GetCustomRootNsBlockOk returns a tuple with the CustomRootNsBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRootNsBlock

`func (o *ConfigServerInheritance) SetCustomRootNsBlock(v ConfigInheritedCustomRootNSBlock)`

SetCustomRootNsBlock sets CustomRootNsBlock field to given value.

### HasCustomRootNsBlock

`func (o *ConfigServerInheritance) HasCustomRootNsBlock() bool`

HasCustomRootNsBlock returns a boolean if a field has been set.

### GetDnssecValidationBlock

`func (o *ConfigServerInheritance) GetDnssecValidationBlock() ConfigInheritedDNSSECValidationBlock`

GetDnssecValidationBlock returns the DnssecValidationBlock field if non-nil, zero value otherwise.

### GetDnssecValidationBlockOk

`func (o *ConfigServerInheritance) GetDnssecValidationBlockOk() (*ConfigInheritedDNSSECValidationBlock, bool)`

GetDnssecValidationBlockOk returns a tuple with the DnssecValidationBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecValidationBlock

`func (o *ConfigServerInheritance) SetDnssecValidationBlock(v ConfigInheritedDNSSECValidationBlock)`

SetDnssecValidationBlock sets DnssecValidationBlock field to given value.

### HasDnssecValidationBlock

`func (o *ConfigServerInheritance) HasDnssecValidationBlock() bool`

HasDnssecValidationBlock returns a boolean if a field has been set.

### GetEcsBlock

`func (o *ConfigServerInheritance) GetEcsBlock() ConfigInheritedECSBlock`

GetEcsBlock returns the EcsBlock field if non-nil, zero value otherwise.

### GetEcsBlockOk

`func (o *ConfigServerInheritance) GetEcsBlockOk() (*ConfigInheritedECSBlock, bool)`

GetEcsBlockOk returns a tuple with the EcsBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsBlock

`func (o *ConfigServerInheritance) SetEcsBlock(v ConfigInheritedECSBlock)`

SetEcsBlock sets EcsBlock field to given value.

### HasEcsBlock

`func (o *ConfigServerInheritance) HasEcsBlock() bool`

HasEcsBlock returns a boolean if a field has been set.

### GetFilterAaaaAcl

`func (o *ConfigServerInheritance) GetFilterAaaaAcl() ConfigInheritedACLItems`

GetFilterAaaaAcl returns the FilterAaaaAcl field if non-nil, zero value otherwise.

### GetFilterAaaaAclOk

`func (o *ConfigServerInheritance) GetFilterAaaaAclOk() (*ConfigInheritedACLItems, bool)`

GetFilterAaaaAclOk returns a tuple with the FilterAaaaAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaaAcl

`func (o *ConfigServerInheritance) SetFilterAaaaAcl(v ConfigInheritedACLItems)`

SetFilterAaaaAcl sets FilterAaaaAcl field to given value.

### HasFilterAaaaAcl

`func (o *ConfigServerInheritance) HasFilterAaaaAcl() bool`

HasFilterAaaaAcl returns a boolean if a field has been set.

### GetFilterAaaaOnV4

`func (o *ConfigServerInheritance) GetFilterAaaaOnV4() Inheritance2InheritedString`

GetFilterAaaaOnV4 returns the FilterAaaaOnV4 field if non-nil, zero value otherwise.

### GetFilterAaaaOnV4Ok

`func (o *ConfigServerInheritance) GetFilterAaaaOnV4Ok() (*Inheritance2InheritedString, bool)`

GetFilterAaaaOnV4Ok returns a tuple with the FilterAaaaOnV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaaOnV4

`func (o *ConfigServerInheritance) SetFilterAaaaOnV4(v Inheritance2InheritedString)`

SetFilterAaaaOnV4 sets FilterAaaaOnV4 field to given value.

### HasFilterAaaaOnV4

`func (o *ConfigServerInheritance) HasFilterAaaaOnV4() bool`

HasFilterAaaaOnV4 returns a boolean if a field has been set.

### GetForwardersBlock

`func (o *ConfigServerInheritance) GetForwardersBlock() ConfigInheritedForwardersBlock`

GetForwardersBlock returns the ForwardersBlock field if non-nil, zero value otherwise.

### GetForwardersBlockOk

`func (o *ConfigServerInheritance) GetForwardersBlockOk() (*ConfigInheritedForwardersBlock, bool)`

GetForwardersBlockOk returns a tuple with the ForwardersBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardersBlock

`func (o *ConfigServerInheritance) SetForwardersBlock(v ConfigInheritedForwardersBlock)`

SetForwardersBlock sets ForwardersBlock field to given value.

### HasForwardersBlock

`func (o *ConfigServerInheritance) HasForwardersBlock() bool`

HasForwardersBlock returns a boolean if a field has been set.

### GetGssTsigEnabled

`func (o *ConfigServerInheritance) GetGssTsigEnabled() Inheritance2InheritedBool`

GetGssTsigEnabled returns the GssTsigEnabled field if non-nil, zero value otherwise.

### GetGssTsigEnabledOk

`func (o *ConfigServerInheritance) GetGssTsigEnabledOk() (*Inheritance2InheritedBool, bool)`

GetGssTsigEnabledOk returns a tuple with the GssTsigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigEnabled

`func (o *ConfigServerInheritance) SetGssTsigEnabled(v Inheritance2InheritedBool)`

SetGssTsigEnabled sets GssTsigEnabled field to given value.

### HasGssTsigEnabled

`func (o *ConfigServerInheritance) HasGssTsigEnabled() bool`

HasGssTsigEnabled returns a boolean if a field has been set.

### GetKerberosKeys

`func (o *ConfigServerInheritance) GetKerberosKeys() ConfigInheritedKerberosKeys`

GetKerberosKeys returns the KerberosKeys field if non-nil, zero value otherwise.

### GetKerberosKeysOk

`func (o *ConfigServerInheritance) GetKerberosKeysOk() (*ConfigInheritedKerberosKeys, bool)`

GetKerberosKeysOk returns a tuple with the KerberosKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosKeys

`func (o *ConfigServerInheritance) SetKerberosKeys(v ConfigInheritedKerberosKeys)`

SetKerberosKeys sets KerberosKeys field to given value.

### HasKerberosKeys

`func (o *ConfigServerInheritance) HasKerberosKeys() bool`

HasKerberosKeys returns a boolean if a field has been set.

### GetLameTtl

`func (o *ConfigServerInheritance) GetLameTtl() Inheritance2InheritedUInt32`

GetLameTtl returns the LameTtl field if non-nil, zero value otherwise.

### GetLameTtlOk

`func (o *ConfigServerInheritance) GetLameTtlOk() (*Inheritance2InheritedUInt32, bool)`

GetLameTtlOk returns a tuple with the LameTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLameTtl

`func (o *ConfigServerInheritance) SetLameTtl(v Inheritance2InheritedUInt32)`

SetLameTtl sets LameTtl field to given value.

### HasLameTtl

`func (o *ConfigServerInheritance) HasLameTtl() bool`

HasLameTtl returns a boolean if a field has been set.

### GetLogQueryResponse

`func (o *ConfigServerInheritance) GetLogQueryResponse() Inheritance2InheritedBool`

GetLogQueryResponse returns the LogQueryResponse field if non-nil, zero value otherwise.

### GetLogQueryResponseOk

`func (o *ConfigServerInheritance) GetLogQueryResponseOk() (*Inheritance2InheritedBool, bool)`

GetLogQueryResponseOk returns a tuple with the LogQueryResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogQueryResponse

`func (o *ConfigServerInheritance) SetLogQueryResponse(v Inheritance2InheritedBool)`

SetLogQueryResponse sets LogQueryResponse field to given value.

### HasLogQueryResponse

`func (o *ConfigServerInheritance) HasLogQueryResponse() bool`

HasLogQueryResponse returns a boolean if a field has been set.

### GetMatchRecursiveOnly

`func (o *ConfigServerInheritance) GetMatchRecursiveOnly() Inheritance2InheritedBool`

GetMatchRecursiveOnly returns the MatchRecursiveOnly field if non-nil, zero value otherwise.

### GetMatchRecursiveOnlyOk

`func (o *ConfigServerInheritance) GetMatchRecursiveOnlyOk() (*Inheritance2InheritedBool, bool)`

GetMatchRecursiveOnlyOk returns a tuple with the MatchRecursiveOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchRecursiveOnly

`func (o *ConfigServerInheritance) SetMatchRecursiveOnly(v Inheritance2InheritedBool)`

SetMatchRecursiveOnly sets MatchRecursiveOnly field to given value.

### HasMatchRecursiveOnly

`func (o *ConfigServerInheritance) HasMatchRecursiveOnly() bool`

HasMatchRecursiveOnly returns a boolean if a field has been set.

### GetMaxCacheTtl

`func (o *ConfigServerInheritance) GetMaxCacheTtl() Inheritance2InheritedUInt32`

GetMaxCacheTtl returns the MaxCacheTtl field if non-nil, zero value otherwise.

### GetMaxCacheTtlOk

`func (o *ConfigServerInheritance) GetMaxCacheTtlOk() (*Inheritance2InheritedUInt32, bool)`

GetMaxCacheTtlOk returns a tuple with the MaxCacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheTtl

`func (o *ConfigServerInheritance) SetMaxCacheTtl(v Inheritance2InheritedUInt32)`

SetMaxCacheTtl sets MaxCacheTtl field to given value.

### HasMaxCacheTtl

`func (o *ConfigServerInheritance) HasMaxCacheTtl() bool`

HasMaxCacheTtl returns a boolean if a field has been set.

### GetMaxNegativeTtl

`func (o *ConfigServerInheritance) GetMaxNegativeTtl() Inheritance2InheritedUInt32`

GetMaxNegativeTtl returns the MaxNegativeTtl field if non-nil, zero value otherwise.

### GetMaxNegativeTtlOk

`func (o *ConfigServerInheritance) GetMaxNegativeTtlOk() (*Inheritance2InheritedUInt32, bool)`

GetMaxNegativeTtlOk returns a tuple with the MaxNegativeTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNegativeTtl

`func (o *ConfigServerInheritance) SetMaxNegativeTtl(v Inheritance2InheritedUInt32)`

SetMaxNegativeTtl sets MaxNegativeTtl field to given value.

### HasMaxNegativeTtl

`func (o *ConfigServerInheritance) HasMaxNegativeTtl() bool`

HasMaxNegativeTtl returns a boolean if a field has been set.

### GetMinimalResponses

`func (o *ConfigServerInheritance) GetMinimalResponses() Inheritance2InheritedBool`

GetMinimalResponses returns the MinimalResponses field if non-nil, zero value otherwise.

### GetMinimalResponsesOk

`func (o *ConfigServerInheritance) GetMinimalResponsesOk() (*Inheritance2InheritedBool, bool)`

GetMinimalResponsesOk returns a tuple with the MinimalResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimalResponses

`func (o *ConfigServerInheritance) SetMinimalResponses(v Inheritance2InheritedBool)`

SetMinimalResponses sets MinimalResponses field to given value.

### HasMinimalResponses

`func (o *ConfigServerInheritance) HasMinimalResponses() bool`

HasMinimalResponses returns a boolean if a field has been set.

### GetNotify

`func (o *ConfigServerInheritance) GetNotify() Inheritance2InheritedBool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *ConfigServerInheritance) GetNotifyOk() (*Inheritance2InheritedBool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *ConfigServerInheritance) SetNotify(v Inheritance2InheritedBool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *ConfigServerInheritance) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetQueryAcl

`func (o *ConfigServerInheritance) GetQueryAcl() ConfigInheritedACLItems`

GetQueryAcl returns the QueryAcl field if non-nil, zero value otherwise.

### GetQueryAclOk

`func (o *ConfigServerInheritance) GetQueryAclOk() (*ConfigInheritedACLItems, bool)`

GetQueryAclOk returns a tuple with the QueryAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryAcl

`func (o *ConfigServerInheritance) SetQueryAcl(v ConfigInheritedACLItems)`

SetQueryAcl sets QueryAcl field to given value.

### HasQueryAcl

`func (o *ConfigServerInheritance) HasQueryAcl() bool`

HasQueryAcl returns a boolean if a field has been set.

### GetQueryPort

`func (o *ConfigServerInheritance) GetQueryPort() Inheritance2InheritedUInt32`

GetQueryPort returns the QueryPort field if non-nil, zero value otherwise.

### GetQueryPortOk

`func (o *ConfigServerInheritance) GetQueryPortOk() (*Inheritance2InheritedUInt32, bool)`

GetQueryPortOk returns a tuple with the QueryPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryPort

`func (o *ConfigServerInheritance) SetQueryPort(v Inheritance2InheritedUInt32)`

SetQueryPort sets QueryPort field to given value.

### HasQueryPort

`func (o *ConfigServerInheritance) HasQueryPort() bool`

HasQueryPort returns a boolean if a field has been set.

### GetRecursionAcl

`func (o *ConfigServerInheritance) GetRecursionAcl() ConfigInheritedACLItems`

GetRecursionAcl returns the RecursionAcl field if non-nil, zero value otherwise.

### GetRecursionAclOk

`func (o *ConfigServerInheritance) GetRecursionAclOk() (*ConfigInheritedACLItems, bool)`

GetRecursionAclOk returns a tuple with the RecursionAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursionAcl

`func (o *ConfigServerInheritance) SetRecursionAcl(v ConfigInheritedACLItems)`

SetRecursionAcl sets RecursionAcl field to given value.

### HasRecursionAcl

`func (o *ConfigServerInheritance) HasRecursionAcl() bool`

HasRecursionAcl returns a boolean if a field has been set.

### GetRecursionEnabled

`func (o *ConfigServerInheritance) GetRecursionEnabled() Inheritance2InheritedBool`

GetRecursionEnabled returns the RecursionEnabled field if non-nil, zero value otherwise.

### GetRecursionEnabledOk

`func (o *ConfigServerInheritance) GetRecursionEnabledOk() (*Inheritance2InheritedBool, bool)`

GetRecursionEnabledOk returns a tuple with the RecursionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursionEnabled

`func (o *ConfigServerInheritance) SetRecursionEnabled(v Inheritance2InheritedBool)`

SetRecursionEnabled sets RecursionEnabled field to given value.

### HasRecursionEnabled

`func (o *ConfigServerInheritance) HasRecursionEnabled() bool`

HasRecursionEnabled returns a boolean if a field has been set.

### GetRecursiveClients

`func (o *ConfigServerInheritance) GetRecursiveClients() Inheritance2InheritedUInt32`

GetRecursiveClients returns the RecursiveClients field if non-nil, zero value otherwise.

### GetRecursiveClientsOk

`func (o *ConfigServerInheritance) GetRecursiveClientsOk() (*Inheritance2InheritedUInt32, bool)`

GetRecursiveClientsOk returns a tuple with the RecursiveClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursiveClients

`func (o *ConfigServerInheritance) SetRecursiveClients(v Inheritance2InheritedUInt32)`

SetRecursiveClients sets RecursiveClients field to given value.

### HasRecursiveClients

`func (o *ConfigServerInheritance) HasRecursiveClients() bool`

HasRecursiveClients returns a boolean if a field has been set.

### GetResolverQueryTimeout

`func (o *ConfigServerInheritance) GetResolverQueryTimeout() Inheritance2InheritedUInt32`

GetResolverQueryTimeout returns the ResolverQueryTimeout field if non-nil, zero value otherwise.

### GetResolverQueryTimeoutOk

`func (o *ConfigServerInheritance) GetResolverQueryTimeoutOk() (*Inheritance2InheritedUInt32, bool)`

GetResolverQueryTimeoutOk returns a tuple with the ResolverQueryTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolverQueryTimeout

`func (o *ConfigServerInheritance) SetResolverQueryTimeout(v Inheritance2InheritedUInt32)`

SetResolverQueryTimeout sets ResolverQueryTimeout field to given value.

### HasResolverQueryTimeout

`func (o *ConfigServerInheritance) HasResolverQueryTimeout() bool`

HasResolverQueryTimeout returns a boolean if a field has been set.

### GetSecondaryAxfrQueryLimit

`func (o *ConfigServerInheritance) GetSecondaryAxfrQueryLimit() Inheritance2InheritedUInt32`

GetSecondaryAxfrQueryLimit returns the SecondaryAxfrQueryLimit field if non-nil, zero value otherwise.

### GetSecondaryAxfrQueryLimitOk

`func (o *ConfigServerInheritance) GetSecondaryAxfrQueryLimitOk() (*Inheritance2InheritedUInt32, bool)`

GetSecondaryAxfrQueryLimitOk returns a tuple with the SecondaryAxfrQueryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAxfrQueryLimit

`func (o *ConfigServerInheritance) SetSecondaryAxfrQueryLimit(v Inheritance2InheritedUInt32)`

SetSecondaryAxfrQueryLimit sets SecondaryAxfrQueryLimit field to given value.

### HasSecondaryAxfrQueryLimit

`func (o *ConfigServerInheritance) HasSecondaryAxfrQueryLimit() bool`

HasSecondaryAxfrQueryLimit returns a boolean if a field has been set.

### GetSecondarySoaQueryLimit

`func (o *ConfigServerInheritance) GetSecondarySoaQueryLimit() Inheritance2InheritedUInt32`

GetSecondarySoaQueryLimit returns the SecondarySoaQueryLimit field if non-nil, zero value otherwise.

### GetSecondarySoaQueryLimitOk

`func (o *ConfigServerInheritance) GetSecondarySoaQueryLimitOk() (*Inheritance2InheritedUInt32, bool)`

GetSecondarySoaQueryLimitOk returns a tuple with the SecondarySoaQueryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySoaQueryLimit

`func (o *ConfigServerInheritance) SetSecondarySoaQueryLimit(v Inheritance2InheritedUInt32)`

SetSecondarySoaQueryLimit sets SecondarySoaQueryLimit field to given value.

### HasSecondarySoaQueryLimit

`func (o *ConfigServerInheritance) HasSecondarySoaQueryLimit() bool`

HasSecondarySoaQueryLimit returns a boolean if a field has been set.

### GetSortList

`func (o *ConfigServerInheritance) GetSortList() ConfigInheritedSortListItems`

GetSortList returns the SortList field if non-nil, zero value otherwise.

### GetSortListOk

`func (o *ConfigServerInheritance) GetSortListOk() (*ConfigInheritedSortListItems, bool)`

GetSortListOk returns a tuple with the SortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortList

`func (o *ConfigServerInheritance) SetSortList(v ConfigInheritedSortListItems)`

SetSortList sets SortList field to given value.

### HasSortList

`func (o *ConfigServerInheritance) HasSortList() bool`

HasSortList returns a boolean if a field has been set.

### GetSynthesizeAddressRecordsFromHttps

`func (o *ConfigServerInheritance) GetSynthesizeAddressRecordsFromHttps() Inheritance2InheritedBool`

GetSynthesizeAddressRecordsFromHttps returns the SynthesizeAddressRecordsFromHttps field if non-nil, zero value otherwise.

### GetSynthesizeAddressRecordsFromHttpsOk

`func (o *ConfigServerInheritance) GetSynthesizeAddressRecordsFromHttpsOk() (*Inheritance2InheritedBool, bool)`

GetSynthesizeAddressRecordsFromHttpsOk returns a tuple with the SynthesizeAddressRecordsFromHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthesizeAddressRecordsFromHttps

`func (o *ConfigServerInheritance) SetSynthesizeAddressRecordsFromHttps(v Inheritance2InheritedBool)`

SetSynthesizeAddressRecordsFromHttps sets SynthesizeAddressRecordsFromHttps field to given value.

### HasSynthesizeAddressRecordsFromHttps

`func (o *ConfigServerInheritance) HasSynthesizeAddressRecordsFromHttps() bool`

HasSynthesizeAddressRecordsFromHttps returns a boolean if a field has been set.

### GetTransferAcl

`func (o *ConfigServerInheritance) GetTransferAcl() ConfigInheritedACLItems`

GetTransferAcl returns the TransferAcl field if non-nil, zero value otherwise.

### GetTransferAclOk

`func (o *ConfigServerInheritance) GetTransferAclOk() (*ConfigInheritedACLItems, bool)`

GetTransferAclOk returns a tuple with the TransferAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAcl

`func (o *ConfigServerInheritance) SetTransferAcl(v ConfigInheritedACLItems)`

SetTransferAcl sets TransferAcl field to given value.

### HasTransferAcl

`func (o *ConfigServerInheritance) HasTransferAcl() bool`

HasTransferAcl returns a boolean if a field has been set.

### GetUpdateAcl

`func (o *ConfigServerInheritance) GetUpdateAcl() ConfigInheritedACLItems`

GetUpdateAcl returns the UpdateAcl field if non-nil, zero value otherwise.

### GetUpdateAclOk

`func (o *ConfigServerInheritance) GetUpdateAclOk() (*ConfigInheritedACLItems, bool)`

GetUpdateAclOk returns a tuple with the UpdateAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAcl

`func (o *ConfigServerInheritance) SetUpdateAcl(v ConfigInheritedACLItems)`

SetUpdateAcl sets UpdateAcl field to given value.

### HasUpdateAcl

`func (o *ConfigServerInheritance) HasUpdateAcl() bool`

HasUpdateAcl returns a boolean if a field has been set.

### GetUseForwardersForSubzones

`func (o *ConfigServerInheritance) GetUseForwardersForSubzones() Inheritance2InheritedBool`

GetUseForwardersForSubzones returns the UseForwardersForSubzones field if non-nil, zero value otherwise.

### GetUseForwardersForSubzonesOk

`func (o *ConfigServerInheritance) GetUseForwardersForSubzonesOk() (*Inheritance2InheritedBool, bool)`

GetUseForwardersForSubzonesOk returns a tuple with the UseForwardersForSubzones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForwardersForSubzones

`func (o *ConfigServerInheritance) SetUseForwardersForSubzones(v Inheritance2InheritedBool)`

SetUseForwardersForSubzones sets UseForwardersForSubzones field to given value.

### HasUseForwardersForSubzones

`func (o *ConfigServerInheritance) HasUseForwardersForSubzones() bool`

HasUseForwardersForSubzones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


