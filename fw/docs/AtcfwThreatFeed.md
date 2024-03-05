# AtcfwThreatFeed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfidenceLevel** | Pointer to **string** | The Confidence Level of the Feed. | [optional] [readonly] 
**Description** | Pointer to **string** | The brief description for the thread feed. | [optional] [readonly] 
**Key** | Pointer to **string** | The TSIG key of the threat feed. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the thread feed. | [optional] [readonly] 
**Source** | Pointer to [**ThreatFeedSource**](ThreatFeedSource.md) |  | [optional] [default to THREATFEEDSOURCE_INFOBLOX]
**ThreatLevel** | Pointer to **string** | The Threat Level of the Feed. | [optional] [readonly] 

## Methods

### NewAtcfwThreatFeed

`func NewAtcfwThreatFeed() *AtcfwThreatFeed`

NewAtcfwThreatFeed instantiates a new AtcfwThreatFeed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtcfwThreatFeedWithDefaults

`func NewAtcfwThreatFeedWithDefaults() *AtcfwThreatFeed`

NewAtcfwThreatFeedWithDefaults instantiates a new AtcfwThreatFeed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfidenceLevel

`func (o *AtcfwThreatFeed) GetConfidenceLevel() string`

GetConfidenceLevel returns the ConfidenceLevel field if non-nil, zero value otherwise.

### GetConfidenceLevelOk

`func (o *AtcfwThreatFeed) GetConfidenceLevelOk() (*string, bool)`

GetConfidenceLevelOk returns a tuple with the ConfidenceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceLevel

`func (o *AtcfwThreatFeed) SetConfidenceLevel(v string)`

SetConfidenceLevel sets ConfidenceLevel field to given value.

### HasConfidenceLevel

`func (o *AtcfwThreatFeed) HasConfidenceLevel() bool`

HasConfidenceLevel returns a boolean if a field has been set.

### GetDescription

`func (o *AtcfwThreatFeed) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AtcfwThreatFeed) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AtcfwThreatFeed) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AtcfwThreatFeed) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *AtcfwThreatFeed) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AtcfwThreatFeed) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AtcfwThreatFeed) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AtcfwThreatFeed) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *AtcfwThreatFeed) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AtcfwThreatFeed) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AtcfwThreatFeed) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AtcfwThreatFeed) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSource

`func (o *AtcfwThreatFeed) GetSource() ThreatFeedSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AtcfwThreatFeed) GetSourceOk() (*ThreatFeedSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AtcfwThreatFeed) SetSource(v ThreatFeedSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *AtcfwThreatFeed) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetThreatLevel

`func (o *AtcfwThreatFeed) GetThreatLevel() string`

GetThreatLevel returns the ThreatLevel field if non-nil, zero value otherwise.

### GetThreatLevelOk

`func (o *AtcfwThreatFeed) GetThreatLevelOk() (*string, bool)`

GetThreatLevelOk returns a tuple with the ThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLevel

`func (o *AtcfwThreatFeed) SetThreatLevel(v string)`

SetThreatLevel sets ThreatLevel field to given value.

### HasThreatLevel

`func (o *AtcfwThreatFeed) HasThreatLevel() bool`

HasThreatLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


