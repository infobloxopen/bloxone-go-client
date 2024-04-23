# BulkCopyView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthZoneConfig** | Pointer to [**AuthZoneConfig**](AuthZoneConfig.md) |  | [optional] 
**ForwardZoneConfig** | Pointer to [**ForwardZoneConfig**](ForwardZoneConfig.md) |  | [optional] 
**Recursive** | Pointer to **bool** | Indicates whether child objects should be copied or not.  Defaults to _false_. Reserved for future use. | [optional] 
**Resources** | **[]string** | The resource identifier. | 
**SecondaryZoneConfig** | Pointer to [**AuthZoneConfig**](AuthZoneConfig.md) |  | [optional] 
**SkipOnError** | Pointer to **bool** | Indicates whether copying should skip object in case of error and continue with next, or abort copying in case of error.  Defaults to _false_. | [optional] 
**Target** | **string** | The resource identifier. | 

## Methods

### NewBulkCopyView

`func NewBulkCopyView(resources []string, target string, ) *BulkCopyView`

NewBulkCopyView instantiates a new BulkCopyView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCopyViewWithDefaults

`func NewBulkCopyViewWithDefaults() *BulkCopyView`

NewBulkCopyViewWithDefaults instantiates a new BulkCopyView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthZoneConfig

`func (o *BulkCopyView) GetAuthZoneConfig() AuthZoneConfig`

GetAuthZoneConfig returns the AuthZoneConfig field if non-nil, zero value otherwise.

### GetAuthZoneConfigOk

`func (o *BulkCopyView) GetAuthZoneConfigOk() (*AuthZoneConfig, bool)`

GetAuthZoneConfigOk returns a tuple with the AuthZoneConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthZoneConfig

`func (o *BulkCopyView) SetAuthZoneConfig(v AuthZoneConfig)`

SetAuthZoneConfig sets AuthZoneConfig field to given value.

### HasAuthZoneConfig

`func (o *BulkCopyView) HasAuthZoneConfig() bool`

HasAuthZoneConfig returns a boolean if a field has been set.

### GetForwardZoneConfig

`func (o *BulkCopyView) GetForwardZoneConfig() ForwardZoneConfig`

GetForwardZoneConfig returns the ForwardZoneConfig field if non-nil, zero value otherwise.

### GetForwardZoneConfigOk

`func (o *BulkCopyView) GetForwardZoneConfigOk() (*ForwardZoneConfig, bool)`

GetForwardZoneConfigOk returns a tuple with the ForwardZoneConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardZoneConfig

`func (o *BulkCopyView) SetForwardZoneConfig(v ForwardZoneConfig)`

SetForwardZoneConfig sets ForwardZoneConfig field to given value.

### HasForwardZoneConfig

`func (o *BulkCopyView) HasForwardZoneConfig() bool`

HasForwardZoneConfig returns a boolean if a field has been set.

### GetRecursive

`func (o *BulkCopyView) GetRecursive() bool`

GetRecursive returns the Recursive field if non-nil, zero value otherwise.

### GetRecursiveOk

`func (o *BulkCopyView) GetRecursiveOk() (*bool, bool)`

GetRecursiveOk returns a tuple with the Recursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursive

`func (o *BulkCopyView) SetRecursive(v bool)`

SetRecursive sets Recursive field to given value.

### HasRecursive

`func (o *BulkCopyView) HasRecursive() bool`

HasRecursive returns a boolean if a field has been set.

### GetResources

`func (o *BulkCopyView) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *BulkCopyView) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *BulkCopyView) SetResources(v []string)`

SetResources sets Resources field to given value.


### GetSecondaryZoneConfig

`func (o *BulkCopyView) GetSecondaryZoneConfig() AuthZoneConfig`

GetSecondaryZoneConfig returns the SecondaryZoneConfig field if non-nil, zero value otherwise.

### GetSecondaryZoneConfigOk

`func (o *BulkCopyView) GetSecondaryZoneConfigOk() (*AuthZoneConfig, bool)`

GetSecondaryZoneConfigOk returns a tuple with the SecondaryZoneConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryZoneConfig

`func (o *BulkCopyView) SetSecondaryZoneConfig(v AuthZoneConfig)`

SetSecondaryZoneConfig sets SecondaryZoneConfig field to given value.

### HasSecondaryZoneConfig

`func (o *BulkCopyView) HasSecondaryZoneConfig() bool`

HasSecondaryZoneConfig returns a boolean if a field has been set.

### GetSkipOnError

`func (o *BulkCopyView) GetSkipOnError() bool`

GetSkipOnError returns the SkipOnError field if non-nil, zero value otherwise.

### GetSkipOnErrorOk

`func (o *BulkCopyView) GetSkipOnErrorOk() (*bool, bool)`

GetSkipOnErrorOk returns a tuple with the SkipOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOnError

`func (o *BulkCopyView) SetSkipOnError(v bool)`

SetSkipOnError sets SkipOnError field to given value.

### HasSkipOnError

`func (o *BulkCopyView) HasSkipOnError() bool`

HasSkipOnError returns a boolean if a field has been set.

### GetTarget

`func (o *BulkCopyView) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *BulkCopyView) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *BulkCopyView) SetTarget(v string)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


