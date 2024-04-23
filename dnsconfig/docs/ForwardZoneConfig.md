# ForwardZoneConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalForwarders** | Pointer to [**[]Forwarder**](Forwarder.md) | Optional. External DNS servers to forward to. Order is not significant. | [optional] 
**Hosts** | Pointer to **[]string** | The resource identifier. | [optional] 
**InternalForwarders** | Pointer to **[]string** | The resource identifier. | [optional] 
**Nsgs** | Pointer to **[]string** | The resource identifier. | [optional] 

## Methods

### NewForwardZoneConfig

`func NewForwardZoneConfig() *ForwardZoneConfig`

NewForwardZoneConfig instantiates a new ForwardZoneConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForwardZoneConfigWithDefaults

`func NewForwardZoneConfigWithDefaults() *ForwardZoneConfig`

NewForwardZoneConfigWithDefaults instantiates a new ForwardZoneConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalForwarders

`func (o *ForwardZoneConfig) GetExternalForwarders() []Forwarder`

GetExternalForwarders returns the ExternalForwarders field if non-nil, zero value otherwise.

### GetExternalForwardersOk

`func (o *ForwardZoneConfig) GetExternalForwardersOk() (*[]Forwarder, bool)`

GetExternalForwardersOk returns a tuple with the ExternalForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalForwarders

`func (o *ForwardZoneConfig) SetExternalForwarders(v []Forwarder)`

SetExternalForwarders sets ExternalForwarders field to given value.

### HasExternalForwarders

`func (o *ForwardZoneConfig) HasExternalForwarders() bool`

HasExternalForwarders returns a boolean if a field has been set.

### GetHosts

`func (o *ForwardZoneConfig) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *ForwardZoneConfig) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *ForwardZoneConfig) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *ForwardZoneConfig) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetInternalForwarders

`func (o *ForwardZoneConfig) GetInternalForwarders() []string`

GetInternalForwarders returns the InternalForwarders field if non-nil, zero value otherwise.

### GetInternalForwardersOk

`func (o *ForwardZoneConfig) GetInternalForwardersOk() (*[]string, bool)`

GetInternalForwardersOk returns a tuple with the InternalForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalForwarders

`func (o *ForwardZoneConfig) SetInternalForwarders(v []string)`

SetInternalForwarders sets InternalForwarders field to given value.

### HasInternalForwarders

`func (o *ForwardZoneConfig) HasInternalForwarders() bool`

HasInternalForwarders returns a boolean if a field has been set.

### GetNsgs

`func (o *ForwardZoneConfig) GetNsgs() []string`

GetNsgs returns the Nsgs field if non-nil, zero value otherwise.

### GetNsgsOk

`func (o *ForwardZoneConfig) GetNsgsOk() (*[]string, bool)`

GetNsgsOk returns a tuple with the Nsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsgs

`func (o *ForwardZoneConfig) SetNsgs(v []string)`

SetNsgs sets Nsgs field to given value.

### HasNsgs

`func (o *ForwardZoneConfig) HasNsgs() bool`

HasNsgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


