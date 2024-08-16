# FederationFederatedBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The address field in form “a.b.c.d/n” where the “/n” may be omitted. In this case, the CIDR value must be defined in the _cidr_ field. When reading, the _address_ field is always in the form “a.b.c.d”. | 
**AllocationV4** | Pointer to [**FederationAllocation**](FederationAllocation.md) | The percentage of the Federated Block’s total address space that is consumed by Leaf Terminals. | [optional] 
**Cidr** | Pointer to **int64** | The CIDR of the federated block. This is required, if _address_ does not specify it in its input. | [optional] 
**Comment** | Pointer to **string** | The description for the federated block. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**FederatedRealm** | **string** | The resource identifier. | 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the federated block. May contain 1 to 256 characters. Can include UTF-8. | [optional] 
**Parent** | Pointer to **string** | The resource identifier. | [optional] 
**Protocol** | Pointer to **string** | The type of protocol of federated block (_ip4_ or _ip6_). | [optional] [readonly] 
**Tags** | Pointer to **map[string]interface{}** | The tags for the federated block in JSON format. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 

## Methods

### NewFederationFederatedBlock

`func NewFederationFederatedBlock(address string, federatedRealm string, ) *FederationFederatedBlock`

NewFederationFederatedBlock instantiates a new FederationFederatedBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationFederatedBlockWithDefaults

`func NewFederationFederatedBlockWithDefaults() *FederationFederatedBlock`

NewFederationFederatedBlockWithDefaults instantiates a new FederationFederatedBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *FederationFederatedBlock) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FederationFederatedBlock) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FederationFederatedBlock) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAllocationV4

`func (o *FederationFederatedBlock) GetAllocationV4() FederationAllocation`

GetAllocationV4 returns the AllocationV4 field if non-nil, zero value otherwise.

### GetAllocationV4Ok

`func (o *FederationFederatedBlock) GetAllocationV4Ok() (*FederationAllocation, bool)`

GetAllocationV4Ok returns a tuple with the AllocationV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationV4

`func (o *FederationFederatedBlock) SetAllocationV4(v FederationAllocation)`

SetAllocationV4 sets AllocationV4 field to given value.

### HasAllocationV4

`func (o *FederationFederatedBlock) HasAllocationV4() bool`

HasAllocationV4 returns a boolean if a field has been set.

### GetCidr

`func (o *FederationFederatedBlock) GetCidr() int64`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *FederationFederatedBlock) GetCidrOk() (*int64, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *FederationFederatedBlock) SetCidr(v int64)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *FederationFederatedBlock) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetComment

`func (o *FederationFederatedBlock) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FederationFederatedBlock) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FederationFederatedBlock) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FederationFederatedBlock) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FederationFederatedBlock) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FederationFederatedBlock) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FederationFederatedBlock) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FederationFederatedBlock) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFederatedRealm

`func (o *FederationFederatedBlock) GetFederatedRealm() string`

GetFederatedRealm returns the FederatedRealm field if non-nil, zero value otherwise.

### GetFederatedRealmOk

`func (o *FederationFederatedBlock) GetFederatedRealmOk() (*string, bool)`

GetFederatedRealmOk returns a tuple with the FederatedRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedRealm

`func (o *FederationFederatedBlock) SetFederatedRealm(v string)`

SetFederatedRealm sets FederatedRealm field to given value.


### GetId

`func (o *FederationFederatedBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FederationFederatedBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FederationFederatedBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FederationFederatedBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FederationFederatedBlock) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FederationFederatedBlock) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FederationFederatedBlock) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FederationFederatedBlock) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *FederationFederatedBlock) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *FederationFederatedBlock) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *FederationFederatedBlock) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *FederationFederatedBlock) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProtocol

`func (o *FederationFederatedBlock) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FederationFederatedBlock) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FederationFederatedBlock) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *FederationFederatedBlock) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTags

`func (o *FederationFederatedBlock) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FederationFederatedBlock) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FederationFederatedBlock) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *FederationFederatedBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FederationFederatedBlock) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FederationFederatedBlock) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FederationFederatedBlock) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FederationFederatedBlock) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


