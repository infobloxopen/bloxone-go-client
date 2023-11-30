# KeysListTSIGKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]KeysTSIGKey**](KeysTSIGKey.md) | The list of TSIGKey objects. | [optional] 

## Methods

### NewKeysListTSIGKeyResponse

`func NewKeysListTSIGKeyResponse() *KeysListTSIGKeyResponse`

NewKeysListTSIGKeyResponse instantiates a new KeysListTSIGKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeysListTSIGKeyResponseWithDefaults

`func NewKeysListTSIGKeyResponseWithDefaults() *KeysListTSIGKeyResponse`

NewKeysListTSIGKeyResponseWithDefaults instantiates a new KeysListTSIGKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *KeysListTSIGKeyResponse) GetResults() []KeysTSIGKey`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *KeysListTSIGKeyResponse) GetResultsOk() (*[]KeysTSIGKey, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *KeysListTSIGKeyResponse) SetResults(v []KeysTSIGKey)`

SetResults sets Results field to given value.

### HasResults

`func (o *KeysListTSIGKeyResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


