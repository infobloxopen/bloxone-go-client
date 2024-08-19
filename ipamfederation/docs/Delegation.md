# Delegation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The address field in form “a.b.c.d/n” where the “/n” may be omitted. In this case, the CIDR value must be defined in the _cidr_ field. When reading, the _address_ field is always in the form “a.b.c.d”. | 
**Cidr** | Pointer to **int64** | The CIDR of the delegation. This is required, if _address_ does not specify it in its input. | [optional] 
**Comment** | Pointer to **string** | The description for the delegation. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**DelegatedTo** | Pointer to **string** | The specific IPAM service the delegation was delegated to. | [optional] 
**FederatedRealms** | **[]string** | The resource identifier. | 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the delegation. May contain 1 to 256 characters. Can include UTF-8. | [optional] 
**Parents** | Pointer to **[]string** | The resource identifier. | [optional] 
**Protocol** | Pointer to **string** | The type of protocol of delegation (_ip4_ or _ip6_). | [optional] [readonly] 
**Tags** | Pointer to **map[string]interface{}** | The tags for the delegation in JSON format. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 

## Methods

### NewDelegation

`func NewDelegation(address string, federatedRealms []string, ) *Delegation`

NewDelegation instantiates a new Delegation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegationWithDefaults

`func NewDelegationWithDefaults() *Delegation`

NewDelegationWithDefaults instantiates a new Delegation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Delegation) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Delegation) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Delegation) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCidr

`func (o *Delegation) GetCidr() int64`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *Delegation) GetCidrOk() (*int64, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *Delegation) SetCidr(v int64)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *Delegation) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetComment

`func (o *Delegation) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Delegation) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Delegation) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Delegation) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Delegation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Delegation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Delegation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Delegation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDelegatedTo

`func (o *Delegation) GetDelegatedTo() string`

GetDelegatedTo returns the DelegatedTo field if non-nil, zero value otherwise.

### GetDelegatedToOk

`func (o *Delegation) GetDelegatedToOk() (*string, bool)`

GetDelegatedToOk returns a tuple with the DelegatedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedTo

`func (o *Delegation) SetDelegatedTo(v string)`

SetDelegatedTo sets DelegatedTo field to given value.

### HasDelegatedTo

`func (o *Delegation) HasDelegatedTo() bool`

HasDelegatedTo returns a boolean if a field has been set.

### GetFederatedRealms

`func (o *Delegation) GetFederatedRealms() []string`

GetFederatedRealms returns the FederatedRealms field if non-nil, zero value otherwise.

### GetFederatedRealmsOk

`func (o *Delegation) GetFederatedRealmsOk() (*[]string, bool)`

GetFederatedRealmsOk returns a tuple with the FederatedRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedRealms

`func (o *Delegation) SetFederatedRealms(v []string)`

SetFederatedRealms sets FederatedRealms field to given value.


### GetId

`func (o *Delegation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Delegation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Delegation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Delegation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Delegation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Delegation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Delegation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Delegation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParents

`func (o *Delegation) GetParents() []string`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *Delegation) GetParentsOk() (*[]string, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *Delegation) SetParents(v []string)`

SetParents sets Parents field to given value.

### HasParents

`func (o *Delegation) HasParents() bool`

HasParents returns a boolean if a field has been set.

### GetProtocol

`func (o *Delegation) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Delegation) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Delegation) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *Delegation) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTags

`func (o *Delegation) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Delegation) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Delegation) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Delegation) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Delegation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Delegation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Delegation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Delegation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


