# ConfigListViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]ConfigView**](ConfigView.md) | List of View objects. | [optional] 

## Methods

### NewConfigListViewResponse

`func NewConfigListViewResponse() *ConfigListViewResponse`

NewConfigListViewResponse instantiates a new ConfigListViewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigListViewResponseWithDefaults

`func NewConfigListViewResponseWithDefaults() *ConfigListViewResponse`

NewConfigListViewResponseWithDefaults instantiates a new ConfigListViewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ConfigListViewResponse) GetResults() []ConfigView`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ConfigListViewResponse) GetResultsOk() (*[]ConfigView, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ConfigListViewResponse) SetResults(v []ConfigView)`

SetResults sets Results field to given value.

### HasResults

`func (o *ConfigListViewResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


