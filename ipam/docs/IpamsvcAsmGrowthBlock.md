# IpamsvcAsmGrowthBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrowthFactor** | Pointer to **int64** | Either the number or percentage of addresses to grow by. | [optional] 
**GrowthType** | Pointer to **string** | The type of factor to use: _percent_ or _count_. | [optional] 

## Methods

### NewIpamsvcAsmGrowthBlock

`func NewIpamsvcAsmGrowthBlock() *IpamsvcAsmGrowthBlock`

NewIpamsvcAsmGrowthBlock instantiates a new IpamsvcAsmGrowthBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcAsmGrowthBlockWithDefaults

`func NewIpamsvcAsmGrowthBlockWithDefaults() *IpamsvcAsmGrowthBlock`

NewIpamsvcAsmGrowthBlockWithDefaults instantiates a new IpamsvcAsmGrowthBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrowthFactor

`func (o *IpamsvcAsmGrowthBlock) GetGrowthFactor() int64`

GetGrowthFactor returns the GrowthFactor field if non-nil, zero value otherwise.

### GetGrowthFactorOk

`func (o *IpamsvcAsmGrowthBlock) GetGrowthFactorOk() (*int64, bool)`

GetGrowthFactorOk returns a tuple with the GrowthFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrowthFactor

`func (o *IpamsvcAsmGrowthBlock) SetGrowthFactor(v int64)`

SetGrowthFactor sets GrowthFactor field to given value.

### HasGrowthFactor

`func (o *IpamsvcAsmGrowthBlock) HasGrowthFactor() bool`

HasGrowthFactor returns a boolean if a field has been set.

### GetGrowthType

`func (o *IpamsvcAsmGrowthBlock) GetGrowthType() string`

GetGrowthType returns the GrowthType field if non-nil, zero value otherwise.

### GetGrowthTypeOk

`func (o *IpamsvcAsmGrowthBlock) GetGrowthTypeOk() (*string, bool)`

GetGrowthTypeOk returns a tuple with the GrowthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrowthType

`func (o *IpamsvcAsmGrowthBlock) SetGrowthType(v string)`

SetGrowthType sets GrowthType field to given value.

### HasGrowthType

`func (o *IpamsvcAsmGrowthBlock) HasGrowthType() bool`

HasGrowthType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


