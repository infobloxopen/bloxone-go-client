# ConfigForwardZoneConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalForwarders** | Pointer to [**[]ConfigForwarder**](ConfigForwarder.md) | Optional. External DNS servers to forward to. Order is not significant. | [optional] 
**Hosts** | Pointer to **[]string** | The resource identifier. | [optional] 
**InternalForwarders** | Pointer to **[]string** | The resource identifier. | [optional] 
**Nsgs** | Pointer to **[]string** | The resource identifier. | [optional] 

## Methods

### NewConfigForwardZoneConfig

`func NewConfigForwardZoneConfig() *ConfigForwardZoneConfig`

NewConfigForwardZoneConfig instantiates a new ConfigForwardZoneConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigForwardZoneConfigWithDefaults

`func NewConfigForwardZoneConfigWithDefaults() *ConfigForwardZoneConfig`

NewConfigForwardZoneConfigWithDefaults instantiates a new ConfigForwardZoneConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalForwarders

`func (o *ConfigForwardZoneConfig) GetExternalForwarders() []ConfigForwarder`

GetExternalForwarders returns the ExternalForwarders field if non-nil, zero value otherwise.

### GetExternalForwardersOk

`func (o *ConfigForwardZoneConfig) GetExternalForwardersOk() (*[]ConfigForwarder, bool)`

GetExternalForwardersOk returns a tuple with the ExternalForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalForwarders

`func (o *ConfigForwardZoneConfig) SetExternalForwarders(v []ConfigForwarder)`

SetExternalForwarders sets ExternalForwarders field to given value.

### HasExternalForwarders

`func (o *ConfigForwardZoneConfig) HasExternalForwarders() bool`

HasExternalForwarders returns a boolean if a field has been set.

### GetHosts

`func (o *ConfigForwardZoneConfig) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *ConfigForwardZoneConfig) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *ConfigForwardZoneConfig) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *ConfigForwardZoneConfig) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetInternalForwarders

`func (o *ConfigForwardZoneConfig) GetInternalForwarders() []string`

GetInternalForwarders returns the InternalForwarders field if non-nil, zero value otherwise.

### GetInternalForwardersOk

`func (o *ConfigForwardZoneConfig) GetInternalForwardersOk() (*[]string, bool)`

GetInternalForwardersOk returns a tuple with the InternalForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalForwarders

`func (o *ConfigForwardZoneConfig) SetInternalForwarders(v []string)`

SetInternalForwarders sets InternalForwarders field to given value.

### HasInternalForwarders

`func (o *ConfigForwardZoneConfig) HasInternalForwarders() bool`

HasInternalForwarders returns a boolean if a field has been set.

### GetNsgs

`func (o *ConfigForwardZoneConfig) GetNsgs() []string`

GetNsgs returns the Nsgs field if non-nil, zero value otherwise.

### GetNsgsOk

`func (o *ConfigForwardZoneConfig) GetNsgsOk() (*[]string, bool)`

GetNsgsOk returns a tuple with the Nsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsgs

`func (o *ConfigForwardZoneConfig) SetNsgs(v []string)`

SetNsgs sets Nsgs field to given value.

### HasNsgs

`func (o *ConfigForwardZoneConfig) HasNsgs() bool`

HasNsgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


