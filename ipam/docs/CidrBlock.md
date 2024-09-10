# CidrBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The address part of the CidrBlock. | [optional] [readonly] 
**Cidr** | Pointer to **int64** | The CIDR part of the CidrBlock. | [optional] [readonly] 
**FederatedRealms** | Pointer to **[]string** | Reserved for future use. | [optional] [readonly] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Parent** | Pointer to **string** | The resource identifier. | [optional] 
**Space** | Pointer to **string** | The resource identifier. | [optional] 

## Methods

### NewCidrBlock

`func NewCidrBlock() *CidrBlock`

NewCidrBlock instantiates a new CidrBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCidrBlockWithDefaults

`func NewCidrBlockWithDefaults() *CidrBlock`

NewCidrBlockWithDefaults instantiates a new CidrBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CidrBlock) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CidrBlock) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CidrBlock) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CidrBlock) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCidr

`func (o *CidrBlock) GetCidr() int64`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *CidrBlock) GetCidrOk() (*int64, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *CidrBlock) SetCidr(v int64)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *CidrBlock) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetFederatedRealms

`func (o *CidrBlock) GetFederatedRealms() []string`

GetFederatedRealms returns the FederatedRealms field if non-nil, zero value otherwise.

### GetFederatedRealmsOk

`func (o *CidrBlock) GetFederatedRealmsOk() (*[]string, bool)`

GetFederatedRealmsOk returns a tuple with the FederatedRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedRealms

`func (o *CidrBlock) SetFederatedRealms(v []string)`

SetFederatedRealms sets FederatedRealms field to given value.

### HasFederatedRealms

`func (o *CidrBlock) HasFederatedRealms() bool`

HasFederatedRealms returns a boolean if a field has been set.

### GetId

`func (o *CidrBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CidrBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CidrBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CidrBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParent

`func (o *CidrBlock) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *CidrBlock) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *CidrBlock) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *CidrBlock) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetSpace

`func (o *CidrBlock) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *CidrBlock) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *CidrBlock) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *CidrBlock) HasSpace() bool`

HasSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


