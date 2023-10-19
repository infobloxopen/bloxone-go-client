# IpamsvcOptionSpace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | The description for the option space. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Name** | **string** | The name of the option space. Must contain 1 to 256 characters. Can include UTF-8. | 
**Protocol** | Pointer to **string** | The type of protocol for the option space (_ip4_ or _ip6_). | [optional] 
**Tags** | Pointer to **map[string]interface{}** | The tags for the option space in JSON format. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 

## Methods

### NewIpamsvcOptionSpace

`func NewIpamsvcOptionSpace(name string, ) *IpamsvcOptionSpace`

NewIpamsvcOptionSpace instantiates a new IpamsvcOptionSpace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcOptionSpaceWithDefaults

`func NewIpamsvcOptionSpaceWithDefaults() *IpamsvcOptionSpace`

NewIpamsvcOptionSpaceWithDefaults instantiates a new IpamsvcOptionSpace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *IpamsvcOptionSpace) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IpamsvcOptionSpace) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IpamsvcOptionSpace) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IpamsvcOptionSpace) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IpamsvcOptionSpace) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IpamsvcOptionSpace) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IpamsvcOptionSpace) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IpamsvcOptionSpace) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *IpamsvcOptionSpace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpamsvcOptionSpace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpamsvcOptionSpace) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpamsvcOptionSpace) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IpamsvcOptionSpace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpamsvcOptionSpace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpamsvcOptionSpace) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *IpamsvcOptionSpace) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IpamsvcOptionSpace) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IpamsvcOptionSpace) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IpamsvcOptionSpace) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTags

`func (o *IpamsvcOptionSpace) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IpamsvcOptionSpace) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IpamsvcOptionSpace) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *IpamsvcOptionSpace) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IpamsvcOptionSpace) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IpamsvcOptionSpace) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IpamsvcOptionSpace) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IpamsvcOptionSpace) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


