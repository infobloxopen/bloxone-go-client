# FederationDelegationBulkItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Type of action to execute on Delegation object.  Valid values are: * _POST_ * _PATCH_ * _DELETE_ | [optional] 
**Delegation** | Pointer to [**FederationDelegation**](FederationDelegation.md) | The Delegation object to execute the action on. | [optional] 

## Methods

### NewFederationDelegationBulkItem

`func NewFederationDelegationBulkItem() *FederationDelegationBulkItem`

NewFederationDelegationBulkItem instantiates a new FederationDelegationBulkItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationDelegationBulkItemWithDefaults

`func NewFederationDelegationBulkItemWithDefaults() *FederationDelegationBulkItem`

NewFederationDelegationBulkItemWithDefaults instantiates a new FederationDelegationBulkItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *FederationDelegationBulkItem) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FederationDelegationBulkItem) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FederationDelegationBulkItem) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *FederationDelegationBulkItem) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDelegation

`func (o *FederationDelegationBulkItem) GetDelegation() FederationDelegation`

GetDelegation returns the Delegation field if non-nil, zero value otherwise.

### GetDelegationOk

`func (o *FederationDelegationBulkItem) GetDelegationOk() (*FederationDelegation, bool)`

GetDelegationOk returns a tuple with the Delegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegation

`func (o *FederationDelegationBulkItem) SetDelegation(v FederationDelegation)`

SetDelegation sets Delegation field to given value.

### HasDelegation

`func (o *FederationDelegationBulkItem) HasDelegation() bool`

HasDelegation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


