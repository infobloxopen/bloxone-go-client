# ConfigBulkCopyError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the resource that was requested to be copied. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Message** | Pointer to **string** | The reason why the copy failed. | [optional] 

## Methods

### NewConfigBulkCopyError

`func NewConfigBulkCopyError() *ConfigBulkCopyError`

NewConfigBulkCopyError instantiates a new ConfigBulkCopyError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigBulkCopyErrorWithDefaults

`func NewConfigBulkCopyErrorWithDefaults() *ConfigBulkCopyError`

NewConfigBulkCopyErrorWithDefaults instantiates a new ConfigBulkCopyError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ConfigBulkCopyError) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigBulkCopyError) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigBulkCopyError) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigBulkCopyError) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ConfigBulkCopyError) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigBulkCopyError) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigBulkCopyError) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigBulkCopyError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *ConfigBulkCopyError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConfigBulkCopyError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConfigBulkCopyError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConfigBulkCopyError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


