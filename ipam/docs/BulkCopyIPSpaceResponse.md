# BulkCopyIPSpaceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]BulkCopyError**](BulkCopyError.md) |  | [optional] 
**Results** | Pointer to [**[]CopyResponse**](CopyResponse.md) |  | [optional] 

## Methods

### NewBulkCopyIPSpaceResponse

`func NewBulkCopyIPSpaceResponse() *BulkCopyIPSpaceResponse`

NewBulkCopyIPSpaceResponse instantiates a new BulkCopyIPSpaceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCopyIPSpaceResponseWithDefaults

`func NewBulkCopyIPSpaceResponseWithDefaults() *BulkCopyIPSpaceResponse`

NewBulkCopyIPSpaceResponseWithDefaults instantiates a new BulkCopyIPSpaceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *BulkCopyIPSpaceResponse) GetErrors() []BulkCopyError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BulkCopyIPSpaceResponse) GetErrorsOk() (*[]BulkCopyError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BulkCopyIPSpaceResponse) SetErrors(v []BulkCopyError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BulkCopyIPSpaceResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetResults

`func (o *BulkCopyIPSpaceResponse) GetResults() []CopyResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BulkCopyIPSpaceResponse) GetResultsOk() (*[]CopyResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BulkCopyIPSpaceResponse) SetResults(v []CopyResponse)`

SetResults sets Results field to given value.

### HasResults

`func (o *BulkCopyIPSpaceResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


