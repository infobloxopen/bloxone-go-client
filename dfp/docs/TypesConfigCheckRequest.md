# TypesConfigCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int32** |  | [optional] 
**ConfigCheckType** | Pointer to **string** |  | [optional] 
**Notify** | Pointer to **bool** | Caller sets to false if wants to suppress notification. | [optional] 

## Methods

### NewTypesConfigCheckRequest

`func NewTypesConfigCheckRequest() *TypesConfigCheckRequest`

NewTypesConfigCheckRequest instantiates a new TypesConfigCheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesConfigCheckRequestWithDefaults

`func NewTypesConfigCheckRequestWithDefaults() *TypesConfigCheckRequest`

NewTypesConfigCheckRequestWithDefaults instantiates a new TypesConfigCheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *TypesConfigCheckRequest) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TypesConfigCheckRequest) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TypesConfigCheckRequest) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *TypesConfigCheckRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetConfigCheckType

`func (o *TypesConfigCheckRequest) GetConfigCheckType() string`

GetConfigCheckType returns the ConfigCheckType field if non-nil, zero value otherwise.

### GetConfigCheckTypeOk

`func (o *TypesConfigCheckRequest) GetConfigCheckTypeOk() (*string, bool)`

GetConfigCheckTypeOk returns a tuple with the ConfigCheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCheckType

`func (o *TypesConfigCheckRequest) SetConfigCheckType(v string)`

SetConfigCheckType sets ConfigCheckType field to given value.

### HasConfigCheckType

`func (o *TypesConfigCheckRequest) HasConfigCheckType() bool`

HasConfigCheckType returns a boolean if a field has been set.

### GetNotify

`func (o *TypesConfigCheckRequest) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *TypesConfigCheckRequest) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *TypesConfigCheckRequest) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *TypesConfigCheckRequest) HasNotify() bool`

HasNotify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


