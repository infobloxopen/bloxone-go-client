# ConfigBulkCopyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]ConfigBulkCopyError**](ConfigBulkCopyError.md) |  | [optional] 
**Results** | Pointer to [**[]ConfigCopyResponse**](ConfigCopyResponse.md) |  | [optional] 

## Methods

### NewConfigBulkCopyResponse

`func NewConfigBulkCopyResponse() *ConfigBulkCopyResponse`

NewConfigBulkCopyResponse instantiates a new ConfigBulkCopyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigBulkCopyResponseWithDefaults

`func NewConfigBulkCopyResponseWithDefaults() *ConfigBulkCopyResponse`

NewConfigBulkCopyResponseWithDefaults instantiates a new ConfigBulkCopyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *ConfigBulkCopyResponse) GetErrors() []ConfigBulkCopyError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ConfigBulkCopyResponse) GetErrorsOk() (*[]ConfigBulkCopyError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ConfigBulkCopyResponse) SetErrors(v []ConfigBulkCopyError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ConfigBulkCopyResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetResults

`func (o *ConfigBulkCopyResponse) GetResults() []ConfigCopyResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ConfigBulkCopyResponse) GetResultsOk() (*[]ConfigCopyResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ConfigBulkCopyResponse) SetResults(v []ConfigCopyResponse)`

SetResults sets Results field to given value.

### HasResults

`func (o *ConfigBulkCopyResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


