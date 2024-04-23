# CreateJoinTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JoinToken** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**JoinToken**](JoinToken.md) |  | [optional] 

## Methods

### NewCreateJoinTokenResponse

`func NewCreateJoinTokenResponse() *CreateJoinTokenResponse`

NewCreateJoinTokenResponse instantiates a new CreateJoinTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateJoinTokenResponseWithDefaults

`func NewCreateJoinTokenResponseWithDefaults() *CreateJoinTokenResponse`

NewCreateJoinTokenResponseWithDefaults instantiates a new CreateJoinTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJoinToken

`func (o *CreateJoinTokenResponse) GetJoinToken() string`

GetJoinToken returns the JoinToken field if non-nil, zero value otherwise.

### GetJoinTokenOk

`func (o *CreateJoinTokenResponse) GetJoinTokenOk() (*string, bool)`

GetJoinTokenOk returns a tuple with the JoinToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinToken

`func (o *CreateJoinTokenResponse) SetJoinToken(v string)`

SetJoinToken sets JoinToken field to given value.

### HasJoinToken

`func (o *CreateJoinTokenResponse) HasJoinToken() bool`

HasJoinToken returns a boolean if a field has been set.

### GetResult

`func (o *CreateJoinTokenResponse) GetResult() JoinToken`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CreateJoinTokenResponse) GetResultOk() (*JoinToken, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CreateJoinTokenResponse) SetResult(v JoinToken)`

SetResult sets Result field to given value.

### HasResult

`func (o *CreateJoinTokenResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


