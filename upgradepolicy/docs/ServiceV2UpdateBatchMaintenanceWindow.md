# ServiceV2UpdateBatchMaintenanceWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeferredWindow** | Pointer to [**ServiceV2DeferredWindow**](ServiceV2DeferredWindow.md) |  | [optional] 
**MwId** | Pointer to **string** |  | [optional] 
**ScheduledWindow** | Pointer to [**ServiceV2ScheduledWindow**](ServiceV2ScheduledWindow.md) |  | [optional] 
**Tags** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewServiceV2UpdateBatchMaintenanceWindow

`func NewServiceV2UpdateBatchMaintenanceWindow() *ServiceV2UpdateBatchMaintenanceWindow`

NewServiceV2UpdateBatchMaintenanceWindow instantiates a new ServiceV2UpdateBatchMaintenanceWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceV2UpdateBatchMaintenanceWindowWithDefaults

`func NewServiceV2UpdateBatchMaintenanceWindowWithDefaults() *ServiceV2UpdateBatchMaintenanceWindow`

NewServiceV2UpdateBatchMaintenanceWindowWithDefaults instantiates a new ServiceV2UpdateBatchMaintenanceWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeferredWindow

`func (o *ServiceV2UpdateBatchMaintenanceWindow) GetDeferredWindow() ServiceV2DeferredWindow`

GetDeferredWindow returns the DeferredWindow field if non-nil, zero value otherwise.

### GetDeferredWindowOk

`func (o *ServiceV2UpdateBatchMaintenanceWindow) GetDeferredWindowOk() (*ServiceV2DeferredWindow, bool)`

GetDeferredWindowOk returns a tuple with the DeferredWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferredWindow

`func (o *ServiceV2UpdateBatchMaintenanceWindow) SetDeferredWindow(v ServiceV2DeferredWindow)`

SetDeferredWindow sets DeferredWindow field to given value.

### HasDeferredWindow

`func (o *ServiceV2UpdateBatchMaintenanceWindow) HasDeferredWindow() bool`

HasDeferredWindow returns a boolean if a field has been set.

### GetMwId

`func (o *ServiceV2UpdateBatchMaintenanceWindow) GetMwId() string`

GetMwId returns the MwId field if non-nil, zero value otherwise.

### GetMwIdOk

`func (o *ServiceV2UpdateBatchMaintenanceWindow) GetMwIdOk() (*string, bool)`

GetMwIdOk returns a tuple with the MwId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMwId

`func (o *ServiceV2UpdateBatchMaintenanceWindow) SetMwId(v string)`

SetMwId sets MwId field to given value.

### HasMwId

`func (o *ServiceV2UpdateBatchMaintenanceWindow) HasMwId() bool`

HasMwId returns a boolean if a field has been set.

### GetScheduledWindow

`func (o *ServiceV2UpdateBatchMaintenanceWindow) GetScheduledWindow() ServiceV2ScheduledWindow`

GetScheduledWindow returns the ScheduledWindow field if non-nil, zero value otherwise.

### GetScheduledWindowOk

`func (o *ServiceV2UpdateBatchMaintenanceWindow) GetScheduledWindowOk() (*ServiceV2ScheduledWindow, bool)`

GetScheduledWindowOk returns a tuple with the ScheduledWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledWindow

`func (o *ServiceV2UpdateBatchMaintenanceWindow) SetScheduledWindow(v ServiceV2ScheduledWindow)`

SetScheduledWindow sets ScheduledWindow field to given value.

### HasScheduledWindow

`func (o *ServiceV2UpdateBatchMaintenanceWindow) HasScheduledWindow() bool`

HasScheduledWindow returns a boolean if a field has been set.

### GetTags

`func (o *ServiceV2UpdateBatchMaintenanceWindow) GetTags() map[string]map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceV2UpdateBatchMaintenanceWindow) GetTagsOk() (*map[string]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceV2UpdateBatchMaintenanceWindow) SetTags(v map[string]map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServiceV2UpdateBatchMaintenanceWindow) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


