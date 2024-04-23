# InheritedForwardersBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Defaults to _inherit_. | [optional] 
**DisplayName** | Pointer to **string** | Human-readable display name for the object referred to by _source_. | [optional] [readonly] 
**Source** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Value** | Pointer to [**ForwardersBlock**](ForwardersBlock.md) |  | [optional] 

## Methods

### NewInheritedForwardersBlock

`func NewInheritedForwardersBlock() *InheritedForwardersBlock`

NewInheritedForwardersBlock instantiates a new InheritedForwardersBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInheritedForwardersBlockWithDefaults

`func NewInheritedForwardersBlockWithDefaults() *InheritedForwardersBlock`

NewInheritedForwardersBlockWithDefaults instantiates a new InheritedForwardersBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *InheritedForwardersBlock) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *InheritedForwardersBlock) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *InheritedForwardersBlock) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *InheritedForwardersBlock) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDisplayName

`func (o *InheritedForwardersBlock) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InheritedForwardersBlock) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InheritedForwardersBlock) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InheritedForwardersBlock) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSource

`func (o *InheritedForwardersBlock) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InheritedForwardersBlock) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InheritedForwardersBlock) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *InheritedForwardersBlock) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetValue

`func (o *InheritedForwardersBlock) GetValue() ForwardersBlock`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InheritedForwardersBlock) GetValueOk() (*ForwardersBlock, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InheritedForwardersBlock) SetValue(v ForwardersBlock)`

SetValue sets Value field to given value.

### HasValue

`func (o *InheritedForwardersBlock) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


