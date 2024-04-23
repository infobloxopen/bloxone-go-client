# BgpNeighbor

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

### NewBgpNeighbor

`func NewBgpNeighbor() *BgpNeighbor`

NewBgpNeighbor instantiates a new BgpNeighbor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpNeighborWithDefaults

`func NewBgpNeighborWithDefaults() *BgpNeighbor`

NewBgpNeighborWithDefaults instantiates a new BgpNeighbor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsn

`func (o *BgpNeighbor) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *BgpNeighbor) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *BgpNeighbor) SetAsn(v int64)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *BgpNeighbor) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetAsnText

`func (o *BgpNeighbor) GetAsnText() string`

GetAsnText returns the AsnText field if non-nil, zero value otherwise.

### GetAsnTextOk

`func (o *BgpNeighbor) GetAsnTextOk() (*string, bool)`

GetAsnTextOk returns a tuple with the AsnText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsnText

`func (o *BgpNeighbor) SetAsnText(v string)`

SetAsnText sets AsnText field to given value.

### HasAsnText

`func (o *BgpNeighbor) HasAsnText() bool`

HasAsnText returns a boolean if a field has been set.

### GetIpAddress

`func (o *BgpNeighbor) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *BgpNeighbor) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *BgpNeighbor) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *BgpNeighbor) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetMaxHopCount

`func (o *BgpNeighbor) GetMaxHopCount() int64`

GetMaxHopCount returns the MaxHopCount field if non-nil, zero value otherwise.

### GetMaxHopCountOk

`func (o *BgpNeighbor) GetMaxHopCountOk() (*int64, bool)`

GetMaxHopCountOk returns a tuple with the MaxHopCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHopCount

`func (o *BgpNeighbor) SetMaxHopCount(v int64)`

SetMaxHopCount sets MaxHopCount field to given value.

### HasMaxHopCount

`func (o *BgpNeighbor) HasMaxHopCount() bool`

HasMaxHopCount returns a boolean if a field has been set.

### GetMultihop

`func (o *BgpNeighbor) GetMultihop() bool`

GetMultihop returns the Multihop field if non-nil, zero value otherwise.

### GetMultihopOk

`func (o *BgpNeighbor) GetMultihopOk() (*bool, bool)`

GetMultihopOk returns a tuple with the Multihop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultihop

`func (o *BgpNeighbor) SetMultihop(v bool)`

SetMultihop sets Multihop field to given value.

### HasMultihop

`func (o *BgpNeighbor) HasMultihop() bool`

HasMultihop returns a boolean if a field has been set.

### GetPassword

`func (o *BgpNeighbor) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BgpNeighbor) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BgpNeighbor) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BgpNeighbor) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


