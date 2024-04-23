# ApplicationFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **time.Time** | The time when this Application Filter object was created. | [optional] [readonly] 
**Criteria** | Pointer to [**[]ApplicationCriterion**](ApplicationCriterion.md) | The array of key-value pairs specifying criteria for the search. | [optional] 
**Description** | Pointer to **string** | The brief description for the application filter. | [optional] 
**Id** | Pointer to **int32** | The Application Filter object identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the application filter. | [optional] 
**Policies** | Pointer to **[]string** | The list of security policy names with which the application filter is associated. | [optional] [readonly] 
**Readonly** | Pointer to **bool** | True if it is a predefined application filter | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Enables tag support for resource where tags attribute contains user-defined key value pairs | [optional] 
**UpdatedTime** | Pointer to **time.Time** | The time when this Application Filter object was last updated. | [optional] [readonly] 

## Methods

### NewApplicationFilter

`func NewApplicationFilter() *ApplicationFilter`

NewApplicationFilter instantiates a new ApplicationFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationFilterWithDefaults

`func NewApplicationFilterWithDefaults() *ApplicationFilter`

NewApplicationFilterWithDefaults instantiates a new ApplicationFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *ApplicationFilter) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ApplicationFilter) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ApplicationFilter) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *ApplicationFilter) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCriteria

`func (o *ApplicationFilter) GetCriteria() []ApplicationCriterion`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *ApplicationFilter) GetCriteriaOk() (*[]ApplicationCriterion, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *ApplicationFilter) SetCriteria(v []ApplicationCriterion)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *ApplicationFilter) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationFilter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationFilter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationFilter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationFilter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ApplicationFilter) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationFilter) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationFilter) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApplicationFilter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationFilter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationFilter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationFilter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicies

`func (o *ApplicationFilter) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ApplicationFilter) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ApplicationFilter) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ApplicationFilter) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetReadonly

`func (o *ApplicationFilter) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *ApplicationFilter) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *ApplicationFilter) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *ApplicationFilter) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetTags

`func (o *ApplicationFilter) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApplicationFilter) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApplicationFilter) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApplicationFilter) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *ApplicationFilter) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *ApplicationFilter) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *ApplicationFilter) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *ApplicationFilter) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


