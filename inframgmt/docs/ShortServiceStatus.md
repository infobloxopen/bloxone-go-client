# ShortServiceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Status message text for the Service. | [optional] 
**Status** | Pointer to **string** | Status of the Service (&#x60;started&#x60;, &#x60;starting&#x60;, &#x60;stopping&#x60;, &#x60;stopped&#x60;, &#x60;error&#x60;). | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp of the latest status update of Service. | [optional] 

## Methods

### NewShortServiceStatus

`func NewShortServiceStatus() *ShortServiceStatus`

NewShortServiceStatus instantiates a new ShortServiceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShortServiceStatusWithDefaults

`func NewShortServiceStatusWithDefaults() *ShortServiceStatus`

NewShortServiceStatusWithDefaults instantiates a new ShortServiceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ShortServiceStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ShortServiceStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ShortServiceStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ShortServiceStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ShortServiceStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ShortServiceStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ShortServiceStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ShortServiceStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ShortServiceStatus) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ShortServiceStatus) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ShortServiceStatus) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ShortServiceStatus) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


