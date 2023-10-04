# InfraDetailHostServiceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentVersion** | Pointer to **string** | The current version of the Service deployed on the Host. | [optional] 
**ServiceId** | Pointer to **string** | The resource identifier. | [optional] 
**ServiceName** | Pointer to **string** | The name of the Service. | [optional] 
**ServiceType** | Pointer to **string** | The type of the Service deployed on the Host (&#x60;dns&#x60;, &#x60;cdc&#x60;, etc.). | [optional] 
**Status** | Pointer to [**InfraShortServiceStatus**](InfraShortServiceStatus.md) |  | [optional] 
**UpgradedAt** | Pointer to **time.Time** | The timestamp of the latest upgrade of the Host-specific Service configuration. | [optional] 

## Methods

### NewInfraDetailHostServiceConfig

`func NewInfraDetailHostServiceConfig() *InfraDetailHostServiceConfig`

NewInfraDetailHostServiceConfig instantiates a new InfraDetailHostServiceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfraDetailHostServiceConfigWithDefaults

`func NewInfraDetailHostServiceConfigWithDefaults() *InfraDetailHostServiceConfig`

NewInfraDetailHostServiceConfigWithDefaults instantiates a new InfraDetailHostServiceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentVersion

`func (o *InfraDetailHostServiceConfig) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *InfraDetailHostServiceConfig) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *InfraDetailHostServiceConfig) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *InfraDetailHostServiceConfig) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetServiceId

`func (o *InfraDetailHostServiceConfig) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *InfraDetailHostServiceConfig) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *InfraDetailHostServiceConfig) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *InfraDetailHostServiceConfig) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *InfraDetailHostServiceConfig) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *InfraDetailHostServiceConfig) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *InfraDetailHostServiceConfig) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *InfraDetailHostServiceConfig) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServiceType

`func (o *InfraDetailHostServiceConfig) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *InfraDetailHostServiceConfig) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *InfraDetailHostServiceConfig) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *InfraDetailHostServiceConfig) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetStatus

`func (o *InfraDetailHostServiceConfig) GetStatus() InfraShortServiceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InfraDetailHostServiceConfig) GetStatusOk() (*InfraShortServiceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InfraDetailHostServiceConfig) SetStatus(v InfraShortServiceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InfraDetailHostServiceConfig) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpgradedAt

`func (o *InfraDetailHostServiceConfig) GetUpgradedAt() time.Time`

GetUpgradedAt returns the UpgradedAt field if non-nil, zero value otherwise.

### GetUpgradedAtOk

`func (o *InfraDetailHostServiceConfig) GetUpgradedAtOk() (*time.Time, bool)`

GetUpgradedAtOk returns a tuple with the UpgradedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradedAt

`func (o *InfraDetailHostServiceConfig) SetUpgradedAt(v time.Time)`

SetUpgradedAt sets UpgradedAt field to given value.

### HasUpgradedAt

`func (o *InfraDetailHostServiceConfig) HasUpgradedAt() bool`

HasUpgradedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


