# FederationAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allocated** | Pointer to **int64** | Percent of total space allocated. | [optional] [readonly] 
**Delegated** | Pointer to **int64** | Percent of total space delegated. | [optional] [readonly] 
**Overlapping** | Pointer to **int64** | Percent of total space in overlapping blocks. | [optional] [readonly] 
**Reserved** | Pointer to **int64** | Percent of total space reserved. | [optional] [readonly] 

## Methods

### NewFederationAllocation

`func NewFederationAllocation() *FederationAllocation`

NewFederationAllocation instantiates a new FederationAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationAllocationWithDefaults

`func NewFederationAllocationWithDefaults() *FederationAllocation`

NewFederationAllocationWithDefaults instantiates a new FederationAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocated

`func (o *FederationAllocation) GetAllocated() int64`

GetAllocated returns the Allocated field if non-nil, zero value otherwise.

### GetAllocatedOk

`func (o *FederationAllocation) GetAllocatedOk() (*int64, bool)`

GetAllocatedOk returns a tuple with the Allocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocated

`func (o *FederationAllocation) SetAllocated(v int64)`

SetAllocated sets Allocated field to given value.

### HasAllocated

`func (o *FederationAllocation) HasAllocated() bool`

HasAllocated returns a boolean if a field has been set.

### GetDelegated

`func (o *FederationAllocation) GetDelegated() int64`

GetDelegated returns the Delegated field if non-nil, zero value otherwise.

### GetDelegatedOk

`func (o *FederationAllocation) GetDelegatedOk() (*int64, bool)`

GetDelegatedOk returns a tuple with the Delegated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegated

`func (o *FederationAllocation) SetDelegated(v int64)`

SetDelegated sets Delegated field to given value.

### HasDelegated

`func (o *FederationAllocation) HasDelegated() bool`

HasDelegated returns a boolean if a field has been set.

### GetOverlapping

`func (o *FederationAllocation) GetOverlapping() int64`

GetOverlapping returns the Overlapping field if non-nil, zero value otherwise.

### GetOverlappingOk

`func (o *FederationAllocation) GetOverlappingOk() (*int64, bool)`

GetOverlappingOk returns a tuple with the Overlapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlapping

`func (o *FederationAllocation) SetOverlapping(v int64)`

SetOverlapping sets Overlapping field to given value.

### HasOverlapping

`func (o *FederationAllocation) HasOverlapping() bool`

HasOverlapping returns a boolean if a field has been set.

### GetReserved

`func (o *FederationAllocation) GetReserved() int64`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *FederationAllocation) GetReservedOk() (*int64, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *FederationAllocation) SetReserved(v int64)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *FederationAllocation) HasReserved() bool`

HasReserved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


