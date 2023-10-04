# ConfigReadGlobalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ConfigGlobal**](ConfigGlobal.md) |  | [optional] 

## Methods

### NewConfigReadGlobalResponse

`func NewConfigReadGlobalResponse() *ConfigReadGlobalResponse`

NewConfigReadGlobalResponse instantiates a new ConfigReadGlobalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigReadGlobalResponseWithDefaults

`func NewConfigReadGlobalResponseWithDefaults() *ConfigReadGlobalResponse`

NewConfigReadGlobalResponseWithDefaults instantiates a new ConfigReadGlobalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ConfigReadGlobalResponse) GetResult() ConfigGlobal`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ConfigReadGlobalResponse) GetResultOk() (*ConfigGlobal, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ConfigReadGlobalResponse) SetResult(v ConfigGlobal)`

SetResult sets Result field to given value.

### HasResult

`func (o *ConfigReadGlobalResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


