# FederationDelegationBulkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delegations** | Pointer to [**[]FederationDelegationBulkItem**](FederationDelegationBulkItem.md) | list of action/delegation requests to execute. | [optional] 
**Fields** | Pointer to [**ProtobufFieldMask**](ProtobufFieldMask.md) |  | [optional] 

## Methods

### NewFederationDelegationBulkRequest

`func NewFederationDelegationBulkRequest() *FederationDelegationBulkRequest`

NewFederationDelegationBulkRequest instantiates a new FederationDelegationBulkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationDelegationBulkRequestWithDefaults

`func NewFederationDelegationBulkRequestWithDefaults() *FederationDelegationBulkRequest`

NewFederationDelegationBulkRequestWithDefaults instantiates a new FederationDelegationBulkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegations

`func (o *FederationDelegationBulkRequest) GetDelegations() []FederationDelegationBulkItem`

GetDelegations returns the Delegations field if non-nil, zero value otherwise.

### GetDelegationsOk

`func (o *FederationDelegationBulkRequest) GetDelegationsOk() (*[]FederationDelegationBulkItem, bool)`

GetDelegationsOk returns a tuple with the Delegations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegations

`func (o *FederationDelegationBulkRequest) SetDelegations(v []FederationDelegationBulkItem)`

SetDelegations sets Delegations field to given value.

### HasDelegations

`func (o *FederationDelegationBulkRequest) HasDelegations() bool`

HasDelegations returns a boolean if a field has been set.

### GetFields

`func (o *FederationDelegationBulkRequest) GetFields() ProtobufFieldMask`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *FederationDelegationBulkRequest) GetFieldsOk() (*ProtobufFieldMask, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *FederationDelegationBulkRequest) SetFields(v ProtobufFieldMask)`

SetFields sets Fields field to given value.

### HasFields

`func (o *FederationDelegationBulkRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


