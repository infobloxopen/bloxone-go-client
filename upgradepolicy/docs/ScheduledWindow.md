# ScheduledWindow

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

### NewScheduledWindow

`func NewScheduledWindow() *ScheduledWindow`

NewScheduledWindow instantiates a new ScheduledWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledWindowWithDefaults

`func NewScheduledWindowWithDefaults() *ScheduledWindow`

NewScheduledWindowWithDefaults instantiates a new ScheduledWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ScheduledWindow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ScheduledWindow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ScheduledWindow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ScheduledWindow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDuration

`func (o *ScheduledWindow) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ScheduledWindow) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ScheduledWindow) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ScheduledWindow) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEnabled

`func (o *ScheduledWindow) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ScheduledWindow) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ScheduledWindow) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ScheduledWindow) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStartTime

`func (o *ScheduledWindow) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ScheduledWindow) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ScheduledWindow) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ScheduledWindow) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ScheduledWindow) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ScheduledWindow) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ScheduledWindow) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ScheduledWindow) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWeekday

`func (o *ScheduledWindow) GetWeekday() int32`

GetWeekday returns the Weekday field if non-nil, zero value otherwise.

### GetWeekdayOk

`func (o *ScheduledWindow) GetWeekdayOk() (*int32, bool)`

GetWeekdayOk returns a tuple with the Weekday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekday

`func (o *ScheduledWindow) SetWeekday(v int32)`

SetWeekday sets Weekday field to given value.

### HasWeekday

`func (o *ScheduledWindow) HasWeekday() bool`

HasWeekday returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


