# InternalDomainsItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedItemsDescribed** | Pointer to [**[]ItemStructs**](ItemStructs.md) | The List of ItemStructs structure which contains the item and its description | [optional] 
**Id** | Pointer to **int32** | The Internal Domain List object identifier. | [optional] [readonly] 
**InsertedItemsDescribed** | Pointer to [**[]ItemStructs**](ItemStructs.md) | The List of ItemStructs structure which contains the item and its description | [optional] 

## Methods

### NewInternalDomainsItems

`func NewInternalDomainsItems() *InternalDomainsItems`

NewInternalDomainsItems instantiates a new InternalDomainsItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalDomainsItemsWithDefaults

`func NewInternalDomainsItemsWithDefaults() *InternalDomainsItems`

NewInternalDomainsItemsWithDefaults instantiates a new InternalDomainsItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedItemsDescribed

`func (o *InternalDomainsItems) GetDeletedItemsDescribed() []ItemStructs`

GetDeletedItemsDescribed returns the DeletedItemsDescribed field if non-nil, zero value otherwise.

### GetDeletedItemsDescribedOk

`func (o *InternalDomainsItems) GetDeletedItemsDescribedOk() (*[]ItemStructs, bool)`

GetDeletedItemsDescribedOk returns a tuple with the DeletedItemsDescribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedItemsDescribed

`func (o *InternalDomainsItems) SetDeletedItemsDescribed(v []ItemStructs)`

SetDeletedItemsDescribed sets DeletedItemsDescribed field to given value.

### HasDeletedItemsDescribed

`func (o *InternalDomainsItems) HasDeletedItemsDescribed() bool`

HasDeletedItemsDescribed returns a boolean if a field has been set.

### GetId

`func (o *InternalDomainsItems) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalDomainsItems) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalDomainsItems) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InternalDomainsItems) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsertedItemsDescribed

`func (o *InternalDomainsItems) GetInsertedItemsDescribed() []ItemStructs`

GetInsertedItemsDescribed returns the InsertedItemsDescribed field if non-nil, zero value otherwise.

### GetInsertedItemsDescribedOk

`func (o *InternalDomainsItems) GetInsertedItemsDescribedOk() (*[]ItemStructs, bool)`

GetInsertedItemsDescribedOk returns a tuple with the InsertedItemsDescribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertedItemsDescribed

`func (o *InternalDomainsItems) SetInsertedItemsDescribed(v []ItemStructs)`

SetInsertedItemsDescribed sets InsertedItemsDescribed field to given value.

### HasInsertedItemsDescribed

`func (o *InternalDomainsItems) HasInsertedItemsDescribed() bool`

HasInsertedItemsDescribed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


