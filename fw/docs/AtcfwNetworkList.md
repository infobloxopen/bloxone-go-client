# AtcfwNetworkList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **time.Time** | The time when this Network List object was created. | [optional] [readonly] 
**Description** | Pointer to **string** | The brief description for the network list. | [optional] 
**Id** | Pointer to **int32** | The Network List object identifier. | [optional] [readonly] 
**Items** | Pointer to **[]string** | The list of networks&#39; CIDRs that are subject for malicious attacks protection. | [optional] 
**Name** | Pointer to **string** | The name of the network list. | [optional] 
**PolicyId** | Pointer to **int32** | The identifier of the security policy with which the network list is associated. | [optional] [readonly] 
**UpdatedTime** | Pointer to **time.Time** | The time when this Network List object was last updated. | [optional] [readonly] 

## Methods

### NewAtcfwNetworkList

`func NewAtcfwNetworkList() *AtcfwNetworkList`

NewAtcfwNetworkList instantiates a new AtcfwNetworkList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtcfwNetworkListWithDefaults

`func NewAtcfwNetworkListWithDefaults() *AtcfwNetworkList`

NewAtcfwNetworkListWithDefaults instantiates a new AtcfwNetworkList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *AtcfwNetworkList) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AtcfwNetworkList) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AtcfwNetworkList) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *AtcfwNetworkList) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *AtcfwNetworkList) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AtcfwNetworkList) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AtcfwNetworkList) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AtcfwNetworkList) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *AtcfwNetworkList) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AtcfwNetworkList) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AtcfwNetworkList) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AtcfwNetworkList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *AtcfwNetworkList) GetItems() []string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AtcfwNetworkList) GetItemsOk() (*[]string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AtcfwNetworkList) SetItems(v []string)`

SetItems sets Items field to given value.

### HasItems

`func (o *AtcfwNetworkList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *AtcfwNetworkList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AtcfwNetworkList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AtcfwNetworkList) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AtcfwNetworkList) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicyId

`func (o *AtcfwNetworkList) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *AtcfwNetworkList) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *AtcfwNetworkList) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *AtcfwNetworkList) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *AtcfwNetworkList) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *AtcfwNetworkList) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *AtcfwNetworkList) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *AtcfwNetworkList) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


