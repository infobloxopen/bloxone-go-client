# GetNextAvailableDelegationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]Delegation**](Delegation.md) |  | [optional] 

## Methods

### NewGetNextAvailableDelegationResponse

`func NewGetNextAvailableDelegationResponse() *GetNextAvailableDelegationResponse`

NewGetNextAvailableDelegationResponse instantiates a new GetNextAvailableDelegationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNextAvailableDelegationResponseWithDefaults

`func NewGetNextAvailableDelegationResponseWithDefaults() *GetNextAvailableDelegationResponse`

NewGetNextAvailableDelegationResponseWithDefaults instantiates a new GetNextAvailableDelegationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *GetNextAvailableDelegationResponse) GetResults() []Delegation`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetNextAvailableDelegationResponse) GetResultsOk() (*[]Delegation, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetNextAvailableDelegationResponse) SetResults(v []Delegation)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetNextAvailableDelegationResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


