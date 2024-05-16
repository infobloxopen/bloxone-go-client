# UpdateBatchMaintenanceWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeferredWindow** | Pointer to [**DeferredWindow**](DeferredWindow.md) |  | [optional] 
**MwId** | Pointer to **string** |  | [optional] 
**ScheduledWindow** | Pointer to [**ScheduledWindow**](ScheduledWindow.md) |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUpdateBatchMaintenanceWindow

`func NewUpdateBatchMaintenanceWindow() *UpdateBatchMaintenanceWindow`

NewUpdateBatchMaintenanceWindow instantiates a new UpdateBatchMaintenanceWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBatchMaintenanceWindowWithDefaults

`func NewUpdateBatchMaintenanceWindowWithDefaults() *UpdateBatchMaintenanceWindow`

NewUpdateBatchMaintenanceWindowWithDefaults instantiates a new UpdateBatchMaintenanceWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeferredWindow

`func (o *UpdateBatchMaintenanceWindow) GetDeferredWindow() DeferredWindow`

GetDeferredWindow returns the DeferredWindow field if non-nil, zero value otherwise.

### GetDeferredWindowOk

`func (o *UpdateBatchMaintenanceWindow) GetDeferredWindowOk() (*DeferredWindow, bool)`

GetDeferredWindowOk returns a tuple with the DeferredWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferredWindow

`func (o *UpdateBatchMaintenanceWindow) SetDeferredWindow(v DeferredWindow)`

SetDeferredWindow sets DeferredWindow field to given value.

### HasDeferredWindow

`func (o *UpdateBatchMaintenanceWindow) HasDeferredWindow() bool`

HasDeferredWindow returns a boolean if a field has been set.

### GetMwId

`func (o *UpdateBatchMaintenanceWindow) GetMwId() string`

GetMwId returns the MwId field if non-nil, zero value otherwise.

### GetMwIdOk

`func (o *UpdateBatchMaintenanceWindow) GetMwIdOk() (*string, bool)`

GetMwIdOk returns a tuple with the MwId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMwId

`func (o *UpdateBatchMaintenanceWindow) SetMwId(v string)`

SetMwId sets MwId field to given value.

### HasMwId

`func (o *UpdateBatchMaintenanceWindow) HasMwId() bool`

HasMwId returns a boolean if a field has been set.

### GetScheduledWindow

`func (o *UpdateBatchMaintenanceWindow) GetScheduledWindow() ScheduledWindow`

GetScheduledWindow returns the ScheduledWindow field if non-nil, zero value otherwise.

### GetScheduledWindowOk

`func (o *UpdateBatchMaintenanceWindow) GetScheduledWindowOk() (*ScheduledWindow, bool)`

GetScheduledWindowOk returns a tuple with the ScheduledWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledWindow

`func (o *UpdateBatchMaintenanceWindow) SetScheduledWindow(v ScheduledWindow)`

SetScheduledWindow sets ScheduledWindow field to given value.

### HasScheduledWindow

`func (o *UpdateBatchMaintenanceWindow) HasScheduledWindow() bool`

HasScheduledWindow returns a boolean if a field has been set.

### GetTags

`func (o *UpdateBatchMaintenanceWindow) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateBatchMaintenanceWindow) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateBatchMaintenanceWindow) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateBatchMaintenanceWindow) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


