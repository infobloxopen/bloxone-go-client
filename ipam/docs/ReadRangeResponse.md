# ReadRangeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**Range**](Range.md) | The Range object. | [optional] 

## Methods

### NewReadRangeResponse

`func NewReadRangeResponse() *ReadRangeResponse`

NewReadRangeResponse instantiates a new ReadRangeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadRangeResponseWithDefaults

`func NewReadRangeResponseWithDefaults() *ReadRangeResponse`

NewReadRangeResponseWithDefaults instantiates a new ReadRangeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ReadRangeResponse) GetResult() Range`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ReadRangeResponse) GetResultOk() (*Range, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ReadRangeResponse) SetResult(v Range)`

SetResult sets Result field to given value.

### HasResult

`func (o *ReadRangeResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


