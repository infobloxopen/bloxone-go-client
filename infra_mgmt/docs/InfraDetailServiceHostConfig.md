# InfraDetailServiceHostConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentVersion** | Pointer to **string** | The current version of the Service deployed on the Host. | [optional] 
**Status** | Pointer to [**InfraShortServiceStatus**](InfraShortServiceStatus.md) |  | [optional] 
**UpgradedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInfraDetailServiceHostConfig

`func NewInfraDetailServiceHostConfig() *InfraDetailServiceHostConfig`

NewInfraDetailServiceHostConfig instantiates a new InfraDetailServiceHostConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfraDetailServiceHostConfigWithDefaults

`func NewInfraDetailServiceHostConfigWithDefaults() *InfraDetailServiceHostConfig`

NewInfraDetailServiceHostConfigWithDefaults instantiates a new InfraDetailServiceHostConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentVersion

`func (o *InfraDetailServiceHostConfig) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *InfraDetailServiceHostConfig) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *InfraDetailServiceHostConfig) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *InfraDetailServiceHostConfig) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetStatus

`func (o *InfraDetailServiceHostConfig) GetStatus() InfraShortServiceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InfraDetailServiceHostConfig) GetStatusOk() (*InfraShortServiceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InfraDetailServiceHostConfig) SetStatus(v InfraShortServiceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InfraDetailServiceHostConfig) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpgradedAt

`func (o *InfraDetailServiceHostConfig) GetUpgradedAt() time.Time`

GetUpgradedAt returns the UpgradedAt field if non-nil, zero value otherwise.

### GetUpgradedAtOk

`func (o *InfraDetailServiceHostConfig) GetUpgradedAtOk() (*time.Time, bool)`

GetUpgradedAtOk returns a tuple with the UpgradedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradedAt

`func (o *InfraDetailServiceHostConfig) SetUpgradedAt(v time.Time)`

SetUpgradedAt sets UpgradedAt field to given value.

### HasUpgradedAt

`func (o *InfraDetailServiceHostConfig) HasUpgradedAt() bool`

HasUpgradedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


