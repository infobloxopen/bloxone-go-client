# ListPoPRegionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]PoPRegion**](PoPRegion.md) |  | [optional] 
**TotalResultCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewListPoPRegionsResponse

`func NewListPoPRegionsResponse() *ListPoPRegionsResponse`

NewListPoPRegionsResponse instantiates a new ListPoPRegionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPoPRegionsResponseWithDefaults

`func NewListPoPRegionsResponseWithDefaults() *ListPoPRegionsResponse`

NewListPoPRegionsResponseWithDefaults instantiates a new ListPoPRegionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ListPoPRegionsResponse) GetResults() []PoPRegion`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListPoPRegionsResponse) GetResultsOk() (*[]PoPRegion, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListPoPRegionsResponse) SetResults(v []PoPRegion)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListPoPRegionsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotalResultCount

`func (o *ListPoPRegionsResponse) GetTotalResultCount() int32`

GetTotalResultCount returns the TotalResultCount field if non-nil, zero value otherwise.

### GetTotalResultCountOk

`func (o *ListPoPRegionsResponse) GetTotalResultCountOk() (*int32, bool)`

GetTotalResultCountOk returns a tuple with the TotalResultCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultCount

`func (o *ListPoPRegionsResponse) SetTotalResultCount(v int32)`

SetTotalResultCount sets TotalResultCount field to given value.

### HasTotalResultCount

`func (o *ListPoPRegionsResponse) HasTotalResultCount() bool`

HasTotalResultCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


