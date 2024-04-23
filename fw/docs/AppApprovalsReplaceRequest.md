# AppApprovalsReplaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approvals** | Pointer to [**[]AppApproval**](AppApproval.md) |  | [optional] 
**Fields** | Pointer to [**ProtobufFieldMask**](ProtobufFieldMask.md) |  | [optional] 

## Methods

### NewAppApprovalsReplaceRequest

`func NewAppApprovalsReplaceRequest() *AppApprovalsReplaceRequest`

NewAppApprovalsReplaceRequest instantiates a new AppApprovalsReplaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppApprovalsReplaceRequestWithDefaults

`func NewAppApprovalsReplaceRequestWithDefaults() *AppApprovalsReplaceRequest`

NewAppApprovalsReplaceRequestWithDefaults instantiates a new AppApprovalsReplaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovals

`func (o *AppApprovalsReplaceRequest) GetApprovals() []AppApproval`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *AppApprovalsReplaceRequest) GetApprovalsOk() (*[]AppApproval, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *AppApprovalsReplaceRequest) SetApprovals(v []AppApproval)`

SetApprovals sets Approvals field to given value.

### HasApprovals

`func (o *AppApprovalsReplaceRequest) HasApprovals() bool`

HasApprovals returns a boolean if a field has been set.

### GetFields

`func (o *AppApprovalsReplaceRequest) GetFields() ProtobufFieldMask`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AppApprovalsReplaceRequest) GetFieldsOk() (*ProtobufFieldMask, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AppApprovalsReplaceRequest) SetFields(v ProtobufFieldMask)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AppApprovalsReplaceRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


