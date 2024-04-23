# SecurityPolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The action for the policy rule that can be either \&quot;action_allow\&quot; or \&quot;action_log\&quot; or \&quot;action_redirect\&quot; or \&quot;action_block\&quot; or \&quot;action_allow_with_local_resolution\&quot;. \&quot;action_allow_with_local_resolution\&quot; only supported for application filter rule with enabled onprem_resolve flag on the Security policy. | [optional] 
**Data** | Pointer to **string** | The data source for the policy rule, that can be either a name of the predefined feed for \&quot;named_feed\&quot;, custom list name for \&quot;custom_list\&quot; type, category filter name for \&quot;category_filter\&quot; type and application filter name for \&quot;application_filter\&quot; type. | [optional] 
**ListId** | Pointer to **int32** | The Custom List object identifier with which the policy rule is associated. 0 value means no custom list is associated with this policy rule. | [optional] [readonly] 
**PolicyId** | Pointer to **int32** | The identifier of the Security Policy object with which the policy rule is associated. | [optional] [readonly] 
**PolicyName** | Pointer to **string** | The name of the security policy with which the policy rule is associated. | [optional] 
**RedirectName** | Pointer to **string** | The name of the redirect address for redirect actions that can be either IPv4 address or a domain name. | [optional] 
**Type** | Pointer to **string** | The policy rule type that can be either \&quot;named_feed\&quot; or \&quot;custom_list\&quot; or \&quot;category_filter\&quot; or \&quot;application_filter\&quot;. | [optional] 

## Methods

### NewSecurityPolicyRule

`func NewSecurityPolicyRule() *SecurityPolicyRule`

NewSecurityPolicyRule instantiates a new SecurityPolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityPolicyRuleWithDefaults

`func NewSecurityPolicyRuleWithDefaults() *SecurityPolicyRule`

NewSecurityPolicyRuleWithDefaults instantiates a new SecurityPolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SecurityPolicyRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SecurityPolicyRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SecurityPolicyRule) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SecurityPolicyRule) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetData

`func (o *SecurityPolicyRule) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SecurityPolicyRule) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SecurityPolicyRule) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *SecurityPolicyRule) HasData() bool`

HasData returns a boolean if a field has been set.

### GetListId

`func (o *SecurityPolicyRule) GetListId() int32`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *SecurityPolicyRule) GetListIdOk() (*int32, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *SecurityPolicyRule) SetListId(v int32)`

SetListId sets ListId field to given value.

### HasListId

`func (o *SecurityPolicyRule) HasListId() bool`

HasListId returns a boolean if a field has been set.

### GetPolicyId

`func (o *SecurityPolicyRule) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *SecurityPolicyRule) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *SecurityPolicyRule) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *SecurityPolicyRule) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *SecurityPolicyRule) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *SecurityPolicyRule) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *SecurityPolicyRule) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *SecurityPolicyRule) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetRedirectName

`func (o *SecurityPolicyRule) GetRedirectName() string`

GetRedirectName returns the RedirectName field if non-nil, zero value otherwise.

### GetRedirectNameOk

`func (o *SecurityPolicyRule) GetRedirectNameOk() (*string, bool)`

GetRedirectNameOk returns a tuple with the RedirectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectName

`func (o *SecurityPolicyRule) SetRedirectName(v string)`

SetRedirectName sets RedirectName field to given value.

### HasRedirectName

`func (o *SecurityPolicyRule) HasRedirectName() bool`

HasRedirectName returns a boolean if a field has been set.

### GetType

`func (o *SecurityPolicyRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityPolicyRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityPolicyRule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SecurityPolicyRule) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


