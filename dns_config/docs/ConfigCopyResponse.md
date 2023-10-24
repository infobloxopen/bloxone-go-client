# ConfigCopyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the resource that was requested to be copied. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**JobId** | Pointer to **string** | An Unique Id to identify copy operation. | [optional] 

## Methods

### NewConfigCopyResponse

`func NewConfigCopyResponse() *ConfigCopyResponse`

NewConfigCopyResponse instantiates a new ConfigCopyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigCopyResponseWithDefaults

`func NewConfigCopyResponseWithDefaults() *ConfigCopyResponse`

NewConfigCopyResponseWithDefaults instantiates a new ConfigCopyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ConfigCopyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigCopyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigCopyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigCopyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ConfigCopyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigCopyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigCopyResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigCopyResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobId

`func (o *ConfigCopyResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ConfigCopyResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ConfigCopyResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *ConfigCopyResponse) HasJobId() bool`

HasJobId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


