# InheritedDHCPOptionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Option** | Pointer to [**OptionItem**](OptionItem.md) |  | [optional] 
**OverridingGroup** | Pointer to **string** | The resource identifier. | [optional] 

## Methods

### NewInheritedDHCPOptionItem

`func NewInheritedDHCPOptionItem() *InheritedDHCPOptionItem`

NewInheritedDHCPOptionItem instantiates a new InheritedDHCPOptionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInheritedDHCPOptionItemWithDefaults

`func NewInheritedDHCPOptionItemWithDefaults() *InheritedDHCPOptionItem`

NewInheritedDHCPOptionItemWithDefaults instantiates a new InheritedDHCPOptionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOption

`func (o *InheritedDHCPOptionItem) GetOption() OptionItem`

GetOption returns the Option field if non-nil, zero value otherwise.

### GetOptionOk

`func (o *InheritedDHCPOptionItem) GetOptionOk() (*OptionItem, bool)`

GetOptionOk returns a tuple with the Option field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption

`func (o *InheritedDHCPOptionItem) SetOption(v OptionItem)`

SetOption sets Option field to given value.

### HasOption

`func (o *InheritedDHCPOptionItem) HasOption() bool`

HasOption returns a boolean if a field has been set.

### GetOverridingGroup

`func (o *InheritedDHCPOptionItem) GetOverridingGroup() string`

GetOverridingGroup returns the OverridingGroup field if non-nil, zero value otherwise.

### GetOverridingGroupOk

`func (o *InheritedDHCPOptionItem) GetOverridingGroupOk() (*string, bool)`

GetOverridingGroupOk returns a tuple with the OverridingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridingGroup

`func (o *InheritedDHCPOptionItem) SetOverridingGroup(v string)`

SetOverridingGroup sets OverridingGroup field to given value.

### HasOverridingGroup

`func (o *InheritedDHCPOptionItem) HasOverridingGroup() bool`

HasOverridingGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


