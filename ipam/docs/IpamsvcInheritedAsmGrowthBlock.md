# IpamsvcInheritedAsmGrowthBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_. | [optional] [readonly] 
**DisplayName** | Pointer to **string** | The human-readable display name for the object referred to by _source_. | [optional] [readonly] 
**Source** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Value** | Pointer to [**IpamsvcAsmGrowthBlock**](IpamsvcAsmGrowthBlock.md) |  | [optional] 

## Methods

### NewIpamsvcInheritedAsmGrowthBlock

`func NewIpamsvcInheritedAsmGrowthBlock() *IpamsvcInheritedAsmGrowthBlock`

NewIpamsvcInheritedAsmGrowthBlock instantiates a new IpamsvcInheritedAsmGrowthBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcInheritedAsmGrowthBlockWithDefaults

`func NewIpamsvcInheritedAsmGrowthBlockWithDefaults() *IpamsvcInheritedAsmGrowthBlock`

NewIpamsvcInheritedAsmGrowthBlockWithDefaults instantiates a new IpamsvcInheritedAsmGrowthBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *IpamsvcInheritedAsmGrowthBlock) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IpamsvcInheritedAsmGrowthBlock) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IpamsvcInheritedAsmGrowthBlock) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *IpamsvcInheritedAsmGrowthBlock) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDisplayName

`func (o *IpamsvcInheritedAsmGrowthBlock) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IpamsvcInheritedAsmGrowthBlock) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IpamsvcInheritedAsmGrowthBlock) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IpamsvcInheritedAsmGrowthBlock) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSource

`func (o *IpamsvcInheritedAsmGrowthBlock) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IpamsvcInheritedAsmGrowthBlock) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IpamsvcInheritedAsmGrowthBlock) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *IpamsvcInheritedAsmGrowthBlock) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetValue

`func (o *IpamsvcInheritedAsmGrowthBlock) GetValue() IpamsvcAsmGrowthBlock`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IpamsvcInheritedAsmGrowthBlock) GetValueOk() (*IpamsvcAsmGrowthBlock, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IpamsvcInheritedAsmGrowthBlock) SetValue(v IpamsvcAsmGrowthBlock)`

SetValue sets Value field to given value.

### HasValue

`func (o *IpamsvcInheritedAsmGrowthBlock) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


