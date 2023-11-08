# ConfigWarning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Warning message. | [optional] 
**Name** | Pointer to **string** | Name of a warning. | [optional] 

## Methods

### NewConfigWarning

`func NewConfigWarning() *ConfigWarning`

NewConfigWarning instantiates a new ConfigWarning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigWarningWithDefaults

`func NewConfigWarningWithDefaults() *ConfigWarning`

NewConfigWarningWithDefaults instantiates a new ConfigWarning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ConfigWarning) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConfigWarning) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConfigWarning) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConfigWarning) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *ConfigWarning) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigWarning) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigWarning) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigWarning) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


