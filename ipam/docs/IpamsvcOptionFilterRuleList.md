# IpamsvcOptionFilterRuleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Match** | Pointer to **string** | Indicates if this list should match if any or all rules match (_any_ or _all_). | [optional] 
**Rules** | Pointer to [**[]IpamsvcOptionFilterRule**](IpamsvcOptionFilterRule.md) | The list of child rules. | [optional] 

## Methods

### NewIpamsvcOptionFilterRuleList

`func NewIpamsvcOptionFilterRuleList() *IpamsvcOptionFilterRuleList`

NewIpamsvcOptionFilterRuleList instantiates a new IpamsvcOptionFilterRuleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcOptionFilterRuleListWithDefaults

`func NewIpamsvcOptionFilterRuleListWithDefaults() *IpamsvcOptionFilterRuleList`

NewIpamsvcOptionFilterRuleListWithDefaults instantiates a new IpamsvcOptionFilterRuleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatch

`func (o *IpamsvcOptionFilterRuleList) GetMatch() string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *IpamsvcOptionFilterRuleList) GetMatchOk() (*string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *IpamsvcOptionFilterRuleList) SetMatch(v string)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *IpamsvcOptionFilterRuleList) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetRules

`func (o *IpamsvcOptionFilterRuleList) GetRules() []IpamsvcOptionFilterRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *IpamsvcOptionFilterRuleList) GetRulesOk() (*[]IpamsvcOptionFilterRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *IpamsvcOptionFilterRuleList) SetRules(v []IpamsvcOptionFilterRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *IpamsvcOptionFilterRuleList) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


