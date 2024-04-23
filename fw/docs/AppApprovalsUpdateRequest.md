# AppApprovalsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**ProtobufFieldMask**](ProtobufFieldMask.md) |  | [optional] 
**InsertedApprovals** | Pointer to [**[]AppApproval**](AppApproval.md) |  | [optional] 
**RemovedApprovals** | Pointer to [**[]AppApprovalRemovalRequest**](AppApprovalRemovalRequest.md) |  | [optional] 

## Methods

### NewAppApprovalsUpdateRequest

`func NewAppApprovalsUpdateRequest() *AppApprovalsUpdateRequest`

NewAppApprovalsUpdateRequest instantiates a new AppApprovalsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppApprovalsUpdateRequestWithDefaults

`func NewAppApprovalsUpdateRequestWithDefaults() *AppApprovalsUpdateRequest`

NewAppApprovalsUpdateRequestWithDefaults instantiates a new AppApprovalsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *AppApprovalsUpdateRequest) GetFields() ProtobufFieldMask`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AppApprovalsUpdateRequest) GetFieldsOk() (*ProtobufFieldMask, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AppApprovalsUpdateRequest) SetFields(v ProtobufFieldMask)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AppApprovalsUpdateRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetInsertedApprovals

`func (o *AppApprovalsUpdateRequest) GetInsertedApprovals() []AppApproval`

GetInsertedApprovals returns the InsertedApprovals field if non-nil, zero value otherwise.

### GetInsertedApprovalsOk

`func (o *AppApprovalsUpdateRequest) GetInsertedApprovalsOk() (*[]AppApproval, bool)`

GetInsertedApprovalsOk returns a tuple with the InsertedApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertedApprovals

`func (o *AppApprovalsUpdateRequest) SetInsertedApprovals(v []AppApproval)`

SetInsertedApprovals sets InsertedApprovals field to given value.

### HasInsertedApprovals

`func (o *AppApprovalsUpdateRequest) HasInsertedApprovals() bool`

HasInsertedApprovals returns a boolean if a field has been set.

### GetRemovedApprovals

`func (o *AppApprovalsUpdateRequest) GetRemovedApprovals() []AppApprovalRemovalRequest`

GetRemovedApprovals returns the RemovedApprovals field if non-nil, zero value otherwise.

### GetRemovedApprovalsOk

`func (o *AppApprovalsUpdateRequest) GetRemovedApprovalsOk() (*[]AppApprovalRemovalRequest, bool)`

GetRemovedApprovalsOk returns a tuple with the RemovedApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedApprovals

`func (o *AppApprovalsUpdateRequest) SetRemovedApprovals(v []AppApprovalRemovalRequest)`

SetRemovedApprovals sets RemovedApprovals field to given value.

### HasRemovedApprovals

`func (o *AppApprovalsUpdateRequest) HasRemovedApprovals() bool`

HasRemovedApprovals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


