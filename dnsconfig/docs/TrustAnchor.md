# TrustAnchor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | **int64** |  | 
**ProtocolZone** | Pointer to **string** | Zone FQDN in punycode. | [optional] [readonly] 
**PublicKey** | **string** | DNSSEC key data. Non-empty, valid base64 string. | 
**Sep** | Pointer to **bool** | Optional. Secure Entry Point flag.  Defaults to _true_. | [optional] 
**Zone** | **string** | Zone FQDN. | 

## Methods

### NewTrustAnchor

`func NewTrustAnchor(algorithm int64, publicKey string, zone string, ) *TrustAnchor`

NewTrustAnchor instantiates a new TrustAnchor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustAnchorWithDefaults

`func NewTrustAnchorWithDefaults() *TrustAnchor`

NewTrustAnchorWithDefaults instantiates a new TrustAnchor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *TrustAnchor) GetAlgorithm() int64`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *TrustAnchor) GetAlgorithmOk() (*int64, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *TrustAnchor) SetAlgorithm(v int64)`

SetAlgorithm sets Algorithm field to given value.


### GetProtocolZone

`func (o *TrustAnchor) GetProtocolZone() string`

GetProtocolZone returns the ProtocolZone field if non-nil, zero value otherwise.

### GetProtocolZoneOk

`func (o *TrustAnchor) GetProtocolZoneOk() (*string, bool)`

GetProtocolZoneOk returns a tuple with the ProtocolZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolZone

`func (o *TrustAnchor) SetProtocolZone(v string)`

SetProtocolZone sets ProtocolZone field to given value.

### HasProtocolZone

`func (o *TrustAnchor) HasProtocolZone() bool`

HasProtocolZone returns a boolean if a field has been set.

### GetPublicKey

`func (o *TrustAnchor) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *TrustAnchor) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *TrustAnchor) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetSep

`func (o *TrustAnchor) GetSep() bool`

GetSep returns the Sep field if non-nil, zero value otherwise.

### GetSepOk

`func (o *TrustAnchor) GetSepOk() (*bool, bool)`

GetSepOk returns a tuple with the Sep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSep

`func (o *TrustAnchor) SetSep(v bool)`

SetSep sets Sep field to given value.

### HasSep

`func (o *TrustAnchor) HasSep() bool`

HasSep returns a boolean if a field has been set.

### GetZone

`func (o *TrustAnchor) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *TrustAnchor) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *TrustAnchor) SetZone(v string)`

SetZone sets Zone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


