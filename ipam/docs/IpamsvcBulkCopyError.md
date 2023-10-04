# IpamsvcBulkCopyError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the resource that was requested to be copied. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Message** | Pointer to **string** | The reason why the copy failed. | [optional] 

## Methods

### NewIpamsvcBulkCopyError

`func NewIpamsvcBulkCopyError() *IpamsvcBulkCopyError`

NewIpamsvcBulkCopyError instantiates a new IpamsvcBulkCopyError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcBulkCopyErrorWithDefaults

`func NewIpamsvcBulkCopyErrorWithDefaults() *IpamsvcBulkCopyError`

NewIpamsvcBulkCopyErrorWithDefaults instantiates a new IpamsvcBulkCopyError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IpamsvcBulkCopyError) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IpamsvcBulkCopyError) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IpamsvcBulkCopyError) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IpamsvcBulkCopyError) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *IpamsvcBulkCopyError) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpamsvcBulkCopyError) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpamsvcBulkCopyError) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpamsvcBulkCopyError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *IpamsvcBulkCopyError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IpamsvcBulkCopyError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IpamsvcBulkCopyError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IpamsvcBulkCopyError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


