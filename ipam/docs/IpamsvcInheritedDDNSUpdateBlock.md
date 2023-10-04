# IpamsvcInheritedDDNSUpdateBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_. | [optional] 
**DisplayName** | Pointer to **string** | The human-readable display name for the object referred to by _source_. | [optional] [readonly] 
**Source** | Pointer to **string** | The resource identifier. | [optional] 
**Value** | Pointer to [**IpamsvcDDNSUpdateBlock**](IpamsvcDDNSUpdateBlock.md) |  | [optional] 

## Methods

### NewIpamsvcInheritedDDNSUpdateBlock

`func NewIpamsvcInheritedDDNSUpdateBlock() *IpamsvcInheritedDDNSUpdateBlock`

NewIpamsvcInheritedDDNSUpdateBlock instantiates a new IpamsvcInheritedDDNSUpdateBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcInheritedDDNSUpdateBlockWithDefaults

`func NewIpamsvcInheritedDDNSUpdateBlockWithDefaults() *IpamsvcInheritedDDNSUpdateBlock`

NewIpamsvcInheritedDDNSUpdateBlockWithDefaults instantiates a new IpamsvcInheritedDDNSUpdateBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *IpamsvcInheritedDDNSUpdateBlock) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IpamsvcInheritedDDNSUpdateBlock) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IpamsvcInheritedDDNSUpdateBlock) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *IpamsvcInheritedDDNSUpdateBlock) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDisplayName

`func (o *IpamsvcInheritedDDNSUpdateBlock) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IpamsvcInheritedDDNSUpdateBlock) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IpamsvcInheritedDDNSUpdateBlock) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IpamsvcInheritedDDNSUpdateBlock) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSource

`func (o *IpamsvcInheritedDDNSUpdateBlock) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IpamsvcInheritedDDNSUpdateBlock) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IpamsvcInheritedDDNSUpdateBlock) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *IpamsvcInheritedDDNSUpdateBlock) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetValue

`func (o *IpamsvcInheritedDDNSUpdateBlock) GetValue() IpamsvcDDNSUpdateBlock`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IpamsvcInheritedDDNSUpdateBlock) GetValueOk() (*IpamsvcDDNSUpdateBlock, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IpamsvcInheritedDDNSUpdateBlock) SetValue(v IpamsvcDDNSUpdateBlock)`

SetValue sets Value field to given value.

### HasValue

`func (o *IpamsvcInheritedDDNSUpdateBlock) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


