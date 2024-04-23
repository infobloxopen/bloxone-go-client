# CreateNextAvailableABResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]AddressBlock**](AddressBlock.md) | The list of Next Available Address Block objects. | [optional] 

## Methods

### NewCreateNextAvailableABResponse

`func NewCreateNextAvailableABResponse() *CreateNextAvailableABResponse`

NewCreateNextAvailableABResponse instantiates a new CreateNextAvailableABResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNextAvailableABResponseWithDefaults

`func NewCreateNextAvailableABResponseWithDefaults() *CreateNextAvailableABResponse`

NewCreateNextAvailableABResponseWithDefaults instantiates a new CreateNextAvailableABResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *CreateNextAvailableABResponse) GetResults() []AddressBlock`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CreateNextAvailableABResponse) GetResultsOk() (*[]AddressBlock, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CreateNextAvailableABResponse) SetResults(v []AddressBlock)`

SetResults sets Results field to given value.

### HasResults

`func (o *CreateNextAvailableABResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


