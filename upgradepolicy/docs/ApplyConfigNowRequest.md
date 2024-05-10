# ApplyConfigNowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | Pointer to [**[]OnpremDetails**](OnpremDetails.md) |  | [optional] 

## Methods

### NewApplyConfigNowRequest

`func NewApplyConfigNowRequest() *ApplyConfigNowRequest`

NewApplyConfigNowRequest instantiates a new ApplyConfigNowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyConfigNowRequestWithDefaults

`func NewApplyConfigNowRequestWithDefaults() *ApplyConfigNowRequest`

NewApplyConfigNowRequestWithDefaults instantiates a new ApplyConfigNowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *ApplyConfigNowRequest) GetPayload() []OnpremDetails`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ApplyConfigNowRequest) GetPayloadOk() (*[]OnpremDetails, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ApplyConfigNowRequest) SetPayload(v []OnpremDetails)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ApplyConfigNowRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


