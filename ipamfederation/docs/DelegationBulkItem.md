# DelegationBulkItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Type of action to execute on Delegation object.  Valid values are: * _POST_ * _PATCH_ * _DELETE_ | [optional] 
**Delegation** | Pointer to [**Delegation**](Delegation.md) | The Delegation object to execute the action on. | [optional] 

## Methods

### NewDelegationBulkItem

`func NewDelegationBulkItem() *DelegationBulkItem`

NewDelegationBulkItem instantiates a new DelegationBulkItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegationBulkItemWithDefaults

`func NewDelegationBulkItemWithDefaults() *DelegationBulkItem`

NewDelegationBulkItemWithDefaults instantiates a new DelegationBulkItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *DelegationBulkItem) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DelegationBulkItem) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DelegationBulkItem) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *DelegationBulkItem) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDelegation

`func (o *DelegationBulkItem) GetDelegation() Delegation`

GetDelegation returns the Delegation field if non-nil, zero value otherwise.

### GetDelegationOk

`func (o *DelegationBulkItem) GetDelegationOk() (*Delegation, bool)`

GetDelegationOk returns a tuple with the Delegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegation

`func (o *DelegationBulkItem) SetDelegation(v Delegation)`

SetDelegation sets Delegation field to given value.

### HasDelegation

`func (o *DelegationBulkItem) HasDelegation() bool`

HasDelegation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


