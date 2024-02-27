# AtcfwNamedListItemsPartialUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedItemsDescribed** | Pointer to [**[]AtcfwItemStructs**](AtcfwItemStructs.md) | The List of ItemStructs structure which contains the item and its description | [optional] 
**Id** | Pointer to **int32** | The Named List object identifier. | [optional] [readonly] 
**InsertedItemsDescribed** | Pointer to [**[]AtcfwItemStructs**](AtcfwItemStructs.md) | The List of ItemStructs structure which contains the item and its description | [optional] 

## Methods

### NewAtcfwNamedListItemsPartialUpdate

`func NewAtcfwNamedListItemsPartialUpdate() *AtcfwNamedListItemsPartialUpdate`

NewAtcfwNamedListItemsPartialUpdate instantiates a new AtcfwNamedListItemsPartialUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtcfwNamedListItemsPartialUpdateWithDefaults

`func NewAtcfwNamedListItemsPartialUpdateWithDefaults() *AtcfwNamedListItemsPartialUpdate`

NewAtcfwNamedListItemsPartialUpdateWithDefaults instantiates a new AtcfwNamedListItemsPartialUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedItemsDescribed

`func (o *AtcfwNamedListItemsPartialUpdate) GetDeletedItemsDescribed() []AtcfwItemStructs`

GetDeletedItemsDescribed returns the DeletedItemsDescribed field if non-nil, zero value otherwise.

### GetDeletedItemsDescribedOk

`func (o *AtcfwNamedListItemsPartialUpdate) GetDeletedItemsDescribedOk() (*[]AtcfwItemStructs, bool)`

GetDeletedItemsDescribedOk returns a tuple with the DeletedItemsDescribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedItemsDescribed

`func (o *AtcfwNamedListItemsPartialUpdate) SetDeletedItemsDescribed(v []AtcfwItemStructs)`

SetDeletedItemsDescribed sets DeletedItemsDescribed field to given value.

### HasDeletedItemsDescribed

`func (o *AtcfwNamedListItemsPartialUpdate) HasDeletedItemsDescribed() bool`

HasDeletedItemsDescribed returns a boolean if a field has been set.

### GetId

`func (o *AtcfwNamedListItemsPartialUpdate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AtcfwNamedListItemsPartialUpdate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AtcfwNamedListItemsPartialUpdate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AtcfwNamedListItemsPartialUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsertedItemsDescribed

`func (o *AtcfwNamedListItemsPartialUpdate) GetInsertedItemsDescribed() []AtcfwItemStructs`

GetInsertedItemsDescribed returns the InsertedItemsDescribed field if non-nil, zero value otherwise.

### GetInsertedItemsDescribedOk

`func (o *AtcfwNamedListItemsPartialUpdate) GetInsertedItemsDescribedOk() (*[]AtcfwItemStructs, bool)`

GetInsertedItemsDescribedOk returns a tuple with the InsertedItemsDescribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertedItemsDescribed

`func (o *AtcfwNamedListItemsPartialUpdate) SetInsertedItemsDescribed(v []AtcfwItemStructs)`

SetInsertedItemsDescribed sets InsertedItemsDescribed field to given value.

### HasInsertedItemsDescribed

`func (o *AtcfwNamedListItemsPartialUpdate) HasInsertedItemsDescribed() bool`

HasInsertedItemsDescribed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


