# InfraService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configs** | Pointer to [**[]InfraServiceHostConfig**](InfraServiceHostConfig.md) | List of Host-specific configurations of this Service. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | Timestamp of creation of Service. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the Service (optional). | [optional] 
**DesiredState** | Pointer to **string** | The desired state of the Service. Should either be &#x60;\&quot;start\&quot;&#x60; or &#x60;\&quot;stop\&quot;&#x60;. | [optional] 
**DesiredVersion** | Pointer to **string** | The desired version of the Service. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InterfaceLabels** | Pointer to **[]string** | List of interfaces on which this Service can operate. Note: The list can contain custom interface labels (Example: &#x60;[\&quot;WAN\&quot;,\&quot;LAN\&quot;,\&quot;label1\&quot;,\&quot;label2\&quot;]&#x60;) | [optional] 
**Name** | **string** | The name of the Service (unique). | 
**PoolId** | **string** | The resource identifier. | 
**ServiceType** | **string** | The type of the Service deployed on the Host (&#x60;dns&#x60;, &#x60;cdc&#x60;, etc.). | 
**Tags** | Pointer to **map[string]map[string]interface{}** | Tags associated with this Service. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp of the latest update on Service. | [optional] [readonly] 

## Methods

### NewInfraService

`func NewInfraService(name string, poolId string, serviceType string, ) *InfraService`

NewInfraService instantiates a new InfraService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfraServiceWithDefaults

`func NewInfraServiceWithDefaults() *InfraService`

NewInfraServiceWithDefaults instantiates a new InfraService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigs

`func (o *InfraService) GetConfigs() []InfraServiceHostConfig`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *InfraService) GetConfigsOk() (*[]InfraServiceHostConfig, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *InfraService) SetConfigs(v []InfraServiceHostConfig)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *InfraService) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InfraService) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InfraService) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InfraService) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InfraService) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *InfraService) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InfraService) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InfraService) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InfraService) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDesiredState

`func (o *InfraService) GetDesiredState() string`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *InfraService) GetDesiredStateOk() (*string, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *InfraService) SetDesiredState(v string)`

SetDesiredState sets DesiredState field to given value.

### HasDesiredState

`func (o *InfraService) HasDesiredState() bool`

HasDesiredState returns a boolean if a field has been set.

### GetDesiredVersion

`func (o *InfraService) GetDesiredVersion() string`

GetDesiredVersion returns the DesiredVersion field if non-nil, zero value otherwise.

### GetDesiredVersionOk

`func (o *InfraService) GetDesiredVersionOk() (*string, bool)`

GetDesiredVersionOk returns a tuple with the DesiredVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredVersion

`func (o *InfraService) SetDesiredVersion(v string)`

SetDesiredVersion sets DesiredVersion field to given value.

### HasDesiredVersion

`func (o *InfraService) HasDesiredVersion() bool`

HasDesiredVersion returns a boolean if a field has been set.

### GetId

`func (o *InfraService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InfraService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InfraService) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InfraService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaceLabels

`func (o *InfraService) GetInterfaceLabels() []string`

GetInterfaceLabels returns the InterfaceLabels field if non-nil, zero value otherwise.

### GetInterfaceLabelsOk

`func (o *InfraService) GetInterfaceLabelsOk() (*[]string, bool)`

GetInterfaceLabelsOk returns a tuple with the InterfaceLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceLabels

`func (o *InfraService) SetInterfaceLabels(v []string)`

SetInterfaceLabels sets InterfaceLabels field to given value.

### HasInterfaceLabels

`func (o *InfraService) HasInterfaceLabels() bool`

HasInterfaceLabels returns a boolean if a field has been set.

### GetName

`func (o *InfraService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InfraService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InfraService) SetName(v string)`

SetName sets Name field to given value.


### GetPoolId

`func (o *InfraService) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *InfraService) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *InfraService) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetServiceType

`func (o *InfraService) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *InfraService) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *InfraService) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.


### GetTags

`func (o *InfraService) GetTags() map[string]map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InfraService) GetTagsOk() (*map[string]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InfraService) SetTags(v map[string]map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *InfraService) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InfraService) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InfraService) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InfraService) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InfraService) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


