# AuthZoneInheritance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GssTsigEnabled** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**Notify** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**QueryAcl** | Pointer to [**InheritedACLItems**](InheritedACLItems.md) |  | [optional] 
**TransferAcl** | Pointer to [**InheritedACLItems**](InheritedACLItems.md) |  | [optional] 
**UpdateAcl** | Pointer to [**InheritedACLItems**](InheritedACLItems.md) |  | [optional] 
**UseForwardersForSubzones** | Pointer to [**Inheritance2InheritedBool**](Inheritance2InheritedBool.md) |  | [optional] 
**ZoneAuthority** | Pointer to [**InheritedZoneAuthority**](InheritedZoneAuthority.md) |  | [optional] 

## Methods

### NewAuthZoneInheritance

`func NewAuthZoneInheritance() *AuthZoneInheritance`

NewAuthZoneInheritance instantiates a new AuthZoneInheritance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthZoneInheritanceWithDefaults

`func NewAuthZoneInheritanceWithDefaults() *AuthZoneInheritance`

NewAuthZoneInheritanceWithDefaults instantiates a new AuthZoneInheritance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGssTsigEnabled

`func (o *AuthZoneInheritance) GetGssTsigEnabled() Inheritance2InheritedBool`

GetGssTsigEnabled returns the GssTsigEnabled field if non-nil, zero value otherwise.

### GetGssTsigEnabledOk

`func (o *AuthZoneInheritance) GetGssTsigEnabledOk() (*Inheritance2InheritedBool, bool)`

GetGssTsigEnabledOk returns a tuple with the GssTsigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigEnabled

`func (o *AuthZoneInheritance) SetGssTsigEnabled(v Inheritance2InheritedBool)`

SetGssTsigEnabled sets GssTsigEnabled field to given value.

### HasGssTsigEnabled

`func (o *AuthZoneInheritance) HasGssTsigEnabled() bool`

HasGssTsigEnabled returns a boolean if a field has been set.

### GetNotify

`func (o *AuthZoneInheritance) GetNotify() Inheritance2InheritedBool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *AuthZoneInheritance) GetNotifyOk() (*Inheritance2InheritedBool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *AuthZoneInheritance) SetNotify(v Inheritance2InheritedBool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *AuthZoneInheritance) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetQueryAcl

`func (o *AuthZoneInheritance) GetQueryAcl() InheritedACLItems`

GetQueryAcl returns the QueryAcl field if non-nil, zero value otherwise.

### GetQueryAclOk

`func (o *AuthZoneInheritance) GetQueryAclOk() (*InheritedACLItems, bool)`

GetQueryAclOk returns a tuple with the QueryAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryAcl

`func (o *AuthZoneInheritance) SetQueryAcl(v InheritedACLItems)`

SetQueryAcl sets QueryAcl field to given value.

### HasQueryAcl

`func (o *AuthZoneInheritance) HasQueryAcl() bool`

HasQueryAcl returns a boolean if a field has been set.

### GetTransferAcl

`func (o *AuthZoneInheritance) GetTransferAcl() InheritedACLItems`

GetTransferAcl returns the TransferAcl field if non-nil, zero value otherwise.

### GetTransferAclOk

`func (o *AuthZoneInheritance) GetTransferAclOk() (*InheritedACLItems, bool)`

GetTransferAclOk returns a tuple with the TransferAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAcl

`func (o *AuthZoneInheritance) SetTransferAcl(v InheritedACLItems)`

SetTransferAcl sets TransferAcl field to given value.

### HasTransferAcl

`func (o *AuthZoneInheritance) HasTransferAcl() bool`

HasTransferAcl returns a boolean if a field has been set.

### GetUpdateAcl

`func (o *AuthZoneInheritance) GetUpdateAcl() InheritedACLItems`

GetUpdateAcl returns the UpdateAcl field if non-nil, zero value otherwise.

### GetUpdateAclOk

`func (o *AuthZoneInheritance) GetUpdateAclOk() (*InheritedACLItems, bool)`

GetUpdateAclOk returns a tuple with the UpdateAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAcl

`func (o *AuthZoneInheritance) SetUpdateAcl(v InheritedACLItems)`

SetUpdateAcl sets UpdateAcl field to given value.

### HasUpdateAcl

`func (o *AuthZoneInheritance) HasUpdateAcl() bool`

HasUpdateAcl returns a boolean if a field has been set.

### GetUseForwardersForSubzones

`func (o *AuthZoneInheritance) GetUseForwardersForSubzones() Inheritance2InheritedBool`

GetUseForwardersForSubzones returns the UseForwardersForSubzones field if non-nil, zero value otherwise.

### GetUseForwardersForSubzonesOk

`func (o *AuthZoneInheritance) GetUseForwardersForSubzonesOk() (*Inheritance2InheritedBool, bool)`

GetUseForwardersForSubzonesOk returns a tuple with the UseForwardersForSubzones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForwardersForSubzones

`func (o *AuthZoneInheritance) SetUseForwardersForSubzones(v Inheritance2InheritedBool)`

SetUseForwardersForSubzones sets UseForwardersForSubzones field to given value.

### HasUseForwardersForSubzones

`func (o *AuthZoneInheritance) HasUseForwardersForSubzones() bool`

HasUseForwardersForSubzones returns a boolean if a field has been set.

### GetZoneAuthority

`func (o *AuthZoneInheritance) GetZoneAuthority() InheritedZoneAuthority`

GetZoneAuthority returns the ZoneAuthority field if non-nil, zero value otherwise.

### GetZoneAuthorityOk

`func (o *AuthZoneInheritance) GetZoneAuthorityOk() (*InheritedZoneAuthority, bool)`

GetZoneAuthorityOk returns a tuple with the ZoneAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAuthority

`func (o *AuthZoneInheritance) SetZoneAuthority(v InheritedZoneAuthority)`

SetZoneAuthority sets ZoneAuthority field to given value.

### HasZoneAuthority

`func (o *AuthZoneInheritance) HasZoneAuthority() bool`

HasZoneAuthority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


