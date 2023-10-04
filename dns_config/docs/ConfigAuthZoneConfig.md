# ConfigAuthZoneConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalPrimaries** | Pointer to [**[]ConfigExternalPrimary**](ConfigExternalPrimary.md) | Optional. DNS primaries external to BloxOne DDI. Order is not significant. | [optional] 
**ExternalSecondaries** | Pointer to [**[]ConfigExternalSecondary**](ConfigExternalSecondary.md) | DNS secondaries external to BloxOne DDI. Order is not significant. | [optional] 
**InternalSecondaries** | Pointer to [**[]ConfigInternalSecondary**](ConfigInternalSecondary.md) | Optional. BloxOne DDI hosts acting as internal secondaries. Order is not significant. | [optional] 
**Nsgs** | Pointer to **[]string** | The resource identifier. | [optional] 

## Methods

### NewConfigAuthZoneConfig

`func NewConfigAuthZoneConfig() *ConfigAuthZoneConfig`

NewConfigAuthZoneConfig instantiates a new ConfigAuthZoneConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigAuthZoneConfigWithDefaults

`func NewConfigAuthZoneConfigWithDefaults() *ConfigAuthZoneConfig`

NewConfigAuthZoneConfigWithDefaults instantiates a new ConfigAuthZoneConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalPrimaries

`func (o *ConfigAuthZoneConfig) GetExternalPrimaries() []ConfigExternalPrimary`

GetExternalPrimaries returns the ExternalPrimaries field if non-nil, zero value otherwise.

### GetExternalPrimariesOk

`func (o *ConfigAuthZoneConfig) GetExternalPrimariesOk() (*[]ConfigExternalPrimary, bool)`

GetExternalPrimariesOk returns a tuple with the ExternalPrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPrimaries

`func (o *ConfigAuthZoneConfig) SetExternalPrimaries(v []ConfigExternalPrimary)`

SetExternalPrimaries sets ExternalPrimaries field to given value.

### HasExternalPrimaries

`func (o *ConfigAuthZoneConfig) HasExternalPrimaries() bool`

HasExternalPrimaries returns a boolean if a field has been set.

### GetExternalSecondaries

`func (o *ConfigAuthZoneConfig) GetExternalSecondaries() []ConfigExternalSecondary`

GetExternalSecondaries returns the ExternalSecondaries field if non-nil, zero value otherwise.

### GetExternalSecondariesOk

`func (o *ConfigAuthZoneConfig) GetExternalSecondariesOk() (*[]ConfigExternalSecondary, bool)`

GetExternalSecondariesOk returns a tuple with the ExternalSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSecondaries

`func (o *ConfigAuthZoneConfig) SetExternalSecondaries(v []ConfigExternalSecondary)`

SetExternalSecondaries sets ExternalSecondaries field to given value.

### HasExternalSecondaries

`func (o *ConfigAuthZoneConfig) HasExternalSecondaries() bool`

HasExternalSecondaries returns a boolean if a field has been set.

### GetInternalSecondaries

`func (o *ConfigAuthZoneConfig) GetInternalSecondaries() []ConfigInternalSecondary`

GetInternalSecondaries returns the InternalSecondaries field if non-nil, zero value otherwise.

### GetInternalSecondariesOk

`func (o *ConfigAuthZoneConfig) GetInternalSecondariesOk() (*[]ConfigInternalSecondary, bool)`

GetInternalSecondariesOk returns a tuple with the InternalSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalSecondaries

`func (o *ConfigAuthZoneConfig) SetInternalSecondaries(v []ConfigInternalSecondary)`

SetInternalSecondaries sets InternalSecondaries field to given value.

### HasInternalSecondaries

`func (o *ConfigAuthZoneConfig) HasInternalSecondaries() bool`

HasInternalSecondaries returns a boolean if a field has been set.

### GetNsgs

`func (o *ConfigAuthZoneConfig) GetNsgs() []string`

GetNsgs returns the Nsgs field if non-nil, zero value otherwise.

### GetNsgsOk

`func (o *ConfigAuthZoneConfig) GetNsgsOk() (*[]string, bool)`

GetNsgsOk returns a tuple with the Nsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsgs

`func (o *ConfigAuthZoneConfig) SetNsgs(v []string)`

SetNsgs sets Nsgs field to given value.

### HasNsgs

`func (o *ConfigAuthZoneConfig) HasNsgs() bool`

HasNsgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


