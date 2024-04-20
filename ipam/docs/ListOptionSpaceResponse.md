# ListOptionSpaceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]OptionSpace**](OptionSpace.md) | The list of OptionSpace objects. | [optional] 

## Methods

### NewListOptionSpaceResponse

`func NewListOptionSpaceResponse() *ListOptionSpaceResponse`

NewListOptionSpaceResponse instantiates a new ListOptionSpaceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOptionSpaceResponseWithDefaults

`func NewListOptionSpaceResponseWithDefaults() *ListOptionSpaceResponse`

NewListOptionSpaceResponseWithDefaults instantiates a new ListOptionSpaceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ListOptionSpaceResponse) GetResults() []OptionSpace`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListOptionSpaceResponse) GetResultsOk() (*[]OptionSpace, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListOptionSpaceResponse) SetResults(v []OptionSpace)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListOptionSpaceResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


