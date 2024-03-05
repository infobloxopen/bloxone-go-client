# AtcfwListPoPRegionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]AtcfwPoPRegion**](AtcfwPoPRegion.md) |  | [optional] 
**TotalResultCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewAtcfwListPoPRegionsResponse

`func NewAtcfwListPoPRegionsResponse() *AtcfwListPoPRegionsResponse`

NewAtcfwListPoPRegionsResponse instantiates a new AtcfwListPoPRegionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtcfwListPoPRegionsResponseWithDefaults

`func NewAtcfwListPoPRegionsResponseWithDefaults() *AtcfwListPoPRegionsResponse`

NewAtcfwListPoPRegionsResponseWithDefaults instantiates a new AtcfwListPoPRegionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *AtcfwListPoPRegionsResponse) GetResults() []AtcfwPoPRegion`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AtcfwListPoPRegionsResponse) GetResultsOk() (*[]AtcfwPoPRegion, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AtcfwListPoPRegionsResponse) SetResults(v []AtcfwPoPRegion)`

SetResults sets Results field to given value.

### HasResults

`func (o *AtcfwListPoPRegionsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotalResultCount

`func (o *AtcfwListPoPRegionsResponse) GetTotalResultCount() int32`

GetTotalResultCount returns the TotalResultCount field if non-nil, zero value otherwise.

### GetTotalResultCountOk

`func (o *AtcfwListPoPRegionsResponse) GetTotalResultCountOk() (*int32, bool)`

GetTotalResultCountOk returns a tuple with the TotalResultCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultCount

`func (o *AtcfwListPoPRegionsResponse) SetTotalResultCount(v int32)`

SetTotalResultCount sets TotalResultCount field to given value.

### HasTotalResultCount

`func (o *AtcfwListPoPRegionsResponse) HasTotalResultCount() bool`

HasTotalResultCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


