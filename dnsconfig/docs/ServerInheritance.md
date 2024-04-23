# ServerInheritance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddEdnsOptionInOutgoingQuery** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**CustomRootNsBlock** | Pointer to [**InheritedCustomRootNSBlock**](InheritedCustomRootNSBlock.md) |  | [optional] 
**DnssecValidationBlock** | Pointer to [**InheritedDNSSECValidationBlock**](InheritedDNSSECValidationBlock.md) |  | [optional] 
**EcsBlock** | Pointer to [**InheritedECSBlock**](InheritedECSBlock.md) |  | [optional] 
**FilterAaaaAcl** | Pointer to [**InheritedACLItems**](InheritedACLItems.md) |  | [optional] 
**FilterAaaaOnV4** | Pointer to [**Inheritance2InheritedString**](Inheritance2InheritedString.md) |  | [optional] 
**ForwardersBlock** | Pointer to [**InheritedForwardersBlock**](InheritedForwardersBlock.md) |  | [optional] 
**GssTsigEnabled** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**KerberosKeys** | Pointer to [**InheritedKerberosKeys**](InheritedKerberosKeys.md) |  | [optional] 
**LameTtl** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**LogQueryResponse** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**MatchRecursiveOnly** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**MaxCacheTtl** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**MaxNegativeTtl** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**MinimalResponses** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**Notify** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**QueryAcl** | Pointer to [**InheritedACLItems**](InheritedACLItems.md) |  | [optional] 
**QueryPort** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**RecursionAcl** | Pointer to [**InheritedACLItems**](InheritedACLItems.md) |  | [optional] 
**RecursionEnabled** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**RecursiveClients** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**ResolverQueryTimeout** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**SecondaryAxfrQueryLimit** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**SecondarySoaQueryLimit** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**SortList** | Pointer to [**InheritedSortListItems**](InheritedSortListItems.md) |  | [optional] 
**SynthesizeAddressRecordsFromHttps** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**TransferAcl** | Pointer to [**InheritedACLItems**](InheritedACLItems.md) |  | [optional] 
**UpdateAcl** | Pointer to [**InheritedACLItems**](InheritedACLItems.md) |  | [optional] 
**UseForwardersForSubzones** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 

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

### GetAddEdnsOptionInOutgoingQuery

`func (o *ServerInheritance) GetAddEdnsOptionInOutgoingQuery() Inheritance2InheritedBool`

GetAddEdnsOptionInOutgoingQuery returns the AddEdnsOptionInOutgoingQuery field if non-nil, zero value otherwise.

### GetAddEdnsOptionInOutgoingQueryOk

`func (o *ServerInheritance) GetAddEdnsOptionInOutgoingQueryOk() (*Inheritance2InheritedBool, bool)`

GetAddEdnsOptionInOutgoingQueryOk returns a tuple with the AddEdnsOptionInOutgoingQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddEdnsOptionInOutgoingQuery

`func (o *ServerInheritance) SetAddEdnsOptionInOutgoingQuery(v Inheritance2InheritedBool)`

SetAddEdnsOptionInOutgoingQuery sets AddEdnsOptionInOutgoingQuery field to given value.

### HasAddEdnsOptionInOutgoingQuery

`func (o *ServerInheritance) HasAddEdnsOptionInOutgoingQuery() bool`

HasAddEdnsOptionInOutgoingQuery returns a boolean if a field has been set.

### GetCustomRootNsBlock

`func (o *ServerInheritance) GetCustomRootNsBlock() InheritedCustomRootNSBlock`

GetCustomRootNsBlock returns the CustomRootNsBlock field if non-nil, zero value otherwise.

### GetCustomRootNsBlockOk

`func (o *ServerInheritance) GetCustomRootNsBlockOk() (*InheritedCustomRootNSBlock, bool)`

GetCustomRootNsBlockOk returns a tuple with the CustomRootNsBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRootNsBlock

`func (o *ServerInheritance) SetCustomRootNsBlock(v InheritedCustomRootNSBlock)`

SetCustomRootNsBlock sets CustomRootNsBlock field to given value.

### HasCustomRootNsBlock

`func (o *ServerInheritance) HasCustomRootNsBlock() bool`

HasCustomRootNsBlock returns a boolean if a field has been set.

### GetDnssecValidationBlock

`func (o *ServerInheritance) GetDnssecValidationBlock() InheritedDNSSECValidationBlock`

GetDnssecValidationBlock returns the DnssecValidationBlock field if non-nil, zero value otherwise.

### GetDnssecValidationBlockOk

`func (o *ServerInheritance) GetDnssecValidationBlockOk() (*InheritedDNSSECValidationBlock, bool)`

GetDnssecValidationBlockOk returns a tuple with the DnssecValidationBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecValidationBlock

`func (o *ServerInheritance) SetDnssecValidationBlock(v InheritedDNSSECValidationBlock)`

SetDnssecValidationBlock sets DnssecValidationBlock field to given value.

### HasDnssecValidationBlock

`func (o *ServerInheritance) HasDnssecValidationBlock() bool`

HasDnssecValidationBlock returns a boolean if a field has been set.

### GetEcsBlock

`func (o *ServerInheritance) GetEcsBlock() InheritedECSBlock`

GetEcsBlock returns the EcsBlock field if non-nil, zero value otherwise.

### GetEcsBlockOk

`func (o *ServerInheritance) GetEcsBlockOk() (*InheritedECSBlock, bool)`

GetEcsBlockOk returns a tuple with the EcsBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsBlock

`func (o *ServerInheritance) SetEcsBlock(v InheritedECSBlock)`

SetEcsBlock sets EcsBlock field to given value.

### HasEcsBlock

`func (o *ServerInheritance) HasEcsBlock() bool`

HasEcsBlock returns a boolean if a field has been set.

### GetFilterAaaaAcl

`func (o *ServerInheritance) GetFilterAaaaAcl() InheritedACLItems`

GetFilterAaaaAcl returns the FilterAaaaAcl field if non-nil, zero value otherwise.

### GetFilterAaaaAclOk

`func (o *ServerInheritance) GetFilterAaaaAclOk() (*InheritedACLItems, bool)`

GetFilterAaaaAclOk returns a tuple with the FilterAaaaAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaaAcl

`func (o *ServerInheritance) SetFilterAaaaAcl(v InheritedACLItems)`

SetFilterAaaaAcl sets FilterAaaaAcl field to given value.

### HasFilterAaaaAcl

`func (o *ServerInheritance) HasFilterAaaaAcl() bool`

HasFilterAaaaAcl returns a boolean if a field has been set.

### GetFilterAaaaOnV4

`func (o *ServerInheritance) GetFilterAaaaOnV4() Inheritance2InheritedString`

GetFilterAaaaOnV4 returns the FilterAaaaOnV4 field if non-nil, zero value otherwise.

### GetFilterAaaaOnV4Ok

`func (o *ServerInheritance) GetFilterAaaaOnV4Ok() (*Inheritance2InheritedString, bool)`

GetFilterAaaaOnV4Ok returns a tuple with the FilterAaaaOnV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAaaaOnV4

`func (o *ServerInheritance) SetFilterAaaaOnV4(v Inheritance2InheritedString)`

SetFilterAaaaOnV4 sets FilterAaaaOnV4 field to given value.

### HasFilterAaaaOnV4

`func (o *ServerInheritance) HasFilterAaaaOnV4() bool`

HasFilterAaaaOnV4 returns a boolean if a field has been set.

### GetForwardersBlock

`func (o *ServerInheritance) GetForwardersBlock() InheritedForwardersBlock`

GetForwardersBlock returns the ForwardersBlock field if non-nil, zero value otherwise.

### GetForwardersBlockOk

`func (o *ServerInheritance) GetForwardersBlockOk() (*InheritedForwardersBlock, bool)`

GetForwardersBlockOk returns a tuple with the ForwardersBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardersBlock

`func (o *ServerInheritance) SetForwardersBlock(v InheritedForwardersBlock)`

SetForwardersBlock sets ForwardersBlock field to given value.

### HasForwardersBlock

`func (o *ServerInheritance) HasForwardersBlock() bool`

HasForwardersBlock returns a boolean if a field has been set.

### GetGssTsigEnabled

`func (o *ServerInheritance) GetGssTsigEnabled() Inheritance2InheritedBool`

GetGssTsigEnabled returns the GssTsigEnabled field if non-nil, zero value otherwise.

### GetGssTsigEnabledOk

`func (o *ServerInheritance) GetGssTsigEnabledOk() (*Inheritance2InheritedBool, bool)`

GetGssTsigEnabledOk returns a tuple with the GssTsigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigEnabled

`func (o *ServerInheritance) SetGssTsigEnabled(v Inheritance2InheritedBool)`

SetGssTsigEnabled sets GssTsigEnabled field to given value.

### HasGssTsigEnabled

`func (o *ServerInheritance) HasGssTsigEnabled() bool`

HasGssTsigEnabled returns a boolean if a field has been set.

### GetKerberosKeys

`func (o *ServerInheritance) GetKerberosKeys() InheritedKerberosKeys`

GetKerberosKeys returns the KerberosKeys field if non-nil, zero value otherwise.

### GetKerberosKeysOk

`func (o *ServerInheritance) GetKerberosKeysOk() (*InheritedKerberosKeys, bool)`

GetKerberosKeysOk returns a tuple with the KerberosKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosKeys

`func (o *ServerInheritance) SetKerberosKeys(v InheritedKerberosKeys)`

SetKerberosKeys sets KerberosKeys field to given value.

### HasKerberosKeys

`func (o *ServerInheritance) HasKerberosKeys() bool`

HasKerberosKeys returns a boolean if a field has been set.

### GetLameTtl

`func (o *ServerInheritance) GetLameTtl() Inheritance2InheritedUInt32`

GetLameTtl returns the LameTtl field if non-nil, zero value otherwise.

### GetLameTtlOk

`func (o *ServerInheritance) GetLameTtlOk() (*Inheritance2InheritedUInt32, bool)`

GetLameTtlOk returns a tuple with the LameTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLameTtl

`func (o *ServerInheritance) SetLameTtl(v Inheritance2InheritedUInt32)`

SetLameTtl sets LameTtl field to given value.

### HasLameTtl

`func (o *ServerInheritance) HasLameTtl() bool`

HasLameTtl returns a boolean if a field has been set.

### GetLogQueryResponse

`func (o *ServerInheritance) GetLogQueryResponse() Inheritance2InheritedBool`

GetLogQueryResponse returns the LogQueryResponse field if non-nil, zero value otherwise.

### GetLogQueryResponseOk

`func (o *ServerInheritance) GetLogQueryResponseOk() (*Inheritance2InheritedBool, bool)`

GetLogQueryResponseOk returns a tuple with the LogQueryResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogQueryResponse

`func (o *ServerInheritance) SetLogQueryResponse(v Inheritance2InheritedBool)`

SetLogQueryResponse sets LogQueryResponse field to given value.

### HasLogQueryResponse

`func (o *ServerInheritance) HasLogQueryResponse() bool`

HasLogQueryResponse returns a boolean if a field has been set.

### GetMatchRecursiveOnly

`func (o *ServerInheritance) GetMatchRecursiveOnly() Inheritance2InheritedBool`

GetMatchRecursiveOnly returns the MatchRecursiveOnly field if non-nil, zero value otherwise.

### GetMatchRecursiveOnlyOk

`func (o *ServerInheritance) GetMatchRecursiveOnlyOk() (*Inheritance2InheritedBool, bool)`

GetMatchRecursiveOnlyOk returns a tuple with the MatchRecursiveOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchRecursiveOnly

`func (o *ServerInheritance) SetMatchRecursiveOnly(v Inheritance2InheritedBool)`

SetMatchRecursiveOnly sets MatchRecursiveOnly field to given value.

### HasMatchRecursiveOnly

`func (o *ServerInheritance) HasMatchRecursiveOnly() bool`

HasMatchRecursiveOnly returns a boolean if a field has been set.

### GetMaxCacheTtl

`func (o *ServerInheritance) GetMaxCacheTtl() Inheritance2InheritedUInt32`

GetMaxCacheTtl returns the MaxCacheTtl field if non-nil, zero value otherwise.

### GetMaxCacheTtlOk

`func (o *ServerInheritance) GetMaxCacheTtlOk() (*Inheritance2InheritedUInt32, bool)`

GetMaxCacheTtlOk returns a tuple with the MaxCacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCacheTtl

`func (o *ServerInheritance) SetMaxCacheTtl(v Inheritance2InheritedUInt32)`

SetMaxCacheTtl sets MaxCacheTtl field to given value.

### HasMaxCacheTtl

`func (o *ServerInheritance) HasMaxCacheTtl() bool`

HasMaxCacheTtl returns a boolean if a field has been set.

### GetMaxNegativeTtl

`func (o *ServerInheritance) GetMaxNegativeTtl() Inheritance2InheritedUInt32`

GetMaxNegativeTtl returns the MaxNegativeTtl field if non-nil, zero value otherwise.

### GetMaxNegativeTtlOk

`func (o *ServerInheritance) GetMaxNegativeTtlOk() (*Inheritance2InheritedUInt32, bool)`

GetMaxNegativeTtlOk returns a tuple with the MaxNegativeTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNegativeTtl

`func (o *ServerInheritance) SetMaxNegativeTtl(v Inheritance2InheritedUInt32)`

SetMaxNegativeTtl sets MaxNegativeTtl field to given value.

### HasMaxNegativeTtl

`func (o *ServerInheritance) HasMaxNegativeTtl() bool`

HasMaxNegativeTtl returns a boolean if a field has been set.

### GetMinimalResponses

`func (o *ServerInheritance) GetMinimalResponses() Inheritance2InheritedBool`

GetMinimalResponses returns the MinimalResponses field if non-nil, zero value otherwise.

### GetMinimalResponsesOk

`func (o *ServerInheritance) GetMinimalResponsesOk() (*Inheritance2InheritedBool, bool)`

GetMinimalResponsesOk returns a tuple with the MinimalResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimalResponses

`func (o *ServerInheritance) SetMinimalResponses(v Inheritance2InheritedBool)`

SetMinimalResponses sets MinimalResponses field to given value.

### HasMinimalResponses

`func (o *ServerInheritance) HasMinimalResponses() bool`

HasMinimalResponses returns a boolean if a field has been set.

### GetNotify

`func (o *ServerInheritance) GetNotify() Inheritance2InheritedBool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *ServerInheritance) GetNotifyOk() (*Inheritance2InheritedBool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *ServerInheritance) SetNotify(v Inheritance2InheritedBool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *ServerInheritance) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetQueryAcl

`func (o *ServerInheritance) GetQueryAcl() InheritedACLItems`

GetQueryAcl returns the QueryAcl field if non-nil, zero value otherwise.

### GetQueryAclOk

`func (o *ServerInheritance) GetQueryAclOk() (*InheritedACLItems, bool)`

GetQueryAclOk returns a tuple with the QueryAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryAcl

`func (o *ServerInheritance) SetQueryAcl(v InheritedACLItems)`

SetQueryAcl sets QueryAcl field to given value.

### HasQueryAcl

`func (o *ServerInheritance) HasQueryAcl() bool`

HasQueryAcl returns a boolean if a field has been set.

### GetQueryPort

`func (o *ServerInheritance) GetQueryPort() Inheritance2InheritedUInt32`

GetQueryPort returns the QueryPort field if non-nil, zero value otherwise.

### GetQueryPortOk

`func (o *ServerInheritance) GetQueryPortOk() (*Inheritance2InheritedUInt32, bool)`

GetQueryPortOk returns a tuple with the QueryPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryPort

`func (o *ServerInheritance) SetQueryPort(v Inheritance2InheritedUInt32)`

SetQueryPort sets QueryPort field to given value.

### HasQueryPort

`func (o *ServerInheritance) HasQueryPort() bool`

HasQueryPort returns a boolean if a field has been set.

### GetRecursionAcl

`func (o *ServerInheritance) GetRecursionAcl() InheritedACLItems`

GetRecursionAcl returns the RecursionAcl field if non-nil, zero value otherwise.

### GetRecursionAclOk

`func (o *ServerInheritance) GetRecursionAclOk() (*InheritedACLItems, bool)`

GetRecursionAclOk returns a tuple with the RecursionAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursionAcl

`func (o *ServerInheritance) SetRecursionAcl(v InheritedACLItems)`

SetRecursionAcl sets RecursionAcl field to given value.

### HasRecursionAcl

`func (o *ServerInheritance) HasRecursionAcl() bool`

HasRecursionAcl returns a boolean if a field has been set.

### GetRecursionEnabled

`func (o *ServerInheritance) GetRecursionEnabled() Inheritance2InheritedBool`

GetRecursionEnabled returns the RecursionEnabled field if non-nil, zero value otherwise.

### GetRecursionEnabledOk

`func (o *ServerInheritance) GetRecursionEnabledOk() (*Inheritance2InheritedBool, bool)`

GetRecursionEnabledOk returns a tuple with the RecursionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursionEnabled

`func (o *ServerInheritance) SetRecursionEnabled(v Inheritance2InheritedBool)`

SetRecursionEnabled sets RecursionEnabled field to given value.

### HasRecursionEnabled

`func (o *ServerInheritance) HasRecursionEnabled() bool`

HasRecursionEnabled returns a boolean if a field has been set.

### GetRecursiveClients

`func (o *ServerInheritance) GetRecursiveClients() Inheritance2InheritedUInt32`

GetRecursiveClients returns the RecursiveClients field if non-nil, zero value otherwise.

### GetRecursiveClientsOk

`func (o *ServerInheritance) GetRecursiveClientsOk() (*Inheritance2InheritedUInt32, bool)`

GetRecursiveClientsOk returns a tuple with the RecursiveClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursiveClients

`func (o *ServerInheritance) SetRecursiveClients(v Inheritance2InheritedUInt32)`

SetRecursiveClients sets RecursiveClients field to given value.

### HasRecursiveClients

`func (o *ServerInheritance) HasRecursiveClients() bool`

HasRecursiveClients returns a boolean if a field has been set.

### GetResolverQueryTimeout

`func (o *ServerInheritance) GetResolverQueryTimeout() Inheritance2InheritedUInt32`

GetResolverQueryTimeout returns the ResolverQueryTimeout field if non-nil, zero value otherwise.

### GetResolverQueryTimeoutOk

`func (o *ServerInheritance) GetResolverQueryTimeoutOk() (*Inheritance2InheritedUInt32, bool)`

GetResolverQueryTimeoutOk returns a tuple with the ResolverQueryTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolverQueryTimeout

`func (o *ServerInheritance) SetResolverQueryTimeout(v Inheritance2InheritedUInt32)`

SetResolverQueryTimeout sets ResolverQueryTimeout field to given value.

### HasResolverQueryTimeout

`func (o *ServerInheritance) HasResolverQueryTimeout() bool`

HasResolverQueryTimeout returns a boolean if a field has been set.

### GetSecondaryAxfrQueryLimit

`func (o *ServerInheritance) GetSecondaryAxfrQueryLimit() Inheritance2InheritedUInt32`

GetSecondaryAxfrQueryLimit returns the SecondaryAxfrQueryLimit field if non-nil, zero value otherwise.

### GetSecondaryAxfrQueryLimitOk

`func (o *ServerInheritance) GetSecondaryAxfrQueryLimitOk() (*Inheritance2InheritedUInt32, bool)`

GetSecondaryAxfrQueryLimitOk returns a tuple with the SecondaryAxfrQueryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAxfrQueryLimit

`func (o *ServerInheritance) SetSecondaryAxfrQueryLimit(v Inheritance2InheritedUInt32)`

SetSecondaryAxfrQueryLimit sets SecondaryAxfrQueryLimit field to given value.

### HasSecondaryAxfrQueryLimit

`func (o *ServerInheritance) HasSecondaryAxfrQueryLimit() bool`

HasSecondaryAxfrQueryLimit returns a boolean if a field has been set.

### GetSecondarySoaQueryLimit

`func (o *ServerInheritance) GetSecondarySoaQueryLimit() Inheritance2InheritedUInt32`

GetSecondarySoaQueryLimit returns the SecondarySoaQueryLimit field if non-nil, zero value otherwise.

### GetSecondarySoaQueryLimitOk

`func (o *ServerInheritance) GetSecondarySoaQueryLimitOk() (*Inheritance2InheritedUInt32, bool)`

GetSecondarySoaQueryLimitOk returns a tuple with the SecondarySoaQueryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySoaQueryLimit

`func (o *ServerInheritance) SetSecondarySoaQueryLimit(v Inheritance2InheritedUInt32)`

SetSecondarySoaQueryLimit sets SecondarySoaQueryLimit field to given value.

### HasSecondarySoaQueryLimit

`func (o *ServerInheritance) HasSecondarySoaQueryLimit() bool`

HasSecondarySoaQueryLimit returns a boolean if a field has been set.

### GetSortList

`func (o *ServerInheritance) GetSortList() InheritedSortListItems`

GetSortList returns the SortList field if non-nil, zero value otherwise.

### GetSortListOk

`func (o *ServerInheritance) GetSortListOk() (*InheritedSortListItems, bool)`

GetSortListOk returns a tuple with the SortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortList

`func (o *ServerInheritance) SetSortList(v InheritedSortListItems)`

SetSortList sets SortList field to given value.

### HasSortList

`func (o *ServerInheritance) HasSortList() bool`

HasSortList returns a boolean if a field has been set.

### GetSynthesizeAddressRecordsFromHttps

`func (o *ServerInheritance) GetSynthesizeAddressRecordsFromHttps() Inheritance2InheritedBool`

GetSynthesizeAddressRecordsFromHttps returns the SynthesizeAddressRecordsFromHttps field if non-nil, zero value otherwise.

### GetSynthesizeAddressRecordsFromHttpsOk

`func (o *ServerInheritance) GetSynthesizeAddressRecordsFromHttpsOk() (*Inheritance2InheritedBool, bool)`

GetSynthesizeAddressRecordsFromHttpsOk returns a tuple with the SynthesizeAddressRecordsFromHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthesizeAddressRecordsFromHttps

`func (o *ServerInheritance) SetSynthesizeAddressRecordsFromHttps(v Inheritance2InheritedBool)`

SetSynthesizeAddressRecordsFromHttps sets SynthesizeAddressRecordsFromHttps field to given value.

### HasSynthesizeAddressRecordsFromHttps

`func (o *ServerInheritance) HasSynthesizeAddressRecordsFromHttps() bool`

HasSynthesizeAddressRecordsFromHttps returns a boolean if a field has been set.

### GetTransferAcl

`func (o *ServerInheritance) GetTransferAcl() InheritedACLItems`

GetTransferAcl returns the TransferAcl field if non-nil, zero value otherwise.

### GetTransferAclOk

`func (o *ServerInheritance) GetTransferAclOk() (*InheritedACLItems, bool)`

GetTransferAclOk returns a tuple with the TransferAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAcl

`func (o *ServerInheritance) SetTransferAcl(v InheritedACLItems)`

SetTransferAcl sets TransferAcl field to given value.

### HasTransferAcl

`func (o *ServerInheritance) HasTransferAcl() bool`

HasTransferAcl returns a boolean if a field has been set.

### GetUpdateAcl

`func (o *ServerInheritance) GetUpdateAcl() InheritedACLItems`

GetUpdateAcl returns the UpdateAcl field if non-nil, zero value otherwise.

### GetUpdateAclOk

`func (o *ServerInheritance) GetUpdateAclOk() (*InheritedACLItems, bool)`

GetUpdateAclOk returns a tuple with the UpdateAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAcl

`func (o *ServerInheritance) SetUpdateAcl(v InheritedACLItems)`

SetUpdateAcl sets UpdateAcl field to given value.

### HasUpdateAcl

`func (o *ServerInheritance) HasUpdateAcl() bool`

HasUpdateAcl returns a boolean if a field has been set.

### GetUseForwardersForSubzones

`func (o *ServerInheritance) GetUseForwardersForSubzones() Inheritance2InheritedBool`

GetUseForwardersForSubzones returns the UseForwardersForSubzones field if non-nil, zero value otherwise.

### GetUseForwardersForSubzonesOk

`func (o *ServerInheritance) GetUseForwardersForSubzonesOk() (*Inheritance2InheritedBool, bool)`

GetUseForwardersForSubzonesOk returns a tuple with the UseForwardersForSubzones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForwardersForSubzones

`func (o *ServerInheritance) SetUseForwardersForSubzones(v Inheritance2InheritedBool)`

SetUseForwardersForSubzones sets UseForwardersForSubzones field to given value.

### HasUseForwardersForSubzones

`func (o *ServerInheritance) HasUseForwardersForSubzones() bool`

HasUseForwardersForSubzones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


