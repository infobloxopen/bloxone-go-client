# NetworkList

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

### NewNetworkList

`func NewNetworkList() *NetworkList`

NewNetworkList instantiates a new NetworkList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkListWithDefaults

`func NewNetworkListWithDefaults() *NetworkList`

NewNetworkListWithDefaults instantiates a new NetworkList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *NetworkList) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *NetworkList) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *NetworkList) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *NetworkList) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkList) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkList) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkList) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkList) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *NetworkList) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkList) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkList) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *NetworkList) GetItems() []string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NetworkList) GetItemsOk() (*[]string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NetworkList) SetItems(v []string)`

SetItems sets Items field to given value.

### HasItems

`func (o *NetworkList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *NetworkList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkList) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkList) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicyId

`func (o *NetworkList) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *NetworkList) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *NetworkList) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *NetworkList) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *NetworkList) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *NetworkList) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *NetworkList) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *NetworkList) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


