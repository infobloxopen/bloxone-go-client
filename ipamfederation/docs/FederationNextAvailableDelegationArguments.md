# FederationNextAvailableDelegationArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | Pointer to **int64** | The CIDR of the delegation to be created. This is required, if _address_ does not specify it in its input. | [optional] 
**Count** | Pointer to **int64** | The count of __Delegation__ required. If not provided, it will default to 1. | [optional] 
**FederatedRealms** | Pointer to **[]string** | The resource identifier. | [optional] 
**Predicates** | Pointer to **map[string]interface{}** | It  contains the map of attributes and the associated value. | [optional] 
**ReadOnly** | Pointer to **bool** | The field which denotes if the __Delegation__ will be provisioned or just calculated and returned. | [optional] 

## Methods

### NewFederationNextAvailableDelegationArguments

`func NewFederationNextAvailableDelegationArguments() *FederationNextAvailableDelegationArguments`

NewFederationNextAvailableDelegationArguments instantiates a new FederationNextAvailableDelegationArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationNextAvailableDelegationArgumentsWithDefaults

`func NewFederationNextAvailableDelegationArgumentsWithDefaults() *FederationNextAvailableDelegationArguments`

NewFederationNextAvailableDelegationArgumentsWithDefaults instantiates a new FederationNextAvailableDelegationArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *FederationNextAvailableDelegationArguments) GetCidr() int64`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *FederationNextAvailableDelegationArguments) GetCidrOk() (*int64, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *FederationNextAvailableDelegationArguments) SetCidr(v int64)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *FederationNextAvailableDelegationArguments) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCount

`func (o *FederationNextAvailableDelegationArguments) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FederationNextAvailableDelegationArguments) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FederationNextAvailableDelegationArguments) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *FederationNextAvailableDelegationArguments) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetFederatedRealms

`func (o *FederationNextAvailableDelegationArguments) GetFederatedRealms() []string`

GetFederatedRealms returns the FederatedRealms field if non-nil, zero value otherwise.

### GetFederatedRealmsOk

`func (o *FederationNextAvailableDelegationArguments) GetFederatedRealmsOk() (*[]string, bool)`

GetFederatedRealmsOk returns a tuple with the FederatedRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedRealms

`func (o *FederationNextAvailableDelegationArguments) SetFederatedRealms(v []string)`

SetFederatedRealms sets FederatedRealms field to given value.

### HasFederatedRealms

`func (o *FederationNextAvailableDelegationArguments) HasFederatedRealms() bool`

HasFederatedRealms returns a boolean if a field has been set.

### GetPredicates

`func (o *FederationNextAvailableDelegationArguments) GetPredicates() map[string]interface{}`

GetPredicates returns the Predicates field if non-nil, zero value otherwise.

### GetPredicatesOk

`func (o *FederationNextAvailableDelegationArguments) GetPredicatesOk() (*map[string]interface{}, bool)`

GetPredicatesOk returns a tuple with the Predicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicates

`func (o *FederationNextAvailableDelegationArguments) SetPredicates(v map[string]interface{})`

SetPredicates sets Predicates field to given value.

### HasPredicates

`func (o *FederationNextAvailableDelegationArguments) HasPredicates() bool`

HasPredicates returns a boolean if a field has been set.

### GetReadOnly

`func (o *FederationNextAvailableDelegationArguments) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *FederationNextAvailableDelegationArguments) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *FederationNextAvailableDelegationArguments) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *FederationNextAvailableDelegationArguments) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


