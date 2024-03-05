# AtcfwAppApprovalsReplaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approvals** | Pointer to [**[]AtcfwAppApproval**](AtcfwAppApproval.md) |  | [optional] 
**Fields** | Pointer to [**ProtobufFieldMask**](ProtobufFieldMask.md) |  | [optional] 

## Methods

### NewAtcfwAppApprovalsReplaceRequest

`func NewAtcfwAppApprovalsReplaceRequest() *AtcfwAppApprovalsReplaceRequest`

NewAtcfwAppApprovalsReplaceRequest instantiates a new AtcfwAppApprovalsReplaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtcfwAppApprovalsReplaceRequestWithDefaults

`func NewAtcfwAppApprovalsReplaceRequestWithDefaults() *AtcfwAppApprovalsReplaceRequest`

NewAtcfwAppApprovalsReplaceRequestWithDefaults instantiates a new AtcfwAppApprovalsReplaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovals

`func (o *AtcfwAppApprovalsReplaceRequest) GetApprovals() []AtcfwAppApproval`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *AtcfwAppApprovalsReplaceRequest) GetApprovalsOk() (*[]AtcfwAppApproval, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *AtcfwAppApprovalsReplaceRequest) SetApprovals(v []AtcfwAppApproval)`

SetApprovals sets Approvals field to given value.

### HasApprovals

`func (o *AtcfwAppApprovalsReplaceRequest) HasApprovals() bool`

HasApprovals returns a boolean if a field has been set.

### GetFields

`func (o *AtcfwAppApprovalsReplaceRequest) GetFields() ProtobufFieldMask`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AtcfwAppApprovalsReplaceRequest) GetFieldsOk() (*ProtobufFieldMask, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AtcfwAppApprovalsReplaceRequest) SetFields(v ProtobufFieldMask)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AtcfwAppApprovalsReplaceRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


