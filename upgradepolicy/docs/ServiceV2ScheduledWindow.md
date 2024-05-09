# ServiceV2ScheduledWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Duration** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**StartTime** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Weekday** | Pointer to **int32** |  | [optional] 

## Methods

### NewServiceV2ScheduledWindow

`func NewServiceV2ScheduledWindow() *ServiceV2ScheduledWindow`

NewServiceV2ScheduledWindow instantiates a new ServiceV2ScheduledWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceV2ScheduledWindowWithDefaults

`func NewServiceV2ScheduledWindowWithDefaults() *ServiceV2ScheduledWindow`

NewServiceV2ScheduledWindowWithDefaults instantiates a new ServiceV2ScheduledWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ServiceV2ScheduledWindow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceV2ScheduledWindow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceV2ScheduledWindow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceV2ScheduledWindow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDuration

`func (o *ServiceV2ScheduledWindow) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ServiceV2ScheduledWindow) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ServiceV2ScheduledWindow) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ServiceV2ScheduledWindow) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEnabled

`func (o *ServiceV2ScheduledWindow) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServiceV2ScheduledWindow) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServiceV2ScheduledWindow) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ServiceV2ScheduledWindow) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStartTime

`func (o *ServiceV2ScheduledWindow) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ServiceV2ScheduledWindow) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ServiceV2ScheduledWindow) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ServiceV2ScheduledWindow) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServiceV2ScheduledWindow) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceV2ScheduledWindow) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceV2ScheduledWindow) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServiceV2ScheduledWindow) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWeekday

`func (o *ServiceV2ScheduledWindow) GetWeekday() int32`

GetWeekday returns the Weekday field if non-nil, zero value otherwise.

### GetWeekdayOk

`func (o *ServiceV2ScheduledWindow) GetWeekdayOk() (*int32, bool)`

GetWeekdayOk returns a tuple with the Weekday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekday

`func (o *ServiceV2ScheduledWindow) SetWeekday(v int32)`

SetWeekday sets Weekday field to given value.

### HasWeekday

`func (o *ServiceV2ScheduledWindow) HasWeekday() bool`

HasWeekday returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


