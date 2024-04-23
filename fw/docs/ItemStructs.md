# ItemStructs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the item | [optional] 
**Item** | Pointer to **string** | The data of the Item | [optional] 

## Methods

### NewItemStructs

`func NewItemStructs() *ItemStructs`

NewItemStructs instantiates a new ItemStructs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemStructsWithDefaults

`func NewItemStructsWithDefaults() *ItemStructs`

NewItemStructsWithDefaults instantiates a new ItemStructs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ItemStructs) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ItemStructs) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ItemStructs) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ItemStructs) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetItem

`func (o *ItemStructs) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ItemStructs) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ItemStructs) SetItem(v string)`

SetItem sets Item field to given value.

### HasItem

`func (o *ItemStructs) HasItem() bool`

HasItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


