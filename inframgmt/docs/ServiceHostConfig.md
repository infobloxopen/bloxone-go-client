# ServiceHostConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentVersion** | Pointer to **string** | The current version of the Service deployed on the Host. | [optional] 
**ExtraData** | Pointer to **string** | The field to carry any extra data specific to this configuration. | [optional] 
**HostId** | Pointer to **string** | The resource identifier. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**ServiceId** | Pointer to **string** | The resource identifier. | [optional] 
**ServiceType** | Pointer to **string** | The type of the Service deployed on the Host (&#x60;dns&#x60;, &#x60;cdc&#x60;, etc.). | [optional] 
**UpgradedAt** | Pointer to **time.Time** | The timestamp of the latest upgrade of the Host-specific Service configuration. | [optional] 

## Methods

### NewServiceHostConfig

`func NewServiceHostConfig() *ServiceHostConfig`

NewServiceHostConfig instantiates a new ServiceHostConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceHostConfigWithDefaults

`func NewServiceHostConfigWithDefaults() *ServiceHostConfig`

NewServiceHostConfigWithDefaults instantiates a new ServiceHostConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentVersion

`func (o *ServiceHostConfig) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *ServiceHostConfig) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *ServiceHostConfig) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *ServiceHostConfig) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetExtraData

`func (o *ServiceHostConfig) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *ServiceHostConfig) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *ServiceHostConfig) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *ServiceHostConfig) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### GetHostId

`func (o *ServiceHostConfig) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *ServiceHostConfig) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *ServiceHostConfig) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *ServiceHostConfig) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetId

`func (o *ServiceHostConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceHostConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceHostConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceHostConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetServiceId

`func (o *ServiceHostConfig) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServiceHostConfig) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServiceHostConfig) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ServiceHostConfig) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceType

`func (o *ServiceHostConfig) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *ServiceHostConfig) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *ServiceHostConfig) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *ServiceHostConfig) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetUpgradedAt

`func (o *ServiceHostConfig) GetUpgradedAt() time.Time`

GetUpgradedAt returns the UpgradedAt field if non-nil, zero value otherwise.

### GetUpgradedAtOk

`func (o *ServiceHostConfig) GetUpgradedAtOk() (*time.Time, bool)`

GetUpgradedAtOk returns a tuple with the UpgradedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradedAt

`func (o *ServiceHostConfig) SetUpgradedAt(v time.Time)`

SetUpgradedAt sets UpgradedAt field to given value.

### HasUpgradedAt

`func (o *ServiceHostConfig) HasUpgradedAt() bool`

HasUpgradedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


