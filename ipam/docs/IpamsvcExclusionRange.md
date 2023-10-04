# IpamsvcExclusionRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | The description for the exclusion range. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**End** | **string** | The end address of the exclusion range. | 
**Start** | **string** | The start address of the exclusion range. | 

## Methods

### NewIpamsvcExclusionRange

`func NewIpamsvcExclusionRange(end string, start string, ) *IpamsvcExclusionRange`

NewIpamsvcExclusionRange instantiates a new IpamsvcExclusionRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcExclusionRangeWithDefaults

`func NewIpamsvcExclusionRangeWithDefaults() *IpamsvcExclusionRange`

NewIpamsvcExclusionRangeWithDefaults instantiates a new IpamsvcExclusionRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *IpamsvcExclusionRange) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IpamsvcExclusionRange) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IpamsvcExclusionRange) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IpamsvcExclusionRange) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEnd

`func (o *IpamsvcExclusionRange) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *IpamsvcExclusionRange) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *IpamsvcExclusionRange) SetEnd(v string)`

SetEnd sets End field to given value.


### GetStart

`func (o *IpamsvcExclusionRange) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *IpamsvcExclusionRange) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *IpamsvcExclusionRange) SetStart(v string)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


