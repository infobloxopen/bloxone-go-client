# ConfigInheritedZoneAuthority

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultTtl** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**Expire** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**MnameBlock** | Pointer to [**ConfigInheritedZoneAuthorityMNameBlock**](ConfigInheritedZoneAuthorityMNameBlock.md) |  | [optional] 
**NegativeTtl** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**ProtocolRname** | Pointer to [**Inheritance2InheritedString**](Inheritance2InheritedString.md) |  | [optional] 
**Refresh** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**Retry** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**Rname** | Pointer to [**Inheritance2InheritedString**](Inheritance2InheritedString.md) |  | [optional] 

## Methods

### NewConfigInheritedZoneAuthority

`func NewConfigInheritedZoneAuthority() *ConfigInheritedZoneAuthority`

NewConfigInheritedZoneAuthority instantiates a new ConfigInheritedZoneAuthority object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigInheritedZoneAuthorityWithDefaults

`func NewConfigInheritedZoneAuthorityWithDefaults() *ConfigInheritedZoneAuthority`

NewConfigInheritedZoneAuthorityWithDefaults instantiates a new ConfigInheritedZoneAuthority object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultTtl

`func (o *ConfigInheritedZoneAuthority) GetDefaultTtl() Inheritance2InheritedUInt32`

GetDefaultTtl returns the DefaultTtl field if non-nil, zero value otherwise.

### GetDefaultTtlOk

`func (o *ConfigInheritedZoneAuthority) GetDefaultTtlOk() (*Inheritance2InheritedUInt32, bool)`

GetDefaultTtlOk returns a tuple with the DefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTtl

`func (o *ConfigInheritedZoneAuthority) SetDefaultTtl(v Inheritance2InheritedUInt32)`

SetDefaultTtl sets DefaultTtl field to given value.

### HasDefaultTtl

`func (o *ConfigInheritedZoneAuthority) HasDefaultTtl() bool`

HasDefaultTtl returns a boolean if a field has been set.

### GetExpire

`func (o *ConfigInheritedZoneAuthority) GetExpire() Inheritance2InheritedUInt32`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *ConfigInheritedZoneAuthority) GetExpireOk() (*Inheritance2InheritedUInt32, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *ConfigInheritedZoneAuthority) SetExpire(v Inheritance2InheritedUInt32)`

SetExpire sets Expire field to given value.

### HasExpire

`func (o *ConfigInheritedZoneAuthority) HasExpire() bool`

HasExpire returns a boolean if a field has been set.

### GetMnameBlock

`func (o *ConfigInheritedZoneAuthority) GetMnameBlock() ConfigInheritedZoneAuthorityMNameBlock`

GetMnameBlock returns the MnameBlock field if non-nil, zero value otherwise.

### GetMnameBlockOk

`func (o *ConfigInheritedZoneAuthority) GetMnameBlockOk() (*ConfigInheritedZoneAuthorityMNameBlock, bool)`

GetMnameBlockOk returns a tuple with the MnameBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnameBlock

`func (o *ConfigInheritedZoneAuthority) SetMnameBlock(v ConfigInheritedZoneAuthorityMNameBlock)`

SetMnameBlock sets MnameBlock field to given value.

### HasMnameBlock

`func (o *ConfigInheritedZoneAuthority) HasMnameBlock() bool`

HasMnameBlock returns a boolean if a field has been set.

### GetNegativeTtl

`func (o *ConfigInheritedZoneAuthority) GetNegativeTtl() Inheritance2InheritedUInt32`

GetNegativeTtl returns the NegativeTtl field if non-nil, zero value otherwise.

### GetNegativeTtlOk

`func (o *ConfigInheritedZoneAuthority) GetNegativeTtlOk() (*Inheritance2InheritedUInt32, bool)`

GetNegativeTtlOk returns a tuple with the NegativeTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeTtl

`func (o *ConfigInheritedZoneAuthority) SetNegativeTtl(v Inheritance2InheritedUInt32)`

SetNegativeTtl sets NegativeTtl field to given value.

### HasNegativeTtl

`func (o *ConfigInheritedZoneAuthority) HasNegativeTtl() bool`

HasNegativeTtl returns a boolean if a field has been set.

### GetProtocolRname

`func (o *ConfigInheritedZoneAuthority) GetProtocolRname() Inheritance2InheritedString`

GetProtocolRname returns the ProtocolRname field if non-nil, zero value otherwise.

### GetProtocolRnameOk

`func (o *ConfigInheritedZoneAuthority) GetProtocolRnameOk() (*Inheritance2InheritedString, bool)`

GetProtocolRnameOk returns a tuple with the ProtocolRname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolRname

`func (o *ConfigInheritedZoneAuthority) SetProtocolRname(v Inheritance2InheritedString)`

SetProtocolRname sets ProtocolRname field to given value.

### HasProtocolRname

`func (o *ConfigInheritedZoneAuthority) HasProtocolRname() bool`

HasProtocolRname returns a boolean if a field has been set.

### GetRefresh

`func (o *ConfigInheritedZoneAuthority) GetRefresh() Inheritance2InheritedUInt32`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *ConfigInheritedZoneAuthority) GetRefreshOk() (*Inheritance2InheritedUInt32, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *ConfigInheritedZoneAuthority) SetRefresh(v Inheritance2InheritedUInt32)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *ConfigInheritedZoneAuthority) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetRetry

`func (o *ConfigInheritedZoneAuthority) GetRetry() Inheritance2InheritedUInt32`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *ConfigInheritedZoneAuthority) GetRetryOk() (*Inheritance2InheritedUInt32, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *ConfigInheritedZoneAuthority) SetRetry(v Inheritance2InheritedUInt32)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *ConfigInheritedZoneAuthority) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### GetRname

`func (o *ConfigInheritedZoneAuthority) GetRname() Inheritance2InheritedString`

GetRname returns the Rname field if non-nil, zero value otherwise.

### GetRnameOk

`func (o *ConfigInheritedZoneAuthority) GetRnameOk() (*Inheritance2InheritedString, bool)`

GetRnameOk returns a tuple with the Rname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRname

`func (o *ConfigInheritedZoneAuthority) SetRname(v Inheritance2InheritedString)`

SetRname sets Rname field to given value.

### HasRname

`func (o *ConfigInheritedZoneAuthority) HasRname() bool`

HasRname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


