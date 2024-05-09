# ServiceV2CreateMaintenanceWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeferredWindow** | Pointer to [**ServiceV2DeferredWindow**](ServiceV2DeferredWindow.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ScheduledWindow** | Pointer to [**ServiceV2ScheduledWindow**](ServiceV2ScheduledWindow.md) |  | [optional] 
**Tags** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**WindowType** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceV2CreateMaintenanceWindow

`func NewServiceV2CreateMaintenanceWindow() *ServiceV2CreateMaintenanceWindow`

NewServiceV2CreateMaintenanceWindow instantiates a new ServiceV2CreateMaintenanceWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceV2CreateMaintenanceWindowWithDefaults

`func NewServiceV2CreateMaintenanceWindowWithDefaults() *ServiceV2CreateMaintenanceWindow`

NewServiceV2CreateMaintenanceWindowWithDefaults instantiates a new ServiceV2CreateMaintenanceWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeferredWindow

`func (o *ServiceV2CreateMaintenanceWindow) GetDeferredWindow() ServiceV2DeferredWindow`

GetDeferredWindow returns the DeferredWindow field if non-nil, zero value otherwise.

### GetDeferredWindowOk

`func (o *ServiceV2CreateMaintenanceWindow) GetDeferredWindowOk() (*ServiceV2DeferredWindow, bool)`

GetDeferredWindowOk returns a tuple with the DeferredWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferredWindow

`func (o *ServiceV2CreateMaintenanceWindow) SetDeferredWindow(v ServiceV2DeferredWindow)`

SetDeferredWindow sets DeferredWindow field to given value.

### HasDeferredWindow

`func (o *ServiceV2CreateMaintenanceWindow) HasDeferredWindow() bool`

HasDeferredWindow returns a boolean if a field has been set.

### GetId

`func (o *ServiceV2CreateMaintenanceWindow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceV2CreateMaintenanceWindow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceV2CreateMaintenanceWindow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceV2CreateMaintenanceWindow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScheduledWindow

`func (o *ServiceV2CreateMaintenanceWindow) GetScheduledWindow() ServiceV2ScheduledWindow`

GetScheduledWindow returns the ScheduledWindow field if non-nil, zero value otherwise.

### GetScheduledWindowOk

`func (o *ServiceV2CreateMaintenanceWindow) GetScheduledWindowOk() (*ServiceV2ScheduledWindow, bool)`

GetScheduledWindowOk returns a tuple with the ScheduledWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledWindow

`func (o *ServiceV2CreateMaintenanceWindow) SetScheduledWindow(v ServiceV2ScheduledWindow)`

SetScheduledWindow sets ScheduledWindow field to given value.

### HasScheduledWindow

`func (o *ServiceV2CreateMaintenanceWindow) HasScheduledWindow() bool`

HasScheduledWindow returns a boolean if a field has been set.

### GetTags

`func (o *ServiceV2CreateMaintenanceWindow) GetTags() map[string]map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceV2CreateMaintenanceWindow) GetTagsOk() (*map[string]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceV2CreateMaintenanceWindow) SetTags(v map[string]map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServiceV2CreateMaintenanceWindow) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetWindowType

`func (o *ServiceV2CreateMaintenanceWindow) GetWindowType() string`

GetWindowType returns the WindowType field if non-nil, zero value otherwise.

### GetWindowTypeOk

`func (o *ServiceV2CreateMaintenanceWindow) GetWindowTypeOk() (*string, bool)`

GetWindowTypeOk returns a tuple with the WindowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowType

`func (o *ServiceV2CreateMaintenanceWindow) SetWindowType(v string)`

SetWindowType sets WindowType field to given value.

### HasWindowType

`func (o *ServiceV2CreateMaintenanceWindow) HasWindowType() bool`

HasWindowType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


