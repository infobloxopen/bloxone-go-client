# ACL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional. Comment for ACL. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**List** | Pointer to [**[]ACLItem**](ACLItem.md) | Optional. Ordered list of access control elements.  Elements are evaluated in order to determine access. If evaluation reaches the end of the list then access is denied. | [optional] 
**Name** | **string** | ACL object name. | 
**Tags** | Pointer to **map[string]interface{}** | Tagging specifics. | [optional] 

## Methods

### NewACL

`func NewACL(name string, ) *ACL`

NewACL instantiates a new ACL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACLWithDefaults

`func NewACLWithDefaults() *ACL`

NewACLWithDefaults instantiates a new ACL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ACL) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ACL) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ACL) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ACL) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetId

`func (o *ACL) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ACL) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ACL) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ACL) HasId() bool`

HasId returns a boolean if a field has been set.

### GetList

`func (o *ACL) GetList() []ACLItem`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *ACL) GetListOk() (*[]ACLItem, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *ACL) SetList(v []ACLItem)`

SetList sets List field to given value.

### HasList

`func (o *ACL) HasList() bool`

HasList returns a boolean if a field has been set.

### GetName

`func (o *ACL) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ACL) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ACL) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *ACL) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ACL) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ACL) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ACL) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


