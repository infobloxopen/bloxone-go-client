# ApiPageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to **int32** | The service may optionally include the offset of the next page of resources. A null value indicates no more pages. | [optional] 
**PageToken** | Pointer to **string** | The service response should contain a string to indicate the next page of resources. A null value indicates no more pages. | [optional] 
**Size** | Pointer to **int32** | The service may optionally include the total number of resources being paged. | [optional] 

## Methods

### NewApiPageInfo

`func NewApiPageInfo() *ApiPageInfo`

NewApiPageInfo instantiates a new ApiPageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPageInfoWithDefaults

`func NewApiPageInfoWithDefaults() *ApiPageInfo`

NewApiPageInfoWithDefaults instantiates a new ApiPageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *ApiPageInfo) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ApiPageInfo) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ApiPageInfo) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ApiPageInfo) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetPageToken

`func (o *ApiPageInfo) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *ApiPageInfo) GetPageTokenOk() (*string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageToken

`func (o *ApiPageInfo) SetPageToken(v string)`

SetPageToken sets PageToken field to given value.

### HasPageToken

`func (o *ApiPageInfo) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### GetSize

`func (o *ApiPageInfo) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ApiPageInfo) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ApiPageInfo) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ApiPageInfo) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


