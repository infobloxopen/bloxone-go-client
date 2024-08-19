# ReadDelegationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**Delegation**](Delegation.md) | The Delegation object. | [optional] 

## Methods

### NewReadDelegationResponse

`func NewReadDelegationResponse() *ReadDelegationResponse`

NewReadDelegationResponse instantiates a new ReadDelegationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadDelegationResponseWithDefaults

`func NewReadDelegationResponseWithDefaults() *ReadDelegationResponse`

NewReadDelegationResponseWithDefaults instantiates a new ReadDelegationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ReadDelegationResponse) GetResult() Delegation`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ReadDelegationResponse) GetResultOk() (*Delegation, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ReadDelegationResponse) SetResult(v Delegation)`

SetResult sets Result field to given value.

### HasResult

`func (o *ReadDelegationResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


