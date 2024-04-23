# InheritedZoneAuthority

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultTtl** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**Expire** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**MnameBlock** | Pointer to [**InheritedZoneAuthorityMNameBlock**](InheritedZoneAuthorityMNameBlock.md) |  | [optional] 
**NegativeTtl** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**ProtocolRname** | Pointer to [**Inheritance2InheritedString**](Inheritance2InheritedString.md) |  | [optional] 
**Refresh** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**Retry** | Pointer to [**Inheritance2InheritedUInt32**](Inheritance2InheritedUInt32.md) |  | [optional] 
**Rname** | Pointer to [**Inheritance2InheritedString**](Inheritance2InheritedString.md) |  | [optional] 

## Methods

### NewInheritedZoneAuthority

`func NewInheritedZoneAuthority() *InheritedZoneAuthority`

NewInheritedZoneAuthority instantiates a new InheritedZoneAuthority object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInheritedZoneAuthorityWithDefaults

`func NewInheritedZoneAuthorityWithDefaults() *InheritedZoneAuthority`

NewInheritedZoneAuthorityWithDefaults instantiates a new InheritedZoneAuthority object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultTtl

`func (o *InheritedZoneAuthority) GetDefaultTtl() Inheritance2InheritedUInt32`

GetDefaultTtl returns the DefaultTtl field if non-nil, zero value otherwise.

### GetDefaultTtlOk

`func (o *InheritedZoneAuthority) GetDefaultTtlOk() (*Inheritance2InheritedUInt32, bool)`

GetDefaultTtlOk returns a tuple with the DefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTtl

`func (o *InheritedZoneAuthority) SetDefaultTtl(v Inheritance2InheritedUInt32)`

SetDefaultTtl sets DefaultTtl field to given value.

### HasDefaultTtl

`func (o *InheritedZoneAuthority) HasDefaultTtl() bool`

HasDefaultTtl returns a boolean if a field has been set.

### GetExpire

`func (o *InheritedZoneAuthority) GetExpire() Inheritance2InheritedUInt32`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *InheritedZoneAuthority) GetExpireOk() (*Inheritance2InheritedUInt32, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *InheritedZoneAuthority) SetExpire(v Inheritance2InheritedUInt32)`

SetExpire sets Expire field to given value.

### HasExpire

`func (o *InheritedZoneAuthority) HasExpire() bool`

HasExpire returns a boolean if a field has been set.

### GetMnameBlock

`func (o *InheritedZoneAuthority) GetMnameBlock() InheritedZoneAuthorityMNameBlock`

GetMnameBlock returns the MnameBlock field if non-nil, zero value otherwise.

### GetMnameBlockOk

`func (o *InheritedZoneAuthority) GetMnameBlockOk() (*InheritedZoneAuthorityMNameBlock, bool)`

GetMnameBlockOk returns a tuple with the MnameBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnameBlock

`func (o *InheritedZoneAuthority) SetMnameBlock(v InheritedZoneAuthorityMNameBlock)`

SetMnameBlock sets MnameBlock field to given value.

### HasMnameBlock

`func (o *InheritedZoneAuthority) HasMnameBlock() bool`

HasMnameBlock returns a boolean if a field has been set.

### GetNegativeTtl

`func (o *InheritedZoneAuthority) GetNegativeTtl() Inheritance2InheritedUInt32`

GetNegativeTtl returns the NegativeTtl field if non-nil, zero value otherwise.

### GetNegativeTtlOk

`func (o *InheritedZoneAuthority) GetNegativeTtlOk() (*Inheritance2InheritedUInt32, bool)`

GetNegativeTtlOk returns a tuple with the NegativeTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeTtl

`func (o *InheritedZoneAuthority) SetNegativeTtl(v Inheritance2InheritedUInt32)`

SetNegativeTtl sets NegativeTtl field to given value.

### HasNegativeTtl

`func (o *InheritedZoneAuthority) HasNegativeTtl() bool`

HasNegativeTtl returns a boolean if a field has been set.

### GetProtocolRname

`func (o *InheritedZoneAuthority) GetProtocolRname() Inheritance2InheritedString`

GetProtocolRname returns the ProtocolRname field if non-nil, zero value otherwise.

### GetProtocolRnameOk

`func (o *InheritedZoneAuthority) GetProtocolRnameOk() (*Inheritance2InheritedString, bool)`

GetProtocolRnameOk returns a tuple with the ProtocolRname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolRname

`func (o *InheritedZoneAuthority) SetProtocolRname(v Inheritance2InheritedString)`

SetProtocolRname sets ProtocolRname field to given value.

### HasProtocolRname

`func (o *InheritedZoneAuthority) HasProtocolRname() bool`

HasProtocolRname returns a boolean if a field has been set.

### GetRefresh

`func (o *InheritedZoneAuthority) GetRefresh() Inheritance2InheritedUInt32`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *InheritedZoneAuthority) GetRefreshOk() (*Inheritance2InheritedUInt32, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *InheritedZoneAuthority) SetRefresh(v Inheritance2InheritedUInt32)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *InheritedZoneAuthority) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetRetry

`func (o *InheritedZoneAuthority) GetRetry() Inheritance2InheritedUInt32`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *InheritedZoneAuthority) GetRetryOk() (*Inheritance2InheritedUInt32, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *InheritedZoneAuthority) SetRetry(v Inheritance2InheritedUInt32)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *InheritedZoneAuthority) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### GetRname

`func (o *InheritedZoneAuthority) GetRname() Inheritance2InheritedString`

GetRname returns the Rname field if non-nil, zero value otherwise.

### GetRnameOk

`func (o *InheritedZoneAuthority) GetRnameOk() (*Inheritance2InheritedString, bool)`

GetRnameOk returns a tuple with the Rname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRname

`func (o *InheritedZoneAuthority) SetRname(v Inheritance2InheritedString)`

SetRname sets Rname field to given value.

### HasRname

`func (o *InheritedZoneAuthority) HasRname() bool`

HasRname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


