# InheritedKerberosKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Optional. Inheritance setting for a field. Defaults to _inherit_. | [optional] 
**DisplayName** | Pointer to **string** | Human-readable display name for the object referred to by _source_. | [optional] [readonly] 
**Source** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Value** | Pointer to [**[]KerberosKey**](KerberosKey.md) | Inherited value. | [optional] [readonly] 

## Methods

### NewInheritedKerberosKeys

`func NewInheritedKerberosKeys() *InheritedKerberosKeys`

NewInheritedKerberosKeys instantiates a new InheritedKerberosKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInheritedKerberosKeysWithDefaults

`func NewInheritedKerberosKeysWithDefaults() *InheritedKerberosKeys`

NewInheritedKerberosKeysWithDefaults instantiates a new InheritedKerberosKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *InheritedKerberosKeys) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *InheritedKerberosKeys) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *InheritedKerberosKeys) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *InheritedKerberosKeys) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDisplayName

`func (o *InheritedKerberosKeys) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InheritedKerberosKeys) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InheritedKerberosKeys) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InheritedKerberosKeys) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSource

`func (o *InheritedKerberosKeys) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InheritedKerberosKeys) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InheritedKerberosKeys) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *InheritedKerberosKeys) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetValue

`func (o *InheritedKerberosKeys) GetValue() []KerberosKey`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InheritedKerberosKeys) GetValueOk() (*[]KerberosKey, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InheritedKerberosKeys) SetValue(v []KerberosKey)`

SetValue sets Value field to given value.

### HasValue

`func (o *InheritedKerberosKeys) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


