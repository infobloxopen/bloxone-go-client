# SubAccountListResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**ApiPageInfo**](ApiPageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]SubAccountV2**](SubAccountV2.md) |  | [optional] 

## Methods

### NewSubAccountListResponseV2

`func NewSubAccountListResponseV2() *SubAccountListResponseV2`

NewSubAccountListResponseV2 instantiates a new SubAccountListResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubAccountListResponseV2WithDefaults

`func NewSubAccountListResponseV2WithDefaults() *SubAccountListResponseV2`

NewSubAccountListResponseV2WithDefaults instantiates a new SubAccountListResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *SubAccountListResponseV2) GetPage() ApiPageInfo`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SubAccountListResponseV2) GetPageOk() (*ApiPageInfo, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SubAccountListResponseV2) SetPage(v ApiPageInfo)`

SetPage sets Page field to given value.

### HasPage

`func (o *SubAccountListResponseV2) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetResults

`func (o *SubAccountListResponseV2) GetResults() []SubAccountV2`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SubAccountListResponseV2) GetResultsOk() (*[]SubAccountV2, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SubAccountListResponseV2) SetResults(v []SubAccountV2)`

SetResults sets Results field to given value.

### HasResults

`func (o *SubAccountListResponseV2) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


