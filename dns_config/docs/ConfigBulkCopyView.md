# ConfigBulkCopyView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthZoneConfig** | Pointer to [**ConfigAuthZoneConfig**](ConfigAuthZoneConfig.md) |  | [optional] 
**ForwardZoneConfig** | Pointer to [**ConfigForwardZoneConfig**](ConfigForwardZoneConfig.md) |  | [optional] 
**Recursive** | Pointer to **bool** | Indicates whether child objects should be copied or not.  Defaults to _false_. Reserved for future use. | [optional] 
**Resources** | **[]string** | The resource identifier. | 
**SecondaryZoneConfig** | Pointer to [**ConfigAuthZoneConfig**](ConfigAuthZoneConfig.md) |  | [optional] 
**SkipOnError** | Pointer to **bool** | Indicates whether copying should skip object in case of error and continue with next, or abort copying in case of error.  Defaults to _false_. | [optional] 
**Target** | **string** | The resource identifier. | 

## Methods

### NewConfigBulkCopyView

`func NewConfigBulkCopyView(resources []string, target string, ) *ConfigBulkCopyView`

NewConfigBulkCopyView instantiates a new ConfigBulkCopyView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigBulkCopyViewWithDefaults

`func NewConfigBulkCopyViewWithDefaults() *ConfigBulkCopyView`

NewConfigBulkCopyViewWithDefaults instantiates a new ConfigBulkCopyView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthZoneConfig

`func (o *ConfigBulkCopyView) GetAuthZoneConfig() ConfigAuthZoneConfig`

GetAuthZoneConfig returns the AuthZoneConfig field if non-nil, zero value otherwise.

### GetAuthZoneConfigOk

`func (o *ConfigBulkCopyView) GetAuthZoneConfigOk() (*ConfigAuthZoneConfig, bool)`

GetAuthZoneConfigOk returns a tuple with the AuthZoneConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthZoneConfig

`func (o *ConfigBulkCopyView) SetAuthZoneConfig(v ConfigAuthZoneConfig)`

SetAuthZoneConfig sets AuthZoneConfig field to given value.

### HasAuthZoneConfig

`func (o *ConfigBulkCopyView) HasAuthZoneConfig() bool`

HasAuthZoneConfig returns a boolean if a field has been set.

### GetForwardZoneConfig

`func (o *ConfigBulkCopyView) GetForwardZoneConfig() ConfigForwardZoneConfig`

GetForwardZoneConfig returns the ForwardZoneConfig field if non-nil, zero value otherwise.

### GetForwardZoneConfigOk

`func (o *ConfigBulkCopyView) GetForwardZoneConfigOk() (*ConfigForwardZoneConfig, bool)`

GetForwardZoneConfigOk returns a tuple with the ForwardZoneConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardZoneConfig

`func (o *ConfigBulkCopyView) SetForwardZoneConfig(v ConfigForwardZoneConfig)`

SetForwardZoneConfig sets ForwardZoneConfig field to given value.

### HasForwardZoneConfig

`func (o *ConfigBulkCopyView) HasForwardZoneConfig() bool`

HasForwardZoneConfig returns a boolean if a field has been set.

### GetRecursive

`func (o *ConfigBulkCopyView) GetRecursive() bool`

GetRecursive returns the Recursive field if non-nil, zero value otherwise.

### GetRecursiveOk

`func (o *ConfigBulkCopyView) GetRecursiveOk() (*bool, bool)`

GetRecursiveOk returns a tuple with the Recursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursive

`func (o *ConfigBulkCopyView) SetRecursive(v bool)`

SetRecursive sets Recursive field to given value.

### HasRecursive

`func (o *ConfigBulkCopyView) HasRecursive() bool`

HasRecursive returns a boolean if a field has been set.

### GetResources

`func (o *ConfigBulkCopyView) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ConfigBulkCopyView) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ConfigBulkCopyView) SetResources(v []string)`

SetResources sets Resources field to given value.


### GetSecondaryZoneConfig

`func (o *ConfigBulkCopyView) GetSecondaryZoneConfig() ConfigAuthZoneConfig`

GetSecondaryZoneConfig returns the SecondaryZoneConfig field if non-nil, zero value otherwise.

### GetSecondaryZoneConfigOk

`func (o *ConfigBulkCopyView) GetSecondaryZoneConfigOk() (*ConfigAuthZoneConfig, bool)`

GetSecondaryZoneConfigOk returns a tuple with the SecondaryZoneConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryZoneConfig

`func (o *ConfigBulkCopyView) SetSecondaryZoneConfig(v ConfigAuthZoneConfig)`

SetSecondaryZoneConfig sets SecondaryZoneConfig field to given value.

### HasSecondaryZoneConfig

`func (o *ConfigBulkCopyView) HasSecondaryZoneConfig() bool`

HasSecondaryZoneConfig returns a boolean if a field has been set.

### GetSkipOnError

`func (o *ConfigBulkCopyView) GetSkipOnError() bool`

GetSkipOnError returns the SkipOnError field if non-nil, zero value otherwise.

### GetSkipOnErrorOk

`func (o *ConfigBulkCopyView) GetSkipOnErrorOk() (*bool, bool)`

GetSkipOnErrorOk returns a tuple with the SkipOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOnError

`func (o *ConfigBulkCopyView) SetSkipOnError(v bool)`

SetSkipOnError sets SkipOnError field to given value.

### HasSkipOnError

`func (o *ConfigBulkCopyView) HasSkipOnError() bool`

HasSkipOnError returns a boolean if a field has been set.

### GetTarget

`func (o *ConfigBulkCopyView) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ConfigBulkCopyView) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ConfigBulkCopyView) SetTarget(v string)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


