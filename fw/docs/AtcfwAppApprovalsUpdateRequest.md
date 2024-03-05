# AtcfwAppApprovalsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**ProtobufFieldMask**](ProtobufFieldMask.md) |  | [optional] 
**InsertedApprovals** | Pointer to [**[]AtcfwAppApproval**](AtcfwAppApproval.md) |  | [optional] 
**RemovedApprovals** | Pointer to [**[]AtcfwAppApprovalRemovalRequest**](AtcfwAppApprovalRemovalRequest.md) |  | [optional] 

## Methods

### NewAtcfwAppApprovalsUpdateRequest

`func NewAtcfwAppApprovalsUpdateRequest() *AtcfwAppApprovalsUpdateRequest`

NewAtcfwAppApprovalsUpdateRequest instantiates a new AtcfwAppApprovalsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtcfwAppApprovalsUpdateRequestWithDefaults

`func NewAtcfwAppApprovalsUpdateRequestWithDefaults() *AtcfwAppApprovalsUpdateRequest`

NewAtcfwAppApprovalsUpdateRequestWithDefaults instantiates a new AtcfwAppApprovalsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *AtcfwAppApprovalsUpdateRequest) GetFields() ProtobufFieldMask`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AtcfwAppApprovalsUpdateRequest) GetFieldsOk() (*ProtobufFieldMask, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AtcfwAppApprovalsUpdateRequest) SetFields(v ProtobufFieldMask)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AtcfwAppApprovalsUpdateRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetInsertedApprovals

`func (o *AtcfwAppApprovalsUpdateRequest) GetInsertedApprovals() []AtcfwAppApproval`

GetInsertedApprovals returns the InsertedApprovals field if non-nil, zero value otherwise.

### GetInsertedApprovalsOk

`func (o *AtcfwAppApprovalsUpdateRequest) GetInsertedApprovalsOk() (*[]AtcfwAppApproval, bool)`

GetInsertedApprovalsOk returns a tuple with the InsertedApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertedApprovals

`func (o *AtcfwAppApprovalsUpdateRequest) SetInsertedApprovals(v []AtcfwAppApproval)`

SetInsertedApprovals sets InsertedApprovals field to given value.

### HasInsertedApprovals

`func (o *AtcfwAppApprovalsUpdateRequest) HasInsertedApprovals() bool`

HasInsertedApprovals returns a boolean if a field has been set.

### GetRemovedApprovals

`func (o *AtcfwAppApprovalsUpdateRequest) GetRemovedApprovals() []AtcfwAppApprovalRemovalRequest`

GetRemovedApprovals returns the RemovedApprovals field if non-nil, zero value otherwise.

### GetRemovedApprovalsOk

`func (o *AtcfwAppApprovalsUpdateRequest) GetRemovedApprovalsOk() (*[]AtcfwAppApprovalRemovalRequest, bool)`

GetRemovedApprovalsOk returns a tuple with the RemovedApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedApprovals

`func (o *AtcfwAppApprovalsUpdateRequest) SetRemovedApprovals(v []AtcfwAppApprovalRemovalRequest)`

SetRemovedApprovals sets RemovedApprovals field to given value.

### HasRemovedApprovals

`func (o *AtcfwAppApprovalsUpdateRequest) HasRemovedApprovals() bool`

HasRemovedApprovals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


