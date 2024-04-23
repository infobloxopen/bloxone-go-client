# SecurityPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessCodes** | Pointer to **[]string** | Access codes assigned to Security Policy | [optional] 
**CreatedTime** | Pointer to **time.Time** | The time when this Security Policy object was created. | [optional] [readonly] 
**DefaultAction** | Pointer to **string** | The policy-level action gets applied when none of the policy rules apply/match. The default value for default_action is \&quot;action_allow\&quot;. | [optional] 
**DefaultRedirectName** | Pointer to **string** | Name of the custom redirect, if the default_action is \&quot;action_redirect\&quot;. | [optional] 
**Description** | Pointer to **string** | The brief description for the security policy. | [optional] 
**DfpServices** | Pointer to **[]string** | The list of DNS Forwarding Proxy Services object identifiers. For Internal Use only. | [optional] 
**Dfps** | Pointer to **[]int32** | The list of DNS Forwarding Proxy object identifiers. | [optional] 
**Ecs** | Pointer to **bool** | Use ECS for handling policy | [optional] 
**Id** | Pointer to **int32** | The Security Policy object identifier. | [optional] [readonly] 
**IsDefault** | Pointer to **bool** | Flag that indicates whether this is a default security policy. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the security policy. | [optional] 
**NetAddressDfps** | Pointer to [**[]NetAddrDfpAssignment**](NetAddrDfpAssignment.md) | List of DFPs associated with this policy via network address (with corresponding network address) | [optional] 
**NetworkLists** | Pointer to **[]int64** | The list of Network Lists identifiers that represents networks that you want to protect from malicious attacks. | [optional] 
**OnpremResolve** | Pointer to **bool** | Use DNS resolve on onprem side | [optional] 
**Precedence** | Pointer to **int32** | Security precedence enable selection of the highest priority policy, in cases where a query matches multiple policies. | [optional] 
**RoamingDeviceGroups** | Pointer to **[]int32** | The list of BloxOne Endpoint groups identifiers. | [optional] 
**Rules** | Pointer to [**[]SecurityPolicyRule**](SecurityPolicyRule.md) | The list of Security Policy Rules objects that represent the set of rules and actions that you define to balance access and constraints so you can mitigate malicious attacks and provide security for your networks. | [optional] 
**SafeSearch** | Pointer to **bool** | Apply automated rules to enforce safe search | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Enables tag support for resource where tags attribute contains user-defined key value pairs | [optional] 
**UpdatedTime** | Pointer to **time.Time** | The time when this Security Policy object was last updated. | [optional] [readonly] 
**UserGroups** | Pointer to **[]string** | List of user groups associated with this policy | [optional] 

## Methods

### NewSecurityPolicy

`func NewSecurityPolicy() *SecurityPolicy`

NewSecurityPolicy instantiates a new SecurityPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityPolicyWithDefaults

`func NewSecurityPolicyWithDefaults() *SecurityPolicy`

NewSecurityPolicyWithDefaults instantiates a new SecurityPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessCodes

`func (o *SecurityPolicy) GetAccessCodes() []string`

GetAccessCodes returns the AccessCodes field if non-nil, zero value otherwise.

### GetAccessCodesOk

`func (o *SecurityPolicy) GetAccessCodesOk() (*[]string, bool)`

GetAccessCodesOk returns a tuple with the AccessCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCodes

`func (o *SecurityPolicy) SetAccessCodes(v []string)`

SetAccessCodes sets AccessCodes field to given value.

### HasAccessCodes

`func (o *SecurityPolicy) HasAccessCodes() bool`

HasAccessCodes returns a boolean if a field has been set.

### GetCreatedTime

`func (o *SecurityPolicy) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *SecurityPolicy) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *SecurityPolicy) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *SecurityPolicy) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDefaultAction

`func (o *SecurityPolicy) GetDefaultAction() string`

GetDefaultAction returns the DefaultAction field if non-nil, zero value otherwise.

### GetDefaultActionOk

`func (o *SecurityPolicy) GetDefaultActionOk() (*string, bool)`

GetDefaultActionOk returns a tuple with the DefaultAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAction

`func (o *SecurityPolicy) SetDefaultAction(v string)`

SetDefaultAction sets DefaultAction field to given value.

### HasDefaultAction

`func (o *SecurityPolicy) HasDefaultAction() bool`

HasDefaultAction returns a boolean if a field has been set.

### GetDefaultRedirectName

`func (o *SecurityPolicy) GetDefaultRedirectName() string`

GetDefaultRedirectName returns the DefaultRedirectName field if non-nil, zero value otherwise.

### GetDefaultRedirectNameOk

`func (o *SecurityPolicy) GetDefaultRedirectNameOk() (*string, bool)`

GetDefaultRedirectNameOk returns a tuple with the DefaultRedirectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRedirectName

`func (o *SecurityPolicy) SetDefaultRedirectName(v string)`

SetDefaultRedirectName sets DefaultRedirectName field to given value.

### HasDefaultRedirectName

`func (o *SecurityPolicy) HasDefaultRedirectName() bool`

HasDefaultRedirectName returns a boolean if a field has been set.

### GetDescription

`func (o *SecurityPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDfpServices

`func (o *SecurityPolicy) GetDfpServices() []string`

GetDfpServices returns the DfpServices field if non-nil, zero value otherwise.

### GetDfpServicesOk

`func (o *SecurityPolicy) GetDfpServicesOk() (*[]string, bool)`

GetDfpServicesOk returns a tuple with the DfpServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfpServices

`func (o *SecurityPolicy) SetDfpServices(v []string)`

SetDfpServices sets DfpServices field to given value.

### HasDfpServices

`func (o *SecurityPolicy) HasDfpServices() bool`

HasDfpServices returns a boolean if a field has been set.

### GetDfps

`func (o *SecurityPolicy) GetDfps() []int32`

GetDfps returns the Dfps field if non-nil, zero value otherwise.

### GetDfpsOk

`func (o *SecurityPolicy) GetDfpsOk() (*[]int32, bool)`

GetDfpsOk returns a tuple with the Dfps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfps

`func (o *SecurityPolicy) SetDfps(v []int32)`

SetDfps sets Dfps field to given value.

### HasDfps

`func (o *SecurityPolicy) HasDfps() bool`

HasDfps returns a boolean if a field has been set.

### GetEcs

`func (o *SecurityPolicy) GetEcs() bool`

GetEcs returns the Ecs field if non-nil, zero value otherwise.

### GetEcsOk

`func (o *SecurityPolicy) GetEcsOk() (*bool, bool)`

GetEcsOk returns a tuple with the Ecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcs

`func (o *SecurityPolicy) SetEcs(v bool)`

SetEcs sets Ecs field to given value.

### HasEcs

`func (o *SecurityPolicy) HasEcs() bool`

HasEcs returns a boolean if a field has been set.

### GetId

`func (o *SecurityPolicy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityPolicy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityPolicy) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *SecurityPolicy) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *SecurityPolicy) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *SecurityPolicy) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *SecurityPolicy) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *SecurityPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetAddressDfps

`func (o *SecurityPolicy) GetNetAddressDfps() []NetAddrDfpAssignment`

GetNetAddressDfps returns the NetAddressDfps field if non-nil, zero value otherwise.

### GetNetAddressDfpsOk

`func (o *SecurityPolicy) GetNetAddressDfpsOk() (*[]NetAddrDfpAssignment, bool)`

GetNetAddressDfpsOk returns a tuple with the NetAddressDfps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAddressDfps

`func (o *SecurityPolicy) SetNetAddressDfps(v []NetAddrDfpAssignment)`

SetNetAddressDfps sets NetAddressDfps field to given value.

### HasNetAddressDfps

`func (o *SecurityPolicy) HasNetAddressDfps() bool`

HasNetAddressDfps returns a boolean if a field has been set.

### GetNetworkLists

`func (o *SecurityPolicy) GetNetworkLists() []int64`

GetNetworkLists returns the NetworkLists field if non-nil, zero value otherwise.

### GetNetworkListsOk

`func (o *SecurityPolicy) GetNetworkListsOk() (*[]int64, bool)`

GetNetworkListsOk returns a tuple with the NetworkLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLists

`func (o *SecurityPolicy) SetNetworkLists(v []int64)`

SetNetworkLists sets NetworkLists field to given value.

### HasNetworkLists

`func (o *SecurityPolicy) HasNetworkLists() bool`

HasNetworkLists returns a boolean if a field has been set.

### GetOnpremResolve

`func (o *SecurityPolicy) GetOnpremResolve() bool`

GetOnpremResolve returns the OnpremResolve field if non-nil, zero value otherwise.

### GetOnpremResolveOk

`func (o *SecurityPolicy) GetOnpremResolveOk() (*bool, bool)`

GetOnpremResolveOk returns a tuple with the OnpremResolve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnpremResolve

`func (o *SecurityPolicy) SetOnpremResolve(v bool)`

SetOnpremResolve sets OnpremResolve field to given value.

### HasOnpremResolve

`func (o *SecurityPolicy) HasOnpremResolve() bool`

HasOnpremResolve returns a boolean if a field has been set.

### GetPrecedence

`func (o *SecurityPolicy) GetPrecedence() int32`

GetPrecedence returns the Precedence field if non-nil, zero value otherwise.

### GetPrecedenceOk

`func (o *SecurityPolicy) GetPrecedenceOk() (*int32, bool)`

GetPrecedenceOk returns a tuple with the Precedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedence

`func (o *SecurityPolicy) SetPrecedence(v int32)`

SetPrecedence sets Precedence field to given value.

### HasPrecedence

`func (o *SecurityPolicy) HasPrecedence() bool`

HasPrecedence returns a boolean if a field has been set.

### GetRoamingDeviceGroups

`func (o *SecurityPolicy) GetRoamingDeviceGroups() []int32`

GetRoamingDeviceGroups returns the RoamingDeviceGroups field if non-nil, zero value otherwise.

### GetRoamingDeviceGroupsOk

`func (o *SecurityPolicy) GetRoamingDeviceGroupsOk() (*[]int32, bool)`

GetRoamingDeviceGroupsOk returns a tuple with the RoamingDeviceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamingDeviceGroups

`func (o *SecurityPolicy) SetRoamingDeviceGroups(v []int32)`

SetRoamingDeviceGroups sets RoamingDeviceGroups field to given value.

### HasRoamingDeviceGroups

`func (o *SecurityPolicy) HasRoamingDeviceGroups() bool`

HasRoamingDeviceGroups returns a boolean if a field has been set.

### GetRules

`func (o *SecurityPolicy) GetRules() []SecurityPolicyRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SecurityPolicy) GetRulesOk() (*[]SecurityPolicyRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SecurityPolicy) SetRules(v []SecurityPolicyRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *SecurityPolicy) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetSafeSearch

`func (o *SecurityPolicy) GetSafeSearch() bool`

GetSafeSearch returns the SafeSearch field if non-nil, zero value otherwise.

### GetSafeSearchOk

`func (o *SecurityPolicy) GetSafeSearchOk() (*bool, bool)`

GetSafeSearchOk returns a tuple with the SafeSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeSearch

`func (o *SecurityPolicy) SetSafeSearch(v bool)`

SetSafeSearch sets SafeSearch field to given value.

### HasSafeSearch

`func (o *SecurityPolicy) HasSafeSearch() bool`

HasSafeSearch returns a boolean if a field has been set.

### GetTags

`func (o *SecurityPolicy) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SecurityPolicy) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SecurityPolicy) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *SecurityPolicy) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *SecurityPolicy) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *SecurityPolicy) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *SecurityPolicy) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *SecurityPolicy) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetUserGroups

`func (o *SecurityPolicy) GetUserGroups() []string`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *SecurityPolicy) GetUserGroupsOk() (*[]string, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *SecurityPolicy) SetUserGroups(v []string)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *SecurityPolicy) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


