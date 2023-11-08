# ConfigECSBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EcsEnabled** | Pointer to **bool** | Optional. Field config for _ecs_enabled_ field. | [optional] 
**EcsForwarding** | Pointer to **bool** | Optional. Field config for _ecs_forwarding_ field. | [optional] 
**EcsPrefixV4** | Pointer to **int64** | Optional. Field config for _ecs_prefix_v4_ field. | [optional] 
**EcsPrefixV6** | Pointer to **int64** | Optional. Field config for _ecs_prefix_v6_ field. | [optional] 
**EcsZones** | Pointer to [**[]ConfigECSZone**](ConfigECSZone.md) | Optional. Field config for _ecs_zones_ field. | [optional] 

## Methods

### NewConfigECSBlock

`func NewConfigECSBlock() *ConfigECSBlock`

NewConfigECSBlock instantiates a new ConfigECSBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigECSBlockWithDefaults

`func NewConfigECSBlockWithDefaults() *ConfigECSBlock`

NewConfigECSBlockWithDefaults instantiates a new ConfigECSBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEcsEnabled

`func (o *ConfigECSBlock) GetEcsEnabled() bool`

GetEcsEnabled returns the EcsEnabled field if non-nil, zero value otherwise.

### GetEcsEnabledOk

`func (o *ConfigECSBlock) GetEcsEnabledOk() (*bool, bool)`

GetEcsEnabledOk returns a tuple with the EcsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsEnabled

`func (o *ConfigECSBlock) SetEcsEnabled(v bool)`

SetEcsEnabled sets EcsEnabled field to given value.

### HasEcsEnabled

`func (o *ConfigECSBlock) HasEcsEnabled() bool`

HasEcsEnabled returns a boolean if a field has been set.

### GetEcsForwarding

`func (o *ConfigECSBlock) GetEcsForwarding() bool`

GetEcsForwarding returns the EcsForwarding field if non-nil, zero value otherwise.

### GetEcsForwardingOk

`func (o *ConfigECSBlock) GetEcsForwardingOk() (*bool, bool)`

GetEcsForwardingOk returns a tuple with the EcsForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsForwarding

`func (o *ConfigECSBlock) SetEcsForwarding(v bool)`

SetEcsForwarding sets EcsForwarding field to given value.

### HasEcsForwarding

`func (o *ConfigECSBlock) HasEcsForwarding() bool`

HasEcsForwarding returns a boolean if a field has been set.

### GetEcsPrefixV4

`func (o *ConfigECSBlock) GetEcsPrefixV4() int64`

GetEcsPrefixV4 returns the EcsPrefixV4 field if non-nil, zero value otherwise.

### GetEcsPrefixV4Ok

`func (o *ConfigECSBlock) GetEcsPrefixV4Ok() (*int64, bool)`

GetEcsPrefixV4Ok returns a tuple with the EcsPrefixV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsPrefixV4

`func (o *ConfigECSBlock) SetEcsPrefixV4(v int64)`

SetEcsPrefixV4 sets EcsPrefixV4 field to given value.

### HasEcsPrefixV4

`func (o *ConfigECSBlock) HasEcsPrefixV4() bool`

HasEcsPrefixV4 returns a boolean if a field has been set.

### GetEcsPrefixV6

`func (o *ConfigECSBlock) GetEcsPrefixV6() int64`

GetEcsPrefixV6 returns the EcsPrefixV6 field if non-nil, zero value otherwise.

### GetEcsPrefixV6Ok

`func (o *ConfigECSBlock) GetEcsPrefixV6Ok() (*int64, bool)`

GetEcsPrefixV6Ok returns a tuple with the EcsPrefixV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsPrefixV6

`func (o *ConfigECSBlock) SetEcsPrefixV6(v int64)`

SetEcsPrefixV6 sets EcsPrefixV6 field to given value.

### HasEcsPrefixV6

`func (o *ConfigECSBlock) HasEcsPrefixV6() bool`

HasEcsPrefixV6 returns a boolean if a field has been set.

### GetEcsZones

`func (o *ConfigECSBlock) GetEcsZones() []ConfigECSZone`

GetEcsZones returns the EcsZones field if non-nil, zero value otherwise.

### GetEcsZonesOk

`func (o *ConfigECSBlock) GetEcsZonesOk() (*[]ConfigECSZone, bool)`

GetEcsZonesOk returns a tuple with the EcsZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsZones

`func (o *ConfigECSBlock) SetEcsZones(v []ConfigECSZone)`

SetEcsZones sets EcsZones field to given value.

### HasEcsZones

`func (o *ConfigECSBlock) HasEcsZones() bool`

HasEcsZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


