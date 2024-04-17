# ProtoBgpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asn** | Pointer to **int64** |  | [optional] 
**AsnText** | Pointer to **string** | Examples:     ASDOT        ASPLAIN     INTEGER     VALID/INVALID     0.1          1           1           Valid     1            1           1           Valid     65535        65535       65535       Valid     0.65535      65535       65535       Valid     1.0          65536       65536       Valid     1.1          65537       65537       Valid     1.65535      131071      131071      Valid     65535.0      4294901760  4294901760  Valid     65535.1      4294901761  4294901761  Valid     65535.65535  4294967295  4294967295  Valid      0.65536                              Invalid     65535.655536                         Invalid     65536.0                              Invalid     65536.65535                          Invalid                  4294967296              Invalid | [optional] 
**Fields** | Pointer to [**ProtobufFieldMask**](ProtobufFieldMask.md) |  | [optional] 
**HolddownSecs** | Pointer to **int64** |  | [optional] 
**KeepAliveSecs** | Pointer to **int64** |  | [optional] 
**LinkDetect** | Pointer to **bool** |  | [optional] 
**Neighbors** | Pointer to [**[]ProtoBgpNeighbor**](ProtoBgpNeighbor.md) |  | [optional] 
**Preamble** | Pointer to **string** | Any predefined BGP configuration, with embedded new lines; the preamble will be prepended to the generated BGP configuration. | [optional] 

## Methods

### NewProtoBgpConfig

`func NewProtoBgpConfig() *ProtoBgpConfig`

NewProtoBgpConfig instantiates a new ProtoBgpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtoBgpConfigWithDefaults

`func NewProtoBgpConfigWithDefaults() *ProtoBgpConfig`

NewProtoBgpConfigWithDefaults instantiates a new ProtoBgpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsn

`func (o *ProtoBgpConfig) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *ProtoBgpConfig) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *ProtoBgpConfig) SetAsn(v int64)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *ProtoBgpConfig) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetAsnText

`func (o *ProtoBgpConfig) GetAsnText() string`

GetAsnText returns the AsnText field if non-nil, zero value otherwise.

### GetAsnTextOk

`func (o *ProtoBgpConfig) GetAsnTextOk() (*string, bool)`

GetAsnTextOk returns a tuple with the AsnText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsnText

`func (o *ProtoBgpConfig) SetAsnText(v string)`

SetAsnText sets AsnText field to given value.

### HasAsnText

`func (o *ProtoBgpConfig) HasAsnText() bool`

HasAsnText returns a boolean if a field has been set.

### GetFields

`func (o *ProtoBgpConfig) GetFields() ProtobufFieldMask`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ProtoBgpConfig) GetFieldsOk() (*ProtobufFieldMask, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ProtoBgpConfig) SetFields(v ProtobufFieldMask)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ProtoBgpConfig) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetHolddownSecs

`func (o *ProtoBgpConfig) GetHolddownSecs() int64`

GetHolddownSecs returns the HolddownSecs field if non-nil, zero value otherwise.

### GetHolddownSecsOk

`func (o *ProtoBgpConfig) GetHolddownSecsOk() (*int64, bool)`

GetHolddownSecsOk returns a tuple with the HolddownSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolddownSecs

`func (o *ProtoBgpConfig) SetHolddownSecs(v int64)`

SetHolddownSecs sets HolddownSecs field to given value.

### HasHolddownSecs

`func (o *ProtoBgpConfig) HasHolddownSecs() bool`

HasHolddownSecs returns a boolean if a field has been set.

### GetKeepAliveSecs

`func (o *ProtoBgpConfig) GetKeepAliveSecs() int64`

GetKeepAliveSecs returns the KeepAliveSecs field if non-nil, zero value otherwise.

### GetKeepAliveSecsOk

`func (o *ProtoBgpConfig) GetKeepAliveSecsOk() (*int64, bool)`

GetKeepAliveSecsOk returns a tuple with the KeepAliveSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAliveSecs

`func (o *ProtoBgpConfig) SetKeepAliveSecs(v int64)`

SetKeepAliveSecs sets KeepAliveSecs field to given value.

### HasKeepAliveSecs

`func (o *ProtoBgpConfig) HasKeepAliveSecs() bool`

HasKeepAliveSecs returns a boolean if a field has been set.

### GetLinkDetect

`func (o *ProtoBgpConfig) GetLinkDetect() bool`

GetLinkDetect returns the LinkDetect field if non-nil, zero value otherwise.

### GetLinkDetectOk

`func (o *ProtoBgpConfig) GetLinkDetectOk() (*bool, bool)`

GetLinkDetectOk returns a tuple with the LinkDetect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDetect

`func (o *ProtoBgpConfig) SetLinkDetect(v bool)`

SetLinkDetect sets LinkDetect field to given value.

### HasLinkDetect

`func (o *ProtoBgpConfig) HasLinkDetect() bool`

HasLinkDetect returns a boolean if a field has been set.

### GetNeighbors

`func (o *ProtoBgpConfig) GetNeighbors() []ProtoBgpNeighbor`

GetNeighbors returns the Neighbors field if non-nil, zero value otherwise.

### GetNeighborsOk

`func (o *ProtoBgpConfig) GetNeighborsOk() (*[]ProtoBgpNeighbor, bool)`

GetNeighborsOk returns a tuple with the Neighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighbors

`func (o *ProtoBgpConfig) SetNeighbors(v []ProtoBgpNeighbor)`

SetNeighbors sets Neighbors field to given value.

### HasNeighbors

`func (o *ProtoBgpConfig) HasNeighbors() bool`

HasNeighbors returns a boolean if a field has been set.

### GetPreamble

`func (o *ProtoBgpConfig) GetPreamble() string`

GetPreamble returns the Preamble field if non-nil, zero value otherwise.

### GetPreambleOk

`func (o *ProtoBgpConfig) GetPreambleOk() (*string, bool)`

GetPreambleOk returns a tuple with the Preamble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreamble

`func (o *ProtoBgpConfig) SetPreamble(v string)`

SetPreamble sets Preamble field to given value.

### HasPreamble

`func (o *ProtoBgpConfig) HasPreamble() bool`

HasPreamble returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


