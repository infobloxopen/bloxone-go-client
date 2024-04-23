# ZoneAuthority

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultTtl** | Pointer to **int64** | Optional. ZoneAuthority default ttl for resource records in zone (value in seconds).  Defaults to 28800. | [optional] 
**Expire** | Pointer to **int64** | Optional. ZoneAuthority expire time in seconds.  Defaults to 2419200. | [optional] 
**Mname** | Pointer to **string** | Defaults to empty. | [optional] 
**NegativeTtl** | Pointer to **int64** | Optional. ZoneAuthority negative caching (minimum) ttl in seconds.  Defaults to 900. | [optional] 
**ProtocolMname** | Pointer to **string** | Optional. ZoneAuthority master name server in punycode.  Defaults to empty. | [optional] [readonly] 
**ProtocolRname** | Pointer to **string** | Optional. A domain name which specifies the mailbox of the person responsible for this zone.  Defaults to empty. | [optional] [readonly] 
**Refresh** | Pointer to **int64** | Optional. ZoneAuthority refresh.  Defaults to 10800. | [optional] 
**Retry** | Pointer to **int64** | Optional. ZoneAuthority retry.  Defaults to 3600. | [optional] 
**Rname** | Pointer to **string** | Optional. ZoneAuthority rname.  Defaults to empty. | [optional] 
**UseDefaultMname** | Pointer to **bool** | Optional. Use default value for master name server.  Defaults to true. | [optional] 

## Methods

### NewZoneAuthority

`func NewZoneAuthority() *ZoneAuthority`

NewZoneAuthority instantiates a new ZoneAuthority object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneAuthorityWithDefaults

`func NewZoneAuthorityWithDefaults() *ZoneAuthority`

NewZoneAuthorityWithDefaults instantiates a new ZoneAuthority object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultTtl

`func (o *ZoneAuthority) GetDefaultTtl() int64`

GetDefaultTtl returns the DefaultTtl field if non-nil, zero value otherwise.

### GetDefaultTtlOk

`func (o *ZoneAuthority) GetDefaultTtlOk() (*int64, bool)`

GetDefaultTtlOk returns a tuple with the DefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTtl

`func (o *ZoneAuthority) SetDefaultTtl(v int64)`

SetDefaultTtl sets DefaultTtl field to given value.

### HasDefaultTtl

`func (o *ZoneAuthority) HasDefaultTtl() bool`

HasDefaultTtl returns a boolean if a field has been set.

### GetExpire

`func (o *ZoneAuthority) GetExpire() int64`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *ZoneAuthority) GetExpireOk() (*int64, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *ZoneAuthority) SetExpire(v int64)`

SetExpire sets Expire field to given value.

### HasExpire

`func (o *ZoneAuthority) HasExpire() bool`

HasExpire returns a boolean if a field has been set.

### GetMname

`func (o *ZoneAuthority) GetMname() string`

GetMname returns the Mname field if non-nil, zero value otherwise.

### GetMnameOk

`func (o *ZoneAuthority) GetMnameOk() (*string, bool)`

GetMnameOk returns a tuple with the Mname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMname

`func (o *ZoneAuthority) SetMname(v string)`

SetMname sets Mname field to given value.

### HasMname

`func (o *ZoneAuthority) HasMname() bool`

HasMname returns a boolean if a field has been set.

### GetNegativeTtl

`func (o *ZoneAuthority) GetNegativeTtl() int64`

GetNegativeTtl returns the NegativeTtl field if non-nil, zero value otherwise.

### GetNegativeTtlOk

`func (o *ZoneAuthority) GetNegativeTtlOk() (*int64, bool)`

GetNegativeTtlOk returns a tuple with the NegativeTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeTtl

`func (o *ZoneAuthority) SetNegativeTtl(v int64)`

SetNegativeTtl sets NegativeTtl field to given value.

### HasNegativeTtl

`func (o *ZoneAuthority) HasNegativeTtl() bool`

HasNegativeTtl returns a boolean if a field has been set.

### GetProtocolMname

`func (o *ZoneAuthority) GetProtocolMname() string`

GetProtocolMname returns the ProtocolMname field if non-nil, zero value otherwise.

### GetProtocolMnameOk

`func (o *ZoneAuthority) GetProtocolMnameOk() (*string, bool)`

GetProtocolMnameOk returns a tuple with the ProtocolMname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolMname

`func (o *ZoneAuthority) SetProtocolMname(v string)`

SetProtocolMname sets ProtocolMname field to given value.

### HasProtocolMname

`func (o *ZoneAuthority) HasProtocolMname() bool`

HasProtocolMname returns a boolean if a field has been set.

### GetProtocolRname

`func (o *ZoneAuthority) GetProtocolRname() string`

GetProtocolRname returns the ProtocolRname field if non-nil, zero value otherwise.

### GetProtocolRnameOk

`func (o *ZoneAuthority) GetProtocolRnameOk() (*string, bool)`

GetProtocolRnameOk returns a tuple with the ProtocolRname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolRname

`func (o *ZoneAuthority) SetProtocolRname(v string)`

SetProtocolRname sets ProtocolRname field to given value.

### HasProtocolRname

`func (o *ZoneAuthority) HasProtocolRname() bool`

HasProtocolRname returns a boolean if a field has been set.

### GetRefresh

`func (o *ZoneAuthority) GetRefresh() int64`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *ZoneAuthority) GetRefreshOk() (*int64, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *ZoneAuthority) SetRefresh(v int64)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *ZoneAuthority) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetRetry

`func (o *ZoneAuthority) GetRetry() int64`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *ZoneAuthority) GetRetryOk() (*int64, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *ZoneAuthority) SetRetry(v int64)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *ZoneAuthority) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### GetRname

`func (o *ZoneAuthority) GetRname() string`

GetRname returns the Rname field if non-nil, zero value otherwise.

### GetRnameOk

`func (o *ZoneAuthority) GetRnameOk() (*string, bool)`

GetRnameOk returns a tuple with the Rname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRname

`func (o *ZoneAuthority) SetRname(v string)`

SetRname sets Rname field to given value.

### HasRname

`func (o *ZoneAuthority) HasRname() bool`

HasRname returns a boolean if a field has been set.

### GetUseDefaultMname

`func (o *ZoneAuthority) GetUseDefaultMname() bool`

GetUseDefaultMname returns the UseDefaultMname field if non-nil, zero value otherwise.

### GetUseDefaultMnameOk

`func (o *ZoneAuthority) GetUseDefaultMnameOk() (*bool, bool)`

GetUseDefaultMnameOk returns a tuple with the UseDefaultMname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefaultMname

`func (o *ZoneAuthority) SetUseDefaultMname(v bool)`

SetUseDefaultMname sets UseDefaultMname field to given value.

### HasUseDefaultMname

`func (o *ZoneAuthority) HasUseDefaultMname() bool`

HasUseDefaultMname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


