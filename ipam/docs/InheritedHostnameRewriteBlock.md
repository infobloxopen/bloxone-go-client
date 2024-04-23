# InheritedHostnameRewriteBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_. | [optional] 
**DisplayName** | Pointer to **string** | The human-readable display name for the object referred to by _source_. | [optional] [readonly] 
**Source** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Value** | Pointer to [**HostnameRewriteBlock**](HostnameRewriteBlock.md) |  | [optional] 

## Methods

### NewInheritedHostnameRewriteBlock

`func NewInheritedHostnameRewriteBlock() *InheritedHostnameRewriteBlock`

NewInheritedHostnameRewriteBlock instantiates a new InheritedHostnameRewriteBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInheritedHostnameRewriteBlockWithDefaults

`func NewInheritedHostnameRewriteBlockWithDefaults() *InheritedHostnameRewriteBlock`

NewInheritedHostnameRewriteBlockWithDefaults instantiates a new InheritedHostnameRewriteBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *InheritedHostnameRewriteBlock) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *InheritedHostnameRewriteBlock) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *InheritedHostnameRewriteBlock) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *InheritedHostnameRewriteBlock) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDisplayName

`func (o *InheritedHostnameRewriteBlock) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InheritedHostnameRewriteBlock) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InheritedHostnameRewriteBlock) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InheritedHostnameRewriteBlock) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSource

`func (o *InheritedHostnameRewriteBlock) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InheritedHostnameRewriteBlock) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InheritedHostnameRewriteBlock) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *InheritedHostnameRewriteBlock) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetValue

`func (o *InheritedHostnameRewriteBlock) GetValue() HostnameRewriteBlock`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InheritedHostnameRewriteBlock) GetValueOk() (*HostnameRewriteBlock, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InheritedHostnameRewriteBlock) SetValue(v HostnameRewriteBlock)`

SetValue sets Value field to given value.

### HasValue

`func (o *InheritedHostnameRewriteBlock) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


