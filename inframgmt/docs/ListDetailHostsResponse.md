# ListDetailHostsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**ApiPageInfo**](ApiPageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]DetailHost**](DetailHost.md) |  | [optional] 

## Methods

### NewListDetailHostsResponse

`func NewListDetailHostsResponse() *ListDetailHostsResponse`

NewListDetailHostsResponse instantiates a new ListDetailHostsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDetailHostsResponseWithDefaults

`func NewListDetailHostsResponseWithDefaults() *ListDetailHostsResponse`

NewListDetailHostsResponseWithDefaults instantiates a new ListDetailHostsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListDetailHostsResponse) GetPage() ApiPageInfo`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListDetailHostsResponse) GetPageOk() (*ApiPageInfo, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListDetailHostsResponse) SetPage(v ApiPageInfo)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListDetailHostsResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetResults

`func (o *ListDetailHostsResponse) GetResults() []DetailHost`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListDetailHostsResponse) GetResultsOk() (*[]DetailHost, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListDetailHostsResponse) SetResults(v []DetailHost)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListDetailHostsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


