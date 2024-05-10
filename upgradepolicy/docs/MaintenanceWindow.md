# MaintenanceWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeferredWindow** | Pointer to [**DeferredWindow**](DeferredWindow.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ScheduledWindow** | Pointer to [**ScheduledWindow**](ScheduledWindow.md) |  | [optional] 
**Tags** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**WindowType** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewMaintenanceWindow

`func NewMaintenanceWindow() *MaintenanceWindow`

NewMaintenanceWindow instantiates a new MaintenanceWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceWindowWithDefaults

`func NewMaintenanceWindowWithDefaults() *MaintenanceWindow`

NewMaintenanceWindowWithDefaults instantiates a new MaintenanceWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeferredWindow

`func (o *MaintenanceWindow) GetDeferredWindow() DeferredWindow`

GetDeferredWindow returns the DeferredWindow field if non-nil, zero value otherwise.

### GetDeferredWindowOk

`func (o *MaintenanceWindow) GetDeferredWindowOk() (*DeferredWindow, bool)`

GetDeferredWindowOk returns a tuple with the DeferredWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferredWindow

`func (o *MaintenanceWindow) SetDeferredWindow(v DeferredWindow)`

SetDeferredWindow sets DeferredWindow field to given value.

### HasDeferredWindow

`func (o *MaintenanceWindow) HasDeferredWindow() bool`

HasDeferredWindow returns a boolean if a field has been set.

### GetId

`func (o *MaintenanceWindow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MaintenanceWindow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MaintenanceWindow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MaintenanceWindow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScheduledWindow

`func (o *MaintenanceWindow) GetScheduledWindow() ScheduledWindow`

GetScheduledWindow returns the ScheduledWindow field if non-nil, zero value otherwise.

### GetScheduledWindowOk

`func (o *MaintenanceWindow) GetScheduledWindowOk() (*ScheduledWindow, bool)`

GetScheduledWindowOk returns a tuple with the ScheduledWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledWindow

`func (o *MaintenanceWindow) SetScheduledWindow(v ScheduledWindow)`

SetScheduledWindow sets ScheduledWindow field to given value.

### HasScheduledWindow

`func (o *MaintenanceWindow) HasScheduledWindow() bool`

HasScheduledWindow returns a boolean if a field has been set.

### GetTags

`func (o *MaintenanceWindow) GetTags() map[string]map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MaintenanceWindow) GetTagsOk() (*map[string]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MaintenanceWindow) SetTags(v map[string]map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *MaintenanceWindow) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetWindowType

`func (o *MaintenanceWindow) GetWindowType() string`

GetWindowType returns the WindowType field if non-nil, zero value otherwise.

### GetWindowTypeOk

`func (o *MaintenanceWindow) GetWindowTypeOk() (*string, bool)`

GetWindowTypeOk returns a tuple with the WindowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowType

`func (o *MaintenanceWindow) SetWindowType(v string)`

SetWindowType sets WindowType field to given value.

### HasWindowType

`func (o *MaintenanceWindow) HasWindowType() bool`

HasWindowType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


