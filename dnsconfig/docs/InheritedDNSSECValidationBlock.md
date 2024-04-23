# InheritedDNSSECValidationBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Defaults to _inherit_. | [optional] 
**DisplayName** | Pointer to **string** | Human-readable display name for the object referred to by _source_. | [optional] [readonly] 
**Source** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Value** | Pointer to [**DNSSECValidationBlock**](DNSSECValidationBlock.md) |  | [optional] 

## Methods

### NewInheritedDNSSECValidationBlock

`func NewInheritedDNSSECValidationBlock() *InheritedDNSSECValidationBlock`

NewInheritedDNSSECValidationBlock instantiates a new InheritedDNSSECValidationBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInheritedDNSSECValidationBlockWithDefaults

`func NewInheritedDNSSECValidationBlockWithDefaults() *InheritedDNSSECValidationBlock`

NewInheritedDNSSECValidationBlockWithDefaults instantiates a new InheritedDNSSECValidationBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *InheritedDNSSECValidationBlock) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *InheritedDNSSECValidationBlock) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *InheritedDNSSECValidationBlock) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *InheritedDNSSECValidationBlock) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDisplayName

`func (o *InheritedDNSSECValidationBlock) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InheritedDNSSECValidationBlock) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InheritedDNSSECValidationBlock) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InheritedDNSSECValidationBlock) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSource

`func (o *InheritedDNSSECValidationBlock) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InheritedDNSSECValidationBlock) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InheritedDNSSECValidationBlock) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *InheritedDNSSECValidationBlock) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetValue

`func (o *InheritedDNSSECValidationBlock) GetValue() DNSSECValidationBlock`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InheritedDNSSECValidationBlock) GetValueOk() (*DNSSECValidationBlock, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InheritedDNSSECValidationBlock) SetValue(v DNSSECValidationBlock)`

SetValue sets Value field to given value.

### HasValue

`func (o *InheritedDNSSECValidationBlock) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


