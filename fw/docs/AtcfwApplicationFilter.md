# AtcfwApplicationFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **time.Time** | The time when this Application Filter object was created. | [optional] [readonly] 
**Criteria** | Pointer to [**[]AtcfwApplicationCriterion**](AtcfwApplicationCriterion.md) | The array of key-value pairs specifying criteria for the search. | [optional] 
**Description** | Pointer to **string** | The brief description for the application filter. | [optional] 
**Id** | Pointer to **int32** | The Application Filter object identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the application filter. | [optional] 
**Policies** | Pointer to **[]string** | The list of security policy names with which the application filter is associated. | [optional] [readonly] 
**Readonly** | Pointer to **bool** | True if it is a predefined application filter | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Enables tag support for resource where tags attribute contains user-defined key value pairs | [optional] 
**UpdatedTime** | Pointer to **time.Time** | The time when this Application Filter object was last updated. | [optional] [readonly] 

## Methods

### NewAtcfwApplicationFilter

`func NewAtcfwApplicationFilter() *AtcfwApplicationFilter`

NewAtcfwApplicationFilter instantiates a new AtcfwApplicationFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtcfwApplicationFilterWithDefaults

`func NewAtcfwApplicationFilterWithDefaults() *AtcfwApplicationFilter`

NewAtcfwApplicationFilterWithDefaults instantiates a new AtcfwApplicationFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *AtcfwApplicationFilter) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AtcfwApplicationFilter) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AtcfwApplicationFilter) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *AtcfwApplicationFilter) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCriteria

`func (o *AtcfwApplicationFilter) GetCriteria() []AtcfwApplicationCriterion`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *AtcfwApplicationFilter) GetCriteriaOk() (*[]AtcfwApplicationCriterion, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *AtcfwApplicationFilter) SetCriteria(v []AtcfwApplicationCriterion)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *AtcfwApplicationFilter) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetDescription

`func (o *AtcfwApplicationFilter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AtcfwApplicationFilter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AtcfwApplicationFilter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AtcfwApplicationFilter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *AtcfwApplicationFilter) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AtcfwApplicationFilter) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AtcfwApplicationFilter) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AtcfwApplicationFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AtcfwApplicationFilter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AtcfwApplicationFilter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AtcfwApplicationFilter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AtcfwApplicationFilter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicies

`func (o *AtcfwApplicationFilter) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *AtcfwApplicationFilter) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *AtcfwApplicationFilter) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *AtcfwApplicationFilter) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetReadonly

`func (o *AtcfwApplicationFilter) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *AtcfwApplicationFilter) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *AtcfwApplicationFilter) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *AtcfwApplicationFilter) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetTags

`func (o *AtcfwApplicationFilter) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AtcfwApplicationFilter) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AtcfwApplicationFilter) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *AtcfwApplicationFilter) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *AtcfwApplicationFilter) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *AtcfwApplicationFilter) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *AtcfwApplicationFilter) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *AtcfwApplicationFilter) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


