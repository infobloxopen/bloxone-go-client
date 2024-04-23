# ListCSRsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**ApiPageInfo**](ApiPageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]CSR**](CSR.md) |  | [optional] 

## Methods

### NewListCSRsResponse

`func NewListCSRsResponse() *ListCSRsResponse`

NewListCSRsResponse instantiates a new ListCSRsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCSRsResponseWithDefaults

`func NewListCSRsResponseWithDefaults() *ListCSRsResponse`

NewListCSRsResponseWithDefaults instantiates a new ListCSRsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListCSRsResponse) GetPage() ApiPageInfo`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListCSRsResponse) GetPageOk() (*ApiPageInfo, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListCSRsResponse) SetPage(v ApiPageInfo)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListCSRsResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetResults

`func (o *ListCSRsResponse) GetResults() []CSR`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListCSRsResponse) GetResultsOk() (*[]CSR, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListCSRsResponse) SetResults(v []CSR)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListCSRsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


