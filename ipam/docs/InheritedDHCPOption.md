# InheritedDHCPOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _block_: Don&#39;t use the inherited value.  Defaults to _inherit_. | [optional] 
**DisplayName** | Pointer to **string** | The human-readable display name for the object referred to by _source_. | [optional] [readonly] 
**Source** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Value** | Pointer to [**InheritedDHCPOptionItem**](InheritedDHCPOptionItem.md) | The inherited value for the DHCP option. | [optional] [readonly] 

## Methods

### NewInheritedDHCPOption

`func NewInheritedDHCPOption() *InheritedDHCPOption`

NewInheritedDHCPOption instantiates a new InheritedDHCPOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInheritedDHCPOptionWithDefaults

`func NewInheritedDHCPOptionWithDefaults() *InheritedDHCPOption`

NewInheritedDHCPOptionWithDefaults instantiates a new InheritedDHCPOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *InheritedDHCPOption) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *InheritedDHCPOption) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *InheritedDHCPOption) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *InheritedDHCPOption) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDisplayName

`func (o *InheritedDHCPOption) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InheritedDHCPOption) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InheritedDHCPOption) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InheritedDHCPOption) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSource

`func (o *InheritedDHCPOption) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InheritedDHCPOption) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InheritedDHCPOption) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *InheritedDHCPOption) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetValue

`func (o *InheritedDHCPOption) GetValue() InheritedDHCPOptionItem`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InheritedDHCPOption) GetValueOk() (*InheritedDHCPOptionItem, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InheritedDHCPOption) SetValue(v InheritedDHCPOptionItem)`

SetValue sets Value field to given value.

### HasValue

`func (o *InheritedDHCPOption) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


