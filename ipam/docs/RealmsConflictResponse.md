# RealmsConflictResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]RealmsConflict**](RealmsConflict.md) | List of conflicts across _ipam/ip_space_ objects. | [optional] 

## Methods

### NewRealmsConflictResponse

`func NewRealmsConflictResponse() *RealmsConflictResponse`

NewRealmsConflictResponse instantiates a new RealmsConflictResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealmsConflictResponseWithDefaults

`func NewRealmsConflictResponseWithDefaults() *RealmsConflictResponse`

NewRealmsConflictResponseWithDefaults instantiates a new RealmsConflictResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *RealmsConflictResponse) GetResults() []RealmsConflict`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RealmsConflictResponse) GetResultsOk() (*[]RealmsConflict, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RealmsConflictResponse) SetResults(v []RealmsConflict)`

SetResults sets Results field to given value.

### HasResults

`func (o *RealmsConflictResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


