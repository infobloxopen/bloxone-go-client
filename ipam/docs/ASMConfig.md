# ASMConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsmThreshold** | Pointer to **int64** | ASM shows the number of addresses forecast to be used _forecast_period_ days in the future, if it is greater than _asm_threshold_ percent * _dhcp_total_ (see _dhcp_utilization_) then the subnet is flagged. | [optional] [default to 90]
**Enable** | Pointer to **bool** | Indicates if Automated Scope Management is enabled. | [optional] [default to true]
**EnableNotification** | Pointer to **bool** | Indicates if ASM should send notifications to the user. | [optional] [default to true]
**ForecastPeriod** | Pointer to **int64** | The forecast period in days. | [optional] [default to 14]
**GrowthFactor** | Pointer to **int64** | Indicates the growth in the number or percentage of IP addresses. | [optional] [default to 20]
**GrowthType** | Pointer to **string** | The type of factor to use: _percent_ or _count_. | [optional] [default to "percent"]
**History** | Pointer to **int64** | The minimum amount of history needed before ASM can run on this subnet. | [optional] [default to 30]
**MinTotal** | Pointer to **int64** | The minimum size of range needed for ASM to run on this subnet. | [optional] [default to 10]
**MinUnused** | Pointer to **int64** | The minimum percentage of addresses that must be available outside of the DHCP ranges and fixed addresses when making a suggested change.. | [optional] [default to 10]
**ReenableDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewASMConfig

`func NewASMConfig() *ASMConfig`

NewASMConfig instantiates a new ASMConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewASMConfigWithDefaults

`func NewASMConfigWithDefaults() *ASMConfig`

NewASMConfigWithDefaults instantiates a new ASMConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsmThreshold

`func (o *ASMConfig) GetAsmThreshold() int64`

GetAsmThreshold returns the AsmThreshold field if non-nil, zero value otherwise.

### GetAsmThresholdOk

`func (o *ASMConfig) GetAsmThresholdOk() (*int64, bool)`

GetAsmThresholdOk returns a tuple with the AsmThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmThreshold

`func (o *ASMConfig) SetAsmThreshold(v int64)`

SetAsmThreshold sets AsmThreshold field to given value.

### HasAsmThreshold

`func (o *ASMConfig) HasAsmThreshold() bool`

HasAsmThreshold returns a boolean if a field has been set.

### GetEnable

`func (o *ASMConfig) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ASMConfig) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ASMConfig) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *ASMConfig) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableNotification

`func (o *ASMConfig) GetEnableNotification() bool`

GetEnableNotification returns the EnableNotification field if non-nil, zero value otherwise.

### GetEnableNotificationOk

`func (o *ASMConfig) GetEnableNotificationOk() (*bool, bool)`

GetEnableNotificationOk returns a tuple with the EnableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNotification

`func (o *ASMConfig) SetEnableNotification(v bool)`

SetEnableNotification sets EnableNotification field to given value.

### HasEnableNotification

`func (o *ASMConfig) HasEnableNotification() bool`

HasEnableNotification returns a boolean if a field has been set.

### GetForecastPeriod

`func (o *ASMConfig) GetForecastPeriod() int64`

GetForecastPeriod returns the ForecastPeriod field if non-nil, zero value otherwise.

### GetForecastPeriodOk

`func (o *ASMConfig) GetForecastPeriodOk() (*int64, bool)`

GetForecastPeriodOk returns a tuple with the ForecastPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastPeriod

`func (o *ASMConfig) SetForecastPeriod(v int64)`

SetForecastPeriod sets ForecastPeriod field to given value.

### HasForecastPeriod

`func (o *ASMConfig) HasForecastPeriod() bool`

HasForecastPeriod returns a boolean if a field has been set.

### GetGrowthFactor

`func (o *ASMConfig) GetGrowthFactor() int64`

GetGrowthFactor returns the GrowthFactor field if non-nil, zero value otherwise.

### GetGrowthFactorOk

`func (o *ASMConfig) GetGrowthFactorOk() (*int64, bool)`

GetGrowthFactorOk returns a tuple with the GrowthFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrowthFactor

`func (o *ASMConfig) SetGrowthFactor(v int64)`

SetGrowthFactor sets GrowthFactor field to given value.

### HasGrowthFactor

`func (o *ASMConfig) HasGrowthFactor() bool`

HasGrowthFactor returns a boolean if a field has been set.

### GetGrowthType

`func (o *ASMConfig) GetGrowthType() string`

GetGrowthType returns the GrowthType field if non-nil, zero value otherwise.

### GetGrowthTypeOk

`func (o *ASMConfig) GetGrowthTypeOk() (*string, bool)`

GetGrowthTypeOk returns a tuple with the GrowthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrowthType

`func (o *ASMConfig) SetGrowthType(v string)`

SetGrowthType sets GrowthType field to given value.

### HasGrowthType

`func (o *ASMConfig) HasGrowthType() bool`

HasGrowthType returns a boolean if a field has been set.

### GetHistory

`func (o *ASMConfig) GetHistory() int64`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *ASMConfig) GetHistoryOk() (*int64, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *ASMConfig) SetHistory(v int64)`

SetHistory sets History field to given value.

### HasHistory

`func (o *ASMConfig) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetMinTotal

`func (o *ASMConfig) GetMinTotal() int64`

GetMinTotal returns the MinTotal field if non-nil, zero value otherwise.

### GetMinTotalOk

`func (o *ASMConfig) GetMinTotalOk() (*int64, bool)`

GetMinTotalOk returns a tuple with the MinTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTotal

`func (o *ASMConfig) SetMinTotal(v int64)`

SetMinTotal sets MinTotal field to given value.

### HasMinTotal

`func (o *ASMConfig) HasMinTotal() bool`

HasMinTotal returns a boolean if a field has been set.

### GetMinUnused

`func (o *ASMConfig) GetMinUnused() int64`

GetMinUnused returns the MinUnused field if non-nil, zero value otherwise.

### GetMinUnusedOk

`func (o *ASMConfig) GetMinUnusedOk() (*int64, bool)`

GetMinUnusedOk returns a tuple with the MinUnused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUnused

`func (o *ASMConfig) SetMinUnused(v int64)`

SetMinUnused sets MinUnused field to given value.

### HasMinUnused

`func (o *ASMConfig) HasMinUnused() bool`

HasMinUnused returns a boolean if a field has been set.

### GetReenableDate

`func (o *ASMConfig) GetReenableDate() time.Time`

GetReenableDate returns the ReenableDate field if non-nil, zero value otherwise.

### GetReenableDateOk

`func (o *ASMConfig) GetReenableDateOk() (*time.Time, bool)`

GetReenableDateOk returns a tuple with the ReenableDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReenableDate

`func (o *ASMConfig) SetReenableDate(v time.Time)`

SetReenableDate sets ReenableDate field to given value.

### HasReenableDate

`func (o *ASMConfig) HasReenableDate() bool`

HasReenableDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


