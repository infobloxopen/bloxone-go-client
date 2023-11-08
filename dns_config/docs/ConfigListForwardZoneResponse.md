# ConfigListForwardZoneResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]ConfigForwardZone**](ConfigForwardZone.md) | List of Forward Zone objects. | [optional] 

## Methods

### NewConfigListForwardZoneResponse

`func NewConfigListForwardZoneResponse() *ConfigListForwardZoneResponse`

NewConfigListForwardZoneResponse instantiates a new ConfigListForwardZoneResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigListForwardZoneResponseWithDefaults

`func NewConfigListForwardZoneResponseWithDefaults() *ConfigListForwardZoneResponse`

NewConfigListForwardZoneResponseWithDefaults instantiates a new ConfigListForwardZoneResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ConfigListForwardZoneResponse) GetResults() []ConfigForwardZone`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ConfigListForwardZoneResponse) GetResultsOk() (*[]ConfigForwardZone, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ConfigListForwardZoneResponse) SetResults(v []ConfigForwardZone)`

SetResults sets Results field to given value.

### HasResults

`func (o *ConfigListForwardZoneResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


