# FederationFederatedRealm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocationV4** | Pointer to [**FederationAllocation**](FederationAllocation.md) | The aggregate of all Federated Blocks within the Realm. | [optional] 
**Comment** | Pointer to **string** | The description of the federated realm. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Name** | **string** | The name of the federated realm. May contain 1 to 256 characters; can include UTF-8. | 
**Tags** | Pointer to **map[string]interface{}** | The tags for the federated realm in JSON format. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 

## Methods

### NewFederationFederatedRealm

`func NewFederationFederatedRealm(name string, ) *FederationFederatedRealm`

NewFederationFederatedRealm instantiates a new FederationFederatedRealm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationFederatedRealmWithDefaults

`func NewFederationFederatedRealmWithDefaults() *FederationFederatedRealm`

NewFederationFederatedRealmWithDefaults instantiates a new FederationFederatedRealm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocationV4

`func (o *FederationFederatedRealm) GetAllocationV4() FederationAllocation`

GetAllocationV4 returns the AllocationV4 field if non-nil, zero value otherwise.

### GetAllocationV4Ok

`func (o *FederationFederatedRealm) GetAllocationV4Ok() (*FederationAllocation, bool)`

GetAllocationV4Ok returns a tuple with the AllocationV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationV4

`func (o *FederationFederatedRealm) SetAllocationV4(v FederationAllocation)`

SetAllocationV4 sets AllocationV4 field to given value.

### HasAllocationV4

`func (o *FederationFederatedRealm) HasAllocationV4() bool`

HasAllocationV4 returns a boolean if a field has been set.

### GetComment

`func (o *FederationFederatedRealm) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FederationFederatedRealm) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FederationFederatedRealm) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FederationFederatedRealm) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FederationFederatedRealm) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FederationFederatedRealm) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FederationFederatedRealm) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FederationFederatedRealm) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *FederationFederatedRealm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FederationFederatedRealm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FederationFederatedRealm) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FederationFederatedRealm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FederationFederatedRealm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FederationFederatedRealm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FederationFederatedRealm) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *FederationFederatedRealm) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FederationFederatedRealm) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FederationFederatedRealm) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *FederationFederatedRealm) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FederationFederatedRealm) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FederationFederatedRealm) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FederationFederatedRealm) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FederationFederatedRealm) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


