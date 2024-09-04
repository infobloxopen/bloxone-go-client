# ListOverlappingBlockResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]OverlappingBlock**](OverlappingBlock.md) | A list of OverlappingBlock objects. | [optional] 

## Methods

### NewListOverlappingBlockResponse

`func NewListOverlappingBlockResponse() *ListOverlappingBlockResponse`

NewListOverlappingBlockResponse instantiates a new ListOverlappingBlockResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOverlappingBlockResponseWithDefaults

`func NewListOverlappingBlockResponseWithDefaults() *ListOverlappingBlockResponse`

NewListOverlappingBlockResponseWithDefaults instantiates a new ListOverlappingBlockResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ListOverlappingBlockResponse) GetResults() []OverlappingBlock`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListOverlappingBlockResponse) GetResultsOk() (*[]OverlappingBlock, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListOverlappingBlockResponse) SetResults(v []OverlappingBlock)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListOverlappingBlockResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


