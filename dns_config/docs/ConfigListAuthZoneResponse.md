# ConfigListAuthZoneResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]ConfigAuthZone**](ConfigAuthZone.md) | The list of Auth Zone objects. | [optional] 

## Methods

### NewConfigListAuthZoneResponse

`func NewConfigListAuthZoneResponse() *ConfigListAuthZoneResponse`

NewConfigListAuthZoneResponse instantiates a new ConfigListAuthZoneResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigListAuthZoneResponseWithDefaults

`func NewConfigListAuthZoneResponseWithDefaults() *ConfigListAuthZoneResponse`

NewConfigListAuthZoneResponseWithDefaults instantiates a new ConfigListAuthZoneResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ConfigListAuthZoneResponse) GetResults() []ConfigAuthZone`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ConfigListAuthZoneResponse) GetResultsOk() (*[]ConfigAuthZone, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ConfigListAuthZoneResponse) SetResults(v []ConfigAuthZone)`

SetResults sets Results field to given value.

### HasResults

`func (o *ConfigListAuthZoneResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


