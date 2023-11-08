# ConfigACL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional. Comment for ACL. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**List** | Pointer to [**[]ConfigACLItem**](ConfigACLItem.md) | Optional. Ordered list of access control elements.  Elements are evaluated in order to determine access. If evaluation reaches the end of the list then access is denied. | [optional] 
**Name** | **string** | ACL object name. | 
**Tags** | Pointer to **map[string]interface{}** | Tagging specifics. | [optional] 

## Methods

### NewConfigACL

`func NewConfigACL(name string, ) *ConfigACL`

NewConfigACL instantiates a new ConfigACL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigACLWithDefaults

`func NewConfigACLWithDefaults() *ConfigACL`

NewConfigACLWithDefaults instantiates a new ConfigACL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ConfigACL) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ConfigACL) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ConfigACL) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ConfigACL) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetId

`func (o *ConfigACL) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigACL) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigACL) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigACL) HasId() bool`

HasId returns a boolean if a field has been set.

### GetList

`func (o *ConfigACL) GetList() []ConfigACLItem`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *ConfigACL) GetListOk() (*[]ConfigACLItem, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *ConfigACL) SetList(v []ConfigACLItem)`

SetList sets List field to given value.

### HasList

`func (o *ConfigACL) HasList() bool`

HasList returns a boolean if a field has been set.

### GetName

`func (o *ConfigACL) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigACL) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigACL) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *ConfigACL) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConfigACL) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConfigACL) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConfigACL) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


