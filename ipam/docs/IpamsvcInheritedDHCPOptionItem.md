# IpamsvcInheritedDHCPOptionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Option** | Pointer to [**IpamsvcOptionItem**](IpamsvcOptionItem.md) |  | [optional] 
**OverridingGroup** | Pointer to **string** | The resource identifier. | [optional] 

## Methods

### NewIpamsvcInheritedDHCPOptionItem

`func NewIpamsvcInheritedDHCPOptionItem() *IpamsvcInheritedDHCPOptionItem`

NewIpamsvcInheritedDHCPOptionItem instantiates a new IpamsvcInheritedDHCPOptionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcInheritedDHCPOptionItemWithDefaults

`func NewIpamsvcInheritedDHCPOptionItemWithDefaults() *IpamsvcInheritedDHCPOptionItem`

NewIpamsvcInheritedDHCPOptionItemWithDefaults instantiates a new IpamsvcInheritedDHCPOptionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOption

`func (o *IpamsvcInheritedDHCPOptionItem) GetOption() IpamsvcOptionItem`

GetOption returns the Option field if non-nil, zero value otherwise.

### GetOptionOk

`func (o *IpamsvcInheritedDHCPOptionItem) GetOptionOk() (*IpamsvcOptionItem, bool)`

GetOptionOk returns a tuple with the Option field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption

`func (o *IpamsvcInheritedDHCPOptionItem) SetOption(v IpamsvcOptionItem)`

SetOption sets Option field to given value.

### HasOption

`func (o *IpamsvcInheritedDHCPOptionItem) HasOption() bool`

HasOption returns a boolean if a field has been set.

### GetOverridingGroup

`func (o *IpamsvcInheritedDHCPOptionItem) GetOverridingGroup() string`

GetOverridingGroup returns the OverridingGroup field if non-nil, zero value otherwise.

### GetOverridingGroupOk

`func (o *IpamsvcInheritedDHCPOptionItem) GetOverridingGroupOk() (*string, bool)`

GetOverridingGroupOk returns a tuple with the OverridingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridingGroup

`func (o *IpamsvcInheritedDHCPOptionItem) SetOverridingGroup(v string)`

SetOverridingGroup sets OverridingGroup field to given value.

### HasOverridingGroup

`func (o *IpamsvcInheritedDHCPOptionItem) HasOverridingGroup() bool`

HasOverridingGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


