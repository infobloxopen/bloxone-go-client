# ListSeverityLevels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfidenceLevel** | Pointer to **string** | The confidence level for a TI list. The possible values are [LOW\&quot;, \&quot;MEDIUM\&quot;, \&quot;HIGH\&quot;] | [optional] 
**Id** | Pointer to **int32** | The Named List object identifier. | [optional] [readonly] 
**Tags** | Pointer to **map[string]interface{}** | Enables tag support for resource where tags attribute contains user-defined key value pairs | [optional] 
**ThreatLevel** | Pointer to **string** | The threat level for a TI list. The possible values are [\&quot;INFO\&quot;, \&quot;LOW\&quot;, \&quot;MEDIUM\&quot;, \&quot;HIGH\&quot;] | [optional] 

## Methods

### NewListSeverityLevels

`func NewListSeverityLevels() *ListSeverityLevels`

NewListSeverityLevels instantiates a new ListSeverityLevels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSeverityLevelsWithDefaults

`func NewListSeverityLevelsWithDefaults() *ListSeverityLevels`

NewListSeverityLevelsWithDefaults instantiates a new ListSeverityLevels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfidenceLevel

`func (o *ListSeverityLevels) GetConfidenceLevel() string`

GetConfidenceLevel returns the ConfidenceLevel field if non-nil, zero value otherwise.

### GetConfidenceLevelOk

`func (o *ListSeverityLevels) GetConfidenceLevelOk() (*string, bool)`

GetConfidenceLevelOk returns a tuple with the ConfidenceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceLevel

`func (o *ListSeverityLevels) SetConfidenceLevel(v string)`

SetConfidenceLevel sets ConfidenceLevel field to given value.

### HasConfidenceLevel

`func (o *ListSeverityLevels) HasConfidenceLevel() bool`

HasConfidenceLevel returns a boolean if a field has been set.

### GetId

`func (o *ListSeverityLevels) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListSeverityLevels) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListSeverityLevels) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ListSeverityLevels) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTags

`func (o *ListSeverityLevels) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListSeverityLevels) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListSeverityLevels) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListSeverityLevels) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThreatLevel

`func (o *ListSeverityLevels) GetThreatLevel() string`

GetThreatLevel returns the ThreatLevel field if non-nil, zero value otherwise.

### GetThreatLevelOk

`func (o *ListSeverityLevels) GetThreatLevelOk() (*string, bool)`

GetThreatLevelOk returns a tuple with the ThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLevel

`func (o *ListSeverityLevels) SetThreatLevel(v string)`

SetThreatLevel sets ThreatLevel field to given value.

### HasThreatLevel

`func (o *ListSeverityLevels) HasThreatLevel() bool`

HasThreatLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


