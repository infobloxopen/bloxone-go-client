# AtcfwMultiListUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]int32** | The Named List object identifier. | [optional] 
**InsertedItemsDescribed** | Pointer to [**[]AtcfwItemStructs**](AtcfwItemStructs.md) | The List of ItemStructs structure which contains the item and its description | [optional] 

## Methods

### NewAtcfwMultiListUpdate

`func NewAtcfwMultiListUpdate() *AtcfwMultiListUpdate`

NewAtcfwMultiListUpdate instantiates a new AtcfwMultiListUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtcfwMultiListUpdateWithDefaults

`func NewAtcfwMultiListUpdateWithDefaults() *AtcfwMultiListUpdate`

NewAtcfwMultiListUpdateWithDefaults instantiates a new AtcfwMultiListUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *AtcfwMultiListUpdate) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *AtcfwMultiListUpdate) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *AtcfwMultiListUpdate) SetIds(v []int32)`

SetIds sets Ids field to given value.

### HasIds

`func (o *AtcfwMultiListUpdate) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetInsertedItemsDescribed

`func (o *AtcfwMultiListUpdate) GetInsertedItemsDescribed() []AtcfwItemStructs`

GetInsertedItemsDescribed returns the InsertedItemsDescribed field if non-nil, zero value otherwise.

### GetInsertedItemsDescribedOk

`func (o *AtcfwMultiListUpdate) GetInsertedItemsDescribedOk() (*[]AtcfwItemStructs, bool)`

GetInsertedItemsDescribedOk returns a tuple with the InsertedItemsDescribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertedItemsDescribed

`func (o *AtcfwMultiListUpdate) SetInsertedItemsDescribed(v []AtcfwItemStructs)`

SetInsertedItemsDescribed sets InsertedItemsDescribed field to given value.

### HasInsertedItemsDescribed

`func (o *AtcfwMultiListUpdate) HasInsertedItemsDescribed() bool`

HasInsertedItemsDescribed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


