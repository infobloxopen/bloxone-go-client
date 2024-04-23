# NamedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfidenceLevel** | Pointer to **string** | The confidence level for a custom list. The possible values are [\&quot;LOW\&quot;, \&quot;MEDIUM\&quot;, \&quot;HIGH\&quot;] | [optional] 
**CreatedTime** | Pointer to **time.Time** | The time when this Named List object was created. | [optional] [readonly] 
**Description** | Pointer to **string** | The brief description for the named list. | [optional] 
**Id** | Pointer to **int32** | The Named List object identifier. | [optional] [readonly] 
**ItemCount** | Pointer to **int32** | The number of items in this named list. | [optional] [readonly] 
**Items** | Pointer to **[]string** | The list of the FQDN or IPv4/IPv6 CIDRs to define whitelists and blacklists for additional protection. | [optional] 
**ItemsDescribed** | Pointer to [**[]ItemStructs**](ItemStructs.md) | The List of ItemStructs structure which contains the item and its description | [optional] 
**Name** | Pointer to **string** | The name of the named list. | [optional] 
**Policies** | Pointer to **[]string** | The list of the security policy names with which the named list is associated. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Enables tag support for resource where tags attribute contains user-defined key value pairs | [optional] 
**ThreatLevel** | Pointer to **string** | The threat level for a custom list. The possible values are [\&quot;INFO\&quot;, \&quot;LOW\&quot;, \&quot;MEDIUM\&quot;, \&quot;HIGH\&quot;] | [optional] 
**Type** | Pointer to **string** | The type of the named list, that can be \&quot;custom_list\&quot;, \&quot;threat_insight\&quot;, \&quot;fast_flux\&quot;, \&quot;dga\&quot;, \&quot;dnsm\&quot;, \&quot;threat_insight_nde\&quot;, \&quot;default_allow\&quot;, \&quot;default_block\&quot; or \&quot;threat_insight_nde\&quot;. | [optional] 
**UpdatedTime** | Pointer to **time.Time** | The time when this Named List object was last updated. | [optional] [readonly] 

## Methods

### NewNamedList

`func NewNamedList() *NamedList`

NewNamedList instantiates a new NamedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamedListWithDefaults

`func NewNamedListWithDefaults() *NamedList`

NewNamedListWithDefaults instantiates a new NamedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfidenceLevel

`func (o *NamedList) GetConfidenceLevel() string`

GetConfidenceLevel returns the ConfidenceLevel field if non-nil, zero value otherwise.

### GetConfidenceLevelOk

`func (o *NamedList) GetConfidenceLevelOk() (*string, bool)`

GetConfidenceLevelOk returns a tuple with the ConfidenceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceLevel

`func (o *NamedList) SetConfidenceLevel(v string)`

SetConfidenceLevel sets ConfidenceLevel field to given value.

### HasConfidenceLevel

`func (o *NamedList) HasConfidenceLevel() bool`

HasConfidenceLevel returns a boolean if a field has been set.

### GetCreatedTime

`func (o *NamedList) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *NamedList) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *NamedList) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *NamedList) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *NamedList) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NamedList) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NamedList) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NamedList) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *NamedList) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NamedList) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NamedList) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NamedList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItemCount

`func (o *NamedList) GetItemCount() int32`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *NamedList) GetItemCountOk() (*int32, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *NamedList) SetItemCount(v int32)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *NamedList) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.

### GetItems

`func (o *NamedList) GetItems() []string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NamedList) GetItemsOk() (*[]string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NamedList) SetItems(v []string)`

SetItems sets Items field to given value.

### HasItems

`func (o *NamedList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetItemsDescribed

`func (o *NamedList) GetItemsDescribed() []ItemStructs`

GetItemsDescribed returns the ItemsDescribed field if non-nil, zero value otherwise.

### GetItemsDescribedOk

`func (o *NamedList) GetItemsDescribedOk() (*[]ItemStructs, bool)`

GetItemsDescribedOk returns a tuple with the ItemsDescribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsDescribed

`func (o *NamedList) SetItemsDescribed(v []ItemStructs)`

SetItemsDescribed sets ItemsDescribed field to given value.

### HasItemsDescribed

`func (o *NamedList) HasItemsDescribed() bool`

HasItemsDescribed returns a boolean if a field has been set.

### GetName

`func (o *NamedList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NamedList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NamedList) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NamedList) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicies

`func (o *NamedList) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *NamedList) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *NamedList) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *NamedList) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetTags

`func (o *NamedList) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NamedList) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NamedList) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *NamedList) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThreatLevel

`func (o *NamedList) GetThreatLevel() string`

GetThreatLevel returns the ThreatLevel field if non-nil, zero value otherwise.

### GetThreatLevelOk

`func (o *NamedList) GetThreatLevelOk() (*string, bool)`

GetThreatLevelOk returns a tuple with the ThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLevel

`func (o *NamedList) SetThreatLevel(v string)`

SetThreatLevel sets ThreatLevel field to given value.

### HasThreatLevel

`func (o *NamedList) HasThreatLevel() bool`

HasThreatLevel returns a boolean if a field has been set.

### GetType

`func (o *NamedList) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NamedList) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NamedList) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NamedList) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *NamedList) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *NamedList) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *NamedList) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *NamedList) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


