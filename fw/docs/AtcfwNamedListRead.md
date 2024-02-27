# AtcfwNamedListRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfidenceLevel** | Pointer to **string** | The confidence level for a custom list. The possible values are [\&quot;LOW\&quot;, \&quot;MEDIUM\&quot;, \&quot;HIGH\&quot;] | [optional] [readonly] 
**CreatedTime** | Pointer to **time.Time** | The time when this Named List object was created. | [optional] [readonly] 
**Description** | Pointer to **string** | The brief description for the named list. | [optional] [readonly] 
**Id** | Pointer to **int32** | The Named List object identifier. | [optional] [readonly] 
**ItemCount** | Pointer to **int32** | The number of items in this named list. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the named list. | [optional] [readonly] 
**Policies** | Pointer to **[]string** | The list of the security policy names with which the named list is associated. | [optional] [readonly] 
**Tags** | Pointer to **map[string]interface{}** | Tags associated with this Named List | [optional] 
**ThreatLevel** | Pointer to **string** | The threat level for a custom list. The possible values are [\&quot;INFO\&quot;, \&quot;LOW\&quot;, \&quot;MEDIUM\&quot;, \&quot;HIGH\&quot;] | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the named list, that can be \&quot;custom_list\&quot;, \&quot;threat_insight\&quot;, \&quot;fast_flux\&quot;, \&quot;dga\&quot;, \&quot;dnsm\&quot;, \&quot;threat_insight_nde\&quot;, \&quot;default_allow\&quot;, \&quot;default_block\&quot;. | [optional] [readonly] 
**UpdatedTime** | Pointer to **time.Time** | The time when this Named List object was last updated. | [optional] [readonly] 

## Methods

### NewAtcfwNamedListRead

`func NewAtcfwNamedListRead() *AtcfwNamedListRead`

NewAtcfwNamedListRead instantiates a new AtcfwNamedListRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtcfwNamedListReadWithDefaults

`func NewAtcfwNamedListReadWithDefaults() *AtcfwNamedListRead`

NewAtcfwNamedListReadWithDefaults instantiates a new AtcfwNamedListRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfidenceLevel

`func (o *AtcfwNamedListRead) GetConfidenceLevel() string`

GetConfidenceLevel returns the ConfidenceLevel field if non-nil, zero value otherwise.

### GetConfidenceLevelOk

`func (o *AtcfwNamedListRead) GetConfidenceLevelOk() (*string, bool)`

GetConfidenceLevelOk returns a tuple with the ConfidenceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceLevel

`func (o *AtcfwNamedListRead) SetConfidenceLevel(v string)`

SetConfidenceLevel sets ConfidenceLevel field to given value.

### HasConfidenceLevel

`func (o *AtcfwNamedListRead) HasConfidenceLevel() bool`

HasConfidenceLevel returns a boolean if a field has been set.

### GetCreatedTime

`func (o *AtcfwNamedListRead) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AtcfwNamedListRead) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AtcfwNamedListRead) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *AtcfwNamedListRead) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *AtcfwNamedListRead) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AtcfwNamedListRead) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AtcfwNamedListRead) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AtcfwNamedListRead) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *AtcfwNamedListRead) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AtcfwNamedListRead) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AtcfwNamedListRead) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AtcfwNamedListRead) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItemCount

`func (o *AtcfwNamedListRead) GetItemCount() int32`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *AtcfwNamedListRead) GetItemCountOk() (*int32, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *AtcfwNamedListRead) SetItemCount(v int32)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *AtcfwNamedListRead) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.

### GetName

`func (o *AtcfwNamedListRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AtcfwNamedListRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AtcfwNamedListRead) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AtcfwNamedListRead) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicies

`func (o *AtcfwNamedListRead) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *AtcfwNamedListRead) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *AtcfwNamedListRead) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *AtcfwNamedListRead) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetTags

`func (o *AtcfwNamedListRead) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AtcfwNamedListRead) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AtcfwNamedListRead) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *AtcfwNamedListRead) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThreatLevel

`func (o *AtcfwNamedListRead) GetThreatLevel() string`

GetThreatLevel returns the ThreatLevel field if non-nil, zero value otherwise.

### GetThreatLevelOk

`func (o *AtcfwNamedListRead) GetThreatLevelOk() (*string, bool)`

GetThreatLevelOk returns a tuple with the ThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLevel

`func (o *AtcfwNamedListRead) SetThreatLevel(v string)`

SetThreatLevel sets ThreatLevel field to given value.

### HasThreatLevel

`func (o *AtcfwNamedListRead) HasThreatLevel() bool`

HasThreatLevel returns a boolean if a field has been set.

### GetType

`func (o *AtcfwNamedListRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AtcfwNamedListRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AtcfwNamedListRead) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AtcfwNamedListRead) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *AtcfwNamedListRead) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *AtcfwNamedListRead) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *AtcfwNamedListRead) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *AtcfwNamedListRead) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


