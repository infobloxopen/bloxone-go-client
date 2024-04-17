# AtcfwSecurityPolicy

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
**NetAddressDfps** | Pointer to [**[]AtcfwNetAddrDfpAssignment**](AtcfwNetAddrDfpAssignment.md) | List of DFPs associated with this policy via network address (with corresponding network address) | [optional] 
**NetworkLists** | Pointer to **[]int64** | The list of Network Lists identifiers that represents networks that you want to protect from malicious attacks. | [optional] 
**OnpremResolve** | Pointer to **bool** | Use DNS resolve on onprem side | [optional] 
**Precedence** | Pointer to **int32** | Security precedence enable selection of the highest priority policy, in cases where a query matches multiple policies. | [optional] 
**RoamingDeviceGroups** | Pointer to **[]int32** | The list of BloxOne Endpoint groups identifiers. | [optional] 
**Rules** | Pointer to [**[]AtcfwSecurityPolicyRule**](AtcfwSecurityPolicyRule.md) | The list of Security Policy Rules objects that represent the set of rules and actions that you define to balance access and constraints so you can mitigate malicious attacks and provide security for your networks. | [optional] 
**SafeSearch** | Pointer to **bool** | Apply automated rules to enforce safe search | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Enables tag support for resource where tags attribute contains user-defined key value pairs | [optional] 
**UpdatedTime** | Pointer to **time.Time** | The time when this Security Policy object was last updated. | [optional] [readonly] 
**UserGroups** | Pointer to **[]string** | List of user groups associated with this policy | [optional] 

## Methods

### NewAtcfwSecurityPolicy

`func NewAtcfwSecurityPolicy() *AtcfwSecurityPolicy`

NewAtcfwSecurityPolicy instantiates a new AtcfwSecurityPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtcfwSecurityPolicyWithDefaults

`func NewAtcfwSecurityPolicyWithDefaults() *AtcfwSecurityPolicy`

NewAtcfwSecurityPolicyWithDefaults instantiates a new AtcfwSecurityPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessCodes

`func (o *AtcfwSecurityPolicy) GetAccessCodes() []string`

GetAccessCodes returns the AccessCodes field if non-nil, zero value otherwise.

### GetAccessCodesOk

`func (o *AtcfwSecurityPolicy) GetAccessCodesOk() (*[]string, bool)`

GetAccessCodesOk returns a tuple with the AccessCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCodes

`func (o *AtcfwSecurityPolicy) SetAccessCodes(v []string)`

SetAccessCodes sets AccessCodes field to given value.

### HasAccessCodes

`func (o *AtcfwSecurityPolicy) HasAccessCodes() bool`

HasAccessCodes returns a boolean if a field has been set.

### GetCreatedTime

`func (o *AtcfwSecurityPolicy) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AtcfwSecurityPolicy) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AtcfwSecurityPolicy) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *AtcfwSecurityPolicy) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDefaultAction

`func (o *AtcfwSecurityPolicy) GetDefaultAction() string`

GetDefaultAction returns the DefaultAction field if non-nil, zero value otherwise.

### GetDefaultActionOk

`func (o *AtcfwSecurityPolicy) GetDefaultActionOk() (*string, bool)`

GetDefaultActionOk returns a tuple with the DefaultAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAction

`func (o *AtcfwSecurityPolicy) SetDefaultAction(v string)`

SetDefaultAction sets DefaultAction field to given value.

### HasDefaultAction

`func (o *AtcfwSecurityPolicy) HasDefaultAction() bool`

HasDefaultAction returns a boolean if a field has been set.

### GetDefaultRedirectName

`func (o *AtcfwSecurityPolicy) GetDefaultRedirectName() string`

GetDefaultRedirectName returns the DefaultRedirectName field if non-nil, zero value otherwise.

### GetDefaultRedirectNameOk

`func (o *AtcfwSecurityPolicy) GetDefaultRedirectNameOk() (*string, bool)`

GetDefaultRedirectNameOk returns a tuple with the DefaultRedirectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRedirectName

`func (o *AtcfwSecurityPolicy) SetDefaultRedirectName(v string)`

SetDefaultRedirectName sets DefaultRedirectName field to given value.

### HasDefaultRedirectName

`func (o *AtcfwSecurityPolicy) HasDefaultRedirectName() bool`

HasDefaultRedirectName returns a boolean if a field has been set.

### GetDescription

`func (o *AtcfwSecurityPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AtcfwSecurityPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AtcfwSecurityPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AtcfwSecurityPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDfpServices

`func (o *AtcfwSecurityPolicy) GetDfpServices() []string`

GetDfpServices returns the DfpServices field if non-nil, zero value otherwise.

### GetDfpServicesOk

`func (o *AtcfwSecurityPolicy) GetDfpServicesOk() (*[]string, bool)`

GetDfpServicesOk returns a tuple with the DfpServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfpServices

`func (o *AtcfwSecurityPolicy) SetDfpServices(v []string)`

SetDfpServices sets DfpServices field to given value.

### HasDfpServices

`func (o *AtcfwSecurityPolicy) HasDfpServices() bool`

HasDfpServices returns a boolean if a field has been set.

### GetDfps

`func (o *AtcfwSecurityPolicy) GetDfps() []int32`

GetDfps returns the Dfps field if non-nil, zero value otherwise.

### GetDfpsOk

`func (o *AtcfwSecurityPolicy) GetDfpsOk() (*[]int32, bool)`

GetDfpsOk returns a tuple with the Dfps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfps

`func (o *AtcfwSecurityPolicy) SetDfps(v []int32)`

SetDfps sets Dfps field to given value.

### HasDfps

`func (o *AtcfwSecurityPolicy) HasDfps() bool`

HasDfps returns a boolean if a field has been set.

### GetEcs

`func (o *AtcfwSecurityPolicy) GetEcs() bool`

GetEcs returns the Ecs field if non-nil, zero value otherwise.

### GetEcsOk

`func (o *AtcfwSecurityPolicy) GetEcsOk() (*bool, bool)`

GetEcsOk returns a tuple with the Ecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcs

`func (o *AtcfwSecurityPolicy) SetEcs(v bool)`

SetEcs sets Ecs field to given value.

### HasEcs

`func (o *AtcfwSecurityPolicy) HasEcs() bool`

HasEcs returns a boolean if a field has been set.

### GetId

`func (o *AtcfwSecurityPolicy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AtcfwSecurityPolicy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AtcfwSecurityPolicy) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AtcfwSecurityPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *AtcfwSecurityPolicy) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *AtcfwSecurityPolicy) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *AtcfwSecurityPolicy) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *AtcfwSecurityPolicy) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *AtcfwSecurityPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AtcfwSecurityPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AtcfwSecurityPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AtcfwSecurityPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetAddressDfps

`func (o *AtcfwSecurityPolicy) GetNetAddressDfps() []AtcfwNetAddrDfpAssignment`

GetNetAddressDfps returns the NetAddressDfps field if non-nil, zero value otherwise.

### GetNetAddressDfpsOk

`func (o *AtcfwSecurityPolicy) GetNetAddressDfpsOk() (*[]AtcfwNetAddrDfpAssignment, bool)`

GetNetAddressDfpsOk returns a tuple with the NetAddressDfps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAddressDfps

`func (o *AtcfwSecurityPolicy) SetNetAddressDfps(v []AtcfwNetAddrDfpAssignment)`

SetNetAddressDfps sets NetAddressDfps field to given value.

### HasNetAddressDfps

`func (o *AtcfwSecurityPolicy) HasNetAddressDfps() bool`

HasNetAddressDfps returns a boolean if a field has been set.

### GetNetworkLists

`func (o *AtcfwSecurityPolicy) GetNetworkLists() []int64`

GetNetworkLists returns the NetworkLists field if non-nil, zero value otherwise.

### GetNetworkListsOk

`func (o *AtcfwSecurityPolicy) GetNetworkListsOk() (*[]int64, bool)`

GetNetworkListsOk returns a tuple with the NetworkLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLists

`func (o *AtcfwSecurityPolicy) SetNetworkLists(v []int64)`

SetNetworkLists sets NetworkLists field to given value.

### HasNetworkLists

`func (o *AtcfwSecurityPolicy) HasNetworkLists() bool`

HasNetworkLists returns a boolean if a field has been set.

### GetOnpremResolve

`func (o *AtcfwSecurityPolicy) GetOnpremResolve() bool`

GetOnpremResolve returns the OnpremResolve field if non-nil, zero value otherwise.

### GetOnpremResolveOk

`func (o *AtcfwSecurityPolicy) GetOnpremResolveOk() (*bool, bool)`

GetOnpremResolveOk returns a tuple with the OnpremResolve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnpremResolve

`func (o *AtcfwSecurityPolicy) SetOnpremResolve(v bool)`

SetOnpremResolve sets OnpremResolve field to given value.

### HasOnpremResolve

`func (o *AtcfwSecurityPolicy) HasOnpremResolve() bool`

HasOnpremResolve returns a boolean if a field has been set.

### GetPrecedence

`func (o *AtcfwSecurityPolicy) GetPrecedence() int32`

GetPrecedence returns the Precedence field if non-nil, zero value otherwise.

### GetPrecedenceOk

`func (o *AtcfwSecurityPolicy) GetPrecedenceOk() (*int32, bool)`

GetPrecedenceOk returns a tuple with the Precedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedence

`func (o *AtcfwSecurityPolicy) SetPrecedence(v int32)`

SetPrecedence sets Precedence field to given value.

### HasPrecedence

`func (o *AtcfwSecurityPolicy) HasPrecedence() bool`

HasPrecedence returns a boolean if a field has been set.

### GetRoamingDeviceGroups

`func (o *AtcfwSecurityPolicy) GetRoamingDeviceGroups() []int32`

GetRoamingDeviceGroups returns the RoamingDeviceGroups field if non-nil, zero value otherwise.

### GetRoamingDeviceGroupsOk

`func (o *AtcfwSecurityPolicy) GetRoamingDeviceGroupsOk() (*[]int32, bool)`

GetRoamingDeviceGroupsOk returns a tuple with the RoamingDeviceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamingDeviceGroups

`func (o *AtcfwSecurityPolicy) SetRoamingDeviceGroups(v []int32)`

SetRoamingDeviceGroups sets RoamingDeviceGroups field to given value.

### HasRoamingDeviceGroups

`func (o *AtcfwSecurityPolicy) HasRoamingDeviceGroups() bool`

HasRoamingDeviceGroups returns a boolean if a field has been set.

### GetRules

`func (o *AtcfwSecurityPolicy) GetRules() []AtcfwSecurityPolicyRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *AtcfwSecurityPolicy) GetRulesOk() (*[]AtcfwSecurityPolicyRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *AtcfwSecurityPolicy) SetRules(v []AtcfwSecurityPolicyRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *AtcfwSecurityPolicy) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetSafeSearch

`func (o *AtcfwSecurityPolicy) GetSafeSearch() bool`

GetSafeSearch returns the SafeSearch field if non-nil, zero value otherwise.

### GetSafeSearchOk

`func (o *AtcfwSecurityPolicy) GetSafeSearchOk() (*bool, bool)`

GetSafeSearchOk returns a tuple with the SafeSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeSearch

`func (o *AtcfwSecurityPolicy) SetSafeSearch(v bool)`

SetSafeSearch sets SafeSearch field to given value.

### HasSafeSearch

`func (o *AtcfwSecurityPolicy) HasSafeSearch() bool`

HasSafeSearch returns a boolean if a field has been set.

### GetTags

`func (o *AtcfwSecurityPolicy) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AtcfwSecurityPolicy) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AtcfwSecurityPolicy) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *AtcfwSecurityPolicy) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *AtcfwSecurityPolicy) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *AtcfwSecurityPolicy) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *AtcfwSecurityPolicy) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *AtcfwSecurityPolicy) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetUserGroups

`func (o *AtcfwSecurityPolicy) GetUserGroups() []string`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *AtcfwSecurityPolicy) GetUserGroupsOk() (*[]string, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *AtcfwSecurityPolicy) SetUserGroups(v []string)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *AtcfwSecurityPolicy) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


