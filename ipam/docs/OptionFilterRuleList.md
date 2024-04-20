# OptionFilterRuleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Match** | Pointer to **string** | Indicates if this list should match if any or all rules match (_any_ or _all_). | [optional] 
**Rules** | Pointer to [**[]OptionFilterRule**](OptionFilterRule.md) | The list of child rules. | [optional] 

## Methods

### NewOptionFilterRuleList

`func NewOptionFilterRuleList() *OptionFilterRuleList`

NewOptionFilterRuleList instantiates a new OptionFilterRuleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionFilterRuleListWithDefaults

`func NewOptionFilterRuleListWithDefaults() *OptionFilterRuleList`

NewOptionFilterRuleListWithDefaults instantiates a new OptionFilterRuleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatch

`func (o *OptionFilterRuleList) GetMatch() string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *OptionFilterRuleList) GetMatchOk() (*string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *OptionFilterRuleList) SetMatch(v string)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *OptionFilterRuleList) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetRules

`func (o *OptionFilterRuleList) GetRules() []OptionFilterRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *OptionFilterRuleList) GetRulesOk() (*[]OptionFilterRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *OptionFilterRuleList) SetRules(v []OptionFilterRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *OptionFilterRuleList) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


