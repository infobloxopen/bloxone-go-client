# AtcfwCategoryFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **[]string** | The list of content category names that falls into this category filter. | [optional] 
**CreatedTime** | Pointer to **time.Time** | The time when this Category Filter object was created. | [optional] [readonly] 
**Description** | Pointer to **string** | The brief description for the category filter. | [optional] 
**Id** | Pointer to **int32** | The Category Filter object identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the category filter. | [optional] 
**Policies** | Pointer to **[]string** | The list of security policy names with which the category filter is associated. | [optional] [readonly] 
**Tags** | Pointer to **map[string]interface{}** | Enables tag support for resource where tags attribute contains user-defined key value pairs | [optional] 
**UpdatedTime** | Pointer to **time.Time** | The time when this Category Filter object was last updated. | [optional] [readonly] 

## Methods

### NewAtcfwCategoryFilter

`func NewAtcfwCategoryFilter() *AtcfwCategoryFilter`

NewAtcfwCategoryFilter instantiates a new AtcfwCategoryFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtcfwCategoryFilterWithDefaults

`func NewAtcfwCategoryFilterWithDefaults() *AtcfwCategoryFilter`

NewAtcfwCategoryFilterWithDefaults instantiates a new AtcfwCategoryFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *AtcfwCategoryFilter) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *AtcfwCategoryFilter) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *AtcfwCategoryFilter) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *AtcfwCategoryFilter) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCreatedTime

`func (o *AtcfwCategoryFilter) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AtcfwCategoryFilter) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AtcfwCategoryFilter) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *AtcfwCategoryFilter) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *AtcfwCategoryFilter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AtcfwCategoryFilter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AtcfwCategoryFilter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AtcfwCategoryFilter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *AtcfwCategoryFilter) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AtcfwCategoryFilter) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AtcfwCategoryFilter) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AtcfwCategoryFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AtcfwCategoryFilter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AtcfwCategoryFilter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AtcfwCategoryFilter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AtcfwCategoryFilter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicies

`func (o *AtcfwCategoryFilter) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *AtcfwCategoryFilter) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *AtcfwCategoryFilter) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *AtcfwCategoryFilter) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetTags

`func (o *AtcfwCategoryFilter) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AtcfwCategoryFilter) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AtcfwCategoryFilter) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *AtcfwCategoryFilter) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *AtcfwCategoryFilter) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *AtcfwCategoryFilter) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *AtcfwCategoryFilter) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *AtcfwCategoryFilter) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


