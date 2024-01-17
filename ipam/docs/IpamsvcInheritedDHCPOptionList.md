# IpamsvcInheritedDHCPOptionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _block_: Don&#39;t use the inherited value.  Defaults to _inherit_. | [optional] 
**Value** | Pointer to [**[]IpamsvcInheritedDHCPOption**](IpamsvcInheritedDHCPOption.md) | The inherited DHCP option values. | [optional] 

## Methods

### NewIpamsvcInheritedDHCPOptionList

`func NewIpamsvcInheritedDHCPOptionList() *IpamsvcInheritedDHCPOptionList`

NewIpamsvcInheritedDHCPOptionList instantiates a new IpamsvcInheritedDHCPOptionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcInheritedDHCPOptionListWithDefaults

`func NewIpamsvcInheritedDHCPOptionListWithDefaults() *IpamsvcInheritedDHCPOptionList`

NewIpamsvcInheritedDHCPOptionListWithDefaults instantiates a new IpamsvcInheritedDHCPOptionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *IpamsvcInheritedDHCPOptionList) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IpamsvcInheritedDHCPOptionList) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IpamsvcInheritedDHCPOptionList) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *IpamsvcInheritedDHCPOptionList) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetValue

`func (o *IpamsvcInheritedDHCPOptionList) GetValue() []IpamsvcInheritedDHCPOption`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IpamsvcInheritedDHCPOptionList) GetValueOk() (*[]IpamsvcInheritedDHCPOption, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IpamsvcInheritedDHCPOptionList) SetValue(v []IpamsvcInheritedDHCPOption)`

SetValue sets Value field to given value.

### HasValue

`func (o *IpamsvcInheritedDHCPOptionList) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


