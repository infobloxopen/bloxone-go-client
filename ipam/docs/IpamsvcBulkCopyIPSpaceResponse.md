# IpamsvcBulkCopyIPSpaceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]IpamsvcBulkCopyError**](IpamsvcBulkCopyError.md) |  | [optional] 
**Results** | Pointer to [**[]IpamsvcCopyResponse**](IpamsvcCopyResponse.md) |  | [optional] 

## Methods

### NewIpamsvcBulkCopyIPSpaceResponse

`func NewIpamsvcBulkCopyIPSpaceResponse() *IpamsvcBulkCopyIPSpaceResponse`

NewIpamsvcBulkCopyIPSpaceResponse instantiates a new IpamsvcBulkCopyIPSpaceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcBulkCopyIPSpaceResponseWithDefaults

`func NewIpamsvcBulkCopyIPSpaceResponseWithDefaults() *IpamsvcBulkCopyIPSpaceResponse`

NewIpamsvcBulkCopyIPSpaceResponseWithDefaults instantiates a new IpamsvcBulkCopyIPSpaceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *IpamsvcBulkCopyIPSpaceResponse) GetErrors() []IpamsvcBulkCopyError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *IpamsvcBulkCopyIPSpaceResponse) GetErrorsOk() (*[]IpamsvcBulkCopyError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *IpamsvcBulkCopyIPSpaceResponse) SetErrors(v []IpamsvcBulkCopyError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *IpamsvcBulkCopyIPSpaceResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetResults

`func (o *IpamsvcBulkCopyIPSpaceResponse) GetResults() []IpamsvcCopyResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *IpamsvcBulkCopyIPSpaceResponse) GetResultsOk() (*[]IpamsvcCopyResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *IpamsvcBulkCopyIPSpaceResponse) SetResults(v []IpamsvcCopyResponse)`

SetResults sets Results field to given value.

### HasResults

`func (o *IpamsvcBulkCopyIPSpaceResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


