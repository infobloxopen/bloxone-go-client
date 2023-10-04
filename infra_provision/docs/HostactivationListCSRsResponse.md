# HostactivationListCSRsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**ApiPageInfo**](ApiPageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]HostactivationCSR**](HostactivationCSR.md) |  | [optional] 

## Methods

### NewHostactivationListCSRsResponse

`func NewHostactivationListCSRsResponse() *HostactivationListCSRsResponse`

NewHostactivationListCSRsResponse instantiates a new HostactivationListCSRsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostactivationListCSRsResponseWithDefaults

`func NewHostactivationListCSRsResponseWithDefaults() *HostactivationListCSRsResponse`

NewHostactivationListCSRsResponseWithDefaults instantiates a new HostactivationListCSRsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *HostactivationListCSRsResponse) GetPage() ApiPageInfo`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *HostactivationListCSRsResponse) GetPageOk() (*ApiPageInfo, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *HostactivationListCSRsResponse) SetPage(v ApiPageInfo)`

SetPage sets Page field to given value.

### HasPage

`func (o *HostactivationListCSRsResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetResults

`func (o *HostactivationListCSRsResponse) GetResults() []HostactivationCSR`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *HostactivationListCSRsResponse) GetResultsOk() (*[]HostactivationCSR, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *HostactivationListCSRsResponse) SetResults(v []HostactivationCSR)`

SetResults sets Results field to given value.

### HasResults

`func (o *HostactivationListCSRsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


