# ListOptionGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]OptionGroup**](OptionGroup.md) | A list of OptionGroup objects. | [optional] 

## Methods

### NewListOptionGroupResponse

`func NewListOptionGroupResponse() *ListOptionGroupResponse`

NewListOptionGroupResponse instantiates a new ListOptionGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOptionGroupResponseWithDefaults

`func NewListOptionGroupResponseWithDefaults() *ListOptionGroupResponse`

NewListOptionGroupResponseWithDefaults instantiates a new ListOptionGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ListOptionGroupResponse) GetResults() []OptionGroup`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListOptionGroupResponse) GetResultsOk() (*[]OptionGroup, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListOptionGroupResponse) SetResults(v []OptionGroup)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListOptionGroupResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


