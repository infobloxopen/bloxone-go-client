# MacAddressItemUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Append** | Pointer to **bool** | If data needs to be appended or overwritten. Defaults to _true_. | [optional] 
**Content** | **string** | The content in plain text of the mac addresses to be uploaded to a large scale hardware filter. | 
**HardwareFilterId** | **string** | The resource identifier. | 

## Methods

### NewMacAddressItemUpload

`func NewMacAddressItemUpload(content string, hardwareFilterId string, ) *MacAddressItemUpload`

NewMacAddressItemUpload instantiates a new MacAddressItemUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacAddressItemUploadWithDefaults

`func NewMacAddressItemUploadWithDefaults() *MacAddressItemUpload`

NewMacAddressItemUploadWithDefaults instantiates a new MacAddressItemUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppend

`func (o *MacAddressItemUpload) GetAppend() bool`

GetAppend returns the Append field if non-nil, zero value otherwise.

### GetAppendOk

`func (o *MacAddressItemUpload) GetAppendOk() (*bool, bool)`

GetAppendOk returns a tuple with the Append field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppend

`func (o *MacAddressItemUpload) SetAppend(v bool)`

SetAppend sets Append field to given value.

### HasAppend

`func (o *MacAddressItemUpload) HasAppend() bool`

HasAppend returns a boolean if a field has been set.

### GetContent

`func (o *MacAddressItemUpload) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MacAddressItemUpload) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MacAddressItemUpload) SetContent(v string)`

SetContent sets Content field to given value.


### GetHardwareFilterId

`func (o *MacAddressItemUpload) GetHardwareFilterId() string`

GetHardwareFilterId returns the HardwareFilterId field if non-nil, zero value otherwise.

### GetHardwareFilterIdOk

`func (o *MacAddressItemUpload) GetHardwareFilterIdOk() (*string, bool)`

GetHardwareFilterIdOk returns a tuple with the HardwareFilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareFilterId

`func (o *MacAddressItemUpload) SetHardwareFilterId(v string)`

SetHardwareFilterId sets HardwareFilterId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


