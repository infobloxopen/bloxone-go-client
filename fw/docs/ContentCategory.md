# ContentCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryCode** | Pointer to **int32** | The category code. | [optional] 
**CategoryName** | Pointer to **string** | The name of the category. | [optional] 
**FunctionalGroup** | Pointer to **string** | The functional group name of the category. | [optional] 

## Methods

### NewContentCategory

`func NewContentCategory() *ContentCategory`

NewContentCategory instantiates a new ContentCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentCategoryWithDefaults

`func NewContentCategoryWithDefaults() *ContentCategory`

NewContentCategoryWithDefaults instantiates a new ContentCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryCode

`func (o *ContentCategory) GetCategoryCode() int32`

GetCategoryCode returns the CategoryCode field if non-nil, zero value otherwise.

### GetCategoryCodeOk

`func (o *ContentCategory) GetCategoryCodeOk() (*int32, bool)`

GetCategoryCodeOk returns a tuple with the CategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryCode

`func (o *ContentCategory) SetCategoryCode(v int32)`

SetCategoryCode sets CategoryCode field to given value.

### HasCategoryCode

`func (o *ContentCategory) HasCategoryCode() bool`

HasCategoryCode returns a boolean if a field has been set.

### GetCategoryName

`func (o *ContentCategory) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *ContentCategory) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *ContentCategory) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.

### HasCategoryName

`func (o *ContentCategory) HasCategoryName() bool`

HasCategoryName returns a boolean if a field has been set.

### GetFunctionalGroup

`func (o *ContentCategory) GetFunctionalGroup() string`

GetFunctionalGroup returns the FunctionalGroup field if non-nil, zero value otherwise.

### GetFunctionalGroupOk

`func (o *ContentCategory) GetFunctionalGroupOk() (*string, bool)`

GetFunctionalGroupOk returns a tuple with the FunctionalGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionalGroup

`func (o *ContentCategory) SetFunctionalGroup(v string)`

SetFunctionalGroup sets FunctionalGroup field to given value.

### HasFunctionalGroup

`func (o *ContentCategory) HasFunctionalGroup() bool`

HasFunctionalGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


