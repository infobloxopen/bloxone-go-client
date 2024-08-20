# DdidnsrickettsAccountListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**ApiPageInfo**](ApiPageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]DdidnsrickettsAccount**](DdidnsrickettsAccount.md) |  | [optional] 

## Methods

### NewDdidnsrickettsAccountListResponse

`func NewDdidnsrickettsAccountListResponse() *DdidnsrickettsAccountListResponse`

NewDdidnsrickettsAccountListResponse instantiates a new DdidnsrickettsAccountListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdidnsrickettsAccountListResponseWithDefaults

`func NewDdidnsrickettsAccountListResponseWithDefaults() *DdidnsrickettsAccountListResponse`

NewDdidnsrickettsAccountListResponseWithDefaults instantiates a new DdidnsrickettsAccountListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *DdidnsrickettsAccountListResponse) GetPage() ApiPageInfo`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DdidnsrickettsAccountListResponse) GetPageOk() (*ApiPageInfo, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DdidnsrickettsAccountListResponse) SetPage(v ApiPageInfo)`

SetPage sets Page field to given value.

### HasPage

`func (o *DdidnsrickettsAccountListResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetResults

`func (o *DdidnsrickettsAccountListResponse) GetResults() []DdidnsrickettsAccount`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DdidnsrickettsAccountListResponse) GetResultsOk() (*[]DdidnsrickettsAccount, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DdidnsrickettsAccountListResponse) SetResults(v []DdidnsrickettsAccount)`

SetResults sets Results field to given value.

### HasResults

`func (o *DdidnsrickettsAccountListResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


