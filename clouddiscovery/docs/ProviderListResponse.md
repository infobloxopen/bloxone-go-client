# ProviderListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**ApiPageInfo**](ApiPageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]DiscoveryConfig**](DiscoveryConfig.md) |  | [optional] 

## Methods

### NewProviderListResponse

`func NewProviderListResponse() *ProviderListResponse`

NewProviderListResponse instantiates a new ProviderListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderListResponseWithDefaults

`func NewProviderListResponseWithDefaults() *ProviderListResponse`

NewProviderListResponseWithDefaults instantiates a new ProviderListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ProviderListResponse) GetPage() ApiPageInfo`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ProviderListResponse) GetPageOk() (*ApiPageInfo, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ProviderListResponse) SetPage(v ApiPageInfo)`

SetPage sets Page field to given value.

### HasPage

`func (o *ProviderListResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetResults

`func (o *ProviderListResponse) GetResults() []DiscoveryConfig`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ProviderListResponse) GetResultsOk() (*[]DiscoveryConfig, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ProviderListResponse) SetResults(v []DiscoveryConfig)`

SetResults sets Results field to given value.

### HasResults

`func (o *ProviderListResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


