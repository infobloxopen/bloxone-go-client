# IpamsvcInheritedASMConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsmEnableBlock** | Pointer to [**IpamsvcInheritedAsmEnableBlock**](IpamsvcInheritedAsmEnableBlock.md) |  | [optional] 
**AsmGrowthBlock** | Pointer to [**IpamsvcInheritedAsmGrowthBlock**](IpamsvcInheritedAsmGrowthBlock.md) |  | [optional] 
**AsmThreshold** | Pointer to [**InheritanceInheritedUInt32**](InheritanceInheritedUInt32.md) |  | [optional] 
**ForecastPeriod** | Pointer to [**InheritanceInheritedUInt32**](InheritanceInheritedUInt32.md) |  | [optional] 
**History** | Pointer to [**InheritanceInheritedUInt32**](InheritanceInheritedUInt32.md) |  | [optional] 
**MinTotal** | Pointer to [**InheritanceInheritedUInt32**](InheritanceInheritedUInt32.md) |  | [optional] 
**MinUnused** | Pointer to [**InheritanceInheritedUInt32**](InheritanceInheritedUInt32.md) |  | [optional] 

## Methods

### NewIpamsvcInheritedASMConfig

`func NewIpamsvcInheritedASMConfig() *IpamsvcInheritedASMConfig`

NewIpamsvcInheritedASMConfig instantiates a new IpamsvcInheritedASMConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcInheritedASMConfigWithDefaults

`func NewIpamsvcInheritedASMConfigWithDefaults() *IpamsvcInheritedASMConfig`

NewIpamsvcInheritedASMConfigWithDefaults instantiates a new IpamsvcInheritedASMConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsmEnableBlock

`func (o *IpamsvcInheritedASMConfig) GetAsmEnableBlock() IpamsvcInheritedAsmEnableBlock`

GetAsmEnableBlock returns the AsmEnableBlock field if non-nil, zero value otherwise.

### GetAsmEnableBlockOk

`func (o *IpamsvcInheritedASMConfig) GetAsmEnableBlockOk() (*IpamsvcInheritedAsmEnableBlock, bool)`

GetAsmEnableBlockOk returns a tuple with the AsmEnableBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmEnableBlock

`func (o *IpamsvcInheritedASMConfig) SetAsmEnableBlock(v IpamsvcInheritedAsmEnableBlock)`

SetAsmEnableBlock sets AsmEnableBlock field to given value.

### HasAsmEnableBlock

`func (o *IpamsvcInheritedASMConfig) HasAsmEnableBlock() bool`

HasAsmEnableBlock returns a boolean if a field has been set.

### GetAsmGrowthBlock

`func (o *IpamsvcInheritedASMConfig) GetAsmGrowthBlock() IpamsvcInheritedAsmGrowthBlock`

GetAsmGrowthBlock returns the AsmGrowthBlock field if non-nil, zero value otherwise.

### GetAsmGrowthBlockOk

`func (o *IpamsvcInheritedASMConfig) GetAsmGrowthBlockOk() (*IpamsvcInheritedAsmGrowthBlock, bool)`

GetAsmGrowthBlockOk returns a tuple with the AsmGrowthBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmGrowthBlock

`func (o *IpamsvcInheritedASMConfig) SetAsmGrowthBlock(v IpamsvcInheritedAsmGrowthBlock)`

SetAsmGrowthBlock sets AsmGrowthBlock field to given value.

### HasAsmGrowthBlock

`func (o *IpamsvcInheritedASMConfig) HasAsmGrowthBlock() bool`

HasAsmGrowthBlock returns a boolean if a field has been set.

### GetAsmThreshold

`func (o *IpamsvcInheritedASMConfig) GetAsmThreshold() InheritanceInheritedUInt32`

GetAsmThreshold returns the AsmThreshold field if non-nil, zero value otherwise.

### GetAsmThresholdOk

`func (o *IpamsvcInheritedASMConfig) GetAsmThresholdOk() (*InheritanceInheritedUInt32, bool)`

GetAsmThresholdOk returns a tuple with the AsmThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmThreshold

`func (o *IpamsvcInheritedASMConfig) SetAsmThreshold(v InheritanceInheritedUInt32)`

SetAsmThreshold sets AsmThreshold field to given value.

### HasAsmThreshold

`func (o *IpamsvcInheritedASMConfig) HasAsmThreshold() bool`

HasAsmThreshold returns a boolean if a field has been set.

### GetForecastPeriod

`func (o *IpamsvcInheritedASMConfig) GetForecastPeriod() InheritanceInheritedUInt32`

GetForecastPeriod returns the ForecastPeriod field if non-nil, zero value otherwise.

### GetForecastPeriodOk

`func (o *IpamsvcInheritedASMConfig) GetForecastPeriodOk() (*InheritanceInheritedUInt32, bool)`

GetForecastPeriodOk returns a tuple with the ForecastPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastPeriod

`func (o *IpamsvcInheritedASMConfig) SetForecastPeriod(v InheritanceInheritedUInt32)`

SetForecastPeriod sets ForecastPeriod field to given value.

### HasForecastPeriod

`func (o *IpamsvcInheritedASMConfig) HasForecastPeriod() bool`

HasForecastPeriod returns a boolean if a field has been set.

### GetHistory

`func (o *IpamsvcInheritedASMConfig) GetHistory() InheritanceInheritedUInt32`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *IpamsvcInheritedASMConfig) GetHistoryOk() (*InheritanceInheritedUInt32, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *IpamsvcInheritedASMConfig) SetHistory(v InheritanceInheritedUInt32)`

SetHistory sets History field to given value.

### HasHistory

`func (o *IpamsvcInheritedASMConfig) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetMinTotal

`func (o *IpamsvcInheritedASMConfig) GetMinTotal() InheritanceInheritedUInt32`

GetMinTotal returns the MinTotal field if non-nil, zero value otherwise.

### GetMinTotalOk

`func (o *IpamsvcInheritedASMConfig) GetMinTotalOk() (*InheritanceInheritedUInt32, bool)`

GetMinTotalOk returns a tuple with the MinTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTotal

`func (o *IpamsvcInheritedASMConfig) SetMinTotal(v InheritanceInheritedUInt32)`

SetMinTotal sets MinTotal field to given value.

### HasMinTotal

`func (o *IpamsvcInheritedASMConfig) HasMinTotal() bool`

HasMinTotal returns a boolean if a field has been set.

### GetMinUnused

`func (o *IpamsvcInheritedASMConfig) GetMinUnused() InheritanceInheritedUInt32`

GetMinUnused returns the MinUnused field if non-nil, zero value otherwise.

### GetMinUnusedOk

`func (o *IpamsvcInheritedASMConfig) GetMinUnusedOk() (*InheritanceInheritedUInt32, bool)`

GetMinUnusedOk returns a tuple with the MinUnused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUnused

`func (o *IpamsvcInheritedASMConfig) SetMinUnused(v InheritanceInheritedUInt32)`

SetMinUnused sets MinUnused field to given value.

### HasMinUnused

`func (o *IpamsvcInheritedASMConfig) HasMinUnused() bool`

HasMinUnused returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


