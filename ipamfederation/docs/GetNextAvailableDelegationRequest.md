# GetNextAvailableDelegationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | Pointer to [**NextAvailableDelegationArguments**](NextAvailableDelegationArguments.md) | The arguments which will be used as parameters while creating __Delegation__ object. | [optional] 
**Properties** | Pointer to [**NextAvailableDelegationProperties**](NextAvailableDelegationProperties.md) | The properties, which will be used to fill the created __Delegation__ object. | [optional] 

## Methods

### NewGetNextAvailableDelegationRequest

`func NewGetNextAvailableDelegationRequest() *GetNextAvailableDelegationRequest`

NewGetNextAvailableDelegationRequest instantiates a new GetNextAvailableDelegationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNextAvailableDelegationRequestWithDefaults

`func NewGetNextAvailableDelegationRequestWithDefaults() *GetNextAvailableDelegationRequest`

NewGetNextAvailableDelegationRequestWithDefaults instantiates a new GetNextAvailableDelegationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *GetNextAvailableDelegationRequest) GetArguments() NextAvailableDelegationArguments`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *GetNextAvailableDelegationRequest) GetArgumentsOk() (*NextAvailableDelegationArguments, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *GetNextAvailableDelegationRequest) SetArguments(v NextAvailableDelegationArguments)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *GetNextAvailableDelegationRequest) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetProperties

`func (o *GetNextAvailableDelegationRequest) GetProperties() NextAvailableDelegationProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *GetNextAvailableDelegationRequest) GetPropertiesOk() (*NextAvailableDelegationProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *GetNextAvailableDelegationRequest) SetProperties(v NextAvailableDelegationProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *GetNextAvailableDelegationRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


