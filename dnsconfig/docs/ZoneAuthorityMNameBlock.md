# ZoneAuthorityMNameBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mname** | Pointer to **string** | Defaults to empty. | [optional] 
**ProtocolMname** | Pointer to **string** | Optional. Master name server in punycode.  Defaults to empty. | [optional] [readonly] 
**UseDefaultMname** | Pointer to **bool** | Optional. Use default value for master name server.  Defaults to true. | [optional] 

## Methods

### NewZoneAuthorityMNameBlock

`func NewZoneAuthorityMNameBlock() *ZoneAuthorityMNameBlock`

NewZoneAuthorityMNameBlock instantiates a new ZoneAuthorityMNameBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneAuthorityMNameBlockWithDefaults

`func NewZoneAuthorityMNameBlockWithDefaults() *ZoneAuthorityMNameBlock`

NewZoneAuthorityMNameBlockWithDefaults instantiates a new ZoneAuthorityMNameBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMname

`func (o *ZoneAuthorityMNameBlock) GetMname() string`

GetMname returns the Mname field if non-nil, zero value otherwise.

### GetMnameOk

`func (o *ZoneAuthorityMNameBlock) GetMnameOk() (*string, bool)`

GetMnameOk returns a tuple with the Mname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMname

`func (o *ZoneAuthorityMNameBlock) SetMname(v string)`

SetMname sets Mname field to given value.

### HasMname

`func (o *ZoneAuthorityMNameBlock) HasMname() bool`

HasMname returns a boolean if a field has been set.

### GetProtocolMname

`func (o *ZoneAuthorityMNameBlock) GetProtocolMname() string`

GetProtocolMname returns the ProtocolMname field if non-nil, zero value otherwise.

### GetProtocolMnameOk

`func (o *ZoneAuthorityMNameBlock) GetProtocolMnameOk() (*string, bool)`

GetProtocolMnameOk returns a tuple with the ProtocolMname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolMname

`func (o *ZoneAuthorityMNameBlock) SetProtocolMname(v string)`

SetProtocolMname sets ProtocolMname field to given value.

### HasProtocolMname

`func (o *ZoneAuthorityMNameBlock) HasProtocolMname() bool`

HasProtocolMname returns a boolean if a field has been set.

### GetUseDefaultMname

`func (o *ZoneAuthorityMNameBlock) GetUseDefaultMname() bool`

GetUseDefaultMname returns the UseDefaultMname field if non-nil, zero value otherwise.

### GetUseDefaultMnameOk

`func (o *ZoneAuthorityMNameBlock) GetUseDefaultMnameOk() (*bool, bool)`

GetUseDefaultMnameOk returns a tuple with the UseDefaultMname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefaultMname

`func (o *ZoneAuthorityMNameBlock) SetUseDefaultMname(v bool)`

SetUseDefaultMname sets UseDefaultMname field to given value.

### HasUseDefaultMname

`func (o *ZoneAuthorityMNameBlock) HasUseDefaultMname() bool`

HasUseDefaultMname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


