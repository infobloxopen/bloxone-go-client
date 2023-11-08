# IpamsvcInheritedDHCPOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _block_: Don&#39;t use the inherited value.  Defaults to _inherit_. | [optional] 
**DisplayName** | Pointer to **string** | The human-readable display name for the object referred to by _source_. | [optional] [readonly] 
**Source** | Pointer to **string** | The resource identifier. | [optional] 
**Value** | Pointer to [**IpamsvcInheritedDHCPOptionItem**](IpamsvcInheritedDHCPOptionItem.md) |  | [optional] 

## Methods

### NewIpamsvcInheritedDHCPOption

`func NewIpamsvcInheritedDHCPOption() *IpamsvcInheritedDHCPOption`

NewIpamsvcInheritedDHCPOption instantiates a new IpamsvcInheritedDHCPOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcInheritedDHCPOptionWithDefaults

`func NewIpamsvcInheritedDHCPOptionWithDefaults() *IpamsvcInheritedDHCPOption`

NewIpamsvcInheritedDHCPOptionWithDefaults instantiates a new IpamsvcInheritedDHCPOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *IpamsvcInheritedDHCPOption) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IpamsvcInheritedDHCPOption) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IpamsvcInheritedDHCPOption) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *IpamsvcInheritedDHCPOption) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDisplayName

`func (o *IpamsvcInheritedDHCPOption) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IpamsvcInheritedDHCPOption) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IpamsvcInheritedDHCPOption) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IpamsvcInheritedDHCPOption) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSource

`func (o *IpamsvcInheritedDHCPOption) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IpamsvcInheritedDHCPOption) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IpamsvcInheritedDHCPOption) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *IpamsvcInheritedDHCPOption) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetValue

`func (o *IpamsvcInheritedDHCPOption) GetValue() IpamsvcInheritedDHCPOptionItem`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IpamsvcInheritedDHCPOption) GetValueOk() (*IpamsvcInheritedDHCPOptionItem, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IpamsvcInheritedDHCPOption) SetValue(v IpamsvcInheritedDHCPOptionItem)`

SetValue sets Value field to given value.

### HasValue

`func (o *IpamsvcInheritedDHCPOption) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


