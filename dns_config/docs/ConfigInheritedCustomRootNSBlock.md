# ConfigInheritedCustomRootNSBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Defaults to _inherit_. | [optional] [readonly] 
**DisplayName** | Pointer to **string** | Human-readable display name for the object referred to by _source_. | [optional] [readonly] 
**Source** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Value** | Pointer to [**ConfigCustomRootNSBlock**](ConfigCustomRootNSBlock.md) |  | [optional] 

## Methods

### NewConfigInheritedCustomRootNSBlock

`func NewConfigInheritedCustomRootNSBlock() *ConfigInheritedCustomRootNSBlock`

NewConfigInheritedCustomRootNSBlock instantiates a new ConfigInheritedCustomRootNSBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigInheritedCustomRootNSBlockWithDefaults

`func NewConfigInheritedCustomRootNSBlockWithDefaults() *ConfigInheritedCustomRootNSBlock`

NewConfigInheritedCustomRootNSBlockWithDefaults instantiates a new ConfigInheritedCustomRootNSBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ConfigInheritedCustomRootNSBlock) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ConfigInheritedCustomRootNSBlock) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ConfigInheritedCustomRootNSBlock) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ConfigInheritedCustomRootNSBlock) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDisplayName

`func (o *ConfigInheritedCustomRootNSBlock) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ConfigInheritedCustomRootNSBlock) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ConfigInheritedCustomRootNSBlock) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ConfigInheritedCustomRootNSBlock) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSource

`func (o *ConfigInheritedCustomRootNSBlock) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ConfigInheritedCustomRootNSBlock) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ConfigInheritedCustomRootNSBlock) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ConfigInheritedCustomRootNSBlock) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetValue

`func (o *ConfigInheritedCustomRootNSBlock) GetValue() ConfigCustomRootNSBlock`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConfigInheritedCustomRootNSBlock) GetValueOk() (*ConfigCustomRootNSBlock, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConfigInheritedCustomRootNSBlock) SetValue(v ConfigCustomRootNSBlock)`

SetValue sets Value field to given value.

### HasValue

`func (o *ConfigInheritedCustomRootNSBlock) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


