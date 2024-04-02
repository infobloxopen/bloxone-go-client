# ProtoBgpNeighbor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asn** | Pointer to **int64** |  | [optional] 
**AsnText** | Pointer to **string** | Examples:     ASDOT        ASPLAIN     INTEGER     VALID/INVALID     0.1          1           1           Valid     1            1           1           Valid     65535        65535       65535       Valid     0.65535      65535       65535       Valid     1.0          65536       65536       Valid     1.1          65537       65537       Valid     1.65535      131071      131071      Valid     65535.0      4294901760  4294901760  Valid     65535.1      4294901761  4294901761  Valid     65535.65535  4294967295  4294967295  Valid      0.65536                              Invalid     65535.655536                         Invalid     65536.0                              Invalid     65536.65535                          Invalid                  4294967296              Invalid | [optional] 
**IpAddress** | Pointer to **string** | IPv4 address of the BGP neighbor | [optional] 
**MaxHopCount** | Pointer to **int64** |  | [optional] 
**Multihop** | Pointer to **bool** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 

## Methods

### NewProtoBgpNeighbor

`func NewProtoBgpNeighbor() *ProtoBgpNeighbor`

NewProtoBgpNeighbor instantiates a new ProtoBgpNeighbor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtoBgpNeighborWithDefaults

`func NewProtoBgpNeighborWithDefaults() *ProtoBgpNeighbor`

NewProtoBgpNeighborWithDefaults instantiates a new ProtoBgpNeighbor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsn

`func (o *ProtoBgpNeighbor) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *ProtoBgpNeighbor) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *ProtoBgpNeighbor) SetAsn(v int64)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *ProtoBgpNeighbor) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetAsnText

`func (o *ProtoBgpNeighbor) GetAsnText() string`

GetAsnText returns the AsnText field if non-nil, zero value otherwise.

### GetAsnTextOk

`func (o *ProtoBgpNeighbor) GetAsnTextOk() (*string, bool)`

GetAsnTextOk returns a tuple with the AsnText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsnText

`func (o *ProtoBgpNeighbor) SetAsnText(v string)`

SetAsnText sets AsnText field to given value.

### HasAsnText

`func (o *ProtoBgpNeighbor) HasAsnText() bool`

HasAsnText returns a boolean if a field has been set.

### GetIpAddress

`func (o *ProtoBgpNeighbor) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ProtoBgpNeighbor) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ProtoBgpNeighbor) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ProtoBgpNeighbor) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetMaxHopCount

`func (o *ProtoBgpNeighbor) GetMaxHopCount() int64`

GetMaxHopCount returns the MaxHopCount field if non-nil, zero value otherwise.

### GetMaxHopCountOk

`func (o *ProtoBgpNeighbor) GetMaxHopCountOk() (*int64, bool)`

GetMaxHopCountOk returns a tuple with the MaxHopCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHopCount

`func (o *ProtoBgpNeighbor) SetMaxHopCount(v int64)`

SetMaxHopCount sets MaxHopCount field to given value.

### HasMaxHopCount

`func (o *ProtoBgpNeighbor) HasMaxHopCount() bool`

HasMaxHopCount returns a boolean if a field has been set.

### GetMultihop

`func (o *ProtoBgpNeighbor) GetMultihop() bool`

GetMultihop returns the Multihop field if non-nil, zero value otherwise.

### GetMultihopOk

`func (o *ProtoBgpNeighbor) GetMultihopOk() (*bool, bool)`

GetMultihopOk returns a tuple with the Multihop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultihop

`func (o *ProtoBgpNeighbor) SetMultihop(v bool)`

SetMultihop sets Multihop field to given value.

### HasMultihop

`func (o *ProtoBgpNeighbor) HasMultihop() bool`

HasMultihop returns a boolean if a field has been set.

### GetPassword

`func (o *ProtoBgpNeighbor) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ProtoBgpNeighbor) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ProtoBgpNeighbor) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ProtoBgpNeighbor) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


