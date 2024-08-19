# DelegationBulkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delegations** | Pointer to [**[]DelegationBulkItem**](DelegationBulkItem.md) | list of action/delegation requests to execute. | [optional] 
**Fields** | Pointer to [**ProtobufFieldMask**](ProtobufFieldMask.md) |  | [optional] 

## Methods

### NewDelegationBulkRequest

`func NewDelegationBulkRequest() *DelegationBulkRequest`

NewDelegationBulkRequest instantiates a new DelegationBulkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegationBulkRequestWithDefaults

`func NewDelegationBulkRequestWithDefaults() *DelegationBulkRequest`

NewDelegationBulkRequestWithDefaults instantiates a new DelegationBulkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegations

`func (o *DelegationBulkRequest) GetDelegations() []DelegationBulkItem`

GetDelegations returns the Delegations field if non-nil, zero value otherwise.

### GetDelegationsOk

`func (o *DelegationBulkRequest) GetDelegationsOk() (*[]DelegationBulkItem, bool)`

GetDelegationsOk returns a tuple with the Delegations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegations

`func (o *DelegationBulkRequest) SetDelegations(v []DelegationBulkItem)`

SetDelegations sets Delegations field to given value.

### HasDelegations

`func (o *DelegationBulkRequest) HasDelegations() bool`

HasDelegations returns a boolean if a field has been set.

### GetFields

`func (o *DelegationBulkRequest) GetFields() ProtobufFieldMask`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DelegationBulkRequest) GetFieldsOk() (*ProtobufFieldMask, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DelegationBulkRequest) SetFields(v ProtobufFieldMask)`

SetFields sets Fields field to given value.

### HasFields

`func (o *DelegationBulkRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


