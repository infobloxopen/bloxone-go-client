# ExclusionRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | The description for the exclusion range. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**End** | **string** | The end address of the exclusion range. | 
**Start** | **string** | The start address of the exclusion range. | 

## Methods

### NewExclusionRange

`func NewExclusionRange(end string, start string, ) *ExclusionRange`

NewExclusionRange instantiates a new ExclusionRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExclusionRangeWithDefaults

`func NewExclusionRangeWithDefaults() *ExclusionRange`

NewExclusionRangeWithDefaults instantiates a new ExclusionRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ExclusionRange) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ExclusionRange) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ExclusionRange) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ExclusionRange) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEnd

`func (o *ExclusionRange) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ExclusionRange) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ExclusionRange) SetEnd(v string)`

SetEnd sets End field to given value.


### GetStart

`func (o *ExclusionRange) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ExclusionRange) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ExclusionRange) SetStart(v string)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


