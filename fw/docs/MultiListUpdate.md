# MultiListUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]int32** | The Named List object identifier. | [optional] 
**InsertedItemsDescribed** | Pointer to [**[]ItemStructs**](ItemStructs.md) | The List of ItemStructs structure which contains the item and its description | [optional] 

## Methods

### NewMultiListUpdate

`func NewMultiListUpdate() *MultiListUpdate`

NewMultiListUpdate instantiates a new MultiListUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiListUpdateWithDefaults

`func NewMultiListUpdateWithDefaults() *MultiListUpdate`

NewMultiListUpdateWithDefaults instantiates a new MultiListUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *MultiListUpdate) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *MultiListUpdate) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *MultiListUpdate) SetIds(v []int32)`

SetIds sets Ids field to given value.

### HasIds

`func (o *MultiListUpdate) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetInsertedItemsDescribed

`func (o *MultiListUpdate) GetInsertedItemsDescribed() []ItemStructs`

GetInsertedItemsDescribed returns the InsertedItemsDescribed field if non-nil, zero value otherwise.

### GetInsertedItemsDescribedOk

`func (o *MultiListUpdate) GetInsertedItemsDescribedOk() (*[]ItemStructs, bool)`

GetInsertedItemsDescribedOk returns a tuple with the InsertedItemsDescribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertedItemsDescribed

`func (o *MultiListUpdate) SetInsertedItemsDescribed(v []ItemStructs)`

SetInsertedItemsDescribed sets InsertedItemsDescribed field to given value.

### HasInsertedItemsDescribed

`func (o *MultiListUpdate) HasInsertedItemsDescribed() bool`

HasInsertedItemsDescribed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


