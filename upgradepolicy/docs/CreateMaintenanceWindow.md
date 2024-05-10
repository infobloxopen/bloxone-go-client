# CreateMaintenanceWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeferredWindow** | Pointer to [**DeferredWindow**](DeferredWindow.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ScheduledWindow** | Pointer to [**ScheduledWindow**](ScheduledWindow.md) |  | [optional] 
**Tags** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**WindowType** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateMaintenanceWindow

`func NewCreateMaintenanceWindow() *CreateMaintenanceWindow`

NewCreateMaintenanceWindow instantiates a new CreateMaintenanceWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMaintenanceWindowWithDefaults

`func NewCreateMaintenanceWindowWithDefaults() *CreateMaintenanceWindow`

NewCreateMaintenanceWindowWithDefaults instantiates a new CreateMaintenanceWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeferredWindow

`func (o *CreateMaintenanceWindow) GetDeferredWindow() DeferredWindow`

GetDeferredWindow returns the DeferredWindow field if non-nil, zero value otherwise.

### GetDeferredWindowOk

`func (o *CreateMaintenanceWindow) GetDeferredWindowOk() (*DeferredWindow, bool)`

GetDeferredWindowOk returns a tuple with the DeferredWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferredWindow

`func (o *CreateMaintenanceWindow) SetDeferredWindow(v DeferredWindow)`

SetDeferredWindow sets DeferredWindow field to given value.

### HasDeferredWindow

`func (o *CreateMaintenanceWindow) HasDeferredWindow() bool`

HasDeferredWindow returns a boolean if a field has been set.

### GetId

`func (o *CreateMaintenanceWindow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateMaintenanceWindow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateMaintenanceWindow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateMaintenanceWindow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScheduledWindow

`func (o *CreateMaintenanceWindow) GetScheduledWindow() ScheduledWindow`

GetScheduledWindow returns the ScheduledWindow field if non-nil, zero value otherwise.

### GetScheduledWindowOk

`func (o *CreateMaintenanceWindow) GetScheduledWindowOk() (*ScheduledWindow, bool)`

GetScheduledWindowOk returns a tuple with the ScheduledWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledWindow

`func (o *CreateMaintenanceWindow) SetScheduledWindow(v ScheduledWindow)`

SetScheduledWindow sets ScheduledWindow field to given value.

### HasScheduledWindow

`func (o *CreateMaintenanceWindow) HasScheduledWindow() bool`

HasScheduledWindow returns a boolean if a field has been set.

### GetTags

`func (o *CreateMaintenanceWindow) GetTags() map[string]map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateMaintenanceWindow) GetTagsOk() (*map[string]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateMaintenanceWindow) SetTags(v map[string]map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateMaintenanceWindow) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetWindowType

`func (o *CreateMaintenanceWindow) GetWindowType() string`

GetWindowType returns the WindowType field if non-nil, zero value otherwise.

### GetWindowTypeOk

`func (o *CreateMaintenanceWindow) GetWindowTypeOk() (*string, bool)`

GetWindowTypeOk returns a tuple with the WindowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowType

`func (o *CreateMaintenanceWindow) SetWindowType(v string)`

SetWindowType sets WindowType field to given value.

### HasWindowType

`func (o *CreateMaintenanceWindow) HasWindowType() bool`

HasWindowType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


