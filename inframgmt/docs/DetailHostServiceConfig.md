# DetailHostServiceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentVersion** | Pointer to **string** | The current version of the Service deployed on the Host. | [optional] 
**ServiceId** | Pointer to **string** | The resource identifier. | [optional] 
**ServiceName** | Pointer to **string** | The name of the Service. | [optional] 
**ServiceType** | Pointer to **string** | The type of the Service deployed on the Host (&#x60;dns&#x60;, &#x60;cdc&#x60;, etc.). | [optional] 
**Status** | Pointer to [**ShortServiceStatus**](ShortServiceStatus.md) | Service status information. | [optional] 
**UpgradedAt** | Pointer to **time.Time** | The timestamp of the latest upgrade of the Host-specific Service configuration. | [optional] 

## Methods

### NewDetailHostServiceConfig

`func NewDetailHostServiceConfig() *DetailHostServiceConfig`

NewDetailHostServiceConfig instantiates a new DetailHostServiceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailHostServiceConfigWithDefaults

`func NewDetailHostServiceConfigWithDefaults() *DetailHostServiceConfig`

NewDetailHostServiceConfigWithDefaults instantiates a new DetailHostServiceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentVersion

`func (o *DetailHostServiceConfig) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *DetailHostServiceConfig) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *DetailHostServiceConfig) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *DetailHostServiceConfig) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetServiceId

`func (o *DetailHostServiceConfig) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DetailHostServiceConfig) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DetailHostServiceConfig) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *DetailHostServiceConfig) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *DetailHostServiceConfig) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *DetailHostServiceConfig) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *DetailHostServiceConfig) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *DetailHostServiceConfig) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServiceType

`func (o *DetailHostServiceConfig) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *DetailHostServiceConfig) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *DetailHostServiceConfig) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *DetailHostServiceConfig) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetStatus

`func (o *DetailHostServiceConfig) GetStatus() ShortServiceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DetailHostServiceConfig) GetStatusOk() (*ShortServiceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DetailHostServiceConfig) SetStatus(v ShortServiceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DetailHostServiceConfig) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpgradedAt

`func (o *DetailHostServiceConfig) GetUpgradedAt() time.Time`

GetUpgradedAt returns the UpgradedAt field if non-nil, zero value otherwise.

### GetUpgradedAtOk

`func (o *DetailHostServiceConfig) GetUpgradedAtOk() (*time.Time, bool)`

GetUpgradedAtOk returns a tuple with the UpgradedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradedAt

`func (o *DetailHostServiceConfig) SetUpgradedAt(v time.Time)`

SetUpgradedAt sets UpgradedAt field to given value.

### HasUpgradedAt

`func (o *DetailHostServiceConfig) HasUpgradedAt() bool`

HasUpgradedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


