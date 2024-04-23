# OptionFilterRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compare** | **string** | Indicates how to compare the _option_value_ to the client option.  Success by comparison:  * _equals_: value and client option are the same,  * _not_equals_: value and client option are not the same,  * _exists_: client option exists,  * _not_exists_: client option does not exist,  * _text_substring_: value is the specified substring of the option,  * _not_text_substring_: value is not the specified substring of the option.  * _hex_substring_: value is the specified hexadecimal substring of the option,  * _not_hex_substring_: value is not the specified hexadecimal substring of the option. | 
**OptionCode** | **string** | The resource identifier. | 
**OptionValue** | Pointer to **string** | The value to match against. | [optional] 
**SubstringOffset** | Pointer to **int64** | The offset where the substring match starts. This is used only if comparing the _option_value_ using any of the substring modes. | [optional] 

## Methods

### NewOptionFilterRule

`func NewOptionFilterRule(compare string, optionCode string, ) *OptionFilterRule`

NewOptionFilterRule instantiates a new OptionFilterRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionFilterRuleWithDefaults

`func NewOptionFilterRuleWithDefaults() *OptionFilterRule`

NewOptionFilterRuleWithDefaults instantiates a new OptionFilterRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompare

`func (o *OptionFilterRule) GetCompare() string`

GetCompare returns the Compare field if non-nil, zero value otherwise.

### GetCompareOk

`func (o *OptionFilterRule) GetCompareOk() (*string, bool)`

GetCompareOk returns a tuple with the Compare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompare

`func (o *OptionFilterRule) SetCompare(v string)`

SetCompare sets Compare field to given value.


### GetOptionCode

`func (o *OptionFilterRule) GetOptionCode() string`

GetOptionCode returns the OptionCode field if non-nil, zero value otherwise.

### GetOptionCodeOk

`func (o *OptionFilterRule) GetOptionCodeOk() (*string, bool)`

GetOptionCodeOk returns a tuple with the OptionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionCode

`func (o *OptionFilterRule) SetOptionCode(v string)`

SetOptionCode sets OptionCode field to given value.


### GetOptionValue

`func (o *OptionFilterRule) GetOptionValue() string`

GetOptionValue returns the OptionValue field if non-nil, zero value otherwise.

### GetOptionValueOk

`func (o *OptionFilterRule) GetOptionValueOk() (*string, bool)`

GetOptionValueOk returns a tuple with the OptionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionValue

`func (o *OptionFilterRule) SetOptionValue(v string)`

SetOptionValue sets OptionValue field to given value.

### HasOptionValue

`func (o *OptionFilterRule) HasOptionValue() bool`

HasOptionValue returns a boolean if a field has been set.

### GetSubstringOffset

`func (o *OptionFilterRule) GetSubstringOffset() int64`

GetSubstringOffset returns the SubstringOffset field if non-nil, zero value otherwise.

### GetSubstringOffsetOk

`func (o *OptionFilterRule) GetSubstringOffsetOk() (*int64, bool)`

GetSubstringOffsetOk returns a tuple with the SubstringOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstringOffset

`func (o *OptionFilterRule) SetSubstringOffset(v int64)`

SetSubstringOffset sets SubstringOffset field to given value.

### HasSubstringOffset

`func (o *OptionFilterRule) HasSubstringOffset() bool`

HasSubstringOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


