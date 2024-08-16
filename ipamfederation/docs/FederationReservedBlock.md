# FederationReservedBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The address field in form “a.b.c.d/n” where the “/n” may be omitted. In this case, the CIDR value must be defined in the _cidr_ field. When reading, the _address_ field is always in the form “a.b.c.d”. | 
**Cidr** | Pointer to **int64** | The CIDR of the reserved block. This is required field, if _address_ does not specify it in its input. | [optional] 
**Comment** | Pointer to **string** | The description for the reserved block. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**FederatedRealm** | **string** | The resource identifier. | 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the reserved block. May contain 1 to 256 characters. Can include UTF-8. | [optional] 
**Parent** | Pointer to **string** | The resource identifier. | [optional] 
**Protocol** | Pointer to **string** | The type of protocol of reserved block (_ip4_ or _ip6_). | [optional] [readonly] 
**Tags** | Pointer to **map[string]interface{}** | The tags for the reserved block in JSON format. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 

## Methods

### NewFederationReservedBlock

`func NewFederationReservedBlock(address string, federatedRealm string, ) *FederationReservedBlock`

NewFederationReservedBlock instantiates a new FederationReservedBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationReservedBlockWithDefaults

`func NewFederationReservedBlockWithDefaults() *FederationReservedBlock`

NewFederationReservedBlockWithDefaults instantiates a new FederationReservedBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *FederationReservedBlock) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FederationReservedBlock) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FederationReservedBlock) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCidr

`func (o *FederationReservedBlock) GetCidr() int64`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *FederationReservedBlock) GetCidrOk() (*int64, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *FederationReservedBlock) SetCidr(v int64)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *FederationReservedBlock) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetComment

`func (o *FederationReservedBlock) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FederationReservedBlock) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FederationReservedBlock) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FederationReservedBlock) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FederationReservedBlock) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FederationReservedBlock) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FederationReservedBlock) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FederationReservedBlock) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFederatedRealm

`func (o *FederationReservedBlock) GetFederatedRealm() string`

GetFederatedRealm returns the FederatedRealm field if non-nil, zero value otherwise.

### GetFederatedRealmOk

`func (o *FederationReservedBlock) GetFederatedRealmOk() (*string, bool)`

GetFederatedRealmOk returns a tuple with the FederatedRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedRealm

`func (o *FederationReservedBlock) SetFederatedRealm(v string)`

SetFederatedRealm sets FederatedRealm field to given value.


### GetId

`func (o *FederationReservedBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FederationReservedBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FederationReservedBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FederationReservedBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FederationReservedBlock) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FederationReservedBlock) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FederationReservedBlock) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FederationReservedBlock) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *FederationReservedBlock) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *FederationReservedBlock) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *FederationReservedBlock) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *FederationReservedBlock) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProtocol

`func (o *FederationReservedBlock) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FederationReservedBlock) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FederationReservedBlock) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *FederationReservedBlock) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTags

`func (o *FederationReservedBlock) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FederationReservedBlock) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FederationReservedBlock) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *FederationReservedBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FederationReservedBlock) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FederationReservedBlock) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FederationReservedBlock) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FederationReservedBlock) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


