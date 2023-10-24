# IpamsvcOptionFilterRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compare** | **string** | Indicates how to compare the _option_value_ to the client option.  Success by comparison:  * _equals_: value and client option are the same,  * _not_equals_: value and client option are not the same,  * _exists_: client option exists,  * _not_exists_: client option does not exist,  * _text_substring_: value is the specified substring of the option,  * _not_text_substring_: value is not the specified substring of the option.  * _hex_substring_: value is the specified hexadecimal substring of the option,  * _not_hex_substring_: value is not the specified hexadecimal substring of the option. | 
**OptionCode** | **string** | The resource identifier. | 
**OptionValue** | Pointer to **string** | The value to match against. | [optional] 
**SubstringOffset** | Pointer to **int64** | The offset where the substring match starts. This is used only if comparing the _option_value_ using any of the substring modes. | [optional] 

## Methods

### NewIpamsvcOptionFilterRule

`func NewIpamsvcOptionFilterRule(compare string, optionCode string, ) *IpamsvcOptionFilterRule`

NewIpamsvcOptionFilterRule instantiates a new IpamsvcOptionFilterRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcOptionFilterRuleWithDefaults

`func NewIpamsvcOptionFilterRuleWithDefaults() *IpamsvcOptionFilterRule`

NewIpamsvcOptionFilterRuleWithDefaults instantiates a new IpamsvcOptionFilterRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompare

`func (o *IpamsvcOptionFilterRule) GetCompare() string`

GetCompare returns the Compare field if non-nil, zero value otherwise.

### GetCompareOk

`func (o *IpamsvcOptionFilterRule) GetCompareOk() (*string, bool)`

GetCompareOk returns a tuple with the Compare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompare

`func (o *IpamsvcOptionFilterRule) SetCompare(v string)`

SetCompare sets Compare field to given value.


### GetOptionCode

`func (o *IpamsvcOptionFilterRule) GetOptionCode() string`

GetOptionCode returns the OptionCode field if non-nil, zero value otherwise.

### GetOptionCodeOk

`func (o *IpamsvcOptionFilterRule) GetOptionCodeOk() (*string, bool)`

GetOptionCodeOk returns a tuple with the OptionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionCode

`func (o *IpamsvcOptionFilterRule) SetOptionCode(v string)`

SetOptionCode sets OptionCode field to given value.


### GetOptionValue

`func (o *IpamsvcOptionFilterRule) GetOptionValue() string`

GetOptionValue returns the OptionValue field if non-nil, zero value otherwise.

### GetOptionValueOk

`func (o *IpamsvcOptionFilterRule) GetOptionValueOk() (*string, bool)`

GetOptionValueOk returns a tuple with the OptionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionValue

`func (o *IpamsvcOptionFilterRule) SetOptionValue(v string)`

SetOptionValue sets OptionValue field to given value.

### HasOptionValue

`func (o *IpamsvcOptionFilterRule) HasOptionValue() bool`

HasOptionValue returns a boolean if a field has been set.

### GetSubstringOffset

`func (o *IpamsvcOptionFilterRule) GetSubstringOffset() int64`

GetSubstringOffset returns the SubstringOffset field if non-nil, zero value otherwise.

### GetSubstringOffsetOk

`func (o *IpamsvcOptionFilterRule) GetSubstringOffsetOk() (*int64, bool)`

GetSubstringOffsetOk returns a tuple with the SubstringOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstringOffset

`func (o *IpamsvcOptionFilterRule) SetSubstringOffset(v int64)`

SetSubstringOffset sets SubstringOffset field to given value.

### HasSubstringOffset

`func (o *IpamsvcOptionFilterRule) HasSubstringOffset() bool`

HasSubstringOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


