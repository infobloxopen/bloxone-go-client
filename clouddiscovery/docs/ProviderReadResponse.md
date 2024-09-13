# ProviderReadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**DiscoveryConfig**](DiscoveryConfig.md) |  | [optional] 

## Methods

### NewProviderReadResponse

`func NewProviderReadResponse() *ProviderReadResponse`

NewProviderReadResponse instantiates a new ProviderReadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderReadResponseWithDefaults

`func NewProviderReadResponseWithDefaults() *ProviderReadResponse`

NewProviderReadResponseWithDefaults instantiates a new ProviderReadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ProviderReadResponse) GetResult() DiscoveryConfig`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ProviderReadResponse) GetResultOk() (*DiscoveryConfig, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ProviderReadResponse) SetResult(v DiscoveryConfig)`

SetResult sets Result field to given value.

### HasResult

`func (o *ProviderReadResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


