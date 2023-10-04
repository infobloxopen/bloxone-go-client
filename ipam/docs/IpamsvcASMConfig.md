# IpamsvcASMConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsmThreshold** | Pointer to **int64** | ASM shows the number of addresses forecast to be used _forecast_period_ days in the future, if it is greater than _asm_threshold_ percent * _dhcp_total_ (see _dhcp_utilization_) then the subnet is flagged. | [optional] 
**Enable** | Pointer to **NullableBool** | Indicates if Automated Scope Management is enabled. | [optional] 
**EnableNotification** | Pointer to **NullableBool** | Indicates if ASM should send notifications to the user. | [optional] 
**ForecastPeriod** | Pointer to **int64** | The forecast period in days. | [optional] 
**GrowthFactor** | Pointer to **int64** | Indicates the growth in the number or percentage of IP addresses. | [optional] 
**GrowthType** | Pointer to **string** | The type of factor to use: _percent_ or _count_. | [optional] 
**History** | Pointer to **int64** | The minimum amount of history needed before ASM can run on this subnet. | [optional] 
**MinTotal** | Pointer to **int64** | The minimum size of range needed for ASM to run on this subnet. | [optional] 
**MinUnused** | Pointer to **int64** | The minimum percentage of addresses that must be available outside of the DHCP ranges and fixed addresses when making a suggested change.. | [optional] 
**ReenableDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewIpamsvcASMConfig

`func NewIpamsvcASMConfig() *IpamsvcASMConfig`

NewIpamsvcASMConfig instantiates a new IpamsvcASMConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcASMConfigWithDefaults

`func NewIpamsvcASMConfigWithDefaults() *IpamsvcASMConfig`

NewIpamsvcASMConfigWithDefaults instantiates a new IpamsvcASMConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsmThreshold

`func (o *IpamsvcASMConfig) GetAsmThreshold() int64`

GetAsmThreshold returns the AsmThreshold field if non-nil, zero value otherwise.

### GetAsmThresholdOk

`func (o *IpamsvcASMConfig) GetAsmThresholdOk() (*int64, bool)`

GetAsmThresholdOk returns a tuple with the AsmThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmThreshold

`func (o *IpamsvcASMConfig) SetAsmThreshold(v int64)`

SetAsmThreshold sets AsmThreshold field to given value.

### HasAsmThreshold

`func (o *IpamsvcASMConfig) HasAsmThreshold() bool`

HasAsmThreshold returns a boolean if a field has been set.

### GetEnable

`func (o *IpamsvcASMConfig) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IpamsvcASMConfig) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IpamsvcASMConfig) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *IpamsvcASMConfig) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### SetEnableNil

`func (o *IpamsvcASMConfig) SetEnableNil(b bool)`

 SetEnableNil sets the value for Enable to be an explicit nil

### UnsetEnable
`func (o *IpamsvcASMConfig) UnsetEnable()`

UnsetEnable ensures that no value is present for Enable, not even an explicit nil
### GetEnableNotification

`func (o *IpamsvcASMConfig) GetEnableNotification() bool`

GetEnableNotification returns the EnableNotification field if non-nil, zero value otherwise.

### GetEnableNotificationOk

`func (o *IpamsvcASMConfig) GetEnableNotificationOk() (*bool, bool)`

GetEnableNotificationOk returns a tuple with the EnableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNotification

`func (o *IpamsvcASMConfig) SetEnableNotification(v bool)`

SetEnableNotification sets EnableNotification field to given value.

### HasEnableNotification

`func (o *IpamsvcASMConfig) HasEnableNotification() bool`

HasEnableNotification returns a boolean if a field has been set.

### SetEnableNotificationNil

`func (o *IpamsvcASMConfig) SetEnableNotificationNil(b bool)`

 SetEnableNotificationNil sets the value for EnableNotification to be an explicit nil

### UnsetEnableNotification
`func (o *IpamsvcASMConfig) UnsetEnableNotification()`

UnsetEnableNotification ensures that no value is present for EnableNotification, not even an explicit nil
### GetForecastPeriod

`func (o *IpamsvcASMConfig) GetForecastPeriod() int64`

GetForecastPeriod returns the ForecastPeriod field if non-nil, zero value otherwise.

### GetForecastPeriodOk

`func (o *IpamsvcASMConfig) GetForecastPeriodOk() (*int64, bool)`

GetForecastPeriodOk returns a tuple with the ForecastPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastPeriod

`func (o *IpamsvcASMConfig) SetForecastPeriod(v int64)`

SetForecastPeriod sets ForecastPeriod field to given value.

### HasForecastPeriod

`func (o *IpamsvcASMConfig) HasForecastPeriod() bool`

HasForecastPeriod returns a boolean if a field has been set.

### GetGrowthFactor

`func (o *IpamsvcASMConfig) GetGrowthFactor() int64`

GetGrowthFactor returns the GrowthFactor field if non-nil, zero value otherwise.

### GetGrowthFactorOk

`func (o *IpamsvcASMConfig) GetGrowthFactorOk() (*int64, bool)`

GetGrowthFactorOk returns a tuple with the GrowthFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrowthFactor

`func (o *IpamsvcASMConfig) SetGrowthFactor(v int64)`

SetGrowthFactor sets GrowthFactor field to given value.

### HasGrowthFactor

`func (o *IpamsvcASMConfig) HasGrowthFactor() bool`

HasGrowthFactor returns a boolean if a field has been set.

### GetGrowthType

`func (o *IpamsvcASMConfig) GetGrowthType() string`

GetGrowthType returns the GrowthType field if non-nil, zero value otherwise.

### GetGrowthTypeOk

`func (o *IpamsvcASMConfig) GetGrowthTypeOk() (*string, bool)`

GetGrowthTypeOk returns a tuple with the GrowthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrowthType

`func (o *IpamsvcASMConfig) SetGrowthType(v string)`

SetGrowthType sets GrowthType field to given value.

### HasGrowthType

`func (o *IpamsvcASMConfig) HasGrowthType() bool`

HasGrowthType returns a boolean if a field has been set.

### GetHistory

`func (o *IpamsvcASMConfig) GetHistory() int64`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *IpamsvcASMConfig) GetHistoryOk() (*int64, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *IpamsvcASMConfig) SetHistory(v int64)`

SetHistory sets History field to given value.

### HasHistory

`func (o *IpamsvcASMConfig) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetMinTotal

`func (o *IpamsvcASMConfig) GetMinTotal() int64`

GetMinTotal returns the MinTotal field if non-nil, zero value otherwise.

### GetMinTotalOk

`func (o *IpamsvcASMConfig) GetMinTotalOk() (*int64, bool)`

GetMinTotalOk returns a tuple with the MinTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTotal

`func (o *IpamsvcASMConfig) SetMinTotal(v int64)`

SetMinTotal sets MinTotal field to given value.

### HasMinTotal

`func (o *IpamsvcASMConfig) HasMinTotal() bool`

HasMinTotal returns a boolean if a field has been set.

### GetMinUnused

`func (o *IpamsvcASMConfig) GetMinUnused() int64`

GetMinUnused returns the MinUnused field if non-nil, zero value otherwise.

### GetMinUnusedOk

`func (o *IpamsvcASMConfig) GetMinUnusedOk() (*int64, bool)`

GetMinUnusedOk returns a tuple with the MinUnused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUnused

`func (o *IpamsvcASMConfig) SetMinUnused(v int64)`

SetMinUnused sets MinUnused field to given value.

### HasMinUnused

`func (o *IpamsvcASMConfig) HasMinUnused() bool`

HasMinUnused returns a boolean if a field has been set.

### GetReenableDate

`func (o *IpamsvcASMConfig) GetReenableDate() time.Time`

GetReenableDate returns the ReenableDate field if non-nil, zero value otherwise.

### GetReenableDateOk

`func (o *IpamsvcASMConfig) GetReenableDateOk() (*time.Time, bool)`

GetReenableDateOk returns a tuple with the ReenableDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReenableDate

`func (o *IpamsvcASMConfig) SetReenableDate(v time.Time)`

SetReenableDate sets ReenableDate field to given value.

### HasReenableDate

`func (o *IpamsvcASMConfig) HasReenableDate() bool`

HasReenableDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


