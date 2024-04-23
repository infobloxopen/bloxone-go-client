# AssignTagsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** | The resource identifier. | [optional] 
**Override** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAssignTagsRequest

`func NewAssignTagsRequest() *AssignTagsRequest`

NewAssignTagsRequest instantiates a new AssignTagsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignTagsRequestWithDefaults

`func NewAssignTagsRequestWithDefaults() *AssignTagsRequest`

NewAssignTagsRequestWithDefaults instantiates a new AssignTagsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *AssignTagsRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *AssignTagsRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *AssignTagsRequest) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *AssignTagsRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetOverride

`func (o *AssignTagsRequest) GetOverride() bool`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *AssignTagsRequest) GetOverrideOk() (*bool, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *AssignTagsRequest) SetOverride(v bool)`

SetOverride sets Override field to given value.

### HasOverride

`func (o *AssignTagsRequest) HasOverride() bool`

HasOverride returns a boolean if a field has been set.

### GetTags

`func (o *AssignTagsRequest) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AssignTagsRequest) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AssignTagsRequest) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *AssignTagsRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


