# FederationReadDelegationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**FederationDelegation**](FederationDelegation.md) | The Delegation object. | [optional] 

## Methods

### NewFederationReadDelegationResponse

`func NewFederationReadDelegationResponse() *FederationReadDelegationResponse`

NewFederationReadDelegationResponse instantiates a new FederationReadDelegationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationReadDelegationResponseWithDefaults

`func NewFederationReadDelegationResponseWithDefaults() *FederationReadDelegationResponse`

NewFederationReadDelegationResponseWithDefaults instantiates a new FederationReadDelegationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *FederationReadDelegationResponse) GetResult() FederationDelegation`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *FederationReadDelegationResponse) GetResultOk() (*FederationDelegation, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *FederationReadDelegationResponse) SetResult(v FederationDelegation)`

SetResult sets Result field to given value.

### HasResult

`func (o *FederationReadDelegationResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


