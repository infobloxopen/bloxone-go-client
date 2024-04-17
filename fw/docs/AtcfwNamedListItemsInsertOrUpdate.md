# AtcfwNamedListItemsInsertOrUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The Named List object identifier. | [optional] [readonly] 
**Items** | Pointer to **[]string** | The list of the FQDN or IPv4/IPv6 addresses or IPv4/IPv6 CIDRs to define whitelists and blacklists for additional protection. | [optional] 
**ItemsDescribed** | Pointer to [**[]AtcfwItemStructs**](AtcfwItemStructs.md) | The List of ItemStructs structure which contains the item and its description | [optional] 

## Methods

### NewAtcfwNamedListItemsInsertOrUpdate

`func NewAtcfwNamedListItemsInsertOrUpdate() *AtcfwNamedListItemsInsertOrUpdate`

NewAtcfwNamedListItemsInsertOrUpdate instantiates a new AtcfwNamedListItemsInsertOrUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtcfwNamedListItemsInsertOrUpdateWithDefaults

`func NewAtcfwNamedListItemsInsertOrUpdateWithDefaults() *AtcfwNamedListItemsInsertOrUpdate`

NewAtcfwNamedListItemsInsertOrUpdateWithDefaults instantiates a new AtcfwNamedListItemsInsertOrUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AtcfwNamedListItemsInsertOrUpdate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AtcfwNamedListItemsInsertOrUpdate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AtcfwNamedListItemsInsertOrUpdate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AtcfwNamedListItemsInsertOrUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *AtcfwNamedListItemsInsertOrUpdate) GetItems() []string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AtcfwNamedListItemsInsertOrUpdate) GetItemsOk() (*[]string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AtcfwNamedListItemsInsertOrUpdate) SetItems(v []string)`

SetItems sets Items field to given value.

### HasItems

`func (o *AtcfwNamedListItemsInsertOrUpdate) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetItemsDescribed

`func (o *AtcfwNamedListItemsInsertOrUpdate) GetItemsDescribed() []AtcfwItemStructs`

GetItemsDescribed returns the ItemsDescribed field if non-nil, zero value otherwise.

### GetItemsDescribedOk

`func (o *AtcfwNamedListItemsInsertOrUpdate) GetItemsDescribedOk() (*[]AtcfwItemStructs, bool)`

GetItemsDescribedOk returns a tuple with the ItemsDescribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsDescribed

`func (o *AtcfwNamedListItemsInsertOrUpdate) SetItemsDescribed(v []AtcfwItemStructs)`

SetItemsDescribed sets ItemsDescribed field to given value.

### HasItemsDescribed

`func (o *AtcfwNamedListItemsInsertOrUpdate) HasItemsDescribed() bool`

HasItemsDescribed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


