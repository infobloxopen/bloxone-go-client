# AsmEnableBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Indicates whether Automated Scope Management is enabled or not. | [optional] 
**EnableNotification** | Pointer to **bool** | Indicates whether sending notifications to the users is enabled or not. | [optional] 
**ReenableDate** | Pointer to **time.Time** | The date at which notifications will be re-enabled automatically. | [optional] 

## Methods

### NewAsmEnableBlock

`func NewAsmEnableBlock() *AsmEnableBlock`

NewAsmEnableBlock instantiates a new AsmEnableBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsmEnableBlockWithDefaults

`func NewAsmEnableBlockWithDefaults() *AsmEnableBlock`

NewAsmEnableBlockWithDefaults instantiates a new AsmEnableBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *AsmEnableBlock) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *AsmEnableBlock) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *AsmEnableBlock) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *AsmEnableBlock) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableNotification

`func (o *AsmEnableBlock) GetEnableNotification() bool`

GetEnableNotification returns the EnableNotification field if non-nil, zero value otherwise.

### GetEnableNotificationOk

`func (o *AsmEnableBlock) GetEnableNotificationOk() (*bool, bool)`

GetEnableNotificationOk returns a tuple with the EnableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNotification

`func (o *AsmEnableBlock) SetEnableNotification(v bool)`

SetEnableNotification sets EnableNotification field to given value.

### HasEnableNotification

`func (o *AsmEnableBlock) HasEnableNotification() bool`

HasEnableNotification returns a boolean if a field has been set.

### GetReenableDate

`func (o *AsmEnableBlock) GetReenableDate() time.Time`

GetReenableDate returns the ReenableDate field if non-nil, zero value otherwise.

### GetReenableDateOk

`func (o *AsmEnableBlock) GetReenableDateOk() (*time.Time, bool)`

GetReenableDateOk returns a tuple with the ReenableDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReenableDate

`func (o *AsmEnableBlock) SetReenableDate(v time.Time)`

SetReenableDate sets ReenableDate field to given value.

### HasReenableDate

`func (o *AsmEnableBlock) HasReenableDate() bool`

HasReenableDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


