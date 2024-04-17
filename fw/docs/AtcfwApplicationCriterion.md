# AtcfwApplicationCriterion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** | Name for the application. Since the name of application is unique it may be used as alternate key for the application. The &#39;name&#39; is used for import-export workflow and should be resolved to the &#39;id&#39; before continue processing Create/Update operations. | [optional] 
**Subcategory** | Pointer to **string** |  | [optional] 

## Methods

### NewAtcfwApplicationCriterion

`func NewAtcfwApplicationCriterion() *AtcfwApplicationCriterion`

NewAtcfwApplicationCriterion instantiates a new AtcfwApplicationCriterion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtcfwApplicationCriterionWithDefaults

`func NewAtcfwApplicationCriterionWithDefaults() *AtcfwApplicationCriterion`

NewAtcfwApplicationCriterionWithDefaults instantiates a new AtcfwApplicationCriterion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *AtcfwApplicationCriterion) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AtcfwApplicationCriterion) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AtcfwApplicationCriterion) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AtcfwApplicationCriterion) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetId

`func (o *AtcfwApplicationCriterion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AtcfwApplicationCriterion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AtcfwApplicationCriterion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AtcfwApplicationCriterion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AtcfwApplicationCriterion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AtcfwApplicationCriterion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AtcfwApplicationCriterion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AtcfwApplicationCriterion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubcategory

`func (o *AtcfwApplicationCriterion) GetSubcategory() string`

GetSubcategory returns the Subcategory field if non-nil, zero value otherwise.

### GetSubcategoryOk

`func (o *AtcfwApplicationCriterion) GetSubcategoryOk() (*string, bool)`

GetSubcategoryOk returns a tuple with the Subcategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcategory

`func (o *AtcfwApplicationCriterion) SetSubcategory(v string)`

SetSubcategory sets Subcategory field to given value.

### HasSubcategory

`func (o *AtcfwApplicationCriterion) HasSubcategory() bool`

HasSubcategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


