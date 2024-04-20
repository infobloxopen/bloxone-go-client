# InheritedDDNSHostnameBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_. | [optional] 
**DisplayName** | Pointer to **string** | The human-readable display name for the object referred to by _source_. | [optional] [readonly] 
**Source** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Value** | Pointer to [**DDNSHostnameBlock**](DDNSHostnameBlock.md) |  | [optional] 

## Methods

### NewInheritedDDNSHostnameBlock

`func NewInheritedDDNSHostnameBlock() *InheritedDDNSHostnameBlock`

NewInheritedDDNSHostnameBlock instantiates a new InheritedDDNSHostnameBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInheritedDDNSHostnameBlockWithDefaults

`func NewInheritedDDNSHostnameBlockWithDefaults() *InheritedDDNSHostnameBlock`

NewInheritedDDNSHostnameBlockWithDefaults instantiates a new InheritedDDNSHostnameBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *InheritedDDNSHostnameBlock) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *InheritedDDNSHostnameBlock) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *InheritedDDNSHostnameBlock) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *InheritedDDNSHostnameBlock) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDisplayName

`func (o *InheritedDDNSHostnameBlock) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InheritedDDNSHostnameBlock) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InheritedDDNSHostnameBlock) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InheritedDDNSHostnameBlock) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSource

`func (o *InheritedDDNSHostnameBlock) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InheritedDDNSHostnameBlock) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InheritedDDNSHostnameBlock) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *InheritedDDNSHostnameBlock) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetValue

`func (o *InheritedDDNSHostnameBlock) GetValue() DDNSHostnameBlock`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InheritedDDNSHostnameBlock) GetValueOk() (*DDNSHostnameBlock, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InheritedDDNSHostnameBlock) SetValue(v DDNSHostnameBlock)`

SetValue sets Value field to given value.

### HasValue

`func (o *InheritedDDNSHostnameBlock) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


