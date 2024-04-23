# InheritedDDNSBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_. | [optional] 
**DisplayName** | Pointer to **string** | The human-readable display name for the object referred to by _source_. | [optional] [readonly] 
**Source** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Value** | Pointer to [**DDNSBlock**](DDNSBlock.md) |  | [optional] 

## Methods

### NewInheritedDDNSBlock

`func NewInheritedDDNSBlock() *InheritedDDNSBlock`

NewInheritedDDNSBlock instantiates a new InheritedDDNSBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInheritedDDNSBlockWithDefaults

`func NewInheritedDDNSBlockWithDefaults() *InheritedDDNSBlock`

NewInheritedDDNSBlockWithDefaults instantiates a new InheritedDDNSBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *InheritedDDNSBlock) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *InheritedDDNSBlock) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *InheritedDDNSBlock) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *InheritedDDNSBlock) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDisplayName

`func (o *InheritedDDNSBlock) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InheritedDDNSBlock) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InheritedDDNSBlock) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InheritedDDNSBlock) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSource

`func (o *InheritedDDNSBlock) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InheritedDDNSBlock) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InheritedDDNSBlock) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *InheritedDDNSBlock) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetValue

`func (o *InheritedDDNSBlock) GetValue() DDNSBlock`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InheritedDDNSBlock) GetValueOk() (*DDNSBlock, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InheritedDDNSBlock) SetValue(v DDNSBlock)`

SetValue sets Value field to given value.

### HasValue

`func (o *InheritedDDNSBlock) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


