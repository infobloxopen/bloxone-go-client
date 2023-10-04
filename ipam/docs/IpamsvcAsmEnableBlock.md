# IpamsvcAsmEnableBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Indicates whether Automated Scope Management is enabled or not. | [optional] 
**EnableNotification** | Pointer to **bool** | Indicates whether sending notifications to the users is enabled or not. | [optional] 
**ReenableDate** | Pointer to **time.Time** | The date at which notifications will be re-enabled automatically. | [optional] 

## Methods

### NewIpamsvcAsmEnableBlock

`func NewIpamsvcAsmEnableBlock() *IpamsvcAsmEnableBlock`

NewIpamsvcAsmEnableBlock instantiates a new IpamsvcAsmEnableBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcAsmEnableBlockWithDefaults

`func NewIpamsvcAsmEnableBlockWithDefaults() *IpamsvcAsmEnableBlock`

NewIpamsvcAsmEnableBlockWithDefaults instantiates a new IpamsvcAsmEnableBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *IpamsvcAsmEnableBlock) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IpamsvcAsmEnableBlock) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IpamsvcAsmEnableBlock) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *IpamsvcAsmEnableBlock) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableNotification

`func (o *IpamsvcAsmEnableBlock) GetEnableNotification() bool`

GetEnableNotification returns the EnableNotification field if non-nil, zero value otherwise.

### GetEnableNotificationOk

`func (o *IpamsvcAsmEnableBlock) GetEnableNotificationOk() (*bool, bool)`

GetEnableNotificationOk returns a tuple with the EnableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNotification

`func (o *IpamsvcAsmEnableBlock) SetEnableNotification(v bool)`

SetEnableNotification sets EnableNotification field to given value.

### HasEnableNotification

`func (o *IpamsvcAsmEnableBlock) HasEnableNotification() bool`

HasEnableNotification returns a boolean if a field has been set.

### GetReenableDate

`func (o *IpamsvcAsmEnableBlock) GetReenableDate() time.Time`

GetReenableDate returns the ReenableDate field if non-nil, zero value otherwise.

### GetReenableDateOk

`func (o *IpamsvcAsmEnableBlock) GetReenableDateOk() (*time.Time, bool)`

GetReenableDateOk returns a tuple with the ReenableDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReenableDate

`func (o *IpamsvcAsmEnableBlock) SetReenableDate(v time.Time)`

SetReenableDate sets ReenableDate field to given value.

### HasReenableDate

`func (o *IpamsvcAsmEnableBlock) HasReenableDate() bool`

HasReenableDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


