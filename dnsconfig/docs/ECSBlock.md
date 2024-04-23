# ECSBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EcsEnabled** | Pointer to **bool** | Optional. Field config for _ecs_enabled_ field. | [optional] 
**EcsForwarding** | Pointer to **bool** | Optional. Field config for _ecs_forwarding_ field. | [optional] 
**EcsPrefixV4** | Pointer to **int64** | Optional. Field config for _ecs_prefix_v4_ field. | [optional] 
**EcsPrefixV6** | Pointer to **int64** | Optional. Field config for _ecs_prefix_v6_ field. | [optional] 
**EcsZones** | Pointer to [**[]ECSZone**](ECSZone.md) | Optional. Field config for _ecs_zones_ field. | [optional] 

## Methods

### NewECSBlock

`func NewECSBlock() *ECSBlock`

NewECSBlock instantiates a new ECSBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewECSBlockWithDefaults

`func NewECSBlockWithDefaults() *ECSBlock`

NewECSBlockWithDefaults instantiates a new ECSBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEcsEnabled

`func (o *ECSBlock) GetEcsEnabled() bool`

GetEcsEnabled returns the EcsEnabled field if non-nil, zero value otherwise.

### GetEcsEnabledOk

`func (o *ECSBlock) GetEcsEnabledOk() (*bool, bool)`

GetEcsEnabledOk returns a tuple with the EcsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsEnabled

`func (o *ECSBlock) SetEcsEnabled(v bool)`

SetEcsEnabled sets EcsEnabled field to given value.

### HasEcsEnabled

`func (o *ECSBlock) HasEcsEnabled() bool`

HasEcsEnabled returns a boolean if a field has been set.

### GetEcsForwarding

`func (o *ECSBlock) GetEcsForwarding() bool`

GetEcsForwarding returns the EcsForwarding field if non-nil, zero value otherwise.

### GetEcsForwardingOk

`func (o *ECSBlock) GetEcsForwardingOk() (*bool, bool)`

GetEcsForwardingOk returns a tuple with the EcsForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsForwarding

`func (o *ECSBlock) SetEcsForwarding(v bool)`

SetEcsForwarding sets EcsForwarding field to given value.

### HasEcsForwarding

`func (o *ECSBlock) HasEcsForwarding() bool`

HasEcsForwarding returns a boolean if a field has been set.

### GetEcsPrefixV4

`func (o *ECSBlock) GetEcsPrefixV4() int64`

GetEcsPrefixV4 returns the EcsPrefixV4 field if non-nil, zero value otherwise.

### GetEcsPrefixV4Ok

`func (o *ECSBlock) GetEcsPrefixV4Ok() (*int64, bool)`

GetEcsPrefixV4Ok returns a tuple with the EcsPrefixV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsPrefixV4

`func (o *ECSBlock) SetEcsPrefixV4(v int64)`

SetEcsPrefixV4 sets EcsPrefixV4 field to given value.

### HasEcsPrefixV4

`func (o *ECSBlock) HasEcsPrefixV4() bool`

HasEcsPrefixV4 returns a boolean if a field has been set.

### GetEcsPrefixV6

`func (o *ECSBlock) GetEcsPrefixV6() int64`

GetEcsPrefixV6 returns the EcsPrefixV6 field if non-nil, zero value otherwise.

### GetEcsPrefixV6Ok

`func (o *ECSBlock) GetEcsPrefixV6Ok() (*int64, bool)`

GetEcsPrefixV6Ok returns a tuple with the EcsPrefixV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsPrefixV6

`func (o *ECSBlock) SetEcsPrefixV6(v int64)`

SetEcsPrefixV6 sets EcsPrefixV6 field to given value.

### HasEcsPrefixV6

`func (o *ECSBlock) HasEcsPrefixV6() bool`

HasEcsPrefixV6 returns a boolean if a field has been set.

### GetEcsZones

`func (o *ECSBlock) GetEcsZones() []ECSZone`

GetEcsZones returns the EcsZones field if non-nil, zero value otherwise.

### GetEcsZonesOk

`func (o *ECSBlock) GetEcsZonesOk() (*[]ECSZone, bool)`

GetEcsZonesOk returns a tuple with the EcsZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsZones

`func (o *ECSBlock) SetEcsZones(v []ECSZone)`

SetEcsZones sets EcsZones field to given value.

### HasEcsZones

`func (o *ECSBlock) HasEcsZones() bool`

HasEcsZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


